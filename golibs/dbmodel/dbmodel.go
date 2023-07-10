package dbmodel

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"errors"
	"flag"
	"fmt"
	"reflect"
	"strings"
	"time"

	"cqserver/golibs/logger"

	"math/rand"

	"github.com/astaxie/beego/orm"
	"github.com/bradfitz/gomemcache/memcache"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

var log = logger.Get("default", true)

var (
	cache           *memcache.Client
	cacheExpiration int32
	enableDbTrace   bool

	errNoCacheServer = errors.New("no cache server found")

	modelMap  = make(map[string][]modelMapItem)
	callbacks = make([]func(), 0)
	inited    = false
)

type DbConnectionStringGetter interface {
	GetDbConnectionString(dbKey string) (string, int, int)
}

type modelMapItem struct {
	model  Model
	initer func(dbMap *gorp.DbMap)
}

type Model interface {
	SetDbMap(dbMap *gorp.DbMap)
	DbMap() *gorp.DbMap
	SetDb(db *sql.DB)
	Db() *sql.DB
}

type CommonModel struct {
	dbMap *gorp.DbMap
	db    *sql.DB
}

func init() {
	flag.BoolVar(&enableDbTrace, "dbtrace", false, "whether enable gorp db trace")
}

func (this *CommonModel) SetDbMap(dbMap *gorp.DbMap) {
	this.dbMap = dbMap
}

func (this *CommonModel) DbMap() *gorp.DbMap {
	return this.dbMap
}

func (this *CommonModel) SetDb(db *sql.DB) {
	this.db = db
}

func (this *CommonModel) Db() *sql.DB {
	return this.db
}

type DbLogger struct {
}

func (this *DbLogger) Printf(format string, v ...interface{}) {
	log.Info(fmt.Sprintf(format, v...))
}

func RegisterCallback(callback func()) {
	if inited {
		callback()
		return
	}
	callbacks = append(callbacks, callback)
}

func Register(dbKey string, model Model, initer func(dbMap *gorp.DbMap)) {
	mapItems, ok := modelMap[dbKey]
	if ok {
		for _, mi := range mapItems {
			if model == mi.model {
				return
			}
			modelMap[dbKey] = append(mapItems, modelMapItem{model: model, initer: initer})
		}
	} else {
		mapItems := make([]modelMapItem, 0, 5)
		modelMap[dbKey] = append(mapItems, modelMapItem{model: model, initer: initer})
	}
}

func InitDb(connStrGetter DbConnectionStringGetter, tableAutoCheckDbKey []string, ormDefaultDbKey string,columnCheck []string) error {

	for dbKey, _ := range modelMap {
		dbUrl, maxIdle, maxOpenCon := connStrGetter.GetDbConnectionString(dbKey)
		if dbUrl == "" {
			continue
		}
		//检查判断是否需要表自动检查
		tableAutoCheck := false
		if tableAutoCheckDbKey != nil {
			for _, v := range tableAutoCheckDbKey {
				if dbKey == v {
					tableAutoCheck = true
					break
				}
			}
		}

		isCheckColumn := false
		if columnCheck != nil {
			for _,v := range columnCheck {
				if dbKey == v {
					isCheckColumn = true
					break
				}
			}
		}
		err := dbInit(dbUrl, dbKey, tableAutoCheck, dbKey == ormDefaultDbKey, maxIdle, maxOpenCon,isCheckColumn)
		if err != nil {
			logger.Error("dbInit dbkey:%v,dbUrl:%v,err:%v",dbKey,dbUrl,err)
			return err
		}

	}
	// modelMap = nil

	inited = true
	for _, callback := range callbacks {
		callback()
	}
	return nil
}

func InitDbByKey(connStrGetter DbConnectionStringGetter, key string, tableAutoCheck bool, ormDefaultDb bool,isCheckColumn bool) error {
	for dbKey, _ := range modelMap {
		if dbKey != key {
			continue
		}
		dbUrl, maxIdle, maxOpenCon := connStrGetter.GetDbConnectionString(dbKey)
		if dbUrl == "" {
			continue
		}
		err := dbInit(dbUrl, dbKey, tableAutoCheck, ormDefaultDb, maxIdle, maxOpenCon,isCheckColumn)
		if err != nil {
			logger.Error("dbInit dbkey:%v,dbUrl:%v,err:%v",dbKey,dbUrl,err)
			return err
		}
	}
	delete(modelMap, key)
	return nil
}

