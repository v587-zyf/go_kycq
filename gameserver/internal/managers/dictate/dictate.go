package dictate

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewDictateManager(m managersI.IModule) *DictateManager {
	return &DictateManager{
		IModule: m,
	}
}

type DictateManager struct {
	util.DefaultModule
	managersI.IModule
}

/**
 *  @Description: 主宰装备激活丶升级
 *  @param user
 *  @param heroIndex
 *  @param body	装备部位
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *DictateManager) DictateUpLv(user *objs.User, heroIndex, body int, op *ophelper.OpBagHelperDefault, ack *pb.DictateUpAck) error {
	if heroIndex <= 0 || body <= 0 {
		return gamedb.ERRPARAM
	}
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	if _, ok := pb.DICTATETYPE_MAP[body]; !ok {
		return gamedb.ERREQUIPTYPE
	}
	dictates := hero.Dictates
	maxLv := gamedb.GetMaxValById(body, constMax.MAX_DICTATE_LEVEL)
	if dictates[body] >= maxLv {
		return gamedb.ERRLVENOUGH
	}

	dictateCfg := gamedb.GetDictateByBodyAndGrade(body, dictates[body])
	if dictateCfg == nil {
		return gamedb.ERRPARAM
	}
	if check := this.GetCondition().CheckMulti(user, heroIndex, dictateCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}

	err := this.GetBag().Remove(user, op, dictateCfg.Consume.ItemId, dictateCfg.Consume.Count)
	if err != nil {
		return err
	}
	dictates[body]++

	ack.HeroIndex = int32(heroIndex)
	ack.DictateInfo = builder.BuildDictateInfo(body, dictates[body])

	kyEvent.ZhuZaiEquipUp(user, heroIndex, body, body, dictates[body]-1, dictates[body])
	this.GetUserManager().UpdateCombat(user, heroIndex)
	this.GetTask().AddTaskProcess(user, pb.CONDITION_UPGRADE_EQUIP, -1)
	this.GetCondition().RecordCondition(user, pb.CONDITION_UPGRADE_EQUIP, []int{})
	return nil
}
