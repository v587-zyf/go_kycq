package bag

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constConsume"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gamelibs/rmodelCross"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"math"
	"strconv"
)

//道具增加
func (this *BagManager) AddItem(user *objs.User, helper *ophelper.OpBagHelperDefault, itemId, count int) error {
	if count <= 0 {
		logger.Error("玩家：%v,通过：%v 添加道具:%v 数量为0", user.IdName(), helper.GetOpType(), itemId)
		return gamedb.ERRITEMZERO
	}
	itemT := gamedb.GetItemBaseCfg(itemId)
	if itemT == nil {
		logger.Error("玩家：%v,通过：%v 添加道具:%v 配置异常", user.IdName(), helper.GetOpType(), itemId)
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("item" + strconv.Itoa(itemId))
	}

	if !this.CheckHasEnoughPos(user, []*gamedb.ItemInfo{{itemId, count}}) {
		//通过邮件发放
		itemSource := ophelper.GetBagFullMailItemSource(helper.GetOpType(), helper.OpTypeSecond())
		this.GetMail().SendSystemMail(user.Id, constMail.BAG_FULL, nil, []*model.Item{{ItemId: itemId, Count: count}}, itemSource)
		logger.Info("背包已满，道具通过邮件发放,玩家：%v,道具：%v-%v，来源：%v", user.IdName(), itemId, count, helper.GetOpType(), helper.OpTypeSecond())
		helper.BuildItemGetDisplay(itemId, count)
		return nil
	}
	user.Dirty = true
	return this.directAddItem(user, helper, itemId, count)
}

/**
 *  @Description: 直接添加道具到背包
 *  @param user
 *  @param helper
 *  @param itemId
 *  @param count
 *  @return error
 */
func (this *BagManager) directAddItem(user *objs.User, helper *ophelper.OpBagHelperDefault, itemId, count int) error {
	itemT := gamedb.GetItemBaseCfg(itemId)
	var err error
	switch itemT.Type {
	case pb.ITEMTYPE_TOP:
		err = this.topItemAdd(user, helper, itemId, count)
	case pb.ITEMTYPE_EQUIP:
		err = this.equipItemAdd(user, helper, itemId, count)
	case pb.ITEMTYPE_CONTRIBUTION:
		err = this.addContribution(user, helper, itemId, count)
	default:
		err = this.addIntoBag(user, helper, itemId, count)
		//this.GoodsChangeToLog(user, itemT.Type, itemId, count, user.Items.GetCount(itemId), helper, true)
	}

	if err != nil {
		logger.Info("背包添加道具异常：玩家：%v-%v,itemId:%v,count:%v，err:%v", user.Id, user.NickName, itemId, count, err)
		return err
	}
	user.Dirty = true

	// 添加日志
	hasNum, _ := this.GetItemNum(user, itemId)
	this.GetTlog().ItemFlow(user, itemId, count, hasNum, helper.GetOpType(), helper.OpTypeSecond(), true)
	kyEvent.ItemChange(user, itemId, count, hasNum, helper.GetOpType(), helper.OpTypeSecond(), true)
	this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_GET_ITEM, []int{count, itemId})
	return nil
}

func (this *BagManager) AddItems(user *objs.User, items gamedb.ItemInfos, helper *ophelper.OpBagHelperDefault) bool {

	if !this.CheckHasEnoughPos(user, items) {
		//通过邮件发放
		itemSource := ophelper.GetBagFullMailItemSource(helper.GetOpType(), helper.OpTypeSecond())
		this.GetMail().SendSystemMailWithItemInfosSignItemSource(user.Id, constMail.BAG_FULL, nil, items, itemSource)
		for _, v := range items {
			helper.BuildItemGetDisplay(v.ItemId, v.Count)
		}
		logger.Info("背包空间不足，通过邮件发放奖励,玩家：%v,来源：%v", user.NickName, helper.GetOpType(), helper.OpTypeSecond())
		return true
	} else {
		for _, v := range items {
			if v.Count > 0 {
				err := this.directAddItem(user, helper, v.ItemId, v.Count)
				if err != nil {
					logger.Error("添加道具异常，玩家：%v,道具：%v-%v，来源：%v,异常：%v", user.IdName(), v.ItemId, v.Count, helper.GetOpType())
				}
			}
		}
	}
	return false
}

//道具增加
func (this *BagManager) Add(user *objs.User, helper *ophelper.OpBagHelperDefault, itemId, count int) error {
	return this.AddItem(user, helper, itemId, count)
}

