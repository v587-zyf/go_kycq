package managers

import (
	"cqserver/crosscenterserver/internal/conf"
	"cqserver/crosscenterserver/internal/managers/activeUser"
	"cqserver/crosscenterserver/internal/managers/challengeCcs"
	"cqserver/crosscenterserver/internal/managers/httpManager"
	"cqserver/crosscenterserver/internal/managers/servers"
	"cqserver/crosscenterserver/internal/managers/shabake"
	"cqserver/crosscenterserver/internal/managers/user"
	"cqserver/crosscenterserver/internal/managersI"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/ptsdk"
	"cqserver/gamelibs/rmodelCross"
	"cqserver/golibs/dbmodel"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"flag"
)

var (
	gameDbBasePath = flag.String("gamedb", "../../yulong/data/configs", "specify gamedb path")
)

var (
	_gameDb *gamedb.GameDb
)

func gameDb() *gamedb.GameDb {
	return _gameDb
}

// 各manager管理
type ModuleManager struct {
	*util.DefaultModuleManager

	ServerListManager managersI.IGsServers
	UserManager       managersI.IUser
	ChallengeCcs      managersI.IChallengeCcs
	ShaBakeCcs        managersI.IShaBake
	ActiveUser        managersI.IActiveUser
	Cron              *CronManager
	GSManager         *servers.GSManager
}

var m = &ModuleManager{
	DefaultModuleManager: util.NewDefaultModuleManager(),
}

func Get() *ModuleManager {
	return m
}

func (this *ModuleManager) init() error {
	var err error
	//*gameDbBasePath = "E:/chuanqi/config/gamedb.dat"

	err = gamedb.Load(*gameDbBasePath)
	if err != nil {
		logger.Debug("err:%v  *gameDbBasePath:%v", err, *gameDbBasePath)
		return err
	}
	err = dbmodel.InitDb(conf.Conf, nil, "", nil)
	//err = dbmodel.InitDb(conf.Conf, []string{model.DB_ACCOUNT}, model.DB_ACCOUNT, nil)
	if err != nil {
		return err
	}

	//初始化redis
	confs, err := modelCross.GetCrossRedisModel().GetCrossRedisConfs()
	if err != nil {
		return err
	}
	for _, conf := range confs {
		err = rmodelCross.InitCrossMap(conf.Id, conf.Network, conf.Address, conf.Password, conf.Db)
		if err != nil {
			logger.Info("ModuleManager.InitCrossRedis err=%v", err)
			return err
		}
	}

	//system_setting system_param 检查
	err = rmodelCross.GetSystemSeting().InitCheck()
	if err != nil {
		return err
	}
	idsName := make([]string, 0)
	idsName = append(idsName, "equip", "guild", "user")
	modelCross.GetIdModel().CheckIdsCfg(idsName)

	//平台sdk初始化
	//platform := rmodelCross.GetSystemSeting().GetSystemSetting(rmodelCross.SYSTEM_SETTING_AREA_NAME)
	ptsdk.InitSdk(conf.Conf.Sdkconfig,conf.Conf.Sandbox)
	return nil
}

func (this *ModuleManager) Init() error {
	err := this.init()
	if err != nil {
		return err
	}
	this.ServerListManager = this.AppendModule(servers.NewServerListManager(this)).(managersI.IGsServers)
	this.UserManager = this.AppendModule(user.NewUserManager(this)).(managersI.IUser)
	this.ChallengeCcs = this.AppendModule(challengeCcs.NewChallengeCcsManager(this)).(managersI.IChallengeCcs)
	this.ShaBakeCcs = this.AppendModule(shabake.NewShaBakeManager(this)).(managersI.IShaBake)
	this.ActiveUser = this.AppendModule(activeUser.NewActiveUserManager(this)).(managersI.IActiveUser)
	this.Cron = this.AppendModule(NewCronManager()).(*CronManager)

	//放在最后
	this.GSManager = this.AppendModule(servers.NewGSManager(this)).(*servers.GSManager)

	err = this.DefaultModuleManager.Init()
	if err != nil {
		return err
	}
	err = httpManager.HttpServerInit(this)
	if err != nil {
		return err
	}
	return nil
}

func (this *ModuleManager) Start() error {
	return this.DefaultModuleManager.Start()
}

func (this *ModuleManager) Stop() {
	this.DefaultModuleManager.Stop()
}

func (this *ModuleManager) GetGsServers() managersI.IGsServers {
	return this.ServerListManager
}
func (this *ModuleManager) GetUser() managersI.IUser {
	return this.UserManager
}

func (this *ModuleManager) GetCcsChallenge() managersI.IChallengeCcs {
	return this.ChallengeCcs
}

func (this *ModuleManager) GetShaBakeCcs() managersI.IShaBake {
	return this.ShaBakeCcs
}

func (this *ModuleManager) GetActiveUser() managersI.IActiveUser {
	return this.ActiveUser
}
