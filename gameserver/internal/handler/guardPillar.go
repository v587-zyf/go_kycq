package handler

import (
	"cqserver/gameserver/internal/managers"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdEnterGuardPillarReqId, HandlerEnterGuardPillarReq)
}

func HandlerEnterGuardPillarReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EnterGuardPillarReq)
	user := conn.GetSession().(*managers.ClientSession).User
	if err := m.GuardPillar.In(user, int(req.StageId)); err != nil {
		return nil, nil, err
	}
	return nil, nil, nil
}
