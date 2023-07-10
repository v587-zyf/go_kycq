package managersI

import (
	"cqserver/gamelibs/modelCross"
)

type IActiveUser interface {
	ActiveUsersAdd(message *modelCross.UserCrossInfo)

	//day:活跃天数
	GetUserIdsByActiveDay(day int) map[int]bool
}
