package builder

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuilderFieldBoss(user *objs.User) *pb.FieldBossInfo {
	userFieldBoss := user.FieldBoss
	return &pb.FieldBossInfo{
		DareNum:      int32(userFieldBoss.DareNum),
		BuyNum:       int32(userFieldBoss.BuyNum),
		FirstReceive: userFieldBoss.FirstReceive,
	}
}
