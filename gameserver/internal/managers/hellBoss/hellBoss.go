package hellBoss

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

func NewHellManager(m managersI.IModule) *HellBoss {
	return &HellBoss{IModule: m}
}

type HellBoss struct {
	util.DefaultModule
	managersI.IModule
}

func (this *HellBoss) Online(user *objs.User) {
	date := common.GetResetTime(time.Now())
	this.ResetHellBoss(user, date)
}

func (this *HellBoss) ResetHellBoss(user *objs.User, date int) {
	hellBoss := user.HellBoss
	if hellBoss.ResetTime != date {
		hellBoss.ResetTime = date
		hellBoss.DareNum = 0
		hellBoss.BuyNum = 0
		hellBoss.HelpNum = 0
	}
}

func (this *HellBoss) Load(user *objs.User, floor int) []*pb.HellBossNtf {
	pbSlice := make([]*pb.HellBossNtf, 0)
	bossCfgs := gamedb.GetHellBossByFloor(floor)
	nowTime := time.Now().Unix()
	for stageId := range bossCfgs {
		stageInfo := this.GetFight().GetHellBossInfos(stageId)
		if stageInfo == nil {
			logger.Warn("获取炼狱首领信息错误")
			continue
		}
		var reliveTime int64 = 0
		if nowTime < stageInfo.ReliveTime {
			reliveTime = stageInfo.ReliveTime
		}
		pbSlice = append(pbSlice, &pb.HellBossNtf{
			StageId:    int32(stageId),
			Blood:      stageInfo.Hp,
			ReliveTime: reliveTime,
		})
	}
	return pbSlice
}

/**
 *  @Description: 炼狱首领购买次数
 *  @param user
 *  @param use
 *  @param buyNum
 *  @param op
 *  @return error
 */
func (this *HellBoss) BuyNum(user *objs.User, use bool, buyNum int, op *ophelper.OpBagHelperDefault) error {
	if buyNum < 1 {
		return gamedb.ERRPARAM
	}
	userHellBoss := user.HellBoss
	vipPrivilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_DARKPALACE_BUYNUM)
	if userHellBoss.BuyNum+buyNum > gamedb.GetConf().HellBossAdd + vipPrivilege {
		return gamedb.ERRENOUGHTIMES
	}
	cost := gamedb.GetConf().HellBossBuy
	consume := cost[0]
	hasNum, _ := this.GetBag().GetItemNum(user, consume.ItemId)
	if hasNum < 1 && use {
		consume = cost[1]
	}
	if err := this.GetBag().Remove(user, op, consume.ItemId, consume.Count*buyNum); err != nil {
		return err
	}
	userHellBoss.BuyNum += buyNum
	user.Dirty = true
	return nil
}

/**
 *  @Description: 获取剩余次数
 *  @param user
 *  @return int
 */
func (this *HellBoss) GetSurplusNum(user *objs.User) int {
	userHellBoss := user.HellBoss
	defFightNum := gamedb.GetConf().HellBossTime
	vipPrivilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_HELLBOSS_FREENUM)
	return (defFightNum + userHellBoss.BuyNum + vipPrivilege) - userHellBoss.DareNum
}

/**
 *  @Description: 是否开启跨服
 *  @return bool
 */
func (this *HellBoss) checkCross() bool {
	crossFightServerId := this.GetSystem().GetCrossFightServerId()
	if crossFightServerId > 1 {
		return true
	}
	return false
}
