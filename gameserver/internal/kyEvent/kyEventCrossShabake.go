package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_cross_shaBake struct {
	*kyEventLog.KyEventPropBase
	WinGuildId      int             `json:"info"`       //占领门派id
	WinGuildUserIds gamedb.IntSlice `json:"log_event_parm1"` //占领帮会的 会长 和副会长 id
	ServerScores    gamedb.IntMap   `json:"num1"`      //区服积分
}

//跨服沙巴克记录
func CrossShaBake(user *objs.User, winGuildId int, winGuildUserIds []int, serverScores map[int]int) {
	data := &KyEvent_cross_shaBake{
		KyEventPropBase: getEventBaseProp(user),
		WinGuildId:      winGuildId,
		WinGuildUserIds: winGuildUserIds,
		ServerScores:    serverScores,
	}
	writeUserEvent(user, "crossShaBake", data)
}
