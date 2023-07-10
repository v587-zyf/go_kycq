package conn

import (
	"bufio"
	"net"
	"time"

	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/internal/common"
)

type TcpPumper struct {
	done                chan struct{}
	readBuffSize        int
	mergedWriteBuffSize int
	disableMergedWrite  bool
}

const DefaultReadBuffSize = 8 * 1024
const DefaultWriteBuffSize = 16 * 1024

func NewTcpConn(c net.Conn, context *nw.Context) *Conn {
	done := make(chan struct{})
	msgChan := NewMessageChan(context.UseNoneBlockingChan, context.ChanSize, done)
	conn := newConn(c, context, msgChan, done)
	readBuffSize := DefaultReadBuffSize
	writeBuffSize := DefaultWriteBuffSize
	mergedWriteBuffSize := MinMergedWriteBuffSize
	if context.ReadBufferSize > 0 {
		readBuffSize = context.ReadBufferSize
	}
	if context.WriteBufferSize > 0 {
		writeBuffSize = context.WriteBufferSize
	}
	if context.MergedWriteBufferSize > mergedWriteBuffSize {
		mergedWriteBuffSize = context.MergedWriteBufferSize
	}
	c.(*net.TCPConn).SetReadBuffer(readBuffSize)
	c.(*net.TCPConn).SetWriteBuffer(writeBuffSize)
	conn.ioPumper = &TcpPumper{
		done:                done,
		readBuffSize:        readBuffSize,
		mergedWriteBuffSize: mergedWriteBuffSize,
		disableMergedWrite:  context.DisableMergedWrite,
	}
	return conn
}

func (this *TcpPumper) readPump(conn *Conn) {
	context := conn.context
	stat := conn.stat
	scanner := bufio.NewScanner(conn.Conn)
	scanner.Buffer(make([]byte, this.readBuffSize), this.readBuffSize)
	scanner.Split(context.Splitter)
	for {
		if ok := scanner.Scan(); ok {
			data := scanner.Bytes()
			conn.Session.OnRecv(conn, data)
			if stat != nil {
				stat.AddRecvStat(len(data), 1)
			}
		} else {
			logger.Error("tcpserver read error: %v", scanner.Err())
			break
		}
	}
	conn.Close() // finish writePump
}

func (this *TcpPumper) writePump(conn *Conn) {
	tickerPing := time.NewTicker(common.PingPeriod)
	stat := conn.stat
	outChan := conn.msgChan.GetOutChan()
	buff := NewDataBuff(this.mergedWriteBuffSize, !this.disableMergedWrite)
loop:
	for {
		select {
		case data := <-outChan:
			conn.SetWriteDeadline(time.Now().Add(common.WriteWait))
			rb, count := buff.GetData(data, outChan)
			_, err := conn.Conn.Write(rb)
			if err != nil {
				logger.Error("tcpserver write error: %v", err.Error())
				break loop
			}
			if stat != nil {
				stat.AddSendStat(len(rb), count)
				stat.SetSendChanItemCount(len(outChan))
			}
		case <-tickerPing.C:
		case <-this.done:
			break loop
		}
	}
	tickerPing.Stop()
	conn.Close()
}
