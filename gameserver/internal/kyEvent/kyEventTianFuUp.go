package kyEvent

import (
	"cqserver/gameserver/internal/objs"
)

type KyEvent_tian_fu_up struct {
	*KyEventHeroPropBase
	TianFuId int `json:"tian_fu_id"`
	BeforeLv int `json:"before_lv"`
	Lv       int `json:"lv"`
}

//天赋升级记录
func TianFuUp(user *objs.User, heroIndex, tianFuId, beforeLv, afterLv int) {
	data := &KyEvent_tian_fu_up{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		TianFuId:            tianFuId,
		BeforeLv:            beforeLv,
		Lv:                  afterLv,
	}
	writeUserEvent(user, "tian_fu_up", data)
}
