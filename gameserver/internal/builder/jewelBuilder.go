package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuilderJewels(hero *objs.Hero) map[int32]*pb.JewelInfo {
	pbJewels := make(map[int32]*pb.JewelInfo)
	for pos, jewel := range hero.Jewel {
		pbJewels[int32(pos)] = BuildJewel(jewel)
	}
	return pbJewels
}

func BuildJewel(jewel *model.Jewel) *pb.JewelInfo {
	return &pb.JewelInfo{
		Left:  int32(jewel.One),
		Right: int32(jewel.Two),
		Down:  int32(jewel.Three),
	}
}
