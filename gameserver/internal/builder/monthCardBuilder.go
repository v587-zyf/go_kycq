package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
	"time"
)

func BuildMonthCard(user *objs.User) map[int32]*pb.MonthCardUnit {
	pbMap := make(map[int32]*pb.MonthCardUnit)
	userMonthCard := user.MonthCard
	for t, unit := range userMonthCard.MonthCards {
		pbMap[int32(t)] = BuildMonthCardUnit(unit)
	}
	user.Dirty = true
	return pbMap
}

func BuildMonthCardUnit(unit *model.MonthCardUnit) *pb.MonthCardUnit {
	isExpire := false
	if unit.EndTime != 0 && unit.EndTime != -1 && unit.EndTime <= int(time.Now().Unix()) {
		isExpire = true
		unit.EndTime = 0
	}
	return &pb.MonthCardUnit{
		StartTime: int64(unit.StartTime),
		EndTime:   int64(unit.EndTime),
		IsExpire:  isExpire,
	}
}
