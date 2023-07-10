package gater

import (
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pbgt"
	"github.com/gogo/protobuf/proto"
	"runtime/debug"
	"sync"
)

type GateSession struct {
	Id             uint32
	Conn           nw.Conn
	Server         *Server
	clientSessions *nw.DefaultSessionManager // 非线程安全的管理器，只能在网络线程中访问
}

func NewGateSession(conn nw.Conn, server *Server) nw.Session {
	s := &GateSession{
		clientSessions: nw.NewDefaultSessionManager(true),
		Conn:           conn,
		Server:         server,
	}
	return s
}

func (this *GateSession) GetId() uint32 {
	return this.Id
}

func (this *GateSession) GetConn() nw.Conn {
	return this.Conn
}

func (this *GateSession) OnOpen(conn nw.Conn) {
	logger.Info("gate connected")
}

func (this *GateSession) OnClose(conn nw.Conn) {
	this.Server.RemoveSession(this.Id)
	if !this.Server.Stopping {
		logger.Warn("gate%d disconnected", this.Id)
	}
	// 关闭本GateSession管理的所有ClientSession
	conns := make(map[nw.Conn]bool)
	this.clientSessions.Range(func(id uint32, session nw.Session) bool {
		conn := session.GetConn()
		conns[conn] = true
		return false
	})
	if len(conns) > 0 {
		var wg sync.WaitGroup
		wg.Add(len(conns))
		pool := getCloserPool()
		for conn := range conns {
			pool.Add(getCloserWorker(conn, &wg))
		}
		wg.Wait()
	}
}

func (this *GateSession) OnRecv(conn nw.Conn, data []byte) {
	defer func() {
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic gate session:%v,%s", r, stackBytes)
		}
	}()

	msgFrame, err := pbgt.Unmarshal(data)
	if err != nil {
		logger.Error("unmarshal gate message error: %s", err.Error())
		conn.Close()
		return
	}

	handler := pbgt.GetHandler(msgFrame.CmdId)
	if handler != nil {
		handler(this.Conn, msgFrame)
	} else {
		logger.Warn("unknown gate cmdId: %d", msgFrame.CmdId)
	}
}

func (this *GateSession) SendMessage(cmdId uint16, sessionId uint32, msg proto.Message) error {
	rb, err := pbgt.Marshal(cmdId, sessionId, 0, msg)
	if err != nil {
		logger.Error("gate SendMessage marshal err:%v,cmdId:%v,sessionId:%v",err,cmdId,sessionId)
		return err
	}
	_, err = this.Conn.Write(rb)
	if err != nil {
		logger.Error("gate SendMessage send err:%v,cmdId:%v,sessionId:%v",err,cmdId,sessionId)
	}
	return err
}
