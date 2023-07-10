package handler

import (
	"cqserver/gamelibs/errex"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constServer"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

func init() {
	pb.Register(pb.CmdEnterPublicCopyReqId, HandlerEnterPublicCopy)
	pb.Register(pb.CmdFightPickUpReqId, HandlerPickUpReq)
	pb.Register(pb.CmdFightUserReliveReqId, HandlerUserReliveReq)
	pb.Register(pb.CmdFightCheerReqId, HandlerCheerReq)
	pb.Register(pb.CmdFightGetCheerNumReqId, HandlerGetCheerNumReq)
	pb.Register(pb.CmdFightPotionReqId, HandlerPotionReq)
	pb.Register(pb.CmdFightPotionCdReqId, HandlerPotionCdReq)
	pb.Register(pb.CmdFightCollectionReqId, HandlerCollectionReq)
	pb.Register(pb.CmdFightCollectionCancelReqId, HandlerCollectionCancelReq)
	pb.Register(pb.CmdFightApplyForHelpReqId, HandlerAskForHelpReq)
	pb.Register(pb.CmdFightAskForHelpResultReqId, HandlerAskForHelpResultReq)
	pb.Register(pb.CmdFightNpcEventReqId, HandlerFightNpcEventReq)
	pb.Register(pb.CmdGetBossFamilyInfoReqId, HandlerGetBossFamilyInfo)
	pb.Register(pb.CmdEnterBossFamilyReqId, HandlerEnterBossFamily)

	pbserver.Register(pbserver.CmdFSFightEndNtfId, HandlerFightEnd)
	pbserver.Register(pbserver.CmdFsResidentFightNtfId, HandlerResidentFight)
	pbserver.Register(pbserver.CmdFsFieldBossInfoNtfId, HandlerFieldBossInfoNtf)
	pbserver.Register(pbserver.CmdFsFieldBossDieUserInfoNtfId, HandlerFieldLeaderDieTimeNtf)
	pbserver.Register(pbserver.CmdWorldBossStatusNtfId, HandlerWorldBossStatusNtf)
	pbserver.Register(pbserver.CmdFSAddItemReqId, HandlerAddItemReq)
	pbserver.Register(pbserver.CmdFsSkillUseNtfId, HandlerUserSkillUse)
	pbserver.Register(pbserver.CmdExpStageKillMonsterNtfId, HandlerExpStageKillMonsterNtf)
	pbserver.Register(pbserver.CmdHangUpKillWaveNtfId, HandlerHangUpKillWaveNtf)
	pbserver.Register(pbserver.CmdGuildbonfireExpAddNtfId, HandlerGuildBonfireAddUserExpNtf)
	pbserver.Register(pbserver.CmdPaodianGoodsAddNtfId, HandlerPaodianGoodsAddNtf)
	pbserver.Register(pbserver.CmdFsToGsCollectionNtfId, HandlerCollectionNtf)
	pbserver.Register(pbserver.CmdFsTOGsClearSkillCdNtfId, HandlerClearSkillCdNtf)
	pbserver.Register(pbserver.CmdWorldLeaderFightRankNtfId, HandlerWorldLeaderRankNtf)
	pbserver.Register(pbserver.CmdDaBaoKillMonsterNtfId, HandlerDaBaoKillMonsterNtf)
	pbserver.Register(pbserver.CmdFsToGsShabakeKillBossNtfId, HandlerFsToGsShabakeKillBossNtf)
	pbserver.Register(pbserver.CmdFsToGsActorKillNtfId,HandlerActorKillNtf)
}

func HandlerEnterPublicCopy(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.EnterPublicCopyReq)
	user := conn.GetSession().(*managers.ClientSession).User
	ack, err := m.Fight.ClientEnterPublicCopy(user, int(req.StageId), int(req.Condition))
	return ack, nil, err
}

func HandlerPickUpReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.FightPickUpReq)
	user := conn.GetSession().(*managers.ClientSession).User
	ack, op, err := m.Fight.GsToFsPickUp(user, req.DropItemIds)
	if err != nil {
		err1 := errex.BuildClientErrorAck(err)
		ack = &pb.FightPickUpAck{
			Items: make(map[int32]*pb.ItemUnit),
			Err:   err1,
		}
		return ack, nil, nil
	}
	return ack, op, nil
}

func HandlerUserReliveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.FightUserReliveReq)
	user := conn.GetSession().(*managers.ClientSession).User
	ack, op, err := m.Fight.GsToFsRelive(user, req.SafeRelive)
	if err != nil {
		return nil, nil, err
	}
	return ack, op, err
}

func HandlerCheerReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFightCheer)
	op.SetOpTypeSecond(user.FightStageId)
	num, guildNum, err := m.Fight.CheerReq(user, op)
	if err != nil {
		return nil, nil, err
	}
	ack := &pb.FightCheerAck{
		CheerNum:      int32(num),
		GuildCheerNum: int32(guildNum),
	}
	return ack, op, nil
}

func HandlerGetCheerNumReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	num, guildNum, err := m.Fight.CheerGetUseNum(user)
	if err != nil {
		return nil, nil, err
	}
	return &pb.FightGetCheerNumNtf{
		CheerNum:      int32(num),
		GuildCheerNum: int32(guildNum),
	}, nil, nil

}

func HandlerPotionReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFightPotion)
	op.SetOpTypeSecond(user.FightStageId)
	coolDown, err := m.Fight.UsePotion(user, op)
	if err != nil {
		return nil, nil, err
	}
	ack := &pb.FightPotionAck{
		CoolDown:   int32(coolDown),
		ServerTime: int32(time.Now().Unix()),
		EndTime:    int32(time.Now().Unix() + int64(coolDown)),
	}
	return ack, op, nil
}

func HandlerPotionCdReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	coolDown, err := m.Fight.UsePotionCdReq(user)
	if err != nil {
		return nil, nil, err
	}
	ack := &pb.FightPotionAck{
		CoolDown:   int32(coolDown),
		ServerTime: int32(time.Now().Unix()),
		EndTime:    int32(time.Now().Unix() + int64(coolDown)),
	}

	return ack, nil, nil
}

func HandlerCollectionReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.FightCollectionReq)
	ack := &pb.FightCollectionAck{}
	err := m.Fight.CollectionReq(user, int(req.ObjId), ack)
	if err != nil {
		return nil, nil, err
	}
	return ack, nil, nil
}

func HandlerCollectionCancelReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.FightCollectionCancelReq)
	ack := &pb.FightCollectionCancelAck{}
	err := m.Fight.CollectionCancelReq(user, int(req.ObjId), ack)
	if err != nil {
		return nil, nil, err
	}
	ack.Result = true
	return ack, nil, nil
}

func HandlerAskForHelpReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.FightApplyForHelpReq)
	ack := &pb.FightApplyForHelpAck{
		HelpUserId: req.HelpUserId,
		Result:     pb.RESULTFLAG_SUCCESS,
	}
	err := m.Fight.ApplyForHelp(user, int(req.HelpUserId), int(req.Source))
	if err != nil {
		ack.Result = pb.RESULTFLAG_FAIL
		ack.FailReason = err.(*errex.ErrorItem).Message
		return ack, nil, err
	}
	return ack, nil, nil
}

func HandlerAskForHelpResultReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.FightAskForHelpResultReq)

	err := m.Fight.AskForHelpResult(user, req.IsAgree, int(req.ReqHelpUserId), int(req.HelpStageId))
	ack := &pb.FightAskForHelpResultAck{
		req.IsAgree, req.ReqHelpUserId, req.HelpStageId, "",
	}
	if err != nil {
		if ei, ok := err.(*errex.ErrorItem); ok {
			ack.EnterErr = ei.Message
		} else {
			return nil, nil, err
		}
	}

	return ack, nil, nil
}

func HandlerFightNpcEventReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.FightNpcEventReq)

	err := m.Fight.FightNpcEventReq(user, int(req.NpcId))

	return nil, nil, err
}

//bossfamily 打宝相关
func HandlerGetBossFamilyInfo(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	//user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.GetBossFamilyInfoReq)
	info, err := m.Fight.GetBossFamilyInfo(int(req.BossFamilyType))
	if err != nil {
		return nil, nil, err
	}
	return &pb.GetBossFamilyInfoAck{BossFamilyInfo: info}, nil, nil
}

