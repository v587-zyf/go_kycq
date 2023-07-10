package prop

import (
	//	"fmt"

	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pbserver"
	"fmt"

	"cqserver/protobuf/pb"
)

const PropBaseRate = 10000.0

type PropEntry struct {
	Total       int
	Base        int
	BasePercent int
}

//calc 单个属性的总值
func (this *PropEntry) calc() {
	base := float64(this.Base) * (1.0 + float64(this.BasePercent)/PropBaseRate)
	this.Base = int(base)
}

var PropEntryType = map[int]bool{
	pb.PROPERTY_HP:       true,
	pb.PROPERTY_PATT_MAX: true,
	pb.PROPERTY_PATT_MIN: true,
	pb.PROPERTY_MATT_MAX: true,
	pb.PROPERTY_MATT_MIN: true,
	pb.PROPERTY_TATT_MAX: true,
	pb.PROPERTY_TATT_MIN: true,
	pb.PROPERTY_DEF_MAX:  true,
	pb.PROPERTY_DEF_MIN:  true,
	pb.PROPERTY_ADF_MAX:  true,
	pb.PROPERTY_ADF_MIN:  true,
	pb.PROPERTY_HIT:      true,
	pb.PROPERTY_MISS:     true,
}

type Prop struct {
	Combat      int
	CombatFixed int //固定直接战力

	props     map[int]*PropEntry //参与百分比计算
	otherProp map[int]int        //不参与百分比计算
}

func NewProp() *Prop {
	p := &Prop{}

	p.Reset()
	return p
}

func (this *Prop) Reset() {
	this.props = make(map[int]*PropEntry)
	for t := range PropEntryType {
		this.props[t] = &PropEntry{}
	}
	this.otherProp = make(map[int]int)
	this.Combat = 0
	this.CombatFixed = 0
}

//增加一个全属性
func (this *Prop) AddAllProp(prop *Prop) {

	this.Add(prop.baseValueMap())
	this.CombatFixed += prop.CombatFixed
}

func (this *Prop) Add(props map[int]int) {
	for k, v := range props {
		this.AddOne(k, v)
	}
}

func (this *Prop) Calc(job int) {
	this.calcEach(job)
	this.calcCombat(job)
}

/**
 *  @Description: 获取单属性值(各个模块提供的属性)
 *  @param id	属性Id
 *  @return int
 */
func (this *Prop) getBaseValue(id int) int {
	if id == pb.PROPERTY_COMBAT {
		return this.Combat
	}
	if _, ok := PropEntryType[id]; ok {
		p := this.props[id]
		if p == nil {
			return 0
		}
		return p.Base
	}
	if p, ok := this.otherProp[id]; ok {
		return p
	}
	return 0
}

/**
 *  @Description: 获取单属性值
 *  @param id	属性Id
 *  @param totalValue 是否最终值（含百分比加成）
 *  @return int
 */
func (this *Prop) Get(id int) int {
	if id == pb.PROPERTY_COMBAT {
		return this.Combat
	}
	if _, ok := PropEntryType[id]; ok {
		p := this.props[id]
		if p == nil {
			return 0
		}
		return p.Total
	}
	if p, ok := this.otherProp[id]; ok {
		return p
	}
	return 0
}

func (this *Prop) baseValueMap() map[int]int {
	m := make(map[int]int, len(pb.PROPERTY_ARRAY))
	for _, pId := range pb.PROPERTY_ARRAY {
		if pId == pb.PROPERTY_COMBAT {
			continue
		}
		v := this.getBaseValue(pId)
		if v > 0 {
			m[pId] = v
		}
	}
	return m
}

func (this *Prop) BuildClient() map[int32]int64 {
	m := make(map[int32]int64, len(pb.PROPERTY_ARRAY))
	for _, pId := range pb.PROPERTY_ARRAY {
		if pId == pb.PROPERTY_COMBAT {
			continue
		}
		if _, ok := PropEntryType[pId]; ok {
			p := this.props[pId]
			if p != nil {
				m[int32(pId)] = int64(p.Total)
			}
		}
		if p, ok := this.otherProp[pId]; ok {
			m[int32(pId)] = int64(p)
		}
	}
	m[int32(pb.PROPERTY_COMBAT)] = int64(this.Combat)
	return m
}

func (this *Prop) calcCombatForMaxMin(max, min int, combat float64) float64 {
	return float64(min+max) / 2 * combat
}

