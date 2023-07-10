package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type IFirstRecharge interface {
	FirstRechargePayCheckPay(user *objs.User, typeId, payNum int) error
	FirstRechargePayOperation(user *objs.User, typeId int, op *ophelper.OpBagHelperDefault)
	Reward(user *objs.User, day int, op *ophelper.OpBagHelperDefault) error
	UpdateFirstRechargeStatus(user *objs.User)
}
