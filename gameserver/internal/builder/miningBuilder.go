package builder

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildMining(user *objs.User) *pb.MiningInfo {
	userMining := user.Mining
	return &pb.MiningInfo{
		WorkTime: int64(userMining.WorkTime),
		WorkNum:  int32(userMining.WorkNum),
		RobNum:   int32(userMining.RobNum),
		BuyNum:   int32(userMining.BuyNum),
		Miner:    int32(userMining.Miner),
		Luck:     int32(userMining.Luck),
	}
}