func (this *Prop) calcCombat(job int) {

	combat := 0.0
	hpFix := 0.0
	if job == pb.JOB_ZHANSHI {
		combat += this.calcCombatForMaxMin(this.Get(pb.PROPERTY_PATT_MAX), this.Get(pb.PROPERTY_PATT_MIN), gamedb.GetPropertyPropertyCfg(pb.PROPERTY_PATT_MAX).Combat[0])
		hpFix = gamedb.GetPropertyPropertyCfg(pb.PROPERTY_HP).Combat[0]
	} else if job == pb.JOB_FASHI {
		combat += this.calcCombatForMaxMin(this.Get(pb.PROPERTY_MATT_MAX), this.Get(pb.PROPERTY_MATT_MIN), gamedb.GetPropertyPropertyCfg(pb.PROPERTY_MATT_MAX).Combat[0])
		hpFix = gamedb.GetPropertyPropertyCfg(pb.PROPERTY_HP).Combat[1]
	} else {
		combat += this.calcCombatForMaxMin(this.Get(pb.PROPERTY_TATT_MAX), this.Get(pb.PROPERTY_TATT_MIN), gamedb.GetPropertyPropertyCfg(pb.PROPERTY_TATT_MAX).Combat[0])
		hpFix = gamedb.GetPropertyPropertyCfg(pb.PROPERTY_HP).Combat[2]
	}
	//防御
	combat += this.calcCombatForMaxMin(this.Get(pb.PROPERTY_DEF_MAX), this.Get(pb.PROPERTY_DEF_MIN), gamedb.GetPropertyPropertyCfg(pb.PROPERTY_DEF_MAX).Combat[0])
	//魔法防御
	combat += this.calcCombatForMaxMin(this.Get(pb.PROPERTY_ADF_MAX), this.Get(pb.PROPERTY_ADF_MIN), gamedb.GetPropertyPropertyCfg(pb.PROPERTY_ADF_MAX).Combat[0])
	for _, pId := range pb.PROPERTY_ARRAY {
		if pId == pb.PROPERTY_COMBAT || pId == pb.PROPERTY_PATT_MAX || pId == pb.PROPERTY_PATT_MIN ||
			pId == pb.PROPERTY_MATT_MAX || pId == pb.PROPERTY_MATT_MIN ||
			pId == pb.PROPERTY_TATT_MAX || pId == pb.PROPERTY_TATT_MIN ||
			pId == pb.PROPERTY_DEF_MAX || pId == pb.PROPERTY_DEF_MIN ||
			pId == pb.PROPERTY_ADF_MAX || pId == pb.PROPERTY_ADF_MIN {
			continue
		}
		propertiesConf := gamedb.GetPropertyPropertyCfg(pId)
		if propertiesConf == nil {
			continue
		}
		if len(propertiesConf.Combat) == 0 {
			continue
		}
		v := this.Get(pId)
		if v < 1 {
			continue
		}
		//其他 的数值 * 表中的数值
		if pId == pb.PROPERTY_HP {
			combat += float64(hpFix) * float64(v)
		} else {
			combat += float64(propertiesConf.Combat[0]) * float64(v)
		}
	}
	this.Combat = int(combat) + this.CombatFixed
}

