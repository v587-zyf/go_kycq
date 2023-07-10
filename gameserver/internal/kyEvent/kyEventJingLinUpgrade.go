package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_jin_lin struct {
	*kyEventLog.KyEventPropBase
	BeforeLv int `json:"before_lv"` //升级前等级
	Lv       int `json:"lv"`        //升级后等级
}

func JinLinUp(user *objs.User, beforeLv int, afterLv int) {
	data := &KyEvent_jin_lin{
		KyEventPropBase: getEventBaseProp(user),
		BeforeLv:        beforeLv,
		Lv:              afterLv,
	}
	writeUserEvent(user, "jin_lin_up", data)
}
