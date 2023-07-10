package rmodel

import "fmt"

const (
	first_drop_item_all_get_num = "first_drop_item_all_get_num:%d"
)

type FirstDropItemModel struct{}

var FirstDrop = &FirstDropItemModel{}

//首爆装备 全服人员领取次数限制
func (this *FirstDropItemModel) GetFirstDropItemGetKey(types int) string {
	return fmt.Sprintf(first_drop_item_all_get_num, types)
}
func (this *FirstDropItemModel) SetFirstDropItemGetNum(types, itemId, value int) {
	key := this.GetFirstDropItemGetKey(types)
	redisDb.HIncrBy(key, itemId, value)
}
func (this *FirstDropItemModel) GetFirstDropItemGetNum(types, ItemId int) int {
	key := this.GetFirstDropItemGetKey(types)
	v, _ := redisDb.HgetInt(key, ItemId)
	return v
}

func (this *FirstDropItemModel) GetFirstDropAllItemGetNum(types int) map[int]int {
	key := this.GetFirstDropItemGetKey(types)
	v, _ := redisDb.HgetallIntMap(key)
	return v
}
