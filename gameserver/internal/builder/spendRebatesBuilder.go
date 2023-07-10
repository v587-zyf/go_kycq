package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildSpendRebates(user *objs.User) *pb.SpendRebates {
	userSpendRebates := user.SpendRebates
	return &pb.SpendRebates{
		CountIngot: int32(userSpendRebates.CountIngot),
		Ingot:      int32(userSpendRebates.Ingot),
		Reward:     BuildSpendRebatesReward(userSpendRebates.Reward),
		Cycle:      int32(userSpendRebates.Cycle),
	}
}

func BuildSpendRebatesReward(reward model.IntKv) []int32 {
	pbSlice := make([]int32, 0)
	for id := range reward {
		pbSlice = append(pbSlice, int32(id))
	}
	return pbSlice
}
