package darkPalace

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

/**
 *  @Description: 暗殿boss进入战斗
 *  @param user
 *  @param floor	层数
 *  @param stageId
 *  @return error
 */
func (this *DarkPalaceManager) EnterDarkPalaceFight(user *objs.User, stageId int, helpUserId int) error {
	if stageId <= 0 {
		return gamedb.ERRPARAM
	}
	//if this.GetSurplusNum(user) <= 0 {
	//	return gamedb.ERRNOTENOUGHTIMES
	//}
	darkPalaceBossCfg := gamedb.GetDarkPalaceStageCfg(stageId)
	if darkPalaceBossCfg == nil {
		return gamedb.ERRPARAM
	}
	if check := this.GetCondition().CheckMulti(user, -1, darkPalaceBossCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}
	if bossInfo := this.GetDarkPalaceBossInfo(stageId); bossInfo.Blood <= 0 {
		return gamedb.ERRFIGHTEND
	}

	err := this.GetFight().EnterResidentFightByStageId(user, stageId, helpUserId)
	if err != nil {
		logger.Error("进入暗殿boss战斗异常,玩家：%v,关卡：%v，err:%v", user.IdName(), stageId, err)
		return nil
	}
	this.GetCondition().RecordCondition(user, pb.CONDITION_TIAO_ZHAN_AN_DIAN_BOSS, []int{1})
	return nil
}

/**
 *  @Description: 暗殿boss战斗回调
 *  @param user
 *  @param floor		层数
 *  @param stageId
 *  @param winUserId	归属者用户id
 *  @param op
 */
func (this *DarkPalaceManager) DarkPalaceFightResultNtf(user *objs.User, stageId, winUserId int, items map[int]int, toHelpUserId int) {
	ntf := &pb.DarkPalaceFightResultNtf{
		StageId: int32(stageId),
		Result:  pb.RESULTFLAG_FAIL,
	}
	winnerId := user.Id
	if user.Id != winUserId {
		winnerId = winUserId
		if toHelpUserId > 0 {
			ntf.IsHelper = true
			if toHelpUserId == winUserId {
				if user.DarkPalace.HelpNum < gamedb.GetConf().DarkBossHelp {
					user.Dirty = true
					user.DarkPalace.HelpNum++
					darkPalaceBossCfg := gamedb.GetDarkPalaceStageCfg(stageId)
					op := ophelper.NewOpBagHelperDefault(constBag.OpTypeHelp)
					op.SetOpTypeSecond(stageId)
					this.GetBag().AddItems(user, darkPalaceBossCfg.HelpDrop, op)
					this.GetUserManager().SendItemChangeNtf(user, op)
					ntf.Goods = op.ToChangeItems()
				}
				ntf.Result = pb.RESULTFLAG_SUCCESS
			}
		}

	} else {
		user.Dirty = true
		ntf.Result = pb.RESULTFLAG_SUCCESS
		ntf.Goods = ophelper.CreateGoodsChangeNtf(items)
		this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_KILL_MONSTER, []int{1, constFight.FIGHT_TYPE_DARKPALACE_BOSS})
		this.GetCondition().RecordCondition(user, pb.CONDITION_KILL_BOSS_NUM, []int{1})
		this.GetFirstDrop().CheckIsFirstDrop(user, items)
	}
	this.GetDailyTask().CompletionOfTask(user, pb.DAILYTASKACTIVITYTYPE_DUO_REN_AN_DIAN_BOSS, 1)
	//this.GetAnnouncement().FightSendSystemChat(user, items, stageId, pb.SCROLINGTYPE_AN_DIAN_BOSS_DROP)
	ntf.Winner = this.GetUserManager().BuilderBrieUserInfo(winnerId)
	ntf.DareNum = int32(user.DarkPalace.DareNum)
	ntf.HelpNum = int32(user.DarkPalace.HelpNum)
	this.GetUserManager().SendMessage(user, ntf, true)
}
