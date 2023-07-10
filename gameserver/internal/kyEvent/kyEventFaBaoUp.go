package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_fa_bao_up struct {
	*kyEventLog.KyEventPropBase
	FaBaoId  int `json:"config_id1"` //法宝id
	BeforeLv int `json:"before_lv"`  //升级前等级
	Lv       int `json:"lv"`         //升级后等级
}

//法宝升级记录
func FaBaoUp(user *objs.User, faBaoId, beforeLv, afterLv int) {
	data := &KyEvent_fa_bao_up{
		KyEventPropBase: getEventBaseProp(user),
		FaBaoId:         faBaoId,
		BeforeLv:        beforeLv,
		Lv:              afterLv,
	}
	writeUserEvent(user, "fa_bao_up", data)
}
