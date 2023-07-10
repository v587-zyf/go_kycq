package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_jue_xue_up struct {
	*kyEventLog.KyEventPropBase
	JueXueId int `json:"config_id1"` //绝学id
	BeforeLv int `json:"before_lv"`  //升级前等级
	Lv       int `json:"lv"`         //升级后等级
}

//绝学升级记录
func JueXueUp(user *objs.User, jueXueId, beforeLv int, afterLv int) {
	data := &KyEvent_jue_xue_up{
		KyEventPropBase: getEventBaseProp(user),
		JueXueId:        jueXueId,
		BeforeLv:        beforeLv,
		Lv:              afterLv,
	}
	writeUserEvent(user, "jue_xue_up", data)
}
