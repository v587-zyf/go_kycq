package kyEvent

import (
	"cqserver/gameserver/internal/objs"
)

type KyEvent_clear struct {
	*KyEventHeroPropBase
	EquipPos   int         `json:"equip_pos"`   //装备部为
	BeforeProp map[int]int `json:"before_prop"` //洗练前属性
	Prop       map[int]int `json:"prop"`        //洗练后属性
}

//洗练
func Clear(user *objs.User, heroIndex, equipPos int, beforeProp, prop map[int]int) {
	data := &KyEvent_clear{
		KyEventHeroPropBase: getEventHeroBaseProp(user,heroIndex),
		EquipPos:        equipPos,
		BeforeProp:      beforeProp,
		Prop:            prop,
	}
	writeUserEvent(user, "clear", data)
}
