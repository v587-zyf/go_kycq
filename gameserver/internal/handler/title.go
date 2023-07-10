package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdTitleActiveReqId, HandlerTitleActiveReq)
	pb.Register(pb.CmdTitleWearReqId, HandlerTitleWearReq)
	pb.Register(pb.CmdTitleRemoveReqId, HandlerTitleRemoveReq)
	pb.Register(pb.CmdTitleLookReqId, HandlerTitleLookReq)
}

func HandlerTitleActiveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.TitleActiveReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeTitleActive)

	err := m.Title.Active(user, int(req.TitleId), op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.TitleActiveAck{Title: builder.BuildTitle(int(req.TitleId), user.Title[int(req.TitleId)])}, op, nil
}

func HandlerTitleWearReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.TitleWearReq)
	user := conn.GetSession().(*managers.ClientSession).User

	err := m.Title.Wear(user, int(req.HeroIndex), int(req.TitleId))
	if err != nil {
		return nil, nil, err
	}

	return &pb.TitleWearAck{HeroIndex: req.HeroIndex, TitleId: req.TitleId}, nil, nil
}

func HandlerTitleRemoveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.TitleRemoveReq)
	user := conn.GetSession().(*managers.ClientSession).User

	err, titleId := m.Title.Remove(user, int(req.HeroIndex))
	if err != nil {
		return nil, nil, err
	}

	return &pb.TitleRemoveAck{HeroIndex: req.HeroIndex, TitleId: int32(titleId)}, nil, nil
}

func HandlerTitleLookReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.TitleLookReq)
	user := conn.GetSession().(*managers.ClientSession).User

	err := m.Title.Look(user, int(req.TitleId))
	if err != nil {
		return nil, nil, err
	}

	return &pb.TitleLookAck{TitleId: req.TitleId}, nil, nil
}
