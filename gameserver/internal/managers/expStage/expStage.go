package expStage

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

const (
	APPRAOSE_F = iota
	APPRAOSE_E
	APPRAOSE_D
	APPRAOSE_C
	APPRAOSE_B
	APPRAOSE_A
	APPRAOSE_S
	APPRAOSE_SS
	APPRAOSE_SSS
)

func NewExpStageManager(module managersI.IModule) *ExpStageManager {
	return &ExpStageManager{IModule: module}
}

type ExpStageManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *ExpStageManager) OnLine(user *objs.User) {
	date := common.GetResetTime(time.Now())
	this.ResetExpStage(user, date)
}

func (this *ExpStageManager) ResetExpStage(user *objs.User, date int) {
	userExpStage := user.ExpStage
	if userExpStage.ResetTime != date {
		userExpStage.ResetTime = date
		defNum := gamedb.GetConf().ExpStageFightNum
		if userExpStage.DareNum > defNum {
			userExpStage.BuyNum -= userExpStage.DareNum - defNum
			if userExpStage.BuyNum < 0 {
				userExpStage.BuyNum = 0
			}
		}
		userExpStage.DareNum = 0
	}
}

/**
 *  @Description:经验副本领取双倍经验
 *  @param user
 *  @param op
 *  @param stageId
 *  @param ack
 *  @return error
 */
func (this *ExpStageManager) Double(user *objs.User, op *ophelper.OpBagHelperDefault, stageId int, ack *pb.ExpStageDoubleAck) error {
	if stageId <= 0 {
		return gamedb.ERRPARAM
	}
	exp, ok := user.ExpStage.ExpStages[stageId]
	if !ok || exp == 0 {
		return gamedb.ERRNOTFIGHT
	}

	num := this.getKillExpStageNum(user)
	if num > 1 {
		consumes := gamedb.GetConf().ExpStageDoubleCost
		if err := this.GetBag().RemoveItemsInfos(user, op, consumes); err != nil {
			return err
		}
	}

	this.GetBag().Add(user, op, pb.ITEMID_EXP, exp)
	user.ExpStage.ExpStages[stageId] = 0
	user.Dirty = true

	ack.StageId = int32(stageId)
	ack.Exp = int64(exp)
	return nil
}

func (this *ExpStageManager) getKillExpStageNum(user *objs.User) int {
	data := user.Conditions[pb.CONDITION_ALL_KILL_STAGE]
	num := 0
	for i := 0; i < len(data); i += 2 {
		if cfg := gamedb.GetExpStageByStageId(data[i]); cfg != nil {
			num += data[i+1]
		}
	}
	return num
}

/**
 *  @Description: 获取剩余次数
 *  @return int
 */
func (this *ExpStageManager) GetSurplusNum(user *objs.User) int {
	userExpStage := user.ExpStage
	freeNum := gamedb.GetConf().ExpStageFightNum
	privilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_EXPSTAGE_FIGHTNUM)
	return freeNum + userExpStage.BuyNum + privilege - userExpStage.DareNum
}

/**
 *  @Description: 经验副本使用增加次数的道具校验
 *  @param user
 *  @param itemId	道具id
 *  @return error
 */
func (this *ExpStageManager) ExpStageBuyNumCheck(user *objs.User) (int, error) {
	buyConf := gamedb.GetConf().LevelAddTimes
	buyItemId, buyMaxNum := buyConf[0][0], buyConf[0][1]
	if user.ExpStage.BuyNum >= buyMaxNum {
		return buyItemId, gamedb.ERRBUYTIMESLIMIT
	}
	return buyItemId, nil
}

/**
 *  @Description:经验副本使用增加次数的道具
 *  @param user
 *  @param itemId 道具id
 *  @return error
 */
func (this *ExpStageManager) ExpStageBuyNumNtf(user *objs.User) error {
	user.ExpStage.BuyNum++

	this.GetUserManager().SendMessage(user, &pb.ExpStageBuyNumNtf{
		BuyNum: int32(user.ExpStage.BuyNum),
	}, true)
	return nil
}

/**
 *  @Description: 购买并使用道具增加次数
 *  @param user
 *  @param use
 *  @param op
 *  @return error
 */
func (this *ExpStageManager) ExpStageBuyNum(user *objs.User, use bool, op *ophelper.OpBagHelperDefault) error {
	buyItemId, err := this.ExpStageBuyNumCheck(user)
	if err != nil {
		return err
	}
	if use {
		err = this.GetBag().RemoveItemsInfos(user, op, gamedb.GetConf().ExperienceCost)
	} else {
		err = this.GetBag().Remove(user, op, buyItemId, 1)
	}
	if err != nil {
		return gamedb.ERRNOTENOUGHGOODS
	}

	user.ExpStage.BuyNum++
	user.Dirty = true
	return nil
}
