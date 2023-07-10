package juexue

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

type JuexueManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewJuexueManager(module managersI.IModule) *JuexueManager {
	return &JuexueManager{IModule: module}
}

/**
 *  @Description: 绝学激活丶升级
 *  @param user
 *  @param id	绝学id
 *  @param op
 *  @return error
 */
func (this *JuexueManager) JuexueUpLevel(user *objs.User, id int, op *ophelper.OpBagHelperDefault) error {
	if id < 1 {
		return gamedb.ERRPARAM
	}
	juexue := user.Juexues[id]
	if juexue == nil {
		juexue = &model.Juexue{
			Id: id,
			Lv: 0,
		}
	}
	nowLvConf := gamedb.GetJuexueLevelConfCfg(gamedb.GetRealId(juexue.Id, juexue.Lv))
	nextLvConf := gamedb.GetJuexueLevelConfCfg(gamedb.GetRealId(juexue.Id, juexue.Lv+1))
	if nowLvConf == nil || nextLvConf == nil {
		return gamedb.ERRSETTINGNOTFOUND
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, nowLvConf.Consume); err != nil {
		return err
	}
	juexue.Lv += 1
	if user.Juexues[id] == nil {
		user.Juexues[id] = juexue
	}
	kyEvent.JueXueUp(user, id, juexue.Lv-1, juexue.Lv)
	this.GetUserManager().UpdateCombat(user, -1)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ACTIVE_JUE_XUE_SKILL, []int{})
	this.GetTask().UpdateTaskProcess(user, false, false)
	return nil
}
