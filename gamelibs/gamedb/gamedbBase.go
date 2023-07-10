package gamedb


var fileInfos = []fileInfo{

	fileInfo{"game.xlsx", []sheetInfo{
		{"game", (*GameDb).loadGameConf, GameBaseCfg{}},
	}},
	
	fileInfo{"PaoDianReward.xlsx", []sheetInfo{
			{SheetName: "PaoDianReward", Initer: mapLoader("PaoDianRewardPaoDianRewardCfgs", "Id"), ObjProptype: PaoDianRewardPaoDianRewardCfg{}},
	}},
	fileInfo{"Spendrebates.xlsx", []sheetInfo{
			{SheetName: "spendrebates", Initer: mapLoader("SpendrebatesSpendrebatesCfgs", "Id"), ObjProptype: SpendrebatesSpendrebatesCfg{}},
	}},
	fileInfo{"accumulate.xlsx", []sheetInfo{
			{SheetName: "accumulate", Initer: mapLoader("AccumulateAccumulateCfgs", "Id"), ObjProptype: AccumulateAccumulateCfg{}},
	}},
	fileInfo{"achievement.xlsx", []sheetInfo{
			{SheetName: "achievement", Initer: mapLoader("AchievementAchievementCfgs", "Id"), ObjProptype: AchievementAchievementCfg{}},
	}},
	fileInfo{"achievementMedal.xlsx", []sheetInfo{
			{SheetName: "medal", Initer: mapLoader("AchievementMedalMedalCfgs", "Id"), ObjProptype: AchievementMedalMedalCfg{}},
	}},
	fileInfo{"ancientBoss.xlsx", []sheetInfo{
			{SheetName: "ancientBoss", Initer: mapLoader("AncientBossAncientBossCfgs", "StageId"), ObjProptype: AncientBossAncientBossCfg{}},
	}},
	fileInfo{"ancientSkillGrade.xlsx", []sheetInfo{
			{SheetName: "ancientSkillGrade", Initer: mapLoader("AncientSkillGradeAncientSkillGradeCfgs", "Level"), ObjProptype: AncientSkillGradeAncientSkillGradeCfg{}},
	}},
	fileInfo{"ancientSkillLevel.xlsx", []sheetInfo{
			{SheetName: "ancientSkillLevel", Initer: mapLoader("AncientSkillLevelAncientSkillLevelCfgs", "Level"), ObjProptype: AncientSkillLevelAncientSkillLevelCfg{}},
	}},
	fileInfo{"archer.xlsx", []sheetInfo{
			{SheetName: "magic", Initer: mapLoader("ArcherMagicCfgs", "Id"), ObjProptype: ArcherMagicCfg{}},
	}},
	fileInfo{"archerElement.xlsx", []sheetInfo{
			{SheetName: "magicElement", Initer: mapLoader("ArcherElementMagicElementCfgs", "Id"), ObjProptype: ArcherElementMagicElementCfg{}},
	}},
	fileInfo{"area.xlsx", []sheetInfo{
			{SheetName: "area", Initer: mapLoader("AreaAreaCfgs", "Id"), ObjProptype: AreaAreaCfg{}},
	}},
	fileInfo{"areaLevel.xlsx", []sheetInfo{
			{SheetName: "areaLevel", Initer: mapLoader("AreaLevelAreaLevelCfgs", "Id"), ObjProptype: AreaLevelAreaLevelCfg{}},
	}},
	fileInfo{"arenaBuy.xlsx", []sheetInfo{
			{SheetName: "arenaBuy", Initer: mapLoader("ArenaBuyArenaBuyCfgs", "Num"), ObjProptype: ArenaBuyArenaBuyCfg{}},
	}},
	fileInfo{"arenaMatch.xlsx", []sheetInfo{
			{SheetName: "arenaMatch", Initer: mapLoader("ArenaMatchArenaMatchCfgs", "RankMin"), ObjProptype: ArenaMatchArenaMatchCfg{}},
	}},
	fileInfo{"arenaRank.xlsx", []sheetInfo{
			{SheetName: "arenaRank", Initer: mapLoader("ArenaRankArenaRankCfgs", "Id"), ObjProptype: ArenaRankArenaRankCfg{}},
	}},
	fileInfo{"aspd.xlsx", []sheetInfo{
			{SheetName: "aspd", Initer: mapLoader("AspdAspdCfgs", "Id"), ObjProptype: AspdAspdCfg{}},
	}},
	fileInfo{"atlas.xlsx", []sheetInfo{
			{SheetName: "atlas", Initer: mapLoader("AtlasAtlasCfgs", "Id"), ObjProptype: AtlasAtlasCfg{}},
	}},
	fileInfo{"atlasGather.xlsx", []sheetInfo{
			{SheetName: "atlasGather", Initer: mapLoader("AtlasGatherAtlasGatherCfgs", "Id"), ObjProptype: AtlasGatherAtlasGatherCfg{}},
	}},
	fileInfo{"atlasPos.xlsx", []sheetInfo{
			{SheetName: "atlasPos", Initer: mapLoader("AtlasPosAtlasPosCfgs", "Pos"), ObjProptype: AtlasPosAtlasPosCfg{}},
	}},
	fileInfo{"atlasStar.xlsx", []sheetInfo{
			{SheetName: "atlasStar", Initer: mapLoader("AtlasStarAtlasStarCfgs", "Id"), ObjProptype: AtlasStarAtlasStarCfg{}},
	}},
	fileInfo{"atlasUpgrade.xlsx", []sheetInfo{
			{SheetName: "atlasUpgrade", Initer: mapLoader("AtlasUpgradeAtlasUpgradeCfgs", "Id"), ObjProptype: AtlasUpgradeAtlasUpgradeCfg{}},
	}},
	fileInfo{"attackEnemy.xlsx", []sheetInfo{
			{SheetName: "attackEnemy", Initer: mapLoader("AttackEnemyAttackEnemyCfgs", "Id"), ObjProptype: AttackEnemyAttackEnemyCfg{}},
	}},
	fileInfo{"attackEnemyCard.xlsx", []sheetInfo{
			{SheetName: "attackEnemyCard", Initer: mapLoader("AttackEnemyCardAttackEnemyCardCfgs", "Type"), ObjProptype: AttackEnemyCardAttackEnemyCardCfg{}},
	}},
	fileInfo{"auction.xlsx", []sheetInfo{
			{SheetName: "auctioin", Initer: mapLoader("AuctionAuctioinCfgs", "Id"), ObjProptype: AuctionAuctioinCfg{}},
	}},
	fileInfo{"awaken.xlsx", []sheetInfo{
			{SheetName: "awaken", Initer: mapLoader("AwakenAwakenCfgs", "Id"), ObjProptype: AwakenAwakenCfg{}},
	}},
	fileInfo{"awakenTitle.xlsx", []sheetInfo{
			{SheetName: "awakenTitle", Initer: mapLoader("AwakenTitleAwakenTitleCfgs", "Rank"), ObjProptype: AwakenTitleAwakenTitleCfg{}},
	}},
	fileInfo{"bag.xlsx", []sheetInfo{
			{SheetName: "spaceAdd", Initer: mapLoader("BagSpaceAddCfgs", "Id"), ObjProptype: BagSpaceAddCfg{}},
	}},
	fileInfo{"bindGroup.xlsx", []sheetInfo{
			{SheetName: "bindGroup", Initer: mapLoader("BindGroupBindGroupCfgs", "BindGroup"), ObjProptype: BindGroupBindGroupCfg{}},
	}},
	fileInfo{"bless.xlsx", []sheetInfo{
			{SheetName: "bless", Initer: mapLoader("BlessBlessCfgs", "Id"), ObjProptype: BlessBlessCfg{}},
	}},
	fileInfo{"bossFamily.xlsx", []sheetInfo{
			{SheetName: "bossFamily", Initer: mapLoader("BossFamilyBossFamilyCfgs", "StageId"), ObjProptype: BossFamilyBossFamilyCfg{}},
	}},
	fileInfo{"buff.xlsx", []sheetInfo{
			{SheetName: "buff", Initer: mapLoader("BuffBuffCfgs", "Id"), ObjProptype: BuffBuffCfg{}},
	}},
	fileInfo{"chat.xlsx", []sheetInfo{
			{SheetName: "clear", Initer: mapLoader("ChatClearCfgs", "Type"), ObjProptype: ChatClearCfg{}},
	}},
	fileInfo{"chuanShiEquip.xlsx", []sheetInfo{
			{SheetName: "chuanShiEquip", Initer: mapLoader("ChuanShiEquipChuanShiEquipCfgs", "Id"), ObjProptype: ChuanShiEquipChuanShiEquipCfg{}},
	}},
	fileInfo{"chuanShiEquipType.xlsx", []sheetInfo{
			{SheetName: "chuanShiEquipType", Initer: mapLoader("ChuanShiEquipTypeChuanShiEquipTypeCfgs", "Id"), ObjProptype: ChuanShiEquipTypeChuanShiEquipTypeCfg{}},
	}},
	fileInfo{"chuanShiStrengthen.xlsx", []sheetInfo{
			{SheetName: "chuanShiStrengthen", Initer: mapLoader("ChuanShiStrengthenChuanShiStrengthenCfgs", "Id"), ObjProptype: ChuanShiStrengthenChuanShiStrengthenCfg{}},
	}},
	fileInfo{"chuanShiStrengthenLink.xlsx", []sheetInfo{
			{SheetName: "chuanShiStrengthenLink", Initer: mapLoader("ChuanShiStrengthenLinkChuanShiStrengthenLinkCfgs", "Id"), ObjProptype: ChuanShiStrengthenLinkChuanShiStrengthenLinkCfg{}},
	}},
	fileInfo{"chuanShiSuit.xlsx", []sheetInfo{
			{SheetName: "chuanShiSuit", Initer: mapLoader("ChuanShiSuitChuanShiSuitCfgs", "Id"), ObjProptype: ChuanShiSuitChuanShiSuitCfg{}},
	}},
	fileInfo{"chuanShiSuitType.xlsx", []sheetInfo{
			{SheetName: "chuanShiSuitType", Initer: mapLoader("ChuanShiSuitTypeChuanShiSuitTypeCfgs", "Id"), ObjProptype: ChuanShiSuitTypeChuanShiSuitTypeCfg{}},
	}},
	fileInfo{"clear.xlsx", []sheetInfo{
			{SheetName: "clear", Initer: mapLoader("ClearClearCfgs", "Type"), ObjProptype: ClearClearCfg{}},
	}},
	fileInfo{"clearRate.xlsx", []sheetInfo{
			{SheetName: "clearRate", Initer: mapLoader("ClearRateClearRateCfgs", "Id"), ObjProptype: ClearRateClearRateCfg{}},
	}},
	fileInfo{"collection.xlsx", []sheetInfo{
			{SheetName: "collection", Initer: mapLoader("CollectionCollectionCfgs", "Id"), ObjProptype: CollectionCollectionCfg{}},
	}},
	fileInfo{"competitve.xlsx", []sheetInfo{
			{SheetName: "competitve", Initer: mapLoader("CompetitveCompetitveCfgs", "Id"), ObjProptype: CompetitveCompetitveCfg{}},
	}},
	fileInfo{"competitveReward.xlsx", []sheetInfo{
			{SheetName: "rankReward", Initer: mapLoader("CompetitveRewardRankRewardCfgs", "Id"), ObjProptype: CompetitveRewardRankRewardCfg{}},
	}},
	fileInfo{"composeChuanShiSub.xlsx", []sheetInfo{
			{SheetName: "composeChuanShiSub", Initer: mapLoader("ComposeChuanShiSubComposeChuanShiSubCfgs", "Id"), ObjProptype: ComposeChuanShiSubComposeChuanShiSubCfg{}},
	}},
	fileInfo{"composeChuanShiType.xlsx", []sheetInfo{
			{SheetName: "composeChuanShiType", Initer: mapLoader("ComposeChuanShiTypeComposeChuanShiTypeCfgs", "Id"), ObjProptype: ComposeChuanShiTypeComposeChuanShiTypeCfg{}},
	}},
	fileInfo{"composeEquipSub.xlsx", []sheetInfo{
			{SheetName: "composeEquipSub", Initer: mapLoader("ComposeEquipSubComposeEquipSubCfgs", "Id"), ObjProptype: ComposeEquipSubComposeEquipSubCfg{}},
	}},
	fileInfo{"composeEquipType.xlsx", []sheetInfo{
			{SheetName: "composeEquipType", Initer: mapLoader("ComposeEquipTypeComposeEquipTypeCfgs", "Id"), ObjProptype: ComposeEquipTypeComposeEquipTypeCfg{}},
	}},
	fileInfo{"composeSub.xlsx", []sheetInfo{
			{SheetName: "composeSub", Initer: mapLoader("ComposeSubComposeSubCfgs", "Id"), ObjProptype: ComposeSubComposeSubCfg{}},
	}},
	fileInfo{"composeType.xlsx", []sheetInfo{
			{SheetName: "composeType", Initer: mapLoader("ComposeTypeComposeTypeCfgs", "Id"), ObjProptype: ComposeTypeComposeTypeCfg{}},
	}},
	fileInfo{"condition.xlsx", []sheetInfo{
			{SheetName: "condition", Initer: mapLoader("ConditionConditionCfgs", "Id"), ObjProptype: ConditionConditionCfg{}},
	}},
	fileInfo{"contRecharge.xlsx", []sheetInfo{
			{SheetName: "contRecharge", Initer: mapLoader("ContRechargeContRechargeCfgs", "Id"), ObjProptype: ContRechargeContRechargeCfg{}},
	}},
	fileInfo{"crossArena.xlsx", []sheetInfo{
			{SheetName: "crossArena", Initer: mapLoader("CrossArenaCrossArenaCfgs", "Id"), ObjProptype: CrossArenaCrossArenaCfg{}},
	}},
	fileInfo{"crossArenaReward.xlsx", []sheetInfo{
			{SheetName: "crossArenaReward", Initer: mapLoader("CrossArenaRewardCrossArenaRewardCfgs", "Id"), ObjProptype: CrossArenaRewardCrossArenaRewardCfg{}},
	}},
	fileInfo{"crossArenaRobot.xlsx", []sheetInfo{
			{SheetName: "crossArenaRobot", Initer: mapLoader("CrossArenaRobotCrossArenaRobotCfgs", "Id"), ObjProptype: CrossArenaRobotCrossArenaRobotCfg{}},
	}},
	fileInfo{"crossArenaTime.xlsx", []sheetInfo{
			{SheetName: "crossArenaTime", Initer: mapLoader("CrossArenaTimeCrossArenaTimeCfgs", "Id"), ObjProptype: CrossArenaTimeCrossArenaTimeCfg{}},
	}},
	fileInfo{"cumulationsign.xlsx", []sheetInfo{
			{SheetName: "cumulationsign", Initer: mapLoader("CumulationsignCumulationsignCfgs", "Id"), ObjProptype: CumulationsignCumulationsignCfg{}},
	}},
	fileInfo{"cutTreasure.xlsx", []sheetInfo{
			{SheetName: "cutTreasure", Initer: mapLoader("CutTreasureCutTreasureCfgs", "Id"), ObjProptype: CutTreasureCutTreasureCfg{}},
	}},
	fileInfo{"daBaoEquip.xlsx", []sheetInfo{
			{SheetName: "daBaoEquip", Initer: mapLoader("DaBaoEquipDaBaoEquipCfgs", "Id"), ObjProptype: DaBaoEquipDaBaoEquipCfg{}},
	}},
	fileInfo{"daBaoEquipAddition.xlsx", []sheetInfo{
			{SheetName: "daBaoEquipAddition", Initer: mapLoader("DaBaoEquipAdditionDaBaoEquipAdditionCfgs", "Id"), ObjProptype: DaBaoEquipAdditionDaBaoEquipAdditionCfg{}},
	}},
	fileInfo{"daBaoMystery.xlsx", []sheetInfo{
			{SheetName: "daBaoMystery", Initer: mapLoader("DaBaoMysteryDaBaoMysteryCfgs", "Id"), ObjProptype: DaBaoMysteryDaBaoMysteryCfg{}},
	}},
	fileInfo{"dailyActivity.xlsx", []sheetInfo{
			{SheetName: "dailyActivity", Initer: mapLoader("DailyActivityDailyActivityCfgs", "Id"), ObjProptype: DailyActivityDailyActivityCfg{}},
	}},
	fileInfo{"dailyReward.xlsx", []sheetInfo{
			{SheetName: "dailyReward", Initer: mapLoader("DailyRewardDailyRewardCfgs", "Id"), ObjProptype: DailyRewardDailyRewardCfg{}},
	}},
	fileInfo{"dailyTask.xlsx", []sheetInfo{
			{SheetName: "dailytask", Initer: mapLoader("DailyTaskDailytaskCfgs", "Id"), ObjProptype: DailyTaskDailytaskCfg{}},
	}},
	fileInfo{"dailypack.xlsx", []sheetInfo{
			{SheetName: "dailypack", Initer: mapLoader("DailypackDailypackCfgs", "Id"), ObjProptype: DailypackDailypackCfg{}},
	}},
	fileInfo{"darkPalaceBoss.xlsx", []sheetInfo{
			{SheetName: "darkPalaceBoss", Initer: mapLoader("DarkPalaceBossDarkPalaceBossCfgs", "Id"), ObjProptype: DarkPalaceBossDarkPalaceBossCfg{}},
	}},
	fileInfo{"dayRanking.xlsx", []sheetInfo{
			{SheetName: "dayRanking", Initer: mapLoader("DayRankingDayRankingCfgs", "Day"), ObjProptype: DayRankingDayRankingCfg{}},
	}},
	fileInfo{"dayRankingGift.xlsx", []sheetInfo{
			{SheetName: "dayRankingGift", Initer: mapLoader("DayRankingGiftDayRankingGiftCfgs", "Id"), ObjProptype: DayRankingGiftDayRankingGiftCfg{}},
	}},
	fileInfo{"dayRankingMark.xlsx", []sheetInfo{
			{SheetName: "dayRankingMark", Initer: mapLoader("DayRankingMarkDayRankingMarkCfgs", "Id"), ObjProptype: DayRankingMarkDayRankingMarkCfg{}},
	}},
	fileInfo{"dayRankingReward.xlsx", []sheetInfo{
			{SheetName: "dayRankingReward", Initer: mapLoader("DayRankingRewardDayRankingRewardCfgs", "Id"), ObjProptype: DayRankingRewardDayRankingRewardCfg{}},
	}},
	fileInfo{"dictateEquip.xlsx", []sheetInfo{
			{SheetName: "dictateEquip", Initer: mapLoader("DictateEquipDictateEquipCfgs", "Id"), ObjProptype: DictateEquipDictateEquipCfg{}},
	}},
	fileInfo{"dictateSuit.xlsx", []sheetInfo{
			{SheetName: "dictateSuit", Initer: mapLoader("DictateSuitDictateSuitCfgs", "Grade"), ObjProptype: DictateSuitDictateSuitCfg{}},
	}},
	fileInfo{"dragonEquip.xlsx", []sheetInfo{
			{SheetName: "dragonEquip", Initer: mapLoader("DragonEquipDragonEquipCfgs", "Id"), ObjProptype: DragonEquipDragonEquipCfg{}},
	}},
	fileInfo{"dragonEquipLevel.xlsx", []sheetInfo{
			{SheetName: "dragonEquipLevel", Initer: mapLoader("DragonEquipLevelDragonEquipLevelCfgs", "Id"), ObjProptype: DragonEquipLevelDragonEquipLevelCfg{}},
	}},
	fileInfo{"dragonarms.xlsx", []sheetInfo{
			{SheetName: "dragonarms", Initer: mapLoader("DragonarmsDragonarmsCfgs", "Id"), ObjProptype: DragonarmsDragonarmsCfg{}},
	}},
	fileInfo{"draw.xlsx", []sheetInfo{
			{SheetName: "draw", Initer: mapLoader("DrawDrawCfgs", "Id"), ObjProptype: DrawDrawCfg{}},
	}},
	fileInfo{"drawShop.xlsx", []sheetInfo{
			{SheetName: "drawShop", Initer: mapLoader("DrawShopDrawShopCfgs", "Id"), ObjProptype: DrawShopDrawShopCfg{}},
	}},
	fileInfo{"drop.xlsx", []sheetInfo{
			{SheetName: "drop", Initer: mapLoader("DropDropCfgs", "Id"), ObjProptype: DropDropCfg{}},
	}},
	fileInfo{"dropSpecial.xlsx", []sheetInfo{
			{SheetName: "dropSpecial", Initer: mapLoader("DropSpecialDropSpecialCfgs", "Id"), ObjProptype: DropSpecialDropSpecialCfg{}},
	}},
	fileInfo{"effect.xlsx", []sheetInfo{
			{SheetName: "effect", Initer: mapLoader("EffectEffectCfgs", "Id"), ObjProptype: EffectEffectCfg{}},
	}},
	fileInfo{"elfGrow.xlsx", []sheetInfo{
			{SheetName: "elfGrow", Initer: mapLoader("ElfGrowElfGrowCfgs", "Level"), ObjProptype: ElfGrowElfGrowCfg{}},
	}},
	fileInfo{"elfRecover.xlsx", []sheetInfo{
			{SheetName: "elfRecover", Initer: mapLoader("ElfRecoverElfRecoverCfgs", "Id"), ObjProptype: ElfRecoverElfRecoverCfg{}},
	}},
	fileInfo{"elfSkill.xlsx", []sheetInfo{
			{SheetName: "elfGrow", Initer: mapLoader("ElfSkillElfGrowCfgs", "Id"), ObjProptype: ElfSkillElfGrowCfg{}},
	}},
	fileInfo{"equip.xlsx", []sheetInfo{
			{SheetName: "equip", Initer: mapLoader("EquipEquipCfgs", "Id"), ObjProptype: EquipEquipCfg{}},
	}},
	fileInfo{"equipsuit.xlsx", []sheetInfo{
			{SheetName: "equipsuit", Initer: mapLoader("EquipsuitEquipsuitCfgs", "Id"), ObjProptype: EquipsuitEquipsuitCfg{}},
	}},
	fileInfo{"expLevel.xlsx", []sheetInfo{
			{SheetName: "level", Initer: mapLoader("ExpLevelLevelCfgs", "Id"), ObjProptype: ExpLevelLevelCfg{}},
	}},
	fileInfo{"expPill.xlsx", []sheetInfo{
			{SheetName: "expPill", Initer: mapLoader("ExpPillExpPillCfgs", "Id"), ObjProptype: ExpPillExpPillCfg{}},
	}},
	fileInfo{"expPool.xlsx", []sheetInfo{
			{SheetName: "expPool", Initer: mapLoader("ExpPoolExpPoolCfgs", "Id"), ObjProptype: ExpPoolExpPoolCfg{}},
	}},
	fileInfo{"expStage.xlsx", []sheetInfo{
			{SheetName: "expStage", Initer: mapLoader("ExpStageExpStageCfgs", "Id"), ObjProptype: ExpStageExpStageCfg{}},
	}},
	fileInfo{"fabao.xlsx", []sheetInfo{
			{SheetName: "fabao", Initer: mapLoader("FabaoFabaoCfgs", "Id"), ObjProptype: FabaoFabaoCfg{}},
	}},
	fileInfo{"fabaoSkill.xlsx", []sheetInfo{
			{SheetName: "fabaoSkill", Initer: mapLoader("FabaoSkillFabaoSkillCfgs", "Id"), ObjProptype: FabaoSkillFabaoSkillCfg{}},
	}},
	fileInfo{"fabaolevel.xlsx", []sheetInfo{
			{SheetName: "fabaolevel", Initer: mapLoader("FabaolevelFabaolevelCfgs", "Id"), ObjProptype: FabaolevelFabaolevelCfg{}},
	}},
	fileInfo{"fashion.xlsx", []sheetInfo{
			{SheetName: "fashion", Initer: mapLoader("FashionFashionCfgs", "Id"), ObjProptype: FashionFashionCfg{}},
	}},
	fileInfo{"fieldBoss.xlsx", []sheetInfo{
			{SheetName: "fieldBoss", Initer: mapLoader("FieldBossFieldBossCfgs", "StageId"), ObjProptype: FieldBossFieldBossCfg{}},
	}},
	fileInfo{"fieldFight.xlsx", []sheetInfo{
			{SheetName: "fieldBase", Initer: mapLoader("FieldFightFieldBaseCfgs", "Id"), ObjProptype: FieldFightFieldBaseCfg{}},
	}},
	fileInfo{"firstBloodPer.xlsx", []sheetInfo{
			{SheetName: "firstBloodper", Initer: mapLoader("FirstBloodPerFirstBloodperCfgs", "Id"), ObjProptype: FirstBloodPerFirstBloodperCfg{}},
	}},
	fileInfo{"firstBloodmil.xlsx", []sheetInfo{
			{SheetName: "firstBloodmil", Initer: mapLoader("FirstBloodmilFirstBloodmilCfgs", "Id"), ObjProptype: FirstBloodmilFirstBloodmilCfg{}},
	}},
	fileInfo{"firstBlooduni.xlsx", []sheetInfo{
			{SheetName: "firstBlooduni", Initer: mapLoader("FirstBlooduniFirstBlooduniCfgs", "Id"), ObjProptype: FirstBlooduniFirstBlooduniCfg{}},
	}},
	fileInfo{"firstDrop.xlsx", []sheetInfo{
			{SheetName: "firstDrop", Initer: mapLoader("FirstDropFirstDropCfgs", "Id"), ObjProptype: FirstDropFirstDropCfg{}},
	}},
	fileInfo{"firstRechargType.xlsx", []sheetInfo{
			{SheetName: "firstRechargType", Initer: mapLoader("FirstRechargTypeFirstRechargTypeCfgs", "Id"), ObjProptype: FirstRechargTypeFirstRechargTypeCfg{}},
	}},
	fileInfo{"firstRecharge.xlsx", []sheetInfo{
			{SheetName: "firstRecharg", Initer: mapLoader("FirstRechargeFirstRechargCfgs", "Day"), ObjProptype: FirstRechargeFirstRechargCfg{}},
	}},
	fileInfo{"fitFashion.xlsx", []sheetInfo{
			{SheetName: "fitFashion", Initer: mapLoader("FitFashionFitFashionCfgs", "Id"), ObjProptype: FitFashionFitFashionCfg{}},
	}},
	fileInfo{"fitFashionLevel.xlsx", []sheetInfo{
			{SheetName: "fitFashionLevel", Initer: mapLoader("FitFashionLevelFitFashionLevelCfgs", "Id"), ObjProptype: FitFashionLevelFitFashionLevelCfg{}},
	}},
	fileInfo{"fitHolyEquip.xlsx", []sheetInfo{
			{SheetName: "fitHolyEquip", Initer: mapLoader("FitHolyEquipFitHolyEquipCfgs", "Id"), ObjProptype: FitHolyEquipFitHolyEquipCfg{}},
	}},
	fileInfo{"fitHolyEquipSuit.xlsx", []sheetInfo{
			{SheetName: "fitHolyEquipSuit", Initer: mapLoader("FitHolyEquipSuitFitHolyEquipSuitCfgs", "Grade"), ObjProptype: FitHolyEquipSuitFitHolyEquipSuitCfg{}},
	}},
	fileInfo{"fitLevel.xlsx", []sheetInfo{
			{SheetName: "fitLevel", Initer: mapLoader("FitLevelFitLevelCfgs", "Id"), ObjProptype: FitLevelFitLevelCfg{}},
	}},
	fileInfo{"fitSkill.xlsx", []sheetInfo{
			{SheetName: "fitSkill", Initer: mapLoader("FitSkillFitSkillCfgs", "Id"), ObjProptype: FitSkillFitSkillCfg{}},
	}},
	fileInfo{"fitSkillLevel.xlsx", []sheetInfo{
			{SheetName: "fitSkillLevel", Initer: mapLoader("FitSkillLevelFitSkillLevelCfgs", "Id"), ObjProptype: FitSkillLevelFitSkillLevelCfg{}},
	}},
	fileInfo{"fitSkillSlot.xlsx", []sheetInfo{
			{SheetName: "fitSkillSlot", Initer: mapLoader("FitSkillSlotFitSkillSlotCfgs", "Id"), ObjProptype: FitSkillSlotFitSkillSlotCfg{}},
	}},
	fileInfo{"fitSkillStar.xlsx", []sheetInfo{
			{SheetName: "fitSkillStar", Initer: mapLoader("FitSkillStarFitSkillStarCfgs", "Id"), ObjProptype: FitSkillStarFitSkillStarCfg{}},
	}},
	fileInfo{"function.xlsx", []sheetInfo{
			{SheetName: "function", Initer: mapLoader("FunctionFunctionCfgs", "Id"), ObjProptype: FunctionFunctionCfg{}},
	}},
	fileInfo{"gameDot.xlsx", []sheetInfo{
			{SheetName: "gameDot", Initer: mapLoader("GameDotGameDotCfgs", "Id"), ObjProptype: GameDotGameDotCfg{}},
	}},
	fileInfo{"gameText.xlsx", []sheetInfo{
			{SheetName: "errorText", Initer: mapLoader("GameTextErrorTextCfgs", "Id"), ObjProptype: GameTextErrorTextCfg{}},{SheetName: "codeText", Initer: mapLoader("GameTextCodeTextCfgs", "Id"), ObjProptype: GameTextCodeTextCfg{}},
	}},
	fileInfo{"gameword.xlsx", []sheetInfo{
			{SheetName: "game", Initer: mapLoader("GamewordGameCfgs", "Id"), ObjProptype: GamewordGameCfg{}},
	}},
	fileInfo{"gift.xlsx", []sheetInfo{
			{SheetName: "gift", Initer: mapLoader("GiftGiftCfgs", "Id"), ObjProptype: GiftGiftCfg{}},
	}},
	fileInfo{"giftCode.xlsx", []sheetInfo{
			{SheetName: "giftCode", Initer: mapLoader("GiftCodeGiftCodeCfgs", "Id"), ObjProptype: GiftCodeGiftCodeCfg{}},
	}},
	fileInfo{"godBlood.xlsx", []sheetInfo{
			{SheetName: "godBlood", Initer: mapLoader("GodBloodGodBloodCfgs", "Id"), ObjProptype: GodBloodGodBloodCfg{}},
	}},
	fileInfo{"godEquip.xlsx", []sheetInfo{
			{SheetName: "conf", Initer: mapLoader("GodEquipConfCfgs", "Id"), ObjProptype: GodEquipConfCfg{}},
	}},
	fileInfo{"godEquipLevel.xlsx", []sheetInfo{
			{SheetName: "conf", Initer: mapLoader("GodEquipLevelConfCfgs", "Id"), ObjProptype: GodEquipLevelConfCfg{}},
	}},
	fileInfo{"godEquipSuit.xlsx", []sheetInfo{
			{SheetName: "conf", Initer: mapLoader("GodEquipSuitConfCfgs", "Level"), ObjProptype: GodEquipSuitConfCfg{}},
	}},
	fileInfo{"growFund.xlsx", []sheetInfo{
			{SheetName: "growFund", Initer: mapLoader("GrowFundGrowFundCfgs", "Id"), ObjProptype: GrowFundGrowFundCfg{}},
	}},
	fileInfo{"guardRank.xlsx", []sheetInfo{
			{SheetName: "guardRank", Initer: mapLoader("GuardRankGuardRankCfgs", "Id"), ObjProptype: GuardRankGuardRankCfg{}},
	}},
	fileInfo{"guardRounds.xlsx", []sheetInfo{
			{SheetName: "guardRounds", Initer: mapLoader("GuardRoundsGuardRoundsCfgs", "Rounds"), ObjProptype: GuardRoundsGuardRoundsCfg{}},
	}},
	fileInfo{"guild.xlsx", []sheetInfo{
			{SheetName: "guild", Initer: mapLoader("GuildGuildCfgs", "Id"), ObjProptype: GuildGuildCfg{}},
	}},
	fileInfo{"guildActivity.xlsx", []sheetInfo{
			{SheetName: "guildActivity", Initer: mapLoader("GuildActivityGuildActivityCfgs", "Id"), ObjProptype: GuildActivityGuildActivityCfg{}},
	}},
	fileInfo{"guildAuction.xlsx", []sheetInfo{
			{SheetName: "guildAuction", Initer: mapLoader("GuildAuctionGuildAuctionCfgs", "Id"), ObjProptype: GuildAuctionGuildAuctionCfg{}},
	}},
	fileInfo{"guildAutoCreate.xlsx", []sheetInfo{
			{SheetName: "guildAutoCreate", Initer: mapLoader("GuildAutoCreateGuildAutoCreateCfgs", "Id"), ObjProptype: GuildAutoCreateGuildAutoCreateCfg{}},
	}},
	fileInfo{"guildBonfire.xlsx", []sheetInfo{
			{SheetName: "guildBonfire", Initer: mapLoader("GuildBonfireGuildBonfireCfgs", "Id"), ObjProptype: GuildBonfireGuildBonfireCfg{}},
	}},
	fileInfo{"guildLevel.xlsx", []sheetInfo{
			{SheetName: "guildLevel", Initer: mapLoader("GuildLevelGuildLevelCfgs", "Id"), ObjProptype: GuildLevelGuildLevelCfg{}},
	}},
	fileInfo{"guildName.xlsx", []sheetInfo{
			{SheetName: "guildName", Initer: mapLoader("GuildNameGuildNameCfgs", "Id"), ObjProptype: GuildNameGuildNameCfg{}},
	}},
	fileInfo{"guildRobot.xlsx", []sheetInfo{
			{SheetName: "guildRobot", Initer: mapLoader("GuildRobotGuildRobotCfgs", "Id"), ObjProptype: GuildRobotGuildRobotCfg{}},
	}},
	fileInfo{"hellBoss.xlsx", []sheetInfo{
			{SheetName: "hellBoss", Initer: mapLoader("HellBossHellBossCfgs", "Id"), ObjProptype: HellBossHellBossCfg{}},
	}},
	fileInfo{"hellBossFloor.xlsx", []sheetInfo{
			{SheetName: "hellBossFloor", Initer: mapLoader("HellBossFloorHellBossFloorCfgs", "Map"), ObjProptype: HellBossFloorHellBossFloorCfg{}},
	}},
	fileInfo{"holyArms.xlsx", []sheetInfo{
			{SheetName: "holyArms", Initer: mapLoader("HolyArmsHolyArmsCfgs", "Id"), ObjProptype: HolyArmsHolyArmsCfg{}},
	}},
	fileInfo{"holyBeast.xlsx", []sheetInfo{
			{SheetName: "HolyBeast", Initer: mapLoader("HolyBeastHolyBeastCfgs", "Id"), ObjProptype: HolyBeastHolyBeastCfg{}},
	}},
	fileInfo{"holySkill.xlsx", []sheetInfo{
			{SheetName: "holySkill", Initer: mapLoader("HolySkillHolySkillCfgs", "Id"), ObjProptype: HolySkillHolySkillCfg{}},
	}},
	fileInfo{"holylevel.xlsx", []sheetInfo{
			{SheetName: "holylevel", Initer: mapLoader("HolylevelHolylevelCfgs", "Id"), ObjProptype: HolylevelHolylevelCfg{}},
	}},
	fileInfo{"hookMap.xlsx", []sheetInfo{
			{SheetName: "hookMap", Initer: mapLoader("HookMapHookMapCfgs", "Stage_id"), ObjProptype: HookMapHookMapCfg{}},
	}},
	fileInfo{"insideArt.xlsx", []sheetInfo{
			{SheetName: "insideArt", Initer: mapLoader("InsideArtInsideArtCfgs", "Id"), ObjProptype: InsideArtInsideArtCfg{}},
	}},
	fileInfo{"insideGrade.xlsx", []sheetInfo{
			{SheetName: "insideGrade", Initer: mapLoader("InsideGradeInsideGradeCfgs", "Grade"), ObjProptype: InsideGradeInsideGradeCfg{}},
	}},
	fileInfo{"insideSkill.xlsx", []sheetInfo{
			{SheetName: "insideSkill", Initer: mapLoader("InsideSkillInsideSkillCfgs", "Id"), ObjProptype: InsideSkillInsideSkillCfg{}},
	}},
	fileInfo{"insideStar.xlsx", []sheetInfo{
			{SheetName: "insideStar", Initer: mapLoader("InsideStarInsideStarCfgs", "Id"), ObjProptype: InsideStarInsideStarCfg{}},
	}},
	fileInfo{"item.xlsx", []sheetInfo{
			{SheetName: "base", Initer: mapLoader("ItemBaseCfgs", "Id"), ObjProptype: ItemBaseCfg{}},
	}},
	fileInfo{"jewel.xlsx", []sheetInfo{
			{SheetName: "jewel", Initer: mapLoader("JewelJewelCfgs", "Id"), ObjProptype: JewelJewelCfg{}},
	}},
	fileInfo{"jewelBody.xlsx", []sheetInfo{
			{SheetName: "jewelBody", Initer: mapLoader("JewelBodyJewelBodyCfgs", "Body"), ObjProptype: JewelBodyJewelBodyCfg{}},
	}},
	fileInfo{"jewelSuit.xlsx", []sheetInfo{
			{SheetName: "jewelSuit", Initer: mapLoader("JewelSuitJewelSuitCfgs", "Sum"), ObjProptype: JewelSuitJewelSuitCfg{}},
	}},
	fileInfo{"juexueLevel.xlsx", []sheetInfo{
			{SheetName: "conf", Initer: mapLoader("JuexueLevelConfCfgs", "Id"), ObjProptype: JuexueLevelConfCfg{}},
	}},
	fileInfo{"kingarms.xlsx", []sheetInfo{
			{SheetName: "kingarms", Initer: mapLoader("KingarmsKingarmsCfgs", "Id"), ObjProptype: KingarmsKingarmsCfg{}},
	}},
	fileInfo{"kuafushabakeRewardserver.xlsx", []sheetInfo{
			{SheetName: "kuafushabakeRewardserver", Initer: mapLoader("KuafushabakeRewardserverKuafushabakeRewardserverCfgs", "Id"), ObjProptype: KuafushabakeRewardserverKuafushabakeRewardserverCfg{}},
	}},
	fileInfo{"kuafushabakeRewarduni.xlsx", []sheetInfo{
			{SheetName: "kuafushabakeRewarduni", Initer: mapLoader("KuafushabakeRewarduniKuafushabakeRewarduniCfgs", "Id"), ObjProptype: KuafushabakeRewarduniKuafushabakeRewarduniCfg{}},
	}},
	fileInfo{"label.xlsx", []sheetInfo{
			{SheetName: "label", Initer: mapLoader("LabelLabelCfgs", "Id"), ObjProptype: LabelLabelCfg{}},
	}},
	fileInfo{"labelTask.xlsx", []sheetInfo{
			{SheetName: "labelTask", Initer: mapLoader("LabelTaskLabelTaskCfgs", "Id"), ObjProptype: LabelTaskLabelTaskCfg{}},
	}},
	fileInfo{"limitedGift.xlsx", []sheetInfo{
			{SheetName: "limitedGift", Initer: mapLoader("LimitedGiftLimitedGiftCfgs", "Id"), ObjProptype: LimitedGiftLimitedGiftCfg{}},
	}},
	fileInfo{"lucky.xlsx", []sheetInfo{
			{SheetName: "lucky", Initer: mapLoader("LuckyLuckyCfgs", "Id"), ObjProptype: LuckyLuckyCfg{}},
	}},
	fileInfo{"magicCircle.xlsx", []sheetInfo{
			{SheetName: "magicCircle", Initer: mapLoader("MagicCircleMagicCircleCfgs", "Id"), ObjProptype: MagicCircleMagicCircleCfg{}},
	}},
	fileInfo{"magicCircleLevel.xlsx", []sheetInfo{
			{SheetName: "magicCircleLevel", Initer: mapLoader("MagicCircleLevelMagicCircleLevelCfgs", "Id"), ObjProptype: MagicCircleLevelMagicCircleLevelCfg{}},
	}},
	fileInfo{"magicTower.xlsx", []sheetInfo{
			{SheetName: "magicTower", Initer: mapLoader("MagicTowerMagicTowerCfgs", "Id"), ObjProptype: MagicTowerMagicTowerCfg{}},
	}},
	fileInfo{"magicTowerReward.xlsx", []sheetInfo{
			{SheetName: "magicTowerReward", Initer: mapLoader("MagicTowerRewardMagicTowerRewardCfgs", "Id"), ObjProptype: MagicTowerRewardMagicTowerRewardCfg{}},
	}},
	fileInfo{"mail.xlsx", []sheetInfo{
			{SheetName: "mail", Initer: mapLoader("MailMailCfgs", "Id"), ObjProptype: MailMailCfg{}},
	}},
	fileInfo{"mainPr.xlsx", []sheetInfo{
			{SheetName: "mainPr", Initer: mapLoader("MainPrMainPrCfgs", "Id"), ObjProptype: MainPrMainPrCfg{}},
	}},
	fileInfo{"map.xlsx", []sheetInfo{
			{SheetName: "map", Initer: mapLoader("MapMapCfgs", "Id"), ObjProptype: MapMapCfg{}},
	}},
	fileInfo{"maptype.xlsx", []sheetInfo{
			{SheetName: "game", Initer: mapLoader("MaptypeGameCfgs", "Id"), ObjProptype: MaptypeGameCfg{}},
	}},
	fileInfo{"materialCost.xlsx", []sheetInfo{
			{SheetName: "materialCost", Initer: mapLoader("MaterialCostMaterialCostCfgs", "Number"), ObjProptype: MaterialCostMaterialCostCfg{}},
	}},
	fileInfo{"materialHome.xlsx", []sheetInfo{
			{SheetName: "materialHome", Initer: mapLoader("MaterialHomeMaterialHomeCfgs", "Type"), ObjProptype: MaterialHomeMaterialHomeCfg{}},
	}},
	fileInfo{"materialStage.xlsx", []sheetInfo{
			{SheetName: "materialStage", Initer: mapLoader("MaterialStageMaterialStageCfgs", "Id"), ObjProptype: MaterialStageMaterialStageCfg{}},
	}},
	fileInfo{"miji.xlsx", []sheetInfo{
			{SheetName: "miji", Initer: mapLoader("MijiMijiCfgs", "Id"), ObjProptype: MijiMijiCfg{}},
	}},
	fileInfo{"mijiLevel.xlsx", []sheetInfo{
			{SheetName: "mijiLevel", Initer: mapLoader("MijiLevelMijiLevelCfgs", "Id"), ObjProptype: MijiLevelMijiLevelCfg{}},
	}},
	fileInfo{"mijiType.xlsx", []sheetInfo{
			{SheetName: "mijiType", Initer: mapLoader("MijiTypeMijiTypeCfgs", "Id"), ObjProptype: MijiTypeMijiTypeCfg{}},
	}},
	fileInfo{"mining.xlsx", []sheetInfo{
			{SheetName: "mining", Initer: mapLoader("MiningMiningCfgs", "Id"), ObjProptype: MiningMiningCfg{}},
	}},
	fileInfo{"monster.xlsx", []sheetInfo{
			{SheetName: "monster", Initer: mapLoader("MonsterMonsterCfgs", "Monsterid"), ObjProptype: MonsterMonsterCfg{}},
	}},
	fileInfo{"monsterdrop.xlsx", []sheetInfo{
			{SheetName: "drop", Initer: mapLoader("MonsterdropDropCfgs", "Dropid"), ObjProptype: MonsterdropDropCfg{}},
	}},
	fileInfo{"monstergroup.xlsx", []sheetInfo{
			{SheetName: "monstergroup", Initer: mapLoader("MonstergroupMonstergroupCfgs", "Groupid"), ObjProptype: MonstergroupMonstergroupCfg{}},
	}},
	fileInfo{"monthCard.xlsx", []sheetInfo{
			{SheetName: "monthCard", Initer: mapLoader("MonthCardMonthCardCfgs", "Id"), ObjProptype: MonthCardMonthCardCfg{}},
	}},
	fileInfo{"monthCardPrivilege.xlsx", []sheetInfo{
			{SheetName: "monthCardPrivilege", Initer: mapLoader("MonthCardPrivilegeMonthCardPrivilegeCfgs", "Id"), ObjProptype: MonthCardPrivilegeMonthCardPrivilegeCfg{}},
	}},
	fileInfo{"npc.xlsx", []sheetInfo{
			{SheetName: "monster", Initer: mapLoader("NpcMonsterCfgs", "Id"), ObjProptype: NpcMonsterCfg{}},
	}},
	fileInfo{"official.xlsx", []sheetInfo{
			{SheetName: "official", Initer: mapLoader("OfficialOfficialCfgs", "Id"), ObjProptype: OfficialOfficialCfg{}},
	}},
	fileInfo{"openGift.xlsx", []sheetInfo{
			{SheetName: "openGift", Initer: mapLoader("OpenGiftOpenGiftCfgs", "Id"), ObjProptype: OpenGiftOpenGiftCfg{}},
	}},
	fileInfo{"panacea.xlsx", []sheetInfo{
			{SheetName: "panacea", Initer: mapLoader("PanaceaPanaceaCfgs", "Id"), ObjProptype: PanaceaPanaceaCfg{}},
	}},
	fileInfo{"personalBoss.xlsx", []sheetInfo{
			{SheetName: "personalBoss", Initer: mapLoader("PersonalBossPersonalBossCfgs", "Id"), ObjProptype: PersonalBossPersonalBossCfg{}},
	}},
	fileInfo{"pets.xlsx", []sheetInfo{
			{SheetName: "conf", Initer: mapLoader("PetsConfCfgs", "Id"), ObjProptype: PetsConfCfg{}},
	}},
	fileInfo{"petsAdd.xlsx", []sheetInfo{
			{SheetName: "petsAdd", Initer: mapLoader("PetsAddPetsAddCfgs", "Id"), ObjProptype: PetsAddPetsAddCfg{}},
	}},
	fileInfo{"petsAddSkill.xlsx", []sheetInfo{
			{SheetName: "petsAddSkill", Initer: mapLoader("PetsAddSkillPetsAddSkillCfgs", "Id"), ObjProptype: PetsAddSkillPetsAddSkillCfg{}},
	}},
	fileInfo{"petsBreak.xlsx", []sheetInfo{
			{SheetName: "conf", Initer: mapLoader("PetsBreakConfCfgs", "Id"), ObjProptype: PetsBreakConfCfg{}},
	}},
	fileInfo{"petsGrade.xlsx", []sheetInfo{
			{SheetName: "conf", Initer: mapLoader("PetsGradeConfCfgs", "Id"), ObjProptype: PetsGradeConfCfg{}},
	}},
	fileInfo{"petsLevel.xlsx", []sheetInfo{
			{SheetName: "conf", Initer: mapLoader("PetsLevelConfCfgs", "Id"), ObjProptype: PetsLevelConfCfg{}},
	}},
	fileInfo{"phantom.xlsx", []sheetInfo{
			{SheetName: "phantom", Initer: mapLoader("PhantomPhantomCfgs", "Phantom"), ObjProptype: PhantomPhantomCfg{}},
	}},
	fileInfo{"phantomLevel.xlsx", []sheetInfo{
			{SheetName: "phantomLevel", Initer: mapLoader("PhantomLevelPhantomLevelCfgs", "Id"), ObjProptype: PhantomLevelPhantomLevelCfg{}},
	}},
	fileInfo{"powerRoll.xlsx", []sheetInfo{
			{SheetName: "powerRoll", Initer: mapLoader("PowerRollPowerRollCfgs", "Id"), ObjProptype: PowerRollPowerRollCfg{}},
	}},
	fileInfo{"preFunction.xlsx", []sheetInfo{
			{SheetName: "preFunction", Initer: mapLoader("PreFunctionPreFunctionCfgs", "Id"), ObjProptype: PreFunctionPreFunctionCfg{}},
	}},
	fileInfo{"privilege.xlsx", []sheetInfo{
			{SheetName: "privilege", Initer: mapLoader("PrivilegePrivilegeCfgs", "Id"), ObjProptype: PrivilegePrivilegeCfg{}},
	}},
	fileInfo{"property.xlsx", []sheetInfo{
			{SheetName: "property", Initer: mapLoader("PropertyPropertyCfgs", "Id"), ObjProptype: PropertyPropertyCfg{}},
	}},
	fileInfo{"publicCopy.xlsx", []sheetInfo{
			{SheetName: "stage", Initer: mapLoader("PublicCopyStageCfgs", "StageId"), ObjProptype: PublicCopyStageCfg{}},
	}},
	fileInfo{"rand.xlsx", []sheetInfo{
			{SheetName: "rand", Initer: mapLoader("RandRandCfgs", "Id"), ObjProptype: RandRandCfg{}},
	}},
	fileInfo{"recharge.xlsx", []sheetInfo{
			{SheetName: "recharge", Initer: mapLoader("RechargeRechargeCfgs", "Id"), ObjProptype: RechargeRechargeCfg{}},
	}},
	fileInfo{"redDayMax.xlsx", []sheetInfo{
			{SheetName: "redDayMax", Initer: mapLoader("RedDayMaxRedDayMaxCfgs", "Id"), ObjProptype: RedDayMaxRedDayMaxCfg{}},
	}},
	fileInfo{"redRecovery.xlsx", []sheetInfo{
			{SheetName: "redRecovery", Initer: mapLoader("RedRecoveryRedRecoveryCfgs", "Id"), ObjProptype: RedRecoveryRedRecoveryCfg{}},
	}},
	fileInfo{"rein.xlsx", []sheetInfo{
			{SheetName: "rein", Initer: mapLoader("ReinReinCfgs", "Id"), ObjProptype: ReinReinCfg{}},
	}},
	fileInfo{"reinCost.xlsx", []sheetInfo{
			{SheetName: "reinCost", Initer: mapLoader("ReinCostReinCostCfgs", "Id"), ObjProptype: ReinCostReinCostCfg{}},
	}},
	fileInfo{"rewardsOnline.xlsx", []sheetInfo{
			{SheetName: "award", Initer: mapLoader("RewardsOnlineAwardCfgs", "Id"), ObjProptype: RewardsOnlineAwardCfg{}},
	}},
	fileInfo{"ring.xlsx", []sheetInfo{
			{SheetName: "ring", Initer: mapLoader("RingRingCfgs", "Ringid"), ObjProptype: RingRingCfg{}},
	}},
	fileInfo{"ringPhantom.xlsx", []sheetInfo{
			{SheetName: "ringPhantom", Initer: mapLoader("RingPhantomRingPhantomCfgs", "Id"), ObjProptype: RingPhantomRingPhantomCfg{}},
	}},
	fileInfo{"ringStrengthen.xlsx", []sheetInfo{
			{SheetName: "ringStrengthen", Initer: mapLoader("RingStrengthenRingStrengthenCfgs", "Level"), ObjProptype: RingStrengthenRingStrengthenCfg{}},
	}},
	fileInfo{"robot.xlsx", []sheetInfo{
			{SheetName: "robot", Initer: mapLoader("RobotRobotCfgs", "Id"), ObjProptype: RobotRobotCfg{}},
	}},
	fileInfo{"roleFirstname.xlsx", []sheetInfo{
			{SheetName: "roleFirstname", Initer: mapLoader("RoleFirstnameRoleFirstnameCfgs", "Id"), ObjProptype: RoleFirstnameRoleFirstnameCfg{}},
	}},
	fileInfo{"roleName.xlsx", []sheetInfo{
			{SheetName: "base", Initer: mapLoader("RoleNameBaseCfgs", "Id"), ObjProptype: RoleNameBaseCfg{}},
	}},
	fileInfo{"scrolling.xlsx", []sheetInfo{
			{SheetName: "scrolling", Initer: mapLoader("ScrollingScrollingCfgs", "Type"), ObjProptype: ScrollingScrollingCfg{}},
	}},
	fileInfo{"set.xlsx", []sheetInfo{
			{SheetName: "type", Initer: mapLoader("SetTypeCfgs", "Id"), ObjProptype: SetTypeCfg{}},
	}},
	fileInfo{"sevenDayInvest.xlsx", []sheetInfo{
			{SheetName: "sevenDayInvest", Initer: mapLoader("SevenDayInvestSevenDayInvestCfgs", "Id"), ObjProptype: SevenDayInvestSevenDayInvestCfg{}},
	}},
	fileInfo{"shabakeRewardper.xlsx", []sheetInfo{
			{SheetName: "shabakeRewardper", Initer: mapLoader("ShabakeRewardperShabakeRewardperCfgs", "Id"), ObjProptype: ShabakeRewardperShabakeRewardperCfg{}},
	}},
	fileInfo{"shabakeRewarduni.xlsx", []sheetInfo{
			{SheetName: "shabakeRewarduni", Initer: mapLoader("ShabakeRewarduniShabakeRewarduniCfgs", "Id"), ObjProptype: ShabakeRewarduniShabakeRewarduniCfg{}},
	}},
	fileInfo{"shop.xlsx", []sheetInfo{
			{SheetName: "type", Initer: mapLoader("ShopTypeCfgs", "Id"), ObjProptype: ShopTypeCfg{}},
	}},
	fileInfo{"shopItem.xlsx", []sheetInfo{
			{SheetName: "unit", Initer: mapLoader("ShopItemUnitCfgs", "Id"), ObjProptype: ShopItemUnitCfg{}},
	}},
	fileInfo{"sign.xlsx", []sheetInfo{
			{SheetName: "sign", Initer: mapLoader("SignSignCfgs", "Id"), ObjProptype: SignSignCfg{}},
	}},
	fileInfo{"skill.xlsx", []sheetInfo{
			{SheetName: "skill", Initer: mapLoader("SkillSkillCfgs", "Skillid"), ObjProptype: SkillSkillCfg{}},
	}},
	fileInfo{"skillAttackEffect.xlsx", []sheetInfo{
			{SheetName: "skillAttackEffect", Initer: mapLoader("SkillAttackEffectSkillAttackEffectCfgs", "Id"), ObjProptype: SkillAttackEffectSkillAttackEffectCfg{}},
	}},
	fileInfo{"skillLevel.xlsx", []sheetInfo{
			{SheetName: "skill", Initer: mapLoader("SkillLevelSkillCfgs", "Skillid"), ObjProptype: SkillLevelSkillCfg{}},
	}},
	fileInfo{"stage.xlsx", []sheetInfo{
			{SheetName: "stage", Initer: mapLoader("StageStageCfgs", "Id"), ObjProptype: StageStageCfg{}},
	}},
	fileInfo{"strengthen.xlsx", []sheetInfo{
			{SheetName: "strengthen", Initer: mapLoader("StrengthenStrengthenCfgs", "Id"), ObjProptype: StrengthenStrengthenCfg{}},
	}},
	fileInfo{"strengthenlink.xlsx", []sheetInfo{
			{SheetName: "strengthen", Initer: mapLoader("StrengthenlinkStrengthenCfgs", "Id"), ObjProptype: StrengthenlinkStrengthenCfg{}},
	}},
	fileInfo{"summon.xlsx", []sheetInfo{
			{SheetName: "conf", Initer: mapLoader("SummonConfCfgs", "Id"), ObjProptype: SummonConfCfg{}},
	}},
	fileInfo{"talent.xlsx", []sheetInfo{
			{SheetName: "talent", Initer: mapLoader("TalentTalentCfgs", "Id"), ObjProptype: TalentTalentCfg{}},
	}},
	fileInfo{"talentGet.xlsx", []sheetInfo{
			{SheetName: "talentGet", Initer: mapLoader("TalentGetTalentGetCfgs", "Id"), ObjProptype: TalentGetTalentGetCfg{}},
	}},
	fileInfo{"talentLevel.xlsx", []sheetInfo{
			{SheetName: "talentLevel", Initer: mapLoader("TalentLevelTalentLevelCfgs", "Id"), ObjProptype: TalentLevelTalentLevelCfg{}},
	}},
	fileInfo{"talentStage.xlsx", []sheetInfo{
			{SheetName: "talengStage", Initer: mapLoader("TalentStageTalengStageCfgs", "Id"), ObjProptype: TalentStageTalengStageCfg{}},
	}},
	fileInfo{"talentWay.xlsx", []sheetInfo{
			{SheetName: "talengWay", Initer: mapLoader("TalentWayTalengWayCfgs", "Id"), ObjProptype: TalentWayTalengWayCfg{}},
	}},
	fileInfo{"talenteffect.xlsx", []sheetInfo{
			{SheetName: "talent", Initer: mapLoader("TalenteffectTalentCfgs", "Id"), ObjProptype: TalenteffectTalentCfg{}},
	}},
	fileInfo{"talentgeneral.xlsx", []sheetInfo{
			{SheetName: "talent", Initer: mapLoader("TalentgeneralTalentCfgs", "Id"), ObjProptype: TalentgeneralTalentCfg{}},
	}},
	fileInfo{"task.xlsx", []sheetInfo{
			{SheetName: "condition", Initer: mapLoader("TaskConditionCfgs", "Id"), ObjProptype: TaskConditionCfg{}},
	}},
	fileInfo{"title.xlsx", []sheetInfo{
			{SheetName: "title", Initer: mapLoader("TitleTitleCfgs", "Id"), ObjProptype: TitleTitleCfg{}},
	}},
	fileInfo{"tower.xlsx", []sheetInfo{
			{SheetName: "tower", Initer: mapLoader("TowerTowerCfgs", "Id"), ObjProptype: TowerTowerCfg{}},
	}},
	fileInfo{"towerLotteryCircle.xlsx", []sheetInfo{
			{SheetName: "towerLotteryCircle", Initer: mapLoader("TowerLotteryCircleTowerLotteryCircleCfgs", "Id"), ObjProptype: TowerLotteryCircleTowerLotteryCircleCfg{}},
	}},
	fileInfo{"towerRankReward.xlsx", []sheetInfo{
			{SheetName: "towerRankReward", Initer: mapLoader("TowerRankRewardTowerRankRewardCfgs", "Id"), ObjProptype: TowerRankRewardTowerRankRewardCfg{}},
	}},
	fileInfo{"towerReward.xlsx", []sheetInfo{
			{SheetName: "towerReward", Initer: mapLoader("TowerRewardTowerRewardCfgs", "Id"), ObjProptype: TowerRewardTowerRewardCfg{}},
	}},
	fileInfo{"treasure.xlsx", []sheetInfo{
			{SheetName: "treasure", Initer: mapLoader("TreasureTreasureCfgs", "Id"), ObjProptype: TreasureTreasureCfg{}},
	}},
	fileInfo{"treasureArt.xlsx", []sheetInfo{
			{SheetName: "treasureArt", Initer: mapLoader("TreasureArtTreasureArtCfgs", "Id"), ObjProptype: TreasureArtTreasureArtCfg{}},
	}},
	fileInfo{"treasureAwaken.xlsx", []sheetInfo{
			{SheetName: "treasureAwaken", Initer: mapLoader("TreasureAwakenTreasureAwakenCfgs", "Id"), ObjProptype: TreasureAwakenTreasureAwakenCfg{}},
	}},
	fileInfo{"treasureDiscount.xlsx", []sheetInfo{
			{SheetName: "treasureDiscount", Initer: mapLoader("TreasureDiscountTreasureDiscountCfgs", "Id"), ObjProptype: TreasureDiscountTreasureDiscountCfg{}},
	}},
	fileInfo{"treasureShop.xlsx", []sheetInfo{
			{SheetName: "treasureShop", Initer: mapLoader("TreasureShopTreasureShopCfgs", "Id"), ObjProptype: TreasureShopTreasureShopCfg{}},
	}},
	fileInfo{"treasureStars.xlsx", []sheetInfo{
			{SheetName: "treasureStars", Initer: mapLoader("TreasureStarsTreasureStarsCfgs", "Id"), ObjProptype: TreasureStarsTreasureStarsCfg{}},
	}},
	fileInfo{"treasureSuit.xlsx", []sheetInfo{
			{SheetName: "treasureSuit", Initer: mapLoader("TreasureSuitTreasureSuitCfgs", "Id"), ObjProptype: TreasureSuitTreasureSuitCfg{}},
	}},
	fileInfo{"trialTask.xlsx", []sheetInfo{
			{SheetName: "trialTask", Initer: mapLoader("TrialTaskTrialTaskCfgs", "Id"), ObjProptype: TrialTaskTrialTaskCfg{}},
	}},
	fileInfo{"trialTotalReward.xlsx", []sheetInfo{
			{SheetName: "trialTotalReward", Initer: mapLoader("TrialTotalRewardTrialTotalRewardCfgs", "Id"), ObjProptype: TrialTotalRewardTrialTotalRewardCfg{}},
	}},
	fileInfo{"vip.xlsx", []sheetInfo{
			{SheetName: "lvl", Initer: mapLoader("VipLvlCfgs", "Lvl"), ObjProptype: VipLvlCfg{}},
	}},
	fileInfo{"vipBoss.xlsx", []sheetInfo{
			{SheetName: "vipBoss", Initer: mapLoader("VipBossVipBossCfgs", "Id"), ObjProptype: VipBossVipBossCfg{}},
	}},
	fileInfo{"warOrderCondition.xlsx", []sheetInfo{
			{SheetName: "warOrderCondition", Initer: mapLoader("WarOrderConditionWarOrderConditionCfgs", "Id"), ObjProptype: WarOrderConditionWarOrderConditionCfg{}},
	}},
	fileInfo{"warOrderCycle.xlsx", []sheetInfo{
			{SheetName: "warOrderCycle", Initer: mapLoader("WarOrderCycleWarOrderCycleCfgs", "Id"), ObjProptype: WarOrderCycleWarOrderCycleCfg{}},
	}},
	fileInfo{"warOrderCycleTask.xlsx", []sheetInfo{
			{SheetName: "warOrderCycleTask", Initer: mapLoader("WarOrderCycleTaskWarOrderCycleTaskCfgs", "Id"), ObjProptype: WarOrderCycleTaskWarOrderCycleTaskCfg{}},
	}},
	fileInfo{"warOrderExchange.xlsx", []sheetInfo{
			{SheetName: "warOrderExchange", Initer: mapLoader("WarOrderExchangeWarOrderExchangeCfgs", "Id"), ObjProptype: WarOrderExchangeWarOrderExchangeCfg{}},
	}},
	fileInfo{"warOrderLevel.xlsx", []sheetInfo{
			{SheetName: "warOrderLevel", Initer: mapLoader("WarOrderLevelWarOrderLevelCfgs", "Id"), ObjProptype: WarOrderLevelWarOrderLevelCfg{}},
	}},
	fileInfo{"warOrderWeekTask.xlsx", []sheetInfo{
			{SheetName: "warOrderWeekTask", Initer: mapLoader("WarOrderWeekTaskWarOrderWeekTaskCfgs", "Id"), ObjProptype: WarOrderWeekTaskWarOrderWeekTaskCfg{}},
	}},
	fileInfo{"wash.xlsx", []sheetInfo{
			{SheetName: "wash", Initer: mapLoader("WashWashCfgs", "Id"), ObjProptype: WashWashCfg{}},
	}},
	fileInfo{"washrand.xlsx", []sheetInfo{
			{SheetName: "rand", Initer: mapLoader("WashrandRandCfgs", "Id"), ObjProptype: WashrandRandCfg{}},
	}},
	fileInfo{"wingNew.xlsx", []sheetInfo{
			{SheetName: "wingNew", Initer: mapLoader("WingNewWingNewCfgs", "Id"), ObjProptype: WingNewWingNewCfg{}},
	}},
	fileInfo{"wingSpecial.xlsx", []sheetInfo{
			{SheetName: "wingSpecial", Initer: mapLoader("WingSpecialWingSpecialCfgs", "Id"), ObjProptype: WingSpecialWingSpecialCfg{}},
	}},
	fileInfo{"worldBoss.xlsx", []sheetInfo{
			{SheetName: "worldBoss", Initer: mapLoader("WorldBossWorldBossCfgs", "Id"), ObjProptype: WorldBossWorldBossCfg{}},
	}},
	fileInfo{"worldLeader.xlsx", []sheetInfo{
			{SheetName: "conf", Initer: mapLoader("WorldLeaderConfCfgs", "Id"), ObjProptype: WorldLeaderConfCfg{}},
	}},
	fileInfo{"worldLeaderReward.xlsx", []sheetInfo{
			{SheetName: "worldLeaderReward", Initer: mapLoader("WorldLeaderRewardWorldLeaderRewardCfgs", "Id"), ObjProptype: WorldLeaderRewardWorldLeaderRewardCfg{}},
	}},
	fileInfo{"worldLevel.xlsx", []sheetInfo{
			{SheetName: "worldLevel", Initer: mapLoader("WorldLevelWorldLevelCfgs", "Id"), ObjProptype: WorldLevelWorldLevelCfg{}},
	}},
	fileInfo{"worldLevelBuff.xlsx", []sheetInfo{
			{SheetName: "worldLevelBuff", Initer: mapLoader("WorldLevelBuffWorldLevelBuffCfgs", "Id"), ObjProptype: WorldLevelBuffWorldLevelBuffCfg{}},
	}},
	fileInfo{"worldRank.xlsx", []sheetInfo{
			{SheetName: "worldRank", Initer: mapLoader("WorldRankWorldRankCfgs", "Rank"), ObjProptype: WorldRankWorldRankCfg{}},
	}},
	fileInfo{"xiaoyouxiTower.xlsx", []sheetInfo{
			{SheetName: "xiaoyouxiTower", Initer: mapLoader("XiaoyouxiTowerXiaoyouxiTowerCfgs", "Id"), ObjProptype: XiaoyouxiTowerXiaoyouxiTowerCfg{}},
	}},
	fileInfo{"xunlong.xlsx", []sheetInfo{
			{SheetName: "xunlong", Initer: mapLoader("XunlongXunlongCfgs", "Id"), ObjProptype: XunlongXunlongCfg{}},
	}},
	fileInfo{"xunlongPr.xlsx", []sheetInfo{
			{SheetName: "xunlongPr", Initer: mapLoader("XunlongPrXunlongPrCfgs", "Time"), ObjProptype: XunlongPrXunlongPrCfg{}},
	}},
	fileInfo{"xunlongRounds.xlsx", []sheetInfo{
			{SheetName: "xunlongRounds", Initer: mapLoader("XunlongRoundsXunlongRoundsCfgs", "Id"), ObjProptype: XunlongRoundsXunlongRoundsCfg{}},
	}},
	fileInfo{"zodiacEquip.xlsx", []sheetInfo{
			{SheetName: "zodiacEquip", Initer: mapLoader("ZodiacEquipZodiacEquipCfgs", "Id"), ObjProptype: ZodiacEquipZodiacEquipCfg{}},
	}},
}


type GameDbBase struct {
	Ver         string
	FileModTime map[string]int64

	//NOTE 关于client的配置：
	//client:对象名,对象类型　，对象名要小写．
	// mapKey 即对应的我们的结构里的key, 要看具体的型中key是什么　，一段是大写的
	InitConf         *InitConf
	
    PaoDianRewardPaoDianRewardCfgs		map[int]*PaoDianRewardPaoDianRewardCfg
    SpendrebatesSpendrebatesCfgs		map[int]*SpendrebatesSpendrebatesCfg
    AccumulateAccumulateCfgs		map[int]*AccumulateAccumulateCfg
    AchievementAchievementCfgs		map[int]*AchievementAchievementCfg
    AchievementMedalMedalCfgs		map[int]*AchievementMedalMedalCfg
    AncientBossAncientBossCfgs		map[int]*AncientBossAncientBossCfg
    AncientSkillGradeAncientSkillGradeCfgs		map[int]*AncientSkillGradeAncientSkillGradeCfg
    AncientSkillLevelAncientSkillLevelCfgs		map[int]*AncientSkillLevelAncientSkillLevelCfg
    ArcherMagicCfgs		map[int]*ArcherMagicCfg
    ArcherElementMagicElementCfgs		map[int]*ArcherElementMagicElementCfg
    AreaAreaCfgs		map[int]*AreaAreaCfg
    AreaLevelAreaLevelCfgs		map[int]*AreaLevelAreaLevelCfg
    ArenaBuyArenaBuyCfgs		map[int]*ArenaBuyArenaBuyCfg
    ArenaMatchArenaMatchCfgs		map[int]*ArenaMatchArenaMatchCfg
    ArenaRankArenaRankCfgs		map[int]*ArenaRankArenaRankCfg
    AspdAspdCfgs		map[int]*AspdAspdCfg
    AtlasAtlasCfgs		map[int]*AtlasAtlasCfg
    AtlasGatherAtlasGatherCfgs		map[int]*AtlasGatherAtlasGatherCfg
    AtlasPosAtlasPosCfgs		map[int]*AtlasPosAtlasPosCfg
    AtlasStarAtlasStarCfgs		map[int]*AtlasStarAtlasStarCfg
    AtlasUpgradeAtlasUpgradeCfgs		map[int]*AtlasUpgradeAtlasUpgradeCfg
    AttackEnemyAttackEnemyCfgs		map[int]*AttackEnemyAttackEnemyCfg
    AttackEnemyCardAttackEnemyCardCfgs		map[int]*AttackEnemyCardAttackEnemyCardCfg
    AuctionAuctioinCfgs		map[int]*AuctionAuctioinCfg
    AwakenAwakenCfgs		map[int]*AwakenAwakenCfg
    AwakenTitleAwakenTitleCfgs		map[int]*AwakenTitleAwakenTitleCfg
    BagSpaceAddCfgs		map[int]*BagSpaceAddCfg
    BindGroupBindGroupCfgs		map[int]*BindGroupBindGroupCfg
    BlessBlessCfgs		map[int]*BlessBlessCfg
    BossFamilyBossFamilyCfgs		map[int]*BossFamilyBossFamilyCfg
    BuffBuffCfgs		map[int]*BuffBuffCfg
    ChatClearCfgs		map[int]*ChatClearCfg
    ChuanShiEquipChuanShiEquipCfgs		map[int]*ChuanShiEquipChuanShiEquipCfg
    ChuanShiEquipTypeChuanShiEquipTypeCfgs		map[int]*ChuanShiEquipTypeChuanShiEquipTypeCfg
    ChuanShiStrengthenChuanShiStrengthenCfgs		map[int]*ChuanShiStrengthenChuanShiStrengthenCfg
    ChuanShiStrengthenLinkChuanShiStrengthenLinkCfgs		map[int]*ChuanShiStrengthenLinkChuanShiStrengthenLinkCfg
    ChuanShiSuitChuanShiSuitCfgs		map[int]*ChuanShiSuitChuanShiSuitCfg
    ChuanShiSuitTypeChuanShiSuitTypeCfgs		map[int]*ChuanShiSuitTypeChuanShiSuitTypeCfg
    ClearClearCfgs		map[int]*ClearClearCfg
    ClearRateClearRateCfgs		map[int]*ClearRateClearRateCfg
    CollectionCollectionCfgs		map[int]*CollectionCollectionCfg
    CompetitveCompetitveCfgs		map[int]*CompetitveCompetitveCfg
    CompetitveRewardRankRewardCfgs		map[int]*CompetitveRewardRankRewardCfg
    ComposeChuanShiSubComposeChuanShiSubCfgs		map[int]*ComposeChuanShiSubComposeChuanShiSubCfg
    ComposeChuanShiTypeComposeChuanShiTypeCfgs		map[int]*ComposeChuanShiTypeComposeChuanShiTypeCfg
    ComposeEquipSubComposeEquipSubCfgs		map[int]*ComposeEquipSubComposeEquipSubCfg
    ComposeEquipTypeComposeEquipTypeCfgs		map[int]*ComposeEquipTypeComposeEquipTypeCfg
    ComposeSubComposeSubCfgs		map[int]*ComposeSubComposeSubCfg
    ComposeTypeComposeTypeCfgs		map[int]*ComposeTypeComposeTypeCfg
    ConditionConditionCfgs		map[int]*ConditionConditionCfg
    ContRechargeContRechargeCfgs		map[int]*ContRechargeContRechargeCfg
    CrossArenaCrossArenaCfgs		map[int]*CrossArenaCrossArenaCfg
    CrossArenaRewardCrossArenaRewardCfgs		map[int]*CrossArenaRewardCrossArenaRewardCfg
    CrossArenaRobotCrossArenaRobotCfgs		map[int]*CrossArenaRobotCrossArenaRobotCfg
    CrossArenaTimeCrossArenaTimeCfgs		map[int]*CrossArenaTimeCrossArenaTimeCfg
    CumulationsignCumulationsignCfgs		map[int]*CumulationsignCumulationsignCfg
    CutTreasureCutTreasureCfgs		map[int]*CutTreasureCutTreasureCfg
    DaBaoEquipDaBaoEquipCfgs		map[int]*DaBaoEquipDaBaoEquipCfg
    DaBaoEquipAdditionDaBaoEquipAdditionCfgs		map[int]*DaBaoEquipAdditionDaBaoEquipAdditionCfg
    DaBaoMysteryDaBaoMysteryCfgs		map[int]*DaBaoMysteryDaBaoMysteryCfg
    DailyActivityDailyActivityCfgs		map[int]*DailyActivityDailyActivityCfg
    DailyRewardDailyRewardCfgs		map[int]*DailyRewardDailyRewardCfg
    DailyTaskDailytaskCfgs		map[int]*DailyTaskDailytaskCfg
    DailypackDailypackCfgs		map[int]*DailypackDailypackCfg
    DarkPalaceBossDarkPalaceBossCfgs		map[int]*DarkPalaceBossDarkPalaceBossCfg
    DayRankingDayRankingCfgs		map[int]*DayRankingDayRankingCfg
    DayRankingGiftDayRankingGiftCfgs		map[int]*DayRankingGiftDayRankingGiftCfg
    DayRankingMarkDayRankingMarkCfgs		map[int]*DayRankingMarkDayRankingMarkCfg
    DayRankingRewardDayRankingRewardCfgs		map[int]*DayRankingRewardDayRankingRewardCfg
    DictateEquipDictateEquipCfgs		map[int]*DictateEquipDictateEquipCfg
    DictateSuitDictateSuitCfgs		map[int]*DictateSuitDictateSuitCfg
    DragonEquipDragonEquipCfgs		map[int]*DragonEquipDragonEquipCfg
    DragonEquipLevelDragonEquipLevelCfgs		map[int]*DragonEquipLevelDragonEquipLevelCfg
    DragonarmsDragonarmsCfgs		map[int]*DragonarmsDragonarmsCfg
    DrawDrawCfgs		map[int]*DrawDrawCfg
    DrawShopDrawShopCfgs		map[int]*DrawShopDrawShopCfg
    DropDropCfgs		map[int]*DropDropCfg
    DropSpecialDropSpecialCfgs		map[int]*DropSpecialDropSpecialCfg
    EffectEffectCfgs		map[int]*EffectEffectCfg
    ElfGrowElfGrowCfgs		map[int]*ElfGrowElfGrowCfg
    ElfRecoverElfRecoverCfgs		map[int]*ElfRecoverElfRecoverCfg
    ElfSkillElfGrowCfgs		map[int]*ElfSkillElfGrowCfg
    EquipEquipCfgs		map[int]*EquipEquipCfg
    EquipsuitEquipsuitCfgs		map[int]*EquipsuitEquipsuitCfg
    ExpLevelLevelCfgs		map[int]*ExpLevelLevelCfg
    ExpPillExpPillCfgs		map[int]*ExpPillExpPillCfg
    ExpPoolExpPoolCfgs		map[int]*ExpPoolExpPoolCfg
    ExpStageExpStageCfgs		map[int]*ExpStageExpStageCfg
    FabaoFabaoCfgs		map[int]*FabaoFabaoCfg
    FabaoSkillFabaoSkillCfgs		map[int]*FabaoSkillFabaoSkillCfg
    FabaolevelFabaolevelCfgs		map[int]*FabaolevelFabaolevelCfg
    FashionFashionCfgs		map[int]*FashionFashionCfg
    FieldBossFieldBossCfgs		map[int]*FieldBossFieldBossCfg
    FieldFightFieldBaseCfgs		map[int]*FieldFightFieldBaseCfg
    FirstBloodPerFirstBloodperCfgs		map[int]*FirstBloodPerFirstBloodperCfg
    FirstBloodmilFirstBloodmilCfgs		map[int]*FirstBloodmilFirstBloodmilCfg
    FirstBlooduniFirstBlooduniCfgs		map[int]*FirstBlooduniFirstBlooduniCfg
    FirstDropFirstDropCfgs		map[int]*FirstDropFirstDropCfg
    FirstRechargTypeFirstRechargTypeCfgs		map[int]*FirstRechargTypeFirstRechargTypeCfg
    FirstRechargeFirstRechargCfgs		map[int]*FirstRechargeFirstRechargCfg
    FitFashionFitFashionCfgs		map[int]*FitFashionFitFashionCfg
    FitFashionLevelFitFashionLevelCfgs		map[int]*FitFashionLevelFitFashionLevelCfg
    FitHolyEquipFitHolyEquipCfgs		map[int]*FitHolyEquipFitHolyEquipCfg
    FitHolyEquipSuitFitHolyEquipSuitCfgs		map[int]*FitHolyEquipSuitFitHolyEquipSuitCfg
    FitLevelFitLevelCfgs		map[int]*FitLevelFitLevelCfg
    FitSkillFitSkillCfgs		map[int]*FitSkillFitSkillCfg
    FitSkillLevelFitSkillLevelCfgs		map[int]*FitSkillLevelFitSkillLevelCfg
    FitSkillSlotFitSkillSlotCfgs		map[int]*FitSkillSlotFitSkillSlotCfg
    FitSkillStarFitSkillStarCfgs		map[int]*FitSkillStarFitSkillStarCfg
    FunctionFunctionCfgs		map[int]*FunctionFunctionCfg
    GameDotGameDotCfgs		map[int]*GameDotGameDotCfg
    GameTextErrorTextCfgs		map[int]*GameTextErrorTextCfg
    GameTextCodeTextCfgs		map[int]*GameTextCodeTextCfg
    GamewordGameCfgs		map[int]*GamewordGameCfg
    GiftGiftCfgs		map[int]*GiftGiftCfg
    GiftCodeGiftCodeCfgs		map[int]*GiftCodeGiftCodeCfg
    GodBloodGodBloodCfgs		map[int]*GodBloodGodBloodCfg
    GodEquipConfCfgs		map[int]*GodEquipConfCfg
    GodEquipLevelConfCfgs		map[int]*GodEquipLevelConfCfg
    GodEquipSuitConfCfgs		map[int]*GodEquipSuitConfCfg
    GrowFundGrowFundCfgs		map[int]*GrowFundGrowFundCfg
    GuardRankGuardRankCfgs		map[int]*GuardRankGuardRankCfg
    GuardRoundsGuardRoundsCfgs		map[int]*GuardRoundsGuardRoundsCfg
    GuildGuildCfgs		map[int]*GuildGuildCfg
    GuildActivityGuildActivityCfgs		map[int]*GuildActivityGuildActivityCfg
    GuildAuctionGuildAuctionCfgs		map[int]*GuildAuctionGuildAuctionCfg
    GuildAutoCreateGuildAutoCreateCfgs		map[int]*GuildAutoCreateGuildAutoCreateCfg
    GuildBonfireGuildBonfireCfgs		map[int]*GuildBonfireGuildBonfireCfg
    GuildLevelGuildLevelCfgs		map[int]*GuildLevelGuildLevelCfg
    GuildNameGuildNameCfgs		map[int]*GuildNameGuildNameCfg
    GuildRobotGuildRobotCfgs		map[int]*GuildRobotGuildRobotCfg
    HellBossHellBossCfgs		map[int]*HellBossHellBossCfg
    HellBossFloorHellBossFloorCfgs		map[int]*HellBossFloorHellBossFloorCfg
    HolyArmsHolyArmsCfgs		map[int]*HolyArmsHolyArmsCfg
    HolyBeastHolyBeastCfgs		map[int]*HolyBeastHolyBeastCfg
    HolySkillHolySkillCfgs		map[int]*HolySkillHolySkillCfg
    HolylevelHolylevelCfgs		map[int]*HolylevelHolylevelCfg
    HookMapHookMapCfgs		map[int]*HookMapHookMapCfg
    InsideArtInsideArtCfgs		map[int]*InsideArtInsideArtCfg
    InsideGradeInsideGradeCfgs		map[int]*InsideGradeInsideGradeCfg
    InsideSkillInsideSkillCfgs		map[int]*InsideSkillInsideSkillCfg
    InsideStarInsideStarCfgs		map[int]*InsideStarInsideStarCfg
    ItemBaseCfgs		map[int]*ItemBaseCfg
    JewelJewelCfgs		map[int]*JewelJewelCfg
    JewelBodyJewelBodyCfgs		map[int]*JewelBodyJewelBodyCfg
    JewelSuitJewelSuitCfgs		map[int]*JewelSuitJewelSuitCfg
    JuexueLevelConfCfgs		map[int]*JuexueLevelConfCfg
    KingarmsKingarmsCfgs		map[int]*KingarmsKingarmsCfg
    KuafushabakeRewardserverKuafushabakeRewardserverCfgs		map[int]*KuafushabakeRewardserverKuafushabakeRewardserverCfg
    KuafushabakeRewarduniKuafushabakeRewarduniCfgs		map[int]*KuafushabakeRewarduniKuafushabakeRewarduniCfg
    LabelLabelCfgs		map[int]*LabelLabelCfg
    LabelTaskLabelTaskCfgs		map[int]*LabelTaskLabelTaskCfg
    LimitedGiftLimitedGiftCfgs		map[int]*LimitedGiftLimitedGiftCfg
    LuckyLuckyCfgs		map[int]*LuckyLuckyCfg
    MagicCircleMagicCircleCfgs		map[int]*MagicCircleMagicCircleCfg
    MagicCircleLevelMagicCircleLevelCfgs		map[int]*MagicCircleLevelMagicCircleLevelCfg
    MagicTowerMagicTowerCfgs		map[int]*MagicTowerMagicTowerCfg
    MagicTowerRewardMagicTowerRewardCfgs		map[int]*MagicTowerRewardMagicTowerRewardCfg
    MailMailCfgs		map[int]*MailMailCfg
    MainPrMainPrCfgs		map[int]*MainPrMainPrCfg
    MapMapCfgs		map[int]*MapMapCfg
    MaptypeGameCfgs		map[int]*MaptypeGameCfg
    MaterialCostMaterialCostCfgs		map[int]*MaterialCostMaterialCostCfg
    MaterialHomeMaterialHomeCfgs		map[int]*MaterialHomeMaterialHomeCfg
    MaterialStageMaterialStageCfgs		map[int]*MaterialStageMaterialStageCfg
    MijiMijiCfgs		map[int]*MijiMijiCfg
    MijiLevelMijiLevelCfgs		map[int]*MijiLevelMijiLevelCfg
    MijiTypeMijiTypeCfgs		map[int]*MijiTypeMijiTypeCfg
    MiningMiningCfgs		map[int]*MiningMiningCfg
    MonsterMonsterCfgs		map[int]*MonsterMonsterCfg
    MonsterdropDropCfgs		map[int]*MonsterdropDropCfg
    MonstergroupMonstergroupCfgs		map[int]*MonstergroupMonstergroupCfg
    MonthCardMonthCardCfgs		map[int]*MonthCardMonthCardCfg
    MonthCardPrivilegeMonthCardPrivilegeCfgs		map[int]*MonthCardPrivilegeMonthCardPrivilegeCfg
    NpcMonsterCfgs		map[int]*NpcMonsterCfg
    OfficialOfficialCfgs		map[int]*OfficialOfficialCfg
    OpenGiftOpenGiftCfgs		map[int]*OpenGiftOpenGiftCfg
    PanaceaPanaceaCfgs		map[int]*PanaceaPanaceaCfg
    PersonalBossPersonalBossCfgs		map[int]*PersonalBossPersonalBossCfg
    PetsConfCfgs		map[int]*PetsConfCfg
    PetsAddPetsAddCfgs		map[int]*PetsAddPetsAddCfg
    PetsAddSkillPetsAddSkillCfgs		map[int]*PetsAddSkillPetsAddSkillCfg
    PetsBreakConfCfgs		map[int]*PetsBreakConfCfg
    PetsGradeConfCfgs		map[int]*PetsGradeConfCfg
    PetsLevelConfCfgs		map[int]*PetsLevelConfCfg
    PhantomPhantomCfgs		map[int]*PhantomPhantomCfg
    PhantomLevelPhantomLevelCfgs		map[int]*PhantomLevelPhantomLevelCfg
    PowerRollPowerRollCfgs		map[int]*PowerRollPowerRollCfg
    PreFunctionPreFunctionCfgs		map[int]*PreFunctionPreFunctionCfg
    PrivilegePrivilegeCfgs		map[int]*PrivilegePrivilegeCfg
    PropertyPropertyCfgs		map[int]*PropertyPropertyCfg
    PublicCopyStageCfgs		map[int]*PublicCopyStageCfg
    RandRandCfgs		map[int]*RandRandCfg
    RechargeRechargeCfgs		map[int]*RechargeRechargeCfg
    RedDayMaxRedDayMaxCfgs		map[int]*RedDayMaxRedDayMaxCfg
    RedRecoveryRedRecoveryCfgs		map[int]*RedRecoveryRedRecoveryCfg
    ReinReinCfgs		map[int]*ReinReinCfg
    ReinCostReinCostCfgs		map[int]*ReinCostReinCostCfg
    RewardsOnlineAwardCfgs		map[int]*RewardsOnlineAwardCfg
    RingRingCfgs		map[int]*RingRingCfg
    RingPhantomRingPhantomCfgs		map[int]*RingPhantomRingPhantomCfg
    RingStrengthenRingStrengthenCfgs		map[int]*RingStrengthenRingStrengthenCfg
    RobotRobotCfgs		map[int]*RobotRobotCfg
    RoleFirstnameRoleFirstnameCfgs		map[int]*RoleFirstnameRoleFirstnameCfg
    RoleNameBaseCfgs		map[int]*RoleNameBaseCfg
    ScrollingScrollingCfgs		map[int]*ScrollingScrollingCfg
    SetTypeCfgs		map[int]*SetTypeCfg
    SevenDayInvestSevenDayInvestCfgs		map[int]*SevenDayInvestSevenDayInvestCfg
    ShabakeRewardperShabakeRewardperCfgs		map[int]*ShabakeRewardperShabakeRewardperCfg
    ShabakeRewarduniShabakeRewarduniCfgs		map[int]*ShabakeRewarduniShabakeRewarduniCfg
    ShopTypeCfgs		map[int]*ShopTypeCfg
    ShopItemUnitCfgs		map[int]*ShopItemUnitCfg
    SignSignCfgs		map[int]*SignSignCfg
    SkillSkillCfgs		map[int]*SkillSkillCfg
    SkillAttackEffectSkillAttackEffectCfgs		map[int]*SkillAttackEffectSkillAttackEffectCfg
    SkillLevelSkillCfgs		map[int]*SkillLevelSkillCfg
    StageStageCfgs		map[int]*StageStageCfg
    StrengthenStrengthenCfgs		map[int]*StrengthenStrengthenCfg
    StrengthenlinkStrengthenCfgs		map[int]*StrengthenlinkStrengthenCfg
    SummonConfCfgs		map[int]*SummonConfCfg
    TalentTalentCfgs		map[int]*TalentTalentCfg
    TalentGetTalentGetCfgs		map[int]*TalentGetTalentGetCfg
    TalentLevelTalentLevelCfgs		map[int]*TalentLevelTalentLevelCfg
    TalentStageTalengStageCfgs		map[int]*TalentStageTalengStageCfg
    TalentWayTalengWayCfgs		map[int]*TalentWayTalengWayCfg
    TalenteffectTalentCfgs		map[int]*TalenteffectTalentCfg
    TalentgeneralTalentCfgs		map[int]*TalentgeneralTalentCfg
    TaskConditionCfgs		map[int]*TaskConditionCfg
    TitleTitleCfgs		map[int]*TitleTitleCfg
    TowerTowerCfgs		map[int]*TowerTowerCfg
    TowerLotteryCircleTowerLotteryCircleCfgs		map[int]*TowerLotteryCircleTowerLotteryCircleCfg
    TowerRankRewardTowerRankRewardCfgs		map[int]*TowerRankRewardTowerRankRewardCfg
    TowerRewardTowerRewardCfgs		map[int]*TowerRewardTowerRewardCfg
    TreasureTreasureCfgs		map[int]*TreasureTreasureCfg
    TreasureArtTreasureArtCfgs		map[int]*TreasureArtTreasureArtCfg
    TreasureAwakenTreasureAwakenCfgs		map[int]*TreasureAwakenTreasureAwakenCfg
    TreasureDiscountTreasureDiscountCfgs		map[int]*TreasureDiscountTreasureDiscountCfg
    TreasureShopTreasureShopCfgs		map[int]*TreasureShopTreasureShopCfg
    TreasureStarsTreasureStarsCfgs		map[int]*TreasureStarsTreasureStarsCfg
    TreasureSuitTreasureSuitCfgs		map[int]*TreasureSuitTreasureSuitCfg
    TrialTaskTrialTaskCfgs		map[int]*TrialTaskTrialTaskCfg
    TrialTotalRewardTrialTotalRewardCfgs		map[int]*TrialTotalRewardTrialTotalRewardCfg
    VipLvlCfgs		map[int]*VipLvlCfg
    VipBossVipBossCfgs		map[int]*VipBossVipBossCfg
    WarOrderConditionWarOrderConditionCfgs		map[int]*WarOrderConditionWarOrderConditionCfg
    WarOrderCycleWarOrderCycleCfgs		map[int]*WarOrderCycleWarOrderCycleCfg
    WarOrderCycleTaskWarOrderCycleTaskCfgs		map[int]*WarOrderCycleTaskWarOrderCycleTaskCfg
    WarOrderExchangeWarOrderExchangeCfgs		map[int]*WarOrderExchangeWarOrderExchangeCfg
    WarOrderLevelWarOrderLevelCfgs		map[int]*WarOrderLevelWarOrderLevelCfg
    WarOrderWeekTaskWarOrderWeekTaskCfgs		map[int]*WarOrderWeekTaskWarOrderWeekTaskCfg
    WashWashCfgs		map[int]*WashWashCfg
    WashrandRandCfgs		map[int]*WashrandRandCfg
    WingNewWingNewCfgs		map[int]*WingNewWingNewCfg
    WingSpecialWingSpecialCfgs		map[int]*WingSpecialWingSpecialCfg
    WorldBossWorldBossCfgs		map[int]*WorldBossWorldBossCfg
    WorldLeaderConfCfgs		map[int]*WorldLeaderConfCfg
    WorldLeaderRewardWorldLeaderRewardCfgs		map[int]*WorldLeaderRewardWorldLeaderRewardCfg
    WorldLevelWorldLevelCfgs		map[int]*WorldLevelWorldLevelCfg
    WorldLevelBuffWorldLevelBuffCfgs		map[int]*WorldLevelBuffWorldLevelBuffCfg
    WorldRankWorldRankCfgs		map[int]*WorldRankWorldRankCfg
    XiaoyouxiTowerXiaoyouxiTowerCfgs		map[int]*XiaoyouxiTowerXiaoyouxiTowerCfg
    XunlongXunlongCfgs		map[int]*XunlongXunlongCfg
    XunlongPrXunlongPrCfgs		map[int]*XunlongPrXunlongPrCfg
    XunlongRoundsXunlongRoundsCfgs		map[int]*XunlongRoundsXunlongRoundsCfg
    ZodiacEquipZodiacEquipCfgs		map[int]*ZodiacEquipZodiacEquipCfg
}	


func GetPaoDianRewardPaoDianRewardCfg( Id int) *PaoDianRewardPaoDianRewardCfg {
	return gameDb.PaoDianRewardPaoDianRewardCfgs[Id]
}

func RangPaoDianRewardPaoDianRewardCfgs(f func(conf *PaoDianRewardPaoDianRewardCfg)bool){
	for _,v := range gameDb.PaoDianRewardPaoDianRewardCfgs{
		if !f(v){
			return
		}
	}
}

func GetSpendrebatesSpendrebatesCfg( Id int) *SpendrebatesSpendrebatesCfg {
	return gameDb.SpendrebatesSpendrebatesCfgs[Id]
}

func RangSpendrebatesSpendrebatesCfgs(f func(conf *SpendrebatesSpendrebatesCfg)bool){
	for _,v := range gameDb.SpendrebatesSpendrebatesCfgs{
		if !f(v){
			return
		}
	}
}

func GetAccumulateAccumulateCfg( Id int) *AccumulateAccumulateCfg {
	return gameDb.AccumulateAccumulateCfgs[Id]
}

func RangAccumulateAccumulateCfgs(f func(conf *AccumulateAccumulateCfg)bool){
	for _,v := range gameDb.AccumulateAccumulateCfgs{
		if !f(v){
			return
		}
	}
}

func GetAchievementAchievementCfg( Id int) *AchievementAchievementCfg {
	return gameDb.AchievementAchievementCfgs[Id]
}

func RangAchievementAchievementCfgs(f func(conf *AchievementAchievementCfg)bool){
	for _,v := range gameDb.AchievementAchievementCfgs{
		if !f(v){
			return
		}
	}
}

func GetAchievementMedalMedalCfg( Id int) *AchievementMedalMedalCfg {
	return gameDb.AchievementMedalMedalCfgs[Id]
}

func RangAchievementMedalMedalCfgs(f func(conf *AchievementMedalMedalCfg)bool){
	for _,v := range gameDb.AchievementMedalMedalCfgs{
		if !f(v){
			return
		}
	}
}

func GetAncientBossAncientBossCfg( StageId int) *AncientBossAncientBossCfg {
	return gameDb.AncientBossAncientBossCfgs[StageId]
}

func RangAncientBossAncientBossCfgs(f func(conf *AncientBossAncientBossCfg)bool){
	for _,v := range gameDb.AncientBossAncientBossCfgs{
		if !f(v){
			return
		}
	}
}

func GetAncientSkillGradeAncientSkillGradeCfg( Level int) *AncientSkillGradeAncientSkillGradeCfg {
	return gameDb.AncientSkillGradeAncientSkillGradeCfgs[Level]
}

func RangAncientSkillGradeAncientSkillGradeCfgs(f func(conf *AncientSkillGradeAncientSkillGradeCfg)bool){
	for _,v := range gameDb.AncientSkillGradeAncientSkillGradeCfgs{
		if !f(v){
			return
		}
	}
}

func GetAncientSkillLevelAncientSkillLevelCfg( Level int) *AncientSkillLevelAncientSkillLevelCfg {
	return gameDb.AncientSkillLevelAncientSkillLevelCfgs[Level]
}

func RangAncientSkillLevelAncientSkillLevelCfgs(f func(conf *AncientSkillLevelAncientSkillLevelCfg)bool){
	for _,v := range gameDb.AncientSkillLevelAncientSkillLevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetArcherMagicCfg( Id int) *ArcherMagicCfg {
	return gameDb.ArcherMagicCfgs[Id]
}

func RangArcherMagicCfgs(f func(conf *ArcherMagicCfg)bool){
	for _,v := range gameDb.ArcherMagicCfgs{
		if !f(v){
			return
		}
	}
}

func GetArcherElementMagicElementCfg( Id int) *ArcherElementMagicElementCfg {
	return gameDb.ArcherElementMagicElementCfgs[Id]
}

func RangArcherElementMagicElementCfgs(f func(conf *ArcherElementMagicElementCfg)bool){
	for _,v := range gameDb.ArcherElementMagicElementCfgs{
		if !f(v){
			return
		}
	}
}

func GetAreaAreaCfg( Id int) *AreaAreaCfg {
	return gameDb.AreaAreaCfgs[Id]
}

func RangAreaAreaCfgs(f func(conf *AreaAreaCfg)bool){
	for _,v := range gameDb.AreaAreaCfgs{
		if !f(v){
			return
		}
	}
}

func GetAreaLevelAreaLevelCfg( Id int) *AreaLevelAreaLevelCfg {
	return gameDb.AreaLevelAreaLevelCfgs[Id]
}

func RangAreaLevelAreaLevelCfgs(f func(conf *AreaLevelAreaLevelCfg)bool){
	for _,v := range gameDb.AreaLevelAreaLevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetArenaBuyArenaBuyCfg( Num int) *ArenaBuyArenaBuyCfg {
	return gameDb.ArenaBuyArenaBuyCfgs[Num]
}

func RangArenaBuyArenaBuyCfgs(f func(conf *ArenaBuyArenaBuyCfg)bool){
	for _,v := range gameDb.ArenaBuyArenaBuyCfgs{
		if !f(v){
			return
		}
	}
}

func GetArenaMatchArenaMatchCfg( RankMin int) *ArenaMatchArenaMatchCfg {
	return gameDb.ArenaMatchArenaMatchCfgs[RankMin]
}

func RangArenaMatchArenaMatchCfgs(f func(conf *ArenaMatchArenaMatchCfg)bool){
	for _,v := range gameDb.ArenaMatchArenaMatchCfgs{
		if !f(v){
			return
		}
	}
}

func GetArenaRankArenaRankCfg( Id int) *ArenaRankArenaRankCfg {
	return gameDb.ArenaRankArenaRankCfgs[Id]
}

func RangArenaRankArenaRankCfgs(f func(conf *ArenaRankArenaRankCfg)bool){
	for _,v := range gameDb.ArenaRankArenaRankCfgs{
		if !f(v){
			return
		}
	}
}

func GetAspdAspdCfg( Id int) *AspdAspdCfg {
	return gameDb.AspdAspdCfgs[Id]
}

func RangAspdAspdCfgs(f func(conf *AspdAspdCfg)bool){
	for _,v := range gameDb.AspdAspdCfgs{
		if !f(v){
			return
		}
	}
}

func GetAtlasAtlasCfg( Id int) *AtlasAtlasCfg {
	return gameDb.AtlasAtlasCfgs[Id]
}

func RangAtlasAtlasCfgs(f func(conf *AtlasAtlasCfg)bool){
	for _,v := range gameDb.AtlasAtlasCfgs{
		if !f(v){
			return
		}
	}
}

func GetAtlasGatherAtlasGatherCfg( Id int) *AtlasGatherAtlasGatherCfg {
	return gameDb.AtlasGatherAtlasGatherCfgs[Id]
}

func RangAtlasGatherAtlasGatherCfgs(f func(conf *AtlasGatherAtlasGatherCfg)bool){
	for _,v := range gameDb.AtlasGatherAtlasGatherCfgs{
		if !f(v){
			return
		}
	}
}

func GetAtlasPosAtlasPosCfg( Pos int) *AtlasPosAtlasPosCfg {
	return gameDb.AtlasPosAtlasPosCfgs[Pos]
}

func RangAtlasPosAtlasPosCfgs(f func(conf *AtlasPosAtlasPosCfg)bool){
	for _,v := range gameDb.AtlasPosAtlasPosCfgs{
		if !f(v){
			return
		}
	}
}

func GetAtlasStarAtlasStarCfg( Id int) *AtlasStarAtlasStarCfg {
	return gameDb.AtlasStarAtlasStarCfgs[Id]
}

func RangAtlasStarAtlasStarCfgs(f func(conf *AtlasStarAtlasStarCfg)bool){
	for _,v := range gameDb.AtlasStarAtlasStarCfgs{
		if !f(v){
			return
		}
	}
}

func GetAtlasUpgradeAtlasUpgradeCfg( Id int) *AtlasUpgradeAtlasUpgradeCfg {
	return gameDb.AtlasUpgradeAtlasUpgradeCfgs[Id]
}

func RangAtlasUpgradeAtlasUpgradeCfgs(f func(conf *AtlasUpgradeAtlasUpgradeCfg)bool){
	for _,v := range gameDb.AtlasUpgradeAtlasUpgradeCfgs{
		if !f(v){
			return
		}
	}
}

func GetAttackEnemyAttackEnemyCfg( Id int) *AttackEnemyAttackEnemyCfg {
	return gameDb.AttackEnemyAttackEnemyCfgs[Id]
}

func RangAttackEnemyAttackEnemyCfgs(f func(conf *AttackEnemyAttackEnemyCfg)bool){
	for _,v := range gameDb.AttackEnemyAttackEnemyCfgs{
		if !f(v){
			return
		}
	}
}

func GetAttackEnemyCardAttackEnemyCardCfg( Type int) *AttackEnemyCardAttackEnemyCardCfg {
	return gameDb.AttackEnemyCardAttackEnemyCardCfgs[Type]
}

func RangAttackEnemyCardAttackEnemyCardCfgs(f func(conf *AttackEnemyCardAttackEnemyCardCfg)bool){
	for _,v := range gameDb.AttackEnemyCardAttackEnemyCardCfgs{
		if !f(v){
			return
		}
	}
}

func GetAuctionAuctioinCfg( Id int) *AuctionAuctioinCfg {
	return gameDb.AuctionAuctioinCfgs[Id]
}

func RangAuctionAuctioinCfgs(f func(conf *AuctionAuctioinCfg)bool){
	for _,v := range gameDb.AuctionAuctioinCfgs{
		if !f(v){
			return
		}
	}
}

func GetAwakenAwakenCfg( Id int) *AwakenAwakenCfg {
	return gameDb.AwakenAwakenCfgs[Id]
}

func RangAwakenAwakenCfgs(f func(conf *AwakenAwakenCfg)bool){
	for _,v := range gameDb.AwakenAwakenCfgs{
		if !f(v){
			return
		}
	}
}

func GetAwakenTitleAwakenTitleCfg( Rank int) *AwakenTitleAwakenTitleCfg {
	return gameDb.AwakenTitleAwakenTitleCfgs[Rank]
}

func RangAwakenTitleAwakenTitleCfgs(f func(conf *AwakenTitleAwakenTitleCfg)bool){
	for _,v := range gameDb.AwakenTitleAwakenTitleCfgs{
		if !f(v){
			return
		}
	}
}

func GetBagSpaceAddCfg( Id int) *BagSpaceAddCfg {
	return gameDb.BagSpaceAddCfgs[Id]
}

func RangBagSpaceAddCfgs(f func(conf *BagSpaceAddCfg)bool){
	for _,v := range gameDb.BagSpaceAddCfgs{
		if !f(v){
			return
		}
	}
}

func GetBindGroupBindGroupCfg( BindGroup int) *BindGroupBindGroupCfg {
	return gameDb.BindGroupBindGroupCfgs[BindGroup]
}

func RangBindGroupBindGroupCfgs(f func(conf *BindGroupBindGroupCfg)bool){
	for _,v := range gameDb.BindGroupBindGroupCfgs{
		if !f(v){
			return
		}
	}
}

func GetBlessBlessCfg( Id int) *BlessBlessCfg {
	return gameDb.BlessBlessCfgs[Id]
}

func RangBlessBlessCfgs(f func(conf *BlessBlessCfg)bool){
	for _,v := range gameDb.BlessBlessCfgs{
		if !f(v){
			return
		}
	}
}

func GetBossFamilyBossFamilyCfg( StageId int) *BossFamilyBossFamilyCfg {
	return gameDb.BossFamilyBossFamilyCfgs[StageId]
}

func RangBossFamilyBossFamilyCfgs(f func(conf *BossFamilyBossFamilyCfg)bool){
	for _,v := range gameDb.BossFamilyBossFamilyCfgs{
		if !f(v){
			return
		}
	}
}

func GetBuffBuffCfg( Id int) *BuffBuffCfg {
	return gameDb.BuffBuffCfgs[Id]
}

func RangBuffBuffCfgs(f func(conf *BuffBuffCfg)bool){
	for _,v := range gameDb.BuffBuffCfgs{
		if !f(v){
			return
		}
	}
}

func GetChatClearCfg( Type int) *ChatClearCfg {
	return gameDb.ChatClearCfgs[Type]
}

func RangChatClearCfgs(f func(conf *ChatClearCfg)bool){
	for _,v := range gameDb.ChatClearCfgs{
		if !f(v){
			return
		}
	}
}

func GetChuanShiEquipChuanShiEquipCfg( Id int) *ChuanShiEquipChuanShiEquipCfg {
	return gameDb.ChuanShiEquipChuanShiEquipCfgs[Id]
}

func RangChuanShiEquipChuanShiEquipCfgs(f func(conf *ChuanShiEquipChuanShiEquipCfg)bool){
	for _,v := range gameDb.ChuanShiEquipChuanShiEquipCfgs{
		if !f(v){
			return
		}
	}
}

func GetChuanShiEquipTypeChuanShiEquipTypeCfg( Id int) *ChuanShiEquipTypeChuanShiEquipTypeCfg {
	return gameDb.ChuanShiEquipTypeChuanShiEquipTypeCfgs[Id]
}

func RangChuanShiEquipTypeChuanShiEquipTypeCfgs(f func(conf *ChuanShiEquipTypeChuanShiEquipTypeCfg)bool){
	for _,v := range gameDb.ChuanShiEquipTypeChuanShiEquipTypeCfgs{
		if !f(v){
			return
		}
	}
}

func GetChuanShiStrengthenChuanShiStrengthenCfg( Id int) *ChuanShiStrengthenChuanShiStrengthenCfg {
	return gameDb.ChuanShiStrengthenChuanShiStrengthenCfgs[Id]
}

func RangChuanShiStrengthenChuanShiStrengthenCfgs(f func(conf *ChuanShiStrengthenChuanShiStrengthenCfg)bool){
	for _,v := range gameDb.ChuanShiStrengthenChuanShiStrengthenCfgs{
		if !f(v){
			return
		}
	}
}

func GetChuanShiStrengthenLinkChuanShiStrengthenLinkCfg( Id int) *ChuanShiStrengthenLinkChuanShiStrengthenLinkCfg {
	return gameDb.ChuanShiStrengthenLinkChuanShiStrengthenLinkCfgs[Id]
}

func RangChuanShiStrengthenLinkChuanShiStrengthenLinkCfgs(f func(conf *ChuanShiStrengthenLinkChuanShiStrengthenLinkCfg)bool){
	for _,v := range gameDb.ChuanShiStrengthenLinkChuanShiStrengthenLinkCfgs{
		if !f(v){
			return
		}
	}
}

func GetChuanShiSuitChuanShiSuitCfg( Id int) *ChuanShiSuitChuanShiSuitCfg {
	return gameDb.ChuanShiSuitChuanShiSuitCfgs[Id]
}

func RangChuanShiSuitChuanShiSuitCfgs(f func(conf *ChuanShiSuitChuanShiSuitCfg)bool){
	for _,v := range gameDb.ChuanShiSuitChuanShiSuitCfgs{
		if !f(v){
			return
		}
	}
}

func GetChuanShiSuitTypeChuanShiSuitTypeCfg( Id int) *ChuanShiSuitTypeChuanShiSuitTypeCfg {
	return gameDb.ChuanShiSuitTypeChuanShiSuitTypeCfgs[Id]
}

func RangChuanShiSuitTypeChuanShiSuitTypeCfgs(f func(conf *ChuanShiSuitTypeChuanShiSuitTypeCfg)bool){
	for _,v := range gameDb.ChuanShiSuitTypeChuanShiSuitTypeCfgs{
		if !f(v){
			return
		}
	}
}

func GetClearClearCfg( Type int) *ClearClearCfg {
	return gameDb.ClearClearCfgs[Type]
}

func RangClearClearCfgs(f func(conf *ClearClearCfg)bool){
	for _,v := range gameDb.ClearClearCfgs{
		if !f(v){
			return
		}
	}
}

func GetClearRateClearRateCfg( Id int) *ClearRateClearRateCfg {
	return gameDb.ClearRateClearRateCfgs[Id]
}

func RangClearRateClearRateCfgs(f func(conf *ClearRateClearRateCfg)bool){
	for _,v := range gameDb.ClearRateClearRateCfgs{
		if !f(v){
			return
		}
	}
}

func GetCollectionCollectionCfg( Id int) *CollectionCollectionCfg {
	return gameDb.CollectionCollectionCfgs[Id]
}

func RangCollectionCollectionCfgs(f func(conf *CollectionCollectionCfg)bool){
	for _,v := range gameDb.CollectionCollectionCfgs{
		if !f(v){
			return
		}
	}
}

func GetCompetitveCompetitveCfg( Id int) *CompetitveCompetitveCfg {
	return gameDb.CompetitveCompetitveCfgs[Id]
}

func RangCompetitveCompetitveCfgs(f func(conf *CompetitveCompetitveCfg)bool){
	for _,v := range gameDb.CompetitveCompetitveCfgs{
		if !f(v){
			return
		}
	}
}

func GetCompetitveRewardRankRewardCfg( Id int) *CompetitveRewardRankRewardCfg {
	return gameDb.CompetitveRewardRankRewardCfgs[Id]
}

func RangCompetitveRewardRankRewardCfgs(f func(conf *CompetitveRewardRankRewardCfg)bool){
	for _,v := range gameDb.CompetitveRewardRankRewardCfgs{
		if !f(v){
			return
		}
	}
}

func GetComposeChuanShiSubComposeChuanShiSubCfg( Id int) *ComposeChuanShiSubComposeChuanShiSubCfg {
	return gameDb.ComposeChuanShiSubComposeChuanShiSubCfgs[Id]
}

func RangComposeChuanShiSubComposeChuanShiSubCfgs(f func(conf *ComposeChuanShiSubComposeChuanShiSubCfg)bool){
	for _,v := range gameDb.ComposeChuanShiSubComposeChuanShiSubCfgs{
		if !f(v){
			return
		}
	}
}

func GetComposeChuanShiTypeComposeChuanShiTypeCfg( Id int) *ComposeChuanShiTypeComposeChuanShiTypeCfg {
	return gameDb.ComposeChuanShiTypeComposeChuanShiTypeCfgs[Id]
}

func RangComposeChuanShiTypeComposeChuanShiTypeCfgs(f func(conf *ComposeChuanShiTypeComposeChuanShiTypeCfg)bool){
	for _,v := range gameDb.ComposeChuanShiTypeComposeChuanShiTypeCfgs{
		if !f(v){
			return
		}
	}
}

func GetComposeEquipSubComposeEquipSubCfg( Id int) *ComposeEquipSubComposeEquipSubCfg {
	return gameDb.ComposeEquipSubComposeEquipSubCfgs[Id]
}

func RangComposeEquipSubComposeEquipSubCfgs(f func(conf *ComposeEquipSubComposeEquipSubCfg)bool){
	for _,v := range gameDb.ComposeEquipSubComposeEquipSubCfgs{
		if !f(v){
			return
		}
	}
}

func GetComposeEquipTypeComposeEquipTypeCfg( Id int) *ComposeEquipTypeComposeEquipTypeCfg {
	return gameDb.ComposeEquipTypeComposeEquipTypeCfgs[Id]
}

func RangComposeEquipTypeComposeEquipTypeCfgs(f func(conf *ComposeEquipTypeComposeEquipTypeCfg)bool){
	for _,v := range gameDb.ComposeEquipTypeComposeEquipTypeCfgs{
		if !f(v){
			return
		}
	}
}

func GetComposeSubComposeSubCfg( Id int) *ComposeSubComposeSubCfg {
	return gameDb.ComposeSubComposeSubCfgs[Id]
}

func RangComposeSubComposeSubCfgs(f func(conf *ComposeSubComposeSubCfg)bool){
	for _,v := range gameDb.ComposeSubComposeSubCfgs{
		if !f(v){
			return
		}
	}
}

func GetComposeTypeComposeTypeCfg( Id int) *ComposeTypeComposeTypeCfg {
	return gameDb.ComposeTypeComposeTypeCfgs[Id]
}

func RangComposeTypeComposeTypeCfgs(f func(conf *ComposeTypeComposeTypeCfg)bool){
	for _,v := range gameDb.ComposeTypeComposeTypeCfgs{
		if !f(v){
			return
		}
	}
}

func GetConditionConditionCfg( Id int) *ConditionConditionCfg {
	return gameDb.ConditionConditionCfgs[Id]
}

func RangConditionConditionCfgs(f func(conf *ConditionConditionCfg)bool){
	for _,v := range gameDb.ConditionConditionCfgs{
		if !f(v){
			return
		}
	}
}

func GetContRechargeContRechargeCfg( Id int) *ContRechargeContRechargeCfg {
	return gameDb.ContRechargeContRechargeCfgs[Id]
}

func RangContRechargeContRechargeCfgs(f func(conf *ContRechargeContRechargeCfg)bool){
	for _,v := range gameDb.ContRechargeContRechargeCfgs{
		if !f(v){
			return
		}
	}
}

func GetCrossArenaCrossArenaCfg( Id int) *CrossArenaCrossArenaCfg {
	return gameDb.CrossArenaCrossArenaCfgs[Id]
}

func RangCrossArenaCrossArenaCfgs(f func(conf *CrossArenaCrossArenaCfg)bool){
	for _,v := range gameDb.CrossArenaCrossArenaCfgs{
		if !f(v){
			return
		}
	}
}

func GetCrossArenaRewardCrossArenaRewardCfg( Id int) *CrossArenaRewardCrossArenaRewardCfg {
	return gameDb.CrossArenaRewardCrossArenaRewardCfgs[Id]
}

func RangCrossArenaRewardCrossArenaRewardCfgs(f func(conf *CrossArenaRewardCrossArenaRewardCfg)bool){
	for _,v := range gameDb.CrossArenaRewardCrossArenaRewardCfgs{
		if !f(v){
			return
		}
	}
}

func GetCrossArenaRobotCrossArenaRobotCfg( Id int) *CrossArenaRobotCrossArenaRobotCfg {
	return gameDb.CrossArenaRobotCrossArenaRobotCfgs[Id]
}

func RangCrossArenaRobotCrossArenaRobotCfgs(f func(conf *CrossArenaRobotCrossArenaRobotCfg)bool){
	for _,v := range gameDb.CrossArenaRobotCrossArenaRobotCfgs{
		if !f(v){
			return
		}
	}
}

func GetCrossArenaTimeCrossArenaTimeCfg( Id int) *CrossArenaTimeCrossArenaTimeCfg {
	return gameDb.CrossArenaTimeCrossArenaTimeCfgs[Id]
}

func RangCrossArenaTimeCrossArenaTimeCfgs(f func(conf *CrossArenaTimeCrossArenaTimeCfg)bool){
	for _,v := range gameDb.CrossArenaTimeCrossArenaTimeCfgs{
		if !f(v){
			return
		}
	}
}

func GetCumulationsignCumulationsignCfg( Id int) *CumulationsignCumulationsignCfg {
	return gameDb.CumulationsignCumulationsignCfgs[Id]
}

func RangCumulationsignCumulationsignCfgs(f func(conf *CumulationsignCumulationsignCfg)bool){
	for _,v := range gameDb.CumulationsignCumulationsignCfgs{
		if !f(v){
			return
		}
	}
}

func GetCutTreasureCutTreasureCfg( Id int) *CutTreasureCutTreasureCfg {
	return gameDb.CutTreasureCutTreasureCfgs[Id]
}

func RangCutTreasureCutTreasureCfgs(f func(conf *CutTreasureCutTreasureCfg)bool){
	for _,v := range gameDb.CutTreasureCutTreasureCfgs{
		if !f(v){
			return
		}
	}
}

func GetDaBaoEquipDaBaoEquipCfg( Id int) *DaBaoEquipDaBaoEquipCfg {
	return gameDb.DaBaoEquipDaBaoEquipCfgs[Id]
}

func RangDaBaoEquipDaBaoEquipCfgs(f func(conf *DaBaoEquipDaBaoEquipCfg)bool){
	for _,v := range gameDb.DaBaoEquipDaBaoEquipCfgs{
		if !f(v){
			return
		}
	}
}

func GetDaBaoEquipAdditionDaBaoEquipAdditionCfg( Id int) *DaBaoEquipAdditionDaBaoEquipAdditionCfg {
	return gameDb.DaBaoEquipAdditionDaBaoEquipAdditionCfgs[Id]
}

func RangDaBaoEquipAdditionDaBaoEquipAdditionCfgs(f func(conf *DaBaoEquipAdditionDaBaoEquipAdditionCfg)bool){
	for _,v := range gameDb.DaBaoEquipAdditionDaBaoEquipAdditionCfgs{
		if !f(v){
			return
		}
	}
}

func GetDaBaoMysteryDaBaoMysteryCfg( Id int) *DaBaoMysteryDaBaoMysteryCfg {
	return gameDb.DaBaoMysteryDaBaoMysteryCfgs[Id]
}

func RangDaBaoMysteryDaBaoMysteryCfgs(f func(conf *DaBaoMysteryDaBaoMysteryCfg)bool){
	for _,v := range gameDb.DaBaoMysteryDaBaoMysteryCfgs{
		if !f(v){
			return
		}
	}
}

func GetDailyActivityDailyActivityCfg( Id int) *DailyActivityDailyActivityCfg {
	return gameDb.DailyActivityDailyActivityCfgs[Id]
}

func RangDailyActivityDailyActivityCfgs(f func(conf *DailyActivityDailyActivityCfg)bool){
	for _,v := range gameDb.DailyActivityDailyActivityCfgs{
		if !f(v){
			return
		}
	}
}

func GetDailyRewardDailyRewardCfg( Id int) *DailyRewardDailyRewardCfg {
	return gameDb.DailyRewardDailyRewardCfgs[Id]
}

func RangDailyRewardDailyRewardCfgs(f func(conf *DailyRewardDailyRewardCfg)bool){
	for _,v := range gameDb.DailyRewardDailyRewardCfgs{
		if !f(v){
			return
		}
	}
}

func GetDailyTaskDailytaskCfg( Id int) *DailyTaskDailytaskCfg {
	return gameDb.DailyTaskDailytaskCfgs[Id]
}

func RangDailyTaskDailytaskCfgs(f func(conf *DailyTaskDailytaskCfg)bool){
	for _,v := range gameDb.DailyTaskDailytaskCfgs{
		if !f(v){
			return
		}
	}
}

func GetDailypackDailypackCfg( Id int) *DailypackDailypackCfg {
	return gameDb.DailypackDailypackCfgs[Id]
}

func RangDailypackDailypackCfgs(f func(conf *DailypackDailypackCfg)bool){
	for _,v := range gameDb.DailypackDailypackCfgs{
		if !f(v){
			return
		}
	}
}

func GetDarkPalaceBossDarkPalaceBossCfg( Id int) *DarkPalaceBossDarkPalaceBossCfg {
	return gameDb.DarkPalaceBossDarkPalaceBossCfgs[Id]
}

func RangDarkPalaceBossDarkPalaceBossCfgs(f func(conf *DarkPalaceBossDarkPalaceBossCfg)bool){
	for _,v := range gameDb.DarkPalaceBossDarkPalaceBossCfgs{
		if !f(v){
			return
		}
	}
}

func GetDayRankingDayRankingCfg( Day int) *DayRankingDayRankingCfg {
	return gameDb.DayRankingDayRankingCfgs[Day]
}

func RangDayRankingDayRankingCfgs(f func(conf *DayRankingDayRankingCfg)bool){
	for _,v := range gameDb.DayRankingDayRankingCfgs{
		if !f(v){
			return
		}
	}
}

func GetDayRankingGiftDayRankingGiftCfg( Id int) *DayRankingGiftDayRankingGiftCfg {
	return gameDb.DayRankingGiftDayRankingGiftCfgs[Id]
}

func RangDayRankingGiftDayRankingGiftCfgs(f func(conf *DayRankingGiftDayRankingGiftCfg)bool){
	for _,v := range gameDb.DayRankingGiftDayRankingGiftCfgs{
		if !f(v){
			return
		}
	}
}

func GetDayRankingMarkDayRankingMarkCfg( Id int) *DayRankingMarkDayRankingMarkCfg {
	return gameDb.DayRankingMarkDayRankingMarkCfgs[Id]
}

func RangDayRankingMarkDayRankingMarkCfgs(f func(conf *DayRankingMarkDayRankingMarkCfg)bool){
	for _,v := range gameDb.DayRankingMarkDayRankingMarkCfgs{
		if !f(v){
			return
		}
	}
}

func GetDayRankingRewardDayRankingRewardCfg( Id int) *DayRankingRewardDayRankingRewardCfg {
	return gameDb.DayRankingRewardDayRankingRewardCfgs[Id]
}

func RangDayRankingRewardDayRankingRewardCfgs(f func(conf *DayRankingRewardDayRankingRewardCfg)bool){
	for _,v := range gameDb.DayRankingRewardDayRankingRewardCfgs{
		if !f(v){
			return
		}
	}
}

func GetDictateEquipDictateEquipCfg( Id int) *DictateEquipDictateEquipCfg {
	return gameDb.DictateEquipDictateEquipCfgs[Id]
}

func RangDictateEquipDictateEquipCfgs(f func(conf *DictateEquipDictateEquipCfg)bool){
	for _,v := range gameDb.DictateEquipDictateEquipCfgs{
		if !f(v){
			return
		}
	}
}

func GetDictateSuitDictateSuitCfg( Grade int) *DictateSuitDictateSuitCfg {
	return gameDb.DictateSuitDictateSuitCfgs[Grade]
}

func RangDictateSuitDictateSuitCfgs(f func(conf *DictateSuitDictateSuitCfg)bool){
	for _,v := range gameDb.DictateSuitDictateSuitCfgs{
		if !f(v){
			return
		}
	}
}

func GetDragonEquipDragonEquipCfg( Id int) *DragonEquipDragonEquipCfg {
	return gameDb.DragonEquipDragonEquipCfgs[Id]
}

func RangDragonEquipDragonEquipCfgs(f func(conf *DragonEquipDragonEquipCfg)bool){
	for _,v := range gameDb.DragonEquipDragonEquipCfgs{
		if !f(v){
			return
		}
	}
}

func GetDragonEquipLevelDragonEquipLevelCfg( Id int) *DragonEquipLevelDragonEquipLevelCfg {
	return gameDb.DragonEquipLevelDragonEquipLevelCfgs[Id]
}

func RangDragonEquipLevelDragonEquipLevelCfgs(f func(conf *DragonEquipLevelDragonEquipLevelCfg)bool){
	for _,v := range gameDb.DragonEquipLevelDragonEquipLevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetDragonarmsDragonarmsCfg( Id int) *DragonarmsDragonarmsCfg {
	return gameDb.DragonarmsDragonarmsCfgs[Id]
}

func RangDragonarmsDragonarmsCfgs(f func(conf *DragonarmsDragonarmsCfg)bool){
	for _,v := range gameDb.DragonarmsDragonarmsCfgs{
		if !f(v){
			return
		}
	}
}

func GetDrawDrawCfg( Id int) *DrawDrawCfg {
	return gameDb.DrawDrawCfgs[Id]
}

func RangDrawDrawCfgs(f func(conf *DrawDrawCfg)bool){
	for _,v := range gameDb.DrawDrawCfgs{
		if !f(v){
			return
		}
	}
}

func GetDrawShopDrawShopCfg( Id int) *DrawShopDrawShopCfg {
	return gameDb.DrawShopDrawShopCfgs[Id]
}

func RangDrawShopDrawShopCfgs(f func(conf *DrawShopDrawShopCfg)bool){
	for _,v := range gameDb.DrawShopDrawShopCfgs{
		if !f(v){
			return
		}
	}
}

func GetDropDropCfg( Id int) *DropDropCfg {
	return gameDb.DropDropCfgs[Id]
}

func RangDropDropCfgs(f func(conf *DropDropCfg)bool){
	for _,v := range gameDb.DropDropCfgs{
		if !f(v){
			return
		}
	}
}

func GetDropSpecialDropSpecialCfg( Id int) *DropSpecialDropSpecialCfg {
	return gameDb.DropSpecialDropSpecialCfgs[Id]
}

func RangDropSpecialDropSpecialCfgs(f func(conf *DropSpecialDropSpecialCfg)bool){
	for _,v := range gameDb.DropSpecialDropSpecialCfgs{
		if !f(v){
			return
		}
	}
}

func GetEffectEffectCfg( Id int) *EffectEffectCfg {
	return gameDb.EffectEffectCfgs[Id]
}

func RangEffectEffectCfgs(f func(conf *EffectEffectCfg)bool){
	for _,v := range gameDb.EffectEffectCfgs{
		if !f(v){
			return
		}
	}
}

func GetElfGrowElfGrowCfg( Level int) *ElfGrowElfGrowCfg {
	return gameDb.ElfGrowElfGrowCfgs[Level]
}

func RangElfGrowElfGrowCfgs(f func(conf *ElfGrowElfGrowCfg)bool){
	for _,v := range gameDb.ElfGrowElfGrowCfgs{
		if !f(v){
			return
		}
	}
}

func GetElfRecoverElfRecoverCfg( Id int) *ElfRecoverElfRecoverCfg {
	return gameDb.ElfRecoverElfRecoverCfgs[Id]
}

func RangElfRecoverElfRecoverCfgs(f func(conf *ElfRecoverElfRecoverCfg)bool){
	for _,v := range gameDb.ElfRecoverElfRecoverCfgs{
		if !f(v){
			return
		}
	}
}

func GetElfSkillElfGrowCfg( Id int) *ElfSkillElfGrowCfg {
	return gameDb.ElfSkillElfGrowCfgs[Id]
}

func RangElfSkillElfGrowCfgs(f func(conf *ElfSkillElfGrowCfg)bool){
	for _,v := range gameDb.ElfSkillElfGrowCfgs{
		if !f(v){
			return
		}
	}
}

func GetEquipEquipCfg( Id int) *EquipEquipCfg {
	return gameDb.EquipEquipCfgs[Id]
}

func RangEquipEquipCfgs(f func(conf *EquipEquipCfg)bool){
	for _,v := range gameDb.EquipEquipCfgs{
		if !f(v){
			return
		}
	}
}

func GetEquipsuitEquipsuitCfg( Id int) *EquipsuitEquipsuitCfg {
	return gameDb.EquipsuitEquipsuitCfgs[Id]
}

func RangEquipsuitEquipsuitCfgs(f func(conf *EquipsuitEquipsuitCfg)bool){
	for _,v := range gameDb.EquipsuitEquipsuitCfgs{
		if !f(v){
			return
		}
	}
}

func GetExpLevelLevelCfg( Id int) *ExpLevelLevelCfg {
	return gameDb.ExpLevelLevelCfgs[Id]
}

func RangExpLevelLevelCfgs(f func(conf *ExpLevelLevelCfg)bool){
	for _,v := range gameDb.ExpLevelLevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetExpPillExpPillCfg( Id int) *ExpPillExpPillCfg {
	return gameDb.ExpPillExpPillCfgs[Id]
}

func RangExpPillExpPillCfgs(f func(conf *ExpPillExpPillCfg)bool){
	for _,v := range gameDb.ExpPillExpPillCfgs{
		if !f(v){
			return
		}
	}
}

func GetExpPoolExpPoolCfg( Id int) *ExpPoolExpPoolCfg {
	return gameDb.ExpPoolExpPoolCfgs[Id]
}

func RangExpPoolExpPoolCfgs(f func(conf *ExpPoolExpPoolCfg)bool){
	for _,v := range gameDb.ExpPoolExpPoolCfgs{
		if !f(v){
			return
		}
	}
}

func GetExpStageExpStageCfg( Id int) *ExpStageExpStageCfg {
	return gameDb.ExpStageExpStageCfgs[Id]
}

func RangExpStageExpStageCfgs(f func(conf *ExpStageExpStageCfg)bool){
	for _,v := range gameDb.ExpStageExpStageCfgs{
		if !f(v){
			return
		}
	}
}

func GetFabaoFabaoCfg( Id int) *FabaoFabaoCfg {
	return gameDb.FabaoFabaoCfgs[Id]
}

func RangFabaoFabaoCfgs(f func(conf *FabaoFabaoCfg)bool){
	for _,v := range gameDb.FabaoFabaoCfgs{
		if !f(v){
			return
		}
	}
}

func GetFabaoSkillFabaoSkillCfg( Id int) *FabaoSkillFabaoSkillCfg {
	return gameDb.FabaoSkillFabaoSkillCfgs[Id]
}

func RangFabaoSkillFabaoSkillCfgs(f func(conf *FabaoSkillFabaoSkillCfg)bool){
	for _,v := range gameDb.FabaoSkillFabaoSkillCfgs{
		if !f(v){
			return
		}
	}
}

func GetFabaolevelFabaolevelCfg( Id int) *FabaolevelFabaolevelCfg {
	return gameDb.FabaolevelFabaolevelCfgs[Id]
}

func RangFabaolevelFabaolevelCfgs(f func(conf *FabaolevelFabaolevelCfg)bool){
	for _,v := range gameDb.FabaolevelFabaolevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetFashionFashionCfg( Id int) *FashionFashionCfg {
	return gameDb.FashionFashionCfgs[Id]
}

func RangFashionFashionCfgs(f func(conf *FashionFashionCfg)bool){
	for _,v := range gameDb.FashionFashionCfgs{
		if !f(v){
			return
		}
	}
}

func GetFieldBossFieldBossCfg( StageId int) *FieldBossFieldBossCfg {
	return gameDb.FieldBossFieldBossCfgs[StageId]
}

func RangFieldBossFieldBossCfgs(f func(conf *FieldBossFieldBossCfg)bool){
	for _,v := range gameDb.FieldBossFieldBossCfgs{
		if !f(v){
			return
		}
	}
}

func GetFieldFightFieldBaseCfg( Id int) *FieldFightFieldBaseCfg {
	return gameDb.FieldFightFieldBaseCfgs[Id]
}

func RangFieldFightFieldBaseCfgs(f func(conf *FieldFightFieldBaseCfg)bool){
	for _,v := range gameDb.FieldFightFieldBaseCfgs{
		if !f(v){
			return
		}
	}
}

func GetFirstBloodPerFirstBloodperCfg( Id int) *FirstBloodPerFirstBloodperCfg {
	return gameDb.FirstBloodPerFirstBloodperCfgs[Id]
}

func RangFirstBloodPerFirstBloodperCfgs(f func(conf *FirstBloodPerFirstBloodperCfg)bool){
	for _,v := range gameDb.FirstBloodPerFirstBloodperCfgs{
		if !f(v){
			return
		}
	}
}

func GetFirstBloodmilFirstBloodmilCfg( Id int) *FirstBloodmilFirstBloodmilCfg {
	return gameDb.FirstBloodmilFirstBloodmilCfgs[Id]
}

func RangFirstBloodmilFirstBloodmilCfgs(f func(conf *FirstBloodmilFirstBloodmilCfg)bool){
	for _,v := range gameDb.FirstBloodmilFirstBloodmilCfgs{
		if !f(v){
			return
		}
	}
}

func GetFirstBlooduniFirstBlooduniCfg( Id int) *FirstBlooduniFirstBlooduniCfg {
	return gameDb.FirstBlooduniFirstBlooduniCfgs[Id]
}

func RangFirstBlooduniFirstBlooduniCfgs(f func(conf *FirstBlooduniFirstBlooduniCfg)bool){
	for _,v := range gameDb.FirstBlooduniFirstBlooduniCfgs{
		if !f(v){
			return
		}
	}
}

func GetFirstDropFirstDropCfg( Id int) *FirstDropFirstDropCfg {
	return gameDb.FirstDropFirstDropCfgs[Id]
}

func RangFirstDropFirstDropCfgs(f func(conf *FirstDropFirstDropCfg)bool){
	for _,v := range gameDb.FirstDropFirstDropCfgs{
		if !f(v){
			return
		}
	}
}

func GetFirstRechargTypeFirstRechargTypeCfg( Id int) *FirstRechargTypeFirstRechargTypeCfg {
	return gameDb.FirstRechargTypeFirstRechargTypeCfgs[Id]
}

func RangFirstRechargTypeFirstRechargTypeCfgs(f func(conf *FirstRechargTypeFirstRechargTypeCfg)bool){
	for _,v := range gameDb.FirstRechargTypeFirstRechargTypeCfgs{
		if !f(v){
			return
		}
	}
}

func GetFirstRechargeFirstRechargCfg( Day int) *FirstRechargeFirstRechargCfg {
	return gameDb.FirstRechargeFirstRechargCfgs[Day]
}

func RangFirstRechargeFirstRechargCfgs(f func(conf *FirstRechargeFirstRechargCfg)bool){
	for _,v := range gameDb.FirstRechargeFirstRechargCfgs{
		if !f(v){
			return
		}
	}
}

func GetFitFashionFitFashionCfg( Id int) *FitFashionFitFashionCfg {
	return gameDb.FitFashionFitFashionCfgs[Id]
}

func RangFitFashionFitFashionCfgs(f func(conf *FitFashionFitFashionCfg)bool){
	for _,v := range gameDb.FitFashionFitFashionCfgs{
		if !f(v){
			return
		}
	}
}

func GetFitFashionLevelFitFashionLevelCfg( Id int) *FitFashionLevelFitFashionLevelCfg {
	return gameDb.FitFashionLevelFitFashionLevelCfgs[Id]
}

func RangFitFashionLevelFitFashionLevelCfgs(f func(conf *FitFashionLevelFitFashionLevelCfg)bool){
	for _,v := range gameDb.FitFashionLevelFitFashionLevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetFitHolyEquipFitHolyEquipCfg( Id int) *FitHolyEquipFitHolyEquipCfg {
	return gameDb.FitHolyEquipFitHolyEquipCfgs[Id]
}

func RangFitHolyEquipFitHolyEquipCfgs(f func(conf *FitHolyEquipFitHolyEquipCfg)bool){
	for _,v := range gameDb.FitHolyEquipFitHolyEquipCfgs{
		if !f(v){
			return
		}
	}
}

func GetFitHolyEquipSuitFitHolyEquipSuitCfg( Grade int) *FitHolyEquipSuitFitHolyEquipSuitCfg {
	return gameDb.FitHolyEquipSuitFitHolyEquipSuitCfgs[Grade]
}

func RangFitHolyEquipSuitFitHolyEquipSuitCfgs(f func(conf *FitHolyEquipSuitFitHolyEquipSuitCfg)bool){
	for _,v := range gameDb.FitHolyEquipSuitFitHolyEquipSuitCfgs{
		if !f(v){
			return
		}
	}
}

func GetFitLevelFitLevelCfg( Id int) *FitLevelFitLevelCfg {
	return gameDb.FitLevelFitLevelCfgs[Id]
}

func RangFitLevelFitLevelCfgs(f func(conf *FitLevelFitLevelCfg)bool){
	for _,v := range gameDb.FitLevelFitLevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetFitSkillFitSkillCfg( Id int) *FitSkillFitSkillCfg {
	return gameDb.FitSkillFitSkillCfgs[Id]
}

func RangFitSkillFitSkillCfgs(f func(conf *FitSkillFitSkillCfg)bool){
	for _,v := range gameDb.FitSkillFitSkillCfgs{
		if !f(v){
			return
		}
	}
}

func GetFitSkillLevelFitSkillLevelCfg( Id int) *FitSkillLevelFitSkillLevelCfg {
	return gameDb.FitSkillLevelFitSkillLevelCfgs[Id]
}

func RangFitSkillLevelFitSkillLevelCfgs(f func(conf *FitSkillLevelFitSkillLevelCfg)bool){
	for _,v := range gameDb.FitSkillLevelFitSkillLevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetFitSkillSlotFitSkillSlotCfg( Id int) *FitSkillSlotFitSkillSlotCfg {
	return gameDb.FitSkillSlotFitSkillSlotCfgs[Id]
}

func RangFitSkillSlotFitSkillSlotCfgs(f func(conf *FitSkillSlotFitSkillSlotCfg)bool){
	for _,v := range gameDb.FitSkillSlotFitSkillSlotCfgs{
		if !f(v){
			return
		}
	}
}

func GetFitSkillStarFitSkillStarCfg( Id int) *FitSkillStarFitSkillStarCfg {
	return gameDb.FitSkillStarFitSkillStarCfgs[Id]
}

func RangFitSkillStarFitSkillStarCfgs(f func(conf *FitSkillStarFitSkillStarCfg)bool){
	for _,v := range gameDb.FitSkillStarFitSkillStarCfgs{
		if !f(v){
			return
		}
	}
}

func GetFunctionFunctionCfg( Id int) *FunctionFunctionCfg {
	return gameDb.FunctionFunctionCfgs[Id]
}

func RangFunctionFunctionCfgs(f func(conf *FunctionFunctionCfg)bool){
	for _,v := range gameDb.FunctionFunctionCfgs{
		if !f(v){
			return
		}
	}
}

func GetGameDotGameDotCfg( Id int) *GameDotGameDotCfg {
	return gameDb.GameDotGameDotCfgs[Id]
}

func RangGameDotGameDotCfgs(f func(conf *GameDotGameDotCfg)bool){
	for _,v := range gameDb.GameDotGameDotCfgs{
		if !f(v){
			return
		}
	}
}

func GetGameTextErrorTextCfg( Id int) *GameTextErrorTextCfg {
	return gameDb.GameTextErrorTextCfgs[Id]
}

func RangGameTextErrorTextCfgs(f func(conf *GameTextErrorTextCfg)bool){
	for _,v := range gameDb.GameTextErrorTextCfgs{
		if !f(v){
			return
		}
	}
}

func GetGameTextCodeTextCfg( Id int) *GameTextCodeTextCfg {
	return gameDb.GameTextCodeTextCfgs[Id]
}

func RangGameTextCodeTextCfgs(f func(conf *GameTextCodeTextCfg)bool){
	for _,v := range gameDb.GameTextCodeTextCfgs{
		if !f(v){
			return
		}
	}
}

func GetGamewordGameCfg( Id int) *GamewordGameCfg {
	return gameDb.GamewordGameCfgs[Id]
}

func RangGamewordGameCfgs(f func(conf *GamewordGameCfg)bool){
	for _,v := range gameDb.GamewordGameCfgs{
		if !f(v){
			return
		}
	}
}

func GetGiftGiftCfg( Id int) *GiftGiftCfg {
	return gameDb.GiftGiftCfgs[Id]
}

func RangGiftGiftCfgs(f func(conf *GiftGiftCfg)bool){
	for _,v := range gameDb.GiftGiftCfgs{
		if !f(v){
			return
		}
	}
}

func GetGiftCodeGiftCodeCfg( Id int) *GiftCodeGiftCodeCfg {
	return gameDb.GiftCodeGiftCodeCfgs[Id]
}

func RangGiftCodeGiftCodeCfgs(f func(conf *GiftCodeGiftCodeCfg)bool){
	for _,v := range gameDb.GiftCodeGiftCodeCfgs{
		if !f(v){
			return
		}
	}
}

func GetGodBloodGodBloodCfg( Id int) *GodBloodGodBloodCfg {
	return gameDb.GodBloodGodBloodCfgs[Id]
}

func RangGodBloodGodBloodCfgs(f func(conf *GodBloodGodBloodCfg)bool){
	for _,v := range gameDb.GodBloodGodBloodCfgs{
		if !f(v){
			return
		}
	}
}

func GetGodEquipConfCfg( Id int) *GodEquipConfCfg {
	return gameDb.GodEquipConfCfgs[Id]
}

func RangGodEquipConfCfgs(f func(conf *GodEquipConfCfg)bool){
	for _,v := range gameDb.GodEquipConfCfgs{
		if !f(v){
			return
		}
	}
}

func GetGodEquipLevelConfCfg( Id int) *GodEquipLevelConfCfg {
	return gameDb.GodEquipLevelConfCfgs[Id]
}

func RangGodEquipLevelConfCfgs(f func(conf *GodEquipLevelConfCfg)bool){
	for _,v := range gameDb.GodEquipLevelConfCfgs{
		if !f(v){
			return
		}
	}
}

func GetGodEquipSuitConfCfg( Level int) *GodEquipSuitConfCfg {
	return gameDb.GodEquipSuitConfCfgs[Level]
}

func RangGodEquipSuitConfCfgs(f func(conf *GodEquipSuitConfCfg)bool){
	for _,v := range gameDb.GodEquipSuitConfCfgs{
		if !f(v){
			return
		}
	}
}

func GetGrowFundGrowFundCfg( Id int) *GrowFundGrowFundCfg {
	return gameDb.GrowFundGrowFundCfgs[Id]
}

func RangGrowFundGrowFundCfgs(f func(conf *GrowFundGrowFundCfg)bool){
	for _,v := range gameDb.GrowFundGrowFundCfgs{
		if !f(v){
			return
		}
	}
}

func GetGuardRankGuardRankCfg( Id int) *GuardRankGuardRankCfg {
	return gameDb.GuardRankGuardRankCfgs[Id]
}

func RangGuardRankGuardRankCfgs(f func(conf *GuardRankGuardRankCfg)bool){
	for _,v := range gameDb.GuardRankGuardRankCfgs{
		if !f(v){
			return
		}
	}
}

func GetGuardRoundsGuardRoundsCfg( Rounds int) *GuardRoundsGuardRoundsCfg {
	return gameDb.GuardRoundsGuardRoundsCfgs[Rounds]
}

func RangGuardRoundsGuardRoundsCfgs(f func(conf *GuardRoundsGuardRoundsCfg)bool){
	for _,v := range gameDb.GuardRoundsGuardRoundsCfgs{
		if !f(v){
			return
		}
	}
}

func GetGuildGuildCfg( Id int) *GuildGuildCfg {
	return gameDb.GuildGuildCfgs[Id]
}

func RangGuildGuildCfgs(f func(conf *GuildGuildCfg)bool){
	for _,v := range gameDb.GuildGuildCfgs{
		if !f(v){
			return
		}
	}
}

func GetGuildActivityGuildActivityCfg( Id int) *GuildActivityGuildActivityCfg {
	return gameDb.GuildActivityGuildActivityCfgs[Id]
}

func RangGuildActivityGuildActivityCfgs(f func(conf *GuildActivityGuildActivityCfg)bool){
	for _,v := range gameDb.GuildActivityGuildActivityCfgs{
		if !f(v){
			return
		}
	}
}

func GetGuildAuctionGuildAuctionCfg( Id int) *GuildAuctionGuildAuctionCfg {
	return gameDb.GuildAuctionGuildAuctionCfgs[Id]
}

func RangGuildAuctionGuildAuctionCfgs(f func(conf *GuildAuctionGuildAuctionCfg)bool){
	for _,v := range gameDb.GuildAuctionGuildAuctionCfgs{
		if !f(v){
			return
		}
	}
}

func GetGuildAutoCreateGuildAutoCreateCfg( Id int) *GuildAutoCreateGuildAutoCreateCfg {
	return gameDb.GuildAutoCreateGuildAutoCreateCfgs[Id]
}

func RangGuildAutoCreateGuildAutoCreateCfgs(f func(conf *GuildAutoCreateGuildAutoCreateCfg)bool){
	for _,v := range gameDb.GuildAutoCreateGuildAutoCreateCfgs{
		if !f(v){
			return
		}
	}
}

func GetGuildBonfireGuildBonfireCfg( Id int) *GuildBonfireGuildBonfireCfg {
	return gameDb.GuildBonfireGuildBonfireCfgs[Id]
}

func RangGuildBonfireGuildBonfireCfgs(f func(conf *GuildBonfireGuildBonfireCfg)bool){
	for _,v := range gameDb.GuildBonfireGuildBonfireCfgs{
		if !f(v){
			return
		}
	}
}

func GetGuildLevelGuildLevelCfg( Id int) *GuildLevelGuildLevelCfg {
	return gameDb.GuildLevelGuildLevelCfgs[Id]
}

func RangGuildLevelGuildLevelCfgs(f func(conf *GuildLevelGuildLevelCfg)bool){
	for _,v := range gameDb.GuildLevelGuildLevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetGuildNameGuildNameCfg( Id int) *GuildNameGuildNameCfg {
	return gameDb.GuildNameGuildNameCfgs[Id]
}

func RangGuildNameGuildNameCfgs(f func(conf *GuildNameGuildNameCfg)bool){
	for _,v := range gameDb.GuildNameGuildNameCfgs{
		if !f(v){
			return
		}
	}
}

func GetGuildRobotGuildRobotCfg( Id int) *GuildRobotGuildRobotCfg {
	return gameDb.GuildRobotGuildRobotCfgs[Id]
}

func RangGuildRobotGuildRobotCfgs(f func(conf *GuildRobotGuildRobotCfg)bool){
	for _,v := range gameDb.GuildRobotGuildRobotCfgs{
		if !f(v){
			return
		}
	}
}

func GetHellBossHellBossCfg( Id int) *HellBossHellBossCfg {
	return gameDb.HellBossHellBossCfgs[Id]
}

func RangHellBossHellBossCfgs(f func(conf *HellBossHellBossCfg)bool){
	for _,v := range gameDb.HellBossHellBossCfgs{
		if !f(v){
			return
		}
	}
}

func GetHellBossFloorHellBossFloorCfg( Map int) *HellBossFloorHellBossFloorCfg {
	return gameDb.HellBossFloorHellBossFloorCfgs[Map]
}

func RangHellBossFloorHellBossFloorCfgs(f func(conf *HellBossFloorHellBossFloorCfg)bool){
	for _,v := range gameDb.HellBossFloorHellBossFloorCfgs{
		if !f(v){
			return
		}
	}
}

func GetHolyArmsHolyArmsCfg( Id int) *HolyArmsHolyArmsCfg {
	return gameDb.HolyArmsHolyArmsCfgs[Id]
}

func RangHolyArmsHolyArmsCfgs(f func(conf *HolyArmsHolyArmsCfg)bool){
	for _,v := range gameDb.HolyArmsHolyArmsCfgs{
		if !f(v){
			return
		}
	}
}

func GetHolyBeastHolyBeastCfg( Id int) *HolyBeastHolyBeastCfg {
	return gameDb.HolyBeastHolyBeastCfgs[Id]
}

func RangHolyBeastHolyBeastCfgs(f func(conf *HolyBeastHolyBeastCfg)bool){
	for _,v := range gameDb.HolyBeastHolyBeastCfgs{
		if !f(v){
			return
		}
	}
}

func GetHolySkillHolySkillCfg( Id int) *HolySkillHolySkillCfg {
	return gameDb.HolySkillHolySkillCfgs[Id]
}

func RangHolySkillHolySkillCfgs(f func(conf *HolySkillHolySkillCfg)bool){
	for _,v := range gameDb.HolySkillHolySkillCfgs{
		if !f(v){
			return
		}
	}
}

func GetHolylevelHolylevelCfg( Id int) *HolylevelHolylevelCfg {
	return gameDb.HolylevelHolylevelCfgs[Id]
}

func RangHolylevelHolylevelCfgs(f func(conf *HolylevelHolylevelCfg)bool){
	for _,v := range gameDb.HolylevelHolylevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetHookMapHookMapCfg( Stage_id int) *HookMapHookMapCfg {
	return gameDb.HookMapHookMapCfgs[Stage_id]
}

func RangHookMapHookMapCfgs(f func(conf *HookMapHookMapCfg)bool){
	for _,v := range gameDb.HookMapHookMapCfgs{
		if !f(v){
			return
		}
	}
}

func GetInsideArtInsideArtCfg( Id int) *InsideArtInsideArtCfg {
	return gameDb.InsideArtInsideArtCfgs[Id]
}

func RangInsideArtInsideArtCfgs(f func(conf *InsideArtInsideArtCfg)bool){
	for _,v := range gameDb.InsideArtInsideArtCfgs{
		if !f(v){
			return
		}
	}
}

func GetInsideGradeInsideGradeCfg( Grade int) *InsideGradeInsideGradeCfg {
	return gameDb.InsideGradeInsideGradeCfgs[Grade]
}

func RangInsideGradeInsideGradeCfgs(f func(conf *InsideGradeInsideGradeCfg)bool){
	for _,v := range gameDb.InsideGradeInsideGradeCfgs{
		if !f(v){
			return
		}
	}
}

func GetInsideSkillInsideSkillCfg( Id int) *InsideSkillInsideSkillCfg {
	return gameDb.InsideSkillInsideSkillCfgs[Id]
}

func RangInsideSkillInsideSkillCfgs(f func(conf *InsideSkillInsideSkillCfg)bool){
	for _,v := range gameDb.InsideSkillInsideSkillCfgs{
		if !f(v){
			return
		}
	}
}

func GetInsideStarInsideStarCfg( Id int) *InsideStarInsideStarCfg {
	return gameDb.InsideStarInsideStarCfgs[Id]
}

func RangInsideStarInsideStarCfgs(f func(conf *InsideStarInsideStarCfg)bool){
	for _,v := range gameDb.InsideStarInsideStarCfgs{
		if !f(v){
			return
		}
	}
}

func GetItemBaseCfg( Id int) *ItemBaseCfg {
	return gameDb.ItemBaseCfgs[Id]
}

func RangItemBaseCfgs(f func(conf *ItemBaseCfg)bool){
	for _,v := range gameDb.ItemBaseCfgs{
		if !f(v){
			return
		}
	}
}

func GetJewelJewelCfg( Id int) *JewelJewelCfg {
	return gameDb.JewelJewelCfgs[Id]
}

func RangJewelJewelCfgs(f func(conf *JewelJewelCfg)bool){
	for _,v := range gameDb.JewelJewelCfgs{
		if !f(v){
			return
		}
	}
}

func GetJewelBodyJewelBodyCfg( Body int) *JewelBodyJewelBodyCfg {
	return gameDb.JewelBodyJewelBodyCfgs[Body]
}

func RangJewelBodyJewelBodyCfgs(f func(conf *JewelBodyJewelBodyCfg)bool){
	for _,v := range gameDb.JewelBodyJewelBodyCfgs{
		if !f(v){
			return
		}
	}
}

func GetJewelSuitJewelSuitCfg( Sum int) *JewelSuitJewelSuitCfg {
	return gameDb.JewelSuitJewelSuitCfgs[Sum]
}

func RangJewelSuitJewelSuitCfgs(f func(conf *JewelSuitJewelSuitCfg)bool){
	for _,v := range gameDb.JewelSuitJewelSuitCfgs{
		if !f(v){
			return
		}
	}
}

func GetJuexueLevelConfCfg( Id int) *JuexueLevelConfCfg {
	return gameDb.JuexueLevelConfCfgs[Id]
}

func RangJuexueLevelConfCfgs(f func(conf *JuexueLevelConfCfg)bool){
	for _,v := range gameDb.JuexueLevelConfCfgs{
		if !f(v){
			return
		}
	}
}

func GetKingarmsKingarmsCfg( Id int) *KingarmsKingarmsCfg {
	return gameDb.KingarmsKingarmsCfgs[Id]
}

func RangKingarmsKingarmsCfgs(f func(conf *KingarmsKingarmsCfg)bool){
	for _,v := range gameDb.KingarmsKingarmsCfgs{
		if !f(v){
			return
		}
	}
}

func GetKuafushabakeRewardserverKuafushabakeRewardserverCfg( Id int) *KuafushabakeRewardserverKuafushabakeRewardserverCfg {
	return gameDb.KuafushabakeRewardserverKuafushabakeRewardserverCfgs[Id]
}

func RangKuafushabakeRewardserverKuafushabakeRewardserverCfgs(f func(conf *KuafushabakeRewardserverKuafushabakeRewardserverCfg)bool){
	for _,v := range gameDb.KuafushabakeRewardserverKuafushabakeRewardserverCfgs{
		if !f(v){
			return
		}
	}
}

func GetKuafushabakeRewarduniKuafushabakeRewarduniCfg( Id int) *KuafushabakeRewarduniKuafushabakeRewarduniCfg {
	return gameDb.KuafushabakeRewarduniKuafushabakeRewarduniCfgs[Id]
}

func RangKuafushabakeRewarduniKuafushabakeRewarduniCfgs(f func(conf *KuafushabakeRewarduniKuafushabakeRewarduniCfg)bool){
	for _,v := range gameDb.KuafushabakeRewarduniKuafushabakeRewarduniCfgs{
		if !f(v){
			return
		}
	}
}

func GetLabelLabelCfg( Id int) *LabelLabelCfg {
	return gameDb.LabelLabelCfgs[Id]
}

func RangLabelLabelCfgs(f func(conf *LabelLabelCfg)bool){
	for _,v := range gameDb.LabelLabelCfgs{
		if !f(v){
			return
		}
	}
}

func GetLabelTaskLabelTaskCfg( Id int) *LabelTaskLabelTaskCfg {
	return gameDb.LabelTaskLabelTaskCfgs[Id]
}

func RangLabelTaskLabelTaskCfgs(f func(conf *LabelTaskLabelTaskCfg)bool){
	for _,v := range gameDb.LabelTaskLabelTaskCfgs{
		if !f(v){
			return
		}
	}
}

func GetLimitedGiftLimitedGiftCfg( Id int) *LimitedGiftLimitedGiftCfg {
	return gameDb.LimitedGiftLimitedGiftCfgs[Id]
}

func RangLimitedGiftLimitedGiftCfgs(f func(conf *LimitedGiftLimitedGiftCfg)bool){
	for _,v := range gameDb.LimitedGiftLimitedGiftCfgs{
		if !f(v){
			return
		}
	}
}

func GetLuckyLuckyCfg( Id int) *LuckyLuckyCfg {
	return gameDb.LuckyLuckyCfgs[Id]
}

func RangLuckyLuckyCfgs(f func(conf *LuckyLuckyCfg)bool){
	for _,v := range gameDb.LuckyLuckyCfgs{
		if !f(v){
			return
		}
	}
}

func GetMagicCircleMagicCircleCfg( Id int) *MagicCircleMagicCircleCfg {
	return gameDb.MagicCircleMagicCircleCfgs[Id]
}

func RangMagicCircleMagicCircleCfgs(f func(conf *MagicCircleMagicCircleCfg)bool){
	for _,v := range gameDb.MagicCircleMagicCircleCfgs{
		if !f(v){
			return
		}
	}
}

func GetMagicCircleLevelMagicCircleLevelCfg( Id int) *MagicCircleLevelMagicCircleLevelCfg {
	return gameDb.MagicCircleLevelMagicCircleLevelCfgs[Id]
}

func RangMagicCircleLevelMagicCircleLevelCfgs(f func(conf *MagicCircleLevelMagicCircleLevelCfg)bool){
	for _,v := range gameDb.MagicCircleLevelMagicCircleLevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetMagicTowerMagicTowerCfg( Id int) *MagicTowerMagicTowerCfg {
	return gameDb.MagicTowerMagicTowerCfgs[Id]
}

func RangMagicTowerMagicTowerCfgs(f func(conf *MagicTowerMagicTowerCfg)bool){
	for _,v := range gameDb.MagicTowerMagicTowerCfgs{
		if !f(v){
			return
		}
	}
}

func GetMagicTowerRewardMagicTowerRewardCfg( Id int) *MagicTowerRewardMagicTowerRewardCfg {
	return gameDb.MagicTowerRewardMagicTowerRewardCfgs[Id]
}

func RangMagicTowerRewardMagicTowerRewardCfgs(f func(conf *MagicTowerRewardMagicTowerRewardCfg)bool){
	for _,v := range gameDb.MagicTowerRewardMagicTowerRewardCfgs{
		if !f(v){
			return
		}
	}
}

func GetMailMailCfg( Id int) *MailMailCfg {
	return gameDb.MailMailCfgs[Id]
}

func RangMailMailCfgs(f func(conf *MailMailCfg)bool){
	for _,v := range gameDb.MailMailCfgs{
		if !f(v){
			return
		}
	}
}

func GetMainPrMainPrCfg( Id int) *MainPrMainPrCfg {
	return gameDb.MainPrMainPrCfgs[Id]
}

func RangMainPrMainPrCfgs(f func(conf *MainPrMainPrCfg)bool){
	for _,v := range gameDb.MainPrMainPrCfgs{
		if !f(v){
			return
		}
	}
}

func GetMapMapCfg( Id int) *MapMapCfg {
	return gameDb.MapMapCfgs[Id]
}

func RangMapMapCfgs(f func(conf *MapMapCfg)bool){
	for _,v := range gameDb.MapMapCfgs{
		if !f(v){
			return
		}
	}
}

func GetMaptypeGameCfg( Id int) *MaptypeGameCfg {
	return gameDb.MaptypeGameCfgs[Id]
}

func RangMaptypeGameCfgs(f func(conf *MaptypeGameCfg)bool){
	for _,v := range gameDb.MaptypeGameCfgs{
		if !f(v){
			return
		}
	}
}

func GetMaterialCostMaterialCostCfg( Number int) *MaterialCostMaterialCostCfg {
	return gameDb.MaterialCostMaterialCostCfgs[Number]
}

func RangMaterialCostMaterialCostCfgs(f func(conf *MaterialCostMaterialCostCfg)bool){
	for _,v := range gameDb.MaterialCostMaterialCostCfgs{
		if !f(v){
			return
		}
	}
}

func GetMaterialHomeMaterialHomeCfg( Type int) *MaterialHomeMaterialHomeCfg {
	return gameDb.MaterialHomeMaterialHomeCfgs[Type]
}

func RangMaterialHomeMaterialHomeCfgs(f func(conf *MaterialHomeMaterialHomeCfg)bool){
	for _,v := range gameDb.MaterialHomeMaterialHomeCfgs{
		if !f(v){
			return
		}
	}
}

func GetMaterialStageMaterialStageCfg( Id int) *MaterialStageMaterialStageCfg {
	return gameDb.MaterialStageMaterialStageCfgs[Id]
}

func RangMaterialStageMaterialStageCfgs(f func(conf *MaterialStageMaterialStageCfg)bool){
	for _,v := range gameDb.MaterialStageMaterialStageCfgs{
		if !f(v){
			return
		}
	}
}

func GetMijiMijiCfg( Id int) *MijiMijiCfg {
	return gameDb.MijiMijiCfgs[Id]
}

func RangMijiMijiCfgs(f func(conf *MijiMijiCfg)bool){
	for _,v := range gameDb.MijiMijiCfgs{
		if !f(v){
			return
		}
	}
}

func GetMijiLevelMijiLevelCfg( Id int) *MijiLevelMijiLevelCfg {
	return gameDb.MijiLevelMijiLevelCfgs[Id]
}

func RangMijiLevelMijiLevelCfgs(f func(conf *MijiLevelMijiLevelCfg)bool){
	for _,v := range gameDb.MijiLevelMijiLevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetMijiTypeMijiTypeCfg( Id int) *MijiTypeMijiTypeCfg {
	return gameDb.MijiTypeMijiTypeCfgs[Id]
}

func RangMijiTypeMijiTypeCfgs(f func(conf *MijiTypeMijiTypeCfg)bool){
	for _,v := range gameDb.MijiTypeMijiTypeCfgs{
		if !f(v){
			return
		}
	}
}

func GetMiningMiningCfg( Id int) *MiningMiningCfg {
	return gameDb.MiningMiningCfgs[Id]
}

func RangMiningMiningCfgs(f func(conf *MiningMiningCfg)bool){
	for _,v := range gameDb.MiningMiningCfgs{
		if !f(v){
			return
		}
	}
}

func GetMonsterMonsterCfg( Monsterid int) *MonsterMonsterCfg {
	return gameDb.MonsterMonsterCfgs[Monsterid]
}

func RangMonsterMonsterCfgs(f func(conf *MonsterMonsterCfg)bool){
	for _,v := range gameDb.MonsterMonsterCfgs{
		if !f(v){
			return
		}
	}
}

func GetMonsterdropDropCfg( Dropid int) *MonsterdropDropCfg {
	return gameDb.MonsterdropDropCfgs[Dropid]
}

func RangMonsterdropDropCfgs(f func(conf *MonsterdropDropCfg)bool){
	for _,v := range gameDb.MonsterdropDropCfgs{
		if !f(v){
			return
		}
	}
}

func GetMonstergroupMonstergroupCfg( Groupid int) *MonstergroupMonstergroupCfg {
	return gameDb.MonstergroupMonstergroupCfgs[Groupid]
}

func RangMonstergroupMonstergroupCfgs(f func(conf *MonstergroupMonstergroupCfg)bool){
	for _,v := range gameDb.MonstergroupMonstergroupCfgs{
		if !f(v){
			return
		}
	}
}

func GetMonthCardMonthCardCfg( Id int) *MonthCardMonthCardCfg {
	return gameDb.MonthCardMonthCardCfgs[Id]
}

func RangMonthCardMonthCardCfgs(f func(conf *MonthCardMonthCardCfg)bool){
	for _,v := range gameDb.MonthCardMonthCardCfgs{
		if !f(v){
			return
		}
	}
}

func GetMonthCardPrivilegeMonthCardPrivilegeCfg( Id int) *MonthCardPrivilegeMonthCardPrivilegeCfg {
	return gameDb.MonthCardPrivilegeMonthCardPrivilegeCfgs[Id]
}

func RangMonthCardPrivilegeMonthCardPrivilegeCfgs(f func(conf *MonthCardPrivilegeMonthCardPrivilegeCfg)bool){
	for _,v := range gameDb.MonthCardPrivilegeMonthCardPrivilegeCfgs{
		if !f(v){
			return
		}
	}
}

func GetNpcMonsterCfg( Id int) *NpcMonsterCfg {
	return gameDb.NpcMonsterCfgs[Id]
}

func RangNpcMonsterCfgs(f func(conf *NpcMonsterCfg)bool){
	for _,v := range gameDb.NpcMonsterCfgs{
		if !f(v){
			return
		}
	}
}

func GetOfficialOfficialCfg( Id int) *OfficialOfficialCfg {
	return gameDb.OfficialOfficialCfgs[Id]
}

func RangOfficialOfficialCfgs(f func(conf *OfficialOfficialCfg)bool){
	for _,v := range gameDb.OfficialOfficialCfgs{
		if !f(v){
			return
		}
	}
}

func GetOpenGiftOpenGiftCfg( Id int) *OpenGiftOpenGiftCfg {
	return gameDb.OpenGiftOpenGiftCfgs[Id]
}

func RangOpenGiftOpenGiftCfgs(f func(conf *OpenGiftOpenGiftCfg)bool){
	for _,v := range gameDb.OpenGiftOpenGiftCfgs{
		if !f(v){
			return
		}
	}
}

func GetPanaceaPanaceaCfg( Id int) *PanaceaPanaceaCfg {
	return gameDb.PanaceaPanaceaCfgs[Id]
}

func RangPanaceaPanaceaCfgs(f func(conf *PanaceaPanaceaCfg)bool){
	for _,v := range gameDb.PanaceaPanaceaCfgs{
		if !f(v){
			return
		}
	}
}

func GetPersonalBossPersonalBossCfg( Id int) *PersonalBossPersonalBossCfg {
	return gameDb.PersonalBossPersonalBossCfgs[Id]
}

func RangPersonalBossPersonalBossCfgs(f func(conf *PersonalBossPersonalBossCfg)bool){
	for _,v := range gameDb.PersonalBossPersonalBossCfgs{
		if !f(v){
			return
		}
	}
}

func GetPetsConfCfg( Id int) *PetsConfCfg {
	return gameDb.PetsConfCfgs[Id]
}

func RangPetsConfCfgs(f func(conf *PetsConfCfg)bool){
	for _,v := range gameDb.PetsConfCfgs{
		if !f(v){
			return
		}
	}
}

func GetPetsAddPetsAddCfg( Id int) *PetsAddPetsAddCfg {
	return gameDb.PetsAddPetsAddCfgs[Id]
}

func RangPetsAddPetsAddCfgs(f func(conf *PetsAddPetsAddCfg)bool){
	for _,v := range gameDb.PetsAddPetsAddCfgs{
		if !f(v){
			return
		}
	}
}

func GetPetsAddSkillPetsAddSkillCfg( Id int) *PetsAddSkillPetsAddSkillCfg {
	return gameDb.PetsAddSkillPetsAddSkillCfgs[Id]
}

func RangPetsAddSkillPetsAddSkillCfgs(f func(conf *PetsAddSkillPetsAddSkillCfg)bool){
	for _,v := range gameDb.PetsAddSkillPetsAddSkillCfgs{
		if !f(v){
			return
		}
	}
}

func GetPetsBreakConfCfg( Id int) *PetsBreakConfCfg {
	return gameDb.PetsBreakConfCfgs[Id]
}

func RangPetsBreakConfCfgs(f func(conf *PetsBreakConfCfg)bool){
	for _,v := range gameDb.PetsBreakConfCfgs{
		if !f(v){
			return
		}
	}
}

func GetPetsGradeConfCfg( Id int) *PetsGradeConfCfg {
	return gameDb.PetsGradeConfCfgs[Id]
}

func RangPetsGradeConfCfgs(f func(conf *PetsGradeConfCfg)bool){
	for _,v := range gameDb.PetsGradeConfCfgs{
		if !f(v){
			return
		}
	}
}

func GetPetsLevelConfCfg( Id int) *PetsLevelConfCfg {
	return gameDb.PetsLevelConfCfgs[Id]
}

func RangPetsLevelConfCfgs(f func(conf *PetsLevelConfCfg)bool){
	for _,v := range gameDb.PetsLevelConfCfgs{
		if !f(v){
			return
		}
	}
}

func GetPhantomPhantomCfg( Phantom int) *PhantomPhantomCfg {
	return gameDb.PhantomPhantomCfgs[Phantom]
}

func RangPhantomPhantomCfgs(f func(conf *PhantomPhantomCfg)bool){
	for _,v := range gameDb.PhantomPhantomCfgs{
		if !f(v){
			return
		}
	}
}

func GetPhantomLevelPhantomLevelCfg( Id int) *PhantomLevelPhantomLevelCfg {
	return gameDb.PhantomLevelPhantomLevelCfgs[Id]
}

func RangPhantomLevelPhantomLevelCfgs(f func(conf *PhantomLevelPhantomLevelCfg)bool){
	for _,v := range gameDb.PhantomLevelPhantomLevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetPowerRollPowerRollCfg( Id int) *PowerRollPowerRollCfg {
	return gameDb.PowerRollPowerRollCfgs[Id]
}

func RangPowerRollPowerRollCfgs(f func(conf *PowerRollPowerRollCfg)bool){
	for _,v := range gameDb.PowerRollPowerRollCfgs{
		if !f(v){
			return
		}
	}
}

func GetPreFunctionPreFunctionCfg( Id int) *PreFunctionPreFunctionCfg {
	return gameDb.PreFunctionPreFunctionCfgs[Id]
}

func RangPreFunctionPreFunctionCfgs(f func(conf *PreFunctionPreFunctionCfg)bool){
	for _,v := range gameDb.PreFunctionPreFunctionCfgs{
		if !f(v){
			return
		}
	}
}

func GetPrivilegePrivilegeCfg( Id int) *PrivilegePrivilegeCfg {
	return gameDb.PrivilegePrivilegeCfgs[Id]
}

func RangPrivilegePrivilegeCfgs(f func(conf *PrivilegePrivilegeCfg)bool){
	for _,v := range gameDb.PrivilegePrivilegeCfgs{
		if !f(v){
			return
		}
	}
}

func GetPropertyPropertyCfg( Id int) *PropertyPropertyCfg {
	return gameDb.PropertyPropertyCfgs[Id]
}

func RangPropertyPropertyCfgs(f func(conf *PropertyPropertyCfg)bool){
	for _,v := range gameDb.PropertyPropertyCfgs{
		if !f(v){
			return
		}
	}
}

func GetPublicCopyStageCfg( StageId int) *PublicCopyStageCfg {
	return gameDb.PublicCopyStageCfgs[StageId]
}

func RangPublicCopyStageCfgs(f func(conf *PublicCopyStageCfg)bool){
	for _,v := range gameDb.PublicCopyStageCfgs{
		if !f(v){
			return
		}
	}
}

func GetRandRandCfg( Id int) *RandRandCfg {
	return gameDb.RandRandCfgs[Id]
}

func RangRandRandCfgs(f func(conf *RandRandCfg)bool){
	for _,v := range gameDb.RandRandCfgs{
		if !f(v){
			return
		}
	}
}

func GetRechargeRechargeCfg( Id int) *RechargeRechargeCfg {
	return gameDb.RechargeRechargeCfgs[Id]
}

func RangRechargeRechargeCfgs(f func(conf *RechargeRechargeCfg)bool){
	for _,v := range gameDb.RechargeRechargeCfgs{
		if !f(v){
			return
		}
	}
}

func GetRedDayMaxRedDayMaxCfg( Id int) *RedDayMaxRedDayMaxCfg {
	return gameDb.RedDayMaxRedDayMaxCfgs[Id]
}

func RangRedDayMaxRedDayMaxCfgs(f func(conf *RedDayMaxRedDayMaxCfg)bool){
	for _,v := range gameDb.RedDayMaxRedDayMaxCfgs{
		if !f(v){
			return
		}
	}
}

func GetRedRecoveryRedRecoveryCfg( Id int) *RedRecoveryRedRecoveryCfg {
	return gameDb.RedRecoveryRedRecoveryCfgs[Id]
}

func RangRedRecoveryRedRecoveryCfgs(f func(conf *RedRecoveryRedRecoveryCfg)bool){
	for _,v := range gameDb.RedRecoveryRedRecoveryCfgs{
		if !f(v){
			return
		}
	}
}

func GetReinReinCfg( Id int) *ReinReinCfg {
	return gameDb.ReinReinCfgs[Id]
}

func RangReinReinCfgs(f func(conf *ReinReinCfg)bool){
	for _,v := range gameDb.ReinReinCfgs{
		if !f(v){
			return
		}
	}
}

func GetReinCostReinCostCfg( Id int) *ReinCostReinCostCfg {
	return gameDb.ReinCostReinCostCfgs[Id]
}

func RangReinCostReinCostCfgs(f func(conf *ReinCostReinCostCfg)bool){
	for _,v := range gameDb.ReinCostReinCostCfgs{
		if !f(v){
			return
		}
	}
}

func GetRewardsOnlineAwardCfg( Id int) *RewardsOnlineAwardCfg {
	return gameDb.RewardsOnlineAwardCfgs[Id]
}

func RangRewardsOnlineAwardCfgs(f func(conf *RewardsOnlineAwardCfg)bool){
	for _,v := range gameDb.RewardsOnlineAwardCfgs{
		if !f(v){
			return
		}
	}
}

func GetRingRingCfg( Ringid int) *RingRingCfg {
	return gameDb.RingRingCfgs[Ringid]
}

func RangRingRingCfgs(f func(conf *RingRingCfg)bool){
	for _,v := range gameDb.RingRingCfgs{
		if !f(v){
			return
		}
	}
}

func GetRingPhantomRingPhantomCfg( Id int) *RingPhantomRingPhantomCfg {
	return gameDb.RingPhantomRingPhantomCfgs[Id]
}

func RangRingPhantomRingPhantomCfgs(f func(conf *RingPhantomRingPhantomCfg)bool){
	for _,v := range gameDb.RingPhantomRingPhantomCfgs{
		if !f(v){
			return
		}
	}
}

func GetRingStrengthenRingStrengthenCfg( Level int) *RingStrengthenRingStrengthenCfg {
	return gameDb.RingStrengthenRingStrengthenCfgs[Level]
}

func RangRingStrengthenRingStrengthenCfgs(f func(conf *RingStrengthenRingStrengthenCfg)bool){
	for _,v := range gameDb.RingStrengthenRingStrengthenCfgs{
		if !f(v){
			return
		}
	}
}

func GetRobotRobotCfg( Id int) *RobotRobotCfg {
	return gameDb.RobotRobotCfgs[Id]
}

func RangRobotRobotCfgs(f func(conf *RobotRobotCfg)bool){
	for _,v := range gameDb.RobotRobotCfgs{
		if !f(v){
			return
		}
	}
}

func GetRoleFirstnameRoleFirstnameCfg( Id int) *RoleFirstnameRoleFirstnameCfg {
	return gameDb.RoleFirstnameRoleFirstnameCfgs[Id]
}

func RangRoleFirstnameRoleFirstnameCfgs(f func(conf *RoleFirstnameRoleFirstnameCfg)bool){
	for _,v := range gameDb.RoleFirstnameRoleFirstnameCfgs{
		if !f(v){
			return
		}
	}
}

func GetRoleNameBaseCfg( Id int) *RoleNameBaseCfg {
	return gameDb.RoleNameBaseCfgs[Id]
}

func RangRoleNameBaseCfgs(f func(conf *RoleNameBaseCfg)bool){
	for _,v := range gameDb.RoleNameBaseCfgs{
		if !f(v){
			return
		}
	}
}

func GetScrollingScrollingCfg( Type int) *ScrollingScrollingCfg {
	return gameDb.ScrollingScrollingCfgs[Type]
}

func RangScrollingScrollingCfgs(f func(conf *ScrollingScrollingCfg)bool){
	for _,v := range gameDb.ScrollingScrollingCfgs{
		if !f(v){
			return
		}
	}
}

func GetSetTypeCfg( Id int) *SetTypeCfg {
	return gameDb.SetTypeCfgs[Id]
}

func RangSetTypeCfgs(f func(conf *SetTypeCfg)bool){
	for _,v := range gameDb.SetTypeCfgs{
		if !f(v){
			return
		}
	}
}

func GetSevenDayInvestSevenDayInvestCfg( Id int) *SevenDayInvestSevenDayInvestCfg {
	return gameDb.SevenDayInvestSevenDayInvestCfgs[Id]
}

func RangSevenDayInvestSevenDayInvestCfgs(f func(conf *SevenDayInvestSevenDayInvestCfg)bool){
	for _,v := range gameDb.SevenDayInvestSevenDayInvestCfgs{
		if !f(v){
			return
		}
	}
}

func GetShabakeRewardperShabakeRewardperCfg( Id int) *ShabakeRewardperShabakeRewardperCfg {
	return gameDb.ShabakeRewardperShabakeRewardperCfgs[Id]
}

func RangShabakeRewardperShabakeRewardperCfgs(f func(conf *ShabakeRewardperShabakeRewardperCfg)bool){
	for _,v := range gameDb.ShabakeRewardperShabakeRewardperCfgs{
		if !f(v){
			return
		}
	}
}

func GetShabakeRewarduniShabakeRewarduniCfg( Id int) *ShabakeRewarduniShabakeRewarduniCfg {
	return gameDb.ShabakeRewarduniShabakeRewarduniCfgs[Id]
}

func RangShabakeRewarduniShabakeRewarduniCfgs(f func(conf *ShabakeRewarduniShabakeRewarduniCfg)bool){
	for _,v := range gameDb.ShabakeRewarduniShabakeRewarduniCfgs{
		if !f(v){
			return
		}
	}
}

func GetShopTypeCfg( Id int) *ShopTypeCfg {
	return gameDb.ShopTypeCfgs[Id]
}

func RangShopTypeCfgs(f func(conf *ShopTypeCfg)bool){
	for _,v := range gameDb.ShopTypeCfgs{
		if !f(v){
			return
		}
	}
}

func GetShopItemUnitCfg( Id int) *ShopItemUnitCfg {
	return gameDb.ShopItemUnitCfgs[Id]
}

func RangShopItemUnitCfgs(f func(conf *ShopItemUnitCfg)bool){
	for _,v := range gameDb.ShopItemUnitCfgs{
		if !f(v){
			return
		}
	}
}

func GetSignSignCfg( Id int) *SignSignCfg {
	return gameDb.SignSignCfgs[Id]
}

func RangSignSignCfgs(f func(conf *SignSignCfg)bool){
	for _,v := range gameDb.SignSignCfgs{
		if !f(v){
			return
		}
	}
}

func GetSkillSkillCfg( Skillid int) *SkillSkillCfg {
	return gameDb.SkillSkillCfgs[Skillid]
}

func RangSkillSkillCfgs(f func(conf *SkillSkillCfg)bool){
	for _,v := range gameDb.SkillSkillCfgs{
		if !f(v){
			return
		}
	}
}

func GetSkillAttackEffectSkillAttackEffectCfg( Id int) *SkillAttackEffectSkillAttackEffectCfg {
	return gameDb.SkillAttackEffectSkillAttackEffectCfgs[Id]
}

func RangSkillAttackEffectSkillAttackEffectCfgs(f func(conf *SkillAttackEffectSkillAttackEffectCfg)bool){
	for _,v := range gameDb.SkillAttackEffectSkillAttackEffectCfgs{
		if !f(v){
			return
		}
	}
}

func GetSkillLevelSkillCfg( Skillid int) *SkillLevelSkillCfg {
	return gameDb.SkillLevelSkillCfgs[Skillid]
}

func RangSkillLevelSkillCfgs(f func(conf *SkillLevelSkillCfg)bool){
	for _,v := range gameDb.SkillLevelSkillCfgs{
		if !f(v){
			return
		}
	}
}

func GetStageStageCfg( Id int) *StageStageCfg {
	return gameDb.StageStageCfgs[Id]
}

func RangStageStageCfgs(f func(conf *StageStageCfg)bool){
	for _,v := range gameDb.StageStageCfgs{
		if !f(v){
			return
		}
	}
}

func GetStrengthenStrengthenCfg( Id int) *StrengthenStrengthenCfg {
	return gameDb.StrengthenStrengthenCfgs[Id]
}

func RangStrengthenStrengthenCfgs(f func(conf *StrengthenStrengthenCfg)bool){
	for _,v := range gameDb.StrengthenStrengthenCfgs{
		if !f(v){
			return
		}
	}
}

func GetStrengthenlinkStrengthenCfg( Id int) *StrengthenlinkStrengthenCfg {
	return gameDb.StrengthenlinkStrengthenCfgs[Id]
}

func RangStrengthenlinkStrengthenCfgs(f func(conf *StrengthenlinkStrengthenCfg)bool){
	for _,v := range gameDb.StrengthenlinkStrengthenCfgs{
		if !f(v){
			return
		}
	}
}

func GetSummonConfCfg( Id int) *SummonConfCfg {
	return gameDb.SummonConfCfgs[Id]
}

func RangSummonConfCfgs(f func(conf *SummonConfCfg)bool){
	for _,v := range gameDb.SummonConfCfgs{
		if !f(v){
			return
		}
	}
}

func GetTalentTalentCfg( Id int) *TalentTalentCfg {
	return gameDb.TalentTalentCfgs[Id]
}

func RangTalentTalentCfgs(f func(conf *TalentTalentCfg)bool){
	for _,v := range gameDb.TalentTalentCfgs{
		if !f(v){
			return
		}
	}
}

func GetTalentGetTalentGetCfg( Id int) *TalentGetTalentGetCfg {
	return gameDb.TalentGetTalentGetCfgs[Id]
}

func RangTalentGetTalentGetCfgs(f func(conf *TalentGetTalentGetCfg)bool){
	for _,v := range gameDb.TalentGetTalentGetCfgs{
		if !f(v){
			return
		}
	}
}

func GetTalentLevelTalentLevelCfg( Id int) *TalentLevelTalentLevelCfg {
	return gameDb.TalentLevelTalentLevelCfgs[Id]
}

func RangTalentLevelTalentLevelCfgs(f func(conf *TalentLevelTalentLevelCfg)bool){
	for _,v := range gameDb.TalentLevelTalentLevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetTalentStageTalengStageCfg( Id int) *TalentStageTalengStageCfg {
	return gameDb.TalentStageTalengStageCfgs[Id]
}

func RangTalentStageTalengStageCfgs(f func(conf *TalentStageTalengStageCfg)bool){
	for _,v := range gameDb.TalentStageTalengStageCfgs{
		if !f(v){
			return
		}
	}
}

func GetTalentWayTalengWayCfg( Id int) *TalentWayTalengWayCfg {
	return gameDb.TalentWayTalengWayCfgs[Id]
}

func RangTalentWayTalengWayCfgs(f func(conf *TalentWayTalengWayCfg)bool){
	for _,v := range gameDb.TalentWayTalengWayCfgs{
		if !f(v){
			return
		}
	}
}

func GetTalenteffectTalentCfg( Id int) *TalenteffectTalentCfg {
	return gameDb.TalenteffectTalentCfgs[Id]
}

func RangTalenteffectTalentCfgs(f func(conf *TalenteffectTalentCfg)bool){
	for _,v := range gameDb.TalenteffectTalentCfgs{
		if !f(v){
			return
		}
	}
}

func GetTalentgeneralTalentCfg( Id int) *TalentgeneralTalentCfg {
	return gameDb.TalentgeneralTalentCfgs[Id]
}

func RangTalentgeneralTalentCfgs(f func(conf *TalentgeneralTalentCfg)bool){
	for _,v := range gameDb.TalentgeneralTalentCfgs{
		if !f(v){
			return
		}
	}
}

func GetTaskConditionCfg( Id int) *TaskConditionCfg {
	return gameDb.TaskConditionCfgs[Id]
}

func RangTaskConditionCfgs(f func(conf *TaskConditionCfg)bool){
	for _,v := range gameDb.TaskConditionCfgs{
		if !f(v){
			return
		}
	}
}

func GetTitleTitleCfg( Id int) *TitleTitleCfg {
	return gameDb.TitleTitleCfgs[Id]
}

func RangTitleTitleCfgs(f func(conf *TitleTitleCfg)bool){
	for _,v := range gameDb.TitleTitleCfgs{
		if !f(v){
			return
		}
	}
}

func GetTowerTowerCfg( Id int) *TowerTowerCfg {
	return gameDb.TowerTowerCfgs[Id]
}

func RangTowerTowerCfgs(f func(conf *TowerTowerCfg)bool){
	for _,v := range gameDb.TowerTowerCfgs{
		if !f(v){
			return
		}
	}
}

func GetTowerLotteryCircleTowerLotteryCircleCfg( Id int) *TowerLotteryCircleTowerLotteryCircleCfg {
	return gameDb.TowerLotteryCircleTowerLotteryCircleCfgs[Id]
}

func RangTowerLotteryCircleTowerLotteryCircleCfgs(f func(conf *TowerLotteryCircleTowerLotteryCircleCfg)bool){
	for _,v := range gameDb.TowerLotteryCircleTowerLotteryCircleCfgs{
		if !f(v){
			return
		}
	}
}

func GetTowerRankRewardTowerRankRewardCfg( Id int) *TowerRankRewardTowerRankRewardCfg {
	return gameDb.TowerRankRewardTowerRankRewardCfgs[Id]
}

func RangTowerRankRewardTowerRankRewardCfgs(f func(conf *TowerRankRewardTowerRankRewardCfg)bool){
	for _,v := range gameDb.TowerRankRewardTowerRankRewardCfgs{
		if !f(v){
			return
		}
	}
}

func GetTowerRewardTowerRewardCfg( Id int) *TowerRewardTowerRewardCfg {
	return gameDb.TowerRewardTowerRewardCfgs[Id]
}

func RangTowerRewardTowerRewardCfgs(f func(conf *TowerRewardTowerRewardCfg)bool){
	for _,v := range gameDb.TowerRewardTowerRewardCfgs{
		if !f(v){
			return
		}
	}
}

func GetTreasureTreasureCfg( Id int) *TreasureTreasureCfg {
	return gameDb.TreasureTreasureCfgs[Id]
}

func RangTreasureTreasureCfgs(f func(conf *TreasureTreasureCfg)bool){
	for _,v := range gameDb.TreasureTreasureCfgs{
		if !f(v){
			return
		}
	}
}

func GetTreasureArtTreasureArtCfg( Id int) *TreasureArtTreasureArtCfg {
	return gameDb.TreasureArtTreasureArtCfgs[Id]
}

func RangTreasureArtTreasureArtCfgs(f func(conf *TreasureArtTreasureArtCfg)bool){
	for _,v := range gameDb.TreasureArtTreasureArtCfgs{
		if !f(v){
			return
		}
	}
}

func GetTreasureAwakenTreasureAwakenCfg( Id int) *TreasureAwakenTreasureAwakenCfg {
	return gameDb.TreasureAwakenTreasureAwakenCfgs[Id]
}

func RangTreasureAwakenTreasureAwakenCfgs(f func(conf *TreasureAwakenTreasureAwakenCfg)bool){
	for _,v := range gameDb.TreasureAwakenTreasureAwakenCfgs{
		if !f(v){
			return
		}
	}
}

func GetTreasureDiscountTreasureDiscountCfg( Id int) *TreasureDiscountTreasureDiscountCfg {
	return gameDb.TreasureDiscountTreasureDiscountCfgs[Id]
}

func RangTreasureDiscountTreasureDiscountCfgs(f func(conf *TreasureDiscountTreasureDiscountCfg)bool){
	for _,v := range gameDb.TreasureDiscountTreasureDiscountCfgs{
		if !f(v){
			return
		}
	}
}

func GetTreasureShopTreasureShopCfg( Id int) *TreasureShopTreasureShopCfg {
	return gameDb.TreasureShopTreasureShopCfgs[Id]
}

func RangTreasureShopTreasureShopCfgs(f func(conf *TreasureShopTreasureShopCfg)bool){
	for _,v := range gameDb.TreasureShopTreasureShopCfgs{
		if !f(v){
			return
		}
	}
}

func GetTreasureStarsTreasureStarsCfg( Id int) *TreasureStarsTreasureStarsCfg {
	return gameDb.TreasureStarsTreasureStarsCfgs[Id]
}

func RangTreasureStarsTreasureStarsCfgs(f func(conf *TreasureStarsTreasureStarsCfg)bool){
	for _,v := range gameDb.TreasureStarsTreasureStarsCfgs{
		if !f(v){
			return
		}
	}
}

func GetTreasureSuitTreasureSuitCfg( Id int) *TreasureSuitTreasureSuitCfg {
	return gameDb.TreasureSuitTreasureSuitCfgs[Id]
}

func RangTreasureSuitTreasureSuitCfgs(f func(conf *TreasureSuitTreasureSuitCfg)bool){
	for _,v := range gameDb.TreasureSuitTreasureSuitCfgs{
		if !f(v){
			return
		}
	}
}

func GetTrialTaskTrialTaskCfg( Id int) *TrialTaskTrialTaskCfg {
	return gameDb.TrialTaskTrialTaskCfgs[Id]
}

func RangTrialTaskTrialTaskCfgs(f func(conf *TrialTaskTrialTaskCfg)bool){
	for _,v := range gameDb.TrialTaskTrialTaskCfgs{
		if !f(v){
			return
		}
	}
}

func GetTrialTotalRewardTrialTotalRewardCfg( Id int) *TrialTotalRewardTrialTotalRewardCfg {
	return gameDb.TrialTotalRewardTrialTotalRewardCfgs[Id]
}

func RangTrialTotalRewardTrialTotalRewardCfgs(f func(conf *TrialTotalRewardTrialTotalRewardCfg)bool){
	for _,v := range gameDb.TrialTotalRewardTrialTotalRewardCfgs{
		if !f(v){
			return
		}
	}
}

func GetVipLvlCfg( Lvl int) *VipLvlCfg {
	return gameDb.VipLvlCfgs[Lvl]
}

func RangVipLvlCfgs(f func(conf *VipLvlCfg)bool){
	for _,v := range gameDb.VipLvlCfgs{
		if !f(v){
			return
		}
	}
}

func GetVipBossVipBossCfg( Id int) *VipBossVipBossCfg {
	return gameDb.VipBossVipBossCfgs[Id]
}

func RangVipBossVipBossCfgs(f func(conf *VipBossVipBossCfg)bool){
	for _,v := range gameDb.VipBossVipBossCfgs{
		if !f(v){
			return
		}
	}
}

func GetWarOrderConditionWarOrderConditionCfg( Id int) *WarOrderConditionWarOrderConditionCfg {
	return gameDb.WarOrderConditionWarOrderConditionCfgs[Id]
}

func RangWarOrderConditionWarOrderConditionCfgs(f func(conf *WarOrderConditionWarOrderConditionCfg)bool){
	for _,v := range gameDb.WarOrderConditionWarOrderConditionCfgs{
		if !f(v){
			return
		}
	}
}

func GetWarOrderCycleWarOrderCycleCfg( Id int) *WarOrderCycleWarOrderCycleCfg {
	return gameDb.WarOrderCycleWarOrderCycleCfgs[Id]
}

func RangWarOrderCycleWarOrderCycleCfgs(f func(conf *WarOrderCycleWarOrderCycleCfg)bool){
	for _,v := range gameDb.WarOrderCycleWarOrderCycleCfgs{
		if !f(v){
			return
		}
	}
}

func GetWarOrderCycleTaskWarOrderCycleTaskCfg( Id int) *WarOrderCycleTaskWarOrderCycleTaskCfg {
	return gameDb.WarOrderCycleTaskWarOrderCycleTaskCfgs[Id]
}

func RangWarOrderCycleTaskWarOrderCycleTaskCfgs(f func(conf *WarOrderCycleTaskWarOrderCycleTaskCfg)bool){
	for _,v := range gameDb.WarOrderCycleTaskWarOrderCycleTaskCfgs{
		if !f(v){
			return
		}
	}
}

func GetWarOrderExchangeWarOrderExchangeCfg( Id int) *WarOrderExchangeWarOrderExchangeCfg {
	return gameDb.WarOrderExchangeWarOrderExchangeCfgs[Id]
}

func RangWarOrderExchangeWarOrderExchangeCfgs(f func(conf *WarOrderExchangeWarOrderExchangeCfg)bool){
	for _,v := range gameDb.WarOrderExchangeWarOrderExchangeCfgs{
		if !f(v){
			return
		}
	}
}

func GetWarOrderLevelWarOrderLevelCfg( Id int) *WarOrderLevelWarOrderLevelCfg {
	return gameDb.WarOrderLevelWarOrderLevelCfgs[Id]
}

func RangWarOrderLevelWarOrderLevelCfgs(f func(conf *WarOrderLevelWarOrderLevelCfg)bool){
	for _,v := range gameDb.WarOrderLevelWarOrderLevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetWarOrderWeekTaskWarOrderWeekTaskCfg( Id int) *WarOrderWeekTaskWarOrderWeekTaskCfg {
	return gameDb.WarOrderWeekTaskWarOrderWeekTaskCfgs[Id]
}

func RangWarOrderWeekTaskWarOrderWeekTaskCfgs(f func(conf *WarOrderWeekTaskWarOrderWeekTaskCfg)bool){
	for _,v := range gameDb.WarOrderWeekTaskWarOrderWeekTaskCfgs{
		if !f(v){
			return
		}
	}
}

func GetWashWashCfg( Id int) *WashWashCfg {
	return gameDb.WashWashCfgs[Id]
}

func RangWashWashCfgs(f func(conf *WashWashCfg)bool){
	for _,v := range gameDb.WashWashCfgs{
		if !f(v){
			return
		}
	}
}

func GetWashrandRandCfg( Id int) *WashrandRandCfg {
	return gameDb.WashrandRandCfgs[Id]
}

func RangWashrandRandCfgs(f func(conf *WashrandRandCfg)bool){
	for _,v := range gameDb.WashrandRandCfgs{
		if !f(v){
			return
		}
	}
}

func GetWingNewWingNewCfg( Id int) *WingNewWingNewCfg {
	return gameDb.WingNewWingNewCfgs[Id]
}

func RangWingNewWingNewCfgs(f func(conf *WingNewWingNewCfg)bool){
	for _,v := range gameDb.WingNewWingNewCfgs{
		if !f(v){
			return
		}
	}
}

func GetWingSpecialWingSpecialCfg( Id int) *WingSpecialWingSpecialCfg {
	return gameDb.WingSpecialWingSpecialCfgs[Id]
}

func RangWingSpecialWingSpecialCfgs(f func(conf *WingSpecialWingSpecialCfg)bool){
	for _,v := range gameDb.WingSpecialWingSpecialCfgs{
		if !f(v){
			return
		}
	}
}

func GetWorldBossWorldBossCfg( Id int) *WorldBossWorldBossCfg {
	return gameDb.WorldBossWorldBossCfgs[Id]
}

func RangWorldBossWorldBossCfgs(f func(conf *WorldBossWorldBossCfg)bool){
	for _,v := range gameDb.WorldBossWorldBossCfgs{
		if !f(v){
			return
		}
	}
}

func GetWorldLeaderConfCfg( Id int) *WorldLeaderConfCfg {
	return gameDb.WorldLeaderConfCfgs[Id]
}

func RangWorldLeaderConfCfgs(f func(conf *WorldLeaderConfCfg)bool){
	for _,v := range gameDb.WorldLeaderConfCfgs{
		if !f(v){
			return
		}
	}
}

func GetWorldLeaderRewardWorldLeaderRewardCfg( Id int) *WorldLeaderRewardWorldLeaderRewardCfg {
	return gameDb.WorldLeaderRewardWorldLeaderRewardCfgs[Id]
}

func RangWorldLeaderRewardWorldLeaderRewardCfgs(f func(conf *WorldLeaderRewardWorldLeaderRewardCfg)bool){
	for _,v := range gameDb.WorldLeaderRewardWorldLeaderRewardCfgs{
		if !f(v){
			return
		}
	}
}

func GetWorldLevelWorldLevelCfg( Id int) *WorldLevelWorldLevelCfg {
	return gameDb.WorldLevelWorldLevelCfgs[Id]
}

func RangWorldLevelWorldLevelCfgs(f func(conf *WorldLevelWorldLevelCfg)bool){
	for _,v := range gameDb.WorldLevelWorldLevelCfgs{
		if !f(v){
			return
		}
	}
}

func GetWorldLevelBuffWorldLevelBuffCfg( Id int) *WorldLevelBuffWorldLevelBuffCfg {
	return gameDb.WorldLevelBuffWorldLevelBuffCfgs[Id]
}

func RangWorldLevelBuffWorldLevelBuffCfgs(f func(conf *WorldLevelBuffWorldLevelBuffCfg)bool){
	for _,v := range gameDb.WorldLevelBuffWorldLevelBuffCfgs{
		if !f(v){
			return
		}
	}
}

func GetWorldRankWorldRankCfg( Rank int) *WorldRankWorldRankCfg {
	return gameDb.WorldRankWorldRankCfgs[Rank]
}

func RangWorldRankWorldRankCfgs(f func(conf *WorldRankWorldRankCfg)bool){
	for _,v := range gameDb.WorldRankWorldRankCfgs{
		if !f(v){
			return
		}
	}
}

func GetXiaoyouxiTowerXiaoyouxiTowerCfg( Id int) *XiaoyouxiTowerXiaoyouxiTowerCfg {
	return gameDb.XiaoyouxiTowerXiaoyouxiTowerCfgs[Id]
}

func RangXiaoyouxiTowerXiaoyouxiTowerCfgs(f func(conf *XiaoyouxiTowerXiaoyouxiTowerCfg)bool){
	for _,v := range gameDb.XiaoyouxiTowerXiaoyouxiTowerCfgs{
		if !f(v){
			return
		}
	}
}

func GetXunlongXunlongCfg( Id int) *XunlongXunlongCfg {
	return gameDb.XunlongXunlongCfgs[Id]
}

func RangXunlongXunlongCfgs(f func(conf *XunlongXunlongCfg)bool){
	for _,v := range gameDb.XunlongXunlongCfgs{
		if !f(v){
			return
		}
	}
}

func GetXunlongPrXunlongPrCfg( Time int) *XunlongPrXunlongPrCfg {
	return gameDb.XunlongPrXunlongPrCfgs[Time]
}

func RangXunlongPrXunlongPrCfgs(f func(conf *XunlongPrXunlongPrCfg)bool){
	for _,v := range gameDb.XunlongPrXunlongPrCfgs{
		if !f(v){
			return
		}
	}
}

func GetXunlongRoundsXunlongRoundsCfg( Id int) *XunlongRoundsXunlongRoundsCfg {
	return gameDb.XunlongRoundsXunlongRoundsCfgs[Id]
}

func RangXunlongRoundsXunlongRoundsCfgs(f func(conf *XunlongRoundsXunlongRoundsCfg)bool){
	for _,v := range gameDb.XunlongRoundsXunlongRoundsCfgs{
		if !f(v){
			return
		}
	}
}

func GetZodiacEquipZodiacEquipCfg( Id int) *ZodiacEquipZodiacEquipCfg {
	return gameDb.ZodiacEquipZodiacEquipCfgs[Id]
}

func RangZodiacEquipZodiacEquipCfgs(f func(conf *ZodiacEquipZodiacEquipCfg)bool){
	for _,v := range gameDb.ZodiacEquipZodiacEquipCfgs{
		if !f(v){
			return
		}
	}
}

type PaoDianRewardPaoDianRewardCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Stage IntSlice	`col:"stage" client:"stage"`	 // 地图id
    Interval int	`col:"interval" client:"interval"`	 // 收益间隔：秒
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 泡点奖励
    Times int	`col:"times" client:"times"`	 // 占领者倍数
    TopUserNum int	`col:"topUserNum" client:"topUserNum"`	 // 占领者人数
}

type SpendrebatesSpendrebatesCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type1 int	`col:"type1" client:"type1"`	 // 活动期数
    Show int	`col:"show" client:"show"`	 // 档位展示
    Condition IntMap	`col:"condition" client:"condition"`	 // 累消条件
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
    Time IntSlice	`col:"time" client:"time"`	 // 开启时间
}

type AccumulateAccumulateCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Condition IntMap	`col:"condition" client:"condition"`	 // 累充条件
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type AchievementAchievementCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    NextId int	`col:"nextId"`	 // 下一个任务id
    Condition int	`col:"condition" client:"condition"`	 // 成就条件
    ConditionId int	`col:"conditionId" client:"conditionId"`	 // 成就id
    Level int	`col:"level" client:"level"`	 // 成就参数
    Type int	`col:"Type" client:"Type"`	 // 成就类型
    Drop ItemInfos	`col:"drop" client:"drop" checker:"item"`	 // 奖励
    Point int	`col:"Point"`	 // 成就积分
    Jump int	`col:"jump" client:"jump"`	 // 跳转
}

type AchievementMedalMedalCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Icon string	`col:"icon"`	 // 图标
    Name string	`col:"name"`	 // 勋章名称
    PointOut int	`col:"pointOut" client:"pointOut"`	 // 消耗积分
    Buff IntMap	`col:"buff" client:"buff"`	 // BUFF加成
}

type AncientBossAncientBossCfg struct {
    StageId int	`col:"stageId" client:"stageId"`	 // 关卡信息
    Condition IntMap	`col:"condition" client:"condition"`	 // boss开启条件
    Area int	`col:"area" client:"area"`	 // 远古区域
    JoinDrop ItemInfos	`col:"joinDrop" client:"joinDrop" checker:"item"`	 // 参与奖励
}

type AncientSkillGradeAncientSkillGradeCfg struct {
    Level int	`col:"level" client:"level"`	 // 等阶（技能进阶）
    Condition IntMap	`col:"condition" client:"condition"`	 // 升级条件
    Cost ItemInfos	`col:"cost" client:"cost" checker:"item"`	 // 升级消耗
    ZhanEffect int	`col:"zhanEffect" client:"zhanEffect"`	 // 战士技能（填effectID）
    FaEffect int	`col:"faEffect" client:"faEffect"`	 // 法师技能（填effectID）
    DaoEffect int	`col:"daoEffect" client:"daoEffect"`	 // 道士技能（填effectID）
}

type AncientSkillLevelAncientSkillLevelCfg struct {
    Level int	`col:"level" client:"level"`	 // 等级（技能修炼）
    Condition IntMap	`col:"condition" client:"condition"`	 // 升级条件
    Cost ItemInfos	`col:"cost" client:"cost" checker:"item"`	 // 升级消耗
    Effect1 int	`col:"effect1" client:"effect1"`	 // 战士属性（填effectID）
    Effect2 int	`col:"effect2" client:"effect2"`	 // 法师属性（填effectID）
    Effect3 int	`col:"effect3" client:"effect3"`	 // 道士属性（填effectID）
}

type ArcherMagicCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    StageId int	`col:"stageId" client:"stageId"`	 // 地图id
    Map int	`col:"map" client:"map"`	 // 所属地图
    InsistReward IntSlice2	`col:"insistReward" client:"insistReward"`	 // 坚持奖励
    PassReward ItemInfos	`col:"passReward" client:"passReward" checker:"item"`	 // 通关奖励
    BeginBuff IntSlice	`col:"beginBuff" client:"beginBuff"`	 // 初始元素
}

type ArcherElementMagicElementCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Type int	`col:"type" client:"type"`	 // 元素类型
    Param IntSlice2	`col:"param" client:"param"`	 // 参数
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type AreaAreaCfg struct {
    Id int	`col:"id" client:"id"`	 // ID(不可改）
    Name string	`col:"name" client:"name"`	 // 名字
    Special_id IntSlice	`col:"special_id" client:"special_id"`	 // 等阶id
}

type AreaLevelAreaLevelCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Condition IntMap	`col:"condition" client:"condition"`	 // 升级条件
    Item ItemInfos	`col:"item" client:"item" checker:"item"`	 // 消耗材料(道具ID,数量|道具ID,数量）
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
    Attribute1 IntMap	`col:"attribute1"`	 // 特殊属性加成
    Up int	`col:"up" client:"up"`	 // 升级或突破（前端用）1激活、2升级、3突破
}

type ArenaBuyArenaBuyCfg struct {
    Num int	`col:"num" client:"num"`	 // 购买次数
    Cost ItemInfo	`col:"cost" client:"cost" checker:"item"`	 // 花费
}

type ArenaMatchArenaMatchCfg struct {
    RankMin int	`col:"rankMin" client:"rankMin"`	 // 玩家最低名次
    RankMax int	`col:"rankMax" client:"rankMax"`	 // 玩家最高名次
    RangeHigh int	`col:"rangeHigh" client:"rangeHigh"`	 // 高名次匹配范围
    RangeLow int	`col:"rangeLow" client:"rangeLow"`	 // 低名次匹配范围
}

type ArenaRankArenaRankCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    RankMin int	`col:"rankMin" client:"rankMin"`	 // 最小排名
    RankMax int	`col:"rankMax" client:"rankMax"`	 // 最大排名
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type AspdAspdCfg struct {
    Id int	`col:"id" client:"id"`	 // 档位
    AspdMin int	`col:"aspdMin"`	 // 最小攻速(包含）
    AspdMax int	`col:"aspdMax" client:"aspdMax"`	 // 最大攻速
    Time int	`col:"time" client:"time"`	 // 攻击间隔（毫秒）
    OnceTime int	`col:"onceTime" client:"onceTime"`	 // 攻击一次消耗时间(毫秒)
}

type AtlasAtlasCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 图鉴页签
    Name string	`col:"name" client:"name"`	 // 图鉴名称
    Icon string	`col:"icon" client:"icon"`	 // 图鉴图片
    Quality int	`col:"quality" client:"quality"`	 // 图鉴品质
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 激活消耗
    Desc string	`col:"desc" client:"desc"`	 // 图鉴说明
}

type AtlasGatherAtlasGatherCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 图鉴页签
    Type1 int	`col:"type1"`	 // 图鉴连携页签
    Name string	`col:"name" client:"name"`	 // 图鉴连携名称
    Atlas_id IntSlice	`col:"atlas_id" client:"atlas_id"`	 // 连携图鉴id
}

type AtlasPosAtlasPosCfg struct {
    Pos int	`col:"pos" client:"pos"`	 // 位置
    Condition IntMap	`col:"condition" client:"condition"`	 // 开启条件
}

type AtlasStarAtlasStarCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 图鉴id
    Star int	`col:"star" client:"star"`	 // 图鉴星级
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 升星消耗
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 激活属性
    Attribute1 IntMap	`col:"attribute1" client:"attribute1"`	 // 穿戴属性
}

type AtlasUpgradeAtlasUpgradeCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    AtlasGather int	`col:"atlasGather" client:"atlasGather"`	 // 连携id
    Star int	`col:"star" client:"star"`	 // 星数
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
}

type AttackEnemyAttackEnemyCfg struct {
    Id int	`col:"id" client:"id"`	 // 地图id
    Resource int	`col:"resource" client:"resource"`	 // 资源id
    Card IntSlice2	`col:"card" client:"card"`	 // 卡牌配置（卡牌类型，卡牌数|卡牌类型，卡牌数）
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 通关奖励
    Condition IntMap	`col:"condition" client:"condition"`	 // 进入条件
}

type AttackEnemyCardAttackEnemyCardCfg struct {
    Type int	`col:"type" client:"type"`	 // 卡牌类型
}

type AuctionAuctioinCfg struct {
    Id int	`col:"id" client:"id"`	 // 商品id
    Name string	`col:"name" client:"name"`	 // 商品名称
    Price1 int	`col:"price1" client:"price1"`	 // 商品一口价
    Price2 int	`col:"price2" client:"price2"`	 // 推荐价格
    LowPrice int	`col:"lowPrice" client:"lowPrice"`	 // 最低价格
    HighPrice int	`col:"highPrice" client:"highPrice"`	 // 最高价格
    MaxNum int	`col:"maxNum" client:"maxNum"`	 // 最大数量
    Price3 int	`col:"price3" client:"price3"`	 // 每次加价价格
    Time int	`col:"time"`	 // 竞拍时间（分钟）
}

type AwakenAwakenCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 龙器类型
    Name string	`col:"name" client:"name"`	 // 龙器名字
    Level int	`col:"level" client:"level"`	 // 龙器等级
    AttributeP IntMap	`col:"attributeP" client:"attributeP"`	 // 战士属性
    AttributeM IntMap	`col:"attributeM" client:"attributeM"`	 // 法师属性
    AttributeT IntMap	`col:"attributeT" client:"attributeT"`	 // 道士属性
    Condition IntMap	`col:"condition" client:"condition"`	 // 开启条件
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 消耗
    AttributeS IntMap	`col:"attributeS" client:"attributeS"`	 // 特殊属性
    Icon string	`col:"icon" client:"icon"`	 // 特殊属性图标
    AttributeName string	`col:"attributeName" client:"attributeName"`	 // 特殊属性名称
    Decline int	`col:"decline" client:"decline"`	 // 每周神威下降值
}

type AwakenTitleAwakenTitleCfg struct {
    Rank int	`col:"rank" client:"rank"`	 // 排名
    Sattribute IntMap	`col:"Sattribute" client:"Sattribute"`	 // 特殊属性
}

type BagSpaceAddCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 类型
    Cost int	`col:"cost"`	 // 消耗
    Value int	`col:"value" client:"value"`	 // 值
    BagNumber int	`col:"bagNumber" client:"bagNumber"`	 // 背包格子
}

type BindGroupBindGroupCfg struct {
    BindGroup int	`col:"bindGroup" client:"bindGroup"`	 // 绑定组
    Itemid IntSlice	`col:"itemid" client:"itemid"`	 // 道具id
}

type BlessBlessCfg struct {
    Id int	`col:"id" client:"id"`	 // 武器幸运等级
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 消耗材料
    Success int	`col:"success" client:"success"`	 // 成功权重
    Lose int	`col:"lose" client:"lose"`	 // 失败权重
    Invalid int	`col:"invalid" client:"invalid"`	 // 无效权重
}

type BossFamilyBossFamilyCfg struct {
    StageId int	`col:"stageId" client:"stageId"`	 // 关卡信息
    Condition IntMap	`col:"condition" client:"condition"`	 // 进入条件
    Function int	`col:"function" client:"function"`	 // 打宝地图
}

type BuffBuffCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    BuffType int	`col:"buffType" client:"buffType"`	 // buff类型
    Group int	`col:"group"`	 // 分组
    BuffValue int	`col:"buffValue"`	 // 同类BUFF等级
    BuffRule int	`col:"buffRule"`	 // 同类型叠加规则
    Layer int	`col:"layer"`	 // 最大层数
    Debuff int	`col:"debuff"`	 // 是否负面buff
    Target int	`col:"target"`	 // 目标（敌方，自身）
    Probability int	`col:"probability"`	 // 触发概率(万分比）
    Effect IntMap	`col:"effect" client:"effect"`	 // buff效果
    Time int	`col:"time" client:"time"`	 // 持续时间（毫秒）
    Resurrection int	`col:"resurrection"`	 // 是否复活清空
    Remove int	`col:"remove"`	 // 是否可移除
    FitPurge int	`col:"fitPurge"`	 // 切合体状态是否清除
}

type ChatClearCfg struct {
    Type int	`col:"type" client:"type"`	 // 聊天类型
    Conditon IntMap	`col:"conditon" client:"conditon"`	 // 发言条件
}

type ChuanShiEquipChuanShiEquipCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 部位
    Level int	`col:"level" client:"level"`	 // 阶数
    Condition IntMap	`col:"condition" client:"condition"`	 // 穿戴条件
    Properties IntMap	`col:"properties" client:"properties"`	 // 基础属性
    DeComposeItem ItemInfos	`col:"deComposeItem" client:"deComposeItem" checker:"item"`	 // 分解
}

type ChuanShiEquipTypeChuanShiEquipTypeCfg struct {
    Id int	`col:"id" client:"id"`	 // 装备位
    Name string	`col:"name" client:"name"`	 // 装备位名字
}

type ChuanShiStrengthenChuanShiStrengthenCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Position int	`col:"position" client:"position"`	 // 部位
    Level int	`col:"level" client:"level"`	 // 等级
    Rate int	`col:"rate" client:"rate"`	 // 基础成功率(万分比)
    DemoteRate int	`col:"demoteRate" client:"demoteRate"`	 // 失败降级几率
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 消耗材料
    Condition IntMap	`col:"condition" client:"condition"`	 // 强化条件
    LuckyStone ItemInfos	`col:"luckyStone" client:"luckyStone" checker:"item"`	 // 幸运石
    Relegation ItemInfos	`col:"relegation" client:"relegation" checker:"item"`	 // 保级石
    Definitely ItemInfos	`col:"Definitely" client:"Definitely" checker:"item"`	 // 必升石
    AttributeP IntMap	`col:"attributeP" client:"attributeP"`	 // 战士属性
    AttributeM IntMap	`col:"attributeM" client:"attributeM"`	 // 法师属性
    AttributeT IntMap	`col:"attributeT" client:"attributeT"`	 // 道士属性
}

type ChuanShiStrengthenLinkChuanShiStrengthenLinkCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 类型
    Level int	`col:"level" client:"level"`	 // 加成等级
    Condition int	`col:"condition" client:"condition"`	 // 解锁等级
    Effect int	`col:"effect" client:"effect"`	 // 加成
}

type ChuanShiSuitChuanShiSuitCfg struct {
    Id int	`col:"id" client:"id"`	 // 类型
    Type IntSlice	`col:"type" client:"type"`	 // 装备位
}

type ChuanShiSuitTypeChuanShiSuitTypeCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    SuitType int	`col:"suitType" client:"suitType"`	 // 类型：1刀甲套装、2首饰套装、3防具套装
    Level int	`col:"level" client:"level"`	 // 阶数
    Attribute1 int	`col:"attribute1" client:"attribute1"`	 // 2件套装效果
    Attribute2 int	`col:"attribute2" client:"attribute2"`	 // 4件套装效果
}

type ClearClearCfg struct {
    Type int	`col:"type" client:"type"`	 // 装备部位
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 消耗材料
    AttributeP IntMap	`col:"attributeP" client:"attributeP"`	 // 战士最高属性
    AttributeM IntMap	`col:"attributeM" client:"attributeM"`	 // 法师最高属性
    AttributeT IntMap	`col:"attributeT" client:"attributeT"`	 // 道士最高属性
}

type ClearRateClearRateCfg struct {
    Id int	`col:"id" client:"id"`	 // 档位
    Section1 float64	`col:"section1" client:"section1"`	 // 最低区间系数
    Section2 float64	`col:"section2" client:"section2"`	 // 最高区间系数
    Rate int	`col:"rate" client:"rate"`	 // 万分概率
}

type CollectionCollectionCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Collection IntSlice	`col:"collection" client:"collection"`	 // 采集刷新区域
    Goods int	`col:"goods" client:"goods"`	 // 刷新采集物数量
    Model int	`col:"model" client:"model"`	 // 摆件模型
    Type int	`col:"type" client:"type"`	 // 刷新类型
    Time int	`col:"time" client:"time"`	 // 刷新间隔(秒）
    Effect IntSlice	`col:"effect" client:"effect"`	 // 采集效果
    Success int	`col:"success" client:"success"`	 // 采集成功时间（秒）
}

type CompetitveCompetitveCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Mark int	`col:"mark" client:"mark"`	 // 积分要求
    RewardDay ItemInfos	`col:"rewardDay" client:"rewardDay" checker:"item"`	 // 每日奖励
    MarkLoss int	`col:"markLoss"`	 // 结算积分-输
    MarkWin int	`col:"markWin"`	 // 结算积分-赢
    RewardLoss ItemInfos	`col:"rewardLoss" checker:"item"`	 // 输-奖励
    RewardWin ItemInfos	`col:"rewardWin" checker:"item"`	 // 赢-奖励
}

type CompetitveRewardRankRewardCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Rank IntSlice	`col:"rank" client:"rank"`	 // 排名
    RewardWeek ItemInfos	`col:"rewardWeek" client:"rewardWeek" checker:"item"`	 // 赛季奖励
}

type ComposeChuanShiSubComposeChuanShiSubCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    SubTab int	`col:"subTab" client:"subTab"`	 // 子页签
    Consume1 ItemInfos	`col:"consume1" client:"consume1" checker:"item"`	 // 消耗道具
    Composeid IntSlice2	`col:"composeid" client:"composeid"`	 // 合成道具
    ComCondition IntMap	`col:"ComCondition" client:"ComCondition"`	 // 道具合成条件
}

type ComposeChuanShiTypeComposeChuanShiTypeCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 主页id
    SubTab int	`col:"subTab" client:"subTab"`	 // 子页签
}

type ComposeEquipSubComposeEquipSubCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    SubTab int	`col:"subTab" client:"subTab"`	 // 子页签
    Consume1 ItemInfos	`col:"consume1" client:"consume1" checker:"item"`	 // 消耗道具
    Composeid ItemInfo	`col:"composeid" client:"composeid" checker:"item"`	 // 合成道具（道具id|数量)
    ComCondition IntMap	`col:"ComCondition" client:"ComCondition"`	 // 道具合成条件
    ReplaceItem ItemInfo	`col:"replaceItem" client:"replaceItem" checker:"item"`	 // 需要/代替多少碎片
    CanLucky int	`col:"canLucky" client:"canLucky"`	 // 是否可使用幸运石
    LuckyStone1 ItemInfo	`col:"luckyStone1" client:"luckyStone1" checker:"item"`	 // 小幸运石
    Composeid2 IntSlice2	`col:"composeid2" client:"composeid2"`	 // 使用小幸运石合成道具（道具id，数量，权重|道具id，数量，权重）
    LuckyStone2 ItemInfo	`col:"luckyStone2" client:"luckyStone2" checker:"item"`	 // 大幸运石
    Composeid3 IntSlice2	`col:"composeid3" client:"composeid3"`	 // 使用大幸运石合成道具（道具id，数量，权重|道具id，数量，权重）
}

type ComposeEquipTypeComposeEquipTypeCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 主页id
    SubTab int	`col:"subTab" client:"subTab"`	 // 子页签
}

type ComposeSubComposeSubCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    SubTab int	`col:"subTab" client:"subTab"`	 // 子页签
    Composeid ItemInfo	`col:"composeid" client:"composeid" checker:"item"`	 // 合成道具
    Consume1 ItemInfos	`col:"consume1" client:"consume1" checker:"item"`	 // 消耗道具
    Consume2 ItemInfo	`col:"consume2" client:"consume2" checker:"item"`	 // 消耗货币
    Is_buy bool	`col:"is_buy"`	 // 是否可以直接购买
}

type ComposeTypeComposeTypeCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 主页id
    SubTab int	`col:"subTab" client:"subTab"`	 // 子页签
}

type ConditionConditionCfg struct {
    Id int	`col:"id" client:"id"`	 // 条件id(100-200为打开界面 模块操作)
}

type ContRechargeContRechargeCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Day int	`col:"day" client:"day"`	 // 达成天数
    Type int	`col:"type" client:"type"`	 // 活动期数
    Time IntSlice	`col:"time" client:"time"`	 // 持续时间
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 活动奖励
}

type CrossArenaCrossArenaCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Match int	`col:"match" client:"match"`	 // 赛程
    Round int	`col:"round" client:"round"`	 // 轮次
    RoundName string	`col:"roundName"`	 // 比赛轮次名称
    PlayerNumber int	`col:"playerNumber"`	 // 参与人数
    IntervalTime HmsTime	`col:"intervalTime" client:"intervalTime"`	 // 比赛开启时间
}

type CrossArenaRewardCrossArenaRewardCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type"`	 // 类型：1战败方，2战胜方
    Round int	`col:"round"`	 // 比赛场次
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type CrossArenaRobotCrossArenaRobotCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Name string	`col:"name" client:"name"`	 // 名字
    Icon string	`col:"icon" client:"icon"`	 // 头像
    Combat int	`col:"combat" client:"combat"`	 // 基础战力
}

type CrossArenaTimeCrossArenaTimeCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    OpenDayMin int	`col:"openDayMin" client:"openDayMin"`	 // 开服最小天数
    OpenDayMax int	`col:"openDayMax" client:"openDayMax"`	 // 开服最大天数
    MergeDayMin int	`col:"mergeDayMin" client:"mergeDayMin"`	 // 合服最小天数
    MergeDayMax int	`col:"mergeDayMax" client:"mergeDayMax"`	 // 合服最大天数
    SignUpBegin int	`col:"signUpBegin" client:"signUpBegin"`	 // 报名开始日
    SignUpBeginTime HmsTime	`col:"signUpBeginTime" client:"signUpBeginTime"`	 // 报名开始时间
    SignUpEnd int	`col:"signUpEnd" client:"signUpEnd"`	 // 报名结束日（比赛开始日）
    SignUpEndTime HmsTime	`col:"signUpEndTime" client:"signUpEndTime"`	 // 报名结束时间
    OpenTime HmsTime	`col:"openTime" client:"openTime"`	 // 比赛开始时间
    CloseTime HmsTime	`col:"closeTime" client:"closeTime"`	 // 比赛结束时间
    Condition IntMap	`col:"condition" client:"condition"`	 // 报名条件
}

type CumulationsignCumulationsignCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Days int	`col:"days" client:"days"`	 // 签到天数
    Reward ItemInfo	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type CutTreasureCutTreasureCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Level int	`col:"level" client:"level"`	 // 等级
    Condition IntMap	`col:"condition"`	 // 升级条件
    Item ItemInfos	`col:"item" client:"item" checker:"item"`	 // 升级材料
    SkillZhan int	`col:"skillZhan" client:"skillZhan"`	 // 技能（战士）
    SkillFa int	`col:"skillFa" client:"skillFa"`	 // 技能（法师）
    SkillDao int	`col:"skillDao" client:"skillDao"`	 // 技能（道士）
    Buffzhan int	`col:"buffzhan" client:"buffzhan"`	 // 切割效果
    Bufffa int	`col:"bufffa" client:"bufffa"`	 // 切割效果
    Buffdao int	`col:"buffdao" client:"buffdao"`	 // 切割效果
    CooldownTime int	`col:"cooldownTime" client:"cooldownTime"`	 // 技能冷却时间
}

type DaBaoEquipDaBaoEquipCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 类型
    Class int	`col:"class" client:"class"`	 // 等阶
    Cost ItemInfos	`col:"cost" client:"cost" checker:"item"`	 // 消耗
    Condition IntSlice2	`col:"condition" client:"condition"`	 // 升级条件
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
}

type DaBaoEquipAdditionDaBaoEquipAdditionCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    EquipType int	`col:"equipType" client:"equipType"`	 // 神器类型
    AddType int	`col:"addType" client:"addType"`	 // 加成类型
    AddLevel int	`col:"addLevel" client:"addLevel"`	 // 加成等级
    EquipClass int	`col:"equipClass" client:"equipClass"`	 // 激活条件：打宝神器等级
    Effect IntSlice	`col:"effect" client:"effect"`	 // 效果|生效关卡id（需要支持多个关卡id，以及全部关卡）
}

type DaBaoMysteryDaBaoMysteryCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    StageID int	`col:"stageID" client:"stageID"`	 // 地图ID
    Condition IntSlice2	`col:"condition" client:"condition"`	 // 进入条件
    LimitEnergy int	`col:"limitEnergy" client:"limitEnergy"`	 // 最低体力值
}

type DailyActivityDailyActivityCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    ActivityType int	`col:"activityType" client:"activityType"`	 // 活动类型
    OpenDayMin int	`col:"openDayMin" client:"openDayMin"`	 // 开服最小天数
    OpenDayMax int	`col:"openDayMax" client:"openDayMax"`	 // 开服最大天数
    MergeDayMin int	`col:"mergeDayMin" client:"mergeDayMin"`	 // 合服最小天数
    MergeDayMax int	`col:"mergeDayMax" client:"mergeDayMax"`	 // 合服最大天数
    Week IntSlice	`col:"week" client:"week"`	 // 周几
    OpenTime HmsTime	`col:"openTime" client:"openTime"`	 // 活动开启时间
    CloseTime HmsTime	`col:"closeTime" client:"closeTime"`	 // 活动结束时间
    Stage IntSlice	`col:"stage" client:"stage"`	 // 地图id
    Condition IntMap	`col:"condition" client:"condition"`	 // 进入条件
}

type DailyRewardDailyRewardCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 类型
    Active int	`col:"active" client:"active"`	 // 活跃度
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type DailyTaskDailytaskCfg struct {
    Id int	`col:"id" client:"id"`	 // 日常任务id
    Type int	`col:"type" client:"type"`	 // 任务类型
    Name string	`col:"name" client:"name"`	 // 任务名称
    Icon int	`col:"icon" client:"icon"`	 // 图标
    Active int	`col:"active" client:"active"`	 // 每次活跃度
    Num int	`col:"num" client:"num"`	 // 次数
    Limit int	`col:"limit" client:"limit"`	 // 购买上限
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 购买消耗道具
    Jump int	`col:"jump" client:"jump"`	 // 跳转
    Is_retrieve bool	`col:"is_retrieve" client:"is_retrieve"`	 // 是否可以找回
    Active_consume int	`col:"active_consume" client:"active_consume"`	 // 找回消耗活跃度
    FindReward ItemInfos	`col:"findReward" client:"findReward" checker:"item"`	 // 找回奖励
    Condition IntMap	`col:"condition" client:"condition"`	 // 开启条件
    Reward ItemInfos	`col:"Reward" client:"Reward" checker:"item"`	 // 任务奖励
    Gold_consume int	`col:"gold_consume" client:"gold_consume"`	 // 找回消耗元宝
}

type DailypackDailypackCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 类型
    Condition int	`col:"condition" client:"condition"`	 // 可购买次数
    Type2 int	`col:"type2" client:"type2"`	 // 消耗类型
    Price1 int	`col:"price1" client:"price1"`	 // 购买价格
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
    Display string	`col:"display" client:"display"`	 // 充值名称
}

type DarkPalaceBossDarkPalaceBossCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Level int	`col:"level" client:"level"`	 // BOSS等级
    Condition IntMap	`col:"condition" client:"condition"`	 // boss开启条件
    Floor int	`col:"floor" client:"floor"`	 // 暗殿层
    StageId int	`col:"stageId" client:"stageId"`	 // 关卡信息
    HelpDrop ItemInfos	`col:"helpDrop" client:"helpDrop" checker:"item"`	 // 协助奖励
}

type DayRankingDayRankingCfg struct {
    Day int	`col:"day" client:"day"`	 // 开服天数
    Type int	`col:"type" client:"type"`	 // 排行榜类型
    Name string	`col:"name" client:"name"`	 // 名字
}

type DayRankingGiftDayRankingGiftCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type1 int	`col:"type1" client:"type1"`	 // 排行榜类型
    Type2 int	`col:"type2" client:"type2"`	 // 礼包类型
    Consume int	`col:"consume" client:"consume"`	 // 消耗
    Time int	`col:"time" client:"time"`	 // 限购次数
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
    Display string	`col:"display" client:"display"`	 // 充值名称
}

type DayRankingMarkDayRankingMarkCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Type int	`col:"type" client:"type"`	 // 排行榜类型
    Mark int	`col:"mark" client:"mark"`	 // 分数
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type DayRankingRewardDayRankingRewardCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Ranking IntSlice	`col:"ranking" client:"ranking"`	 // 排名
    Type int	`col:"type" client:"type"`	 // 榜单类型
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
    Least int	`col:"least" client:"least"`	 // 上榜战力、充值金额要求
}

type DictateEquipDictateEquipCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Body int	`col:"body" client:"body"`	 // 部位
    Grade int	`col:"grade" client:"grade"`	 // 阶数
    Condition IntMap	`col:"condition" client:"condition"`	 // 升阶条件
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 升阶材料 id|数量
    Properties IntMap	`col:"properties" client:"properties"`	 // 属性
}

type DictateSuitDictateSuitCfg struct {
    Grade int	`col:"grade" client:"grade"`	 // 阶数
    Effectid1 IntMap	`col:"effectid1" client:"effectid1"`	 // 左套装效果（件数,套装效果|件数,套装效果|件数,套装效果）
    Effectid2 IntMap	`col:"effectid2" client:"effectid2"`	 // 右套装效果（件数,套装效果|件数,套装效果|件数,套装效果）
}

type DragonEquipDragonEquipCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Name string	`col:"name" client:"name"`	 // 名称
}

type DragonEquipLevelDragonEquipLevelCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Item ItemInfos	`col:"item" client:"item" checker:"item"`	 // 升级材料
    Condition IntMap	`col:"condition" client:"condition"`	 // 升级条件
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
    Name string	`col:"name" client:"name"`	 // 名字
}

type DragonarmsDragonarmsCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 装备部位
    Condition IntMap	`col:"condition" client:"condition"`	 // 穿戴条件
    Properties IntMap	`col:"properties" client:"properties"`	 // 固定属性
    Sproperties IntMap	`col:"Sproperties" client:"Sproperties"`	 // 特殊属性
}

type DrawDrawCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type1 int	`col:"type1" client:"type1"`	 // 活动期数
    Type2 int	`col:"type2"`	 // 奖池类型
    Weight int	`col:"weight"`	 // 奖池权重
    Probability IntSlice2	`col:"probability" client:"probability"`	 // 奖池物品概率
    Weight1 int	`col:"weight1"`	 // 奖池权重增加
    Time IntSlice	`col:"time" client:"time"`	 // 开启时间
}

type DrawShopDrawShopCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Reward ItemInfo	`col:"reward" client:"reward" checker:"item"`	 // 奖励
    Num int	`col:"num" client:"num"`	 // 领取积分条件
    Type1 int	`col:"type1" client:"type1"`	 // 活动期数
}

type DropDropCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Dropid int	`col:"dropid" client:"dropid"`	 // 掉落组
    Drop IntSlice2	`col:"drop" client:"drop"`	 // 掉落（物品Id,数量min,数量max,权重|物品Id,数量min,数量max,权重）
    Number IntSlice	`col:"number"`	 // 礼包数量（随机数量1|随机数量2）
    Rate int	`col:"rate"`	 // 概率
}

type DropSpecialDropSpecialCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Type int	`col:"type" client:"type"`	 // 类型
    Num IntSlice	`col:"num" client:"num"`	 // 掉落次数
    DropSpecial IntSlice2	`col:"dropSpecial" client:"dropSpecial"`	 // 掉落库
}

type EffectEffectCfg struct {
    Id int	`col:"id" client:"id"`	 // 效果id
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
    Buffid IntSlice	`col:"buffid" client:"buffid"`	 // buff
    Skillevelid int	`col:"skillevelid" client:"skillevelid"`	 // 技能
}

type ElfGrowElfGrowCfg struct {
    Level int	`col:"level" client:"level"`	 // 等级
    Experience int	`col:"experience" client:"experience"`	 // 经验值
    Buff IntMap	`col:"buff" client:"buff"`	 // BUFF加成
    Condition IntMap	`col:"condition" client:"condition"`	 // 升级条件
}

type ElfRecoverElfRecoverCfg struct {
    Id int	`col:"id" client:"id"`	 // 道具id
    Name string	`col:"name" client:"name"`	 // 道具名称
    Type int	`col:"type" client:"type"`	 // 类型（3装备，其他道具不填）
    Quality int	`col:"quality" client:"quality"`	 // 品质
    Class IntSlice	`col:"class" client:"class"`	 // 阶数
}

type ElfSkillElfGrowCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    SkillId int	`col:"skillId" client:"skillId"`	 // 技能
    SkillLv int	`col:"skillLv" client:"skillLv"`	 // 技能等级
    Condition IntMap	`col:"condition" client:"condition"`	 // 技能激活条件
    Name string	`col:"name" client:"name"`	 // 技能名称
    Way int	`col:"way" client:"way"`	 // 技能激活方式（1.手动激活；2.自动激活）
}

type EquipEquipCfg struct {
    Id int	`col:"id" client:"id"`	 // 装备id
    Type int	`col:"type" client:"type"`	 // 装备部位
    Quality int	`col:"quality" client:"quality"`	 // 品质
    Class int	`col:"class" client:"class"`	 // 阶数
    Star int	`col:"star" client:"star"`	 // 星级
    Condition IntMap	`col:"condition" client:"condition"`	 // 穿戴条件
    Properties IntMap	`col:"properties" client:"properties"`	 // 固定属性
    PropertiesStar IntMap	`col:"propertiesStar" client:"propertiesStar"`	 // 星级属性
    Rand_group IntSlice	`col:"rand_group" client:"rand_group"`	 // 随机属性
}

type EquipsuitEquipsuitCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Effect2_id int	`col:"effect2_id" client:"effect2_id"`	 // 2件套效果
    Effect4_id int	`col:"effect4_id" client:"effect4_id"`	 // 4件套效果
    Effect6_id int	`col:"effect6_id" client:"effect6_id"`	 // 6件套效果
}

type ExpLevelLevelCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Player int	`col:"player" client:"player"`	 // 角色
    Level int	`col:"level" client:"level"`	 // 等级
    ShowText string	`col:"showText" client:"showText"`	 // 等级显示
    Condition IntMap	`col:"condition" client:"condition"`	 // 升级条件
    Item ItemInfos	`col:"item" client:"item" checker:"item"`	 // 消耗道具
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性(战士）
    Attribute2 IntMap	`col:"attribute2" client:"attribute2"`	 // 属性(法师）
    Attribute3 IntMap	`col:"attribute3" client:"attribute3"`	 // 属性(道士）
    Rein int	`col:"rein" client:"rein"`	 // 是否是转生
}

type ExpPillExpPillCfg struct {
    Id int	`col:"id" client:"id"`	 // 经验丹id
    Condition IntSlice2	`col:"condition" client:"condition"`	 // 限制条件
}

type ExpPoolExpPoolCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Limit int	`col:"limit" client:"limit"`	 // 存储上限
}

type ExpStageExpStageCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Stage_id int	`col:"stage_id" client:"stage_id"`	 // 关卡id
    Layer int	`col:"layer" client:"layer"`	 // 层数
    Condition IntMap	`col:"condition" client:"condition"`	 // 开启条件
    Killandappraise IntSlice2	`col:"killandappraise" client:"killandappraise"`	 // 杀敌数量和评价
    Appraiseexp IntSlice	`col:"appraiseexp" client:"appraiseexp"`	 // 评价经验
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 进入消耗
}

type FabaoFabaoCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    ActiveItem int	`col:"activeItem" client:"activeItem"`	 // 激活道具
    Condition PropInfos	`col:"condition" client:"condition"`	 // 开放条件
    UpLvCostItem int	`col:"upLvCostItem" client:"upLvCostItem"`	 // 升级消耗的道具
}

type FabaoSkillFabaoSkillCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Fabao_id int	`col:"fabao_id" client:"fabao_id"`	 // 法宝id
    Fabao_level int	`col:"fabao_level" client:"fabao_level"`	 // 法宝等级
    SkillCostItem ItemInfo	`col:"skillCostItem" client:"skillCostItem" checker:"item"`	 // 技能消耗道具
    AttP IntMap	`col:"attP" client:"attP"`	 // 战士属性
    AttM IntMap	`col:"attM" client:"attM"`	 // 法师属性
    AttT IntMap	`col:"attT" client:"attT"`	 // 道士属性
}

type FabaolevelFabaolevelCfg struct {
    Id int	`col:"id"`	 // id
    Fabaoid int	`col:"fabaoid" client:"fabaoid"`	 // 法宝id
    Level int	`col:"level" client:"level"`	 // 等级
    Exp int	`col:"exp" client:"exp"`	 // 升级经验
    AttributeP IntMap	`col:"attributeP" client:"attributeP"`	 // 战士属性
    AttributeM IntMap	`col:"attributeM" client:"attributeM"`	 // 法师属性
    AttributeT IntMap	`col:"attributeT" client:"attributeT"`	 // 道士属性
    Condition IntMap	`col:"condition" client:"condition"`	 // 升级条件
}

type FashionFashionCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Fashion_id int	`col:"fashion_id" client:"fashion_id"`	 // 时装id
    FashionType int	`col:"fashionType" client:"fashionType"`	 // 类型
    Model IntSlice	`col:"model" client:"model"`	 // 模型
    Level int	`col:"level" client:"level"`	 // 时装等级
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 基础属性
    AttributeS IntMap	`col:"attributeS" client:"attributeS"`	 // 特殊属性
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 消耗材料
}

type FieldBossFieldBossCfg struct {
    StageId int	`col:"stageId" client:"stageId"`	 // 关卡信息
    Level int	`col:"level" client:"level"`	 // BOSS等级
    Condition IntMap	`col:"condition" client:"condition"`	 // boss开启条件
    Area int	`col:"area" client:"area"`	 // 野外区域
    JoinDrop ItemInfos	`col:"joinDrop" client:"joinDrop" checker:"item"`	 // 参与奖励
}

type FieldFightFieldBaseCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type FirstBloodPerFirstBloodperCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Stageid int	`col:"stageid" client:"stageid"`	 // 关卡ID
    RewardFirst ItemInfos	`col:"rewardFirst" client:"rewardFirst" checker:"item"`	 // 首杀奖励
}

type FirstBloodmilFirstBloodmilCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 类别
    Level int	`col:"level" client:"level"`	 // 等级
    Stageid int	`col:"stageid" client:"stageid"`	 // 关卡ID
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
    Num int	`col:"num" client:"num"`	 // 本服击杀数量
}

type FirstBlooduniFirstBlooduniCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Stageid int	`col:"stageid" client:"stageid"`	 // 关卡ID
    RewardFirst ItemInfos	`col:"rewardFirst" client:"rewardFirst" checker:"item"`	 // 首杀奖励
    RewardUni ItemInfos	`col:"rewardUni" client:"rewardUni" checker:"item"`	 // 全服奖励
}

type FirstDropFirstDropCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 类型（function中id）
    Item int	`col:"item" client:"item"`	 // 装备道具
    Count int	`col:"count" client:"count"`	 // 领取数量（-1，无限）
    Reward ItemInfo	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type FirstRechargTypeFirstRechargTypeCfg struct {
    Id int	`col:"id" client:"id"`	 // 档位
    Money int	`col:"money" client:"money"`	 // 金额
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 获得元宝
    Money2 int	`col:"money2" client:"money2"`	 // 优惠券后金额
    Discount int	`col:"discount" client:"discount"`	 // 优惠券显示折扣(百分比
    Display string	`col:"display" client:"display"`	 // 充值名称
}

type FirstRechargeFirstRechargCfg struct {
    Day int	`col:"day" client:"day"`	 // 天数
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type FitFashionFitFashionCfg struct {
    Id int	`col:"id" client:"id"`	 // id
}

type FitFashionLevelFitFashionLevelCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Item ItemInfos	`col:"item" client:"item" checker:"item"`	 // 升级消耗
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
}

type FitHolyEquipFitHolyEquipCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    SuitType int	`col:"suitType" client:"suitType"`	 // 套装类型
    Type int	`col:"type" client:"type"`	 // 部位
    Grade int	`col:"grade" client:"grade"`	 // 阶数
    Condition IntMap	`col:"condition" client:"condition"`	 // 穿戴条件（也是合成条件）
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
    ComposeItem ItemInfos	`col:"composeItem" client:"composeItem" checker:"item"`	 // 合成消耗道具
    DeComposeItem ItemInfos	`col:"deComposeItem" client:"deComposeItem" checker:"item"`	 // 分解
}

type FitHolyEquipSuitFitHolyEquipSuitCfg struct {
    Grade int	`col:"grade" client:"grade"`	 // 阶数
    Effect int	`col:"effect" client:"effect"`	 // 套装效果
}

type FitLevelFitLevelCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Level1 int	`col:"level1" client:"level1"`	 // 境界
    Level2 int	`col:"level2" client:"level2"`	 // 一到十重天
    Level3 int	`col:"level3" client:"level3"`	 // 级数
    Name string	`col:"name" client:"name"`	 // 等级名字
    Condition IntMap	`col:"condition" client:"condition"`	 // 升级条件
    Item ItemInfos	`col:"item" client:"item" checker:"item"`	 // 消耗材料
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
    Duration int	`col:"Duration" client:"Duration"`	 // 合体持续时间
    CooldownTime int	`col:"cooldownTime" client:"cooldownTime"`	 // 合体冷却时间
}

type FitSkillFitSkillCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 类型
    Icon int	`col:"icon" client:"icon"`	 // 技能图标
    SkillName string	`col:"skillName" client:"skillName"`	 // 技能名字
    SkillDescribe string	`col:"skillDescribe" client:"skillDescribe"`	 // 技能概述
    Item ItemInfos	`col:"item" client:"item" checker:"item"`	 // 激活道具
}

type FitSkillLevelFitSkillLevelCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Item ItemInfos	`col:"item" client:"item" checker:"item"`	 // 升级材料
    Condition IntMap	`col:"condition" client:"condition"`	 // 升级限制：合体等级
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
}

type FitSkillSlotFitSkillSlotCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Condition IntMap	`col:"condition" client:"condition"`	 // 合体多少级时解锁
}

type FitSkillStarFitSkillStarCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Item ItemInfos	`col:"item" client:"item" checker:"item"`	 // 升星材料
    Condition IntMap	`col:"condition" client:"condition"`	 // 升级限制：合体等级
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
    Skill int	`col:"skill" client:"skill"`	 // 技能
}

type FunctionFunctionCfg struct {
    Id int	`col:"id" client:"id"`	 // id(不可修改)
    Activation int	`col:"activation" client:"activation"`	 // 默认是否开启（0：达到条件才可开启，1：默认开启）
    Desc string	`col:"desc" client:"desc"`	 // 描述
    Condition IntMap	`col:"condition" client:"condition"`	 // 开启条件
    ConditionTime IntSlice	`col:"conditionTime" client:"conditionTime"`	 // 开启时间（时间戳）|结束时间（时间戳）
}

type GameDotGameDotCfg struct {
    Id int	`col:"id"`	 // id
    ClinetType string	`col:"clinetType"`	 // 客户端读取类型
}

type GameTextErrorTextCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    ConstName string	`col:"constName"`	 // 代码常量
    Chinese string	`col:"chinese" client:"chinese"`	 // 中文错误消息
    Language string	`col:"language"`	 // 错误消息
}

type GameTextCodeTextCfg struct {
    Id int	`col:"id"`	 // id
    ConstName string	`col:"constName"`	 // 代码常量
    Chinese string	`col:"chinese"`	 // 中文
    Language string	`col:"language"`	 // 其他语言
}

type GamewordGameCfg struct {
    Id int	`col:"id"`	 // id
    ClinetType string	`col:"clinetType"`	 // 客户端读取类型
}

type GiftGiftCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 类型
    Reward IntSlice2	`col:"reward" client:"reward"`	 // 物品
    Choose int	`col:"choose" client:"choose"`	 // 自选数量
}

type GiftCodeGiftCodeCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Code string	`col:"code" client:"code"`	 // 礼包码
    Gift ItemInfos	`col:"gift" client:"gift" checker:"item"`	 // 礼包
}

type GodBloodGodBloodCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 升重消耗
    Level int	`col:"level" client:"level"`	 // 阶数
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性加成
}

type GodEquipConfCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 激活消耗
}

type GodEquipLevelConfCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 升级消耗
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性加成
    Condition IntMap	`col:"condition" client:"condition"`	 // 突破条件
}

type GodEquipSuitConfCfg struct {
    Level int	`col:"level" client:"level"`	 // 阶数
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性加成
}

type GrowFundGrowFundCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Condition IntMap	`col:"condition" client:"condition"`	 // 领取条件
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
    Jump int	`col:"jump" client:"jump"`	 // 跳转
}

type GuardRankGuardRankCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Rank IntSlice	`col:"rank" client:"rank"`	 // 排名
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type GuardRoundsGuardRoundsCfg struct {
    Rounds int	`col:"rounds" client:"rounds"`	 // 波数
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type GuildGuildCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Position string	`col:"position" client:"position"`	 // 职位
    Count int	`col:"count" client:"count"`	 // 人数限制
    ApplyMassage int	`col:"applyMassage" client:"applyMassage"`	 // 处理申请消息权限
    ChangePosition int	`col:"changePosition" client:"changePosition"`	 // 任命职位权限
    OustGuild int	`col:"oustGuild" client:"oustGuild"`	 // 踢出公会权限
    Activity IntSlice	`col:"activity" client:"activity"`	 // 开启活动权限
}

type GuildActivityGuildActivityCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    ActivityType int	`col:"activityType" client:"activityType"`	 // 活动类型
    OpenDayMin int	`col:"openDayMin" client:"openDayMin"`	 // 开服最小天数
    OpenDayMax int	`col:"openDayMax" client:"openDayMax"`	 // 开服最大天数
    MergeDayMin int	`col:"mergeDayMin" client:"mergeDayMin"`	 // 合服最小天数
    MergeDayMax int	`col:"mergeDayMax" client:"mergeDayMax"`	 // 合服最大天数
    Week IntSlice	`col:"week" client:"week"`	 // 周几
    OpenTime HmsTime	`col:"openTime" client:"openTime"`	 // 活动开启时间
    CloseTime HmsTime	`col:"closeTime" client:"closeTime"`	 // 活动结束时间
    Stage IntSlice	`col:"stage" client:"stage"`	 // 地图id
    Condition IntMap	`col:"condition" client:"condition"`	 // 进入条件
}

type GuildAuctionGuildAuctionCfg struct {
    Id int	`col:"id" client:"id"`	 // 商品id
    Name string	`col:"name" client:"name"`	 // 商品名称
    Price1 int	`col:"price1" client:"price1"`	 // 商品一口价
    Price2 int	`col:"price2" client:"price2"`	 // 商品初始竞价
    Price3 int	`col:"price3" client:"price3"`	 // 每次加价价格
    Time int	`col:"time" client:"time"`	 // 竞拍时间（分钟）
}

type GuildAutoCreateGuildAutoCreateCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    RobotNum int	`col:"robotNum" client:"robotNum"`	 // 机器人数量
}

type GuildBonfireGuildBonfireCfg struct {
    Id int	`col:"id" client:"id"`	 // 提升类型（1木材，2元宝）
    Item ItemInfo	`col:"item" client:"item" checker:"item"`	 // 消耗道具（道具ID#数量）
    Promote int	`col:"promote" client:"promote"`	 // 每次捐献提升经验万分比
    Times int	`col:"times" client:"times"`	 // 每次活动最多捐献次数
}

type GuildLevelGuildLevelCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Value int	`col:"value" client:"value"`	 // 所需贡献值
    NumberLimit int	`col:"numberLimit" client:"numberLimit"`	 // 人数上限
}

type GuildNameGuildNameCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Name string	`col:"name" client:"name"`	 // 名字
}

type GuildRobotGuildRobotCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Name string	`col:"name" client:"name"`	 // 名字
    Level int	`col:"level" client:"level"`	 // 等级
}

type HellBossHellBossCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Condition IntMap	`col:"condition" client:"condition"`	 // boss开启条件
    Floor int	`col:"floor" client:"floor"`	 // 炼狱层
    StageId int	`col:"stageId" client:"stageId"`	 // 关卡信息
    HelpDrop ItemInfos	`col:"helpDrop" client:"helpDrop" checker:"item"`	 // 协助奖励
}

type HellBossFloorHellBossFloorCfg struct {
    Map int	`col:"map" client:"map"`	 // 炼狱层
    Stageid int	`col:"stageid" client:"stageid"`	 // 关卡信息
    Condition IntMap	`col:"condition" client:"condition"`	 // 炼狱层条件
}

type HolyArmsHolyArmsCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Condition IntMap	`col:"condition" client:"condition"`	 // 开放条件
    UpLvCostItem IntSlice	`col:"upLvCostItem" client:"upLvCostItem"`	 // 消耗道具
}

type HolyBeastHolyBeastCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 类型id
    Star int	`col:"star" client:"star"`	 // 圣兽星数
    Properties IntMap	`col:"properties" client:"properties"`	 // 普通属性
    SelectProperties IntSlice2	`col:"selectProperties" client:"selectProperties"`	 // 属性选择
    Effect IntSlice	`col:"effect" client:"effect"`	 // 技能
    Active ItemInfos	`col:"active" client:"active" checker:"item"`	 // 激活消耗
    ActiveNum int	`col:"activeNum" client:"activeNum"`	 // 消耗圣灵点数量
}

type HolySkillHolySkillCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Holy_id int	`col:"holy_id" client:"holy_id"`	 // 神兵id
    Holy_level int	`col:"holy_level" client:"holy_level"`	 // 神兵等级
    Skill_level int	`col:"skill_level" client:"skill_level"`	 // 技能等级
    SkillCostItem ItemInfo	`col:"skillCostItem" client:"skillCostItem" checker:"item"`	 // 技能消耗道具
    Effect int	`col:"effect" client:"effect"`	 // 效果
}

type HolylevelHolylevelCfg struct {
    Id int	`col:"id"`	 // id
    Holy_id int	`col:"holy_id" client:"holy_id"`	 // 神兵id
    Level int	`col:"level" client:"level"`	 // 等级
    Exp int	`col:"exp" client:"exp"`	 // 升级经验
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
    Condition IntMap	`col:"condition" client:"condition"`	 // 升级条件
}

type HookMapHookMapCfg struct {
    Stage_id int	`col:"stage_id" client:"stage_id"`	 // 关卡id
    StageId2 int	`col:"stageId2" client:"stageId2"`	 // 闯关BOSS
    StageId3 int	`col:"stageId3" client:"stageId3"`	 // 开启下一地图
    Name int	`col:"name" client:"name"`	 // 挂机经验（分钟）
    Drop int	`col:"drop" client:"drop"`	 // 关卡奖励
    Num int	`col:"num" client:"num"`	 // 进入BOSS关卡所需击杀小怪波数
}

type InsideArtInsideArtCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Grade int	`col:"grade" client:"grade"`	 // 阶数
    Order int	`col:"order" client:"order"`	 // 重数
    Star int	`col:"star" client:"star"`	 // 星数
    Attribute1 IntMap	`col:"attribute1" client:"attribute1"`	 // 穴位1属性
    Attribute2 IntMap	`col:"attribute2" client:"attribute2"`	 // 穴位2属性
    Attribute3 IntMap	`col:"attribute3" client:"attribute3"`	 // 穴位3属性
    Attribute4 IntMap	`col:"attribute4" client:"attribute4"`	 // 穴位4属性
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 消耗
    Condition IntMap	`col:"condition" client:"condition"`	 // 突破条件
}

type InsideGradeInsideGradeCfg struct {
    Grade int	`col:"grade" client:"grade"`	 // 阶数
    Success int	`col:"success" client:"success"`	 // 成功率
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 每次消耗
    Condition IntMap	`col:"condition" client:"condition"`	 // 突破条件
}

type InsideSkillInsideSkillCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Skill_id int	`col:"skill_id" client:"skill_id"`	 // 技能
    Type int	`col:"type" client:"type"`	 // 技能类型
    Name string	`col:"name" client:"name"`	 // 技能名称
    Icon string	`col:"icon" client:"icon"`	 // 技能图标
    Describe string	`col:"describe" client:"describe"`	 // 技能描述
    Condition IntMap	`col:"condition" client:"condition"`	 // 激活条件
    Level int	`col:"level" client:"level"`	 // 等级
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 消耗
    Effect int	`col:"effect" client:"effect"`	 // 效果
    InsideUp int	`col:"insideUp"`	 // 增加内功属性万分比
}

type InsideStarInsideStarCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Grade int	`col:"grade" client:"grade"`	 // 阶数
    Order int	`col:"order" client:"order"`	 // 重数
    Star int	`col:"star" client:"star"`	 // 星数
    Weight int	`col:"weight" client:"weight"`	 // 权重
}

type ItemBaseCfg struct {
    Id int	`col:"id" client:"id"`	 // id（0-100占位留给顶级道具）
    Name string	`col:"name" client:"name"`
    Type int	`col:"type" client:"type"`	 // 类型（1货币道具、2药水、3装备、4礼包、5生肖装备、6普通道具、7帝器、8龙器、9特戒，10 技能书、11公会贡献值、12、自选宝箱、13、充值优惠券20持续回血药、21持续回蓝药、23瞬回药,24:随机石，25:回城卷、26随机宝箱、27充值代币、28红包、29圣装、30传世装备（神装）、31称号、32秘籍、33神兵血炼、34打宝秘境体力卷轴、35特权体验卡、36经验丹、37特权卡激活、101转生材料、102技能材料、103天赋材料、104主宰装备材料、105强化材料、106宝石、107内功材料、108灵丹材料、109法宝材料、110龙器材料、111神兵材料、112羽翼材料、113法阵材料、114领域材料、115绝学材料、116战宠材料、117官职材料、118合体材料、119图鉴材料、120时装材料、121圣兽材料、122小精灵材料、123神刀技材料、124洗练材料、125公会材料、126装备合成材料、127合体圣装材料、128特戒材料、129野战、130竞技场、131寻龙探宝、132神机宝库、133挖矿、134首领卷轴、135战令材料、136副本材料、137日常任务活跃度、138传世装备材料、139远古神技材料、140远古宝物材料、141传世强化材料、142打宝神器）
    Quality int	`col:"quality" client:"quality"`	 // 品质
    Quality1 int	`col:"quality1" client:"quality1"`	 // 特殊品质框
    Bind int	`col:"bind" client:"bind"`	 // 是否绑定
    BindGruonp int	`col:"bindGruonp" client:"bindGruonp"`	 // 绑定组
    CountLimit int	`col:"countLimit" client:"countLimit"`	 // 堆叠上限
    EffectVal int	`col:"effectVal" client:"effectVal"`	 // 获取效果值
    Destroy_items ItemInfos	`col:"destroy_items" client:"destroy_items" checker:"item"`	 // 销毁物品，不配是不可销毁
    Recover_items ItemInfos	`col:"recover_items" client:"recover_items" checker:"item"`	 // 回收物品
    ElfExperience int	`col:"elfExperience" client:"elfExperience"`	 // 精灵回收经验值
    ElfRecover ItemInfos	`col:"elfRecover" client:"elfRecover" checker:"item"`	 // 精灵回收物品
    ItemValues int	`col:"itemValues"`	 // 道具价值（暂只用于整理排序）
    Window int	`col:"window" client:"window"`	 // 是否弹窗（1弹窗，0不弹窗）
}

type JewelJewelCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 宝石种类
    Level int	`col:"level" client:"level"`	 // 宝石等级
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 宝石属性
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 消耗
}

type JewelBodyJewelBodyCfg struct {
    Body int	`col:"body" client:"body"`	 // 装备部位
    Type int	`col:"type" client:"type"`	 // 宝石种类
    Condition1 IntMap	`col:"condition1" client:"condition1"`	 // 条件1
    Condition2 IntMap	`col:"condition2" client:"condition2"`	 // 条件2
    Condition3 IntMap	`col:"condition3" client:"condition3"`	 // 条件3
}

type JewelSuitJewelSuitCfg struct {
    Sum int	`col:"sum" client:"sum"`	 // 宝石等级总和
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
}

type JuexueLevelConfCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 消耗
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性加成
    Buff int	`col:"buff" client:"buff"`	 // 技能
}

type KingarmsKingarmsCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 装备部位
    Condition IntMap	`col:"condition" client:"condition"`	 // 穿戴条件
    Star int	`col:"star" client:"star"`	 // 显示星级
    Properties IntMap	`col:"properties" client:"properties"`	 // 固定属性
    Sproperties IntMap	`col:"Sproperties" client:"Sproperties"`	 // 特殊属性
}

type KuafushabakeRewardserverKuafushabakeRewardserverCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Rank IntSlice	`col:"rank" client:"rank"`	 // 排名
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type KuafushabakeRewarduniKuafushabakeRewarduniCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Rank IntSlice	`col:"rank" client:"rank"`	 // 排名
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type LabelLabelCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    AllCondition IntSlice	`col:"allCondition" client:"allCondition"`	 // 晋升条件
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
    DailyReward ItemInfos	`col:"dailyReward" client:"dailyReward" checker:"item"`	 // 每日俸禄
    NormalAttr IntMap	`col:"normalAttr" client:"normalAttr"`	 // 基础属性
    IsChoice int	`col:"isChoice" client:"isChoice"`	 // 是否需要选择晋升之路
    WenAttr IntMap	`col:"wenAttr" client:"wenAttr"`	 // 文官属性
    WuAttr IntMap	`col:"wuAttr" client:"wuAttr"`	 // 武将属性
    WenEffect int	`col:"wenEffect" client:"wenEffect"`	 // 文官技能
    WuEffect int	`col:"wuEffect" client:"wuEffect"`	 // 武将技能
}

type LabelTaskLabelTaskCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Condition int	`col:"condition" client:"condition"`	 // 条件类型
    Value IntSlice	`col:"value" client:"value"`	 // 参数
}

type LimitedGiftLimitedGiftCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Type int	`col:"type" client:"type"`	 // 礼包模块
    Lv int	`col:"lv" client:"lv"`	 // 模块等级
    Condition IntMap	`col:"condition" client:"condition"`	 // 激活条件
    Type1 int	`col:"Type1" client:"Type1"`	 // 档位1类型
    Reward1 ItemInfos	`col:"reward1" client:"reward1" checker:"item"`	 // 档位1道具
    Consume1 int	`col:"consume1" client:"consume1"`	 // 档位1价格
    Type2 int	`col:"Type2" client:"Type2"`	 // 档位2类型
    Reward2 ItemInfos	`col:"reward2" client:"reward2" checker:"item"`	 // 档位2道具
    Consume2 int	`col:"consume2" client:"consume2"`	 // 档位2价格
    Type3 int	`col:"Type3" client:"Type3"`	 // 档位3类型
    Reward3 ItemInfos	`col:"reward3" client:"reward3" checker:"item"`	 // 档位3道具
    Consume3 int	`col:"consume3" client:"consume3"`	 // 档位3价格
    Type4 int	`col:"Type4" client:"Type4"`	 // 档位4类型
    Reward4 ItemInfos	`col:"reward4" client:"reward4" checker:"item"`	 // 档位4道具
    Consume4 int	`col:"consume4" client:"consume4"`	 // 档位4价格
    Type5 int	`col:"Type5" client:"Type5"`	 // 档位5类型
    Reward5 ItemInfos	`col:"reward5" client:"reward5" checker:"item"`	 // 档位5道具
    Consume5 int	`col:"consume5" client:"consume5"`	 // 档位5价格
    Time int	`col:"time" client:"time"`	 // 有效时间（秒）
    Display1 string	`col:"display1" client:"display1"`	 // 充值名称1
    Display2 string	`col:"display2" client:"display2"`	 // 充值名称2
    Display3 string	`col:"display3" client:"display3"`	 // 充值名称3
    Display4 string	`col:"display4" client:"display4"`	 // 充值名称4
    Display5 string	`col:"display5" client:"display5"`	 // 充值名称5
}

type LuckyLuckyCfg struct {
    Id int	`col:"id" client:"id"`	 // 幸运值
    AddAtk float64	`col:"addAtk" client:"addAtk"`	 // 提高攻击下限比例
}

type MagicCircleMagicCircleCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Name string	`col:"name" client:"name"`	 // 名字
}

type MagicCircleLevelMagicCircleLevelCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Type int	`col:"type" client:"type"`	 // 法阵类型
    Rank int	`col:"rank" client:"rank"`	 // 阶数
    Level int	`col:"level" client:"level"`	 // 星级
    Condition IntMap	`col:"condition" client:"condition"`	 // 升级条件
    Item ItemInfos	`col:"item" client:"item" checker:"item"`	 // 升级消耗
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性加成
}

type MagicTowerMagicTowerCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    StageId1 int	`col:"stageId1" client:"stageId1"`	 // 关卡ID
    Rewards ItemInfos	`col:"rewards" client:"rewards" checker:"item"`	 // 奖励
    MarkConsume int	`col:"markConsume" client:"markConsume"`	 // 传送消耗积分
    StageId2 int	`col:"stageId2" client:"stageId2"`	 // 安全复活地图ID
    MarkGet int	`col:"markGet" client:"markGet"`	 // 击杀怪物得分
    RewardsSpecial ItemInfos	`col:"rewardsSpecial" client:"rewardsSpecial" checker:"item"`	 // 地图存活奖励
    RewardsSpecialtime int	`col:"rewardsSpecialtime" client:"rewardsSpecialtime"`	 // 存活时间获得奖励(秒）
}

type MagicTowerRewardMagicTowerRewardCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Rank IntSlice	`col:"rank" client:"rank"`	 // 积分排名
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 排名奖励
}

type MailMailCfg struct {
    Id int	`col:"id"`	 // 序号
    FromName string	`col:"fromName"`	 // 发件人
    Title string	`col:"title"`	 // 标题
    Content string	`col:"content"`	 // 内容
    ExpireDays int	`col:"expireDays"`	 // 过期天数
    DelHours int	`col:"delHours"`	 // 操作后删除小时
    Explain string	`col:"explain"`	 // 说明
}

type MainPrMainPrCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Type int	`col:"type" client:"type"`	 // 装备种类（3装备、5生肖装备、7帝器、8龙器、9特戒）
    Body int	`col:"body" client:"body"`	 // 装备部位
    MainPrP int	`col:"mainPrP" client:"mainPrP"`	 // 战士主属性
    MainPrM int	`col:"mainPrM" client:"mainPrM"`	 // 法师主属性
    MainPrT int	`col:"mainPrT" client:"mainPrT"`	 // 道士主属性
}

type MapMapCfg struct {
    Id int	`col:"id" client:"id"`	 // 地图id
    Resource int	`col:"resource" client:"resource"`	 // 资源途径
    Born IntSlice	`col:"born" client:"born"`	 // 出生区域( 对应地图配置数据的id )
    Revive IntSlice	`col:"revive" client:"revive"`	 // 复活区域( 对应地图配置数据的id )
    SafeRect IntSlice	`col:"safeRect" client:"safeRect"`	 // 安全区
    Special IntSlice	`col:"special" client:"special"`	 // 特殊区域（不同关卡类型做不同处理）
}

type MaptypeGameCfg struct {
    Id int	`col:"id" client:"id"`	 // 类型ID
    CanUseDDQG IntSlice	`col:"canUseDDQG" client:"canUseDDQG"`	 // 是否可以使用刀刀切割[第一字段表是否可以使用刀刀切割，第二字段标识是否可以自动释放刀刀切割]
    CanUseHt int	`col:"canUseHt" client:"canUseHt"`	 // 是否可以使用合体
    ResetCutCD int	`col:"resetCutCD" client:"resetCutCD"`	 // 刷新神刀技CD
    CountDown int	`col:"countDown" client:"countDown"`	 // 开场倒计时（秒）
    Tower int	`col:"tower"`	 // 灯塔
    BagFull int	`col:"bagFull" client:"bagFull"`	 // 背包满了是否拣掉落（1，不捡   0或者不填，捡）
}

type MaterialCostMaterialCostCfg struct {
    Number int	`col:"number" client:"number"`	 // 扫荡次数
    Cost ItemInfo	`col:"cost" client:"cost" checker:"item"`	 // 扫荡消耗
}

type MaterialHomeMaterialHomeCfg struct {
    Type int	`col:"type" client:"type"`	 // 副本类型
    Name string	`col:"name" client:"name"`	 // 副本名称
    Challenge int	`col:"challenge" client:"challenge"`	 // 每日挑战次数
    VIPID int	`col:"VIPID" client:"VIPID"`	 // 特权ID
}

type MaterialStageMaterialStageCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 副本类型
    Level int	`col:"level" client:"level"`	 // 副本等级
    Stageid int	`col:"stageid" client:"stageid"`	 // 副本关卡id
    Conditon IntMap	`col:"conditon" client:"conditon"`	 // 开启条件
    Reward ItemInfo	`col:"reward" client:"reward" checker:"item"`	 // 通关奖励
}

type MijiMijiCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Type int	`col:"type" client:"type"`	 // 秘籍所属
}

type MijiLevelMijiLevelCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 消耗
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性加成
    SkillLevel int	`col:"skillLevel" client:"skillLevel"`	 // 技能
}

type MijiTypeMijiTypeCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
}

type MiningMiningCfg struct {
    Id int	`col:"id" client:"id"`	 // 矿工ID
    Lv int	`col:"lv" client:"lv"`	 // 矿工等级
    Reward IntMap	`col:"reward" client:"reward"`	 // 挖矿奖励
    Time int	`col:"time" client:"time"`	 // 挖矿时间
    Lose IntMap	`col:"lose" client:"lose"`	 // 被抢夺丢失
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 单次升级消耗
    ConsumeMax ItemInfos	`col:"consumeMax" client:"consumeMax" checker:"item"`	 // 一键升级消耗
    Lucky int	`col:"lucky" client:"lucky"`	 // 幸运值上限
    Probability int	`col:"probability"`	 // 升级成功概率
}

type MonsterMonsterCfg struct {
    Monsterid int	`col:"monsterid" client:"monsterid"`	 // 怪物id
    Name string	`col:"name" client:"name"`	 // 名称
    Type int	`col:"type" client:"type"`	 // 类型
    Level int	`col:"level" client:"level"`	 // 等级
    Job int	`col:"job" client:"job"`	 // 职业(1战士2，法师，3道士）
    Attr IntMap	`col:"attr" client:"attr"`	 // 属性
    Skills IntSlice	`col:"skills" client:"skills"`	 // 技能
    DropType int	`col:"dropType" client:"dropType"`	 // 掉落归属（1.击杀者2.归属者）
    DropId int	`col:"dropId"`	 // 掉落
    DropSpecial IntSlice	`col:"dropSpecial"`	 // 特殊掉落
    FirstDropId int	`col:"firstDropId"`	 // 首爆掉落
    ImmuneBuff IntSlice	`col:"immuneBuff" client:"immuneBuff"`	 // 免疫buff（buff类型）
    Aitype int	`col:"aitype" client:"aitype"`	 // AI类型(1主动怪 2被动 )
    Toattackarea int	`col:"toattackarea" client:"toattackarea"`	 // 警戒范围(格子数)
    ChaseDis int	`col:"chaseDis" client:"chaseDis"`	 // 追击距离(格子数)
    ReliveDelay int	`col:"reliveDelay" client:"reliveDelay"`	 // 复活时间 毫秒 0 不复活
    Refresh StringSlice	`col:"refresh"`	 // 定点刷新时间
    ReliveAddrType int	`col:"reliveAddrType" client:"reliveAddrType"`	 // 复活地类型（1原地复活，2出生点复活）
    Move int	`col:"move" client:"move"`	 // 是否可移动（1：可移动，0：不可移动）
    Protect IntSlice	`col:"protect" client:"protect"`	 // 保护条出现剩余血量(百分比)|持续时间(秒）
    Maxatt int	`col:"maxatt"`	 // 最大伤害系数（万分比；-1不限制最大伤害）
    DaBaoEnergy int	`col:"daBaoEnergy" client:"daBaoEnergy"`	 // 击杀扣除体力
    Buff IntSlice	`col:"buff" client:"buff"`	 // buff
}

type MonsterdropDropCfg struct {
    Dropid int	`col:"dropid" client:"dropid"`	 // 掉落组
    CommonDrop IntSlice2	`col:"commonDrop" client:"commonDrop"`	 // 普通掉落库（物品Id,数量min,数量max,权重|物品Id,数量min,数量max,权重）
    BestDrop IntSlice2	`col:"bestDrop" client:"bestDrop"`	 // 极品掉落（物品Id,数量min,数量max,权重|物品Id,数量min,数量max,权重）
}

type MonstergroupMonstergroupCfg struct {
    Groupid int	`col:"groupid" client:"groupid"`	 // 怪物组id
    Monsterid IntSlice2	`col:"monsterid" client:"monsterid"`	 // 怪物(怪物Id,数量,权重|怪物Id,数量,权重)
}

type MonthCardMonthCardCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 档次（1白银，2黄金）
    Day int	`col:"day" client:"day"`	 // 时效(天数)
    Cost int	`col:"cost" client:"cost"`	 // 激活消耗(人民币)
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 激活奖励
    DailyReward ItemInfos	`col:"dailyReward" client:"dailyReward" checker:"item"`	 // 每日礼包
    Privilege IntMap	`col:"privilege" client:"privilege"`	 // 特权
    IsShow int	`col:"isShow" client:"isShow"`	 // 状态(1启用,2不启用)
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
    Display string	`col:"display" client:"display"`	 // 充值显示item
}

type MonthCardPrivilegeMonthCardPrivilegeCfg struct {
    Id int	`col:"id" client:"id"`	 // 类型
}

type NpcMonsterCfg struct {
    Id int	`col:"id" client:"id"`	 // npcid
}

type OfficialOfficialCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Condition IntMap	`col:"condition" client:"condition"`	 // 升级条件
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 消耗材料
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
}

type OpenGiftOpenGiftCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Item ItemInfos	`col:"item" client:"item" checker:"item"`	 // 礼包道具
    Price int	`col:"price" client:"price"`	 // 礼包价格（元）
    Time int	`col:"time" client:"time"`	 // 限购次数
    Display string	`col:"display" client:"display"`	 // 充值显示item
}

type PanaceaPanaceaCfg struct {
    Id int	`col:"id" client:"id"`	 // 灵丹id
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
    Condition IntSlice2	`col:"condition" client:"condition"`	 // 使用上限解锁条件
    Limit IntSlice	`col:"limit" client:"limit"`	 // 使用数量上限
}

type PersonalBossPersonalBossCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Level int	`col:"level" client:"level"`	 // BOSS等级
    StageId int	`col:"stageId" client:"stageId"`	 // 关卡id
    MaxNum int	`col:"maxNum" client:"maxNum"`	 // 每日次数
    Condition IntMap	`col:"condition" client:"condition"`	 // 开启条件
}

type PetsConfCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Name string	`col:"name" client:"name"`	 // 战宠名字
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 激活消耗
    Skill IntSlice	`col:"skill" client:"skill"`	 // 战宠技能
}

type PetsAddPetsAddCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    PetsId int	`col:"petsId" client:"petsId"`	 // 战宠ID
    Condition IntSlice2	`col:"condition" client:"condition"`	 // 附体条件
    Level int	`col:"level" client:"level"`	 // 附体等级
    AttributePets IntMap	`col:"attributePets" client:"attributePets"`	 // 附体加成
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 升级消耗
}

type PetsAddSkillPetsAddSkillCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Type int	`col:"type" client:"type"`	 // 技能类型
    Effect int	`col:"effect" client:"effect"`	 // 激活技能
    Condition IntSlice2	`col:"condition" client:"condition"`	 // 激活条件
}

type PetsBreakConfCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 突破材料
    AttributePets IntMap	`col:"attributePets" client:"attributePets"`	 // 宠物属性加成
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 角色属性加成
    Effect IntSlice	`col:"effect" client:"effect"`	 // 被动技能
}

type PetsGradeConfCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 进阶消耗
    NeedLv int	`col:"needLv" client:"needLv"`	 // 进阶条件(等级)
    AttributePets IntMap	`col:"attributePets" client:"attributePets"`	 // 宠物属性加成
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 角色属性加成
    Skill IntMap	`col:"skill" client:"skill"`	 // 主动技能
}

type PetsLevelConfCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Exp int	`col:"exp" client:"exp"`	 // 升级经验值
    AttributePets IntMap	`col:"attributePets" client:"attributePets"`	 // 宠物属性加成
}

type PhantomPhantomCfg struct {
    Phantom int	`col:"phantom" client:"phantom"`	 // 幻灵种类
    Name string	`col:"name" client:"name"`	 // 幻灵名称
    Icon int	`col:"icon" client:"icon"`	 // 幻灵icon
    Model int	`col:"model" client:"model"`	 // 幻灵模型
    Effect int	`col:"effect" client:"effect"`	 // 基础技能
}

type PhantomLevelPhantomLevelCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Phantom int	`col:"phantom" client:"phantom"`	 // 幻灵种类
    Position1 int	`col:"position1" client:"position1"`	 // 技能位置1
    Skill_level1 int	`col:"skill_level1" client:"skill_level1"`	 // 技能等级1
    Talent1 int	`col:"talent1" client:"talent1"`	 // 消耗天赋点1
    Effect1 int	`col:"effect1" client:"effect1"`	 // 技能效果1
    Position2 int	`col:"position2" client:"position2"`	 // 技能位置2
    Skill_level2 int	`col:"skill_level2" client:"skill_level2"`	 // 技能等级2
    Talent2 int	`col:"talent2" client:"talent2"`	 // 消耗天赋点2
    Effect2 int	`col:"effect2" client:"effect2"`	 // 技能效果2
}

type PowerRollPowerRollCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Min float64	`col:"min" client:"min"`	 // 最小系数
    Max float64	`col:"max" client:"max"`	 // 最大系数
    TowerReduce int	`col:"towerReduce" client:"towerReduce"`	 // BOSS受到伤害减免
    PVPReduce int	`col:"PVPReduce" client:"PVPReduce"`	 // pvp伤害减免（万分比）
}

type PreFunctionPreFunctionCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 功能类型
    Condition int	`col:"condition" client:"condition"`	 // 开启条件
    Item ItemInfos	`col:"item" client:"item" checker:"item"`	 // 商品
    BuyTime int	`col:"buyTime" client:"buyTime"`	 // 购买波数
    Price ItemInfo	`col:"price" client:"price" checker:"item"`	 // 购买价格
    Show ItemInfos	`col:"show" client:"show" checker:"item"`	 // 弹窗物品
    Jump int	`col:"jump" client:"jump"`	 // 跳转
}

type PrivilegePrivilegeCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Name string	`col:"name" client:"name"`	 // 特权名称
    Need int	`col:"need" client:"need"`	 // 需激活前置
    Cost ItemInfo	`col:"cost" client:"cost" checker:"item"`	 // 消耗元宝
    Privilege IntMap	`col:"privilege" client:"privilege"`	 // 特权
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 激活奖励
}

type PropertyPropertyCfg struct {
    Id int	`col:"id" client:"id"`	 // 属性id
    Definition string	`col:"definition"`	 // 名称
    Type int	`col:"type" client:"type"`	 // 属性类型
    Name string	`col:"name" client:"name"`	 // 属性名称
    Combat FloatSlice	`col:"combat" client:"combat"`	 // 战斗力系数-固定值int
    Mark int	`col:"mark" client:"mark"`	 // 属性标记
    Display int	`col:"display" client:"display"`	 // 显示
}

type PublicCopyStageCfg struct {
    StageId int	`col:"stageId" client:"stageId"`	 // 关卡id（对应stage表Id）
    Condition IntMap	`col:"condition" client:"condition"`	 // 进入条件
    ConditionType int	`col:"conditionType" client:"conditionType"`	 // 进入条件类型(1:全部满足,2：满足条件之一)
}

type RandRandCfg struct {
    Id int	`col:"id"`	 // 随机id
    Group int	`col:"group"`	 // 随机组
    Quality int	`col:"quality"`	 // 颜色
    Weight int	`col:"weight"`	 // 颜色权重
    Attribute IntSlice2	`col:"attribute"`	 // 属性(属性Id,权重,最小值,最大值|属性Id,权重,最小值,最大值)
}

type RechargeRechargeCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Money int	`col:"money" client:"money"`	 // 金额
    Reward1 ItemInfos	`col:"reward1" client:"reward1" checker:"item"`	 // 正常获得物品
    Reward2 ItemInfos	`col:"reward2" client:"reward2" checker:"item"`	 // 双倍额外获得
    Display string	`col:"display" client:"display"`	 // 充值名称
}

type RedDayMaxRedDayMaxCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Day IntSlice	`col:"day" client:"day"`	 // 开服天数
    Max int	`col:"max" client:"max"`	 // 红包掉落上限（分）
}

type RedRecoveryRedRecoveryCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Item int	`col:"item" client:"item"`	 // 装备道具
}

type ReinReinCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Level int	`col:"level" client:"level"`	 // 转生等级
    Exp int	`col:"exp" client:"exp"`	 // 消耗修为
    AttributeP IntMap	`col:"attributeP" client:"attributeP"`	 // 战士属性
    AttributeM IntMap	`col:"attributeM" client:"attributeM"`	 // 法师属性
    AttributeT IntMap	`col:"attributeT" client:"attributeT"`	 // 道士属性
}

type ReinCostReinCostCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Cost ItemInfo	`col:"cost" client:"cost" checker:"item"`	 // 花费
    Reward ItemInfo	`col:"reward" client:"reward" checker:"item"`	 // 获得道具
    Number int	`col:"number" client:"number"`	 // 每日购买上限
}

type RewardsOnlineAwardCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Time int	`col:"time" client:"time"`	 // 在线时长
    Rewards ItemInfos	`col:"rewards" client:"rewards" checker:"item"`	 // 奖励物品
}

type RingRingCfg struct {
    Ringid int	`col:"ringid" client:"ringid"`	 // 特戒id
    Model int	`col:"model" client:"model"`	 // 特戒模型
    Attr IntMap	`col:"attr" client:"attr"`	 // 属性
    Phantom IntSlice	`col:"phantom" client:"phantom"`	 // 幻灵种类
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 融合消耗戒指
    Consume1 ItemInfo	`col:"consume1" client:"consume1" checker:"item"`	 // 融合消耗材料
}

type RingPhantomRingPhantomCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Order int	`col:"order" client:"order"`	 // 阶数
    Star int	`col:"star" client:"star"`	 // 星数
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 消耗材料
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
    Talent int	`col:"talent" client:"talent"`	 // 累计获得天赋点
}

type RingStrengthenRingStrengthenCfg struct {
    Level int	`col:"level" client:"level"`	 // 等级
    Type int	`col:"type" client:"type"`	 // 升级类型
    Condition IntMap	`col:"condition" client:"condition"`	 // 升级条件
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 消耗
}

type RobotRobotCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Name string	`col:"name" client:"name"`	 // 名字
    Level int	`col:"level" client:"level"`	 // 等级
    Icon IntSlice	`col:"icon" client:"icon"`	 // 头像
    Gender IntSlice	`col:"gender" client:"gender"`	 // 性别
    Job IntSlice	`col:"job" client:"job"`	 // 职业
    Fighting IntSlice	`col:"fighting"`	 // 战力区间
    LevelCompetive int	`col:"levelCompetive" client:"levelCompetive"`	 // 段位
    Property1 IntMap	`col:"property1" client:"property1"`	 // 属性1
    Property2 IntMap	`col:"property2" client:"property2"`	 // 属性2
    Property3 IntMap	`col:"property3" client:"property3"`	 // 属性3
    Skills IntSlice2	`col:"skills" client:"skills"`	 // 技能
    Model1 IntSlice	`col:"model1" client:"model1"`	 // 外显武器
    Model2 IntSlice	`col:"model2" client:"model2"`	 // 外显衣服
    MiningLevel int	`col:"miningLevel" client:"miningLevel"`	 // 矿工等级
    FightingShow int	`col:"fightingShow" client:"fightingShow"`	 // 战力
}

type RoleFirstnameRoleFirstnameCfg struct {
    Id int	`col:"id"`	 // 编号
    FirstName string	`col:"firstName"`	 // 姓
}

type RoleNameBaseCfg struct {
    Id int	`col:"id"`	 // 编号
    Name string	`col:"name"`	 // 名字
    Sex int	`col:"sex"`	 // 标识男名、女名（1男，2女 0不区分）
    Whole int	`col:"whole"`	 // 全名
}

type ScrollingScrollingCfg struct {
    Type int	`col:"type" client:"type"`	 // 活动类型：1.购买白银卡2.购买黄金卡3.购买战令4.购买七日投资5.神机宝库抽到XX品质以上6.多人野外BOSS掉落XX品质以上7.沙巴克开启信息8.暗殿BOSS掉落XX品质以上9.世界首领开启10.购买开服直购11.领取连充豪礼12.协助13.摇彩14.炼狱首领15.远古首领16.打宝地图17.击杀沙巴克城门18.击杀沙巴克BOSS
    Condition int	`col:"condition" client:"condition"`	 // 触发条件
    Txt string	`col:"txt" client:"txt"`	 // 显示文案（ 1：代表玩家，2：道具 ）
}

type SetTypeCfg struct {
    Id int	`col:"id" client:"id"`	 // id
}

type SevenDayInvestSevenDayInvestCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Rewards ItemInfos	`col:"rewards" client:"rewards" checker:"item"`	 // 奖励
}

type ShabakeRewardperShabakeRewardperCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Rank IntSlice	`col:"rank" client:"rank"`	 // 排名
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type ShabakeRewarduniShabakeRewarduniCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Rank IntSlice	`col:"rank" client:"rank"`	 // 排名
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
    RewardShow ItemInfos	`col:"rewardShow" client:"rewardShow" checker:"item"`	 // 奖励展示
}

type ShopTypeCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 商城分类
    Item ItemInfo	`col:"item" client:"item" checker:"item"`	 // 商品
    Price ItemInfo	`col:"price" client:"price" checker:"item"`	 // 消耗货币
    BuyType int	`col:"buyType" client:"buyType"`	 // 购买类型
    Discount int	`col:"discount" client:"discount"`	 // 折扣
    JobId int	`col:"jobId" client:"jobId"`	 // 职业id
    Time int	`col:"time" client:"time"`	 // 购买次数
}

type ShopItemUnitCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    ShopType int	`col:"shopType" client:"shopType"`	 // 商场分类
    Item ItemInfo	`col:"item" client:"item" checker:"item"`	 // 道具
    BuyType int	`col:"buyType" client:"buyType"`	 // 购买类型
    Purchase ItemInfo	`col:"purchase" client:"purchase" checker:"item"`	 // 购买道具
    LimitType int	`col:"limitType" client:"limitType"`	 // 限购类型
    LimitCount int	`col:"limitCount" client:"limitCount"`	 // 限购次数
    OpenDay int	`col:"openDay" client:"openDay"`	 // 开服天数
    Weight int	`col:"weight" client:"weight"`	 // 权重
}

type SignSignCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Reward ItemInfo	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type SkillSkillCfg struct {
    Skillid int	`col:"skillid" client:"skillid"`	 // 技能id
    Role int	`col:"role" client:"role"`	 // 角色
    Job int	`col:"job" client:"job"`	 // 职业
    Type int	`col:"type" client:"type"`	 // 类型
    SkillEffectType int	`col:"skillEffectType" client:"skillEffectType"`	 // 技能效果类型1：攻击技能，2治疗技能
    Target int	`col:"target" client:"target"`	 // 目标类型
    Num_max int	`col:"num_max" client:"num_max"`	 // 目标最大数量
    TargetRoleNum int	`col:"targetRoleNum" client:"targetRoleNum"`	 // 目标角色数量
    MonsterNum int	`col:"monsterNum" client:"monsterNum"`	 // 目标怪物数量
    RangeType int	`col:"rangeType" client:"rangeType"`	 // 施法类型
    Range int	`col:"range" client:"range"`	 // 范围
    Distance int	`col:"distance" client:"distance"`	 // 施法距离
    Position int	`col:"position" client:"position"`	 // 施法朝向 1、自己的朝向，2、目标的朝向（客户端发过来的朝向，可能是鼠标位置，可能是目标位置）
    Aspd bool	`col:"Aspd" client:"Aspd"`	 // 是否受攻速影响
    Untarget bool	`col:"untarget" client:"untarget"`	 // 是否可以无目标释放
    Use bool	`col:"use" client:"use"`	 // 是否可持续使用
    Selectcondition IntSlice	`col:"Selectcondition" client:"Selectcondition"`	 // 选取条件
    Cast int	`col:"Cast" client:"Cast"`	 // 被动是否释放
    CastPreSkill int	`col:"castPreSkill"`	 // 被动施法前置技能
    Skillpoint int	`col:"skillpoint" client:"skillpoint"`	 // 技能作用点
    UsePriority int	`col:"usePriority" client:"usePriority"`	 // 释放优先级
    CastDuration int	`col:"castDuration" client:"castDuration"`	 // 进入战斗&毫秒后触发（仅被动技能和被动释放技能）
    Infinite int	`col:"Infinite" client:"Infinite"`	 // 无限刀
}

type SkillAttackEffectSkillAttackEffectCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Description string	`col:"description" client:"description"`	 // 描述
    Type int	`col:"type"`	 // 类型
    Probability int	`col:"probability"`	 // 概率值
    Target int	`col:"target"`	 // 目标判断
    BuffType IntSlice	`col:"buffType"`	 // 目标buff(buff类型|buffId),buffId未配置或-1时，只判断buff类型
    TargetType int	`col:"targetType"`	 // 判断类型
    BuffLeave int	`col:"buffLeave"`	 // 判断目标BUFF后是否移除
    EnmyAddProp IntSlice	`col:"enmyAddProp"`	 // 敌方效果属性（只在类型1中生效，且配合enmyEffectValue）,降低属性填负值
    EnmyEffectValue IntSlice	`col:"enmyEffectValue"`	 // 敌方值(1类型配合enmyAddProp，其他在3 5类型生效)
    EnmyBuffRemoveLayer int	`col:"enmyBuffRemoveLayer"`	 // 敌方取消层数
    EnmybuffRemoveNum IntSlice	`col:"enmybuffRemoveNum"`	 // 敌方取消buff类型数量（配置buff类型，-1时消除所有增益buff）
    EnmyAddBuff IntSlice	`col:"enmyAddBuff"`	 // 添加buff
    SelfAddProp IntSlice	`col:"selfAddProp"`	 // 己方效果属性（只在类型1中生效，且配合selfEffectValue）,降低属性填负值
    SelfEffectValue IntSlice	`col:"selfEffectValue"`	 // 己方值(1类型配合selfAddProp，其他在3 5类型生效)
    SelfbuffRemoveLayer int	`col:"selfbuffRemoveLayer"`	 // 己方取消层数
    SelfbuffRemoveNum IntSlice	`col:"selfbuffRemoveNum"`	 // 己方取消buff类型数量（配置buff类型，-1时消除所有delbuff）
    SelfaddBuff IntSlice	`col:"selfaddBuff"`	 // 添加buff
    SeckillPro int	`col:"seckillPro"`	 // 秒杀概率
}

type SkillLevelSkillCfg struct {
    Skillid int	`col:"skillid" client:"skillid"`	 // 技能id
    Txt string	`col:"txt" client:"txt"`	 // 文字描述
    Level_open int	`col:"level_open" client:"level_open"`	 // 开放等级
    Level int	`col:"level" client:"level"`	 // 技能等级
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 升级消耗
    Type int	`col:"type" client:"type"`	 // 升级类型
    CD int	`col:"CD" client:"CD"`	 // 技能CD（毫秒）
    MP int	`col:"MP" client:"MP"`	 // 技能蓝量
    Atk FloatSlice	`col:"atk" client:"atk"`	 // 攻击系数
    Atk2 int	`col:"atk2" client:"atk2"`	 // 攻击固定值
    Effects IntSlice	`col:"effects" client:"effects"`	 // 技能效果
    Prebuff IntSlice	`col:"prebuff" client:"prebuff"`	 // 技能释放前buff
    Buff IntSlice	`col:"buff" client:"buff"`	 // buff效果
    Num_max int	`col:"num_max" client:"num_max"`	 // 目标最大数量
    TargetRoleNum int	`col:"targetRoleNum" client:"targetRoleNum"`	 // 目标角色数量
    MonsterNum int	`col:"monsterNum" client:"monsterNum"`	 // 目标怪物数量
    RangeType int	`col:"rangeType" client:"rangeType"`	 // 施法类型
    Range int	`col:"range" client:"range"`	 // 范围
    Distance int	`col:"distance" client:"distance"`	 // 施法距离
    Combat int	`col:"combat" client:"combat"`	 // 战斗力
    Passiveconditions1 int	`col:"Passiveconditions1" client:"Passiveconditions1"`	 // 被动条件1
    Times int	`col:"times" client:"times"`	 // 被动条件1的次数限制
    Passiveconditions2 IntMap	`col:"Passiveconditions2" client:"Passiveconditions2"`	 // 被动条件2
    Perskill int	`col:"perskill"`	 // 触发该技能的概率
    Summonid int	`col:"summonid" client:"summonid"`	 // 召唤物ID
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 常驻属性
}

type StageStageCfg struct {
    Id int	`col:"id" client:"id"`	 // 关卡id
    Name string	`col:"name" client:"name"`	 // 关卡名称
    Type int	`col:"type" client:"type"`	 // 关卡类型1：界面关卡战斗2：个人boss3：爬塔4：野外boss5.世界BOSS6.材料副本7.vipBoss8.主城9.经验副本10.竞技场11.公共打宝地图12.暗殿主界面13.世界BOSS新14.野战15.矿洞战斗16 暗殿boss17.挂机BOSS18.新手村19.公会篝火20.泡点PK21.沙巴克22.世界首领23.跨服沙巴克24.主线插入动画地图25.远古首领26.守卫龙柱27.九重魔塔28.炼狱首领29.炼狱大地图30新版沙巴克31打宝个人地图 33.新手关复古地图
    Mapid int	`col:"mapid" client:"mapid"`	 // 地图id
    Monster_group IntSlice2	`col:"monster_group" client:"monster_group"`	 // 怪物组(怪物组，地图数据怪物区域id|...)
    Monster_num IntMap	`col:"monster_num"`	 // 区域怪物低于多少数量立即刷新（出生点，低于数量|出生点，低于数量）
    ReliveDelay int	`col:"reliveDelay" client:"reliveDelay"`	 // 玩家复活时间 毫秒 0 不复活
    ReliveAddrType IntSlice	`col:"reliveAddrType" client:"reliveAddrType"`	 // 复活地类型（1原地复活，2出生点复活，3回城复活）[自然复活地类型|元宝复活地类型】
    Consume IntMap	`col:"consume" client:"consume"`	 // 复活消耗
    UseItem IntMap	`col:"useItem" client:"useItem"`	 // 可使用道具（道具Id,次数|道具Id,次数...）
    Sound int	`col:"sound" client:"sound"`	 // 背景音乐
    LifeTime int	`col:"lifeTime" client:"lifeTime"`	 // 副本生命周期
    Door IntSlice2	`col:"door" client:"door"`	 // 传送门 地图数据npc标识Id,npcId,stageId ）
    Npc IntSlice2	`col:"npc" client:"npc"`	 // npc（ 地图数据npc标识id,npcId ）
    Get int	`col:"get" client:"get"`	 // 物品拾取
    DropItemDisappearTime int	`col:"dropItemDisappearTime"`	 // 物品掉落存在周期（秒数）
    Collection IntSlice	`col:"collection" client:"collection"`	 // 采集类型
    Camera IntSlice	`col:"camera" client:"camera"`	 // 摄像机固定
    IsPk bool	`col:"isPk" client:"isPk"`	 // 是否可PK
}

type StrengthenStrengthenCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Position int	`col:"position" client:"position"`	 // 部位
    Level int	`col:"level" client:"level"`	 // 等级
    Rate int	`col:"rate" client:"rate"`	 // 成功率(万分比)
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 消耗材料
    Condition IntMap	`col:"condition" client:"condition"`	 // 解锁条件
    AttributeP IntMap	`col:"attributeP" client:"attributeP"`	 // 战士属性
    AttributeM IntMap	`col:"attributeM" client:"attributeM"`	 // 法师属性
    AttributeT IntMap	`col:"attributeT" client:"attributeT"`	 // 道士属性
    IsBreak int	`col:"isBreak" client:"isBreak"`	 // 是否是突破
}

type StrengthenlinkStrengthenCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Condition IntMap	`col:"condition" client:"condition"`	 // 解锁条件
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 连携属性
}

type SummonConfCfg struct {
    Id int	`col:"id" client:"id"`	 // 召唤物id
    Name string	`col:"name" client:"name"`	 // 名称
    Level int	`col:"level" client:"level"`	 // 等级
    Attr IntMap	`col:"attr" client:"attr"`	 // 属性
    HpFix int	`col:"hpFix"`	 // 生命系数（万分比）
    Skills IntSlice	`col:"skills" client:"skills"`	 // 技能（对应技能lvId）
    ChaseDis int	`col:"chaseDis" client:"chaseDis"`	 // 追击距离(格子数)
    Time int	`col:"time" client:"time"`	 // 持续时间
    Max int	`col:"max" client:"max"`	 // 最大召唤数量
    Group IntSlice	`col:"group"`	 // 互斥组别
}

type TalentTalentCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Name string	`col:"name" client:"name"`	 // 名称
}

type TalentGetTalentGetCfg struct {
    Id int	`col:"id" client:"id"`	 // 等级
    TalentCount int	`col:"talentCount" client:"talentCount"`	 // 获得天赋点
}

type TalentLevelTalentLevelCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Requirement1 IntMap	`col:"requirement1" client:"requirement1"`	 // 限制1：单路线消耗点数（天赋路线,点数）
    Requirement2 int	`col:"requirement2" client:"requirement2"`	 // 限制2：总消耗点数
    Count int	`col:"count" client:"count"`	 // 消耗天赋点
    Skill int	`col:"skill" client:"skill"`	 // 触发技能
    Effect int	`col:"effect" client:"effect"`	 // 效果
}

type TalentStageTalengStageCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    TalentID IntSlice	`col:"talentID" client:"talentID"`	 // 天赋
}

type TalentWayTalengWayCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Profession int	`col:"profession" client:"profession"`	 // 职业
    Name string	`col:"name" client:"name"`	 // 名称
}

type TalenteffectTalentCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Des string	`col:"des" client:"des"`	 // 描述
}

type TalentgeneralTalentCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Condition IntMap	`col:"condition" client:"condition"`	 // 达成条件
    Type int	`col:"type" client:"type"`	 // 模块类型
    Icon IntMap	`col:"icon" client:"icon"`	 // 达成效果
    Talentway int	`col:"talentway" client:"talentway"`	 // 天赋ID
    TalentID int	`col:"talentID" client:"talentID"`	 // 天赋ID
    Level int	`col:"level" client:"level"`	 // 技能等级
}

type TaskConditionCfg struct {
    Id int	`col:"id" client:"id"`	 // 任务id
    NextId int	`col:"nextId" client:"nextId"`	 // 下一个任务id
    ConditionType int	`col:"conditionType" client:"conditionType"`	 // 条件类型（对应condition）
    ConditionValue IntSlice	`col:"conditionValue" client:"conditionValue"`	 // 条件值
    AwardZhan ItemInfos	`col:"awardZhan" client:"awardZhan" checker:"item"`	 // 任务奖励
    AwardFa ItemInfos	`col:"awardFa" client:"awardFa" checker:"item"`	 // 任务奖励
    AwardDao ItemInfos	`col:"awardDao" client:"awardDao" checker:"item"`	 // 任务奖励
}

type TitleTitleCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Item ItemInfos	`col:"item" client:"item" checker:"item"`	 // 激活道具
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
    AchievementId int	`col:"achievementId" client:"achievementId"`	 // 成就id
    Type int	`col:"Type" client:"Type"`	 // 类型
    Time int	`col:"time" client:"time"`	 // 时间
}

type TowerTowerCfg struct {
    Id int	`col:"id" client:"id"`	 // 试练塔id
    Stage int	`col:"stage" client:"stage"`	 // 关卡id
    Condition PropInfos	`col:"condition" client:"condition"`	 // 通关条件
    Recommend string	`col:"recommend" client:"recommend"`	 // 推荐战力
    Show ItemInfos	`col:"show" client:"show" checker:"item"`	 // 奖励预览
    DayReward ItemInfos	`col:"dayReward" client:"dayReward" checker:"item"`	 // 每日奖励
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 通关奖励
    Crush string	`col:"Crush" client:"Crush"`	 // 碾压战力
}

type TowerLotteryCircleTowerLotteryCircleCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Circle int	`col:"circle" client:"circle"`	 // 轮数
    ItemId ItemInfos	`col:"itemId" client:"itemId" checker:"item"`	 // 道具id
    Line int	`col:"line" client:"line"`	 // 抽取顺序
}

type TowerRankRewardTowerRankRewardCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Rank IntSlice	`col:"rank" client:"rank"`	 // 排名
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type TowerRewardTowerRewardCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Floor IntSlice	`col:"floor" client:"floor"`	 // 层数区间
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type TreasureTreasureCfg struct {
    Id int	`col:"id" client:"id"`	 // 宝物ID
    Quality int	`col:"quality" client:"quality"`	 // 宝物品质
    Item ItemInfos	`col:"item" client:"item" checker:"item"`	 // 宝物激活碎片
    Type int	`col:"type" client:"type"`	 // 宝物分类
}

type TreasureArtTreasureArtCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    TreasureId int	`col:"treasureId" client:"treasureId"`	 // 宝物ID
    Level int	`col:"level" client:"level"`	 // 注灵等级
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性加成
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 升级消耗
}

type TreasureAwakenTreasureAwakenCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    TruesureId int	`col:"truesureId" client:"truesureId"`	 // 宝物ID
    Consume IntSlice2	`col:"consume" client:"consume"`	 // 觉醒消耗
    Type int	`col:"type" client:"type"`	 // 属性加成类型
    Attribute IntSlice2	`col:"attribute" client:"attribute"`	 // 效果
}

type TreasureDiscountTreasureDiscountCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Condition int	`col:"condition" client:"condition"`	 // 满减金额
    Discount int	`col:"discount" client:"discount"`	 // 优惠价格
}

type TreasureShopTreasureShopCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    ItemId int	`col:"itemId" client:"itemId"`	 // 物品id
    Count int	`col:"count" client:"count"`	 // 数量
    Weight int	`col:"weight" client:"weight"`	 // 权重
    Price ItemInfo	`col:"price" client:"price" checker:"item"`	 // 价格
}

type TreasureStarsTreasureStarsCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    TreasureId int	`col:"treasureId" client:"treasureId"`	 // 宝物ID
    Level int	`col:"level" client:"level"`	 // 宝物星级
    Consume ItemInfos	`col:"consume" client:"consume" checker:"item"`	 // 升星消耗
    Type int	`col:"type" client:"type"`	 // 属性加成类型
    Attribute IntSlice2	`col:"attribute" client:"attribute"`	 // 效果
}

type TreasureSuitTreasureSuitCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    TruesureId IntSlice	`col:"truesureId" client:"truesureId"`	 // 宝物ID
    Type1 int	`col:"type1" client:"type1"`	 // 激活属性加成类型
    Attribute1 IntSlice2	`col:"attribute1" client:"attribute1"`	 // 效果
    Type2 int	`col:"type2" client:"type2"`	 // 3星属性加成类型
    Attribute2 IntSlice2	`col:"attribute2" client:"attribute2"`	 // 效果
    Type3 int	`col:"type3" client:"type3"`	 // 觉醒属性加成类型
    Attribute3 IntSlice2	`col:"attribute3" client:"attribute3"`	 // 效果
}

type TrialTaskTrialTaskCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Condition int	`col:"condition" client:"condition"`	 // 任务类型
    Value IntSlice	`col:"value" client:"value"`	 // 参数
    Rewards ItemInfos	`col:"rewards" client:"rewards" checker:"item"`	 // 奖励
}

type TrialTotalRewardTrialTotalRewardCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    TaskNum int	`col:"taskNum" client:"taskNum"`	 // 需完成任务数量
    TotalReward ItemInfos	`col:"totalReward" client:"totalReward" checker:"item"`	 // 奖励
}

type VipLvlCfg struct {
    Lvl int	`col:"lvl" client:"lvl"`	 // 等级
    Exp int	`col:"exp" client:"exp"`	 // 要求经验
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 礼包物品
    Privilege IntMap	`col:"privilege" client:"privilege"`	 // 特权
    Cost1 ItemInfos	`col:"cost1" client:"cost1" checker:"item"`	 // 礼包原价
    Cost2 ItemInfos	`col:"cost2" client:"cost2" checker:"item"`	 // 礼包购买价格
    Discount int	`col:"discount" client:"discount"`	 // 折扣
}

type VipBossVipBossCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Stageid int	`col:"stageid" client:"stageid"`	 // 关卡id
    MaxNum int	`col:"maxNum" client:"maxNum"`	 // 每日次数
    Condition IntMap	`col:"condition" client:"condition"`	 // 开启条件
}

type WarOrderConditionWarOrderConditionCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Name string	`col:"name" client:"name"`	 // 名称
}

type WarOrderCycleWarOrderCycleCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    StartTime StringSlice	`col:"startTime" client:"startTime"`	 // 开始日期
    Duration int	`col:"duration" client:"duration"`	 // 持续天数
}

type WarOrderCycleTaskWarOrderCycleTaskCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    ConditionType int	`col:"conditionType" client:"conditionType"`	 // 任务类型
    ConditionValue IntSlice	`col:"conditionValue" client:"conditionValue"`	 // 任务参数
    ConditionName string	`col:"conditionName" client:"conditionName"`	 // 任务名字
    Reward ItemInfo	`col:"reward" client:"reward" checker:"item"`	 // 奖励（战令经验）
    TaskCard ItemInfo	`col:"taskCard" client:"taskCard" checker:"item"`	 // 消耗任务卡直接完成
}

type WarOrderExchangeWarOrderExchangeCfg struct {
    Id int	`col:"id" client:"id"`	 // id（周期*10000+序号）
    Item ItemInfo	`col:"item" client:"item" checker:"item"`	 // 道具id
    WarOrderExp ItemInfo	`col:"warOrderExp" client:"warOrderExp" checker:"item"`	 // 兑换所需战令经验数量
    Number int	`col:"number" client:"number"`	 // 兑换限制数量
}

type WarOrderLevelWarOrderLevelCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    WarOrderExp ItemInfo	`col:"warOrderExp" client:"warOrderExp" checker:"item"`	 // 解锁所需战令经验
    Reward1 ItemInfos	`col:"reward1" client:"reward1" checker:"item"`	 // 精英奖励
    Reward2 ItemInfos	`col:"reward2" client:"reward2" checker:"item"`	 // 豪华奖励
}

type WarOrderWeekTaskWarOrderWeekTaskCfg struct {
    Id int	`col:"id" client:"id"`	 // id（周*10000+序号）
    ConditionType int	`col:"conditionType" client:"conditionType"`	 // 任务类型
    ConditionValue IntSlice	`col:"conditionValue" client:"conditionValue"`	 // 任务参数
    ConditionName string	`col:"conditionName" client:"conditionName"`	 // 任务名字
    Reward ItemInfo	`col:"reward" client:"reward" checker:"item"`	 // 奖励（战令经验）
    TaskCard ItemInfo	`col:"taskCard" client:"taskCard" checker:"item"`	 // 消耗任务卡直接完成
}

type WashWashCfg struct {
    Id int	`col:"id" client:"id"`	 // 装备id
    Type int	`col:"type" client:"type"`	 // 装备类型
    Order int	`col:"order" client:"order"`	 // 阶数
    Washrand_group IntSlice	`col:"washrand_group" client:"washrand_group"`	 // 随机属性
}

type WashrandRandCfg struct {
    Id int	`col:"id" client:"id"`	 // 随机id
    Quality int	`col:"quality"`	 // 颜色
    Weight int	`col:"weight"`	 // 颜色权重
    Attribute IntSlice2	`col:"attribute" client:"attribute"`	 // 属性(属性Id,权重,最小值,最大值|属性Id,权重,最小值,最大值)
}

type WingNewWingNewCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Order int	`col:"order" client:"order"`	 // 阶数
    Star int	`col:"star" client:"star"`	 // 星数
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 消耗材料
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
    Resource string	`col:"resource" client:"resource"`	 // 外观
    Condition IntMap	`col:"condition" client:"condition"`	 // 升级条件
}

type WingSpecialWingSpecialCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 技能类型
    Order int	`col:"order" client:"order"`	 // 神翼阶数
    Level int	`col:"level" client:"level"`	 // 等级
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 消耗材料
    Attribute IntMap	`col:"attribute" client:"attribute"`	 // 属性
    Effect int	`col:"effect" client:"effect"`	 // 特殊技能
}

type WorldBossWorldBossCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Stageid int	`col:"stageid" client:"stageid"`	 // 关卡id
    OpenTime HmsTime	`col:"openTime" client:"openTime"`	 // 开启时间
    PrepareTime HmsTime	`col:"prepareTime" client:"prepareTime"`	 // 预开启时间
    Consume IntSlice2	`col:"consume" client:"consume"`	 // 鼓舞消耗
    Limit int	`col:"limit" client:"limit"`	 // 鼓舞次数上限
    Addition IntMap	`col:"Addition" client:"Addition"`	 // 鼓舞加成
    Continue int	`col:"continue" client:"continue"`	 // 持续时间
    Lucky ItemInfos	`col:"lucky" client:"lucky" checker:"item"`	 // 幸运奖励
}

type WorldLeaderConfCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Time HmsTimes	`col:"time" client:"time"`	 // 挑战开启时间
    StageId int	`col:"stageId" client:"stageId"`	 // 关卡信息
    Dropshow ItemInfos	`col:"dropshow" client:"dropshow" checker:"item"`	 // 奖励展示
    LastDrop ItemInfos	`col:"lastDrop" client:"lastDrop" checker:"item"`	 // 最后一击奖励
}

type WorldLeaderRewardWorldLeaderRewardCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    StageId int	`col:"stageId" client:"stageId"`	 // 关卡信息
    Rank IntSlice	`col:"rank" client:"rank"`	 // 排名
    Reward ItemInfos	`col:"Reward" client:"Reward" checker:"item"`	 // 奖励
}

type WorldLevelWorldLevelCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    MinDay int	`col:"minDay" client:"minDay"`	 // 最小开服天数
    MaxDay int	`col:"maxDay" client:"maxDay"`	 // 最大开服天数
    WorldLevel int	`col:"worldLevel" client:"worldLevel"`	 // 世界等级
}

type WorldLevelBuffWorldLevelBuffCfg struct {
    Id int	`col:"id" client:"id"`	 // ID
    Type int	`col:"type" client:"type"`	 // 类型：1大于世界等级，2小于
    LowLevel int	`col:"lowLevel" client:"lowLevel"`	 // 相差最小等级
    TopLevel int	`col:"topLevel" client:"topLevel"`	 // 相差最大等级
    Effect int	`col:"effect" client:"effect"`	 // 增益减益万分比
}

type WorldRankWorldRankCfg struct {
    Rank int	`col:"rank" client:"rank"`	 // 排名
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type XiaoyouxiTowerXiaoyouxiTowerCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type1 int	`col:"type1" client:"type1"`	 // 章节
    Type2 int	`col:"type2" client:"type2"`	 // 关卡
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 通关奖励
    Tower1 IntSlice2	`col:"tower1" client:"tower1"`	 // 己方塔配置(层数，放置类型，模型，数字，运算法则）放置类型：1-空 2-角色 3-装备 4-怪物 模型：运算法则1-加法 2-减法 3-乘法 4-除法 5-开根号 6-平方 7-幂方 8-阶层！
    Tower2 IntSlice2	`col:"tower2" client:"tower2"`	 // 敌方塔配置
    Condition IntMap	`col:"condition" client:"condition"`	 // 开启条件
}

type XunlongXunlongCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 奖池类型
    Type1 int	`col:"type1" client:"type1"`	 // 活动期数
    Time IntSlice	`col:"time" client:"time"`	 // 开启时间
    Reward IntSlice2	`col:"reward" client:"reward"`	 // 奖池选择
    Reward1 IntSlice	`col:"reward1" client:"reward1"`	 // 推荐物品
}

type XunlongPrXunlongPrCfg struct {
    Time int	`col:"time" client:"time"`	 // 每一轮抽取次数
    Probability IntSlice2	`col:"probability"`	 // 权重
    Consume ItemInfo	`col:"consume" client:"consume" checker:"item"`	 // 消耗
}

type XunlongRoundsXunlongRoundsCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type1 int	`col:"type1" client:"type1"`	 // 活动期数
    Rounds int	`col:"rounds" client:"rounds"`	 // 轮数
    Reward ItemInfos	`col:"reward" client:"reward" checker:"item"`	 // 奖励
}

type ZodiacEquipZodiacEquipCfg struct {
    Id int	`col:"id" client:"id"`	 // id
    Type int	`col:"type" client:"type"`	 // 部位
    Condition IntMap	`col:"condition" client:"condition"`	 // 穿戴条件
    Properties IntMap	`col:"properties" client:"properties"`	 // 属性
    Effectid IntSlice	`col:"effectid" client:"effectid"`	 // 套装效果（6件套效果|12件套效果）
}



type InitConf struct {

    CellWH            IntSlice     `conf:"cellWH"  default:"48|32"`
    RoleMaxLv            int     `conf:"roleMaxLv"  default:"0"`
    StageInitId            int     `conf:"stageInitId"  default:"100"`
    ControllDelayTime            int     `conf:"ControllDelayTime"  default:"2000"`
    PixelToJoy            int     `conf:"pixelToJoy"  default:"50"`
    Autopickuptime            int     `conf:"autopickuptime"  default:"1000"`
    SingleBossRefresh            int     `conf:"singleBossRefresh"  default:"10"`
    SingleBossLntervalTime            int     `conf:"singleBossLntervalTime"  default:"7200"`
    FightHitBase            float64     `conf:"fightHitBase"  default:"0.95"`
    FightHitFix            int     `conf:"fightHitFix"  default:"1"`
    FightCritRateLimit            FloatSlice     `conf:"fightCritRateLimit"  default:"0|0.7"`
    FightCritBase            float64     `conf:"fightCritBase"  default:"1.5"`
    FightDeathRate            FloatSlice     `conf:"fightDeathRate"  default:"0|0.5"`
    FightDeathBase            float64     `conf:"fightDeathBase"  default:"1.2"`
    FightCalcTypeRatio            float64     `conf:"fightCalcTypeRatio"  default:"1.2"`
    FightCalcRatio            float64     `conf:"fightCalcRatio"  default:"0.2"`
    HookMapDrop            int     `conf:"hookMapDrop"  default:"60"`
    FieldBossFightNum            int     `conf:"fieldBossFightNum"  default:"10"`
    Speedparam            int     `conf:"speedparam"  default:"350"`
    ExpStageDoubleCost            ItemInfos     `conf:"expStageDoubleCost"  default:"5,200"`
    ExpStageAddNumItemId            int     `conf:"expStageAddNumItemId"  default:"41310"`
    ExpStageFightNum            int     `conf:"expStageFightNum"  default:"3"`
    ExpStageRefreshMonster            int     `conf:"expStageRefreshMonster"  default:"5"`
    ActIdleFrameRate            int     `conf:"actIdleFrameRate"  default:"200"`
    ActRunFrameRate            int     `conf:"actRunFrameRate"  default:"60"`
    ActWalkFrameRate            int     `conf:"actWalkFrameRate"  default:"70"`
    ActAttackFrameRate            int     `conf:"actAttackFrameRate"  default:"60"`
    ActAttack2FrameRate            int     `conf:"actAttack2FrameRate"  default:"60"`
    ActDieFrameRate            int     `conf:"actDieFrameRate"  default:"60"`
    ActMiningFrameRate            int     `conf:"actMiningFrameRate"  default:"60"`
    ActAwaitFrameRate            int     `conf:"actAwaitFrameRate"  default:"60"`
    FieldBossFightCd            int     `conf:"fieldBossFightCd"  default:"30"`
    DropItemOwnerProtectTime            int     `conf:"dropItemOwnerProtectTime"  default:"60"`
    BaseModelScale            int     `conf:"baseModelScale"  default:"100"`
    ArenaRefreshTime            int     `conf:"ArenaRefreshTime"  default:"5"`
    ArenaChallengeNum            int     `conf:"ArenaChallengeNum"  default:"10"`
    ArenaBattleTime            int     `conf:"ArenaBattleTime"  default:"120"`
    ArenaReward            ItemInfos     `conf:"ArenaReward"  default:"12000,10|6,1000"`
    BaseModel            FloatSlice     `conf:"baseModel"  default:"10110000|10120000|20110000|20120000|30110000|30120000"`
    Skillmovevalue            int     `conf:"skillmovevalue"  default:"300"`
    Jobview            FloatSlice     `conf:"jobview"  default:"10|10|10"`
    TaskInitId            int     `conf:"taskInitId"  default:"1"`
    Pickupwaittime            int     `conf:"pickupwaittime"  default:"300"`
    CdHpRecover            int     `conf:"cdHpRecover"  default:"1000"`
    CdMpRecover            int     `conf:"cdMpRecover"  default:"1000"`
    CdPotion            int     `conf:"cdPotion"  default:"1000"`
    CdBackCity            int     `conf:"cdBackCity"  default:"1000"`
    CdRandomStone            int     `conf:"cdRandomStone"  default:"1000"`
    RepairSign            ItemInfo     `conf:"repairSign"  default:"5|10"`
    RankWorshipAward            ItemInfos     `conf:"rankWorshipAward"  default:"5,50"`
    MiningWorkMaxNum            int     `conf:"miningWorkMaxNum"  default:"3"`
    MiningBuyMaxNum            int     `conf:"miningBuyMaxNum"  default:"2"`
    MiningRobMaxNum            int     `conf:"miningRobMaxNum"  default:"5"`
    MiningBuy            ItemInfos     `conf:"miningBuy"  default:"5,10"`
    SkillResetConsume            ItemInfos     `conf:"skillResetConsume"  default:"4030,1"`
    DarkPalaceChallengeTimes            int     `conf:"darkPalaceChallengeTimes"  default:"5"`
    FollowParam            FloatSlice     `conf:"followParam"  default:"1|5|2|6"`
    PetUpLvConsume            IntSlice     `conf:"PetUpLvConsume"  default:"41400"`
    CompetitveTimes            IntSlice     `conf:"competitveTimes"  default:"10|10"`
    CompetitveCost            ItemInfos     `conf:"competitveCost"  default:"41100,1|5,20"`
    CompetitveLevel            IntSlice     `conf:"competitveLevel"  default:"750|4000"`
    CompetitveRank            IntSlice     `conf:"competitveRank"  default:"1|50"`
    HeroBirthOffset            IntSlice2     `conf:"heroBirthOffset"  default:"4,1|1,4"`
    FieldFightTips            float64     `conf:"fieldFightTips"  default:"0.2"`
    FieldFightMaxNum            IntSlice     `conf:"fieldFightMaxNum"  default:"5|5"`
    FieldFightBuyMaxNum            ItemInfos     `conf:"fieldFightBuyMaxNum"  default:"4000,1|5,20"`
    FieldFightCd            int     `conf:"fieldFightCd"  default:"600"`
    FieldFightLevel            IntSlice2     `conf:"fieldFightLevel"  default:"15000,1|11000,2|10000,2"`
    FieldFightBack            ItemInfos     `conf:"fieldFightBack"  default:"17000,3|18000,3|10010,3"`
    EquipClearGradeMin            int     `conf:"equipClearGradeMin"  default:"610"`
    EquipClearConsume            ItemInfos     `conf:"equipClearConsume"  default:"10800,1"`
    EquipClearPropMax            int     `conf:"equipClearPropMax"  default:"5"`
    EquipClearRedMark            int     `conf:"equipClearRedMark"  default:"100"`
    FriendsMaxNum            int     `conf:"friendsMaxNum"  default:"50"`
    BlackListMaxNum            int     `conf:"blackListMaxNum"  default:"30"`
    DarkPalaceBuy            ItemInfos     `conf:"darkPalaceBuy"  default:"109000,1|10,400"`
    NewRewardMaxTimes            int     `conf:"newRewardMaxTimes"  default:"5"`
    ExpPoolLimitMultiple            int     `conf:"expPoolLimitMultiple"  default:"10"`
    FieldBossTimes            IntSlice     `conf:"fieldBossTimes"  default:"8|8"`
    FieldBossCost            ItemInfos     `conf:"fieldBossCost"  default:"110000,1|10,400"`
    OfficialHurtParameter            IntSlice     `conf:"officialHurtParameter"  default:"2|3"`
    PersonalBossTimes            int     `conf:"personalBossTimes"  default:"1"`
    VipBossTimes            int     `conf:"vipBossTimes"  default:"1"`
    CreateGuild            ItemInfo     `conf:"createGuild"  default:"11321|1"`
    ResetTalentPart            ItemInfo     `conf:"resetTalentPart"  default:"11310|1"`
    ResetTalentAll            ItemInfo     `conf:"resetTalentAll"  default:"11311|1"`
    ChangeGuildInterval            int     `conf:"changeGuildInterval"  default:"1800"`
    GuildOpenLv            int     `conf:"guildOpenLv"  default:"1"`
    Impeach            int     `conf:"impeach"  default:"4320"`
    ImpeachSuccess            int     `conf:"impeachSuccess"  default:"30"`
    Chat            IntSlice     `conf:"chat"  default:"2|30"`
    BonefireTime            int     `conf:"bonefireTime"  default:"3"`
    BonefireRewards            ItemInfos     `conf:"bonefireRewards"  default:"1,666|6,1000|11,1"`
    ShabakeTime1            int     `conf:"shabakeTime1"  default:"3"`
    ShabakeTime2            IntSlice     `conf:"shabakeTime2"  default:"3|6"`
    ShabakeTime3            HmsTimes     `conf:"shabakeTime3"  default:"20:00|20:30"`
    ShabakeBuff            IntSlice     `conf:"shabakeBuff"  default:"5|100|3|10000001|50"`
    ShabakePotion            IntSlice     `conf:"shabakePotion"  default:"5|50|30|30|120"`
    ShabakeReward1            ItemInfos     `conf:"shabakeReward1"  default:"100502,1|101000,1|17006,1|880006,1"`
    ShabakeReward2            ItemInfos     `conf:"shabakeReward2"  default:"17004,1|5,100|10112,2|500013,1"`
    ShabakeTime4            int     `conf:"shabakeTime4"  default:"5"`
    ShabakeScore            IntSlice     `conf:"shabakeScore"  default:"1|2|200"`
    AuctionNum            int     `conf:"auctionNum"  default:"5"`
    AuctionWorldTax            int     `conf:"auctionWorldTax"  default:"1500"`
    AuctionUnionTax            int     `conf:"auctionUnionTax"  default:"800"`
    AuctionShare            int     `conf:"auctionShare"  default:"2000"`
    ChatRecord            int     `conf:"chatRecord"  default:"1"`
    Vip            int     `conf:"vip"  default:"10"`
    BagMaxNum            int     `conf:"bagMaxNum"  default:"100"`
    WarehouseMaxNum            int     `conf:"warehouseMaxNum"  default:"400"`
    ResetFitSkill            ItemInfo     `conf:"resetFitSkill"  default:"11606|1"`
    FitPropFix            IntSlice     `conf:"fitPropFix"  default:"10000|10000|10000"`
    FitPropHpFix            int     `conf:"fitPropHpFix"  default:"8000"`
    DayRankOpenAndEndTime            HmsTimes     `conf:"dayRankOpenAndEndTime"  default:"0:00|23:59:59"`
    DayRankSendRewardTime            HmsTime     `conf:"dayRankSendRewardTime"  default:"23:00"`
    FitModel            int     `conf:"fitModel"  default:"50000005"`
    FitGeneralSkill            int     `conf:"fitGeneralSkill"  default:"13001"`
    OffLine            int     `conf:"offLine"  default:"480"`
    FitChangeEffect            IntSlice     `conf:"fitChangeEffect"  default:"69001|69002"`
    ResurrectionEffect            IntSlice     `conf:"resurrectionEffect"  default:"68006"`
    CrossArenaGamble            ItemInfos     `conf:"crossArenaGamble"  default:"6,1000000|6,2000000"`
    BornMap            IntSlice     `conf:"bornMap"  default:"21|100000"`
    Dailypack            HmsTime     `conf:"dailypack"  default:"00:00:00"`
    Weekpack            int     `conf:"weekpack"  default:"1"`
    Pickup            int     `conf:"pickup"  default:"1600"`
    GrowFund            int     `conf:"growFund"  default:"98"`
    Recharge            int     `conf:"recharge"  default:"15"`
    GrowFundGet            ItemInfo     `conf:"growFundGet"  default:"5|100"`
    MainTaskGuide            IntSlice     `conf:"mainTaskGuide"  default:"21|71|30"`
    WarOrderLuxury            int     `conf:"warOrderLuxury"  default:"98"`
    WarOrderExpBuy            IntSlice     `conf:"warOrderExpBuy"  default:"6|1000"`
    PaoDianEffect            IntSlice     `conf:"paoDianEffect"  default:"70011"`
    WorldBossNoonOpenAndEndTime            HmsTimes     `conf:"worldBossNoonOpenAndEndTime"  default:"12:00|14:00"`
    WorldBossNightOpenAndEndTime            HmsTimes     `conf:"worldBossNightOpenAndEndTime"  default:"18:00|20:00"`
    RedTime            int     `conf:"redTime"  default:"5"`
    RedTime1            int     `conf:"redTime1"  default:"3"`
    KuafushabakeTime1            int     `conf:"kuafushabakeTime1"  default:"13"`
    KuafushabakeTime2            IntSlice     `conf:"kuafushabakeTime2"  default:"4|7"`
    KuafushabakeTime3            HmsTimes     `conf:"kuafushabakeTime3"  default:"20:00|20:30"`
    KuafushabakeBuff            IntSlice     `conf:"kuafushabakeBuff"  default:"5|100|3|10000001|50"`
    KuafushabakePotion            IntSlice     `conf:"kuafushabakePotion"  default:"5|50|30|30|120"`
    KuafushabakeReward1            ItemInfos     `conf:"kuafushabakeReward1"  default:"17006,1|5,500|41408,100|500014,20"`
    KuafushabakeReward2            ItemInfos     `conf:"kuafushabakeReward2"  default:"10010,40|16000,40|6,1200000"`
    KuafushabakeTime4            int     `conf:"kuafushabakeTime4"  default:"5"`
    KuafushabakeScore            IntSlice     `conf:"kuafushabakeScore"  default:"1|2|100"`
    Draw            ItemInfos     `conf:"draw"  default:"4010,1|4010,10|5,1000|5,10000"`
    DrawMax            int     `conf:"drawMax"  default:"1000"`
    DrawMark            int     `conf:"drawMark"  default:"10"`
    DrawRecord            IntSlice     `conf:"drawRecord"  default:"5|6"`
    ItemTipsLimit            IntSlice     `conf:"itemTipsLimit"  default:"5|10"`
    XunlongConsume            IntSlice     `conf:"xunlongConsume"  default:"4020|10"`
    XunlongBuy            IntSlice2     `conf:"xunlongBuy"  default:"1,1,1000,10|2,1,8,999|2,10,68,999"`
    XunlongShow            int     `conf:"xunlongShow"  default:"20"`
    DrawShow            int     `conf:"drawShow"  default:"20"`
    WarOrderChallengeCard            int     `conf:"warOrderChallengeCard"  default:"11711"`
    XunlongRecord            IntSlice     `conf:"xunlongRecord"  default:"5|6"`
    PropTipsLimit            IntSlice     `conf:"propTipsLimit"  default:"10|100|1000"`
    ShengLingDian            ItemInfos     `conf:"shengLingDian"  default:"51000,1|51001,1|51002,1"`
    WorldLeaderBuff            IntSlice     `conf:"worldLeaderBuff"  default:"5|100|3|10000007|50"`
    ShengLingDian1            ItemInfos     `conf:"shengLingDian1"  default:"51000,30|51001,60|51002,100"`
    ResetHolyBeast            ItemInfos     `conf:"resetHolyBeast"  default:"50000,1|50001,1|50002,1|50003,1"`
    BagInterval            int     `conf:"bagInterval"  default:"3"`
    BackMusic            int     `conf:"backMusic"  default:"3003"`
    NoFashion            IntSlice     `conf:"noFashion"  default:"20110100|20120100|10100010"`
    ShengLingGuo            IntSlice     `conf:"ShengLingGuo"  default:"51003"`
    CompetitveSeason            int     `conf:"competitveSeason"  default:"7"`
    TowSweepMax            int     `conf:"towSweepMax"  default:"50"`
    PlayerCountLimit            IntSlice     `conf:"playerCountLimit"  default:"25|10"`
    FitPVP            int     `conf:"FitPVP"  default:"3000"`
    CompetitveVip            int     `conf:"competitveVip"  default:"2"`
    EffectCountLimit            int     `conf:"effectCountLimit"  default:"10"`
    StrongerGuidance            IntSlice     `conf:"StrongerGuidance"  default:"121|123|124|125|126|127"`
    DayRankingshow            int     `conf:"dayRankingshow"  default:"50"`
    AuctionMoney            int     `conf:"auctionMoney"  default:"14"`
    HookMapMax            int     `conf:"hookMapMax"  default:"28800"`
    FirstRecharge            int     `conf:"firstRecharge"  default:"193"`
    BossWindowTime            int     `conf:"bossWindowTime"  default:"10"`
    LevelAddTimes            IntSlice2     `conf:"levelAddTimes"  default:"41310,3|41302,3|41301,3|41303,3|41304,3"`
    SpendrebatesRefresh            HmsTime     `conf:"spendrebatesRefresh"  default:"0:00:00"`
    DayRankngRefresh            HmsTime     `conf:"dayRankngRefresh"  default:"0:00:00"`
    ExperienceCost            ItemInfos     `conf:"experienceCost"  default:"5,50"`
    WingCost            ItemInfos     `conf:"wingCost"  default:"5,50"`
    CoinCost            ItemInfos     `conf:"coinCost"  default:"5,50"`
    StrengthCost            ItemInfos     `conf:"strengthCost"  default:"5,50"`
    FirstRechargeRefresh            HmsTime     `conf:"firstRechargeRefresh"  default:"0:00:00"`
    FindResourcesConsume            int     `conf:"findResourcesConsume"  default:"5"`
    InvestCost            int     `conf:"investCost"  default:"98"`
    ContRecharge            int     `conf:"contRecharge"  default:"30"`
    OpenGiftTime            IntSlice     `conf:"openGiftTime"  default:"15|0|0"`
    VipServiceExplain            string     `conf:"vipServiceExplain"  default:"添加专属客服之后，可直接联系咨询和处理游戏问题，还能获得免费礼包哦"`
    VipServiceTips            string     `conf:"vipServiceTips"  default:"客服在线时间：9:00-18:00"`
    NoticeRemoveSpeed            int     `conf:"NoticeRemoveSpeed"  default:"200"`
    ExpStageCondition            int     `conf:"expStageCondition"  default:"7"`
    FireSkills            IntSlice     `conf:"fireSkills"  default:"20100|20200"`
    GuildCombat            int     `conf:"guildCombat"  default:"100000"`
    GuildNotice            string     `conf:"guildNotice"  default:"欢迎各位加入本公会，兄弟们努力提升，抢夺皇城。击杀首领、闯关通天塔，可获得大量材料。参与排行榜比拼，竞争排名奖励。开通特权卡，自动挑战挂机BOSS。"`
    JinDingRate            IntSlice     `conf:"JinDingRate"  default:"1|500"`
    YuanBaoRate            IntSlice     `conf:"YuanBaoRate"  default:"10|1"`
    GuajiBuff            int     `conf:"guajiBuff"  default:"21"`
    ElfPosDelta            IntSlice2     `conf:"elfPosDelta"  default:"20,-150|-20,-150|-20,-150|-20,-150|-20,-150,|20,-150|20,-150|20,-150"`
    ResetCutCD            IntSlice     `conf:"resetCutCD"  default:"90004|90005|90006"`
    AncientBossTimes            IntSlice     `conf:"ancientBossTimes"  default:"3|3"`
    AncientBossCost            ItemInfos     `conf:"ancientBossCost"  default:"130000,1|10,400"`
    BossRevive            IntSlice     `conf:"BossRevive"  default:"1|8|11|12|18|20"`
    AncientSkillCost            ItemInfos     `conf:"ancientSkillCost"  default:"11601,1|11611,1|11612,1|11613,1"`
    Elo            int     `conf:"elo"  default:"2"`
    GuardBuff            IntSlice     `conf:"guardBuff"  default:"5|100|3|20000001|50"`
    StartTime            int     `conf:"startTime"  default:"60"`
    NextTime            int     `conf:"nextTime"  default:"20"`
    MiningRobot            int     `conf:"miningRobot"  default:"1800"`
    MiningRobotcondition            int     `conf:"miningRobotcondition"  default:"10"`
    TruesureReset            ItemInfos     `conf:"truesureReset"  default:"5,100"`
    FitMonster            int     `conf:"FitMonster"  default:"5"`
    TreasureNum            int     `conf:"treasureNum"  default:"6"`
    TreasureBuyTime            int     `conf:"treasureBuyTime"  default:"5"`
    TreasureTime            int     `conf:"treasureTime"  default:"60"`
    TreasureCost            ItemInfo     `conf:"treasureCost"  default:"5|20"`
    DarkBossHelp            int     `conf:"darkBossHelp"  default:"5"`
    HellBossHelp            int     `conf:"hellBossHelp"  default:"3"`
    HelpTime            int     `conf:"helpTime"  default:"10"`
    AcceptHelpTime            int     `conf:"acceptHelpTime"  default:"10"`
    MagicTowerStageId            IntSlice2     `conf:"magicTowerStageId"  default:"508,60,10"`
    MagicTowerMark            int     `conf:"magicTowerMark"  default:"50"`
    PVPfight            int     `conf:"PVPfight"  default:"3300"`
    HellBossTime            int     `conf:"hellBossTime"  default:"3"`
    HellBossBuy            ItemInfos     `conf:"hellBossBuy"  default:"120000,1|10,400"`
    HellBossAdd            int     `conf:"hellBossAdd"  default:"3"`
    HelpWait            int     `conf:"helpWait"  default:"60"`
    AllBuy            int     `conf:"allBuy"  default:"1000"`
    BuyNum            int     `conf:"buyNum"  default:"100"`
    BuyTime            HmsTimes     `conf:"buyTime"  default:"10:00:00|15:00:00"`
    RewardTime            HmsTimes     `conf:"rewardTime"  default:"20:00:00"`
    OnePrice            ItemInfos     `conf:"onePrice"  default:"5,10"`
    LuckyItem            ItemInfos     `conf:"luckyItem"  default:"5,10"`
    LuckyGold            ItemInfos     `conf:"luckyGold"  default:"610000,1"`
    LuckyReward            ItemInfos     `conf:"luckyReward"  default:"610001,1"`
    NotLuckyGold            float64     `conf:"notLuckyGold"  default:"1.2"`
    NotLuckReward            ItemInfos     `conf:"notLuckReward"  default:"610002,1"`
    BulletScreenNum            int     `conf:"bulletScreenNum"  default:"6"`
    BulletScreenSpeed            int     `conf:"bulletScreenSpeed"  default:"200"`
    LimitCount            int     `conf:"limitCount"  default:"3000"`
    GrowCount            int     `conf:"growCount"  default:"100"`
    YaoCaiTime            int     `conf:"yaoCaiTime"  default:"2"`
    YaoCaiTime1            IntSlice     `conf:"yaoCaiTime1"  default:"100|0|0"`
    TreasureTime1            IntSlice     `conf:"treasureTime1"  default:"100|0|0"`
    TrialPlayerModel            int     `conf:"trialPlayerModel"  default:"40000039"`
    CrossArenaGroping            int     `conf:"crossArenaGroping"  default:"300"`
    TrialDuration            int     `conf:"trialDuration"  default:"7"`
    LotteryChance            int     `conf:"lotteryChance"  default:"5"`
    TrialWeaponModel            IntSlice     `conf:"trialWeaponModel"  default:"34|18"`
    MagicTowerMonster            IntSlice2     `conf:"magicTowerMonster"  default:"100502,1|100507,2|100508,3|100509,4|100510,5|100511,5|100512,5|100513,5|100503,500|100504,500|100505,500|100506,500|100514,500|100515,500|100516,500|100517,500|100518,500"`
    JewelGet            int     `conf:"jewelGet"  default:"166"`
    ScrollingEquipt            IntSlice2     `conf:"scrollingEquipt"  default:"1,4,50|3,5,200|7,6,300"`
    HookMapShow	            int     `conf:"hookMapShow"  default:"5"`
    AncientBossTime            int     `conf:"ancientBossTime"  default:"3"`
    AncientBossBuy            ItemInfos     `conf:"ancientBossBuy"  default:"130000,1|10,400"`
    AncientBossAdd            int     `conf:"ancientBossAdd"  default:"3"`
    ElfRecoverLimit            ItemInfos     `conf:"elfRecoverLimit"  default:"5,500"`
    DaBaoMysteryEnergy            int     `conf:"daBaoMysteryEnergy"  default:"100"`
    DaBaoMysteryEnergyResume            IntSlice     `conf:"daBaoMysteryEnergyResume"  default:"600|1"`
    DaBaoMysteryEnergyItem            ItemInfo     `conf:"daBaoMysteryEnergyItem"  default:"10201|1"`
    ShabakeMonsterBUff            IntSlice2     `conf:"shabakeMonsterBUff"  default:"102000,10000005,1000|102001,10000006,1000"`
    ShabakeStop            int     `conf:"shabakeStop"  default:"6"`
    ShabakeNPC1            IntSlice     `conf:"shabakeNPC1"  default:"81|0|29"`
    ShabakeNPC2            IntSlice     `conf:"shabakeNPC2"  default:"82|500|2000"`
    ShabakeScore2            int     `conf:"shabakeScore2"  default:"200"`
    ShabakeTime6            int     `conf:"shabakeTime6"  default:"5"`
    ShabakeTime7            int     `conf:"shabakeTime7"  default:"3"`
    ShabakeScore3            IntSlice2     `conf:"shabakeScore3"  default:"1,15,30|30,15,45"`
    TreasureTime2            int     `conf:"treasureTime2"  default:"2"`
    XiaoYouXiEnergy            int     `conf:"xiaoYouXiEnergy"  default:"100"`
    XiaoYouXiEnergyResume            IntSlice     `conf:"xiaoYouXiEnergyResume"  default:"3600|1"`
    XiaoYouXiCostEnergy            int     `conf:"xiaoYouXiCostEnergy"  default:"5"`
    BeginShotSpeed            int     `conf:"beginShotSpeed"  default:"250"`
    BulletMoveSpeed            int     `conf:"bulletMoveSpeed"  default:"10"`
    RoleMoveSpeed            int     `conf:"roleMoveSpeed"  default:"5"`
    RoleWeight            IntSlice     `conf:"roleWeight"  default:"100|100"`
    MagicTowerView            int     `conf:"magicTowerView"  default:"20"`
    MonsterRangeTime            int     `conf:"monsterRangeTime"  default:"5"`
    RoleAcrossSpeed            int     `conf:"roleAcrossSpeed"  default:"8"`
    ArcherMovePixel            int     `conf:"archerMovePixel"  default:"1"`
    BulletSize            IntSlice     `conf:"bulletSize"  default:"10|20"`
    BulletShoot            int     `conf:"bulletShoot"  default:"-100"`
    DaBaoMysteryOut            int     `conf:"daBaoMysteryOut"  default:"15"`
    XiaoyouxiTowerSounds            IntMap     `conf:"xiaoyouxiTowerSounds"  default:"1,10000000|2,10002003|3,10002002|4,10002001|5,10002008|6,0|7,10002006|8,10002007"`
    ShabakeGateRange            int     `conf:"shabakeGateRange"  default:"5"`
    ShabakeBossRange            int     `conf:"shabakeBossRange"  default:"10"`
    ShabakeRange            int     `conf:"shabakeRange"  default:"20"`
    ShabakeReward3            ItemInfos     `conf:"shabakeReward3"  default:"11723,20|11723,20|11723,20|11723,20|11723,20|11723,20|11723,20|11723,20|11723,20|11723,20|11821,5|11821,5|11821,5|11821,5|11821,5|11821,5|11822,5|11822,5|11822,5|11822,5|11822,5|11822,5|11823,20|11823,20|11823,20|11823,20|11823,20|11823,20|11823,20|11823,20|11823,20|11823,20|16000,20|16000,20|16000,20|16000,20|16000,20|16000,20|16000,20|16000,20|16000,20|16000,20|27501,50|27501,50|27503,5|27503,5|27503,5|27503,5|51003,5|51003,5|51003,5|51003,5|51003,5|51003,5|60002,5|60002,5|60002,5|60002,5|41000,1|41000,1|10012,5|10012,5|10012,5|10012,5|17007,1|17007,1"`
    ShabakeReward4            ItemInfos     `conf:"shabakeReward4"  default:"100502,1|101000,1|880006,1"`
    ShabakeReward5            ItemInfos     `conf:"shabakeReward5"  default:"880006,1"`
    TreasureType            IntSlice     `conf:"treasureType"  default:"601|603|600|219|217|218|618|1131|227|1061|228|235|240|614|245|1151|1021|607|1011|615|605|606|1141|604|221"`
    DaBaoMystery            int     `conf:"daBaoMystery"  default:"80"`
    ShabakePalace            int     `conf:"shabakePalace"  default:"30"`
    ShabakeNPC3            IntSlice     `conf:"shabakeNPC3"  default:"80|100|7"`
    TreasureReset            ItemInfos     `conf:"treasureReset"  default:"5,500"`
    AttackEnemySounds            IntMap     `conf:"attackEnemySounds"  default:"1,10000000|2,10001001|3,10001002|4,10001003|5,10001004|6,10001005|7,10001006"`
    WashText            string     `conf:"washText"  default:"转生装备才可洗练"`
    AcherSounds            IntMap     `conf:"acherSounds"  default:"1,10000000|2,10003001|3,10003002|4,10003003|5,10003004|6,10003005|7,10003006|8,10003007|9,10003008|10,10003009|11,10003010|12,10003011|13,10003012|14,10001005|15,10001006"`
    DailyTime            IntMap     `conf:"dailyTime"  default:"20,155"`
    WindowAiTime            int     `conf:"windowAiTime"  default:"3"`
    HookMapBossTime1            int     `conf:"hookMapBossTime1"  default:"750"`
    HookMapBossTime2            int     `conf:"hookMapBossTime2"  default:"1000"`
    HookMapButton            IntSlice     `conf:"hookMapButton"  default:"10000|1000"`
    OpenGiftDisplay            string     `conf:"openGiftDisplay"  default:"成长基金"`
    WarOrderLuxuryDisplay            string     `conf:"warOrderLuxuryDisplay"  default:"豪华战令"`
    WarOrderExpBuyDisplay            string     `conf:"warOrderExpBuyDisplay"  default:"战令经验"`
    InvestCostDisplay            string     `conf:"investCostDisplay"  default:"七日投资"`
    XunlongDisplay            string     `conf:"xunlongDisplay"  default:"寻龙探宝"`
    IosShieldFuction            IntSlice     `conf:"iosShieldFuction"  default:"146|103|162|193|101|104|102|182|106|269"`
    FitUpgradeRedDot            int     `conf:"fitUpgradeRedDot"  default:"5"`
    PetsLevelButton            IntSlice     `conf:"petsLevelButton"  default:"2|4|10"`
    BagRedDot            int     `conf:"bagRedDot"  default:"10"`
    ResetFightStage            int     `conf:"resetFightStage"  default:"5000"`
    BossRemind            int     `conf:"BossRemind"  default:"3"`
    Weixinhookmap1            string     `conf:"weixinhookmap1"  default:"8小时挂机上限已完成，请返回游戏领取奖励"`
    Weixinhookmap2            string     `conf:"weixinhookmap2"  default:"挂机奖励弹窗"`
    Weixinhookmap3            int     `conf:"weixinhookmap3"  default:"8"`
    ElfNumTips            int     `conf:"elfNumTips"  default:"20"`
    StrengthRedMark            int     `conf:"strengthRedMark"  default:"70"`
    InsideRedMark            int     `conf:"insideRedMark"  default:"3"`
    FirstDropOpenTime            IntSlice2     `conf:"firstDropOpenTime"  default:"306,1,25920000"`
    RedRecoveryCount            int     `conf:"redRecoveryCount"  default:"30"`
    RedRecoveryTime            IntSlice     `conf:"redRecoveryTime"  default:"1|-1"`
    BossFamilyShowRecharge            IntSlice     `conf:"bossFamilyShowRecharge"  default:""`
    EscapeTime            int     `conf:"escapeTime"  default:"10000"`
    AutoTime            int     `conf:"autoTime"  default:"10000"`
    PrivilegeRedMarkTime            int     `conf:"privilegeRedMarkTime"  default:"600"`
    BossFamilyView            int     `conf:"bossFamilyView"  default:"20"`
    PickupView            int     `conf:"pickupView"  default:"6"`
    TaskLimit            int     `conf:"taskLimit"  default:"21"`
    FashionCost            ItemInfos     `conf:"fashionCost"  default:"5,50"`
    BossFamilyShowRechargeTotal            IntSlice     `conf:"bossFamilyShowRechargeTotal"  default:"312"`
    RecoverEffective            int     `conf:"recoverEffective"  default:"1"`
    SoundInterval            int     `conf:"soundInterval"  default:"800"`
    RedRate            int     `conf:"redRate"  default:"100"`
    AutoCD            int     `conf:"autoCD"  default:"3"`


}
