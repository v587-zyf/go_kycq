package actorPkg

import (
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/errex"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type PetActor struct {
	*DefaultActor
	PetLv          int
	PetBreak       int
	PetId          int
	PetGrade       int
	PetAddSkill    []int
	PetAddAttr     map[int32]int64
	leader         base.Actor //归属玩家
	birthAreaIndex int        //出生区域索引（对应stage表monster_group下角标）
}

func NewPetActor(leader base.Actor, pet *pbserver.ActorPet, aiCreator base.AICreator) *PetActor {
	petConf := gamedb.GetPetsConfCfg(int(pet.PetId))
	actor := &PetActor{
		PetId:    int(pet.PetId),
		PetBreak: int(pet.Break),
		PetGrade: int(pet.Grade),
		PetLv:    int(pet.Lv),
		leader:   leader,
	}
	if len(pet.AddSkill) > 0 {
		addSkill := make([]int, 0)
		for _, skillId := range pet.AddSkill {
			addSkill = append(addSkill, int(skillId))
		}
		actor.PetAddSkill = addSkill
	}
	if len(pet.AddAttr) > 0 {
		addAttr := make(map[int32]int64)
		for pId, pVal := range pet.AddAttr {
			addAttr[pId] = pVal
		}
		actor.PetAddAttr = addAttr
	}
	actor.DefaultActor = NewDefaultActor(pb.SCENEOBJTYPE_PET, petConf.Name, "", pb.JOB_ZHANSHI, &pbserver.ActorDisplayInfo{}, aiCreator(actor), actor)
	actor.teamIndex = leader.TeamIndex()
	actor.InitPet()
	return actor
}

func (this *PetActor) InitPet() {

	logger.Info("初始化战宠,id：%v，阶级：%v,突破：%v,等级：%v", this.PetId, this.PetGrade, this.PetBreak, this.PetLv)
	petBreakConf := gamedb.GetPetsBreakConfCfg(gamedb.GetRealId(this.PetId, this.PetBreak))
	this.GetProp().Add(petBreakConf.AttributePets)
	petLvConf := gamedb.GetPetsLevelConfCfg(gamedb.GetRealId(this.PetId, this.PetLv))
	this.GetProp().Add(petLvConf.AttributePets)
	petGradeConf := gamedb.GetPetsGradeConfCfg(gamedb.GetRealId(this.PetId, this.PetGrade))
	this.GetProp().Add(petGradeConf.AttributePets)

	var skills []*base.Skill
	for k, v := range petGradeConf.Skill {
		if v <= 0 {
			continue
		}
		//初始化怪物技能
		newSkill, err := base.NewSkill(k, v, 0)
		if err != nil {
			logger.Error("Skill Error %v", err.Error())
			continue
		}
		skills = append(skills, newSkill)
	}

	for _, v := range petBreakConf.Effect {
		effectConf := gamedb.GetEffectEffectCfg(v)
		if len(effectConf.Attribute) > 0 {
			this.GetProp().Add(effectConf.Attribute)
		}
		if len(effectConf.Buffid) > 0 {
			for _, v := range effectConf.Buffid {
				this.AddBuff(v, this, true)
			}
		}
		if effectConf.Skillevelid > 0 {
			skillId, skillLv := gamedb.GetSkillIdAndLv(effectConf.Skillevelid)
			newSkill, err := base.NewSkill(skillId, skillLv, 0)
			if err != nil {
				logger.Error("Skill Error %v", err.Error())
				continue
			}
			skills = append(skills, newSkill)
		}
	}

	if len(this.PetAddSkill) > 0 {
		for _, skillLvId := range this.PetAddSkill {
			id, lv := gamedb.GetSkillIdAndLv(skillLvId)
			newSkill, err := base.NewSkill(id, lv, 0)
			if err != nil {
				logger.Error("Skill Error %v", err.Error())
				continue
			}
			skills = append(skills, newSkill)
		}
	}

	if len(this.PetAddAttr) > 0 {
		for pId, pVal := range this.PetAddAttr {
			this.GetProp().AddOne(int(pId), int(pVal))
		}
	}

	//属性计算
	this.GetProp().AddOne(pb.PROPERTY_HP, 1)
	this.GetProp().Calc(pb.JOB_ZHANSHI)
	this.GetProp().SetHpNow(1)
	//技能设置
	this.SetSkills(skills)

}

func (this *PetActor) TeamIndex() int {
	return this.leader.TeamIndex()
}

func (this *PetActor) BuildSceneObjMessage() nw.ProtoMessage {
	r := base.BuildDefaltSceneObjMessage(this)
	r.Pet = &pb.ScenePet{UserId: int32(this.leader.GetUserId()), Idx: int32(this.PetId)}
	return r
}

func (this *PetActor) BuildAppearMessage() nw.ProtoMessage {
	appearNtf := &pb.SceneEnterNtf{}
	appearNtf.Objs = append(appearNtf.Objs, this.BuildSceneObjMessage().(*pb.SceneObj))
	return appearNtf
}

func (this *PetActor) GetBirthAreaIndex() int {
	return this.birthAreaIndex
}

func (this *PetActor) IsEnemy(target base.Actor) bool {
	return this.leader.IsEnemy(target)
}
func (this *PetActor) IsFriend(target base.Actor) bool {
	return this.leader.IsFriend(target)
}

func (this *PetActor) GetUserId() int {
	return this.leader.GetUserId()
}

func (this *PetActor) GetLeader() base.Actor {
	return this.leader
}

func (this *PetActor) GetPlayer() *base.PlayerActor {
	return this.leader.(base.ActorPlayer).GetPlayer()
}

func (this *PetActor) UpdateInfo(pet *pbserver.ActorPet) {

	this.PetBreak = int(pet.Break)
	this.PetGrade = int(pet.Grade)
	this.PetLv = int(pet.Lv)
	if len(pet.AddSkill) > 0 {
		addSkill := make([]int, 0)
		for _, skillId := range pet.AddSkill {
			addSkill = append(addSkill, int(skillId))
		}
		this.PetAddSkill = addSkill
	}
	if len(pet.AddAttr) > 0 {
		addAttr := make(map[int32]int64)
		for pId, pVal := range pet.AddAttr {
			addAttr[pId] = pVal
		}
		this.PetAddAttr = addAttr
	}
	this.GetProp().Reset()
	this.InitPet()
}

func (this *PetActor) PetAttack(skillid int, point *pb.Point, dir int, targetIds []int) nw.ProtoMessage {

	//以下为原有的技能攻击逻辑增加进入技能碰撞模式检测
	skill := this.GetSkill(skillid)
	if skill == nil {
		logger.Error("UserActor:UserAttack:nil skill:%d", skillid)
		return base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), gamedb.ERRSKILLNOTFOUND.Code)
	}

	//普通技能
	if err, _ := this.CanUseSkill(skill); err != nil {
		return base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), err.(*errex.ErrorItem).Code)
	}

	attackEffectNtf, err := CastSkill(this, skill, dir, targetIds, false)
	if err != nil {
		attackEffectNtf = base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), err.(*errex.ErrorItem).Code)
	}
	return attackEffectNtf
}
