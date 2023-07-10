package dailyTask

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constDailyTask"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"strconv"
	"strings"
	"time"
)

func NewDailyTaskManager(module managersI.IModule) *DailyTaskManager {
	return &DailyTaskManager{IModule: module}
}

type DailyTaskManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *DailyTaskManager) DailyTaskLoad(user *objs.User, ack *pb.DailyTaskLoadAck) error {

	ack.DayExp = int32(user.DailyTask.DayExp)
	ack.WeekExp = int32(user.DailyTask.WeekExp)
	ack.DayResourcesBackExp = int32(user.DailyTask.ResourcesBackExp)
	ack.HaveChallengeTimes = this.buildHaveChallengeTimesInfo(user)
	for _, v := range user.DailyTask.GetDayRewardIds {
		ack.GetDayRewardIds = append(ack.GetDayRewardIds, int32(v))
	}
	for _, v := range user.DailyTask.GetWeekRewardIds {
		ack.GetWeekRewardIds = append(ack.GetWeekRewardIds, int32(v))
	}
	ack.ResourcesBackInfos = this.buildResourcesBackInfo(user, -1)
	return nil
}

func (this *DailyTaskManager) BuyChallengeTime(user *objs.User, activityId int, ack *pb.BuyChallengeTimeAck, op *ophelper.OpBagHelperDefault) error {
	if activityId < 0 {
		return gamedb.ERRPARAM
	}
	cfg := gamedb.GetDailyTaskDailytaskCfg(activityId)
	if cfg == nil {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("item" + strconv.Itoa(activityId))
	}

	if cfg.Limit == 0 {
		return gamedb.ERRCANBUYTIMES
	}

	if user.DailyTask.DailyTask[cfg.Type] == nil {

		cfg1 := gamedb.GetDailyTaskDailyTaskCfgByType(cfg.Type)
		if cfg1 == nil {
			logger.Error("GetDailyTaskDailyTaskCfgByType type:%v nil", cfg.Type)
			return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("item" + strconv.Itoa(cfg.Type))

		}
		user.DailyTask.DailyTask[cfg.Type] = &model.DailyTaskActivityInfo{ActivityId: cfg.Id}
	}

	if user.DailyTask.DailyTask[cfg.Type].BuyChallengeTimes >= cfg.Limit {
		return gamedb.ERRBUYTIMESLIMIT
	}

	if ok, _ := this.GetBag().HasEnough(user, cfg.Consume.ItemId, cfg.Consume.Count); !ok {
		return gamedb.ERRNOTENOUGHGOODS
	}

	err := this.GetBag().Remove(user, op, cfg.Consume.ItemId, cfg.Consume.Count)
	if err != nil {
		return err
	}

	user.DailyTask.DailyTask[cfg.Type].BuyChallengeTimes += 1
	ack.ActivityId = int32(activityId)
	ack.BuyChallengTimes = int32(user.DailyTask.DailyTask[cfg.Type].BuyChallengeTimes)
	ack.HaveChallengeTime = int32(user.DailyTask.DailyTask[cfg.Type].HaveChallengeTimes)

	return nil
}

