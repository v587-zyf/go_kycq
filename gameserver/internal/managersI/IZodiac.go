package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IZodiacManager interface {
	// 穿戴
	ZodiacChange(user *objs.User, heroIndex, pos, bagPos int, op *ophelper.OpBagHelperDefault) (error, *pb.SpecialEquipUnit)
	// 卸下
	ZodiacRemove(user *objs.User, heroIndex, pos int, op *ophelper.OpBagHelperDefault) error
}