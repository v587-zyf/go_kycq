package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdDailyRankLoadReqId, HandlerDailyRankLoadReq)
	pb.Register(pb.CmdDailyRankGetMarkRewardReqId, HandlerDailyRankGetMarkRewardReq)
	pb.Register(pb.CmdDailyRankBuyGiftReqId, HandlerDailyRankBuyGiftReq)
}

//获取指定类型排行榜
func HandlerDailyRankLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.DailyRankLoadAck{}
	m.DailyRank.LoadRankReq(user, ack)
	return ack, nil, nil
}

//获取积分奖励
func HandlerDailyRankGetMarkRewardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.DailyRankGetMarkRewardReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeDailyRankGetMarkReward)

	ack := &pb.DailyRankGetMarkRewardAck{}
	err := m.DailyRank.GetMarkReward(user, int(req.Id), op, ack)
	if err != nil {
		return nil, nil, err
	}
	return ack, nil, nil
}

//礼包购买 消耗元宝
func HandlerDailyRankBuyGiftReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.DailyRankBuyGiftReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeDailyRankBuyGift)

	ack := &pb.DailyRankBuyGiftAck{}
	err := m.DailyRank.BuyDailyRankGift(user, int(req.Id), op, ack)
	if err != nil {
		return nil, nil, err
	}
	return ack, nil, nil
}
