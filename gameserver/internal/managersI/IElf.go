package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IElfManager interface {
	Online(user *objs.User)
	Feed(user *objs.User, op *ophelper.OpBagHelperDefault, positions []int32, ack *pb.ElfFeedAck) error
	SkillUpLv(user *objs.User, skillId int, ack *pb.ElfSkillUpLvAck) error
	SkillChangePos(user *objs.User, skillId, pos int, ack *pb.ElfSkillChangePosAck) error
	//Active(user *objs.User, ack *pb.ElfActiveAck) error
}
