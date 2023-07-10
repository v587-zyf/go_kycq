package kyEvent

import (
	"cqserver/gameserver/internal/objs"
)

type KyEvent_shen_shou_skill_up struct {
	*KyEventHeroPropBase
	ShenShouId      int `json:"shen_shou_id"`
	ShenShouSkillId int `json:"shen_shou_skill_id"`
	BeforeLv        int `json:"before_lv"`
	Lv              int `json:"lv"`
}

//圣兽技能升级记录
func ShenShouSkillUp(user *objs.User, heroIndex, shenShouId, shenShouSkillId, beforeLv, afterLv int) {
	data := &KyEvent_shen_shou_skill_up{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		ShenShouId:          shenShouId,
		ShenShouSkillId:     shenShouSkillId,
		BeforeLv:            beforeLv,
		Lv:                  afterLv,
	}
	writeUserEvent(user, "shen_shou_skill_up", data)
}
