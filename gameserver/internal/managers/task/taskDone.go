package task

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constTask"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

func (this *TaskManager) TaskDone(user *objs.User, op *ophelper.OpBagHelperDefault) error {
	logger.Debug("TaskDone userId:%v  user.MainLineTask.TaskId:%v  process:%v", user.Id, user.MainLineTask.TaskId, user.MainLineTask.Process)
	taskConf := gamedb.GetTaskConditionCfg(user.MainLineTask.TaskId)
	awards := this.getTaskJobAward(user, taskConf)
	if !this.GetBag().CheckHasEnoughPos(user, awards) {
		return gamedb.ERRBAGHAVENOENOUGHSPACE
	}
	isCanNext := false
	oldTaskId := user.MainLineTask.TaskId

	if taskConf.ConditionType == pb.CONDITION_LEARN_TO_WEAR_SKILL || taskConf.ConditionType == pb.CONDITION_UPGRADE_SKILL || taskConf.ConditionType == pb.CONDITION_UPGRADE_NEI_GONG || taskConf.ConditionType == pb.CONDITION_WEAR_CHUAN_SHI_EQUIP {
		isCanNext, _ = this.SpecialCheck(user, taskConf.ConditionType, taskConf.ConditionValue)
	} else if taskConf.ConditionType == pb.CONDITION_TWO_HERO_LV || taskConf.ConditionType == pb.CONDITION_THREE_HERO_LV {
		if user.MainLineTask.Process == 1 {
			isCanNext = true
		}
	} else {
		maxProcess := this.getMaxProcess(taskConf, user)
		nowProcess := user.MainLineTask.Process
		if maxProcess != nowProcess {
			nowProcess = this.getNowProcess(user)
		}
		if nowProcess >= maxProcess {

			if awards != nil {
				this.GetBag().AddItems(user, awards, op)
			}
			//下一个任务
			nextTask := taskConf.NextId
			if nextTask > 0 {
				user.MainLineTask.TaskId = nextTask
			} else {
				user.MainLineTask.TaskId += 1
			}
			user.MainLineTask.Process = 0
			user.MainLineTask.MarkProcess = 0
			//通知任务系统
			this.GetTask().AddTaskProcess(user, pb.CONDITION_TASK_NUMS, 1)
		}
	}
	if isCanNext {
		if awards != nil {
			this.GetBag().AddItems(user, awards, op)
		}
		//下一个任务
		nextTask := taskConf.NextId
		logger.Debug("getTaskJobAward taskId:%v  nextTaskId:%v", taskConf.Id, nextTask)
		if nextTask > 0 {
			user.MainLineTask.TaskId = nextTask
		} else {
			user.MainLineTask.TaskId += 1
		}
		user.MainLineTask.Process = 0
		user.MainLineTask.MarkProcess = 0
		//通知任务系统
		this.GetTask().AddTaskProcess(user, pb.CONDITION_TASK_NUMS, 1)
		logger.Debug("userId:%v isCanNext:%v TaskId:%v  process:%v  MarkProcess:%v", user.Id, isCanNext, user.MainLineTask.TaskId, user.MainLineTask.Process, user.MainLineTask.MarkProcess)
	}
	user.Dirty = true
	this.UpdateTaskProcess(user, true, true)
	if oldTaskId != user.MainLineTask.TaskId {

		kyEvent.UserTask(user, constTask.TASK_TYPE_MAIN, oldTaskId, user.MainLineTask.TaskId)
	}
	return nil
}

func (this *TaskManager) getTaskJobAward(user *objs.User, taskConf *gamedb.TaskConditionCfg) gamedb.ItemInfos {

	job := user.Heros[constUser.USER_HERO_MAIN_INDEX].Job
	if job == pb.JOB_ZHANSHI {
		return taskConf.AwardZhan
	} else if job == pb.JOB_FASHI {
		return taskConf.AwardFa
	} else {
		return taskConf.AwardDao
	}
}

