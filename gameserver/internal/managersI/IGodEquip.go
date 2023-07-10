package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type IGodEquipManager interface {
	GodEquipActive(user *objs.User, heroIndex, id int, op *ophelper.OpBagHelperDefault) error
	GodEquipUpLevel(user *objs.User, heroIndex, id int, op *ophelper.OpBagHelperDefault) error

	GodEquipBlood(user *objs.User, heroIndex, godEquipId int, op *ophelper.OpBagHelperDefault) error
}
