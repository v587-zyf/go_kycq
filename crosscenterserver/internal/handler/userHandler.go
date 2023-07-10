package handler

import (
	"cqserver/gamelibs/rmodelCross"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pbserver"
)

func init() {
	pbserver.Register(pbserver.CmdSyncUserInfoNtfId, HandleUserInfoSync)
	pbserver.Register(pbserver.CmdSetDayRechargeNumNtfId, HandleSetDayRechargeNumNtf)
	pbserver.Register(pbserver.CmdChallengeAppuserUpNtfId, HandleChallengeAppuserUpNtf)
}

func HandleUserInfoSync(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	req := msgFrame.Body.(*pbserver.SyncUserInfoNtf)
	m.GetUser().UserInfoSync(req)
	return nil, nil
}

func HandleSetDayRechargeNumNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	req := msgFrame.Body.(*pbserver.SetDayRechargeNumNtf)
	logger.Info("SetDayRechargeNum  serverId:%v  rechargeNum:%v,", req.ServerId, req.RechargeNum)
	rmodelCross.GetUserCrossInfoRmodle().SetDayRechargeNum(req.ServerId, req.RechargeNum)
	return nil, nil
}

func HandleChallengeAppuserUpNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	req := msgFrame.Body.(*pbserver.ChallengeAppuserUpNtf)
	logger.Info("HandleChallengeAppuserUpNtf  crossFsId:%v ", req.CrossFsId)
	m.GetCcsChallenge().ToGsUp(int(req.CrossFsId))
	return nil, nil
}
