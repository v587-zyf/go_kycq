package main

import (
	"cqserver/tools/serverMerge/internal/managers"
	"flag"
	"math/rand"
	_ "net/http/pprof"
	"runtime"
	"time"

	_ "cqserver/gamelibs/retry"
	"cqserver/golibs/logger"
	"cqserver/tools/serverMerge/internal/base"
	"cqserver/tools/serverMerge/internal/model"
)

const (
	APP_NAME    = "serverMerge"
)

var (
	APP_VERSION = "1.0.0"
)


var (
	confFile        = flag.String("conf", "./config.conf", "specify config file")
	mergeServerIds  = flag.String("server_ids", "", "合服的serverid（不填写默认读config的mergeServerIds）")
	dbName          = flag.String("db_name", "", "合并到的新数据库dbname（不填写默认读config的dbconfigs.serverdbNew.dbname）")
	dbLogName       = flag.String("db_log_name", "", "合并到的新数据库dbname（不填写默认读config的dbconfigs.serverdbNew.dbname）")
	dbHost          = flag.String("db_host", "", "合并到的新数据库host（不填写默认读config的dbconfigs.serverdbNew.host）")
	dbUser          = flag.String("db_user", "", "合并到的新数据库db_user（不填写默认读config的dbconfigs.serverdbNew.uid）")
	dbPass          = flag.String("db_pass", "", "合并到的新数据库db_pass（不填写默认读config的dbconfigs.serverdbNew.pwd）")
	dbPort          = flag.Int("db_port", -1, "合并到的新数据库db_port（不填写默认读config的dbconfigs.serverdbNew.port）")
	redisAddr       = flag.String("redis_address", "", "合并到的新redis redis_address（不填写默认读config的dbconfigs.redis.address）")
	redisDb         = flag.Int("redis_db", -1, "合并到的新redis redisDb（不填写默认读config的dbconfigs.redis.db）")
	redisPwd        = flag.String("redis_pass", "", "合并到的新redis redisAddr（不填写默认读config的dbconfigs.redis.password）")
	csvConf         = flag.String("csvConf", "", "表配置路径")
	showHelp        = flag.Bool("help", false, "show help")
)


func main() {
	logger.Init()
	defer logger.Close()

	flag.Parse()
	if *showHelp {
		flag.Usage()
		return
	}
	var err error
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())

	err = base.Conf.Load(*confFile)
	if err != nil {
		logger.Error("program start error when loading config: %s", err.Error())
		return
	}
	//非沙箱模式检查
	if !notSandBoxConfCheck() {
		return
	}
	//配置赋值
	confValue()

	m := managers.Get()
	err = m.Init()
	if err != nil {
		logger.Error("serverMerge start error at managers.Init: %s", err.Error())
		return
	}
	m.StartMerge()

}

func notSandBoxConfCheck() bool {

	if !base.Conf.SandBox {
		if *mergeServerIds == "" {
			logger.Error("合并服务器配置错误")
			return false
		}
		if *dbName == "" {
			logger.Error("新数据库dbName配置错误")
			return false
		}
		if *dbLogName == "" {
			logger.Error("新数据库dbLogName配置错误")
			return false
		}
		if *dbHost == "" {
			logger.Error("新数据库db_host配置错误")
			return false
		}
		if *dbPort == -1 {
			logger.Error("新数据库db_port配置错误")
			return false
		}
		if *dbUser == "" {
			logger.Error("新数据库db_user配置错误")
			return false
		}
		if *dbPass == "" {
			logger.Error("新数据库db_pass配置错误")
			return false
		}
		if *redisAddr == "" {
			logger.Error("新数据库redis_address配置错误")
			return false
		}
		if *redisPwd == "" {
			logger.Error("新数据库redis_pass配置错误")
			return false
		}
		if *redisDb == -1 {
			logger.Error("新数据库redis_db配置错误")
			return false
		}
	}
	return true
}

func confValue() {

	if *mergeServerIds != "" {
		base.Conf.MergeServerIds = *mergeServerIds
	}
	if *dbName != "" {
		base.Conf.DbConfigs[model.NEW_SERVER].DbName = *dbName
	}
	if *dbHost != "" {
		base.Conf.DbConfigs[model.NEW_SERVER].Host = *dbHost
	}
	if *dbPort != -1 {
		base.Conf.DbConfigs[model.NEW_SERVER].Port = *dbPort
	}
	if *dbUser != "" {
		base.Conf.DbConfigs[model.NEW_SERVER].Uid = *dbUser
	}
	if *dbPass != "" {
		base.Conf.DbConfigs[model.NEW_SERVER].Pwd = *dbPass
	}
	if *redisAddr != "" {
		base.Conf.Redis.Address = *redisAddr
	}
	if *redisPwd != "" {
		base.Conf.Redis.Password = *redisPwd
	}
	if *redisDb != -1 {
		base.Conf.Redis.DB = *redisDb
	}
	if *dbLogName != "" {
		base.Conf.DbLogName = *dbLogName
	}
	if *csvConf != "" {
		base.Conf.CsvConf = *csvConf
	}
}
