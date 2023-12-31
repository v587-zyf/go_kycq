package constBag

var (
	OpTypeCommonShopBuy     = opTypeInit(1, "普通商城")
	OpTypeCreateRole        = opTypeInit(2, "创建角色")
	OpTypeGoodsSell         = opTypeInit(3, "物品出售")
	OpTypeDebugAddGoods     = opTypeInit(4, "debug添加")
	OpTypeTaskDone          = opTypeInit(5, "任务获取")
	OpTypeBagSpaceAdd       = opTypeInit(6, "背包格子扩充")
	OpTypeNormalStage       = opTypeInit(7, "普通关卡")
	OpTypeEquipChange       = opTypeInit(8, "装备替换")
	OpTypeEquipRecover      = opTypeInit(9, "装备回收")
	OpTypeEquipStrengthen   = opTypeInit(10, "装备强化")
	OpTypeFabaoActive       = opTypeInit(11, "法宝激活")
	OpTypeFabaoUpLevel      = opTypeInit(12, "法宝升级")
	OpTypeFabaoSkillActive  = opTypeInit(13, "法宝技能激活")
	OpTypeArtifactActive    = opTypeInit(14, "神兵激活")
	OpTypeArtifactUpAdvance = opTypeInit(15, "神兵升阶")
	OpTypeArtifactUpLevel   = opTypeInit(16, "神兵升级")
	OpTypePersonBossDare    = opTypeInit(17, "个人boss挑战")
	OpTypeWingUpLevel       = opTypeInit(18, "羽翼升级")
	OpTypeMail              = opTypeInit(19, "邮件")
	OpTypeMailAll           = opTypeInit(20, "邮件一键领取")
	OpTypeReinCostBuy       = opTypeInit(21, "购买转生修为丹")
	OpTypeReinCostUse       = opTypeInit(22, "使用转生修为丹")
	OpTypeShopBuy           = opTypeInit(23, "商城购买")
	OpTypeShopFlush         = opTypeInit(24, "商城刷新")
	OpTypeAtlasActive       = opTypeInit(25, "图鉴激活")
	OpTypeAtlasUpStar       = opTypeInit(26, "图鉴升星")
	OpTypeAtlasGatherActive = opTypeInit(27, "图鉴集合激活")
	OpTypeAtlasGatherUpStar = opTypeInit(28, "图鉴集合升星")
	OpTypeTowerFight        = opTypeInit(30, "爬塔")
	OpTypeTowerLottery      = opTypeInit(31, "爬塔抽奖")
	OpTypeTowerDayAward     = opTypeInit(32, "爬塔每日奖励")
	OpTypeFieldBossFight    = opTypeInit(33, "野外首领")
	OpTypeWorldBossFight    = opTypeInit(34, "世界boss")
	OpTypeMaterialStage     = opTypeInit(35, "材料副本")
	OpTypeEquipRemove       = opTypeInit(36, "装备卸下")
	OpTypeVipBossFight      = opTypeInit(37, "vipboss")
	OpTypeExpStage          = opTypeInit(38, "经验副本")

	OpTypeArena              = opTypeInit(40, "竞技场")
	OpTypeZodiac             = opTypeInit(41, "生肖装备")
	OpTypeKingarms           = opTypeInit(42, "帝器")
	OpTypeCompose            = opTypeInit(43, "合成")
	OpTypeDragonEquip        = opTypeInit(44, "龙器")
	OpTypeItemUse            = opTypeInit(45, "道具使用")
	OpTypeAwaken             = opTypeInit(45, "觉醒")
	OpTypeTask               = opTypeInit(50, "任务")
	OpTypeBless              = opTypeInit(46, "祝福油")
	OpTypeClear              = opTypeInit(47, "洗练")
	OpTypeDictate            = opTypeInit(48, "主宰装备")
	OpTypeWingSpecialUp      = opTypeInit(49, "羽翼特殊属性升级")
	OpTypeEnterPublicCopy    = opTypeInit(50, "进入公共副本")
	OpTypePanacea            = opTypeInit(51, "灵丹")
	OpTypeJewel              = opTypeInit(52, "宝石")
	OpTypeSign               = opTypeInit(53, "签到")
	OpTypeInside             = opTypeInit(54, "内功")
	OpTypeHolyarms           = opTypeInit(55, "神兵")
	OpTypeRing               = opTypeInit(56, "特戒")
	OpTypeMining             = opTypeInit(57, "挖矿")
	OpTypeMiningRobFight     = opTypeInit(58, "抢矿")
	OpTypeMiningRobBackFight = opTypeInit(59, "挖矿夺回奖励")
	OpTypeOnlineAward        = opTypeInit(60, "在线奖励")
	OpTypeRankWorship        = opTypeInit(61, "排行榜膜拜")
	OpTypeSKillUpLv          = opTypeInit(62, "技能升级")
	OpTypeSkillReset         = opTypeInit(63, "技能重置")
	OpTypePetActive          = opTypeInit(64, "战宠激活")
	OpTypePetUpLv            = opTypeInit(65, "战宠升级")
	OpTypePetUpGrade         = opTypeInit(66, "战宠升阶")
	OpTypePetBreak           = opTypeInit(67, "战宠突破")
	OpTypeArea               = opTypeInit(68, "领域")
	OpTypeDarkPalace         = opTypeInit(69, "暗殿")
	OpTypeMagicCircle        = opTypeInit(70, "法阵")
	OpTypeWareHouseSpaceAdd  = opTypeInit(71, "仓库扩容")
	OpTypeEquipMoveToWear    = opTypeInit(72, "装备移动到仓库")
	OpTypeEquipMoveToBag     = opTypeInit(73, "装备移动到背包")
	OpTypeEquipDestroy       = opTypeInit(74, "装备销毁")
	OpTypeNormalStageBoss    = opTypeInit(75, "挂机BOSS")
	OpTypeTalent             = opTypeInit(76, "天赋")
	OpTypePaoDian            = opTypeInit(77, "泡点pk")
	OpTypeVip                = opTypeInit(78, "vip")
	OpTypeFit                = opTypeInit(79, "合体")
	OpTypeOfficialUpLv       = opTypeInit(80, "官职")
	OpTypeFashionUpLv        = opTypeInit(81, "时装（激活升级）")

	OpTypeMonthCard                          = opTypeInit(83, "月卡")
	OpTypeGodEquipActive                     = opTypeInit(90, "神兵激活")
	OpTypeGodEquipUpLv                       = opTypeInit(91, "神兵升级")
	OpTypeJuexueUpLv                         = opTypeInit(92, "绝学激活升级")
	OpTypeCompetitveBuyChallenge             = opTypeInit(93, "竞技场购买挑战次数")
	OpTypeCompetitveDailyReward              = opTypeInit(94, "领取每日奖励")
	OpTypeFieldFightBuyChallenge             = opTypeInit(95, "野战购买挑战次数")
	OpTypeFieldFightChallengeWinReward       = opTypeInit(96, "野战挑战胜利奖励")
	OpTypeFieldBackFightChallengeWinReward   = opTypeInit(97, "野战反击奖励")
	OpTypeFirstRecharge                      = opTypeInit(98, "首充")
	OpTypeSpendRebates                       = opTypeInit(99, "消费返利")
	OpTypeFightRelive                        = opTypeInit(100, "复活")
	OpTypeCreateGuild                        = opTypeInit(101, "创建门派")
	OpTypeGuildBonfireAddExp                 = opTypeInit(102, "篝火增加经验")
	OpTypeAuctionPutAwayItem                 = opTypeInit(103, "拍卖行上架物品")
	OpTypeAuctionGotItem                     = opTypeInit(104, "拍卖行获得道具")
	OpTypeAuctionbidding                     = opTypeInit(105, "拍卖行竞价")
	OpTypeDailyTaskBuyChallenge              = opTypeInit(106, "每日任务 购买挑战次数")
	OpTypeDailyTaskGetAward                  = opTypeInit(107, "每日任务 领取奖励")
	OpTypeDailyTaskResourcesBackGetReward    = opTypeInit(108, "每日资源找回 领取奖励")
	OpTypeOpenGiftReward                     = opTypeInit(109, "礼包奖励")
	OpTypeOfflineReward                      = opTypeInit(110, "离线奖励")
	OpTypeGiftCodeReward                     = opTypeInit(111, "礼包码")
	OpTypeAchieveGetReward                   = opTypeInit(112, "成就奖励")
	OpTypeComposeEquip                       = opTypeInit(113, "合成装备")
	OpTypeLimitedGift                        = opTypeInit(114, "限时礼包")
	OpTypeDailyPack                          = opTypeInit(115, "每日礼包")
	OpTypeRecharge                           = opTypeInit(116, "充值")
	OpTypeGrowFund                           = opTypeInit(116, "成长基金")
	OpTypeChallengeBottom                    = opTypeInit(116, "擂台赛下注")
	OpTypeWarOrderTaskFininsh                = opTypeInit(117, "战令任务完成")
	OpTypeWarOrderTaskReward                 = opTypeInit(118, "战令任务奖励")
	OpTypeWarOrderBuyLuxury                  = opTypeInit(119, "战令购买豪华")
	OpTypeWarOrderBuyExp                     = opTypeInit(120, "战令购买经验")
	OpTypeWarOrderLvReward                   = opTypeInit(121, "战令等级奖励")
	OpTypeWarOrderExchange                   = opTypeInit(122, "战令兑换")
	OpTypeCardActivity                       = opTypeInit(123, " 抽卡")
	OpTypeActivityCardReward                 = opTypeInit(124, "抽卡积分兑换物品")
	OpTypeActivityCardGetReward              = opTypeInit(125, "抽卡获得物品")
	OpTypeCrossShaBakeEndReward              = opTypeInit(126, "跨服沙巴克结算奖励个人")
	OpTypeCrossShaBakeGuildEndReward         = opTypeInit(127, "跨服沙巴克结算奖励门派")
	OpTypeElf                                = opTypeInit(128, "精灵")
	OpTreasureBuyXunLongLin                  = opTypeInit(129, "元宝购买获得寻龙令")
	OpTreasureApplyGetItem                   = opTypeInit(130, "寻龙探宝转盘获得奖励")
	OpTreasureBuyXunLongLinCost              = opTypeInit(131, "元宝购买获得寻龙令消耗")
	OpTreasureApplyGet                       = opTypeInit(132, "寻龙探宝抽奖")
	OpTreasureIntegralAward                  = opTypeInit(133, "寻龙探宝完成多少轮奖励")
	OpTypeCutTreasure                        = opTypeInit(134, "切割")
	OpTypeHolyBeastActive                    = opTypeInit(135, "圣兽激活")
	OpTypeHolyBeastUpStar                    = opTypeInit(136, "圣兽升星")
	OpTypeHolyBeastUpPoint                   = opTypeInit(137, "圣灵点升级")
	OpTypeHolyBeastChooseProp                = opTypeInit(138, "圣兽选择属性")
	OpTypeHolyBeastRest                      = opTypeInit(139, "圣灵点重置消耗")
	OpTypeFitHolyEquipCompose                = opTypeInit(140, "合体圣装合成")
	OpTypeFitHolyEquipDeCompose              = opTypeInit(141, "合体圣装分解")
	OpTypeFitHolyEquipWear                   = opTypeInit(142, "合体圣装穿")
	OpTypeFitHolyEquipRemove                 = opTypeInit(143, "合体圣装卸下")
	OpTypeHolyBeastRestCost                  = opTypeInit(144, "圣灵点重置返回圣灵点")
	OpTypeHolyBeastAddPoint                  = opTypeInit(145, "圣灵点增加")
	OpTypePayToken                           = opTypeInit(146, "充值代币使用")
	OpTypeCompetitveEndReward                = opTypeInit(147, "竞技场结算奖励")
	OpTypeMaterialStageSweep                 = opTypeInit(148, "材料副本扫荡")
	OpTypePersonBossSweep                    = opTypeInit(149, "个人boss扫荡")
	OpTypeVipBossSweep                       = opTypeInit(150, "个人boss扫荡")
	OpTypeInsideAutoUp                       = opTypeInit(151, "内功一键升级")
	OpTypeCompetitveMultipleClaim            = opTypeInit(152, "竞技场多倍领取")
	OpTypeTowerSweep                         = opTypeInit(153, "试炼塔碾压")
	OpTypeDailyRankGetMarkReward             = opTypeInit(154, "每日排名积分奖励")
	OpTypeDailyRankBuyGift                   = opTypeInit(155, "每日排名礼包购买")
	OpTypeChuanShiEquipWear                  = opTypeInit(156, "传世装备穿戴")
	OpTypeChuanShiEquipRemove                = opTypeInit(157, "传世装备卸下")
	OpTypeChuanShiEquipDeCompose             = opTypeInit(158, "传世装备分解")
	OpTypeComposeChuanShiEquip               = opTypeInit(159, "合成传世装备")
	OpTypeExpStageSweep                      = opTypeInit(160, "经验副本扫荡")
	OpTypePreviewFunction                    = opTypeInit(161, "功能预览购买")
	OpTypeDailyTaskResourcesBackAllGetReward = opTypeInit(162, "每日资源找回 一键找回领取奖励")
	OpTypeMaterialStageBuyNum                = opTypeInit(161, "材料副本购买")
	OpTypeExpStageBuyNum                     = opTypeInit(162, "经验副本购买次数")
	OpTypeSevenInvestment                    = opTypeInit(163, "七日投资奖励")
	OpTypeContReceive                        = opTypeInit(164, "连续充值")
	OpTypeEquipDestroy1                      = opTypeInit(165, "装备销毁")
	OpTypeConversionGoldIngot                = opTypeInit(166, "兑换金锭")
	OpTypeAncientBuyNum                      = opTypeInit(167, "远古首领购买次数")
	OpTypeAncientFight                       = opTypeInit(168, "远古首领战斗")
	OpTypeAncientSkillActive                 = opTypeInit(169, "远古神技激活")
	OpTypeAncientSkillUpLv                   = opTypeInit(170, "远古神技升级")
	OpTypeAncientSkillUpGrade                = opTypeInit(171, "远古神技升阶")
	OpTypeTitleActive                        = opTypeInit(172, "称号激活")
	OpTypeMiJiUp                             = opTypeInit(173, "秘籍升級")
	OpTypeAncientTreasureActive              = opTypeInit(174, "远古宝物激活")
	OpTypeAncientTreasureZhuLin              = opTypeInit(175, "远古宝物注灵")
	OpTypeAncientTreasureUpStar              = opTypeInit(176, "远古宝物升星")
	OpTypeAncientTreasureJueXin              = opTypeInit(177, "远古宝物觉醒")
	OpTypeAncientTreasureReset               = opTypeInit(178, "远古宝物重置")
	OpTypeKillMonsterUniFirstDraw            = opTypeInit(179, "首领首杀本服首杀奖励")
	OpTypeKillMonsterUniDraw                 = opTypeInit(180, "首领首杀本服奖励")
	OpTypeKillMonsterPerDraw                 = opTypeInit(181, "首领首杀个人首通奖励")
	OpTypeKillMonsterMilDraw                 = opTypeInit(182, "首领首杀里程碑奖励")
	OpTypeTreasureShopBuy                    = opTypeInit(183, "多宝阁购买")
	OpTypeTreasureShopRefresh                = opTypeInit(184, "多宝阁刷新")
	OpTypeChuanShiStrengthen                 = opTypeInit(185, "传世装备强化")
	OpTypePetAppendage                       = opTypeInit(186, "战宠附体")
	OpTypeGodEquipBlood                      = opTypeInit(187, "神兵血炼")
	OpTypeHellBossBuyNum                     = opTypeInit(188, "炼狱首领购买次数")
	OpTypeLotteryBuyNum                      = opTypeInit(189, "摇彩购买份额")
	OpTypeGetGoodLuckNum                     = opTypeInit(190, "摇彩接好运")
	OpTypeGetGetEndAward                     = opTypeInit(191, "摇彩主动领取奖励")
	OpTypeGetTrialTaskAward                  = opTypeInit(192, "试炼之路任务奖励")
	OpTypeGetTrialTaskStageAward             = opTypeInit(193, "试炼之路阶段奖励")
	OpTypeTowerRankReward                    = opTypeInit(194, "试炼塔排行奖励")
	OpTypeFieldBossFirstReceive              = opTypeInit(195, "野外首领首次奖励")
	OpTypeMagicLayerAward                    = opTypeInit(196, "九层魔塔层奖励")
	OpTypeMagicLayerTimeAward                = opTypeInit(197, "九层魔塔层计时奖励")
	OpTypeDaBaoEquipUp                       = opTypeInit(198, "打宝神器升级")
	OpTypeDaBaoMysteryEnergyItemBuy          = opTypeInit(199, "购买打宝秘境体力道具")
	OpTypeDaBaoMysteryEnergyAdd              = opTypeInit(200, "打宝秘境体力道具使用")
	OpTypeAppletsReceive                     = opTypeInit(201, "领取魔法射击杀怪奖励")
	OpTypeCronGetAwardReq                    = opTypeInit(202, "魔法射击定时奖励获取")
	OpTypeEndResult                          = opTypeInit(203, "小游戏通关奖励")
	OpTypeFirstRechargeDiscount              = opTypeInit(204, "首充优惠券使用")
	OpTypeLabelUp                            = opTypeInit(205, "头衔升级")
	OpTypeLabelDayReward                     = opTypeInit(206, "头衔每日奖励")
	OpTypFirstDropReward                     = opTypeInit(207, "首暴奖励领取")
	OpTypAllFirstDropReward                  = opTypeInit(208, "首暴奖励一键领取")
	OpTypePrivilegeBuy                       = opTypeInit(209, "特权购买")
	OpTypePrivilegeItemActive                = opTypeInit(210, "特权道具激活")

	//1000+战斗用
	OpTypeFightCheer  = opTypeInit(1000, "战斗鼓舞")
	OpTypeFightPotion = opTypeInit(1001, "战斗吃药")
	OpTypeCollection  = opTypeInit(1002, "采集")
	OpTypePickUp      = opTypeInit(1003, "拾取")
	OpTypeHelp        = opTypeInit(1004, "协助")

	//10000+为邮件

	/***************************************************************/
	/******************   1000+战斗用           *********************/
	/******************   10000+为邮件          *********************/
	/******************   普通来源不要超过10000  *********************/
	/***************************************************************/
)

var OPTYPE_SOURCE_MAP = make(map[int]string)

func opTypeInit(typeId int, typeMsg string) int {
	OPTYPE_SOURCE_MAP[typeId] = typeMsg
	return typeId
}
