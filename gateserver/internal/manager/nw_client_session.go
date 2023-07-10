package manager

import (
	"cqserver/gamelibs/errex"
	"cqserver/gamelibs/publicCon/constServer"
	"cqserver/gamelibs/rmodel"
	"cqserver/gateserver/conf"
	"runtime/debug"
	"time"

	"cqserver/golibs/nw"

	"sync"

	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbgt"
)

type ClientSession struct {
	Id                 uint32
	Messages           chan interface{}
	Conn               nw.Conn
	openId             string //玩家账号
	curFightId         int32  // 当前的fightId
	fsSessionNo        int    //战斗服服务器ID
	isClose            bool
	loginKeyUpdateTime int64
	broadcast          bool
}

const (
	FIGHT_CMDID_MIN = 9001
	FIGHT_CMDID_MAX = 9999
	ALLOCCATE_RANGE = 1000000
)

var (
	WITHOUT_LOG_MSG_CMDID = map[int]bool{
		pb.CmdAttackRptId:    true,
		pb.CmdSceneMoveRptId: true,
		pb.CmdPingReqId:      true,
	}
)

var (
	allMu      sync.RWMutex
	allocateId = uint32(0)
)

func isRouteToFS(cmdId uint16) bool {
	return cmdId >= FIGHT_CMDID_MIN && cmdId <= FIGHT_CMDID_MAX
}

func NewClientSession(conn nw.Conn) *ClientSession {
	s := &ClientSession{
		Id:       getAllocatorSession(),
		Messages: make(chan interface{}, 100),
		Conn:     conn,
	}
	return s
}

func getAllocatorSession() uint32 {
	allMu.Lock()
	allocateId++
	if allocateId > ALLOCCATE_RANGE {
		allocateId = 0
	}
	newSessionId := uint32(conf.Conf.ServerId*ALLOCCATE_RANGE) + allocateId
	allMu.Unlock()
	return newSessionId
}

func (this *ClientSession) SetOpenId(openId string) {
	this.openId = openId
}

func (this *ClientSession) GetOpenId() string {
	return this.openId
}

func (this *ClientSession) GetId() uint32 {
	return this.Id
}

func (this *ClientSession) GetConn() nw.Conn {
	return this.Conn
}

func (this *ClientSession) OnOpen(conn nw.Conn) {
	m.ClientManager.AddSession(this.Id, this)
}

func (this *ClientSession) OnClose(conn nw.Conn) {
	m.ClientManager.RemoveSession(this.Id)
	this.SendUserQuitToGS()
}

func (this *ClientSession) OnRecv(conn nw.Conn, data []byte) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("ClientSession OnRecv panic: %v, %s", err, (debug.Stack()))
		}
	}()

	if this.isClose {
		logger.Info("session链接已标记关闭")
		return
	}

	// 注意: 此处data直接引用的网络缓冲区的slice，如果把data发送给其他goroutine处理，需要注意缓冲区覆盖问题
	cmdId := pb.GetCmdId(data)
	if msgPrototype := pb.GetMsgPrototype(cmdId); msgPrototype == nil {
		logger.Error("route client message to server error, CmdId error:%v",cmdId)
		conn.Close()
		return
	}
	//标识玩家是否可广播
	if cmdId == pb.CmdUserInGameOkReqId{
		this.broadcast = true
		return
	}

	if cmdId == pb.CmdEnterGameReqId {
		msgFrame, err := pb.Unmarshal(data)
		if err != nil {
			logger.Error("unmarshal client message error, CmdId: %d, error: %s", cmdId, err.Error())
			conn.Close()
			return
		}
		req := msgFrame.Body.(*pb.EnterGameReq)
		if !conf.Conf.Sandbox && !rmodel.GetLoginKeyModel().ValidateLoginKey(req.OpenId, req.LoginKey) {
			this.SendMessageToClient(0, &pb.ErrorAck{Code: int32(errex.ErrLoginKeyError.Code), Message: errex.ErrLoginKeyError.Message})
			return
		}

		req.Ip = this.Conn.RemoteAddr().String()
		data, err = pb.Marshal(cmdId, msgFrame.TransId, req)
		if err != nil {
			logger.Error("unmarshal client message error, CmdId: %d, error: %s", cmdId, err.Error())
			conn.Close()
		}
		this.loginKeyUpdateTime = time.Now().Unix()
		this.SetOpenId(req.OpenId)
	} else {
		if time.Now().Unix()-this.loginKeyUpdateTime > 300 {
			rmodel.GetLoginKeyModel().ValidateLoginUpdate(this.openId)
			this.loginKeyUpdateTime = time.Now().Unix()
		}
	}
	if !WITHOUT_LOG_MSG_CMDID[int(cmdId)] {
		logger.Debug("收到客户端发送来的消息,openId:%v，cmdId:%v", this.openId, cmdId)
	}
	var err error
	if isRouteToFS(cmdId) {
		err = this.routeToFs(data)
	} else {
		err = m.GsManager.RouteToGS(this.GetId(), data)
	}
	if err != nil {
		logger.Error("route client message to server error, CmdId: %d, error: %s,session len:%v,s:%v", cmdId, err, len(m.FsManager.sessions), m.FsManager.sessions)
		conn.Close()
		return
	}
}

func (this *ClientSession) routeToFs(data []byte) error {
	rb, err := pbgt.Marshal(pbgt.CmdRouteMessageId, this.GetId(), uint32(this.curFightId), data)
	if err != nil {
		return err
	}
	return m.FsManager.GateRouteMsgToFs(this.fsSessionNo, rb)
}

// 每次CopyReadAck都会绑定当前的战斗Session
func (this *ClientSession) BindFightSession(fightId int32, crossFightServerId int) {
	this.curFightId = fightId
	// 优先动态战斗跨服
	if crossFightServerId > 0 {
		this.fsSessionNo = crossFightServerId
		logger.Info("绑定动态跨服 FSSession fightId:%v,crossGroup:%v,dynamicFightServerId:%v fsNo=%v", fightId, crossFightServerId)
	} else {
		this.fsSessionNo = constServer.FIGHT_SESSIONID_LOCAL
		logger.Info("绑定本地战斗服 FSSession fightId:%v,crossGroup:%v,dynamicFightServerId:%v fsNo=nil", fightId, crossFightServerId)
	}
}

func (this *ClientSession) WriteToClient(data []byte) error {
	_, err := this.Conn.Write(data)
	return err
}

func (this *ClientSession) SendMessageToClient(transId uint32, msg nw.ProtoMessage) error {
	rb, err := pb.Marshal(pb.GetCmdIdFromType(msg), transId, msg)
	if err != nil {
		return err
	}
	_, err = this.Conn.Write(rb)
	return err
}

func (this *ClientSession) SendUserQuitToGS() error {
	rb, err := pbgt.Marshal(pbgt.CmdUserQuitRptId, this.GetId(), 0, &pbgt.UserQuitRpt{})
	if err != nil {
		logger.Error("UserQuitReq marshal error: %s", err.Error())
		return err
	}
	logger.Info("user in or out openId:%v quit,clientSessionId:%v", this.openId, this.Id)
	err = m.GsManager.WriteToGS(rb)
	if err != nil {
		logger.Error("UserQuitReq send error: %s", err.Error())
		return err
	}
	return nil
}
