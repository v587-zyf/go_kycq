package builder

import (
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/objs"
)

func BuildPersonBoss(user *objs.User) map[int32]int32 {
	userPersonBoss := user.PersonBosses.DareNum
	pbPersonBoss := make(map[int32]int32)
	for stageId := range userPersonBoss {
		pbPersonBoss[int32(stageId)] = int32(rmodel.Boss.GetBossKillNum(user.Id, constFight.FIGHT_TYPE_PERSON_BOSS, stageId))
	}
	return pbPersonBoss
}