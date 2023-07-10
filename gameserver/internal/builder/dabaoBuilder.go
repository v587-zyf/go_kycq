package builder

import "cqserver/gameserver/internal/objs"

func BuildDaBaoEquip(user *objs.User) map[int32]int32 {
	pbMap := make(map[int32]int32)
	for t, lv := range user.DaBaoEquip {
		pbMap[int32(t)] = int32(lv)
	}
	return pbMap
}
