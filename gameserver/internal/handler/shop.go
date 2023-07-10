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
	pb.Register(pb.CmdShopListReqId, HandlerShopListReq)
	pb.Register(pb.CmdShopBuyReqId, HandlerShopBuyReq)
}

func HandlerShopListReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.ShopListReq)
	return &pb.ShopListAck{
		ShopType: req.ShopType,
		ShopList: builder.BuildShopByType(user.Shops.ShopItem[int(req.ShopType)]),
	}, nil, nil
}

func HandlerShopBuyReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ShopBuyReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeShopBuy)

	ack := &pb.ShopBuyAck{}
	err := m.Shop.Buy(user, int(req.Id), int(req.BuyNum), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerShopBuyReq ack:%v", ack)

	return ack, op, nil
}
