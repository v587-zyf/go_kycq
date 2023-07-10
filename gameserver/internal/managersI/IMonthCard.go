package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type IMonthCard interface {
	Online(user *objs.User)
	ResetMonthCard(user *objs.User, date int)
	DailyReward(user *objs.User, monthCardType int, op *ophelper.OpBagHelperDefault) error
	GetPrivilege(user *objs.User, privilege int) int

	MonthCardCheckBuy(user *objs.User, typeId, payNum int) error
	MonthCardBuyOperation(user *objs.User, payModuleId int, op *ophelper.OpBagHelperDefault)

	ItemActiveMonthCardCheck(user *objs.User, itemId int) error
	ItemActiveMonthCard(user *objs.User, itemId int) error
	CheckExpire(user *objs.User)
}
