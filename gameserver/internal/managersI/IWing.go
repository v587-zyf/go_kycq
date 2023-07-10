package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IWingManager interface {
	Online(user *objs.User)
	// 羽翼升级|自动升阶
	UpLevel(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault, ack *pb.WingUpLevelAck) error
	// 羽翼特殊属性升级
	UpSpecialLevel(user *objs.User, heroIndex, specialT int, op *ophelper.OpBagHelperDefault, ack *pb.WingSpecialUpAck) error
	// 羽翼使用1个材料升级
	UseMaterial(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault, ack *pb.WingUseMaterialAck) error
	// 穿戴
	Wear(user *objs.User, heroIndex, wingId int, ack *pb.WingWearAck) error
}
