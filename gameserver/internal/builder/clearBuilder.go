package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildEquipClears(hero *objs.Hero) map[int32]*pb.EquipClearArr {
	pbEquipClears := make(map[int32]*pb.EquipClearArr)
	heroClears := hero.EquipClear
	for pos, clearUnit := range heroClears {
		pbEquipClears[int32(pos)] = &pb.EquipClearArr{
			EquipClearInfo: BuildEquipClearUnit(clearUnit),
		}
	}
	return pbEquipClears
}

func BuildEquipClearUnit(clearUnit []*model.EquipClearUnit) []*pb.EquipClearInfo {
	pbArr := make([]*pb.EquipClearInfo, 0)
	for _, equipClearUnit := range clearUnit {
		pbArr = append(pbArr, &pb.EquipClearInfo{
			Grade:  int32(equipClearUnit.Grade),
			Color:  int32(equipClearUnit.Color),
			PropId: int32(equipClearUnit.PropId),
			Value:  int32(equipClearUnit.Value),
		})
	}
	return pbArr
}


