package managersI

import "cqserver/gameserver/internal/objs"

type IGuardPillar interface {
	In(user *objs.User, stageId int) error
	GuardPillarResult(userId, stageId, rounds, rank int)
}
