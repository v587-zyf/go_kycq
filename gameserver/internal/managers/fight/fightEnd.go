package fight

import (
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"strconv"
)

func (this *Fight) FightEnd(endMsg *pbserver.FSFightEndNtf, isCross bool) {
	logger.Info("接收到战斗服发送来的战斗结果：%v", *endMsg)
	switch endMsg.FightType {
	case constFight.FIGHT_TYPE_STAGE_BOSS:
		this.huangUpBossFightEnd(endMsg.CpData)
	case constFight.FIGHT_TYPE_PERSON_BOSS:
		this.personFightEnd(int(endMsg.StageId), endMsg.CpData)
	case constFight.FIGHT_TYPE_TOWERBOSS:
		this.towerFightEnd(int(endMsg.StageId), endMsg.CpData)
	case constFight.FIGHT_TYPE_VIPBOSS:
		this.vipBossFightEnd(int(endMsg.StageId), endMsg.CpData)
	case constFight.FIGHT_TYPE_MINING:
		this.miningFightEnd(endMsg.CpData)
	case constFight.FIGHT_TYPE_MATERIAL,
		constFight.FIGHT_TYPE_EXPBOSS, constFight.FIGHT_TYPE_ARENA, constFight.FIGHT_TYPE_FIELD:
		//单人个人战斗结算
		this.userSingleFight(endMsg)
	case constFight.FIGHT_TYPE_FIELDBOSS:
		this.fieldBossFightResult(endMsg)
	case constFight.FIGHT_TYPE_WORLDBOSS:
		rank := common.ConvertInt32SliceToIntSlice(endMsg.Winners)
		lucker, _ := strconv.Atoi(string(endMsg.CpData))
		logger.Info("接收到战斗服发来的世界boss战斗结果，排名：%v,幸运者：%v", rank, lucker)
		err := this.GetWorldBoss().WorldBossFightEndAck(rank, lucker, int(endMsg.StageId))
		if err != nil {
			logger.Error("接收到战斗服发来的世界boss战斗结果,异常：%v", err)
		}
	case constFight.FIGHT_TYPE_DARKPALACE_BOSS:
		this.darkPalaceBossFightResult(endMsg)
	case constFight.FIGHT_TYPE_PAODIAN:
		this.paodianFightEnd(endMsg)
	case constFight.FIGHT_TYPE_GUILD_BONFIRE:
		this.guildBonfireEnd(endMsg)
	case constFight.FIGHT_TYPE_SHABAKE, constFight.FIGHT_TYPE_SHABAKE_NEW:
		this.guildShabakeEnd(endMsg)
	case constFight.FIGHT_TYPE_CROSS_WORLD_LEADER:
		this.crossWorldLeaderEnt(endMsg)
	case constFight.FIGHT_TYPE_CROSS_SHABAKE:
		this.guildShabakeCrossEnd(endMsg)
	case constFight.FIGHT_TYPE_ANCIENT_BOSS:
		this.ancientBossFightResult(endMsg)
	case constFight.FIGHT_TYPE_GUARDPILLAR:
		this.guardPillarFightEnd(endMsg)
	case constFight.FIGHT_TYPE_HELL_BOSS:
		this.hellBossFightResult(endMsg)
	case constFight.FIGHT_TYPE_MAGIC_TOWER:
		if this.GetSystem().IsCross() == isCross {
			this.magicTowerFightResult(endMsg)
		} else {
			logger.Debug("接收九层魔塔战斗结果，是来自跨服：%v,当前是否跨服：%v", isCross, this.GetSystem().IsCross())
		}
	case constFight.FIGHT_TYPE_DABAO:
		this.daBaoFightResult(endMsg)
	}
}

/**
 *  @Description: 单人个人战斗结算
 *  @param endMsg
 */
