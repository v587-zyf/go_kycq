package builder

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildUserWear(user *objs.User) *pb.UserWear {
	userWear := user.Wear
	return &pb.UserWear{
		Petid:        int32(userWear.PetId),
		FitFashionId: int32(userWear.FitFashionId),
	}
}

func BuildWear(hero *objs.Hero) *pb.Wears {
	heroWear := hero.Wear
	return &pb.Wears{
		FashionWeaponId: int32(heroWear.FashionWeaponId),
		FashionClothId:  int32(heroWear.FashionClothId),
		AtlasWear:       BuildAtlasWear(heroWear.AtlasWear),
		WingId:          int32(heroWear.WingId),
		MagicCircleLvId: int32(heroWear.MagicCircleLvId),
		TitleId:         int32(heroWear.TitleId),
	}
}

func BuildAtlasWear(atlasWear map[int]int) []int32 {
	pbSlice := make([]int32, 0)
	for id := range atlasWear {
		pbSlice = append(pbSlice, int32(id))
	}
	return pbSlice
}
