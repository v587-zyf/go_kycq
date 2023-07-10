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
	pb.Register(pb.CmdVipGiftGetReqId, HandlerVipGiftGetReq)
	pb.Register(pb.CmdRechargeAllGetReqId, RechargeAllGetReq)
}

func HandlerVipGiftGetReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.VipGiftGetReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeVip)

	var err error
	ack := &pb.VipGiftGetAck{}
	err = m.VipManager.GetGift(user, int(req.Lv), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerVipGiftGetReq ack is %v", ack)

	return ack, op, nil
}

func RechargeAllGetReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.RechargeAllGetReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeVip)

	var err error
	ack := &pb.RechargeAllGetAck{}
	err = m.VipManager.GetRechargeAllGift(user,int(req.Id), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("RechargeAllGetReq ack is %v", ack)

	return ack, op, nil
}
