package net

import (
	"cqserver/fightserver/conf"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbgt"
	"runtime/debug"
)

type GateSession struct {
	Conn nw.Conn
}

func NewGateSession(conn nw.Conn) nw.Session {
	s := &GateSession{
		Conn: conn,
	}
	return s
}

func (this *GateSession) GetId() uint32 {
	return uint32(conf.Conf.ServerId)
}

func (this *GateSession) GetConn() nw.Conn {
	return this.Conn
}

func (this *GateSession) OnOpen(conn nw.Conn) {
	logger.Info("nw_gate_session.go:OnOpen")
}

func (this *GateSession) OnClose(conn nw.Conn) {
	logger.Info("nw_gate_session.go:OnClose")
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
		logger.Error("nw_gate_session unmarshal err:%v", err)
		return
	}

	msgId = int(msgFrame.CmdId)
	if msgFrame.CmdId == pbgt.CmdHandShakeReqId {
		this.OnHandShakeReq(msgFrame)
		return
	}

	if conf.Conf.IsCorssFightServer() && msgId == pbgt.CmdGateMessageToFSId { //动态战斗服
		msgFrame, _, err = GetFccRouteFromGateMsgFrame(msgFrame)
		if err != nil {
			logger.Error("nw_gate_session OnRecv GetGateNewMsgFrame err:%v", err)
			return
		}
	}

	handler := pbgt.GetHandler(msgFrame.CmdId)
	if handler != nil {
		handler(this.Conn, msgFrame)
	} else {
		logger.Error("未注册的cmd：%v", msgFrame.CmdId)
	}
}

func (this *GateSession) OnHandShakeReq(msgFrame *pbgt.MessageFrame) {
	req := msgFrame.Body.(*pbgt.HandShakeReq)
	if int(req.GateSeq) != conf.Conf.ServerId {
		logger.Error("连接错误，本服serverId：%v，连接上来的为：%v", conf.Conf.ServerId, req.GateSeq)
		this.Conn.Close()
		return
	}

	msg := &pbgt.HandShakeAck{ServerSeq: int32(conf.Conf.ServerId)}
	rb, err := pbgt.Marshal(pbgt.GetCmdIdFromType(msg), 0, 0, msg)
	if err != nil {
		logger.Error("GateSession marshal error: %s", err.Error())
		return
	}
	this.Conn.Write(rb)
	if conf.Conf.IsCorssFightServer() {
		logger.Info("与战斗中心握手成功")
	} else {
		logger.Info("与gate握手成功：%v", req.GateSeq)
	}
}

func (this *GateSession) SendMessage(serverId, sessionId, transId uint32, msg nw.ProtoMessage) error {

	//编译发送客户端消息
	cmdId := pb.GetCmdIdFromType(msg)
	r, err := pb.Marshal(cmdId, transId, msg)
	if err != nil {
		return err
	}
	//需要经过gate网关转发，编译一层给gate
	rb, err := pbgt.Marshal(pbgt.CmdRouteMessageId, sessionId, 0, r)
	if err != nil {
		return err
	}
	//跨发战斗服，需要经战斗中心转发
	if conf.Conf.IsCorssFightServer() {
		rb, err = GetFccRouteToGateMessageBytes(serverId, rb)
		if err != nil {
			return err
		}
	}
	_, err = this.Conn.Write(rb)
	if err != nil {
		logger.Error("gate推送消息错误：%v", err)
	}
	return err
}

func (this *GateSession) BroadcastToGate(ids map[int]map[int]int, msg nw.ProtoMessage) {

	cmdId := pb.GetCmdIdFromType(msg)
	r, err := pb.Marshal(cmdId, 0, msg)
	if err != nil {
		return
	}
	for k, v := range ids {
		sessionIds := make(map[uint32]uint32)
		for kk, vv := range v {
			sessionIds[uint32(kk)] = uint32(vv)
		}
		gateMsg := &pbgt.BroadcastByFS{
			SessionIds: sessionIds,
			Msg:        r,
			MsgId:      int32(cmdId),
		}

		rb, err := pbgt.Marshal(pbgt.GetCmdIdFromType(gateMsg), 0, 0, gateMsg)
		if err != nil {
			logger.Error("BroadcastToGate err:%v", err)
			return
		}
		if conf.Conf.IsCorssFightServer() {
			rb, err = GetGateNewMsgBytesNormal(int(k), rb)
		}

		if err != nil {
			logger.Error("BroadcastToGate err:%v", err)
			return
		}
		this.Conn.Write(rb)
	}
}
