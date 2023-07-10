package builder

import "cqserver/gameserver/internal/objs"

func BuildVipBoss(user *objs.User) map[int32]int32 {
	userVipBoss := user.VipBosses
	pbVipBoss := make(map[int32]int32)
	for stageId, dareNum := range userVipBoss.DareNum {
		pbVipBoss[int32(stageId)] = int32(dareNum)
	}
	return pbVipBoss
}
