package gater

import (
	"cqserver/protobuf/pb"
	"time"

	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pbgt"
)

func init() {
	// cmd register begin
	pbgt.Register(pbgt.CmdHandShakeReqId, HandleHandShakeReq)
	pbgt.Register(pbgt.CmdUserQuitRptId, HandleUserQuitRpt)
	pbgt.Register(pbgt.CmdRouteMessageId, HandleRouteMessage)
	pbgt.Register(pbgt.CmdServerPingReqId, HandleGatePingReq)
	// cmd register end
}

func HandleGatePingReq(conn nw.Conn, msgFrame *pbgt.MessageFrame) {
	gateSession := conn.GetSession().(*GateSession)
	gateSession.Server.PingTime = int(time.Now().Unix())
	gateSession.SendMessage(pbgt.CmdServerPingAckId, 0, &pbgt.ServerPingAck{})
}

func HandleHandShakeReq(conn nw.Conn, msgFrame *pbgt.MessageFrame) {
	body := msgFrame.Body.(*pbgt.HandShakeReq)
	gateSession := conn.GetSession().(*GateSession)
	gateSession.Id = uint32(body.GateSeq)
	if gateSession.Server.GetSession(gateSession.Id) != nil {
		logger.Error("gate handshake error: %d already exists", gateSession.Id)
		conn.Close()
		return
	}
	gateSession.Server.AddSession(gateSession.Id, gateSession)
	logger.Info("gate%d handshake ok...............................................", gateSession.Id)
	msg := &pbgt.HandShakeAck{}
	gateSession.SendMessage(pbgt.CmdHandShakeAckId, 0, msg)
}

func HandleUserQuitRpt(conn nw.Conn, msgFrame *pbgt.MessageFrame) {
	gateSession := conn.GetSession().(*GateSession)
	clientSession := gateSession.clientSessions.GetSession(msgFrame.SessionId)
	if clientSession != nil {
		clientSession.GetConn().Close()
	} else {
		logger.Error("HandleUserQuitRpt clientSession nil:%v", msgFrame.SessionId)
	}
}

func HandleRouteMessage(conn nw.Conn, msgFrame *pbgt.MessageFrame) {
	gateSession := conn.GetSession().(*GateSession)

	body := msgFrame.Body.([]byte)
	bodyCmdId := pb.GetCmdId(body)
	if bodyCmdId == pb.CmdEnterGameReqId {
		enterGameFrame, _ := pb.Unmarshal(body)
		req := enterGameFrame.Body.(*pb.EnterGameReq)
		clientConn := NewClientConn(conn, msgFrame.SessionId, req.Ip, gateSession.Server.ClientContext)
		clientConn.ServeIO()
	}


	clientSession := gateSession.clientSessions.GetSession(msgFrame.SessionId)
	if clientSession == nil {
		logger.Debug("HandleRouteMessage gateSession not found: %d", msgFrame.SessionId)
		return
	}
	clientConn := clientSession.GetConn().(*ClientConn)
	data := make([]byte, len(body))
	copy(data, body) // body引用的直接是底层的缓存区，这里为了防止底层接收消息过快，而上层处理速度慢，而导致缓冲区覆盖
	clientConn.putMessage(data)
}
