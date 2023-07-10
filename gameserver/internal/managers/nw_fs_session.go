package managers

import (
	"cqserver/gamelibs/errex"
	"cqserver/gameserver/internal/base"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/rpc"
	"cqserver/protobuf/pbserver"
	"errors"
	"fmt"
	"runtime/debug"
)

type FSSession struct {
	Conn nw.Conn
	FSNo int //fs编号
	Addr string
	rpc  *rpc.RpcWrapper
}

func NewFSSession(fsNo int, host string, port int) *FSSession {
	s := &FSSession{
		FSNo: fsNo,
		Addr: fmt.Sprintf("%s:%d", host, port),
	}
	return s
}

func (this *FSSession) GetId() uint32 {
	return uint32(this.FSNo)
}

func (this *FSSession) GetConn() nw.Conn {
	return this.Conn
}

func (this *FSSession) OnOpen(conn nw.Conn) {
	msg := &pbserver.HandShakeReq{}
	msg.ShakeNo = int32(base.Conf.ServerId)
	rb, err := pbserver.Marshal(pbserver.CmdHandShakeReqId, 0, false, msg)
	if err != nil {
		logger.Error("marshal HandShakeReq error: %s", err.Error())
		this.Conn.Close()
		return
	}
	this.Conn.Write(rb)
	this.rpc = rpc.NewRpcWrapper(conn, 0, pbserver.RpcMarshaler, pbserver.RpcUnmarshaler)
}

func (this *FSSession) OnClose(conn nw.Conn) {
	this.Conn = nil
	logger.Info("FS%d disconnected", this.FSNo)
}

func (this *FSSession) OnRecv(conn nw.Conn, data []byte) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("nw_fs OnRecv panic: %v, %s", err, debug.Stack())
		}
	}()

	msgFrame, err := this.rpc.HookRecv(data)
	if err != nil {
		logger.Info("unmarshal gs data error: %s", err.Error())
		return
	} else if msgFrame == nil {
		return
	}
	msg := msgFrame.(*pbserver.MessageFrame)

	//reqMsg := msgFrame.(*pbserver.MessageFrame)
	if msg.CmdId == pbserver.CmdHandShakeAckId {
		logger.Info("game fight建立链接，收到握手成功")
		return
	}

	handler := pbserver.GetHandler(msg.CmdId)
	if handler != nil {
		ack, err := handler(conn, msg)
		if err != nil {
			ack = errex.BuildServerErrorAck(err)
		}
		if ack != nil {
			this.SendMessage(0, msg.TransId, ack)
		}
	} else {
		logger.Info("unhandled cmdId from gs: %d", msg.CmdId)
	}
}

func (this *FSSession) IsConnected() bool {
	return this.Conn != nil
}

func (this *FSSession) SendMessage(crossFightServerId int, transId uint32, msg nw.ProtoMessage) error {
	cmdId := pbserver.GetCmdIdFromType(msg)
	rb, err := pbserver.Marshal(cmdId, transId, false, msg)
	if err != nil {
		return err
	}
	if this.Conn == nil {
		return nil
	}
	//logger.Info("=====FSSession==== SendMessage GSMessageToFS cmdId:%v; ", cmdId)
	if crossFightServerId > 0 { //动态战斗跨服
		m := &pbserver.GSMessageToFS{
			Msg:           rb,
			CrossServerId: int32(crossFightServerId),
			ServerId:      int32(base.Conf.ServerId),
		}
		rb, err = pbserver.Marshal(pbserver.CmdGSMessageToFSId, transId, false, m)
		if err != nil {
			return err
		}
	}

	_, err = this.Conn.Write(rb)
	return err
}

func (this *FSSession) Call(crossFightServerId int, req nw.ProtoMessage, resp nw.ProtoMessage) error {
	if this.rpc == nil || this.Conn == nil {
		if this.rpc == nil {
			logger.Error("rpcs nil")
		} else {
			logger.Error("conn nil")
		}
		return errors.New("connection closed10")
	}

	if crossFightServerId > 0 {

		cmdId := pbserver.GetCmdIdFromType(req)
		rb, err := pbserver.Marshal(cmdId, 0, false, req)
		if err != nil {
			return err
		}
		logger.Info("=====FSSession==== Call GSMessageToFS cmdId:%v;CrossServerId:%v; ", cmdId, crossFightServerId)
		m := &pbserver.GSMessageToFS{
			Msg:           rb,
			CrossServerId: int32(crossFightServerId),
			ServerId:      int32(base.Conf.ServerId),
		}
		req = m
	}

	return this.rpc.Call(this.rpc.NewContext(resp, nil), req)
}

func (this *FSSession) Close() {
	if this.Conn != nil {
		this.Conn.Close()
		this.Conn.Wait()
	} else {
		logger.Error("close fs session:%v.con nil", this.GetId())
	}
}
