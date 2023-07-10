package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IComposeManager interface {
	//普通合成
	Compose(user *objs.User, heroIndex, subTab, composeNum int, op *ophelper.OpBagHelperDefault, ack *pb.ComposeAck) error
	//合成装备
	ComposeEquip(user *objs.User, req *pb.ComposeEquipReq, op *ophelper.OpBagHelperDefault) error
	//合成传世装备
	ComposeChuanShiEquip(user *objs.User, subId int, op *ophelper.OpBagHelperDefault) error
}
