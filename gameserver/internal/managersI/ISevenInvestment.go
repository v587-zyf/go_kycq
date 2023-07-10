package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type ISevenInvestmentManager interface {

	Load(user *objs.User, ack *pb.SevenInvestmentLoadAck)

	GetAward(user *objs.User, op *ophelper.OpBagHelperDefault, id int, ack *pb.GetSevenInvestmentAwardAck) error

	SevenPayCheck(user *objs.User, payNum int) error

	SevenPayCallBack(user *objs.User)
}
