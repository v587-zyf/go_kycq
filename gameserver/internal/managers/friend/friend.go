package friend

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	FRIEND_ADD_NUM   = 60
	SEARCH_NUM       = 20
	MSG_LOG_STR      = "%d[-]%s[-]%d" //userId,msg,time
	MSG_LOG_STR_TOO  = "%s{-}%s"
	FRIEND_APPLY_TOO = "%s{-}%s"
)

type FriendManager struct {
	util.DefaultModule
	managersI.IModule

	updateChan chan FriendChange
	loopChan   chan struct{}
}
type FriendChange struct {
	UserId   int
	FriendId int
}

func NewFriendManager(module managersI.IModule) *FriendManager {
	f := &FriendManager{IModule: module}
	return f
}

func (this *FriendManager) Init() error {
	this.updateChan = make(chan FriendChange, 1000)
	go this.updateService()
	return nil
}

func (this *FriendManager) Online(user *objs.User) {
	if user.Friend == nil {
		user.Friend = make(model.Friend)
	}
	userId := user.Id

	friendAddApply := rmodel.Friend.GetFriendApplyAdd(userId)
	if len([]rune(friendAddApply)) >= 1 {
		friendApplyArr := strings.Split(friendAddApply, "{-}")
		for _, fid := range friendApplyArr {
			this.Add(user, fid, &pb.FriendAddAck{})
		}
		rmodel.Friend.DelFriendApplyAdd(userId)
	}

	friendDelApply := rmodel.Friend.GetFriendApplyDel(userId)
	if len([]rune(friendDelApply)) >= 1 {
		friendApplyArr := strings.Split(friendDelApply, "{-}")
		for _, fid := range friendApplyArr {
			fId, _ := strconv.Atoi(fid)
			this.Del(user, fId)
		}
		rmodel.Friend.DelFriendApplyDel(userId)
	}

	taskMsgLog := rmodel.Friend.GetTaskMsgLog(base.Conf.ServerId, userId)
	if len([]rune(taskMsgLog)) >= 1 {
		userFriend := user.Friend
		msgStrArr := strings.Split(taskMsgLog, "{-}")
		timeNow := int(time.Now().Unix())
		for _, msgStr := range msgStrArr {
			msgLogArr := strings.Split(msgStr, "[-]")
			friendId, _ := strconv.Atoi(msgLogArr[0])
			if this.CheckFriendBlock(userId, friendId) {
				continue
			}
			if userFriend[friendId] == nil {
				userFriend[friendId] = &model.FriendUnit{
					MsgLog:    make(model.MsgLogs),
					CreatedAt: timeNow,
					DeletedAt: timeNow,
				}
			}
			time, _ := strconv.Atoi(msgLogArr[2])
			userFriend[friendId].MsgLog[time] = &model.MsgLog{
				Msg:  msgLogArr[1],
				Time: time,
				IsMy: false,
			}
			userFriend[friendId].IsRead = true
		}
		rmodel.Friend.DelTaskMsgLog(base.Conf.ServerId, userId)
	}
}

/**
 *  @Description: 获取好友私聊信息
 *  @param user
 *  @param friendId	好友id
 *  @param ack
 *  @return error
 */
func (this *FriendManager) FriendMsg(user *objs.User, friendId int, ack *pb.FriendMsgAck) error {
	if friendId < 1 {
		return gamedb.ERRPARAM
	}
	userFriend := user.Friend
	if userFriend[friendId] == nil {
		return gamedb.ERRFRIENDNOTADD
	}
	ack.FriendId = int32(friendId)
	ack.MsgLog = builder.BuildFriendMsgLog(userFriend[friendId].MsgLog, this.GetMsgNotShowTime())
	return nil
}

func (this *FriendManager) GetMsgNotShowTime() int {
	notShowMonth := gamedb.GetConf().ChatRecord
	return int(common.GetZeroClockTimestamp(time.Now().AddDate(0, -notShowMonth, 0)))
}

func (this *FriendManager) BuildPbFriendInfo(friendId int, msgLogs model.MsgLogs) *pb.FriendInfo {
	notShowTime := this.GetMsgNotShowTime()
	isOnline, outTime := this.GetUserManager().GetUserOnlineStatus(friendId)
	var lastMsg *pb.MsgLog
	if len(msgLogs) > 0 {
		maxTime := math.MinInt32
		for sendTime, msg := range msgLogs {
			if maxTime < sendTime && notShowTime < sendTime {
				maxTime = sendTime
				lastMsg = builder.BuildFriendMsg(msg)
			}
		}
	}
	return &pb.FriendInfo{
		UserInfo: this.GetUserManager().BuilderBrieUserInfo(friendId),
		IsOnline: isOnline,
		OutTime:  outTime,
		//MsgLog:   builder.BuildFriendMsgLog(msgLogs, int(notShowTime)),
		//IsRead:    false,
		LastMsg: lastMsg,
	}
}

/**
 *  @Description: 是否把该用户拉黑
 *  @param userId
 *  @param friendId
 *  @return bool
 */
func (this *FriendManager) CheckFriendBlock(userId, friendId int) bool {
	var friend model.Friend
	if user := this.GetUserManager().GetUser(userId); user != nil {
		friend = user.Friend
	} else {
		friend, _ = modelGame.GetUserModel().GetFriendInfoByUserId(userId)
	}
	if friend == nil {
		return false
	}
	if f, ok := friend[friendId]; ok {
		if f.BlockTime != 0 {
			return true
		}
	}
	return false
}

func (this *FriendManager) GetFriendNum(user *objs.User) int {
	userFriend := user.Friend
	num := 0
	for _, friend := range userFriend {
		if friend.BlockTime != 0 || friend.DeletedAt != 0 {
			continue
		}
		num++
	}
	return num
}

func (this *FriendManager) MakeFriendDbDate() *model.FriendUnit {
	timeNow := time.Now().Unix()
	return &model.FriendUnit{
		//MsgLog:    make(model.MsgLogs),
		CreatedAt: int(timeNow),
	}
}
