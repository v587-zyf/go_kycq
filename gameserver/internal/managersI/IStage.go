package managersI

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type IStageManger interface {
	Online(user *objs.User)
	StageFightEndReq(user *objs.User, op *ophelper.OpBagHelperDefault) error
	StageFightStartReq(user *objs.User)error
	LeaveStage(user *objs.User)
	/**
	 *  @Description: 挂机boss战斗结果
	 *  @param user
	 *  @param stageId
	 *  @param result
	 */
	StageBossKillResult(user *objs.User, stageId, result int,items map[int]int)

	//挂机boss击杀
	StageBossKill(user *objs.User)
	/**
	 *  @Description: 进入挂机地图boss战斗
	 *  @param user
	 */
	StartStageBossFight(user *objs.User) error

	/**
	 *  @Description: 获取boss掉落
	 *  @param stageId
	 *  @return gamedb.ItemInfos
	 *  @return error
	 */
	GetBossDropItem(user *objs.User, stageId int) (gamedb.ItemInfos, map[int]int, error)


	/**
    *  @Description: 领取挂机奖励
    *  @param user
    *  @param op
    *  @return error
    **/
	GetStageHookMapReward(user *objs.User,op *ophelper.OpBagHelperDefault)error
}
