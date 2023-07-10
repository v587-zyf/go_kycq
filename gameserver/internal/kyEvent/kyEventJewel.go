package kyEvent

import "cqserver/gameserver/internal/objs"

type KyEvent_jewel_make struct {
	*KyEventHeroPropBase
	EquipPos  int `json:"type"`  //装备部位
	JewelPos  int `json:"config_id1"`  //宝石部位
	JewelId   int `json:"config_id2"`   //宝石id
}

//宝石镶嵌
func JewelMake(user *objs.User, heroIndex, equipPos, jewelPos, jewelId int) {
	data := &KyEvent_jewel_make{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		EquipPos:            equipPos,
		JewelPos:            jewelPos,
		JewelId:             jewelId,
	}
	writeUserEvent(user, "jewel_make", data)
}
