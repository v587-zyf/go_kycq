package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_official_lv_up struct {
	*kyEventLog.KyEventPropBase
	BeforeLv int `json:"before_lv"` //升级前等级
	Lv       int `json:"lv"`        //升级后等级
}

//官职升级记录
func OfficialLvUp(user *objs.User, lv int) {
	data := &KyEvent_official_lv_up{
		KyEventPropBase: getEventBaseProp(user),
		BeforeLv:        lv - 1,
		Lv:              lv,
	}
	writeUserEvent(user, "official_lv_up", data)
}
