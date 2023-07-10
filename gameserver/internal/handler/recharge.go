package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdRechargeApplyPayReqId, HandleRechargeApplyPayReq)
	pb.Register(pb.CmdMoneyPayReqId, HandlerMoneyPayReq)

	pb.Register(pb.CmdContRechargeReceiveReqId, HandlerContRechargeReceiveReq)
}

// 充值相关-发起订单
func HandleRechargeApplyPayReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.RechargeApplyPayReq)
	user := conn.GetSession().((*managers.ClientSession)).User
	payData, err, isPayToken := m.Recharge.ApplyPay(user, req.RechargeId, int(req.GetPayNum()))
	if err != nil {
		return nil, nil, err
	}
	return &pb.RechargeApplyPayAck{Result: true, PayData: payData, RechargeId: req.RechargeId, IsPayToken: isPayToken}, nil, nil
}

// 人民币购买
func HandlerMoneyPayReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.MoneyPayReq)
	user := conn.GetSession().((*managers.ClientSession)).User
	payData, err, isPayToken, order := m.Recharge.Pay(user, int(req.GetPayNum()), int(req.GetPayType()), int(req.GetTypeId()), false)
	if err != nil {
		return nil, nil, err
	}
	return &pb.MoneyPayAck{Result: true, PayData: payData, PayType: req.PayType, PayNum: int32(order.PayMoney), TypeId: req.TypeId, IsPayToken: isPayToken}, nil, nil
}

func HandlerContRechargeReceiveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ContRechargeReceiveReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeContReceive)

	err := m.Recharge.ContRechargeReceive(user, int(req.ContRechargeId), op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.ContRechargeReceiveAck{ContRechargeId: req.ContRechargeId, Goods: op.ToChangeItems()}, op, nil
}
