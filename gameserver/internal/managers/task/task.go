package task

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

//结构的声明
type TaskManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewTaskManager(module managersI.IModule) *TaskManager {
	return &TaskManager{IModule: module}
}

func (this *TaskManager) InitNewUserTask(user *objs.User) {

	user.MainLineTask = &model.MainLineTask{
		TaskId: gamedb.GetConf().TaskInitId,
	}
}

func (this *TaskManager) UpdateTaskForKillMonster(user *objs.User, monsterId, num int) {
	//成就任务
	this.GetCondition().RecordCondition(user, pb.CONDITION_KILL_MONSTER_NUM, []int{num})
	this.GetWarOrder().WriteWarOrderTaskByKillMonster(user, monsterId, num)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_KILL_MONSTER, []int{monsterId, num})
	taskConf := gamedb.GetTaskConditionCfg(user.MainLineTask.TaskId)
	if taskConf != nil && taskConf.ConditionType == pb.CONDITION_KILL_SMALL_MONSTER {
		this.AddTaskProcess(user, taskConf.ConditionType, num)
	}
	if taskConf != nil && taskConf.ConditionType == pb.CONDITION_KILL_MONSTER {
		if taskConf.ConditionValue[0] == user.FightStageId {
			if taskConf.ConditionValue[1] == 0 || monsterId == taskConf.ConditionValue[1] {
				this.AddTaskProcess(user, taskConf.ConditionType, num)
			}
		}
	}
	if taskConf != nil && taskConf.ConditionType == pb.CONDITION_KILL_STAGE_MONSTER {
		if user.StageId == taskConf.ConditionValue[0] {
			this.AddTaskProcess(user, taskConf.ConditionType, num)
		}
	}

	user.Dirty = true
}

func (this *TaskManager) AddTaskProcess(user *objs.User, condition int, process int) {

	taskConf := gamedb.GetTaskConditionCfg(user.MainLineTask.TaskId)
	if taskConf == nil || condition != taskConf.ConditionType {
		return
	}
	maxProcess := this.getMaxProcess(taskConf, user)
	if user.MainLineTask.Process >= maxProcess {
		return
	}
	if this.GetCondition().CheckConditionType(condition) {
		value, _ := this.GetCondition().Check(user, -1, condition, maxProcess)
		if value > maxProcess {
			value = maxProcess
		}
		user.MainLineTask.Process = value
	} else {
		if user.MainLineTask.Process < maxProcess {
			user.MainLineTask.Process += process
		}
	}
	this.sendUserTaskNewProcess(user)
}

func (this *TaskManager) UpdateTaskProcess(user *objs.User, sendClient, isTaskDone bool) {

	taskConf := gamedb.GetTaskConditionCfg(user.MainLineTask.TaskId)
	if taskConf == nil {
		if sendClient && isTaskDone {
			//最后一个任务
			this.sendUserTaskNewProcess(user)
		}
		return
	}
	maxProcess := this.getMaxProcess(taskConf, user)
	oldProcess := user.MainLineTask.Process
	//logger.Debug("maxProcess :%v  oldProcess:%v", maxProcess, oldProcess)
	if maxProcess == oldProcess {
		return
	}
	nowProcess := this.getNowProcess(user)
	//logger.Debug("nowProcess:%v", nowProcess)
	if sendClient || nowProcess != oldProcess {
		user.MainLineTask.Process = nowProcess
		if user.MainLineTask.Process > maxProcess {
			user.MainLineTask.Process = maxProcess
		}
		user.Dirty = true
		this.sendUserTaskNewProcess(user)
	}
}

func (this *TaskManager) getNowProcess(user *objs.User) int {
	taskConf := gamedb.GetTaskConditionCfg(user.MainLineTask.TaskId)
	newProcess := user.MainLineTask.Process
	switch taskConf.ConditionType {
	case pb.CONDITION_GOTO_NPC:
		if user.FightStageId == taskConf.ConditionValue[0] {
			newProcess = 1
		}
	case pb.CONDITION_NPC_CHAT:
		newProcess = 1
	case pb.CONDITION_ALL_HEROS_WEAR_ASSIGN_EQUIP, pb.CONDITION_UPGRADE_NEI_GONG, pb.CONDITION_WEAR_CHUAN_SHI_EQUIP, pb.CONDITION_UPGRADE_SKILL, pb.CONDITION_LEARN_TO_WEAR_SKILL:
		_, num := this.GetTask().SpecialCheck(user, taskConf.ConditionType, taskConf.ConditionValue)
		newProcess = num
	case pb.CONDITION_KILL_SHOU_LIN:
		newProcess, _ = this.GetCondition().Check(user, -1, taskConf.ConditionType, taskConf.ConditionValue[0])
		if newProcess >= taskConf.ConditionValue[0] {
			newProcess = 1
		} else {
			newProcess = 0
		}
	default:
		state := false
		ok := this.GetCondition().CheckConditionType(taskConf.ConditionType)
		if ok {
			if len(taskConf.ConditionValue) == 1 {
				newProcess, state = this.GetCondition().Check(user, -1, taskConf.ConditionType, taskConf.ConditionValue[0])
				if taskConf.ConditionType == pb.CONDITION_TWO_HERO_LV || taskConf.ConditionType == pb.CONDITION_THREE_HERO_LV {
					if state {
						newProcess = 1
					} else {
						newProcess = 0
					}
				}
			}
		}
	}
	return newProcess
}

