package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdSpendRebatesRewardReqId, HandlerSpendRebatesRewardReq)
}

func HandlerSpendRebatesRewardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.SpendRebatesRewardReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeSpendRebates)

	var err error
	err = m.SpendRebates.Reward(user, int(req.Id), op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.SpendRebatesRewardAck{
		Id:    req.Id,
		Goods: op.ToChangeItems(),
	}, op, nil
}
