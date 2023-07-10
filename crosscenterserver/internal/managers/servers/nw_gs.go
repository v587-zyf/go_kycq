package servers

import (
	"cqserver/crosscenterserver/internal/managersI"
	"cqserver/gamelibs/modelCross"
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/tcpserver"
	"cqserver/golibs/util"
	"cqserver/protobuf/pbserver"
	"fmt"
)

type GSManager struct {
	util.DefaultModule
	managersI.IModule
	*nw.DefaultSessionManager
	server nw.Server
}

var gsManager *GSManager

func NewGSManager(m managersI.IModule) *GSManager {
	gsManager = &GSManager{
		IModule:               m,
		DefaultSessionManager: nw.NewDefaultSessionManager(true),
	}
	return gsManager
}

func (this *GSManager) Init() error {
	//获取启动端口号
	infos, err := modelCross.GetServerPortInfoModel().GetServerPortInfo(modelCross.GAME_TO_CROSSCENTER)
	if err != nil || infos == nil {
		return nil
	}

	context := &nw.Context{
		SessionCreator: NewGsSession,
		Splitter:       pbserver.Split,
		ReadBufferSize: 1024 * 64,
	}
	server := tcpserver.NewServer(context)
	err = server.Start(fmt.Sprintf(":%d", infos.Port))
	if err != nil {
		return err
	}
	this.server = server
	return nil
}

func (this *GSManager) GetGsSession(sessionId uint32) *GsSession {
	session := this.GetSession(sessionId)
	if session != nil {
		return session.(*GsSession)
	}
	return nil
}

func (this *GSManager) Stop() {
	if this.server != nil {
		this.server.Stop()
	}
}
