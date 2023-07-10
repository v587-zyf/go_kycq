package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdMaterialStageLoadReqId, HandlerMaterialStageLoadReq)
	pb.Register(pb.CmdEnterMaterialStageFightReqId, HandlerEnterMaterialStageFightReq)
	pb.Register(pb.CmdMaterialStageSweepReqId, HandlerMaterialStageSweepReq)
	pb.Register(pb.CmdMaterialStageBuyNumReqId, HandlerMaterialStageBuyNumReq)
}

func HandlerMaterialStageLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	return &pb.MaterialStageLoadAck{MaterialStage: builder.BuildMaterialStage(user)}, nil, nil
}

func HandlerEnterMaterialStageFightReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EnterMaterialStageFightReq)
	user := conn.GetSession().(*managers.ClientSession).User

	err := m.MaterialStage.EnterMaterialStageFight(user, int(req.StageId))
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}

func HandlerMaterialStageSweepReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.MaterialStageSweepReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeMaterialStageSweep)

	ack := &pb.MaterialStageSweepAck{}
	err := m.MaterialStage.MaterialStageSweep(user, int(req.StageId), op, ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, op, nil
}

func HandlerMaterialStageBuyNumReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.MaterialStageBuyNumReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeMaterialStageBuyNum)

	err := m.MaterialStage.MaterialBuyNum(user, int(req.MaterialType), req.Use, op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.MaterialStageBuyNumAck{BuyNum: int32(user.MaterialStage.MaterialStages[int(req.MaterialType)].BuyNum), MaterialType: req.MaterialType}, op, nil
}
