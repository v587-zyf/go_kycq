package kyEvent

import (
	"cqserver/gameserver/internal/objs"
)

type KyEvent_ling_yu_up struct {
	*KyEventHeroPropBase
	LingYuType int `json:"type"`      //领域类型
	BeforeLv   int `json:"before_lv"` //升级前等级
	Lv         int `json:"lv"`        //升级后等级
}

//领域升级记录
func LingYuUp(user *objs.User, heroIndex int, lingYuType, beforeLv, afterLv int) {
	data := &KyEvent_ling_yu_up{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		LingYuType:          lingYuType,
		BeforeLv:            beforeLv,
		Lv:                  afterLv,
	}
	writeUserEvent(user, "ling_yu_up", data)
}