func (this *TaskManager) sendUserTaskNewProcess(user *objs.User) {

	taskInfo := &pb.TaskInfoNtf{
		TaskId:      int32(user.MainLineTask.TaskId),
		Process:     int32(user.MainLineTask.Process),
		MarkProcess: int32(user.MainLineTask.MarkProcess),
	}
	this.GetUserManager().SendMessage(user, taskInfo, true)

}

func (this *TaskManager) getMaxProcess(taskConf *gamedb.TaskConditionCfg, user *objs.User) int {

	switch taskConf.ConditionType {
	case pb.CONDITION_GOTO_NPC, pb.CONDITION_KILL_UNKNOWN_BOSS:
		return 1
	case pb.CONDITION_KILL_MONSTER:
		return taskConf.ConditionValue[2]
	case pb.CONDITION_KILL_SHOU_LIN, pb.CONDITION_NPC_CHAT:
		return 1
	case pb.CONDITION_KILL_STAGE_MONSTER:
		return taskConf.ConditionValue[1]
	default:
		return taskConf.ConditionValue[0]
	}
}

/**
 *  @Description: 获取战斗进入区域
 *  @param user
 *  @param stageId
 *  @return int
 */
func (this *TaskManager) GetFightEnterArea(user *objs.User, stageId int) int {

	taskConf := gamedb.GetTaskConditionCfg(user.MainLineTask.TaskId)
	if taskConf == nil {
		return 0
	}
	if taskConf.ConditionType == pb.CONDITION_KILL_MONSTER && len(taskConf.ConditionValue) >= 4 && taskConf.ConditionValue[0] == stageId {
		return taskConf.ConditionValue[3]
	}
	return 0
}

