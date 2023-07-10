package builder

import (
	"cqserver/gamelibs/publicCon/constShop"
)

func BuildTreasureShop(shop map[int]int) map[int32]bool {
	pbMap := make(map[int32]bool)
	if shop != nil {
		for id, addFlag := range shop {
			flag := false
			if addFlag == constShop.SHOP_ADD_YES {
				flag = true
			}
			pbMap[int32(id)] = flag
		}
	}
	return pbMap
}

func BuildTreasureCar(car map[int]int) map[int32]int32 {
	pbMap := make(map[int32]int32)
	if car != nil {
		for shopId, num := range car {
			pbMap[int32(shopId)] = int32(num)
		}
	}
	return pbMap
}
