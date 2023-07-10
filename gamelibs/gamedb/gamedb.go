package gamedb

import (
	"cqserver/gamelibs/errex"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"math/rand"
	"sort"
	"time"
)

func GetRealId(id, lv int) int {
	if lv >= constConstant.COMPUTE_TEN_THOUSAND {
		panic("Id lv异常了")
	}
	return id*constConstant.COMPUTE_TEN_THOUSAND + lv
}

func GetSkillIdAndLv(skillLvId int) (int, int) {
	skillId := skillLvId / 100
	skillLv := skillLvId % 100
	return skillId, skillLv
}

func GetSkillLvId(skillId, skillLv int) int {
	return skillId*100 + skillLv
}

// 客户端要的Json数据
var (
	onDemandData map[string]map[string]interface{}
)

var errSlice = make([]*errex.ErrorItem, 0)
var TextMap = make(map[string]string)

func initError(id int, message string) *errex.ErrorItem {
	e := errex.Create(id, message)
	errSlice = append(errSlice, e)
	return e
}

func codeTextSign(constName, text string) string {
	return text
}

type GameBaseCfg struct {
	Id    int    `col:"id" client:"id"`
	Type  string `col:"clinetType" client:"clinetType"`
	Name  string `col:"name" client:"name"`
	Value string `col:"value" client:"value"`
}

var gameDb *GameDb

func InitGameDb(gamedbPath string) {
	gameDb = &GameDb{GameDbBase: &GameDbBase{}}
	gameDb.gamedbPath = gamedbPath
	gameDb.FileModTime = make(map[string]int64)
	gameDb.InitConf = &InitConf{}
}

type GameDb struct {
	*GameDbBase
	gamedbPath                                 string
	BagSpaceAddCfgsMap                         map[int][]*BagSpaceAddCfg
	EquipRandPropsMap                          map[int][]*RandRandCfg
	EquipStrengthMap                           map[int]map[int]*StrengthenStrengthenCfg
	EquipStrengthMaxLvMap                      map[int]int
	DropMap                                    map[int][]*DropDropCfg
	DropSpecial                                map[int][]*DropSpecialDropSpecialCfg
	FabaoLvMap                                 map[int]map[int]*FabaolevelFabaolevelCfg
	FabaoMaxLvMap                              map[int]int
	FabaoSkillMap                              map[int]map[int]*FabaoSkillFabaoSkillCfg
	ArtifactMaxLvMap                           map[int]int
	PersonalBossMap                            map[int]*PersonalBossPersonalBossCfg
	WingMaxStarMap                             map[int]int
	AtlasStarMap                               map[int]map[int]*AtlasStarAtlasStarCfg
	AtlasMaxStarMap                            map[int]int
	AtlasUpgradeMap                            map[int]map[int]*AtlasUpgradeAtlasUpgradeCfg
	AtlasMaxUpgradeMap                         map[int]int
	AtlasWearMax                               int
	FieldBossMap                               map[int]map[int]*FieldBossFieldBossCfg
	WorldBossMap                               map[int]*WorldBossWorldBossCfg
	WorldRankMap                               map[int]*WorldRankWorldRankCfg
	worldRankMax                               int
	MaterialStageMap                           map[int]map[int]*MaterialStageMaterialStageCfg
	MaterialStageIdMap                         map[int]*MaterialStageMaterialStageCfg
	MaterialMaxLvMap                           map[int]int
	MaterialStageTypeMap                       map[int]map[int]*MaterialStageMaterialStageCfg
	MaleRoleNameMap                            []string
	FeMaleRoleNameMap                          []string
	VipBossMap                                 map[int]*VipBossVipBossCfg
	ExpStageMap                                map[int]*ExpStageExpStageCfg
	ExpStageLayerMap                           map[int]*ExpStageExpStageCfg
	ArenaMaxBuyNum                             int
	ArenaBuyNumMap                             map[int]*ArenaBuyArenaBuyCfg
	SkillLvMap                                 map[int]map[int]*SkillLevelSkillCfg
	SkillMaxLvMap                              map[int]int
	ComposeTypeMap                             map[int]*ComposeTypeComposeTypeCfg
	ComposeItemMap                             map[int]*ComposeSubComposeSubCfg
	AwakenMap                                  map[int]map[int]*AwakenAwakenCfg
	AwakenMaxLvMap                             map[int]int
	MainPrMap                                  map[int]map[int]*MainPrMainPrCfg
	BlessMaxLv                                 int
	DictateMap                                 map[int]map[int]*DictateEquipDictateEquipCfg
	DictateMaxLvMap                            map[int]int
	WingSpecialMap                             map[int]map[int]*WingSpecialWingSpecialCfg
	WingSpecialMaxLvMap                        map[int]int
	JewelMap                                   map[int]map[int]*JewelJewelCfg
	JewelMaxLvMap                              map[int]int
	JewelTypeLvMap                             map[int][]int
	JewelItemIdMap                             map[int]int
	FashionMap                                 map[int]map[int]*FashionFashionCfg
	FashionMaxLv                               map[int]int
	CumulativeSignMap                          map[int]*CumulationsignCumulationsignCfg
	InsideStarMap                              map[int]map[int]map[int]int
	InsideArtMap                               map[int]map[int]map[int]*InsideArtInsideArtCfg
	InsideMaxStar                              map[int]map[int]int
	InsideSkillMap                             map[int]map[int]*InsideSkillInsideSkillCfg
	InsideSkillMaxLv                           map[int]int
	HolyMaxLv                                  map[int]int
	HolyLvMap                                  map[int]map[int]*HolylevelHolylevelCfg
	HolySkillMap                               map[int]map[int]map[int]*HolySkillHolySkillCfg
	HolySkillMaxLv                             map[int]map[int]int
	PhantomSkill1Map                           map[int]map[int]map[int]*PhantomLevelPhantomLevelCfg
	PhantomSkill2Map                           map[int]map[int]map[int]*PhantomLevelPhantomLevelCfg
	MiningLvMap                                map[int]*MiningMiningCfg
	MiningMaxLv                                int
	WashMap                                    map[int]map[int]*WashWashCfg
	DarkPalaceBossMap                          map[int]map[int]*DarkPalaceBossDarkPalaceBossCfg
	DarkPalaceStageMap                         map[int]*DarkPalaceBossDarkPalaceBossCfg
	ExpPoolLvlAssembleCfg                      map[int]map[int]*ExpLevelLevelCfg          //经验池组装数据 k1:heroIndex k2:lvl
	ExpPoolLvlBuff                             map[int][]*WorldLevelBuffWorldLevelBuffCfg //经验池增益万分比表组装
	MagicCircleLvMap                           map[int]map[int]map[int]*MagicCircleLevelMagicCircleLevelCfg
	TalentMap                                  map[int]map[int]int
	DailyTaskDailytaskCfgsByType               map[int]*DailyTaskDailytaskCfg
	DailyRewardByTypes                         map[int][]*DailyRewardDailyRewardCfg //每日任务奖励数据组装
	MonthCardMap                               map[int]*MonthCardMonthCardCfg
	DayRankingRewardDayRankingRewardCfgsByType map[int][]*DayRankingRewardDayRankingRewardCfg
	GiftCodeMap                                map[string]*GiftCodeGiftCodeCfg
	AchievementAchievementCfgsByConditionType  map[int][]*AchievementAchievementCfg
	AchievementConditionState                  map[int]bool
	AchievementConditionIdState                map[int]bool
	AchievementConditionIdAndCondition         map[int]int
	LimitedGiftMap                             map[int]map[int]*LimitedGiftLimitedGiftCfg
	WorldLeaderSlice                           []*WorldLeaderConfCfg
	WorldLeaderRewardSlice                     map[int][]*WorldLeaderRewardWorldLeaderRewardCfg
	WarOrderMaxLv                              map[int]int
	WarOrderTaskMap                            map[int][]*WarOrderCycleTaskWarOrderCycleTaskCfg
	WarOrderWeekTaskMap                        map[int][]*WarOrderWeekTaskWarOrderWeekTaskCfg
	DrawCfgBySeasonAndType                     map[int]map[int]*DrawDrawCfg
	DrawShopDrawShopCfgByShop                  map[int]map[int]*DrawShopDrawShopCfg
	ElfRecoverMap                              map[int]map[int]*ElfRecoverElfRecoverCfg
	ElfSkillMap                                map[int]map[int]*ElfSkillElfGrowCfg
	ElfSkillMaxLv                              map[int]int
	CutTreasureMap                             map[int]*CutTreasureCutTreasureCfg
	TreasureCfgBySeasonAndType                 map[int]map[int]*XunlongXunlongCfg
	TreasureCfgBySeasonAndRound                map[int]map[int]*XunlongRoundsXunlongRoundsCfg
	HolyBeastCfgByTypeAndStar                  map[int]map[int]*HolyBeastHolyBeastCfg
	HolyBeastCfgByTypeMaxStar                  map[int]int
	FitHolyEquipMap                            map[int]map[int]map[int]*FitHolyEquipFitHolyEquipCfg
	RobotCfgSlice                              []*RobotRobotCfg
	towerMaxLv                                 int
	XunLongTypeTime                            map[int]int
	XunLongRoundSliceCfg                       map[int][]*XunlongRoundsXunlongRoundsCfg
	DayRankingMarkSliceCfg                     map[int][]*DayRankingMarkDayRankingMarkCfg
	ChuanShiSuitMap                            map[int]map[int]*ChuanShiSuitTypeChuanShiSuitTypeCfg
	SpendRebateMap                             map[int]*SpendrebatesSpendrebatesCfg
	ContRechargeMap                            map[int]*ContRechargeContRechargeCfg
	ContRechargeTypeMap                        map[int][]*ContRechargeContRechargeCfg
	AncientBossMap                             map[int]map[int]*AncientBossAncientBossCfg
	TitleAutoActiveMap                         map[int]*TitleTitleCfg
	KillMonsterUniMap                          map[int]*FirstBlooduniFirstBlooduniCfg
	KillMonsterPerMap                          map[int]*FirstBloodPerFirstBloodperCfg
	KillMonsterMilMaxLv                        map[int]int
	KillMonsterMilMap                          map[int]map[int]*FirstBloodmilFirstBloodmilCfg
	AncientTreasureZhuLinMap                   map[int]map[int]*TreasureArtTreasureArtCfg
	AncientTreasureZhuLinMaxLvMap              map[int]int
	AncientTreasureStarMap                     map[int]map[int]*TreasureStarsTreasureStarsCfg
	AncientTreasureMaxStarMap                  map[int]int
	AncientTreasureJueXinMap                   map[int]*TreasureAwakenTreasureAwakenCfg
	ChuanShiStrengthenMap                      map[int]map[int]*ChuanShiStrengthenChuanShiStrengthenCfg
	ChuanShiStrengthenSuitMap                  map[int]map[int]*ChuanShiStrengthenLinkChuanShiStrengthenLinkCfg
	PetAppendageMap                            map[int]map[int]*PetsAddPetsAddCfg
	HellBossMap                                map[int]map[int]*HellBossHellBossCfg
	HellBossStageMap                           map[int]*HellBossHellBossCfg
	DaBaoEquipMap                              map[int]map[int]*DaBaoEquipDaBaoEquipCfg
	DaBaoMysteryMap                            map[int]*DaBaoMysteryDaBaoMysteryCfg
	DaBaoEquipAdditionMap                      map[int]map[int]map[int]*DaBaoEquipAdditionDaBaoEquipAdditionCfg
	ComposeEquipMap                            map[int]*ComposeEquipSubComposeEquipSubCfg
	LabelTaskMap                               map[int][]int
	FirstDropByItemId                          map[int]*FirstDropFirstDropCfg
	FirstDropByTypes                           map[int][]*FirstDropFirstDropCfg
}

