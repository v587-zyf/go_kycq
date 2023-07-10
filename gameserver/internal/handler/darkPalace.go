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
	pb.Register(pb.CmdDarkPalaceLoadReqId, HandlerDarkPalaceLoadReq)
	pb.Register(pb.CmdDarkPalaceBuyNumReqId, HandlerDarkPalaceBuyNumReq)
	pb.Register(pb.CmdEnterDarkPalaceFightReqId, HandlerEnterDarkPalaceFightReq)
}

func HandlerDarkPalaceLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.DarkPalaceLoadReq)
	ack := &pb.DarkPalaceLoadAck{}
	m.DarkPalace.Load(int(req.Floor), ack)
	logger.Debug("HandlerDarkPalaceLoadReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerDarkPalaceBuyNumReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.DarkPalaceBuyNumReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeDarkPalace)

	var err error
	ack := &pb.DarkPalaceBuyNumAck{}

	err = m.DarkPalace.BuyNum(user, req.Use, int(req.BuyNum), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerDarkPalaceBuyNumReq ack is %v", ack)

	return ack, op, nil
}

func HandlerEnterDarkPalaceFightReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EnterDarkPalaceFightReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	err = m.DarkPalace.EnterDarkPalaceFight(user, int(req.StageId), 0)
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}
