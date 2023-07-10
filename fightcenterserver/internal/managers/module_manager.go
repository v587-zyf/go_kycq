package managers

import (
	"cqserver/fightcenterserver/internal/base"
	"cqserver/golibs/dbmodel"
	"cqserver/golibs/util"
)

type ModuleManager struct {
	*util.DefaultModuleManager

	gateManager *GateManager
	gsManager   *GSManager
	fsManager   *FSManager
}

var m = &ModuleManager{
	DefaultModuleManager: util.NewDefaultModuleManager(),
}

func Get() *ModuleManager {
	return m
}

func (this *ModuleManager) init() error {
	var err error

	err = dbmodel.InitDb(base.Conf, nil, "", nil)
	if err != nil {
		return err
	}

	return nil
}

func (this *ModuleManager) Init() error {
	err := this.init()
	if err != nil {
		return err
	}
	this.gateManager = this.AppendModule(NewGateManager()).(*GateManager)
	this.gsManager = this.AppendModule(NewGSManager()).(*GSManager)
	this.fsManager = this.AppendModule(NewFSManager()).(*FSManager)

	return this.DefaultModuleManager.Init()
}

func (this *ModuleManager) Start() error {
	return this.DefaultModuleManager.Start()
}
