package user

import (
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/protobuf/pb"
)

func (this *UserManager) BuildUserRankInfo(userId int, heroIndex int, rank, score int) *pb.RankInfo {

	rankInfo := &pb.RankInfo{
		Rank:     int32(rank),
		Score:    int64(score),
		UserInfo: &pb.BriefUserInfo{},
		Display:  &pb.Display{},
	}
	user := this.GetUser(userId)
	if user != nil {
		rankInfo.UserInfo = &pb.BriefUserInfo{}
		rankInfo.UserInfo.Id = int32(user.Id)
		rankInfo.UserInfo.Name = user.NickName
		rankInfo.UserInfo.Avatar = user.Avatar

		if heroIndex == - 1 {
			heroIndex = constUser.USER_HERO_MAIN_INDEX
		}

		if _, ok := user.Heros[heroIndex]; !ok {
			return rankInfo
		}

		rankInfo.UserInfo.Sex = int32(user.Heros[heroIndex].Sex)
		rankInfo.UserInfo.Job = int32(user.Heros[heroIndex].Job)
		rankInfo.Display = &pb.Display{
			ClothItemId:     int32(user.Heros[heroIndex].Display.ClothItemId),
			ClothType:       int32(user.Heros[heroIndex].Display.ClothType),
			WeaponItemId:    int32(user.Heros[heroIndex].Display.WeaponItemId),
			WeaponType:      int32(user.Heros[heroIndex].Display.WeaponType),
			WingId:          int32(user.Heros[heroIndex].Display.WingId),
			MagicCircleLvId: int32(user.Heros[heroIndex].Display.MagicCircleLvId),
		}
	} else {
		userBaseInfo := this.GetUserBasicInfo(userId)
		if userBaseInfo != nil {
			rankInfo.UserInfo = &pb.BriefUserInfo{}
			rankInfo.UserInfo.Id = int32(userBaseInfo.Id)
			rankInfo.UserInfo.Name = userBaseInfo.NickName
			//rankInfo.UserInfo.Lvl = int32(userBaseInfo.Lvl)
			rankInfo.UserInfo.Avatar = userBaseInfo.Avatar

			if heroIndex == - 1 {
				heroIndex = constUser.USER_HERO_MAIN_INDEX
			}

			if _, ok := userBaseInfo.HeroDisplay[heroIndex]; !ok {
				return rankInfo
			}
			rankInfo.UserInfo.Sex = int32(userBaseInfo.HeroDisplay[heroIndex].Sex)
			rankInfo.UserInfo.Job = int32(userBaseInfo.HeroDisplay[heroIndex].Job)
			rankInfo.Display = &pb.Display{
				ClothItemId:     int32(userBaseInfo.HeroDisplay[heroIndex].Display.ClothItemId),
				ClothType:       int32(userBaseInfo.HeroDisplay[heroIndex].Display.ClothType),
				WeaponItemId:    int32(userBaseInfo.HeroDisplay[heroIndex].Display.WeaponItemId),
				WeaponType:      int32(userBaseInfo.HeroDisplay[heroIndex].Display.WeaponType),
				WingId:          int32(userBaseInfo.HeroDisplay[heroIndex].Display.WingId),
				MagicCircleLvId: int32(userBaseInfo.HeroDisplay[heroIndex].Display.MagicCircleLvId),
			}
		}
	}
	return rankInfo

}
