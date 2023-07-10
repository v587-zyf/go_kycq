package guardPillar

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constGuild"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"strconv"
	"time"
)

func NewGuardPillarManager(m managersI.IModule) *GuardPillar {
	return &GuardPillar{
		IModule: m,
	}
}

type GuardPillar struct {
	util.DefaultModule
	managersI.IModule
}

/**
 *  @Description: 守卫龙柱进入
 *  @param user
 *  @param stageId
 *  @return error
 */
func (this *GuardPillar) In(user *objs.User, stageId int) error {
	if err := this.checkOpen(user); err != nil {
		return err
	}
	if err := this.GetFight().EnterGuardPillarFight(user); err != nil {
		return err
	}
	this.GetGuild().SetGuildActivityInfo(user.GuildData.NowGuildId, constGuild.GUILD_ACTIVITY_GUARD_PILLAR, constGuild.GUILD_ACTIVITY_OPEN)
	return nil
}

/**
 *  @Description: 守卫龙柱结算
 *  @param user
 *  @param rounds	波数
 *  @param rank		伤害排名
 */
func (this *GuardPillar) GuardPillarResult(userId, stageId, rounds, rank int) {
	roundGoods := make(map[int]int)
	rankGoods := make(map[int]int)
	//波数奖励
	items := make([]*model.Item, 0)
	roundsCfg := gamedb.GetGuardRoundsGuardRoundsCfg(rounds)
	if roundsCfg != nil {
		for _, info := range roundsCfg.Reward {
			items = append(items, &model.Item{ItemId: info.ItemId, Count: info.Count})
			roundGoods[info.ItemId] += info.Count
		}
		this.GetMail().SendSystemMail(userId, constMail.MAILTYPE_GUARDPILLAR_ROUNDS, []string{strconv.Itoa(rounds)}, items, 0)
	}
	//伤害排行奖励
	rankCfgs := gamedb.GetGuardRankCfgs()
	var rankReward gamedb.ItemInfos
	for _, cfg := range rankCfgs {
		if rank <= cfg.Rank[1] && rank >= cfg.Rank[0] {
			rankReward = cfg.Reward
			break
		}
	}
	if rankReward != nil {
		items = make([]*model.Item, 0)
		for _, info := range rankReward {
			items = append(items, &model.Item{ItemId: info.ItemId, Count: info.Count})
			rankGoods[info.ItemId] += info.Count
		}
		this.GetMail().SendSystemMail(userId, constMail.MAILTYPE_GUARDPILLAR_RANK, []string{strconv.Itoa(rank)}, items, 0)
	}
	guildId := 0
	if user := this.GetUserManager().GetUser(userId); user != nil {
		ntf := &pb.GuardPillarResultNtf{
			StageId:    int32(stageId),
			Rounds:     int32(rounds),
			Rank:       int32(rank),
			RoundGoods: ophelper.CreateGoodsChangeNtf(roundGoods),
			RankGoods:  ophelper.CreateGoodsChangeNtf(rankGoods),
		}
		this.GetUserManager().SendMessage(user, ntf, true)
		guildId = user.GuildData.NowGuildId
	} else {
		userInfo := this.GetUserManager().GetUserBasicInfo(userId)
		guildId = userInfo.GuildData.NowGuildId
	}
	this.GetGuild().SetGuildActivityInfo(guildId, constGuild.GUILD_ACTIVITY_GUARD_PILLAR, constGuild.GUILD_ACTIVITY_CLOSE)
}

func (this *GuardPillar) checkOpen(user *objs.User) error {
	guildActivityId := constGuild.GUILD_ACTIVITY_GUARD_PILLAR
	activityCfg := gamedb.GetGuildActivityGuildActivityCfg(guildActivityId)
	if activityCfg == nil {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("guildActivity id = %v", guildActivityId)
	}

	//如果没合服,判断开服时间,合过服判断合服时间
	if !this.GetSystem().IsMerge() {
		serverOpenDays := this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
		if serverOpenDays < activityCfg.OpenDayMin || serverOpenDays > activityCfg.OpenDayMax {
			return gamedb.ERRACTIVITYNOTOPEN
		}
	} else {
		serverMergeDays := this.GetSystem().GetServerMergeDayByServerId(user.ServerId)
		if serverMergeDays < activityCfg.MergeDayMin || serverMergeDays > activityCfg.MergeDayMax {
			return gamedb.ERRACTIVITYNOTOPEN
		}
	}

	oTime, eTime := gamedb.GetActiveTime(activityCfg.OpenTime, activityCfg.CloseTime, activityCfg.Week)
	openTime := time.Unix(int64(oTime), 0)
	endTime := time.Unix(int64(eTime), 0)

	//校验公会
	guildId := user.GuildData.NowGuildId
	guildInfo := this.GetGuild().GetGuildInfo(guildId)
	if guildInfo == nil {
		return gamedb.ERRHAVENOGUILD
	}
	//校验公会创建时间
	//if guildInfo.CreatedAt.Unix() >= openTime.Unix() {
	//	logger.Debug("公会创建在守护龙柱开放后")
	//	return gamedb.ERRNOTTOTIME
	//}
	//是否已开启过
	timeNow := time.Now()
	guardPillarEndTime := this.GetFight().GetGuardPillarFightEndTime(guildId)
	hasPower := this.GetGuild().CheckActivityOpenPower(user.GuildData.Position, guildActivityId)
	if guardPillarEndTime != 0 {
		gTime := time.Unix(int64(guardPillarEndTime), 0)
		if gTime.Day() == openTime.Day() {
			logger.Debug("守卫龙柱已开启 进入检查是否重复开启")
			//if timeNow.Hour() < gTime.Hour() && timeNow.Minute() < gTime.Minute() && timeNow.Second() < gTime.Second() {
			if this.GetGuild().GetGuildActivityInfo(guildId, constGuild.GUILD_ACTIVITY_GUARD_PILLAR) {
				return gamedb.ERRREPEATOPEN
			}
			//}
		} else {
			logger.Debug("1===守卫龙柱未开启 检查开启时间和用户权限")
			if timeNow.Unix() > endTime.Unix() || timeNow.Unix() < openTime.Unix() {
				return gamedb.ERRACTIVITYNOTOPEN
			}
			if !hasPower {
				return gamedb.ERRNOPOWER
			}
		}
	} else {
		logger.Debug("2===守卫龙柱未开启 检查开启时间和用户权限")
		if timeNow.Unix() > endTime.Unix() || timeNow.Unix() < openTime.Unix() {
			return gamedb.ERRACTIVITYNOTOPEN
		}
		if !hasPower {
			return gamedb.ERRNOPOWER
		}
	}
	return nil
}
