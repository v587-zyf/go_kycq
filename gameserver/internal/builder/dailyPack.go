package builder

import "cqserver/gameserver/internal/objs"

func BuildDailyPack(user *objs.User) map[int32]int32 {
	pbMap := make(map[int32]int32)
	for _, unit := range user.DailyPack {
		for k, v := range unit.BuyIds {
			pbMap[int32(k)] = int32(v)
		}
	}
	return pbMap
}
