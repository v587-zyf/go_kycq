package gater

import (
	"time"

	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/tcpserver"
	"cqserver/protobuf/pbgt"
)

type ServerStatus struct {
	OnlineNum int
}

type ExtraContext struct {
	ReadBufferSize      int
	UseNoneBlockingChan bool // use none blocking chan
	ChanSize            int  // chan size for bufferring
	SessionFinder       func(id uint32) nw.Session
	StatusGetter        func() *ServerStatus
}

type Server struct {
	*nw.DefaultSessionManager
	ClientContext      *nw.Context   // 创建clientsession客户端连接需要的Context
	ClientExtraContext *ExtraContext // ClientContext中的Extra字段å
	gateServer         nw.Server
	Stopping           bool
	PingTime           int
	acceptConnect      bool
}

const (
	ReadBufferSize  = 4 * 1024 * 1024
	DefaultChanSize = 10000
)

func NewServer(context *nw.Context) *Server {
	if context.Extra == nil {
		panic("no Extra specified when starting a gater")
	}
	extra := context.Extra.(*ExtraContext)
	return &Server{
		DefaultSessionManager: nw.NewDefaultSessionManager(true),
		ClientContext:         context,
		ClientExtraContext:    extra,
	}
}

func (this *Server) SetAcceptConnect(acceptConnect bool) {
	this.acceptConnect = acceptConnect
}

func (this *Server) Start(addr string) error {
	// 此处需重新创建一个Context，用于与Gate的通信
	readBufferSize := this.ClientExtraContext.ReadBufferSize
	if readBufferSize <= 0 {
		readBufferSize = ReadBufferSize
	}
	chanSize := this.ClientExtraContext.ChanSize
	if chanSize <= 0 {
		chanSize = DefaultChanSize
	}
	context := &nw.Context{
		SessionCreator:      func(conn nw.Conn) nw.Session { return NewGateSession(conn, this) },
		Splitter:            pbgt.Split,
		ReadBufferSize:      readBufferSize,
		UseNoneBlockingChan: this.ClientExtraContext.UseNoneBlockingChan,
		ChanSize:            chanSize,
		IdleTimeAfterOpen:   time.Second * 10, // 必需，由于需要等待HandShake才能分配SessionId，故使用conn.Activate来解决Hub中Id获取问题
	}
	server := tcpserver.NewServer(context)
	err := server.Start(addr)
	if err != nil {
		return err
	}
	this.gateServer = server
	return nil
}

func (this *Server) Stop() {
	this.Stopping = true
	if this.gateServer != nil {
		this.gateServer.Stop()
	}
}

func (this *Server) Broadcast(sessionIds []uint32, rb []byte) {
	if this.gateServer == nil {
		logger.Error("Broadcast error this gate server nil")
		return
	}
	if this.Stopping {
		return
	}
	this.gateServer.Broadcast(sessionIds, rb)
}