func (this *Fight) userSingleFight(endMsg *pbserver.FSFightEndNtf) {

	userId := 0
	isWin := false
	if len(endMsg.Winners) > 0 {
		userId = int(endMsg.Winners[0])
		isWin = true
	} else if len(endMsg.Losers) > 0 {
		userId = int(endMsg.Losers[0])
	}
	if userId > 0 {
		this.DispatchEvent(userId, endMsg, func(userId int, user *objs.User, data interface{}) {
			if user == nil {
				logger.Warn("接收到战斗服发送来的战斗结果：%v，玩家不在线", data)
				return
			}
			msgData := data.(*pbserver.FSFightEndNtf)
			if msgData.FightType == constFight.FIGHT_TYPE_MATERIAL {

				this.GetMaterialStage().MaterialStageFightResultNtf(user, isWin, int(msgData.StageId))
				result := pb.RESULTFLAG_FAIL
				if isWin {
					result = pb.RESULTFLAG_SUCCESS
				}
				kyEvent.StageEnd(user, int(msgData.StageId), result, user.FightStartTime, nil)
			} else if msgData.FightType == constFight.FIGHT_TYPE_EXPBOSS {
				cpData, _ := common.IntSliceFromString(string(msgData.CpData), ",")
				killMonster := 0
				getExp := 0
				if len(cpData) == 2 {
					killMonster = cpData[0]
					getExp = cpData[1]
				}
				this.GetExpStage().ExpStageFightResultNtf(user, killMonster, getExp, int(msgData.StageId))
				kyEvent.StageEnd(user, int(msgData.StageId), pb.RESULTFLAG_SUCCESS, user.FightStartTime, map[int]int{pb.ITEMID_EXP: getExp})
			} else if msgData.FightType == constFight.FIGHT_TYPE_ARENA {
				//challengeRanking, _ := strconv.Atoi(string(endMsg.CpData))
				//this.GetArena().ArenaFightResult(user, isWin, challengeRanking)
				this.GetCompetitve().CompetitveFightEndResult(user, isWin)
			} else if msgData.FightType == constFight.FIGHT_TYPE_FIELD {

				cpData, _ := common.IntSliceFromString(string(msgData.CpData), ",")
				challengeUid := 0
				isBeatBack := 0
				if len(cpData) == 2 {
					challengeUid = cpData[0]
					isBeatBack = cpData[1]
				}

				this.GetFieldFight().FieldFightFightEndResult(user, isWin, challengeUid, isBeatBack)
			}
		})
	} else {
		//if endMsg.FightType == constFight.FIGHT_TYPE_ARENA {
		//	challengeRanking, _ := strconv.Atoi(string(endMsg.CpData))
		//	this.GetArena().LeaveFight(challengeRanking)
		//}
		logger.Error("个人boss发送来的战斗结果消息异常，未找到玩家")
	}
}

func (this *Fight) miningFightEnd(data []byte) {

	fightResultMsg := &pbserver.MiningFightResultNtf{}
	err := fightResultMsg.Unmarshal(data)
	if err != nil {
		logger.Error("解析矿洞战斗结果异常：%v", err)
		return
	}

	userId := int(fightResultMsg.UserId)
	this.DispatchEvent(userId, nil, func(userId int, user *objs.User, data interface{}) {

		if fightResultMsg.IsRetake {
			this.GetMining().RobBackFightBack(user, int(fightResultMsg.MiningId), fightResultMsg.Result == pb.RESULTFLAG_SUCCESS)
		} else {
			this.GetMining().RobFightBack(user, int(fightResultMsg.MiningId), fightResultMsg.Result == pb.RESULTFLAG_SUCCESS)
		}
	})
}

func (this *Fight) towerFightEnd(stageId int, data []byte) {

	towerFightResult := &pbserver.TowerFightResult{}
	err := towerFightResult.Unmarshal(data)
	if err != nil {
		logger.Error("爬塔战斗结果解析异常：%v", err)
		return
	}

	userId := int(towerFightResult.UserId)
	this.DispatchEvent(userId, towerFightResult, func(userId int, user *objs.User, data interface{}) {

		result := towerFightResult.Result == pb.RESULTFLAG_SUCCESS
		items := common.ConvertMapInt32ToInt(towerFightResult.Items)
		this.GetTower().TowerFightEndAck(user, result, items)
		this.GetFirstDrop().CheckIsFirstDrop(user, items)
		kyEvent.StageEnd(user, stageId, int(towerFightResult.Result), user.FightStartTime, items)

	})
}

