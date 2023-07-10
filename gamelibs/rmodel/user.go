package rmodel

import (
	"cqserver/golibs/logger"
	"fmt"
	"time"

	"cqserver/golibs/redisdb"
)

const (
	UserOnline                 = "user_online:%d"             //得到玩家的最后一次上线的时间key,string,(userId)
	UserLastLoginTime          = "last_login_time:%s:%d"      //得到玩家上次登录时间key,string,(yyMMdd,userId)
	TaskOpenTime               = "task_open_time:%d"          //得到玩家任务开启时间 key,hash,{userId},(index)
	UserDailyOnlineSeconds     = "user_daily_ol_%s_%d"        //玩家当天在线的秒数string,{yymmdd,userId}
	UserOfflineAvaliabeSeconds = "user_offline_avb_secs_%d"   //玩家当前用的离线的秒数string,{userId}
	UserOfflineAvaliabeResetAt = "user_offline_avb_restAt_%d" //玩家离线可用时间重置点string,{userId}
	UserOfflineAvaliabeMax     = "user_offline_avb_max_%d"    //玩家离线可用时间增加后的最大值string,{userId}
	UserFightDisConnectAward   = "user_dight_disc_award:%d"   //玩家战斗异常掉线奖励保存hash,{userId},(fightType)
	USER_LOGIN_DAY             = "user_login_day:%d:%d"       //userID:day  玩家在开服的这一天是否登录过
	USER_DIE_TIMES             = "user_die_times:%d"          //玩家死亡时间集合
	USER_DIE_DROP_ITEM_TIME    = "user_die_drop_item_time:%d" //玩家死亡掉落时间
	SendHour5MailState         = "send_hour_5_mail_state"
	UserSubscribe              = "user_subscribe_%d" //订阅玩家
)

type UserModel struct {
}

var User = &UserModel{}

//getOnLineAtKey 得到玩家的最后一次上线的时间
// string ts seconds
func (this *UserModel) getOnLineAtKey(userId int) string {
	return fmt.Sprintf(UserOnline, userId)
}

//GetUserLasTimesKey 得到玩家上次登录时间key
func (this *UserModel) getLasLoginTimesKey(userId int) string {
	return fmt.Sprintf(UserLastLoginTime, time.Now().Format("060102"), userId)
}

func (this *UserModel) GetFightDiscAwardKey(userId int) string {
	return fmt.Sprintf(UserFightDisConnectAward, userId)
}

//getTaskOpenTimeKey 得到玩家任务开启时间 key
func (this *UserModel) getTaskOpenTimeKey(userId int) string {
	return fmt.Sprintf(TaskOpenTime, userId)
}

//GetUserLasTimesKey 得到玩家上次登录时间key
func (this *UserModel) GetLasLoginTime(userId int) int {
	key := this.getLasLoginTimesKey(userId)
	ts, _ := redisDb.Get(key).Int()
	return ts
}

func (this *UserModel) SetLasLoginTime(userId int) {
	key := this.getLasLoginTimesKey(userId)
	redisDb.SetWithExpireSecond(key, 1, 24*60*60)
}

func (this *UserModel) HsetTaskOpenTime(userId int, index string, openTime int32) {
	key := this.getTaskOpenTimeKey(userId)
	redisDb.Hmset(key, index, openTime)
}

func (this *UserModel) ExpireTaskOpenTime(userId int, deadLine int) {
	key := this.getTaskOpenTimeKey(userId)
	redisDb.ExpireSecond(key, deadLine)
}

func (this *UserModel) HGetTaskOpenTime(userId int) (map[string]string, error) {
	key := this.getTaskOpenTimeKey(userId)
	return redisDb.Hgetall(key)
}

func (this *UserModel) GetOnLineAt(userId int) (int64, error) {
	key := this.getOnLineAtKey(userId)
	return redisDb.Get(key).Int64()
}

func (this *UserModel) SetOnLineAt(userId, deadLine int) *redisdb.RedisReply {
	key := this.getOnLineAtKey(userId)
	return redisDb.SetWithExpireSecond(key, time.Now().Unix(), deadLine)
}

func (this *UserModel) IncrDailyOnlineSeconds(day string, userId, delta int) (newOnlineSeconds int, err error) {
	key := fmt.Sprintf(UserDailyOnlineSeconds, day, userId)
	newOnlineSeconds, err = redisDb.IncrBy(key, delta)
	if err != nil {
		return
	}
	redisDb.Expire(key, 2*24*time.Hour)
	return
}

