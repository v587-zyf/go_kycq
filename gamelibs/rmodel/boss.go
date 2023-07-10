package rmodel

import "fmt"

const (
	Boss_Kill_Num = "boss_kill_num_%d_%d" //boss类型,stageId
	Boss_Owner    = "boss_owner_%d"       //stageId
)

type BossModel struct{}

var Boss = &BossModel{}

//boss击杀次数
func (this *BossModel) GetBossKillNumKey(bossType, bossId int) string {
	return fmt.Sprintf(Boss_Kill_Num, bossType, bossId)
}
func (this *BossModel) SetBossKillNum(userId, bossType, bossId, value int) {
	key := this.GetBossKillNumKey(bossType, bossId)
	redisDb.HIncrBy(key, userId, value)
}
func (this *BossModel) GetBossKillNum(userId, bossType, bossId int) int {
	key := this.GetBossKillNumKey(bossType, bossId)
	v, _ := redisDb.HgetInt(key, userId)
	return v
}
func (this *BossModel) DelBossKillNum(userId, bossType, bossId int) {
	key := this.GetBossKillNumKey(bossType, bossId)
	redisDb.HDel(key, userId)
}

//boss归属者
func (this *BossModel) GetBossOwnerKey(stageId int) string {
	return fmt.Sprintf(Boss_Owner, stageId)
}
func (this *BossModel) GetBossOwner(stageId int) string {
	key := this.GetBossOwnerKey(stageId)
	data, _ := redisDb.Get(key).String()
	return data
}
func (this *BossModel) SetBossOwner(stageId int, value string) {
	key := this.GetBossOwnerKey(stageId)
	redisDb.SetWithExpire(key, value, AutoExpireTime)
}
func (this *BossModel) DelBossOwner(stageId int) {
	key := this.GetBossOwnerKey(stageId)
	redisDb.Del(key)
}
