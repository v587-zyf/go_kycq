package achievement

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"strconv"
)

type AchievementManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewAchievementManager(m managersI.IModule) *AchievementManager {
	return &AchievementManager{
		IModule: m,
	}
}

func (this *AchievementManager) Online(user *objs.User) {
	if user.Achievement.Medal == nil {
		user.Achievement.Medal = make([]int, 0)
	}

	if user.Achievement.Task == nil {
		user.Achievement.Task = make(map[int]*model.AchievementInfo)
	}

	allConditionTypes := gamedb.GetAchievementAchievementConditionState()

	for conditionType, state := range allConditionTypes {
		if !state {
			continue
		}
		cfg := gamedb.GetAchievementAchievementCfgByConditionType(conditionType)
		if cfg == nil || len(cfg) <= 0 {
			logger.Error("GetAchievementAchievementCfgByConditionType nil conditionType:%v", conditionType)
			continue
		}
		achieveCfg := gamedb.GetAchievementAchievementCfg(cfg[0].Id)
		if achieveCfg == nil {
			logger.Error("GetAchievementAchievementCfg NowTaskId:%v  nil", cfg[0].Id)
			continue
		}

		if user.Achievement.Task[conditionType] == nil {
			user.Achievement.Task[conditionType] = &model.AchievementInfo{
				NowTaskId:  cfg[0].Id,
				NextTaskId: cfg[0].NextId,
				Process:    0,
			}
		}
	}
	user.Dirty = true
	return
}

func (this *AchievementManager) Load(user *objs.User, ack *pb.AchievementLoadAck) error {

	info, medalList := this.BuildInfo(user)
	ack.AchievementInfo = info
	ack.Medal = medalList
	ack.AllPoint = int32(user.Achievement.Point)
	return nil
}

func (this *AchievementManager) GetAward(user *objs.User, ids []int32, ack *pb.AchievementGetAwardAck, op *ophelper.OpBagHelperDefault, ) error {

	if len(ids) <= 0 {
		return gamedb.ERRPARAM
	}
	getCount := 0
	for _, idNum := range ids {
		id := int(idNum)

		cfg := gamedb.GetAchievementAchievementCfg(id)
		if cfg == nil {
			break
		}

		if user.Achievement.Task[cfg.Condition].NowTaskId == cfg.NextId {
			break
		}

		if user.Achievement.Task[cfg.Condition] == nil {
			break
		}

		if user.Achievement.Task[cfg.Condition].IsGetAll == 1 {
			break
		}

		num, _ := this.GetCondition().Check(user, -1, cfg.ConditionId, cfg.Level)

		if num < cfg.Level {
			break
		}

		for _, item := range cfg.Drop {
			this.GetBag().AddItem(user, op, item.ItemId, item.Count)
		}

		if cfg.NextId > 0 {
			user.Achievement.Task[cfg.Condition].NowTaskId = cfg.NextId
			cfg1 := gamedb.GetAchievementAchievementCfg(cfg.NextId)
			if cfg1 != nil {
				user.Achievement.Task[cfg.Condition].NextTaskId = cfg1.NextId
			}
		} else {
			user.Achievement.Task[cfg.Condition].NextTaskId = cfg.NextId
			user.Achievement.Task[cfg.Condition].IsGetAll = 1
		}
		user.Achievement.Point += cfg.Point
		getCount++
		kyEvent.AchievementInfo(user, cfg.Condition, user.Achievement.Task[cfg.Condition].NowTaskId)
	}
	user.Dirty = true
	info, _ := this.BuildInfo(user)
	ack.AllPoint = int32(user.Achievement.Point)
	ack.AchievementInfo = info
	this.GetCondition().RecordCondition(user, pb.CONDITION_GET_ONE_TIME_CHENG_JIU_AWARD, []int{getCount})
	this.GetTask().AddTaskProcess(user, pb.CONDITION_GET_ONE_TIME_CHENG_JIU_AWARD, -1)
	return nil
}

func (this *AchievementManager) ActiveMedal(user *objs.User, id int, ack *pb.ActiveMedalAck) error {

	if id <= 0 {
		return gamedb.ERRPARAM
	}

	cfg := gamedb.GetAchievementMedalMedalCfg(id)
	if cfg == nil {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("AchievementMedal" + strconv.Itoa(id))
	}

	for _, v := range user.Achievement.Medal {
		if v == id {
			return gamedb.ERRGETCONDITIONERR4
		}
	}

	if user.Achievement.Point < cfg.PointOut {
		return gamedb.ERRGETCONDITIONERR3
	}

	user.Achievement.Point -= cfg.PointOut
	user.Achievement.Medal = append(user.Achievement.Medal, id)
	user.Dirty = true
	_, medal := this.BuildInfo(user)
	this.GetUserManager().UpdateCombat(user, -1)
	ack.Medal = medal
	ack.AllPoint = int32(user.Achievement.Point)
	return nil
}