func (this *Fight) personFightEnd(stageId int, data []byte) {

	personBossResult := &pbserver.PersonFightResult{}
	err := personBossResult.Unmarshal(data)
	if err != nil {
		logger.Error("个人boss战斗结果解析异常：%v", err)
		return
	}

	userId := int(personBossResult.UserId)
	this.DispatchEvent(userId, personBossResult, func(userId int, user *objs.User, data interface{}) {

		result := personBossResult.Result == pb.RESULTFLAG_SUCCESS
		items := common.ConvertMapInt32ToInt(personBossResult.Items)
		this.GetPersonBoss().PersonBossFightResultNtf(user, stageId, result, items)
		this.GetFirstDrop().CheckIsFirstDrop(user, items)
		kyEvent.StageEnd(user, stageId, int(personBossResult.Result), user.FightStartTime, items)
	})
}

func (this *Fight) vipBossFightEnd(stageId int, data []byte) {

	resultMsg := &pbserver.VipBossFightResult{}
	err := resultMsg.Unmarshal(data)
	if err != nil {
		logger.Error("个人boss战斗结果解析异常：%v", err)
		return
	}

	userId := int(resultMsg.UserId)
	this.DispatchEvent(userId, resultMsg, func(userId int, user *objs.User, data interface{}) {

		result := resultMsg.Result == pb.RESULTFLAG_SUCCESS
		items := common.ConvertMapInt32ToInt(resultMsg.Items)
		this.GetVipBoss().VipBossFightResultNtf(user, result, stageId, items)
		this.GetFirstDrop().CheckIsFirstDrop(user, items)
		kyEvent.StageEnd(user, stageId, int(resultMsg.Result), user.FightStartTime, items)
	})
}

/**
 *  @Description: 野外boss战斗结果
 *  @param endMsg
 */
func (this *Fight) fieldBossFightResult(endMsg *pbserver.FSFightEndNtf) {
	winUserId := 0
	if len(endMsg.Winners) == 1 {
		winUserId = int(endMsg.Winners[0])
	} else {
		logger.Error("接收到战斗服发来的野外boss战斗结果,归属异常：%v", endMsg.Winners)
	}
	fieldResult := &pbserver.FieldBossResult{}
	err := fieldResult.Unmarshal(endMsg.CpData)
	if err != nil {
		logger.Error("解析战斗服发送来玩家野外boss战斗结果异常", err)
	}
	if fieldResult.SendWinner {
		this.DispatchEvent(winUserId, endMsg.StageId, func(userId int, user *objs.User, data interface{}) {
			if user == nil {
				logger.Warn("接收到战斗服发来的战斗结束，为找到相应玩家:%v", userId)
				return
			}

			items := make(map[int]int)
			for _, v := range fieldResult.UserPickItems[int32(userId)].Items {
				items[int(v.ItemId)] = int(v.ItemNum)
			}
			err := this.GetFieldBoss().FieldBossFightEndAck(user, winUserId, int(endMsg.StageId), items)
			if err != nil {
				logger.Error("接收到战斗服发来的野外boss战斗结果,玩家：%v,异常：%v", user.Id, err)
			}
			this.GetFirstDrop().CheckIsFirstDrop(user, items)
			kyEvent.StageEnd(user, int(endMsg.StageId), pb.RESULTFLAG_SUCCESS, user.FightStartTime, items)

		})
	} else {
		for _, v := range endMsg.Losers {
			userId := int(v)
			this.DispatchEvent(userId, endMsg.StageId, func(userId int, user *objs.User, data interface{}) {
				if user == nil {
					logger.Warn("接收到战斗服发来的战斗结束，为找到相应玩家:%v", userId)
					return
				}
				err := this.GetFieldBoss().FieldBossFightEndAck(user, winUserId, int(endMsg.StageId), nil)
				if err != nil {
					logger.Error("接收到战斗服发来的野外boss战斗结果,玩家：%v,异常：%v", user.Id, err)
				}
				kyEvent.StageEnd(user, int(endMsg.StageId), pb.RESULTFLAG_FAIL, user.FightStartTime, nil)
			})
		}
	}
}

