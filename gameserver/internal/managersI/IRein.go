package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IReinManager interface {
	OnLine(user *objs.User)
	// 激活
	ReinActive(user *objs.User, ack *pb.ReinActiveAck) error
	// 转生
	Reincarnation(user *objs.User, ack *pb.ReincarnationAck) error
	// 购买修为丹
	ReinCostBuy(user *objs.User, id, num int, use bool, op *ophelper.OpBagHelperDefault, ack *pb.ReinCostBuyAck) error
	// 使用修为丹
	ReinCostUse(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.ReinCostUseAck) error
	// 重置
	ResetReinCostBuyNum(user *objs.User)
}