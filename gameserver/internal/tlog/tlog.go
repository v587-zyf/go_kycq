package tlog

import (
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
	"runtime/debug"
	"strconv"
	"time"

	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/util"
)

type TLog struct {
	util.DefaultModule
}

func NewTLogManager() *TLog {
	return &TLog{}
}

func getFormatTime() string {
	currentTime := time.Now().Local()
	return currentTime.Format("2006-01-02 15:04:05")
}

//玩家注册
func (this *TLog) PlayerRegister(user *objs.User) {
	log := &LogPlayerRegister{}
	log.LogCommon = getCommon(user)
	DbLog(LogModelManager.DbMap(), log)
}

//玩家登陆
func (this *TLog) PlayerLogin(user *objs.User) {
	log := &LogPlayerLogin{}
	log.LogCommon = getCommon(user)
	DbLog(LogModelManager.DbMap(), log)
}

//玩家登出
func (this *TLog) PlayerLogout(user *objs.User) {
	log := &LogPlayerLogout{}
	log.LogCommon = getCommon(user)
	DbLog(LogModelManager.DbMap(), log)
}

//货币流水
func (this *TLog) MoneyFlow(user *objs.User, moneyType, count, afterMoney, Reason int, addOrReduce bool) {
	if user.Id < 0 {
		return
	}
	log := &LogMoneyFlow{}
	log.UserId = user.Id
	log.GameSvrId = strconv.Itoa(user.ServerId)
	log.DtEventTime = getFormatTime()
	log.IZoneAreaId = user.ServerId
	log.Vopenid = user.OpenId
	log.Sequence = 0
	//log.Level = user.Lvl
	log.IMoneyType = moneyType
	log.AfterMoney = afterMoney
	log.IMoney = count
	log.Reason = Reason
	log.SubReason = 0
	DbLog(LogModelManager.DbMap(), log)
}

//道具流水表
func (this *TLog) ItemFlow(user *objs.User, itemId, count, afterCount, Reason int, reason2 int, addOrReduce bool) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("ItemFlow panic: %v, %s", err, debug.Stack())
		}
	}()

	// 省略经验记录
	if itemId == pb.ITEMID_EXP {
		return
	}

	if user.Id < 0 {
		return
	}

	iAddOrReduce := 0
	if !addOrReduce {
		iAddOrReduce = 1
	}
	reasonLv1, reasonLv2 := ophelper.GetResaon(Reason, reason2)
	log := &LogItemFlow{}
	log.LogCommon = getCommon(user)
	//log.UserLv = user.Lvl
	log.IGoodsId = itemId
	log.Count = count
	log.AfterCount = afterCount
	log.Reason = reasonLv1
	log.Reason2 = reasonLv2
	log.AddOrReduce = iAddOrReduce
	DbLog(LogModelManager.DbMap(), log)
}

func getCommon(user *objs.User) LogCommon {
	return LogCommon{
		DtEventTime: getFormatTime(),
		UserId:      user.Id,
		ServerId:    user.ServerId,
		Openid:      user.OpenId,
	}
}
