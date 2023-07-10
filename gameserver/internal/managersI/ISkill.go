package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type ISkillManager interface {
	//普通技能
	UpLv(user *objs.User, op *ophelper.OpBagHelperDefault, req *pb.SkillUpLvReq, ack *pb.SkillUpLvAck) error
	ChangePos(user *objs.User, req *pb.SkillChangePosReq, ack *pb.SkillChangePosAck) error
	ChangeWear(user *objs.User, req *pb.SkillChangeWearReq, ack *pb.SkillChangeWearAck) error
	Reset(user *objs.User, op *ophelper.OpBagHelperDefault, req *pb.SkillResetReq, ack *pb.SkillResetAck) error
	UseSkill(user *objs.User, req *pb.SkillUseReq) error

	//切割技能
	CutTreasureUpLv(user *objs.User, op *ophelper.OpBagHelperDefault) error
	CutTreasureSkillUse(user *objs.User) (int32, error)

	//远古神技
	AncientSkillActive(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault) error
	AncientSkillUpLv(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault) error
	AncientSkillUpGrade(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault) error

	/**
	 *  @Description: 清理切割技能Cd
	 *  @param user
	 *  @param stageId
	 **/
	ClearTreasureCd(user *objs.User, stageId int)
}
