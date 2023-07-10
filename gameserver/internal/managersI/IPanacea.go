package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IPanaceaManager interface {
	// 使用
	PanaceaUse(user *objs.User, kind int, op *ophelper.OpBagHelperDefault,  ack *pb.PanaceaUseAck) error
	// 变换相应使用次数
	PanaceaUpUseNum(user *objs.User)
}
