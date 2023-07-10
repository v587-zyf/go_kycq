package manager

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/robots/conf"
	"fmt"
	"strings"
)

func init() {
	var cmdIds = []uint16{
		pb.CmdErrorAckId,
		pb.CmdResetNtfId,
		pb.CmdChatMessageNtfId,
		pb.CmdUserOnlineNtfId, pb.CmdUserOffLineNtfId,
		pb.CmdEnterGameAckId, pb.CmdCreateUserAckId,
		pb.CmdBagInfoNtfId, pb.CmdTopDataChangeNtfId, pb.CmdBagDataChangeNtfId, pb.CmdBagEquipDataChangeNtfId,
		pb.CmdStageFightEndNtfId, pb.CmdSceneEnterNtfId, pb.CmdSceneMoveNtfId, pb.CmdAttackEffectNtfId,
		pb.CmdFitHolyEquipComposeAckId, pb.CmdFitHolyEquipDeComposeAckId, pb.CmdFitHolyEquipWearAckId, pb.CmdFitHolyEquipRemoveAckId, pb.CmdFitHolyEquipSuitSkillChangeAckId,
		pb.CmdContRechargeNtfId,
		pb.CmdCreateGuildAckId, pb.CmdApplyJoinGuildAckId, pb.CmdGuildLoadInfoAckId,
		pb.CmdSkillUpLvAckId, pb.CmdSkillChangeWearAckId,
		pb.CmdCreateHeroAckId,
		pb.CmdEnterShaBaKeFightReqId,
		pb.CmdKickUserNtfId,
		pb.CmdAncientBossLoadAckId, pb.CmdAncientBossNtfId,
		pb.CmdAncientSkillActiveAckId, pb.CmdAncientSkillUpLvReqId, pb.CmdAncientSkillUpGradeReqId,
		pb.CmdTreasureShopRefreshNtfId,
		pb.CmdDaBaoMysteryEnergyNtfId,
		pb.CmdRankLoadAckId,
		pb.CmdItemUseAckId,
	}
	for _, cmdId := range cmdIds {
		pb.Register(cmdId, handlerFunc)
	}
}

