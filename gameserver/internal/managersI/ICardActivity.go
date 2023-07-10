package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type ICardActivityManager interface {
	Load(user *objs.User, ack *pb.CardActivityInfosAck) error

	//抽卡
	Draw(user *objs.User, times int, ack *pb.CardActivityApplyGetAck, op *ophelper.OpBagHelperDefault) error

	//抽卡积分兑换物品
	GetReward(user *objs.User, id, times int, ack *pb.GetIntegralAwardAck, op *ophelper.OpBagHelperDefault) error

	Rest(user *objs.User, isEnter bool)
}
