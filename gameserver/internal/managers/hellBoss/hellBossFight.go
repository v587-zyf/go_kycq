package hellBoss

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

func (this *HellBoss) EnterHellBossFight(user *objs.User, stageId, helpUserId int) error {
	if stageId <= 0 {
		return gamedb.ERRPARAM
	}
	//if this.GetSurplusNum(user) <= 0 {
	//	return gamedb.ERRNOTENOUGHTIMES
	//}
	bossCfg := gamedb.GetHellBossByStage(stageId)
	if bossCfg == nil {
		return gamedb.ERRPARAM
	}
	for k, v := range bossCfg.Condition {
		if _, check := this.GetCondition().CheckBySlice(user, -1, []int{k, v}); !check {
			return gamedb.ERRCONDITION
		}
	}
	if bossInfo := this.GetFight().GetHellBossInfos(stageId); bossInfo.Hp <= 0 {
		return gamedb.ERRFIGHTEND
	}
	err := this.GetFight().EnterResidentFightByStageId(user, stageId, helpUserId)
	if err != nil {
		logger.Error("进入炼狱boss战斗异常,玩家：%v,关卡：%v，err:%v", user.IdName(), stageId, err)
	}
	return nil
}

func (this *HellBoss) HellBossFightResult(user *objs.User, stageId, winUserId int, items map[int]int, toHelpUserId int) {
	ntf := &pb.HellBossFightResultNtf{
		StageId: int32(stageId),
		Result:  pb.RESULTFLAG_FAIL,
	}
	winnerId := user.Id
	userHellBoss := user.HellBoss
	if winUserId != user.Id {
		winnerId = winUserId
		if toHelpUserId > 0 {
			ntf.IsHelper = true
			if toHelpUserId == winUserId {
				if userHellBoss.HelpNum < gamedb.GetConf().HellBossHelp {
					user.Dirty = true
					userHellBoss.HelpNum++
					bossCfg := gamedb.GetHellBossByStage(stageId)
					op := ophelper.NewOpBagHelperDefault(constBag.OpTypeHelp)
					op.SetOpTypeSecond(stageId)
					this.GetBag().AddItems(user, bossCfg.HelpDrop, op)
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
		this.GetFirstDrop().CheckIsFirstDrop(user, items)
	}
	ntf.Winner = this.GetUserManager().BuilderBrieUserInfo(winnerId)
	ntf.DareNum = int32(userHellBoss.DareNum)
	ntf.HelpNum = int32(userHellBoss.HelpNum)
	this.GetUserManager().SendMessage(user, ntf, true)
}
