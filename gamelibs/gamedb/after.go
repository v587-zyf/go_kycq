package gamedb

import (
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/protobuf/pb"
	"sort"
	"strings"
)

// 组装表数据放这里
func (this *GameDb) Patch() {
	this.genPropsCombat()
	this.genGameText()
	this.genBagAddSpace()
	this.genEquipPropRand()
	this.genEquipStrength()
	this.genDropMap()
	this.genFabaoLevel()
	this.genFabaoSkill()
	this.genPersonalBoss()
	this.genWing()
	this.genAtlas()
	this.genFieldBoss()
	this.genWorldBoss()
	this.genMaterial()
	this.genRoleName()
	this.genVipBoss()
	this.genExpStage()
	this.genArenaBuy()
	this.genSkill()
	this.genCompose()
	this.genAwaken()
	this.genMainPr()
	this.genBless()
	this.genDictate()
	this.genJewel()
	this.genFashion()
	this.genSign()
	this.genInside()
	this.genHolyarms()
	this.genRing()
	this.genMining()
	this.genWash()
	this.genDarkPalace()
	this.genExpPool()
	this.genMagicCircle()
	this.genTalent()
	this.genDailyTask()
	this.genDailyReward()
	this.genMonthCard()
	this.genDayRankReward()
	this.genGiftCode()
	this.genAchievement()
	this.genLimitedGift()
	this.genWorldLeader()
	this.genWarOrder()
	this.genCard()
	this.genElf()
	this.genCutTreasure()
	this.genTreasure()
	this.genHolyBeast()
	this.genFitHolyEquip()
	this.genRobot()
	this.genTower()
	this.genXunLongType()
	this.genDailyRank()
	this.genChuanShi()
	this.genSpendRebate()
	this.genContRecharge()
	this.genAncientBoss()
	this.genTitle()
	this.genKillMonster()
	this.genAncientTreasure()
	this.genPet()
	this.genHellBoss()
	this.genDaBao()
	this.genLabel()
	this.genFirstDrop()
}

func (this *GameDb) genPropsCombat() {
	//this.PropertiesCombat = make(map[int]float32)
	//for _, property := range this.Properties {
	//	if property.Combat > 0 {
	//		this.PropertiesCombat[property.Id] = property.Combat
	//	}
	//}
}

func (this *GameDb) genGameText() {

	for _, v := range errSlice {
		if this.GameTextErrorTextCfgs[v.Code] == nil {
			continue
		}
		v.Code, v.Message = v.Code, this.GameTextErrorTextCfgs[v.Code].Chinese
		if strings.TrimSpace(this.GameTextErrorTextCfgs[v.Code].Language) != "" {
			v.Message = this.GameTextErrorTextCfgs[v.Code].Language
		}
	}

	for _, v := range this.GameTextCodeTextCfgs {
		key := strings.ToUpper(strings.TrimSpace(v.ConstName))
		TextMap[key] = v.Chinese
		if strings.TrimSpace(v.Language) != "" {
			TextMap[key] = v.Language
		}
	}
}

func (this *GameDb) genBagAddSpace() {

	this.BagSpaceAddCfgsMap = make(map[int][]*BagSpaceAddCfg)
	dataByType := make(map[int][]int, 0)
	for _, v := range this.BagSpaceAddCfgs {
		if dataByType[v.Type] == nil {
			dataByType[v.Type] = make([]int, 0)
		}
		dataByType[v.Type] = append(dataByType[v.Type], v.Id)
	}
	for types, data := range dataByType {
		if this.BagSpaceAddCfgsMap[types] == nil {
			this.BagSpaceAddCfgsMap[types] = make([]*BagSpaceAddCfg, 0)
		}
		sort.Ints(data)
		for _, id := range data {
			this.BagSpaceAddCfgsMap[types] = append(this.BagSpaceAddCfgsMap[types], this.BagSpaceAddCfgs[id])
		}
	}

}

func (this *GameDb) genEquipPropRand() {
	this.EquipRandPropsMap = make(map[int][]*RandRandCfg)
	for _, v := range this.RandRandCfgs {
		if this.EquipRandPropsMap[v.Group] == nil {
			this.EquipRandPropsMap[v.Group] = make([]*RandRandCfg, 0)
		}
		this.EquipRandPropsMap[v.Group] = append(this.EquipRandPropsMap[v.Group], v)
	}
}

func (this *GameDb) genEquipStrength() {
	this.EquipStrengthMap = make(map[int]map[int]*StrengthenStrengthenCfg)
	this.EquipStrengthMaxLvMap = make(map[int]int)
	for _, v := range this.StrengthenStrengthenCfgs {
		if this.EquipStrengthMap[v.Position] == nil {
			this.EquipStrengthMap[v.Position] = make(map[int]*StrengthenStrengthenCfg)
		}
		this.EquipStrengthMap[v.Position][v.Level] = v
		level := v.Level
		if this.EquipStrengthMaxLvMap[v.Position] < level {
			this.EquipStrengthMaxLvMap[v.Position] = level
		}
	}
}

func (this *GameDb) genDropMap() {
	this.DropMap = make(map[int][]*DropDropCfg)
	for _, v := range this.DropDropCfgs {
		if this.DropMap[v.Dropid] == nil {
			this.DropMap[v.Dropid] = make([]*DropDropCfg, 0)
		}
		this.DropMap[v.Dropid] = append(this.DropMap[v.Dropid], v)
	}

	this.DropSpecial = make(map[int][]*DropSpecialDropSpecialCfg)
	for _, v := range this.DropSpecialDropSpecialCfgs {
		if this.DropSpecial[v.Type] == nil {
			this.DropSpecial[v.Type] = make([]*DropSpecialDropSpecialCfg, 0)
		}
		this.DropSpecial[v.Type] = append(this.DropSpecial[v.Type], v)
	}
}

func (this *GameDb) genFabaoLevel() {
	this.FabaoLvMap = make(map[int]map[int]*FabaolevelFabaolevelCfg)
	this.FabaoMaxLvMap = make(map[int]int)
	for _, v := range this.FabaolevelFabaolevelCfgs {
		if this.FabaoLvMap[v.Fabaoid] == nil {
			this.FabaoLvMap[v.Fabaoid] = make(map[int]*FabaolevelFabaolevelCfg)
		}
		this.FabaoLvMap[v.Fabaoid][v.Level] = v
		level := v.Level
		if this.FabaoMaxLvMap[v.Fabaoid] < level {
			this.FabaoMaxLvMap[v.Fabaoid] = level
		}
	}
}

func (this *GameDb) genFabaoSkill() {
	this.FabaoSkillMap = make(map[int]map[int]*FabaoSkillFabaoSkillCfg)
	for _, v := range this.FabaoSkillFabaoSkillCfgs {
		if this.FabaoSkillMap[v.Fabao_id] == nil {
			this.FabaoSkillMap[v.Fabao_id] = make(map[int]*FabaoSkillFabaoSkillCfg)
		}
		this.FabaoSkillMap[v.Fabao_id][v.Id] = v
	}
}

