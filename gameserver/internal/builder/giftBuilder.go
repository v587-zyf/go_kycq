package builder

import "cqserver/gameserver/internal/objs"

func BuildOpenGift(user *objs.User) map[int32]int32 {
	pbMap := make(map[int32]int32)
	for id, buyNum := range user.OpenGift {
		pbMap[int32(id)] = int32(buyNum)
	}
	return pbMap
}
