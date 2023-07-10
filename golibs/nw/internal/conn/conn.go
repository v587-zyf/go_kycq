package conn

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"cqserver/golibs/nw"
	//"cqserver/golibs/nw/internal/hub"
	"cqserver/golibs/nw/netstat"
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

type IOPumper interface {
	writePump(conn *Conn)
	readPump(conn *Conn)
}

type MessageChan interface {
	GetInChan() chan<- []byte
	GetOutChan() <-chan []byte
	Len() int
	Size() int
}

type Conn struct {
	net.Conn
	msgChan    MessageChan
	inChan     chan<- []byte
	context    *nw.Context
	ioPumper   IOPumper
	wg         sync.WaitGroup
	Session    nw.Session
	ConnTime   time.Time
	stat       *netstat.NetStat
	done       chan struct{}
	closeOnce  sync.Once
	remoteAddr *NetAddr
}

var errConnClosed = errors.New("connection already closed")

func newConn(c net.Conn, context *nw.Context, msgChan MessageChan, done chan struct{}) *Conn {
	conn := &Conn{
		Conn:    c,
		msgChan: msgChan,
		inChan:  msgChan.GetInChan(),
		context: context,
		ConnTime: time.Now(),
		done:     done,
	}

	if context.SessionCreator != nil {
		conn.Session = context.SessionCreator(conn)
	}
	if context.EnableStatistics {
		conn.stat = netstat.NewNetStat()
		conn.stat.Start()
	}
	return conn
}

func (this *Conn) GetSession() nw.Session {
	return this.Session
}

func (this *Conn) GetConnTime() time.Time {
	return this.ConnTime
}

func (this *Conn) ServeIO() {
	this.wg.Add(2)
	go func() {
		this.ioPumper.writePump(this)
		this.wg.Done()
	}()

	go func() {
		this.Session.OnOpen(this)
		this.ioPumper.readPump(this)
		this.Session.OnClose(this)
		if this.stat != nil {
			this.stat.Stop()
		}
		this.wg.Done()
	}()
}

func (this *Conn) Write(data []byte) (int, error) {
	length, size := this.msgChan.Len(), this.msgChan.Size()
	if length >= size {
		fmt.Println(time.Now(), "nw/internal/conn/conn.go Write msgChan.Len()>msgChan.Size()", length, size)
	}
	select {
	case this.inChan <- data:
		return len(data), nil
	case <-this.done:
		return 0, errConnClosed
	}
}

func (this *Conn) SetCloseReason(reason string){

}

func (this *Conn) Close() error {
	var err error
	this.closeOnce.Do(func() {
		close(this.done)
		err = this.Conn.Close()
	})
	return err
}

func (this *Conn) Wait() {
	this.wg.Wait()
}

func (this *Conn) SetRemoteAddr(ip string) {
	this.remoteAddr = &NetAddr{ip: ip}
}

func (this *Conn) RemoteAddr() net.Addr {
	return this.remoteAddr
}
