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
	pb.Register(pb.CmdAuctionInfoReqId, AuctionInfo)                       //物品信息
	pb.Register(pb.CmdBidInfoReqId, HandleBidInfoReq)                      //请求指定物品信息
	pb.Register(pb.CmdAuctionPutawayItemReqId, AuctionPutAwayItem)         //上架物品到拍卖行
	pb.Register(pb.CmdBidReqId, HandleBidReq)                              // 请求竞价
	pb.Register(pb.CmdMyBidReqId, HandleMyBidReq)                          // 请求我的竞拍信息
	pb.Register(pb.CmdMyPutAwayItemInfoReqId, HandleMyPutAwayItemInfo)     // 请求我的上架物品信息 非竞拍中
	pb.Register(pb.CmdMyBidInfoItemReqId, HandleGetMyBidInfo)              // 请求我的竞拍过的信息
	pb.Register(pb.CmdConversionGoldIngotReqId, HandleConversionGoldIngot) // 兑换金锭
}

func AuctionInfo(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.AuctionInfoReq)
	ack := &pb.AuctionInfoNtf{}
	err := m.Auction.ProcessAuctionInfoReq(user, int(req.AuctionType), ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("AuctionInfo ack is %v", ack)
	return ack, nil, nil
}

func AuctionPutAwayItem(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.AuctionPutawayItemReq)
	ack := &pb.AuctionPutawayItemNtf{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAuctionPutAwayItem)
	err := m.Auction.UpItemToAuction(user, req, ack, op)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("AuctionPutAwayItem ack is %v", ack)
	return ack, op, nil
}

func HandleBidInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.BidInfoReq)
	ack := &pb.BidInfoNtf{}
	err := m.Auction.ProcessBidInfo(user, int(req.AuctionId), int(req.AuctionType), ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandleBidInfoReq ack is %v", ack)
	return ack, nil, nil
}

func HandleBidReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.BidReq)
	ack := &pb.BidNtf{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAuctionbidding)
	err := m.Auction.ProcessBid(user, req, ack, op)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandleBidReq ack is %v", ack)
	return ack, op, nil
}

// 处理获取我的竞拍物品请求
func HandleMyBidReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.MyBidNtf{}
	m.Auction.MyBidNtf(user, ack)
	return ack, nil, nil
}

// 处理获取我上架的物品信息
func HandleMyPutAwayItemInfo(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.MyPutAwayItemInfoAck{}
	m.Auction.MyPutAwayItemInfo(user, ack)
	return ack, nil, nil
}

// 处理获取我竞拍过的物品信息
func HandleGetMyBidInfo(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.MyBidInfoItemAck{}
	m.Auction.GetUserBidInfos(user, ack)
	return ack, nil, nil
}

func HandleConversionGoldIngot(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.ConversionGoldIngotReq)
	ack := &pb.ConversionGoldIngotAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeConversionGoldIngot)
	err := m.Auction.ConversionGoldIngot(user, int(req.Num), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandleConversionGoldIngot ack is %v", ack)
	return ack, op, nil
}
