package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IKillMonsterManager interface {
	WriteKillMonster(user *objs.User, stageId int)

	LoadUni(user *objs.User) (error, []*pb.KillMonsterUniInfo)
	LoadPer(user *objs.User) (error, []*pb.KillMonsterPerInfo)
	LoadMil(user *objs.User) (error, []*pb.KillMonsterMilInfo)

	DrawUni(user *objs.User, stageId int, op *ophelper.OpBagHelperDefault) error
	DrawUniFirst(user *objs.User, stageId int, op *ophelper.OpBagHelperDefault) error
	DrawPer(user *objs.User, stageId int, op *ophelper.OpBagHelperDefault) error
	DrawMil(user *objs.User, t int, op *ophelper.OpBagHelperDefault, ack *pb.KillMonsterMilDrawAck) error
}