func (this *DailyTaskManager) GetReward(user *objs.User, id, types int, ack *pb.GetAwardAck, op *ophelper.OpBagHelperDefault) error {
	if id <= 0 {
		return gamedb.ERRPARAM
	}

	if types <= 0 || types > 2 {
		return gamedb.ERRPARAM
	}
	cfg := gamedb.GetDailyRewardDailyRewardCfg(id)
	if cfg == nil {
		return gamedb.ERRPARAM
	}
	haveGetId := make(map[int]bool, 0)
	for _, ids := range user.DailyTask.GetDayRewardIds {
		haveGetId[ids] = true
	}
	for _, ids := range user.DailyTask.GetWeekRewardIds {
		haveGetId[ids] = true
	}

	if haveGetId[id] {
		return gamedb.ERRHAVEGETREWARD
	}
	exp := user.DailyTask.DayExp
	if types == constDailyTask.WeekExp {
		exp = user.DailyTask.WeekExp
	}
	ids := gamedb.GetDailyTaskRewardCfgByType(types, exp)
	logger.Debug("每日任务 领取周或者日 阶段奖励id:%v  ids:%v  exp:%v  types:%v", id, ids, exp, types)
	for _, aId := range ids {
		if aId.Id == id {
			if aId.Reward != nil {
				this.GetBag().AddItems(user, aId.Reward, op)
				if types == constDailyTask.DayExp {
					user.DailyTask.GetDayRewardIds = append(user.DailyTask.GetDayRewardIds, aId.Id)
				} else if types == constDailyTask.WeekExp {
					user.DailyTask.GetWeekRewardIds = append(user.DailyTask.GetWeekRewardIds, aId.Id)
				}
				//通知任务系统
				this.GetCondition().RecordCondition(user, pb.CONDITION_GET_DAILY_TASK, []int{1})
				this.GetTask().AddTaskProcess(user, pb.CONDITION_GET_DAILY_TASK, -1)
			}
		}
	}
	for _, v := range user.DailyTask.GetDayRewardIds {
		ack.GetDayRewardIds = append(ack.GetDayRewardIds, int32(v))
	}

	for _, v := range user.DailyTask.GetWeekRewardIds {
		ack.GetWeekRewardIds = append(ack.GetWeekRewardIds, int32(v))
	}

	return nil
}

//
//  CompletionOfTask
//  @Description: 任务完成通知
//  @receiver this
//

func (this *DailyTaskManager) CompletionOfTask(user *objs.User, types, times int) {
	logger.Debug("CompletionOfTask  userId:%v  type:%v  times:%v", user.Id, types, times)
	cfg := gamedb.GetDailyTaskDailyTaskCfgByType(types)
	if cfg == nil {
		logger.Error("DailyTaskManager CompletionOfTask  配置错误 userId:%v types:%v ", user.Id, types)
		return
	}
	taskInfo := user.DailyTask.DailyTask[types]
	if taskInfo == nil {
		user.DailyTask.DailyTask[types] = &model.DailyTaskActivityInfo{ActivityId: cfg.Id}
		taskInfo = user.DailyTask.DailyTask[types]
	}

	if taskInfo.HaveChallengeTimes+times > taskInfo.BuyChallengeTimes+cfg.Num {
		logger.Debug("玩家挑战次数到达上限 userId:%v  type:%v  HaveChallengeTimes:%v  BuyChallengeTimes:%v  Num:%v times:%v", user.Id, types, taskInfo.HaveChallengeTimes, taskInfo.BuyChallengeTimes, cfg.Num, times)
		times = taskInfo.BuyChallengeTimes + cfg.Num - taskInfo.HaveChallengeTimes
	}
	if times <= 0 {
		logger.Debug("taskInfo.HaveChallengeTimes:%v times:%v  taskInfo.BuyChallengeTimes:%v  cfg.Num:%v", taskInfo.HaveChallengeTimes, times, taskInfo.BuyChallengeTimes, cfg.Num)
		return
	}

	taskInfo.HaveChallengeTimes += times
	user.DailyTask.DailyTask[types] = taskInfo
	user.DailyTask.ResourcesBackExp += cfg.Active * times
	user.DailyTask.DayExp += cfg.Active * times
	user.DailyTask.WeekExp += cfg.Active * times
	user.Dirty = true

	ack := &pb.DailyTaskLoadAck{}
	ack.DayExp = int32(user.DailyTask.DayExp)
	ack.WeekExp = int32(user.DailyTask.WeekExp)
	ack.DayResourcesBackExp = int32(user.DailyTask.ResourcesBackExp)
	_ = this.DailyTaskLoad(user, ack)
	_ = this.GetUserManager().SendMessage(user, ack, false)
	if time.Now().Hour() == 4 && time.Now().Minute() == 59 {
		_ = this.GetUserManager().Save(user, true)
	}

	this.GetTask().AddTaskProcess(user, pb.CONDITION_DAILY_TASK_LIVENESS, -1)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_FINISH_DAILY_TASK, []int{times})
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_DAILY_TASK_LIVENESS, []int{cfg.Active * times})
	return
}

