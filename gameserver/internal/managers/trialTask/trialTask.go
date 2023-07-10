package trialTask

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewTrialTaskManager(module managersI.IModule) *TrialTaskManager {
	return &TrialTaskManager{IModule: module}
}

type TrialTaskManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *TrialTaskManager) TrialTaskLoad(user *objs.User, ack *pb.TrialTaskInfoAck) {

	serverOpenTime := this.GetSystem().GetServerOpenTimeByServerId(base.Conf.ServerId)
	openTime := []int{gamedb.GetConf().TrialDuration, 0, 0}
	openDay := this.GetSystem().GetServerOpenDaysByServerId(base.Conf.ServerId)
	allInfo := gamedb.GetAllTrialTask()
	showBeforeNum := false
	if openDay > gamedb.GetConf().TrialDuration {
		showBeforeNum = true
	}
	ack.TrialTaskInfos = make(map[int32]*pb.TrialTaskInfo)
	ack.HaveGetStageId = make([]int32, 0)
	logger.Debug("showBeforeNum:%v   allInfo:%v", showBeforeNum, allInfo)
	for _, info := range allInfo {
		num := this.getTrialTaskInfoNum(user, info)
		if showBeforeNum {
			num = user.TrialTaskInfo[info.Condition].MarkNum
		}
		if num > info.Value[0] {
			num = info.Value[0]
		}
		logger.Debug("info.Id:%v, info.Condition:%v, info.Value:%v, num:%v ,info.Value[0]:%v", info.Id, info.Condition, info.Value, num, info.Value[0])
		ack.TrialTaskInfos[int32(info.Id)] = &pb.TrialTaskInfo{
			NowNum: int32(num),
			IsGet:  int32(user.TrialTaskInfo[info.Condition].IsGetAward),
		}
	}
	for id, state := range user.TrialTaskInfoStage {
		if state == 1 {
			ack.HaveGetStageId = append(ack.HaveGetStageId, int32(id))
		}
	}
	ack.EndTime = common.CalcEndTime(serverOpenTime, openTime, 0)
	return
}

//领取试炼之路奖励
func (this *TrialTaskManager) GetTrialTaskAward(user *objs.User, id int, ack *pb.TrialTaskGetAwardAck, op *ophelper.OpBagHelperDefault) error {

	if !this.checkIsOpen(user) {
		return gamedb.ERRACTIVITYCLOSE
	}

	cfg := gamedb.GetTrialTaskTrialTaskCfg(id)
	if cfg == nil {
		return gamedb.ERRPARAM
	}

	if user.TrialTaskInfo[cfg.Condition].IsGetAward > 0 {
		return gamedb.ERRHAVEGETREWARD
	}

	num := this.getTrialTaskInfoNum(user, cfg)
	if num < cfg.Value[0] {
		return gamedb.ERRCONDITION
	}

	this.GetBag().AddItems(user, cfg.Rewards, op)

	user.TrialTaskInfo[cfg.Condition].IsGetAward = 1
	ack.IsGet = 1
	ack.Id = int32(id)
	return nil
}

//领取阶段奖励
func (this *TrialTaskManager) GetStageAward(user *objs.User, id int, ack *pb.TrialTaskGetStageAwardAck, op *ophelper.OpBagHelperDefault) error {

	if !this.checkIsOpen(user) {
		return gamedb.ERRACTIVITYCLOSE
	}

	cfg := gamedb.GetTrialTotalRewardTrialTotalRewardCfg(id)
	if cfg == nil {
		return gamedb.ERRPARAM
	}

	if user.TrialTaskInfoStage[id] >= 1 {
		return gamedb.ERRAWARDGET
	}

	allInfo := gamedb.GetAllTrialTask()
	allNum := 0

	for _, info := range allInfo {
		num := this.getTrialTaskInfoNum(user, info)
		if num >= info.Value[0] {
			allNum++
		}
	}

	if allNum < cfg.TaskNum {
		return gamedb.ERRCONDITION
	}

	this.GetBag().AddItems(user, cfg.TotalReward, op)
	ack.Id = int32(id)
	ack.IsGet = 1
	user.TrialTaskInfoStage[id] = 1
	user.Dirty = true
	return nil
}

