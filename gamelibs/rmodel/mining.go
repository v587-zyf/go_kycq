package rmodel

import (
	"cqserver/golibs/logger"
	"fmt"
	"time"
)

const (
	MiningWork  = "mining_work_%d"  //正在挖矿的玩家(ServerId)
	MiningFight = "mining_fight:%d" //挖矿掠夺中
)

type MiningModel struct {
}

var Mining = &MiningModel{}

//getMineWorkKey 得到玩家挖矿 key
func (this *MiningModel) GetMiningWorkKey(serverId int) string {
	return fmt.Sprintf(MiningWork, serverId)
}

// 获取挖矿玩家总人数
func (this *MiningModel) ZcardMiningWork(serverId int) int {
	key := this.GetMiningWorkKey(serverId)
	return redisDb.ZCard(key)
}

// 获取所有挖矿玩家（从小到大排序）
func (this *MiningModel) ZrevrangeMiningWork(serverId int) []int {
	key := this.GetMiningWorkKey(serverId)
	values, _ := redisDb.ZRevrange(key, 0, this.ZcardMiningWork(serverId)).ValuesIntSlice()
	return values
}

// 添加挖矿玩家
func (this *MiningModel) ZaddMiningWork(userId, serverId, expire int) {
	key := this.GetMiningWorkKey(serverId)
	redisDb.ZAdd(key, expire, userId)
}

// 删除挖矿玩家
func (this *MiningModel) ZremMiningWork(userId, serverId int) {
	key := this.GetMiningWorkKey(serverId)
	redisDb.ZRem(key, userId)
}

// 获取挖矿玩家失效时间
func (this *MiningModel) ZscoreMiningWork(userId, serverId int) int {
	key := this.GetMiningWorkKey(serverId)
	score, _ := redisDb.ZScore(key, userId)
	return score
}

func (this *MiningModel) HsetMiningWork(userId, serverId, expire int) {
	key := this.GetMiningWorkKey(serverId)
	err := redisDb.HSet(key, userId, expire)
	if err != nil {
		logger.Error("HsetMiningWork HSet err:%v", err)
	}
}
func (this *MiningModel) HdelMiningWork(userId, serverId int) {
	key := this.GetMiningWorkKey(serverId)
	_, err := redisDb.HDel(key, userId)
	if err != nil {
		logger.Error("HdelMiningWork err:%v", err)
	}
}
func (this *MiningModel) HgetAllMiningWork(serverId int) map[string]string {
	key := this.GetMiningWorkKey(serverId)
	values, err := redisDb.HgetAll(key).HashValues()
	if err != nil {
		logger.Error("GetMining err:%v", err)
	}
	return values
}

func (this *MiningModel) SetMiningFight(miningId int, fightUserId int) {
	key := fmt.Sprintf(MiningFight, miningId)
	redisDb.SetWithExpire(key, fightUserId, 20*time.Minute)
}
func (this *MiningModel) GetMiningFight(miningId int) (int, error) {
	key := fmt.Sprintf(MiningFight, miningId)
	return redisDb.Get(key).Int()
}
func (this *MiningModel) DelMiningFight(miningId int) {
	key := fmt.Sprintf(MiningFight, miningId)
	redisDb.Del(key)
}