//func (this *TaskManager) CheckAndMarkMainLineTaskDone(task *model.MainLineTask) bool {
//	taskConf := gameDb().GetTask(task.TaskId)
//	if taskConf == nil {
//		return false
//	}
//	return task.Award == pb.TASKSTATUS_HAS_GOT || task.Award == pb.TASKSTATUS_CAN_GOT
//}
//
//func (this *TaskManager) KilledAMonsterCheckMainlineTask(user *objs.User, copyId, monsterId int) {
//	this.markMainLineTaskAllType(user, pb.TASKTYPE_KILL_MONSTER, copyId, monsterId)
//}
//func (this *TaskManager) ShopBuy(user *objs.User, itemId, count int) {
//	this.markMainLineTaskAllType(user, pb.TASKTYPE_SHOP_BUY, itemId, count)
//}
//func (this *TaskManager) MarkMainLineTask(user *objs.User, taskType, targetId int) {
//	this.markMainLineTaskAllType(user, taskType, targetId, 0)
//}
//
//func (this *TaskManager) Collect(user *objs.User, count int) {
//	this.markMainLineTaskAllType(user, pb.TASKTYPE_COLLECT, count, 0)
//}
//
//func (this *TaskManager) MarkEquipStrength(user *objs.User) {
//	//一键强化也算一次
//	this.MarkMainLineTask(user, pb.TASKTYPE_EQUIP_STRENGTH, 1)
//}
//func (this *TaskManager) MarkEquipWear(user *objs.User, owner, count int) {
//	this.MarkMainLineTask(user, pb.TASKTYPE_EQUIP_WEAR, count)
//	if owner == pb.OWNERTYPE_USER {
//		this.MarkMainLineTask(user, pb.TASKTYPE_EQUIP_USER_WEAR, count)
//	} else {
//		this.MarkMainLineTask(user, pb.TASKTYPE_EQUIP_HERO_WEAR, count)
//	}
//}
//
//func (this *TaskManager) markMainLineTaskAllType(user *objs.User, taskType, targetId, arg2 int) {
//	if user.MainLineTask.Award != pb.TASKSTATUS_GOING_ON {
//		return
//	}
//	taskConf := gameDb().GetTask(user.MainLineTask.TaskId)
//	if taskConf == nil || len(taskConf.TargetValue) < 1 {
//		return
//	}
//	if _, reached := m.Condition.Check(user, &taskConf.OpenCondition); !reached {
//		return
//	}
//
//	//这里杀boss和杀小怪都当作杀怪任务处理，不再判断taskType了
//	if taskType != taskConf.TaskType && !(taskType == pb.TASKTYPE_KILL_MONSTER && (taskConf.TaskType == pb.TASKTYPE_KILL_BOSS || taskConf.TaskType == pb.TASKTYPE_KILL_MONSTER || taskConf.TaskType == pb.TASKTYPE_KILL_COPY_BOSS)) {
//		//logger.Info("markMainLineTaskAllType:taskType:%d，conftype:%d", taskType, taskConf.TaskType)
//		return
//	}
//
//	switch taskConf.TaskType {
//	case pb.TASKTYPE_KILL_BOSS, pb.TASKTYPE_KILL_MONSTER, pb.TASKTYPE_KILL_COPY_BOSS:
//		if len(taskConf.TargetValue) < 3 {
//			logger.Warn("杀怪任务完成目标配置错误", taskConf.Id)
//			return
//		}
//		needCopyId := taskConf.TargetValue[0]
//		targetCount := taskConf.TargetValue[1]
//		targetMonsterId := taskConf.TargetValue[2]
//		copyId := targetId
//		monsterId := arg2
//		if needCopyId == copyId && monsterId == targetMonsterId && user.MainLineTask.Process < targetCount {
//			user.MainLineTask.Process++
//			user.Dirty = true
//			if user.MainLineTask.Process >= targetCount {
//				user.MainLineTask.Award = pb.TASKSTATUS_CAN_GOT
//			}
//		}
//	case
//		pb.TASKTYPE_EQUIP_WEAR,
//		pb.TASKTYPE_EQUIP_STRENGTH,
//		pb.TASKTYPE_EQUIP_SMELT,      // 5
//		pb.TASKTYPE_EQUIP_UPSTAR,     // 6
//		pb.TASKTYPE_STRATAGEM_ACTIVE, // 7
//		pb.TASKTYPE_WEAPON_UPGRADE,   // 10
//		pb.TASKTYPE_EQUIP_GEM,        // 11
//		pb.TASKTYPE_REALM_WEAR,       // 12
//		pb.TASKTYPE_HERO_UPGRADE,     // 14
//		pb.TASKTYPE_EQUIP_USER_WEAR,  // 16
//		pb.TASKTYPE_EQUIP_HERO_WEAR,  // 17
//		pb.TASKTYPE_COLLECT:
//
//		user.MainLineTask.Process += targetId
//		targetCount := taskConf.TargetValue[0]
//		if targetCount <= user.MainLineTask.Process {
//			user.MainLineTask.Award = pb.TASKSTATUS_CAN_GOT
//		}
//
//	case pb.TASKTYPE_SHOP_BUY: // 8
//		if len(taskConf.TargetValue) < 3 {
//			logger.Warn("购买任务完成目标配置错误", taskConf.Id)
//			return
//		}
//		if taskConf.TargetValue[0] > 0 && targetId != taskConf.TargetValue[0] {
//			return
//		}
//		if taskConf.TargetValue[2] > 0 {
//			itemT := gameDb().GetItem(targetId)
//			if itemT == nil || itemT.Type != taskConf.TargetValue[2] {
//				return
//			}
//		}
//
//		user.MainLineTask.Process += arg2
//		if taskConf.TargetValue[1] <= user.MainLineTask.Process {
//			user.MainLineTask.Award = pb.TASKSTATUS_CAN_GOT
//		}
//
//	case
//		pb.TASKTYPE_TALK, pb.TASKTYPE_OPEN_UI: // 18
//		user.MainLineTask.Process = 1
//		user.MainLineTask.Award = pb.TASKSTATUS_CAN_GOT
//
//	case
//		pb.TASKTYPE_ENTER_COPY,    // 15
//		pb.TASKTYPE_WEAPON_ACTIVE, // 9
//		pb.TASKTYPE_HERO_ACTIVE:   // 13
//		if taskConf.TargetValue[0] != targetId {
//			return
//		}
//		user.MainLineTask.Process = 1
//		user.MainLineTask.Award = pb.TASKSTATUS_CAN_GOT
//		//fmt.Printf("targetId = %+v\n", targetId)
//	case pb.TASKTYPE_CONDITION:
//		if len(taskConf.TargetValue) < 2 || taskConf.TargetValue[0] == pb.CONDITIONTYPE_FINISH_MAIN_LINE_TASK {
//			logger.Warn("条件类任务配置错误或者不能配任务完成", taskConf.Id)
//			return
//		}
//		if targetId != taskConf.Condition.K {
//			return
//		}
//		reachedValue, done := m.Condition.Check(user, &taskConf.Condition)
//		user.MainLineTask.Process = reachedValue
//		if done {
//			user.MainLineTask.Award = pb.TASKSTATUS_CAN_GOT
//		}
//	default:
//		logger.Warn("markMainLineTask:not implement task type %d", taskConf.TaskType)
//
//	}
//	//m.Condition.Mark(user, pb.CONDITIONTYPE_FINISH_MAIN_LINE_TASK)
//	user.MarkSyncStatus(objs.SyncStatusTaskInfoNtf)
//	user.Dirty = true
//}
//
//func (this *TaskManager) CheckMainLineTask(user *objs.User) {
//
//	//最后一个任务检查
//	taskConf := gameDb().GetTask(user.MainLineTask.TaskId)
//	if taskConf != nil && taskConf.NextTaskId == 0 && GameDb().InitConf.RestartMainlineTaskId > 0 {
//		user.MainLineTask.TaskId = GameDb().InitConf.RestartMainlineTaskId
//		user.MainLineTask.Process = 0
//		user.MainLineTask.Award = pb.TASKSTATUS_GOING_ON
//	}
//
//	//第一个主线任务，在配置里
//	if user.MainLineTask.TaskId <= 0 {
//		user.MainLineTask.TaskId = gameDb().GetConf().FirstMainLineTaskId
//		user.MainLineTask.Process = 0
//		user.MainLineTask.Award = pb.TASKSTATUS_GOING_ON
//		user.Dirty = true
//	} else if user.MainLineTask.Award == pb.TASKSTATUS_HAS_GOT {
//		this.GoNextMainLineTask(user)
//	} else if user.MainLineTask.Award == pb.TASKSTATUS_GOING_ON {
//		taskConf := gameDb().GetTask(user.MainLineTask.TaskId)
//		if taskConf != nil {
//			if taskConf.TaskType == pb.TASKTYPE_CONDITION && len(taskConf.TargetValue) == 2 {
//				this.MarkMainLineTask(user, pb.TASKTYPE_CONDITION, taskConf.TargetValue[0])
//			} else if taskConf.TaskType == pb.TASKTYPE_COLLECT {
//				//如果是收集类任务,要检查一下是不是完成了,有可能策划调整了,但是服务器记录的还是没有完成.
//				m.Task.Collect(user, 0)
//			}
//		}
//	}
//}
//
//func (this *TaskManager) GoNextMainLineTask(user *objs.User) {
//
//	taskConf := gameDb().GetTask(user.MainLineTask.TaskId)
//	if taskConf != nil {
//		nextId := taskConf.NextTaskId
//		if nextId > 0 {
//			nextTaskT := gameDb().GetTask(nextId)
//			if nextTaskT == nil {
//				logger.Error("next task id :%d,not exist", nextId)
//				return
//			}
//			if nextTaskT.NextTaskId <= 0 && GameDb().InitConf.RestartMainlineTaskId > 0 {
//				nextId = GameDb().InitConf.RestartMainlineTaskId
//				nextTaskT = gameDb().GetTask(nextId)
//				if nextTaskT == nil {
//					logger.Error("next task id :%d,not exist", nextId)
//					return
//				}
//			}
//			user.MainLineTask.TaskId = nextId
//			user.MainLineTask.Process = 0
//			user.MainLineTask.Award = pb.TASKSTATUS_GOING_ON
//
//			switch nextTaskT.TaskType {
//			case pb.TASKTYPE_ENTER_COPY:
//			case pb.TASKTYPE_CONDITION:
//				if len(nextTaskT.TargetValue) == 2 {
//					this.MarkMainLineTask(user, pb.TASKTYPE_CONDITION, nextTaskT.TargetValue[0])
//				}
//			case pb.TASKTYPE_AWARD:
//				user.MainLineTask.Process = 1
//				user.MainLineTask.Award = pb.TASKSTATUS_CAN_GOT
//			default:
//				user.MainLineTask.Award = pb.TASKSTATUS_GOING_ON
//			}
//
//		} else {
//			user.MainLineTask.Process = 0
//			user.MainLineTask.Award = pb.TASKSTATUS_HAS_GOT
//		}
//		user.Dirty = true
//	}
//
//}
