package manager

import (
	"cqserver/gamelibs/publicCon/constServer"
	"cqserver/gateserver/conf"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pbgt"
	"errors"
)

type FSSession struct {
	Conn nw.Conn
	FSNo int //fs编号 1-本地战斗服， 2-动态战斗消息中心
	Host string
	Port int
}

func NewFSSession(fsNo int, host string, port int) *FSSession {
	s := &FSSession{
		FSNo: fsNo,
		Host: host,
		Port: port,
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
	msg := &pbgt.HandShakeReq{}
	msg.GateSeq = int32(conf.Conf.ServerId)
	rb, err := pbgt.Marshal(pbgt.CmdHandShakeReqId, 0, 0, msg)
	if err != nil {
		logger.Error("marshal HandShakeReq error: %s", err.Error())
		this.Conn.Close()
		return
	}
	if this.FSNo == constServer.FIGHT_SESSIONID_LOCAL {
		logger.Info("和战斗服建立连接，发送握手协议 gateSId=%v", msg.GateSeq)
	} else {
		logger.Info("和战斗中心建立连接，发送握手协议 gateSId=%v", msg.GateSeq)
	}
	this.Conn.Write(rb)
}

func (this *FSSession) OnClose(conn nw.Conn) {
	m.ClientManager.Range(func(id uint32, session nw.Session) bool {
		//clientSession := session.(*ClientSession)
		//if clientSession.curFightId%10 == uint32(this.FSNo) {
		//	ntf := &pb.KickUserNtf{
		//		Reason: "和服务器断开连接，请重新登录",
		//	}
		//	clientSession.SendMessageToClient(0, ntf)
		//}
		return false
	})
	this.Conn = nil
	logger.Info("FS%d disconnected", this.FSNo)
}

func (this *FSSession) OnRecv(conn nw.Conn, data []byte) {
	msgFrame, err := pbgt.Unmarshal(data)
	if err != nil {
		logger.Info("unmarshal gs data error: %s", err.Error())
		return
	}
	if msgFrame.CmdId == pbgt.CmdHandShakeAckId {
		if this.FSNo == constServer.FIGHT_SESSIONID_LOCAL {
			logger.Info("和战斗服建立连接，收到握手协议:%v", this.FSNo)
		} else {
			logger.Info("和战斗中心建立连接，收到握手协议:%v", this.FSNo)
		}
		return
	}
	handler := pbgt.GetHandler(msgFrame.CmdId)
	if handler != nil {
		handler(conn, msgFrame)
	} else {
		logger.Info("unhandled cmdId from gs: %d", msgFrame.CmdId)
	}
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

func (this *FSSession) WriteToFS(fightServerId int, data []byte) error {

	sendData := data
	//动态跨服战斗 需要通过消息中心转发
	if this.FSNo == constServer.FIGHT_SESSIONID_CENTER {

		m := &pbgt.GateMessageToFS{
			Msg:           data,
			CrossServerId: int32(fightServerId),
			ServerId:      int32(conf.Conf.ServerId),
		}
		var err error
		sendData, err = pbgt.Marshal(pbgt.CmdGateMessageToFSId, 0, 0, m)
		if err != nil {
			logger.Error("RouteToFS CrossServerId:%v,err:%v", fightServerId, err)
			return err
		}
	}

	conn := this.GetConn()
	if conn == nil {
		return errors.New("fight server session error")
	}
	_, err := conn.Write(sendData)
	return err
}