func (this *DailyTaskManager) InfoInit(user *objs.User) {
	if user.DailyTask == nil {
		user.DailyTask = &model.DailyTaskInfo{}
	}
	if user.DailyTask.DailyTask == nil {
		user.DailyTask.DailyTask = make(map[int]*model.DailyTaskActivityInfo)
	}
	if user.DailyTask.GetDayRewardIds == nil {
		user.DailyTask.GetDayRewardIds = make([]int, 0)
	}

	if user.DailyTask.GetWeekRewardIds == nil {
		user.DailyTask.GetWeekRewardIds = make([]int, 0)
	}

	for i := pb.DAILYTASKACTIVITYTYPE_JING_YAN_FU_BEN; i <= pb.DAILYTASKACTIVITYTYPE_QIANG_HUA_FU_BEN; i++ {
		if user.DailyTask.DailyTask[i] == nil {
			cfg := gamedb.GetDailyTaskDailyTaskCfgByType(i)
			if cfg == nil {
				logger.Error("GetDailyTaskDailyTaskCfgByType type:%v nil", i)
				continue
			}
			user.DailyTask.DailyTask[i] = &model.DailyTaskActivityInfo{ActivityId: cfg.Id}
		}
	}
	user.Dirty = true
}

func (this *DailyTaskManager) OnLine(user *objs.User) {
	for taskType, data := range user.DailyTask.DailyTask {
		taskInfo := gamedb.GetDailyTaskDailyTaskCfgByType(taskType)
		if taskInfo == nil {
			delete(user.DailyTask.DailyTask, taskType)
			continue
		}
		if taskInfo.Id != data.ActivityId {
			user.DailyTask.DailyTask[taskType].ActivityId = taskInfo.Id
		}
	}
	date := common.GetResetTime(time.Now())
	this.Reset(user, date, false)
}

func (this *DailyTaskManager) Reset(user *objs.User, date int, isReset bool) {
	userMaterial := user.DailyTask
	if userMaterial.ResetTime != date {
		userMaterial.ResetTime = date
		user.DailyTask.DayExp = 0
		user.DailyTask.ResourcesBackExp = 0
		user.DailyTask.GetDayRewardIds = make([]int, 0)
		for i := pb.DAILYTASKACTIVITYTYPE_JING_YAN_FU_BEN; i <= pb.DAILYTASKACTIVITYTYPE_QIANG_HUA_FU_BEN; i++ {
			cfg := gamedb.GetDailyTaskDailyTaskCfgByType(i)
			if cfg == nil {
				logger.Error("GetDailyTaskDailyTaskCfgByType type:%v nil", i)
				continue
			}
			user.DailyTask.DailyTask[i] = &model.DailyTaskActivityInfo{ActivityId: cfg.Id}
		}
		user.DailyTask.ResourcesHaveBackTimes = make(map[int]int)
		if time.Now().Weekday() == 1 {
			user.DailyTask.WeekExp = 0
			user.DailyTask.GetWeekRewardIds = make([]int, 0)
		}
		this.GetTask().AddTaskProcess(user, pb.CONDITION_DAILY_TASK_LIVENESS, -1)
		this.resetResourceBackInfo(user)
		user.Dirty = true
	}

	if isReset {
		ack := &pb.DailyTaskLoadAck{}
		err := this.DailyTaskLoad(user, ack)
		if err == nil {
			_ = this.GetUserManager().SendMessage(user, ack, false)
		}
	}
}

