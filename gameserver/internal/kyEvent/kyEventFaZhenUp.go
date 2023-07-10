package kyEvent

import (
	"cqserver/gameserver/internal/objs"
)

type KyEvent_fa_zhen_up struct {
	*KyEventHeroPropBase
	FaZhenType int `json:"config_id1"` //法阵id
	BeforeLv   int `json:"before_lv"`  //升级前等级
	Lv         int `json:"lv"`         //升级后等级
}

type KyEvent_fa_zhen_change struct {
	*KyEventHeroPropBase
	FaZhenId int `json:"config_id1"` //法阵id
}

//法阵升级记录
func FaZhenUp(user *objs.User, heroIndex, faZhenId, beforeLv, afterLv int) {
	data := &KyEvent_fa_zhen_up{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		FaZhenType:          faZhenId,
		BeforeLv:            beforeLv,
		Lv:                  afterLv,
	}
	writeUserEvent(user, "fa_zhen_up", data)
}

//法阵穿戴
func FazhenChange(user *objs.User, heroIndex, id int) {
	data := &KyEvent_fa_zhen_change{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		FaZhenId:            id,
	}
	writeUserEvent(user, "fa_zhen_change", data)
}
