package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IPersonBoss interface {
	Online(user *objs.User)
	ResetPersonBoss(user *objs.User, date int)
	EnterPersonBossFightReq(user *objs.User, stageId int, ack *pb.EnterPersonBossFightAck) error
	PersonBossFightResultNtf(user *objs.User, stageId int, isWin bool, items map[int]int)
	PersonBossSweep(user *objs.User, stageId int, op *ophelper.OpBagHelperDefault, ack *pb.PersonBossSweepAck) error
	KillMonsterChangeDareNum(user *objs.User, stageId, hasFightNum int)
	//redis
	GetBossKillNum(userId, stageId int) int
}
