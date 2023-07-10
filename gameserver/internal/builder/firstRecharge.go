package builder

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildFirstRecharge(user *objs.User) *pb.FirstRecharge {
	userFirstRecharge := user.FirstRecharge
	return &pb.FirstRecharge{
		IsRecharge: userFirstRecharge.IsRecharge,
		Days:       BuildFirstRechargeDay(userFirstRecharge.Days),
		OpenDay:    int64(userFirstRecharge.OpenDay),
	}
}

func BuildFirstRechargeDay(days map[int]int) []int32 {
	pbSlice := make([]int32, 0)
	for day := range days {
		pbSlice = append(pbSlice, int32(day))
	}
	return pbSlice
}
