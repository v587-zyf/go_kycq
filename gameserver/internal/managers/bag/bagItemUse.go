package bag

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constConsume"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

func (this *BagManager) ItemUse(user *objs.User, heroIndex, itemId, itemNum int, helperDefault *ophelper.OpBagHelperDefault) error {
	itemConf := gamedb.GetItemBaseCfg(itemId)
	if itemConf == nil {
		return gamedb.ERRPARAM
	}
	if isEnough, _ := this.HasEnough(user, itemId, itemNum); !isEnough {
		return gamedb.ERRNOTENOUGHGOODS
	}

	if itemConf.Type == pb.ITEMTYPE_EXP_PILL {
		return this.GetExpPool().ExpPillUse(user, itemId, itemNum, helperDefault)
	}

	if err := this.checkUseBefore(user, heroIndex, itemId); err != nil {
		logger.Error("道具不可使用：玩家：%v,道具：%v, 原因：%v", user.IdName(), itemId, err)
		return err
	}
	if err := this.Remove(user, helperDefault, itemId, itemNum); err != nil {
		return err
	}

	logger.Debug("Type:%v itemId:%v", itemConf.Type, itemId)
	switch itemConf.Type {
	case pb.ITEMTYPE_RECHARGE_ITEM, pb.ITEMTYPE_RED_PACKET: //红包
		return this.useRedPacket(user, itemConf, itemNum, helperDefault)
	case pb.ITEMTYPE_SKILL:
	//技能
	//this.GetSkill().ActiveSkill(user, itemConf.EffectVal)
	case pb.ITEMTYPE_GIFT_BAG: //礼包
		return this.giftUse(user, itemId, itemNum, helperDefault)
	case pb.ITEMTYPE_PROP: //道具
		return this.useProp(user, heroIndex, itemId)
	case pb.ITEMTYPE_FIELD_FIGHT_ITEM: //野战
		return this.useFightItem(user, itemId)
	case pb.ITEMTYPE_COMPETITVE_ITEM: //竞技场
		return this.useCompetitveItem(user, itemId)
	case pb.ITEMTYPE_RANDOM_STONE, pb.ITEMTYPE_BACK_CITY: //回城石
		user.QuickCd[itemConf.Type] = int(common.GetNowMillisecond())
		return this.GetFight().UseItem(user, itemId)
	case pb.ITEMTYPE_POTION, pb.ITEMTYPE_HP_RECOVER, pb.ITEMTYPE_MP_RECOVER: //药水
		user.QuickCd[itemConf.Type] = int(common.GetNowMillisecond())
		return this.GetFight().UseItem(user, itemId)
	case pb.ITEMTYPE_COPY_FIGHT_ITEM: //副本卷轴
		return this.useCopyFightItem(user, itemId)
	case pb.ITEMTYPE_MONTH_CARD_ITEM: //月卡
		return this.GetMonthCard().ItemActiveMonthCard(user, itemId)
	case pb.ITEMTYPE_PRIVILEGE_ACTIVE: //特权激活
		return this.GetPrivilegeModule().ItemActive(user, itemConf.EffectVal)
	default:
		return gamedb.ERRITEMCANNOTUSE
	}
	user.Dirty = true
	return nil
}

