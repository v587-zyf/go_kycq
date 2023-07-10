package kyEvent

import "cqserver/gameserver/internal/base"

type KyEvent_Ranking_List struct {
	Plat     int    `json:"plat"`            //平台	数值	是
	Channel  int    `json:"channel"`         //渠道	数值	是
	Serverid int    `json:"_serverid"`       //	区服	数值	是
	Type     int    `json:"type"`            //排行榜类型 数值
	RankName string `json:"log_event_parm1"` //排行榜名称
	Info     string `json:"info"`            //排行榜信息 如：[{"roleid":"222","score":"100","rank":1},{"roleid":"222","score":"100","rank":2}]
}

func RankingList(rankT int, rankName, info string) {
	data := &KyEvent_Ranking_List{
		Plat:     base.Conf.Sdkconfig.KyPlatformId,
		Channel:  0,
		Serverid: base.Conf.ServerId,
		Type:     rankT,
		RankName: rankName,
		Info:     info,
	}
	writeUserEvent(nil, "ranking_list", data)
}