func handlerFunc(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	robot := conn.GetSession().(*ClientSession).Robot
	var module string
	var msg nw.ProtoMessage
	switch p.(type) {
	case *pb.ErrorAck:
		module = "错误"
		msg = p.(*pb.ErrorAck)
		logger.Info("ERRORACK code:%v msg:%v", p.(*pb.ErrorAck).Code, p.(*pb.ErrorAck).Message)
		return nil, nil, nil
	case *pb.ResetNtf:
		module = "重置任务"
		msg = p.(*pb.ResetNtf)
	case *pb.ChatMessageNtf:
		module = "聊天信息"
		msg = p.(*pb.ChatMessageNtf)
		if p.(*pb.ChatMessageNtf).ToId == robot.user.Userid && p.(*pb.ChatMessageNtf).Type == pb.CHATTYPE_PRIVATE {
			msgArr := strings.Split(p.(*pb.ChatMessageNtf).Msg, "#")
			if len(msgArr) > 1 && msgArr[0] == "changeguildgm" {
				guildId := common.Str2Int(msgArr[1])
				msgs := make([]nw.ProtoMessage, 0)
				if guildId == 0 {
					userCfg := conf.Conf.Create["user"]
					guildName := common.Interface2Str(userCfg["guildName"])
					createGuildItem := gamedb.GetConf().CreateGuild
					msgs = append(msgs, robot.DebugAdd([]int32{int32(createGuildItem.ItemId)}, []int32{int32(createGuildItem.Count)}))
					msgs = append(msgs, robot.MakeMsg(pb.CmdCreateGuildReqId, fmt.Sprintf(`%s,%s,%s`, guildName, "", "")))
					msgs = append(msgs, robot.MakeMsg(pb.CmdJoinGuildCombatLimitReqId, fmt.Sprintf(`%d,%d`, 0, 1)))
				} else {
					msgs = append(msgs, robot.MakeMsg(pb.CmdApplyJoinGuildReqId, fmt.Sprintf(`%d`, guildId)))
				}
				robot.initDataPb = append(robot.initDataPb, msgs...)
			}
		}
	case *pb.UserOnlineNtf:
		module = "用户上线通知"
		msg = p.(*pb.UserOnlineNtf)
	case *pb.UserOffLineNtf:
		module = "用户下线通知"
		msg = p.(*pb.UserOffLineNtf)
	case *pb.EnterGameAck:
		module = "EnterGame进入游戏"
		msg = p.(*pb.EnterGameAck)
		robot.status = conf.STATUS_ENTERGAME_DONE
		robot.user = p.(*pb.EnterGameAck).User
	case *pb.CreateUserAck:
		module = "创建User"
		CreateUserAck := p.(*pb.CreateUserAck)
		if CreateUserAck.User == nil || CreateUserAck.FailReason != "" {
			logger.Error("创建玩家异常：%v", CreateUserAck.FailReason)
		} else {
			robot.user.NickName = p.(*pb.CreateUserAck).GetUser().GetNickName()
			robot.user = CreateUserAck.User
			robot.initData()
			robot.status = conf.STATUS_IN_GAME
		}

	case *pb.CreateHeroAck:
		module = "创建Hero"
		ack := p.(*pb.CreateHeroAck)
		robot.user.Heros = append(robot.user.Heros, ack.Hero)
		logger.Debug("-----------CreateHeroAck", len(robot.user.Heros))
	case *pb.BagInfoNtf:
		module = "背包数据"
		msg = p.(*pb.BagInfoNtf)
		data := p.(*pb.BagInfoNtf)
		logger.Info("背包格子数量：%v", data.BagMax)
		for _, v := range data.Items {
			logger.Info("bag Info %v_%v_%v ", v.Position, v.ItemId, v.Count)
		}
	case *pb.TopDataChangeNtf:
		module = "顶级道具变化"
		msg = p.(*pb.TopDataChangeNtf)
		data := p.(*pb.TopDataChangeNtf)
		for _, v := range data.ChangeInfos {
			logger.Info("topDataChange id:%v,change:%v,now:%v", v.Id, v.Change, v.NowNum)
		}
	case *pb.BagDataChangeNtf:
		module = "背包道具变化"
		msg = p.(*pb.BagDataChangeNtf)
		data := p.(*pb.BagDataChangeNtf)
		for _, v := range data.ChangeInfos {
			logger.Info("bagChange pos:%v itemId:%v,change:%v,now:%v", v.Position, v.ItemId, v.Change, v.NowNum)
		}
	case *pb.BagEquipDataChangeNtf:
		module = "背包装备道具变化"
		msg = p.(*pb.BagEquipDataChangeNtf)
		data := p.(*pb.BagEquipDataChangeNtf)
		for _, v := range data.ChangeInfos {
			logger.Info("bagEquipChange pos:%v itemId:%v,change:%v,now:%v", v.Position, v.ItemId, v.Change, v.NowNum)
		}
	case *pb.StageFightEndNtf:
		module = "Stage战斗结束"
		msg = p.(*pb.StageFightEndNtf)
	case *pb.SceneMoveNtf:
		module = "SceneMoveNtf"
		msg = p.(*pb.SceneMoveNtf)
	case *pb.AttackEffectNtf:
		module = "AttackEffectNtf"
		msg = p.(*pb.AttackEffectNtf)
	case *pb.FitHolyEquipComposeAck:
		module = "合体圣装升级"
		msg = p.(*pb.FitHolyEquipComposeAck)
	case *pb.FitHolyEquipDeComposeAck:
		module = "合体圣装分解"
		msg = p.(*pb.FitHolyEquipDeComposeAck)
	case *pb.FitHolyEquipWearAck:
		module = "合体圣装穿戴"
		msg = p.(*pb.FitHolyEquipWearAck)
	case *pb.FitHolyEquipRemoveAck:
		module = "合体圣装卸下"
		msg = p.(*pb.FitHolyEquipRemoveAck)
	case *pb.FitHolyEquipSuitSkillChangeAck:
		module = "合体圣装套装技能更换"
		msg = p.(*pb.FitHolyEquipSuitSkillChangeAck)
	case *pb.ContRechargeNtf:
		module = "连续充值变化"
		msg = p.(*pb.ContRechargeNtf)
	case *pb.CreateGuildAck:
		module = "创建公会"
		msg = p.(*pb.CreateGuildAck)
		robot.guildStatue = 2
	case *pb.ApplyJoinGuildAck:
		module = "加入公会"
		msg = p.(*pb.ApplyJoinGuildAck)
		robot.guildStatue = 2
	case *pb.SkillUpLvAck:
		module = "技能升级"
		msg = p.(*pb.SkillUpLvAck)
		ack := p.(*pb.SkillUpLvAck)
		has := false
		for _, h := range m.robot.user.Heros {
			if h.Index != ack.HeroIndex {
				continue
			}
			for k, v := range h.Skills {
				if v.SkillId == ack.Skill.SkillId {
					h.Skills[k] = v
					has = true
					break
				}
			}
			if !has {
				h.Skills = append(h.Skills, ack.Skill)
			}
		}
	case *pb.SkillChangeWearAck:
		module = "技能穿戴、卸下"
		msg = p.(*pb.SkillChangeWearAck)
		ack := p.(*pb.SkillChangeWearAck)
		for _, h := range m.robot.user.Heros {
			if h.Index != ack.HeroIndex {
				continue
			}
			h.SkillBag = ack.SkillBags
			logger.Debug("----------------SkillChangeWearAck", ack.HeroIndex, ack.SkillBags)
		}
	case *pb.GuildLoadInfoAck:
		module = "公会信息加载"
		msg = p.(*pb.GuildLoadInfoAck)
		ack := p.(*pb.GuildLoadInfoAck)
		if ack.GuildInfo != nil && ack.GuildInfo.GuildId > 0 {
			m.robot.guildInfo = ack.GuildInfo
			m.robot.guildStatue = 2
		} else {
			m.robot.changeGuild()
		}
	case *pb.KickUserNtf:
		logger.Info("机器人被踢下线：%v", m.robot.openId)
		m.robot.status = conf.STATUS_OFF
	case *pb.AncientBossLoadAck:
		module = "远古首领列表"
		msg = p.(*pb.AncientBossLoadAck)
	case *pb.AncientBossNtf:
		module = "远古首领详细信息"
		msg = p.(*pb.AncientBossNtf)
	case *pb.AncientSkillActiveAck:
		module = "远古神技激活"
		msg = p.(*pb.AncientSkillActiveAck)
	case *pb.AncientSkillUpLvAck:
		module = "远古神技升级"
		msg = p.(*pb.AncientSkillUpLvAck)
	case *pb.AncientSkillUpGradeAck:
		module = "远古神技升阶"
		msg = p.(*pb.AncientSkillUpGradeAck)
	case *pb.TreasureShopRefreshNtf:
		module = "多宝阁刷新"
		msg = p.(*pb.TreasureShopRefreshNtf)
	case *pb.DaBaoMysteryEnergyNtf:
		module = "打宝秘境"
		msg = p.(*pb.DaBaoMysteryEnergyNtf)
	case *pb.RankLoadAck:
		module = "排行榜"
		msg = p.(*pb.RankLoadAck)
	case *pb.ItemUseAck:
		module = "道具使用"
		msg = p.(*pb.ItemUseAck)
	}
	logger.Info("%v msg:%v", module, msg)
	return nil, nil, nil
}