func (this *TrialTaskManager) OfflineSaveTrialTaskInfo(user *objs.User) {

	openDay := this.GetSystem().GetServerOpenDaysByServerId(base.Conf.ServerId)
	if openDay > gamedb.GetConf().TrialDuration || user.GetMaxHeroLv() < 90 {
		return
	}

	allInfo := gamedb.GetAllTrialTask()
	for _, info := range allInfo {
		num := this.getTrialTaskInfoNum(user, info)
		user.TrialTaskInfo[info.Condition].MarkNum = num
	}
	return
}

func (this *TrialTaskManager) SendTrialTaskInfoNtf(user *objs.User, conditionType int) {
	if user.TrialTaskInfo[conditionType] == nil {
		return
	}

	openDay := this.GetSystem().GetServerOpenDaysByServerId(base.Conf.ServerId)
	if openDay > gamedb.GetConf().TrialDuration {
		return
	}

	if user.TrialTaskInfo[conditionType].IsGetAward > 0 {
		return
	}
	allData := gamedb.GetAllTrialTask()
	num := 0
	for _, data := range allData {
		if data.Condition == conditionType {
			if len(data.Value) > 1 {
				_, num = this.GetTask().SpecialCheck(user, conditionType, data.Value)
			} else {
				num, _ = this.GetCondition().Check(user, -1, conditionType, data.Value[0])
			}
			if num > data.Value[0] {
				num = data.Value[0]
			}
			this.GetUserManager().SendMessage(user, &pb.TrialTaskInfoNtf{Id: int32(data.Id), Num: int32(num)}, true)
			return
		}
	}
}

//活动结束未领取奖励的发邮件
func (this *TrialTaskManager) SendReward() {
	openDay := this.GetSystem().GetServerOpenDaysByServerId(base.Conf.ServerId)
	if openDay != gamedb.GetConf().TrialDuration+1 {
		return
	}
	allBaseUserInfo := this.GetUserManager().GetAllUsersBasicInfo()
	for _, info := range allBaseUserInfo {
		if info.Level < 90 {
			continue
		}
		userInfo := this.GetUserManager().GetUser(info.Id)
		if userInfo == nil {
			userInfo = this.GetUserManager().GetOfflineUserInfo(info.Id)
		}
		if userInfo == nil {
			continue
		}

		items := this.getUserMailReward(userInfo)
		if len(items) > 0 {
			this.GetMail().SendSystemMailWithItemInfos(userInfo.Id, constMail.MAILTYPE_TRIAL_TASK, []string{}, items)
		}
	}

}

func (this *TrialTaskManager) OnlineCheck(user *objs.User) {

	allInfo := gamedb.GetAllTrialTask()
	for _, info := range allInfo {
		if user.TrialTaskInfo[info.Condition] == nil {
			user.TrialTaskInfo[info.Condition] = &model.TrialTaskInfo{}
		}
	}
}

func (this *TrialTaskManager) getUserMailReward(user *objs.User) gamedb.ItemInfos {

	items := gamedb.ItemInfos{}
	allInfo := gamedb.GetAllTrialTask()
	allNum := 0
	for _, info := range allInfo {
		num := this.getTrialTaskInfoNum(user, info)
		if num >= info.Value[0] {
			allNum++
			if user.TrialTaskInfo[info.Condition] == nil {
				user.TrialTaskInfo[info.Condition] = &model.TrialTaskInfo{}
			}
			if user.TrialTaskInfo[info.Condition].IsGetAward == 0 {
				items = append(items, info.Rewards...)
			}
		}
	}

	stageInfos := gamedb.GetAllTrialTaskStageAward()
	for _, info := range stageInfos {
		if allNum >= info.TaskNum {
			if user.TrialTaskInfoStage[info.Id] == 0 {
				items = append(items, info.TotalReward...)
			}
		}
	}
	return items
}

func (this *TrialTaskManager) getTrialTaskInfoNum(user *objs.User, info *gamedb.TrialTaskTrialTaskCfg) int {
	num := 0
	if len(info.Value) > 1 {
		_, num = this.GetTask().SpecialCheck(user, info.Condition, info.Value)
	} else {
		num, _ = this.GetCondition().Check(user, -1, info.Condition, info.Value[0])
	}
	return num
}

func (this *TrialTaskManager) checkIsOpen(user *objs.User) bool {

	openDay := this.GetSystem().GetServerOpenDaysByServerId(base.Conf.ServerId)
	if openDay > gamedb.GetConf().TrialDuration || user.GetMaxHeroLv() < 90 {
		return false
	}
	return true
}
