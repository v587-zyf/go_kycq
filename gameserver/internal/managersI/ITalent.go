package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type ITalentManager interface {
	Online(user *objs.User)
	UpLv(user *objs.User, heroIndex, id int, isMax bool, ack *pb.TalentUpLvAck) error
	Reset(user *objs.User, heroIndex, id int, op *ophelper.OpBagHelperDefault, ack *pb.TalentResetAck) error

	TalentGeneral(user *objs.User)
}
