package magicCircle

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewMagicCircleManagerManager(module managersI.IModule) *MagicCircleManager {
	return &MagicCircleManager{IModule: module}
}

type MagicCircleManager struct {
	util.DefaultModule
	managersI.IModule
}

/**
 *  @Description: 法阵升级
 *  @param user
 *  @param heroIndex
 *  @param magicCircleId	法阵类型
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *MagicCircleManager) UpLv(user *objs.User, heroIndex, magicCircleId int, op *ophelper.OpBagHelperDefault, ack *pb.MagicCircleUpLvAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}

	if gamedb.GetMagicCircleMagicCircleCfg(magicCircleId) == nil {
		return gamedb.ERRPARAM
	}
	excelId, ok := hero.MagicCircle[magicCircleId]
	if !ok {
		excelId = 0
	}

	var nextLvCfg, circleLvCfg *gamedb.MagicCircleLevelMagicCircleLevelCfg
	if excelId == 0 {
		circleLvCfg = gamedb.GetMagicCircleLvCfg(magicCircleId, excelId, excelId)
		nextLvCfg = gamedb.GetMagicCircleLevelMagicCircleLevelCfg(circleLvCfg.Id + 1)
	} else {
		nextLvCfg = gamedb.GetMagicCircleLevelMagicCircleLevelCfg(excelId + 1)
		if nextLvCfg == nil || nextLvCfg.Type != magicCircleId {
			return gamedb.ERRLVENOUGH
		}
		circleLvCfg = gamedb.GetMagicCircleLevelMagicCircleLevelCfg(excelId)
	}
	if !this.GetCondition().CheckMulti(user, heroIndex, circleLvCfg.Condition) {
		return gamedb.ERRCONDITION
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, circleLvCfg.Item); err != nil {
		return err
	}
	kyEvent.FaZhenUp(user, heroIndex, magicCircleId, hero.MagicCircle[magicCircleId], nextLvCfg.Id)
	hero.MagicCircle[magicCircleId] = nextLvCfg.Id

	ack.MagicCircleType = int32(magicCircleId)
	ack.HeroIndex = int32(heroIndex)
	ack.ExcelId = int32(hero.MagicCircle[magicCircleId])
	this.GetUserManager().UpdateCombat(user, heroIndex)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_HERO_MAGIC_CIRCLE_UP_STAR, []int{1})
	this.GetTask().UpdateTaskProcess(user, false, false)
	return nil
}

/**
 *  @Description: 法阵穿戴
 *  @param user
 *  @param heroIndex
 *  @param magicCircleLvId	法阵等级表id
 *  @param ack
 *  @return error
 */
func (this *MagicCircleManager) ChangeWear(user *objs.User, heroIndex, magicCircleLvId int, ack *pb.MagicCircleChangeWearAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	magicCircleLevelCfg := gamedb.GetMagicCircleLevelMagicCircleLevelCfg(magicCircleLvId)
	if magicCircleLevelCfg == nil {
		return gamedb.ERRPARAM
	}
	magicCircleId, ok := hero.MagicCircle[magicCircleLevelCfg.Type]
	if !ok || magicCircleId < magicCircleLvId {
		return gamedb.ERRNOTACTIVE
	}

	if hero.Wear.MagicCircleLvId == magicCircleLvId {
		hero.Wear.MagicCircleLvId = 0
	} else {
		hero.Wear.MagicCircleLvId = magicCircleLvId
		kyEvent.FazhenChange(user, heroIndex, magicCircleLvId)
	}
	user.Dirty = true

	this.GetUserManager().SendDisplay(user)
	ack.HeroIndex = int32(heroIndex)
	ack.MagicCircleLvId = int32(hero.Wear.MagicCircleLvId)
	return nil
}
