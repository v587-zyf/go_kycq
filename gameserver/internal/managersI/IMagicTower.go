package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pbserver"
)

type IMagicTowerManager interface {

	/**
    *  @Description: 进入九层魔塔
    *  @param user
    *  @return int
    *  @return error
    **/
	EnterMagicTower(user *objs.User)(int,error)
	//九重魔塔 结算
	EndMagicTowerNtf(userRank []*pbserver.ShabakeRankScore)

	/**
	 *  @Description: 领取九层魔塔玩家信息（是否领奖 当前积分）
	 *  @param user
	 *  @param op
	 **/
	MagicTowerGetUserInfo(user *objs.User)(int,int,error)

	/**
    *  @Description: 领取九层魔塔层奖励
    *  @param user
    *  @param op
    **/
	MagicTowerlayerAward(user *objs.User,op *ophelper.OpBagHelperDefault)error
}
