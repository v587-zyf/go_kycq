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
	pb.Register(pb.CmdGuildLoadInfoReqId, HandlerGuildLoadInfoReq)
	pb.Register(pb.CmdCreateGuildReqId, HandlerCreateGuildReq)                   //创建公会
	pb.Register(pb.CmdJoinGuildCombatLimitReqId, HandlerJoinGuildCombatLimitReq) //设置加入工会战力限制
	pb.Register(pb.CmdApplyJoinGuildReqId, HandlerApplyJoinGuildReq)             //加入公会
	pb.Register(pb.CmdQuitGuildReqId, HandlerQuitGuildReq)                       //退出公会
	pb.Register(pb.CmdGuildAssignReqId, HandlerGuildAssignReq)                   //任命
	pb.Register(pb.CmdJoinGuildDisposeReqId, HandlerJoinGuildDisposeReq)         //处理申请列表是否同意玩家加入门派
	pb.Register(pb.CmdGetApplyUserListReqId, HandlerGetApplyUserListReq)         //申请列表玩家
	pb.Register(pb.CmdAllGuildInfosReqId, HandlerGetAllGuildInfosReq)            //门派列表
	pb.Register(pb.CmdKickOutReqId, HandlerKickOutReq)                           //踢人
	pb.Register(pb.CmdDissolveGuildReqId, HandlerDissolveGuildReq)               //解散公会
	pb.Register(pb.CmdModifyBulletinReqId, HandlerModifyBulletinReq)             //修改公告
	pb.Register(pb.CmdImpeachPresidentReqId, HandlerImpeachPresidentReq)         //弹劾会长
	pb.Register(pb.CmdAllJoinGuildDisposeReqId, HandlerAllJoinGuildDisposeReq)   //一键处理申请列表是否同意玩家加入门派

	pb.Register(pb.CmdGuildActivityLoadReqId, HandlerGuildActivityLoadReq) //公会活动加载
}

func HandlerGuildLoadInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.GuildLoadInfoAck{}

	err = m.Guild.LoadGuild(user, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerGuildLoadInfoReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerCreateGuildReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.CreateGuildReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeCreateGuild)

	var err error
	ack := &pb.CreateGuildAck{}

	err = m.Guild.CreateGuild(user, op, req, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerCreateGuildReq ack is %v", ack)

	return ack, op, nil
}

func HandlerJoinGuildCombatLimitReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.JoinGuildCombatLimitReq)
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.JoinGuildCombatLimitAck{}
	err := m.Guild.SetJoinGuildCombatLimit(user, ack, int(req.Combat), int(req.IsAgree))
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerJoinGuildCombatLimitReq ack is %v", ack)

	return ack, nil, nil

}

func HandlerApplyJoinGuildReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ApplyJoinGuildReq)
	user := conn.GetSession().(*managers.ClientSession).User
	var err error
	ack := &pb.ApplyJoinGuildAck{}

	err = m.Guild.ApplyJoinGuild(user, int(req.GuildId), ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerApplyJoinGuildReq ack is %v", ack)

	return ack, nil, nil
}

func HandlerQuitGuildReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	var err error
	ack := &pb.QuitGuildAck{}

	err = m.Guild.QuitGuild(user, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerApplyJoinGuildReq ack is %v", ack)

	return ack, nil, nil

}

func HandlerGetApplyUserListReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	var err error
	ack := &pb.GetApplyUserListAck{}

	err = m.Guild.GetAllApplyUserLists(user, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerGetApplyUserListReq ack is %v", ack)

	return ack, nil, nil

}

func HandlerJoinGuildDisposeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.JoinGuildDisposeReq)
	user := conn.GetSession().(*managers.ClientSession).User
	var err error
	ack := &pb.JoinGuildDisposeAck{}

	err = m.Guild.JoinGuildDispose(user, ack, req.IsAgree, int(req.ApplyUserId))
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerJoinGuildDisposeReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerGuildAssignReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.GuildAssignReq)
	user := conn.GetSession().(*managers.ClientSession).User
	var err error
	ack := &pb.GuildAssignAck{}
	err = m.Guild.GuildAssign(user, ack, int(req.Id), int(req.Position))
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerGuildAssignReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerGetAllGuildInfosReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	var err error
	ack := &pb.AllGuildInfosAck{}
	err = m.Guild.GetAllGuildInfo(user, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerGetAllGuildInfosReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerKickOutReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.KickOutReq)
	user := conn.GetSession().(*managers.ClientSession).User
	var err error
	ack := &pb.KickOutAck{}
	err = m.Guild.KickOut(user, ack, int(req.KickUserId))
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerGetAllGuildInfosReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerDissolveGuildReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	var err error
	ack := &pb.DissolveGuildAck{}
	err = m.Guild.DissolveGuild(user, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerDissolveGuildReq ack is %v", ack)
	return ack, nil, nil
}

//修改公告
func HandlerModifyBulletinReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ModifyBulletinReq)
	user := conn.GetSession().(*managers.ClientSession).User
	var err error
	ack := &pb.ModifyBulletinAck{}
	err = m.Guild.ModifyBulletin(user, ack, req.Content)
	if err != nil {
		return nil, nil, err
	}
	ack.Content = req.Content
	logger.Debug("HandlerDissolveGuildReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerImpeachPresidentReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	var err error
	ack := &pb.ImpeachPresidentAck{}
	err = m.Guild.ImpeachPresident(user, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerImpeachPresidentReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerAllJoinGuildDisposeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AllJoinGuildDisposeReq)
	user := conn.GetSession().(*managers.ClientSession).User
	var err error
	ack := &pb.AllJoinGuildDisposeAck{}

	err = m.Guild.AllJoinGuildDispose(user, ack, req.IsAgree)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerAllJoinGuildDisposeReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerGuildActivityLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.GuildActivityLoadReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.GuildActivityLoadAck{}
	if err := m.Guild.GuildActivityLoad(user, int(req.GuildActivityId), ack); err != nil {
		return nil, nil, err
	}
	return ack, nil, nil
}