func (this *BagManager) checkUseBefore(user *objs.User, heroIndex, itemId int) error {
	switch itemId {
	case constConsume.BlessItemId, constConsume.BlessItemId1:
		return this.GetEquip().CheckWeaponLucky(user, heroIndex, itemId)
	}

	var err error
	itemConf := gamedb.GetItemBaseCfg(itemId)
	switch itemConf.Type {
	case pb.ITEMTYPE_MONTH_CARD_ITEM: //月卡
		err = this.GetMonthCard().ItemActiveMonthCardCheck(user, itemId)
	case pb.ITEMTYPE_RECHARGE_ITEM:
	//充值道具
	case pb.ITEMTYPE_RED_PACKET:
	//红包
	case pb.ITEMTYPE_SKILL: //技能
		//err = this.GetSkill().CheckSkill(user, itemConf.EffectVal, 0)
	case pb.ITEMTYPE_GIFT_BAG: //礼包
		if itemConf.EffectVal <= 0 {
			logger.Error("礼包未配置效果值，随机道具异常,玩家：%v,itemId:%v，itemNum:%v", user.IdName(), itemId)
			err = gamedb.ERRGIFTDROPID
		}
	case pb.ITEMTYPE_POTION, pb.ITEMTYPE_MP_RECOVER, pb.ITEMTYPE_HP_RECOVER, pb.ITEMTYPE_BACK_CITY, pb.ITEMTYPE_RANDOM_STONE: //药水|回城石
		if !this.quickCd(user, itemId) {
			err = gamedb.ERRREFCD
		}
	case pb.ITEMTYPE_PROP: //道具
		switch itemId {
		case constConsume.BlessItemId, constConsume.BlessItemId1, constConsume.CompetitveScroll, constConsume.FieldScroll:
		default:
			err = gamedb.ERRITEMCANNOTUSE
		}
	case pb.ITEMTYPE_FIELD_FIGHT_ITEM: //野战
		switch itemId {
		case constConsume.FieldScroll:
		default:
			err = gamedb.ERRITEMCANNOTUSE
		}
	case pb.ITEMTYPE_COMPETITVE_ITEM: //竞技场
		switch itemId {
		case constConsume.CompetitveScroll:
			err = this.GetCompetitve().BagUserItemAddChallengeNumCheck(user)
		default:
			err = gamedb.ERRITEMCANNOTUSE
		}
	case pb.ITEMTYPE_COPY_FIGHT_ITEM: //副本道具
		buyNumCfg := gamedb.GetConf().LevelAddTimes
		mateType := 0
		switch itemId {
		case buyNumCfg[0][0]:
			_, err = this.GetExpStage().ExpStageBuyNumCheck(user)
		case buyNumCfg[1][0]:
			mateType = pb.MATERIALSTAGETYPE_WING
			fallthrough
		case buyNumCfg[2][0]:
			mateType = pb.MATERIALSTAGETYPE_GOLD
			fallthrough
		case buyNumCfg[3][0]:
			mateType = pb.MATERIALSTAGETYPE_STRENGTH
			_, _, err = this.GetMaterialStage().BuyNumCheck(user, mateType)
		default:
			err = gamedb.ERRITEMCANNOTUSE
		}
	case pb.ITEMTYPE_PRIVILEGE_ACTIVE: //特权激活
		err = this.GetPrivilegeModule().ItemActiveCheck(user, itemConf.EffectVal)
	default:
		err = gamedb.ERRGIFTDROPID
	}
	return err
}

func (this *BagManager) useCopyFightItem(user *objs.User, itemId int) error {
	buyNumCfg := gamedb.GetConf().LevelAddTimes
	mateType := 0
	switch itemId {
	case buyNumCfg[0][0]:
		return this.GetExpStage().ExpStageBuyNumNtf(user)
	case buyNumCfg[1][0]:
		mateType = pb.MATERIALSTAGETYPE_WING
		fallthrough
	case buyNumCfg[2][0]:
		mateType = pb.MATERIALSTAGETYPE_GOLD
		fallthrough
	case buyNumCfg[3][0]:
		mateType = pb.MATERIALSTAGETYPE_STRENGTH
		return this.GetMaterialStage().BuyNum(user, mateType)
	default:
		return gamedb.ERRITEMCANNOTUSE
	}
	return nil
}

func (this *BagManager) useCompetitveItem(user *objs.User, itemId int) error {
	switch itemId {
	case constConsume.CompetitveScroll: //竞技场卷轴增加挑战次数
		return this.GetCompetitve().BagUserItemAddChallengeNum(user)
	default:
		return gamedb.ERRITEMCANNOTUSE
	}
	return nil
}

