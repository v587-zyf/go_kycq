package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IFieldBoss interface {
	Online(user *objs.User)
	ResetFieldBoss(user *objs.User, date int)
	Load(user *objs.User, area int, ack *pb.FieldBossLoadAck) error
	BuyNum(user *objs.User, use bool, buyNum int, op *ophelper.OpBagHelperDefault, ack *pb.FieldBossBuyNumAck) error
	SendFieldBossNtf(fieldBossNtf *pb.FieldBossNtf)
	EnterFieldBossFight(user *objs.User, stageId int) error
	FieldBossFightEndAck(user *objs.User, winUserId, stageId int, items map[int]int) error
	UserLeave(user *objs.User, stageId int)

	FirstReceive(user *objs.User, op *ophelper.OpBagHelperDefault) error
}
