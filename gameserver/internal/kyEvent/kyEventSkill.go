package kyEvent

import (
	"cqserver/gameserver/internal/objs"
)

type KyEvent_skill struct {
	*KyEventHeroPropBase
	SkillId  int `json:"config_id1"`	//技能id
	BeforeLv int `json:"before_lv"`	//升级前等级
	Lv       int `json:"lv"`		//升级后等级
}

func SkillLvUp(user *objs.User, heroIndex int, skillId, beforeLv, afterLv int) {
	data := &KyEvent_skill{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		SkillId:             skillId,
		BeforeLv:            beforeLv,
		Lv:                  afterLv,
	}
	writeUserEvent(user, "skill_level_up", data)
}
