package handler

import (
	"cqserver/gameserver/internal/managers"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdLoadWorldLeaderReqId, HandleLoadWorldLeaderReq)
	pb.Register(pb.CmdWorldLeaderEnterReqId, HandleWorldLeaderEnterReq)
	pb.Register(pb.CmdGetWorldLeaderRankInfoReqId, HandleGetWorldLeaderRankInfoReq)
}

func HandleLoadWorldLeaderReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.LoadWorldLeaderAck{}
	err := m.WorldLeader.LoadWorldLeader(user, ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, nil, nil
}

//进入
func HandleWorldLeaderEnterReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.WorldLeaderEnterReq)
	user := conn.GetSession().((*managers.ClientSession)).User
	ack := &pb.WorldLeaderEnterAck{}
	err := m.WorldLeader.WorldLeaderEnter(user, int(req.StageId), ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, nil, nil
}

///获取伤害排名
func HandleGetWorldLeaderRankInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.GetWorldLeaderRankInfoReq)
	user := conn.GetSession().((*managers.ClientSession)).User
	ack := &pb.GetWorldLeaderRankInfoAck{}
	err := m.WorldLeader.WorldLeaderRankInfo(user, int(req.StageId), ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, nil, nil
}
