package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IFabaoManager interface {
	// 法宝激活
	ActiveChange(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.FabaoActiveAck) error
	// 法宝升级
	UpLevel(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.FabaoUpLevelAck) error
	// 法宝技能激活
	ActiveSkillChange(user *objs.User, id, skillId int, op *ophelper.OpBagHelperDefault, ack *pb.FabaoSkillActiveAck) error
}
