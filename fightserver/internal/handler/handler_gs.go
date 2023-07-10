package handler

import (
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/fightModule"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pbserver"
	"time"
)

func init() {
	pbserver.Register(pbserver.CmdHandShakeReqId, HanlderShakeReq)
	pbserver.Register(pbserver.CmdHandCloseNtfId, HanlderGameClose)
	pbserver.Register(pbserver.CmdFSCreateFightReqId, OnFSCreateFightReq)
	pbserver.Register(pbserver.CmdGsRouteMessageToFightId, RouteGameMessage)
	pbserver.Register(pbserver.CmdGSTOFSCheckFightReqId, HandlerCheckFightExist)
	pbserver.Register(pbserver.CmdGSTOFSGetFightIdReqId, HandlerGetFightId)
	pbserver.Register(pbserver.CmdGsToFsResidentFightReqId, handlerResidentfight)
	pbserver.Register(pbserver.CmdGsToFsGamedbReloadReqId, handlerReloadGamedb)
	pbserver.Register(pbserver.CmdBossFamilyBossInfoReqId, handlerBossFamilyInfoReq)

	RegisterServerHandler(pbserver.CmdHandCloseNtfId, GameClose)
	RegisterServerHandler(pbserver.CmdFSEnterFightReqId, EnterFightReq)
	RegisterServerHandler(pbserver.CmdFSLeaveFightReqId, LeaveFightReq)
	RegisterServerHandler(pbserver.CmdFSContinueFightReqId, ContinueFightReq)
	RegisterServerHandler(pbserver.CmdFSUpdateUserInfoNtfId, FSUpdateUserInfoNtf)
	RegisterServerHandler(pbserver.CmdFsRandomDeliveryNtfId, RandomDeliveryNtf)
	RegisterServerHandler(pbserver.CmdGsToFsUseItemNtfId, GsToFsUseItemNtf)
	RegisterServerHandler(pbserver.CmdGsTOFsPickUpReqId, GSToFsPickUpReq)
	RegisterServerHandler(pbserver.CmdGsToFSCheckUserReliveReqId, GsToFSCheckUserReliveReq)
	RegisterServerHandler(pbserver.CmdGsToFSUserReliveReqId, GsToFSkUserReliveReq)
	RegisterServerHandler(pbserver.CmdGSToFsUpdateUserFightModelId, GSToFsUpdateUserFightModel)
	RegisterServerHandler(pbserver.CmdMiningNewFightInfoReqId, MiningNewFightInfoReq)
	RegisterServerHandler(pbserver.CmdGsToFsGetCheerNumReqId, GetCheerNumReq)
	RegisterServerHandler(pbserver.CmdGsToFsCheerReqId, GetCheerReq)
	RegisterServerHandler(pbserver.CmdGsToFsGetPotionCdReqId, GetPotionCdReq)
	RegisterServerHandler(pbserver.CmdGsToFsUsePotionReqId, UsePotionReq)
	RegisterServerHandler(pbserver.CmdGsToFsCollectionReqId, CollectionReq)
	RegisterServerHandler(pbserver.CmdGsToFsCollectionCancelReqId, CollectionCancelReq)
	RegisterServerHandler(pbserver.CmdGsToFsUseFitReqId, HandlerUseFit)
	RegisterServerHandler(pbserver.CmdGsToFsFitCacelReqId, HandlerFitCacel)
	RegisterServerHandler(pbserver.CmdGsToFsUpdatePetReqId, handlerUpdatePet)
	RegisterServerHandler(pbserver.CmdGsToFsPickRedPacketInfoId, handlerUserRedPacketInfo)
	RegisterServerHandler(pbserver.CmdGsToFsGmReqId, handlerUserGmReq)
	RegisterServerHandler(pbserver.CmdGsToFsUseCutTreasureReqId, handlerUseCutTreasureReq)
	RegisterServerHandler(pbserver.CmdGsToFsUpdateUserElfReqId, handlerUpdateElfReq)
	RegisterServerHandler(pbserver.CmdGsToFsChangeToHelperReqId, handlerChangeToHelperReq)
	RegisterServerHandler(pbserver.CmdGsToFsFightNumChangeReqId, handlerFightNumChangeReq)
	RegisterServerHandler(pbserver.CmdGsToFsFightScoreLessReqId, handlerFightScoreLessReq)
	RegisterServerHandler(pbserver.CmdMagicTowerGetUserInfoReqId, handlerMagicTowerGetUserInfoReq)
	RegisterServerHandler(pbserver.CmdGsToFsFightNpcEventReqId, handlerFightNpcEventReq)
	RegisterServerHandler(pbserver.CmdDaBaoResumeEnergyReqId, handlerDaBaoResumeEnergyReq)

}

