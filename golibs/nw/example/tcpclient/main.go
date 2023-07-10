package main

import (
	"time"

	"cqserver/golibs/nw"
	"cqserver/golibs/nw/tcpclient"
	"cqserver/golibs/util"
)

type Session struct {
	Conn nw.Conn
}

func NewSession(conn nw.Conn) nw.Session {
	return &Session{
		Conn: conn,
	}
}

func (this *Session) GetId() uint32 {
	return 0
}

func (this *Session) GetConn() nw.Conn {
	return this.Conn
}

func (this *Session) OnOpen(conn nw.Conn) {

}

func (this *Session) OnClose(conn nw.Conn) {

}

func (this *Session) OnRecv(conn nw.Conn, data []byte) {

}

func main() {
	context := &nw.Context{
		SessionCreator: NewSession,
		Splitter:       Split,
	}

	ch := tcpclient.DialEx("127.0.0.1:7008", context, 3*time.Second)
	util.WaitForTerminate()
	close(ch)
}

func Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	return len(data), data, nil
}
