package builder

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuilderEquipStrength( hero *objs.Hero )[]*pb.EquipGrid{

	equipGrids := make([]*pb.EquipGrid,0)
	for k,v := range hero.EquipsStrength {
		equipGrids = append(equipGrids,&pb.EquipGrid{
			Pos: int32(k),
			Strength:int32(v),
		})
	}
	return equipGrids
}
