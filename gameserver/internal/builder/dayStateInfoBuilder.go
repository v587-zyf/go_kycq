package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildDayStateInfo(user *objs.User) *pb.DayStateInfo {
	userDayStateInfo := user.DayStateRecord
	return &pb.DayStateInfo{
		RankWorship:      int32(userDayStateInfo.RankWorship),
		MonthCardReceive: BuildMonthCardReceive(userDayStateInfo.MonthCardReceive),
	}
}

func BuildMonthCardReceive(data model.IntKv) []int32 {
	pbSlice := make([]int32, 0)
	for id := range data {
		pbSlice = append(pbSlice, int32(id))
	}
	return pbSlice
}
