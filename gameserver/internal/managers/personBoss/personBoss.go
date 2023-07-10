package personBoss

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"time"
)

type PersonBoss struct {
	util.DefaultModule
	managersI.IModule
}

func NewPersonBoss(module managersI.IModule) *PersonBoss {
	p := &PersonBoss{IModule: module}
	return p
}

func (this *PersonBoss) Online(user *objs.User) {
	date := common.GetResetTime(time.Now())
	this.ResetPersonBoss(user, date)
}

func (this *PersonBoss) ResetPersonBoss(user *objs.User, date int) {
	userPersonBoss := user.PersonBosses
	if userPersonBoss.ResetTime != date {
		userPersonBoss.ResetTime = date
		//userPersonBoss.DareNum = make(model.IntKv)
		cfgs := gamedb.GetSingleBossBaseCfgs()
		for _, cfg := range cfgs {
			this.delBossKillNum(user.Id, cfg.StageId)
		}
	}
}

func (this *PersonBoss) delBossKillNum(userId, stageId int) {
	rmodel.Boss.DelBossKillNum(userId, constFight.FIGHT_TYPE_PERSON_BOSS, stageId)
}
func (this *PersonBoss) GetBossKillNum(userId, stageId int) int {
	return rmodel.Boss.GetBossKillNum(userId, constFight.FIGHT_TYPE_PERSON_BOSS, stageId)
}
func (this *PersonBoss) writeBossKillNum(userId, stageId, value int) {
	rmodel.Boss.SetBossKillNum(userId, constFight.FIGHT_TYPE_PERSON_BOSS, stageId, value)
}
