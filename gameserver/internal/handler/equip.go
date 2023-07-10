package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdEquipChangeReqId, HandlerEquipChangeReq)
	pb.Register(pb.CmdEquipRemoveReqId, HandlerEquipRemoveReq)
	pb.Register(pb.CmdEquipStrengthenReqId, HandlerEquipStrengthenReq)
	pb.Register(pb.CmdEquipStrengthenAutoReqId, HandlerEquipStrengthenAutoReq)

	pb.Register(pb.CmdClearReqId, HandlerClearReq)
}

func HandlerEquipStrengthenReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EquipStrengthenReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeEquipStrengthen)

	var err error
	ack := &pb.EquipStrengthenAck{}

	err = m.Equip.EquipsStrengthen(user, int(req.HeroIndex), int(req.Pos), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("ack is %v", ack)

	return ack, op, nil
}

func HandlerEquipRemoveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EquipRemoveReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeEquipRemove)

	var err error
	ack := &pb.EquipRemoveAck{}
	err = m.Equip.EquipRemove(user, int(req.HeroIndex), int(req.Pos), op, ack)
	if err != nil {
		return nil, nil, err
	}

	m.UserManager.SendItemChangeNtf(user, op)
	return ack, nil, nil
}

func HandlerEquipChangeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.EquipChangeReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeEquipChange)

	var err error
	ack := &pb.EquipChangeAck{}
	var changePos []int
	if req.Pos == -1 {
		changePos, err = m.Equip.EquipChangeOneKey(user, int(req.HeroIndex), op)
		if err != nil {
			return nil, nil, err
		}

	} else {
		err = m.Equip.EquipChange(user, int(req.HeroIndex), op, int(req.Pos), int(req.BagPos))
		if err != nil {
			return nil, nil, err
		}
		changePos = append(changePos, int(req.Pos))
	}
	ack.HeroIndex = req.HeroIndex
	ack.Equips = buildEquipUnit(user.Heros[int(req.HeroIndex)], changePos)

	m.UserManager.SendItemChangeNtf(user, op)
	return ack, nil, nil
}

func buildEquipUnit(user *objs.Hero, pos []int) map[int32]*pb.EquipUnit {

	equips := make(map[int32]*pb.EquipUnit)
	for _, v := range pos {
		equip := user.Equips[v]
		equips[int32(v)] = builder.BuildPbEquipUnit(equip)
	}

	return equips
}

func HandlerClearReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ClearReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeClear)

	var err error
	ack := &pb.ClearAck{}

	err = m.Equip.Clear(user, int(req.HeroIndex), int(req.Pos), int(req.PropIndex), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerClearReq ack is %v", ack)

	return ack, op, nil
}

func HandlerEquipStrengthenAutoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EquipStrengthenAutoReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeEquipStrengthen)

	var err error
	ack := &pb.EquipStrengthenAutoAck{}

	err = m.Equip.EquipStrengthenOneKey(user, int(req.HeroIndex), req.IsBreak, op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerEquipStrengthenAutoReq ack is %v", ack)

	return ack, op, nil
}