//成就任务更新
func (this *AchievementManager) AddAchievementTaskProcess(user *objs.User, conditionType, num int) {

	if user.Achievement.Task[conditionType] == nil {
		//logger.Error("AddAchievementTask conditionType:%v nil", conditionType)
		return
	}

	if user.Achievement.Task[conditionType].IsGetAll == 1 {
		return
	}

	switch conditionType {

	case pb.ACHIEVEMENTTYPE_CHENG_JIU_HERO_LV:
		user.Achievement.Task[conditionType].Process = this.GetExpPool().GetHeroMaxLv(user)
	case pb.ACHIEVEMENTTYPE_CHENG_JIU_EQUIP_MAX_LV:
		if num > user.Achievement.Task[conditionType].Process {
			user.Achievement.Task[conditionType].Process = num
		}

	default:
		user.Achievement.Task[conditionType].Process += num
	}
	user.Dirty = true
	this.sendAchievementTaskProcess(user, user.Achievement.Task[conditionType].NowTaskId, conditionType)
}

//
//  UpdateAchievementTaskProcess
//  @Description:
//  @receiver this
//  @param user
//  @param conditionType  condition表的id
//
func (this *AchievementManager) UpdateAchievementTaskProcess(user *objs.User, conditionId int) {

	conditionCfg := gamedb.GetAchievementConditionIdAndCondition()
	conditionType := conditionCfg[conditionId]

	if user.Achievement.Task[conditionType] == nil {
		//logger.Error("AddAchievementTask conditionType:%v nil", conditionType)
		return
	}

	if user.Achievement.Task[conditionType].IsGetAll == 1 {
		return
	}

	this.sendAchievementTaskProcess(user, user.Achievement.Task[conditionType].NowTaskId, conditionType)
}

func (this *AchievementManager) sendAchievementTaskProcess(user *objs.User, id, types int) {

	if err := this.GetCondition().CheckFunctionOpen(user, 160); err != nil {
		return
	}

	cfg := gamedb.GetAchievementAchievementCfg(id)
	if cfg == nil {
		return
	}
	num, _ := this.GetCondition().Check(user, -1, cfg.ConditionId, cfg.Level)
	num = this.buildNum(cfg.ConditionId, num)
	if num > cfg.Level {
		return
	}
	taskInfo := &pb.AchievementTaskInfoNtf{
		TaskId:        int32(id),
		Process:       int32(num),
		ConditionType: int32(types),
	}
	this.GetUserManager().SendMessage(user, taskInfo, true)

}

func (this *AchievementManager) BuildInfo(user *objs.User) ([]*pb.AchievementInfo, []int32) {

	achievementInfo := make([]*pb.AchievementInfo, 0)

	for conditionType, data := range user.Achievement.Task {
		cfg := gamedb.GetAchievementAchievementCfgByConditionType(conditionType)
		if cfg == nil || len(cfg) <= 0 {
			logger.Error("GetAchievementAchievementCfgByConditionType nil conditionType:%v", conditionType)
			continue
		}

		achieveCfg := gamedb.GetAchievementAchievementCfg(data.NowTaskId)
		if achieveCfg == nil {
			logger.Error("GetAchievementAchievementCfg NowTaskId:%v  nil", data.NowTaskId)
			continue
		}

		num, _ := this.GetCondition().Check(user, -1, achieveCfg.ConditionId, achieveCfg.Level)
		num = this.buildNum(achieveCfg.ConditionId, num)
		achievementInfo = append(achievementInfo, &pb.AchievementInfo{
			ConditionType: int32(conditionType),
			CanGetId:      int32(data.NowTaskId),
			Process:       int32(num),
			IsGetAllAward: int32(data.IsGetAll),
		})
	}

	medalLists := make([]int32, 0)
	for _, v := range user.Achievement.Medal {
		medalLists = append(medalLists, int32(v))
	}
	return achievementInfo, medalLists
}

func (this *AchievementManager) buildNum(conditionId, num int) int {

	if num < 0 {
		num = 0
	}
	if conditionId == pb.CONDITION_FIT_LV {
		if num < constConstant.COMPUTE_TEN_THOUSAND {
			num += constConstant.COMPUTE_TEN_THOUSAND
		}
	}
	return num
}
