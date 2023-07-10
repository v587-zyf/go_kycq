package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type IMiJi interface {
	Up(user *objs.User, id int, ack *pb.MiJiUpAck, op *ophelper.OpBagHelperDefault) error

	GetMiJiSkillInfo(user *objs.User) *pbserver.MiJiInfo

	GetMiJiInfos(user *objs.User) []*pb.MiJiInfo
}
