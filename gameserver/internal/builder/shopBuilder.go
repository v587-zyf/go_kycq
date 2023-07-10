package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildShop(user *objs.User) map[int32]*pb.ShopInfo {
	userShop := user.Shops
	pbShopInfo := make(map[int32]*pb.ShopInfo)
	for shopType, ShopInfo := range userShop.ShopItem {
		pbShopInfo[int32(shopType)] = BuildShopByType(ShopInfo)
	}
	return pbShopInfo
}

func BuildShopByType(ShopInfo model.IntKv) *pb.ShopInfo {
	pbShopInfo := make(map[int32]int32)
	for id, buyNum := range ShopInfo {
		pbShopInfo[int32(id)] = int32(buyNum)
	}
	return &pb.ShopInfo{ShopItem: pbShopInfo}
}
