package managers

import (
	"cqserver/gamelibs/gamedb"
	"errors"
	"runtime/debug"

	"time"

	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/rpc"
	"cqserver/protobuf/pbserver"
)

type GSSession struct {
	Conn  nw.Conn
	GSSeq int //gs端传过来的 serverId
	rpc   *rpc.RpcWrapper
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
	this.rpc = rpc.NewRpcWrapper(conn, 0, pbserver.RpcMarshaler, pbserver.RpcUnmarshaler)
}

func (this *GSSession) OnClose(conn nw.Conn) {
	startAt := time.Now()
	logger.Info("GSSession:OnClose start GSSeq=%v curSessionCount=%v costTime:%v", this.GSSeq, m.gsManager.GetSessionCount())
	defer func() {
		costTime := time.Now().Sub(startAt)
		logger.Info("GSSession:OnClose stop GSSeq=%v curSessionCount=%v costTime:%v", this.GSSeq, m.gsManager.GetSessionCount(), costTime)
	}()

	if this == m.gsManager.GetSession(this.GetId()) {
		m.gsManager.RemoveSession(uint32(this.GSSeq))
	} else {
		logger.Error("GSSession:OnClose error GSSeq=%v curSessionCount=%v,this:%v,gsManager:%v", this.GSSeq, m.gsManager.GetSessionCount(), this, m.gsManager.GetSession(this.GetId()))
	}
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

	switch msgFrame.CmdId {
	case pbserver.CmdHandShakeReqId:
		this.OnFSHandShakeReq(msgFrame)

	case pbserver.CmdGSMessageToFSId:
		this.GsMessageToFS(data, msgFrame)
	default:
		logger.Warn("nw_gs_sesson:OnRecv not implement:cmdId:%d", msgFrame.CmdId)
	}
}

func (this *GSSession) GsMessageToFS(data []byte, msgFrame *pbserver.MessageFrame) error {
	msg := msgFrame.Body.(*pbserver.GSMessageToFS)

	if msgFrame.RpcFlag == 1 { //如果是rpc请求,则也发起rpc
		req := msg
		req.GsTransId = int32(msgFrame.TransId)
		resp := &pbserver.FSMessageToGS{}
		err := m.fsManager.RouteGsMessageCallToFs(int(msg.CrossServerId), req, resp)
		if err != nil {
			logger.Error("转发rpc消息[%v]到fs异常：%v", req, err)
			return err
		} else {
			if int(resp.ServerId) != this.GSSeq {
				logger.Error("接收到异常返回消息，返回消息服务器Id:%v，当前sessionId:%v", resp.ServerId, this.GSSeq)
				return gamedb.ERRPARAM.CloneWithMsg("return msg serverId err")
			}
			data := resp.GetMsg()
			this.Conn.Write(data)
			return nil
		}

	} else {
		return m.fsManager.RouteGsMessageToFs(int(msg.CrossServerId), data)
	}
}

func (this *GSSession) SendMessage(transId uint32, msg nw.ProtoMessage) {
	rb, err := pbserver.Marshal(pbserver.GetCmdIdFromType(msg), transId, false, msg)
	if err != nil {
		return
	}

	this.Conn.Write(rb)
}

func (this *GSSession) OnFSHandShakeReq(msgFrame *pbserver.MessageFrame) {
	req := msgFrame.Body.(*pbserver.HandShakeReq)
	this.GSSeq = int(req.ShakeNo)
	m.gsManager.AddSession(uint32(this.GSSeq), this)
	ack := &pbserver.HandShakeReq{}
	this.SendMessage(msgFrame.TransId, ack)
	logger.Info("gs handshake ok. gsseq:%v curSessionCount=%v", this.GSSeq, m.gsManager.GetSessionCount())

}

func (this *GSSession) Call(cmdId uint16, req nw.ProtoMessage, resp nw.ProtoMessage) error {
	if this.rpc == nil || this.Conn == nil {
		if this.rpc == nil {
			logger.Error("rpcs nil")
		} else {
			logger.Error("conn nil")
		}
		return errors.New("connection closed7")
	}
	return this.rpc.Call(this.rpc.NewContext(resp, nil), req)
}
