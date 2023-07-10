package conf

import (
	"fmt"
	"os"

	"cqserver/gamelibs/beans"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
)

type Config struct {
	ServerId  int
	RobotNum  int
	callbacks []func()

	DbConfigs   map[string]*beans.DbConfig //如果配置多个数据库源，则用逗号分隔源的名字
	Create		map[string]map[string]interface{}
}

var Conf = &Config{}
var InitConfStr = `
serverid=1

[dbconfigs.accountdb]
  dbname="yulong_accountdb"
  host="127.0.0.1"
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
	err := util.UnmarshalTomlStr(InitConfStr, this)
	if err != nil {
		return err
	}
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		logger.Info("config.Load file:%s not exist,use default config!", fileName)
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
