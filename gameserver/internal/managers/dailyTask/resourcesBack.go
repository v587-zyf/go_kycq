package dailyTask

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"fmt"
	"strconv"
	"time"
)

func (this *DailyTaskManager) GetResourcesBackReward(user *objs.User, useIngot, activityId, backTimes int, ack *pb.ResourcesBackGetRewardAck, op *ophelper.OpBagHelperDefault) error {
	if useIngot > 1 || useIngot < 0 {
		return gamedb.ERRPARAM
	}

	cfg := gamedb.GetDailyTaskDailytaskCfg(activityId)
	if cfg == nil {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("item" + strconv.Itoa(activityId))
	}

	if cfg.Is_retrieve == false || cfg.Active_consume == 0 {
		return gamedb.ERRCANNOTRETRIEVE
	}

	if check := this.GetCondition().CheckMulti(user, -1, cfg.Condition); !check {
		return gamedb.ERRCONDITION
	}

	openDay := this.GetDailyTaskOpenDay(user.ServerId)

	differDay := this.GetDiffDaysBySecond(time.Now().Unix(), user.Heros[constUser.USER_HERO_MAIN_INDEX].CreatTime.Unix()) + 1
	allTimes := this.GetDailyTaskThreeDayLastChallengeTimes(user, openDay, activityId, differDay)
	dayHaveBackTimes := user.DailyTask.ResourcesHaveBackTimes[activityId]
	if backTimes > allTimes || backTimes <= 0 || backTimes+dayHaveBackTimes > allTimes+dayHaveBackTimes {
		logger.Error("GetResourcesBackReward  backTimes:%v > allTimes:%v  openDay:%v userId:%v  activityId:%v  haveBackTimes:%v", backTimes, allTimes, openDay, user.Id, activityId, dayHaveBackTimes)
		return gamedb.ERRNOTENOUGHTIMES
	}

	user.DailyTask.ResourcesHaveBackTimes[activityId] += backTimes
	num := backTimes * cfg.Active_consume
	if useIngot > 0 {
		num = cfg.Gold_consume * backTimes
	}

	monthCardPrivilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_DAILYTASK_RESOURCES_BACK)
	afterNum := num
	vipNum := 0
	if monthCardPrivilege != 0 {
		vipNum = common.CalcTenThousand(monthCardPrivilege, num)
		matchNum := vipNum - num
		afterNum = num - matchNum
	}
	logger.Debug(" userId:%v user.DailyTask.ResourcesBackExp:%v  backTimes:%v cfg.Active_consume:%v  num:%v monthCardPrivilege:%v   vipNum:%v  afterNum:%v", user.Id, user.DailyTask.ResourcesBackExp, backTimes, cfg.Active_consume, num, monthCardPrivilege, vipNum, afterNum)
	if useIngot == 0 {
		if user.DailyTask.ResourcesBackExp < afterNum {
			return gamedb.ERRHAVENOENOUGHEXP
		}
		user.DailyTask.ResourcesBackExp -= afterNum
	} else {
		enough, _ := this.GetBag().HasEnough(user, pb.ITEMID_INGOT, afterNum)
		if !enough {
			return gamedb.ERRNOTENOUGHGOODS
		}
		_ = this.GetBag().Remove(user, op, pb.ITEMID_INGOT, afterNum)
	}

	cfgs := make(gamedb.ItemInfos, 0)
	if len(cfg.FindReward) > 0 {
		for _, info := range cfg.FindReward {
			cfgs = append(cfgs, &gamedb.ItemInfo{ItemId: info.ItemId, Count: info.Count * backTimes})
		}
		this.GetBag().AddItems(user, cfgs, op)
	}
	logger.Debug("userId:%v user.DailyTask.ResourcesBackExp:%v", user.Id, user.DailyTask.ResourcesBackExp)
	this.DelCanBackTimes(user, activityId, backTimes)
	user.Dirty = true
	ack.ResourcesBackInfos = this.buildResourcesBackInfo(user, activityId)
	ack.DayResourcesBackExp = int32(user.DailyTask.ResourcesBackExp)
	return nil
}

