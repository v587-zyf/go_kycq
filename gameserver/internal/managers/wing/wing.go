package wing

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewWingManager(module managersI.IModule) *WingManager {
	return &WingManager{IModule: module}
}

const (
	wingSpecialDefLv = 1
)

type WingManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *WingManager) Online(user *objs.User) {
	for _, hero := range user.Heros {
		wing := hero.Wings[0]
		curId, curExp := wing.Id, wing.Exp
		wingCfg := gamedb.GetWingNewWingNewCfg(curId)
		if wingCfg == nil {
			return
		}
		nextCfg := gamedb.GetWingNewWingNewCfg(curId + 1)
		if nextCfg == nil {
			return
		}
		for wingCfg != nil && nextCfg != nil && curExp >= wingCfg.Consume.Count {
			if !this.GetCondition().CheckMulti(user, -1, wingCfg.Condition) {
				break
			}
			curId++
			curExp -= wingCfg.Consume.Count
			this.WingSpecialAutoUpLv(user, wingCfg.Star, gamedb.GetMaxValById(wingCfg.Order, constMax.MAX_WING_STAR), wing, hero)
			wingCfg = gamedb.GetWingNewWingNewCfg(curId)
			nextCfg = gamedb.GetWingNewWingNewCfg(curId + 1)
		}
		wing.Id = curId
		wing.Exp = curExp
	}
}

/**
 *  @Description: 羽翼升级
 *  @param user
 *  @param heroIndex
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *WingManager) UpLevel(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault, ack *pb.WingUpLevelAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	if err := this.GetCondition().CheckFunctionOpen(user, pb.FUNCTIONID_WING_OPEN); err != nil {
		return err
	}

	allItemCount := 0
	wing := hero.Wings[0]
	curId := wing.Id
	beforeLv := wing.Id
	wingConf := gamedb.GetWingCfg(curId)
	if !this.GetCondition().CheckMulti(user, heroIndex, wingConf.Condition) {
		return gamedb.ERRCONDITION
	}
	if gamedb.GetWingCfg(curId+1) == nil {
		itemId := gamedb.GetWingCfg(curId - 1).Consume.ItemId
		removeNum := this.GetBag().RemoveAllByItemId(user, op, itemId)
		wing.Exp += removeNum
		allItemCount += removeNum
	} else {
		curStar := wingConf.Star
		maxStar := gamedb.GetMaxValById(wingConf.Order, constMax.MAX_WING_STAR)
		for ; curStar < maxStar; curStar++ {
			wingConf := gamedb.GetWingCfg(wing.Id)
			if !this.GetCondition().CheckMulti(user, heroIndex, wingConf.Condition) {
				break
			}
			itemId, needCount := wingConf.Consume.ItemId, wingConf.Consume.Count
			if wing.Exp > 0 {
				needCount -= wing.Exp
			}
			itemCount, err := this.GetBag().GetItemNum(user, itemId)
			if err != nil {
				return err
			}
			if itemCount >= needCount {
				err := this.GetBag().Remove(user, op, itemId, needCount)
				if err != nil {
					return err
				}
				wing.Id = wingConf.Id + 1
				wing.Exp = 0
				allItemCount += needCount
			} else {
				this.GetBag().Remove(user, op, itemId, itemCount)
				wing.Exp += itemCount
				allItemCount += itemCount
				break
			}
		}
		this.WingSpecialAutoUpLv(user, curStar, maxStar, wing, hero)
	}

	kyEvent.WingUp(user, heroIndex, beforeLv, wing.Id)

	ack.HeroIndex = int32(heroIndex)
	ack.Wing = builder.BuildWing(wing)

	this.GetUserManager().UpdateCombat(user, heroIndex)
	if cfg := gamedb.GetWingNewWingNewCfg(wing.Id); cfg != nil {
		this.GetTask().AddTaskProcess(user, pb.CONDITION_UPGRADE_SHEN_YI_2, -1)
	}
	this.GetCondition().RecordCondition(user, pb.CONDITION_THREE_WING_GRADE, []int{})

	this.GetCondition().RecordCondition(user, pb.CONDITION_UPGRADE_SHEN_YI_1, []int{allItemCount})
	this.GetTask().UpdateTaskProcess(user, false, false)

	return nil
}

/**
 *  @Description: 羽翼特殊技能升级
 *  @param user
 *  @param heroIndex
 *  @param specialT
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *WingManager) UpSpecialLevel(user *objs.User, heroIndex, specialT int, op *ophelper.OpBagHelperDefault, ack *pb.WingSpecialUpAck) error {
	if specialT < 1 {
		return gamedb.ERRPARAM
	}
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	if err := this.GetCondition().CheckFunctionOpen(user, pb.FUNCTIONID_WING_OPEN); err != nil {
		return err
	}

	wingSpecialLv := hero.WingSpecial[specialT]
	if wingSpecialLv == 0 {
		return gamedb.ERRNOTACTIVE
	}

	wingSpecialMaxLv := gamedb.GetMaxValById(specialT, constMax.MAX_WING_SPECIAL_LEVEL)
	if wingSpecialLv >= wingSpecialMaxLv {
		return gamedb.ERRLVENOUGH
	}

	wingSpecialCfg := gamedb.GetWingSpecialByOrderAndLv(specialT, wingSpecialLv)
	err := this.GetBag().Remove(user, op, wingSpecialCfg.Consume.ItemId, wingSpecialCfg.Consume.Count)
	if err != nil {
		return err
	}
	hero.WingSpecial[specialT]++

	ack.HeroIndex = int32(heroIndex)
	ack.WingSpecial = builder.BuilderWingSpecial(specialT, hero.WingSpecial[specialT])
	kyEvent.WingSpecialLvUp(user, heroIndex, hero.WingSpecial[specialT])

	this.GetUserManager().UpdateCombat(user, heroIndex)
	return nil
}

/**
 *  @Description: 使用材料单次升级羽翼
 *  @param user
 *  @param heroIndex
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *WingManager) UseMaterial(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault, ack *pb.WingUseMaterialAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	wing := hero.Wings[0]
	wingConf := gamedb.GetWingCfg(wing.Id)
	nextConf := gamedb.GetWingCfg(wing.Id + 1)

	itemId := 0
	if nextConf == nil {
		itemId = gamedb.GetWingCfg(wing.Id - 1).Consume.ItemId
	} else {
		itemId = wingConf.Consume.ItemId
	}
	if !this.GetCondition().CheckMulti(user, heroIndex, wingConf.Condition) {
		return gamedb.ERRCONDITION
	}
	allDelCount := 0
	needItemNum := wingConf.Consume.Count - wing.Exp
	if nextConf == nil {
		hasNum, _ := this.GetBag().GetItemNum(user, itemId)
		if err := this.GetBag().Remove(user, op, itemId, hasNum); err != nil {
			return err
		}
		wing.Exp += hasNum
		allDelCount += hasNum
	} else {
		if needItemNum > 0 {
			if hasNum, _ := this.GetBag().GetItemNum(user, itemId); hasNum >= needItemNum {
				if err := this.GetBag().Remove(user, op, itemId, needItemNum); err != nil {
					return err
				}
				wing.Id++
				wing.Exp = 0
				allDelCount += needItemNum
			} else {
				if err := this.GetBag().Remove(user, op, itemId, hasNum); err != nil {
					return err
				}
				wing.Exp += hasNum
				allDelCount += hasNum
			}
		}
	}

	curStar := gamedb.GetWingCfg(wing.Id).Star
	maxStar := gamedb.GetMaxValById(wingConf.Order, constMax.MAX_WING_STAR)
	this.WingSpecialAutoUpLv(user, curStar, maxStar, wing, hero)

	ack.HeroIndex = int32(heroIndex)
	ack.Wing = builder.BuildWing(wing)

	this.GetCondition().RecordCondition(user, pb.CONDITION_UPGRADE_SHEN_YI_1, []int{allDelCount})
	this.GetUserManager().UpdateCombat(user, heroIndex)
	this.GetTask().UpdateTaskProcess(user, false, false)
	cfg := gamedb.GetWingNewWingNewCfg(wing.Id)
	if cfg != nil {
		this.GetTask().AddTaskProcess(user, pb.CONDITION_UPGRADE_SHEN_YI_2, -1)
	}
	this.GetCondition().RecordCondition(user, pb.CONDITION_THREE_WING_GRADE, []int{})
	return nil
}

/**
 *  @Description: 羽翼穿戴
 *  @param user
 *  @param heroIndex
 *  @param wingId
 *  @param ack
 *  @return error
 */