func (this *GameDb) genPersonalBoss() {
	this.PersonalBossMap = make(map[int]*PersonalBossPersonalBossCfg)
	for _, v := range this.PersonalBossPersonalBossCfgs {
		this.PersonalBossMap[v.StageId] = v
	}
}

func (this *GameDb) genWing() {
	this.WingMaxStarMap = make(map[int]int)
	this.WingSpecialMaxLvMap = make(map[int]int)
	this.WingSpecialMap = make(map[int]map[int]*WingSpecialWingSpecialCfg)
	for _, v := range this.WingNewWingNewCfgs {
		star := v.Star
		if this.WingMaxStarMap[v.Order] < star {
			this.WingMaxStarMap[v.Order] = star
		}
	}
	for _, wingSpecialCfg := range this.WingSpecialWingSpecialCfgs {
		specialT, lv := wingSpecialCfg.Type, wingSpecialCfg.Level
		if this.WingSpecialMap[specialT] == nil {
			this.WingSpecialMap[specialT] = make(map[int]*WingSpecialWingSpecialCfg)
		}
		this.WingSpecialMap[specialT][lv] = wingSpecialCfg
		if this.WingSpecialMaxLvMap[specialT] < lv {
			this.WingSpecialMaxLvMap[specialT] = lv
		}
	}
}

func (this *GameDb) genAtlas() {
	this.AtlasStarMap = make(map[int]map[int]*AtlasStarAtlasStarCfg)
	this.AtlasMaxStarMap = make(map[int]int)
	for _, v := range this.AtlasStarAtlasStarCfgs {
		if this.AtlasStarMap[v.Type] == nil {
			this.AtlasStarMap[v.Type] = make(map[int]*AtlasStarAtlasStarCfg)
		}
		this.AtlasStarMap[v.Type][v.Star] = v
		star := v.Star
		if this.AtlasMaxStarMap[v.Type] < star {
			this.AtlasMaxStarMap[v.Type] = star
		}
	}

	this.AtlasUpgradeMap = make(map[int]map[int]*AtlasUpgradeAtlasUpgradeCfg)
	this.AtlasMaxUpgradeMap = make(map[int]int)
	for _, v := range this.AtlasUpgradeAtlasUpgradeCfgs {
		if this.AtlasUpgradeMap[v.AtlasGather] == nil {
			this.AtlasUpgradeMap[v.AtlasGather] = make(map[int]*AtlasUpgradeAtlasUpgradeCfg)
		}
		this.AtlasUpgradeMap[v.AtlasGather][v.Star] = v
		star := v.Star
		if this.AtlasMaxUpgradeMap[v.AtlasGather] < star {
			this.AtlasMaxUpgradeMap[v.AtlasGather] = star
		}
	}

	for pos := range this.AtlasPosAtlasPosCfgs {
		if this.AtlasWearMax < pos {
			this.AtlasWearMax = pos
		}
	}
}

func (this *GameDb) genFieldBoss() {
	this.FieldBossMap = make(map[int]map[int]*FieldBossFieldBossCfg)
	for _, cfg := range this.FieldBossFieldBossCfgs {
		area, stageId := cfg.Area, cfg.StageId
		if this.FieldBossMap[area] == nil {
			this.FieldBossMap[area] = make(map[int]*FieldBossFieldBossCfg)
		}
		this.FieldBossMap[area][stageId] = cfg
	}
}

func (this *GameDb) genWorldBoss() {
	this.WorldBossMap = make(map[int]*WorldBossWorldBossCfg)
	for _, v := range this.WorldBossWorldBossCfgs {
		this.WorldBossMap[v.Stageid] = v
	}
	// 排行奖励
	this.WorldRankMap = make(map[int]*WorldRankWorldRankCfg)
	for _, v := range this.WorldRankWorldRankCfgs {
		this.WorldRankMap[v.Rank] = v
		if this.worldRankMax < v.Rank {
			this.worldRankMax = v.Rank
		}
	}
	rankMax := this.worldRankMax
	var tmpConf *WorldRankWorldRankCfg
	for ; rankMax >= 1; rankMax-- {
		if cfg, ok := this.WorldRankMap[rankMax]; ok {
			tmpConf = cfg
		} else {
			this.WorldRankMap[rankMax] = tmpConf
		}
	}
}

func (this *GameDb) genMaterial() {
	this.MaterialMaxLvMap = make(map[int]int)
	this.MaterialStageIdMap = make(map[int]*MaterialStageMaterialStageCfg)
	this.MaterialStageMap = make(map[int]map[int]*MaterialStageMaterialStageCfg)
	this.MaterialStageTypeMap = make(map[int]map[int]*MaterialStageMaterialStageCfg)
	for _, v := range this.MaterialStageMaterialStageCfgs {
		materialType := v.Type
		if this.MaterialStageMap[materialType] == nil {
			this.MaterialStageMap[materialType] = make(map[int]*MaterialStageMaterialStageCfg)
		}
		this.MaterialStageMap[materialType][v.Level] = v
		this.MaterialStageIdMap[v.Stageid] = v

		if this.MaterialMaxLvMap[materialType] < v.Level {
			this.MaterialMaxLvMap[materialType] = v.Level
		}

		if this.MaterialStageTypeMap[materialType] == nil {
			this.MaterialStageTypeMap[materialType] = make(map[int]*MaterialStageMaterialStageCfg, 0)
		}
		this.MaterialStageTypeMap[materialType][v.Stageid] = v
	}
}

func (this *GameDb) genRoleName() {
	var firstSlice, maleSlice, feMaleSlice []string
	for _, roleFirstName := range this.RoleFirstnameRoleFirstnameCfgs {
		firstSlice = append(firstSlice, roleFirstName.FirstName)
	}
	for _, roleName := range this.RoleNameBaseCfgs {
		if roleName.Whole == 0 {
			if roleName.Sex == pb.SEX_MALE {
				maleSlice = append(maleSlice, roleName.Name)
			} else if roleName.Sex == pb.SEX_FEMALE {
				feMaleSlice = append(feMaleSlice, roleName.Name)
			} else {
				maleSlice = append(maleSlice, roleName.Name)
				feMaleSlice = append(feMaleSlice, roleName.Name)
			}
		} else {
			if roleName.Sex == pb.SEX_MALE {
				this.MaleRoleNameMap = append(this.MaleRoleNameMap, roleName.Name)
			} else if roleName.Sex == pb.SEX_FEMALE {
				this.FeMaleRoleNameMap = append(this.FeMaleRoleNameMap, roleName.Name)
			} else {
				this.MaleRoleNameMap = append(this.MaleRoleNameMap, roleName.Name)
				this.FeMaleRoleNameMap = append(this.FeMaleRoleNameMap, roleName.Name)
			}
		}
	}
	for _, firstItem := range firstSlice {
		for _, maleItem := range maleSlice {
			this.MaleRoleNameMap = append(this.MaleRoleNameMap, firstItem+maleItem)
		}
		for _, femaleItem := range feMaleSlice {
			this.FeMaleRoleNameMap = append(this.FeMaleRoleNameMap, firstItem+femaleItem)
		}
	}
}

