package wsserver

import (
	"cqserver/golibs/common"
	"net/http"
	"sync"

	"cqserver/golibs/nw"
	"cqserver/golibs/nw/httpserver"
	"cqserver/golibs/nw/internal/conn"
	"github.com/gorilla/websocket"

	"cqserver/golibs/logger"
)

type Server struct {
	Context    *nw.Context
	upgrader   *websocket.Upgrader
	httpServer *httpserver.HttpServer

	conns         map[uint32]nw.Conn
	mu            sync.RWMutex
	acceptConnect bool
}

func (this *Server) SetAcceptConnect(acceptConnect bool) {
	this.acceptConnect = acceptConnect
}

// NewServer create a new websocket server
func NewServer(context *nw.Context) *Server {
	if context == nil || context.SessionCreator == nil {
		panic("wsserver.NewServer: context is nil or context.SessionCreator is nil")
	}
	server := &Server{
		Context: context,
		//hub:     hub.NewHub(context.IdleTimeAfterOpen),
		conns: make(map[uint32]nw.Conn),
	}
	server.upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024 * 2,
		WriteBufferSize: 1024 * 4,
		CheckOrigin:     func(r *http.Request) bool { return true }, // disable check
	}
	if context.ReadBufferSize > 0 {
		server.upgrader.ReadBufferSize = context.ReadBufferSize
	}
	if context.WriteBufferSize > 0 {
		server.upgrader.WriteBufferSize = context.WriteBufferSize
	}
	return server
}

// Start start websocket server, and start default http server if addr is not empty
func (this *Server) Start(addr string) error {
	if len(addr) > 0 {
		httpServer := httpserver.DefaultHttpServer
		err := httpServer.Start(addr)
		if err != nil {
			return err
		}
		httpServer.Handle("/ws", this)
		this.httpServer = httpServer
	}
	return nil
}

// ServeHTTP serve http request
func (this *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if !this.acceptConnect{
		logger.Error("server Startup incomplete ")
		return
	}

	c, err := this.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Info("wsserver.ServeHTTP upgrade error: %s", err)
		return
	}

	conn := conn.NewWsConn(c, this.Context)

	go func() {
		this.AddConn(conn)
		conn.SetRemoteAddr(common.GetIpAddress(r))
		conn.ServeIO()
		conn.Wait()
		this.DeleteConn(conn)

	}()
}

func (this *Server) getIp(r *http.Request) string {
	ip := r.Header.Get("Remote_addr")
	if ip == "" {
		ip = r.RemoteAddr
	}
	logger.Info("get ip:%v", ip)
	return ip
}

func (this *Server) AddConn(conn nw.Conn) {
	this.mu.Lock()
	defer this.mu.Unlock()

	this.conns[conn.GetSession().GetId()] = conn
}

func (this *Server) DeleteConn(conn nw.Conn) {
	this.mu.Lock()
	defer this.mu.Unlock()
	delete(this.conns, conn.GetSession().GetId())
}

// Stop stop websocket server, and the underline default http server
func (this *Server) Stop() {
	if this.httpServer != nil {
		this.httpServer.Stop()
	}
	//this.hub.Stop()

	this.mu.Lock()
	defer this.mu.Unlock()

	this.conns = make(map[uint32]nw.Conn)
}

// Broadcast broadcast data to all active connections
func (this *Server) Broadcast(sessionIds []uint32, data []byte) {
	//this.hub.Broadcast(sessionIds, data)

	this.mu.RLock()
	defer this.mu.RUnlock()

	if len(sessionIds) == 0 {
		for _, conn := range this.conns {
			conn.Write(data)
		}
		return
	}

	for i := 0; i < len(sessionIds); i++ {
		conn := this.conns[sessionIds[i]]
		if conn != nil {
			conn.Write(data)
		}
	}

}

// GetActiveConnNum get count of active connections
// func (this *Server) GetActiveConnNum() int {
// 	return 0 //this.hub.GetActiveConnNum()
// }