func HandlerEnterBossFamily(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.EnterBossFamilyReq)
	err := m.Fight.EnterBossFamily(user, int(req.StageId))
	if err != nil {
		return nil, nil, err
	}
	return nil, nil, nil
}

/*********************************************************************************************************************/
/********************************************   以下为战斗服通讯     ****************************************************/
/*********************************************************************************************************************/

func HandlerFightEnd(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {

	endMsg := msgFrame.Body.(*pbserver.FSFightEndNtf)
	m.Fight.FightEnd(endMsg, conn.GetSession().GetId() == constServer.FIGHT_SESSIONID_CENTER)
	return nil, nil
}

func HandlerResidentFight(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	msg := msgFrame.Body.(*pbserver.FsResidentFightNtf)
	m.Fight.RecordResidentFight(msg, conn.GetSession().GetId() == constServer.FIGHT_SESSIONID_CENTER)
	return nil, nil
}

func HandlerFieldBossInfoNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	msg := msgFrame.Body.(*pbserver.FsFieldBossInfoNtf)
	m.Fight.HandlerFieldBossInfoNtf(msg, conn.GetSession().GetId() == constServer.FIGHT_SESSIONID_CENTER)
	return nil, nil
}

func HandlerWorldBossStatusNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	msg := msgFrame.Body.(*pbserver.WorldBossStatusNtf)
	m.WorldBoss.WorldBossInfoNtf()
	m.WorldLeader.SendClientWorldLeaderStart(int(msg.StageId))
	return nil, nil
}

func HandlerAddItemReq(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	msg := msgFrame.Body.(*pbserver.FSAddItemReq)
	ack, err := m.Fight.FsToGsAddItem(msg)
	return ack, err
}

func HandlerUserSkillUse(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	msg := msgFrame.Body.(*pbserver.FsSkillUseNtf)
	m.Fight.HandlerUserSkillUse(msg)
	return nil, nil
}

func HandlerActorKillNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	msg := msgFrame.Body.(*pbserver.FsToGsActorKillNtf)
	m.Fight.ActorKillNtf(msg)
	return nil, nil
}

func HandlerExpStageKillMonsterNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	msg := msgFrame.Body.(*pbserver.ExpStageKillMonsterNtf)
	m.Fight.HandlerExpStageKillMonsterNtf(int(msg.UserId))
	return nil, nil
}

func HandlerHangUpKillWaveNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	msg := msgFrame.Body.(*pbserver.HangUpKillWaveNtf)
	m.Fight.HangUpUserKillWave(msg)
	return nil, nil
}

func HandlerDaBaoKillMonsterNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	msg := msgFrame.Body.(*pbserver.DaBaoKillMonsterNtf)
	m.Fight.DaBaoKillMonster(msg)
	return nil, nil
}

func HandlerGuildBonfireAddUserExpNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {

	msg := msgFrame.Body.(*pbserver.GuildbonfireExpAddNtf)
	m.Fight.GuildBonfireUserAddExp(msg)
	return nil, nil
}

func HandlerPaodianGoodsAddNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	msg := msgFrame.Body.(*pbserver.PaodianGoodsAddNtf)
	m.Fight.PaodianGoodsAdd(msg)
	return nil, nil
}

func HandlerCollectionNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	msg := msgFrame.Body.(*pbserver.FsToGsCollectionNtf)
	m.Fight.CollectionOverNtf(msg)
	return nil, nil
}

func HandlerClearSkillCdNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	msg := msgFrame.Body.(*pbserver.FsTOGsClearSkillCdNtf)
	m.Fight.ClearSkillCD(int(msg.UserId), int(msg.HeroIndex))
	return nil, nil
}

func HandlerWorldLeaderRankNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	msg := msgFrame.Body.(*pbserver.WorldLeaderFightRankNtf)
	m.Fight.CrossWorldLeaderRankNtf(msg)
	return nil, nil
}

func HandlerFieldLeaderDieTimeNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	msg := msgFrame.Body.(*pbserver.FsFieldBossDieUserInfoNtf)
	m.Fight.HandlerFieldBossDieUserInfoNtf(msg)
	return nil, nil
}

func HandlerFsToGsShabakeKillBossNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	endMsg := msgFrame.Body.(*pbserver.FsToGsShabakeKillBossNtf)
	m.Announcement.BroadCastFsKillInfo(endMsg.Infos)
	return nil, nil
}
