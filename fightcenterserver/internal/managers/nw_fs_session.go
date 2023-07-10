package managers

import (
	"runtime/debug"

	"errors"

	"cqserver/fightcenterserver/internal/base"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/rpc"
	"cqserver/protobuf/pbgt"
	"cqserver/protobuf/pbserver"
)

type FSSession struct {
	IsGs bool //是否是gs-fs连接
	Conn nw.Conn
	FSNo int //fs编号 1-本地战斗服，2-跨服战斗服  动态战斗服为:crossserverId
	Host string
	Port int
	rpc  *rpc.RpcWrapper
	linkServers map[int]bool		//服务器id=>bool
}

func NewFSSession(fsNo int, host string, port int, IsGs bool) *FSSession {
	s := &FSSession{
		IsGs: IsGs,
		FSNo: fsNo,
		Host: host,
		Port: port,
		linkServers: make(map[int]bool),
	}
	return s
}

func (this *FSSession) GetId() uint32 {
	return uint32(this.FSNo)
}

func (this *FSSession) GetConn() nw.Conn {
	return this.Conn
}

func (this *FSSession) SetLinkServers(linkServers map[int]bool){
	this.linkServers = linkServers
}

func (this *FSSession) OnOpen(conn nw.Conn) {

	logger.Info("FS%d OnOpen connected, isgs:%v", this.FSNo, this.IsGs)

	if this.IsGs {
		msg := &pbserver.HandShakeReq{}
		msg.ShakeNo = int32(this.FSNo)
		rb, err := pbserver.Marshal(pbserver.CmdHandShakeReqId, 0, false, msg)
		if err != nil {
			logger.Error("marshal HandShakeReq error: %s", err.Error())
			this.Conn.Close()
			return
		}
		this.Conn.Write(rb)
		this.rpc = rpc.NewRpcWrapper(conn, 0, pbserver.RpcMarshaler, pbserver.RpcUnmarshaler)

	} else {
		msg := &pbgt.HandShakeReq{}
		msg.GateSeq = int32(this.FSNo)
		rb, err := pbgt.Marshal(pbgt.CmdHandShakeReqId, 0, 0, msg)
		if err != nil {
			logger.Error("marshal HandShakeReq error: %s", err.Error())
			this.Conn.Close()
			return
		}
		logger.Info("start FS handshake gateSId=%v", msg.GateSeq)
		this.Conn.Write(rb)

	}
}

func (this *FSSession) OnClose(conn nw.Conn) {
	logger.Info("FS%d disconnected", this.FSNo)
}

//fs->fsc->gate
func (this *FSSession) OnRecv(conn nw.Conn, data []byte) {

	var msgId int

	defer func() {
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic nw_gate recv msg[%v]:%v,%s", msgId, r, stackBytes)
		}
	}()

	if this.IsGs {
		msg, err := this.rpc.HookRecv(data)
		if err != nil {
			logger.Error("unmarshal gs OnRecv unmarshal error: %s", err.Error())
			return
		} else if msg == nil {
			return
		}
		msgFrame := msg.(*pbserver.MessageFrame)
		msgId = int(msgFrame.CmdId)
		switch msgFrame.CmdId {
		case pbserver.CmdFSMessageToGSId:
			this.FSMessageToGS(msgFrame.Body.(*pbserver.FSMessageToGS))
		case pbserver.CmdHandShakeAckId:
			logger.Info("CmdFSHandShakeAckId")
		case pbserver.CmdFSCallMessageToGSId:
			this.FSCallMessageToGS(msgFrame)

		default:
			logger.Error("gs cmd:%v not found", msgId)
		}

	} else {
		msgFrame, err := pbgt.Unmarshal(data)
		if err != nil {
			logger.Error("nw_fs_session gate OnRecv unmarshal err:%v", err)
			return
		}
		msgId = int(msgFrame.CmdId)
		//logger.Info("nw_fs_session gate OnRecv msgId:%v", msgId)
		switch msgFrame.CmdId {
		case pbgt.CmdFSMessageToGateId:
			this.FSMessageToGate(msgFrame)
		case pbgt.CmdHandShakeAckId:
			logger.Info("CmdHandShakeAckId")
		default:
			logger.Error("gate cmd:%v not found", msgId)
		}
	}

}

func (this *FSSession) FSCallMessageToGS(msgFrame *pbserver.MessageFrame) error {

	msg := msgFrame.Body.(*pbserver.FSCallMessageToGS)

	gs := m.gsManager.GetGSSession(uint32(msg.ServerId))
	if gs != nil {
		if msgFrame.RpcFlag == 1 { //如果是rpc请求,则也发起rpc
			req := msg
			req.FsTransId = int32(msgFrame.TransId)

			resp := &pbserver.GSCallMessageToFS{}
			err := gs.Call(msgFrame.CmdId, req, resp) //这里重新生成transId
			if err != nil {
				logger.Error("gs.RouteGsMsgCallFs err:%v", err)
				return err
			}

			data := resp.GetMsg()
			_, err = this.GetConn().Write(data)

			return err
		}
	}
	return base.ErrNoFsServerAssigned
}

