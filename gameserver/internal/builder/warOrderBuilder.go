package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildWarOrder(user *objs.User) *pb.WarOrder {
	userWarOrder := user.WarOrder
	return &pb.WarOrder{
		Lv:        int32(userWarOrder.Lv),
		Exp:       int32(userWarOrder.Exp),
		Season:    int32(userWarOrder.Season),
		StartTime: int64(userWarOrder.StartTime),
		EndTime:   int64(userWarOrder.EndTIme),
		IsLuxury:  userWarOrder.IsLuxury,
		Task:      &pb.WarOrderTask{Task: BuildWarOrderTask(userWarOrder.Task)},
		Exchange:  BuildWarOrderExchange(userWarOrder.Exchange),
		WeekTask:  BuildWarOrderWeekTask(userWarOrder.WeekTask),
		Reward:    BuildWarOrderReward(userWarOrder.Reward),
	}
}

func BuildWarOrderExchange(data model.IntKv) map[int32]int32 {
	pbMap := make(map[int32]int32)
	for k, v := range data {
		pbMap[int32(k)] = int32(v)
	}
	return pbMap
}

func BuildWarOrderTask(task map[int]*model.WarOrderTask) map[int32]*pb.WarOrderTaskInfo {
	pbMap := make(map[int32]*pb.WarOrderTaskInfo)
	for k, t := range task {
		pbMap[int32(k)] = BuildWarOrderInfo(t)
	}
	return pbMap
}

func BuildWarOrderInfo(task *model.WarOrderTask) *pb.WarOrderTaskInfo {
	return &pb.WarOrderTaskInfo{
		Val:    BuildWarOrderUnit(task.Val),
		Finish: task.Finish,
		Reward: task.Reward,
	}
}

func BuildWarOrderUnit(val model.WarOrderTaskUnit) *pb.WarOrderTaskUnit {
	return &pb.WarOrderTaskUnit{
		One: int32(val.One),
		Two: BuildWarOrderUnitTwo(val.Two),
		Three: BuildWarOrderUnitTwo(val.Three),
	}
}

func BuildWarOrderUnitTwo(two model.IntKv) map[int32]int32 {
	pbMap := make(map[int32]int32)
	for k, v := range two {
		pbMap[int32(k)] = int32(v)
	}
	return pbMap
}

func BuildWarOrderWeekTask(weekTask map[int]map[int]*model.WarOrderTask) map[int32]*pb.WarOrderTask {
	pbMap := make(map[int32]*pb.WarOrderTask)
	for week, t := range weekTask {
		pbMap[int32(week)] = &pb.WarOrderTask{Task: BuildWarOrderTask(t)}
	}
	return pbMap
}

func BuildWarOrderReward(reward map[int]*model.WarOrderReward) map[int32]*pb.WarOrderReward {
	pbMap := make(map[int32]*pb.WarOrderReward)
	for id, orderReward := range reward {
		pbMap[int32(id)] = &pb.WarOrderReward{
			Elite:  orderReward.Elite,
			Luxury: orderReward.Luxury,
		}
	}
	return pbMap
}