func (this *GameDb) genVipBoss() {
	this.VipBossMap = make(map[int]*VipBossVipBossCfg)
	for _, v := range this.VipBossVipBossCfgs {
		this.VipBossMap[v.Stageid] = v
	}
}

func (this *GameDb) genExpStage() {
	this.ExpStageMap = make(map[int]*ExpStageExpStageCfg)
	this.ExpStageLayerMap = make(map[int]*ExpStageExpStageCfg)
	for _, v := range this.ExpStageExpStageCfgs {
		this.ExpStageMap[v.Stage_id] = v
		this.ExpStageLayerMap[v.Layer] = v
	}
}

func (this *GameDb) genArenaBuy() {
	this.ArenaBuyNumMap = make(map[int]*ArenaBuyArenaBuyCfg)
	for _, cfg := range this.ArenaBuyArenaBuyCfgs {
		this.ArenaBuyNumMap[cfg.Num] = cfg
		if this.ArenaMaxBuyNum < cfg.Num {
			this.ArenaMaxBuyNum = cfg.Num
		}
	}
}

func (this *GameDb) genSkill() {
	this.SkillLvMap = make(map[int]map[int]*SkillLevelSkillCfg)
	this.SkillMaxLvMap = make(map[int]int)
	for _, skillLevelSkillCfg := range this.SkillLevelSkillCfgs {
		skillLvId := skillLevelSkillCfg.Skillid
		skillLv := skillLevelSkillCfg.Level
		skillId := skillLvId / 100
		if this.SkillLvMap[skillId] == nil {
			this.SkillLvMap[skillId] = make(map[int]*SkillLevelSkillCfg)
		}
		this.SkillLvMap[skillId][skillLv] = skillLevelSkillCfg

		if this.SkillMaxLvMap[skillId] < skillLv {
			this.SkillMaxLvMap[skillId] = skillLv
		}
	}
}

func (this *GameDb) genCompose() {
	this.ComposeTypeMap = make(map[int]*ComposeTypeComposeTypeCfg)
	this.ComposeItemMap = make(map[int]*ComposeSubComposeSubCfg)
	for _, composeTypeCfg := range this.ComposeTypeComposeTypeCfgs {
		this.ComposeTypeMap[composeTypeCfg.SubTab] = composeTypeCfg
	}
	for _, composeSubCfg := range this.ComposeSubComposeSubCfgs {
		this.ComposeItemMap[composeSubCfg.Composeid.ItemId] = composeSubCfg
	}

	this.ComposeEquipMap = make(map[int]*ComposeEquipSubComposeEquipSubCfg)
	for _, cfg := range this.ComposeEquipSubComposeEquipSubCfgs {
		this.ComposeEquipMap[cfg.Composeid.ItemId] = cfg
	}
}

func (this *GameDb) genAwaken() {
	this.AwakenMaxLvMap = make(map[int]int)
	this.AwakenMap = make(map[int]map[int]*AwakenAwakenCfg)
	for _, awakenCfg := range this.AwakenAwakenCfgs {
		t, lv := awakenCfg.Type, awakenCfg.Level
		if this.AwakenMap[t] == nil {
			this.AwakenMap[t] = make(map[int]*AwakenAwakenCfg)
		}
		this.AwakenMap[t][lv] = awakenCfg

		if this.AwakenMaxLvMap[t] < lv {
			this.AwakenMaxLvMap[t] = lv
		}
	}
}

func (this *GameDb) genMainPr() {
	this.MainPrMap = make(map[int]map[int]*MainPrMainPrCfg)
	for _, mainPrCfg := range this.MainPrMainPrCfgs {
		if this.MainPrMap[mainPrCfg.Type] == nil {
			this.MainPrMap[mainPrCfg.Type] = make(map[int]*MainPrMainPrCfg)
		}
		this.MainPrMap[mainPrCfg.Type][mainPrCfg.Body] = mainPrCfg
	}
}

func (this *GameDb) genBless() {
	for id, _ := range this.BlessBlessCfgs {
		if this.BlessMaxLv < id {
			this.BlessMaxLv = id
		}
	}
}

func (this *GameDb) genDictate() {
	this.DictateMaxLvMap = make(map[int]int)
	this.DictateMap = make(map[int]map[int]*DictateEquipDictateEquipCfg)
	for _, cfg := range this.DictateEquipDictateEquipCfgs {
		body := cfg.Body
		grade := cfg.Grade
		if this.DictateMap[body] == nil {
			this.DictateMap[body] = make(map[int]*DictateEquipDictateEquipCfg)
		}
		this.DictateMap[body][grade] = cfg
		if this.DictateMaxLvMap[body] < grade {
			this.DictateMaxLvMap[body] = grade
		}
	}
}

func (this *GameDb) genJewel() {
	this.JewelItemIdMap = make(map[int]int)
	this.JewelMaxLvMap = make(map[int]int)
	this.JewelMap = make(map[int]map[int]*JewelJewelCfg)
	this.JewelTypeLvMap = make(map[int][]int)
	for _, jewelCfg := range this.JewelJewelCfgs {
		kind, lv := jewelCfg.Type, jewelCfg.Level
		if this.JewelMap[kind] == nil {
			this.JewelMap[kind] = make(map[int]*JewelJewelCfg)
		}
		this.JewelTypeLvMap[kind] = append(this.JewelTypeLvMap[kind], jewelCfg.Id)
		this.JewelMap[kind][lv] = jewelCfg
		if this.JewelMaxLvMap[jewelCfg.Type] < lv {
			this.JewelMaxLvMap[jewelCfg.Type] = lv
		}
		this.JewelItemIdMap[jewelCfg.Id] = 0
	}
	for _, slice := range this.JewelTypeLvMap {
		sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	}
}

func (this *GameDb) genFashion() {
	this.FashionMaxLv = make(map[int]int)
	this.FashionMap = make(map[int]map[int]*FashionFashionCfg)
	for _, v := range this.FashionFashionCfgs {
		if this.FashionMap[v.Fashion_id] == nil {
			this.FashionMap[v.Fashion_id] = make(map[int]*FashionFashionCfg)
		}
		if v.Level > this.FashionMaxLv[v.Fashion_id] {
			this.FashionMaxLv[v.Fashion_id] = v.Level
		}
		this.FashionMap[v.Fashion_id][v.Level] = v
	}
}

func (this *GameDb) genSign() {
	this.CumulativeSignMap = make(map[int]*CumulationsignCumulationsignCfg)
	for _, cfg := range this.CumulationsignCumulationsignCfgs {
		this.CumulativeSignMap[cfg.Days] = cfg
	}
}