func GetDb() *GameDb {
	return gameDb
}

func GetConf() *InitConf {
	return gameDb.InitConf
}

func GetBagInitNum() int {
	return gameDb.BagSpaceAddCfgsMap[constBag.BAG_ADD_TYPE_INIT][0].BagNumber
}

func GetBagAddCfgLen(types int) int {
	return len(gameDb.BagSpaceAddCfgsMap[types])
}

func GetWareHouseBagInitNum() int {
	return gameDb.BagSpaceAddCfgsMap[constBag.WAREHOUSE_BAG_ADD_TYPE_INIT][0].BagNumber
}

func GetBagSpaceAddCost(index int) (int, int, int) {
	return gameDb.BagSpaceAddCfgsMap[constBag.BAG_ADD_TYPE_ITEM][index].Cost,
		gameDb.BagSpaceAddCfgsMap[constBag.BAG_ADD_TYPE_ITEM][index].Value,
		gameDb.BagSpaceAddCfgsMap[constBag.BAG_ADD_TYPE_ITEM][index].BagNumber
}

func GetWarehouseBagSpaceAddCost(index int) (int, int, int) {
	return gameDb.BagSpaceAddCfgsMap[constBag.WAREHOUSE_BAG_ADD_TYPE_ITEM][index].Cost,
		gameDb.BagSpaceAddCfgsMap[constBag.WAREHOUSE_BAG_ADD_TYPE_ITEM][index].Value,
		gameDb.BagSpaceAddCfgsMap[constBag.WAREHOUSE_BAG_ADD_TYPE_ITEM][index].BagNumber
}

func GetVipMaxLv() int {

	return len(gameDb.VipLvlCfgs) - 1
}

//根据vip获取增加对应vip配置,增加格子数
func GetBagSapceAddByType(addType, getValue, nowValue int) (int, int) {

	addNum := 0
	var maxConf *BagSpaceAddCfg
	for _, v := range gameDb.BagSpaceAddCfgsMap[addType] {
		if v.Value > getValue && v.Value <= nowValue {
			addNum += v.BagNumber
			if maxConf == nil || v.Value > maxConf.Value {
				maxConf = v
			}
		}
	}
	if maxConf == nil {
		return 0, 0
	}
	return maxConf.Value, addNum
}

func RandEquipProp(equipId int) ([]*model.EquipRandProp, error) {

	conf := GetEquipEquipCfg(equipId)
	if conf == nil {
		return nil, ERRSETTINGNOTFOUND
	}

	randProp := make([]*model.EquipRandProp, len(conf.Rand_group))
	for k, v := range conf.Rand_group {

		randPropSetting := gameDb.EquipRandPropsMap[v]
		if randPropSetting == nil {
			return nil, ERRSETTINGNOTFOUND
		}

		weightSlice := make([]int, len(randPropSetting))
		for kk, vv := range randPropSetting {
			weightSlice[kk] = vv.Weight
		}
		randIndex := common.RandWeightByIntSlice(weightSlice)

		weightSlice = make([]int, len(randPropSetting[randIndex].Attribute))
		for kk, vv := range randPropSetting[randIndex].Attribute {
			weightSlice[kk] = vv[1]
		}

		randPropIndex := common.RandWeightByIntSlice(weightSlice)

		randProp[k] = &model.EquipRandProp{
			PropId: randPropSetting[randIndex].Attribute[randPropIndex][0],
			Color:  randPropSetting[randIndex].Quality,
			Value:  common.RandNum(randPropSetting[randIndex].Attribute[randPropIndex][2], randPropSetting[randIndex].Attribute[randPropIndex][3]),
		}
	}
	return randProp, nil
}

//怪物id， 用户总拾取次数， 根据开服天数获取的次数,  用户类型拾取次数
func GetMonsterDrop(monsterId, userPickNum, pickMax int, pickInfos map[int]int, isFirst bool, stageType int) ([]*ItemInfo, map[int]int, error) {
	monsterConf := GetMonsterMonsterCfg(monsterId)
	dropItems := make(ItemInfos, 0)
	//如果是首杀，并且配置了首爆，掉落首爆,否则掉落普通掉落和特殊掉落
	if isFirst && monsterConf.FirstDropId > 0 {
		firstDropItems, err := GetDropItems(monsterConf.FirstDropId)
		if err != nil {
			logger.Error("首爆奖励获取异常：%v", monsterConf.FirstDropId, err)
		}
		if len(firstDropItems) > 0 {
			dropItems = append(dropItems, firstDropItems...)
		}
	} else {
		normalDrop, err := GetDropItems(monsterConf.DropId)
		if err != nil {
			return nil, nil, err
		}
		if len(normalDrop) <= 0 {
			logger.Error("申请怪物掉落，随机物品为空, monsterId:%v", monsterId)
			return nil, nil, ERRUNKNOW
		}
		dropItems = append(dropItems, normalDrop...)
		if len(monsterConf.DropSpecial) > 1 {
			userDropTimes := pickInfos[monsterConf.DropSpecial[0]] + 1
			specialItem, err := GetDropSpecialItems(monsterId, userDropTimes, userPickNum, pickMax)
			if err != nil {
				return nil, nil, err
			}
			if specialItem != nil {
				dropItems = append(dropItems, specialItem)
				pickInfos[monsterConf.DropSpecial[0]] += 1
			}
		}
		if stageType == constFight.FIGHT_TYPE_PERSON_BOSS && isFirst {
			copySlice := make([]*ItemInfo, len(dropItems))
			copy(copySlice, dropItems)
			dropItems = append(dropItems, copySlice...)
		}
	}
	if len(dropItems) <= 0 {
		logger.Error("申请怪物掉落，随机物品为空, monsterId:%v", monsterId)
		return nil, nil, ERRUNKNOW
	}
	return dropItems, pickInfos, nil
}

