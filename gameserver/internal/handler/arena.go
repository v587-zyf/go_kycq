package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdArenaOpenReqId, HandlerArenaOpenReq)
	pb.Register(pb.CmdEnterArenaFightReqId, HandlerEnterArenaFightReq)
	pb.Register(pb.CmdBuyArenaFightNumReqId, HandlerBuyArenaFightNumReq)
	pb.Register(pb.CmdRefArenaRankReqId, HandlerRefArenaRankReq)
}

func HandlerArenaOpenReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.ArenaOpenAck{}
	err := m.Arena.ArenaOpen(user, ack)
	if err != nil {
		return nil, nil, err
	}

	logger.Debug("HandlerArenaOpenReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerEnterArenaFightReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EnterArenaFightReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error

	err = m.Arena.EnterArenaFight(user, int(req.ChallengeUid), int(req.ChallengeRanking))
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}

func HandlerBuyArenaFightNumReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeArena)

	var err error
	ack := &pb.BuyArenaFightNumAck{}

	err = m.Arena.BuyArenaFightNum(user, op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerBuyArenaFightNumReq ack is %v", ack)

	return ack, op, nil
}

func HandlerRefArenaRankReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.RefArenaRankAck{}

	err = m.Arena.RefArenaRank(user, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerRefArenaRankReq ack is %v", ack)

	return ack, nil, nil
}