/**
 *  @Description: 野外boss战斗结果
 *  @param endMsg
 */
func (this *Fight) darkPalaceBossFightResult(endMsg *pbserver.FSFightEndNtf) {
	winUserId := 0
	if len(endMsg.Winners) == 1 {
		winUserId = int(endMsg.Winners[0])
	} else {
		logger.Error("接收到战斗服发来的野外boss战斗结果,归属异常：%v", endMsg.Winners)
	}
	fieldResult := &pbserver.DarkPalaceBossResult{}
	err := fieldResult.Unmarshal(endMsg.CpData)
	if err != nil {
		logger.Error("解析战斗服发送来玩家野外boss战斗结果异常", err)
	}
	stageId := int(endMsg.StageId)
	if fieldResult.SendWinner {
		this.DispatchEvent(winUserId, 0, func(userId int, user *objs.User, data interface{}) {
			if user == nil {
				logger.Warn("接收到战斗服发来的战斗结束，为找到相应玩家:%v", userId)
				return
			}
			items := make(map[int]int)
			for _, v := range fieldResult.UserPickItems[int32(userId)].Items {
				items[int(v.ItemId)] = int(v.ItemNum)
			}
			this.GetDarkPalace().DarkPalaceFightResultNtf(user, stageId, winUserId, items, 0)
			this.GetFirstDrop().CheckIsFirstDrop(user, items)
			kyEvent.StageEnd(user, stageId, pb.RESULTFLAG_SUCCESS, user.FightStartTime, items)

		})
	} else {
		for _, v := range endMsg.Losers {
			userId := int(v)
			this.DispatchEvent(userId, 0, func(userId int, user *objs.User, data interface{}) {
				if user == nil {
					logger.Warn("接收到战斗服发来的战斗结束，为找到相应玩家:%v", userId)
					return
				}
				this.GetDarkPalace().DarkPalaceFightResultNtf(user, stageId, winUserId, nil, 0)
				kyEvent.StageEnd(user, stageId, pb.RESULTFLAG_FAIL, user.FightStartTime, nil)
			})
		}

		if len(fieldResult.Helper) > 0 {
			for k, v := range fieldResult.Helper {
				userId := int(k)
				this.DispatchEvent(userId, int(v), func(userId int, user *objs.User, data interface{}) {
					if user == nil {
						logger.Warn("接收到战斗服发来的战斗结束，为找到相应玩家:%v", userId)
						return
					}
					toHelpUserId := data.(int)
					this.GetDarkPalace().DarkPalaceFightResultNtf(user, stageId, winUserId, nil, int(toHelpUserId))
					kyEvent.StageEnd(user, stageId, pb.RESULTFLAG_SUCCESS, user.FightStartTime, nil)
				})
			}
		}
	}
}

/**
 *  @Description: 远古首领战斗结果
 *  @param endMsg
 */
