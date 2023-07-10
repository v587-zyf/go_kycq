package dabao

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

/**
 *  @Description: 进入打宝秘境
 *  @param user
 *  @param stageId
 *  @return error
 */
func (this *DaBao) EnterMystery(user *objs.User, stageId int) error {
	mysteryCfg := gamedb.GetDaBaoMysteryByStageId(stageId)
	if mysteryCfg == nil {
		return gamedb.ERRPARAM
	}
	if check := this.GetCondition().CheckMultiBySlice2(user, -1, mysteryCfg.Condition); !check {
		return gamedb.ERRPARAM
	}
	if user.DaBaoMystery.Energy < mysteryCfg.LimitEnergy {
		return gamedb.ERRENERGY
	}

	fightId, err := this.GetFight().CreateFight(stageId, nil)
	if err != nil {
		return err
	}
	err = this.GetFight().EnterFightByFightId(user, stageId, fightId)
	if err != nil {
		return err
	}
	return nil
}

/**
 *  @Description: 打宝秘境结算
 *  @param user
 *  @param stageId
 *  @param isWin
 *  @param items
 *  @param op
 */
//func (this *DaBao) MysteryResult(user *objs.User, stageId int, isWin bool, items map[int]int) {
//	if user == nil {
//		return
//	}
//	ntf := &pb.DaBaoMysteryResultNtf{
//		StageId: int32(stageId),
//		Result:  pb.RESULTFLAG_FAIL,
//	}
//	if isWin {
//		ntf.Result = pb.RESULTFLAG_SUCCESS
//		ntf.Goods = ophelper.CreateGoodsChangeNtf(items)
//	}
//	this.GetUserManager().SendMessage(user, ntf, true)
//}

/**
 *  @Description: 恢复体力
 *  @param user
 */
func (this *DaBao) ResumeEnergy(user *objs.User) {
	hasEnergy := user.DaBaoMystery.Energy
	energyTop := gamedb.GetConf().DaBaoMysteryEnergy
	if hasEnergy >= energyTop {
		return
	}

	nowTime := int(time.Now().Unix())
	resumeCfg := gamedb.GetConf().DaBaoMysteryEnergyResume
	if n := (nowTime - user.DaBaoMystery.ResumeTime) / resumeCfg[0]; n > 0 {
		addEnergy := resumeCfg[1] * n
		if hasEnergy+addEnergy > energyTop {
			addEnergy = energyTop - hasEnergy
		}
		this.SyncEnergy(user, addEnergy)
	}
}

func (this *DaBao) SendSystemDropItem(user *objs.User, stageId int, replyItems map[int32]*pbserver.ItemUnitForPickUp) {
	items := make(map[int]int)
	for _, v := range replyItems {
		items[int(v.ItemId)] += int(v.ItemNum)
	}
	if len(items) <= 0 {
		return
	}
	//this.GetAnnouncement().FightSendSystemChat(user, items, stageId, pb.SCROLINGTYPE_DA_BAO)
}

/**
 *  @Description: 使用恢复体力道具
 *  @param user
 *  @param itemId
 *  @param op
 *  @return error
 */
func (this *DaBao) EnergyItemUse(user *objs.User, itemId int, op *ophelper.OpBagHelperDefault) error {
	itemCfg := gamedb.GetItemBaseCfg(itemId)
	if itemCfg == nil || itemCfg.Type != pb.ITEMTYPE_DABAO_MY_STERY_ITEM {
		return gamedb.ERRPARAM
	}
	if err := this.GetBag().Remove(user, op, itemId, 1); err != nil {
		return err
	}
	this.SyncEnergy(user, itemCfg.EffectVal)
	return nil
}

/**
 *  @Description: 同步体力
 *  @param user
 *  @param changeNum
 */
func (this *DaBao) SyncEnergy(user *objs.User, changeNum int) {
	userMystery := user.DaBaoMystery
	userMystery.Energy += changeNum
	if userMystery.Energy < 0 {
		userMystery.Energy = 0
	}
	userMystery.ResumeTime = int(time.Now().Unix())
	user.Dirty = true
	this.GetUserManager().SendMessage(user, &pb.DaBaoMysteryEnergyNtf{Energy: int32(userMystery.Energy)}, true)

	energyTop := gamedb.GetConf().DaBaoMysteryEnergy
	if userMystery.Energy >= energyTop {
		user.CheckDaBaoMysteryEnergy = false
	} else {
		user.CheckDaBaoMysteryEnergy = true
	}

	this.FSSendMessage(user.FightId, user.FightStageId, &pbserver.DaBaoResumeEnergyReq{UserId: int32(user.Id), Energy: int32(userMystery.Energy)})
}
