package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IDictateManager interface {
	// 主宰装备升级
	DictateUpLv(user *objs.User, heroIndex, body int, op *ophelper.OpBagHelperDefault,  ack *pb.DictateUpAck) error
}
