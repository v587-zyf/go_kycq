package managersI

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IEquipManager interface {
	Online(user *objs.User)
	EquipChange(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault, equipPos int, EquipBagPos int) error
	EquipChangeOneKey(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault) ([]int, error)
	EquipsStrengthen(user *objs.User, heroIndex int, pos int, op *ophelper.OpBagHelperDefault, ack *pb.EquipStrengthenAck) error
	EquipStrengthenOneKey(user *objs.User, heroIndex int, isBreak bool, op *ophelper.OpBagHelperDefault, ack *pb.EquipStrengthenAutoAck) error
	EquipRemove(user *objs.User, heroIndex, pos int, op *ophelper.OpBagHelperDefault, ack *pb.EquipRemoveAck) error
	EquipCompareCombat(job int, equip1, equip2 *model.Equip) bool
	CheckWeaponLucky(user *objs.User, heroIndex, itemId int) error
	WeaponBless(user *objs.User, heroIndex int) error

	// 洗练
	Clear(user *objs.User, heroIndex, pos, itemId int, op *ophelper.OpBagHelperDefault, ack *pb.ClearAck) error
}
