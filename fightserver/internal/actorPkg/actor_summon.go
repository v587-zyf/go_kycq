package actorPkg

import (
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/errex"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"fmt"
	"time"
)

type SummonActor struct {
	*DefaultActor
	summonId       int
	leader         base.Actor //归属玩家
	birthAreaIndex int        //出生区域索引（对应stage表monster_group下角标）
	createTime     int64
	leftTime       int
}

func NewSummonActor(owner base.Actor, summonId int, aiCreator base.AICreator) *SummonActor {
	actor := &SummonActor{
		leader:     owner,
		summonId:   summonId,
		createTime: time.Now().Unix(),
	}
	summonConf := gamedb.GetSummonConfCfg(summonId)
	actor.DefaultActor = NewDefaultActor(pb.SCENEOBJTYPE_SUMMON, fmt.Sprintf(owner.NickName()+".%s", summonConf.Name), "", pb.JOB_ZHANSHI, &pbserver.ActorDisplayInfo{}, aiCreator(actor), actor)
	actor.teamIndex = owner.TeamIndex()
	actor.InitSummonActor()
	return actor
}

func (this *SummonActor) InitSummonActor() {

	summonConf := gamedb.GetSummonConfCfg(this.summonId)
	if att, ok := summonConf.Attr[0]; ok {
		attValue := int(float64(this.leader.GetProp().Get(pb.PROPERTY_TATT_MAX)) * float64(att) / base.BASE_RATE)
		this.GetProp().AddOne(pb.PROPERTY_PATT_MAX, attValue)
		this.GetProp().AddOne(pb.PROPERTY_PATT_MIN, attValue)
	}
	if att, ok := summonConf.Attr[1]; ok {
		attValue := int(float64(this.leader.GetProp().Get(pb.PROPERTY_TATT_MAX)+this.leader.GetProp().Get(pb.PROPERTY_TATT_MIN)) * 0.5 * float64(att) / base.BASE_RATE)
		this.GetProp().AddOne(pb.PROPERTY_PATT_MAX, attValue)
		this.GetProp().AddOne(pb.PROPERTY_PATT_MIN, attValue)
	}
	if att, ok := summonConf.Attr[2]; ok {
		attValue := int(float64(this.leader.GetProp().Get(pb.PROPERTY_DEF_MAX)+this.leader.GetProp().Get(pb.PROPERTY_DEF_MIN)) * 0.5 * float64(att) / base.BASE_RATE)
		this.GetProp().AddOne(pb.PROPERTY_DEF_MAX, attValue)
		this.GetProp().AddOne(pb.PROPERTY_DEF_MIN, attValue)
	}

	if att, ok := summonConf.Attr[3]; ok {
		attValue := int(float64(this.leader.GetProp().Get(pb.PROPERTY_ADF_MAX)+this.leader.GetProp().Get(pb.PROPERTY_ADF_MIN)) * 0.5 * float64(att) / base.BASE_RATE)
		this.GetProp().AddOne(pb.PROPERTY_ADF_MAX, attValue)
		this.GetProp().AddOne(pb.PROPERTY_ADF_MIN, attValue)
	}
	this.GetProp().AddOne(pb.PROPERTY_HP, int(float64(this.leader.GetProp().Get(pb.PROPERTY_HP))*(float64(summonConf.HpFix)/base.BASE_RATE)))
	this.GetProp().Calc(pb.JOB_ZHANSHI)
	this.GetProp().SetHpNow(this.GetProp().Get(pb.PROPERTY_HP))

	skills := make([]*base.Skill, 0)
	for _, v := range summonConf.Skills {
		skillId, skillLv := gamedb.GetSkillIdAndLv(v)
		newSkill, err := base.NewSkill(skillId, skillLv, 0)
		if err != nil {
			logger.Error("Skill Error %v", err.Error())
			continue
		}
		skills = append(skills, newSkill)
	}

	this.leftTime = summonConf.Time
	this.createTime = time.Now().Unix()
	this.SetSkills(skills)
}

func (this *SummonActor) BuildSceneObjMessage() nw.ProtoMessage {
	r := base.BuildDefaltSceneObjMessage(this)
	r.Summon = &pb.SceneSummon{UserId: int32(this.leader.GetUserId()), SummonId: int32(this.summonId), ObjId: int32(this.leader.GetObjId())}
	return r
}

func (this *SummonActor) BuildAppearMessage() nw.ProtoMessage {
	appearNtf := &pb.SceneEnterNtf{}
	appearNtf.Objs = append(appearNtf.Objs, this.BuildSceneObjMessage().(*pb.SceneObj))
	return appearNtf
}

func (this *SummonActor) GetBirthAreaIndex() int {
	return this.birthAreaIndex
}

func (this *SummonActor) TeamIndex() int {
	return this.leader.TeamIndex()
}

func (this *SummonActor) RunAI() {

	now := time.Now().Unix()
	if this.leftTime > 0 && now-this.createTime > int64(this.leftTime) {
		//宠物消失离开
		this.GetFight().Leave(this)
	} else {
		this.DefaultActor.RunAI()
	}
}

func (this *SummonActor) IsEnemy(target base.Actor) bool {
	return this.leader.IsEnemy(target)
}
func (this *SummonActor) IsFriend(target base.Actor) bool {
	return this.leader.IsFriend(target)
}

func (this *SummonActor) GetUserId() int {
	return this.leader.GetUserId()
}

func (this *SummonActor) GetLeader() base.Actor {
	return this.leader
}

func (this *SummonActor) GetPlayer() *base.PlayerActor {
	return this.leader.(base.ActorPlayer).GetPlayer()
}

func (this *SummonActor) SummonAttack(skillid int, point *pb.Point, dir int, targetIds []int) nw.ProtoMessage {

	//以下为原有的技能攻击逻辑增加进入技能碰撞模式检测
	skill := this.GetSkill(skillid)
	if skill == nil {
		logger.Error("UserActor:UserAttack:nil skill:%d", skillid)
		return base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), gamedb.ERRSKILLNOTFOUND.Code)
	}

	//普通技能
	if err, hasEffect := this.CanUseSkill(skill); err != nil {
		return base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), err.(*errex.ErrorItem).Code)
	} else if !hasEffect {
		//技能标记使用，记录cd
		this.UseSkill(skill, false)
		return base.SkillAttackEffect(this, skill, this.GetDir(), make([]*pb.HurtEffect, 0), 0)
	}

	if this.GetProp().HpNow() <= 0 {
		return base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), gamedb.ERRPLAYERDIE.Code)
	}

	attackEffectNtf, err := CastSkill(this, skill, dir, targetIds, false)
	if err != nil {
		attackEffectNtf = base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), err.(*errex.ErrorItem).Code)
	}
	return attackEffectNtf
}
