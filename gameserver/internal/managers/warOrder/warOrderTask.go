package warOrder

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

/**
 *  @Description: 战令任务直接完成
 *  @param user
 *  @param op
 *  @param req
 *  @param ack
 *  @return error
 */
func (this *WarOrderManager) TaskFinish(user *objs.User, op *ophelper.OpBagHelperDefault, req *pb.WarOrderTaskFinishReq, ack *pb.WarOrderTaskFinishAck) error {
	userWarOrder := user.WarOrder
	taskId := int(req.TaskId)
	week := int(req.Week)
	if taskId < 1 {
		return gamedb.ERRPARAM
	}
	if req.IsWeekTask {
		if week < 0 {
			return gamedb.ERRPARAM
		}
		nowWeek := this.GetNowWeek(userWarOrder)
		if week > nowWeek {
			return gamedb.ERRWARORDERWEEK
		}
		if userWarOrder.WeekTask[week] == nil {
			userWarOrder.WeekTask[week] = make(map[int]*model.WarOrderTask)
		}
		taskId = taskId % constConstant.COMPUTE_TEN_THOUSAND
		warOrderTask, ok := userWarOrder.WeekTask[week][taskId]
		if !ok {
			userWarOrder.WeekTask[week][taskId] = &model.WarOrderTask{}
			warOrderTask = userWarOrder.WeekTask[week][taskId]
		}
		weekTaskCfg := gamedb.GetWarOrderWeekTaskWarOrderWeekTaskCfg(gamedb.GetRealId(week, taskId))
		if err := this.GetBag().Remove(user, op, weekTaskCfg.TaskCard.ItemId, weekTaskCfg.TaskCard.Count); err != nil {
			return gamedb.ERRNOTENOUGHGOODS
		}
		if warOrderTask.Finish {
			return gamedb.ERRTASKOK
		}
		warOrderTask.Finish = true
	} else {
		warOrderTask, ok := userWarOrder.Task[taskId]
		if !ok {
			userWarOrder.Task[taskId] = &model.WarOrderTask{}
			warOrderTask = userWarOrder.Task[taskId]
		}
		taskCfg := gamedb.GetWarOrderCycleTaskWarOrderCycleTaskCfg(taskId)
		if err := this.GetBag().Remove(user, op, taskCfg.TaskCard.ItemId, taskCfg.TaskCard.Count); err != nil {
			return gamedb.ERRNOTENOUGHGOODS
		}
		warOrderTask.Finish = true
	}
	user.Dirty = true

	ack.TaskId = req.TaskId
	ack.Week = req.Week
	ack.IsWeekTask = req.IsWeekTask
	return nil
}

/**
 *  @Description: 战令任务领取奖励
 *  @param user
 *  @param op
 *  @param req
 *  @param ack
 *  @return error
 */
func (this *WarOrderManager) TaskReward(user *objs.User, op *ophelper.OpBagHelperDefault, req *pb.WarOrderTaskRewardReq, ack *pb.WarOrderTaskRewardAck) error {
	userWarOrder := user.WarOrder
	taskId := int(req.TaskId)
	week := int(req.Week)
	if taskId < 1 {
		return gamedb.ERRPARAM
	}
	if req.IsWeekTask {
		if week == 0 {
			return gamedb.ERRPARAM
		}
		nowWeek := this.GetNowWeek(userWarOrder)
		if week > nowWeek {
			return gamedb.ERRWARORDERWEEK
		}
		if userWarOrder.WeekTask[week] == nil {
			userWarOrder.WeekTask[week] = make(map[int]*model.WarOrderTask)
		}
		warOrderTask, ok := userWarOrder.WeekTask[week][taskId]
		if !ok || !warOrderTask.Finish {
			return gamedb.ERRTASKISNOTOVER
		}
		if warOrderTask.Reward {
			return gamedb.ERRREPEATRECEIVE
		}
		weekTaskCfg := gamedb.GetWarOrderWeekTaskWarOrderWeekTaskCfg(gamedb.GetRealId(week, taskId))
		this.GetBag().Add(user, op, weekTaskCfg.Reward.ItemId, weekTaskCfg.Reward.Count)
		warOrderTask.Reward = true
	} else {
		warOrderTask, ok := userWarOrder.Task[taskId]
		if !ok || !warOrderTask.Finish {
			return gamedb.ERRTASKISNOTOVER
		}
		if warOrderTask.Reward {
			return gamedb.ERRREPEATRECEIVE
		}
		taskCfg := gamedb.GetWarOrderCycleTaskWarOrderCycleTaskCfg(taskId)
		this.GetBag().Add(user, op, taskCfg.Reward.ItemId, taskCfg.Reward.Count)
		warOrderTask.Reward = true
	}
	this.AutoUpLv(user, userWarOrder, false)
	ack.IsWeekTask = req.IsWeekTask
	ack.Week = req.Week
	ack.TaskId = req.TaskId
	ack.Lv = int32(userWarOrder.Lv)
	ack.Exp = int32(userWarOrder.Exp)
	return nil
}
