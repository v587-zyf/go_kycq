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
	pb.Register(pb.CmdDailyPackBuyReqId, HandlerDailyPackBuyReq)
}

func HandlerDailyPackBuyReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.DailyPackBuyReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeDailyPack)

	var err error
	ack := &pb.DailyPackBuyAck{}

	err = m.DailyPack.Buy(user, int(req.Id), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerDailyPackBuyReq ack is %v", ack)

	return ack, op, nil
}