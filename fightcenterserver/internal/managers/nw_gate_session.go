package managers

import (
	"runtime/debug"

	"cqserver/golibs/nw"
	"github.com/gogo/protobuf/proto"

	//"cqserver/protobuf/pb"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pbgt"
)

type GateSession struct {
	Conn    nw.Conn
	GateSeq int //客户端传过来的serverId
}

func NewGateSession(conn nw.Conn) nw.Session {
	s := &GateSession{
		Conn: conn,
	}
	return s
}

func (this *GateSession) GetId() uint32 {
	return uint32(this.GateSeq)
}

func (this *GateSession) GetConn() nw.Conn {
	return this.Conn
}

func (this *GateSession) OnOpen(conn nw.Conn) {
	logger.Info("nw_gate_session.go:OnOpen")
}

func (this *GateSession) OnClose(conn nw.Conn) {

	if this == m.gateManager.GetSession(uint32(this.GateSeq)) {
		m.gateManager.RemoveSession(uint32(this.GateSeq))
		logger.Info("nw_gate_session.go:OnClose GateSeq=%v curSessionCount=%v", this.GateSeq, m.gateManager.GetSessionCount())
	} else {
		logger.Info("nw_gate_session.go:OnClose error GateSeq=%v curSessionCount=%v", this.GateSeq, m.gateManager.GetSessionCount())
	}
}

func (this *GateSession) OnRecv(conn nw.Conn, data []byte) {
	var msgId int

	defer func() {
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic nw_gate recv msg[%v]:%v,%s", msgId, r, stackBytes)
		}
	}()

	msgFrame, err := pbgt.Unmarshal(data)
	if err != nil {
		logger.Error("nw_gate_session OnRecv unmarshal err:%v", err)
		return
	}
	msgId = int(msgFrame.CmdId)
	switch msgFrame.CmdId {
	case pbgt.CmdHandShakeReqId:
		this.OnHandShakeReq(msgFrame)
	case pbgt.CmdGateMessageToFSId: //新的转发到fs
		this.GateMessageToFS(data, msgFrame)
	}
}

func (this *GateSession) GateMessageToFS(data []byte, msgFrame *pbgt.MessageFrame) error {
	msg := msgFrame.Body.(*pbgt.GateMessageToFS)
	return m.fsManager.RouteGateMessageToFs(int(msg.CrossServerId), data)
}

func (this *GateSession) OnHandShakeReq(msgFrame *pbgt.MessageFrame) {
	req := msgFrame.Body.(*pbgt.HandShakeReq)
	this.GateSeq = int(req.GateSeq)

	this.SendMessage(&pbgt.HandShakeAck{})
	m.gateManager.AddSession(uint32(this.GateSeq), this)
	logger.Info("gate handshake ok. gateSeq=%v curSessionCount=%v", req.GateSeq, m.gateManager.GetSessionCount())

}

func (this *GateSession) SendMessage(msg proto.Message) {
	rb, err := pbgt.Marshal(pbgt.GetCmdIdFromType(msg), 0, 0, msg)
	if err != nil {
		logger.Error("GateSession marshal error: %s", err.Error())
		return
	}
	this.Conn.Write(rb)
}
