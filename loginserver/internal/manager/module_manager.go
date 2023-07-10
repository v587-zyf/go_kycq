package manager

import (
	"cqserver/gamelibs/ptsdk"
	"cqserver/gamelibs/rmodel"
	"cqserver/golibs/nw/httpserver"
	"cqserver/loginserver/conf"
	//"cqserver/gamelibs"

	"cqserver/golibs/dbmodel"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
)

const (
	SUPER_LOGIN_KEY = "1da5720b7a4e3e9f2de3c007804d4a02"
)
type ModuleManager struct {
	*util.DefaultModuleManager

	ServerList *ServerListManager
	//Gate       *GateManager

	//ServerLogin  *ServerLoginManager
	Announcement *AnnouncementManager
}

var m = &ModuleManager{
	DefaultModuleManager: util.NewDefaultModuleManager(),
}

func Get() *ModuleManager {
	return m
}

func (this *ModuleManager) init() error {

	//init db
	err := dbmodel.InitDb(conf.Conf, nil, "", nil)
	if err != nil {
		return err
	}

	rc := conf.Conf.Redis
	err = rmodel.Init(rc, 0)
	if err != nil {
		logger.Error("redis 连接错误:%v", err)
	}

	ptsdk.InitSdk(conf.Conf.Sdkconfig, conf.Conf.Sandbox)
	return err
}

func (this *ModuleManager) Init() error {

	err := this.init()
	if err != nil {
		return err
	}

	this.ServerList = this.AppendModule(NewServerListManager()).(*ServerListManager)
	//this.Gate = this.AppendModule(NewGateManager()).(*GateManager)
	//this.ServerLogin = this.AppendModule(NewServerLoginManager()).(*ServerLoginManager)
	this.Announcement = this.AppendModule(NewAnnouncementManager()).(*AnnouncementManager)

	err = this.DefaultModuleManager.Init()
	if err != nil {
		logger.Error("模块初始化错误：%v", err)
		return err
	}
	//启动http服务
	err = HttpServerInit()
	if err != nil {
		logger.Error("http启动错误：%v", err)
		return err
	}
	return nil
}

func (this *ModuleManager) Stop() {
	httpserver.Stop()
	this.DefaultModuleManager.Stop()
}
