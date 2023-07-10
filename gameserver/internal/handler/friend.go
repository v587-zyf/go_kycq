package handler

import (
	"cqserver/gameserver/internal/managers"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdFriendListReqId, HandlerFriendListReq)
	pb.Register(pb.CmdFriendAddReqId, HandlerFriendAddReq)
	pb.Register(pb.CmdFriendDelReqId, HandlerFriendDelReq)
	pb.Register(pb.CmdFriendBlockAddReqId, HandlerFriendBlockAddReq)
	pb.Register(pb.CmdFriendBlockDelReqId, HandlerFriendBlockDelReq)
	pb.Register(pb.CmdFriendBlockListReqId, HandlerFriendBlockListReq)
	pb.Register(pb.CmdFriendSearchReqId, HandlerFriendSearchReq)
	pb.Register(pb.CmdFriendMsgReadReqId, HandlerFriendMsgReadReq)
	pb.Register(pb.CmdFriendUserInfoReqId, HandlerFriendUserInfoReq)
	pb.Register(pb.CmdFriendMsgReqId, HandlerFriendMsgReq)

	pb.Register(pb.CmdFriendApplyListReqId, HandlerFriendApplyListReq)
	pb.Register(pb.CmdFriendApplyAddReqId, HandlerFriendApplyAddReq)
	pb.Register(pb.CmdFriendApplyAgreeReqId, HandlerFriendApplyAgreeReq)
	pb.Register(pb.CmdFriendApplyRefuseReqId, HandlerFriendApplyRefuseReq)
}

func HandlerFriendListReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	friendInfo := m.Friend.List(user, false)

	return &pb.FriendListAck{FriendList: friendInfo}, nil, nil
}

func HandlerFriendAddReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FriendAddReq)
	user := conn.GetSession().(*managers.ClientSession).User
	var err error
	ack := &pb.FriendAddAck{}

	err = m.Friend.Add(user, req.UserId, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFriendAddReq ack:%v", ack)

	return ack, nil, nil
}

func HandlerFriendDelReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FriendDelReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	err = m.Friend.Del(user, int(req.UserId))
	if err != nil {
		return nil, nil, err
	}

	return &pb.FriendDelAck{UserId: req.UserId}, nil, nil
}

func HandlerFriendBlockAddReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FriendBlockAddReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	err = m.Friend.BlockAdd(user, int(req.UserId))
	if err != nil {
		return nil, nil, err
	}

	return &pb.FriendBlockAddAck{UserId: req.UserId, FriendList: m.Friend.List(user, true)}, nil, nil
}

func HandlerFriendBlockDelReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FriendBlockDelReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	err = m.Friend.BlockDel(user, int(req.UserId))
	if err != nil {
		return nil, nil, err
	}

	return &pb.FriendBlockDelAck{UserId: req.UserId}, nil, nil
}

func HandlerFriendBlockListReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	friendInfo := m.Friend.List(user, true)

	return &pb.FriendBlockListAck{FriendList: friendInfo}, nil, nil
}

func HandlerFriendSearchReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FriendSearchReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.FriendSearchAck{}
	err = m.Friend.Search(user, req.Name, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFriendSearchReq ack:%v", ack)

	return ack, nil, nil
}

func HandlerFriendMsgReadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FriendMsgReadReq)
	user := conn.GetSession().(*managers.ClientSession).User

	err := m.Friend.ReadMsg(user, int(req.FriendId))
	if err != nil {
		return nil, nil, err
	}

	return &pb.FriendMsgReadAck{FriendId: req.FriendId}, nil, nil
}

func HandlerFriendUserInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FriendUserInfoReq)
	return &pb.FriendUserInfoAck{FriendUserInfo: m.Friend.GetFriendUserInfo(int(req.FriendId))}, nil, nil
}

func HandlerFriendMsgReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FriendMsgReq)

	user := conn.GetSession().(*managers.ClientSession).User
	var err error
	ack := &pb.FriendMsgAck{}

	err = m.Friend.FriendMsg(user, int(req.FriendId), ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFriendMsgReq ack:%v", ack)

	return ack, nil, nil
}

func HandlerFriendApplyListReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	return &pb.FriendApplyListAck{ApplyList: m.Friend.ApplyList(user)}, nil, nil
}

func HandlerFriendApplyAddReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FriendApplyAddReq)
	user := conn.GetSession().(*managers.ClientSession).User

	err := m.Friend.ApplyAdd(user, int(req.FriendId))
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}

func HandlerFriendApplyAgreeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FriendApplyAgreeReq)
	user := conn.GetSession().(*managers.ClientSession).User

	err := m.Friend.ApplyAgree(user, int(req.FriendId))
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}

func HandlerFriendApplyRefuseReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FriendApplyRefuseReq)
	user := conn.GetSession().(*managers.ClientSession).User

	err := m.Friend.ApplyRefuse(user, int(req.FriendId))
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}
