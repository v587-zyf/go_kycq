package handler

import (
	"cqserver/gameserver/internal/managers"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdEnterDailyActivityReqId, HandlerEnterDailyActivityReq)
	pb.Register(pb.CmdDailyActivityListReqId, HandlerDailyActivityListReq)
}

func HandlerEnterDailyActivityReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EnterDailyActivityReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	err = m.DailyActivity.EnterDailyActivity(user, int(req.ActivityId),int(req.StageId))
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}

func HandlerDailyActivityListReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	ack := &pb.DailyActivityListAck{}
	m.DailyActivity.List(ack)
	logger.Debug("HandlerDailyActivityListReq ack:%v", ack)

	return ack, nil, nil
}