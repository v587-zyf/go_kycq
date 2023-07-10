package friend

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
)

/**
 *  @Description: 获取列表信息
 *  @param user
 *  @param block	是否是黑名单列表
 *  @return []*pb.FriendInfo
 */
func (this *FriendManager) List(user *objs.User, block bool) []*pb.FriendInfo {
	userFriend := user.Friend
	pbFriendInfo := make([]*pb.FriendInfo, 0)
	for friendId, friend := range userFriend {
		if friendId <= 0{
			continue
		}
		if block {
			if friend.BlockTime == 0 {
				continue
			}
		} else {
			if friend.DeletedAt != 0 || friend.BlockTime != 0 {
				continue
			}
		}
		friendInfo := this.BuildPbFriendInfo(friendId, friend.MsgLog)
		friendInfo.IsRead = friend.IsRead
		pbFriendInfo = append(pbFriendInfo, friendInfo)
	}

	return pbFriendInfo
}

/**
 *  @Description: 添加好友
 *  @param user
 *  @param fId
 *  @return error
 */
func (this *FriendManager) Add(user *objs.User, friendIdStr string, ack *pb.FriendAddAck) error {
	userId := user.Id
	addFriendMaxNum := gamedb.GetConf().FriendsMaxNum
	hasFriendNum := this.GetFriendNum(user)
	if hasFriendNum >= addFriendMaxNum {
		return gamedb.ERRFRIENDADDENOUGH
	}
	if len([]rune(friendIdStr)) < 1 {
		return gamedb.ERRPARAM
	}
	friendIdArr := strings.Split(friendIdStr, ",")
	pbFriendInfo := make([]*pb.FriendInfo, 0)
	for _, friendId := range friendIdArr {
		fId, _ := strconv.Atoi(friendId)
		if fId == userId || fId <= 0 {
			continue
		}
		friend, ok := user.Friend[fId]
		if ok {
			if friend.DeletedAt != 0 || friend.BlockTime != 0 {
				friend.DeletedAt = 0
				friend.BlockTime = 0
			} else {
				continue
			}
		} else {
			user.Friend[fId] = this.MakeFriendDbDate()
		}
		pbFriendInfo = append(pbFriendInfo, this.BuildPbFriendInfo(fId, nil))
		hasFriendNum += 1
		if hasFriendNum >= addFriendMaxNum {
			break
		}
	}
	if len(pbFriendInfo) == 0 {
		return gamedb.ERRFRIENDADD
	}
	user.Dirty = true

	ack.FriendInfo = pbFriendInfo
	return nil
}

/**
 *  @Description: 删除好友
 *  @param user
 *  @param fId
 *  @return error
 */
func (this *FriendManager) Del(user *objs.User, fId int) error {
	userId := user.Id
	if fId == userId {
		return gamedb.ERRPARAM
	}
	friend, ok := user.Friend[fId]
	if !ok {
		return gamedb.ERRFRIENDNOTADD
	}
	friend.DeletedAt = int(time.Now().Unix())
	this.friendDelUser(fId, userId)
	user.Dirty = true
	return nil
}

/**
 *  @Description: 添加黑名单
 *  @param user
 *  @param fId
 *  @return error
 */
func (this *FriendManager) BlockAdd(user *objs.User, fId int) error {
	userId := user.Id
	if fId == userId {
		return gamedb.ERRPARAM
	}
	blockFriendMaxNum := gamedb.GetConf().BlackListMaxNum
	if len(this.List(user, true)) >= blockFriendMaxNum {
		return gamedb.ERRFRIENDBLOCKMAX
	}
	friend, ok := user.Friend[fId]
	timeNow := int(time.Now().Unix())
	if !ok {
		friendDbDate := this.MakeFriendDbDate()
		friendDbDate.BlockTime = timeNow
		user.Friend[fId] = friendDbDate
	} else {
		if friend.BlockTime != 0 {
			return gamedb.ERRFRIENDBLOCKADD
		}
		friend.BlockTime = timeNow
	}
	user.Dirty = true
	return nil
}

/**
 *  @Description: 删除黑名单
 *  @param user
 *  @param fId	删除所有传-1
 *  @return error
 */
