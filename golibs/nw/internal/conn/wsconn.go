package conn

import (
	"errors"
	"time"

	"cqserver/golibs/nw"
	"cqserver/golibs/nw/internal/common"
	"github.com/gorilla/websocket"

	"cqserver/golibs/logger"
)

// wrap websocket.Conn to adopt net.Conn
type WsConnWrapper struct {
	*websocket.Conn
}

type WsPumper struct {
	wsConn              *websocket.Conn
	done                chan struct{}
	mergedWriteBuffSize int
	disableMergedWrite  bool
}

func NewWsConn(c *websocket.Conn, context *nw.Context) *Conn {
	done := make(chan struct{})
	msgChan := NewMessageChan(context.UseNoneBlockingChan, context.ChanSize, done)
	conn := newConn(&WsConnWrapper{Conn: c}, context, msgChan, done)
	mergedWriteBuffSize := MinMergedWriteBuffSize
	if context.MergedWriteBufferSize > mergedWriteBuffSize {
		mergedWriteBuffSize = context.MergedWriteBufferSize
	}
	if context.MaxMessageSize > 0 && mergedWriteBuffSize > context.MaxMessageSize {
		mergedWriteBuffSize = context.MaxMessageSize
	}
	conn.ioPumper = &WsPumper{
		wsConn:              c,
		done:                done,
		mergedWriteBuffSize: mergedWriteBuffSize,
		disableMergedWrite:  context.DisableMergedWrite,
	}
	return conn
}

func (this *WsConnWrapper) Read(b []byte) (int, error) {
	return 0, errors.New("not implemented")
}

func (this *WsConnWrapper) Write(data []byte) (int, error) {
	return 0, errors.New("not implemented")
}

func (this *WsConnWrapper) SetDeadline(t time.Time) error {
	return errors.New("not implemented")
}

func (this *WsPumper) readPump(conn *Conn) {
	wsConn := this.wsConn
	context := conn.context
	stat := conn.stat
	if context.MaxMessageSize > 0 {
		wsConn.SetReadLimit(int64(context.MaxMessageSize))
	}

	wsConn.SetReadDeadline(time.Now().Add(common.PongWait)) // client must send ping, or conn will be closed after PongWait
	wsConn.SetPongHandler(func(string) error {
		wsConn.SetReadDeadline(time.Now().Add(common.PongWait))
		return nil
	})

	for {
		_, data, err := wsConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNoStatusReceived, websocket.CloseAbnormalClosure) {
				logger.Info("wsserver read error: %s", err.Error())
			}
			break
		}
		conn.Session.OnRecv(conn, data)
		if stat != nil {
			stat.AddRecvStat(len(data), 1)
		}
	}
	conn.Close() // finish writePump
}

func (this *WsPumper) writePump(conn *Conn) {
	wsConn := this.wsConn
	pingTicker := time.NewTicker(common.PingPeriod)
	stat := conn.stat
	outChan := conn.msgChan.GetOutChan()
	buff := NewDataBuff(this.mergedWriteBuffSize, !this.disableMergedWrite)
loop:
	for {
		select {
		case data := <-outChan:
			wsConn.SetWriteDeadline(time.Now().Add(common.WriteWait))
			rb, count := buff.GetData(data, outChan)
			err := wsConn.WriteMessage(websocket.BinaryMessage, rb)
			if err != nil {
				logger.Info("wsserver write error: %s", err.Error())
				break loop
			}
			if stat != nil {
				stat.AddSendStat(len(rb), count)
				stat.SetSendChanItemCount(len(outChan))
			}
		case <-pingTicker.C:
			wsConn.SetWriteDeadline(time.Now().Add(common.WriteWait))
			err := wsConn.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				logger.Info("wsserver send ping error: %s", err.Error())
				break loop
			}
		case <-this.done:
			break loop
		}
	}
	pingTicker.Stop()
	conn.Close()
}