//普通掉落
func GetDropItems(dropId int) ([]*ItemInfo, error) {
	dropConfs := gameDb.DropMap[dropId]
	if dropConfs == nil {
		return nil, ERRSETTINGNOTFOUND
	}

	dropSetting := make([]IntSlice2, 0)
	for _, v := range dropConfs {

		isOk := common.RandByTenShousand(v.Rate)
		if !isOk {
			continue
		}

		dropNum := common.RandNum(v.Number[0], v.Number[1])
		for i := 0; i < dropNum; i++ {
			dropSetting = append(dropSetting, v.Drop)
		}
	}

	dropItems := make([]*ItemInfo, 0)
	var weightSlice []int
	for _, v := range dropSetting {
		if len(v) < 1 {
			continue
		}
		weightSlice = make([]int, len(v))
		for kk, vv := range v {
			weightSlice[kk] = vv[3]
		}
		dropIndex := common.RandWeightByIntSlice(weightSlice)

		//dropItems[v[dropIndex][0]] += common.RandNum(v[dropIndex][1], v[dropIndex][2])
		dropItems = append(dropItems, &ItemInfo{
			ItemId: v[dropIndex][0],
			Count:  common.RandNum(v[dropIndex][1], v[dropIndex][2]),
		})

	}
	return dropItems, nil
}

//特殊掉落
func GetDropSpecialItems(monsterId, userDropTimes, userPickNum, pickMax int) (items *ItemInfo, err error) {
	monsterCfg := GetMonsterMonsterCfg(monsterId)
	if len(monsterCfg.DropSpecial) <= 1 {
		return
	}
	if !common.RandByTenShousand(monsterCfg.DropSpecial[1]) {
		return
	}

	dropCfs := GetDropSpecialsByType(monsterCfg.DropSpecial[0])
	if dropCfs == nil {
		logger.Error("怪物配置特殊掉落，掉落配置未找到,怪物：%v，掉落：%v", monsterCfg.Monsterid, monsterCfg.DropSpecial)
		return
	}

	var dropSetting *DropSpecialDropSpecialCfg
	for _, v := range dropCfs {
		if v.Num[0] <= userDropTimes && v.Num[1] >= userDropTimes {
			dropSetting = v
			break
		}
	}
	if dropSetting == nil {
		return
	}

	weightSlice := make(map[int]int, 0)
	for kk, vv := range dropSetting.DropSpecial {
		itemConf := GetItemBaseCfg(vv[0])
		if itemConf == nil {
			logger.Error("特殊掉落 未配置itemId:%v", vv[0])
			continue
		}
		if userPickNum+(itemConf.EffectVal*vv[1]) > pickMax {
			continue
		}
		weightSlice[kk] = vv[2]
	}
	if len(weightSlice) <= 0 {
		return
	}
	dropIndex := common.RandWeightByMap(weightSlice)
	if dropIndex == -1 {
		return
	}
	logger.Debug("特殊掉落,怪物：%v，掉落id：%v,掉落：%v", monsterId, dropSetting.Id, dropSetting.DropSpecial[dropIndex])
	items = &ItemInfo{
		ItemId: dropSetting.DropSpecial[dropIndex][0],
		Count:  dropSetting.DropSpecial[dropIndex][1],
	}
	return
}

func GetDropSpecialsByType(t int) []*DropSpecialDropSpecialCfg {
	return gameDb.DropSpecial[t]
}

func GetRedPacketDropMax(day int) int {

	for _, v := range gameDb.RedDayMaxRedDayMaxCfgs {
		if v.Day[0] <= day && v.Day[1] >= day {
			return v.Max
		}
	}
	return 0
}

func GetEquipStrengthConfByLvAndPos(pos, lv int) *StrengthenStrengthenCfg {
	return gameDb.EquipStrengthMap[pos][lv]
}
func GetEquipStrengthenLinkCfgs() map[int]*StrengthenlinkStrengthenCfg {
	return gameDb.StrengthenlinkStrengthenCfgs
}

func GetFabaoById(id int) *FabaoFabaoCfg {
	return gameDb.FabaoFabaoCfgs[id]
}

func GetFabaoLvByIdAndLv(id, lv int) *FabaolevelFabaolevelCfg {
	return gameDb.FabaoLvMap[id][lv]
}

func GetFabaoSkillByIdAndSkillId(id, skillId int) *FabaoSkillFabaoSkillCfg {
	return gameDb.FabaoSkillMap[id][skillId]
}

func GetMaxValById(val, model int) int {
	var backVal int
	switch model {
	case constMax.MAX_FABAO_LEVEL:
		backVal = gameDb.FabaoMaxLvMap[val]
	case constMax.MAX_ARTIFACT_LEVEL:
		backVal = gameDb.ArtifactMaxLvMap[val]
	case constMax.MAX_ATLAS_STAR:
		backVal = gameDb.AtlasMaxStarMap[val]
	case constMax.MAX_ATLAS_UPGRADE:
		backVal = gameDb.AtlasMaxUpgradeMap[val]
	case constMax.MAX_WING_STAR:
		backVal = gameDb.WingMaxStarMap[val]
	case constMax.MAX_ARENA_BUY_NUM:
		backVal = gameDb.ArenaMaxBuyNum
	case constMax.MAX_SKILL_LEVEL:
		backVal = gameDb.SkillMaxLvMap[val]
	case constMax.MAX_AWAKEN_LEVEL:
		backVal = gameDb.AwakenMaxLvMap[val]
	case constMax.MAX_BLESS_LEVEL:
		backVal = gameDb.BlessMaxLv
	case constMax.MAX_MATERIAL_LEVEL:
		backVal = gameDb.MaterialMaxLvMap[val]
	case constMax.MAX_DICTATE_LEVEL:
		backVal = gameDb.DictateMaxLvMap[val]
	case constMax.MAX_WING_SPECIAL_LEVEL:
		backVal = gameDb.WingSpecialMaxLvMap[val]
	case constMax.MAX_JEWEL_LEVEL:
		backVal = gameDb.JewelMaxLvMap[val]
	case constMax.MAX_FASHION_LEVEL:
		backVal = gameDb.FashionMaxLv[val]
	case constMax.MAX_INSIDE_SKILL_LEVEL:
		backVal = gameDb.InsideSkillMaxLv[val]
	case constMax.MAX_HOLY_LEVEL:
		backVal = gameDb.HolyMaxLv[val]
	case constMax.MAX_ATLAS_WEAR:
		backVal = gameDb.AtlasWearMax
	case constMax.MAX_MINING_LEVEL:
		backVal = gameDb.MiningMaxLv
	case constMax.MAX_WARORDER_LEVEL:
		backVal = gameDb.WarOrderMaxLv[val]
	case constMax.MAX_ELF_SKILL_LEVEL:
		backVal = gameDb.ElfSkillMaxLv[val]
	case constMax.MAX_TOWER_LEVEL:
		backVal = gameDb.towerMaxLv
	case constMax.MAX_KILL_MONSTER_MIL_LEVEL:
		backVal = gameDb.KillMonsterMilMaxLv[val]
	case constMax.MAX_EQUIP_STRENGTHEN_LEVEL:
		backVal = gameDb.EquipStrengthMaxLvMap[val]
	}
	return backVal
}

func GetSingleBossBaseCfgs() map[int]*PersonalBossPersonalBossCfg {
	return gameDb.PersonalBossMap
}

func GetSingleBossByStage(stageId int) *PersonalBossPersonalBossCfg {
	if val, ok := gameDb.PersonalBossMap[stageId]; ok {
		return val
	}
	return nil
}

func GetWingCfg(id int) *WingNewWingNewCfg {
	return gameDb.WingNewWingNewCfgs[id]
}
func GetWingSpecialByOrderAndLv(specialT, lv int) *WingSpecialWingSpecialCfg {
	return gameDb.WingSpecialMap[specialT][lv]
}
func GetWingSpecialCfgs() map[int]*WingSpecialWingSpecialCfg {
	return gameDb.WingSpecialWingSpecialCfgs
}

func GetReinCfg(id int) *ReinReinCfg {
	return gameDb.ReinReinCfgs[id]
}

func GetReinCostCfg(id int) *ReinCostReinCostCfg {
	return gameDb.ReinCostReinCostCfgs[id]
}

func GetAllStageCfg() map[int]*StageStageCfg {
	return gameDb.StageStageCfgs
}

func GetWorldBossConfByStageId(stageId int) *WorldBossWorldBossCfg {

	for _, v := range gameDb.WorldBossWorldBossCfgs {
		if v.Stageid == stageId {
			return v
		}
	}
	return nil
}

