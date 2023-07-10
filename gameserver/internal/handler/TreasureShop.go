package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdTreasureShopLoadReqId, HandlerTreasureShopLoadReq)
	pb.Register(pb.CmdTreasureShopCarChangeReqId, HandlerTreasureCarChangeReq)
	pb.Register(pb.CmdTreasureShopBuyReqId, HandlerTreasureShopBuyReq)
	pb.Register(pb.CmdTreasureShopRefreshReqId, HandlerTreasureShopRefreshReq)
}

func HandlerTreasureShopLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	return m.TreasureShop.Load(user), nil, nil
}

func HandlerTreasureCarChangeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.TreasureShopCarChangeReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.TreasureShopCarChangeAck{}
	if err := m.TreasureShop.CarChange(user, int(req.ShopId), req.IsAdd, ack); err != nil {
		return nil, nil, err
	}
	return ack, nil, nil
}

func HandlerTreasureShopBuyReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.TreasureShopBuyReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeTreasureShopBuy)

	ack := &pb.TreasureShopBuyAck{}
	if err := m.TreasureShop.Buy(user, req.Shop, op, ack); err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

func HandlerTreasureShopRefreshReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeTreasureShopRefresh)

	if err := m.TreasureShop.RefreshShop(user, op); err != nil {
		return nil, nil, err
	}
	return nil, op, nil
}
