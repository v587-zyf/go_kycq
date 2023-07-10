package vipBoss

import (
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"time"
)

func NewVipBossManager(module managersI.IModule) *VipBossManager {
	return &VipBossManager{IModule: module}
}

type VipBossManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *VipBossManager) OnLine(user *objs.User) {
	date := common.GetResetTime(time.Now())
	this.ResetVipBossFightNum(user, date)
}

func (this *VipBossManager) ResetVipBossFightNum(user *objs.User, date int) {
	userVipBoss := user.VipBosses
	if userVipBoss.ResetTime != date {
		userVipBoss.ResetTime = date
		for stageId := range userVipBoss.DareNum {
			userVipBoss.DareNum[stageId] = 0
		}
	}
}
