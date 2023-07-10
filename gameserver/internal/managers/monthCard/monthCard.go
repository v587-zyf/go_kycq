package monthCard

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"fmt"
	"time"
)

func NewMonthCardManager(m managersI.IModule) *MonthCardManager {
	return &MonthCardManager{IModule: m}
}

type MonthCardManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *MonthCardManager) Online(user *objs.User) {
	date := common.GetResetTime(time.Now())
	this.ResetMonthCard(user, date)
	this.updateMonthCardEffect(user)
}

func (this *MonthCardManager) ResetMonthCard(user *objs.User, date int) {
	userMonthCard := user.MonthCard
	if userMonthCard.ResetTime != date {
		userMonthCard.ResetTime = date
		user.DayStateRecord.MonthCardReceive = make(model.IntKv)
	}
}

/**
 *  @Description: 月卡，校验是否开放，支付金额
 *  @param typeId	monthCard表id
 *  @param payNum	支付金额
 *  @return error
 */
func (this *MonthCardManager) MonthCardCheckBuy(user *objs.User, typeId, payNum int) error {
	monthCardCfg := gamedb.GetMonthCardMonthCardCfg(typeId)
	if monthCardCfg == nil || monthCardCfg.IsShow == pb.MONTHCARDSTATUS_CLOSE {
		return gamedb.ERRPARAM
	}
	if userMonthCard, ok := user.MonthCard.MonthCards[monthCardCfg.Type]; ok && userMonthCard.EndTime == -1 {
		return gamedb.ERRREPEATBUY
	}
	payMoney := monthCardCfg.Cost
	if payMoney != payNum {
		return gamedb.ERRBUYNUM
	}
	return nil
}

/**
 *  @Description: 月卡购买后续操作
 *  @param user
 *  @param payModuleId	monthCard表id
 *  @param op
 */
func (this *MonthCardManager) MonthCardBuyOperation(user *objs.User, payModuleId int, op *ophelper.OpBagHelperDefault) {
	cfg := gamedb.GetMonthCardMonthCardCfg(payModuleId)
	monthCardT := cfg.Type
	monthCard, ok := user.MonthCard.MonthCards[monthCardT]
	if !ok {
		user.MonthCard.MonthCards[monthCardT] = &model.MonthCardUnit{}
		monthCard = user.MonthCard.MonthCards[monthCardT]
	}
	timeNow := int(time.Now().Unix())
	if monthCard.StartTime == 0 || monthCard.EndTime <= timeNow {
		monthCard.StartTime = timeNow
	}
	if cfg.Day == -1 {
		monthCard.EndTime = -1
	} else {
		if monthCard.EndTime == 0 || monthCard.EndTime <= timeNow {
			monthCard.EndTime = int(time.Now().AddDate(0, 0, cfg.Day).Unix())
		} else {
			dayDuration := time.Duration(cfg.Day) * 24 * time.Hour
			monthCard.EndTime += int(dayDuration.Seconds())
		}
	}
	this.GetBag().AddItems(user, cfg.Reward, op)
	this.activeOperation(user, payModuleId, monthCardT)
}

/**
 *  @Description: 领取月卡每日礼包
 *  @param user
 *  @param id
 *  @param op
 *  @return error
 */
func (this *MonthCardManager) DailyReward(user *objs.User, monthCardType int, op *ophelper.OpBagHelperDefault) error {
	cfg := gamedb.GetMonthCardByType(monthCardType)
	if cfg == nil {
		return gamedb.ERRPARAM
	}
	monthCard, ok := user.MonthCard.MonthCards[monthCardType]
	if !ok {
		return gamedb.ERRMONTHCARDNOTBUY
	}
	if _, ok := user.DayStateRecord.MonthCardReceive[monthCardType]; ok {
		return gamedb.ERRREPEATRECEIVE
	}
	if monthCard.EndTime != -1 && monthCard.EndTime < int(time.Now().Unix()) {
		return gamedb.ERROVERDUECOLLECTIONTIME
	}

	this.GetBag().AddItems(user, cfg.DailyReward, op)
	user.DayStateRecord.MonthCardReceive[monthCardType] = 0
	user.Dirty = true
	return nil
}

/**
 *  @Description: 获取月卡特权
 *  @param user
 *  @param privilege 特权常量
 *  @return int
 */
