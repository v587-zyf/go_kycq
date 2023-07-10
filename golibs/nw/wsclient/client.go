package wsclient

import (
	"net"

	"cqserver/golibs/nw"
	"cqserver/golibs/nw/internal/conn"
	"github.com/gorilla/websocket"
)

func Dial(addr string, context *nw.Context) (net.Conn, error) {
	c, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		return nil, err
	}
	conn := conn.NewWsConn(c, context)
	conn.ServeIO()
	return conn, nil
}
