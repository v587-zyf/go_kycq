package area

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewAreaManager(m managersI.IModule) *AreaManager {
	return &AreaManager{
		IModule: m,
	}
}

type AreaManager struct {
	util.DefaultModule
	managersI.IModule
}

/**
 *  @Description: 领域升级
 *  @param user
 *  @param heroIndex
 *  @param id	领域id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *AreaManager) UpLv(user *objs.User, heroIndex, id int, op *ophelper.OpBagHelperDefault) error {
	if id <= 0 {
		return gamedb.ERRPARAM
	}
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	if gamedb.GetAreaAreaCfg(id) == nil {
		return gamedb.ERRPARAM
	}
	_, ok := pb.AREATYPE_MAP[id]
	if !ok {
		return gamedb.ERRPARAM
	}

	if gamedb.GetAreaLevelAreaLevelCfg(gamedb.GetRealId(id, hero.Area[id]+1)) == nil {
		return gamedb.ERRLVENOUGH
	}
	lvCfg := gamedb.GetAreaLevelAreaLevelCfg(gamedb.GetRealId(id, hero.Area[id]))
	if !this.GetCondition().CheckMulti(user, heroIndex, lvCfg.Condition) {
		return gamedb.ERRCONDITION
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, lvCfg.Item); err != nil {
		return err
	}
	hero.Area[id]++
	user.Dirty = true

	kyEvent.LingYuUp(user, heroIndex, id, hero.Area[id]-1, hero.Area[id])

	this.GetUserManager().UpdateCombat(user, heroIndex)
	this.GetTask().UpdateTaskProcess(user, false, false)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_HERO_AREA_UP_LV, []int{})
	return nil
}
