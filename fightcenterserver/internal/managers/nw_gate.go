package managers

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/tcpserver"
	"cqserver/golibs/util"
	"cqserver/protobuf/pbserver"
	"fmt"
)

type GateManager struct {
	util.DefaultModule
	*nw.DefaultSessionManager

	server nw.Server
}

func NewGateManager() *GateManager {
	return &GateManager{
		DefaultSessionManager: nw.NewDefaultSessionManager(true),
	}
}

func (this *GateManager) Init() error {
	serverPortInfo, err := modelCross.GetServerPortInfoModel().GetServerPortInfo(modelCross.GATE_TO_FIGHTCENTER)
	if err != nil {
		return err
	}
	context := &nw.Context{
		SessionCreator:  NewGateSession,
		Splitter:        pbserver.Split,
		ReadBufferSize:  10 * 1024,
		WriteBufferSize: 2 * 1024 * 1024,
		ChanSize:        1024 * 8,
	}
	server := tcpserver.NewServer(context)
	err = server.Start(fmt.Sprintf(":%d", serverPortInfo.Port))
	if err != nil {
		return err
	}
	this.server = server
	logger.Info("端口监听【%v】,客户端连接使用", serverPortInfo.Port)
	return nil
}

func (this *GateManager) GetGateSession(sessionId uint32) *GateSession {
	session := this.GetSession(sessionId)
	if session != nil {
		return session.(*GateSession)
	}
	return nil
}

func (this *GateManager) Stop() {
	if this.server != nil {
		this.server.Stop()
	}
}
