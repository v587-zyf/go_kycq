package ancientBoss

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

const (
	OWNER_STR     = "%d[-]%d"
	OWNER_STR_TOO = "%s{-}%s"
)

type AncientBoss struct {
	util.DefaultModule
	managersI.IModule
}

func NewAncientBoss(module managersI.IModule) *AncientBoss {
	return &AncientBoss{IModule: module}
}

func (this *AncientBoss) Online(user *objs.User) {
	date := common.GetResetTime(time.Now())
	this.ResetAncientBoss(user, date)
}

/**
 *  @Description: 重置用户远古首领信息
 *  @param user
 *  @param date
 */
func (this *AncientBoss) ResetAncientBoss(user *objs.User, date int) {
	if user.AncientBoss == nil {
		user.AncientBoss = &model.AncientBoss{}
	}
	userAncient := user.AncientBoss
	if userAncient.ResetTime != date {
		userAncient.ResetTime = date
		userAncient.DareNum = 0
		userAncient.BuyNum = 0
	}
}

/**
 *  @Description: 重置远古首领归属者
 */
func (this *AncientBoss) ResetAncientBossOwner() {
	cfgs := gamedb.GetAncientBossCfgs()
	for stageId := range cfgs {
		rmodel.Boss.DelBossOwner(stageId)
	}
}

/**
 *  @Description: 加载boss列表
 *  @param user
 *  @param area		区域
 *  @param ack
 *  @return error
 */
func (this *AncientBoss) Load(user *objs.User, area int, ack *pb.AncientBossLoadAck) error {
	ancientBossCfgs := gamedb.GetAncientBossCfgsByArea(area)
	if len(ancientBossCfgs) < 1 {
		return gamedb.ERRPARAM
	}
	pbSlice := make([]*pb.AncientBossNtf, 0)
	for stageId := range ancientBossCfgs {
		info := this.getBossInfo(stageId)
		info.Area = int32(area)
		pbSlice = append(pbSlice, info)
	}
	ack.AncientBoss = pbSlice
	return nil
}

/**
 *  @Description: 购买次数
 *  @param user
 *  @param use		是否购买并使用
 *  @param op
 *  @return error
 */
func (this *AncientBoss) BuyNum(user *objs.User, use bool, buyNum int, op *ophelper.OpBagHelperDefault) error {
	if buyNum < 1 {
		return gamedb.ERRPARAM
	}
	userAncient := user.AncientBoss

	vipPrivilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_ANCIENT_BOSS_BUYNUM)
	if userAncient.BuyNum+buyNum > gamedb.GetConf().AncientBossTimes[1] + vipPrivilege {
		return gamedb.ERRENOUGHTIMES
	}

	cost := gamedb.GetConf().AncientBossCost
	consume := cost[0]
	hasNum, _ := this.GetBag().GetItemNum(user, consume.ItemId)
	if hasNum < 1 && use {
		consume = cost[1]
	}
	if err := this.GetBag().Remove(user, op, consume.ItemId, consume.Count*buyNum); err != nil {
		return err
	}
	userAncient.BuyNum += buyNum
	user.Dirty = true

	return nil
}

/**
 *  @Description: 推送boss复活

 *  @param ancientBossInfo
 */
func (this *AncientBoss) SendBossInfo(ancientBossInfo *pb.AncientBossNtf) {
	bossInfo := gamedb.GetAncientBossAncientBossCfg(int(ancientBossInfo.StageId))
	if bossInfo != nil {
		ancientBossInfo.Area = int32(bossInfo.Area)
	}
	this.BroadcastAll(ancientBossInfo)
}

func (this *AncientBoss) getBossInfo(stageId int) *pb.AncientBossNtf {
	bossInfo := this.GetFight().GetAncientBossInfos(stageId)
	return &pb.AncientBossNtf{
		StageId:    int32(stageId),
		Blood:      bossInfo.Hp * 100,
		ReliveTime: bossInfo.ReliveTime,
		UserCount:  bossInfo.UserCount,
	}
}

func (this *AncientBoss) getSurplusNum(user *objs.User) int {
	userAncient := user.AncientBoss
	defFightNum := gamedb.GetConf().AncientBossTimes[0]
	vipPrivilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_ANCIENT_BOSS_FREENUM)
	return defFightNum + (userAncient.BuyNum + vipPrivilege) - userAncient.DareNum
}
