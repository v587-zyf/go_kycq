package main

import (
	"cqserver/fightserver/conf"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/fightModule"
	_ "cqserver/fightserver/internal/handler"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/dbmodel"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"errors"
)

type ModuleManager struct {
	*util.DefaultModuleManager
	serverSeq int

	gateManager  *net.GateManager
	gsManager    *net.GSManager
	fightManager *fightModule.FightManager
	CronManager  *fightModule.CronManager
}

var m = &ModuleManager{
	DefaultModuleManager: util.NewDefaultModuleManager(),
}

func (this *ModuleManager) init() error {
	var err error
	err = gamedb.Load(*gameDbBasePath)
	if err != nil {
		return err
	}

	this.readyBefore()
	err = dbmodel.InitDb(conf.Conf, nil, "", nil)
	if err != nil {
		return err
	}

	return err
}

func (this *ModuleManager) Init() error {
	err := this.init()
	if err != nil {
		return err
	}
	gatePort, gsPort := this.getPort()
	if gatePort == -1 || gsPort == -1 {
		return errors.New("程序端口获取失败")
	}
	defer func() {
		logger.Info("程序模块初始化完成，启动端口，gatePort：%v,gsPort:%v", gatePort, gsPort)
	}()

	this.fightManager = this.AppendModule(fightModule.NewFightManager()).(*fightModule.FightManager)
	this.CronManager = this.AppendModule(fightModule.NewCronManager()).(*fightModule.CronManager)
	this.gateManager = this.AppendModule(net.NewGateManager(gatePort)).(*net.GateManager)
	this.gsManager = this.AppendModule(net.NewGSManager(gsPort)).(*net.GSManager)

	return this.DefaultModuleManager.Init()
}

func (this *ModuleManager) readyBefore(){
	//移动间隔计算
	base.MoveInterval = int64(float64(constFight.SCENE_SIZE) / float64(gamedb.GetConf().Speedparam) * constFight.MOVE_CHECK_FIX * 1000000000)
}

//获取启动端口
func (this *ModuleManager) getPort() (gatePort int, gsPort int) {

	if conf.Conf.IsCorssFightServer() {
		serverInfo, err := modelCross.GetCrossFightServerInfoModel().GetCrossFightServerInfo(conf.Conf.ServerId)
		if err != nil {
			return -1, -1
		}
		return serverInfo.GatePort, serverInfo.GsPort

	} else {
		serverInfo, err := modelCross.GetServerInfoModel().GetServerInfoByServerId(conf.Conf.ServerId)
		if err != nil || serverInfo == nil {
			return -1, -1
		}
		gatePort = serverInfo.GatefsPort
		gsPort = serverInfo.GsfsPort
	}
	return gatePort, gsPort
}
