package worldLeader

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constAuction"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gamelibs/rmodel"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"encoding/json"
)

func (this *WorldLeader) WorldLeaderFightRankNtf(msg *pbserver.WorldLeaderFightRankNtf) {
	logger.Debug("推送世界首领排行  msg:%v", msg)
	this.setStageHp(int(msg.StageId), int(msg.BossHp))
	this.setStageRankInfos(int(msg.StageId), msg.Ranks)
	data, _ := json.Marshal(msg.Ranks)
	rmodel.WorldLeader.SetWorldLeaderRankInfo(int(msg.StageId), string(data))
	rmodel.WorldLeader.SetWorldLeaderHpInfo(int(msg.StageId), int(msg.BossHp))
	this.BroadcastAll(&pb.WorldLeaderBossHpNtf{StageId: msg.StageId, BossHp: msg.BossHp})
}

func (this *WorldLeader) EndWorldLeaderBossNtf(msg *pbserver.WorldLeaderFightEndNtf) {

	logger.Info("世界首领结算  msg:%v", msg)

	cfg := gamedb.GetWorldLeaderByStageId(int(msg.StageId))
	if cfg == nil {
		logger.Error("获取配置错误  stageId:%v", msg.StageId)
		return
	}
	for index, data := range msg.Ranks {
		guildInfo := this.GetGuild().GetGuildInfo(int(data.GuildId))
		if guildInfo == nil {
			//非本服拍卖行
			continue
		}
		canGetRewardUserIds := make(model.IntSlice, 0)
		for _, userId := range data.Users {
			canGetRewardUserIds = append(canGetRewardUserIds, int(userId))
		}

		logger.Debug("data.Users  :%v", data.Users)
		//结算推送
		for _, joinUserId := range data.Users {
			this.sendClientFightEnd(msg.StageId, joinUserId, data.GuildId, msg.LastAttacker, int32(index+1))
		}
		awardCfg := gamedb.GetWorldLeaderRewardByRank(int(msg.StageId), index+1)
		if awardCfg == nil {
			logger.Error("GetWorldLeaderRewardByRank nil stageId:%v  rank:%v", int(msg.StageId), index+1)
			continue
		}

		logger.Debug("世界首领 发放排名奖励 guildId:%v   cfg.LastDrop:%v   cfgId:%v  rank:%v", data.GuildId, awardCfg.Reward, awardCfg.Id, index+1)
		for _, item := range awardCfg.Reward {
			this.GetAuction().DropItemToGuildAuction(int(data.GuildId), item.ItemId, item.Count, constAuction.DropWorldLeader, canGetRewardUserIds)
		}
		this.GetAuction().BroadcastAuctionRedPointNtf(int(data.GuildId), pb.REDPOINTTYPE_OWN_GUILD_AUCTION_HAVE_ITEMS, pb.REDPOINTSTATE_BRIGHT)
	}

	lastAttacker := int(msg.LastAttacker)
	if lastAttacker > 0 {
		info := this.GetUserManager().BuilderBrieUserInfo(lastAttacker)
		//不是本服玩家
		if info == nil {
			return
		}
		//发送最后一击奖励
		this.GetMail().SendSystemMailWithItemInfos(lastAttacker, constMail.MAILTYPE_WORLD_LEADER, nil, cfg.LastDrop)
	}
	logger.Info("EndWorldLeaderBossNtf lastAttacker:%v",lastAttacker)
}