func (this *FSSession) FSMessageToGS(msg *pbserver.FSMessageToGS) error {

	if msg.ServerId == 0 {
		sessions := m.gsManager.GetAllSessionId()
		if len(sessions) == 0 {
			return base.ErrNoFsServerAssigned
		}

		for _, v := range sessions {
			ss := m.gsManager.GetGSSession(v)
			if ss != nil && this.linkServers[int(ss.GetId())]{

				data := msg.GetMsg()
				err := this.WriteToGS(ss, data)
				if err != nil {
					logger.Error("跨服广播消息：%v，异常：%v",msg.Msg,err)
				}
			}
		}

	}else if msg.ServerId == -1 {
		sessions := m.gsManager.GetAllSessionId()
		if len(sessions) == 0 {
			return base.ErrNoFsServerAssigned
		}

		for _, v := range sessions {
			ss := m.gsManager.GetGSSession(v)
			if ss != nil {

				data := msg.GetMsg()
				err := this.WriteToGS(ss, data)
				if err != nil {
					logger.Error("跨服广播消息：%v，异常：%v",msg.Msg,err)
				}
			}
		}
	} else {
		ss := m.gsManager.GetGSSession(uint32(msg.ServerId))
		if ss != nil {

			data := msg.GetMsg()
			return this.WriteToGS(ss, data)
		}
		return base.ErrNoFsServerAssigned
	}
	return nil
}

func (this *FSSession) WriteToGS(ss *GSSession, data []byte) error {
	if ss == nil {
		return base.ErrNoFsServerAssigned
	}
	conn := ss.GetConn()
	if conn == nil {
		return base.ErrServerNotConnected
	}
	//logger.Info("WriteToGS  GSSeq:%v", ss.GSSeq)
	_, err := conn.Write(data)
	if err != nil {
		logger.Error("WriteToGS conn.Write err:%v", err)
	}
	return err
}

func (this *FSSession) FSMessageToGate(msgFrame *pbgt.MessageFrame) error {
	msg := msgFrame.Body.(*pbgt.FSMessageToGate)

	ss := m.gateManager.GetGateSession(uint32(msg.ServerId))
	if ss != nil {
		data := msg.GetMsg()
		return this.WriteToGate(ss, data)
	}
	return base.ErrNoFsServerAssigned
}

func (this *FSSession) WriteToGate(ss *GateSession, data []byte) error {
	if ss == nil {
		return base.ErrNoFsServerAssigned
	}
	conn := ss.GetConn()
	if conn == nil {
		return base.ErrServerNotConnected
	}
	_, err := conn.Write(data)
	if err != nil {
		logger.Info("WriteToGate  GateSeq:%v,err:%v", ss.GateSeq, err)
	}
	return err
}

func (this *FSSession) SendMessage(transId uint32, msg nw.ProtoMessage) error {
	rb, err := pbserver.Marshal(pbserver.GetCmdIdFromType(msg), transId, false, msg)
	if err != nil {
		return err
	}
	if this.Conn == nil {
		return nil
	}
	_, err = this.Conn.Write(rb)
	return err
}

func (this *FSSession) IsConnected() bool {
	return this.Conn != nil
}

func (this *FSSession) Close() {
	if this.Conn != nil {
		this.Conn.Close()
		if this.Conn != nil {
			this.Conn.Wait()
		}
	}
}

//-------------------------------------------------

func (this *FSSession) RoutoGateMsgTofs(data []byte) error {
	conn := this.GetConn()
	if conn == nil {
		return base.ErrServerNotConnected
	}
	_, err := conn.Write(data)
	if err != nil {
		logger.Error("WriteToFS isgs:%v, fsno:%v,err:%v", this.IsGs, this.FSNo, err)
	}
	return err
}

func (this *FSSession) RoutoGsMsgTofs(data []byte) error {
	conn := this.GetConn()
	if conn == nil {
		return base.ErrServerNotConnected
	}
	_, err := conn.Write(data)
	if err != nil {
		logger.Error("WriteToFS isgs:%v, fsno:%v,err:%v", this.IsGs, this.FSNo, err)
	}
	return err
}

func (this *FSSession) RouteGsMsgCallFs(req nw.ProtoMessage, resp nw.ProtoMessage) error {
	if this.rpc == nil || this.Conn == nil {
		if this.rpc == nil {
			logger.Error("rpcs nil")
		} else {
			logger.Error("conn nil")
		}
		return errors.New("FSSession connection closed7")
	}
	return this.rpc.Call(this.rpc.NewContext(resp, nil), req)
}
