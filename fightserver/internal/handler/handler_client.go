package handler

import (
	"cqserver/fightserver/conf"
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/fightModule"
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbgt"
	"cqserver/protobuf/pbserver"
)

func init() {
	pbgt.Register(pbgt.CmdRouteMessageId, RouteClientMessage)
	RegisterClientHandler(pb.CmdSceneMoveRptId, HandleSceneMoveRpt)
	RegisterClientHandler(pb.CmdAttackRptId, HandleAttackRpt)
	RegisterClientHandler(pb.CmdFightHurtRankReqId, HandleFightHurtRank)
	RegisterClientHandler(pb.CmdGetBossOwnerChangReqId, HandleGetBossOwnerChangReq)
	RegisterClientHandler(pb.CmdFightEnterOkReqId, HandleFightEnterOk)
	RegisterClientHandler(pb.CmdFightStartCountDownOkReqId, HandleFightStartCountDownOk)
	RegisterClientHandler(pb.CmdPaodianTopUserReqId, HandlerPaodianTopUser)
	RegisterClientHandler(pb.CmdGetShabakeScoresReqId, HandlerShabakeScoreReq)
	RegisterClientHandler(pb.CmdGetFightBossInfosReqId, HandlerGetFightBossInfos)
}

func GateClose() {

}

func RouteClientMessage(conn nw.Conn, msgFrame *pbgt.MessageFrame) {

	pbserver.GetMsgPrototype(0)
	clientData := msgFrame.Body.([]byte)
	clientMsgFrame, err := pb.Unmarshal(clientData)
	if err != nil {
		logger.Error("unknown message, cmdId: %d,err:%v", pb.GetCmdId(clientData), err)
		return
	}
	cmdId := pb.GetCmdId(clientData)
	fight := fightModule.GetFightMgr().GetFight(msgFrame.Reserved) // Reserved为fightId
	if fight == nil {
		logger.Error("fight err fight not found, fightId: %d cmdId: %d", msgFrame.Reserved, pb.GetMsgName(cmdId))
		return
	}
	//if fight.GetOnlineActor(msgFrame.SessionId) == nil {
	//	logger.Error("fight err actor not found, fightId: %d cmdId: %d", msgFrame.Reserved, pb.GetMsgName(cmdId))
	//	return
	//}

	fight.SendMessage(NewClientMessage(fight, msgFrame.SessionId, clientMsgFrame))
}