func (this *Fight) ancientBossFightResult(endMsg *pbserver.FSFightEndNtf) {
	winUserId := 0
	if len(endMsg.Winners) == 1 {
		winUserId = int(endMsg.Winners[0])
	} else {
		logger.Error("接收到战斗服发来的远古首领战斗结果,归属异常：%v", endMsg.Winners)
	}
	fieldResult := &pbserver.AncientBossResult{}
	err := fieldResult.Unmarshal(endMsg.CpData)
	if err != nil {
		logger.Error("解析战斗服发送来玩家远古首领战斗结果异常", err)
	}
	if fieldResult.SendWinner {
		this.DispatchEvent(winUserId, int(endMsg.StageId), func(userId int, user *objs.User, data interface{}) {
			if user == nil {
				logger.Warn("接收到战斗服发来的战斗结束，未找到相应玩家:%v", userId)
				return
			}
			stageId := data.(int)

			items := make(map[int]int)
			for _, v := range fieldResult.UserPickItems[int32(userId)].Items {
				items[int(v.ItemId)] = int(v.ItemNum)
			}
			this.GetAncientBoss().AncientBossFightResult(user, winUserId, stageId, items)
			this.GetFirstDrop().CheckIsFirstDrop(user, items)
			kyEvent.StageEnd(user, int(endMsg.StageId), pb.RESULTFLAG_SUCCESS, user.FightStartTime, items)
		})
	} else {
		for _, v := range endMsg.Losers {
			userId := int(v)
			this.DispatchEvent(userId, int(endMsg.StageId), func(userId int, user *objs.User, data interface{}) {
				if user == nil {
					logger.Warn("接收到战斗服发来的战斗结束，未找到相应玩家:%v", userId)
					return
				}
				stageId := data.(int)
				this.GetAncientBoss().AncientBossFightResult(user, winUserId, stageId, nil)
				kyEvent.StageEnd(user, int(endMsg.StageId), pb.RESULTFLAG_FAIL, user.FightStartTime, nil)
			})
		}
	}
}

/**
 *  @Description: 炼狱首领战斗结果
 *  @param endMsg
 */
func (this *Fight) hellBossFightResult(endMsg *pbserver.FSFightEndNtf) {
	winUserId := 0
	if len(endMsg.Winners) == 1 {
		winUserId = int(endMsg.Winners[0])
	} else {
		logger.Error("接收到战斗服发来的炼狱首领战斗结果,归属异常：%v", endMsg.Winners)
	}
	fightResult := &pbserver.HellBossResult{}
	err := fightResult.Unmarshal(endMsg.CpData)
	if err != nil {
		logger.Error("解析战斗服发送来玩家炼狱首领战斗结果异常", err)
	}
	stageId := int(endMsg.StageId)
	if fightResult.SendWinner {
		this.DispatchEvent(winUserId, 0, func(userId int, user *objs.User, data interface{}) {
			if user == nil {
				logger.Warn("接收到战斗服发来的战斗结束，未找到相应玩家:%v", userId)
				return
			}
			items := make(map[int]int)
			for _, v := range fightResult.UserPickItems[int32(userId)].Items {
				items[int(v.ItemId)] = int(v.ItemNum)
			}
			this.GetHellBoss().HellBossFightResult(user, stageId, winUserId, items, 0)
			this.GetFirstDrop().CheckIsFirstDrop(user, items)
			kyEvent.StageEnd(user, stageId, pb.RESULTFLAG_SUCCESS, user.FightStartTime, items)
			//this.GetAnnouncement().FightSendSystemChat(user, items, stageId, pb.SCROLINGTYPE_HELL_BOSS)

		})
	} else {
		for _, v := range endMsg.Losers {
			userId := int(v)
			this.DispatchEvent(userId, 0, func(userId int, user *objs.User, data interface{}) {
				if user == nil {
					logger.Warn("接收到战斗服发来的战斗结束，未找到相应玩家:%v", userId)
					return
				}
				this.GetHellBoss().HellBossFightResult(user, stageId, winUserId, nil, 0)
				kyEvent.StageEnd(user, stageId, pb.RESULTFLAG_FAIL, user.FightStartTime, nil)
			})
		}

		if len(fightResult.Helper) > 0 {
			for k, v := range fightResult.Helper {
				userId := int(k)
				this.DispatchEvent(userId, int(v), func(userId int, user *objs.User, data interface{}) {
					if user == nil {
						logger.Warn("接收到战斗服发来的战斗结束，为找到相应玩家:%v", userId)
						return
					}
					toHelpUserId := data.(int)
					this.GetHellBoss().HellBossFightResult(user, stageId, winUserId, nil, toHelpUserId)
					kyEvent.StageEnd(user, stageId, pb.RESULTFLAG_SUCCESS, user.FightStartTime, nil)
				})
			}
		}
	}
}