func HanlderGameClose(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	fightModule.GetFightMgr().Range(func(f base.Fight) bool {
		f.SendMessage(NewGSMessage(f, conn, msgFrame.TransId, pbserver.CmdHandCloseNtfId, msgFrame.Body))
		return false
	})
	return nil, nil
}

func HanlderShakeReq(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {

	fightModule.GetFightMgr().GsServerConnected()

	return nil, nil
}

func RouteGameMessage(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {

	msg := msgFrame.Body.(*pbserver.GsRouteMessageToFight)
	fightId := uint32(msg.FightId)
	cmdId := msg.CmdId
	req, err := pbserver.UnmarshalMsgByCmdId(msg.MsgData, uint16(cmdId))
	if err != nil {
		return nil, err
	}
	fight := fightModule.GetFightMgr().GetFight(fightId)
	if fight == nil {
		logger.Error("RouteToFight:can't found fightId:%v cmdId:%v", fightId, req)
		if cmdId == pbserver.CmdFSLeaveFightReqId {
			return &pbserver.FSLeaveFightAck{}, nil
		}
		return nil, gamedb.ERRFIGHTEND
	}

	logger.Debug("接收到game发来的消息,战斗：%v,命令 Id:%v，消息内容：%v", fight.GetId(), cmdId, req)

	fight.SendMessage(NewGSMessage(fight, conn, msgFrame.TransId, cmdId, req))
	return nil, nil
}

func HandlerCheckFightExist(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {

	req := msgFrame.Body.(*pbserver.GSTOFSCheckFightReq)
	fight := fightModule.GetFightMgr().GetFight(uint32(req.FightId))
	msg := &pbserver.FSTOGSCheckFightAck{}
	if fight != nil {
		msg.FightId = req.FightId
	}
	return msg, nil
}

func HandlerGetFightId(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {

	req := msgFrame.Body.(*pbserver.GSTOFSGetFightIdReq)
	fightId := fightModule.GetFightMgr().GetFightByStageId(int(req.StageId), int(req.Ext))
	msg := &pbserver.FSTOGSGetFightIdAck{
		FightId: int32(fightId),
	}
	return msg, nil
}

func handlerResidentfight(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {

	req := msgFrame.Body.(*pbserver.GsToFsResidentFightReq)
	msg := fightModule.GetFightMgr().GetResidentFightInfo()
	logger.Debug("接收到服务器：%v,发送来的请求常驻战斗信息", req.ServerId)
	if msg != nil {
		return msg, nil
	}
	return nil, gamedb.ERRUNKNOW
}

func handlerReloadGamedb(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {

	gamedb.Reload()
	ack := &pbserver.GsToFsGamedbReloadAck{}
	logger.Info("重新加载配置表")
	return ack, nil
}

func handlerBossFamilyInfoReq(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {

	req := msgFrame.Body.(*pbserver.BossFamilyBossInfoReq)
	info := fightModule.GetFightMgr().GetBossFamilyInfo(int(req.BossFamilyType))
	return &pbserver.BossFamilyBossInfoAck{
		common.ConvertMapIntToInt32(info),
	}, nil
}

func OnFSCreateFightReq(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {

	req := msgFrame.Body.(*pbserver.FSCreateFightReq)
	fightId, err := fightModule.GetFightMgr().CreateFight(int(req.StageId), req.CpData)
	if err != nil {
		return nil, err
	}
	ack := &pbserver.FSCreateFightAck{
		FightId: fightId,
	}
	return ack, nil
}

func GameClose(fight base.Fight, msg nw.ProtoMessage) (nw.ProtoMessage, error) {

	req := msg.(*pbserver.HandCloseNtf)
	fight.Range(func(in base.Actor) bool {
		if in.GetType() == base.ActorTypeUser {
			userActor := in.(*actorPkg.UserActor)
			if userActor.HostId() == int(req.SessionNo) {
				fight.LeaveUser(userActor.GetUserId())
			}
		}
		return false
	})
	return nil, nil
}

func EnterFightReq(fight base.Fight, msg nw.ProtoMessage) (nw.ProtoMessage, error) {

	req := msg.(*pbserver.FSEnterFightReq)

	err := fight.UserEnter(req)
	if err != nil {
		return nil, err
	}

	return &pbserver.FSEnterFightAck{
		Refuse: true,
	}, nil
}

func LeaveFightReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {

	req := msgFrame.(*pbserver.FSLeaveFightReq)
	actor0 := fight.GetOnlineActor(req.ActorSessionId)
	if actor0 != nil {
		fight.LeaveUser(actor0.GetUserId())
		////玩家下线 清理玩家数据
		//if req.Reason == constFight.LEAVE_FIGHT_TYPE_OFFLINE {
		//	fightModule.PlayerM.RemovePlayer(actor0.GetUserId())
		//}
		logger.Info("收到玩家离开战斗 userId:%d_%v", actor0.GetUserId(), actor0.NickName())
	}
	return &pbserver.FSLeaveFightAck{}, nil
}

func ContinueFightReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {

	req := msgFrame.(*pbserver.FSContinueFightReq)
	//TODO 暂时只有爬塔
	towerFight := fight.(*fightModule.TowerFight)
	towerFight.ContinueFight(int(req.StageId))
	return nil, nil
}

func FSUpdateUserInfoNtf(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {

	req := msgFrame.(*pbserver.FSUpdateUserInfoNtf)

	userActors := fight.GetUserByUserId(int(req.UserInfo.GetUserId()))
	if userActors == nil {
		logger.Error("玩家申请更新玩家数据异常，玩家未找到：%v", req.UserInfo.GetUserId())
		return nil, gamedb.ERRUSEENOUGH
	}

	//玩家数据一定在战斗里面的
	for k, _ := range req.UserInfo.Heros {
		if _, ok := userActors[int(k)]; ok {
			fight.UpdateUserFigntInfo(req.UserInfo, int(k))
		}
	}
	//更新主武将总血量
	mainActor := fight.GetUserMainActor(int(req.UserInfo.GetUserId()))
	mainActor.(base.ActorPlayer).GetPlayer().ResetSkill(req.UserInfo.PublicSkill)
	mainActor.(base.ActorPlayer).GetPlayer().CalcUserHpTotal()

	return &pbserver.FSUpdateUserInfoAck{}, nil
}

func RandomDeliveryNtf(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.FsRandomDeliveryNtf)
	fight.RandomDelivery(int(req.UserId), req.Rand)
	return nil, nil
}

func GsToFsUseItemNtf(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	//req := msgFrame.(*pbserver.GsToFsUseItemNtf)
	//fightModule.PlayerM.UseItem(fight, int(req.UserId), int(req.ItemId))
	return nil, nil
}

func GSToFsPickUpReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsTOFsPickUpReq)
	items, err := fight.PickUp(int(req.UseId), req.ObjIds, req.IsPick)
	if err != nil {
		return nil, err
	}
	ack := &pbserver.FsTOGsPickUpAck{
		Items:  items,
		IsPick: req.IsPick,
	}
	return ack, nil
}

func GsToFSCheckUserReliveReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {

	req := msgFrame.(*pbserver.GsToFSCheckUserReliveReq)
	userActors := fight.GetUserByUserId(int(req.UserId))
	if userActors == nil {
		return nil, gamedb.ERRUNFOUNDUSER
	}
	for _, v := range userActors {
		if v.GetProp().HpNow() > 0 {
			logger.Error("请求复活,玩家：%v,角色：%v,未死亡,当前血量：%v", req.UserId, v.NickName(), v.GetProp().HpNow())
			return nil, gamedb.ERRPLAYERNODIE
		}
	}

	mainActor := fight.GetUserMainActor(int(req.UserId))
	player := mainActor.(base.ActorPlayer).GetPlayer()

	//isDie, reliveTimes, reliveByIngotTimes, err := fightModule.PlayerM.GetUserReliveCheckInfo(int(req.UserId))
	return &pbserver.FsToGSCheckUserReliveAck{
		IsDie:              true,
		ReliveTimes:        int32(player.ReliveTimes()),
		ReliveByIngotTimes: int32(player.ReliveByIngotTimes()),
	}, nil

}

func GsToFSkUserReliveReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {

	req := msgFrame.(*pbserver.GsToFSUserReliveReq)

	err := fight.UserRelive(int(req.UserId), int(req.ReliveType))
	if err != nil {
		return nil, err
	}
	mainActor := fight.GetUserMainActor(int(req.UserId))
	player := mainActor.(base.ActorPlayer).GetPlayer()
	if req.ReliveType == constFight.RELIVE_ADDR_TYPE_SITU {
		player.SetReliveByIngotTimes(player.ReliveByIngotTimes() + 1)
	}
	player.SetReliveTimes(player.ReliveTimes() + 1)

	ack := &pbserver.FSToGsUserReliveAck{
		ReliveTimes:        int32(player.ReliveTimes()),
		ReliveByIngotTimes: int32(player.ReliveByIngotTimes()),
	}
	return ack, nil
}

func GSToFsUpdateUserFightModel(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GSToFsUpdateUserFightModel)
	userActors := fight.GetUserByUserId(int(req.UserId))
	if userActors == nil {
		logger.Error("玩家申请更新玩家数据异常，玩家未找到：%v", req.UserId)
		return nil, gamedb.ERRUNFOUNDUSER
	}
	for _, userActor := range userActors {
		userActor.(base.ActorUser).SetFightModel(int(req.FigthModel))
	}
	return &pbserver.FSUpdateUserInfoAck{Refuse: true}, nil
}

func MiningNewFightInfoReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.MiningNewFightInfoReq)
	ack := &pbserver.MiningNewFightInfoAck{}
	if f, ok := fight.(*fightModule.MiningFight); ok {
		ack.ReadyOk = f.NewFightInfo(req)
	} else {
		ack.ReadyOk = false
	}
	return ack, nil
}

func GetCheerNumReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsGetCheerNumReq)
	ack := &pbserver.FsToGsGetCheerNumAck{
		UserId: req.UserId,
	}
	if f, ok := fight.(base.IFightCheer); ok {
		mainActor := fight.GetUserMainActor(int(req.UserId))
		if mainActor == nil {
			ack.CheerNum = -1
			logger.Error("获取鼓舞次数异常,获取玩家数据异常:%v", req.UserId)
			return ack, nil
		}

		guildId := 0
		if u, ok := mainActor.(base.ActorUser); ok {
			guildId = u.GuildId()
		}
		num, guildNum := f.GetCheerNum(int(req.UserId), guildId)
		ack.CheerNum = int32(num)
		ack.GuildCheerNum = int32(guildNum)
	} else {
		ack.CheerNum = -1
		logger.Error("获取鼓舞次数异常，战斗未实现鼓舞:%v", fight.GetStageConf().Id)
	}
	return ack, nil
}

func GetCheerReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsCheerReq)
	ack := &pbserver.FsToGsCheerAck{
		Result: false,
	}
	if f, ok := fight.(base.IFightCheer); ok {
		f.FightUseCheer(fight, int(req.UserId))
		ack.Result = true
	} else {
		logger.Error("战斗未实现鼓舞", fight.GetStageConf().Id)
	}
	return ack, nil

}

func GetPotionCdReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsGetPotionCdReq)
	ack := &pbserver.FsToGsGetPotionCdAck{
		UserId: req.UserId,
	}
	if f, ok := fight.(base.IFightUsePotion); ok {
		num := f.GetUserPotionCooldown(int(req.UserId))
		ack.UseTime = num
	} else {
		ack.UseTime = -1
		logger.Error("获取药水使用异常，战斗未实现药水使用")
	}
	return ack, nil
}

func UsePotionReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsUsePotionReq)
	ack := &pbserver.FsToGsUsePotionAck{
		Result: false,
	}
	if f, ok := fight.(base.IFightUsePotion); ok {
		f.FightUsePotionFunc(fight, int(req.UserId))
		ack.Result = true
		ack.UseTime = f.GetUserPotionCooldown(int(req.UserId))
	} else {
		logger.Error("战斗未实现药水使用")
	}
	return ack, nil
}

func CollectionReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsCollectionReq)
	endTime, err := fight.Collection(int(req.UserId), int(req.ObjId))
	if err != nil {
		return nil, err
	}
	ack := &pbserver.FsToGsCollectionAck{
		StartTime: time.Now().Unix(),
		EndTime:   int64(endTime),
	}
	return ack, nil
}

func CollectionCancelReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsCollectionCancelReq)
	err := fight.CollectionCancel(int(req.UserId), int(req.ObjId))
	if err != nil {
		return nil, err
	}
	ack := &pbserver.FsToGsCollectionCancelAck{
		Result: true,
	}
	return ack, nil
}

