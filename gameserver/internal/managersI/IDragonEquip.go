package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IDragonEquipManager interface {
	UpLv(user *objs.User, heroIndex, id int, op *ophelper.OpBagHelperDefault, ack *pb.DragonEquipUpLvAck) error
}
