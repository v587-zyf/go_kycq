package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IPetManager interface {
	Online(user *objs.User)
	Active(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.PetActiveAck) error
	UpLv(user *objs.User, id, itemId, itemNum int, op *ophelper.OpBagHelperDefault, ack *pb.PetUpLvAck) error
	OneKeyUpLv(user *objs.User, petId, itemId int, op *ophelper.OpBagHelperDefault, ack *pb.PetUpLvAck) error
	UpGrade(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.PetUpGradeAck) error
	Break(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.PetBreakAck) error
	ChangeWear(user *objs.User, id int, ack *pb.PetChangeWearAck) error

	AppendageStrengthen(user *objs.User, pid int, op *ophelper.OpBagHelperDefault) error
	CalcPetCombat(user *objs.User)
}
