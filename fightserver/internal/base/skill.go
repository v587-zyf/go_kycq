package base

import (
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"fmt"
)

type Skill struct {
	*gamedb.SkillSkillCfg
	LevelT                   *gamedb.SkillLevelSkillCfg
	skillUseTime             int64 //技能使用时间（毫秒）
	nextAttackTime           int64 // 下一次可攻击的时间(毫秒)
	Duration                 int64 // 技能持续时间
	areaOffset               map[int][]int
	talentEffect             []int
	passiveSkillConditionUse int
}

func (this *Skill) PassiveSkillConditionUse() int {
	return this.passiveSkillConditionUse
}

func (this *Skill) SetPassiveSkillConditionUse(passiveSkillConditionUse int) {
	this.passiveSkillConditionUse = passiveSkillConditionUse
}

func NewSkill(id, level int, nextAttackTime int64) (*Skill, error) {
	skillT := gamedb.GetSkillSkillCfg(id)
	if skillT == nil {
		logger.Error("no skill found, id: %d", id)
		return nil, fmt.Errorf("no skill found, id: %d", id)
	}
	levelT := gamedb.GetSkillLvConf(id, level)
	if levelT == nil {
		logger.Error("no skill level found, id: %d, level: %d", id, level)
		return nil, fmt.Errorf("no skill level found, id: %d, level: %d", id, level)
	}

	var areaOffset map[int][]int
	if _, ok := skillRangTypeOffsetFunc[levelT.RangeType]; ok {
		areaOffset = skillRangTypeOffsetFunc[levelT.RangeType](levelT.Range)
		//获取技能释放区域偏移
	} else {
		areaOffset = make(map[int][]int)
		for _, v := range pb.SCENEDIR_ARRAY {
			areaOffset[v] = make([]int, 0)
		}
	}

	logger.Debug("初始话技能：%v", id, level)
	var duration int64
	return &Skill{
		SkillSkillCfg:  skillT,
		LevelT:         levelT,
		Duration:       duration,
		areaOffset:     areaOffset,
		nextAttackTime: nextAttackTime,
	}, nil
}

func (this *Skill) SetTalentEffect(talentEffect []int) {

	this.talentEffect = talentEffect
}

func (this *Skill) GetTalentEffect() []int {
	return this.talentEffect
}

func (this *Skill) ResetSkillLvel(lv int) {
	levelT := gamedb.GetSkillLvConf(this.Skillid, lv)
	if levelT != nil {
		this.LevelT = levelT
	}
}

func (this *Skill) GetLevel() int {

	return this.LevelT.Level
}

func (this *Skill) CanUse(actorMp int) error {
	now := common.GetNowMillisecond()
	//这里-2 是为了容错，客户端提前了发上来了
	if now-this.nextAttackTime < -20 {
		return gamedb.ERRSKILLCASTBYCD
	}
	if this.LevelT.MP > 0 && this.LevelT.MP > actorMp {
		return gamedb.ERRSKILLCASTMP
	}
	return nil
}

func (this *Skill) InAttackArea(dir, offsetX, offsetY int) bool {

	x, y := scene.GetDirOffset(dir)
	for d := 0; d <= this.LevelT.Distance; d++ {
		for i := 0; i < len(this.areaOffset[dir]); i += 2 {

			newX := this.areaOffset[dir][i] + d*x
			newY := this.areaOffset[dir][i+1] + d*y
			//logger.Debug("InAttackArea cal dir:%v,x,y:%v-%v，newX,newY:%v-%v", dir, offsetX, offsetY, newX, newY)
			if offsetX == newX && offsetY == newY {
				return true
			}
		}
	}

	return false
}

func (this *Skill) GetAttackAreaOffset(dir int) []int {

	return this.areaOffset[dir]
}

func (this *Skill) Use() {
	this.skillUseTime = common.GetNowMillisecond()
	this.nextAttackTime = this.skillUseTime + int64(this.LevelT.CD)
}

func (this *Skill) GetSkillUseTime() int64 {
	return this.skillUseTime
}

func (this *Skill) GetNextAttackTime() int64 {
	return this.nextAttackTime
}

func (this *Skill) SetNextAttackTime(t int64) {
	this.nextAttackTime = t
}

func (this *Skill) PassiveSkillCheck(attacker, target Actor, castPreSkill int) bool {

	if this.CastPreSkill != 0 && this.CastPreSkill != castPreSkill {
		return false
	}

	canUse := this.CanUse(attacker.GetProp().MpNow())
	if canUse != nil {
		return false
	}
	check := true
	if len(this.LevelT.Passiveconditions2) > 0 {

		if _, ok := this.LevelT.Passiveconditions2[constFight.SKILL_PASSIVE_HP_CHECK_LESS]; ok {
			return false
		}

		if value, ok := this.LevelT.Passiveconditions2[constFight.SKILL_PASSIVE_HP_CHECK_SELF_LOW]; ok {

			if int(float64(attacker.GetProp().HpNow())/float64(attacker.GetProp().Get(pb.PROPERTY_HP))*BASE_RATE) > value {
				check = false
			}
		}
		if value, ok := this.LevelT.Passiveconditions2[constFight.SKILL_PASSIVE_HP_CHECK_SELF_HIGH]; ok {

			if int(float64(attacker.GetProp().HpNow())/float64(attacker.GetProp().Get(pb.PROPERTY_HP))*BASE_RATE) < value {
				check = false
			}
		}

		if value, ok := this.LevelT.Passiveconditions2[constFight.SKILL_PASSIVE_HP_CHECK_ENMY_LOW]; ok {

			if target == nil {
				check = false
			} else if int(float64(target.GetProp().HpNow())/float64(target.GetProp().Get(pb.PROPERTY_HP))*BASE_RATE) > value {
				check = false
			}
		}

		if value, ok := this.LevelT.Passiveconditions2[constFight.SKILL_PASSIVE_HP_CHECK_ENMY_HIGH]; ok {

			if target == nil {
				check = false
			}
			if int(float64(target.GetProp().HpNow())/float64(target.GetProp().Get(pb.PROPERTY_HP))*BASE_RATE) < value {
				check = false
			}
		}

	}
	return check
}
