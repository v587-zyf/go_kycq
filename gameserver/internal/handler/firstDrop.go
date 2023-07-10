package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdFirstDropLoadReqId, HandlerFirstDropLoadReq)
	pb.Register(pb.CmdGetFirstDropAwardReqId, HandlerGetFirstDropAwardReq)
	pb.Register(pb.CmdGetAllFirstDropAwardReqId, HandlerGetAllFirstDropAwardReq)
	pb.Register(pb.CmdGetAllRedPacketReqId, HandlerGetAllRedPacketReq)
}

//Load
func HandlerFirstDropLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.FirstDropLoadReq)
	user := conn.GetSession().(*managers.ClientSession).User

	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFashionUpLv)
	ack := &pb.FirstDropLoadAck{}
	m.FirstDrop.LoadInfo(user, int(req.Types), ack)
	return ack, op, nil
}

//首爆装备领取
func HandlerGetFirstDropAwardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.GetFirstDropAwardReq)
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.GetFirstDropAwardAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypFirstDropReward)
	err := m.FirstDrop.GetAward(user, int(req.Id), ack, op)
	if err != nil {
		return nil, nil, err
	}
	goods := op.ToChangeItems()
	ack.Goods = goods
	return ack, op, nil
}

//一键领取首爆奖励
func HandlerGetAllFirstDropAwardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.GetAllFirstDropAwardReq)
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.GetAllFirstDropAwardAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypAllFirstDropReward)
	err := m.FirstDrop.GetAllAward(user, int(req.Types), ack, op)
	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

//一键领取红包
func HandlerGetAllRedPacketReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.GetAllRedPacketReq)
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.GetAllRedPacketAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeItemUse)
	err := m.FirstDrop.GetAllRedPacketAward(user, req.Infos, op)
	if err != nil {
		return nil, nil, err
	}
	ack.UsePacketNum = int32(user.RedPacketUseNum)
	return ack, op, nil
}
