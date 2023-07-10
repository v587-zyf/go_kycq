package base

import (
	"fmt"
	"os"

	"cqserver/gamelibs/beans"
	"cqserver/golibs/util"
)

type Config struct {
	SandBox         bool
	//合服id
	MergeServerIds  string
	//N天内活跃(不做条件填写-1)
	ActiveDay int
	//最少充值额(不做条件填写-1)
	RechargeMin int
	//等级需求(不做条件填写-1)
	LevelMin int
	//战力需求(不做条件填写-1)
	CombatMin int
	//表配置路径
	CsvConf string

	DbLogName string
	Redis     *beans.RedisConfig
	DbConfigs map[string]*beans.DbConfig //如果配置多个数据库源，则用逗号分隔源的名字
	callbacks []func()
}

var Conf = &Config{}
var InitConfStr = `
SandBox = false
MergeServerIds="1,2"

activeDay = 7
rechargeMin = 0
levelMin = -1
combatMin = -1
csvConf = "./config/excels"

dbLogName = "yulong_log"
[dbconfigs.accountdb]
  dbname="accountdb"
  host="127.0.0.1"
  port=3306
  pwd = "123456"
  uid="root"
  maxIdle = 3
  maxOpenCon = 5

[dbconfigs.serverdb]
  dbname = "yulong_server"
  host = "127.0.0.1"
  port = 3306
  pwd = "123456"
  uid = "root"
  maxIdle = 3
  maxOpenCon = 5
`

func (this *Config) GetDbConnectionString(name string) (string, int, int) {
	dbConfig, ok := this.DbConfigs[name]
	if ok {
		str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConfig.Uid, dbConfig.Pwd, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
		str += "?charset=utf8&timeout=5s&parseTime=true&loc=Asia%2FShanghai"
		return str, dbConfig.MaxIdle, dbConfig.MaxOpenCon
	}
	return "", 0, 0
}

func (this *Config) Load(fileName string) error {
	err := util.UnmarshalTomlStr(InitConfStr, this)
	if err != nil {
		return err
	}
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		//logger.Info("config.Load file:%s not exist,use default config!", fileName)
		return nil
	}
	err = util.UnmarshalToml(fileName, this)
	if err != nil {
		return err
	}
	for _, callback := range this.callbacks {
		callback()
	}
	return nil
}

func (this *Config) RegisterCallback(callback func()) {
	this.callbacks = append(this.callbacks, callback)
}