func (this *TaskManager) SpecialCheck(user *objs.User, conditionType int, conditionValue gamedb.IntSlice) (bool, int) {
	isCanNext := true
	if conditionType == pb.CONDITION_LEARN_TO_WEAR_SKILL {
		if len(conditionValue) < 4 {
			logger.Error("SpecialCheck CONDITION_LEARN_TO_WEAR_SKILL 配置填错  len(taskConf.ConditionValue):%v", len(conditionValue))
			return false, 0
		}
		for _, hero := range user.Heros {
			if hero.Index == conditionValue[3] {
				flag := false
				skillId := conditionValue[0]
				if hero.Job == 1 {
					skillId = conditionValue[0]
				} else if hero.Job == 2 {
					skillId = conditionValue[1]
				} else {
					skillId = conditionValue[2]
				}
				logger.Debug(" userId:%v hero.Index:%v  job:%v  skillId:%v   hero.SkillBag:%v", user.Id, hero.Index, hero.Job, skillId, hero.SkillBag)
				for _, id := range hero.SkillBag {
					if skillId == id {
						flag = true
					}
				}
				if flag == false {
					isCanNext = false
				}
			}
		}
		if user.MainLineTask.Process >= 1 {
			isCanNext = true
		}
	}

	if conditionType == pb.CONDITION_UPGRADE_SKILL {
		for _, hero := range user.Heros {
			if hero.Index == 1 {
				skillId := conditionValue[0]
				if hero.Job == 1 {
					skillId = conditionValue[0]
				} else if hero.Job == 2 {
					skillId = conditionValue[1]
				} else {
					skillId = conditionValue[2]
				}
				if hero.Skills[skillId] == nil {
					isCanNext = false
					logger.Debug(" isCanNext:%v  skillId:%v  hero:%v", isCanNext, skillId, hero.Job)
					return isCanNext, 0
				}

				if hero.Skills[skillId].Lv < conditionValue[3] {
					isCanNext = false
					logger.Debug(" isCanNext:%v  skillId:%v  hero:%v", isCanNext, skillId, hero.Job)
					return isCanNext, 0
				}

				isCanNext = true
			}
		}
	}

	//	三个角色总共穿戴&件指定阶数品质的装备
	if conditionType == pb.CONDITION_ALL_HEROS_WEAR_ASSIGN_EQUIP {

		maxNum := 0
		for _, hero := range user.Heros {
			if hero == nil {
				continue
			}
			for _, v1 := range hero.Equips {
				if v1.ItemId > 0 {
					cfg := gamedb.GetEquipEquipCfg(v1.ItemId)
					if cfg == nil {
						continue
					}
					logger.Debug(" nickName:%v GetEquipEquipCfg itemId:%v class:%v  quality:%v  taskConf:%v", user.NickName, v1.ItemId, cfg.Class, cfg.Quality, conditionType)
					if cfg.Class > conditionValue[1] || (cfg.Class == conditionValue[1] && cfg.Quality >= conditionValue[2]) {
						maxNum++
					}
					logger.Debug(" maxNum:%v", maxNum)
					if maxNum >= conditionValue[0] {
						return true, maxNum
					}

				}
			}

		}
		return true, maxNum
	}

	//	任意角色内功达到指定等级时，完成任务，填insideArt表的grade和order字段
	if conditionType == pb.CONDITION_UPGRADE_NEI_GONG {
		for _, hero := range user.Heros {
			inside := hero.Inside

			insideCfg := gamedb.GetInsideArtInsideArtCfg(inside.Acupoint[pb.INSIDETYPE_ONE])
			if insideCfg.Grade > conditionValue[0] {
				return true, 1
			}
			if insideCfg.Grade == conditionValue[0] && insideCfg.Order >= conditionValue[1] {
				return true, 1
			}
		}
		return false, 0
	}

	//三角色穿戴&件&阶传世装备
	if conditionType == pb.CONDITION_WEAR_CHUAN_SHI_EQUIP {
		num := 0
		for _, hero := range user.Heros {
			for _, eItem := range hero.ChuanShi {
				cfg := gamedb.GetChuanShiEquipChuanShiEquipCfg(eItem)
				if cfg == nil || len(conditionValue) < 2 {
					continue
				}
				if cfg.Level >= conditionValue[1] {
					num++
				}
			}
		}
		if num > conditionValue[0] {
			num = conditionValue[0]
		}
		return num >= conditionValue[0], num
	}

	num := 1
	if isCanNext == false {
		num = 0
	}
	logger.Debug(" isCanNext:%v", isCanNext)
	return isCanNext, num
}

func (this *TaskManager) SendSpecialCheckTaskInfo(user *objs.User, conditionType int) {

	cfg := gamedb.GetTaskConditionCfg(user.MainLineTask.TaskId)
	if cfg != nil {
		if cfg.ConditionType == conditionType {
			state, num := this.GetTask().SpecialCheck(user, cfg.ConditionType, cfg.ConditionValue)
			if state {
				if user.MainLineTask.Process >= num {
					return
				}
				if num > user.MainLineTask.Process {
					user.MainLineTask.Process = num
				}
				taskInfo := &pb.TaskInfoNtf{
					TaskId:      int32(user.MainLineTask.TaskId),
					Process:     int32(user.MainLineTask.Process),
					MarkProcess: int32(user.MainLineTask.MarkProcess),
				}
				this.GetUserManager().SendMessage(user, taskInfo, true)
			}
		}
	}
}
