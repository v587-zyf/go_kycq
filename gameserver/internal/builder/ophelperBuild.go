package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/protobuf/pb"
)

func BuildTopDataChange(itemId, count, nowNum int) *pb.TopDataChange {

	return &pb.TopDataChange{Id: int32(itemId), Change: int64(count), NowNum: int64(nowNum)}
}

func BuildItemDataChange(itemId, change, nowNum, pos int) *pb.ItemChange {
	return &pb.ItemChange{ItemId: int32(itemId), Change: int64(change), NowNum: int64(nowNum), Position: int32(pos)}
}

func BuildEquipDataChagne(itemId, change, nowNum, pos int,equip *model.Equip) *pb.EquipChange {
	equipChange := &pb.EquipChange{ItemId: int32(itemId), Change: int64(change), NowNum: int64(nowNum), Position: int32(pos)}
	if equip != nil {
		equipChange.Equip = BuildPbEquipUnit(equip)
	}
	return equipChange
}
