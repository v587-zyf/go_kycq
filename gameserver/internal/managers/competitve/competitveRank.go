package competitve

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"math/rand"
	"time"
)

//获取赛季排名
func (this *CompetitveManager) GetSeasonRankInfo(seasonRankInfos []float64) []*pb.CompetitveRankInfo {

	seasonRank := make([]*pb.CompetitveRankInfo, 0)

	rank := 0
	for i, j := 0, len(seasonRankInfos); i < j; i += 2 {
		rank++
		if seasonRankInfos[i] > 0 {
			userId := int(seasonRankInfos[i])
			score := int(seasonRankInfos[i+1])
			userInfo := this.GetUserManager().GetUserBasicInfo(userId)
			if userInfo != nil {
				seasonRank = append(seasonRank, &pb.CompetitveRankInfo{Ranking: int32(rank), Avatar: string(userInfo.Avatar), Score: int32(score), NickName: userInfo.NickName})
			}
		}
	}
	return seasonRank
}

//发放赛季排名奖励
func (this *CompetitveManager) SendSeasonEndReward() {
	if !this.JudgeSeasonIsOver() {
		logger.Info("竞技场赛季还未结束")
		return
	}
	lastSeason, _, _ := this.GetCurrentSeason(base.Conf.ServerId, true)
	lastSeasonRankInfos := rmodel.Competitve.GetSeasonRankInfos(lastSeason, 1000)
	logger.Info("发送竞技场 赛季结算奖励  lastSeason:%v  lastSeasonRankInfos:%v", lastSeason, lastSeasonRankInfos)
	rmodel.Competitve.SetCompetitiveSeasonSendRewardMark(base.Conf.ServerId, lastSeason)
	rank := 0
	for i, j := 0, len(lastSeasonRankInfos); i < j; i += 2 {
		rank++
		reward, _ := gamedb.GetCompetitveSeasonEndReward(rank)
		if len(reward) > 0 {
			//rewards := gamedb.ItemInfos{}
			//rewards = reward
			//users := this.GetUserManager().GetAllUserInfoIncludeOfflineUser(int(lastSeasonRankInfos[i]))
			//if users != nil {
			//	if monthCardPrivilege := this.GetVipManager().GetPrivilege(users, pb.VIPPRIVILEGE_COMPETITVE_DAILY_REWARD); monthCardPrivilege != 0 {
			//		for _, v := range rewards {
			//			count := common.CalcTenThousand(monthCardPrivilege, v.Count)
			//			v.Count = count
			//		}
			//	}
			//}
			_ = this.GetMail().SendSystemMailWithItemInfos(int(lastSeasonRankInfos[i]), constMail.COMPETITVE_RANK_REWARD, nil, reward)
		}
	}
	return
}

func (this *CompetitveManager) buildCompetitiveInfo(ownUserId int, userIds []int, userIdInfos map[int]int, ack *pb.RefCompetitveRankAck) error {

	if len(userIds) > 0 {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(len(userIds))
		randUserId := userIds[randNum]
		userInfo := this.GetUserManager().BuilderBrieUserInfo(randUserId)
		if userInfo == nil {
			this.CheckUser()
			return gamedb.ERRMATCH
		}
		ack.UserInfo = userInfo
		ack.Score = int32(userIdInfos[randUserId])
		ack.UserInfo.Lvl = int32(userInfo.MaxLv)
		rmodel.Competitve.SetLastMarkUserId(ownUserId, randUserId)
	}
	return nil
}