func (this *GameDb) genInside() {
	this.InsideStarMap = make(map[int]map[int]map[int]int)
	for _, cfg := range this.InsideStarInsideStarCfgs {
		g, o, s, w := cfg.Grade, cfg.Order, cfg.Star, cfg.Weight
		if this.InsideStarMap[g] == nil {
			this.InsideStarMap[g] = make(map[int]map[int]int)
		}
		if this.InsideStarMap[g][o] == nil {
			this.InsideStarMap[g][o] = make(map[int]int)
		}
		this.InsideStarMap[g][o][s] = w
	}
	this.InsideArtMap = make(map[int]map[int]map[int]*InsideArtInsideArtCfg)
	this.InsideMaxStar = make(map[int]map[int]int)
	for _, cfg := range this.InsideArtInsideArtCfgs {
		g, o, s := cfg.Grade, cfg.Order, cfg.Star
		if this.InsideArtMap[g] == nil {
			this.InsideArtMap[g] = make(map[int]map[int]*InsideArtInsideArtCfg)
		}
		if this.InsideArtMap[g][o] == nil {
			this.InsideArtMap[g][o] = make(map[int]*InsideArtInsideArtCfg)
		}
		this.InsideArtMap[g][o][s] = cfg

		//阶数
		if this.InsideMaxStar[g] == nil {
			this.InsideMaxStar[g] = make(map[int]int)
		}
		if this.InsideMaxStar[g][o] < s {
			this.InsideMaxStar[g][o] = s
		}
	}

	this.InsideSkillMaxLv = make(map[int]int)
	this.InsideSkillMap = make(map[int]map[int]*InsideSkillInsideSkillCfg)
	for _, cfg := range this.InsideSkillInsideSkillCfgs {
		id, lv := cfg.Skill_id, cfg.Level
		if this.InsideSkillMap[id] == nil {
			this.InsideSkillMap[id] = make(map[int]*InsideSkillInsideSkillCfg)
		}
		this.InsideSkillMap[id][lv] = cfg
		if this.InsideSkillMaxLv[id] < lv {
			this.InsideSkillMaxLv[id] = lv
		}
	}
}

func (this *GameDb) genHolyarms() {
	this.HolyMaxLv = make(map[int]int)
	for _, cfg := range this.HolylevelHolylevelCfgs {
		if this.HolyMaxLv[cfg.Holy_id] < cfg.Level {
			this.HolyMaxLv[cfg.Holy_id] = cfg.Level
		}
	}

	this.HolyLvMap = make(map[int]map[int]*HolylevelHolylevelCfg)
	for _, cfg := range this.HolylevelHolylevelCfgs {
		hid, lv := cfg.Holy_id, cfg.Level
		if this.HolyLvMap[hid] == nil {
			this.HolyLvMap[hid] = make(map[int]*HolylevelHolylevelCfg)
		}
		this.HolyLvMap[hid][lv] = cfg
	}

	this.HolySkillMaxLv = make(map[int]map[int]int)
	this.HolySkillMap = make(map[int]map[int]map[int]*HolySkillHolySkillCfg)
	for _, cfg := range this.HolySkillHolySkillCfgs {
		hid, hlv, lv := cfg.Holy_id, cfg.Holy_level, cfg.Skill_level
		if this.HolySkillMaxLv[hid] == nil {
			this.HolySkillMaxLv[hid] = make(map[int]int)
		}
		if this.HolySkillMaxLv[hid][hlv] < lv {
			this.HolySkillMaxLv[hid][hlv] = lv
		}

		if this.HolySkillMap[hid] == nil {
			this.HolySkillMap[hid] = make(map[int]map[int]*HolySkillHolySkillCfg)
		}
		if this.HolySkillMap[hid][hlv] == nil {
			this.HolySkillMap[hid][hlv] = make(map[int]*HolySkillHolySkillCfg)
		}
		this.HolySkillMap[hid][hlv][lv] = cfg
	}
}

func (this *GameDb) genRing() {
	this.PhantomSkill1Map = make(map[int]map[int]map[int]*PhantomLevelPhantomLevelCfg)
	this.PhantomSkill2Map = make(map[int]map[int]map[int]*PhantomLevelPhantomLevelCfg)
	for _, cfg := range this.PhantomLevelPhantomLevelCfgs {
		phantom, pos1, lv1, pos2, lv2 := cfg.Phantom, cfg.Position1, cfg.Skill_level1, cfg.Position2, cfg.Skill_level2
		if this.PhantomSkill1Map[phantom] == nil {
			this.PhantomSkill1Map[phantom] = make(map[int]map[int]*PhantomLevelPhantomLevelCfg)
		}
		if this.PhantomSkill1Map[phantom][pos1] == nil {
			this.PhantomSkill1Map[phantom][pos1] = make(map[int]*PhantomLevelPhantomLevelCfg)
		}
		this.PhantomSkill1Map[phantom][pos1][lv1] = cfg

		if this.PhantomSkill2Map[phantom] == nil {
			this.PhantomSkill2Map[phantom] = make(map[int]map[int]*PhantomLevelPhantomLevelCfg)
		}
		if this.PhantomSkill2Map[phantom][pos2] == nil {
			this.PhantomSkill2Map[phantom][pos2] = make(map[int]*PhantomLevelPhantomLevelCfg)
		}
		this.PhantomSkill2Map[phantom][pos2][lv2] = cfg
	}
}

func (this *GameDb) genMining() {
	this.MiningLvMap = make(map[int]*MiningMiningCfg)
	for _, cfg := range this.MiningMiningCfgs {
		this.MiningLvMap[cfg.Lv] = cfg
		if this.MiningMaxLv < cfg.Lv {
			this.MiningMaxLv = cfg.Lv
		}
	}
}

func (this *GameDb) genWash() {
	this.WashMap = make(map[int]map[int]*WashWashCfg)
	for _, cfg := range this.WashWashCfgs {
		if this.WashMap[cfg.Type] == nil {
			this.WashMap[cfg.Type] = make(map[int]*WashWashCfg)
		}
		this.WashMap[cfg.Type][cfg.Order] = cfg
	}
}

func (this *GameDb) genDarkPalace() {
	this.DarkPalaceStageMap = make(map[int]*DarkPalaceBossDarkPalaceBossCfg)
	this.DarkPalaceBossMap = make(map[int]map[int]*DarkPalaceBossDarkPalaceBossCfg)
	for _, palaceBossCfg := range this.DarkPalaceBossDarkPalaceBossCfgs {
		floor, stageId := palaceBossCfg.Floor, palaceBossCfg.StageId
		if this.DarkPalaceBossMap[floor] == nil {
			this.DarkPalaceBossMap[floor] = make(map[int]*DarkPalaceBossDarkPalaceBossCfg)
		}
		this.DarkPalaceBossMap[floor][stageId] = palaceBossCfg
		this.DarkPalaceStageMap[stageId] = palaceBossCfg
	}
}

