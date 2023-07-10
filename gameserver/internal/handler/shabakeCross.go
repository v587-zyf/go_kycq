package handler

import (
	"cqserver/gameserver/internal/managers"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdShaBaKeInfoCrossReqId, HandlerShabakeInfoCrossReq)
	pb.Register(pb.CmdEnterCrossShaBaKeFightReqId, HandlerShabakeEnterCrossReq)
}

//打开页面
func HandlerShabakeInfoCrossReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.ShaBaKeInfoCrossAck{}
	err := m.ShaBaKeCross.LoadCross(user, ack)
	if err != nil {
		return nil, nil, err
	}

	logger.Debug("HandlerShabakeInfoReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerShabakeEnterCrossReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	err := m.ShaBaKeCross.EnterCrossShaBakeFight(user)
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}
