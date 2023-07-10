package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdLabelUpReqId, HandlerLabelUpReq)
	pb.Register(pb.CmdLabelTransferReqId, HandlerLabelTransFerReq)
	pb.Register(pb.CmdLabelDayRewardReqId, HandlerLabelDayRewardReq)
	pb.Register(pb.CmdLabelTaskReqId, HandlerLabelTaskReq)
}

func HandlerLabelUpReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeLabelUp)

	if err := m.Label.Up(user, op); err != nil {
		return nil, nil, err
	}
	return &pb.LabelUpAck{Id: int32(user.Label.Id), Goods: op.ToChangeItems()}, op, nil
}

func HandlerLabelTransFerReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.LabelTransferReq)
	user := conn.GetSession().(*managers.ClientSession).User

	if err := m.Label.Transfer(user, int(req.Job)); err != nil {
		return nil, nil, err
	}
	return &pb.LabelTransferAck{Job: req.Job, Transfer: int32(user.Label.Transfer)}, nil, nil
}

func HandlerLabelDayRewardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeLabelDayReward)

	if err := m.Label.DayReward(user, op); err != nil {
		return nil, nil, err
	}
	return &pb.LabelDayRewardAck{DayReward: true, Goods: op.ToChangeItems()}, op, nil
}

func HandlerLabelTaskReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	m.Label.SendLabelTaskNtf(user, -1)
	return nil, nil, nil
}