func (this *UserModel) IncrOfflineAvaliableSeconds(userId, delta, max int) (newValue int, err error) {
	key := fmt.Sprintf(UserOfflineAvaliabeSeconds, userId)
	newValue, err = redisDb.IncrBy(key, delta)
	if err != nil {
		return
	}
	if newValue > max {
		err = redisDb.Set(key, max).Err
		if err != nil {
			return
		}
		newValue = max
	}

	//如果是增加值,记录当前重置的时间点,和最大值新值
	if delta > 0 {
		err = this.SetOfflineAvaliableResetTsAndMax(userId, newValue)
	}

	redisDb.Expire(key, 4*24*time.Hour)
	return
}

func (this *UserModel) GetOfflineAvaliableSeconds(userId int) (int, error) {
	key := fmt.Sprintf(UserOfflineAvaliabeSeconds, userId)
	return redisDb.Get(key).IntDef(0)
}

func (this *UserModel) SetOfflineAvaliableResetTsAndMax(userId, newValue int) error {
	resetAtKey := fmt.Sprintf(UserOfflineAvaliabeResetAt, userId)
	maxKey := fmt.Sprintf(UserOfflineAvaliabeMax, userId)
	err := redisDb.MSet(resetAtKey, time.Now().Unix(), maxKey, newValue).Err
	redisDb.Expire(resetAtKey, 4*24*time.Hour)
	redisDb.Expire(maxKey, 4*24*time.Hour)
	return err
}

func (this *UserModel) SetOfflineAvaliableResetTs(userId int, t time.Time) error {
	resetAtKey := fmt.Sprintf(UserOfflineAvaliabeResetAt, userId)
	return redisDb.Set(resetAtKey, t.Unix()).Err
}
func (this *UserModel) GetOfflineAvaliableResetTs(userId int) (int64, error) {
	resetAtKey := fmt.Sprintf(UserOfflineAvaliabeResetAt, userId)
	return redisDb.Get(resetAtKey).Int64Def(0)
}

func (this *UserModel) GetOfflineAvaliableMax(userId int) (int, error) {
	maxKey := fmt.Sprintf(UserOfflineAvaliabeMax, userId)
	return redisDb.Get(maxKey).IntDef(0)
}

func (this *UserModel) HgetAllFightDiscAward(userId int) (int, error) {
	maxKey := fmt.Sprintf(UserOfflineAvaliabeMax, userId)
	return redisDb.Get(maxKey).IntDef(0)
}

func (this *UserModel) SetLoginDay(userId, day int) {
	key := fmt.Sprintf(USER_LOGIN_DAY, userId, day)
	redisDb.Set(key, 1)
	redisDb.ExpireSecond(key, 7*24*60*60)
}

func (this *UserModel) GetLoginDay(userId, day int) bool {
	key := fmt.Sprintf(USER_LOGIN_DAY, userId, day)
	value, _ := redisDb.Get(key).IntDef(0)
	if value == 0 {
		redisDb.Set(key, 1)
		redisDb.ExpireSecond(key, 7*24*60*60)
		return true
	} else {
		return false
	}
}

func (this *UserModel) SendHour0MailState(day int) error {
	return redisDb.Set(SendHour5MailState, day).Err
}
func (this *UserModel) GetHour0MailState() (int64, error) {
	return redisDb.Get(SendHour5MailState).Int64Def(0)
}

func (this *UserModel) UserSubscribe(userId int, subscribe int) error {
	key := fmt.Sprintf(UserSubscribe, subscribe)
	_, err := redisDb.SAdd(key, userId)
	if err != nil {
		logger.Error("订阅异常,订阅：%v,玩家：%v,err:%v", subscribe, userId, err)
		return err
	}
	return nil
}

func (this *UserModel) UserSubscribeCheck(userId int, subscribe int) bool {
	key := fmt.Sprintf(UserSubscribe, subscribe)
	isSubscribe, err := redisDb.SIsMember(key, userId)
	if err != nil {
		logger.Error("订阅检查异常,订阅：%v,玩家：%v,err:%v", subscribe, userId, err)
		return false
	}
	return isSubscribe
}

func (this *UserModel) GetSubscribeUser(subscribe int) []int {
	key := fmt.Sprintf(UserSubscribe, subscribe)
	all, err := redisDb.SmembersInt(key)
	if err != nil {
		logger.Error("订阅获取集合异常,订阅：%v,err:%v", subscribe, err)
	}
	return all
}

func (this *UserModel) UserSubscribeRemove(subscribe int, userId ...int) {
	key := fmt.Sprintf(UserSubscribe, subscribe)
	arg := make([]interface{}, 0)
	arg = append(arg, key)
	for _, v := range userId {
		arg = append(arg, v)
	}
	_, err := redisDb.SRem(arg...)
	if err != nil {
		logger.Error("订阅移除异常,订阅：%v,玩家：%v,err:%v", subscribe, userId, err)
	}
}
