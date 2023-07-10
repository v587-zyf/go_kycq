package handler

import (
	"cqserver/gateserver/internal/manager"
	"cqserver/gateserver/internal/pbclient"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbgt"
)

func init() {
	// cmd register begin
	pbgt.Register(pbgt.CmdUserQuitNtfId, HandleUserQuitNtf)
	pbgt.Register(pbgt.CmdUserFightInfoNtfId, HandleUserFightInfoNtf)
	pbgt.Register(pbgt.CmdRouteMessageId, HandleRouteMessage)
	pbgt.Register(pbgt.CmdBroadcastNtfId, HandleBroadcastNtf)
	pbgt.Register(pbgt.CmdBroadcastByFSId, HandleBroadcastByFS)
	// cmd register end
}

func HandleUserQuitNtf(gsConn nw.Conn, msgFrame *pbgt.MessageFrame) {
	msg := msgFrame.Body.(*pbgt.UserQuitNtf)
	m.ClientManager.CloseClientSession(msgFrame.SessionId,msg.Reason)
}

func HandleUserFightInfoNtf(gscon nw.Conn,msgFrame *pbgt.MessageFrame){
	m.ClientManager.BindFightInfo(msgFrame)
}

func HandleRouteMessage(serverConn nw.Conn, msgFrame *pbgt.MessageFrame) {
	var gsSession *manager.GSSession
	var serverSession = serverConn.GetSession()
	switch session := serverSession.(type) {
	case *manager.GSSession:
		gsSession = session
	case *manager.FSSession:
		gsSession = m.GsManager.GetSession()
	}
	if gsSession == nil {
		logger.Error("no gsSession found")
		return
	}
	clientSession := m.ClientManager.GetClientSession(msgFrame.SessionId)
	if clientSession == nil {
		logger.Debug("ClientSession not found on RouteMessage")
		return
	}
	body := msgFrame.Body.([]byte)
	data := make([]byte, len(body))
	copy(data, body) // body引用的直接是底层的缓存区，这里为了防止底层接收消息过快，而发给客户端速度慢，而导致缓冲区覆盖
	cmdId := pbclient.GetCmdId(data)

	handler := pbclient.GetServerHandler(cmdId)
	if handler != nil {
		msg, err := pbclient.UnmarshalServer(data)
		if err != nil {
			logger.Error("unmarshal client message(cmdId: %d) error: %s", cmdId, err.Error())
			return
		}
		needRoute := handler(serverConn, clientSession, msg)
		if !needRoute {
			return
		}
	}

	//logger.Debug("收到推送客户端的消息,cmdId:%v",cmdId)

	clientSession.WriteToClient(data)
}

func HandleBroadcastNtf(gsConn nw.Conn, msgFrame *pbgt.MessageFrame) {
	body := msgFrame.Body.(*pbgt.BroadcastNtf)
	m.ClientManager.Broadcast(body.SessionIds, body.Msg)
}

func HandleBroadcastByFS(gsConn nw.Conn, msgFrame *pbgt.MessageFrame) {
	body := msgFrame.Body.(*pbgt.BroadcastByFS)
	if body.MsgId == pb.CmdSceneEnterNtfId {
		logger.Info("msg from fs user enter fight sessionId:%v", body.SessionIds)
	}
	for id, _ := range body.SessionIds {
		clientSession := m.ClientManager.GetClientSession(id)
		if clientSession == nil {
			continue
		}
		clientSession.WriteToClient(body.Msg)
	}
}
