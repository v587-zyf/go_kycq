package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuilderHeroGodEquip(hero *objs.Hero) map[int32]*pb.GodEquip {
	heroGodEquips := make(map[int32]*pb.GodEquip, 0)
	for _, godEquip := range hero.GodEquips {
		heroGodEquips[int32(godEquip.Id)] = BuilderGodEquipUnit(godEquip)
	}
	return heroGodEquips
}

func BuilderGodEquipUnit(godEquip *model.GodEquip) *pb.GodEquip {
	return &pb.GodEquip{
		Id:    int32(godEquip.Id),
		Level: int32(godEquip.Lv),
		Blood: int32(godEquip.Blood),
	}
}
