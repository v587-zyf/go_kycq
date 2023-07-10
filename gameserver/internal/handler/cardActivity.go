package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdCardActivityApplyGetReqId, HandleApplyGetCardReq)
	pb.Register(pb.CmdCardActivityInfosReqId, HandleGetCardAtInfoReq)
	pb.Register(pb.CmdGetIntegralAwardReqId, HandleGetCardFromCardBoxReq)
}

/*申请抽卡*/
func HandleApplyGetCardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.CardActivityApplyGetReq)
	user := conn.GetSession().(*managers.ClientSession).User
	opHelper := ophelper.NewOpBagHelperDefault(constBag.OpTypeCardActivity)

	ack := &pb.CardActivityApplyGetAck{}
	err := m.CardActivity.Draw(user, int(req.Times), ack, opHelper)
	if err != nil {
		return nil, nil, err
	}
	return ack, nil, nil
}

/**获取系统抽卡信息*/
func HandleGetCardAtInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.CardActivityInfosAck{}
	err := m.CardActivity.Load(user, ack)
	if err != nil {
		return nil, nil, err
	}
	return ack, nil, nil
}

/*抽卡券宝箱领取功能*/
func HandleGetCardFromCardBoxReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	req := p.(*pb.GetIntegralAwardReq)
	opHelper := ophelper.NewOpBagHelperDefault(constBag.OpTypeActivityCardReward)

	ack := &pb.GetIntegralAwardAck{}
	err := m.CardActivity.GetReward(user, int(req.Index),int(req.Times), ack, opHelper)
	if err != nil {
		return nil, nil, err
	}
	return ack, opHelper, nil
}
