package panacea

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewPanaceaManager(m managersI.IModule) *PanaceaManager {
	return &PanaceaManager{
		IModule: m,
	}
}

type PanaceaManager struct {
	util.DefaultModule
	managersI.IModule
}

/**
 *  @Description: 灵丹使用
 *  @param user
 *  @param id	灵丹id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *PanaceaManager) PanaceaUse(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.PanaceaUseAck) error {
	if id < 1 {
		return gamedb.ERRPARAM
	}
	panacea := this.UpdateUseNum(user, gamedb.GetPanaceaPanaceaCfg(id), id)
	if panacea.Number >= panacea.Numbers {
		return gamedb.ERRUSEENOUGH
	}
	err := this.GetBag().Remove(user, op, id, 1)
	if err != nil {
		return err
	}
	panacea.Number++
	ack.Id = int32(id)
	ack.Panacea = builder.BuilderPanaceaInfo(panacea)

	this.GetUserManager().UpdateCombat(user, -1)

	kyEvent.PanaceaUse(user, id, panacea.Number)
	this.GetCondition().RecordCondition(user, pb.CONDITION_USE_LIN_DAN, []int{1})
	this.GetTask().AddTaskProcess(user, pb.CONDITION_USE_LIN_DAN, -1)
	return nil
}

//更新所有灵丹使用次数
func (this *PanaceaManager) PanaceaUpUseNum(user *objs.User) {
	panaceaCfgs := gamedb.GetPanaceaCfgs()
	for id, panaceaCfg := range panaceaCfgs {
		this.UpdateUseNum(user, panaceaCfg, id)
	}
	this.GetUserManager().UpdateCombat(user, -1)
}

func (this *PanaceaManager) UpdateUseNum(user *objs.User, panaceaCfg *gamedb.PanaceaPanaceaCfg, id int) *model.PanaceaUnit {
	userPanacea, ok := user.Panaceas[id]
	if !ok {
		user.Panaceas[id] = &model.PanaceaUnit{}
		userPanacea = user.Panaceas[id]
	}
	for cfgIndex, condition := range panaceaCfg.Condition {
		if _, b := this.GetCondition().Check(user, -1, condition[0], condition[1]); !b {
			break
		}
		if userPanacea.Numbers < panaceaCfg.Limit[cfgIndex] {
			userPanacea.Numbers = panaceaCfg.Limit[cfgIndex]
		}
	}
	return userPanacea
}
