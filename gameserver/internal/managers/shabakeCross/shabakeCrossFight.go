package shabakeCross

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"strconv"
)

func NewCrossShaBaKeManager(m managersI.IModule) *CrossShaBaKe {
	return &CrossShaBaKe{
		IModule: m,
	}
}

type CrossShaBaKe struct {
	util.DefaultModule
	managersI.IModule
}

func (this *CrossShaBaKe) EnterCrossShaBakeFight(user *objs.User) error {
	if user.GuildData.NowGuildId <= 0 {
		return gamedb.ERRHAVENOGUILD
	}

	err := this.JudgeCrossIsOpen(user)
	if err != nil {
		return err
	}
	err = this.GetFight().EnterShabakeCrossFight(user)
	if err != nil {
		return err
	}
	return nil
}

//  CrossShaBakeFightEndNtf
//  @Description:结算回调
func (this *CrossShaBaKe) CrossShaBakeFightEndNtf(msg *pbserver.ShabakeCrossFightEndNtf) {
	logger.Debug("跨服沙巴克结算 msg.ServerRank:%v  msg.GuildRank:%v", msg.ServerRank, msg.GuildRank)

	if rmodel.Shabake.GetCrossShaBakeIsEnd(base.Conf.ServerId) == 1 {
		logger.Info("跨服沙巴沙巴克活动已发过奖励")
		return
	}

	firstGuildId := 0
	winUserIds := make([]int, 0)
	firstGuildHuiZID := 0
	serverRank := make(map[int]int)
	needKyEvent := false
	serverMaxScoreByGuild := make(map[int]int)
	serverMaxScore := make(map[int]int)
	for _, data := range msg.GuildRank {
		if int(data.Score) > serverMaxScore[int(data.ServerId)] {
			serverMaxScore[int(data.ServerId)] = int(data.Score)
			serverMaxScoreByGuild[int(data.ServerId)] = int(data.GuildId)
		}
	}

	for serverId, guildId := range serverMaxScoreByGuild {
		logger.Debug("serverId:%v,guildId:%v", serverId, guildId)
	}

	//区服积分结算
	ntf := &pb.CrossShaBaKeFightEndNtf{}
	for index, data := range msg.ServerRank {
		rank := index + 1
		if rank == 1 && int(data.Id) == base.Conf.ServerId {
			needKyEvent = true
			firstGuildId = serverMaxScoreByGuild[int(data.Id)]
			users := this.GetGuild().GetGuildHuiAndFuHuiUserIds(firstGuildId)
			for i, l := 0, len(users); i < l; i += 2 {
				if users[i+1] == pb.GUILDPOSITION_HUIZHANG {
					firstGuildHuiZID = users[i]
				}
				winUserIds = append(winUserIds, users[i])
			}
			this.SendFirstGuildInfoToCcs(firstGuildId,0)
		}
		serverRank[int(data.Id)] = rank
		ntf.ServerRank = append(ntf.ServerRank, &pb.ShabakeRankScore{Id: data.Id, Score: data.Score})
	}

	//门派积分结算
	for index, data := range msg.GuildRank {
		rank := index + 1
		guildInfo := this.GetGuild().GetGuildInfo(int(data.GuildId))
		if guildInfo == nil {
			continue
		}

		cfg := gamedb.GetCrossShaBakePerAwardByRank(serverRank[int(data.ServerId)])
		cfg1 := gamedb.GetCrossShaBakeUniAwardByRank(rank)
		for i, j := 0, len(guildInfo.Positions); i < j; i += 2 {
			userId := guildInfo.Positions[i]
			//个人奖励
			if cfg != nil {
				_ = this.GetMail().SendSystemMailWithItemInfos(userId, constMail.MAILTYPE_CROSS_SHABAKE_PER_REWARD, []string{strconv.Itoa(serverRank[int(data.ServerId)])}, cfg.Reward)
			}
			//公会奖励
			if cfg1 != nil {
				_ = this.GetMail().SendSystemMailWithItemInfos(userId, constMail.MAILTYPE_CROSS_SHABAKE_GUILD_REWARD, []string{strconv.Itoa(rank)}, cfg1.Reward)
			}
			logger.Debug("跨服沙巴克结算 userId:%v serverId:%v 区服排名:%v  门派id:%v 门派排名:%v", userId, data.ServerId, serverRank[int(data.ServerId)], data.GuildId, rank)
		}
	}
	if needKyEvent {
		userInfo := this.GetUserManager().GetAllUserInfoIncludeOfflineUser(firstGuildHuiZID)
		if userInfo != nil {
			kyEvent.CrossShaBake(userInfo, firstGuildId, winUserIds, serverRank)
		}
	}
	this.BroadcastAll(ntf)
	_ = this.CrossShaBakeOpenOrCloseNtf(false)
	rmodel.Shabake.SetCrossShaBakeIsEnd(base.Conf.ServerId, 1)
	return
}
