package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_dragon_equip_lv_up struct {
	*kyEventLog.KyEventPropBase
	Id       int `json:"config_id1"` //龙器id
	BeforeLv int `json:"before_lv"`  //升级前等级
	Lv       int `json:"lv"`         //升级后等级
}

//龙器升级
func DragonEquipLvUp(user *objs.User, id, lv int) {
	data := &KyEvent_dragon_equip_lv_up{
		Id:       id,
		BeforeLv: lv - 1,
		Lv:       lv,
	}
	writeUserEvent(user, "dragon_equip_lv_up", data)
}
