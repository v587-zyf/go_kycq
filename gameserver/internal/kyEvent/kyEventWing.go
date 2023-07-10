package kyEvent

import "cqserver/gameserver/internal/objs"

type KyEvent_wing_lv_up struct {
	*KyEventHeroPropBase
	BeforeId int `json:"config_id1"` //升级前id
	Id       int `json:"config_id2"` //升级后id
}

type KyEvent_wing_special_lv_up struct {
	*KyEventHeroPropBase
	BeforeLv int `json:"before_lv"` //升级前等级
	Lv       int `json:"lv"`        //升级后等级
}

//羽翼升级记录
func WingUp(user *objs.User, heroIndex int, beforeId, afterId int) {
	data := &KyEvent_wing_lv_up{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		BeforeId:            beforeId,
		Id:                  afterId,
	}
	writeUserEvent(user, "wing_lv_up", data)
}

//羽翼技能升级记录
func WingSpecialLvUp(user *objs.User, heroIndex, lv int) {
	data := &KyEvent_wing_special_lv_up{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		BeforeLv:            lv - 1,
		Lv:                  lv,
	}
	writeUserEvent(user, "wing_special_lv_up", data)
}
