package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdEnterAppletsReqId, HandlerEnterAppletsReq)     //进入小程序
	pb.Register(pb.CmdAppletsReceiveReqId, HandlerAppletsReceiveReq) //领取魔法射击杀怪奖励
	pb.Register(pb.CmdCronGetAwardReqId, HandlerCronGetAwardReq)     //魔法射击定时奖励获取
	pb.Register(pb.CmdEndResultReqId, HandlerEndResultReq)           //通关奖励
}

func HandlerEnterAppletsReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EnterAppletsReq)
	user := conn.GetSession().(*managers.ClientSession).User
	ntf := &pb.AppletsEnergyNtf{}
	err := m.Applets.EnterAppletsReq(user, int(req.AppletsType), ntf)
	if err != nil {
		return nil, nil, err
	}
	return ntf, nil, nil
}

func HandlerAppletsReceiveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AppletsReceiveReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAppletsReceive)

	ack := &pb.AppletsReceiveAck{}
	err := m.Applets.AppletsReceiveReq(user, int(req.ReceiveId), ack, op)
	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

func HandlerCronGetAwardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.CronGetAwardReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeCronGetAwardReq)

	ack := &pb.CronGetAwardAck{}
	err := m.Applets.CronGetAwardReq(user, int(req.Id), int(req.Index), ack, op)
	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

func HandlerEndResultReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EndResultReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeEndResult)

	ack := &pb.EndResultAck{}
	err := m.Applets.EndResultReq(user, int(req.AppletsType), int(req.Id), ack, op)
	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}
