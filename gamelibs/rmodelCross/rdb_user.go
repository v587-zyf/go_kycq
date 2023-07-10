package rmodelCross

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/golibs/logger"
	"time"
	"encoding/json"
	"fmt"
)

const (
	cross_user_info              = "cross_user_base_info:%d"            //跨服所有玩家基础信息
	cross_recharge_num_by_server = "cross_recharge_num_by_server:%d:%d" //跨服所有玩家基础信息
)

type UserCrossRmodel struct {
}

var userCrossRmodel = &UserCrossRmodel{}

func GetUserCrossInfoRmodle() *UserCrossRmodel {
	return userCrossRmodel
}

func (this *UserCrossRmodel) Set(userInfo *modelCross.UserCrossInfo) error {

	key := fmt.Sprintf(cross_user_info, userInfo.UserId)
	bytes, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println("set key error:", key, err)
		return err
	}

	err = redisMap[CROSS_USER_REDIS].SetWithExpire(key, bytes, 7*24*time.Hour).Err
	if err != nil {
		fmt.Println("cross user set user error userId:%v,err", userInfo.UserId, err)
	}
	return err
}

func (this *UserCrossRmodel) Get(userId int) *modelCross.UserCrossInfo {

	key := fmt.Sprintf(cross_user_info, userId)
	bytes, err := redisMap[CROSS_USER_REDIS].GetByJson(key)
	if err != nil {
		fmt.Printf("rmodel.GetObj = %+v\n", err)
		return nil
	}
	if len(bytes) == 0 {
		userInfo, err1 := modelCross.GetUserCrossInfoModel().GetUserInfo(userId)
		if err1 != nil || userInfo == nil {
			fmt.Println("crossUser get err,userid:%v,error:%v", userId, err)
			return nil
		}
		this.Set(userInfo)
		logger.Info("DB userId:%d, name:%s, combat:%d", userInfo.UserId, userInfo.NickName, userInfo.Combat)
		return userInfo
	} else {
		var userInfo modelCross.UserCrossInfo
		err = json.Unmarshal(bytes, &userInfo)
		if err != nil {
			fmt.Printf("rmodel.GetObj:Unmarshal = %+v\n", err)
			return nil
		}
		logger.Info("redis userId:%d, name:%s, combat:%d", userInfo.UserId, userInfo.NickName, userInfo.Combat)
		return &userInfo
	}
}

func (this *UserCrossRmodel) SetDayRechargeNum(serverId, num int32) {

	key := fmt.Sprintf(cross_recharge_num_by_server, time.Now().Day(), serverId)

	redisMap[CROSS_USER_REDIS].IncrBy(key, int(num))
	redisMap[CROSS_USER_REDIS].Expire(key, 3*24*time.Hour)
}

func (this *UserCrossRmodel) GetDayRechargeNumByServerId(serverId int) int {

	key := fmt.Sprintf(cross_recharge_num_by_server, time.Now().Day()-1, serverId)
	num, _ := redisMap[CROSS_USER_REDIS].Get(key).IntDef(0)
	return num
}
