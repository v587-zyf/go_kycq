package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constRank"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdRankLoadReqId, HandlerRankLoadReq)
	pb.Register(pb.CmdRankWorshipReqId, HandlerWorshipReq)
}

//获取指定类型排行榜
func HandlerRankLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.RankLoadReq)
	user := conn.GetSession().(*managers.ClientSession).User
	key := rmodel.Rank.GetRankKey(int(req.Type), base.Conf.ServerId)
	max := constRank.MAX
	if int(req.Type) == pb.RANKTYPE_ARENA {
		max = constRank.MAX_ARENA
	}
	ranks := m.GetRank().LoadRank(int(req.Type), max-1)
	ack := &pb.RankLoadAck{
		Ranks: make([]*pb.RankInfo, 0),
		Type:  req.Type,
	}
	for i, l := 0, len(ranks); i < l; i += 2 {
		rankUserId := ranks[i]
		rankScore := ranks[i+1]
		ack.Ranks = append(ack.Ranks, m.UserManager.BuildUserRankInfo(rankUserId, -1, (i/2)+1, rankScore))
	}
	//获取玩家自己排名
	selfRank := rmodel.Rank.GetSelfRank(key, user.Id)
	if selfRank >= 0 {
		selfRank += 1
	}
	ack.Self = int32(selfRank)
	return ack, nil, nil
}

func HandlerWorshipReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeRankWorship)
	err := m.Rank.WorshipReward(user, op)
	if err != nil {
		return nil, nil, err
	}
	ack := &pb.RankWorshipAck{
		Goods: op.ToChangeItems(),
	}
	return ack, op, nil
}
