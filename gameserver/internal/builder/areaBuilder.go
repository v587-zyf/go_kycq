package builder

import "cqserver/gameserver/internal/objs"

func BuildArea(hero *objs.Hero) map[int32]int32 {
	pbArea := make(map[int32]int32)
	heroArea := hero.Area
	for id, lv := range heroArea {
		pbArea[int32(id)] = int32(lv)
	}
	return pbArea
}
