package rmodel

import (
	"fmt"
	"time"
)

const (
	Task_Msg_Log    = "task_msg_log_%d_%d" //得到玩家离线接收的私聊信息key,string(serverId, userId)
	Friend_HeroInfo = "friend_hero_%d"     //好友战力(userId)

	Friend_Apply     = "friend_apply_%d"     //好友申请(userId)
	Friend_Apply_Add = "friend_apply_add_%d" //好友添加(userId)
	Friend_Apply_Del = "friend_apply_del_%d" //好友删除(userId)

	AutoExpireTime = 15 * 24 * time.Hour
)

type FriendModel struct{}

var Friend = &FriendModel{}

//到玩家离线接收的私聊信息key
func (this *FriendModel) GetTaskMsgLogKey(serverId, userId int) string {
	return fmt.Sprintf(Task_Msg_Log, serverId, userId)
}
func (this *FriendModel) GetTaskMsgLog(serverId, userId int) string {
	key := this.GetTaskMsgLogKey(serverId, userId)
	msgLog, _ := redisDb.Get(key).String()
	return msgLog
}
func (this *FriendModel) SetTaskMsgLog(serverId, userId int, value string) {
	key := this.GetTaskMsgLogKey(serverId, userId)
	redisDb.SetWithExpire(key, value, AutoExpireTime)
}
func (this *FriendModel) DelTaskMsgLog(serverId, userId int) {
	key := this.GetTaskMsgLogKey(serverId, userId)
	redisDb.Del(key)
}

//好友装备和属性Key
func (this *FriendModel) GetFriendHeroInfoKey(userId int) string {
	return fmt.Sprintf(Friend_HeroInfo, userId)
}
func (this *FriendModel) SetFriendInfo(userId int, val string) {
	key := this.GetFriendHeroInfoKey(userId)
	redisDb.SetWithExpire(key, val, AutoExpireTime)
}
func (this *FriendModel) GetFriendInfo(userId int) []byte {
	key := this.GetFriendHeroInfoKey(userId)
	bytes, _ := redisDb.Get(key).Bytes()
	return bytes
}
func (this *FriendModel) DelFriendInfo(userId int) {
	key := this.GetFriendHeroInfoKey(userId)
	redisDb.Del(key)
}

/**
 *  @Description: 好友申请
 *  @param friendId	要添加的好友id
 *  @return string
 */
func (this *FriendModel) GetFriendApplyKey(friendId int) string {
	return fmt.Sprintf(Friend_Apply, friendId)
}
func (this *FriendModel) AddFriendApply(friendId, userId int) {
	key := this.GetFriendApplyKey(friendId)
	redisDb.SAdd(key, userId)
}
func (this *FriendModel) GetFriendApply(friendId int) []int {
	key := this.GetFriendApplyKey(friendId)
	data, _ := redisDb.SmembersInt(key)
	return data
}
func (this *FriendModel) DelFriendApply(friendId, userId int) {
	key := this.GetFriendApplyKey(friendId)
	redisDb.SRem(key, userId)
}
func (this *FriendModel) CheckFriendApply(friendId, userId int) bool {
	key := this.GetFriendApplyKey(friendId)
	has, _ := redisDb.SIsMember(key, userId)
	return has
}

/**
 *  @Description: 添加好友，离线状态
 *  @param userId
 *  @return string
 */
func (this *FriendModel) GetFriendApplyAddKey(userId int) string {
	return fmt.Sprintf(Friend_Apply_Add, userId)
}
func (this *FriendModel) SetFriendApplyAdd(userId int, val string) {
	key := this.GetFriendApplyAddKey(userId)
	redisDb.Set(key, val)
	redisDb.Expire(key, AutoExpireTime)
}
func (this *FriendModel) GetFriendApplyAdd(userId int) string {
	key := this.GetFriendApplyAddKey(userId)
	val, _ := redisDb.Get(key).String()
	return val
}
func (this *FriendModel) DelFriendApplyAdd(userId int) {
	key := this.GetFriendApplyAddKey(userId)
	redisDb.Del(key)
}

/**
 *  @Description: 删除好友，离线状态
 *  @param userId
 *  @return string
 */
func (this *FriendModel) GetFriendApplyDelKey(userId int) string {
	return fmt.Sprintf(Friend_Apply_Add, userId)
}
func (this *FriendModel) SetFriendApplyDel(userId int, val string) {
	key := this.GetFriendApplyDelKey(userId)
	redisDb.Set(key, val)
	redisDb.Expire(key, AutoExpireTime)
}
func (this *FriendModel) GetFriendApplyDel(userId int) string {
	key := this.GetFriendApplyDelKey(userId)
	val, _ := redisDb.Get(key).String()
	return val
}
func (this *FriendModel) DelFriendApplyDel(userId int) {
	key := this.GetFriendApplyDelKey(userId)
	redisDb.Del(key)
}
