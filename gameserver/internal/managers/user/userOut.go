package user

import (
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"time"
)

func (this *UserManager) RemoveUser(userId int) {
	this.usersMu.Lock()
	if user, ok := this.users[userId]; ok {
		this.AddLocalServerInfo(user)
		delete(this.users, userId)
		delete(this.usersByOpenId[user.ServerId], user.OpenId)
		logger.Info("usermanager user remove,userId:%v,openId:%v,serverid:%v", user.Id, user.OpenId, user.ServerId)
	} else {
		logger.Error("usermanagerErr remove userId:%v", userId)
	}
	this.usersMu.Unlock()
}

func (this *UserManager) UserDisconnect(user *objs.User) {
	logger.Info("UserDisconnect userId:%v,nickName:%v,openID:%v", user.Id, user.NickName, user.OpenId)
	//this.GetOffline().AutoGetAward(user)
	this.GetMonthCard().CheckExpire(user)
	this.GetOnline().OffLine(user, true)
	this.GetFight().LeaveFight(user, constFight.LEAVE_FIGHT_TYPE_OFFLINE)
	this.saveInDB(user)
	this.GetFriend().WriteFriendUserInfo(user)
	this.GetTlog().PlayerLogout(user)
	this.RemoveUser(user.Id)
	this.SyncUser(user, constUser.USER_STATUS_OFFLINE)
	this.BroadcastAll(&pb.UserOffLineNtf{UserId: int32(user.Id), OffLintTime: time.Now().Unix()})
	kyEvent.UserExit(user)
}

func (this *UserManager) KickUserWithMsg(user *objs.User, reason string) error {
	if user == nil {
		return nil
	}
	logger.Info("踢出玩家：%v-%v,reason:%v", user.Id, user.NickName, reason)
	//ntf := builder.BuildKickUserNtf(reason)
	//this.SendMessage(user, ntf, true)
	//time.Sleep(time.Microsecond * 3)
	user.CloseConn(reason)

	return nil
}