//顶级数据增加（经验 等级 vip经验 vip等级 元宝 金币）
func (this *BagManager) topItemAdd(user *objs.User, op *ophelper.OpBagHelperDefault, itemId, count int) error {

	lvAdd := false
	var err error
	if itemId == pb.ITEMID_EXP {
		_, err = this.addExp(user, op, count)
		if err != nil {
			return err
		}
	} else if itemId == pb.ITEMID_VIP_EXP {
		return this.GetVipManager().AddExp(user, op, count)
	} else if itemId == pb.ITEMID_WAR_ORDER_EXP {
		return this.GetWarOrder().AddExp(user, op, count)
	} else {
		c, err1 := user.AddTopDataByItemId(itemId, count)
		if itemId == pb.ITEMID_LV {
			for _, v := range user.Heros {
				v.ExpLvl += count
			}
			lvAdd = true
		}
		if err1 != nil {
			return err1
		}
		op.OnGoodsChange(&pb.TopDataChange{Id: int32(itemId), Change: int64(count), NowNum: int64(c)}, count)
		////通知竞技场
		//this.GetArena().AddArenaRank(user)
	}
	if lvAdd {
		this.GetUserManager().UpdateCombat(user, -1)
		//this.GetRank().Append(pb.RANKTYPE_LEVEL, user.Id, user.Lvl, false)
		this.GetTask().UpdateTaskProcess(user, true, false)
	}
	return nil
}

//玩家经验等级增加
func (this *BagManager) addExp(user *objs.User, op *ophelper.OpBagHelperDefault, exp int) (bool, error) {
	if exp < 0 {
		logger.Error("玩家：%v,通过：%v 添加经验，经验值为0", user.IdName(), op.GetOpType())
		return false, gamedb.ERRPARAM
	}
	changeExp := exp
	curExp := user.Exp
	worldLv, limit, types, addPec := this.GetExpPool().GetExpWorldLvlAndLvLimit(user)
	if worldLv != -1 && addPec != -1 {
		if types == 1 {
			//减少
			pec := addPec / float64(10000)
			exp -= int(math.Ceil(float64(exp) * pec))
		} else if types == 2 {
			//增加
			pec := addPec / float64(10000)
			exp += int(math.Ceil(float64(exp) * pec))
		}
	}

	curExp += exp

	if curExp > limit*gamedb.GetConf().ExpPoolLimitMultiple {
		logger.Error("升级到达当前等级上线 limit:%v  倍数:%v", limit, gamedb.GetConf().ExpPoolLimitMultiple)
		exp = limit*gamedb.GetConf().ExpPoolLimitMultiple - user.Exp
		curExp = limit * gamedb.GetConf().ExpPoolLimitMultiple
	}

	op.OnGoodsChange(&pb.TopDataChange{int32(pb.ITEMID_EXP), int64(changeExp), int64(curExp)}, exp)
	user.SetTopDataByItemId(pb.ITEMID_EXP, curExp)
	this.GetExpPool().AutoUpExpLv(user)
	return false, nil
}

func (this *BagManager) equipItemAdd(user *objs.User, op *ophelper.OpBagHelperDefault, itemId int, count int) error {

	//获取空位
	emptyPos := this.getEmptyPos(user, count)
	if len(emptyPos) < count {
		logger.Error("玩家：%v,通过：%v 添加道具:%v 背包空位不足,添加数量：%v,空位：%v", user.IdName(), op.GetOpType(), itemId, count, emptyPos)
		return gamedb.ERRBAGENOUGH
	}
	for i := 0; i < count; i++ {
		//生成唯一索引
		equipIndex, err := modelGame.GetUserModel().GetEquipId()
		if err != nil {
			logger.Error("玩家：%v,通过：%v 添加道具:%v 生成装备索引异常：%v", user.IdName(), op.GetOpType(), itemId, err)
			return err
		}
		//随机属性
		randProps, err := gamedb.RandEquipProp(itemId)
		if err != nil {
			logger.Error("玩家：%v,通过：%v 添加道具:%v 生成装随机属性异常：%v", user.IdName(), op.GetOpType(), itemId, err)
			return err
		}
		//添加到背包
		bagItem(user.Bag[emptyPos[i]], itemId, 1)
		user.Bag[emptyPos[i]].EquipIndex = equipIndex
		//添加到装备背包
		user.EquipBag[equipIndex] = &model.Equip{
			Index:     equipIndex,
			ItemId:    itemId,
			RandProps: randProps,
		}
		op.OnGoodsChange(builder.BuildEquipDataChagne(itemId, 1, 1, emptyPos[i], user.EquipBag[equipIndex]), 1)
	}
	return nil
}

