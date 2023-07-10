package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildBagInfoAck(user *objs.User) *pb.BagInfoNtf {

	ack := &pb.BagInfoNtf{}
	for _, v := range user.Bag {
		if v.ItemId > 0 {
			ack.Items = append(ack.Items, BuildPbItem(user, v))
		}
	}
	ack.BagMax = int32(user.BagInfo[constBag.BAG_TYPE_COMMON].MaxNum)
	ack.HaveBuyTimes = int32(user.BagInfo[constBag.BAG_TYPE_COMMON].BuyNum)
	return ack
}

func BuildPbItem(user *objs.User, item *model.Item) *pb.Item {

	pbItem := &pb.Item{
		ItemId:   int32(item.ItemId),
		Count:    int64(item.Count),
		Position: int32(item.Position),
	}
	if item.EquipIndex > 0 {
		if user.EquipBag[item.EquipIndex] != nil {
			pbItem.Equip = BuildPbEquipUnit(user.EquipBag[item.EquipIndex])
		}
	}
	return pbItem
}

func BuildPbEquipUnit(equip *model.Equip) *pb.EquipUnit {
	randProps := make([]*pb.EquipRandProp, len(equip.RandProps))
	for kk, vv := range equip.RandProps {
		randProps[kk] = &pb.EquipRandProp{
			PropId: int32(vv.PropId),
			Color:  int32(vv.Color),
			Value:  int32(vv.Value),
		}
	}
	pbEquipUnit := &pb.EquipUnit{
		ItemId:     int32(equip.ItemId),
		RandProps:  randProps,
		Lock:       equip.IsLock,
		Lucky:      int32(equip.Lucky),
		EquipIndex: int32(equip.Index),
	}
	return pbEquipUnit
}

//--------------仓库相关
func BuildWarehouseBagInfoAck(user *objs.User) *pb.WarehouseInfoNtf {

	ack := &pb.WarehouseInfoNtf{}
	for _, v := range user.WarehouseBag {
		if v.ItemId > 0 {
			ack.Items = append(ack.Items, BuildPbItem(user, v))
		}
	}
	ack.HaveBuyTimes = int32(user.WarehouseBagInfo[constBag.WAREHOUSE_BAG_TYPE_COMMON].BuyNum)
	ack.BagMax = int32(user.WarehouseBagInfo[constBag.WAREHOUSE_BAG_TYPE_COMMON].MaxNum)
	return ack
}
