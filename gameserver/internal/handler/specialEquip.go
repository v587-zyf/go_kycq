package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"errors"
)

func init() {
	pb.Register(pb.CmdSpecialEquipChangeReqId, HandlerSpecialEquipChangeReq)
	pb.Register(pb.CmdSpecialEquipRemoveReqId, HandlerSpecialEquipRemoveReq)
}

func HandlerSpecialEquipChangeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.SpecialEquipChangeReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.SpecialEquipChangeAck{
		HeroIndex: req.HeroIndex,
		Pos:       req.Pos,
		Type:      req.Type,
	}

	var err error
	var op *ophelper.OpBagHelperDefault
	switch int(req.Type) {
	case pb.SPECIALEQUIPTYPE_ZODIAC:
		op = ophelper.NewOpBagHelperDefault(constBag.OpTypeZodiac)
		err, ack.SpecialEquip = m.Zodiac.ZodiacChange(user, int(req.HeroIndex), int(req.Pos), int(req.BagPos), op)
		if err != nil {
			return nil, nil, err
		}
	//case pb.SPECIALEQUIPTYPE_DRAGONARMS:
	//	op = ophelper.NewOpBagHelperDefault(constBag.OpTypeDragonarms)
	//	err, ack.SpecialEquip = m.Dragonarms.DragonarmsChange(user, int(req.HeroIndex), int(req.Pos), int(req.BagPos), op)
	//	if err != nil {
	//		return nil, nil, err
	//	}
	case pb.SPECIALEQUIPTYPE_KINGARMS:
		op = ophelper.NewOpBagHelperDefault(constBag.OpTypeKingarms)
		err, ack.SpecialEquip = m.Kingarms.KingarmsChange(user, int(req.HeroIndex), int(req.Pos), int(req.BagPos), op)
		if err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, errors.New("error specialEquip type")
	}
	ack.Goods = op.ToChangeItems()

	m.UserManager.SendItemChangeNtf(user, op)

	return ack, nil, nil
}

func HandlerSpecialEquipRemoveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.SpecialEquipRemoveReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.SpecialEquipRemoveAck{
		HeroIndex: req.HeroIndex,
		Pos:       req.Pos,
		Type:      req.Type,
	}

	var err error
	var op *ophelper.OpBagHelperDefault
	switch int(req.Type) {
	case pb.SPECIALEQUIPTYPE_ZODIAC:
		op = ophelper.NewOpBagHelperDefault(constBag.OpTypeZodiac)
		err = m.Zodiac.ZodiacRemove(user, int(req.HeroIndex), int(req.Pos), op)
		if err != nil {
			return nil, nil, err
		}
	//case pb.SPECIALEQUIPTYPE_DRAGONARMS:
	//	op = ophelper.NewOpBagHelperDefault(constBag.OpTypeDragonarms)
	//	err = m.Dragonarms.DragonarmsRemove(user, int(req.HeroIndex), int(req.Pos), op)
	//	if err != nil {
	//		return nil, nil, err
	//	}
	case pb.SPECIALEQUIPTYPE_KINGARMS:
		op = ophelper.NewOpBagHelperDefault(constBag.OpTypeKingarms)
		err = m.Kingarms.KingarmsRemove(user, int(req.HeroIndex), int(req.Pos), op)
		if err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, errors.New("error specialEquip type")
	}
	ack.Goods = op.ToChangeItems()

	m.UserManager.SendItemChangeNtf(user, op)

	return ack, nil, nil
}
