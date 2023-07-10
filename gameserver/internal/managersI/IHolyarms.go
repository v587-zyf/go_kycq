package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IHolyarmsManager interface {
	Active(user *objs.User, id int, ack *pb.HolyActiveAck) error
	UpLevel(user *objs.User, id, itemId int, op *ophelper.OpBagHelperDefault, ack *pb.HolyUpLevelAck) error
	AutoUpLv(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.HolyUpLevelAck) error
	ActiveSkill(user *objs.User, id, skillId int, op *ophelper.OpBagHelperDefault, ack *pb.HolySkillActiveAck) error
	SkillUpLv(user *objs.User, hid, hlv int, op *ophelper.OpBagHelperDefault, ack *pb.HolySkillUpLvAck) error
}