func (this *DailyTaskManager) buildResourcesBackInfo(user *objs.User, activityId int) []*pb.ResourcesBackInfo {

	info := make([]*pb.ResourcesBackInfo, 0)
	openDay := this.GetDailyTaskOpenDay(user.ServerId)
	for _, data := range user.DailyTask.DailyTask {
		if activityId > 0 {
			if data.ActivityId != activityId {
				continue
			}
		}
		cfg := gamedb.GetDailyTaskDailytaskCfg(data.ActivityId)
		if cfg == nil {
			logger.Error("ActivityId:%v == nil", data.ActivityId)
			continue
		}
		if !cfg.Is_retrieve {
			continue
		}
		if check := this.GetCondition().CheckMulti(user, -1, cfg.Condition); !check {
			continue
		}
		differDay := this.GetDiffDaysBySecond(time.Now().Unix(), user.Heros[constUser.USER_HERO_MAIN_INDEX].CreatTime.Unix()) + 1
		//logger.Debug("openDay:%v, user.Id:%v, data.ActivityId:%v, differDay:%v", openDay, user.Id, data.ActivityId, differDay)
		if differDay <= 1 {
			continue
		}

		times := this.GetDailyTaskThreeDayLastChallengeTimes(user, openDay, data.ActivityId, differDay)

		logger.Debug("buildResourcesBackInfo openDay:%v  user.Id:%v, data.ActivityId:%v, differDay:%v  times:%v", openDay, user.Id, data.ActivityId, differDay, times)

		if user.DailyTask.ResourcesHaveBackTimes == nil {
			user.DailyTask.ResourcesHaveBackTimes = make(map[int]int)
		}
		info = append(info, &pb.ResourcesBackInfo{
			ActivityId:            int32(data.ActivityId),
			ResidueChallengeTimes: int32(times),
			HaveChallengeTimes:    int32(user.DailyTask.ResourcesHaveBackTimes[data.ActivityId]),
		})
	}

	return info

}

// 获取t1和t2的相差天数，单位：秒，0表同一天，正数表t1>t2，负数表t1<t2
func (this *DailyTaskManager) GetDiffDaysBySecond(t1, t2 int64) int {
	time1 := time.Unix(t1, 0)
	time2 := time.Unix(t2, 0)

	// 调用上面的函数
	return this.GetDiffDays(time1, time2)
}

// 获取两个时间相差的天数，0表同一天，正数表t1>t2，负数表t1<t2
func (this *DailyTaskManager) GetDiffDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

func (this *DailyTaskManager) OfflineSaveDailyTaskInfo(user *objs.User) {
	if user.Heros == nil || len(user.Heros) <= 0 {
		return
	}

	if user.DailyTask.ResourceCanBackTimes == nil {
		user.DailyTask.ResourceCanBackTimes = make(map[string]int)
	}

	//处理玩家每天 每日活动剩余的挑战次数  用于 资源找回
	openDay := this.GetDailyTaskOpenDay(user.ServerId)
	for _, info := range user.DailyTask.DailyTask {
		cfg := gamedb.GetDailyTaskDailytaskCfg(info.ActivityId)
		if cfg == nil {
			logger.Error("GetDailyTaskDailyTaskCfg nil ActivityId:%v", info.ActivityId)
			continue
		}

		if check := this.GetCondition().CheckMulti(user, -1, cfg.Condition); check {
			key := fmt.Sprintf("%v|%v", openDay, info.ActivityId)
			lastNum := cfg.Num + info.BuyChallengeTimes - info.HaveChallengeTimes
			if lastNum <= 0 {
				lastNum = -1
			}
			user.DailyTask.ResourceCanBackTimes[key] = lastNum
		}
	}

}

//每日任务 获取开服天数
func (this *DailyTaskManager) GetDailyTaskOpenDay(serverId int) int {

	openDay := this.GetSystem().GetServerOpenDaysByServerIdByExcursionTime(serverId, 0)
	return openDay
}

