package managersI

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IWarOrderManager interface {
	Online(user *objs.User)
	ResetWarOrder(user *objs.User, reset bool) bool
	LvReward(user *objs.User, lv int, op *ophelper.OpBagHelperDefault, ack *pb.WarOrderLvRewardAck) error
	Exchange(user *objs.User, exchangeId, num int, op *ophelper.OpBagHelperDefault, ack *pb.WarOrderExchangeAck) error
	AddExp(user *objs.User, op *ophelper.OpBagHelperDefault, exp int) error
	AutoUpLv(user *objs.User, userWarOrder *model.WarOrder, isSend bool) (int, int)

	WarOrderCheckBuyLuxury(user *objs.User, payNum int) error
	WarOrderBuyLuxuryOperation(user *objs.User)
	WarOrderCheckBuyExp(user *objs.User, payNum int) error
	WarOrderBuyExpOperation(user *objs.User)

	TaskFinish(user *objs.User, op *ophelper.OpBagHelperDefault, req *pb.WarOrderTaskFinishReq, ack *pb.WarOrderTaskFinishAck) error
	TaskReward(user *objs.User, op *ophelper.OpBagHelperDefault, req *pb.WarOrderTaskRewardReq, ack *pb.WarOrderTaskRewardAck) error
	WriteWarOrderTaskByKillMonster(user *objs.User, monsterId, killNum int)
	WriteWarOrderTask(user *objs.User, t int, val []int)
}
