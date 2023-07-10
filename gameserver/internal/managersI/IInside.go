package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IInsideManager interface {
	// 内功升星
	InsideUpStar(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault) error
	// 内功升阶
	InsideUpGrade(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault, ack *pb.InsideUpGradeAck) error
	// 内功升重
	InsideUpOrder(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault, ack *pb.InsideUpOrderAck) error
	// 内功技能升级
	InsideSkillUpLv(user *objs.User, heroIndex, skillId int, op *ophelper.OpBagHelperDefault, ack *pb.InsideSkillUpLvAck) error

	AutoUp(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault) error
}