func GetAtlasStar(id, star int) *AtlasStarAtlasStarCfg {
	return gameDb.AtlasStarMap[id][star]
}
func GetAtlasGather(id int) *AtlasGatherAtlasGatherCfg {
	return gameDb.AtlasGatherAtlasGatherCfgs[id]
}
func GetAtlasUpgrade(id, star int) *AtlasUpgradeAtlasUpgradeCfg {
	return gameDb.AtlasUpgradeMap[id][star]
}
func GetAtlasCfgs() map[int]*AtlasAtlasCfg {
	return gameDb.AtlasAtlasCfgs
}
func GetAtlasGatherCfgs() map[int]*AtlasGatherAtlasGatherCfg {
	return gameDb.AtlasGatherAtlasGatherCfgs
}

func GetFieldBoss(area, stageId int) *FieldBossFieldBossCfg {
	return gameDb.FieldBossMap[area][stageId]
}

func GetFieldBossByStageId(stageId int) *FieldBossFieldBossCfg {
	for _, v := range gameDb.FieldBossFieldBossCfgs {
		if v.StageId == stageId {
			return v
		}
	}
	return nil
}

func GetFieldBossByArea(area int) map[int]*FieldBossFieldBossCfg {
	return gameDb.FieldBossMap[area]
}

func GetWorldBosses() map[int]*WorldBossWorldBossCfg {
	return gameDb.WorldBossMap
}

func GetWorldBossByStageId(stageId int) *WorldBossWorldBossCfg {
	return gameDb.WorldBossMap[stageId]
}

func GetWorldRank(rank int) *WorldRankWorldRankCfg {
	if val, ok := gameDb.WorldRankMap[rank]; ok {
		return val
	}
	return gameDb.WorldRankMap[gameDb.worldRankMax]
}

func GetMaterialStageCfgsByType(t int) map[int]*MaterialStageMaterialStageCfg {
	return gameDb.MaterialStageTypeMap[t]
}
func GetMaterialByStageId(stageId int) *MaterialStageMaterialStageCfg {
	return gameDb.MaterialStageIdMap[stageId]
}
func GetMaterialByTypeAndLv(t, lv int) *MaterialStageMaterialStageCfg {
	return gameDb.MaterialStageMap[t][lv]
}

func GetRoleName(sex int) []string {
	switch sex {
	case pb.SEX_MALE:
		return gameDb.MaleRoleNameMap
	case pb.SEX_FEMALE:
		return gameDb.FeMaleRoleNameMap
	default:
		slice := gameDb.MaleRoleNameMap
		slice = append(slice, gameDb.FeMaleRoleNameMap...)
		return slice
	}
}

func GetVipBossByStageId(stageId int) *VipBossVipBossCfg {
	return gameDb.VipBossMap[stageId]
}

func GetExpStageByStageId(stageId int) *ExpStageExpStageCfg {
	if cfg, ok := gameDb.ExpStageMap[stageId]; ok {
		return cfg
	} else {
		return nil
	}
}
func GetExpStageByLayer(layer int) *ExpStageExpStageCfg {
	return gameDb.ExpStageLayerMap[layer]
}
func GetExpStageLayers() map[int]*ExpStageExpStageCfg {
	return gameDb.ExpStageLayerMap
}

func GetAttIntervalByAttSpeed(atkSpeed int) int {

	var maxAspd *AspdAspdCfg
	var minAspd *AspdAspdCfg
	for _, v := range gameDb.AspdAspdCfgs {
		if v.AspdMin <= atkSpeed && atkSpeed < v.AspdMax {
			return v.Time
		}
		if maxAspd == nil || v.AspdMax > maxAspd.AspdMax {
			maxAspd = v
		}
		if minAspd == nil || v.AspdMin < minAspd.AspdMin {
			minAspd = v
		}
	}
	if atkSpeed >= maxAspd.AspdMax {
		return maxAspd.Time
	}
	return minAspd.Time
}

func GetArenaRank() map[int]*ArenaRankArenaRankCfg {
	return gameDb.ArenaRankArenaRankCfgs
}

func GetArenaMatch() map[int]*ArenaMatchArenaMatchCfg {
	return gameDb.ArenaMatchArenaMatchCfgs
}
func GetLenCompetitveCompetitveCfg() int {
	return len(gameDb.CompetitveCompetitveCfgs)
}

func GetCompetitveCfgByScore(score int) (bool, ItemInfos, *CompetitveCompetitveCfg) {
	reward := ItemInfos{}
	state := false
	cfg1 := &CompetitveCompetitveCfg{}
	for i := 1; i <= len(gameDb.CompetitveCompetitveCfgs); i++ {
		cfg := gameDb.CompetitveCompetitveCfgs[i]
		if cfg == nil {
			continue
		}
		if score >= cfg.Mark {
			reward = cfg.RewardDay
			state = true
			cfg1 = cfg
		}
	}

	return state, reward, cfg1
}

func GetCompetitveMatchRuleSectionCfg(score int) ([]int, int) {

	//section:当前分数段区间
	//match: 当前段位配置表id
	section := make([]int, 0)
	competitveCfgsLen := len(gameDb.CompetitveCompetitveCfgs)
	matchId := 0
	for i := 1; i <= competitveCfgsLen; i++ {
		if score >= gameDb.CompetitveCompetitveCfgs[competitveCfgsLen].Mark {
			section = append(section, gameDb.CompetitveCompetitveCfgs[competitveCfgsLen].Mark, 99999999)
			matchId = competitveCfgsLen
			return section, matchId
		}
		if score == 0 {
			if i == 1 {
				section = append(section, 0, gameDb.CompetitveCompetitveCfgs[i+1].Mark)
				matchId = 1
				return section, matchId
			}
		}

		cfg := gameDb.CompetitveCompetitveCfgs[i]
		if score < cfg.Mark {
			if i == 1 {
				section = append(section, 0, gameDb.CompetitveCompetitveCfgs[i+1].Mark)
				matchId = 1
				return section, matchId
			}
			section = append(section, gameDb.CompetitveCompetitveCfgs[i-1].Mark, cfg.Mark)
			matchId = i - 1
			return section, matchId

		}
	}
	return section, matchId
}

func GetSkills() map[int]*SkillSkillCfg {
	return gameDb.SkillSkillCfgs
}

func GetSkillLvConf(skillId, lv int) *SkillLevelSkillCfg {
	return gameDb.SkillLvMap[skillId][lv]
}

func GetComposeTypeBySubTab(subTab int) *ComposeTypeComposeTypeCfg {
	return gameDb.ComposeTypeMap[subTab]
}

func GetAwakenByTypeAndLv(t, lv int) *AwakenAwakenCfg {
	return gameDb.AwakenMap[t][lv]
}

func GetMainPrConf(t, body, job int) int {
	mainPrConf := gameDb.MainPrMap[t][body]
	mainPr := 0
	switch job {
	case pb.JOB_ZHANSHI:
		mainPr = mainPrConf.MainPrP
	case pb.JOB_FASHI:
		mainPr = mainPrConf.MainPrM
	case pb.JOB_DAOSHI:
		mainPr = mainPrConf.MainPrT
	}
	return mainPr
}

func GetItemSourceByStageId(stageId int) string {

	monsterName := ""
	stageConf := GetStageStageCfg(stageId)

	if stageConf != nil {

		if len(stageConf.Monster_group) == 1 {

			monsterGroupConf := GetMonstergroupMonstergroupCfg(stageConf.Monster_group[0][0])

			if monsterGroupConf != nil && len(monsterGroupConf.Monsterid) == 1 {
				monsterConf := GetMonsterMonsterCfg(monsterGroupConf.Monsterid[0][0])
				if monsterConf != nil {
					monsterName = monsterConf.Name
				}
			}
		}
	}

	return monsterName
}

func GetDictateByBodyAndGrade(body, grade int) *DictateEquipDictateEquipCfg {
	return gameDb.DictateMap[body][grade]
}

func GetPanaceaCfgs() map[int]*PanaceaPanaceaCfg {
	return gameDb.PanaceaPanaceaCfgs
}

func GetJewelSuitCfgs() map[int]*JewelSuitJewelSuitCfg {
	return gameDb.JewelSuitJewelSuitCfgs
}

func GetJewelByKindAndLv(kind, lv int) *JewelJewelCfg {
	return gameDb.JewelMap[kind][lv]
}

func GetJewelTypeLvCfg(kind int) []int {
	return gameDb.JewelTypeLvMap[kind]
}

func GetJewelItemIdMap() map[int]int {
	return gameDb.JewelItemIdMap
}

func GetFashionConf(fashionId, lv int) *FashionFashionCfg {
	return gameDb.FashionMap[fashionId][lv]
}

func GetCumulativeByDay(day int) *CumulationsignCumulationsignCfg {
	return gameDb.CumulativeSignMap[day]
}

