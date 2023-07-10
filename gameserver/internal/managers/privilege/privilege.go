package privilege

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

func NewPrivilegeManager(m managersI.IModule) *PrivilegeManager {
	return &PrivilegeManager{IModule: m}
}

type PrivilegeManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *PrivilegeManager) Online(user *objs.User) {
	this.updatePrivilegeEffect(user)
}

/**
 *  Buy
 *  @Description: 购买特权
 *  @param user
 *  @param privilegeId	特权id
 *  @param op
 *  @return error
**/
func (this *PrivilegeManager) Buy(user *objs.User, privilegeId int, op *ophelper.OpBagHelperDefault) error {
	privilegeCfg := gamedb.GetPrivilegePrivilegeCfg(privilegeId)
	if privilegeCfg == nil {
		return gamedb.ERRPARAM
	}
	userPrivilege := user.Privilege
	if _, ok := userPrivilege[privilegeId]; ok {
		return gamedb.ERRREPEATBUY
	}
	if privilegeCfg.Need != 0 {
		if _, ok := userPrivilege[privilegeCfg.Need]; !ok {
			return gamedb.ERRCONDITION
		}
	}
	if err := this.GetBag().Remove(user, op, privilegeCfg.Cost.ItemId, privilegeCfg.Cost.Count); err != nil {
		return err
	}

	user.Privilege[privilegeId] = int(time.Now().Unix())
	this.GetBag().AddItems(user, privilegeCfg.Reward, op)

	this.GetAnnouncement().SendSystemChat(user, pb.SCROLINGTYPE_PRIVILEGE_OPEN, privilegeId, -1)
	this.updatePrivilegeEffect(user)
	return nil
}

/**
 *  GetPrivilege
 *  @Description: 获取特权
 *  @param user
 *  @param privilegeId 特权id
**/
func (this *PrivilegeManager) GetPrivilege(user *objs.User, privilegeId int) int {
	num := 0
	userPrivilegeId := user.Privilege
	privilegeCfgs := gamedb.GetPrivilegeCfgs()
	for id, cfg := range privilegeCfgs {
		if _, ok := userPrivilegeId[id]; !ok {
			continue
		}
		num += cfg.Privilege[privilegeId]
	}
	return num
}

func (this *PrivilegeManager) ItemActiveCheck(user *objs.User, privilegeId int) error {
	if privilegeCfg := gamedb.GetPrivilegePrivilegeCfg(privilegeId); privilegeCfg == nil {
		return gamedb.ERRPARAM
	}
	if _, ok := user.Privilege[privilegeId]; ok {
		return gamedb.ERRREPEATBUY
	}
	return nil
}

/**
 *  ItemActive
 *  @Description: 用道具激活
 *  @param user
 *  @param privilegeId	特权id
 *  @return error
**/
func (this *PrivilegeManager) ItemActive(user *objs.User, privilegeId int) error {
	privilegeCfg := gamedb.GetPrivilegePrivilegeCfg(privilegeId)
	user.Privilege[privilegeId] = int(time.Now().Unix())
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypePrivilegeItemActive)
	this.GetBag().AddItems(user, privilegeCfg.Reward, op)
	this.GetAnnouncement().SendSystemChat(user, pb.SCROLINGTYPE_PRIVILEGE_OPEN, privilegeId, -1)
	this.updatePrivilegeEffect(user)
	this.GetUserManager().SendMessage(user, &pb.PrivilegeBuyAck{PrivilegeId: int32(privilegeId), Goods: op.ToChangeItems()}, true)
	this.GetUserManager().SendItemChangeNtf(user, op)
	return nil
}

func (this *PrivilegeManager) updatePrivilegeEffect(user *objs.User) {
	effect := this.GetPrivilege(user, pb.VIPPRIVILEGE_ATTR)
	if effect != 0 {
		effectSlice := make([]int, 0)
		effectSlice = append(effectSlice, effect)
		for _, hero := range user.Heros {
			hero.PrivilegeEffects = effectSlice
		}
	}
	this.GetUserManager().UpdateCombat(user, -1)
}
