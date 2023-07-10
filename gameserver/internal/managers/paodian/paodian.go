package paodian

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

func NewPaoDianManager(module managersI.IModule) *PaoDianManager {
	return &PaoDianManager{IModule: module}
}

type PaoDianManager struct {
	util.DefaultModule
	managersI.IModule
}

/**
 *  @Description: 进入泡点pk
 *  @param user
 *  @return error
 */
func (this *PaoDianManager) EnterPaoDianFight(user *objs.User, stageId int) error {
	err := this.GetFight().EnterPaodian(user, stageId)
	if err != nil {
		return err
	}
	this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_PAODIAN_NUM, []int{1})
	return nil
}

/**
 *  @Description: 泡点收益
 *  @param user
 *  @param paoDianRewardId
 */
func (this *PaoDianManager) PaoDianRewardNtf(user *objs.User, paoDianRewardId int, times int) {
	paoDianRewardCfg := gamedb.GetPaoDianRewardPaoDianRewardCfg(paoDianRewardId)
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypePaoDian)
	addItems := make(gamedb.ItemInfos, 0)
	for _, info := range paoDianRewardCfg.Reward {
		addItems = append(addItems, &gamedb.ItemInfo{
			ItemId: info.ItemId,
			Count:  info.Count * times,
		})
	}
	this.GetBag().AddItems(user, addItems, op)
	this.GetUserManager().SendItemChangeNtf(user, op)
}

/**
 *  @Description: 更新最后进入高倍泡点时间
 *  @param user
 */
func (this *PaoDianManager) UpdateEndTime(user *objs.User) {
	user.PaoDian.EndTime = int(time.Now().Unix())
	user.Dirty = true
}
