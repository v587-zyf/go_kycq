package manager

import (
	"cqserver/golibs/logger"
	"cqserver/protobuf/pbgt"
	"fmt"
	"sync"
	"time"

	"cqserver/golibs/nw"
	"cqserver/golibs/nw/wsserver"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

type ClientManager struct {
	util.DefaultModule

	*nw.DefaultSessionManager
	sessionMu sync.RWMutex
	server    nw.Server
	done      chan struct{}
	port      int
}

func NewClientManager(port int) *ClientManager {
	return &ClientManager{
		DefaultSessionManager: nw.NewDefaultSessionManager(true),
		port:                  port,
	}
}

func (this *ClientManager) Init() error {
	context := &nw.Context{
		SessionCreator: func(conn nw.Conn) nw.Session { return NewClientSession(conn) },
		Splitter:       pb.Split,
		ChanSize:       200,
	}
	server := wsserver.NewServer(context)
	err := server.Start(fmt.Sprintf(":%d", this.port))
	if err != nil {
		return err
	}
	this.server = server
	logger.Info("端口监听【%v】,客户端连接使用", this.port)
	return nil
}

func (this *ClientManager) Started() {
	logger.Info("服务器启动完成，可链接")
	this.server.SetAcceptConnect(true)
}

func (this *ClientManager) GetClientSession(id uint32) *ClientSession {
	session := this.GetSession(id)
	if session == nil {
		return nil
	}
	return session.(*ClientSession)
}

func (this *ClientManager) Broadcast(sessionIds []uint32, data []byte) {
	if sessionIds != nil && len(sessionIds) > 0 {

		for _, v := range sessionIds {
			clientSession := m.ClientManager.GetClientSession(v)
			if clientSession == nil || !clientSession.broadcast {
				continue
			}
			clientSession.WriteToClient(data)
		}

	} else {

		this.Range(func(id uint32, session nw.Session) bool {
			if clientSession, ok := session.(*ClientSession); ok {
				if clientSession.broadcast {
					clientSession.WriteToClient(data)
				}
			}
			return false
		})
	}
}

func (this *ClientManager) Stop() {
	if this.server != nil {
		this.server.Stop()
		logger.Info("nw_client stop")
	}
}

func (this *ClientManager) CloseClientSession(clientSessionId uint32, closeReason string) {
	//session := this.GetSession(clientSessionId)
	session := this.GetClientSession(clientSessionId)
	if session == nil {
		logger.Warn("关闭客户端连接，session不存在")
		return
	}
	session.isClose = true
	logger.Info("关闭客户端连接openId:%v,id:%v", session.GetOpenId(), session.GetId())
	kickMsg := &pb.KickUserNtf{
		Reason: closeReason,
	}
	rb, err1 := pb.Marshal(pb.CmdKickUserNtfId, 0, kickMsg)
	if err1 != nil {
		logger.Debug("marshal kick user openId:%v ntf err:%v", session.GetOpenId(), err1)
		return
	}
	err := session.WriteToClient(rb)
	if err != nil {
		logger.Debug("send kick user openId:%v ntf err:%v", session.GetOpenId(), err)
		return
	}
	time.Sleep(time.Millisecond * 3)
	session.GetConn().Close()
}

func (this *ClientManager) BindFightInfo(msg *pbgt.MessageFrame) {
	session := this.GetSession(msg.SessionId)
	if session == nil {
		logger.Warn("推送gate玩家战斗信息，session不存在")
		return
	}
	fightInfoMsg := msg.Body.(*pbgt.UserFightInfoNtf)
	session.(*ClientSession).BindFightSession(fightInfoMsg.FightId, int(fightInfoMsg.CrossFightServerId))
}

func (this *ClientManager) GsClose() {
	this.DefaultSessionManager.Range(func(id uint32, session nw.Session) bool {
		session.GetConn().Close()
		return false
	})

}
