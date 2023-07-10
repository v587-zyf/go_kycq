package builder

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildAncient(user *objs.User) *pb.AncientBossInfo {
	userAncient := user.AncientBoss
	return &pb.AncientBossInfo{
		DareNum: int32(userAncient.DareNum),
		BuyNum:  int32(userAncient.BuyNum),
	}
}