func GetInsideStarWeightMap(g, o int) map[int]int {
	return gameDb.InsideStarMap[g][o]
}

func GetInsideMaxStar(grade, order int) int {
	return gameDb.InsideMaxStar[grade][order]
}

func GetInsideByGradeAndOrder(grade, order, star int) *InsideArtInsideArtCfg {
	return gameDb.InsideArtMap[grade][order][star]
}

func GetInsideSkillBySidAndLv(sid, lv int) *InsideSkillInsideSkillCfg {
	return gameDb.InsideSkillMap[sid][lv]
}

func GetHolyLvByHidAndLv(hid, lv int) *HolylevelHolylevelCfg {
	return gameDb.HolyLvMap[hid][lv]
}

func GetHolySkillByHidAndLv(hid, hlv, lv int) *HolySkillHolySkillCfg {
	return gameDb.HolySkillMap[hid][hlv][lv]
}

func GetHolySkillMaxLv(hid, hlv int) int {
	return gameDb.HolySkillMaxLv[hid][hlv]
}

func GetPhantomSkill1(phantom, pos, lv int) *PhantomLevelPhantomLevelCfg {
	return gameDb.PhantomSkill1Map[phantom][pos][lv]
}

func GetPhantomSkill2(phantom, pos, lv int) *PhantomLevelPhantomLevelCfg {
	return gameDb.PhantomSkill2Map[phantom][pos][lv]
}

func GetMiningLvCfg(lv int) *MiningMiningCfg {
	return gameDb.MiningLvMap[lv]
}

func GetCompetitveSeasonEndReward(rank int) (ItemInfos, int) {

	for i := 1; i <= len(gameDb.CompetitveRewardRankRewardCfgs); i++ {
		cfg := gameDb.CompetitveRewardRankRewardCfgs[i]
		if len(cfg.Rank) < 2 {
			continue
		}
		if rank >= cfg.Rank[0] && rank <= cfg.Rank[1] {
			return cfg.RewardWeek, cfg.Id
		}
	}
	return gameDb.CompetitveRewardRankRewardCfgs[len(gameDb.CompetitveRewardRankRewardCfgs)].RewardWeek, len(gameDb.CompetitveRewardRankRewardCfgs)
}

func GetWashByPosAndGrade(pos, grade int) *WashWashCfg {
	return gameDb.WashMap[pos][grade]
}

func GetDarkPalaceBossCfg(floor, stageId int) *DarkPalaceBossDarkPalaceBossCfg {
	return gameDb.DarkPalaceBossMap[floor][stageId]
}

func GetDarkPalaceBossByFloor(floor int) map[int]*DarkPalaceBossDarkPalaceBossCfg {
	return gameDb.DarkPalaceBossMap[floor]
}

func GetDarkPalaceStageCfg(stageId int) *DarkPalaceBossDarkPalaceBossCfg {
	return gameDb.DarkPalaceStageMap[stageId]
}

//获取经验池升级限制
func GetExpPoolCfg(heroIndex, lvl int) *ExpLevelLevelCfg {
	return gameDb.ExpPoolLvlAssembleCfg[heroIndex][lvl]
}

func GetExpPoolCfgByHeroIndex(heroIndex int) map[int]*ExpLevelLevelCfg {
	return gameDb.ExpPoolLvlAssembleCfg[heroIndex]
}

//获取经验池世界等级
func GetExpWorldLevel(openDay int) int {
	for i := 1; i <= len(gameDb.WorldLevelWorldLevelCfgs); i++ {
		cfg := gameDb.WorldLevelWorldLevelCfgs[i]
		if openDay >= cfg.MinDay && openDay <= cfg.MaxDay {
			return cfg.WorldLevel
		}
	}
	return -1
}

func GetExpPoolLimitMaxLen() int {
	return len(gameDb.ExpPoolExpPoolCfgs)
}

func GetExpPoolWorldLvBuff(types, level int) float64 {

	if types == -1 {
		return -1
	}

	for _, v := range gameDb.ExpPoolLvlBuff[types] {

		if level >= v.LowLevel && level <= v.TopLevel {
			return float64(v.Effect)
		}
	}
	return -1
}

func GetMagicCircleLvCfg(id, grade, lv int) *MagicCircleLevelMagicCircleLevelCfg {
	return gameDb.MagicCircleLvMap[id][grade][lv]
}

//根据战力区间 筛选出假人  返回0代表没有找到假人
func GetRandomRobotIdCfg(myCombat int, screen map[int]bool) int {
	robotIds := make([]int, 0)
	for i := 1; i < len(gameDb.RobotRobotCfgs); i++ {
		if screen[i] {
			continue
		}
		cfg := gameDb.RobotRobotCfgs[i]
		robotIds = append(robotIds, cfg.Id)
		if cfg != nil {
			if len(cfg.Fighting) >= 2 {
				if myCombat >= cfg.Fighting[0] && myCombat <= cfg.Fighting[1] {
					return cfg.Id
				}
			}
		}
	}
	if len(robotIds) > 0 {
		return robotIds[0]
	}

	return 0
}
func GetRobotCfgs() map[int]*RobotRobotCfg {
	return gameDb.RobotRobotCfgs
}

func GetTalentGetCfgs() map[int]*TalentGetTalentGetCfg {
	return gameDb.TalentGetTalentGetCfgs
}

func GetTalentByJobAndId(job, id int) int {
	if id, ok := gameDb.TalentMap[job][id]; ok {
		return id
	}
	return 0
}

func GetGuildMemberLimit(contribution int) (int, int) {
	lv := 1
	limit := gameDb.GuildLevelGuildLevelCfgs[1].NumberLimit
	for i := 1; i < len(gameDb.GuildLevelGuildLevelCfgs); i++ {
		if contribution >= gameDb.GuildLevelGuildLevelCfgs[len(gameDb.GuildLevelGuildLevelCfgs)-1].Value {
			lv = len(gameDb.GuildLevelGuildLevelCfgs)
			limit = gameDb.GuildLevelGuildLevelCfgs[len(gameDb.GuildLevelGuildLevelCfgs)].NumberLimit
			return lv, limit
		} else {
			if contribution < gameDb.GuildLevelGuildLevelCfgs[i].Value {
				lv = i
				limit = gameDb.GuildLevelGuildLevelCfgs[i].NumberLimit
				return lv, limit
			}
		}
	}

	return lv, limit
}

func GetGuildBonfireGuildBonfire() map[int]*GuildBonfireGuildBonfireCfg {
	return gameDb.GuildBonfireGuildBonfireCfgs
}

func GetPaodianConfByStageId(stageId int) *PaoDianRewardPaoDianRewardCfg {
	for _, v := range gameDb.PaoDianRewardPaoDianRewardCfgs {
		for _, vv := range v.Stage {
			if vv == stageId {
				return v
			}
		}
	}
	return nil
}

func GetDailyActivityCfgs() map[int]*DailyActivityDailyActivityCfg {
	return gameDb.DailyActivityDailyActivityCfgs
}

func GetShabakePerRewardByRank(rank int) ItemInfos {
	for i := 1; i < len(gameDb.ShabakeRewardperShabakeRewardperCfgs); i++ {
		if len(gameDb.ShabakeRewardperShabakeRewardperCfgs[i].Rank) >= 2 {
			if rank >= gameDb.ShabakeRewardperShabakeRewardperCfgs[i].Rank[0] && rank <= gameDb.ShabakeRewardperShabakeRewardperCfgs[i].Rank[1] {
				return gameDb.ShabakeRewardperShabakeRewardperCfgs[i].Reward
			}
		}
	}
	return nil
}

func GetShabakeUniRewardByRank(rank int) ItemInfos {
	for i := 1; i < len(gameDb.ShabakeRewarduniShabakeRewarduniCfgs); i++ {
		if len(gameDb.ShabakeRewarduniShabakeRewarduniCfgs[i].Rank) >= 2 {
			if rank >= gameDb.ShabakeRewarduniShabakeRewarduniCfgs[i].Rank[0] && rank <= gameDb.ShabakeRewarduniShabakeRewarduniCfgs[i].Rank[1] {
				return gameDb.ShabakeRewarduniShabakeRewarduniCfgs[i].Reward
			}
		}
	}
	return nil
}

func GetDailyTaskDailyTaskCfgByType(types int) *DailyTaskDailytaskCfg {
	if gameDb.DailyTaskDailytaskCfgsByType[types] == nil {
		if gameDb.DailyTaskDailytaskCfgsByType != nil {
			for types, data := range gameDb.DailyTaskDailytaskCfgsByType {
				logger.Info("DailyTaskDailytaskCfgsByType types:%v  data:%v", types, data)
			}
		} else {
			logger.Error("gameDb.DailyTaskDailytaskCfgsByType == nil")
		}
	}
	return gameDb.DailyTaskDailytaskCfgsByType[types]
}