func (this *BagManager) useFightItem(user *objs.User, itemId int) error {
	switch itemId {
	case constConsume.FieldScroll:
		return this.GetFieldFight().BuyFieldFightChallengeNumByBag(user)
	default:
		return gamedb.ERRITEMCANNOTUSE
	}
	return nil
}

func (this *BagManager) useProp(user *objs.User, heroIndex int, itemId int) error {
	switch itemId {
	case constConsume.BlessItemId:
		fallthrough
	case constConsume.BlessItemId1:
		return this.GetEquip().WeaponBless(user, heroIndex)
	default:
		return gamedb.ERRITEMCANNOTUSE
	}
	return nil
}

func (this *BagManager) useRedPacket(user *objs.User, itemConf *gamedb.ItemBaseCfg, itemNum int, helperDefault *ophelper.OpBagHelperDefault) error {
	getValue := itemConf.EffectVal * itemNum
	if itemConf.Type == pb.ITEMTYPE_RED_PACKET {
		vipNum := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_RED_PACKET_GET_NUM)
		if user.RedPacketUseNum+getValue > (vipNum+gamedb.GetConf().RedRecoveryCount)*100 {
			return gamedb.ERRUSEENOUGH
		}
		user.RedPacketNum += getValue
		user.RedPacketUseNum += getValue
	}
	user.RechargeAll += getValue
	this.Add(user, helperDefault, pb.ITEMID_INGOT, common.CeilFloat64(float64(getValue)))
	this.Add(user, helperDefault, pb.ITEMID_VIP_EXP, common.CeilFloat64(float64(getValue)/100*float64(gamedb.GetConf().Vip)))
	ntf := &pb.UserRechargeNumNtf{
		RechargeNum:  int32(user.RechargeAll),
		RedPacketNum: int32(user.RedPacketNum),
	}
	this.GetUserManager().SendMessage(user, ntf, true)
	//if itemConf.Type != pb.ITEMTYPE_RED_PACKET {
	//	this.GetFirstRecharge().UpdateFirstRechargeStatus(user)
	//}
	this.GetRecharge().ContRechargeWrite(user, getValue)
	return nil
}

func (this *BagManager) giftUse(user *objs.User, itemId, itemNum int, helperDefault *ophelper.OpBagHelperDefault) error {
	itemConf := gamedb.GetItemBaseCfg(itemId)
	if itemConf.EffectVal <= 0 {
		logger.Error("礼包未配置效果值，随机道具异常,玩家：%v,itemId:%v，itemNum:%v", user.IdName(), itemId)
		return gamedb.ERRGIFTDROPID
	}

	awardItems := make(map[int]int)
	for i := 0; i < itemNum; i++ {
		items, err := gamedb.GetDropItems(itemConf.EffectVal)
		if err != nil {
			return err
		}
		for _, v := range items {
			awardItems[v.ItemId] += v.Count
		}
	}

	for k, v := range awardItems {
		this.Add(user, helperDefault, k, v)
	}
	return nil
}

func (this *BagManager) quickCd(user *objs.User, itemId int) bool {
	cd := 0
	itemConf := gamedb.GetItemBaseCfg(itemId)
	if itemConf.Type == pb.ITEMTYPE_POTION {
		cd = gamedb.GetConf().CdPotion
	} else if itemConf.Type == pb.ITEMTYPE_HP_RECOVER {
		cd = gamedb.GetConf().CdHpRecover
	} else if itemConf.Type == pb.ITEMTYPE_MP_RECOVER {
		cd = gamedb.GetConf().CdMpRecover
	} else if itemConf.Type == pb.ITEMTYPE_RANDOM_STONE {
		cd = gamedb.GetConf().CdRandomStone
	} else if itemConf.Type == pb.ITEMTYPE_BACK_CITY {
		cd = gamedb.GetConf().CdBackCity
	} else {
		return false
	}
	lastUseTime := user.QuickCd[itemConf.Type]
	now := int(common.GetNowMillisecond())
	if now-lastUseTime >= cd {
		return true
	}
	return false
}
