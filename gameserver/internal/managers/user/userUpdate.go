package user

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/combat"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func (this *UserManager) UpdateCombatRobot(user *objs.User, heroIndex int) {
	this.UpdateUserCombatForIsLogin(user, heroIndex, false, true)
}

func (this *UserManager) UpdateCombat(user *objs.User, heroIndex int) {
	this.GetTalent().TalentGeneral(user)
	this.UpdateUserCombatForIsLogin(user, heroIndex, false, false)
	this.GetCondition().RecordCondition(user, pb.CONDITION_USER_COMBAT, []int{0})
}

func (this *UserManager) SetConditionAttributeProp(hero *objs.Hero, types int, prop gamedb.IntMap) {

	hero.ConditionAttribute = make(map[int]map[int]int)
	if hero.ConditionAttribute[types] == nil {
		hero.ConditionAttribute[types] = make(map[int]int)
	}
	if prop != nil {
		for pid, pVal := range prop {
			hero.ConditionAttribute[types][pid] += pVal
		}
	}

}

// 设置远古宝物 条件属性
func (this *UserManager) SetAncientConditionAttributeProp(hero *objs.Hero, types, treasureId int, prop gamedb.IntMap) {

	hero.AncientTreasureConditionAttribute = make(map[int]map[int]map[int]int)
	if hero.AncientTreasureConditionAttribute[types] == nil {
		hero.AncientTreasureConditionAttribute[types] = make(map[int]map[int]int)
	}
	if hero.AncientTreasureConditionAttribute[types][treasureId] == nil {
		hero.AncientTreasureConditionAttribute[types][treasureId] = make(map[int]int)
	}
	if prop != nil {
		for pid, pVal := range prop {
			hero.AncientTreasureConditionAttribute[types][treasureId][pid] += pVal
		}
	}

}

func (this *UserManager) SetEquipStrengthenLink(user *objs.User, hero *objs.Hero, heroIndex int) {
	//装备强化连携属性
	var equipLinkCfg *gamedb.StrengthenlinkStrengthenCfg
	for _, cfg := range gamedb.GetEquipStrengthenLinkCfgs() {
		if check := this.GetCondition().CheckMulti(user, heroIndex, cfg.Condition); !check || (equipLinkCfg != nil && equipLinkCfg.Id > cfg.Id) {
			continue
		}
		equipLinkCfg = cfg
	}
	if equipLinkCfg != nil {
		this.SetConditionAttributeProp(hero, pb.CONDITIONATTRIBUTETYPE_STRENGTHEN_LINK, equipLinkCfg.Attribute)
	}

	//远古宝物升星  条件属性
	starProp := make(map[int]int)
	//远古宝物觉醒  条件属性
	jueXinProp := make(map[int]int)
	//远古宝物套装  条件属性
	taoZhuangProp := make(map[int]int)
	for treasureId, info := range user.AncientTreasure {
		starCfg := gamedb.GetAncientTreasureStarById(treasureId, info.Star)
		if starCfg != nil {
			starProp = this.GetAncientTreasure().GetConditionProp(user, starCfg.Type, starProp, starCfg.Attribute)
			if len(starProp) > 0 {
				this.SetAncientConditionAttributeProp(hero, pb.CONDITIONATTRIBUTETYPE_ANCIENT_TREASURE_STAR, treasureId, starProp)
			}
		}

		//觉醒属性加成
		if info.JueXinLv > 0 {
			jueXinCfg := gamedb.GetAncientTreasureJueXinCfg(treasureId)
			if jueXinCfg != nil {
				jueXinProp = this.GetAncientTreasure().GetConditionProp(user, jueXinCfg.Type, jueXinProp, jueXinCfg.Attribute)
				if len(jueXinProp) > 0 {
					this.SetAncientConditionAttributeProp(hero, pb.CONDITIONATTRIBUTETYPE_ANCIENT_TREASURE_JUE_XING, treasureId, jueXinProp)
				}
			}
		}

		allActiveTaoZhuang := this.GetAncientTreasure().GetAncientTreasureTaoZ(user)
		for taoId, num := range allActiveTaoZhuang {
			tapCfg := gamedb.GetTreasureSuitTreasureSuitCfg(taoId)
			if tapCfg == nil {
				continue
			}
			for i := 1; i <= num; i++ {
				if i == 1 {
					taoZhuangProp = this.GetAncientTreasure().GetConditionProp(user, tapCfg.Type1, taoZhuangProp, tapCfg.Attribute1)
				}

				if i == 2 {
					taoZhuangProp = this.GetAncientTreasure().GetConditionProp(user, tapCfg.Type2, taoZhuangProp, tapCfg.Attribute2)
				}

				if i == 3 {
					taoZhuangProp = this.GetAncientTreasure().GetConditionProp(user, tapCfg.Type3, taoZhuangProp, tapCfg.Attribute3)
				}
			}
			if len(taoZhuangProp) > 0 {
				this.SetConditionAttributeProp(hero, pb.CONDITIONATTRIBUTETYPE_ANCIENT_TREASURE_TAO_ZHUANG, taoZhuangProp)
			}
		}
	}
}