func GetDailyTaskRewardCfgByType(types, exp int) []*DailyRewardDailyRewardCfg {

	canRewardId := make([]*DailyRewardDailyRewardCfg, 0)
	cfg := gameDb.DailyRewardByTypes[types]
	for _, v := range cfg {
		if exp >= v.Active {
			canRewardId = append(canRewardId, v)
		}
	}
	return canRewardId
}

func GetMonthCardByType(t int) *MonthCardMonthCardCfg {
	return gameDb.MonthCardMap[t]
}

func GetDayRankRewardCfgByType(rankType, rankScore, lastPeopleRank, lastCfgId int) (int, int) {

	myRank := lastPeopleRank + 1
	allCfg := gameDb.DayRankingRewardDayRankingRewardCfgsByType[rankType]
	for i := lastCfgId; i < len(allCfg); i++ {
		if len(allCfg[i].Ranking) <= 0 {
			continue
		}
		if len(allCfg[i].Ranking) == 1 {
			if myRank == allCfg[i].Ranking[0] {
				if rankScore >= allCfg[i].Least {
					if myRank-lastPeopleRank > 1 && lastPeopleRank > 1 {
						myRank = allCfg[i].Ranking[0]
					}
					return myRank, i
				} else {
					myRank += 1
				}
			}
		}

		if len(allCfg[i].Ranking) >= 2 {
			if myRank >= allCfg[i].Ranking[0] && myRank <= allCfg[i].Ranking[1] {
				if rankScore >= allCfg[i].Least {
					if myRank-lastPeopleRank > 1 && lastPeopleRank > 1 {
						myRank = allCfg[i].Ranking[0]
					}
					return myRank, i
				} else {
					myRank += allCfg[i].Ranking[1] - allCfg[i].Ranking[0] + 1
				}
			}
		}
	}

	return myRank, len(allCfg) - 1
}

func GetDayRankRewardCfgByRank(rankId, rank int) *DayRankingRewardDayRankingRewardCfg {
	cfgs := gameDb.DayRankingRewardDayRankingRewardCfgsByType[rankId]
	for i := 0; i < len(cfgs); i++ {
		if len(cfgs[i].Ranking) <= 0 {
			continue
		}
		if len(cfgs[i].Ranking) == 1 {
			if rank == cfgs[i].Ranking[0] {
				return cfgs[i]
			}
		}

		if len(cfgs[i].Ranking) >= 2 {
			if rank >= cfgs[i].Ranking[0] && rank <= cfgs[i].Ranking[1] {
				return cfgs[i]
			}
		}
	}
	return nil
}

func GetDayRankRewardCfgByTypes(rankId int) []*DayRankingRewardDayRankingRewardCfg {
	return gameDb.DayRankingRewardDayRankingRewardCfgsByType[rankId]
}

func GetDayRankRewardCfgByTypesAndRank(rankId, rank int) *DayRankingRewardDayRankingRewardCfg {
	for _, data := range gameDb.DayRankingRewardDayRankingRewardCfgsByType[rankId] {
		if len(data.Ranking) >= 1 {
			if rank == data.Ranking[0] {
				return data
			}
		}
		if len(data.Ranking) >= 2 {
			if rank >= data.Ranking[0] && rank <= data.Ranking[1] {
				return data
			}
		}
	}
	return nil
}

func GetGiftCodeByCode(code string) *GiftCodeGiftCodeCfg {
	return gameDb.GiftCodeMap[code]
}
func GetAchievementAchievementCfgByConditionType(conditionType int) []*AchievementAchievementCfg {
	return gameDb.AchievementAchievementCfgsByConditionType[conditionType]
}

func GetAchievementAchievementConditionState() map[int]bool {
	return gameDb.AchievementConditionState
}

func GetAchievementAchievementConditionIdState() map[int]bool {
	return gameDb.AchievementConditionIdState
}

func GetAchievementConditionIdAndCondition() map[int]int {
	return gameDb.AchievementConditionIdAndCondition
}

func GetLimitedGiftByTypeAndLv(t, lv int) *LimitedGiftLimitedGiftCfg {
	return gameDb.LimitedGiftMap[t][lv]
}

func GetWorldLeaderByStageId(stageId int) *WorldLeaderConfCfg {
	for _, v := range gameDb.WorldLeaderConfCfgs {
		if v.StageId == stageId {
			return v
		}
	}
	return nil
}

func GetWorldLeaderCfs() []*WorldLeaderConfCfg {
	return gameDb.WorldLeaderSlice
}

func GetWarOrderCycle() map[int]*WarOrderCycleWarOrderCycleCfg {
	return gameDb.WarOrderCycleWarOrderCycleCfgs
}
func GetWarOrderTaskByConditionType(t int) []*WarOrderCycleTaskWarOrderCycleTaskCfg {
	return gameDb.WarOrderTaskMap[t]
}
func GetWarOrderWeekTaskByConditionType(t int) []*WarOrderWeekTaskWarOrderWeekTaskCfg {
	return gameDb.WarOrderWeekTaskMap[t]
}

func GetRewardOnlineCfgs() map[int]*RewardsOnlineAwardCfg {
	return gameDb.RewardsOnlineAwardCfgs
}

func GetTalentGeneral() map[int]*TalentgeneralTalentCfg {
	return gameDb.TalentgeneralTalentCfgs
}

func GetElfRecoverByTypeAndQuality(t, q int) *ElfRecoverElfRecoverCfg {
	return gameDb.ElfRecoverMap[t][q]
}
func GetElfSkillBySkillIdAndLv(sid, slv int) *ElfSkillElfGrowCfg {
	return gameDb.ElfSkillMap[sid][slv]
}

func GetDrawCfgBySeasonAndType(season, types int) *DrawDrawCfg {
	return gameDb.DrawCfgBySeasonAndType[season][types]
}

func GetCrossShaBakePerAwardByRank(rank int) *KuafushabakeRewardserverKuafushabakeRewardserverCfg {
	for _, v := range gameDb.KuafushabakeRewardserverKuafushabakeRewardserverCfgs {
		if rank >= v.Rank[0] && rank <= v.Rank[1] {
			return v
		}
	}
	return nil
}

func GetCrossShaBakeUniAwardByRank(rank int) *KuafushabakeRewarduniKuafushabakeRewarduniCfg {
	for _, v := range gameDb.KuafushabakeRewarduniKuafushabakeRewarduniCfgs {
		if rank >= v.Rank[0] && rank <= v.Rank[1] {
			return v
		}
	}

	return nil
}

func GetDrawShopDrawShopCfgByShop(season, id int) *DrawShopDrawShopCfg {
	return gameDb.DrawShopDrawShopCfgByShop[season][id]
}

func GetDrawNowSeason(openDay int) (*DrawDrawCfg, int, int) {
	for _, v := range gameDb.DrawDrawCfgs {
		if len(v.Time) >= 2 {
			if v.Time[0] <= openDay && openDay <= v.Time[1] {
				return v, v.Type1, v.Time[1]
			}
		}
	}
	return nil, -1, -1
}

func GetCutTreasureByLv(lv int) *CutTreasureCutTreasureCfg {
	return gameDb.CutTreasureMap[lv]
}

func GetTreasureCfgBySeasonAndType(season, types int) *XunlongXunlongCfg {

	return gameDb.TreasureCfgBySeasonAndType[season][types]
}

func GetTreasureNowSeason(openDay int) *XunlongXunlongCfg {
	for _, v := range gameDb.XunlongXunlongCfgs {
		if len(v.Time) >= 2 {
			if v.Time[0] <= openDay && openDay <= v.Time[1] {
				return v
			}
		}
	}
	return nil
}

func GetHolyBeastByTypesAndStar(types, star int) *HolyBeastHolyBeastCfg {
	return gameDb.HolyBeastCfgByTypeAndStar[types][star]
}
func GetHolyBeastByTypesMaxStar(types int) int {
	return gameDb.HolyBeastCfgByTypeMaxStar[types]
}

func GetReturnHolyPoints(types, star int) ItemInfos {
	items := gameDb.HolyBeastCfgByTypeAndStar[types]
	if items == nil {
		return nil
	}
	allItems := make(map[int]int)
	allItemInfos := make(ItemInfos, 0)
	for i := 1; i <= star; i++ {
		for _, data := range items[i].Active {
			allItems[data.ItemId] += data.Count
		}
	}
	for k, v := range allItems {
		allItemInfos = append(allItemInfos, &ItemInfo{k, v})
	}

	return allItemInfos
}

func GetFitHolyEquipByPosAndGrade(st, t, g int) *FitHolyEquipFitHolyEquipCfg {
	return gameDb.FitHolyEquipMap[st][t][g]
}

