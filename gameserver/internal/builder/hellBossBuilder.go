package builder

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildHellBoss(user *objs.User) *pb.HellBossInfo {
	userHellBoss := user.HellBoss
	return &pb.HellBossInfo{
		DareNum: int32(userHellBoss.DareNum),
		BuyNum:  int32(userHellBoss.BuyNum),
		HelpNum: int32(userHellBoss.HelpNum),
	}
}
