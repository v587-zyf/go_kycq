package shabake

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"encoding/json"
	"time"
)

func NewShabakeManager(m managersI.IModule) *ShabakeManager {
	return &ShabakeManager{
		IModule: m,
	}
}

type ShabakeManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *ShabakeManager) GetShaBakeOpenDay(serverId int) int {
	return this.GetSystem().GetMergerServerOpenDaysByServerId(serverId)
}

func (this *ShabakeManager) JudgeIsOpen(user *objs.User) error {

	state := rmodel.Shabake.GetShaBakeIsEnd(base.Conf.ServerId)
	if state == 1 {
		return gamedb.ERRACTIVITYCLOSE
	}

	limitedOpenDay := gamedb.GetConf().ShabakeTime1
	openDay := this.GetShaBakeOpenDay(user.ServerId)
	//开服天数限制
	if limitedOpenDay > openDay {
		return gamedb.ERRACTIVITYNOTOPEN
	}

	limitedWeekDay := gamedb.GetConf().ShabakeTime2
	if limitedWeekDay != nil && len(limitedWeekDay) > 0 {
		t := int(time.Now().Weekday())
		if t == 0 {
			t = 7
		}
		//周几限制
		canIn := false
		for _, v := range limitedWeekDay {
			if v == t {
				canIn = true
				break
			}
		}
		if !canIn {
			return gamedb.ERRACTIVITYNOTOPEN
		}
	}
	limitedTimes := gamedb.GetConf().ShabakeTime3
	if limitedTimes == nil {
		return gamedb.ERRACTIVITYNOTOPEN
	}
	if len(limitedTimes) < 2 {
		return gamedb.ERRACTIVITYNOTOPEN
	}

	nowTime := time.Now()
	openTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), limitedTimes[0].Hour, limitedTimes[0].Minute, limitedTimes[0].Second, 0, time.Local)
	endTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), limitedTimes[1].Hour, limitedTimes[1].Minute, limitedTimes[1].Second, 0, time.Local)
	//openTime := limitedTimes[0].GetSecondsFromZero()
	//endTime := limitedTimes[1].GetSecondsFromZero()
	//now := common.GetTimeSeconds(nowTime)
	//每天时间段限制
	if nowTime.Unix() < openTime.Unix() || nowTime.Unix() > endTime.Unix() {
		logger.Error("沙巴克不开启时间段内")
		return gamedb.ERRACTIVITYNOTOPEN
	}
	return nil
}

func (this *ShabakeManager) Load(user *objs.User, ack *pb.ShaBaKeInfoAck) error {

	data := rmodel.Shabake.GetShabakeFirstGuildInfo()
	if data == "" {
		return nil
	}
	infos := &pb.ShaBaKeInfoCrossAck{}
	err := json.Unmarshal([]byte(data), infos)
	if err != nil {
		logger.Error("跨服沙巴克 load unMarshal err:%v", err)
		return nil
	}
	ack.IsEnd = int32(rmodel.Shabake.GetShaBakeIsEnd(base.Conf.ServerId))
	ack.WinGuildUserInfo = infos.WinGuildUserInfo
	ack.FirstGuildName = infos.FirstGuildName
	ack.WinGuildServerId = infos.ServerId
	return nil
}

func (this *ShabakeManager) ShabakeOpenOrCloseNtf(isOpen bool) error {
	this.BroadcastAll(&pb.ShabakeIsOpenNtf{IsOpen: isOpen})
	return nil
}
