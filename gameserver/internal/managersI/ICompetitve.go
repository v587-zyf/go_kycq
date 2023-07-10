package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

//new 竞技场
type ICompetitiveManager interface {
	LoadInfo(user *objs.User, ack *pb.CompetitveLoadAck) error

	//购买竞技场挑战次数
	BuyCompetitiveChallengeNum(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.BuyCompetitveChallengeTimesAck) error

	//竞技场每日奖励领取
	GetCompetitiveDailyReward(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.GetCompetitveDailyRewardAck) error

	//定时5点发放赛季奖励
	SendSeasonEndReward()

	//竞技场匹配对手
	RefCompetitiveRival(user *objs.User, ack *pb.RefCompetitveRankAck) error

	// 进入战斗
	EnterCompetitveFight(user *objs.User, challengeUid, challengeRanking int) error

	/**
	 *  @Description: 竞技场战斗结果推送
	 *  @param user		玩家数据
	 *  @param result	结果
	 */
	CompetitveFightEndResult(user *objs.User, result bool)

	//背包使用道具购买挑战次数 检查
	BagUserItemAddChallengeNumCheck(user *objs.User) error

	//背包使用道具购买挑战次数
	BagUserItemAddChallengeNum(user *objs.User) error

	CompetitiveMultipleClaim(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.CompetitveMultipleClaimAck) error

	DayReset(user *objs.User, isReset bool)

	OnLine(user *objs.User)

	RefCompetitiveRivalNew(user *objs.User, ack *pb.RefCompetitveRankAck) error
	//竞技场多倍领取
}
