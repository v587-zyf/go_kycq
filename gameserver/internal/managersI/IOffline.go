package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IOfflineManager interface {

	GetAward(user *objs.User, ack *pb.OfflineAwardGetAck, op *ophelper.OpBagHelperDefault) error

	AutoGetAward(user *objs.User)

}
