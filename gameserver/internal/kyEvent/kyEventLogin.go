package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_roleLogin struct {
	*kyEventLog.KyEventPropBase
	Diamond   int    `json:"diamond"`  //剩余付费货币数量	数值	是
	Client_ip string `json:"client_ip"` //用户的客户端ip
}

func UserLogin(user *objs.User) {
	data := &KyEvent_roleLogin{
		KyEventPropBase: getEventBaseProp(user),
		Diamond:         user.Ingot,
		Client_ip:       user.Ip,
	}
	writeUserEvent(user, "role_login", data)
}

func UserEnterServer(user *objs.User) {
	data := &KyEvent_roleLogin{
		KyEventPropBase: getEventBaseProp(user),
	}
	writeUserEvent(user, "enter_sid", data)
}
