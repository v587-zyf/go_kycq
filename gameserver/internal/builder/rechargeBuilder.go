package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/protobuf/pb"
	"time"
)

func BuildRecharge(user *objs.User) []int32 {
	pbSlice := make([]int32, 0)
	userRecharge := user.Recharge
	for id := range userRecharge {
		pbSlice = append(pbSlice, int32(id))
	}
	return pbSlice
}

func BuildContRecharge(user *objs.User) *pb.ContRecharge {
	userContRecharge := user.ContRecharge
	return &pb.ContRecharge{
		Cycle:    int32(userContRecharge.Cycle),
		Recharge: BuildContRechargeRecharge(userContRecharge.Day),
		Receive:  BuildContRechargeReceive(userContRecharge.Receive),
		TodayPay: int32(userContRecharge.Day[common.GetResetTime(time.Now())]),
	}
}
func BuildContRechargeRecharge(recharge model.IntKv) map[int32]int32 {
	pbMap := make(map[int32]int32)
	for day, payNum := range recharge {
		pbMap[int32(day)] = int32(payNum)
	}
	return pbMap
}
func BuildContRechargeReceive(receive model.IntKv) []int32 {
	pbSlice := make([]int32, 0)
	for id := range receive {
		pbSlice = append(pbSlice, int32(id))
	}
	return pbSlice
}
