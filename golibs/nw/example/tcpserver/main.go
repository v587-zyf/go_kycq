package main

import (
	"bytes"

	"cqserver/golibs/nw"
	"cqserver/golibs/nw/tcpserver"
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
	server := tcpserver.NewServer(context)
	server.Start(":7008")
	util.WaitForTerminate()
	server.Stop()
}

func Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	n := bytes.IndexByte(data, '-')
	if n > 0 {
		return n + 1, data[0:n], nil
	}
	return 0, nil, nil
}
