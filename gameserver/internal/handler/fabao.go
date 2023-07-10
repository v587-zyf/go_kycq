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
	pb.Register(pb.CmdFabaoActiveReqId, HandlerFabaoActiveReq)           // 法宝激活
	pb.Register(pb.CmdFabaoUpLevelReqId, HandlerFabaoUpLevelReq)         // 法宝升级
	pb.Register(pb.CmdFabaoSkillActiveReqId, HandlerFabaoSkillActiveReq) // 法宝技能激活
}

func HandlerFabaoActiveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FabaoActiveReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFabaoActive)

	var err error
	ack := &pb.FabaoActiveAck{}

	err = m.Fabao.ActiveChange(user, int(req.Id), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFabaoActiveReq ack is %v", ack)
	return ack, op, nil
}

func HandlerFabaoUpLevelReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FabaoUpLevelReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.FabaoUpLevelAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFabaoUpLevel)

	err = m.Fabao.UpLevel(user, int(req.Id), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFabaoUpLevelReq ack is %v", ack)
	return ack, op, nil
}

func HandlerFabaoSkillActiveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FabaoSkillActiveReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.FabaoSkillActiveAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFabaoSkillActive)

	err = m.Fabao.ActiveSkillChange(user, int(req.Id), int(req.SkillId), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("handlerFabaoSkillActiveReq ack is %v", ack)
	return ack, op, nil
}