func dbInit(dbUrl string, dbKey string, tableAutoCheck bool, ormDefaultDb bool, maxIdle int, maxOpenCon int,isCheckColumn bool) error {
	//dbUrl := connStrGetter.GetDbConnectionString(dbKey)
	if dbUrl == "" {
		return fmt.Errorf("init db:%v,err,db has not config", dbKey)
	}
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		logger.Error("sql open dburl:%v, err:%v",dbUrl,err)
		return err
	}
	err = db.Ping()
	if err != nil {
		logger.Error("sql  ping dburl:%v, err:%v",dbUrl,err)
		return err
	}
	if maxIdle <= 0 {
		maxIdle = 3
	}
	if maxOpenCon <= 0 {
		maxOpenCon = 5
	}
	logger.Info("openDB %v %v %v,%v,%v,%v", dbKey, dbUrl, tableAutoCheck, ormDefaultDb, maxIdle, maxOpenCon)
	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpenCon)
	//db.SetConnMaxLifetime(110)
	//开启保活
	//go dbPing(db)
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	if enableDbTrace {
		dbMap.TraceOn("[gorp]", &DbLogger{})
	}
	mapItems, ok := modelMap[dbKey]
	if ok {
		for _, mi := range mapItems {
			mi.model.SetDbMap(dbMap)
			mi.model.SetDb(db)
			mi.initer(dbMap)
		}
	}

	//是否注册为orm默认数据库
	if ormDefaultDb {
		err := orm.AddAliasWthDB("default", "mysql", db)
		if err != nil {
			return err
		}
	}

	//表结构自动检测
	if tableAutoCheck {
		err := orm.AddAliasWthDB(dbKey, "mysql", db)
		if err != nil {
			return err
		}
		err = orm.RunSyncdb(dbKey, false, true)
		if err != nil {
			return err
		}
	}

	if isCheckColumn {
		missColumns, err := dbMap.CheckMissColumns()
		if err != nil {
			return err
		}
		if missColumns != nil {
			for k, v := range missColumns {
				for _, vv := range v {
					logger.Error("数据库：[%v],表：[%v],缺少列：[%v]", dbKey,k, vv)
				}
			}
			return errors.New(fmt.Sprintf("数据库缺少列"))
		}
	}
	return nil
}

func dbPing(db *sql.DB) {

	rand.Seed(time.Now().UnixNano())
	timer := time.NewTicker(time.Duration(100+rand.Intn(10)) * time.Second)
	for {
		select {
		case <-timer.C:
			db.Ping()
		}
	}
}

// get all fields of a model, used in select clause to prevent return undefined fields
func GetAllFieldsAsString(obj interface{}) string {
	objT := reflect.TypeOf(obj)
	var fields []string
	for i := 0; i < objT.NumField(); i++ {
		fieldT := objT.Field(i)
		tag := fieldT.Tag.Get("db")
		if tag == "" {
			continue
		}
		fields = append(fields, "`"+tag+"`")
	}
	return strings.Join(fields, ",")
}

func GetCache(key string, v interface{}) error {
	if cache == nil {
		return errNoCacheServer
	}
	mi, err := cache.Get(key)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(bytes.NewReader(mi.Value))
	err = dec.Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func SetCache(key string, v interface{}, expiration int32) error {
	if cache == nil {
		return errNoCacheServer
	}
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return err
	}
	return cache.Set(&memcache.Item{Key: key, Value: buf.Bytes(), Expiration: expiration})
}

func GetBytesCache(key string) ([]byte, error) {
	if cache == nil {
		return nil, errNoCacheServer
	}
	mi, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	return mi.Value, nil
}

func SetBytesCache(key string, value []byte, expiration int32) error {
	if cache == nil {
		return errNoCacheServer
	}
	return cache.Set(&memcache.Item{Key: key, Value: value, Expiration: expiration})
}

func RemoveCache(key string) error {
	if cache == nil {
		return errNoCacheServer
	}
	return cache.Delete(key)
}

func ResetSetModelMap() {
	modelMap = make(map[string][]modelMapItem)
}
