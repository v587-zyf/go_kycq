package builder

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildSignInfo(user *objs.User) *pb.SignInfo {
	return &pb.SignInfo{
		SignCount: int32(user.Sign.Count),
		SignDay:   BuildSignDay(user.Sign.SignDay),
		CumulativeDay: BuildCumulativeDay(user.Sign.Cumulative),
	}
}

func BuildSignDay(signDay map[int]int) map[int32]int32 {
	maps := make(map[int32]int32)
	for day, date := range signDay {
		maps[int32(day)] = int32(date)
	}
	return maps
}

func BuildCumulativeDay(cumulativeDay map[int]int) map[int32]int32 {
	maps := make(map[int32]int32)
	for day := range cumulativeDay {
		maps[int32(day)] = int32(day)
	}
	return maps
}