func HandleSceneMoveRpt(fight base.Fight, actor1 base.Actor, msg nw.ProtoMessage) (nw.ProtoMessage, error) {

	req := msg.(*pb.SceneMoveRpt)
	handleActor := fight.GetActorByObjId(int(req.ObjId))
	if handleActor == nil {
		logger.Info("接收到客户端发来的移动的命令：玩家未找到,战斗类型：%v", fight.GetStageConf().Type)
		return nil, nil
	}
	if handleActor.GetUserId() != actor1.GetUserId() {
		logger.Info("接收到客户端发来的移动的命令：玩家不能移动,玩家错误,战斗类型：%v", fight.GetStageConf().Type)
		return nil, gamedb.ERRMOVEPARAM
	}
	point := fight.GetScene().GetPointByXY(int(req.Point.X), int(req.Point.Y))
	if point == nil || scene.DistanceByPoint(handleActor.Point(), point) > 1 {
		if point != nil {
			logger.Error("玩家移动失败：%v,stage:%v，坐标：(%v,%v),格子：%v,距离：%v", req.ObjId, fight.GetStageConf().Id, point.GetAllObject(), scene.DistanceByPoint(handleActor.Point(), point))
		} else {
			logger.Error("玩家移动失败：%v,stage:%v，坐标：(%v,%v),客户端发送来的坐标未找到", req.ObjId, fight.GetStageConf().Id, req.Point.X, req.Point.Y)
		}

		ack := &pb.SceneMoveNtf{
			Point: handleActor.Point().ToPbPoint(),
			ObjId: int32(handleActor.GetObjId()),
			Force: true,
		}
		return ack, nil
	}
	if !handleActor.CanMove() {
		logger.Info("接收到客户端发来的移动的命令：玩家不能移动,战斗类型：%v", fight.GetStageConf().Type)
		return nil, gamedb.ERRMOVECD
	}

	err := handleActor.MoveTo(point, int(req.MoveType), false, true)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func HandleAttackRpt(fight base.Fight, actor1 base.Actor, msg nw.ProtoMessage) (nw.ProtoMessage, error) {

	req := msg.(*pb.AttackRpt)
	var ack nw.ProtoMessage
	attackActor := fight.GetActorByObjId(int(req.ObjId))
	if attackActor == nil || attackActor.GetUserId() != actor1.GetUserId() {
		logger.Error("客户端发送来的攻击者未找到,请求者：%v,攻击者：%v，技能：%v，requestUserId:%v-attackUserId:%v", actor1.NickName(), req.ObjId, req.SkillId, actor1.GetUserId(), attackActor)
		//return errex.BuildClientErrorAck(gamedb.ERRUNFOUNDUSER), nil
		return base.SkillAttackEffect(actor1, nil, 0, make([]*pb.HurtEffect, 0), gamedb.ERRUNFOUNDUSER.Code), nil
	}
	logger.Debug("接收客户端发来的攻击指令：玩家：%v_%v_%v,技能：%v,方向：%v", actor1.NickName(), attackActor.NickName(), actor1.GetType(), req.SkillId, req.Dir, req.ObjIds)
	if attackActor.GetType() == pb.SCENEOBJTYPE_USER {
		user := attackActor.(*actorPkg.UserActor)
		ack = user.UserAttack(int(req.SkillId), req.Point, int(req.Dir), common.ConvertInt32SliceToIntSlice(req.ObjIds), req.IsElf)
	} else if attackActor.GetType() == pb.SCENEOBJTYPE_PET {
		pet := attackActor.(*actorPkg.PetActor)
		ack = pet.PetAttack(int(req.SkillId), req.Point, int(req.Dir), common.ConvertInt32SliceToIntSlice(req.ObjIds))
	} else if attackActor.GetType() == pb.SCENEOBJTYPE_FIT {
		fit := attackActor.(*actorPkg.FitActor)
		ack = fit.FitAttack(int(req.SkillId), req.Point, int(req.Dir), common.ConvertInt32SliceToIntSlice(req.ObjIds))
	} else if attackActor.GetType() == pb.SCENEOBJTYPE_SUMMON {
		summon := attackActor.(*actorPkg.SummonActor)
		ack = summon.SummonAttack(int(req.SkillId), req.Point, int(req.Dir), common.ConvertInt32SliceToIntSlice(req.ObjIds))
	}
	if conf.Conf.Sandbox {
		if a, ok := ack.(*pb.AttackEffectNtf); ok && a.Err > 0 {
			logger.Error("客户端发送来攻击指令异常,err:%v", a.Err)
		}
	}

	return ack, nil
}

func HandleFightHurtRank(fight base.Fight, actor1 base.Actor, msg nw.ProtoMessage) (nw.ProtoMessage, error) {

	if rFight, ok := fight.(base.IFightDamageRank); ok {
		ack := rFight.FightDamageRankGetRankInfos(actor1)
		return ack, nil
	}

	if fight.GetStageConf().Type == constFight.FIGHT_TYPE_MAGIC_TOWER {
		ack := fight.(*fightModule.MagicTowerFight).GetFightInfos()
		return ack, nil
	}

	logger.Warn("获取战斗伤害排行异常，战斗未实现排行榜信息", fight.GetStageConf().Type)
	return nil, nil
}

func HandleGetBossOwnerChangReq(fight base.Fight, actor1 base.Actor, msg nw.ProtoMessage) (nw.ProtoMessage, error) {
	var bossObjId, owner int
	switch fight.GetStageConf().Type {
	case constFight.FIGHT_TYPE_FIELDBOSS:
		bossObjId, owner = fight.(*fightModule.FieldBossFight).GetBossOwner()
	case constFight.FIGHT_TYPE_DARKPALACE_BOSS:
		bossObjId, owner = fight.(*fightModule.DarkPalaceBossFight).GetBossOwner()
	case constFight.FIGHT_TYPE_ANCIENT_BOSS:
		bossObjId, owner = fight.(*fightModule.AncientBossFight).GetBossOwner()
	case constFight.FIGHT_TYPE_HELL_BOSS:
		bossObjId, owner = fight.(*fightModule.HellBossFight).GetBossOwner()
	default:
		logger.Error("客户端获取boss归属，当前战斗不支持,fightType:%v", fight.GetStageConf().Type)
		return nil, nil
	}
	ack := &pb.BossOwnerChangNtf{
		ObjId:  int32(bossObjId),
		UserId: int32(owner),
	}
	return ack, nil
}

func HandleFightEnterOk(fight base.Fight, actor1 base.Actor, msg nw.ProtoMessage) (nw.ProtoMessage, error) {
	if rFight, ok := fight.GetContext().(base.IFightReady); ok {
		rFight.FightEnterOk(actor1)
	}
	logger.Warn("获取战斗伤害排行异常，战斗未实现排行榜信息", fight.GetStageConf().Type)
	return nil, nil

}

func HandleFightStartCountDownOk(fight base.Fight, actor1 base.Actor, msg nw.ProtoMessage) (nw.ProtoMessage, error) {
	if rFight, ok := fight.GetContext().(base.IFightReady); ok {
		rFight.FightStartCountDownOk(actor1)
	}
	return nil, nil
}

func HandlerPaodianTopUser(fight base.Fight, actor1 base.Actor, msg nw.ProtoMessage) (nw.ProtoMessage, error) {

	if rFight, ok := fight.(*fightModule.PaodianFight); ok {
		msg := rFight.GetTopUser()
		if msg != nil {
			return msg, nil
		}
	}
	return nil, nil
}

func HandlerShabakeScoreReq(fight base.Fight, actor1 base.Actor, msg nw.ProtoMessage) (nw.ProtoMessage, error) {
	if rFight, ok := fight.(*fightModule.ShabakeFight); ok {
		msg := rFight.ShabakeScoreRank(true)
		if msg != nil {
			return msg, nil
		}
	}
	return nil, nil
}

func HandlerGetFightBossInfos(fight base.Fight, actor1 base.Actor, msg nw.ProtoMessage) (nw.ProtoMessage, error) {

	bossInfos := fight.GetBossInfos()
	return &pb.GetFightBossInfosAck{
		BossInfos: bossInfos,
	}, nil
}
