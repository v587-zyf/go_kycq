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
	pb.Register(pb.CmdGuildBonfireLoadReqId, HandlerGuildBonfireLoadReq)
	pb.Register(pb.CmdGuildBonfireAddExpReqId, HandlerGuildBonfireAddExpReq)
	pb.Register(pb.CmdEnterGuildBonfireFightReqId, HandlerEnterGuildBonfireFightReq)
}

func HandlerGuildBonfireLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.GuildBonfireLoadAck{}

	err = m.GuildBonfire.LoadInfo(user, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerGuildLoadInfoReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerGuildBonfireAddExpReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.GuildBonfireAddExpReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeGuildBonfireAddExp)
	var err error
	ack := &pb.GuildBonfireAddExpAck{}

	err = m.GuildBonfire.GuildAddExpPercent(user, op, int(req.ConsumptionType), ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerCreateGuildReq ack is %v", ack)

	return ack, op, nil
}

//进入战斗
func HandlerEnterGuildBonfireFightReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	var err error

	err = m.GuildBonfire.EnterGuildBonfireFight(user)
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}
