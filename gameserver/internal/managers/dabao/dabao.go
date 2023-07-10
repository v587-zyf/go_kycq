package dabao

import (
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/util"
)

func NewDaBaoManager(m managersI.IModule) *DaBao {
	return &DaBao{
		IModule: m,
	}
}

type DaBao struct {
	util.DefaultModule
	managersI.IModule
}

func (this *DaBao) Online(user *objs.User) {
	this.ResumeEnergy(user)
}
