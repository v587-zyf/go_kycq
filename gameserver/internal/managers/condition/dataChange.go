package condition

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"math"
)

var conditionIdMap = map[int]int{
	pb.CONDITION_CONTINUOUS_SIGN: 8991, //48 连续签到天数

	pb.CONDITION_WEAR_EQUIP:               9991, //202 穿戴任意一件装备
	pb.CONDITION_TIAO_ZHAN_SHEN_YI_FU_BEN: 9981, //207 挑战X次神翼副本并获得对应奖励后成功
	pb.CONDITION_UPGRADE_SHEN_YI_1:        9971, //208 提升一次神翼经验
	//pb.CONDITION_CHALLENGE_ONE_TIME_GE_REN_BOSS: 9961, //212 成功挑战X次个人boss后成功
	pb.CONDITION_CHALLENGE_JIN_YAN_FU_BEN: 9951, //214 挑战一次经验副本并获得对应奖励后成功
	pb.CONDITION_CHALLENGE_VIP_BOSS:       9941, //216 挑战X次VIPboss
	pb.CONDITION_HE_CHENG_LIN_DAN:         9931, //217 合成&颗灵丹
	pb.CONDITION_USE_LIN_DAN:              9921, //218 使用X颗灵丹
	//pb.CONDITION_KILL_YE_WAI_BOSS:               9911, //219 击杀对应数量的野外boss
	pb.CONDITION_XIANG_QIAN_BAO_SHI:            9891, //220 镶嵌任意一颗宝石即可
	pb.CONDITION_CHALLENGE_JIN_JI_CHANG:        9881, //222 挑战X次竞技场，无需竞技场胜负
	pb.CONDITION_GET_DAILY_TASK:                9871, //223 领取一次任意一档日常任务奖励
	pb.CONDITION_GO_TO_WA_KUANG:                9861, //228 玩家在挖矿地图点击成功开始挖矿后完成任务
	pb.CONDITION_ACTIVATE_ZHAN_CHONG:           9851, //229 玩家成功激活任意一只战宠后完成任务
	pb.CONDITION_GET_ONE_TIME_CHENG_JIU_AWARD:  9841, //233 玩家成功领取任意成就奖励后完成任务
	pb.CONDITION_CHALLENGE_QIANG_HUA_FU_BEN:    9831, //246 挑战&次强化副本后完成任务
	pb.CONDITION_CHALLENGE_FIELD_LEADER:        9821, //260 进入野外首领地图即可完成任务，不需要击杀首领
	pb.CONDITION_ALL_KILL_USER:                 9811, //600 击败玩家数量	击败&个玩家（三角色都击杀）
	pb.CONDITION_ALL_HERO_MAGIC_CIRCLE_UP_STAR: 9781, //606 全角色法阵累计升到&星
	pb.CONDITION_DAO_BAO_MI_JIN_BOSS_KILL_NUMS: 9791, //615 打宝怪物
}

/**
 *  @Description: 记录condition数据
 *  @param user
 *  @param id		conditionId
 *  @param value	数据
 *  @param check
 */
func (this *ConditionManager) RecordCondition(user *objs.User, id int, value []int) {
	this.writeConditionData(user, id, value)
}

/**
 *  @Description: 获取condition数据
 *  @param user
 *  @param id		conditionId
 *  @param value	多数据查询专用，单数据填0
 *  @return int
 */
func (this *ConditionManager) GetConditionData(user *objs.User, id int, value int) int {
	conditionId, ok := conditionIdMap[id]
	if ok {
		id = conditionId
	}

	result := 0
	dataSlice, ok := user.Conditions[id]
	if ok {
		if len(dataSlice) <= 0 {
			return result
		}
		if id > 1000 {
			endNum := id % 10
			switch endNum {
			case 1, 3:
				result = dataSlice[0]
			case 2:
				index := math.MinInt32
				for i := 0; i < len(dataSlice); i += 2 {
					if dataSlice[i] == value {
						index = i
						break
					}
				}
				if index != math.MinInt32 {
					result = dataSlice[index+1]
				}
			default:
				result = 0
			}
		} else {
			result = dataSlice[0]
		}
	}
	return result
}

func (this *ConditionManager) writeConditionData(user *objs.User, id int, value []int) {
	oldId := id
	conditionId, ok := conditionIdMap[id]
	if ok {
		id = conditionId
	}

	slice := make([]int, 0)
	if user.Conditions[id] != nil {
		slice = user.Conditions[id]
	}
	if id > 1000 {
		//尾数  是 1 是1个参数
		endNum := id % 10
		switch endNum {
		case 1:
			if len(value) > 0 {
				if len(slice) == 0 {
					slice = append(slice, value...)
				} else {
					slice[0] += value[0]
				}
			}
		case 2:
			//尾数  是 2 是2个参数 boosId,数量
			if len(value) > 0 {
				index := math.MinInt32
				for i := 0; i < len(slice); i += 2 {
					if slice[i] == value[0] {
						index = i
						break
					}
				}
				if index != math.MinInt32 {
					slice[index+1] += value[1]
				} else {
					slice = append(slice, value...)
				}
			}
		case 3:
			if len(value) > 0 {
				if len(slice) == 0 {
					slice = append(slice, value...)
				} else {
					if value[0] > slice[0] {
						slice[0] = value[0]
					}
				}
			}
		default:
			logger.Error("记录condition错误 conditionId:%v value:%v", id, value)
			return
		}
		logger.Debug("conditionId:%v value:%v", id, value)
		user.Conditions[id] = slice
	}

	this.GetGift().LimitedGiftSend(user)
	this.GetTitle().AutoActive(user)
	this.GetTrialTask().SendTrialTaskInfoNtf(user, oldId)
	this.GetLabel().SendLabelTaskNtf(user, oldId)
	this.GetAchievement().UpdateAchievementTaskProcess(user, oldId)
}