func (this *MonthCardManager) GetPrivilege(user *objs.User, privilege int) int {
	num := 0
	timeNow := int(time.Now().Unix())
	for t, unit := range user.MonthCard.MonthCards {
		if unit.EndTime >= timeNow || unit.EndTime == -1 {
			cfg := gamedb.GetMonthCardByType(t)
			if cfg != nil {
				if _, ok := cfg.Privilege[privilege]; ok {
					num += cfg.Privilege[privilege]
				}
			}
		}
	}
	return num
}

/**
 *  ItemActiveMonthCardCheck
 *  @Description: 月卡用道具激活检查
 *  @receiver this
 *  @param user
 *  @param itemId
 *  @return error
**/
func (this *MonthCardManager) ItemActiveMonthCardCheck(user *objs.User, itemId int) error {
	monthCardT := pb.MONTHCARDTYPE_GOLD
	userMonthCard, ok := user.MonthCard.MonthCards[monthCardT]
	if ok && userMonthCard.EndTime == -1 {
		return gamedb.ERRREPEATBUY
	}
	itemCfg := gamedb.GetItemBaseCfg(itemId)
	if itemCfg == nil || itemCfg.Type != pb.ITEMTYPE_MONTH_CARD_ITEM {
		return gamedb.ERRITEMUSEFAIL
	}
	if num, err := this.GetBag().GetItemNum(user, itemId); err != nil || num <= 0 {
		return gamedb.ERRITEMUSEFAIL
	}
	return nil
}

/**
 *  ItemActiveMonthCard
 *  @Description: 月卡用道具激活
 *  @receiver this
 *  @param user
 *  @param itemId
 *  @return error
**/
func (this *MonthCardManager) ItemActiveMonthCard(user *objs.User, itemId int) error {
	monthCardT := pb.MONTHCARDTYPE_GOLD
	itemCfg := gamedb.GetItemBaseCfg(itemId)
	if itemCfg == nil {
		return gamedb.ERRITEMUSEFAIL
	}
	monthCard, ok := user.MonthCard.MonthCards[monthCardT]
	if !ok {
		user.MonthCard.MonthCards[monthCardT] = &model.MonthCardUnit{}
		monthCard = user.MonthCard.MonthCards[monthCardT]
	}
	timeNow := time.Now()
	if monthCard.StartTime == 0 || monthCard.EndTime <= int(timeNow.Unix()) {
		monthCard.StartTime = int(timeNow.Unix())
	}
	if monthCard.EndTime == 0 || monthCard.EndTime <= int(timeNow.Unix()) {
		addMinute, _ := time.ParseDuration(fmt.Sprintf(`%dm`, itemCfg.EffectVal))
		monthCard.EndTime = int(timeNow.Add(addMinute).Unix())
	} else {
		addTimeDuration := time.Duration(itemCfg.EffectVal) * 60 * time.Second
		monthCard.EndTime += int(addTimeDuration.Seconds())
	}
	this.activeOperation(user, 0, monthCardT)
	return nil
}

/**
 *  CheckExpire
 *  @Description: 月卡检查是否过期（下线时）
 *  @receiver this
 *  @param user
**/
func (this *MonthCardManager) CheckExpire(user *objs.User) {
	userMonthCard := user.MonthCard
	timeNow := int(time.Now().Unix())
	for _, info := range userMonthCard.MonthCards {
		if info.EndTime != -1 && timeNow >= info.EndTime {
			info.EndTime = 0
		}
	}
}

func (this *MonthCardManager) activeOperation(user *objs.User, payModuleId int, monthCardT int) {
	this.GetFight().UpdateUserfightNum(user)
	this.updateMonthCardEffect(user)
	chatT := pb.SCROLINGTYPE_JI_HUO_BAI_HUANG_JIN_KA
	if monthCardT == pb.MONTHCARDTYPE_SLIVER {
		chatT = pb.SCROLINGTYPE_JI_HUO_BAI_YING_KA
	}
	this.GetAnnouncement().SendSystemChat(user, chatT, -1, -1)
	this.GetUserManager().SendMessage(user, &pb.MonthCardBuyAck{
		Id:            int32(payModuleId),
		MonthCardType: int32(monthCardT),
		MonthCard:     builder.BuildMonthCardUnit(user.MonthCard.MonthCards[monthCardT]),
	}, true)
}

func (this *MonthCardManager) updateMonthCardEffect(user *objs.User) {
	effect := this.GetPrivilege(user, pb.VIPPRIVILEGE_ATTR)
	if effect != 0 {
		effectSlice := make([]int, 0)
		effectSlice = append(effectSlice, effect)
		for _, hero := range user.Heros {
			hero.MonthCardEffects = effectSlice
		}
	}
	this.GetUserManager().UpdateCombat(user, -1)
}