func (this *UserManager) UpdateUserCombatForIsLogin(user *objs.User, heroIndex int, isLogin bool, robot bool) {

	for k, v := range user.Heros {
		if heroIndex > 0 && heroIndex != k {
			continue
		}
		oldCombat := user.Heros[k].Combat
		oldLevelCombat := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_PLAYER_LV]
		//oldEquipCombat := user.ModuleCombat[pb.OPENSYSID_EQUIP]
		//
		////重新计算玩家战斗力

		//装备强化连携属性
		this.SetEquipStrengthenLink(user, v, k)

		combat.UpdatePropCombat(user, k)
		//newCombat := user.Heros[k].Combat

		//玩家战斗变化 超过0.5%，通知战斗服
		//if !isLogin && math.Abs(float64(newCombat-oldCombat)/float64(newCombat)) > 0.005 {
		if !isLogin {
			//通知战斗服
			user.UpdateFightUserHeroIndexFun(heroIndex)
		}

		//玩家武将排行榜
		if !robot {
			this.GetRank().Append(this.getJobCombatType(user.Heros[k].Job), user.Id, user.Heros[k].Combat, false, isLogin, false)
			//总战力榜
			allCombat := this.GetAllCombat(user, pb.RANKTYPE_COMBAT)
			this.GetRank().Append(pb.RANKTYPE_COMBAT, user.Id, allCombat, false, isLogin, false)
			//宝石总战力排行榜
			allCombat = this.GetAllCombat(user, pb.RANKTYPE_COMBAT_JEWEL)
			key := rmodel.Rank.GetDailyRankKey(pb.RANKTYPE_COMBAT_JEWEL, base.Conf.ServerId)
			score := rmodel.Rank.GetDailyRankSelfRankAndScore(key, user.Id)

			this.GetRank().Append(pb.RANKTYPE_COMBAT_JEWEL, user.Id, allCombat, false, isLogin, allCombat > int(score))
			//神翼总战力榜
			allCombat = this.GetAllCombat(user, pb.RANKTYPE_COMBAT_WING)
			this.GetRank().Append(pb.RANKTYPE_COMBAT_WING, user.Id, allCombat, false, isLogin, false)
			//神兵总战力榜
			allCombat = this.GetAllCombat(user, pb.RANKTYPE_COMBAT_SHENG_BING)
			this.GetRank().Append(pb.RANKTYPE_COMBAT_SHENG_BING, user.Id, allCombat, false, isLogin, false)
			//法宝总战力榜
			allCombat = this.GetAllCombat(user, pb.RANKTYPE_COMBAT_FA_BAO)
			this.GetRank().Append(pb.RANKTYPE_COMBAT_FA_BAO, user.Id, allCombat, false, isLogin, false)
			//绝学总战力榜
			allCombat = this.GetAllCombat(user, pb.RANKTYPE_COMBAT_JUE_XUE)
			this.GetRank().Append(pb.RANKTYPE_COMBAT_JUE_XUE, user.Id, allCombat, false, isLogin, false)
			//装备总战力榜
			allCombat = this.GetAllCombat(user, pb.RANKTYPE_COMBAT_EQUIP)
			key = rmodel.Rank.GetDailyRankKey(pb.RANKTYPE_COMBAT_EQUIP, base.Conf.ServerId)
			score = rmodel.Rank.GetDailyRankSelfRankAndScore(key, user.Id)
			this.GetRank().Append(pb.RANKTYPE_COMBAT_EQUIP, user.Id, allCombat, false, isLogin, allCombat > int(score))

			//阵法总战力榜
			allCombat = this.GetAllCombat(user, pb.RANKTYPE_COMBAT_ZHEN_FA)
			this.GetRank().Append(pb.RANKTYPE_COMBAT_ZHEN_FA, user.Id, allCombat, false, isLogin, false)
			//龙器总战力榜
			allCombat = this.GetAllCombat(user, pb.RANKTYPE_COMBAT_LONG_QI)
			this.GetRank().Append(pb.RANKTYPE_COMBAT_LONG_QI, user.Id, allCombat, false, isLogin, false)

			//战宠榜
			allCombat = this.GetAllCombat(user, pb.RANKTYPE_COMBAT_PET)
			this.GetRank().Append(pb.RANKTYPE_COMBAT_PET, user.Id, allCombat, false, isLogin, false)
			//远古宝物榜
			allCombat = this.GetAllCombat(user, pb.RANKTYPE_COMBAT_ANCIENT)
			this.GetRank().Append(pb.RANKTYPE_COMBAT_ANCIENT, user.Id, allCombat, false, isLogin, false)
			//装备和传世装备
			allCombat = this.GetAllCombat(user, pb.RANKTYPE_COMBAT_CHUAN_SHI)
			this.GetRank().Append(pb.RANKTYPE_COMBAT_CHUAN_SHI, user.Id, allCombat, false, isLogin, false)
			//合体
			allCombat = this.GetAllCombat(user, pb.RANKTYPE_COMBAT_FIT)
			this.GetRank().Append(pb.RANKTYPE_COMBAT_FIT, user.Id, allCombat, false, isLogin, false)
			//打宝神器
			allCombat = this.GetAllCombat(user, pb.RANKTYPE_COMBAT_DA_BAO)
			this.GetRank().Append(pb.RANKTYPE_COMBAT_DA_BAO, user.Id, allCombat, false, isLogin, false)
			//特戒榜
			allCombat = this.GetAllCombat(user, pb.RANKTYPE_COMBAT_RING)
			this.GetRank().Append(pb.RANKTYPE_COMBAT_RING, user.Id, allCombat, false, isLogin, false)

			this.combatRankAndLog(user, -1, pb.PROPERTYMODULE_PLAYER_LV, oldCombat, oldLevelCombat, isLogin)
		}
		kyEvent.HeroCombat(user, k)
	}
	userCombat := this.GetAllCombat(user, pb.RANKTYPE_COMBAT)
	user.Combat = userCombat
	user.Dirty = true
	//this.combatRankAndLog(user, -1, pb.OPENSYSID_EQUIP, oldCombat, oldEquipCombat, isLogin)
	//推送客户端战斗力改变
	if !isLogin && !robot {

		this.SendMessage(user, builder.BuildProperMsg(user, heroIndex), true)
	}
}

