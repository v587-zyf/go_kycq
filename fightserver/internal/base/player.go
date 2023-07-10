package base

import (
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

/**
*  @Description: 玩家信息
**/
type PlayerActor struct {
	sessionId          uint32
	userId             int
	heroActors         map[int]Actor
	petActor           Actor
	fitActor           Actor
	summonActors       map[int]Actor
	redPacketInfo      *pbserver.ActorRedPacket
	stageFightNum      int
	userCombat         int //玩家总战力
	userHpTotal        int //玩家总血量
	daBaoEnergy        int //打宝秘境体力
	reliveTimes        int //复活次数
	reliveByIngotTimes int //元宝复活次数
	toHelpUserId       int //被协助者Id
	toHelpTime         int64
	fightNum           int      //当前战斗模式可挑战次数（用于协助判断，有次数时，被协助者退出战斗可已成为归属者，无次数，则不能成为归属者）
	passiveSkill       []*Skill //技能
}

func (this *PlayerActor) GetSkill(skillId int) *Skill {
	for _, v := range this.passiveSkill {
		if v.Skillid == skillId {
			return v
		}
	}
	return nil
}

func (this *PlayerActor) PassiveSkills() []*Skill {
	return this.passiveSkill
}

func (this *PlayerActor) ToHelpTime() int64 {
	return this.toHelpTime
}

func (this *PlayerActor) ToHelpUserId() int {
	return this.toHelpUserId
}

func (this *PlayerActor) SetToHelpUserId(toHelpUserId int) {
	this.toHelpUserId = toHelpUserId
	if toHelpUserId > 0 {
		this.toHelpTime = time.Now().UnixNano()
	} else {
		this.toHelpTime = 0
	}
}

func (this *PlayerActor) FightNum() int {
	return this.fightNum
}

func (this *PlayerActor) SetFightNum(fightNum int) {
	this.fightNum = fightNum
}

func (this *PlayerActor) ReliveTimes() int {
	return this.reliveTimes
}

func (this *PlayerActor) SetReliveTimes(reliveTimes int) {
	this.reliveTimes = reliveTimes
}

func (this *PlayerActor) ReliveByIngotTimes() int {
	return this.reliveByIngotTimes
}

func (this *PlayerActor) SetReliveByIngotTimes(reliveByIngotTimes int) {
	this.reliveByIngotTimes = reliveByIngotTimes
}

func (this *PlayerActor) UserHpTotal() int {
	return this.userHpTotal
}

func (this *PlayerActor) SetUserHpTotal(userHpTotal int) {
	this.userHpTotal = userHpTotal
}

func (this *PlayerActor) CalcUserHpTotal() {
	this.userHpTotal = 0
	for _, v := range this.heroActors {
		this.userHpTotal += v.GetProp().Get(pb.PROPERTY_HP)
	}
}

func (this *PlayerActor) UserCombat() int {
	return this.userCombat
}

func (this *PlayerActor) SetUserCombat(userCombat int) {
	this.userCombat = userCombat
}

func (this *PlayerActor) StageFightNum() int {
	return this.stageFightNum
}

func (this *PlayerActor) SetStageFightNum(stageFightNum int) {
	this.stageFightNum = stageFightNum
}

func (this *PlayerActor) SessionId() uint32 {
	return this.sessionId
}

func (this *PlayerActor) SummonActors() map[int]Actor {
	return this.summonActors
}

func (this *PlayerActor) HeroActors() map[int]Actor {
	return this.heroActors
}

func (this *PlayerActor) GetSummonActor(objId int) Actor {
	return this.summonActors[objId]
}

func (this *PlayerActor) AddSummonActor(objId int, summonActor Actor) {
	this.summonActors[objId] = summonActor
}

func (this *PlayerActor) RemoveSummonActor(objId int) {
	delete(this.summonActors, objId)
}

func (this *PlayerActor) GetFirstReliveHero() Actor {

	if this.fitActor != nil {
		return this.fitActor
	}
	for _, v := range constUser.USER_HERO_INDEX {
		if hero, ok := this.heroActors[v]; ok && hero.GetProp().HpNow() > 0 {
			return hero
		}
	}
	return nil
}

func (this *PlayerActor) GetHeroActor(heroIndex int) Actor {
	return this.heroActors[heroIndex]
}

func (this *PlayerActor) AddHeroActors(heroIndex int, heroActor Actor) {
	this.heroActors[heroIndex] = heroActor
}

func (this *PlayerActor) FitActor() Actor {
	return this.fitActor
}

func (this *PlayerActor) SetFitActor(fitActor Actor) {
	this.fitActor = fitActor
}

func (this *PlayerActor) PetActor() Actor {
	return this.petActor
}

func (this *PlayerActor) SetPetActor(petActor Actor) {
	this.petActor = petActor
}

func (this *PlayerActor) SetRedPacketInfo(redPacketInfo *pbserver.ActorRedPacket) {
	this.redPacketInfo = redPacketInfo
}

func (this *PlayerActor) RedPacketInfo() *pbserver.ActorRedPacket {
	return this.redPacketInfo
}

func (this *PlayerActor) CheckUserDie() bool {

	if this.fitActor != nil && !this.fitActor.IsDeath() {
		return false
	}

	for _, v := range this.heroActors {
		if !v.IsDeath() {
			return false
		}
	}
	return true
}

func (this *PlayerActor) SetDaBaoEnergy(energy int) {
	this.daBaoEnergy = energy
}
func (this *PlayerActor) DaBaoEnergy() int {
	return this.daBaoEnergy
}

func (this *PlayerActor) ChangeTeam(teamIndex int) {
	if this.fitActor != nil {
		this.fitActor.SetTeamIndex(teamIndex)
	}
	if this.petActor != nil {
		this.petActor.SetTeamIndex(teamIndex)
	}
	if this.summonActors != nil && len(this.summonActors) > 0 {
		for _, v := range this.summonActors {
			v.SetTeamIndex(teamIndex)
		}
	}
	for _, v := range this.heroActors {
		v.SetTeamIndex(teamIndex)
	}

}

func (this *PlayerActor) ResetSkill(skills []*pbserver.Skill) {

	tempPassiveSkills := this.passiveSkill
	this.passiveSkill = make([]*Skill, 0)

loop:
	for _, newSkillData := range skills {
		has := false
		for _, baseSkill := range tempPassiveSkills {
			if baseSkill.Skillid == int(newSkillData.Id) {
				has = true
				if baseSkill.LevelT.Level < int(newSkillData.Level) {
					baseSkill.ResetSkillLvel(int(newSkillData.Level))
				}
				this.passiveSkill = append(this.passiveSkill, baseSkill)
				continue loop
			}
		}
		if !has {
			newSkill, err := NewSkill(int(newSkillData.Id), int(newSkillData.Level), newSkillData.CdEndTime)
			if err != nil {
				logger.Error("更新玩家技能数据异常，：%v", err)
				continue
			}
			this.passiveSkill = append(this.passiveSkill, newSkill)
		}
	}
}

func NewPlayer(sessionId uint32, redPacket *pbserver.ActorRedPacket, skills []*pbserver.Skill) *PlayerActor {
	player := &PlayerActor{
		heroActors:    make(map[int]Actor),
		summonActors:  make(map[int]Actor),
		redPacketInfo: redPacket,
		sessionId:     sessionId,
		passiveSkill:  make([]*Skill, len(skills)),
	}
	if len(skills) > 0 {
		for k, v := range skills {
			skill, err := NewSkill(int(v.Id), int(v.Level), 0)
			if skill != nil && err == nil {
				player.passiveSkill[k] = skill
			}
		}
	}
	return player
}
