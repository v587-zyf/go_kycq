package condition

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"strconv"
	"time"
)

type ConditionManager struct {
	util.DefaultModule
	managersI.IModule
	checkers map[int]func(user *objs.User, heroIndex, value int) (int, bool)
	//appointCheckers map[int]func(user *objs.User, heroIndex, appoint, value int) (int, bool)
	tooCheckers map[int]func(user *objs.User, heroIndex int, value []int) ([]int, bool)
}

func NewConditionManager(module managersI.IModule) *ConditionManager {
	return &ConditionManager{IModule: module}
}

func (this *ConditionManager) Init() error {
	this.checkers = make(map[int]func(user *objs.User, heroIndex, value int) (int, bool))
	this.tooCheckers = make(map[int]func(user *objs.User, heroIndex int, value []int) ([]int, bool))
	this.initCheckers()
	return nil
}

func (this *ConditionManager) Check(user *objs.User, heroIndex, condition int, conditionValue int) (int, bool) {
	if fn, ok := this.checkers[condition]; ok {
		return fn(user, heroIndex, conditionValue)
	}
	logger.Error("condition 检查：id %d, 未实现", condition)
	return 0, false
}

func (this *ConditionManager) CheckBySlice(user *objs.User, heroIndex int, condition []int) ([]int, bool) {
	if len(condition) < 2 {
		return []int{0}, false
	}
	conditionName := condition[0]
	conditionValue := condition[1:]

	if fn, ok := this.checkers[conditionName]; ok {
		r, b := fn(user, heroIndex, conditionValue[0])
		return []int{r}, b
	}
	if fn, ok := this.tooCheckers[conditionName]; ok {
		return fn(user, heroIndex, conditionValue)
	}
	logger.Error("condition 检查：id %d, 未实现", condition)
	return []int{0}, false
}
func (this *ConditionManager) CheckMultiBySlice2(user *objs.User, heroIndex int, conditions [][]int) bool {
	for _, conditionSlice := range conditions {
		if _, b := this.CheckBySlice(user, heroIndex, conditionSlice); !b {
			return false
		}
	}
	return true
}
func (this *ConditionManager) CheckMultiByMap(user *objs.User, heroIndex int, conditions map[int]int) bool {
	slices := make([][]int, 0)
	for c, v := range conditions {
		slices = append(slices, []int{c, v})
	}
	return this.CheckMultiBySlice2(user, heroIndex, slices)
}

func (this *ConditionManager) CheckConditionType(id int) bool {
	_, ok := this.checkers[id]
	return ok
}

func (this *ConditionManager) CheckMulti(user *objs.User, heroIndex int, conditions map[int]int) bool {

	return this.CheckMultiByType(user, heroIndex, conditions, pb.CONDITIONTYPE_ALL)
}

func (this *ConditionManager) CheckMultiByType(user *objs.User, heroIndex int, conditions map[int]int, conditionType int) bool {
	for condition, conditionValue := range conditions {
		if fn, ok := this.checkers[condition]; ok {
			_, isPass := fn(user, heroIndex, conditionValue)
			if conditionType == pb.CONDITIONTYPE_ALL && !isPass {
				return false
			} else if conditionType == pb.CONDITIONTYPE_JUST_ONE && isPass {
				return true
			}
		} else {
			logger.Error("condition 检查：id %d, 未实现", condition)
			return false
		}
	}
	if conditionType == pb.CONDITIONTYPE_ALL {
		return true
	} else {
		return false
	}
}

