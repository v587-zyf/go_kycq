package handler

import (
	"cqserver/gameserver/internal/managers"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdPreferenceSetReqId, HandlerPreferenceSetReq)
	pb.Register(pb.CmdPreferenceLoadReqId, HandlerPreferenceLoadReq)
}

func HandlerPreferenceSetReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.PreferenceSetReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.PreferenceSetAck{}
	err := m.System.PreferenceSet(user, req.Preference, ack)
	if err != nil {
		return nil, nil, err
	}

	logger.Debug("HandlerPreferenceSetReq ack is %v", ack)

	return ack, nil, nil
}

func HandlerPreferenceLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.PreferenceLoadAck{}
	err := m.System.PreferenceLoad(user, ack)
	if err != nil {
		return nil, nil, err
	}

	logger.Debug("HandlerPreferenceLoadReq ack is %v", ack)

	return ack, nil, nil
}