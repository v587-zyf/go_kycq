package managers

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constServer"
	"time"

	"cqserver/golibs/nw"

	"fmt"

	"cqserver/golibs/logger"
	"cqserver/golibs/nw/tcpclient"
	"cqserver/golibs/util"
	"cqserver/protobuf/pbserver"
	"github.com/pkg/errors"
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

func (this *FSManager) Init() error {

	this.sessions[constServer.FIGHT_SESSIONID_LOCAL] = NewFSSession(constServer.FIGHT_SESSIONID_LOCAL, "", this.localFightServerPort)

	serverPortInfo, err := modelCross.GetServerPortInfoModel().GetServerPortInfo(modelCross.GAME_TO_FIGHTCENTER)
	if err != nil {
		return err
	}
	session := NewFSSession(constServer.FIGHT_SESSIONID_CENTER, serverPortInfo.Host, serverPortInfo.Port)
	this.sessions[constServer.FIGHT_SESSIONID_CENTER] = session //动态战斗服消息中心

	for _, v := range this.sessions {
		this.connectFs(v)
	}

	return nil
}

func GetFSSessionCreator(session *FSSession) func(conn nw.Conn) nw.Session {
	return func(conn nw.Conn) nw.Session {

		session.Conn = conn
		//logger.Info("======GetFSSessionCreator fsNo=%v addr=%v", fsNo, s.Addr)
		return session
	}
}

func (this *FSManager) connectFs(session *FSSession) {
	context := &nw.Context{
		SessionCreator: GetFSSessionCreator(session),
		Splitter:       pbserver.Split,
		ReadBufferSize: 1024 * 1024,
	}
	addr := session.Addr
	logger.Info("======tcpclient.DialEx, fsNo=%v, addr=%v", session.FSNo, addr)
	dialerCloseChan := tcpclient.DialEx(addr, context, 3*time.Second)
	this.dialerCloseChans[session.FSNo] = dialerCloseChan
}

func (this *FSManager) GetSession(fsNo int) *FSSession {
	if fsNo > 0 {
		fsNo = constServer.FIGHT_SESSIONID_CENTER
	} else {
		fsNo = constServer.FIGHT_SESSIONID_LOCAL
	}
	fs := this.sessions[fsNo]
	return fs
}

func (this *FSManager) GetSessions() map[int]*FSSession {
	return this.sessions
}

func (this *FSManager) Stop() {
	for _, session := range this.sessions {
		conn := session.Conn
		if conn != nil {
			conn.Close()
			conn.Wait()
		}
	}

	for _, dialerCloseChans := range this.dialerCloseChans {
		close(dialerCloseChans)
	}
}

// 本服战斗服+静态战斗跨服 RpcCall
func (this *FSManager) RpcCall(fightId int, stageId int, req nw.ProtoMessage, resp nw.ProtoMessage) error {
	if fightId > 0 {
		var err error
		req, err = this.createGsRouteMessage(fightId, req)
		if err != nil {
			return err
		}
	}
	cmdId := pbserver.GetCmdIdFromType(req)

	crossFightServerId := 0
	if stageId > 0 {
		crossFightServerId = this.GetCrossFightServerId(stageId)
	}
	session := this.GetSession(crossFightServerId)
	if session != nil {
		err := session.Call(crossFightServerId, req, resp)
		return err
	}
	return errors.New(fmt.Sprintf("FSManager RpcCall failed, fsNo=%v  cmdId=%v", 0, pbserver.GetMsgName(cmdId)))
}

// 本服战斗服+静态战斗跨服 RpcCall
func (this *FSManager) RpcCallByFightServerId(crossFightServerId int, req nw.ProtoMessage, resp nw.ProtoMessage) error {
	cmdId := pbserver.GetCmdIdFromType(req)
	session := this.GetSession(crossFightServerId)
	if session != nil {
		err := session.Call(crossFightServerId, req, resp)
		return err
	}
	return errors.New(fmt.Sprintf("FSManager RpcCall failed, fsNo=%v  cmdId=%v", crossFightServerId, pbserver.GetMsgName(cmdId)))
}

func (this *FSManager) SendMessage(fightId int, stageId int, msg nw.ProtoMessage) error {
	if fightId > 0 {
		var err error
		msg, err = this.createGsRouteMessage(fightId, msg)
		if err != nil {
			return err
		}
	}
	crossFightServerId := this.GetCrossFightServerId(stageId)
	session := this.GetSession(crossFightServerId)
	if session != nil {
		if session.IsConnected() {
			return session.SendMessage(crossFightServerId, 0, msg)
		} else {
			logger.Warn("SendMessageToCrossFightServer fail. session is not connected")
			return errors.New("session closed")
		}
	} else {
		logger.Warn("SendMessageToCrossFightServer fail. session is nil")
		return errors.New("session not found")
	}
}

func (this *FSManager) GetCrossFightServerId(stageId int) int {
	crossFightServerId := 0
	stageConf := gamedb.GetStageStageCfg(stageId)
	if stageConf != nil {
		if stageConf.Type == constFight.FIGHT_TYPE_CROSS_WORLD_LEADER ||
			stageConf.Type == constFight.FIGHT_TYPE_CROSS_SHABAKE {
			crossFightServerId = m.GetSystem().GetCrossFightServerId()
		} else if stageConf.Type == constFight.FIGHT_TYPE_HELL_BOSS ||
			stageConf.Type == constFight.FIGHT_TYPE_HELL ||
			stageConf.Type == constFight.FIGHT_TYPE_MAGIC_TOWER ||
			stageConf.Type == constFight.FIGHT_TYPE_SHABAKE_NEW {
			if m.GetSystem().IsCross() {
				crossFightServerId = m.GetSystem().GetCrossFightServerId()
			}
		}
	}
	return crossFightServerId
}

func (this *FSManager) createGsRouteMessage(fightId int, msg nw.ProtoMessage) (nw.ProtoMessage, error) {

	msgData, err := msg.Marshal()
	if err != nil {
		return nil, err
	}
	msg = &pbserver.GsRouteMessageToFight{
		FightId: int32(fightId),
		CmdId:   int32(pbserver.GetCmdIdFromType(msg)),
		MsgData: msgData,
	}
	return msg, nil
}
