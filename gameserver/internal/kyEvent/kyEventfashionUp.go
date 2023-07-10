package kyEvent

import (
	"cqserver/gameserver/internal/objs"
)

type KyEvent_fashion_up struct {
	*KyEventHeroPropBase
	FashionId int `json:"fashion_id"` //时装id
	BeforeLv  int `json:"before_lv"`  //升级前等级
	Lv        int `json:"lv"`         //升级后等级
}

//时装升级记录
func FashionUp(user *objs.User, heroIndex int, fashionId, beforeLv, afterLv int) {
	data := &KyEvent_fashion_up{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		FashionId:           fashionId,
		BeforeLv:            beforeLv,
		Lv:                  afterLv,
	}
	writeUserEvent(user, "fashion_up", data)
}