func (this *FriendManager) BlockDel(user *objs.User, fId int) error {
	userId := user.Id
	if fId == userId {
		return gamedb.ERRPARAM
	}
	timeNow := int(time.Now().Unix())
	if fId == -1 {
		for fid, friend := range user.Friend {
			if friend.BlockTime == 0 {
				continue
			}
			friend.BlockTime = 0
			friend.DeletedAt = timeNow
			this.friendDelUser(fid, userId)
		}
	} else {
		friend, ok := user.Friend[fId]
		if !ok {
			return gamedb.ERRFRIENDNOTBLOCKADD
		}
		friend.BlockTime = 0
		friend.DeletedAt = timeNow
		this.friendDelUser(fId, userId)
	}
	user.Dirty = true
	return nil
}

/**
 *  @Description: 搜索用户
 *  @param user
 *  @param name	搜索名称（不填为推荐好友）
 *  @param ack
 *  @return error
 */
func (this *FriendManager) Search(user *objs.User, name string, ack *pb.FriendSearchAck) error {
	userId := user.Id
	pbFriendInfo := make([]*pb.FriendInfo, 0)
	if len([]rune(name)) <= 0 {
		hasMap := make(map[int]int)
		for friendId, friend := range user.Friend {
			if friend.DeletedAt != 0 || friend.BlockTime != 0 {
				continue
			}
			hasMap[friendId] = 0
		}
		hasMap[userId] = 0
		randUserId := this.GetUserManager().RandGetUserId(SEARCH_NUM, hasMap)
		for rUserId := range randUserId {
			if this.GetUserManager().GetUserBasicInfo(rUserId).HeroDisplay[constUser.USER_HERO_MAIN_INDEX] != nil {
				pbFriendInfo = append(pbFriendInfo, this.BuildPbFriendInfo(rUserId, nil))
			}
		}
	} else {
		searchSlice, err := modelGame.GetUserModel().SearchName(name)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		for _, userBasicInfo := range searchSlice {
			pbFriendInfo = append(pbFriendInfo, this.BuildPbFriendInfo(userBasicInfo.Id, nil))
		}
	}
	ack.FriendList = pbFriendInfo
	return nil
}

/**
 *  @Description: 私聊信息
 *  @param userId
 *  @param friendId
 *  @param msg
 */
func (this *FriendManager) WriteMsgLog(user *objs.User, friendId int, msg string) {
	userId := user.Id
	timeNow := int(time.Now().Unix())
	if friend, ok := user.Friend[friendId]; ok {
		if friend.MsgLog == nil {
			friend.MsgLog = make(model.MsgLogs)
		}
		friend.MsgLog[timeNow] = &model.MsgLog{
			Msg:  msg,
			Time: timeNow,
			IsMy: true,
		}
		user.Dirty = true
	}
	if friendUser := this.GetUserManager().GetUser(friendId); friendUser != nil {
		if this.CheckFriendBlock(friendUser.Id, userId) {
			return
		}
		friend, ok := friendUser.Friend[userId]
		if !ok {
			friendDbDate := this.MakeFriendDbDate()
			friendDbDate.DeletedAt = timeNow
			friendDbDate.MsgLog = make(model.MsgLogs)
			friendDbDate.MsgLog[timeNow] = &model.MsgLog{
				Msg:  msg,
				Time: timeNow,
				IsMy: false,
			}
			friendUser.Friend[userId] = friendDbDate
			friend = friendUser.Friend[userId]
		} else {
			if friend.MsgLog == nil {
				friend.MsgLog = make(model.MsgLogs)
			}
			friend.MsgLog[timeNow] = &model.MsgLog{
				Msg:  msg,
				Time: timeNow,
				IsMy: false,
			}
		}
		friend.IsRead = true
		friendUser.Dirty = true
	} else {
		msgStr := fmt.Sprintf(MSG_LOG_STR, userId, msg, timeNow)
		taskMsgLog := rmodel.Friend.GetTaskMsgLog(base.Conf.ServerId, friendId)
		if len([]rune(taskMsgLog)) > 0 {
			msgStr = fmt.Sprintf(MSG_LOG_STR_TOO, msgStr, taskMsgLog)
		}
		rmodel.Friend.SetTaskMsgLog(base.Conf.ServerId, friendId, msgStr)
	}
}

/**
 *  @Description: 阅读消息
 *  @param user
 *  @param friendId	好友id
 *  @return error
 */
func (this *FriendManager) ReadMsg(user *objs.User, friendId int) error {
	friend, ok := user.Friend[friendId]
	if !ok {
		return gamedb.ERRFRIENDNOTADD
	}
	friend.IsRead = false
	user.Dirty = true
	return nil
}
