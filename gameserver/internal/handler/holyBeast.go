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
	pb.Register(pb.CmdHolyBeastLoadInfoReqId, HolyBeastLoadInfoReq)
	pb.Register(pb.CmdHolyBeastUpStarReqId, HolyBeastUpStarReq)
	pb.Register(pb.CmdHolyBeastChoosePropReqId, HolyBeastChoosePropReq)
	pb.Register(pb.CmdHolyBeastRestReqId, HolyBeastRestReq)
}

func HolyBeastLoadInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.HolyBeastLoadInfoAck{}

	m.HolyBeast.Load(user, ack)

	logger.Debug("HolyBeastLoadInfoReq ack is %v", ack)
	return ack, nil, nil
}



func HolyBeastUpStarReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.HolyBeastUpStarReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.HolyBeastUpStarAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeHolyBeastUpStar)
	if err = m.HolyBeast.UpStar(user, int(req.HeroIndex), int(req.Type), ack, op); err != nil {
		return nil, nil, err
	}
	logger.Debug("HolyBeastUpStarReq ack is %v", ack)
	return ack, op, nil
}

func HolyBeastChoosePropReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.HolyBeastChoosePropReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.HolyBeastChoosePropAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeHolyBeastChooseProp)
	if err = m.HolyBeast.ChooseProp(user, int(req.HeroIndex), int(req.Type), int(req.Index), ack, op); err != nil {
		return nil, nil, err
	}
	logger.Debug("HolyBeastChoosePropReq ack is %v", ack)
	return ack, op, nil
}

func HolyBeastRestReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	//重置功能注销
	//req := p.(*pb.HolyBeastRestReq)
	//user := conn.GetSession().(*managers.ClientSession).User

	//var err error
	ack := &pb.HolyBeastRestAck{}
	//op := ophelper.NewOpBagHelperDefault(constBag.OpTypeHolyBeastRest)
	//if err = m.HolyBeast.Rest(user, int(req.HeroIndex), int(req.Type), ack, op); err != nil {
	//	return nil, nil, err
	//}
	//logger.Debug("HolyBeastRestReq ack is %v", ack)
	return ack, nil, nil
}
