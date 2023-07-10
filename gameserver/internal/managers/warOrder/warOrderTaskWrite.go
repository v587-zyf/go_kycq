package warOrder

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/protobuf/pb"
	"math"
	"time"
)

/**
 *  @Description: 记录杀怪战令任务
 *  @param user
 *  @param monsterId
 *  @param killNum
 */
func (this *WarOrderManager) WriteWarOrderTaskByKillMonster(user *objs.User, monsterId, killNum int) {
	monsterCfg := gamedb.GetMonsterMonsterCfg(monsterId)
	if monsterCfg.Type == constFight.MONSTER_TYPE_NORMAL {
		this.WriteWarOrderTask(user, pb.WARORDERCONDITION_KILL_MONSTERAI_ONE, []int{killNum})
	} else {
		this.WriteWarOrderTask(user, pb.WARORDERCONDITION_KILL_MONSTERAI_TWO, []int{killNum})
	}
}

/**
 *  @Description: 记录战令任务
 *  @param t	类型
 *  @param val	增加数量
 */
func (this *WarOrderManager) WriteWarOrderTask(user *objs.User, t int, val []int) {
	this.WriteMu.Lock()
	defer this.WriteMu.Unlock()

	userWarOrder := user.WarOrder
	//循环这个类型所有任务
	//周期任务
	flag := false
	taskCfgs := gamedb.GetWarOrderTaskByConditionType(t)
	date := common.GetResetTime(time.Now())
	for _, cfg := range taskCfgs {
		task, ok := userWarOrder.Task[cfg.Id]
		if !ok {
			userWarOrder.Task[cfg.Id] = &model.WarOrderTask{}
			task = userWarOrder.Task[cfg.Id]
		}
		if task.Finish {
			continue
		}
		this.WriteTask(userWarOrder, t, date, task, val, cfg.ConditionValue)
		if this.CheckTask(t, task, cfg.ConditionValue) {
			flag = true
		}
	}
	//周任务
	weekTaskCfgs := gamedb.GetWarOrderWeekTaskByConditionType(t)
	for _, cfg := range weekTaskCfgs {
		week := cfg.Id / constConstant.COMPUTE_TEN_THOUSAND
		id := cfg.Id % constConstant.COMPUTE_TEN_THOUSAND
		nowWeek := this.GetNowWeek(userWarOrder)
		if week > nowWeek || (t == pb.WARORDERCONDITION_WEEK_LOGIN_DAY && week != nowWeek) {
			continue
		}
		tasks, ok := userWarOrder.WeekTask[week]
		if !ok {
			userWarOrder.WeekTask[week] = make(map[int]*model.WarOrderTask)
			tasks = userWarOrder.WeekTask[week]
		}
		task, ok := tasks[id]
		if !ok {
			tasks[id] = &model.WarOrderTask{}
			task = tasks[id]
		}
		if task.Finish {
			continue
		}
		this.WriteTask(userWarOrder, t, date, task, val, cfg.ConditionValue)
		if this.CheckTask(t, task, cfg.ConditionValue) {
			flag = true
		}
	}
	user.Dirty = true
	if flag {
		this.GetUserManager().SendMessage(user, &pb.WarOrderTaskNtf{
			Task:     &pb.WarOrderTask{Task: builder.BuildWarOrderTask(userWarOrder.Task)},
			WeekTask: builder.BuildWarOrderWeekTask(userWarOrder.WeekTask),
			Lv:       int32(userWarOrder.Lv),
			Exp:      int32(userWarOrder.Exp),
		}, true)
	}
}

func (this *WarOrderManager) WriteTask(userWarOrder *model.WarOrder, t, date int, task *model.WarOrderTask, val, conditionV []int) bool {
	//特殊类型，每日只记录一次
	if task.Date == nil {
		task.Date = make(model.IntKv)
	}
	switch t {
	case pb.WARORDERCONDITION_SHABAKE_NUM:
		fallthrough
	case pb.WARORDERCONDITION_PAODIAN_NUM:
		fallthrough
	case pb.WARORDERCONDITION_GUILDBONFIRE_NUM:
		fallthrough
	case pb.WARORDERCONDITION_WEEK_LOGIN_DAY:
		if task.Date[t] == date {
			return false
		}
	}

	if task.Val.Two == nil {
		task.Val.Two = make(model.IntKv)
	}
	switch t {
	case pb.WARORDERCONDITION_KILL_MONSTER:
		k, v := val[1], val[0]
		if k != conditionV[1] {
			return false
		}
		task.Val.Two[k] += v
	case pb.WARORDERCONDITION_WEEK_LOGIN_DAY:
		week := this.GetNowWeek(userWarOrder)
		task.Val.Two[week] += val[0]
	case pb.WARORDERCONDITION_BUY_SHOP_GOODS:
		fallthrough
	case pb.WARORDERCONDITION_COMPLETE_COPY:
		k, v := val[1], val[0]
		if k != conditionV[1] {
			return false
		}
		task.Val.Two[k] += v
	case pb.WARORDERCONDITION_GET_ITEM:
		if val[1] != conditionV[1] {
			return false
		}
		if task.Val.Three == nil {
			task.Val.Three = make(model.IntKv)
		}
		k, v := val[1], val[0]
		task.Val.Three[k] += v
		if v < conditionV[0]/50 {
			return false
		}
	default:
		task.Val.One += val[0]
	}
	task.Date[t] = date
	return true
}

func (this *WarOrderManager) CheckTask(t int, task *model.WarOrderTask, conditionV []int) bool {
	flag := false
	switch t {
	case pb.WARORDERCONDITION_KILL_MONSTER:
		fallthrough
	case pb.WARORDERCONDITION_BUY_SHOP_GOODS:
		fallthrough
	case pb.WARORDERCONDITION_COMPLETE_COPY:
		k, v := conditionV[1], conditionV[0]
		if n, ok := task.Val.Two[k]; ok && n >= v {
			flag = true
		}
	case pb.WARORDERCONDITION_GET_ITEM:
		k, v := conditionV[1], conditionV[0]
		if n, ok := task.Val.Three[k]; ok && n >= v {
			flag = true
		}
	case pb.WARORDERCONDITION_WEEK_LOGIN_DAY:
		k, v := conditionV[0], conditionV[1]
		if n, ok := task.Val.Two[k]; ok && n >= v {
			flag = true
		}
	case pb.WARORDERCONDITION_ONLINE:
		if int(math.Ceil(float64(task.Val.One)/float64(60))) >= conditionV[0] {
			flag = true
		}
	default:
		if task.Val.One >= conditionV[0] {
			flag = true
		}
	}
	if flag {
		task.Finish = true
	}
	return flag
}
