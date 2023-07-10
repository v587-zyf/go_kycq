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

type GSManager struct {
	util.DefaultModule

	*nw.DefaultSessionManager
	server nw.Server
}

func NewGSManager() *GSManager {
	return &GSManager{
		DefaultSessionManager: nw.NewDefaultSessionManager(true),
	}
}

func (this *GSManager) Init() error {

	serverPortInfo, err := modelCross.GetServerPortInfoModel().GetServerPortInfo(modelCross.GAME_TO_FIGHTCENTER)
	if err != nil {
		return err
	}

	context := &nw.Context{
		SessionCreator: NewGSSession,
		Splitter:       pbserver.Split,
		ReadBufferSize: 1024 * 1024,
		ChanSize:       1024 * 4,
	}
	server := tcpserver.NewServer(context)
	err = server.Start(fmt.Sprintf(":%d",serverPortInfo.Port))
	if err != nil {
		return err
	}
	this.server = server
	logger.Info("端口监听【%v】,game连接使用", serverPortInfo.Port)
	return nil
}

func (this *GSManager) GetGSSession(sessionId uint32) *GSSession {
	session := this.GetSession(sessionId)
	if session != nil {
		return session.(*GSSession)
	}
	return nil
}

func (this *GSManager) Stop() {
	if this.server != nil {
		this.server.Stop()
	}
}