func HandlerUseFit(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsUseFitReq)
	err := fight.UseFitReq(int(req.UserId), req.Fit)
	if err != nil {
		return nil, err
	}
	return &pbserver.FsToGsUseFitAck{
		IsSuccess: true,
	}, nil
}

func HandlerFitCacel(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsFitCacelReq)
	err := fight.FitCacelReq(int(req.UserId))
	if err != nil {
		return nil, err
	}
	return &pbserver.FsToGsFitCacelAck{
	}, nil
}

func handlerUpdatePet(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsUpdatePetReq)
	err := fight.UpdatePet(int(req.UserId), req.Pet)
	return &pbserver.FsToGsUpdatePetAck{}, err
}

func handlerUpdateElfReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsUpdateUserElfReq)
	err := fight.UpdateElf(int(req.UserId), req.Elf)
	if err != nil {
		return nil, err
	}
	return &pbserver.FsToGsUpdateUserElfAck{
		Result: true,
	}, nil
}

func handlerChangeToHelperReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsChangeToHelperReq)
	fight.ChangeUserToHelper(int(req.UserId), int(req.ToHelpUserId))
	return nil, nil
}

func handlerFightNumChangeReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsFightNumChangeReq)
	fight.PlayerFightNumChange(req)
	return nil, nil
}

func handlerFightScoreLessReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsFightScoreLessReq)
	score, err := fight.FightScoreLess(int(req.UserId), int(req.LessNum))
	if err != nil {
		return nil, err
	}
	ack := &pbserver.FsToGsFightScoreLessAck{
		UserId: req.UserId,
		Score:  int32(score),
	}
	return ack, nil
}

func handlerMagicTowerGetUserInfoReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.MagicTowerGetUserInfoReq)
	score, isGetAward, canGetAward := fight.(*fightModule.MagicTowerFight).GetMagicUserInfo(int(req.UserId), req.IsGetAward)
	ack := &pbserver.MagicTowerGetUserInfoAck{
		Score:       int32(score),
		IsGetAward:  isGetAward,
		CanGetAward: canGetAward,
	}
	return ack, nil
}

func handlerDaBaoResumeEnergyReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.DaBaoResumeEnergyReq)
	mainActor := fight.GetUserMainActor(int(req.UserId))
	mainActor.(base.ActorPlayer).GetPlayer().SetDaBaoEnergy(int(req.Energy))
	return nil, nil
}

func handlerUserRedPacketInfo(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsPickRedPacketInfo)
	fight.UpdateUserReqPacketInfo(int(req.UserId), req.RedPacket)
	return nil, nil
}

func handlerUserGmReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsGmReq)
	reMsg := fight.GmReq(req)
	return &pbserver.FsToGsGmAck{
		Result: reMsg,
	}, nil
}

func handlerUseCutTreasureReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsUseCutTreasureReq)
	mainActor := fight.GetUserMainActor(int(req.UserId))
	var err error
	if mainActor == nil {
		return nil, gamedb.ERRUNFOUNDUSER
	}

	if mainActor.GetProp().HpNow() <= 0 {
		return nil, gamedb.ERRPLAYERDIE
	}

	fitActor := fight.GetUserFitActor(int(req.UserId))
	if fitActor != nil {
		return nil, gamedb.ERRCUTTREASUREBYFIT
	}

	if u, ok := mainActor.(base.ActorUser); ok {

		err = u.UseCutTreasure(int(req.CutTreasureLv))
	} else {
		return nil, gamedb.ERRUNFOUNDUSER
	}
	if err != nil {
		return nil, err
	}
	return &pbserver.FsToGsUseCutTreasureAck{
		Result: true,
	}, nil
}
func handlerFightNpcEventReq(fight base.Fight, msgFrame nw.ProtoMessage) (nw.ProtoMessage, error) {
	req := msgFrame.(*pbserver.GsToFsFightNpcEventReq)
	err := fight.NpcEventReq(int(req.UserId), int(req.NpcId))
	if err != nil {
		return nil, err
	}
	return &pbserver.FsToGsFightNpcEventAck{}, nil
}
