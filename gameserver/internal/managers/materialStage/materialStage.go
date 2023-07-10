package materialStage

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

type MaterialStage struct {
	util.DefaultModule
	managersI.IModule
}

func NewMaterialStage(module managersI.IModule) *MaterialStage {
	m := &MaterialStage{IModule: module}
	return m
}

func (this *MaterialStage) OnLine(user *objs.User) {
	date := common.GetResetTime(time.Now())
	this.ResetMaterialStage(user, date)
}

func (this *MaterialStage) ResetMaterialStage(user *objs.User, date int) {
	userMaterial := user.MaterialStage
	if userMaterial.MaterialStages == nil {
		userMaterial.MaterialStages = make(map[int]*model.MaterialStageUnit)
	}
	for _, mateType := range pb.MATERIALSTAGETYPE_ARRAY {
		mateInfo, ok := userMaterial.MaterialStages[mateType]
		if !ok {
			userMaterial.MaterialStages[mateType] = &model.MaterialStageUnit{}
			mateInfo = userMaterial.MaterialStages[mateType]
		}
		if userMaterial.ResetTime != date {
			mateInfo.DareNum = 0
			mateInfo.BuyNum = 0
			if mateInfo.NowLayer < 1 {
				mateInfo.NowLayer = 1
			}
		}
	}
	if userMaterial.ResetTime != date {
		userMaterial.ResetTime = date
	}
}

func (this *MaterialStage) BuyNumCheck(user *objs.User, mateType int) (int, gamedb.ItemInfos, error) {
	buyConf := gamedb.GetConf().LevelAddTimes
	buyItemId, buyMaxNum, buyItemCost := 0, 0, gamedb.ItemInfos{}
	switch mateType {
	case pb.MATERIALSTAGETYPE_WING:
		buyItemId = buyConf[2][0]
		buyMaxNum = buyConf[2][1]
		buyItemCost = gamedb.GetConf().WingCost
	case pb.MATERIALSTAGETYPE_GOLD:
		buyItemId = buyConf[1][0]
		buyMaxNum = buyConf[1][1]
		buyItemCost = gamedb.GetConf().CoinCost
	case pb.MATERIALSTAGETYPE_STRENGTH:
		buyItemId = buyConf[3][0]
		buyMaxNum = buyConf[3][1]
		buyItemCost = gamedb.GetConf().StrengthCost
	case pb.MATERIALSTAGETYPE_FASHION:
		buyItemId = buyConf[4][0]
		buyMaxNum = buyConf[4][1]
		buyItemCost = gamedb.GetConf().FashionCost
	default:
		return buyItemId, buyItemCost, gamedb.ERRPARAM
	}
	userMaterial := user.MaterialStage
	buyNum := userMaterial.MaterialStages[mateType].BuyNum
	if buyNum >= buyMaxNum {
		return buyItemId, buyItemCost, gamedb.ERRBUYTIMESLIMIT
	}
	return buyItemId, buyItemCost, nil
}

/**
 *  @Description: 用道具购买次数
 *  @param user
 *  @param mateType	材料副本类型
 *  @param op
 *  @return error
 */
func (this *MaterialStage) BuyNum(user *objs.User, mateType int) error {
	user.MaterialStage.MaterialStages[mateType].BuyNum++
	user.Dirty = true

	this.GetUserManager().SendMessage(user, &pb.MaterialStageBuyNumNtf{
		MaterialType: int32(mateType),
		BuyNum:       int32(user.MaterialStage.MaterialStages[mateType].BuyNum),
	}, true)
	return nil
}

/**
 *  @Description: 购买并使用道具增加次数
 *  @param user
 *  @param mateType	材料副本类型
 *  @param use
 *  @param op
 *  @return error
 */
func (this *MaterialStage) MaterialBuyNum(user *objs.User, mateType int, use bool, op *ophelper.OpBagHelperDefault) error {
	buyItemId, buyItemCost, err := this.BuyNumCheck(user, mateType)
	if err != nil {
		return err
	}
	if use {
		err = this.GetBag().RemoveItemsInfos(user, op, buyItemCost)
	} else {
		err = this.GetBag().Remove(user, op, buyItemId, 1)
	}
	if err != nil {
		return gamedb.ERRNOTENOUGHGOODS
	}
	user.MaterialStage.MaterialStages[mateType].BuyNum++
	user.Dirty = true
	return nil
}
