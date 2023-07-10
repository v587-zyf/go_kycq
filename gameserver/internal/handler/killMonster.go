package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdKillMonsterUniLoadReqId, HandlerKillMonsterUniLoadReq)
	pb.Register(pb.CmdKillMonsterUniFirstDrawReqId, HandlerKillMonsterUniFirstDrawReq)
	pb.Register(pb.CmdKillMonsterUniDrawReqId, HandlerKillMonsterUniDrawReq)
	pb.Register(pb.CmdKillMonsterPerLoadReqId, HandlerKillMonsterPerLoadReq)
	pb.Register(pb.CmdKillMonsterPerDrawReqId, HandlerKillMonsterPerDrawReq)
	pb.Register(pb.CmdKillMonsterMilLoadReqId, HandlerKillMonsterMilLoadReq)
	pb.Register(pb.CmdKillMonsterMilDrawReqId, HandlerKillMonsterMilDrawReq)
}

func HandlerKillMonsterUniLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	err, infos := m.KillMonster.LoadUni(user)
	if err != nil {
		return nil, nil, err
	}
	return &pb.KillMonsterUniLoadAck{List: infos}, nil, nil
}

func HandlerKillMonsterUniFirstDrawReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.KillMonsterUniFirstDrawReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeKillMonsterUniFirstDraw)

	if err := m.KillMonster.DrawUniFirst(user, int(req.StageId), op); err != nil {
		return nil, nil, err
	}
	return &pb.KillMonsterUniFirstDrawAck{StageId: req.StageId, Goods: op.ToChangeItems()}, op, nil
}

func HandlerKillMonsterUniDrawReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.KillMonsterUniDrawReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeKillMonsterUniDraw)

	if err := m.KillMonster.DrawUni(user, int(req.StageId), op); err != nil {
		return nil, nil, err
	}
	return &pb.KillMonsterUniDrawAck{StageId: req.StageId, Goods: op.ToChangeItems()}, op, nil
}

func HandlerKillMonsterPerLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	err, infos := m.KillMonster.LoadPer(user)
	if err != nil {
		return nil, nil, err
	}
	return &pb.KillMonsterPerLoadAck{List: infos}, nil, nil
}

func HandlerKillMonsterPerDrawReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.KillMonsterPerDrawReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeKillMonsterPerDraw)

	if err := m.KillMonster.DrawPer(user, int(req.StageId), op); err != nil {
		return nil, nil, err
	}
	return &pb.KillMonsterPerDrawAck{StageId: req.StageId, Goods: op.ToChangeItems()}, op, nil
}

func HandlerKillMonsterMilLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	err, infos := m.KillMonster.LoadMil(user)
	if err != nil {
		return nil, nil, err
	}
	return &pb.KillMonsterMilLoadAck{List: infos}, nil, nil
}

func HandlerKillMonsterMilDrawReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.KillMonsterMilDrawReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeKillMonsterMilDraw)

	ack := &pb.KillMonsterMilDrawAck{}
	if err := m.KillMonster.DrawMil(user, int(req.Type), op, ack); err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}
