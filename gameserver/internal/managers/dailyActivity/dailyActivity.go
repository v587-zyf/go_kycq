package dailyActivity

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

type DailyActivityManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewDailyActivityManager(module managersI.IModule) *DailyActivityManager {
	return &DailyActivityManager{IModule: module}
}

/**
 *  @Description: 日常活动入口
 *  @param user
 *  @param activityType	活动类型
 *  @return error
 */
func (this *DailyActivityManager) EnterDailyActivity(user *objs.User, activityId int, stageId int) error {
	if activityId <= 0 || stageId <= 0 {
		return gamedb.ERRPARAM
	}
	dailyActivityCfg := gamedb.GetDailyActivityDailyActivityCfg(activityId)
	if dailyActivityCfg == nil {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("activityId:%:v", activityId)
	}
	//如果没合服,判断开服时间,合过服判断合服时间
	if !this.GetSystem().IsMerge() {
		serverOpenDays := this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
		if serverOpenDays < dailyActivityCfg.OpenDayMin || serverOpenDays > dailyActivityCfg.OpenDayMax {
			return gamedb.ERRACTIVITYNOTOPEN
		}
	} else {
		serverMergeDays := this.GetSystem().GetServerMergeDayByServerId(user.ServerId)
		if serverMergeDays < dailyActivityCfg.MergeDayMin || serverMergeDays > dailyActivityCfg.MergeDayMax {
			return gamedb.ERRACTIVITYNOTOPEN
		}
	}

	nowTime := time.Now()
	openTime, closeTime := gamedb.GetActiveTime(dailyActivityCfg.OpenTime, dailyActivityCfg.CloseTime, dailyActivityCfg.Week)
	if int(nowTime.Unix()) < openTime {
		return gamedb.ERRACTIVITYNOTOPEN
	}
	if int(nowTime.Unix()) > closeTime {
		return gamedb.ERRACTIVITYCLOSE
	}
	//进入条件
	if check := this.GetCondition().CheckMulti(user, -1, dailyActivityCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}

	var err error
	switch dailyActivityCfg.ActivityType {
	case pb.DAILYACTIVITYTYPE_PAODIAN:
		err = this.GetPaoDian().EnterPaoDianFight(user, stageId)
	case pb.DAILYACTIVITYTYPE_MAGIC_TOWER:
		_, err = this.GetMagicTower().EnterMagicTower(user)
	}
	return err
}

func (this *DailyActivityManager) List(ack *pb.DailyActivityListAck) {
	ack.List = make([]*pb.DailyActivityInfo, 0)
	listCfgs := gamedb.GetDailyActivityCfgs()
	zeroTime := common.GetZeroTimeUnixFrom1970()
	for id, cfg := range listCfgs {
		ack.List = append(ack.List, &pb.DailyActivityInfo{
			ActivityId: int32(id),
			StartTime:  int64(zeroTime + cfg.OpenTime.GetSecondsFromZero()),
			EndTime:    int64(zeroTime + cfg.CloseTime.GetSecondsFromZero()),
		})
	}
}
