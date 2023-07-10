package tcpclient

import (
	"net"
	"runtime"
	"time"

	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/internal/conn"
)

var log = logger.Get("default", true)

func Dial(addr string, context *nw.Context) (net.Conn, error) {
	c, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		return nil, err
	}
	nwConn := conn.NewTcpConn(c, context)
	nwConn.ServeIO()
	return nwConn, nil
}

func DialEx(addr string, context *nw.Context, retryWait time.Duration) chan struct{} {
	doneChan := make(chan struct{})
	go func() {
		retryChan := make(chan bool, 1)
		retryChan <- true
		needWait := false
		for {
			select {
			case <-doneChan:
				return
			case <-retryChan:
				if needWait && retryWait > 0 {
					select {
					case <-time.NewTimer(retryWait).C:
					case <-doneChan:
						return
					}
				} else {
					needWait = true
				}
				c, err := net.DialTimeout("tcp", addr, 5*time.Second)
				if err != nil {
					//logger.Debug("connect error: %s", err.Error())
					//fmt.Printf("connect error: %s\n", err.Error())
					retryChan <- true
					continue
				}
				nwConn := conn.NewTcpConn(c, context)
				nwConn.ServeIO()
				nwConn.Wait() // this will block the for loop, the application layer need to close the conn to let it go along
				runtime.Gosched()
				retryChan <- true
			}
		}
	}()
	return doneChan
}