// 经验池数据组装
func (this *GameDb) genExpPool() {
	this.ExpPoolLvlAssembleCfg = make(map[int]map[int]*ExpLevelLevelCfg)
	for _, cfg := range this.ExpLevelLevelCfgs {
		if this.ExpPoolLvlAssembleCfg[cfg.Player] == nil {
			this.ExpPoolLvlAssembleCfg[cfg.Player] = make(map[int]*ExpLevelLevelCfg)
		}
		this.ExpPoolLvlAssembleCfg[cfg.Player][cfg.Level] = cfg
	}

	this.ExpPoolLvlBuff = make(map[int][]*WorldLevelBuffWorldLevelBuffCfg)
	for _, randCfg := range this.WorldLevelBuffWorldLevelBuffCfgs {
		if this.ExpPoolLvlBuff[randCfg.Type] == nil {
			this.ExpPoolLvlBuff[randCfg.Type] = make([]*WorldLevelBuffWorldLevelBuffCfg, 0)
		}
		this.ExpPoolLvlBuff[randCfg.Type] = append(this.ExpPoolLvlBuff[randCfg.Type], randCfg)
	}

}

func (this *GameDb) genMagicCircle() {
	this.MagicCircleLvMap = make(map[int]map[int]map[int]*MagicCircleLevelMagicCircleLevelCfg)
	for _, cfg := range this.MagicCircleLevelMagicCircleLevelCfgs {
		id, grade, lv := cfg.Type, cfg.Rank, cfg.Level
		if this.MagicCircleLvMap[id] == nil {
			this.MagicCircleLvMap[id] = make(map[int]map[int]*MagicCircleLevelMagicCircleLevelCfg)
		}
		if this.MagicCircleLvMap[id][grade] == nil {
			this.MagicCircleLvMap[id][grade] = make(map[int]*MagicCircleLevelMagicCircleLevelCfg)
		}
		this.MagicCircleLvMap[id][grade][lv] = cfg
	}
}

func (this *GameDb) genTalent() {
	this.TalentMap = make(map[int]map[int]int)
	for id, cfg := range this.TalentWayTalengWayCfgs {
		job := cfg.Profession
		if this.TalentMap[job] == nil {
			this.TalentMap[job] = make(map[int]int)
		}
		for i := 1; i < constConstant.COMPUTE_TEN_THOUSAND; i++ {
			stageCfg := this.TalentStageTalengStageCfgs[GetRealId(id, i)]
			if stageCfg == nil {
				break
			}
			for _, talentId := range stageCfg.TalentID {
				this.TalentMap[job][talentId] = id
			}
		}
	}
}
func (this *GameDb) genDailyTask() {
	this.DailyTaskDailytaskCfgsByType = make(map[int]*DailyTaskDailytaskCfg)
	for _, cfg := range this.DailyTaskDailytaskCfgs {
		this.DailyTaskDailytaskCfgsByType[cfg.Type] = cfg
	}
}

func (this *GameDb) genDailyReward() {

	this.DailyRewardByTypes = make(map[int][]*DailyRewardDailyRewardCfg)

	dataByType := make(map[int][]int, 0)
	for _, v := range this.DailyRewardDailyRewardCfgs {
		if dataByType[v.Type] == nil {
			dataByType[v.Type] = make([]int, 0)
		}
		dataByType[v.Type] = append(dataByType[v.Type], v.Id)
	}

	for types, data := range dataByType {
		if this.DailyRewardByTypes[types] == nil {
			this.DailyRewardByTypes[types] = make([]*DailyRewardDailyRewardCfg, 0)
		}
		sort.Ints(data)
		for _, id := range data {
			this.DailyRewardByTypes[types] = append(this.DailyRewardByTypes[types], this.DailyRewardDailyRewardCfgs[id])
		}
	}

}

func (this *GameDb) genMonthCard() {
	this.MonthCardMap = make(map[int]*MonthCardMonthCardCfg)
	for _, cfg := range this.MonthCardMonthCardCfgs {
		this.MonthCardMap[cfg.Type] = cfg
	}
}

func (this *GameDb) genDayRankReward() {

	this.DayRankingRewardDayRankingRewardCfgsByType = make(map[int][]*DayRankingRewardDayRankingRewardCfg)
	dataByType := make(map[int][]int, 0)
	for _, v := range this.DayRankingRewardDayRankingRewardCfgs {
		if dataByType[v.Type] == nil {
			dataByType[v.Type] = make([]int, 0)
		}
		dataByType[v.Type] = append(dataByType[v.Type], v.Id)
	}

	for types, data := range dataByType {
		if this.DayRankingRewardDayRankingRewardCfgsByType[types] == nil {
			this.DayRankingRewardDayRankingRewardCfgsByType[types] = make([]*DayRankingRewardDayRankingRewardCfg, 0)
		}
		sort.Ints(data)
		for _, id := range data {
			this.DayRankingRewardDayRankingRewardCfgsByType[types] = append(this.DayRankingRewardDayRankingRewardCfgsByType[types], this.DayRankingRewardDayRankingRewardCfgs[id])
		}
	}

}

func (this *GameDb) genGiftCode() {
	this.GiftCodeMap = make(map[string]*GiftCodeGiftCodeCfg)
	for _, cfg := range this.GiftCodeGiftCodeCfgs {
		this.GiftCodeMap[cfg.Code] = cfg
	}
}

func (this *GameDb) genAchievement() {
	this.AchievementAchievementCfgsByConditionType = make(map[int][]*AchievementAchievementCfg)
	this.AchievementConditionState = make(map[int]bool)
	this.AchievementConditionIdState = make(map[int]bool)
	this.AchievementConditionIdAndCondition = make(map[int]int)
	dataByType := make(map[int][]int, 0)
	for _, v := range this.AchievementAchievementCfgs {
		if dataByType[v.Condition] == nil {
			dataByType[v.Condition] = make([]int, 0)
		}
		dataByType[v.Condition] = append(dataByType[v.Condition], v.Id)
		if !this.AchievementConditionState[v.Condition] {
			this.AchievementConditionState[v.Condition] = true
		}
		if !this.AchievementConditionIdState[v.ConditionId] {
			this.AchievementConditionIdState[v.ConditionId] = true
		}
		if this.AchievementConditionIdAndCondition[v.ConditionId] <= 0 {
			this.AchievementConditionIdAndCondition[v.ConditionId] = v.Condition
		}
	}

	for condition, data := range dataByType {
		if this.AchievementAchievementCfgsByConditionType[condition] == nil {
			this.AchievementAchievementCfgsByConditionType[condition] = make([]*AchievementAchievementCfg, 0)
		}
		sort.Ints(data)
		for _, id := range data {
			this.AchievementAchievementCfgsByConditionType[condition] = append(this.AchievementAchievementCfgsByConditionType[condition], this.AchievementAchievementCfgs[id])
		}
	}

}

func (this *GameDb) genLimitedGift() {
	this.LimitedGiftMap = make(map[int]map[int]*LimitedGiftLimitedGiftCfg)
	for _, cfg := range this.LimitedGiftLimitedGiftCfgs {
		t, tlv := cfg.Type, cfg.Lv
		if this.LimitedGiftMap[t] == nil {
			this.LimitedGiftMap[t] = make(map[int]*LimitedGiftLimitedGiftCfg)
		}
		this.LimitedGiftMap[t][tlv] = cfg
	}
}

