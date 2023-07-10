package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type IPrivilege interface {
	Online(user *objs.User)
	Buy(user *objs.User, privilegeId int, op *ophelper.OpBagHelperDefault) error
	GetPrivilege(user *objs.User, privilegeId int) int

	ItemActive(user *objs.User, privilegeId int) error
	ItemActiveCheck(user *objs.User, privilegeId int) error
}
