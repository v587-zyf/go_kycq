package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IFitManager interface {
	UpLv(user *objs.User, fitId int, op *ophelper.OpBagHelperDefault, ack *pb.FitUpLvAck) error
	SkillActive(user *objs.User, fitSkillId int, op *ophelper.OpBagHelperDefault, ack *pb.FitSkillActiveAck) error
	SkillUpLv(user *objs.User, fitSkillId int, op *ophelper.OpBagHelperDefault, ack *pb.FitSkillUpLvAck) error
	SkillUpStar(user *objs.User, fitSkillId int, op *ophelper.OpBagHelperDefault, ack *pb.FitSkillUpStarAck) error
	SkillChange(user *objs.User, skillId, slot int, ack *pb.FitSkillChangeAck) error
	SkillReset(user *objs.User, skillId int, op *ophelper.OpBagHelperDefault, ack *pb.FitSkillResetAck) error
	FashionUpLv(user *objs.User, fashionId int, op *ophelper.OpBagHelperDefault, ack *pb.FitFashionUpLvAck) error
	FashionChange(user *objs.User, fashionId int, ack *pb.FitFashionChangeAck) error
	EnterFit(user *objs.User, fitId int, ack *pb.FitEnterAck) error
	FitCancel(user *objs.User) error

	HolyEquipCompose(user *objs.User, equipId, equipPos int, op *ophelper.OpBagHelperDefault, ack *pb.FitHolyEquipComposeAck) error
	HolyEquipDeCompose(user *objs.User, bagPos int, op *ophelper.OpBagHelperDefault, ack *pb.FitHolyEquipDeComposeAck) error
	HolyEquipWear(user *objs.User, bagPos, equipPos int, op *ophelper.OpBagHelperDefault, ack *pb.FitHolyEquipWearAck) error
	HolyEquipRemove(user *objs.User, pos, suitType int, op *ophelper.OpBagHelperDefault, ack *pb.FitHolyEquipRemoveAck) error
	HolyEquipSuitSkillChange(user *objs.User, suitId int) error
}
