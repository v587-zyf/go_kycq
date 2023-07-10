package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdMiJiUpReqId, HandlerMiJiUpReq)
}

func HandlerMiJiUpReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.MiJiUpReq)
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.MiJiUpAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeMiJiUp)
	err := m.GetMiJi().Up(user, int(req.Id), ack, op)
	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}