func (this *WingManager) Wear(user *objs.User, heroIndex, wingId int, ack *pb.WingWearAck) error {
	if wingId < 1 {
		return gamedb.ERRPARAM
	}
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}

	if hero.Wings[0].Id < wingId {
		return gamedb.ERRPARAM
	}
	if hero.Wear.WingId == wingId {
		hero.Wear.WingId = 0
	} else {
		hero.Wear.WingId = wingId
	}
	user.Dirty = true

	ack.HeroIndex = int32(heroIndex)
	ack.WingId = int32(hero.Wear.WingId)
	this.GetUserManager().SendDisplay(user)
	return nil
}

// 羽翼特殊技能自动激活
func (this *WingManager) WingSpecialAutoUpLv(user *objs.User, curStar int, maxStar int, wing *model.Wing, hero *objs.Hero) {
	if curStar == maxStar {
		wingConf := gamedb.GetWingNewWingNewCfg(wing.Id)
		if !this.GetCondition().CheckMulti(user, hero.Index, wingConf.Condition) {
			return
		}
		if gamedb.GetWingCfg(wing.Id+1) != nil {
			wing.Id += 1
		}
		wingOrder := gamedb.GetWingNewWingNewCfg(wing.Id).Order
		wingSpecialCfgs := gamedb.GetWingSpecialCfgs()
		for _, cfg := range wingSpecialCfgs {
			specialT, lv := cfg.Type, cfg.Level
			if lv != wingSpecialDefLv || hero.WingSpecial[specialT] > 0 || wingOrder < cfg.Order {
				continue
			}
			hero.WingSpecial[specialT] = wingSpecialDefLv
			this.GetUserManager().SendMessage(user, &pb.WingSpecialUpAck{
				HeroIndex:   int32(hero.Index),
				WingSpecial: builder.BuilderWingSpecial(specialT, wingSpecialDefLv),
			}, true)
		}
	}
}
