package base

import (
	"cqserver/gamelibs/prop"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"fmt"
	"sync"
)

type ActorProp struct {
	hpNow         int         //当前血量
	mpNow         int         //当前蓝量
	baseProp      *prop.Prop  //玩家基础属性
	buffPropLess  map[int]int //玩家buff改变的属性 减少
	buffPropAdd   map[int]int //玩家buff改变的属性 增加
	skillTempProp map[int]int //本回合技能临时属性
	propMu        sync.RWMutex
}

func NewActorProp() *ActorProp {
	actorProp := &ActorProp{
		baseProp:      prop.NewProp(),
		buffPropLess:  make(map[int]int),
		buffPropAdd:   make(map[int]int),
		skillTempProp: make(map[int]int),
	}
	return actorProp
}

func (this *ActorProp) GetActorBaseProp(id int) int {
	this.propMu.RLock()
	defer func() {
		this.propMu.RUnlock()
	}()
	return this.baseProp.Get(id)
}

/**
 *  @Description: 获取单属性值
 *  @param id	属性Id
 *  @param totalValue 是否最终值（含百分比加成）
 *  @return int
 */
func (this *ActorProp) Get(id int) int {
	if id == pb.PROPERTY_COMBAT {
		return this.baseProp.Combat
	}
	this.propMu.RLock()
	defer func() {
		this.propMu.RUnlock()
	}()
	return this.baseProp.Get(id) + this.buffPropLess[id] + this.buffPropAdd[id] + this.skillTempProp[id]
}

/**
 *  @Description: buff属性改变
 *  @param id
 *  @param change
 */
func (this *ActorProp) BuffPropChange(id int, change int, isRecover bool) {
	this.propMu.Lock()
	defer func() {
		this.propMu.Unlock()
	}()
	if change > 0 {

		if isRecover {
			this.buffPropLess[id] += change
			if this.buffPropLess[id] > 0 {
				this.buffPropLess[id] = 0
			}
		} else {
			this.buffPropAdd[id] += change
		}

	} else {

		if isRecover {
			this.buffPropAdd[id] += change
			if this.buffPropAdd[id] <= 0 {
				this.buffPropAdd[id] = 0
			}
		} else {
			this.buffPropLess[id] += change
		}

	}
}

func (this *ActorProp) SkillTempPropChange(propId, value int) {

	this.propMu.Lock()
	defer func() {
		this.propMu.Unlock()
	}()
	this.skillTempProp[propId] += value

}

func (this *ActorProp) SkillTempPropReset() {
	this.propMu.Lock()
	defer func() {
		this.propMu.Unlock()
	}()
	this.skillTempProp = make(map[int]int)
}

func (this *ActorProp) SkillTempPropToString() string {
	return fmt.Sprintf("技能临时属性：%v", this.skillTempProp)
}

func (this *ActorProp) ToString(propIds []int) string {

	propIdMaps := make(map[int]bool, 0)
	for _, v := range propIds {
		propIdMaps[v] = true
	}
	propstr := fmt.Sprintf("战力：%v，当前血量：%v,最大血量：%v;当前蓝量：%v,最大蓝量：%v；", this.baseProp.Combat, this.HpNow(), this.Get(pb.PROPERTY_HP), this.MpNow(), this.Get(pb.PROPERTY_MP))
	for _, v := range pb.PROPERTY_ARRAY {
		if propIds != nil && len(propIds) > 0 {
			if propIdMaps[v] {
				propstr += fmt.Sprintf("%v:%v;", v, this.Get(v))
			}
		} else {
			if prop := this.Get(v); prop > 0 {
				propstr += fmt.Sprintf("%v:%v;", v, prop)
			}
		}
	}
	return propstr
}

func (this *ActorProp) AddOne(propId, value int) {
	this.propMu.RLock()
	defer func() {
		this.propMu.RUnlock()
	}()
	this.baseProp.AddOne(propId, value)
}

func (this *ActorProp) Add(props map[int]int) {
	this.propMu.RLock()
	defer func() {
		this.propMu.RUnlock()
	}()
	this.baseProp.Add(props)
}

func (this *ActorProp) Calc(job int) {
	this.propMu.RLock()
	defer func() {
		this.propMu.RUnlock()
	}()
	this.baseProp.Calc(job)
}

func (this *ActorProp) Combat() int {
	return this.baseProp.Combat
}
func ( this *ActorProp) Reset(){
	this.propMu.RLock()
	defer func() {
		this.propMu.RUnlock()
	}()
	this.baseProp.Reset()
}

func (this *ActorProp) MpNow() int {
	return this.mpNow
}

func (this *ActorProp) SetMpNow(mpNow int) {
	this.mpNow = mpNow
}

func (this *ActorProp) HpNow() int {
	return this.hpNow
}

func (this *ActorProp) SetHpNow(hpNow int) {
	this.hpNow = hpNow
}

func (this *ActorProp) DecHP(value int) int {

	now := this.HpNow()
	changeHp := value
	if now+value < 0 {
		changeHp = -now
	}
	this.SetHpNow(now + changeHp)
	return changeHp
}

func (this *ActorProp) AddHP(add int) int {

	max := this.Get(pb.PROPERTY_HP)
	now := this.HpNow()
	changeHp := add
	if now+add > max {
		changeHp = max - now
	}
	this.SetHpNow(now + changeHp)
	return changeHp
}

func (this *ActorProp) ByFightActorProp(prop *pbserver.ActorProp) {
	this.propMu.RLock()
	defer func() {
		this.propMu.RUnlock()
	}()
	this.baseProp.ByFightActorProp(prop)
}
