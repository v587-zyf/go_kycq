package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IDragonarmsManager interface {
	// 穿戴
	DragonarmsChange(user *objs.User, heroIndex, pos, bagPos int, op *ophelper.OpBagHelperDefault) (error, *pb.SpecialEquipUnit)
	// 卸下
	DragonarmsRemove(user *objs.User, heroIndex, pos int, op *ophelper.OpBagHelperDefault) error
}
