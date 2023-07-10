package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdAncientTreasuresActivateReqId, HandlerAncientTreasuresActivateReq)
	pb.Register(pb.CmdAncientTreasuresZhuLinReqId, HandlerAncientTreasuresZhuLinReq)
	pb.Register(pb.CmdAncientTreasuresUpStarReqId, HandlerAncientTreasuresUpStarReq)
	pb.Register(pb.CmdAncientTreasuresJueXingReqId, HandlerAncientTreasuresJueXingReq)
	pb.Register(pb.CmdAncientTreasuresResertReqId, HandlerAncientTreasuresResertReq)
	pb.Register(pb.CmdAncientTreasuresCondotionInfosReqId, HandlerAncientTreasuresCondotionInfosReq)

}

//激活远古宝物
func HandlerAncientTreasuresActivateReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AncientTreasuresActivateReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.AncientTreasuresActivateAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAncientBuyNum)
	err := m.AncientTreasure.Active(user, int(req.TreasureId), ack, op)

	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

//远古宝物  注灵
func HandlerAncientTreasuresZhuLinReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AncientTreasuresZhuLinReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.AncientTreasuresZhuLinAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAncientBuyNum)
	err := m.AncientTreasure.ZhuLin(user, int(req.TreasureId), ack, op)

	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

//远古宝物  升星
func HandlerAncientTreasuresUpStarReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AncientTreasuresUpStarReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.AncientTreasuresUpStarAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAncientBuyNum)
	err := m.AncientTreasure.UpStar(user, int(req.TreasureId), ack, op)

	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

//远古宝物  觉醒
func HandlerAncientTreasuresJueXingReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AncientTreasuresJueXingReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.AncientTreasuresJueXingAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAncientBuyNum)
	err := m.AncientTreasure.JueXin(user, int(req.TreasureId), int(req.Index), req.ChooseItemInfos, ack, op)

	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

//远古宝物  重置
func HandlerAncientTreasuresResertReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AncientTreasuresResertReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.AncientTreasuresResertAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAncientBuyNum)
	err := m.AncientTreasure.Reset(user, int(req.TreasureId), ack, op)

	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

//远古宝物  获取条件值
func HandlerAncientTreasuresCondotionInfosReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.AncientTreasuresCondotionInfosAck{}
	m.AncientTreasure.GetAncientTreasureConditionValue(user, ack)
	return ack, nil, nil
}
