package net

import (
	"cqserver/fightserver/conf"
	"cqserver/gamelibs/errex"
	"errors"
	"runtime/debug"

	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/rpc"
	"cqserver/protobuf/pbserver"
)

type GSSession struct {
	Conn      nw.Conn
	GSSeq     int
	rpc       *rpc.RpcWrapper
	gsManager *GSManager
}

func NewGSSession(conn nw.Conn) nw.Session {
	s := &GSSession{
		Conn: conn,
	}
	return s
}

func (this *GSSession) GetId() uint32 {
	return uint32(this.GSSeq)
}

func (this *GSSession) GetConn() nw.Conn {
	return this.Conn
}

func (this *GSSession) OnOpen(conn nw.Conn) {
	logger.Info("GSSession:OnOpen")
	this.rpc = rpc.NewRpcWrapper(conn, 0, pbserver.RpcMarshaler, pbserver.RpcUnmarshaler)
}

func (this *GSSession) OnClose(conn nw.Conn) {
	this.OnGSConnClose(int32(this.GSSeq))
	logger.Info("GSSession:OnClose start GSSeq=%v curSessionCount=%v costTime:%v", this.GSSeq)

}

func (this *GSSession) OnRecv(conn nw.Conn, data []byte) {
	var msgId int

	defer func() {
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic fight recv msg[%v]:%v,%s", msgId, r, stackBytes)
		}
	}()

	msg, err := this.rpc.HookRecv(data)
	if err != nil {
		logger.Error("nw_gs_sesson:OnRecv err or msg is nil::err:%v,msg:%v,data:%v", err, msg, data)
		return
	}
	//是个rpc调用，这里msg是空，不处理了．
	if msg == nil {
		return
	}
	msgFrame := msg.(*pbserver.MessageFrame)
	msgId = int(msgFrame.CmdId)

	if msgFrame.CmdId == pbserver.CmdHandShakeReqId {
		this.OnFSHandShakeReq(msgFrame)
		return
	}

	//如果是动态战斗服,则需要先解码消息
	if conf.Conf.IsCorssFightServer() && msgId == pbserver.CmdGSMessageToFSId {

		msgFrame, err = GetFssRouteFromGsMsgFrame(msgFrame)
		if err != nil {
			logger.Error("Recv gs err: %v", err)
			return
		}
		msgId = int(msgFrame.CmdId)
	}

	handler := pbserver.GetHandler(msgFrame.CmdId)
	if handler != nil {
		ack, err := handler(this.Conn, msgFrame)
		if err != nil {
			logger.Error("消息执行错误：msg:%v,%v", *msgFrame, err)
			errMsg := errex.BuildServerErrorAck(err)
			this.ReplyMessage(msgFrame.TransId, errMsg)
			return
		}
		if ack != nil {
			this.ReplyMessage(msgFrame.TransId, ack)
		}
	} else {
		logger.Error("nw_gs_sesson:OnRecv not implement:cmdId:%d", msgFrame.CmdId)
	}
}


//对rpc返回的,使用这个方法
func (this *GSSession) SendMessage(serverId int32, msg nw.ProtoMessage) error {

	rb, err := pbserver.Marshal(pbserver.GetCmdIdFromType(msg), 0, false, msg)
	if err != nil {
		return err
	}

	if conf.Conf.IsCorssFightServer() {
		resp := &pbserver.FSMessageToGS{
			GsTransId: 0,
			ServerId:  serverId,
			Msg:       rb,
		}

		rb, err = pbserver.Marshal(pbserver.GetCmdIdFromType(resp), 0, false, resp)
		if err != nil {
			return err
		}
	}

	_, err = this.Conn.Write(rb)
	return err
}

func (this *GSSession) ReplyMessage(transId uint32, msg nw.ProtoMessage) error {

	var rb []byte
	var err error
	if !conf.Conf.IsCorssFightServer() {
		rb, err = pbserver.Marshal(pbserver.GetCmdIdFromType(msg), transId, false, msg)
	} else {
		rb, err = GetFccRouteToGsMessageBytes(transId, msg)
	}
	//如果是动态服的,则rpc消息需要封装
	if err != nil {
		logger.Error("GSSession SendMessage err:%v", err)
		return err
	}

	_, err = this.Conn.Write(rb)
	return err
}

func (this *GSSession) OnFSHandShakeReq(msgFrame *pbserver.MessageFrame) {
	req := msgFrame.Body.(*pbserver.HandShakeReq)
	if int(req.ShakeNo) != conf.Conf.ServerId {
		logger.Error("连接错误，本服serverId：%v，连接上来的为：%v", conf.Conf.ServerId, req.ShakeNo)
		this.Conn.Close()
		return
	}
	this.GSSeq = int(req.ShakeNo) //动态战斗服的 GsSeq 为crossserverid
	ack := &pbserver.HandShakeAck{}
	rb, err := pbserver.Marshal(pbserver.GetCmdIdFromType(ack), 0, false, ack)
	if err != nil {
		logger.Error("GateSession marshal error: %s", err.Error())
		return
	}
	this.Conn.Write(rb)
	if conf.Conf.IsCorssFightServer() {
		logger.Info("与战斗中心握手成功")
	} else {
		logger.Info("与game握手成功：%v", req.ShakeNo)
	}
	//抛出gs 和 fs 连接成功
	handler := pbserver.GetHandler(pbserver.CmdHandShakeReqId)
	handler(this.Conn, msgFrame)
}

func (this *GSSession) OnGSConnClose(gsSeq int32) {

	handlder := pbserver.GetHandler(pbserver.CmdHandCloseNtfId)
	handlder(this.Conn, &pbserver.MessageFrame{Body: &pbserver.HandCloseNtf{SessionNo: int32(this.GetId())}})
}

func (this *GSSession) Call(serverId uint32, req nw.ProtoMessage, resp nw.ProtoMessage) error {

	if this.rpc == nil || this.Conn == nil {
		if this.rpc == nil {
			logger.Error("rpcs nil")
		} else {
			logger.Error("conn nil")
		}
		return errors.New("connection closed7")
	}

	if conf.Conf.IsCorssFightServer() {

		cmdId := pbserver.GetCmdIdFromType(req)
		rb, err := pbserver.Marshal(cmdId, 0, false, req)
		if err != nil {
			return err
		}

		m := &pbserver.FSCallMessageToGS{
			Msg:           rb,
			CrossServerId: int32(conf.Conf.ServerId),
			ServerId:      int32(serverId),
		}
		req = m
	}
	return this.rpc.Call(this.rpc.NewContext(resp, nil), req)
}
