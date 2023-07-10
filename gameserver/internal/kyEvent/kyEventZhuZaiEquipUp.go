package kyEvent

import (
	"cqserver/gameserver/internal/objs"
)

type KyEvent_zhu_zai_equip_up struct {
	*KyEventHeroPropBase
	EquipPos int `json:"type"`       //装备部位
	EquipId  int `json:"config_id2"` //装备id
	BeforeLv int `json:"before_lv"`  //升级前等级
	Lv       int `json:"lv"`         //升级后等级
}

//主宰装备升级记录
func ZhuZaiEquipUp(user *objs.User, heroIndex int, equipPos, equipId, beforeLv, afterLv int) {
	data := &KyEvent_zhu_zai_equip_up{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		EquipPos:            equipPos,
		EquipId:             equipId,
		BeforeLv:            beforeLv,
		Lv:                  afterLv,
	}
	writeUserEvent(user, "zhu_zai_equip_up", data)
}
