package tcpserver

import (
	"net"
	"runtime"
	"strings"
	"sync"

	"cqserver/golibs/nw"
	"cqserver/golibs/nw/internal/conn"

	"cqserver/golibs/logger"
)

type Server struct {
	Context *nw.Context
	listener net.Listener

	conns map[nw.Conn]bool //这里不能使用sessionid作为key, 因为Conn中关联的sessionid是在连接完成后才赋值
	mu    sync.RWMutex
	acceptConnect bool
}

// NewServer create a new tcp server
func NewServer(context *nw.Context) *Server {
	if context == nil || context.SessionCreator == nil || context.Splitter == nil {
		panic("tcpserver.NewServer: context.SessionCreator is nil or context.Splitter is nil")
	}

	server := &Server{
		Context: context,
		conns: make(map[nw.Conn]bool),
	}
	return server
}

// Start start tcp server
func (this *Server) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	this.listener = l
	go this.serve()
	return nil
}
func (this *Server) SetAcceptConnect(acceptConnect bool) {
	this.acceptConnect = acceptConnect
}

func (this *Server) serve() {
	l := this.listener
	for {
		c, err := l.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				logger.Info("tcpserver: temporary Accept() error, %s", err)
				runtime.Gosched()
				continue
			}
			// theres no direct way to detect this error because it is not exposed
			if !strings.Contains(err.Error(), "use of closed network connection") {
				logger.Info("tcpserver: listener.Accept() error, %s", err)
			}
			break
		}
		conn := conn.NewTcpConn(c, this.Context)
		go func() {
			this.AddConn(conn)
			conn.ServeIO()
			conn.Wait()
			this.DeleteConn(conn)

		}()
	}
}

func (this *Server) AddConn(conn nw.Conn) {
	this.mu.Lock()
	defer this.mu.Unlock()

	this.conns[conn] = true

}

func (this *Server) DeleteConn(conn nw.Conn) {
	this.mu.Lock()
	defer this.mu.Unlock()

	delete(this.conns, conn)

}

// Stop stop tcp server
func (this *Server) Stop() {
	if this.listener != nil {
		this.listener.Close()
	}
	this.mu.Lock()
	defer this.mu.Unlock()
	for c ,_ := range this.conns{
		c.Close()
	}
	for c ,_ := range this.conns{
		c.Wait()
	}
	if len(this.conns) != 0 {
		this.conns = make(map[nw.Conn]bool)
	}
}

// Broadcast broadcast data to all active connections
//给所有链接广播消息

// Broadcast broadcast data to all active connections
func (this *Server) Broadcast(sessionIds []uint32, data []byte) {
	//this.hub.Broadcast(sessionIds, data)

	this.mu.RLock()
	defer this.mu.RUnlock()

	if len(sessionIds) == 0 {
		for conn := range this.conns {
			conn.Write(data)
		}
		return
	}

	ids := make(map[uint32]uint32)
	for i := 0; i < len(sessionIds); i++ {
		ids[sessionIds[i]] = sessionIds[i]
	}
	for conn := range this.conns {
		if ids[conn.GetSession().GetId()] > 0 {
			conn.Write(data)
		}
	}

}

// GetActiveConnNum get count of active connections
// func (this *Server) GetActiveConnNum() int {
// 	//return this.hub.GetActiveConnNum()
// 	return 0
// }
