package kyEvent

import (
	"cqserver/gameserver/internal/objs"
)

type KyEvent_equip_intensify struct {
	*KyEventHeroPropBase
	EquipType int `json:"type"`       //装备部位
	EquipId   int `json:"config_id2"` //装备id
	BeforeLv  int `json:"before_lv"`  //升级前等级
	Lv        int `json:"lv"`         //升级后等级
}

//装备强化
func EquipIntensify(user *objs.User, heroIndex int, equipType, equipId, beforeLv, afterLv int) {
	data := &KyEvent_equip_intensify{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		EquipType:           equipType,
		EquipId:             equipId,
		BeforeLv:            beforeLv,
		Lv:                  afterLv,
	}
	writeUserEvent(user, "equip_intensify", data)
}
