package manager

import (
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/tcpclient"
	"cqserver/golibs/util"
	"cqserver/protobuf/pbgt"
	"strconv"
	"time"
)

type GSManager struct {
	util.DefaultModule

	sessions         *GSSession // 固定创建各个GSSession，不需要加锁
	dialerCloseChans chan struct{}
	port             int
}

func NewGSManager(port int) *GSManager {
	connector := &GSManager{
		port: port,
	}
	connector.sessions = NewGSSession()
	return connector
}

func GetGSSessionCreator(gsManager *GSManager) func(conn nw.Conn) nw.Session {
	return func(conn nw.Conn) nw.Session {
		s := gsManager.GetSession()
		s.Conn = conn
		return s
	}
}

func (this *GSManager) Init() error {
	context := &nw.Context{
		SessionCreator: GetGSSessionCreator(this),
		Splitter:       pbgt.Split,
		ReadBufferSize: 1024 * 1024,
		ChanSize:       10000,
	}
	this.dialerCloseChans = tcpclient.DialEx(":"+strconv.Itoa(this.port), context, 3*time.Second)
	logger.Info("端口开始连接【%v】,连接gameserver使用", this.port)
	return nil
}

func (this *GSManager) GetSession() *GSSession {
	return this.sessions
}

func (this *GSManager) Stop() {
	conn := this.sessions.Conn
	if conn != nil {
		conn.Close()
		conn.Wait()
	}
	close(this.dialerCloseChans)
	logger.Info("nw_gs stop")
}

func (this *GSManager) RouteToGS(clientSessionId uint32, data []byte) error {
	rb, err := pbgt.Marshal(pbgt.CmdRouteMessageId, clientSessionId, 0, data)
	if err != nil {
		return err
	}
	return this.sessions.Write(rb)
}

func (this *GSManager) WriteToGS(data []byte) error {
	return this.sessions.Write(data)
}
