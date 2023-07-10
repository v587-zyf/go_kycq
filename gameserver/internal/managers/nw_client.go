package managers

import (
	"errors"
	"fmt"
	"time"

	"cqserver/gameserver/internal/gater"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbgt"
)

type ClientManager struct {
	util.DefaultModule
	*nw.DefaultSessionManager //保存clientsession
	server                    nw.Server
	port                      int
}

func NewClientManager(port int) *ClientManager {
	return &ClientManager{
		DefaultSessionManager: nw.NewDefaultSessionManager(true),
		port:                  port,
	}
}

func (this *ClientManager) Init() error {

	context := &nw.Context{
		SessionCreator: NewClientSession,
		Extra: &gater.ExtraContext{
			SessionFinder: func(id uint32) nw.Session { return this.GetSession(id) },
			StatusGetter:  func() *gater.ServerStatus { return &gater.ServerStatus{} },
		},
	}
	server := gater.NewServer(context)
	err := server.Start(fmt.Sprintf(":%d", this.port))
	if err != nil {
		return err
	}
	this.server = server
	return nil
}

func (this *ClientManager) Broadcast(sessionIds []uint32, msg nw.ProtoMessage) {
	rb, err := pb.Marshal(pb.GetCmdIdFromType(msg), 0, msg)
	if err != nil {
		logger.Error("Broadcast marshal broadcast message error")
		return
	}
	this.BroadcastData(sessionIds, rb)
}

func (this *ClientManager) BroadcastAll(msg nw.ProtoMessage) {
	rb, err := pb.Marshal(pb.GetCmdIdFromType(msg), 0, msg)
	if err != nil {
		logger.Error("BroadcastAll marshal broadcast message error")
		return
	}
	this.BroadcastData(nil, rb)
}

func (this *ClientManager) BroadcastData(sessionIds []uint32, data []byte) {

	if this.server == nil {
		logger.Error("nw_client server is not start")
		return
	}

	msg := &pbgt.BroadcastNtf{SessionIds: sessionIds, Msg: data}
	rb, err := pbgt.Marshal(pbgt.CmdBroadcastNtfId, 0, 0, msg)
	if err != nil {
		logger.Error("BroadcastData error msg marshal error:%v", err)
		return
	}

	this.server.Broadcast(nil, rb)
}

func (this *ClientManager) PutOutMessage(sessionId uint32, msg nw.ProtoMessage, sendNow bool) error {
	session := this.GetSession(sessionId)
	if session == nil {
		logger.Warn("给玩家发送消息，session为空,玩家：%v,msg:%v", sessionId, msg)
		return errors.New("获取玩家session错误")
	}
	session.(*ClientSession).PutOutMessage(msg, sendNow)
	return nil
}

func (this *ClientManager) Stop() {
	if this.server != nil {
		this.server.Stop()
	}
}

//获取保活状态
func (this *ClientManager) GetPingStatus() bool {

	lastPingTime := this.server.(*gater.Server).PingTime
	if int(time.Now().Unix())-lastPingTime > 30 {
		return false
	}
	return true
}

func (this *ClientManager) DispatchEvent(userId int, data interface{}, callback func(userId int, user *objs.User, data interface{})) {
	//logger.Info("开始DispatchEvent：%v",user.Id)
	user := m.GetUserManager().GetUser(userId)
	if user == nil {
		logger.Info("玩家:%v不在线，直接返回操作", userId)
		callback(userId, nil, data)
		return
	}
	session := this.GetSession(user.GateSessionId)
	if session != nil {
		userChangeData := &objs.UserDataChangeEvent{userId, data, callback}
		session.(*ClientSession).DispatchEvent(userChangeData)
	} else {
		logger.Info("玩家:%v session没有找到,直接返回操作", userId)
		callback(userId, nil, data)
	}
}
