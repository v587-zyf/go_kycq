package net

import (
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/tcpserver"
	"cqserver/golibs/util"
	"cqserver/protobuf/pbserver"
	"errors"
	"fmt"
)

type GsConn interface {
	/**
	 *  @Description: 发送消息给gs(普通战斗服发送给链接game，跨服则发送给跨服组所有game服)
	 *  @param msg
	 *  @return error/**
	 */
	SendMessage(msg nw.ProtoMessage) error

	/**
	 *  @Description: 发送消息给gs(普通战斗服发送给链接game，跨服则发送给所有game服)
	 *  @param msg
	 *  @return error/**
	 */
	SendMessageAll(msg nw.ProtoMessage) error


	/**
	 *  @Description: 给指定gs服发送消息（普通战斗服 跨服战斗服都可使用）
	 *  @param msg
	 *  @return error
	 */
	SendMessageToGs(serverId uint32, msg nw.ProtoMessage) error

	///**
	// *  @Description: rpc回调消息
	// *  @param serverId
	// *  @param req
	// *  @param resp
	// *  @return error
	// */
	//CallGs(serverId uint32, req nw.ProtoMessage, resp nw.ProtoMessage) error
}

type GSManager struct {
	util.DefaultModule
	server  nw.Server
	session *GSSession
	port    int
}

var gsManage *GSManager

func GetGsConn() GsConn {
	return gsManage
}

func NewGSManager(port int) *GSManager {
	gsManage = &GSManager{port: port}
	return gsManage
}

func (this *GSManager) Init() error {
	context := &nw.Context{
		SessionCreator: this.crateSession,
		Splitter:       pbserver.Split,
		ReadBufferSize: 1024 * 1024,
		ChanSize:       1024 * 4,
	}
	server := tcpserver.NewServer(context)
	err := server.Start(fmt.Sprintf(":%d", this.port))
	if err != nil {
		return err
	}
	this.server = server
	return nil
}

func (this *GSManager) crateSession(conn nw.Conn) nw.Session {

	this.session = NewGSSession(conn).(*GSSession)
	return this.session
}

func (this *GSManager) CallGs(serverId uint32, req nw.ProtoMessage, resp nw.ProtoMessage) error {
	return this.session.Call(serverId, req, resp)
}

func (this *GSManager) SendMessageToGs(serverId uint32, msg nw.ProtoMessage) error {
	if this.session != nil {
		return this.session.SendMessage(int32(serverId), msg)
	}
	return errors.New("session not found")
}

func (this *GSManager) SendMessage(msg nw.ProtoMessage) error {
	if this.session != nil {
		return this.session.SendMessage(0, msg)
	}
	return errors.New("session not found")
}

func (this *GSManager) SendMessageAll(msg nw.ProtoMessage) error {
	if this.session != nil {
		return this.session.SendMessage(-1, msg)
	}
	return errors.New("session not found")
}

func (this *GSManager) GetGSSession() *GSSession {
	return this.session
}

func (this *GSManager) Stop() {
	if this.server != nil {
		this.server.Stop()
	}
}