func (this *ConditionManager) initCheckers() {
	//等级
	this.checkers[pb.CONDITION_USER_LEVEL] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return this.GetExpPool().GetHeroMaxLv(user), this.GetExpPool().GetHeroMaxLv(user) >= value
	}
	//转生等级
	this.checkers[pb.CONDITION_USER_REBIRTH_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		reinLvl := 0
		if reinConf := gamedb.GetReinCfg(user.Rein.Id); reinConf != nil {
			reinLvl = reinConf.Level
		}
		return reinLvl, reinLvl >= value
	}
	//战力
	this.checkers[pb.CONDITION_USER_COMBAT] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return user.Combat, user.Combat >= value
	}
	//通关关卡
	this.checkers[pb.CONDITION_USER_STAGE_PASS] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return user.StageId, user.StageId > value
	}
	//vip等级
	this.checkers[pb.CONDITION_USER_VIP_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return user.VipLevel, user.VipLevel >= value
	}
	//月卡
	this.checkers[pb.CONDITION_ACTIVE_MONTH_CARD_SLIVER] = func(user *objs.User, heroIndex, value int) (int, bool) {
		flag := true
		if info, ok := user.MonthCard.MonthCards[pb.MONTHCARDTYPE_SLIVER]; !ok || info.EndTime < int(time.Now().Unix()) {
			flag = false
		}
		return 0, flag
	}
	this.checkers[pb.CONDITION_ACTIVE_MONTH_CARD_GOLD] = func(user *objs.User, heroIndex, value int) (int, bool) {
		flag := true
		if info, ok := user.MonthCard.MonthCards[pb.MONTHCARDTYPE_GOLD]; !ok || info.EndTime < int(time.Now().Unix()) {
			flag = false
		}
		return 0, flag
	}
	//开服天数
	this.checkers[pb.CONDITION_SERVER_OPEN_TIME] = func(user *objs.User, heroIndex, value int) (int, bool) {
		openDay := this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
		return openDay, openDay >= value
	}
	//性别
	this.checkers[pb.CONDITION_USER_SEX] = func(user *objs.User, heroIndex, value int) (int, bool) {
		hero := user.Heros[heroIndex]
		flag := hero.Sex == value
		if value == 0 {
			flag = true
		}
		return hero.Sex, flag
	}
	//职业
	this.checkers[pb.CONDITION_JOB] = func(user *objs.User, heroIndex, value int) (int, bool) {
		hero := user.Heros[heroIndex]
		return hero.Job, hero.Job == value
	}
	this.checkers[pb.CONDITION_COST_CHUAN_QI_BI] = func(user *objs.User, heroIndex, value int) (int, bool) {
		enouth, num := this.GetBag().HasEnough(user, pb.ITEMID_CHUANQI_BI, value)
		return num, enouth
	}
	this.checkers[pb.CONDITION_COST_INGOT] = func(user *objs.User, heroIndex, value int) (int, bool) {
		enouth, num := this.GetBag().HasEnough(user, pb.ITEMID_INGOT, value)
		return num, enouth
	}
	this.checkers[pb.CONDITION_COST_GOLD] = func(user *objs.User, heroIndex, value int) (int, bool) {
		enouth, num := this.GetBag().HasEnough(user, pb.ITEMID_GOLD, value)
		return num, enouth
	}
	//单角色全身强化
	this.checkers[pb.CONDITION_ONE_STRENGTHEN_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		for _, hero := range user.Heros {
			flag := true
			for _, lv := range hero.EquipsStrength {
				if lv < value {
					flag = false
					break
				}
			}
			if flag {
				return value, true
			}
		}
		return value, false
	}
	//2角色全身强化等级
	this.checkers[pb.CONDITION_TWO_STRENGTHEN_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := 0
		for _, hero := range user.Heros {
			flag := true
			for _, lv := range hero.EquipsStrength {
				if lv < value {
					flag = false
					break
				}
			}
			if flag {
				num++
			}
		}
		return value, num >= 2
	}
	//三角色全身强化
	this.checkers[pb.CONDITION_THREE_STRENGTHEN_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		flag := true
		for _, hero := range user.Heros {
			for _, lv := range hero.EquipsStrength {
				if lv < value {
					flag = false
					break
				}
			}
		}
		return value, flag
	}
	//战宠等级
	this.checkers[pb.CONDITION_PET_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		for _, pet := range user.Pet {
			if pet.Lv >= value {
				return pet.Lv, true
			}
		}
		return value, false
	}
	//一个角色达到等级的
	this.checkers[pb.CONDITION_ONE_HERO_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		flag := false
		maxLv := 0
		for _, hero := range user.Heros {
			if hero.ExpLvl >= value {
				flag = true
				return value, flag
			}
			if hero.ExpLvl >= maxLv {
				maxLv = hero.ExpLvl
			}
		}
		return maxLv, flag
	}
	//2个角色达到等级的
	this.checkers[pb.CONDITION_TWO_HERO_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := 0
		maxLv := 0
		for _, hero := range user.Heros {
			if hero.ExpLvl >= value {
				num++
			}
			if hero.ExpLvl >= maxLv {
				maxLv = hero.ExpLvl
			}
		}
		return maxLv, num >= 2
	}
	//3个角色达到等级的
	this.checkers[pb.CONDITION_THREE_HERO_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := 0
		maxLv := 0
		for _, hero := range user.Heros {
			if hero.ExpLvl >= value {
				num++
			}
			if hero.ExpLvl >= maxLv {
				maxLv = hero.ExpLvl
			}
		}
		return maxLv, num >= 3
	}
	//当前角色等级
	this.checkers[pb.CONDITION_NOW_HERO_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		if hero, ok := user.Heros[heroIndex]; ok && hero.ExpLvl >= value {
			return hero.ExpLvl, true
		}
		return 0, false
	}
	//通天塔挑战层数
	this.checkers[pb.CONDITION_TOWER] = func(user *objs.User, heroIndex, value int) (int, bool) {
		if user.Tower.TowerLv-1 >= value {
			return value, true
		}
		return user.Tower.TowerLv - 1, false
	}
	//内功阶数
	this.checkers[pb.CONDITION_INSIDE_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		heroInside := user.Heros[heroIndex].Inside
		insideCfg := gamedb.GetInsideArtInsideArtCfg(heroInside.Acupoint[pb.INSIDETYPE_ONE])
		if insideCfg.Grade >= value {
			return value, true
		}
		return insideCfg.Grade, false
	}
	//合体等级
	this.checkers[pb.CONDITION_FIT_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		id := value / constConstant.COMPUTE_TEN_THOUSAND
		lv := value % constConstant.COMPUTE_TEN_THOUSAND
		maxLv := 0
		if fitLv, ok := user.Fit.Lv[id]; ok {
			if fitLv >= maxLv {
				maxLv = fitLv
			}
			if fitLv >= lv {
				return value, true
			}
		}
		return maxLv, false
	}
	//累计消费
	this.checkers[pb.CONDITION_SPEND_INGOT] = func(user *objs.User, heroIndex, value int) (int, bool) {
		i := user.SpendRebates.Ingot
		if i >= value {
			return value, true
		}
		return i, false
	}
	//穿戴任意一件装备
	this.checkers[pb.CONDITION_WEAR_EQUIP] = func(user *objs.User, heroIndex, value int) (int, bool) {
		for _, v := range user.Heros {
			for _, v1 := range v.Equips {
				if v1.ItemId > 0 {
					return value, true
				}
			}
		}
		return 0, false
	}
	//提升神翼至2阶
	this.checkers[pb.CONDITION_UPGRADE_SHEN_YI_2] = func(user *objs.User, heroIndex, value int) (int, bool) {
		maxJie := 0
		for _, v := range user.Heros {
			if v.Wings != nil && len(v.Wings) >= 1 {
				ids := v.Wings[0].Id
				order := gamedb.GetWingNewWingNewCfg(ids).Order
				if order > maxJie {
					maxJie = order
				}
				if order >= value {
					return value, true
				}
			}
		}
		return maxJie, false
	}
	//解锁二角色
	this.checkers[pb.CONDITION_UNLOCK_SECOND_HERO] = func(user *objs.User, heroIndex, value int) (int, bool) {
		if len(user.Heros) >= 2 {
			return value, true
		}
		return 0, false
	}
	//解锁三角色
	this.checkers[pb.CONDITION_UNLOCK_THREE_HERO] = func(user *objs.User, heroIndex, value int) (int, bool) {
		if len(user.Heros) >= 3 {
			return value, true
		}
		return 0, false
	}
	//任意激活xx个图鉴
	this.checkers[pb.CONDITION_ACTIVATE_TU_JIAN] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := len(user.Atlases)
		return num, num >= value
	}
	//挑战一次经验副本
	this.checkers[pb.CONDITION_CHALLENGE_JIN_YAN_FU_BEN] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetConditionData(user, pb.CONDITION_CHALLENGE_JIN_YAN_FU_BEN, 0)
		return num, num >= value
	}
	//提升vip等级 将vip等级提升至1级
	this.checkers[pb.CONDITION_UPGRADE_VIP_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		lv := user.VipLevel
		return lv, lv >= value
	}
	//合成一颗灵丹
	this.checkers[pb.CONDITION_HE_CHENG_LIN_DAN] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetConditionData(user, pb.CONDITION_HE_CHENG_LIN_DAN, 0)
		return num, num >= value
	}
	//使用灵丹
	this.checkers[pb.CONDITION_USE_LIN_DAN] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := 0
		for _, info := range user.Panaceas {
			num += info.Number
		}
		return num, num >= value
	}
	//镶嵌一颗宝石
	this.checkers[pb.CONDITION_XIANG_QIAN_BAO_SHI] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetConditionData(user, pb.CONDITION_XIANG_QIAN_BAO_SHI, 0)
		return num, num >= value
	}
	//全身宝石达到value级
	this.checkers[pb.CONDITION_XIANG_QIAN_BAO_SHI_1] = func(user *objs.User, heroIndex, value int) (int, bool) {
		allLv := 0
		for _, hero := range user.Heros {
			if hero.Jewel != nil {
				for _, v := range hero.Jewel {
					if v.One != 0 {
						oneCfg := gamedb.GetJewelJewelCfg(v.One)
						allLv += oneCfg.Level
					}
					if v.Two != 0 {
						twoCfg := gamedb.GetJewelJewelCfg(v.Two)
						allLv += twoCfg.Level
					}
					if v.Three != 0 {
						threeCfg := gamedb.GetJewelJewelCfg(v.Three)
						allLv += threeCfg.Level
					}
				}
			}
		}
		return allLv, allLv >= value
	}
	//挑战1次竞技场
	this.checkers[pb.CONDITION_CHALLENGE_JIN_JI_CHANG] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetConditionData(user, pb.CONDITION_CHALLENGE_JIN_JI_CHANG, 0)
		return num, num >= value
	}
	//领取一次日常任务奖励
	this.checkers[pb.CONDITION_GET_DAILY_TASK] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetConditionData(user, pb.CONDITION_GET_DAILY_TASK, 0)
		return num, num >= value
	}
	//内功重数达到2阶
	this.checkers[pb.CONDITION_XIU_LIAN_NEI_GONG_1] = func(user *objs.User, heroIndex, value int) (int, bool) {
		allLv := 0
		for _, hero := range user.Heros {
			inside := hero.Inside
			insideCfg := gamedb.GetInsideArtInsideArtCfg(inside.Acupoint[pb.INSIDETYPE_ONE])
			if insideCfg != nil {
				allLv += insideCfg.Grade
			}
		}
		return allLv, allLv >= value
	}
	//提升所有神兵总等级到x阶
	this.checkers[pb.CONDITION_UPGRADE_SHEN_BIN_1] = func(user *objs.User, heroIndex, value int) (int, bool) {
		allLv := 0
		for _, hero := range user.Heros {
			for _, godEquip := range hero.GodEquips {
				if godEquip != nil {
					allLv += godEquip.Lv
				}
			}
		}
		return allLv, allLv >= value
	}
	//前往挖矿一次
	this.checkers[pb.CONDITION_GO_TO_WA_KUANG] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetConditionData(user, pb.CONDITION_GO_TO_WA_KUANG, 0)
		return num, num >= value
	}
	//激活x只战宠
	this.checkers[pb.CONDITION_ACTIVATE_ZHAN_CHONG] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := len(user.Pet)
		return num, num >= value
	}
	//提升任意一只战宠到2阶
	this.checkers[pb.CONDITION_UPGRADE_ZHAN_CHONG_1] = func(user *objs.User, heroIndex, value int) (int, bool) {
		maxLv := 0
		for _, v := range user.Pet {
			if v.Grade >= maxLv {
				maxLv = v.Grade
			}
			if v.Grade >= value {
				return value, true
			}
		}
		return maxLv, false
	}

	//战宠总等级达到&级
	this.checkers[pb.CONDITION_ZHAN_CHONG_ALL_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		allLv := 0
		for _, v := range user.Pet {
			allLv += v.Lv
		}
		return allLv, allLv >= value
	}

	//将官衔提升至官职1
	this.checkers[pb.CONDITION_UPGRADE_GUAN_XIAN] = func(user *objs.User, heroIndex, value int) (int, bool) {
		lv := user.Official
		return lv, lv >= value
	}
	//领取一次成就奖励
	this.checkers[pb.CONDITION_GET_ONE_TIME_CHENG_JIU_AWARD] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetConditionData(user, pb.CONDITION_GET_ONE_TIME_CHENG_JIU_AWARD, 0)
		return num, num >= value
	}
	//挑战1次神翼副本
	this.checkers[pb.CONDITION_TIAO_ZHAN_SHEN_YI_FU_BEN] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetConditionData(user, pb.CONDITION_TIAO_ZHAN_SHEN_YI_FU_BEN, 0)
		return num, num >= value
	}
	//升级神翼
	this.checkers[pb.CONDITION_UPGRADE_SHEN_YI_1] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetConditionData(user, pb.CONDITION_UPGRADE_SHEN_YI_1, 0)
		return num, num >= value
	}
	//通过第一关卡
	this.checkers[pb.CONDITION_CHALLENGE_STAGE] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := user.StageId
		return num, num >= value
	}
	//累计充值奖励领取
	this.checkers[pb.CONDITION_RECHARGE_ALL_CHECK] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := user.RechargeAll / 100
		return num, num >= value
	}
	//打完第几个首领
	this.checkers[pb.CONDITION_KILL_SHOU_LIN] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := user.StageId2
		return num, num >= value
	}
	//主宰装备总等级
	this.checkers[pb.CONDITION_UPGRADE_EQUIP] = func(user *objs.User, heroIndex, value int) (int, bool) {
		allLv := 0
		for _, hero := range user.Heros {
			for _, lv := range hero.Dictates {
				if lv > 0 {
					allLv += lv
				}
			}
		}
		return allLv, allLv >= value
	}
	//装备xx个图鉴
	this.checkers[pb.CONDITION_WEAR_TU_JIAN] = func(user *objs.User, heroIndex, value int) (int, bool) {
		maxNum := 0
		for _, hero := range user.Heros {
			atlasWears := hero.Wear.AtlasWear
			num := 0
			for _, v := range atlasWears {
				if v > 0 {
					num++
				}
				if num > maxNum {
					maxNum = num
				}
				if num >= value {
					return value, true
				}
			}
		}
		return maxNum, false
	}
	//三个角色强化总等级之和，达到指定级数时完成任务
	this.checkers[pb.CONDITION_ALL_HEROS_QIANG_HUA_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		allLv := 0
		for _, hero := range user.Heros {
			for _, lv := range hero.EquipsStrength {
				allLv += lv
			}
		}
		return allLv, allLv >= value
	}
	//完成&个主线任务
	this.checkers[pb.CONDITION_TASK_NUMS] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return user.MainLineTask.TaskId, user.MainLineTask.TaskId >= value
	}
	//至尊法器阶数
	this.checkers[pb.CONDITION_HOLYARMS_GRADE] = func(user *objs.User, heroIndex, value int) (int, bool) {
		allLv := 0
		for _, holyarm := range user.Holyarms {
			allLv += holyarm.Level
		}
		return allLv, allLv >= value
	}
	//玩家图鉴总星级达到&颗
	this.checkers[pb.CONDITION_ATLAS_STAR] = func(user *objs.User, heroIndex, value int) (int, bool) {
		allStar := 0
		for _, star := range user.Atlases {
			allStar += star
		}
		for _, star := range user.AtlasGathers {
			allStar += star
		}
		return allStar, allStar >= value
	}
	//小精灵达到&级
	this.checkers[pb.CONDITION_ELF_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return user.Elf.Lv, user.Elf.Lv >= value
	}
	//挑战1次强化副本
	this.checkers[pb.CONDITION_CHALLENGE_QIANG_HUA_FU_BEN] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetCondition().GetConditionData(user, pb.CONDITION_CHALLENGE_QIANG_HUA_FU_BEN, 0)
		return num, num >= value
	}
	//小精灵升至&级
	this.checkers[pb.CONDITION_UPGRADE_JIN_LIN_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return user.Elf.Lv, user.Elf.Lv >= value
	}

	//小精灵升至&级
	this.checkers[pb.CONDITION_XIAO_JIN_LIN_UP_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return user.Elf.Lv - 1, user.Elf.Lv-1 >= value
	}
	//圣兽升至&级
	this.checkers[pb.CONDITION_UPGRADE_SHEN_SHOU] = func(user *objs.User, heroIndex, value int) (int, bool) {
		taskConf := gamedb.GetTaskConditionCfg(user.MainLineTask.TaskId)
		if taskConf == nil || taskConf.ConditionType != pb.CONDITION_UPGRADE_SHEN_SHOU {
			maxStar := 0
			for _, v := range user.Heros {
				for _, data := range v.HolyBeastInfos {
					if data.Star > maxStar {
						maxStar = data.Star
					}
				}
			}
			return maxStar, maxStar >= value
		}
		markStar := 0
		conditionValue := taskConf.ConditionValue
		if conditionValue != nil && len(conditionValue) >= 2 {
			for _, v := range user.Heros {
				for types, data := range v.HolyBeastInfos {
					if types == conditionValue[0] {
						if data.Star > markStar {
							markStar = data.Star
						}
					}

				}
			}
		}
		return markStar, markStar >= value
	}
	//升级神刀技
	this.checkers[pb.CONDITION_UPGRADE_SHENG_DAO_SKILL] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return user.CutTreasure, user.CutTreasure >= value
	}
	//两角色主宰平均&阶
	this.checkers[pb.CONDITION_TWO_DICTATE_GRADE] = func(user *objs.User, heroIndex, value int) (int, bool) {
		okNum := 0
		for _, hero := range user.Heros {
			flag := true
			for _, grade := range hero.Dictates {
				if grade < value {
					flag = false
					break
				}
			}
			if flag {
				okNum++
			}
		}
		return okNum, okNum >= 2
	}
	//三角色主宰平均&阶
	this.checkers[pb.CONDITION_THREE_DICTATE_GRADE] = func(user *objs.User, heroIndex, value int) (int, bool) {
		g := 0
		for _, hero := range user.Heros {
			for _, grade := range hero.Dictates {
				if grade < value {
					return grade, false
				}
				if g < grade {
					g = grade
				}
			}
		}
		return g, true
	}
	//三角色神翼总阶数达到&阶
	this.checkers[pb.CONDITION_THREE_WING_GRADE] = func(user *objs.User, heroIndex, value int) (int, bool) {
		allLv := 0
		for _, hero := range user.Heros {
			allLv += gamedb.GetWingNewWingNewCfg(hero.Wings[0].Id).Order
		}
		return allLv, allLv >= value
	}
	//两角色内功达到&阶
	this.checkers[pb.CONDITION_TWO_INSIDE_GRADE] = func(user *objs.User, heroIndex, value int) (int, bool) {
		okNum := 0
		for _, hero := range user.Heros {
			grade := gamedb.GetInsideArtInsideArtCfg(hero.Inside.Acupoint[pb.INSIDETYPE_ONE]).Grade
			if grade >= value {
				okNum++
			}
		}
		return okNum, okNum >= 2
	}
	//三角色内功达到&阶
	this.checkers[pb.CONDITION_THREE_INSIDE_GRADE] = func(user *objs.User, heroIndex, value int) (int, bool) {
		okNum := 0
		for _, hero := range user.Heros {
			grade := gamedb.GetInsideArtInsideArtCfg(hero.Inside.Acupoint[pb.INSIDETYPE_ONE]).Grade
			if grade >= value {
				okNum++
			}
		}
		return okNum, okNum >= 3
	}

	//三角色内功总重数
	this.checkers[pb.CONDITION_THREE_INSIDE_CHONG] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetCondition().GetConditionData(user, pb.CONDITION_THREE_INSIDE_CHONG, 0)
		return num, num >= value
	}

	//当前角色强化等级
	this.checkers[pb.CONDITION_STRENGTHEN_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		hero := user.Heros[heroIndex]
		for _, lv := range hero.EquipsStrength {
			if lv < value {
				return lv, false
			}
		}
		return value, true
	}
	//任意两角神翼达到&阶&星
	this.checkers[pb.CONDITION_TWO_WING] = func(user *objs.User, heroIndex, value int) (int, bool) {
		heros := user.Heros
		okNum := 0
		for _, hero := range heros {
			if hero.Wings[0].Id >= value {
				okNum++
			}
		}
		return value, okNum >= 2
	}
	//三角神翼达到&阶&星
	this.checkers[pb.CONDITION_THREE_WING] = func(user *objs.User, heroIndex, value int) (int, bool) {
		heros := user.Heros
		for _, hero := range heros {
			if hero.Wings[0].Id < value {
				return value, false
			}
		}
		return value, true
	}
	//每日任务(实时)活跃度达到xx
	this.checkers[pb.CONDITION_DAILY_TASK_LIVENESS] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return user.DailyTask.DayExp, user.DailyTask.DayExp >= value
	}
	//挑战x次野外首领  进入地图就算
	this.checkers[pb.CONDITION_CHALLENGE_FIELD_LEADER] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetCondition().GetConditionData(user, pb.CONDITION_CHALLENGE_FIELD_LEADER, 0)
		return num, num >= value
	}
	//解锁&个角色
	this.checkers[pb.CONDITION_UNLOCK_HERO] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := len(user.Heros)
		return conditionData, conditionData > value
	}
	//全角色共穿戴&件装备
	this.checkers[pb.CONDITION_ALL_HERO_WEAR_EQUIP] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := 0
		for _, hero := range user.Heros {
			for _, equip := range hero.Equips {
				if equip.ItemId != 0 {
					conditionData++
				}
			}
		}
		return conditionData, conditionData >= value
	}
	//全角色共解锁&件时装
	this.checkers[pb.CONDITION_ALL_HERO_UNLOCK_FASHION] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := 0
		for _, hero := range user.Heros {
			for _, fashion := range hero.Fashions {
				if fashion.Id != 0 {
					conditionData++
				}
			}
		}
		return conditionData, conditionData >= value
	}
	//全角色共穿戴&件传世装备
	this.checkers[pb.CONDITION_ALL_HERO_WEAR_CHUANSHI_EQUIP] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := 0
		for _, hero := range user.Heros {
			for _, itemId := range hero.ChuanShi {
				if itemId != 0 {
					conditionData++
				}
			}
		}
		return conditionData, conditionData >= value
	}
	//全角色共穿戴&件主宰装备
	this.checkers[pb.CONDITION_ALL_HERO_WEAR_DICTATE] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := 0
		for _, hero := range user.Heros {
			for _, itemId := range hero.Dictates {
				if itemId != 0 {
					conditionData++
				}
			}
		}
		return conditionData, conditionData >= value
	}
	//全角色穿戴&件合体圣装
	this.checkers[pb.CONDITION_ALL_HERO_WEAR_FIT_HOLY_EQUIP] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := 0
		for _, equips := range user.FitHolyEquip.Equips {
			for _, itemId := range equips {
				if itemId != 0 {
					conditionData++
				}
			}
		}
		return conditionData, conditionData >= value
	}
	//激活&个合体技能
	this.checkers[pb.CONDITION_ACTIVE_FIT_SKILL] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := len(user.Fit.Skills)
		return conditionData, conditionData >= value
	}
	//全角色龙器累计升到&阶
	this.checkers[pb.CONDITION_ALL_HERO_DRAGON_EQUIP_GRADE] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := 0
		for _, hero := range user.Heros {
			for _, lv := range hero.DragonEquip {
				if lv != 0 {
					conditionData += lv
				}
			}
		}
		return conditionData, conditionData >= value
	}
	//全角色共佩戴&个特戒
	this.checkers[pb.CONDITION_ALL_HERO_WEAR_RING] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := 0
		for _, hero := range user.Heros {
			for _, ring := range hero.Ring {
				if ring.Rid != 0 {
					conditionData++
				}
			}
		}
		return conditionData, conditionData >= value
	}
	//特戒累计强化到&级
	this.checkers[pb.CONDITION_RING_STRENGTHEN_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := 0
		for _, hero := range user.Heros {
			for _, ring := range hero.Ring {
				conditionData += ring.Strengthen
			}
		}
		return conditionData, conditionData >= value
	}

	//特戒累计强化到xx级
	this.checkers[pb.CONDITION_TE_JIE_QIANG_HUA_TIMES] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetCondition().GetConditionData(user, pb.CONDITION_TE_JIE_QIANG_HUA_TIMES, 0)
		return num, num >= value
	}

	//沙巴克积分排名第&的公会会长
	this.checkers[pb.CONDITION_SHABAKE_GUILD_RANK] = func(user *objs.User, heroIndex, value int) (int, bool) {
		rankData := this.GetRank().GetRankByScore(pb.RANKTYPE_SHABAKE_GUILD, value, value)
		if len(rankData) > 0 && rankData[0] != 0 {
			if user.GuildData.MyCreateId == rankData[0] {
				return rankData[0], true
			}
		}
		return 0, false
	}
	//法宝累计升到&阶
	this.checkers[pb.CONDITION_HOLYARMS_UP_GRADE] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := 0
		for _, info := range user.Holyarms {
			conditionData += info.Level
		}
		return conditionData, conditionData >= value
	}
	//绝学累计升到&级
	this.checkers[pb.CONDITION_JUEXUE_UP_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := 0
		for _, info := range user.Juexues {
			conditionData += info.Lv
		}
		return conditionData, conditionData >= value
	}
	//激活&只战宠
	this.checkers[pb.CONDITION_ACTIVE_PET] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := 0
		for _, info := range user.Pet {
			if info.Lv > 0 {
				conditionData++
			}
		}
		return conditionData, conditionData >= value
	}
	//&战宠升到&级
	this.tooCheckers[pb.CONDITION_PET_UP_LV] = func(user *objs.User, heroIndex int, value []int) ([]int, bool) {
		petId, lvId := value[0], value[1]
		lv := 0
		if cfg := gamedb.GetPetsLevelConfCfg(lvId); cfg != nil {
			lv = cfg.Id % constConstant.COMPUTE_TEN_THOUSAND
			if userPet, ok := user.Pet[petId]; ok && userPet.Lv >= lv {
				return []int{petId, userPet.Lv}, true
			}
		}
		return []int{petId, lv}, false
	}
	//&战宠升到&阶
	this.tooCheckers[pb.CONDITION_PET_UP_GRADE] = func(user *objs.User, heroIndex int, value []int) ([]int, bool) {
		petId, gradeId := value[0], value[1]
		grade := 0
		if cfg := gamedb.GetPetsGradeConfCfg(gradeId); cfg != nil {
			grade = cfg.Id % constConstant.COMPUTE_TEN_THOUSAND
			if userPet, ok := user.Pet[petId]; ok && userPet.Grade >= grade {
				return []int{petId, userPet.Grade}, true
			}
		}
		return []int{petId, grade}, false
	}
	//全角色领域累计升到&级
	this.checkers[pb.CONDITION_ALL_HERO_AREA_UP_LV] = func(user *objs.User, heroIndex int, value int) (int, bool) {
		conditionData := 0
		for _, hero := range user.Heros {
			for _, lv := range hero.Area {
				conditionData += lv
			}
		}
		return conditionData, conditionData >= value
	}
	//全角色法阵累计升到&星
	this.checkers[pb.CONDITION_ALL_HERO_MAGIC_CIRCLE_UP_STAR] = func(user *objs.User, heroIndex int, value int) (int, bool) {
		conditionData := this.GetConditionData(user, pb.CONDITION_ALL_HERO_MAGIC_CIRCLE_UP_STAR, 0)
		return conditionData, conditionData >= value
	}
	//全角色圣兽累计升到&星
	this.checkers[pb.CONDITION_ALL_HERO_HOLY_BEAST_UP_STAR] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := 0
		for _, hero := range user.Heros {
			for _, info := range hero.HolyBeastInfos {
				conditionData += info.Star
			}
		}
		return conditionData, conditionData >= value
	}
	//累计在线&分钟
	this.checkers[pb.CONDITION_ALL_ONLINE] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := this.GetConditionData(user, pb.CONDITION_ALL_ONLINE, 0)
		conditionData = common.FloorFloat64(float64(conditionData) / 60)
		return conditionData, conditionData >= value
	}
	//首次充值
	this.checkers[pb.CONDITION_FIRST_RECHARGE] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return 0, user.FirstRecharge.IsRecharge
	}
	//购买七日投资
	this.checkers[pb.CONDITION_BUY_SEVEN_INVESTMENT] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return 0, user.SevenInvestment.BuyOpenDay != 0
	}
	//加入一个公会
	this.checkers[pb.CONDITION_ADD_GUILD] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return 0, user.GuildData.NowGuildId != 0
	}
	//成为公会会长
	this.checkers[pb.CONDITION_BECOME_GUILD_PRESIDENT] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return 0, user.GuildData.MyCreateId != 0
	}
	//成为公会副会长
	this.checkers[pb.CONDITION_BECOME_GUILD_VICE_PRESIDENT] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return 0, user.GuildData.Position == pb.GUILDPOSITION_FUHUIZHANG
	}
	//成为公会长老
	this.checkers[pb.CONDITION_BECOME_GUILD_ELDERS] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return 0, user.GuildData.Position == pb.GUILDPOSITION_ZHANGLAO
	}
	//全角色技能累计升到&级
	this.checkers[pb.CONDITION_ALL_SKILL_UP_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := 0
		for _, hero := range user.Heros {
			for _, info := range hero.Skills {
				conditionData += info.Lv
			}
		}
		return conditionData, conditionData >= value
	}
	//连续签到天数
	this.checkers[pb.CONDITION_CONTINUOUS_SIGN] = func(user *objs.User, heroIndex, value int) (int, bool) {
		data := user.Sign.ContinuitySign
		return data, data >= value
	}
	//竞技场连胜
	this.checkers[pb.CONDITION_COMPETIVE_WIN] = func(user *objs.User, heroIndex, value int) (int, bool) {
		data := user.CompetitiveInfo.ContinuityWin
		return data, data >= value
	}

	//竞技场all胜
	this.checkers[pb.CONDITION_COMPETIVE_ALL_WIN] = func(user *objs.User, heroIndex, value int) (int, bool) {
		data := user.SeasonWinTimes
		return data, data >= value
	}

	//全身宝石达到value级
	this.checkers[pb.CONDITION_XIANG_QIAN_BAO_SHI_MAX_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return user.ModuleUpMax.BaoSiLv, user.ModuleUpMax.BaoSiLv >= value
	}
	//战宠突破
	this.tooCheckers[pb.CONDITION_PET_UP_BREAK] = func(user *objs.User, heroIndex int, value []int) ([]int, bool) {
		petId, breakId := value[0], value[1]
		b := 0
		if cfg := gamedb.GetPetsBreakConfCfg(breakId); cfg != nil {
			b = cfg.Id % constConstant.COMPUTE_TEN_THOUSAND
			if userPet, ok := user.Pet[petId]; ok && userPet.Break >= b {
				return []int{petId, userPet.Grade}, true
			}
		}
		return []int{petId, b}, false
	}

	//累计充值天数
	this.checkers[pb.CONDITION_ALL_RECHARGE_DAY] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := this.GetConditionData(user, pb.CONDITION_ALL_RECHARGE_DAY, 0)
		data, _ := strconv.Atoi(strconv.Itoa(conditionData)[4:])
		return data, data >= value
	}
	//战宠附体
	this.tooCheckers[pb.CONDITION_PET_APPENDAGE] = func(user *objs.User, heroIndex int, value []int) ([]int, bool) {
		petId, needLv := value[0], value[1]
		lv, ok := user.PetAppendage[petId]
		return []int{petId, lv}, ok && lv >= needLv
	}
	//所有击杀monster
	this.tooCheckers[pb.CONDITION_ALL_KILL_MONSTER] = func(user *objs.User, heroIndex int, value []int) ([]int, bool) {
		monsterId, num := value[0], value[1]
		conditionData := this.GetConditionData(user, pb.CONDITION_ALL_KILL_MONSTER, monsterId)
		return []int{monsterId, conditionData}, conditionData >= num
	}
	//所有击杀stageId
	this.tooCheckers[pb.CONDITION_ALL_KILL_STAGE] = func(user *objs.User, heroIndex int, value []int) ([]int, bool) {
		stageId, num := value[0], value[1]
		conditionData := this.GetConditionData(user, pb.CONDITION_ALL_KILL_STAGE, stageId)
		return []int{stageId, conditionData}, conditionData >= num
	}

	var killStageByType = func(user *objs.User, heroIndex, value, fightType int) (int, bool) {
		num := 0
		state, nums := this.checkConditionNum(user, pb.CONDITION_ALL_KILL_STAGE)
		if state {
			return num, num >= value
		}
		dataSlice := nums
		for i := 0; i < len(dataSlice); i += 2 {
			stageCfg := gamedb.GetStageStageCfg(dataSlice[i])
			if stageCfg.Type == fightType {
				num += dataSlice[i+1]
			}
		}
		return num, num >= value
	}
	//击杀个人首领
	this.checkers[pb.CONDITION_CHALLENGE_ONE_TIME_GE_REN_BOSS] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return killStageByType(user, heroIndex, value, constFight.FIGHT_TYPE_PERSON_BOSS)
	}
	//击杀VIP首领
	this.checkers[pb.CONDITION_CHALLENGE_VIP_BOSS] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return killStageByType(user, heroIndex, value, constFight.FIGHT_TYPE_VIPBOSS)
	}
	//击杀野外首领
	this.checkers[pb.CONDITION_KILL_YE_WAI_BOSS] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return killStageByType(user, heroIndex, value, constFight.FIGHT_TYPE_FIELDBOSS)
	}
	//击杀世界首领
	this.checkers[pb.CONDITION_ALL_KILL_WORLDLEADER] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return killStageByType(user, heroIndex, value, constFight.FIGHT_TYPE_CROSS_WORLD_LEADER)
	}
	//击杀远古首领
	this.checkers[pb.CONDITION_ALL_KILL_ANCIENT_BOSS] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return killStageByType(user, heroIndex, value, constFight.FIGHT_TYPE_ANCIENT_BOSS)
	}

	//击杀打宝秘境怪物
	this.checkers[pb.CONDITION_DAO_BAO_MI_JIN_BOSS_KILL_NUMS] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := this.GetConditionData(user, pb.CONDITION_DAO_BAO_MI_JIN_BOSS_KILL_NUMS, 0)
		return conditionData, conditionData >= value
	}

	//累计击杀玩家
	this.checkers[pb.CONDITION_ALL_KILL_USER] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := this.GetConditionData(user, pb.CONDITION_ALL_KILL_USER, 0)
		return conditionData, conditionData >= value
	}
	//累计签到
	this.checkers[pb.CONDITION_ALL_SIGN] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := this.GetConditionData(user, pb.CONDITION_ALL_SIGN, 0)
		return conditionData, conditionData >= value
	}
	//累计神机宝库
	this.checkers[pb.CONDITION_ALL_CARD_ACTIVITY] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := this.GetConditionData(user, pb.CONDITION_ALL_CARD_ACTIVITY, 0)
		return conditionData, conditionData >= value
	}
	//累计寻龙探宝
	this.checkers[pb.CONDITION_ALL_TREASURE] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := this.GetConditionData(user, pb.CONDITION_ALL_TREASURE, 0)
		return conditionData, conditionData >= value
	}
	//累计购买开服礼包
	this.checkers[pb.CONDITION_ALL_BUY_OPEN_GIFT] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := this.GetConditionData(user, pb.CONDITION_ALL_BUY_OPEN_GIFT, 0)
		return conditionData, conditionData >= value
	}
	//累计竞技场赢
	this.checkers[pb.CONDITION_ALL_COMPETIVE_WIN] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := this.GetConditionData(user, pb.CONDITION_ALL_COMPETIVE_WIN, 0)
		return conditionData, conditionData >= value
	}
	//累计完成每日任务
	this.checkers[pb.CONDITION_ALL_FINISH_DAILY_TASK] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := this.GetConditionData(user, pb.CONDITION_ALL_FINISH_DAILY_TASK, 0)
		return conditionData, conditionData >= value
	}
	//累计完成每日任务活跃度
	this.checkers[pb.CONDITION_ALL_DAILY_TASK_LIVENESS] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := this.GetConditionData(user, pb.CONDITION_ALL_DAILY_TASK_LIVENESS, 0)
		return conditionData, conditionData >= value
	}
	//累计完成副本
	this.checkers[pb.CONDITION_ALL_FINISH_COPY] = func(user *objs.User, heroIndex, value int) (int, bool) {
		conditionData := this.GetConditionData(user, pb.CONDITION_ALL_FINISH_COPY, 0)
		return conditionData, conditionData >= value
	}

	var rankFunc = func(user *objs.User, heroIndex, value, rankType int) (int, bool) {
		conditionData := this.GetRank().GetRankScore(rankType, user.Id)
		return conditionData, conditionData >= value
	}
	//战士排行榜
	this.checkers[pb.CONDITION_ZHANSHI_COMBAT] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return rankFunc(user, heroIndex, value, pb.RANKTYPE_COMBAT_ZHANSHI)
	}
	//法师排行榜
	this.checkers[pb.CONDITION_FASHI_COMBAT] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return rankFunc(user, heroIndex, value, pb.RANKTYPE_COMBAT_FASHI)
	}
	//道士排行榜
	this.checkers[pb.CONDITION_DAOSHI_COMBAT] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return rankFunc(user, heroIndex, value, pb.RANKTYPE_COMBAT_DAOSHI)
	}
	//激活&个绝学技能
	this.checkers[pb.CONDITION_ACTIVE_JUE_XUE_SKILL] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return len(user.Juexues), len(user.Juexues) >= value
	}
	//打宝神器阶
	this.tooCheckers[pb.CONDITION_DABAO_EQUIP_GRADE] = func(user *objs.User, heroIndex int, value []int) ([]int, bool) {
		t, g := value[0], value[1]
		lv := user.DaBaoEquip[t]
		return []int{t, lv}, lv >= g
	}

	//天赋总等级
	this.checkers[pb.CONDITION_TIAN_FU_ALL_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		allLv := 0
		for _, heroInfo := range user.Heros {
			if heroInfo.Talent.TalentList != nil {
				for _, data := range heroInfo.Talent.TalentList {
					if data != nil && data.Talents != nil {
						for _, lv := range data.Talents {
							allLv += lv
						}
					}
				}
			}
		}
		return allLv, allLv >= value
	}

	//头衔升至&级
	this.checkers[pb.CONDITION_TOU_XIAN_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		return user.Label.Id, user.Label.Id >= value
	}

	//打宝神器总等级升至xx级
	this.checkers[pb.CONDITION_DAO_BAO_SHEN_QI_ALL_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {

		allLv := 0
		if user.DaBaoEquip != nil {
			for _, lv := range user.DaBaoEquip {
				allLv += lv
			}
		}
		return allLv, allLv >= value
	}

	//激活（&/&）个远古神技
	this.checkers[pb.CONDITION_ACTIVE_YUAN_GU_SKILL_NUM] = func(user *objs.User, heroIndex, value int) (int, bool) {

		allLv := 0
		for _, heroInfo := range user.Heros {
			if heroInfo.AncientSkill != nil && heroInfo.AncientSkill.Level > 0 {
				allLv++
			}
		}
		return allLv, allLv >= value
	}

	//传世强化总等级达到（&/&）级
	this.checkers[pb.CONDITION_CHUAN_SI_QIANG_HUA_ALL_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		allLv := 0
		for _, heroInfo := range user.Heros {
			if heroInfo.ChuanshiStrengthen != nil {
				for _, lv := range heroInfo.ChuanshiStrengthen {
					allLv += lv
				}
			}
		}
		return allLv, allLv >= value
	}

	//高于&级不可进入
	this.checkers[pb.CONDITION_LEVEL_MAX_LIMIT] = func(user *objs.User, heroIndex, value int) (int, bool) {
		maxLevel := 0
		for _, heroInfo := range user.Heros {
			if heroInfo.ExpLvl > maxLevel {
				maxLevel = heroInfo.ExpLvl
			}
		}
		return maxLevel, maxLevel <= value
	}

	//内功提升
	this.tooCheckers[pb.CONDITION_UPGRADE_NEI_GONG] = func(user *objs.User, heroIndex int, value []int) ([]int, bool) {
		grade, order := value[0], value[1]
		maxGrade, maxOrder := 0, 0
		for _, hero := range user.Heros {
			for _, id := range hero.Inside.Acupoint {
				insideCfg := gamedb.GetInsideArtInsideArtCfg(id)
				if insideCfg.Grade > grade || insideCfg.Grade == grade && insideCfg.Order >= order {
					return []int{insideCfg.Grade, insideCfg.Order}, true
				} else {
					if insideCfg.Grade > maxGrade {
						maxGrade = insideCfg.Grade
					}
					if insideCfg.Order > maxOrder {
						maxOrder = insideCfg.Order
					}
				}
			}
		}
		return []int{maxGrade, maxOrder}, false
	}

	//挑战（&/&)次暗殿首领
	this.checkers[pb.CONDITION_TIAO_ZHAN_AN_DIAN_BOSS] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetCondition().GetConditionData(user, pb.CONDITION_TIAO_ZHAN_AN_DIAN_BOSS, 0)
		return num, num >= value
	}

	//挑战（&/&）次远古首领
	this.checkers[pb.CONDITION_TIAO_ZHAN_YUAN_GU_BOSS] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetCondition().GetConditionData(user, pb.CONDITION_TIAO_ZHAN_YUAN_GU_BOSS, 0)
		return num, num >= value
	}
	//消耗金币
	this.checkers[pb.CONDITION_CONSUME_GOLD] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetCondition().GetConditionData(user, pb.CONDITION_CONSUME_GOLD, 0)
		return num, num >= value
	}
	//加入公会
	this.checkers[pb.CONDITION_JOIN_GUILD] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetCondition().GetConditionData(user, pb.CONDITION_JOIN_GUILD, 0)
		return num, num >= value
	}

	//累计击杀boss 的数量
	this.checkers[pb.CONDITION_KILL_BOSS_NUM] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetCondition().GetConditionData(user, pb.CONDITION_KILL_BOSS_NUM, 0)
		return num, num >= value
	}

	//累计击杀小怪 的数量
	this.checkers[pb.CONDITION_KILL_MONSTER_NUM] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetCondition().GetConditionData(user, pb.CONDITION_KILL_MONSTER_NUM, 0)
		return num, num >= value
	}

	//单件装备强化到Max级
	this.checkers[pb.CONDITION_ONE_EQUIP_INTENSIFY_MAX_LV] = func(user *objs.User, heroIndex, value int) (int, bool) {
		num := this.GetCondition().GetConditionData(user, pb.CONDITION_ONE_EQUIP_INTENSIFY_MAX_LV, 0)
		return num, num >= value
	}

	//	三个角色总共穿戴&件指定阶数品质的装备
	this.tooCheckers[pb.CONDITION_ALL_HEROS_WEAR_ASSIGN_EQUIP] = func(user *objs.User, heroIndex int, value []int) ([]int, bool) {
		maxNum := 0
		for _, hero := range user.Heros {
			for _, v1 := range hero.Equips {
				if v1.ItemId > 0 {
					cfg := gamedb.GetEquipEquipCfg(v1.ItemId)
					if cfg == nil {
						continue
					}
					if cfg.Class > value[1] || (cfg.Class == value[1] && cfg.Quality >= value[2]) {
						maxNum++
					}
				}
			}

		}
		return []int{maxNum}, maxNum >= value[0]
	}
	// 开通特权
	this.checkers[pb.CONDITION_PRIVILEGE_OPEN] = func(user *objs.User, heroIndex, value int) (int, bool) {
		if _, ok := user.Privilege[value]; ok {
			return value, true
		}
		return 0, false
	}
	// 每日充值额度
	this.checkers[pb.CONDITION_RECHARGE_DAILY] = func(user *objs.User, heroIndex, value int) (int, bool) {
		if user.DayStateRecord.DailyRecharge >= value {
			return user.DayStateRecord.DailyRecharge, true
		}
		return 0, false
	}
	// 等级区间
	this.tooCheckers[pb.CONDITION_LEVEL_BETWEEN] = func(user *objs.User, heroIndex int, value []int) ([]int, bool) {
		maxLv := 0
		for _, hero := range user.Heros {
			if hero.ExpLvl > maxLv {
				maxLv = hero.ExpLvl
			}
		}
		return []int{maxLv}, maxLv >= value[0] && maxLv <= value[1]
	}

}

