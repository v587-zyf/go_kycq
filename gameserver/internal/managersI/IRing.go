package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IRingManager interface {
	// 穿戴
	Wear(user *objs.User, heroIndex, ringPos, bagPos int, op *ophelper.OpBagHelperDefault, ack *pb.RingWearAck) error
	// 移除
	Remove(user *objs.User, heroIndex, ringPos int, op *ophelper.OpBagHelperDefault, ack *pb.RingRemoveAck) error
	// 强化(ringStrengthen表)
	Strengthen(user *objs.User, heroIndex, ringPos int, op *ophelper.OpBagHelperDefault, ack *pb.RingStrengthenAck) error
	// 强化(ringPhantom表)
	RingPhantom(user *objs.User, heroIndex, ringPos int, op *ophelper.OpBagHelperDefault, ack *pb.RingPhantomAck) error
	// 幻灵技能升级
	PhantomSkill(user *objs.User, heroIndex, ringPos, phantomPos, skillId int, ack *pb.RingSkillUpAck) error
	// 融合
	Fuse(user *objs.User, id, bagPos1, bagPos2 int, op *ophelper.OpBagHelperDefault, ack *pb.RingFuseAck) error
	// 重置技能点
	ResetSkill(user *objs.User, heroIndex, ringPos, phantomPos int, ack *pb.RingSkillResetAck) error
}
