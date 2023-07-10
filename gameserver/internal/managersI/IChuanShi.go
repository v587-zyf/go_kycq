package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IChuanShiManager interface {
	Wear(user *objs.User, heroIndex, bagPos int, op *ophelper.OpBagHelperDefault, ack *pb.ChuanShiWearAck) error
	Remove(user *objs.User, heroIndex, equipPos int, op *ophelper.OpBagHelperDefault) error
	DeCompose(user *objs.User, bagPos int, op *ophelper.OpBagHelperDefault) error

	Strengthen(user *objs.User, heroIndex, pos, stone int, op *ophelper.OpBagHelperDefault, ack *pb.ChuanshiStrengthenAck) error
}