func GetDailyTaskDailyTaskCfgByTypeTest() {
	if gameDb.DailyTaskDailytaskCfgsByType != nil {
		for types, data := range gameDb.DailyTaskDailytaskCfgsByType {
			logger.Info("Test DailyTaskDailytaskCfgsByType types:%v  data:%v", types, data)
		}
	} else {
		logger.Error("Test gameDb.DailyTaskDailytaskCfgsByType == nil")
	}

}

func GetCompetitiveRobotUserByScore(score int) *RobotRobotCfg {
	for _, info := range gameDb.RobotCfgSlice {
		if info.LevelCompetive > score {
			return info
		}
	}
	for _, info := range gameDb.RobotCfgSlice {
		if info.LevelCompetive <= score {
			return info
		}
	}

	return nil
}

func GetCompetitiveRobotUserByCombat(userCombat, lastMarkUserId int) *RobotRobotCfg {
	if len(gameDb.RobotCfgSlice) <= 0 {
		return nil
	}
	data := make([]*RobotRobotCfg, 0)
	for _, info := range gameDb.RobotCfgSlice {
		if lastMarkUserId < 0 {
			if info.Id == -lastMarkUserId {
				continue
			}
		}
		if userCombat >= info.Fighting[0] && userCombat <= info.Fighting[1] {
			data = append(data, info)
		}
	}

	if len(data) <= 0 {
		return gameDb.RobotCfgSlice[0]
	}
	randNum := rand.Intn(len(data))
	return data[randNum]
}

func GetCompetitiveRobotUserByScoreSlice(nowScore, nextScore int) *RobotRobotCfg {
	if len(gameDb.RobotCfgSlice) <= 0 {
		return nil
	}
	data := make([]*RobotRobotCfg, 0)
	for _, info := range gameDb.RobotCfgSlice {
		if info.LevelCompetive > nowScore && info.LevelCompetive <= nextScore {
			data = append(data, info)
		}
	}

	if len(data) <= 0 {
		return gameDb.RobotCfgSlice[0]
	}
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(len(data))
	return data[randNum]
}

func GetXunLongRoundCfgByType(season, round int, cfgId []int) ItemInfos {

	itemInfos := make(ItemInfos, 0)
	data := gameDb.XunLongRoundSliceCfg[season]
	if data == nil {
		return itemInfos
	}
	haveGet := make(map[int]bool)
	for _, id := range cfgId {
		haveGet[id] = true
	}

	for _, info := range data {
		if haveGet[info.Id] {
			continue
		}
		if info.Rounds > round {
			continue
		}
		itemInfos = append(itemInfos, info.Reward...)
	}
	return itemInfos
}

func GetXunLongTypeTime(season int) int {
	return gameDb.XunLongTypeTime[season]
}

func GetXunLongBeforeDaySeason(openDay int) *XunlongXunlongCfg {
	for _, v := range gameDb.XunlongXunlongCfgs {
		if len(v.Time) >= 2 {
			if openDay == v.Time[1] {
				return v
			}
		}
	}
	return nil
}

func GetDailyRankMarkCfg(types, score int) []*DayRankingMarkDayRankingMarkCfg {
	cfg := gameDb.DayRankingMarkSliceCfg[types]
	if cfg == nil {
		return nil
	}
	info := make([]*DayRankingMarkDayRankingMarkCfg, 0)

	for _, data := range cfg {
		if score >= data.Mark {
			info = append(info, data)
		}
	}
	return info
}

func GetDailyRankBuyGiftCfg(types2, rechargeNum, typeId int) *DayRankingGiftDayRankingGiftCfg {

	cfg := GetDayRankingGiftDayRankingGiftCfg(typeId)
	if cfg == nil {
		return nil
	}

	logger.Debug(" data.Consume:%v == rechargeNum:%v", cfg.Consume, rechargeNum)
	//if int(math.Ceil(float64(cfg.Consume)*(float64(cfg.Discount)/100))) == rechargeNum && cfg.Type2 == types2 {
	if cfg.Consume == rechargeNum && cfg.Type2 == types2 {
		return cfg
	}
	return nil
}

func GetAllDayRankingCfg() map[int]*DayRankingDayRankingCfg {
	return gameDb.DayRankingDayRankingCfgs
}

func GetDailyRankMarkMinScore(types int) int {

	minScore := 0
	cfg := gameDb.DayRankingMarkSliceCfg[types]
	if cfg == nil {
		return minScore
	}

	for index, data := range cfg {
		if index == 0 {
			minScore = data.Mark
		}
		if data.Mark < minScore {
			minScore = data.Mark
		}
	}
	return minScore
}

func GetChuanShiSuitCfgs() map[int]*ChuanShiSuitChuanShiSuitCfg {
	return gameDb.ChuanShiSuitChuanShiSuitCfgs
}
func GetChuanShiSuitByTypeAndLv(t, lv int) *ChuanShiSuitTypeChuanShiSuitTypeCfg {
	return gameDb.ChuanShiSuitMap[t][lv]
}

func GetSpendRebateCfgs() map[int]*SpendrebatesSpendrebatesCfg {
	return gameDb.SpendrebatesSpendrebatesCfgs
}
func GetSpendRebateByType(t int) *SpendrebatesSpendrebatesCfg {
	return gameDb.SpendRebateMap[t]
}

func GetContRechargeTypes() map[int]*ContRechargeContRechargeCfg {
	return gameDb.ContRechargeMap
}
func GetContRechargeByType(t int) []*ContRechargeContRechargeCfg {
	return gameDb.ContRechargeTypeMap[t]
}

func GetComposeItemCfg(itemId int) *ComposeSubComposeSubCfg {
	return gameDb.ComposeItemMap[itemId]
}

func GetTowerByStageId(stageId int) *TowerTowerCfg {
	for _, v := range gameDb.TowerTowerCfgs {
		if v.Stage == stageId {
			return v
		}
	}
	return nil
}

func GetPowerRoll(ratio float64, isPvp bool) int {
	for _, v := range gameDb.PowerRollPowerRollCfgs {
		if ratio >= v.Min && ratio <= v.Max {
			if isPvp {
				return v.PVPReduce
			} else {
				return v.TowerReduce
			}
		}
	}
	return 0
}

func GetWorldLeaderRewardByRank(group, rank int) *WorldLeaderRewardWorldLeaderRewardCfg {
	data := gameDb.WorldLeaderRewardSlice[group]
	if data == nil {
		return nil
	}
	for _, cfg := range data {
		if cfg.Rank == nil || len(cfg.Rank) < 1 {
			continue
		}

		if len(cfg.Rank) == 1 {
			if cfg.Rank[0] == rank {
				return cfg
			}
		}

		if len(cfg.Rank) >= 2 {
			if rank >= cfg.Rank[0] && rank <= cfg.Rank[1] {
				return cfg
			}
		}
	}
	return nil
}

func GetModelBagInfo(items map[int]int, items1 map[int]int) model.Bag {
	infos := model.Bag{}
	if items != nil && len(items) > 0 {
		for itemId, count := range items {
			infos = append(infos, &model.Item{ItemId: itemId, Count: count})
		}
	}
	if items1 != nil && len(items1) > 0 {
		for itemId, count := range items1 {
			infos = append(infos, &model.Item{ItemId: itemId, Count: count})
		}
	}
	return infos
}

func GetAncientBossCfgsByArea(area int) map[int]*AncientBossAncientBossCfg {
	return gameDb.AncientBossMap[area]
}
func GetAncientBossCfgs() map[int]*AncientBossAncientBossCfg {
	return gameDb.AncientBossAncientBossCfgs
}

func GetGuardRankCfgs() map[int]*GuardRankGuardRankCfg {
	return gameDb.GuardRankGuardRankCfgs
}

func GetAutoActiveTitleCfgs() map[int]*TitleTitleCfg {
	return gameDb.TitleAutoActiveMap
}

func GetKillMonsterUniCfgs() map[int]*FirstBlooduniFirstBlooduniCfg {
	return gameDb.KillMonsterUniMap
}
func GetKillMonsterPerCfgs() map[int]*FirstBloodPerFirstBloodperCfg {
	return gameDb.KillMonsterPerMap
}
func GetKillMonsterType() map[int]int {
	return gameDb.KillMonsterMilMaxLv
}
func GetKillMonsterMilByTypeAndLv(t, lv int) *FirstBloodmilFirstBloodmilCfg {
	return gameDb.KillMonsterMilMap[t][lv]
}
func GetKillMonsterUniByStageId(stageId int) *FirstBlooduniFirstBlooduniCfg {
	return gameDb.KillMonsterUniMap[stageId]
}
func GetKillMonsterPerByStageId(stageId int) *FirstBloodPerFirstBloodperCfg {
	return gameDb.KillMonsterPerMap[stageId]
}

