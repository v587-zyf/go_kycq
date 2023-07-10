package vip

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewVipManager(module managersI.IModule) *VipManager {
	return &VipManager{IModule: module}
}

type VipManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *VipManager) Online(user *objs.User) {
	this.updateVipEffect(user)
}

/**
 *  @Description: 增加vip经验
 *  @param user
 *  @param op
 *  @param count	经验数量
 *  @return error
 */
func (this *VipManager) AddExp(user *objs.User, op *ophelper.OpBagHelperDefault, count int) error {
	//充值或消耗元宝兑换的积分
	vipLvl, vipScore := user.VipLevel, user.VipScore
	vipSetting := gamedb.GetVipLvlCfg(vipLvl)
	vipMaxLv := gamedb.GetVipMaxLv()
	if vipLvl >= vipMaxLv {
		user.VipScore += count
		op.OnGoodsChange(builder.BuildTopDataChange(pb.ITEMID_VIP_EXP, count, user.VipScore), count)
		logger.Info("玩家vip已经满级，只增加经验,user:%v-%v,vipSetting:%v,exp:%v,add:%v", user.Id, user.NickName, user.VipLevel, user.VipScore, count)
		return nil
	}
	vipScore += count
	for vipScore >= vipSetting.Exp {
		vipLvl++
		//vipScore -= vipSetting.Exp
		if vipLvl < vipMaxLv {
			vipSetting = gamedb.GetVipLvlCfg(vipLvl)
			if vipSetting == nil {
				break
			}
		} else {
			break
		}
	}

	beforVipLv := user.VipLevel
	user.VipLevel = vipLvl
	user.VipScore = vipScore
	user.Dirty = true
	op.OnGoodsChange(builder.BuildTopDataChange(pb.ITEMID_VIP_EXP, 0, user.VipScore), count)
	if vipLvl > beforVipLv {
		op.OnGoodsChange(builder.BuildTopDataChange(pb.ITEMID_VIP_LV, 0, user.VipLevel), 0)
		this.GetBag().BagSpaceAddByType(user, constBag.BAG_ADD_TYPE_VIP)
		this.GetTask().AddTaskProcess(user, pb.CONDITION_UPGRADE_VIP_LV, -1)
		this.GetFight().UpdateUserfightNum(user)
		this.GetCondition().RecordCondition(user, pb.CONDITION_UPGRADE_VIP_LV, []int{})
	}
	this.updateVipEffect(user)
	return nil
}

/**
 *  @Description: 领取vip礼包
 *  @param user
 *  @param lv	vip等级
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *VipManager) GetGift(user *objs.User, lv int, op *ophelper.OpBagHelperDefault, ack *pb.VipGiftGetAck) error {
	userVipGift := user.VipGift
	if _, ok := userVipGift[lv]; ok {
		return gamedb.ERRREPEATRECEIVE
	}
	if user.VipLevel < lv {
		return gamedb.ERRVIPLVNOTENOUGH
	}

	vipLvlCfg := gamedb.GetVipLvlCfg(lv)
	if vipLvlCfg == nil {
		return gamedb.ERRPARAM
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, vipLvlCfg.Cost2); err != nil {
		return err
	}

	this.GetBag().AddItems(user, vipLvlCfg.Reward, op)
	userVipGift[lv] = 0
	user.Dirty = true

	ack.Lv = int32(lv)
	return nil
}

/**
 *  @Description: 领取累充礼包
 *  @param user
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *VipManager) GetRechargeAllGift(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.RechargeAllGetAck) error {
	cfg := gamedb.GetAccumulateAccumulateCfg(id)
	if cfg == nil {
		return gamedb.ERRSETTINGNOTFOUND
	}

	for _, ids := range user.AccumulativeId {
		if ids == id {
			return gamedb.ERRAWARDGET
		}
	}
	_, ok := this.GetCondition().Check(user, -1, pb.CONDITION_RECHARGE_ALL_CHECK, cfg.Condition[pb.CONDITION_RECHARGE_ALL_CHECK])
	if !ok {
		return gamedb.ERRRECHARGEALL
	}
	this.GetBag().AddItems(user, cfg.Reward, op)
	user.AccumulativeId = append(user.AccumulativeId, id)
	user.Dirty = true

	allIds := make([]int32, 0)
	for _, ids := range user.AccumulativeId {
		allIds = append(allIds, int32(ids))
	}
	ack.RechargeAll = int32(user.RechargeAll)
	ack.RechargeGetGetIds = allIds
	ack.Goods = op.ToChangeItems()
	return nil
}

/**
 *  @Description: 获取vip特权
 *  @param user
 *  @param privilege 特权常量
 *  @return int
 */
func (this *VipManager) GetPrivilege(user *objs.User, privilege int) int {
	p := 0
	if cfg := gamedb.GetVipLvlCfg(user.VipLevel); cfg != nil {
		if num, ok := cfg.Privilege[privilege]; ok {
			p = num
		}
	}
	if privilege != pb.VIPPRIVILEGE_ATTR {
		p += this.GetMonthCard().GetPrivilege(user, privilege)
		p += this.GetPrivilegeModule().GetPrivilege(user, privilege)
	}
	return p
}

func (this *VipManager) updateVipEffect(user *objs.User) {
	effect := this.GetPrivilege(user, pb.VIPPRIVILEGE_ATTR)
	if effect != 0 {
		effectSlice := make([]int, 0)
		effectSlice = append(effectSlice, effect)
		for _, hero := range user.Heros {
			hero.VipEffects = effectSlice
		}
	}
	this.GetUserManager().UpdateCombat(user, -1)
}
