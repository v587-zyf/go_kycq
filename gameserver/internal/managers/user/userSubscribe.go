package user

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/ptsdk"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"time"
)

func (this *UserManager) Subscribe(user *objs.User, subscribeId int) error {

	//判断玩家是否已经订阅
	if rmodel.User.UserSubscribeCheck(user.Id, subscribeId) {
		return nil
	}

	//记录玩家订阅
	err := rmodel.User.UserSubscribe(user.Id, subscribeId)
	return err
}

func (this *UserManager) GetSubscribe(userId int) []int32 {

	subscribe := make([]int32, 0)
	for _, v := range pb.SUBSCRIBE_ARRAY {
		if rmodel.User.UserSubscribeCheck(userId, v) {
			subscribe = append(subscribe, int32(v))
		}
	}
	return subscribe
}

func (this *UserManager) CronSubscribe() {

	this.usersMu.Lock()
	now := time.Now()

	for _, subscribeId := range pb.SUBSCRIBE_ARRAY {

		allUser := rmodel.User.GetSubscribeUser(subscribeId)
		removeUser := make([]int, 0)
		if allUser != nil && len(allUser) > 0 {

			if subscribeId == pb.SUBSCRIBE_HOOK {
				offlineTimeLimit := float64(gamedb.GetConf().Weixinhookmap3 * 3600)
				for _, userId := range allUser {
					if _, ok := this.users[userId]; ok {
						continue
					}
					if u, ok := this.usersBasicInfoMap[userId]; ok {
						if now.Sub(u.LastUpdateTime).Seconds() > offlineTimeLimit {
							removeUser = append(removeUser, userId)
							ptsdk.GetSdk().Subscribe(u.OpenId, subscribeId, gamedb.GetConf().Weixinhookmap1, gamedb.GetConf().Weixinhookmap2)
						}
					} else {
						logger.Error("玩家订阅消息:%v，内存为找到玩家数据:%v", subscribeId, userId)
					}
				}
			}
		}
		if len(removeUser) > 0 {
			rmodel.User.UserSubscribeRemove(subscribeId, removeUser...)
		}
	}
	this.usersMu.Unlock()
}
