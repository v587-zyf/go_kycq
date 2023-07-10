package manager

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/rmodel"
	"cqserver/gateserver/conf"
	"cqserver/golibs/common"
	"cqserver/golibs/dbmodel"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"strconv"
	"time"
)

type ModuleManager struct {
	*util.DefaultModuleManager
	serverInfo *modelCross.ServerInfo

	ClientManager *ClientManager
	GsManager     *GSManager
	FsManager     *FSManager
	//LsManager     *LSManager
}

var m = &ModuleManager{
	DefaultModuleManager: util.NewDefaultModuleManager(),
}

func Get() *ModuleManager {
	return m
}

func (this *ModuleManager) init() error {
	var err error

	err = dbmodel.InitDb(conf.Conf, nil, "", nil)
	if err != nil {
		logger.Error("db model init err:%v", err)
		return err
	}

	serverInfo, err := modelCross.GetServerInfoModel().GetServerInfoByServerId(conf.Conf.ServerId)
	if err != nil {
		return err
	}
	this.serverInfo = serverInfo

	rc := conf.Conf.Redis
	err = rmodel.Init(rc, 0)
	if err != nil {
		logger.Error("redis 连接错误:%v", err)
		return err
	}

	return nil
}

func (this *ModuleManager) Init() error {
	err := this.init()
	if err != nil {
		return err
	}

	//this.LsManager = this.AppendModule(NewLSManager()).(*LSManager)
	this.ClientManager = this.AppendModule(NewClientManager(this.GetServerPort())).(*ClientManager)
	this.GsManager = this.AppendModule(NewGSManager(this.serverInfo.GsPort)).(*GSManager)
	this.FsManager = this.AppendModule(NewFSManager(this.serverInfo.GatefsPort)).(*FSManager)
	err = this.DefaultModuleManager.Init()
	if err != nil {
		return err
	}
	//设置服务器启动完成
	this.ClientManager.Started()
	go this.reloadServerInfo()
	return nil
}

func (this *ModuleManager) GetServerPort() int {

	strSlice := common.NewStringSlice(this.serverInfo.Gates, ":")
	port, _ := strconv.Atoi(strSlice[1])
	return port
}

// 重新加载服务器信息
func (this *ModuleManager) reloadServerInfo() {
	tickerReload := time.NewTicker(time.Second * 10)

	for {
		select {
		case <-tickerReload.C:
			serverInfo, err := modelCross.GetServerInfoModel().GetServerInfoByServerId(conf.Conf.ServerId)
			if err != nil {
				logger.Error("reloadServerTicker DB Error: %v", err)
				return
			}
			if this.serverInfo.CrossFsId != serverInfo.CrossFsId {
				logger.Info("跨服战斗服改变,旧：%v,新：%v", this.serverInfo.CrossFsId, serverInfo.CrossFsId)
				this.serverInfo.CrossFsId = serverInfo.CrossFsId
			}
		}
	}
}