func GetAncientTreasureZhuLinLvById(treasureId, lv int) *TreasureArtTreasureArtCfg {
	if gameDb.AncientTreasureZhuLinMap[treasureId] == nil {
		return nil
	}
	return gameDb.AncientTreasureZhuLinMap[treasureId][lv]
}

func GetAncientTreasureZhuLinMaxLvById(treasureId int) int {
	return gameDb.AncientTreasureZhuLinMaxLvMap[treasureId]
}

func GetAncientTreasureStarById(treasureId, lv int) *TreasureStarsTreasureStarsCfg {
	if gameDb.AncientTreasureStarMap[treasureId] == nil {
		return nil
	}
	return gameDb.AncientTreasureStarMap[treasureId][lv]
}

func GetAncientTreasureStarMaxLvById(treasureId int) int {
	return gameDb.AncientTreasureMaxStarMap[treasureId]
}

func GetAncientTreasureJueXinCfg(treasureId int) *TreasureAwakenTreasureAwakenCfg {
	return gameDb.AncientTreasureJueXinMap[treasureId]
}

func GetAllAncientTreasureSuit() map[int]*TreasureSuitTreasureSuitCfg {
	return gameDb.TreasureSuitTreasureSuitCfgs
}

func RandTreasureShopItem() map[int]*TreasureShopTreasureShopCfg {
	weightMap := make(map[int]int)
	cfgs := gameDb.TreasureShopTreasureShopCfgs
	for id, cfg := range cfgs {
		weightMap[id] = cfg.Weight
	}
	dataMap := make(map[int]*TreasureShopTreasureShopCfg)
	randNum := gameDb.InitConf.TreasureNum
	for i := 0; i < constConstant.COMPUTE_TEN_THOUSAND; i++ {
		randId := common.RandWeightByMap(weightMap)
		if _, ok := dataMap[randId]; ok {
			continue
		}
		dataMap[randId] = cfgs[randId]
		if len(dataMap) >= randNum {
			break
		}
		delete(weightMap, randId)
	}
	return dataMap
}
func GetTreasureDiscount(price int) int {
	cfgs := gameDb.TreasureDiscountTreasureDiscountCfgs
	discountMap := make(map[int]int)
	discountSlice := make([]int, 0)
	for _, cfg := range cfgs {
		discountMap[cfg.Condition] = cfg.Discount
		discountSlice = append(discountSlice, cfg.Condition)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(discountSlice)))
	discount := 0
	for _, condition := range discountSlice {
		if price >= condition {
			discount = discountMap[condition]
			break
		}
	}
	return discount
}

func GetChuanShiStrengthenByPosAndLv(pos, lv int) *ChuanShiStrengthenChuanShiStrengthenCfg {
	return gameDb.ChuanShiStrengthenMap[pos][lv]
}
func GetChuanShiStrengthenSuitByTypeAndLv(t, lv int) *ChuanShiStrengthenLinkChuanShiStrengthenLinkCfg {
	return gameDb.ChuanShiStrengthenSuitMap[t][lv]
}
func GetChuanShiStrengthenLinkCfgs() map[int]*ChuanShiStrengthenLinkChuanShiStrengthenLinkCfg {
	return gameDb.ChuanShiStrengthenLinkChuanShiStrengthenLinkCfgs
}

func GetPetAppendageByPidAndLv(pid, lv int) *PetsAddPetsAddCfg {
	return gameDb.PetAppendageMap[pid][lv]
}
func GetPetAppendageSkillCfgs() map[int]*PetsAddSkillPetsAddSkillCfg {
	return gameDb.PetsAddSkillPetsAddSkillCfgs
}

func GetMagicTowerRewardByRank(rank int) ItemInfos {

	for _, v := range gameDb.MagicTowerRewardMagicTowerRewardCfgs {
		if len(v.Rank) < 2 {
			continue
		}
		if rank >= v.Rank[0] && rank <= v.Rank[1] {
			return v.Reward
		}
	}
	return nil
}

func GetHellBossByStage(stageId int) *HellBossHellBossCfg {
	return gameDb.HellBossStageMap[stageId]
}
func GetHellBossByFloor(floor int) map[int]*HellBossHellBossCfg {
	return gameDb.HellBossMap[floor]
}

func GetAllTrialTask() map[int]*TrialTaskTrialTaskCfg {
	return gameDb.TrialTaskTrialTaskCfgs
}

func GetAllTrialTaskStageAward() map[int]*TrialTotalRewardTrialTotalRewardCfg {
	return gameDb.TrialTotalRewardTrialTotalRewardCfgs
}

/**
*  @Description: 获取活动开启时间
*  @param openTime		开始时间（时分秒）
*  @param closeTime		结束时间（时分秒）
*  @param weeks			开启周日期
*  @return int			开启时间
*  @return int			结束时间
**/
func GetActiveTime(openTime, closeTime HmsTime, weeks []int) (int, int) {

	now := time.Now()
	dayZeroTime := common.GetZeroTimeUnixFrom1970()
	startTime := dayZeroTime + openTime.GetSecondsFromZero()
	stopTime := dayZeroTime + closeTime.GetSecondsFromZero()

	nowWeekNum := time.Now().Weekday()
	if nowWeekNum == time.Sunday {
		nowWeekNum = 7
	}

	nearestDay := 7
	if int(now.Unix()) < stopTime {
		for _, v := range weeks {
			if v == int(nowWeekNum) {
				nearestDay = 0
				break
			}
		}
	}
	if nearestDay != 0 {
		for _, v := range weeks {
			if v <= int(nowWeekNum) {
				v += 7
			}
			dayDiff := v - int(nowWeekNum)
			if nearestDay > dayDiff {
				nearestDay = dayDiff
			}
		}
	}

	daySecond := nearestDay * 86400
	logger.Info("获取活动时间,后面：%v天，开始时间：%v,结束时间：%v", nearestDay, startTime+daySecond, stopTime+daySecond)
	return startTime + daySecond, stopTime + daySecond
}

func GetTowerRankReward(rank int) ItemInfos {
	items := make(ItemInfos, 0)
	for _, cfg := range gameDb.TowerRankRewardTowerRankRewardCfgs {
		min, max := cfg.Rank[0], cfg.Rank[0]
		if len(cfg.Rank) > 1 {
			max = cfg.Rank[1]
		}
		if rank >= min && rank <= max {
			items = cfg.Reward
			break
		}
	}
	return items
}
func GetTowerReward(floor int) ItemInfos {
	items := make(ItemInfos, 0)
	for _, cfg := range gameDb.TowerRewardTowerRewardCfgs {
		if floor >= cfg.Floor[0] && floor <= cfg.Floor[1] {
			items = cfg.Reward
		}
	}
	return items
}

func GetDaBaoEquipByTypeAndLv(t, lv int) *DaBaoEquipDaBaoEquipCfg {
	return gameDb.DaBaoEquipMap[t][lv]
}
func GetDaBaoMysteryByStageId(stageId int) *DaBaoMysteryDaBaoMysteryCfg {
	return gameDb.DaBaoMysteryMap[stageId]
}
func GetDaBaoEquipAdditionByTypeAndLv(equipT, addT, lv int) IntSlice {
	effect := make(IntSlice, 0)
	lvSlice := make([]int, 0)
	for equipClass := range gameDb.DaBaoEquipAdditionMap[equipT][addT] {
		lvSlice = append(lvSlice, equipClass)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(lvSlice)))
	for _, equipClass := range lvSlice {
		if equipClass <= lv {
			effect = gameDb.DaBaoEquipAdditionMap[equipT][addT][equipClass].Effect
			break
		}
	}
	return effect
}

func GetAllGuildAutoCreateGuildAutoCreateCfg() map[int]*GuildAutoCreateGuildAutoCreateCfg {
	return gameDb.GuildAutoCreateGuildAutoCreateCfgs
}

func GetAllGuildNameGuildNameCfgs() map[int]*GuildNameGuildNameCfg {
	return gameDb.GuildNameGuildNameCfgs
}
func GetComposeEquipByComposeId(composeId int) *ComposeEquipSubComposeEquipSubCfg {
	return gameDb.ComposeEquipMap[composeId]
}

func GetLabelTaskConditionSlice(id int) []int {
	return gameDb.LabelTaskMap[id]
}

func GetFirstDropItemInfoByItemId(itemId int) *FirstDropFirstDropCfg {
	return gameDb.FirstDropByItemId[itemId]
}

func GetFirstDropItemInfoByTypes(types int) []*FirstDropFirstDropCfg {
	return gameDb.FirstDropByTypes[types]
}

func GetPrivilegeCfgs() map[int]*PrivilegePrivilegeCfg {
	return gameDb.PrivilegePrivilegeCfgs
}