package managers

import (
	"cqserver/gamelibs/errex"
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/base"
	"errors"

	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/rpc"
	"cqserver/protobuf/pbserver"
)

type CsSession struct {
	Conn nw.Conn
	rpc  *rpc.RpcWrapper
}

func NewCsSession() *CsSession {
	s := &CsSession{}
	return s
}

func (this *CsSession) GetId() uint32 {
	return 1 // 只会创建一个session
}

func (this *CsSession) GetConn() nw.Conn {
	return this.Conn
}

func (this *CsSession) OnOpen(conn nw.Conn) {
	this.rpc = rpc.NewRpcWrapper(conn, 0, pbserver.RpcMarshaler, pbserver.RpcUnmarshaler)
	logger.Info("nw_cs_session:OnOpen: rpcs init.")
	msg := &pbserver.HandShakeReq{}
	msg.ShakeNo = int32(base.Conf.ServerId)
	rb, err := pbserver.Marshal(pbserver.CmdHandShakeReqId, 0, false, msg)
	if err != nil {
		logger.Error("marshal HandShakeReq error: %s", err.Error())
		this.Conn.Close()
		return
	}
	logger.Info("game crosscenter建立链接，发送握手协议")
	this.Conn.Write(rb)
}

func (this *CsSession) OnClose(conn nw.Conn) {
	logger.Info("nw_cs_session:OnClose")
	this.Conn = nil
}

func (this *CsSession) OnRecv(conn nw.Conn, data []byte) {
	msgFrame, err := this.rpc.HookRecv(data)
	if err != nil {
		logger.Error("rpcs.HookRecv error: %s", err.Error())
		return
	} else if msgFrame == nil {
		return
	}
	reqMsg := msgFrame.(*pbserver.MessageFrame)
	if reqMsg.CmdId == pbserver.CmdHandShakeAckId {
		logger.Info("game crosscenter建立链接，收到握手成功")
		return
	}
	var ack nw.ProtoMessage
	handler := pbserver.GetHandler(reqMsg.CmdId)
	if handler != nil {
		ack, err = handler(this.Conn, reqMsg)
	} else {
		logger.Debug("unhandled cmdId from crosscenterserver: %d", reqMsg.CmdId)
	}

	if err != nil {
		ei, ok := err.(*errex.ErrorItem)
		if !ok {
			ei = gamedb.ERRUNKNOW
		}
		logger.Error("消息处理err:%v", err)
		this.SendMessage(reqMsg.TransId, &pbserver.ErrorAck{
			Code:    int32(ei.Code),
			Message: ei.Message,
		})
	} else if ack != nil {
		this.SendMessage(reqMsg.TransId, ack.(nw.ProtoMessage))
	}
}

func (this *CsSession) SendMessage(transId uint32, msg nw.ProtoMessage) error {
	rb, err := pbserver.Marshal(pbserver.GetCmdIdFromType(msg), transId, false, msg)
	if err != nil {
		return err
	}
	if !this.IsConnected() {
		return errors.New("not Connected")
	}
	_, err = this.Conn.Write(rb)
	return err
}

func (this *CsSession) Call(req nw.ProtoMessage, resp nw.ProtoMessage) error {
	if this.rpc == nil || this.Conn == nil {
		if this.rpc == nil {
			logger.Error("rpcs nil")
		} else {
			logger.Error("conn nil")
		}
		return errors.New("connection closed12")
	}
	return this.rpc.Call(this.rpc.NewContext(resp, nil), req)
}

func (this *CsSession) AsnycCall(req nw.ProtoMessage, resp nw.ProtoMessage, callback rpc.AsyncCallBack) error {
	if this.rpc == nil || this.Conn == nil {
		if this.rpc == nil {
			logger.Error("rpcs nil")
		} else {
			logger.Error("conn nil")
		}
		return errors.New("connection closed13")
	}
	return this.rpc.AsyncCall(this.rpc.NewContext(resp, nil), req, callback)
}

func (this *CsSession) IsConnected() bool {
	return this.Conn != nil
}

func (this *CsSession) Close() {
	conn := this.Conn
	if conn == nil {
		return
	}
	conn.Close()
	conn.Wait()
}
