package guildBonfire

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

//
//  EnterGuildBonfireFight
//  @Description: 进入篝火活动
//
func (this *GuildBonfireManager) EnterGuildBonfireFight(user *objs.User) error {

	if user.GuildData.NowGuildId <= 0 {
		return gamedb.ERRHAVENOGUILD
	}

	if !this.JudgeGuildBonfireIsOpen(user) {
		return gamedb.ERRGUILDBONFIREISNOTOPEN
	}

	err := this.GetFight().EnterGuildBonfire(user, user.GuildData.NowGuildId)
	if err != nil {
		return err
	}

	this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_GUILDBONFIRE_NUM, []int{1})
	this.GetDailyTask().CompletionOfTask(user, pb.DAILYTASKACTIVITYTYPE_GONG_HUI_GOU_HUO, 1)
	return nil
}

// 战斗回调
func (this *GuildBonfireManager) GuildBonfireFightResult(userIds []int) {

	//活动结束
	//this.StopAllUserAdd()
	logger.Info("公会篝火 战斗回调 userIds:%v", userIds)

	fightEnd := &pb.GuildBonfireFightNtf{}
	this.BroadcastAll(fightEnd)
	this.GetGuild().ResetGuildBonfireDonateInfo()
}
