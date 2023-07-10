package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type ISignManager interface {
	Online(user *objs.User)
	// 每日重置
	ResetSign(user *objs.User, reset bool)
	// 签到
	Sign(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.SignAck) error
	// 补签
	Repair(user *objs.User, repairDay int, op *ophelper.OpBagHelperDefault, ack *pb.SignRepairAck) error
	// 累计奖励
	Cumulative(user *objs.User, cumulativeDay int, op *ophelper.OpBagHelperDefault, ack *pb.CumulativeSignAck) error
}
