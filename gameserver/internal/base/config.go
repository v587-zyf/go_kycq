package base

import (
	"cqserver/gamelibs/beans"
	"cqserver/golibs/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	ServerId    int
	Sandbox     bool
	GmSwitch    bool
	DbConfigs   map[string]*beans.DbConfig //如果配置多个数据库源，则用逗号分隔源的名字
	Sdkconfig   *beans.Sdkconfig
	callbacks   []func()
}

var Conf = &Config{}
var InitConfStr = `
serverid=1
sandbox=false
gmSwitch=false

[dbconfigs.accountdb]
  dbname="accountdb"
  host="127.0.0.1"
  port=3306
  pwd = ""
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

func GetBeanDb(dblink string, dbLogLink string) (*beans.DbConfig, *beans.DbConfig, error) {

	dbLinkMap := make(map[string]string)
	dbSlice := strings.Split(dblink, ";")
	for _, v := range dbSlice {
		vv := strings.Split(v, "=")
		dbLinkMap[vv[0]] = vv[1]
	}
	if dbLinkMap["server"] == "" ||
		dbLinkMap["database"] == "" ||
		dbLinkMap["uid"] == "" ||
		dbLinkMap["pwd"] == "" ||
		dbLinkMap["port"] == "" {
		return nil, nil, fmt.Errorf("服务器数据库配置错误:%v", dblink)
	}

	dbConfig := &beans.DbConfig{
		Host:   dbLinkMap["server"],
		DbName: dbLinkMap["database"],
		Uid:    dbLinkMap["uid"],
		Pwd:    dbLinkMap["pwd"],
	}
	port, err := strconv.Atoi(dbLinkMap["port"])
	if err != nil {
		return nil, nil, fmt.Errorf("服务器数据库端口配置错误:%v", dblink)
	}
	dbConfig.Port = port

	dbloglinkSlice := strings.Split(dbLogLink, ":")
	if len(dbloglinkSlice) != 2 {
		return nil, nil, fmt.Errorf("服务器log数据库配置错误：%v", dbLogLink)
	}
	dbloglinkSlice1 := strings.Split(dbloglinkSlice[1], "/")
	if len(dbloglinkSlice1) != 2 {
		return nil, nil, fmt.Errorf("服务器log数据库配置错误：%v", dbLogLink)
	}

	dbLogConfig := &beans.DbConfig{
		Host:   dbloglinkSlice[0],
		DbName: dbloglinkSlice1[1],
		Uid:    dbLinkMap["uid"],
		Pwd:    dbLinkMap["pwd"],
	}
	port, err = strconv.Atoi(dbloglinkSlice1[0])
	if err != nil {
		return nil, nil, fmt.Errorf("服务器log数据库端口配置错误:%v", dbLogLink)
	}
	dbLogConfig.Port = port
	return dbConfig, dbLogConfig, nil
}

//获取转化redis配置
func GetRedisConf(redisAddr string) (*beans.RedisConfig, error) {

	redisSlice := strings.Split(redisAddr, ":")
	if len(redisSlice) < 3 {
		return nil, fmt.Errorf("服务器redis配置错误,需配置对应ip:port:pwd:db,db可默认为0,当前配置%v", redisAddr)
	}
	redisC := &beans.RedisConfig{
		Network:  "tcp",
		Address:  redisSlice[0] + ":" + redisSlice[1],
		Password: redisSlice[2],
	}
	db := 0
	if len(redisSlice) == 4 {
		d, err := strconv.Atoi(redisSlice[3])
		if err != nil {
			return nil, fmt.Errorf("服务器redis端口配置错误,需配置对应ip:port:pwd:db,当前配置%v", redisAddr)
		}
		db = d
	}
	redisC.DB = db
	return redisC, nil
}
