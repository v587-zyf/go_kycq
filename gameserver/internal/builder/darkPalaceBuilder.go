package builder

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildDarkPalace(user *objs.User) *pb.DarkPalaceInfo {
	userDarkPalace := user.DarkPalace
	return &pb.DarkPalaceInfo{
		DareNum: int32(userDarkPalace.DareNum),
		BuyNum:  int32(userDarkPalace.BuyNum),
		HelpNum: int32(userDarkPalace.HelpNum),
	}
}
