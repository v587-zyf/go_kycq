package manager

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
)

var log = logger.Get("default", true)

type ServerListManager struct {
	util.DefaultModule
	serverInfo *modelCross.ServerInfo
}

func NewServerListManager(serverInfo *modelCross.ServerInfo) *ServerListManager {
	return &ServerListManager{
		serverInfo: serverInfo,
	}
}

func (this *ServerListManager) Init() error {
	return nil
}

func (this *ServerListManager) Stop() {
	logger.Info("serverListManager stop")
}
