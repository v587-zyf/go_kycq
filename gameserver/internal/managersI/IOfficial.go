package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type IOfficial interface {
	OfficialUpLevel(user *objs.User, op *ophelper.OpBagHelperDefault) error
}
