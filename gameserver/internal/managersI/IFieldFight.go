package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

//野战
type IFieldFightManager interface {
	LoadInfo(user *objs.User, ack *pb.FieldFightLoadAck) error

	//定时0点存储活跃玩家战力 用于匹配对手
	FiveAmRecordUserCombat()

	//长时间不活跃玩家 上线时存储他的当前战力 用于匹配野战玩家 and 判断是否需要刷新对手
	JudgeIsSetUserFiveAmCombat(user *objs.User)

	//野战购买挑战次数
	BuyFieldFightChallengeNum(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.BuyFieldFightChallengeTimesAck) error

	EnterFieldFight(user *objs.User, challengeUid, isBeatBack int) error

	/**
	 *  @Description:
	 *  @param user			玩家
	 *  @param result		结果
	 *  @param challengeUid	被挑战者
	 *  @param isBeatBack	是否反击
	 *  @return error
	 */
	FieldFightFightEndResult(user *objs.User, result bool, challengeUid, isBeatBack int) error

	GetOpenDayByReduceTime() int

	BuyFieldFightChallengeNumByBag(user *objs.User) error

	DayReset(user *objs.User,isReset bool)

	OnLine(user *objs.User)

	//刷新劲敌
	RivalUser(user *objs.User, ack *pb.RefFieldFightRivalUserAck, isLoginCheck bool) error
}
