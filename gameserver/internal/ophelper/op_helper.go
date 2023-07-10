package ophelper

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"strconv"
)

type OpBagHelperDefault struct {
	opType       int
	opTypeSecond int //二级来源
	//items        []*pb.ItemChange
	equips      []*pb.EquipChange
	topData     []*pb.TopDataChange
	changeItems []*pb.ItemUnit
}

func (op *OpBagHelperDefault) ToGoodsChangeMessages() []nw.ProtoMessage {
	msgs := make([]nw.ProtoMessage, 0)
	if len(op.topData) > 0 {
		msg := &pb.TopDataChangeNtf{
			Type:        int32(op.opType),
			ChangeInfos: op.topData,
		}
		msgs = append(msgs, msg)
	}
	if len(op.equips) > 0 {
		msg := &pb.BagEquipDataChangeNtf{
			Type:        int32(op.opType),
			ChangeInfos: op.equips,
		}
		msgs = append(msgs, msg)
	}
	//if len(op.items) > 0 {
	//	msg := &pb.BagDataChangeNtf{
	//		Type:        int32(op.opType),
	//		ChangeInfos: op.items,
	//	}
	//	msgs = append(msgs, msg)
	//}
	return msgs
}

func (op *OpBagHelperDefault) ToChangeItems() *pb.GoodsChangeNtf {

	return &pb.GoodsChangeNtf{
		Items: op.changeItems,
	}

}

func (op *OpBagHelperDefault) SetOpType(opType int) {
	op.opType = opType
}

func (op *OpBagHelperDefault) GetOpType() int {
	return op.opType
}

func (op *OpBagHelperDefault) OpTypeSecond() int {
	return op.opTypeSecond
}

func (op *OpBagHelperDefault) SetOpTypeSecond(opTypeSecond int) {
	op.opTypeSecond = opTypeSecond
}

func (op *OpBagHelperDefault) OnGoodsChange(changeUnit interface{}, count int) {
	switch changeUnit.(type) {
	case *pb.ItemChange:

		itemChange := changeUnit.(*pb.ItemChange)
		has := false
		for _, v := range op.equips {
			if v.Position == itemChange.Position && v.ItemId == itemChange.ItemId {
				v.Change += itemChange.Change
				v.NowNum = itemChange.NowNum
				has = true
			}
		}
		if !has {
			change := builder.BuildEquipDataChagne(int(itemChange.ItemId), int(itemChange.Change), int(itemChange.NowNum), int(itemChange.Position), nil)
			op.equips = append(op.equips, change)
		}
		op.changeItems = append(op.changeItems, &pb.ItemUnit{ItemId: itemChange.ItemId, Count: int64(count)})
	case *pb.EquipChange:

		itemChange := changeUnit.(*pb.EquipChange)
		op.equips = append(op.equips, itemChange)
		op.changeItems = append(op.changeItems, &pb.ItemUnit{ItemId: itemChange.ItemId, Count: int64(count)})

	case *pb.TopDataChange:
		topDataChange := changeUnit.(*pb.TopDataChange)
		has := false
		for _, v := range op.topData {
			if v.Id == topDataChange.Id {
				v.Change += topDataChange.Change
				v.NowNum = topDataChange.NowNum
				has = true
			}
		}
		if !has {
			op.topData = append(op.topData, topDataChange)
		}
		if topDataChange.Id != pb.ITEMID_VIP_LV && topDataChange.Id != pb.ITEMID_LV {
			op.changeItems = append(op.changeItems, &pb.ItemUnit{ItemId: topDataChange.Id, Count: int64(count)})
		}
	}

	return
}

func (op *OpBagHelperDefault) BuildItemGetDisplay(itemId int, count int) {

	op.changeItems = append(op.changeItems, &pb.ItemUnit{ItemId: int32(itemId), Count: int64(count)})
}

func NewOpBagHelperDefault(opType int) *OpBagHelperDefault {
	return &OpBagHelperDefault{opType: opType,
		//items:   make([]*pb.ItemChange, 0),
		topData: make([]*pb.TopDataChange, 0),
	}
}

func CreateGoodsChangeNtf(items map[int]int) *pb.GoodsChangeNtf {

	pbItems := make([]*pb.ItemUnit, 0)
	for k, v := range items {
		pbItems = append(pbItems, &pb.ItemUnit{ItemId: int32(k), Count: int64(v)})
	}

	return &pb.GoodsChangeNtf{
		Items: pbItems,
	}
}

func GetBagFullMailItemSource(reason1, reason2 int) int {
	return reason2*10000 + reason1
}

func GetMailOpType(mailConfId int) int {
	return mailConfId * 10000
}

func GetResaon(reason, reason2 int) (string, string) {

	var reasonLv1 string
	var reasonLv2 string
	reasonFix := reason / 10000
	if reasonFix > 0 {
		mailConf := gamedb.GetMailMailCfg(reasonFix)
		if mailConf != nil {
			reasonLv1 = "邮件_" + mailConf.Title
		}
		if reason2 > 0 {
			reason = reason2 % 10000
			reason2 = reason2 / 10000
		} else {
			return reasonLv1, reasonLv2
		}
	}
	if r, ok := constBag.OPTYPE_SOURCE_MAP[reason]; ok {
		if len(reasonLv1) > 0 {
			reasonLv1 += ":" + r
		} else {
			reasonLv1 += r
		}

		if reason == constBag.OpTypeFightCheer || reason == constBag.OpTypeFightPotion ||
			reason == constBag.OpTypeCollection || reason == constBag.OpTypePickUp || reason == constBag.OpTypeHelp {
			stageConf := gamedb.GetStageStageCfg(reason2)
			if stageConf != nil {
				reasonLv2 = stageConf.Name
			} else {
				reasonLv2 = strconv.Itoa(reason2)
			}
		} else {
			reasonLv2 = strconv.Itoa(reason2)
		}

	} else {
		reasonLv1 = strconv.Itoa(reason)
		reasonLv2 = strconv.Itoa(reason2)
	}
	return reasonLv1, reasonLv2
}
