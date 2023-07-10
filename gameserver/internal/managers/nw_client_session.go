package managers

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/gater"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbgt"
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

var (
	WITHOUT_LOG_MSG_CMDID = map[int]bool{
		2: true,
	}
)

type ClientSession struct {
	Id         uint32
	Messages   chan util.Message
	outMessage chan nw.ProtoMessage //给玩家主动推送的消息
	Conn       nw.Conn
	User       *objs.User
	IsKick     bool
	done       chan struct{}
	event      chan *objs.UserDataChangeEvent
	wg         sync.WaitGroup
}

type ClientSessionMessage struct {
	util.AsyncMessage
	Session  *ClientSession
	MsgFrame *pb.MessageFrame
}

func (this *ClientSessionMessage) Handle() {
	this.Session.HandleMessage(this.MsgFrame)
}

func NewClientSession(conn nw.Conn) nw.Session {
	s := &ClientSession{
		Id:         conn.(*gater.ClientConn).GetGateClientSessionId(),
		Messages:   make(chan util.Message, 100),
		outMessage: make(chan nw.ProtoMessage, 50),
		Conn:       conn,
		IsKick:     false,
		done:       make(chan struct{}),
		event:      make(chan *objs.UserDataChangeEvent, 10),
	}
	return s
}

func (this *ClientSession) GetId() uint32 {
	return this.Id
}

func (this *ClientSession) GetConn() nw.Conn {
	return this.Conn
}

func (this *ClientSession) OnOpen(conn nw.Conn) {
	m.ClientManager.AddSession(this.Id, this)
	logger.Info("user client session sessionId:%v,on open", this.Id)
	this.wg.Add(1)
	go this.messagePump()
}

func (this *ClientSession) OnClose(conn nw.Conn) {
	logger.Info("user client session sessionId:%v,on close", this.Id)

	u := this.User
	if u != nil {
		u.OffLineWg.Add(1)
	}
	close(this.done)
	this.wg.Wait()
	if u != nil {
		u.OffLineWg.Done()
	}
	m.ClientManager.RemoveSession(this.Id)
	logger.Info("user client session sessionId:%v remove", this.Id)
}

func (this *ClientSession) OnRecv(conn nw.Conn, data []byte) {

	//判断服务器是否开起来
	if !m.ServerStarted {
		logger.Error("服务器还未启动完成")
		this.Conn.Close()
		return
	}

	msgFrame, err := pb.Unmarshal(data)
	if err != nil {
		logger.Error("****************unknown message, cmdId: %d************************", pb.GetCmdId(data))
		this.Conn.Close()
		return
	}
	if this.User == nil && msgFrame.CmdId != pb.CmdEnterGameReqId {
		logger.Debug("handler msg but not user. cmdId: %d", msgFrame.CmdId)
		this.Conn.Close()
		return
	}

	this.Messages <- &ClientSessionMessage{Session: this, MsgFrame: msgFrame}
}

func (this *ClientSession) HandleMessage(msgFrame *pb.MessageFrame) {

	if msgFrame.CmdId != pb.CmdEnterGameReqId && this.User == nil {
		logger.Error("消息异常，玩家还未完成登录")
		return
	}

	if this.IsKick {
		err := this.SendMessage(msgFrame.TransId, &pb.ErrorAck{Code: -999, Message: "服务器已断开连接"})
		if err != nil {
			logger.Debug(err.Error())
		}
		return
	}
	handler := pb.GetHandler(msgFrame.CmdId)
	if handler == nil {
		logger.Error("no handler found, cmdId: %d", msgFrame.CmdId)
		this.SendMessage(0, &pb.ErrorAck{-1, fmt.Sprintf("no handler found cmdID:%d", msgFrame.CmdId)})
		return
	}

	startAt := time.Now()
	ack, opGoodsHelper, err := handler(this.Conn, msgFrame.Body)
	if err != nil {
		logger.Error("nw_client_session:HandleMessage:msg:%s,body:%v,err:%v", pb.GetMsgName(msgFrame.CmdId), msgFrame.Body, err)
		ei, ok := err.(util.IError)
		if !ok {
			ei = gamedb.ERRUNKNOW
			if base.Conf.Sandbox {
				ei = gamedb.ERRUNKNOW.CloneWithMsg(fmt.Sprintf("消息号:%d,消息体:%v,异常：%v", msgFrame.CmdId, msgFrame.Body, err.Error()))
			}
		}
		ack = &pb.ErrorAck{Code: int32(ei.ErrCode()), Message: ei.Error()}
	}

	//数据通知
	if opGoodsHelper != nil {
		ntfs := opGoodsHelper.ToGoodsChangeMessages()
		for _, v := range ntfs {
			this.SendMessage(msgFrame.TransId, v)
		}
	}

	if ack != nil {
		err = this.SendMessage(msgFrame.TransId, ack)
		if err != nil {
			logger.Debug(err.Error())
		}
	}
	if WITHOUT_LOG_MSG_CMDID[int(msgFrame.CmdId)] {
		return
	}
	usedTime := time.Now().Sub(startAt)
	userIdName := ""
	if this.User != nil {
		userIdName = this.User.IdName()
	}
	if base.Conf.Sandbox {
		logger.Debug("nw_client_session.玩家%v 消息处理耗时【%v】 接收【%d:%v:%+v】  返回【%+v】", userIdName, usedTime, msgFrame.CmdId, pb.GetMsgName(msgFrame.CmdId), msgFrame.Body, ack)
	} else {
		if usedTime > 100*time.Millisecond {
			logger.Warn("nw_client_session. Handle cost too much time. user:%v msg:%s usedtime:%v", userIdName, pb.GetMsgName(msgFrame.CmdId), usedTime)
		}
	}
}

