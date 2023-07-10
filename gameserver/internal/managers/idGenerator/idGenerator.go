package idGenerator

import (
	"cqserver/gameserver/internal/managersI"
	"cqserver/golibs/util"
	"sync"
)

func NewIdGeneratorManager(module managersI.IModule) *IdGeneratorManager {
	return &IdGeneratorManager{IModule: module}
}

type IdGeneratorManager struct {
	util.DefaultModule
	managersI.IModule
	sync.RWMutex
	//世界拍卖行MAXid
	worldMaxId int
	//门派拍卖行MAXid
	guildMaxId int
}

func (this *IdGeneratorManager) InitWorldNowId(nowId, guildNowId int) {
	this.worldMaxId = nowId
	this.guildMaxId = guildNowId
}

func (this *IdGeneratorManager) GetNextWorldId() int {
	this.Lock()
	defer this.Unlock()
	this.worldMaxId++
	return this.worldMaxId
}

func (this *IdGeneratorManager) GetNextGuildId() int {
	this.Lock()
	defer this.Unlock()
	this.guildMaxId++
	return this.guildMaxId
}