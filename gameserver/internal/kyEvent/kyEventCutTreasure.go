package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_cuttreasure_lv_up struct {
	*kyEventLog.KyEventPropBase
	BeforeLv int `json:"before_lv"` //升级前等级
	Lv       int `json:"lv"`        //升级后等级
}

func CutTreasureLvUp(user *objs.User, beforeLv int, afterLv int) {
	data := &KyEvent_cuttreasure_lv_up{
		KyEventPropBase: getEventBaseProp(user),
		BeforeLv:        beforeLv,
		Lv:              afterLv,
	}
	writeUserEvent(user, "cuttreasure_lv_up", data)
}
