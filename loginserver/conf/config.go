package conf

import (
	"fmt"
	"os"

	"cqserver/gamelibs/beans"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
)

type Config struct {
	Sandbox     bool
	LoginServer int
	Redis       *beans.RedisConfig
	//客户端跳转地址，602用
	ClientAddr      string
	ClientAddrTrail string
	//正式服登录地址
	Wssaddr string
	//提审服登录地址
	WssaddrTrail string

	DbConfigs map[string]*beans.DbConfig
	Sdkconfig *beans.Sdkconfig
	callbacks []func()
}

var Conf = &Config{}
var initConfStr = `
sandbox=false
loginServer=1
clientAddr="http://127.0.0.1:8081/"	
clientAddrTrail="http://127.0.0.1:8081/"
wssaddr = "https://ylwss.hgame.com"
wssaddrTrail = "https://ylwss.hgame.com"

[redis]
	network="tcp"
	address="127.0.0.1:6379"
	db=0
	password=""

[dbconfigs.accountdb]
	host="127.0.0.1"
	port=3306
	uid="root"
	pwd=""
	dbname="accountdb"
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
	err := util.UnmarshalTomlStr(initConfStr, Conf)
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
