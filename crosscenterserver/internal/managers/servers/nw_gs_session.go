package servers

import (
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/rpc"
	"cqserver/protobuf/pbserver"
	"errors"
)

type GsSession struct {
	Conn nw.Conn
	GsNo int //gs端传过来的 serverId
	rpc  *rpc.RpcWrapper
}

func NewGsSession(conn nw.Conn) nw.Session {
	s := &GsSession{
		Conn: conn,
	}
	return s
}

func (this *GsSession) GetId() uint32 {
	return uint32(this.GsNo)
}

func (this *GsSession) GetConn() nw.Conn {
	return this.Conn
}

func (this *GsSession) OnOpen(conn nw.Conn) {
	this.rpc = rpc.NewRpcWrapper(conn, 0, pbserver.RpcMarshaler, pbserver.RpcUnmarshaler)
}

func (this *GsSession) OnClose(conn nw.Conn) {

	if this == gsManager.GetSession(uint32(this.GsNo)) {
		gsManager.RemoveSession(uint32(this.GsNo))
	}
}

func (this *GsSession) OnRecv(conn nw.Conn, data []byte) {
	//msgFrame, err := pbserver.Unmarshal(data)
	//if err != nil {
	//	logger.Error("unmarshal gate data error: %s", err.Error())
	//	return
	//}
	msg, err := this.rpc.HookRecv(data)
	if err != nil {
		logger.Info("unmarshal gs data error: %s", err.Error())
		return
	} else if msg == nil {
		return
	}

	msgFrame := msg.(*pbserver.MessageFrame)
	msgId := msgFrame.CmdId
	handler := pbserver.GetHandler(msgId)

	if msgId == pbserver.CmdHandShakeReqId {
		this.OnHandShakeReq(msgFrame)
		return
	}

	var ack nw.ProtoMessage
	if handler != nil {
		ack, err = handler(conn, msgFrame)
	} else {
		err = errors.New("handler not found")
	}

	if err != nil {
		logger.Error("nw_ccs_session:OnRecvcmdId:%v,:err:%v", msgFrame.CmdId, err)
		this.SendMessage(msgFrame.TransId, &pbserver.ErrorAck{})
	} else if ack != nil {
		this.SendMessage(msgFrame.TransId, ack)
	}
}

func (this *GsSession) SendMessage(transId uint32, msg nw.ProtoMessage) error {
	rb, err := pbserver.Marshal(pbserver.GetCmdIdFromType(msg), transId, false, msg)
	if err != nil {
		return err
	}
	_, err = this.Conn.Write(rb)
	return err
}

func (this *GsSession) CallMessage(requestMsg nw.ProtoMessage, resultMsg nw.ProtoMessage) error {

	return this.rpc.Call(this.rpc.NewContext(resultMsg, nil), requestMsg)
}

func (this *GsSession) OnHandShakeReq(msgFrame *pbserver.MessageFrame) {
	req := msgFrame.Body.(*pbserver.HandShakeReq)
	this.GsNo = int(req.ShakeNo)

	this.SendMessage(0, &pbserver.HandShakeAck{})
	gsManager.AddSession(uint32(this.GsNo), this)
	logger.Info("gs handshake ok. gsNo=%v curSessionCount=%v", req.ShakeNo, gsManager.GetSessionCount())
}
