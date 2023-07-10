package official

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

type OfficialManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewOfficialManager(module managersI.IModule) *OfficialManager {
	return &OfficialManager{IModule: module}
}

func (this *OfficialManager) Init() error {
	return nil
}

/**
 *  @Description: 官职升级
 *  @param user
 *  @param op
 *  @return error
 */
func (this *OfficialManager) OfficialUpLevel(user *objs.User, op *ophelper.OpBagHelperDefault) error {
	officialConf := gamedb.GetOfficialOfficialCfg(user.Official)
	if officialConf == nil {
		return gamedb.ERRSETTINGNOTFOUND
	}
	nextLvOfficialConf := gamedb.GetOfficialOfficialCfg(user.Official + 1)
	if nextLvOfficialConf == nil {
		return gamedb.ERRMAXLV
	}

	ispass := this.GetCondition().CheckMulti(user, -1, officialConf.Condition)
	if !ispass {
		return gamedb.ERRCONDITION
	}
	if officialConf.Consume.ItemId != 0 {
		err := this.GetBag().Remove(user, op, officialConf.Consume.ItemId, officialConf.Consume.Count)
		if err != nil {
			return err
		}
	}
	user.Official += 1
	kyEvent.OfficialLvUp(user, user.Official)
	this.GetTask().AddTaskProcess(user, pb.CONDITION_UPGRADE_GUAN_XIAN, -1)
	this.GetUserManager().UpdateCombat(user, -1)
	this.GetCondition().RecordCondition(user, pb.CONDITION_UPGRADE_GUAN_XIAN, []int{})
	return nil
}
