package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_shaBake struct {
	*kyEventLog.KyEventPropBase
	WinGuildId      int             `json:"info"`       //占领门派id
	WinGuildUserIds gamedb.IntSlice `json:"log_event_parm1"` //占领帮会的 会长 和副会长 id

}

//沙巴克记录
func ShaBake(user *objs.User, winGuildId int, winGuildUserIds []int) {
	data := &KyEvent_shaBake{
		KyEventPropBase: getEventBaseProp(user),
		WinGuildId:      winGuildId,
		WinGuildUserIds: winGuildUserIds,
	}
	writeUserEvent(user, "shaBake", data)
}
