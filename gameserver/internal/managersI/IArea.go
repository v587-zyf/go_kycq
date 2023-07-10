package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type IAreaManager interface {
	UpLv(user *objs.User, heroIndex, areaT int, op *ophelper.OpBagHelperDefault) error
}
