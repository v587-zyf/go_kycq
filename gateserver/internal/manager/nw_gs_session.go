package manager

import (
	"cqserver/gateserver/conf"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pbgt"
	"errors"
	"runtime/debug"
)

type GSSession struct {
	Conn nw.Conn
}

func NewGSSession() *GSSession {
	s := &GSSession{
	}
	return s
}

func (this *GSSession) GetId() uint32 {
	return 0
}

func (this *GSSession) GetConn() nw.Conn {
	return this.Conn
}

func (this *GSSession) OnOpen(conn nw.Conn) {
	msg := &pbgt.HandShakeReq{}
	msg.GateSeq = int32(conf.Conf.ServerId)
	rb, err := pbgt.Marshal(pbgt.CmdHandShakeReqId, 0, 0, msg)
	if err != nil {
		logger.Error("marshal HandShakeReq error: %s", err.Error())
		this.Conn.Close()
		return
	}
	logger.Info("gate game建立链接，发送握手协议")
	this.Conn.Write(rb)
}

func (this *GSSession) OnClose(conn nw.Conn) {
	this.Conn = nil
	m.ClientManager.GsClose()
}

func (this *GSSession) OnRecv(conn nw.Conn, data []byte) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Gs session panic: %v, %s", err, debug.Stack())
		}
	}()

	// 注意: 此处对于中转的消息是直接引用的网络缓冲区的slice，如果把解码后的数据发送给其他goroutine处理，需要注意缓冲区覆盖问题
	msgFrame, err := pbgt.Unmarshal(data)
	if err != nil {
		logger.Info("unmarshal gs data error: %s", err.Error())
		return
	}

	if msgFrame.CmdId == pbgt.CmdHandShakeAckId{
		logger.Info("收到game发送来的握手协议，连接建立完成")
		return
	}

	handler := pbgt.GetHandler(msgFrame.CmdId)
	if handler != nil {
		handler(conn, msgFrame)
	} else {
		logger.Info("unhandled cmdId from gs: %d", msgFrame.CmdId)
	}
}


func (this *GSSession) IsConnected() bool {
	return this.Conn != nil
}

func (this *GSSession) Write(data []byte) error{
	conn := this.GetConn()
	if conn == nil {
		return errors.New("gs session conn unknown")
	}
	_, err := conn.Write(data)
	return err
}