//
//  SendReward
//  @Description: 发放达到条件但未领取的奖励
//  @receiver this
//
func (this *DailyTaskManager) SendReward() {
	logger.Info("开始发送 日常任务 达到条件但未领取的奖励")
	users, err := modelGame.GetUserModel().SearchActivityUsers(86400)
	if err != nil {
		logger.Error("SearchActivityUsers err:%v", err)
		return
	}
	logger.Info("users:%v", users)
	if len(users) == 0 {
		logger.Info("昨天一个玩家都没上过线")
		return
	}

	for _, user := range users {

		if !this.checkMaxNum(user.Id) {
			continue
		}

		if user.DailyTask == nil {
			continue
		}
		if user.DailyTask.DayExp <= 0 {
			continue
		}
		haveGet := make(map[int]bool)
		for _, v := range user.DailyTask.GetDayRewardIds {
			haveGet[v] = true
		}
		infos := gamedb.GetDailyTaskRewardCfgByType(constDailyTask.DayExp, user.DailyTask.DayExp)
		for _, v1 := range infos {
			if haveGet[v1.Id] {
				continue
			}
			if v1.Reward == nil {
				continue
			}
			//returnItem := gamedb.ItemInfos{&gamedb.ItemInfo{ItemId: v1.Reward.ItemId, Count: v1.Reward.Count}}
			this.GetMail().SendSystemMailWithItemInfos(user.Id, constMail.MAILTYPR_DAILY_TASK_Day, nil, v1.Reward)
		}
	}

	if time.Now().Weekday() == time.Monday {
		users, err := modelGame.GetUserModel().SearchActivityUsers(86400 * 5)
		if err != nil {
			logger.Error("SearchActivityUsers err:%v", err)
			return
		}
		if len(users) == 0 {
			logger.Info("前5天一个玩家都没上过线")
			return
		}

		for _, user := range users {
			if !this.checkMaxNum(user.Id) {
				continue
			}

			if user.DailyTask == nil {
				continue
			}

			if user.DailyTask.WeekExp <= 0 {
				continue
			}
			haveGet := make(map[int]bool)
			for _, v := range user.DailyTask.GetWeekRewardIds {
				haveGet[v] = true
			}
			infos := gamedb.GetDailyTaskRewardCfgByType(constDailyTask.WeekExp, user.DailyTask.WeekExp)
			for _, v1 := range infos {
				if haveGet[v1.Id] {
					continue
				}
				if v1.Reward == nil {
					continue
				}
				//returnItem := gamedb.ItemInfos{&gamedb.ItemInfo{ItemId: v1.Reward.ItemId, Count: v1.Reward.Count}}
				this.GetMail().SendSystemMailWithItemInfos(user.Id, constMail.MAILTYPR_DAILY_TASK_WEEK, nil, v1.Reward)
			}
		}
	}

	return
}

func (this *DailyTaskManager) buildHaveChallengeTimesInfo(user *objs.User) []*pb.HaveChallengeTime {

	info := make([]*pb.HaveChallengeTime, 0)

	for _, data := range user.DailyTask.DailyTask {
		info = append(info, &pb.HaveChallengeTime{
			ActivityId:        int32(data.ActivityId),
			HaveChallengeTime: int32(data.HaveChallengeTimes),
			IsGetAward:        int32(data.IsCanGetExp),
			BuyChallengTimes:  int32(data.BuyChallengeTimes),
		})
	}

	return info
}

func (this *DailyTaskManager) resetResourceBackInfo(user *objs.User) {

	if user.DailyTask.ResourceCanBackTimes == nil {
		user.DailyTask.ResourceCanBackTimes = make(map[string]int)
	}

	openDay := this.GetDailyTaskOpenDay(user.ServerId)
	for key := range user.DailyTask.ResourceCanBackTimes {
		data := strings.Split(key, "|")
		day, _ := strconv.Atoi(data[0])
		if openDay-day >= 4 {
			delete(user.DailyTask.ResourceCanBackTimes, key)
		}
	}
}

func (this *DailyTaskManager) checkMaxNum(userId int) bool {

	state := true
	sendCondition := gamedb.GetConf().DailyTime
	num := sendCondition[pb.CONDITION_ONE_HERO_LV]
	baseUserInfo := this.GetUserManager().GetUserBasicInfo(userId)
	if baseUserInfo == nil {
		state = false
		return state
	}
	maxLv := 0
	if baseUserInfo.HeroDisplay != nil {
		for _, data := range baseUserInfo.HeroDisplay {
			if data == nil {
				continue
			}
			if data.ExpLvl > maxLv {
				maxLv = data.ExpLvl
			}
		}
	}

	if maxLv < num {
		state = false
		return state
	}
	return state
}