func (this *Prop) calcEach(job int) {

	if value, ok := this.otherProp[pb.PROPERTY_ATT_ALL]; ok {
		maxPropId, minPropId := GetAtkPropIdByJob(job)
		this.props[maxPropId].Base += value
		this.props[minPropId].Base += value
		this.otherProp[pb.PROPERTY_ATT_ALL] = 0
	}
	if value, ok := this.otherProp[pb.PROPERTY_ATT_MODULE_RATE_ALL]; ok {
		maxPropId, minPropId := GetAtkPropIdByJob(job)
		this.props[maxPropId].BasePercent += value
		this.props[minPropId].BasePercent += value
		this.otherProp[pb.PROPERTY_ATT_MODULE_RATE_ALL] = 0
	}
	if value, ok := this.otherProp[pb.PROPERTY_DEF_ALL]; ok {
		this.props[pb.PROPERTY_DEF_MAX].Base += value
		this.props[pb.PROPERTY_DEF_MIN].Base += value
		this.otherProp[pb.PROPERTY_DEF_ALL] = 0
	}
	if value, ok := this.otherProp[pb.PROPERTY_ADF_ALL]; ok {
		this.props[pb.PROPERTY_ADF_MAX].Base += value
		this.props[pb.PROPERTY_ADF_MIN].Base += value
		this.otherProp[pb.PROPERTY_ADF_ALL] = 0
	}

	if value, ok := this.otherProp[pb.PROPERTY_ATT_RATE_ALL]; ok {
		if job == pb.JOB_ZHANSHI {
			this.otherProp[pb.PROPERTY_PATT_RATE] += value
		} else if job == pb.JOB_FASHI {
			this.otherProp[pb.PROPERTY_MATT_RATE] += value
		} else {
			this.otherProp[pb.PROPERTY_TATT_RATE] += value
		}
		this.otherProp[pb.PROPERTY_ATT_RATE_ALL] = 0
	}

	for k, v := range this.props {
		v.calc()
		percentId := k/10*10 + 2
		percent := this.otherProp[percentId]

		v.Total = int(float64(v.Base) * (1 + float64(percent)/PropBaseRate))
	}

	//this.props[pb.PROPERTY_ATT_ALL].Total = 0
}

func (this *Prop) AddOne(id, value int) {

	fixId := id / 10 * 10
	fix := id % 10
	_, ok := PropEntryType[fixId]
	if fix == 3 && ok {
		if _, ok := PropEntryType[fixId]; ok {
			if this.props[fixId] == nil {
				this.props[fixId] = &PropEntry{}
			}
			this.props[fixId].BasePercent += value
		}
	} else {
		if _, ok := PropEntryType[id]; ok {
			if this.props[id] == nil {
				this.props[id] = &PropEntry{}
			}
			this.props[id].Base += value
		} else {
			this.otherProp[id] += value
		}
	}

}

func (this *Prop) ToFightActorProp() *pbserver.ActorProp {

	props := make(map[int32]int64)
	for _, v := range pb.PROPERTY_ARRAY {
		if v%10 == 2 || v%10 == 3 {
			continue
		}
		props[int32(v)] = int64(this.Get(v))

	}
	props[pb.PROPERTY_COMBAT] = int64(this.Combat)
	actorProp := &pbserver.ActorProp{
		Props: props,
	}
	return actorProp
}

func (this *Prop) ByFightActorProp(prop *pbserver.ActorProp) {

	this.Reset()
	this.Combat = int(prop.Props[pb.PROPERTY_COMBAT])
	for k, v := range prop.Props {
		this.AddOne(int(k), int(v))
		//此处只是为了把base赋值给total
		if _, ok := PropEntryType[int(k)]; ok {
			if this.props[int(k)] != nil {
				this.props[int(k)].Total = this.props[int(k)].Base
			}
		}
	}
}

//---------------------------------------------------------------------------------
//---------------------------------------------------------------------------------
//---------------------------------------------------------------------------------
//以下方法只为了分析

func (this *Prop) AnalyzePercent() string {
	propStr := ""
	for k, v := range this.props {
		propStr += fmt.Sprintf("属性：[%v:%v] 值：%v,base:%v，rate:%v \n", k, gamedb.GetPropertyPropertyCfg(k).Name, v.Total, v.Base, v.BasePercent)
	}
	for k, v := range this.otherProp {
		propStr += fmt.Sprintf("属性：[%v:%v] 值：%v \n", k, gamedb.GetPropertyPropertyCfg(k).Name, v)
	}
	logger.Debug(propStr)
	return propStr
}

//根据职业获取对应 攻击属性Id(最高 最低)
func GetAtkPropIdByJob(job int) (int, int) {

	if job == pb.JOB_DAOSHI {
		return pb.PROPERTY_TATT_MAX, pb.PROPERTY_TATT_MIN
	} else if job == pb.JOB_FASHI {
		return pb.PROPERTY_MATT_MAX, pb.PROPERTY_MATT_MIN
	} else {
		return pb.PROPERTY_PATT_MAX, pb.PROPERTY_PATT_MIN
	}
}

func GetDefPropIdByJob(job int) (int, int) {

	if job == pb.JOB_ZHANSHI {
		return pb.PROPERTY_DEF_MIN, pb.PROPERTY_DEF_MAX
	} else {
		return pb.PROPERTY_ADF_MIN, pb.PROPERTY_ADF_MAX
	}
}
