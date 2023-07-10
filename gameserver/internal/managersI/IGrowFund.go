package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IGrowFundManager interface {
	GrowFundCheckBuy(user *objs.User, payNum int) error
	GrowFundBuyOperation(user *objs.User, op *ophelper.OpBagHelperDefault)

	Reward(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.GrowFundRewardAck) error
}
