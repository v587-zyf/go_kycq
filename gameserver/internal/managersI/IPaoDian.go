package managersI

import "cqserver/gameserver/internal/objs"

type IPaoDian interface {
	EnterPaoDianFight(user *objs.User,stageId int) error
	/**
	 *  @Description: 泡点收益
	 *  @param user
	 *  @param paoDianRewardId
	 */
	PaoDianRewardNtf(user *objs.User, paoDianRewardId int,times int)
	/**
	 *  @Description: 更新最后进入高倍泡点时间
	 *  @param user
	 */
	UpdateEndTime(user *objs.User)
}
