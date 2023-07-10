package manager

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constServer"
	"errors"
	"fmt"
	"time"

	"cqserver/protobuf/pbgt"

	"cqserver/golibs/nw"

	"cqserver/golibs/logger"
	"cqserver/golibs/nw/tcpclient"
	"cqserver/golibs/util"
)

type FSManager struct {
	util.DefaultModule

	sessions             map[int]*FSSession // 固定创建各个FSSession，不需要加锁.共2个:本地,动态中心.
	dialerCloseChans     map[int]chan struct{}
	localFightServerPort int
}

func NewFSManager(localFightServerPort int) *FSManager {
	connector := &FSManager{
		sessions:             make(map[int]*FSSession),
		dialerCloseChans:     make(map[int]chan struct{}),
		localFightServerPort: localFightServerPort,
	}
	return connector
}

func GetFSSessionCreator(session *FSSession) func(conn nw.Conn) nw.Session {
	return func(conn nw.Conn) nw.Session {
		//s := fsManager.GetSession(fsNo)
		session.Conn = conn
		return session
	}
}

func (this *FSManager) Init() error {

	// 本地fightServer
	this.sessions[constServer.FIGHT_SESSIONID_LOCAL] = NewFSSession(constServer.FIGHT_SESSIONID_LOCAL, "", this.localFightServerPort)

	serverPortInfo, err := modelCross.GetServerPortInfoModel().GetServerPortInfo(modelCross.GATE_TO_FIGHTCENTER)
	if err != nil {
		return err
	}
	this.sessions[constServer.FIGHT_SESSIONID_CENTER] = NewFSSession(constServer.FIGHT_SESSIONID_CENTER, serverPortInfo.Host, serverPortInfo.Port) //动态战斗服消息中心

	for _, v := range this.sessions {
		this.connectFs(v)
	}

	return nil
}

// 连接战斗服
func (this *FSManager) connectFs(session *FSSession) {
	context := &nw.Context{
		SessionCreator: GetFSSessionCreator(session),
		Splitter:       pbgt.Split,
		ReadBufferSize: 1024 * 1024,
		ChanSize:       1024 * 8,
	}
	addr := fmt.Sprintf("%s:%d", session.Host, session.Port)
	dialerCloseChan := tcpclient.DialEx(addr, context, 3*time.Second)
	this.dialerCloseChans[session.FSNo] = dialerCloseChan
	logger.Info("tcpclient.DialEx, addr=%v, FSNo=%v dialerCloseChans len=%v", addr, session.FSNo, len(this.dialerCloseChans))
}

func (this *FSManager) Stop() {
	for _, session := range this.sessions {
		conn := session.Conn
		if conn != nil {
			conn.Close()
			conn.Wait()
		}
	}
	for _, v := range this.dialerCloseChans {
		close(v)
	}
	logger.Info("nw_fs stop")
}

// 获取所有的战斗连接
func (this *FSManager) GetSession(fsNo int) *FSSession {

	if fsNo <= 1 {
		return this.sessions[constServer.FIGHT_SESSIONID_LOCAL]
	} else {
		return this.sessions[constServer.FIGHT_SESSIONID_CENTER]
	}
}

func (this *FSManager) GateRouteMsgToFs(fightServerId int, data []byte) error {

	fsSession := this.GetSession(fightServerId)
	if fsSession == nil {
		return errors.New("fs connect not found")
	}
	return fsSession.WriteToFS(fightServerId, data)
}
