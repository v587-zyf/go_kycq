package conf

import (
	"fmt"
	"cqserver/golibs/logger"
	"os"

	"cqserver/gamelibs/beans"
	"cqserver/golibs/util"
)

type Config struct {
	Sandbox   bool
	DbConfigs map[string]*beans.DbConfig //如果配置多个数据库源，则用逗号分隔源的名字
	Sdkconfig *beans.Sdkconfig
}

var Conf = &Config{}
var InitConfStr = `
sandbox=false

[dbconfigs.accountdb]
  dbname="slg_accountdb"
  host="127.0.0.1"
  port=3306
  pwd = "123456"
  uid="root"
  maxIdle = 3
  maxOpenCon = 5

[sdkconfig]
  kyprojectName="chuanqih5"
  kyAppId=7
  kyPlatformId=2
  kyToken="5Lyg5aWHaDUyMjMyM2Rzc2ZzZGZzZAsR"
  kySecretKey="fe02b4b50e7f48e5815a3abe2b3e64f9"
  kyChatReport="https://cmsrv.kingnetdc.com"
  KyGmUrl="https://gm-gateway-ky.kingnet.com"
  platformId=45
  gameid=100106
  gamekey="20220711WOMENDECHUANQI"
  paykey="WOMENDECHUANQI20220711"
  verifyurl="http://verify.junshanggame.com/cp/token/"
  chatUrl="http://cm.602.com/v3/push"
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
	return nil
}
