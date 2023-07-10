package guild

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constGuild"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"time"
)

func (this *GuildManager) SetGuildActivityInfo(guildId, activityId, status int) {
	this.Lock()
	defer this.Unlock()
	if this.guildActivity[guildId] == nil {
		this.guildActivity[guildId] = make(map[int]int)
	}
	if this.guildActivity[guildId][activityId] != status {
		this.guildActivity[guildId][activityId] = status
	}
}

func (this *GuildManager) GetGuildActivityInfo(guildId, activityId int) bool {
	this.Lock()
	defer this.Unlock()
	isClose := false
	if guildActivityInfo, ok := this.guildActivity[guildId]; ok {
		guardPillarEndTime := this.GetFight().GetGuardPillarFightEndTime(guildId)
		if guardPillarEndTime != 0 {
			timeNow := time.Now()
			gTime := time.Unix(int64(guardPillarEndTime), 0)
			if gTime.Day() == timeNow.Day() && guildActivityInfo[activityId] == constGuild.GUILD_ACTIVITY_CLOSE {
				isClose = true
			}
		}
	}
	return isClose
}

/**
 *  @Description: 获取公会活动信息
 *  @param activityId
 *  @param ack
 */
func (this *GuildManager) GuildActivityLoad(user *objs.User, activityId int, ack *pb.GuildActivityLoadAck) error {
	endTime := 0
	switch activityId {
	case constGuild.GUILD_ACTIVITY_GUARD_PILLAR:
		endTime = this.GetFight().GetGuardPillarFightEndTime(user.GuildData.NowGuildId)
	default:
		return gamedb.ERRPARAM
	}
	ack.EndTime = int64(endTime)
	ack.GuildActivityId = int32(activityId)
	ack.IsClose = this.GetGuildActivityInfo(user.GuildData.NowGuildId, activityId)
	return nil
}

/**
 *  @Description: 给公会所有人发送消息
 *  @param guildId			公会id
 *  @param msg				消息
 *  @param notSendUserIds	过滤掉不发送的人
 */
func (this *GuildManager) SendMsgToAllUser(guildId int, msg nw.ProtoMessage, notSendUserIds []int) {
	guildInfo := this.GetGuild().GetGuildInfo(guildId)
	if guildInfo == nil {
		logger.Error("公会推送消息错误 未找到公会 guildId:%v", guildId)
		return
	}

	for i, j := 0, len(guildInfo.Positions); i < j; i += 2 {
		userId := guildInfo.Positions[i]
		isSend := true
		for _, notSendUid := range notSendUserIds {
			if userId == notSendUid {
				isSend = false
				break
			}
		}
		if isSend {
			guildUser := this.GetUserManager().GetUser(userId)
			if guildUser != nil {
				this.GetUserManager().SendMessage(guildUser, msg, true)
			}
		}
	}
}

func (this *GuildManager) CheckActiveOpen(user *objs.User, guildActivityId int) (error, *gamedb.GuildActivityGuildActivityCfg) {
	activityCfg := gamedb.GetGuildActivityGuildActivityCfg(guildActivityId)
	if activityCfg == nil {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("guildActivity id = %v", guildActivityId), activityCfg
	}

	//如果没合服,判断开服时间,合过服判断合服时间
	if !this.GetSystem().IsMerge() {
		serverOpenDays := this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
		if serverOpenDays < activityCfg.OpenDayMin || serverOpenDays > activityCfg.OpenDayMax {
			return gamedb.ERRACTIVITYNOTOPEN, activityCfg
		}
	} else {
		serverMergeDays := this.GetSystem().GetServerMergeDayByServerId(user.ServerId)
		if serverMergeDays < activityCfg.MergeDayMin || serverMergeDays > activityCfg.MergeDayMax {
			return gamedb.ERRACTIVITYNOTOPEN, activityCfg
		}
	}

	nowTime := time.Now()
	openTime, closeTime := gamedb.GetActiveTime(activityCfg.OpenTime, activityCfg.CloseTime, activityCfg.Week)
	if int(nowTime.Unix()) < openTime {
		return gamedb.ERRACTIVITYNOTOPEN, activityCfg
	}
	if int(nowTime.Unix()) > closeTime {
		return gamedb.ERRACTIVITYCLOSE, activityCfg
	}
	//进入条件
	if check := this.GetCondition().CheckMulti(user, -1, activityCfg.Condition); !check {
		return gamedb.ERRCONDITION, activityCfg
	}
	return nil, activityCfg
}

func (this *GuildManager) CheckActivityOpenPower(position, guildActivityId int) bool {
	hasPower := false
	guildCfg := gamedb.GetGuildGuildCfg(position)
	if guildCfg != nil {
		for _, activityId := range guildCfg.Activity {
			if guildActivityId == activityId {
				hasPower = true
				break
			}
		}
	}
	return hasPower
}
