package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildGrowFund(user *objs.User) *pb.GrowFund {
	userGrowFund := user.GrowFund
	return &pb.GrowFund{
		IsBuy: userGrowFund.IsBuy,
		Ids:   BuildGrowFundIds(userGrowFund.Ids),
	}
}

func BuildGrowFundIds(ids model.IntKv) []int32 {
	pbSlice := make([]int32, 0)
	for id := range ids {
		pbSlice = append(pbSlice, int32(id))
	}
	return pbSlice
}
