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
	pb.Register(pb.CmdAtlasActiveReqId, HandlerAtlasActiveReq)
	pb.Register(pb.CmdAtlasUpStarReqId, HandlerAtlasUpStarReq)
	pb.Register(pb.CmdAtlasGatherActiveReqId, HandlerAtlasGatherActiveReq)
	pb.Register(pb.CmdAtlasGatherUpStarReqId, HandlerAtlasGatherUpStarReq)
	pb.Register(pb.CmdAtlasWearChangeReqId, HandlerAtlasWearChangeReq)
}

func HandlerAtlasActiveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AtlasActiveReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAtlasActive)

	var err error
	ack := &pb.AtlasActiveAck{}

	err = m.Atlas.AtlasActive(user, int(req.Id), op, ack)
	if err != nil{
		return nil,nil,err
	}
	logger.Debug("HandlerAtlasActiveReq ack is %v", ack)

	return ack, op, nil
}

func HandlerAtlasUpStarReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AtlasUpStarReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAtlasUpStar)

	var err error
	ack := &pb.AtlasUpStarAck{}

	err = m.Atlas.AtlasUpStar(user,int(req.Id), op, ack)
	if err != nil{
		return nil,nil,err
	}
	logger.Debug("HandlerAtlasUpStarReq ack is %v", ack)

	return ack, op, nil
}

func HandlerAtlasGatherActiveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AtlasGatherActiveReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.AtlasGatherActiveAck{}

	err = m.Atlas.AtlasGatherActive(user, int(req.Id), ack)
	if err != nil{
		return nil,nil,err
	}
	logger.Debug("HandlerAtlasGatherActiveReq ack is %v", ack)

	return ack, nil, nil
}

func HandlerAtlasGatherUpStarReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AtlasGatherUpStarReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.AtlasGatherUpStarAck{}

	err = m.Atlas.AtlasGatherUpStar(user, int(req.Id), ack)
	if err != nil{
		return nil,nil,err
	}
	logger.Debug("HandlerAtlasGatherUpStarReq ack is %v", ack)

	return ack, nil, nil
}

func HandlerAtlasWearChangeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AtlasWearChangeReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.AtlasWearChangeAck{}

	err = m.Atlas.Change(user, int(req.HeroIndex), int(req.Id), ack)
	if err != nil{
		return nil,nil,err
	}
	logger.Debug("HandlerAtlasWearChangeReq ack is %v", ack)

	return ack, nil, nil
}