func (this *GameDb) genWorldLeader() {
	this.WorldLeaderSlice = make([]*WorldLeaderConfCfg, 0)
	this.WorldLeaderRewardSlice = make(map[int][]*WorldLeaderRewardWorldLeaderRewardCfg)
	ids := make([]int, 0)
	for _, cfg := range this.WorldLeaderConfCfgs {
		ids = append(ids, cfg.Id)
	}
	sort.Ints(ids)

	for _, id := range ids {
		this.WorldLeaderSlice = append(this.WorldLeaderSlice, this.WorldLeaderConfCfgs[id])
	}

	for _, cfg := range this.WorldLeaderRewardWorldLeaderRewardCfgs {
		if this.WorldLeaderRewardSlice[cfg.StageId] == nil {
			this.WorldLeaderRewardSlice[cfg.StageId] = make([]*WorldLeaderRewardWorldLeaderRewardCfg, 0)
		}
		this.WorldLeaderRewardSlice[cfg.StageId] = append(this.WorldLeaderRewardSlice[cfg.StageId], cfg)
	}

}

func (this *GameDb) genWarOrder() {
	this.WarOrderMaxLv = make(map[int]int)
	for id := range this.WarOrderLevelWarOrderLevelCfgs {
		season := id / constConstant.COMPUTE_TEN_THOUSAND
		lv := id % constConstant.COMPUTE_TEN_THOUSAND
		if this.WarOrderMaxLv[season] < lv {
			this.WarOrderMaxLv[season] = lv
		}
	}
	this.WarOrderTaskMap = make(map[int][]*WarOrderCycleTaskWarOrderCycleTaskCfg)
	for _, cfg := range this.WarOrderCycleTaskWarOrderCycleTaskCfgs {
		conditionT := cfg.ConditionType
		if this.WarOrderTaskMap[conditionT] == nil {
			this.WarOrderTaskMap[conditionT] = make([]*WarOrderCycleTaskWarOrderCycleTaskCfg, 0)
		}
		this.WarOrderTaskMap[conditionT] = append(this.WarOrderTaskMap[conditionT], cfg)
	}
	this.WarOrderWeekTaskMap = make(map[int][]*WarOrderWeekTaskWarOrderWeekTaskCfg)
	for _, cfg := range this.WarOrderWeekTaskWarOrderWeekTaskCfgs {
		conditionT := cfg.ConditionType
		if this.WarOrderWeekTaskMap[conditionT] == nil {
			this.WarOrderWeekTaskMap[conditionT] = make([]*WarOrderWeekTaskWarOrderWeekTaskCfg, 0)
		}
		this.WarOrderWeekTaskMap[conditionT] = append(this.WarOrderWeekTaskMap[conditionT], cfg)
	}
}

func (this *GameDb) genCard() {
	this.DrawCfgBySeasonAndType = make(map[int]map[int]*DrawDrawCfg)

	for _, cfg := range this.DrawDrawCfgs {
		if this.DrawCfgBySeasonAndType[cfg.Type1] == nil {
			this.DrawCfgBySeasonAndType[cfg.Type1] = make(map[int]*DrawDrawCfg)
		}
		this.DrawCfgBySeasonAndType[cfg.Type1][cfg.Type2] = cfg
	}

	this.DrawShopDrawShopCfgByShop = make(map[int]map[int]*DrawShopDrawShopCfg)
	for _, cfg := range this.DrawShopDrawShopCfgs {
		if this.DrawShopDrawShopCfgByShop[cfg.Type1] == nil {
			this.DrawShopDrawShopCfgByShop[cfg.Type1] = make(map[int]*DrawShopDrawShopCfg)
		}
		this.DrawShopDrawShopCfgByShop[cfg.Type1][cfg.Id] = cfg
	}

}

func (this *GameDb) genElf() {
	this.ElfRecoverMap = make(map[int]map[int]*ElfRecoverElfRecoverCfg)
	for _, cfg := range this.ElfRecoverElfRecoverCfgs {
		t, q := cfg.Type, cfg.Quality
		if this.ElfRecoverMap[t] == nil {
			this.ElfRecoverMap[t] = make(map[int]*ElfRecoverElfRecoverCfg)
		}
		this.ElfRecoverMap[t][q] = cfg
	}
	this.ElfSkillMaxLv = make(map[int]int)
	this.ElfSkillMap = make(map[int]map[int]*ElfSkillElfGrowCfg)
	for _, cfg := range this.ElfSkillElfGrowCfgs {
		sid, slv := cfg.SkillId, cfg.SkillLv
		if this.ElfSkillMap[sid] == nil {
			this.ElfSkillMap[sid] = make(map[int]*ElfSkillElfGrowCfg)
		}
		this.ElfSkillMap[sid][slv] = cfg
		if this.ElfSkillMaxLv[sid] < slv {
			this.ElfSkillMaxLv[sid] = slv
		}
	}
}

func (this *GameDb) genCutTreasure() {
	this.CutTreasureMap = make(map[int]*CutTreasureCutTreasureCfg)
	for _, cfg := range this.CutTreasureCutTreasureCfgs {
		this.CutTreasureMap[cfg.Level] = cfg
	}
}

func (this *GameDb) genTreasure() {
	this.TreasureCfgBySeasonAndType = make(map[int]map[int]*XunlongXunlongCfg)

	for _, cfg := range this.XunlongXunlongCfgs {
		if this.TreasureCfgBySeasonAndType[cfg.Type1] == nil {
			this.TreasureCfgBySeasonAndType[cfg.Type1] = make(map[int]*XunlongXunlongCfg)
		}
		this.TreasureCfgBySeasonAndType[cfg.Type1][cfg.Type] = cfg
	}

	this.TreasureCfgBySeasonAndRound = make(map[int]map[int]*XunlongRoundsXunlongRoundsCfg)
	this.XunLongRoundSliceCfg = make(map[int][]*XunlongRoundsXunlongRoundsCfg)
	for _, cfg := range this.XunlongRoundsXunlongRoundsCfgs {
		if this.TreasureCfgBySeasonAndRound[cfg.Type1] == nil {
			this.TreasureCfgBySeasonAndRound[cfg.Type1] = make(map[int]*XunlongRoundsXunlongRoundsCfg)
		}
		if this.XunLongRoundSliceCfg[cfg.Type1] == nil {
			this.XunLongRoundSliceCfg[cfg.Type1] = make([]*XunlongRoundsXunlongRoundsCfg, 0)
		}
		this.TreasureCfgBySeasonAndRound[cfg.Type1][cfg.Rounds] = cfg
		this.XunLongRoundSliceCfg[cfg.Type1] = append(this.XunLongRoundSliceCfg[cfg.Type1], cfg)
	}
}

