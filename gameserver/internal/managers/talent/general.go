package talent

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
)

func (this *TalentManager) TalentGeneral(user *objs.User) {
	userHero := user.Heros
	for heroIndex, hero := range userHero {
		hero.TalentGeneral = make(map[int]map[int]int)
		heroTalent := hero.Talent
		heroTalentList := heroTalent.TalentList
		if len(heroTalentList) <= 0 {
			continue
		}
		//每个英雄每个模块只记录最高级的
		moduleLv := make(map[int]map[int]int)
		moduleAttr := make(map[int]map[int]map[int]int)
		cfgs := gamedb.GetTalentGeneral()
		for _, cfg := range cfgs {
			mLv := cfg.Level
			way := cfg.Talentway
			moduleType := cfg.Type
			talentId := cfg.TalentID
			if moduleLv[moduleType] == nil {
				moduleLv[moduleType] = make(map[int]int)
			}
			if moduleAttr[moduleType] == nil {
				moduleAttr[moduleType] = make(map[int]map[int]int)
			}
			if moduleAttr[moduleType][talentId] == nil {
				moduleAttr[moduleType][talentId] = make(map[int]int)
			}
			//校验对应condition、校验是否点亮对应天赋
			if heroTalentList[way] == nil || heroTalentList[way].Talents == nil {
				continue
			}
			lv, ok := heroTalentList[way].Talents[talentId]
			if !this.GetCondition().CheckMulti(user, heroIndex, cfg.Condition) || !ok || lv < mLv {
				continue
			}
			if moduleLv[moduleType][talentId] > mLv {
				continue
			}
			moduleAttr[moduleType][talentId] = cfg.Icon
			moduleLv[moduleType][talentId] = mLv
		}
		for moduleT, talentAttr := range moduleAttr {
			if hero.TalentGeneral[moduleT] == nil {
				hero.TalentGeneral[moduleT] = make(map[int]int)
			}
			for _, attr := range talentAttr {
				for talentConditionId, val := range attr {
					hero.TalentGeneral[moduleT][talentConditionId] += val
					logger.Debug("m:%v tid:%v val:%v", moduleT, talentConditionId, val)
				}
			}
		}
	}
}
