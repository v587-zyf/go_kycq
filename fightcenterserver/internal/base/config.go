package base

import (
	"fmt"
	"os"

	"cqserver/gamelibs/beans"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
)

type Config struct {
	Sandbox      bool

	DbConfigs map[string]*beans.DbConfig //如果配置多个数据库源，则用逗号分隔源的名字

	callbacks []func()
}

var Conf = &Config{}
var log = logger.Get("default", true)
var InitConfStr = `
sandbox=true

[dbconfigs.accountdb]
  dbname="yulong_accountdb"
  host="192.168.5.173"
  port=3306
  pwd = "123456"
  uid="root"
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
	err := util.UnmarshalTomlStr(InitConfStr, Conf)
	if err != nil {
		return err
	}
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		logger.Warn("config.Load file:%s not exist,use default config!", fileName)
		return nil
	}
	err = util.UnmarshalToml(fileName, Conf)
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
