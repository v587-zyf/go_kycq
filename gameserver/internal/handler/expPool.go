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

	pb.Register(pb.CmdExpPoolLoadReqId, HandlerExpPoolLoadReq)
	pb.Register(pb.CmdExpPoolUpGradeReqId, HandleExpUpGradeReq)

}

func HandlerExpPoolLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	//req := p.(*pb.ExpPoolLoadReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.ExpPoolLoadAck{}

	err = m.ExpPool.Load(user, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerExpPoolLoadReq ack is %v", ack)

	return ack, nil, nil

}

func HandleExpUpGradeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.ExpPoolUpGradeReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeArea)
	var err error
	ack := &pb.ExpPoolUpGradeAck{}
	err = m.ExpPool.Upgrade(user, op, ack, int(req.Index), int(req.Times))
	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}
