package builder

import "cqserver/gameserver/internal/objs"

func BuildDragonEquip(hero *objs.Hero) map[int32]int32 {
	pbDragonEquip := make(map[int32]int32)
	for id, lv := range hero.DragonEquip {
		pbDragonEquip[int32(id)] = int32(lv)
	}
	return pbDragonEquip
}