func (this *DailyTaskManager) GetResourcesBackAllReward(user *objs.User, ack *pb.ResourcesBackGetAllRewardAck, op *ophelper.OpBagHelperDefault) error {

	resourcesBackInfos := this.buildResourcesBackInfo(user, -1)
	allIngot := 0
	allActive := 0
	allReward := make(map[int]int)
	for _, backInfo := range resourcesBackInfos {
		cfg := gamedb.GetDailyTaskDailytaskCfg(int(backInfo.ActivityId))
		if cfg == nil {
			logger.Error("GetDailyTaskDailyTaskCfg nil  ActivityId:%v", backInfo.ActivityId)
			continue
		}

		if user.DailyTask.ResourcesHaveBackTimes[int(backInfo.ActivityId)] == int(backInfo.ResidueChallengeTimes) {
			continue
		}
		allIngot += int(backInfo.ResidueChallengeTimes) * cfg.Gold_consume
		allActive += int(backInfo.ResidueChallengeTimes) * cfg.Active_consume
		for _, data := range cfg.FindReward {
			allReward[data.ItemId] += int(backInfo.ResidueChallengeTimes) * data.Count
		}
	}
	if allActive <= 0 || allIngot <= 0 {
		return gamedb.ERRHAVEBACKALL
	}

	monthCardPrivilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_DAILYTASK_RESOURCES_BACK)

	afterNum := allIngot
	num := allActive
	if gamedb.GetConf().FindResourcesConsume == pb.ITEMID_INGOT {
		num = allIngot
	}
	vipNum := 0
	if monthCardPrivilege != 0 {
		vipNum = common.CalcTenThousand(monthCardPrivilege, num)
		matchNum := vipNum - num
		logger.Debug("vipNum:%v  num:%v", vipNum, num)
		afterNum = num - matchNum
	}

	if gamedb.GetConf().FindResourcesConsume == pb.ITEMID_INGOT {

		if has, _ := this.GetBag().HasEnough(user, pb.ITEMID_INGOT, afterNum); !has {
			return gamedb.ERRNOTENOUGHGOODS
		}

		err := this.GetBag().Remove(user, op, pb.ITEMID_INGOT, afterNum)
		if err != nil {
			return err
		}

	} else {
		if user.DailyTask.ResourcesBackExp < afterNum {
			return gamedb.ERRNOTENOUGHGOODS
		}
		user.DailyTask.ResourcesBackExp -= afterNum
	}

	for _, backInfo := range resourcesBackInfos {
		user.DailyTask.ResourcesHaveBackTimes[int(backInfo.ActivityId)] = int(backInfo.ResidueChallengeTimes)
		this.DelCanBackTimes(user, int(backInfo.ActivityId), int(backInfo.ResidueChallengeTimes))
	}

	user.Dirty = true
	allItems := make(gamedb.ItemInfos, 0)
	if len(allReward) > 0 {
		for itemId, count := range allReward {
			allItems = append(allItems, &gamedb.ItemInfo{ItemId: itemId, Count: count})
		}
		this.GetBag().AddItems(user, allItems, op)
	}
	logger.Debug("userId:%v user.DailyTask.ResourcesBackExp:%v", user.Id, user.DailyTask.ResourcesBackExp)
	ack.ResourcesBackInfos = this.buildResourcesBackInfo(user, -1)
	ack.DayResourcesBackExp = int32(user.DailyTask.ResourcesBackExp)
	return nil
}

func (this *DailyTaskManager) DelCanBackTimes(user *objs.User, activityId, backTimes int) {

	allNum := 0
	cfg := gamedb.GetDailyTaskDailytaskCfg(activityId)
	if cfg != nil {
		allNum = cfg.Num
	}

	openDay := this.GetDailyTaskOpenDay(user.ServerId)
	diffDay := this.GetDiffDaysBySecond(time.Now().Unix(), user.Heros[constUser.USER_HERO_MAIN_INDEX].CreatTime.Unix()) + 1
	if diffDay <= 1 {
		return
	}

	if diffDay == 2 {
		key := fmt.Sprintf("%v|%v", openDay-1, activityId)
		lastNum := user.DailyTask.ResourceCanBackTimes[key]
		if lastNum == 0 {
			lastNum = allNum
		}

		afterNum := lastNum - backTimes
		if afterNum <= 0 {
			afterNum = -1
		}
		user.DailyTask.ResourceCanBackTimes[key] = afterNum
		return
	}

	if diffDay == 3 {
		this.buildResourceCanBackTimes(user, openDay, activityId, allNum, backTimes)
	}

	if diffDay > 3 {
		this.buildResourceCanBackTimes1(user, openDay, activityId, allNum, backTimes)
	}
	return
}

