package objs

import (
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/prop"
)

type Hero struct {
	*modelGame.Hero
	Prop                              *prop.Prop
	ModuleCombat                      map[int]int
	ConditionAttribute                map[int]map[int]int         //条件属性
	AncientTreasureConditionAttribute map[int]map[int]map[int]int //远古宝物  条件属性 k:哪个模块属性 k1:宝物id k2:属性id

	ZodiacSuit                []int //生肖套装效果id
	DictateSuit               []int //主宰装备套装效果id
	InsideEffects             []int //内功技能效果id
	HolyEffects               []int //神兵技能效果id
	RingPhantomEffects        []int //特戒幻灵效果id
	WingEffects               []int //神翼技能效果id
	VipEffects                []int //vip特权效果id
	HolyBeastEffects          []int //圣兽技能效果id
	TalentEffects             []int //天赋效果id
	JuexueEffects             []int //绝学效果id
	ChuanShiEquipEffects      []int //传世装备套装效果id
	AncientSkillEffects       []int //远古神技效果id
	ChuanShiStrengthenEffects []int //传世强化套装效果id
	MiJiEffects               []int //秘籍效果id
	PetAppendageEffects       []int //战宠附体技能效果id
	LabelEffects              []int //头衔效果id
	MonthCardEffects          []int //月卡效果id
	PrivilegeEffects          []int //特权效果id

	MapEffects map[int]map[int]int //指定地图生效效果id(地体id,effectId,0)

	TalentGeneral map[int]map[int]int //通用天赋
}

func NewHero(hero *modelGame.Hero) *Hero {
	h := &Hero{
		Hero:                              hero,
		Prop:                              prop.NewProp(),
		ModuleCombat:                      make(map[int]int),
		ConditionAttribute:                make(map[int]map[int]int),
		AncientTreasureConditionAttribute: make(map[int]map[int]map[int]int),
	}
	return h
}
