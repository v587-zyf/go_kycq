package builder

import "cqserver/gameserver/internal/objs"

func BuildChuanShiEquip(hero *objs.Hero) map[int32]int32 {
	pbMap := make(map[int32]int32)
	for pos, id := range hero.ChuanShi {
		pbMap[int32(pos)] = int32(id)
	}
	return pbMap
}

func BuildChuanShiStrengthen(hero *objs.Hero) map[int32]int32 {
	pbMap := make(map[int32]int32)
	for pos, lv := range hero.ChuanshiStrengthen {
		pbMap[int32(pos)] = int32(lv)
	}
	return pbMap
}