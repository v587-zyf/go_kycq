package handler

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/errex"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"runtime/debug"
)

type ClientMessageHandler func(fight base.Fight, actor base.Actor, msg nw.ProtoMessage) (nw.ProtoMessage, error)
type ServerMessageHandler func(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error)

var clientHandlers = make(map[uint16]ClientMessageHandler)
var serverHandlers = make(map[uint16]ServerMessageHandler)

func RegisterClientHandler(cmdId uint16, handler ClientMessageHandler) {
	clientHandlers[cmdId] = handler
}

func RegisterServerHandler(cmdId uint16, handler ServerMessageHandler) {
	serverHandlers[cmdId] = handler
}

func GetClientHandler(cmdId uint16) ClientMessageHandler {
	return clientHandlers[cmdId]
}

func GetServerHandler(cmdId uint16) ServerMessageHandler {
	return serverHandlers[cmdId]
}

type ClientMessage struct {
	fight     base.Fight
	sessionId uint32
	msgFrame  *pb.MessageFrame
}

func NewClientMessage(fight base.Fight, sessionId uint32, msgFrame *pb.MessageFrame) *ClientMessage {
	return &ClientMessage{fight, sessionId, msgFrame}
}

type GSMessage struct {
	fight    base.Fight
	conn     nw.Conn
	transId  uint32
	cmdId    int32
	msgFrame nw.ProtoMessage
}

func NewGSMessage(fight base.Fight, conn nw.Conn, transId uint32, cmdId int32, msgFrame nw.ProtoMessage) *GSMessage {
	return &GSMessage{fight, conn, transId, cmdId, msgFrame}
}

func (this *ClientMessage) Handle() {

	/****这里的reqActor只是请求消息的玩家，并不一定是执行指令的玩家****/
	reqActor := this.fight.GetOnlineActor(this.sessionId)
	if reqActor == nil {
		logger.Warn("not found online actor :%d", this.sessionId)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic fight message_client:%v,%s", r, stackBytes)
		}
	}()

	msgFrame := this.msgFrame
	handler := GetClientHandler(msgFrame.CmdId)
	if handler == nil {
		return
	}
	ack, err := handler(this.fight, reqActor, msgFrame.Body.(nw.ProtoMessage))
	if err != nil {
		ack = errex.BuildClientErrorAck(err)
	}
	if ack != nil {
		serverId := uint32(reqActor.HostId())
		net.GetGateConn().SendMessage(serverId, this.sessionId, msgFrame.TransId, ack)
	}
}

func (this *GSMessage) Handle() {

	defer func() {

		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic fight message_gs:%v,%s", r, stackBytes)
		}
	}()
	msgFrame := this.msgFrame
	if msgFrame == nil {
		logger.Error("message_gs:Handle msgFrame is nil")
		return
	}
	h := GetServerHandler(uint16(this.cmdId))
	if h == nil {
		logger.Error("message_gs:No Handle for cmd:%d", this.cmdId)
		return
	}
	ack, err := h(this.fight, msgFrame)
	if err != nil {
		ack = errex.BuildServerErrorAck(err)
	}
	if ack != nil {
		logger.Info("gs return:", ack)
		this.conn.GetSession().(*net.GSSession).ReplyMessage(this.transId, ack)
	}
}
