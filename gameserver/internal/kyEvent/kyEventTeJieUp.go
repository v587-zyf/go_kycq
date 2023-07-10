package kyEvent

import (
	"cqserver/gameserver/internal/objs"
)

type KyEvent_te_jie_up struct {
	*KyEventHeroPropBase
	TeJieId   int `json:"te_jie_id"`
	BeforeLv  int `json:"before_lv"`
	Lv        int `json:"lv"`
}

//特戒升级记录
func TeJieUp(user *objs.User, heroIndex int, teJieId, beforeLv, afterLv int) {
	data := &KyEvent_te_jie_up{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		TeJieId:             teJieId,
		BeforeLv:            beforeLv,
		Lv:                  afterLv,
	}
	writeUserEvent(user, "te_jie_up", data)
}
