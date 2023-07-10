package handler

import (
	"cqserver/gameserver/internal/managers"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdShaBaKeInfoReqId, HandlerShabakeInfoReq)
	pb.Register(pb.CmdEnterShaBaKeFightReqId, HandlerShabakeEnterReq)
}

//打开页面
func HandlerShabakeInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.ShaBaKeInfoAck{}
	err := m.Shabake.Load(user, ack)
	if err != nil {
		return nil, nil, err
	}

	logger.Debug("HandlerShabakeInfoReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerShabakeEnterReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	err := m.Shabake.EnterShabakeFight(user)
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}
