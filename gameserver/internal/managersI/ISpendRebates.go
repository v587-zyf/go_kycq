package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type ISpendRebates interface {
	Online(user *objs.User)
	ResetSpendRebate(user *objs.User, reset bool)
	Reward(user *objs.User, id int, op *ophelper.OpBagHelperDefault) error
	UpdateSpendRebates(user *objs.User, op *ophelper.OpBagHelperDefault, num int)
}
