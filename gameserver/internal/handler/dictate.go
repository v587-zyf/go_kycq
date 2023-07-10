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
	pb.Register(pb.CmdDictateUpReqId, HandlerDictateUpReq)
}

func HandlerDictateUpReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.DictateUpReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeDictate)

	var err error
	ack := &pb.DictateUpAck{}

	err = m.Dictate.DictateUpLv(user, int(req.HeroIndex), int(req.Body), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerDictateUpReq ack is %v", ack)

	return ack, op, nil
}
