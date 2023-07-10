package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type ITower interface {
	//挑战boss
	EnterTowerFight(user *objs.User) error
	//挑战boss后，战斗服务响应结果
	TowerFightEndAck(user *objs.User, isWin bool,items map[int]int) error
	//继续下一场
	TowerFightContinue(user *objs.User) error
	//领取每日奖励
	DayAward(user *objs.User, op *ophelper.OpBagHelperDefault) error
	//爬塔抽奖
	Lottery(user *objs.User, op *ophelper.OpBagHelperDefault) (int, error)

	TowerSweep(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.TowerSweepAck) error

	KillMonsterChangeDareNum(user *objs.User)

	SendRankReward()

}
