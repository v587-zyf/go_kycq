package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_panacea_use struct {
	*kyEventLog.KyEventPropBase
	Id        int `json:"config_id1"` //灵丹id
	BeforeNum int `json:"num1"`       //升级前数量
	Num       int `json:"num2"`       //升级后数量
}

//灵丹使用记录
func PanaceaUse(user *objs.User, id, num int) {
	data := &KyEvent_panacea_use{
		KyEventPropBase: getEventBaseProp(user),
		Id:              id,
		BeforeNum:       num - 1,
		Num:             num,
	}
	writeUserEvent(user, "panacea_use", data)
}
