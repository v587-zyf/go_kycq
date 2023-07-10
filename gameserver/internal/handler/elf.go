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
	pb.Register(pb.CmdElfFeedReqId, HandlerElfFeedReq)
	pb.Register(pb.CmdElfSkillUpLvReqId, HandlerElfSkillUpLvReq)
	pb.Register(pb.CmdElfSkillChangePosReqId, HandlerElfSkillChangePosReq)
	//pb.Register(pb.CmdElfActiveReqId, HandlerElfActiveReq)
}

func HandlerElfFeedReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ElfFeedReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeElf)

	var err error
	ack := &pb.ElfFeedAck{}

	err = m.Elf.Feed(user, op, req.Positions, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerElfFeedReq ack is %v", ack)

	return ack, op, nil
}

func HandlerElfSkillUpLvReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ElfSkillUpLvReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.ElfSkillUpLvAck{}

	err = m.Elf.SkillUpLv(user, int(req.SkillId), ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerElfSkillUpLvReq ack is %v", ack)

	return ack, nil, nil
}

func HandlerElfSkillChangePosReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ElfSkillChangePosReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.ElfSkillChangePosAck{}

	err = m.Elf.SkillChangePos(user, int(req.SkillId), int(req.Pos), ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerElfSkillChangePosReq ack is %v", ack)

	return ack, nil, nil
}

//func HandlerElfActiveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//	user := conn.GetSession().(*managers.ClientSession).User
//
//	var err error
//	ack := &pb.ElfActiveAck{}
//
//	err = m.Elf.Active(user, ack)
//	if err != nil {
//		return nil, nil, err
//	}
//	logger.Debug("HandlerElfActiveReq ack is %v", ack)
//
//	return ack, nil, nil
//}