func (this *DailyTaskManager) buildResourceCanBackTimes(user *objs.User, openDay, activityId, allNum, backTimes int) {

	key := fmt.Sprintf("%v|%v", openDay-2, activityId)
	lastNum := user.DailyTask.ResourceCanBackTimes[key]
	if lastNum == 0 {
		lastNum = allNum
	}
	afterNum := lastNum - backTimes
	twoNeedDelNum := 0
	if lastNum > 0 {
		if afterNum <= 0 {
			afterNum = -1
			twoNeedDelNum = backTimes - lastNum
		}
		user.DailyTask.ResourceCanBackTimes[key] = afterNum
	} else {
		twoNeedDelNum = backTimes
	}
	if twoNeedDelNum <= 0 {
		return
	}

	key2 := fmt.Sprintf("%v|%v", openDay-1, activityId)
	lastNum2 := user.DailyTask.ResourceCanBackTimes[key2]
	if lastNum2 == 0 {
		lastNum2 = allNum
	}

	afterTwoNum := lastNum2 - twoNeedDelNum
	if afterTwoNum <= 0 {
		afterTwoNum = -1
	}
	user.DailyTask.ResourceCanBackTimes[key2] = afterTwoNum
}

func (this *DailyTaskManager) buildResourceCanBackTimes1(user *objs.User, openDay, activityId, allNum, backTimes int) {

	key := fmt.Sprintf("%v|%v", openDay-3, activityId)
	lastNum := user.DailyTask.ResourceCanBackTimes[key]
	if lastNum == 0 {
		lastNum = allNum
	}
	twoNeedDelNum := 0
	if lastNum > 0 {
		afterNum := lastNum - backTimes
		if afterNum <= 0 {
			afterNum = -1
			twoNeedDelNum = backTimes - lastNum
		}
		user.DailyTask.ResourceCanBackTimes[key] = afterNum
	} else {
		twoNeedDelNum = backTimes
	}

	if twoNeedDelNum <= 0 {
		return
	}

	key2 := fmt.Sprintf("%v|%v", openDay-2, activityId)
	lastNum2 := user.DailyTask.ResourceCanBackTimes[key2]
	if lastNum2 == 0 {
		lastNum2 = allNum
	}
	threeNeedDelNum := 0
	if lastNum2 > 0 {
		afterNum2 := lastNum2 - twoNeedDelNum
		if afterNum2 <= 0 {
			afterNum2 = -1
			threeNeedDelNum = twoNeedDelNum - lastNum2
		}
		user.DailyTask.ResourceCanBackTimes[key2] = afterNum2
	} else {
		threeNeedDelNum = backTimes
	}

	if threeNeedDelNum <= 0 {
		return
	}

	key3 := fmt.Sprintf("%v|%v", openDay-1, activityId)
	lastNum3 := user.DailyTask.ResourceCanBackTimes[key3]
	if lastNum3 == 0 {
		lastNum3 = allNum
	}
	afterNum3 := lastNum3 - threeNeedDelNum
	if afterNum3 <= 0 {
		afterNum3 = -1
	}
	user.DailyTask.ResourceCanBackTimes[key3] = afterNum3
}

//前三天剩余挑战次数
func (this *DailyTaskManager) GetDailyTaskThreeDayLastChallengeTimes(user *objs.User, openDay, activityId, diffDay int) int {

	allNum := 0
	cfg := gamedb.GetDailyTaskDailytaskCfg(activityId)
	if cfg != nil {
		allNum = cfg.Num
	}

	allBeforeNum := 0

	for i := 1; i < diffDay; i++ {
		if i > 3 {
			break
		}

		key := fmt.Sprintf("%v|%v", openDay-i, activityId)
		lastNum := user.DailyTask.ResourceCanBackTimes[key]

		if lastNum < 0 {
			continue
		}
		if lastNum == 0 {
			lastNum = allNum
		}

		allBeforeNum += lastNum
	}
	return allBeforeNum
}
