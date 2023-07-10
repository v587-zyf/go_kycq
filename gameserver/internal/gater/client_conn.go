package gater

import (
	"errors"
	"net"
	"sync"
	"time"

	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pbgt"
)

type NetAddr struct {
	ip string
}

func (this *NetAddr) Network() string {
	return "tcp"
}

func (this *NetAddr) String() string {
	return this.ip
}

// ClientConn fake Conn for client, use the underline gate Conn to route message
type ClientConn struct {
	remoteAddr          *NetAddr
	context             *nw.Context
	GateConn            net.Conn
	GateSession         *GateSession
	gateClientSessionId uint32
	Session             nw.Session
	ConnTime            time.Time
	msgOut              chan []byte
	msgIn               chan []byte
	done                chan struct{}
	closeOnce           sync.Once
	wg                  sync.WaitGroup
	closeReason         string
}

const (
	msgInSize  = 100
	msgOutSize = 2000
)

func NewClientConn(gateConn nw.Conn, gateClientSessionId uint32, clientIP string, context *nw.Context) *ClientConn {
	conn := &ClientConn{
		remoteAddr:  &NetAddr{ip: clientIP},
		context:     context,
		GateConn:    gateConn,
		GateSession: gateConn.GetSession().(*GateSession),
		msgOut:      make(chan []byte, msgOutSize),
		msgIn:       make(chan []byte, msgInSize),
		done:        make(chan struct{}),
		ConnTime:    time.Now(),
	}
	/***************************/
	/***下面两行代码顺序不可更改****/
	/***************************/
	conn.gateClientSessionId = gateClientSessionId
	conn.Session = context.SessionCreator(conn)
	return conn
}

func (this *ClientConn) GetSession() nw.Session {
	return this.Session
}

func (this *ClientConn) GetConnTime() time.Time {
	return this.ConnTime
}

func (this *ClientConn) GetGateClientSessionId() uint32 {
	return this.gateClientSessionId
}

func (this *ClientConn) Read(data []byte) (n int, err error) {
	return 0, errors.New("not implemented")
}

func (this *ClientConn) Write(data []byte) (n int, err error) {
	length := len(this.msgOut)
	if length >= msgOutSize {
		logger.Info("Write len(this.msgOut) >= msgOutSize ", length, msgOutSize)
	}

	select {
	case this.msgOut <- data:
		return len(data), nil
	case <-this.done:
		return 0, errors.New("write to a closed connection")
	}
}

func (this *ClientConn) putMessage(data []byte) {
	length := len(this.msgIn)
	if length >= msgInSize {
		logger.Info("putMessage len(this.msgIn) >= msgInSize ", length, msgInSize)
		//TODO 消息chan已满，临时处理，直接先丢掉，需查询问题
		return
	}
	select {
	case this.msgIn <- data:
	case <-this.done:
	}
}

func (this *ClientConn) SetCloseReason(reason string) {
	this.closeReason = reason
}

func (this *ClientConn) Close() error {
	this.closeOnce.Do(func() {
		close(this.done)
		ntf := &pbgt.UserQuitNtf{Reason: this.closeReason}
		this.GateSession.SendMessage(pbgt.CmdUserQuitNtfId, this.Session.GetId(), ntf)
	})
	return nil
}

func (this *ClientConn) LocalAddr() net.Addr {
	return this.GateConn.LocalAddr()
}

func (this *ClientConn) RemoteAddr() net.Addr {
	return this.remoteAddr
}

func (this *ClientConn) SetDeadline(t time.Time) error {
	return errors.New("not implemented")
}

func (this *ClientConn) SetReadDeadline(t time.Time) error {
	return errors.New("not implemented")
}

func (this *ClientConn) SetWriteDeadline(t time.Time) error {
	return errors.New("not implemented")
}

func (this *ClientConn) Activate() {

}

func (this *ClientConn) Wait() {
	this.wg.Wait()
}

func (this *ClientConn) ServeIO() {
	this.wg.Add(2)
	go func() {
		this.writePump()
		this.wg.Done()
	}()

	// AddSession需要在goroutine外层调用，防止gate中转消息过快导致gs找不到对应的session
	this.GateSession.clientSessions.AddSession(this.gateClientSessionId, this.Session)
	// logger.Debug("enter, gateSessionId: %d, gsSessionId: %d, total: %d", this.gateClientSessionId, this.gsClientSessionId, this.GateSession.clientSessions.GetSessionCount())
	go func() {
		this.Session.OnOpen(this)
		this.readPump()
		this.Session.OnClose(this)
		this.GateSession.clientSessions.RemoveSession(this.gateClientSessionId)
		// logger.Debug("quit, gateSessionId: %d, gsSessionId: %d, total: %d", this.gateClientSessionId, this.gsClientSessionId, this.GateSession.clientSessions.GetSessionCount())
		this.wg.Done()
	}()
}

func (this *ClientConn) writePump() {
	for {
		select {
		case message, ok := <-this.msgOut:
			if !ok {
				return
			}
			msg, err := pbgt.Marshal(pbgt.CmdRouteMessageId, this.gateClientSessionId, 0, message)
			if err != nil {
				logger.Error("ClientConn writePump marshal error: %s", err.Error())
				this.Close()
				return
			}
			this.GateConn.Write(msg)
		case <-this.done:
			return
		}
	}
}

func (this *ClientConn) readPump() {
	for {
		select {
		case message, ok := <-this.msgIn:
			if !ok {
				return
			}
			this.Session.OnRecv(this, message)
		case <-this.done:
			return
		}
	}
}
