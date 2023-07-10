package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IVipBossManager interface {
	OnLine(user *objs.User)
	ResetVipBossFightNum(user *objs.User, date int)
	EnterVipBossFight(user *objs.User, stageId int) error
	VipBossFightResultNtf(user *objs.User, isWin bool, stageId int, items map[int]int)
	VipBossSweep(user *objs.User, stageId int, op *ophelper.OpBagHelperDefault, ack *pb.VipBossSweepAck) error
	KillMonsterChangeDareNum(user *objs.User, stageId int)
}
