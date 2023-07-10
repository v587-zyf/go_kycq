package fashion

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

type FashionManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewFashionManager(module managersI.IModule) *FashionManager {
	return &FashionManager{IModule: module}
}

func (this *FashionManager) Init() error {
	return nil
}

/**
 *  @Description: 时装激活丶升级
 *  @param user
 *  @param op
 *  @param heroIndex
 *  @param fashionId	时装id
 *  @return error
 */
func (this *FashionManager) FashionUpLevel(user *objs.User, op *ophelper.OpBagHelperDefault, heroIndex, fashionId int) error {
	if heroIndex <= 0 || fashionId <= 0 {
		return gamedb.ERRPARAM
	}
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	nowFashion := hero.Fashions[fashionId]
	nowfashionLv := 0
	beforeLv := 0
	afterLv := 0
	if nowFashion != nil {
		nowfashionLv = nowFashion.Lv
		beforeLv = nowFashion.Lv
	}
	fashionConf := gamedb.GetFashionConf(fashionId, nowfashionLv)
	if fashionConf == nil {
		return gamedb.ERRSETTINGNOTFOUND
	}
	maxLv := gamedb.GetMaxValById(fashionId, constMax.MAX_FASHION_LEVEL)
	if nowfashionLv >= maxLv {
		return gamedb.ERRMAXLV
	}

	err := this.GetBag().Remove(user, op, fashionConf.Consume.ItemId, fashionConf.Consume.Count)
	if err != nil {
		return err
	}

	if nowFashion == nil {
		hero.Fashions[fashionId] = &model.Fashion{
			Id: fashionId,
			Lv: 1,
		}
		afterLv = 1
		if err := this.FashionWear(user, heroIndex, fashionId, true); err == nil {
			this.GetUserManager().SendMessage(user, &pb.FashionWearAck{HeroIndex: int32(heroIndex), WearFashionId: int32(fashionId), IsWear: true}, true)
		} else {
			logger.Warn("Fashion Active Wear err:%v", err)
		}
	} else {
		nowFashion.Lv += 1
		afterLv = nowFashion.Lv
	}
	kyEvent.FashionUp(user, heroIndex, fashionId, beforeLv, afterLv)

	this.GetUserManager().UpdateCombat(user, heroIndex)
	return nil
}

/**
 *  @Description: 时装穿戴
 *  @param user
 *  @param heroIndex
 *  @param fashionId	时装id
 *  @param isWear		是否穿戴
 *  @return error
 */
func (this *FashionManager) FashionWear(user *objs.User, heroIndex, fashionId int, isWear bool) error {
	if heroIndex <= 0 || fashionId <= 0 {
		return gamedb.ERRPARAM
	}
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	nowFashion := hero.Fashions[fashionId]
	nowfashionLv := 0
	if nowFashion == nil {
		return gamedb.ERRFASHIONNOACTIVE
	}
	fashionConf := gamedb.GetFashionConf(fashionId, nowfashionLv)
	if fashionConf == nil {
		return gamedb.ERRSETTINGNOTFOUND
	}
	newFashionId := fashionId
	if !isWear {
		newFashionId = 0
	}
	if fashionConf.FashionType == pb.FASHIONTYPE_WEAPON {
		hero.Wear.FashionWeaponId = newFashionId
	} else {
		hero.Wear.FashionClothId = newFashionId
	}
	user.Dirty = true
	this.GetUserManager().SendDisplay(user)
	return nil
}
