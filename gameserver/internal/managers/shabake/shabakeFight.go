package shabake

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constAuction"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"strconv"
)

func (this *ShabakeManager) EnterShabakeFight(user *objs.User) error {
	if user.GuildData.NowGuildId <= 0 {
		return gamedb.ERRHAVENOGUILD
	}

	err := this.JudgeIsOpen(user)
	if err != nil {
		return err
	}
	err = this.GetFight().EnterShabakeFight(user)
	if err != nil {
		return err
	}

	if this.GetRank().GetCount(pb.RANKTYPE_SHABAKE_GUILD) > 0 {
		this.GetRank().DelData(pb.RANKTYPE_SHABAKE_GUILD)
	}

	this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_SHABAKE_NUM, []int{1})
	return nil
}

//  WorldBossFightEndAck
//  @Description:战斗之后，回调结果
//  @receiver this
//  @param rank 个人排名
//  @param guildRank 门派排名
//  @return error
//
func (this *ShabakeManager) ShabakeFightEndAck(rank, guildRank, otherUser []int) error {

	if rmodel.Shabake.GetShaBakeIsEnd(base.Conf.ServerId) == 1 {
		logger.Info("沙巴克活动已发过奖励")
		return nil
	}

	guildRankMap := make(map[int]int)
	for k, v := range guildRank {
		guildRankMap[v] = k + 1
	}

	//个人奖励处理
	for index, userId := range rank {

		baseUserInfo := this.GetUserManager().GetUserBasicInfo(userId)
		if baseUserInfo == nil {
			continue
		}
		ranks := index + 1
		rewards := gamedb.GetShabakePerRewardByRank(ranks)
		if rewards == nil {
			logger.Error("GetShabakePerRewardByRank nil ranks:%v", ranks)
			continue
		}

		// 发送邮件
		args := []string{strconv.Itoa(ranks)}
		err := this.GetMail().SendSystemMailWithItemInfos(userId, constMail.SHABAKE_PER_REWARD, args, rewards)
		if err != nil {
			logger.Error("ShabakeFightEndAck send rank reward error err is %v", err)
			continue
		}
		// 给在线玩家推送消息
		if user := this.GetUserManager().GetUser(userId); user != nil {
			userGuild := user.GuildData.NowGuildId
			guildRank := guildRankMap[userGuild]
			this.GetUserManager().SendMessage(user, &pb.ShaBaKeFightResultNtf{
				Rank:     int32(guildRank),
				UserRank: int32(ranks),
			}, true)
		}
	}
	//没有积分玩家
	for _, v := range otherUser {
		if user := this.GetUserManager().GetUser(v); user != nil {
			userGuild := user.GuildData.NowGuildId
			guildRank := guildRankMap[userGuild]
			this.GetUserManager().SendMessage(user, &pb.ShaBaKeFightResultNtf{
				Rank: int32(guildRank),
			}, true)
		}
	}

	//第一门派处理
	if len(guildRank) > 0 {
		firstGuild := guildRank[0] //第一门派id
		guildInfo := this.GetGuild().GetGuildInfo(firstGuild)
		if guildInfo != nil {
			this.dropItemsToGuildAuction(firstGuild)
			this.sendMailToHuiAndFuHui(firstGuild)
			this.GetShaBaKeCross().SendFirstGuildInfoToCcs(firstGuild, 1)
		}

		//沙巴克公会排行
		for rank, guildId := range guildRank {
			this.GetRank().Append(pb.RANKTYPE_SHABAKE_GUILD, guildId, rank, false, false, false)
		}
	}

	rmodel.Shabake.SetShaBakeIsEnd(base.Conf.ServerId, 1)
	this.ShabakeOpenOrCloseNtf(false)
	return nil
}

//掉落物品进公会拍卖行
func (this *ShabakeManager) dropItemsToGuildAuction(firstGuild int) {

	//第一门派公会拍卖行掉落物品
	dropItems := gamedb.GetConf().ShabakeReward3
	_, _, _, guildMember := this.GetGuild().GetGuildMemberInfo(firstGuild)
	for _, data := range dropItems {
		this.GetAuction().DropItemToGuildAuction(firstGuild, data.ItemId, data.Count, constAuction.DropShaBake, guildMember)
	}
	this.GetAuction().BroadcastAuctionRedPointNtf(firstGuild, pb.REDPOINTTYPE_OWN_GUILD_AUCTION_HAVE_ITEMS, pb.REDPOINTSTATE_BRIGHT)
	return
}

func (this *ShabakeManager) sendMailToHuiAndFuHui(guildId int) {

	users := this.GetGuild().GetGuildHuiAndFuHuiUserIds(guildId)
	if users == nil || len(users) <= 0 {
		return
	}

	for i, l := 0, len(users); i < l; i += 2 {
		if users[i+1] == pb.GUILDPOSITION_HUIZHANG {
			huiReward := gamedb.GetConf().ShabakeReward4
			this.GetMail().SendSystemMailWithItemInfos(users[i], constMail.MAILTYPE_SHABAKE_HUIZHANG_OTHER_REWARD, []string{}, huiReward)
		}
		if users[i+1] == pb.GUILDPOSITION_FUHUIZHANG {
			fuHuiReward := gamedb.GetConf().ShabakeReward5
			this.GetMail().SendSystemMailWithItemInfos(users[i], constMail.MAILTYPE_SHABAKE_FUHUIZHANG_OTHER_REWARD, []string{}, fuHuiReward)
		}
	}
	return
}
