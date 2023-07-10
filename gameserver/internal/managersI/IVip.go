package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IVipManager interface {
	Online(user *objs.User)
	AddExp(user *objs.User, op *ophelper.OpBagHelperDefault, count int) error
	GetGift(user *objs.User, lv int, op *ophelper.OpBagHelperDefault, ack *pb.VipGiftGetAck) error
	GetPrivilege(user *objs.User, privilege int) int
	GetRechargeAllGift(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.RechargeAllGetAck) error
}
