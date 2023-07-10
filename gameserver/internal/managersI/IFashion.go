package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type IFashion interface {
	FashionUpLevel(user *objs.User, op *ophelper.OpBagHelperDefault, heroIndex, fashionId int) error

	FashionWear(user *objs.User, heroIndex, fashionId int,isWear bool)error
}
