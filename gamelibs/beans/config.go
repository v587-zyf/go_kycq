package beans

type DbConfig struct {
	Host       string
	Port       int
	Uid        string
	Pwd        string
	DbName     string
	MaxIdle    int
	MaxOpenCon int
}

type RedisConfig struct {
	Network  string
	Address  string
	DB       int
	Password string
}

type Sdkconfig struct {
	KyprojectName string //恺英项目名字
	KyAppId       int    //恺英appId
	KyPlatformId  int    //恺英平台ID
	KyToken       string //恺英token
	KySecretKey   string //恺英加密key
	KyChatReport  string //恺英聊天上报
	KyGmUrl       string //恺英gm后台地址
	PlatformId    int    //平台id
	Gameid        int    //游戏Id
	Gamekey       string //加密key
	Paykey        string //充值加密key
	Verifyurl     string //登录验证地址
	ChatUrl       string //聊天上报地址
}
