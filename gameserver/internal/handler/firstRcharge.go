package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdFirstRechargeRewardReqId, HandlerFirstRechargeRewardReq)
}

func HandlerFirstRechargeRewardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FirstRechargeRewardReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFirstRecharge)

	var err error
	err = m.FirstRecharge.Reward(user, int(req.Day), op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.FirstRechargeRewardAck{
		Day:   req.Day,
		Goods: op.ToChangeItems(),
	}, op, nil
}
