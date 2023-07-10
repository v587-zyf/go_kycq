package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdOpenGiftReqId, HandlerOpenGiftReq)
	pb.Register(pb.CmdGiftCodeRewardReqId, HandlerGiftCodeRewardReq)
	pb.Register(pb.CmdLimitedGiftBuyReqId, HandlerLimitedGiftBuyReq)
	pb.Register(pb.CmdLimitedGiftReqId, HandlerLimitGiftReq)
	pb.Register(pb.CmdOpenGiftEndTimeReqId, HandlerOpenGiftEndTimeReq)
}

func HandlerOpenGiftReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.OpenGiftReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeOpenGiftReward)
	ack := &pb.OpenGiftAck{}
	items := make([]int, 0)
	for _, itemId := range req.ChooseItemId {
		items = append(items, int(itemId))
	}
	err := m.Gift.OpenGift(user, int(req.Type), int(req.ItemId), int(req.Num), items, op)
	if err != nil {
		return nil, nil, err
	}
	ack.Goods = op.ToChangeItems()
	return ack, op, nil
}

func HandlerGiftCodeRewardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.GiftCodeRewardReq)
	user := conn.GetSession().(*managers.ClientSession).User

	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeGiftCodeReward)
	err := m.Gift.GiftCodeReward(user, req.Code, op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.GiftCodeRewardAck{Code: req.Code, Goods: op.ToChangeItems()}, op, nil
}

func HandlerLimitedGiftBuyReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.LimitedGiftBuyReq)
	user := conn.GetSession().(*managers.ClientSession).User

	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeLimitedGift)
	err := m.Gift.LimitedGiftBuy(user, int(req.Type), op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.LimitedGiftBuyAck{Goods: op.ToChangeItems(), Type: req.Type}, op, nil
}

func HandlerLimitGiftReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ntf := m.Gift.SendLimitedGift(user)
	if ntf == nil {
		ntf = &pb.LimitedGiftNtf{}
	}
	return ntf, nil, nil
}

func HandlerOpenGiftEndTimeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	return &pb.OpenGiftEndTimeAck{EndTime: m.Gift.OpenGiftEndTime(user).Unix()}, nil, nil
}
