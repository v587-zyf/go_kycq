package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type ILabelManager interface {
	Online(user *objs.User)
	ResetLabel(user *objs.User, date int)
	Up(user *objs.User, op *ophelper.OpBagHelperDefault) error
	Transfer(user *objs.User, job int) error
	DayReward(user *objs.User, op *ophelper.OpBagHelperDefault) error
	SendLabelTaskNtf(user *objs.User, conditionId int)
}
