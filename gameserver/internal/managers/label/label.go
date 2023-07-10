package label

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

func NewLabelManager(module managersI.IModule) *LabelManager {
	return &LabelManager{IModule: module}
}

type LabelManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *LabelManager) Online(user *objs.User) {
	date := common.GetResetTime(time.Now())
	this.ResetLabel(user, date)
}

func (this *LabelManager) ResetLabel(user *objs.User, date int) {
	userLabel := user.Label
	if userLabel.Id == 0 {
		userLabel.Id = 1
	}
	if userLabel.RefTime != date {
		userLabel.RefTime = date
		userLabel.DayReward = false
		userLabel.Transfer = 0
	}
}

/**
 *  @Description: 升级
 *  @param user
 *  @param op
 *  @return error
 */
func (this *LabelManager) Up(user *objs.User, op *ophelper.OpBagHelperDefault) error {
	userLabel := user.Label
	if gamedb.GetLabelLabelCfg(userLabel.Id+1) == nil {
		return gamedb.ERRLVENOUGH
	}
	labelCfg := gamedb.GetLabelLabelCfg(userLabel.Id)
	if labelCfg == nil {
		return gamedb.ERRPARAM
	}
	for _, tid := range labelCfg.AllCondition {
		if _, ok := userLabel.TaskOver[tid]; !ok {
			taskSlice := gamedb.GetLabelTaskConditionSlice(tid)
			if r, check := this.GetCondition().CheckBySlice(user, -1, taskSlice); !check {
				logger.Debug("conditionId:%v val:%v result:%v", taskSlice[0], taskSlice[1:], r)
				return gamedb.ERRCONDITION
			}
		}
	}
	this.GetBag().AddItems(user, gamedb.GetLabelLabelCfg(userLabel.Id+1).Reward, op)

	userLabel.Id++
	userLabel.TaskOver = make(model.IntKv)

	this.GetUserManager().SendDisplay(user)
	this.GetUserManager().UpdateCombat(user, -1)
	this.SendLabelTaskNtf(user, -1)
	this.GetTask().UpdateTaskProcess(user, false, false)
	return nil
}

/**
 *  @Description: 转职
 *  @param user
 *  @param job
 *  @return error
 */
func (this *LabelManager) Transfer(user *objs.User, job int) error {
	userLabel := user.Label
	if userLabel.Transfer > 0 {
		return gamedb.ERRNOTENOUGHTIMES
	}
	if userLabel.Job == job {
		return gamedb.ERRPARAM
	}

	userLabel.Job = job
	if userLabel.FirstTransfer {
		userLabel.Transfer++
	}
	userLabel.FirstTransfer = true
	this.GetUserManager().SendDisplay(user)
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

/**
 *  @Description: 领取奖励
 *  @param user
 *  @param op
 *  @return error
 */
func (this *LabelManager) DayReward(user *objs.User, op *ophelper.OpBagHelperDefault) error {
	userLabel := user.Label
	if userLabel.DayReward {
		return gamedb.ERRREPEATRECEIVE
	}
	labelCfg := gamedb.GetLabelLabelCfg(userLabel.Id)
	if labelCfg == nil {
		return gamedb.ERRPARAM
	}
	this.GetBag().AddItems(user, labelCfg.DailyReward, op)

	userLabel.DayReward = true
	user.Dirty = true
	return nil
}

/**
 *  @Description: 推送任务进度
 *  @param user
 *  @param conditionId
 */
func (this *LabelManager) SendLabelTaskNtf(user *objs.User, conditionId int) {
	userLabel := user.Label
	labelCfg := gamedb.GetLabelLabelCfg(userLabel.Id)
	if labelCfg != nil {
		if len(labelCfg.AllCondition) < 1 {
			this.GetUserManager().SendMessage(user, &pb.LabelTaskNtf{}, true)
		}
		isSend := false
		pbMap := make(map[int32]*pb.LabelTaskUnit)
		for _, tid := range labelCfg.AllCondition {
			taskSlice := gamedb.GetLabelTaskConditionSlice(tid)
			cid := taskSlice[0]
			if conditionId == cid || conditionId == -1 {
				isSend = true
			}
			isOver := false
			valSlice := make([]int32, 0)
			cfgVal := make([]int32, 0)
			checkVal, flag := this.GetCondition().CheckBySlice(user, -1, taskSlice)
			if _, ok := userLabel.TaskOver[tid]; ok {
				isOver = true
			} else {
				if flag {
					isOver = true
					userLabel.TaskOver[tid] = 0
				}
			}
			for _, v := range checkVal {
				valSlice = append(valSlice, int32(v))
			}
			for k, v := range taskSlice {
				if k == 0 {
					continue
				}
				cfgVal = append(cfgVal, int32(v))
			}
			pbMap[int32(tid)] = &pb.LabelTaskUnit{
				TaskId: int32(tid),
				Value:  valSlice,
				CfgVal: cfgVal,
				IsOver: isOver,
			}
		}
		if isSend {
			ntf := &pb.LabelTaskNtf{LabelId: int32(userLabel.Id), TaskInfo: pbMap}
			this.GetUserManager().SendMessage(user, ntf, false)
		}
	}
}
