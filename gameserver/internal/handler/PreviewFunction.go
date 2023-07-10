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
	pb.Register(pb.CmdPreviewFunctionLoadReqId, PreviewFunctionLoad)
	pb.Register(pb.CmdPreviewFunctionGetReqId, PreviewFunctionGet)
	pb.Register(pb.CmdPreviewFunctionPointReqId, PreviewFunctionPoint)
}

func PreviewFunctionLoad(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.PreviewFunctionLoadAck{}

	m.PreviewFunction.Load(user, ack)

	logger.Debug("PreviewFunctionLoad ack is %v", ack)
	return ack, nil, nil
}

func PreviewFunctionGet(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.PreviewFunctionGetReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypePreviewFunction)

	var err error
	ack := &pb.PreviewFunctionGetAck{}

	err = m.PreviewFunction.GetReward(user, int(req.Id), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("PreviewFunctionGet ack is %v", ack)

	return ack, op, nil
}

func PreviewFunctionPoint(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.PreviewFunctionPointReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.PreviewFunctionPointAck{}
	err = m.PreviewFunction.SetPointId(user, int(req.Id), ack)
	if err != nil{
		return nil, nil, err
	}
	return ack, nil, nil
}
