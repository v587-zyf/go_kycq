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
	pb.Register(pb.CmdSevenInvestmentLoadReqId, SevenInvestmentLoadReq)
	pb.Register(pb.CmdGetSevenInvestmentAwardReqId, GetSevenInvestmentAwardReq)
}

func SevenInvestmentLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.SevenInvestmentLoadAck{}

	m.SevenInvestment.Load(user, ack)
	logger.Debug("SevenInvestmentLoadReq ack is %v", ack)
	return ack, nil, nil
}

func GetSevenInvestmentAwardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.GetSevenInvestmentAwardReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeSevenInvestment)

	var err error
	ack := &pb.GetSevenInvestmentAwardAck{}

	err = m.SevenInvestment.GetAward(user, op, int(req.Id), ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("GetSevenInvestmentAwardReq ack is %v", ack)
	return ack, op, nil
}