func (this *GameDb) genHolyBeast() {
	this.HolyBeastCfgByTypeAndStar = make(map[int]map[int]*HolyBeastHolyBeastCfg)
	this.HolyBeastCfgByTypeMaxStar = make(map[int]int)
	for _, cfg := range this.HolyBeastHolyBeastCfgs {
		if this.HolyBeastCfgByTypeAndStar[cfg.Type] == nil {
			this.HolyBeastCfgByTypeAndStar[cfg.Type] = make(map[int]*HolyBeastHolyBeastCfg)
		}
		this.HolyBeastCfgByTypeAndStar[cfg.Type][cfg.Star] = cfg
		if cfg.Star > this.HolyBeastCfgByTypeMaxStar[cfg.Type] {
			this.HolyBeastCfgByTypeMaxStar[cfg.Type] = cfg.Star
		}
	}

}

func (this *GameDb) genFitHolyEquip() {
	this.FitHolyEquipMap = make(map[int]map[int]map[int]*FitHolyEquipFitHolyEquipCfg)
	for _, cfg := range this.FitHolyEquipFitHolyEquipCfgs {
		st, t, g := cfg.SuitType, cfg.Type, cfg.Grade
		if this.FitHolyEquipMap[st] == nil {
			this.FitHolyEquipMap[st] = make(map[int]map[int]*FitHolyEquipFitHolyEquipCfg)
		}
		if this.FitHolyEquipMap[st][t] == nil {
			this.FitHolyEquipMap[st][t] = make(map[int]*FitHolyEquipFitHolyEquipCfg)
		}
		this.FitHolyEquipMap[st][t][g] = cfg
	}
}

func (this *GameDb) genRobot() {
	this.RobotCfgSlice = make([]*RobotRobotCfg, 0)
	ids := make([]int, 0)
	for _, cfg := range this.RobotRobotCfgs {
		ids = append(ids, cfg.Id)
	}
	sort.Ints(ids)
	for _, id := range ids {
		this.RobotCfgSlice = append(this.RobotCfgSlice, this.RobotRobotCfgs[id])
	}
}

func (this *GameDb) genTower() {
	for id := range this.TowerTowerCfgs {
		if this.towerMaxLv < id {
			this.towerMaxLv = id
		}
	}
}

func (this *GameDb) genXunLongType() {
	this.XunLongTypeTime = make(map[int]int)
	for _, cfg := range this.XunlongXunlongCfgs {
		if this.XunLongTypeTime[cfg.Type1] == 0 {
			if cfg.Time != nil && len(cfg.Time) >= 2 {
				times := (cfg.Time[1] - cfg.Time[0] + 1) * 86400
				this.XunLongTypeTime[cfg.Type1] = times
			}
		}
	}
}

func (this *GameDb) genDailyRank() {
	this.DayRankingMarkSliceCfg = make(map[int][]*DayRankingMarkDayRankingMarkCfg)
	for _, cfg := range this.DayRankingMarkDayRankingMarkCfgs {
		if this.DayRankingMarkSliceCfg[cfg.Type] == nil {
			this.DayRankingMarkSliceCfg[cfg.Type] = make([]*DayRankingMarkDayRankingMarkCfg, 0)
		}
		this.DayRankingMarkSliceCfg[cfg.Type] = append(this.DayRankingMarkSliceCfg[cfg.Type], cfg)
	}
}

func (this *GameDb) genChuanShi() {
	this.ChuanShiSuitMap = make(map[int]map[int]*ChuanShiSuitTypeChuanShiSuitTypeCfg)
	for _, cfg := range this.ChuanShiSuitTypeChuanShiSuitTypeCfgs {
		t, lv := cfg.SuitType, cfg.Level
		if this.ChuanShiSuitMap[t] == nil {
			this.ChuanShiSuitMap[t] = make(map[int]*ChuanShiSuitTypeChuanShiSuitTypeCfg)
		}
		this.ChuanShiSuitMap[t][lv] = cfg
	}

	this.ChuanShiStrengthenMap = make(map[int]map[int]*ChuanShiStrengthenChuanShiStrengthenCfg)
	for _, cfg := range this.ChuanShiStrengthenChuanShiStrengthenCfgs {
		pos, lv := cfg.Position, cfg.Level
		if this.ChuanShiStrengthenMap[pos] == nil {
			this.ChuanShiStrengthenMap[pos] = make(map[int]*ChuanShiStrengthenChuanShiStrengthenCfg)
		}
		this.ChuanShiStrengthenMap[pos][lv] = cfg
	}

	this.ChuanShiStrengthenSuitMap = make(map[int]map[int]*ChuanShiStrengthenLinkChuanShiStrengthenLinkCfg)
	for _, cfg := range this.ChuanShiStrengthenLinkChuanShiStrengthenLinkCfgs {
		t, lv := cfg.Type, cfg.Level
		if this.ChuanShiStrengthenSuitMap[t] == nil {
			this.ChuanShiStrengthenSuitMap[t] = make(map[int]*ChuanShiStrengthenLinkChuanShiStrengthenLinkCfg)
		}
		this.ChuanShiStrengthenSuitMap[t][lv] = cfg
	}
}

func (this *GameDb) genSpendRebate() {
	this.SpendRebateMap = make(map[int]*SpendrebatesSpendrebatesCfg)
	for _, cfg := range this.SpendrebatesSpendrebatesCfgs {
		this.SpendRebateMap[cfg.Type1] = cfg
	}
}

func (this *GameDb) genContRecharge() {
	this.ContRechargeMap = make(map[int]*ContRechargeContRechargeCfg)
	this.ContRechargeTypeMap = make(map[int][]*ContRechargeContRechargeCfg)
	for _, cfg := range this.ContRechargeContRechargeCfgs {
		t := cfg.Type
		this.ContRechargeMap[t] = cfg
		if this.ContRechargeTypeMap[t] == nil {
			this.ContRechargeTypeMap[t] = make([]*ContRechargeContRechargeCfg, 0)
		}
		this.ContRechargeTypeMap[t] = append(this.ContRechargeTypeMap[t], cfg)
	}
}

func (this *GameDb) genAncientBoss() {
	this.AncientBossMap = make(map[int]map[int]*AncientBossAncientBossCfg)
	for stageId, cfg := range this.AncientBossAncientBossCfgs {
		area := cfg.Area
		if this.AncientBossMap[area] == nil {
			this.AncientBossMap[area] = make(map[int]*AncientBossAncientBossCfg)
		}
		this.AncientBossMap[area][stageId] = cfg
	}
}

func (this *GameDb) genTitle() {
	this.TitleAutoActiveMap = make(map[int]*TitleTitleCfg)
	for id, cfg := range this.TitleTitleCfgs {
		if len(cfg.Item) < 1 {
			this.TitleAutoActiveMap[id] = cfg
		}
	}
}

