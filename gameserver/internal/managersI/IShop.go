package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IShop interface {
	OnLine(user *objs.User)
	ResetShop(user *objs.User, date int, reset bool)
	Buy(user *objs.User, id, buyNum int, op *ophelper.OpBagHelperDefault, ack *pb.ShopBuyAck) error
}
