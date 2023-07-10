package managersI

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type ITaskManager interface {
	/**
	 *  @Description: 初始化新玩家任务
	 *  @param user
	 */
	InitNewUserTask(user *objs.User)

	/**
	 *  @Description:	击杀怪物更新任务进度
	 *  @param user
	 *  @param monsterId
	 *  @param num
	 */
	UpdateTaskForKillMonster(user *objs.User, monsterId, num int)
	/**
	 *  @Description:	增加任务进度
	 *  @param user
	 *  @param condition	条件类型
	 *  @param process		增加值
	 */
	AddTaskProcess(user *objs.User, condition int, process int)
	/**
	 *  @Description:	更新任务进度（用于任务自己往玩家身上获取数据的）
	 *  @param user
	 */
	UpdateTaskProcess(user *objs.User, sendclient, isTaskDone bool)

	/**
	 *  @Description:完成任务 领取奖励
	 *  @param user
	 *  @param op
	 */
	TaskDone(user *objs.User, op *ophelper.OpBagHelperDefault) error

	/**
	 *  @Description: 获取战斗进入区域
	 *  @param user
	 *  @param stageId
	 *  @return int
	 */
	GetFightEnterArea(user *objs.User, stageId int) int

	//特殊任务检查
	SpecialCheck(user *objs.User, conditionType int, conditionValue gamedb.IntSlice) (bool, int)

	SendSpecialCheckTaskInfo(user *objs.User, conditionType int)
}
