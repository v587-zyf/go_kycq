package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type ITitleManager interface {
	Online(user *objs.User)
	Active(user *objs.User, titleId int, op *ophelper.OpBagHelperDefault) error
	AutoActive(user *objs.User)
	CheckExpire(user *objs.User)
	Wear(user *objs.User, heroIndex, titleId int) error
	Remove(user *objs.User, heroIndex int) (error, int)
	Look(user *objs.User, titleId int) error
}
