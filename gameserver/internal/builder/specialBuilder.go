package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuilderSpecialEquips(hero *objs.Hero, t int) map[int32]*pb.SpecialEquipUnit {
	specialEquips := make(map[int32]*pb.SpecialEquipUnit, 0)
	var heroInfos map[int]*model.SpecialEquipUnit
	switch t {
	case pb.ITEMTYPE_ZODIAC:
		heroInfos = hero.Zodiacs
	case pb.ITEMTYPE_KINGARMS:
		heroInfos = hero.Kingarms
	//case pb.ITEMTYPE_DRAGONARMS:
	//	heroInfos = hero.Dragonarms
	}
	for pos, unit := range heroInfos {
		specialEquips[int32(pos)] = BuilderSpecialEquipUnit(unit)
	}
	return specialEquips
}

func BuilderSpecialEquipUnit(unit *model.SpecialEquipUnit) *pb.SpecialEquipUnit {
	return &pb.SpecialEquipUnit{
		ItemId:    int32(unit.Id),
	}
}
