package offline

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

type OfflineManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewOfflineManager(module managersI.IModule) *OfflineManager {
	return &OfflineManager{IModule: module}
}

func (this *OfflineManager) GetAward(user *objs.User, ack *pb.OfflineAwardGetAck, op *ophelper.OpBagHelperDefault) error {
	if user.OfflineAwardMark {
		return gamedb.ERRAWARDGET
	}
	offlineTime := time.Now().Unix() - user.OfflineTime.Unix()
	if offlineTime < 60 {
		return gamedb.ERROFFLINE
	}
	minutes := offlineTime / 60
	if offlineTime > int64(gamedb.GetConf().OffLine)*60 {
		minutes = int64(gamedb.GetConf().OffLine)
	}
	stageCfg := gamedb.GetHookMapHookMapCfg(user.StageId)
	count := minutes * int64(stageCfg.Name)

	monthCardPrivilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_HANGUP_EXP)
	if monthCardPrivilege != 0 {
		count = int64(common.CalcTenThousand(monthCardPrivilege, int(count)))
	}

	this.GetBag().AddItem(user, op, pb.ITEMID_EXP, int(count))
	user.OfflineTime = time.Now()
	user.OfflineAwardMark = true
	user.Dirty = true

	ack.Goods = op.ToChangeItems()
	ack.IsGet = true
	return nil
}

//自动领取离线奖励处理
func (this *OfflineManager) AutoGetAward(user *objs.User) {
	if user.OfflineAwardMark {
		return
	}
	if user.OfflineTime.Unix() <= 0 {
		return
	}

	offlineTime := time.Now().Unix() - user.OfflineTime.Unix()
	if offlineTime < 60 {
		return
	}
	minutes := offlineTime / 60
	if offlineTime > int64(gamedb.GetConf().OffLine)*60 {
		minutes = int64(gamedb.GetConf().OffLine)
	}
	stageCfg := gamedb.GetHookMapHookMapCfg(user.StageId)
	count := minutes * int64(stageCfg.Name)
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeOfflineReward)
	this.GetBag().AddItem(user, op, pb.ITEMID_EXP, int(count))
	user.OfflineTime = time.Now()
	user.OfflineAwardMark = true
	user.Dirty = true
	return
}
