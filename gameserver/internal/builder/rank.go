package builder

import (
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/protobuf/pb"
)

func BuildRankUserInfo(user *modelGame.UserBasicInfo, heroIndex int, rank, score int) *pb.RankInfo {

	rankInfo := &pb.RankInfo{
		Rank:  int32(rank),
		Score: int64(score),
	}
	if user == nil {
		return rankInfo
	}

	rankInfo.UserInfo = &pb.BriefUserInfo{}
	rankInfo.UserInfo.Id = int32(user.Id)
	rankInfo.UserInfo.Name = user.NickName
	//rankInfo.UserInfo.Lvl = int32(user.Lvl)
	rankInfo.UserInfo.Avatar = user.Avatar

	if heroIndex == - 1 {
		heroIndex = constUser.USER_HERO_MAIN_INDEX
	}

	if _, ok := user.HeroDisplay[heroIndex]; !ok {
		return rankInfo
	}

	rankInfo.Display = &pb.Display{
		ClothItemId:     int32(user.HeroDisplay[heroIndex].Display.ClothItemId),
		ClothType:       int32(user.HeroDisplay[heroIndex].Display.ClothType),
		WeaponItemId:    int32(user.HeroDisplay[heroIndex].Display.WeaponItemId),
		WeaponType:      int32(user.HeroDisplay[heroIndex].Display.WeaponType),
		WingId:          int32(user.HeroDisplay[heroIndex].Display.WingId),
		MagicCircleLvId: int32(user.HeroDisplay[heroIndex].Display.MagicCircleLvId),
	}
	return rankInfo
}
