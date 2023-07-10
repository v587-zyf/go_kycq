package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	//pb.Register(pb.CmdGodEquipActiveReqId, HandlerGodEquipActiveReq)
	pb.Register(pb.CmdGodEquipUpLevelReqId, HandlerGodEquipUpLevelReq)

	pb.Register(pb.CmdGodEquipBloodReqId, HandlerGodEquipBloodReq)
}

//func HandlerGodEquipActiveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//	req := p.(*pb.GodEquipActiveReq)
//	user := conn.GetSession().(*managers.ClientSession).User
//	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeGodEquipActive)
//
//	var err error
//	err = m.GodEquip.GodEquipActive(user, int(req.HeroIndex), int(req.Id), op)
//	if err != nil {
//		return nil, nil, err
//	}
//	ack := &pb.GodEquipActiveAck{
//		HeroIndex: req.HeroIndex,
//		GodEquip:  builder.BuilderGodEquipUnit(user.Heros[int(req.HeroIndex)].GodEquips[int(req.Id)]),
//	}
//
//	return ack, op, nil
//}

func HandlerGodEquipUpLevelReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.GodEquipUpLevelReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeGodEquipUpLv)

	if err := m.GodEquip.GodEquipUpLevel(user, int(req.HeroIndex), int(req.Id), op); err != nil {
		return nil, nil, err
	}
	return &pb.GodEquipUpLevelAck{
		HeroIndex: req.HeroIndex,
		GodEquip:  builder.BuilderGodEquipUnit(user.Heros[int(req.HeroIndex)].GodEquips[int(req.Id)]),
	}, op, nil
}

func HandlerGodEquipBloodReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.GodEquipBloodReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeGodEquipBlood)

	if err := m.GodEquip.GodEquipBlood(user, int(req.HeroIndex), int(req.GodEquipId), op); err != nil {
		return nil, nil, err
	}
	return &pb.GodEquipBloodAck{
		HeroIndex: req.HeroIndex,
		GodEquip:  builder.BuilderGodEquipUnit(user.Heros[int(req.HeroIndex)].GodEquips[int(req.GodEquipId)]),
	}, op, nil
}
