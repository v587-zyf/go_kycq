package tlog

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/dbmodel"
	"cqserver/golibs/logger"
	"gopkg.in/gorp.v1"
)

const (
	CHAN_MAX_LEN = 100
)

// 数据库操作记录
type DbRecord struct {
	DbMap   *gorp.DbMap // 数据库连接
	Content interface{} // 数据表实体
}

type LogModel struct {
	dbmodel.CommonModel
}

var (
	log          = logger.Get("default", true)
	dbRecordChan chan *DbRecord // 数据库操作通道

	LogModelManager = &LogModel{} // log库管理
)

// 初始化
func Init() {
	dbmodel.Register(model.DB_LOG, LogModelManager, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(LogPlayerRegister{}, "log_player_register")
		dbMap.AddTableWithName(LogPlayerLogin{}, "log_player_login")
		dbMap.AddTableWithName(LogPlayerLogout{}, "log_player_logout")
		dbMap.AddTableWithName(LogItemFlow{}, "log_item_flow")
	})
	// 开启操作线程
	dbRecordChan = make(chan *DbRecord, CHAN_MAX_LEN)
	go dbOperate()
}

func dbOperate() {
	for {
		writeLogToDB()
	}
}

func writeLogToDB() {
	var dbr interface{}
	//回溯错误
	defer func() {
		if err := recover(); err != nil {
			logger.Error("tlog Panic Error. %T, err: %v", dbr, err)
		}
	}()

	select {
	case dbRecord := <-dbRecordChan:
		dbr = dbRecord.Content
		err := dbRecord.DbMap.Insert(dbRecord.Content)
		if err != nil {
			logger.Error("logDbOperate Error. %T error:%v", dbRecord.Content, err)
		}
	}
}

// 记录日志到DB中
func DbLog(dbmap *gorp.DbMap, content interface{}) {
	//fmt.Println("dbRecordChan len", len(dbRecordChan))
	if len(dbRecordChan) < CHAN_MAX_LEN {
		dbRecord := &DbRecord{
			DbMap:   dbmap,
			Content: content,
		}
		dbRecordChan <- dbRecord
	} else {
		logger.Info("dbRecordChan is full %v", content)
	}
}
