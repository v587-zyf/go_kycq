package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type ITreasureShop interface {
	Load(user *objs.User) *pb.TreasureShopLoadAck
	AutoRefreshShop(user *objs.User, sendMsg bool)
	RefreshShop(user *objs.User, op *ophelper.OpBagHelperDefault) error
	CarChange(user *objs.User, shopId int, isAdd bool, ack *pb.TreasureShopCarChangeAck) error
	Buy(user *objs.User, shop []int32, op *ophelper.OpBagHelperDefault, ack *pb.TreasureShopBuyAck) error
}
