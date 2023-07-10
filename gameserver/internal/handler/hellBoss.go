package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdHellBossLoadReqId, HandlerHellBossLoadReq)
	pb.Register(pb.CmdHellBossBuyNumReqId, HandlerHellBossBuyNumReq)
	pb.Register(pb.CmdEnterHellBossFightReqId, HandlerEnterHellBossFightReq)
}

func HandlerHellBossLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.HellBossLoadReq)
	return &pb.HellBossLoadAck{Floor: req.Floor, List: m.HellBoss.Load(user, int(req.Floor))}, nil, nil
}

func HandlerHellBossBuyNumReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.HellBossBuyNumReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeHellBossBuyNum)
	if err := m.HellBoss.BuyNum(user, req.Use, int(req.BuyNum), op); err != nil {
		return nil, nil, err
	}
	return &pb.HellBossBuyNumAck{BuyNum: int32(user.HellBoss.BuyNum)}, op, nil
}

func HandlerEnterHellBossFightReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EnterHellBossFightReq)
	user := conn.GetSession().(*managers.ClientSession).User
	if err := m.HellBoss.EnterHellBossFight(user, int(req.StageId), 0); err != nil {
		return nil, nil, err
	}
	return nil, nil, nil
}