/**
 *  @Description: 推送玩家客户端消息
 *  @param msg
 *  @param rightNow 是否立即发送
 */
func (this *ClientSession) PutOutMessage(msg nw.ProtoMessage, sendNow bool) {

	if sendNow {
		this.SendMessage(0, msg)
	} else {
		if len(this.outMessage) >= 50 {
			if base.Conf.Sandbox {
				panic(fmt.Sprintf("消息过多了:%v", msg))
			}
			return
		}
		this.outMessage <- msg
	}
}

func (this *ClientSession) SendMessage(transId uint32, msg nw.ProtoMessage) error {
	pbgtCmdId := pbgt.GetCmdIdFromType(msg)
	if pbgtCmdId > 0 {
		this.Conn.(*gater.ClientConn).GateSession.SendMessage(pbgtCmdId, this.Id, msg)
		return nil
	}

	cmdId := pb.GetCmdIdFromType(msg)
	rb, err := pb.Marshal(cmdId, transId, msg)
	if err != nil {
		return err
	}
	if cmdId == pb.CmdKickUserNtfId {
		this.IsKick = true
		routemsg, _ := pbgt.Marshal(pbgt.CmdRouteMessageId, this.Id, 0, rb)
		this.Conn.(*gater.ClientConn).GateSession.Conn.Write(routemsg)
	}

	if this.IsKick {
		return nil
	}

	_, err = this.Conn.Write(rb)
	return err
}

func (this *ClientSession) messagePump() {
	tickerSave := time.NewTicker(time.Minute * 5)
	tickerSecond := time.NewTicker(time.Second * 3)
	tickerSecond1 := time.NewTicker(time.Second * 1)
	defer func() {
		this.wg.Done()
		tickerSave.Stop()
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic when messagePump:%v,%s,%v", r, stackBytes, this.User)
			if this.User != nil {
				m.UserManager.UserDisconnect(this.User)
				m.UserManager.KickUserWithMsg(this.User, "程序出错")
			}
		}
	}()
	for {
		select {
		case msg := <-this.Messages:
			msg.Handle()
			msg.Done()
		case msg := <-this.outMessage:
			this.SendMessage(0, msg)
		case <-tickerSave.C:
			user := this.User
			if user != nil {
				m.GetOnline().OffLine(user, false)
				m.UserManager.Save(user, false)
				// //防沉迷
				// m.Health.UpdateUserInfo(user)
			}
		case <-tickerSecond.C:
			user := this.User
			if user != nil {
				if this.User.GetUpdateFightUserHeroState() {
					m.GetFight().UpdateUserInfoToFight(this.User, this.User.GetUpdateFightUserHeroIndex(), false)
					this.User.ResetUpdateFightUserHeroInfo()
				}
				if user.CheckTitleExpire {
					m.Title.CheckExpire(user)
				}
				m.TreasureShop.AutoRefreshShop(user, true)
				m.DaBao.ResumeEnergy(user)
			}
		case <-tickerSecond1.C:
			user := this.User
			if user != nil {
				//小程序 定时检查定时增加体力值
				m.GetApplets().CronAddPhysicalPower(user)
			}
		case eventData, ok := <-this.event:

			user := this.User
			if !ok {
				userId := -1
				if user != nil {
					userId = user.Id
				}
				logger.Error("接收错误事件：%v", userId)
			} else {
				this.eventListener(eventData, this.User)
			}
		case <-this.done:
			if this.User != nil {
				m.UserManager.UserDisconnect(this.User)
			}
			//处理可能剩余的event
			for len(this.event) > 0 {
				if eventData, ok := <-this.event; ok {
					this.eventListener(eventData, nil)
				}
			}
			return
		}
	}
}

func (this *ClientSession) DispatchEvent(userChangeData *objs.UserDataChangeEvent) {
	this.event <- userChangeData
}

func (this *ClientSession) eventListener(userChangeData *objs.UserDataChangeEvent, user *objs.User) {
	userChangeData.Callback(userChangeData.UserId, user, userChangeData.Data)
}
