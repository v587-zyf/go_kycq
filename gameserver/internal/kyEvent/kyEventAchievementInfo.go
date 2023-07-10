package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_achievement_info struct {
	*kyEventLog.KyEventPropBase
	AchievementType int `json:"type"`
	AchievementId   int `json:"config_id1"`
}

//成就完成记录
func AchievementInfo(user *objs.User, achievementType, achievementId int) {
	data := &KyEvent_achievement_info{
		KyEventPropBase: getEventBaseProp(user),
		AchievementType: achievementType,
		AchievementId:   achievementId,
	}
	writeUserEvent(user, "achievement_info", data)
}
