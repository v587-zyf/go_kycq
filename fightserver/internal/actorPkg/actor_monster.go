package actorPkg

import (
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type MonsterActor struct {
	*DefaultActor
	MonsterT       *gamedb.MonsterMonsterCfg //怪物配置表
	owner          int                       //归属玩家userId
	birthAreaIndex int                       //出生区域索引（对应stage表monster_group下角标）
	protectState   bool
}

func (this *MonsterActor) SetReliveTime(reliveTime int64) {
	this.DefaultActor.SetReliveTime(reliveTime)
	//如果是boss 推送客户端boss复活时间
	if reliveTime > 0 && this.MonsterT.Type == constFight.MONSTER_TYPE_BOSS {
		ntf := &pb.BossReliveNtf{
			ObjId:    int32(this.GetObjId()),
			ReliveCd: int32(this.reliveTime / 1000),
		}
		this.GetScene().NotifyAll(ntf)
	}
}

func (this *MonsterActor) Owner() int {
	return this.owner
}

func (this *MonsterActor) SetOwner(ownerObjId int) {

	ownerUserId := 0
	if ownerObjId > 0 {
		ownerObj := this.GetFight().GetActorByObjId(ownerObjId)
		if ownerObj != nil {
			ownerUserId = ownerObj.GetUserId()
		}
	}
	hasChange := this.owner != ownerUserId
	this.owner = ownerUserId

	//推送怪物归属变化
	if hasChange {
		//if this.MonsterT.Type == constFight.MONSTER_TYPE_BOSS {
		ntf := &pb.BossOwnerChangNtf{
			ObjId:  int32(this.GetObjId()),
			UserId: int32(this.owner),
		}
		owner := this.GetFight().GetUserMainActor(this.owner)
		if owner != nil {
			ntf.UserName = owner.NickName()
		}
		this.GetScene().NotifyAll(ntf)
		this.GetFight().OnBossOwnerChange(this.context)
		//}
	}
}

func NewMonsterActor(id int, aiCreator base.AICreator, birthAreaIndex int) *MonsterActor {
	monsterT := gamedb.GetMonsterMonsterCfg(id)
	if monsterT == nil {
		logger.Error("配置的怪物数据错误，怪物Id:%v", id)
		return nil
	}
	actor := &MonsterActor{
		MonsterT:       monsterT,
		birthAreaIndex: birthAreaIndex,
	}
	actor.DefaultActor = NewDefaultActor(pb.SCENEOBJTYPE_MONSTER, monsterT.Name, "", int32(monsterT.Job), &pbserver.ActorDisplayInfo{}, aiCreator(actor), actor)
	actor.InitMonsterProp()
	actor.InitMonsterSkill()
	if len(monsterT.Buff) > 0 {
		for _, v := range monsterT.Buff {
			actor.AddBuff(int(v), actor, true)
		}
	}
	return actor
}

func (this *MonsterActor) InitMonsterProp() {

	//TODO
	this.GetProp().Add(this.MonsterT.Attr)
	this.GetProp().AddOne(pb.PROPERTY_MP, 100000)
	this.GetProp().Calc(this.MonsterT.Job)
	//this.GetProp().AddOne(pb.PROPERTY_SPEED,gamedb.GetConf().Basespeed)
	this.GetProp().SetHpNow(this.GetProp().Get(pb.PROPERTY_HP))
	this.GetProp().SetMpNow(100000)
	//this.GetBaseProp().Add(this.MonsterT.Attr)
	//this.GetBaseProp().Calc(this.MonsterT.Job)
	////this.GetBaseProp().AddOne(pb.PROPERTY_SPEED,gamedb.GetConf().Basespeed)
	//this.GetBaseProp().SetHpNow(this.GetProp().Get(pb.PROPERTY_HP))
	//this.GetBaseProp().SetMpNow(100000)
	//this.GetBaseProp().AddOne(pb.PROPERTY_MP, this.GetProp().MpNow())
}

func (this *MonsterActor) InitMonsterSkill() {
	var skills []*base.Skill
	for _, id := range this.MonsterT.Skills {
		//初始化怪物技能
		newSkill, err := base.NewSkill(id, 1, 0)
		if err != nil {
			logger.Error("Skill Error %v", err.Error())
			continue
		}
		skills = append(skills, newSkill)
	}

	this.SetSkills(skills)
}

func (this *MonsterActor) BuildSceneObjMessage() nw.ProtoMessage {
	r := base.BuildDefaltSceneObjMessage(this)
	r.Monster = &pb.SceneMonster{Idx: int32(this.MonsterT.Monsterid), OwnerUseId: int32(this.owner)}
	if this.owner > 0 {
		owner := this.GetFight().GetUserMainActor(this.owner)
		if owner != nil {
			r.Monster.OwnerUserName = owner.NickName()
		}
	}
	return r
}

func (this *MonsterActor) BuildAppearMessage() nw.ProtoMessage {
	appearNtf := &pb.SceneEnterNtf{}
	appearNtf.Objs = append(appearNtf.Objs, this.BuildSceneObjMessage().(*pb.SceneObj))
	return appearNtf
}

func (this *MonsterActor) BuildRelliveMessage() nw.ProtoMessage {

	reliveNtf := &pb.SceneUserReliveNtf{
		Obj: this.BuildSceneObjMessage().(*pb.SceneObj),
	}
	return reliveNtf
}

func (this *MonsterActor) GetBirthAreaIndex() int {
	return this.birthAreaIndex
}

func (this *MonsterActor) AddOwner(attacker base.Actor, force bool) {
	if this.GetType() == base.ActorTypeMonster {

		if (force && attacker == nil) || attacker.GetObjId() == this.GetObjId() {
			this.SetOwner(0)
			return
		}

		if force || this.Owner() == 0 {
			ownerObjId := attacker.GetObjId()
			if attacker.GetType() == pb.SCENEOBJTYPE_FIT || attacker.GetType() == pb.SCENEOBJTYPE_PET || attacker.GetType() == pb.SCENEOBJTYPE_SUMMON {
				ownerObjId = attacker.(base.ActorLeader).GetLeader().GetObjId()
			}
			this.SetOwner(ownerObjId)
		}
	}
}

func (this *MonsterActor) RunAI() {
	this.DefaultActor.RunAI()
}

func (this *MonsterActor) Relive(reliveAddr int, reliveType int) {
	this.DefaultActor.Relive(reliveAddr, reliveType)
	this.protectState = false
	this.GetFSM().Recover()
}

func (this *MonsterActor) ChangeHp(changeHp int) (realChange int, isDeath bool) {

	oldHp := this.GetProp().HpNow()
	if changeHp < 0 {
		//怪物保护机制
		newChangeHp := this.monsterProtect(changeHp)
		realChange = this.GetProp().DecHP(newChangeHp)
	} else {
		realChange = this.GetProp().AddHP(changeHp)
	}
	if this.GetProp().HpNow() <= 0 {
		isDeath = true
	}
	newHp := this.GetProp().HpNow()

	this.TriggerPassiveSkill(constFight.SKILL_PASSIVE_CONDITION_NORMAL, nil, nil)
	if newHp < oldHp {
		this.TriggerPassiveSkillByHpChange(constFight.SKILL_PASSIVE_CONDITION_NORMAL, nil, newHp, oldHp)
	}

	return
}

func (this *MonsterActor) monsterProtect(changeHp int) int {
	if this.protectState {
		return changeHp
	}
	if this.GetProp().HpNow() <= 0 || len(this.MonsterT.Protect) < 2 {
		return changeHp
	}

	hpRate := int(float64(this.GetProp().HpNow()+changeHp) / float64(this.GetProp().Get(pb.PROPERTY_HP)) * 100)
	if hpRate > this.MonsterT.Protect[0] {
		return changeHp
	}
	//清除怪物身上debuff
	this.buffManager.DelDeBuff()
	//添加buff
	newChangeHp := int(float64(this.GetProp().Get(pb.PROPERTY_HP)*this.MonsterT.Protect[0])/100 - float64(this.GetProp().HpNow()))
	for i, l := 1, len(this.MonsterT.Protect); i < l; i++ {
		this.AddBuff(this.MonsterT.Protect[i], this, false)
	}
	logger.Debug("怪物进入保护，血量：%v/%v,变化血量：%v,保护机制：%v", this.GetProp().HpNow(), this.GetProp().Get(pb.PROPERTY_HP), newChangeHp, this.MonsterT.Protect)
	this.protectState = true
	return newChangeHp
}

func (this *MonsterActor) IsReset() bool {
	//if time.Now().Unix()-this.dieTime > RESET_TIME {
	//	return true
	//}
	return false
}

func (this *MonsterActor) IsEnemy(target base.Actor) bool {

	//角色自己肯定不能是目标
	if target.GetObjId() == this.GetObjId() {
		return false
	}

	if this.GetFight().GetStageConf().Type == constFight.FIGHT_TYPE_GUARDPILLAR {
		if target.GetType() == pb.SCENEOBJTYPE_MONSTER && target.TeamIndex() != this.TeamIndex() {
			return true
		}
		return false
	}

	if target.TeamIndex() != this.TeamIndex() {
		return true
	}
	return false

	//
	//if target.GetType() == base.ActorTypeMonster {
	//	return false
	//}
	//return true
}
func (this *MonsterActor) IsFriend(target base.Actor) bool {
	if target.GetType() == base.ActorTypeMonster {
		return true
	}
	return false
}

func (this *MonsterActor) GetMonsterT() *gamedb.MonsterMonsterCfg {
	return this.MonsterT
}
