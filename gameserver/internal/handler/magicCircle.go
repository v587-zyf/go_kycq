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
	pb.Register(pb.CmdMagicCircleUpLvReqId, HandlerMagicCircleUpLvReq)
	pb.Register(pb.CmdMagicCircleChangeWearReqId, HandlerMagicCircleChangeWearReq)
}

func HandlerMagicCircleUpLvReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.MagicCircleUpLvReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeMagicCircle)

	var err error
	ack := &pb.MagicCircleUpLvAck{}

	err = m.MagicCircle.UpLv(user, int(req.HeroIndex), int(req.MagicCircleType), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerMagicCircleUpLvReq ack is %v", ack)

	return ack, op, nil
}

func HandlerMagicCircleChangeWearReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.MagicCircleChangeWearReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.MagicCircleChangeWearAck{}

	err = m.MagicCircle.ChangeWear(user, int(req.HeroIndex), int(req.MagicCircleLvId), ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerMagicCircleChangeWearReq ack is %v", ack)

	return ack, nil, nil
}
