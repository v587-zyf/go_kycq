package net

import (
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/tcpserver"
	"cqserver/golibs/util"
	"cqserver/protobuf/pbgt"
	"fmt"
)

type GateConn interface {
	SendMessage(serverId, sessionId, transId uint32, msg nw.ProtoMessage) error
	BroadcastToGate(ids map[int]map[int]int, msg nw.ProtoMessage)
}

var gateManager *GateManager

func GetGateConn() GateConn {
	return gateManager
}

type GateManager struct {
	util.DefaultModule
	server  nw.Server
	session *GateSession // 固定创建各个GSSession，不需要加锁
	port    int
}

func NewGateManager(port int) *GateManager {
	gateManager = &GateManager{port: port}
	return gateManager
}

func (this *GateManager) Init() error {

	context := &nw.Context{
		SessionCreator:  this.crateGateSession,
		Splitter:        pbgt.Split,
		ReadBufferSize:  2 * 1024 * 1024,
		WriteBufferSize: 2 * 1024 * 1024,
		ChanSize:        1024 * 8,
	}
	server := tcpserver.NewServer(context)
	err := server.Start(fmt.Sprintf(":%d", this.port))
	if err != nil {
		return err
	}
	this.server = server
	return nil
}

func (this *GateManager) crateGateSession(conn nw.Conn) nw.Session {

	this.session = NewGateSession(conn).(*GateSession)
	return this.session
}

func (this *GateManager) SendMessage(serverId, sessionId, transId uint32, msg nw.ProtoMessage) error {
	return this.session.SendMessage(serverId, sessionId, transId, msg)
}

func (this *GateManager) BroadcastToGate(ids map[int]map[int]int, msg nw.ProtoMessage) {

	this.session.BroadcastToGate(ids, msg)
}

func (this *GateManager) Stop() {
	if this.server != nil {
		this.server.Stop()
	}
}
