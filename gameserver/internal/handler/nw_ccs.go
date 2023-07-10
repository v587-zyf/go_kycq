package handler

import (
	"cqserver/gamelibs/errex"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pbserver"
	"strconv"
)

func init() {
	pbserver.Register(pbserver.CmdCCSToGsCrossFsIdChangeNtfId, handlerCrossFsIdChange)
	//擂台赛相关 ccs->gs
	pbserver.Register(pbserver.CmdChallengeSendLoseRewardNtfId, sendLoseReward)
	pbserver.Register(pbserver.CmdChallengeAppuserUpToGsNtfId, sendApplyUserUpdate)
	//跨服沙巴克相关
	pbserver.Register(pbserver.CmdCcsToGsBroadShaBakeFirstGuildInfoId, setFirstGuildInfo)
	pbserver.Register(pbserver.CmdRechageCcsToGsReqId, handlerRechargeReq)
	pbserver.Register(pbserver.CmdRechargeApplyReqId, handlerRechargeApplyReq)
	pbserver.Register(pbserver.CmdBanInfoCcsToGsReqId, handlerBanAccountInfoReq)
	pbserver.Register(pbserver.CmdMailSendCCsToGsReqId, handlerMailSendCCsToGsReq)
	pbserver.Register(pbserver.CmdFuncStateUpdateReqId, handlerFuncStateUpdate)
	pbserver.Register(pbserver.CmdUpAnnouncementNowReqId, handlerUpAnnouncementNowReq)
	pbserver.Register(pbserver.CmdUpPaoMaDengNowReqId, handlerUpPaoMaDengNowReq)
}

func handlerCrossFsIdChange(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {

	return nil, nil
}

func sendLoseReward(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {

	ntf := msgFrame.Body.(*pbserver.ChallengeSendLoseRewardNtf)
	logger.Info("sendLoseReward  ntf.LoseUsers:%v, int(ntf.RoundIndex:%v), int(ntf.WinUserId:%v)  ntf.WinUsers:%v", ntf.LoseUers, int(ntf.RoundIndex), int(ntf.WinUserId), ntf.WinUsers)
	m.Challenge.SendLoseReward(ntf.LoseUers, ntf.WinUsers, int(ntf.RoundIndex), int(ntf.WinUserId))
	return nil, nil
}

func sendApplyUserUpdate(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {

	m.Challenge.Broadcast()
	return nil, nil
}

func setFirstGuildInfo(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	ntf := msgFrame.Body.(*pbserver.CcsToGsBroadShaBakeFirstGuildInfo)

	m.ShaBaKeCross.SetFirstGuildInfo(ntf)
	return nil, nil
}

func handlerRechargeReq(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {

	req := msgFrame.Body.(*pbserver.RechageCcsToGsReq)
	err := m.GetRecharge().NotifyBuy(req)
	ack := &pbserver.RechageGsToCcsAck{
		Result: 1,
		Msg:    "",
	}
	if err != nil {
		ack.Result = 0
		if ei, ok := err.(*errex.ErrorItem); ok {
			ack.Msg = strconv.Itoa(ei.Code)
		} else {
			ack.Msg = strconv.Itoa(gamedb.ERRUNKNOW.Code)
		}
	}
	return ack, nil
}

func handlerRechargeApplyReq(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {

	req := msgFrame.Body.(*pbserver.RechargeApplyReq)
	user := m.GetUserManager().GetUser(int(req.UserId))
	if user == nil {
		user = objs.NewUser()
		var err error
		user.User, err = modelGame.GetUserModel().GetByUserId(int(req.UserId))
		if err != nil {
			return nil,gamedb.ERRUNFOUNDUSER
		}
	}
	_, err, _, orderData := m.Recharge.Pay(user, int(req.PayNum), int(req.PayType), int(req.PayTypeId), true)
	if err != nil {
		return nil, err
	}
	return &pbserver.RechargeApplyAck{
		OrderId: orderData.OrderNo,
	}, nil

}

func handlerBanAccountInfoReq(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	req := msgFrame.Body.(*pbserver.BanInfoCcsToGsReq)
	m.GetUserManager().UserBanUpdate(req)
	return nil, nil
}

func handlerMailSendCCsToGsReq(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	req := msgFrame.Body.(*pbserver.MailSendCCsToGsReq)
	m.GetMail().MailSendByGM(req)
	return nil, nil
}

func handlerFuncStateUpdate(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	m.GetSystem().UpdateFuncState(true)
	return nil, nil
}

func handlerUpAnnouncementNowReq(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	m.GetAnnouncement().UpAnnouncementInfos()
	return nil, nil
}

func handlerUpPaoMaDengNowReq(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	m.GetAnnouncement().UpPaoMaDengInfoNow()
	return nil, nil
}
