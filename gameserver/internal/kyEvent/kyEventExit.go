package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
	"time"
)

type KyEvent_roleExit struct {
	*kyEventLog.KyEventPropBase
	Diamond       int    `json:"diamond"`       //剩余付费货币数量	数值	是
	Client_ip     string `json:"client_ip"`      //用户的客户端ip
	Page_staytime int    `json:"page_staytime"` //在线时长
}

func UserExit(user *objs.User) {
	data := &KyEvent_roleExit{
		KyEventPropBase: getEventBaseProp(user),
		Diamond:         user.Ingot,
		Client_ip:       user.Ip,
		Page_staytime:   int(time.Now().Sub(user.LoginTime).Seconds()),
	}
	writeUserEvent(user, "role_exit", data)
}
