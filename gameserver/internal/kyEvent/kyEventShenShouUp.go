package kyEvent

import (
	"cqserver/gameserver/internal/objs"
)

type KyEvent_shen_shou_up struct {
	*KyEventHeroPropBase
	ShenShouId int `json:"shen_shou_id"` //圣兽id
	BeforeLv   int `json:"before_lv"`
	Lv         int `json:"lv"`
}

//圣兽升级记录
func ShenShouUp(user *objs.User, heroIndex, shenShouId, beforeLv, afterLv int) {
	data := &KyEvent_shen_shou_up{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		ShenShouId:          shenShouId,
		BeforeLv:            beforeLv,
		Lv:                  afterLv,
	}
	writeUserEvent(user, "shen_shou_up", data)
}