// 玩家战力改变及log
func (this *UserManager) combatRankAndLog(user *objs.User, rankType int, combatType int, userOldCombat int, moduleOldCombat int, isLogin bool) {

	if isLogin && moduleOldCombat == 0 {
		return
	}
	//if moduleOldCombat != user.ModuleCombat[combatType] {
	//	//m.Tlog.CombatFlow(user, userOldCombat, combatType, moduleOldCombat)
	//}
}

// 修改昵称
func (this *UserManager) ChangeHeroName(user *objs.User, heroIndex int, name string) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	err := base.CheckName(name)
	if err != nil {
		return err
	}
	strLen := len([]rune(name))
	if strLen < 2 || strLen > 6 {
		return gamedb.ERRNICKNAMELENGTHINVALID
	}
	hero.Name = name
	user.UpdateFightUserHeroIndexFun(heroIndex)
	return nil
}

func (this *UserManager) GetAllCombat(user *objs.User, rankType int) int {

	allCombat := 0
	switch rankType {
	case pb.RANKTYPE_COMBAT:
		for _, v := range user.Heros {
			allCombat += v.Combat
		}
		allCombat += user.PetCombat
	case pb.RANKTYPE_COMBAT_JEWEL: //宝石战力排行榜
		for k := range user.Heros {
			combat1 := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_JEWEL]
			allCombat += combat1
		}

	case pb.RANKTYPE_COMBAT_WING: //神翼战力排行榜
		for k := range user.Heros {
			combat1 := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_WING]
			allCombat += combat1
		}
	case pb.RANKTYPE_COMBAT_SHENG_BING: //神兵战力排行榜
		for k := range user.Heros {
			combat1 := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_GOD_EEQUIP]
			allCombat += combat1
		}
	case pb.RANKTYPE_COMBAT_FA_BAO: //法宝战力排行榜
		//法宝就是至尊法器
		for k := range user.Heros {
			combat1 := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_HOLYARMS]
			allCombat = combat1
		}
	case pb.RANKTYPE_COMBAT_JUE_XUE: //绝学战力排行榜
		for k := range user.Heros {
			combat1 := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_JUEXUE]
			allCombat += combat1
		}
	case pb.RANKTYPE_COMBAT_EQUIP: //装备战力排行榜 (传世装备战力也算在这里)
		for k := range user.Heros {
			combat1 := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_EQUIP_NORMAL]
			allCombat += combat1
		}

		for k := range user.Heros {
			combat1 := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_CHUAN_SHI_EQUIP]
			allCombat += combat1
		}

	case pb.RANKTYPE_COMBAT_ZHEN_FA: //阵法战力排行榜
		for k := range user.Heros {
			combat1 := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_MAGIC_CIRCLE]
			allCombat += combat1
		}

	case pb.RANKTYPE_COMBAT_LONG_QI: //龙器战力排行榜
		for k := range user.Heros {
			combat1 := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_DRAGON_EQUIP]
			allCombat += combat1
		}

	case pb.RANKTYPE_COMBAT_RING: //特戒榜
		for k := range user.Heros {
			combat1 := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_RING]
			allCombat += combat1
		}
	case pb.RANKTYPE_COMBAT_ANCIENT: //远古榜
		for k := range user.Heros {
			combat1 := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_YUAN_GU_BAO_WU]
			allCombat += combat1
		}
	case pb.RANKTYPE_COMBAT_CHUAN_SHI: //装备和传世装备战力之和排行
		for k := range user.Heros {
			combat1 := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_EQUIP_NORMAL] + user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_CHUAN_SHI_EQUIP]
			allCombat += combat1
		}
	case pb.RANKTYPE_COMBAT_FIT: //合体榜(包含合体圣装)
		for k := range user.Heros {
			combat1 := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_FIT] + user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_FIT_HOLY_EQUIP]
			allCombat += combat1
		}
	case pb.RANKTYPE_COMBAT_DA_BAO: //打宝神器榜
		for k := range user.Heros {
			combat1 := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_DABAO_EQUIP]
			allCombat += combat1
		}
	case pb.RANKTYPE_COMBAT_PET: // 战宠榜
		for k := range user.Heros {
			combat1 := user.Heros[k].ModuleCombat[pb.PROPERTYMODULE_PET]
			allCombat += combat1
		}
	}
	return allCombat
}

func (this *UserManager) updateDailyState(user *objs.User) {
	today := this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
	if user.DayStateRecord == nil || user.DayStateRecord.Day != today {
		dayData := user.DayStateRecord
		user.DayStateRecord = &model.DayStateRecord{
			Day:              today,
			MonthCardReceive: make(model.IntKv),
		}
		if len(dayData.MonthCardReceive) > 0 {
			user.DayStateRecord.MonthCardReceive = dayData.MonthCardReceive
		}
		user.DayStateRecord.RechargeResetTime = dayData.RechargeResetTime
		user.TreasureShop.RefreshFree = true
		user.Elf.RecoverLimit = make(map[int]int)
		user.DayStateRecord.DailyRecharge = 0
	}
}
