package rmodel

import (
	"fmt"
)

const (
	WARORDER_TYPE_NUM = "warOrder_type_num_%d_%d"	//玩家下线后 战令任务完成记录（任务类型，完成次数）
)

type WarOrderModel struct {
}

var WarOrder = &WarOrderModel{}

func (this *WarOrderModel) GetTaskKey(t, userId int) string {
	return fmt.Sprintf(WARORDER_TYPE_NUM, t, userId)
}

func (this *WarOrderModel) GetTask(t, userId int) int {
	key := this.GetTaskKey(t, userId)
	v, _ := redisDb.Get(key).Int()
	return v
}

func (this *WarOrderModel) SetTask(t, userId, value int) {
	key := this.GetTaskKey(t, userId)
	redisDb.SetWithExpire(key, value, AutoExpireTime)
}

func (this *WarOrderModel) DelTask(t, userId int) {
	key := this.GetTaskKey(t, userId)
	redisDb.Del(key)
}