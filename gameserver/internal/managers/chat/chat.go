package chat

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/ptsdk"
	"cqserver/gamelibs/publicCon/constChat"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"strings"
	"sync"
	"time"
)

type chatReport struct {
	chat   *pb.ChatSendReq
	sender int
}

type ChatManager struct {
	util.DefaultModule
	managersI.IModule
	nearestMsg     map[int][]*pb.ChatMessageNtf //保存最近的10条
	chatInterval   map[int]int64                //记录玩家上次发言时间，管理聊天间隔
	mu             sync.RWMutex
	chatReprotChan chan *chatReport // 数据库操作通道
}

func NewChatManager(module managersI.IModule) *ChatManager {
	return &ChatManager{
		IModule:        module,
		nearestMsg:     make(map[int][]*pb.ChatMessageNtf, 10),
		chatInterval:   make(map[int]int64),
		chatReprotChan: make(chan *chatReport, constChat.CHAT_REPORT_MAX_CHAN),
	}
}

func (this *ChatManager) Init() error {
	this.chatReportLoop()
	return nil
}

//保存最近的10条
func (this *ChatManager) addToNearestMsg(one *pb.ChatMessageNtf) {

}

//用户登录后发送最近的10条信息
func (this *ChatManager) GetNearestMsg() *pb.ChatMessageListAck {
	var msgs = &pb.ChatMessageListAck{}
	return msgs
}

//检查频度
func (this *ChatManager) checkCanSpeak(userId int) bool {
	this.mu.RLock()
	defer this.mu.RUnlock()

	c := this.chatInterval[userId]

	interval := int64(time.Duration(gamedb.GetConf().Chat[0]) * 1000 * time.Millisecond)
	if c > 0 && time.Now().UnixNano()-c < interval {
		return false
	}

	return true
}

func (this *ChatManager) addChatTime(userId int) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.chatInterval[userId] = time.Now().UnixNano()
}

//发送聊天
func (this *ChatManager) ChatSendReq(user *objs.User, req *pb.ChatSendReq) (*pb.ChatSendAck, error) {

	if strings.HasPrefix(req.Msg, "GM#") {
		msg := this.GetGm().GmChatCode(user, req.Msg)
		if len(msg) > 0 {
			protoMsg := this.getChat(user, msg, req.Type, 0)
			this.GetUserManager().SendMessage(user, protoMsg, true)
		}
		return nil, nil
	}

	//禁言检查
	if ban, _ := this.GetUserManager().CheckBan(user, constUser.BAN_TYPE_CHAT); ban {
		logger.Info("玩家被禁言：%v", user.IdName())
		return nil, nil
	}

	msgType := int(req.Type)

	err := this.CheckChatSendCondition(user, msgType)
	if err != nil {
		return nil, err
	}

	//类型 长度检查
	logger.Debug("HandleChatSendReq msgType:%v, userId:%v", msgType, user.Id)
	msgLen := len([]rune(req.Msg))
	if msgLen < 1 || msgLen > gamedb.GetConf().Chat[1] {
		return nil, gamedb.ERRCHATLEN
	}
	if !pb.CHATTYPE_MAP[msgType] {
		return nil, gamedb.ERRCHATTYPE
	}

	//发言条件限制检查
	//chatConf := gamedb.GetChatClearCfg(msgType)
	//if chatConf == nil {
	//	return nil, gamedb.ERRCHATTYPE
	//}
	//canSpeak := this.GetCondition().CheckMulti(user, -1, chatConf.Conditon)
	//if !canSpeak {
	//	return nil, gamedb.ERRCHATCONDITION
	//}

	//发言间隔检查
	if !this.checkCanSpeak(user.Id) {
		return nil, gamedb.ERRCHATINTERVAL
	}

	//记录玩家发言时间，用于检查发言间隔
	this.addChatTime(user.Id)

	//敏感字替换 检查
	_, msg := gamedb.CensorAndReplace(req.Msg)
	protoMsg := this.getChat(user, msg, int32(msgType), 0)

	switch msgType {
	case pb.CHATTYPE_WORLD: //世界聊天

		this.BroadcastAll(protoMsg)

	case pb.CHATTYPE_GUILD: //门派
		err := this.GetGuild().BroadcastChatToGuildUsers(user, protoMsg)
		if err != nil {
			return nil, err
		}
	case pb.CHATTYPE_TEAM: //组队

	case pb.CHATTYPE_PRIVATE: //私聊
		msg = common.UnicodeEmojiCode(msg)
		if this.GetFriend().CheckFriendBlock(int(req.ToId), user.Id) {
			return nil, gamedb.ERRFRIENDBLOCK
		}
		this.GetFriend().WriteMsgLog(user, int(req.ToId), msg)
		protoMsg.ToId = req.ToId
		if this.GetUserManager().GetUser(int(req.ToId)) != nil {
			this.GetUserManager().SendMessageByUserId(int(req.ToId), protoMsg)
		}
		this.GetUserManager().SendMessage(user, protoMsg, true)
	}
	//聊天上报
	this.chatReport(user.Id, req)
	return &pb.ChatSendAck{IsBanSpeak: false}, nil
}

