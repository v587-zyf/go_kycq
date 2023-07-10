package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

type KyEvent_cross_challenge struct {
	*kyEventLog.KyEventPropBase
	Round int  `json:"num1"`  //轮次
	IsWin int `json:"status"` //轮次结果   //true 赢
}

//跨服擂台赛记录
func CrossChallenge(user *objs.User, round int, roundResult bool) {
	data := &KyEvent_cross_challenge{
		KyEventPropBase: getEventBaseProp(user),
		Round:           round,
		IsWin:           0,
	}
	if roundResult {
		data.IsWin = pb.RESULTFLAG_SUCCESS
	}
	writeUserEvent(user, "cross_challenge", data)
}
