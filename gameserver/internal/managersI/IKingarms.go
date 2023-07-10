package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IKingarmsManager interface {
	// 穿戴
	KingarmsChange(user *objs.User, heroIndex, pos, bagPos int, op *ophelper.OpBagHelperDefault) (error, *pb.SpecialEquipUnit)
	// 卸下
	KingarmsRemove(user *objs.User, heroIndex, pos int, op *ophelper.OpBagHelperDefault) error
}