//道具添加到背包
func (this *BagManager) addIntoBag(user *objs.User, op *ophelper.OpBagHelperDefault, itemId, count int) error {

	if itemId == constConsume.FIRST_RECHARGE_DISCOUNT {
		trailServer := rmodelCross.GetSystemSeting().GetSystemSettingConverInt(rmodelCross.SYSTEM_SETTING_TRIAL_SERVER)
		if trailServer == 1 {
			logger.Info("提审服 不送首充优惠券,玩家：%v,道具：%v-%v", user.Id, itemId, count)
			return nil
		}
	}

	itemT := gamedb.GetItemBaseCfg(itemId)
	if itemT.CountLimit <= 0 || itemT.CountLimit > 1 {
		count = this.stackAdd(user, op, itemId, count)
	}
	if count <= 0 {
		return nil
	}
	//计算需要位置数
	addNeedEmptyNum := this.GetItemSpaceNum(itemId, count)
	//获取空位
	emptyPosSlice := this.getEmptyPos(user, addNeedEmptyNum)
	if len(emptyPosSlice) < addNeedEmptyNum {
		logger.Debug("添加道具：%v num:%v,计算需要空位数量：%v,获取到空位数量：%v", itemId, count, addNeedEmptyNum, len(emptyPosSlice))
		return gamedb.ERRBAGENOUGH
	}
	//添加到背包
	for k := 0; k < addNeedEmptyNum; k++ {
		cnt := itemT.CountLimit
		if cnt > 0 {
			if cnt > count {
				cnt = count
			}
			count -= cnt
		} else {
			cnt = count
			count = 0
		}
		bagItem(user.Bag[emptyPosSlice[k]], itemId, cnt)
		op.OnGoodsChange(builder.BuildItemDataChange(itemId, cnt, cnt, emptyPosSlice[k]), cnt)
	}
	return nil
}

//已有堆叠添加
func (this *BagManager) stackAdd(user *objs.User, op *ophelper.OpBagHelperDefault, itemId int, count int) int {

	itemT := gamedb.GetItemBaseCfg(itemId)
	for i := 0; i < user.BagInfo[constBag.BAG_TYPE_COMMON].MaxNum; i++ {

		itemUnit := user.Bag[i]
		if itemUnit == nil || itemUnit.ItemId != itemId {
			continue
		}

		if itemT.CountLimit <= 0 || itemUnit.Count+count <= itemT.CountLimit {
			itemUnit.Count = itemUnit.Count + count
			op.OnGoodsChange(builder.BuildItemDataChange(itemId, count, itemUnit.Count, i), count)
			count = 0
			break
		} else {
			diff := itemT.CountLimit - itemUnit.Count
			if diff > 0 {
				count -= diff
				itemUnit.Count = itemT.CountLimit
				op.OnGoodsChange(builder.BuildItemDataChange(itemId, diff, itemUnit.Count, i), diff)
			}
		}
		// 已经分配完了
		if count <= 0 {
			break
		}
	}
	return count
}

//增加贡献值特殊处理
func (this *BagManager) addContribution(user *objs.User, op *ophelper.OpBagHelperDefault, itemId, count int) error {
	if user.GuildData.NowGuildId <= 0 {
		return gamedb.ERRNOGUILD
	}

	guildInfo := this.GetGuild().GetGuildInfo(user.GuildData.NowGuildId)
	if guildInfo == nil {
		logger.Error("GetGuildInfo nil userId:%v  guildId:%v", user.Id, user.GuildData.NowGuildId)
		return gamedb.ERRNOGUILD
	}
	guildInfo.GuildContributionValue += count
	user.GuildData.ContributionValue += count
	this.GetGuild().SetGuildInfo(guildInfo)
	kyEvent.GuildUp(user, guildInfo.GuildName, guildInfo.GuildId, len(guildInfo.Positions)/2, guildInfo.GuildContributionValue)
	op.OnGoodsChange(&pb.TopDataChange{Id: int32(itemId), Change: int64(count), NowNum: int64(guildInfo.GuildContributionValue)}, count)
	return nil
}