//func (this *ConditionManager) checkMainLineTask(user *objs.User, c *gamedb.Condition) (int, bool) {
//	taskT := gameDb().GetTask(user.MainLineTask.TaskId)
//	if taskT == nil {
//		return 0, false
//	}
//	latestDoneTaskId := taskT.Order
//	curTaskT := gameDb().GetTask(c.V)
//	if curTaskT == nil {
//		return 0, false
//	}
//	return latestDoneTaskId, latestDoneTaskId >= curTaskT.Order
//}

func (this *ConditionManager) genFnCheckViaSavedData(conditionId int) func(user *objs.User, c *gamedb.Condition) (int, bool) {
	return func(user *objs.User, c *gamedb.Condition) (int, bool) {
		if savedCount, ok := user.Conditions[c.K]; ok {
			count := 0
			if len(savedCount) > 1 {
				count = savedCount[len(savedCount)-1]
			}
			return count, count >= c.V
		}
		return 0, false
	}
}

func (this *ConditionManager) CheckItemIsCanBeUse(user *objs.User, itemId int) (bool, *gamedb.EquipEquipCfg) {

	cfg := gamedb.GetEquipEquipCfg(itemId)
	if cfg != nil {
		for _, v := range user.Heros {
			if v.Job == cfg.Condition[pb.CONDITION_JOB] {
				return true, cfg
			}
		}
	}
	return false, nil
}

func (this *ConditionManager) CheckFunctionOpen(user *objs.User, moduleType int) error {

	cfg := gamedb.GetFunctionFunctionCfg(moduleType)
	if cfg == nil {
		logger.Error("GetFunctionFunctionCfg nil moduleType:%v", moduleType)
		return gamedb.ERRPARAM
	}
	if !this.CheckMulti(user, -1, cfg.Condition) {
		return gamedb.ERRACTIVITYNOTOPEN
	}

	return nil
}

func (this *ConditionManager) checkConditionNum(user *objs.User, conditionType int) (bool, []int) {

	if user.Conditions[conditionType] == nil || len(user.Conditions[conditionType]) <= 0 {
		return true, []int{}
	}
	return false, user.Conditions[conditionType]
}