func (this *ChatManager) ChatSendSystemMsg(sysMsg string) {
	//敏感字替换 检查
	ntf := &pb.ChatMessageNtf{
		Type: pb.CHATTYPE_SYSTEM, //频道
		Msg:  sysMsg,             //消息内容
		Ts:   int32(time.Now().Unix()),
		ToId: 0,
	}

	this.BroadcastAll(ntf)
}

func (this *ChatManager) getChat(sender *objs.User, msg string, chatType, toUid int32) *pb.ChatMessageNtf {
	ntf := &pb.ChatMessageNtf{
		Sender: this.GetUserManager().BuilderBrieUserInfo(sender.Id), //发送者
		Type:   chatType,                                             //频道
		Msg:    msg,                                                  //消息内容
		Ts:     int32(time.Now().Unix()),
		ToId:   int32(toUid),
	}
	return ntf
}

func (this *ChatManager) chatReport(userId int, chatRep *pb.ChatSendReq) {

	if base.Conf.Sandbox {
		return
	}
	if len(this.chatReprotChan) < constChat.CHAT_REPORT_MAX_CHAN {
		this.chatReprotChan <- &chatReport{
			chat:   chatRep,
			sender: userId,
		}
	} else {
		logger.Error("聊天上报通道已满")
	}
}

func (this *ChatManager) chatReportLoop() {
	go func() {
		for {
			select {
			case c := <-this.chatReprotChan:
				sender := this.GetUserManager().GetUserBasicInfo(c.sender)
				var to *modelGame.UserBasicInfo
				if c.chat.ToId > 0 {
					to = this.GetUserManager().GetUserBasicInfo(int(c.chat.ToId))
				}
				ptsdk.GetSdk().ChatReport(base.Conf.ServerId, int(c.chat.Type), 0, c.chat.Msg, sender, to)
			}
		}
	}()
}

//聊天限制检查
func (this *ChatManager) CheckChatSendCondition(user *objs.User, types int) error {
	chatCfg := gamedb.GetChatClearCfg(types)
	if chatCfg == nil {
		return gamedb.ERRPARAM
	}
	isCanChatFlag := true
	if len(chatCfg.Conditon) > 0 {
		ok := this.GetCondition().CheckMultiByType(user, -1, chatCfg.Conditon, pb.CONDITIONTYPE_JUST_ONE)
		if !ok {
			isCanChatFlag = false
		}
	}

	if !isCanChatFlag {
		return gamedb.ERRCHATCONDITION
	}
	return nil
}

func (this *ChatManager) ChatForFightHelp(user *objs.User) error {

	content := this.GetAnnouncement().BuildContent(user, pb.SCROLINGTYPE_FIGHT_HELP, 0, 0)
	msg := this.getChat(user, content, pb.CHATTYPE_GUILD, 0)
	err := this.GetGuild().BroadcastChatToGuildUsers(user, msg)
	if err != nil {
		logger.Error("发送协助请求异常：%v", err)
		return err
	}
	return nil
}
