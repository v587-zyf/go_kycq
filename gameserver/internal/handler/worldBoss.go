package handler

import (
	"cqserver/gameserver/internal/managers"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdEnterWorldBossFightReqId, HandlerEnterWorldBossFightReq)
}

// 挑战世界boss
func HandlerEnterWorldBossFightReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EnterWorldBossFightReq)
	user := conn.GetSession().(*managers.ClientSession).User

	err := m.WorldBoss.EnterWorldBossFight(user, int(req.StageId))
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}
