package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IDailyPackManager interface {
	Online(user *objs.User)
	ResetDailyPack(user *objs.User, reset bool)
	Buy(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.DailyPackBuyAck) error

	DailyPackCheckBuy(user *objs.User, typeId, payNum int) error
	DailyPackBuyOperation(user *objs.User, payModuleId int, op *ophelper.OpBagHelperDefault)
}
