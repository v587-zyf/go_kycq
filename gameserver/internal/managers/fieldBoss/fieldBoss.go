package fieldBoss

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

type FieldBoss struct {
	util.DefaultModule
	managersI.IModule
}

func NewFieldBoss(module managersI.IModule) *FieldBoss {
	f := &FieldBoss{IModule: module}
	return f
}

func (this *FieldBoss) Online(user *objs.User) {
	date := common.GetResetTime(time.Now())
	this.ResetFieldBoss(user, date)
}

func (this *FieldBoss) ResetFieldBoss(user *objs.User, date int) {
	userFieldBoss := user.FieldBoss
	if userFieldBoss.ResetTime != date {
		userFieldBoss.ResetTime = date
		userFieldBoss.DareNum = 0
		userFieldBoss.BuyNum = 0
	}
}

/**
 *  @Description: 野外boss加载
 *  @param user
 *  @param area	区域
 *  @param ack
 *  @return error
 */
func (this *FieldBoss) Load(user *objs.User, area int, ack *pb.FieldBossLoadAck) error {
	fieldBossMap := gamedb.GetFieldBossByArea(area)
	if len(fieldBossMap) < 1 {
		return gamedb.ERRPARAM
	}
	pbFieldBoss := make([]*pb.FieldBossNtf, 0)
	for stageId := range fieldBossMap {
		bossInfo := this.GetFieldBossInfo(stageId)
		bossInfo.Area = int32(area)
		pbFieldBoss = append(pbFieldBoss, bossInfo)
	}
	ack.FieldBoss = pbFieldBoss
	return nil
}

/**
 *  @Description: 野外boss购买次数
 *  @param user
 *  @param use
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *FieldBoss) BuyNum(user *objs.User, use bool, buyNum int, op *ophelper.OpBagHelperDefault, ack *pb.FieldBossBuyNumAck) error {
	if buyNum < 1 {
		return gamedb.ERRPARAM
	}
	userFieldBoss := user.FieldBoss
	if userFieldBoss.BuyNum+buyNum > gamedb.GetConf().FieldBossTimes[1]+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_FIELDBOSS_BUYNUM) {
		return gamedb.ERRENOUGHTIMES
	}

	cost := gamedb.GetConf().FieldBossCost
	consume := cost[0]
	hasNum, _ := this.GetBag().GetItemNum(user, consume.ItemId)
	if hasNum < 1 && use {
		consume = cost[1]
	}
	if err := this.GetBag().Remove(user, op, consume.ItemId, consume.Count*buyNum); err != nil {
		return err
	}
	user.FieldBoss.BuyNum += buyNum
	user.Dirty = true

	ack.BuyNum = int32(user.FieldBoss.BuyNum)
	return nil
}

/**
 *  @Description: 推送野外boss信息
 *  @param stageId
 *  @param blood		血量
 *  @param reliveTime	复活时间
 */
func (this *FieldBoss) SendFieldBossNtf(fieldBossNtf *pb.FieldBossNtf) {
	fieldBossInfo := gamedb.GetFieldBossByStageId(int(fieldBossNtf.StageId))
	if fieldBossInfo != nil {
		fieldBossNtf.Area = int32(fieldBossInfo.Area)
	}
	this.BroadcastAll(fieldBossNtf)
}

/**
 *  @Description: 客户端动画-击杀野外首领领取奖励
 *  @param user
 *  @param op
 *  @return error
 */
func (this *FieldBoss) FirstReceive(user *objs.User, op *ophelper.OpBagHelperDefault) error {
	userFieldBoss := user.FieldBoss
	if userFieldBoss.FirstReceive {
		return gamedb.ERRREPEATRECEIVE
	}
	dropItem, _, err := this.GetStageManager().GetBossDropItem(user, constFight.FIGHT_TYPE_FIELD_BOSS_STAGE)
	if err != nil {
		return err
	}
	this.GetBag().AddItems(user, dropItem, op)
	userFieldBoss.FirstReceive = true
	this.GetCondition().RecordCondition(user, pb.CONDITION_CHALLENGE_FIELD_LEADER, []int{1})
	this.GetTask().AddTaskProcess(user, pb.CONDITION_CHALLENGE_FIELD_LEADER, -1)
	user.Dirty = true
	return nil
}

/**
 *  @Description: 获取野外boss信息
 *  @param stageId	怪物id
 */
func (this *FieldBoss) GetFieldBossInfo(stageId int) *pb.FieldBossNtf {
	fieldBossInfo := this.GetFight().GetFieldBossInfos(stageId)
	return &pb.FieldBossNtf{
		StageId:    fieldBossInfo.StageId,
		Blood:      fieldBossInfo.Hp * 100,
		ReliveTime: fieldBossInfo.ReliveTime,
	}
}

/**
 *  @Description: 获取剩余次数
 *  @return int
 */
func (this *FieldBoss) GetSurplusNum(user *objs.User) int {
	userFieldBoss := user.FieldBoss
	defFightNum := gamedb.GetConf().FieldBossTimes[0]
	vipPrivilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_FIELDBOSS_FREENUM)
	return defFightNum + userFieldBoss.BuyNum + vipPrivilege - userFieldBoss.DareNum
}