func (this *GameDb) genKillMonster() {
	this.KillMonsterUniMap = make(map[int]*FirstBlooduniFirstBlooduniCfg)
	for _, cfg := range this.FirstBlooduniFirstBlooduniCfgs {
		this.KillMonsterUniMap[cfg.Stageid] = cfg
	}
	this.KillMonsterPerMap = make(map[int]*FirstBloodPerFirstBloodperCfg)
	for _, cfg := range this.FirstBloodPerFirstBloodperCfgs {
		this.KillMonsterPerMap[cfg.Stageid] = cfg
	}
	this.KillMonsterMilMaxLv = make(map[int]int)
	this.KillMonsterMilMap = make(map[int]map[int]*FirstBloodmilFirstBloodmilCfg)
	for _, cfg := range this.FirstBloodmilFirstBloodmilCfgs {
		t, lv := cfg.Type, cfg.Level
		if this.KillMonsterMilMaxLv[t] < lv {
			this.KillMonsterMilMaxLv[t] = lv
		}
		if this.KillMonsterMilMap[t] == nil {
			this.KillMonsterMilMap[t] = make(map[int]*FirstBloodmilFirstBloodmilCfg)
		}
		this.KillMonsterMilMap[t][lv] = cfg
	}
}

func (this *GameDb) genAncientTreasure() {
	this.AncientTreasureZhuLinMap = make(map[int]map[int]*TreasureArtTreasureArtCfg)
	this.AncientTreasureZhuLinMaxLvMap = make(map[int]int)
	this.AncientTreasureStarMap = make(map[int]map[int]*TreasureStarsTreasureStarsCfg)
	this.AncientTreasureMaxStarMap = make(map[int]int)
	for _, cfg := range this.TreasureArtTreasureArtCfgs {
		if this.AncientTreasureZhuLinMap[cfg.TreasureId] == nil {
			this.AncientTreasureZhuLinMap[cfg.TreasureId] = make(map[int]*TreasureArtTreasureArtCfg)
		}
		this.AncientTreasureZhuLinMap[cfg.TreasureId][cfg.Level] = cfg
		if cfg.Level > this.AncientTreasureZhuLinMaxLvMap[cfg.TreasureId] {
			this.AncientTreasureZhuLinMaxLvMap[cfg.TreasureId] = cfg.Level
		}
	}

	for _, cfg := range this.TreasureStarsTreasureStarsCfgs {
		if this.AncientTreasureStarMap[cfg.TreasureId] == nil {
			this.AncientTreasureStarMap[cfg.TreasureId] = make(map[int]*TreasureStarsTreasureStarsCfg)
		}
		this.AncientTreasureStarMap[cfg.TreasureId][cfg.Level] = cfg
		if cfg.Level > this.AncientTreasureMaxStarMap[cfg.TreasureId] {
			this.AncientTreasureMaxStarMap[cfg.TreasureId] = cfg.Level
		}
	}

	this.AncientTreasureJueXinMap = make(map[int]*TreasureAwakenTreasureAwakenCfg)
	for _, cfg := range this.TreasureAwakenTreasureAwakenCfgs {
		this.AncientTreasureJueXinMap[cfg.TruesureId] = cfg
	}

}

func (this *GameDb) genPet() {
	this.PetAppendageMap = make(map[int]map[int]*PetsAddPetsAddCfg)
	for _, cfg := range this.PetsAddPetsAddCfgs {
		pid, lv := cfg.PetsId, cfg.Level
		if this.PetAppendageMap[pid] == nil {
			this.PetAppendageMap[pid] = make(map[int]*PetsAddPetsAddCfg)
		}
		this.PetAppendageMap[pid][lv] = cfg
	}
}

func (this *GameDb) genHellBoss() {
	this.HellBossStageMap = make(map[int]*HellBossHellBossCfg)
	this.HellBossMap = make(map[int]map[int]*HellBossHellBossCfg)
	for _, cfg := range this.HellBossHellBossCfgs {
		stageId, floor := cfg.StageId, cfg.Floor
		if this.HellBossMap[floor] == nil {
			this.HellBossMap[floor] = make(map[int]*HellBossHellBossCfg)
		}
		this.HellBossMap[floor][stageId] = cfg
		this.HellBossStageMap[stageId] = cfg
	}
}

func (this *GameDb) genDaBao() {
	this.DaBaoEquipMap = make(map[int]map[int]*DaBaoEquipDaBaoEquipCfg)
	for _, cfg := range this.DaBaoEquipDaBaoEquipCfgs {
		t, lv := cfg.Type, cfg.Class
		if this.DaBaoEquipMap[t] == nil {
			this.DaBaoEquipMap[t] = make(map[int]*DaBaoEquipDaBaoEquipCfg)
		}
		this.DaBaoEquipMap[t][lv] = cfg
	}
	this.DaBaoMysteryMap = make(map[int]*DaBaoMysteryDaBaoMysteryCfg)
	for _, cfg := range this.DaBaoMysteryDaBaoMysteryCfgs {
		this.DaBaoMysteryMap[cfg.StageID] = cfg
	}
	this.DaBaoEquipAdditionMap = make(map[int]map[int]map[int]*DaBaoEquipAdditionDaBaoEquipAdditionCfg)
	for _, cfg := range this.DaBaoEquipAdditionDaBaoEquipAdditionCfgs {
		equipT, addT, lv := cfg.EquipType, cfg.AddType, cfg.EquipClass
		if this.DaBaoEquipAdditionMap[equipT] == nil {
			this.DaBaoEquipAdditionMap[equipT] = make(map[int]map[int]*DaBaoEquipAdditionDaBaoEquipAdditionCfg)
		}
		if this.DaBaoEquipAdditionMap[equipT][addT] == nil {
			this.DaBaoEquipAdditionMap[equipT][addT] = make(map[int]*DaBaoEquipAdditionDaBaoEquipAdditionCfg)
		}
		this.DaBaoEquipAdditionMap[equipT][addT][lv] = cfg
	}
}

func (this *GameDb) genLabel() {
	this.LabelTaskMap = make(map[int][]int)
	for id, cfg := range this.LabelTaskLabelTaskCfgs {
		if this.LabelTaskMap[id] == nil {
			this.LabelTaskMap[id] = make([]int, 0)
		}
		this.LabelTaskMap[id] = append(this.LabelTaskMap[id], cfg.Condition)
		this.LabelTaskMap[id] = append(this.LabelTaskMap[id], cfg.Value...)
	}
}

func (this *GameDb) genFirstDrop() {
	this.FirstDropByItemId = make(map[int]*FirstDropFirstDropCfg)
	this.FirstDropByTypes = make(map[int][]*FirstDropFirstDropCfg)
	for _, cfg := range this.FirstDropFirstDropCfgs {
		this.FirstDropByItemId[cfg.Item] = cfg
		if this.FirstDropByTypes[cfg.Type] == nil {
			this.FirstDropByTypes[cfg.Type] = make([]*FirstDropFirstDropCfg, 0)
		}
		this.FirstDropByTypes[cfg.Type] = append(this.FirstDropByTypes[cfg.Type], cfg)
	}
}
