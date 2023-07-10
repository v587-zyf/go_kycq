package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IMagicCircleManager interface {
	UpLv(user *objs.User, heroIndex, magicCircleType int, op *ophelper.OpBagHelperDefault, ack *pb.MagicCircleUpLvAck) error
	ChangeWear(user *objs.User, heroIndex, magicCircleLvId int, ack *pb.MagicCircleChangeWearAck) error
}
