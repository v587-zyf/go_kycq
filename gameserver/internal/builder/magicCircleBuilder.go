package builder

import "cqserver/gameserver/internal/objs"

func BuildMagicCircle(hero *objs.Hero) map[int32]int32 {
	magicCircle := hero.MagicCircle
	pbMagicCircle := make(map[int32]int32)
	for id, excelId := range magicCircle {
		pbMagicCircle[int32(id)] = int32(excelId)
	}
	return pbMagicCircle
}