//func init() {
//	pb.Register(pb.CmdFitHolyEquipComposeAckId, handlerFunc)
//
//	pb.Register(pb.CmdLimitedGiftNtfId, handlerLimitedGiftNtf)
//	pb.Register(pb.CmdFriendUserInfoAckId, handleFriendUserInfoAck)
//
//	pb.Register(pb.CmdErrorAckId, handleErrorAck)
//
//	pb.Register(pb.CmdEnterGameAckId, handleEnterGameAck)
//	pb.Register(pb.CmdCreateUserAckId, handleCreateUserAck)
//
//	pb.Register(pb.CmdTopDataChangeNtfId, handlerTopDataChange)
//	pb.Register(pb.CmdBagDataChangeNtfId, handlerBagChange)
//	pb.Register(pb.CmdBagEquipDataChangeNtfId, handlerBagEquipDataChangeNtf)
//
//	pb.Register(pb.CmdBagInfoNtfId, handlerBagInfo)
//	pb.Register(pb.CmdStageFightEndNtfId, func(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//
//		ack := p.(*pb.StageFightEndNtf)
//		fmt.Println("关卡结束数据", *ack, *ack.Goods)
//		return nil, nil, nil
//	})
//
//	pb.Register(pb.CmdSceneEnterNtfId, func(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//
//		msg := p.(*pb.SceneEnterNtf)
//		fmt.Println("接收到玩家进入场景消息：stageId:", msg.StageId)
//		for _, v := range msg.Objs {
//			fmt.Println("场景角色", v.ObjId)
//
//			if v.User != nil {
//				go func() {
//
//					robot := conn.GetSession().(*ClientSession).Robot
//
//					time.Sleep(2 * time.Second)
//					attMsg := &pb.AttackRpt{
//						SkillId: 10000,
//						Dir:     0,
//					}
//					robot.SendMessage(0, attMsg)
//					fmt.Println("推送服务器玩家攻击消息：", *attMsg)
//				}()
//			}
//		}
//		return nil, nil, nil
//	})
//
//	pb.Register(pb.CmdSceneMoveNtfId, func(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//		msg := p.(*pb.SceneMoveNtf)
//		fmt.Println("接收到玩家移动通知：", msg.ObjId, msg.Point)
//		return nil, nil, nil
//	})
//
//	pb.Register(pb.CmdAttackEffectNtfId, func(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//
//		fmt.Println("接收到服务器发来的攻击效果指令", p)
//		return nil, nil, nil
//	})
//}
//func handleErrorAck(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//	logger.Error("error ack:%v", p)
//	return nil, nil, nil
//}
//
//func handleEnterGameAck(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//	ack := p.(*pb.EnterGameAck)
//	robot := conn.GetSession().(*ClientSession).Robot
//	robot.status = STATUS_ENTERGAME_DONE
//	robot.user = ack.User
//	logger.Info("EnterGameAck:%v,%v", *ack, *ack.User)
//	return nil, nil, nil
//}
//
//func handleCreateUserAck(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//	robot := conn.GetSession().(*ClientSession).Robot
//	robot.status = STATUS_IN_GAME
//	ack := p.(*pb.CreateUserAck)
//	robot.user.NickName = ack.GetUser().GetNickName()
//	logger.Info("CreateUserAck, robot.openId:%s", robot.openId)
//	return nil, nil, nil
//}
//
//func handlerBagInfo(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//
//	ntf := p.(*pb.BagInfoNtf)
//	logger.Info("背包格子数量：%v", ntf.BagMax)
//	for _, v := range ntf.Items {
//		logger.Info("bag Info %v_%v_%v ", v.Position, v.ItemId, v.Count)
//	}
//	return nil, nil, nil
//}
//
//func handlerTopDataChange(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//
//	ntf := p.(*pb.TopDataChangeNtf)
//
//	for _, v := range ntf.ChangeInfos {
//		logger.Info("topDataChange id:%v,change:%v,now:%v", v.Id, v.Change, v.NowNum)
//	}
//	return nil, nil, nil
//}
//
//func handlerBagChange(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//
//	ntf := p.(*pb.BagDataChangeNtf)
//
//	for _, v := range ntf.ChangeInfos {
//		logger.Info("bagChange pos:%v itemId:%v,change:%v,now:%v", v.Position, v.ItemId, v.Change, v.NowNum)
//	}
//	return nil, nil, nil
//}
//
//func handlerBagEquipDataChangeNtf(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//	ntf := p.(*pb.BagEquipDataChangeNtf)
//
//	for _, v := range ntf.ChangeInfos {
//		logger.Info("bagEquipChange pos:%v itemId:%v,change:%v,now:%v", v.Position, v.ItemId, v.Change, v.NowNum)
//	}
//	return nil, nil, nil
//}
//
//func handlerLimitedGiftNtf(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//	ntf := p.(*pb.LimitedGiftNtf)
//
//	logger.Debug("ntf:%v", ntf)
//	return nil, nil, nil
//}
//
//func handleFriendUserInfoAck(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//	ack := p.(*pb.FriendUserInfoAck)
//
//	logger.Info("ack:%v", ack)
//	return nil, nil, nil
//}
