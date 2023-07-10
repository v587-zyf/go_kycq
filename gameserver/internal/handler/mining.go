package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdMiningInReqId, HandlerMiningInReq)
	pb.Register(pb.CmdMiningRobReqId, HandlerMiningRobReq)
	pb.Register(pb.CmdMiningRobBackReqId, HandlerMiningRobBackReq)
	pb.Register(pb.CmdMiningLoadReqId, HandlerMiningLoadReq)
	pb.Register(pb.CmdMiningUpMinerReqId, HandlerMiningUpMinerReq)
	pb.Register(pb.CmdMiningBuyNumReqId, HandlerMiningBuyNumReq)
	pb.Register(pb.CmdMiningStartReqId, HandlerMiningStarReq)
	pb.Register(pb.CmdMiningRobListReqId, HandlerMiningRobListReq)
	pb.Register(pb.CmdMiningListReqId, HandlerMiningListReq)
	pb.Register(pb.CmdMiningDrawLoadReqId, HandlerMiningDrawLoadReq)
	pb.Register(pb.CmdMiningDrawReqId, HandlerMiningDrawReq)
}

func HandlerMiningInReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	err := m.Mining.In(user)
	if err != nil {
		return nil, nil, err
	}
	return &pb.MiningInAck{}, nil, nil
}

func HandlerMiningRobReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.MiningRobReq)
	user := conn.GetSession().(*managers.ClientSession).User
	err := m.Mining.Rob(user, int(req.Id))
	if err != nil {
		return nil, nil, err
	}
	return &pb.MiningRobFightAck{Id: req.Id}, nil, nil
}

func HandlerMiningRobBackReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.MiningRobBackReq)
	user := conn.GetSession().(*managers.ClientSession).User
	err := m.Mining.RobBack(user, int(req.Id))
	if err != nil {
		return nil, nil, err
	}
	return &pb.MiningRobBackFightAck{Id: req.Id}, nil, nil
}

func HandlerMiningLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	return &pb.MiningLoadAck{Mining: builder.BuildMining(user)}, nil, nil
}

func HandlerMiningUpMinerReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.MiningUpMinerReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeMining)

	var err error
	ack := &pb.MiningUpMinerAck{}

	err = m.Mining.UpMiner(user, req.IsMax, op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerMiningUpMinerReq ack is %v", ack)

	return ack, op, nil
}

func HandlerMiningBuyNumReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeMining)

	var err error
	ack := &pb.MiningBuyNumAck{}

	err = m.Mining.BuyNum(user, op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerMiningBuyNumReq ack is %v", ack)

	return ack, op, nil
}

func HandlerMiningStarReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.MiningStartAck{}

	err = m.Mining.StartWork(user, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerMiningStarReq ack is %v", ack)

	return ack, nil, nil
}

func HandlerMiningRobListReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.MiningRobListAck{}
	m.Mining.RobList(user, ack)
	logger.Debug("HandlerMiningRobListReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerMiningListReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	ack := &pb.MiningListAck{}

	m.Mining.List(ack)
	logger.Debug("HandlerMiningListReq ack is %v", ack)

	return ack, nil, nil
}

func HandlerMiningDrawLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.MiningDrawLoadAck{}
	err := m.Mining.GetRobInfo(user, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerMiningDrawLoadReq ack is %v", ack)

	return ack, nil, nil
}

func HandlerMiningDrawReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeMining)

	var err error
	ack := &pb.MiningDrawAck{}

	err = m.Mining.Draw(user, op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerMiningDrawReq ack is %v", ack)

	return ack, op, nil
}
