package combat

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/prop"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gamelibs/publicCon/constEquip"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"math"
	"sort"
	"time"
)

type funcCalcProp func(user *objs.User, heroIndex int) *prop.Prop
type funcCalcHeroProp func(user *objs.User, heroIndex int) *prop.Prop

// 武将共享模块战力
var shareModules = make(map[int]funcCalcProp)

// 武将独立模块战力
var heroModules = make(map[int]funcCalcHeroProp)

func init() {
	shareModules[pb.PROPERTYMODULE_FABAO] = calcFabaoLvProps
	shareModules[pb.PROPERTYMODULE_REIN] = calcReinProps
	shareModules[pb.PROPERTYMODULE_ATLAS] = calcAtlasProps
	shareModules[pb.PROPERTYMODULE_ATLAS_WEAR] = calcAtlasWearProps
	shareModules[pb.PROPERTYMODULE_OFFICIAL] = calcOfficialProps
	shareModules[pb.PROPERTYMODULE_HOLYARMS] = calcHolyarmsProps
	shareModules[pb.PROPERTYMODULE_PANACEA] = calcPanaceaProps
	shareModules[pb.PROPERTYMODULE_PET] = calcPetProps
	shareModules[pb.PROPERTYMODULE_JUEXUE] = calcJuexueProps
	shareModules[pb.PROPERTYMODULE_FIT] = calcFitProps
	shareModules[pb.PROPERTYMODULE_MONTH_CARD] = calcMonthCardProps
	shareModules[pb.PROPERTYMODULE_ACHIEVEMENT] = calcAchievementProps
	shareModules[pb.PROPERTYMODULE_ELF] = calcElfProps
	shareModules[pb.PROPERTYMODULE_FIT_HOLY_EQUIP] = calcFitHolyEquipProps
	shareModules[pb.PROPERTYMODULE_ANCIENT_SKILL] = calcAncientSkillProps
	shareModules[pb.PROPERTYMODULE_TITLE] = calcTitleProps
	shareModules[pb.PROPERTYMODULE_MI_JI] = calcMiJiProps
	shareModules[pb.PROPERTYMODULE_YUAN_GU_BAO_WU] = calcAncientTreasure
	shareModules[pb.PROPERTYMODULE_PET_APPENDAGE] = calcPetAppendage
	shareModules[pb.PROPERTYMODULE_DABAO_EQUIP] = calcDaBaoEquipProps
	shareModules[pb.PROPERTYMODULE_LABEL] = calcLabelProps
	shareModules[pb.PROPERTYMODULE_PRIVILEGE] = calcPrivilegeProps

	heroModules[pb.PROPERTYMODULE_EQUIP_STRENGTH] = calcEquipStrengthProps
	heroModules[pb.PROPERTYMODULE_EQUIP_NORMAL] = calcEquipNormalProps
	heroModules[pb.PROPERTYMODULE_WING] = calcWingProps
	heroModules[pb.PROPERTYMODULE_DRAGON_EQUIP] = calcDragonEquipProps
	heroModules[pb.PROPERTYMODULE_KINGARMS] = calcKingarmsProps
	heroModules[pb.PROPERTYMODULE_ZODIAC] = calcZodiacProps
	heroModules[pb.PROPERTYMODULE_CLEAR] = calcClearProps
	heroModules[pb.PROPERTYMODULE_DICTATE] = calcDictateProps
	heroModules[pb.PROPERTYMODULE_JEWEL] = calcJewelProps
	heroModules[pb.PROPERTYMODULE_FASHION] = calcFashionWeaponProps
	heroModules[pb.PROPERTYMODULE_FASHION_CLOTH] = calcFashionClothProps
	heroModules[pb.PROPERTYMODULE_INSIDE] = calcInsideProps
	heroModules[pb.PROPERTYMODULE_RING] = calcRingProps
	heroModules[pb.PROPERTYMODULE_GOD_EEQUIP] = calcGodEquipProps
	heroModules[pb.PROPERTYMODULE_AREA] = calcAreaProps
	heroModules[pb.PROPERTYMODULE_EXP_LV] = calcExpPoolProps
	heroModules[pb.PROPERTYMODULE_MAGIC_CIRCLE] = calcMagicCircleProps
	heroModules[pb.PROPERTYMODULE_TALENT] = calcTalentProps
	heroModules[pb.PROPERTYMODULE_HOLY_BEAST] = calcHolyBeastProps
	heroModules[pb.PROPERTYMODULE_TALENT_GENERAL] = calcTalentGeneralProps
	heroModules[pb.PROPERTYMODULE_VIP] = calcVipProps
	heroModules[pb.PROPERTYMODULE_SKILL] = calcSkillProps
	heroModules[pb.PROPERTYMODULE_CHUAN_SHI_EQUIP] = calcChuanShiEquipProps
	heroModules[pb.PROPERTYMODULE_CHUAN_SHI_STRENGTHEN] = calcChuanShiStrengthen
	heroModules[pb.PROPERTYMODULE_GOD_EQUIP_BLOOD] = calcGodEquipBloodProps
}

func UpdatePropCombat(user *objs.User, heroIndex int) {

	genUserProps(user, heroIndex)
	user.MarkSyncStatus(objs.SyncStatusComatNtf)
}

func genUserProps(user *objs.User, heroIndex int) {
	//propsModules := make(map[int]*prop.Prop)

	var heros []*objs.Hero
	if heroIndex != -1 {
		heros = append(heros, user.Heros[heroIndex])
	} else {
		for _, v := range user.Heros {
			heros = append(heros, v)
		}
	}

	userCombat := 0
	otherAddCombat := user.PetCombat
	for _, v := range heros {
		v.Prop.Reset()
		for k, fun := range shareModules {
			props := fun(user, v.Index)
			props.Calc(v.Job)
			v.Prop.AddAllProp(props)
			v.ModuleCombat[k] = props.Combat
		}

		for k, heroModuleCalc := range heroModules {
			prop := heroModuleCalc(user, v.Index)
			prop.Calc(v.Job)
			v.Prop.AddAllProp(prop)
			v.ModuleCombat[k] = prop.Combat
		}
		v.Prop.Calc(v.Job)
		v.Combat = v.Prop.Combat
		userCombat += v.Combat
		logger.Debug("genUserProps user:%v,武将：%v,武将战力：%v,模块战力：%v,总的详细属性:%v", user.IdName(), v.Index, v.Combat, v.ModuleCombat)
	}

	user.Combat = userCombat + otherAddCombat
	user.Dirty = true
}

func calcTalentGeneral(num, percent int) int {
	return int(math.Ceil(float64(num) * (float64(percent) / 10000.0)))
}

func addEffect(effectMap map[int]int, p *prop.Prop) []int {
	effectSlice := make([]int, 0)
	if len(effectMap) > 0 {
		for effect := range effectMap {
			effectCfg := gamedb.GetEffectEffectCfg(effect)
			if effectCfg != nil {
				p.Add(effectCfg.Attribute)
				effectSlice = append(effectSlice, effect)
			}
		}
	}
	return effectSlice
}

func calcHolyarmsProps(user *objs.User, heroIndex int) *prop.Prop {
	hero := user.Heros[heroIndex]
	p := prop.NewProp()
	effectMap := make(map[int]int)
	addMap := make(map[int]int)
	for id, holyarm := range user.Holyarms {
		lvCfg := gamedb.GetHolyLvByHidAndLv(id, holyarm.Level)
		if lvCfg != nil {
			p.Add(lvCfg.Attribute)
			for pid, pVal := range lvCfg.Attribute {
				addMap[pid] += pVal
			}
		}
		for hlv, slv := range holyarm.Skill {
			skillCfg := gamedb.GetHolySkillByHidAndLv(id, hlv, slv)
			if skillCfg != nil {
				effectMap[skillCfg.Effect] = 0
			}
		}
	}
	talentEffect11 := hero.TalentGeneral[pb.PROPERTYMODULE_HOLYARMS][constConstant.TALENTEFFECT_11]
	if talentEffect11 != 0 {
		for k, v := range addMap {
			p.AddOne(k, calcTalentGeneral(v, talentEffect11))
		}
	}
	effectSlice := addEffect(effectMap, p)
	hero.HolyEffects = effectSlice
	return p
}

func calcJuexueProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	effectMap := make(map[int]int)
	for _, juexue := range user.Juexues {
		lvCfg := gamedb.GetJuexueLevelConfCfg(gamedb.GetRealId(juexue.Id, juexue.Lv))
		if lvCfg != nil {
			p.Add(lvCfg.Attribute)
			effectConf := gamedb.GetEffectEffectCfg(lvCfg.Buff)
			if effectConf != nil {
				if len(effectConf.Attribute) > 0 {
					p.Add(effectConf.Attribute)
				}
				effectMap[lvCfg.Buff] = 0
			}
		}
	}
	effectSlice := addEffect(effectMap, p)
	user.Heros[heroIndex].JuexueEffects = effectSlice
	return p
}

func calcFitProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	userFit := user.Fit
	for id, lv := range userFit.Lv {
		fitLevelCfg := gamedb.GetFitLevelFitLevelCfg(gamedb.GetRealId(id, lv))
		if fitLevelCfg != nil {
			p.Add(fitLevelCfg.Attribute)
		}
	}
	skillSlice := make([]int, 0)
	skillSlice = append(skillSlice, constFight.FIT_SKILL_ZHUDONG_ID)
	for _, skillId := range userFit.SkillBag {
		skillSlice = append(skillSlice, skillId)
	}
	for _, skillId := range skillSlice {
		fitSkill, ok := userFit.Skills[skillId]
		if ok {
			lvCfg := gamedb.GetFitSkillLevelFitSkillLevelCfg(gamedb.GetRealId(skillId, fitSkill.Lv))
			if lvCfg != nil {
				p.Add(lvCfg.Attribute)
			}
			starCfg := gamedb.GetFitSkillStarFitSkillStarCfg(gamedb.GetRealId(skillId, fitSkill.Star))
			if starCfg != nil {
				p.Add(starCfg.Attribute)
			}
		}
	}
	for fashionId, lv := range userFit.Fashion {
		fashionLevelCfg := gamedb.GetFitFashionLevelFitFashionLevelCfg(gamedb.GetRealId(fashionId, lv))
		if fashionLevelCfg != nil {
			p.Add(fashionLevelCfg.Attribute)
		}
	}
	return p
}

func calcMonthCardProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	timeNow := int(time.Now().Unix())
	for t, monthCard := range user.MonthCard.MonthCards {
		if monthCard.EndTime <= timeNow {
			continue
		}
		monthCardCfg := gamedb.GetMonthCardByType(t)
		if monthCardCfg != nil {
			p.Add(monthCardCfg.Attribute)
			if effectId := monthCardCfg.Privilege[pb.VIPPRIVILEGE_ATTR]; effectId != 0 {
				if effectCfg := gamedb.GetEffectEffectCfg(effectId); effectCfg != nil {
					p.Add(effectCfg.Attribute)
				}
			}
		}
	}
	return p
}

func calcAchievementProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()

	for _, id := range user.Achievement.Medal {
		cfg := gamedb.GetAchievementMedalMedalCfg(id)
		if cfg == nil {
			continue
		}
		p.Add(cfg.Buff)
	}
	return p
}

func calcTalentGeneralProps(user *objs.User, heroIndex int) *prop.Prop {
	hero := user.Heros[heroIndex]
	p := prop.NewProp()
	talentGeneral := hero.TalentGeneral[pb.PROPERTYMODULE_TALENT_GENERAL]
	for effectId, effectVal := range talentGeneral {
		switch effectId {
		case constConstant.TALENTEFFECT_5:
			p.AddOne(pb.PROPERTY_ATT_RATE_ALL, effectVal)
		case constConstant.TALENTEFFECT_6:
			p.AddOne(pb.PROPERTY_CRIT, effectVal)
		case constConstant.TALENTEFFECT_7:
			p.AddOne(pb.PROPERTY_ADD_CRIT_HIT, effectVal)
		case constConstant.TALENTEFFECT_9:
			p.AddOne(pb.PROPERTY_RED_HURT, effectVal)
		case constConstant.TALENTEFFECT_10:
			p.AddOne(pb.PROPERTY_UN_PALSY, effectVal)
		case constConstant.TALENTEFFECT_12:
			p.AddOne(pb.PROPERTY_MISS_RATE, effectVal)
		case constConstant.TALENTEFFECT_13:
			p.AddOne(pb.PROPERTY_RED_CRIT_HIT, effectVal)
		case constConstant.TALENTEFFECT_14:
			p.AddOne(pb.PROPERTY_DEF_RATE, effectVal)
		case constConstant.TALENTEFFECT_15:
			p.AddOne(pb.PROPERTY_HP_RATE, effectVal)
		case constConstant.TALENTEFFECT_17:
			p.AddOne(pb.PROPERTY_ADF_RATE, effectVal)
		case constConstant.TALENTEFFECT_19:
			p.AddOne(pb.PROPERTY_DEF_RATE, effectVal)
		}
	}
	return p
}

func calcElfProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	elfGrowCfg := gamedb.GetElfGrowElfGrowCfg(user.Elf.Lv)
	if elfGrowCfg != nil {
		p.Add(elfGrowCfg.Buff)
	}
	return p
}

func calcFitHolyEquipProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	userFitHolyEquip := user.FitHolyEquip
	for _, kv := range userFitHolyEquip.Equips {
		for _, id := range kv {
			fitHolyEquipCfg := gamedb.GetFitHolyEquipFitHolyEquipCfg(id)
			if fitHolyEquipCfg != nil {
				p.Add(fitHolyEquipCfg.Attribute)
			}
		}
	}
	effectMap := make(map[int]int)
	if userFitHolyEquip.SuitId != 0 {
		holyEquipSuitCfg := gamedb.GetFitHolyEquipSuitFitHolyEquipSuitCfg(userFitHolyEquip.SuitId)
		if holyEquipSuitCfg != nil {
			effectMap[holyEquipSuitCfg.Effect] = 0
		}
	}
	effectSlice := addEffect(effectMap, p)
	user.FitHolyEquipEffects = effectSlice
	return p
}

func calcFabaoLvProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	for _, v := range user.Fabaos {
		fabaoConf := gamedb.GetFabaoLvByIdAndLv(v.Id, v.Level)
		if fabaoConf == nil {
			continue
		}
		addProperties(user, heroIndex, p, fabaoConf.AttributeP, fabaoConf.AttributeM, fabaoConf.AttributeT)
		for _, vv := range v.Skill {
			skillConf := gamedb.GetFabaoSkillByIdAndSkillId(v.Id, vv)
			if skillConf == nil {
				continue
			}
			addProperties(user, heroIndex, p, skillConf.AttP, skillConf.AttM, skillConf.AttT)
		}
	}
	return p
}

func calcReinProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	rein := user.Rein
	reinT := gamedb.GetReinCfg(rein.Id)
	if reinT == nil {
		//throwPanic(errors.New("reinConf not found err id = " + strconv.Itoa(rein.Id)))
		return p
	}
	addProperties(user, heroIndex, p, reinT.AttributeP, reinT.AttributeM, reinT.AttributeT)
	return p
}

func calcAtlasProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	//图鉴
	atlases := user.Atlases
	addMap := make(map[int]int)
	for id, star := range atlases {
		conf := gamedb.GetAtlasStar(id, star)
		if conf == nil {
			continue
		}
		p.Add(conf.Attribute)
		for pid, pVal := range conf.Attribute {
			addMap[pid] += pVal
		}
	}
	//图鉴集合
	atlasGathers := user.AtlasGathers
	for id, star := range atlasGathers {
		atlasGatherConf := gamedb.GetAtlasUpgrade(id, star)
		if atlasGatherConf == nil {
			continue
		}
		p.Add(atlasGatherConf.Attribute)
		for pid, pVal := range atlasGatherConf.Attribute {
			addMap[pid] += pVal
		}
	}
	talentEffect16 := hero.TalentGeneral[pb.PROPERTYMODULE_ATLAS][constConstant.TALENTEFFECT_16]
	if talentEffect16 != 0 {
		for k, v := range addMap {
			p.AddOne(k, calcTalentGeneral(v, talentEffect16))
		}
	}
	return p
}

func calcAtlasWearProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	atlases := user.Atlases
	hero := user.Heros[heroIndex]
	atlasWear := hero.Wear.AtlasWear
	for id := range atlasWear {
		atlasCfg := gamedb.GetAtlasStar(id, atlases[id])
		if atlasCfg != nil {
			p.Add(atlasCfg.Attribute1)
		}
	}
	return p
}

func calcOfficialProps(user *objs.User, heroIndex int) *prop.Prop {

	p := prop.NewProp()
	conf := gamedb.GetOfficialOfficialCfg(user.Official)
	if conf != nil {
		p.Add(conf.Attribute)
	}
	return p
}

func calcPanaceaProps(user *objs.User, heroIndex int) *prop.Prop {
	hero := user.Heros[heroIndex]
	p := prop.NewProp()
	panaceas := user.Panaceas
	addMap := make(map[int]int)
	for id, panacea := range panaceas {
		conf := gamedb.GetPanaceaPanaceaCfg(id)
		if conf == nil {
			continue
		}
		for pid, pval := range conf.Attribute {
			p.AddOne(pid, pval*panacea.Number)
			addMap[pid] += pval * panacea.Number
		}
	}
	talentEffect18 := hero.TalentGeneral[pb.PROPERTYMODULE_PANACEA][constConstant.TALENTEFFECT_18]
	if talentEffect18 != 0 {
		for k, v := range addMap {
			p.AddOne(k, calcTalentGeneral(v, talentEffect18))
		}
	}
	return p
}

func calcPetProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	for id, pet := range user.Pet {
		gradeCfg := gamedb.GetPetsGradeConfCfg(gamedb.GetRealId(id, pet.Grade))
		if gradeCfg != nil {
			p.Add(gradeCfg.Attribute)
		}
		breakCfg := gamedb.GetPetsBreakConfCfg(gamedb.GetRealId(id, pet.Break))
		if gradeCfg != nil {
			p.Add(breakCfg.Attribute)
		}
	}
	return p
}

func calcEquipStrengthProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	addMap := make(map[int]int)
	for k, v := range hero.EquipsStrength {
		conf := gamedb.GetEquipStrengthConfByLvAndPos(k, v)
		if conf == nil {
			continue
		}
		if k != pb.EQUIPTYPE_WEAPON_R {
			var attr map[int]int
			switch hero.Job {
			case pb.JOB_ZHANSHI:
				attr = conf.AttributeP
			case pb.JOB_FASHI:
				attr = conf.AttributeM
			case pb.JOB_DAOSHI:
				attr = conf.AttributeT
			}
			for pid, pVal := range attr {
				addMap[pid] += pVal
			}
		}
		addProperties(user, heroIndex, p, conf.AttributeP, conf.AttributeM, conf.AttributeT)
	}
	effect4 := hero.TalentGeneral[pb.PROPERTYMODULE_EQUIP_STRENGTH][constConstant.TALENTEFFECT_4]
	if effect4 != 0 {
		for pid, pVal := range addMap {
			p.AddOne(pid, calcTalentGeneral(pVal, effect4))
		}
	}
	// 强化连携属性
	if hero.ConditionAttribute[pb.CONDITIONATTRIBUTETYPE_STRENGTHEN_LINK] != nil {
		if len(hero.ConditionAttribute[pb.CONDITIONATTRIBUTETYPE_STRENGTHEN_LINK]) > 0 {
			p.Add(hero.ConditionAttribute[pb.CONDITIONATTRIBUTETYPE_STRENGTHEN_LINK])
		}
	}
	return p
}

func calcEquipNormalProps(user *objs.User, heroIndex int) *prop.Prop {
	hero := user.Heros[heroIndex]
	talentGeneral := hero.TalentGeneral[pb.PROPERTYMODULE_EQUIP_NORMAL]

	p := prop.NewProp()
	// 记录套装属性
	weaponAddMap := make(map[int]int)
	otherAddMap := make(map[int]int)
	for pos, v := range hero.Equips {
		conf := gamedb.GetEquipEquipCfg(v.ItemId)
		if conf == nil {
			continue
		}
		for pid, pVal := range conf.Properties {
			if pos == pb.EQUIPTYPE_WEAPON_R {
				weaponAddMap[pid] += pVal
			} else {
				otherAddMap[pid] += pVal
			}
		}
		//幸运值
		p.AddOne(pb.PROPERTY_LUCKY, v.Lucky)
		//添加固定属性
		p.Add(conf.Properties)
		//添加特殊属性
		p.Add(conf.PropertiesStar)
		//添加随机属性
		for _, v := range v.RandProps {
			p.AddOne(v.PropId, v.Value)
		}
	}
	effect1 := talentGeneral[constConstant.TALENTEFFECT_1]
	if effect1 != 0 {
		addPropsArr := []int{pb.PROPERTY_PATT_MIN, pb.PROPERTY_PATT_MAX, pb.PROPERTY_MATT_MIN, pb.PROPERTY_MATT_MAX, pb.PROPERTY_TATT_MIN, pb.PROPERTY_TATT_MAX}
		for _, v := range addPropsArr {
			p.AddOne(v, calcTalentGeneral(weaponAddMap[v], effect1))
		}
	}
	effect2 := talentGeneral[constConstant.TALENTEFFECT_2]
	if effect2 != 0 {
		p.AddOne(pb.PROPERTY_HP, calcTalentGeneral(otherAddMap[pb.PROPERTY_HP], effect2))
	}
	effect3 := talentGeneral[constConstant.TALENTEFFECT_3]
	if effect3 != 0 {
		addPropsArr := []int{pb.PROPERTY_DEF_MIN, pb.PROPERTY_DEF_MAX, pb.PROPERTY_ADF_MIN, pb.PROPERTY_ADF_MAX}
		for _, v := range addPropsArr {
			p.AddOne(v, calcTalentGeneral(otherAddMap[v], effect3))
		}
	}
	return p
}

func calcWingProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	for _, v := range hero.Wings {
		conf := gamedb.GetWingNewWingNewCfg(v.Id)
		if conf != nil {
			p.Add(conf.Attribute)
		}
	}
	effectMap := make(map[int]int)
	for order, lv := range hero.WingSpecial {
		specialCfg := gamedb.GetWingSpecialByOrderAndLv(order, lv)
		if specialCfg != nil {
			p.Add(specialCfg.Attribute)
			effectMap[specialCfg.Effect] = 0
		}
	}
	effectSlice := addEffect(effectMap, p)
	hero.WingEffects = effectSlice
	return p
}

func calcDragonEquipProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	for id, lv := range user.Heros[heroIndex].DragonEquip {
		conf := gamedb.GetDragonEquipLevelDragonEquipLevelCfg(gamedb.GetRealId(id, lv))
		if conf != nil {
			p.Add(conf.Attribute)
		}
	}
	return p
}

func calcKingarmsProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	for _, v := range user.Heros[heroIndex].Kingarms {
		if v.Id <= 0 {
			continue
		}
		conf := gamedb.GetKingarmsKingarmsCfg(v.Id)
		if conf != nil {
			p.Add(conf.Properties)
			p.Add(conf.Sproperties)
		}
	}
	return p
}

func calcZodiacProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()

	qualityMap := make(map[int]int) //品质=》数量
	zodiacMap := make(map[int]int)  //品质=》itemId

	for _, v := range user.Heros[heroIndex].Zodiacs {
		if v.Id <= 0 {
			continue
		}
		conf := gamedb.GetZodiacEquipZodiacEquipCfg(v.Id)
		if conf == nil {
			continue
		}
		p.Add(conf.Properties)

		itemCfg := gamedb.GetItemBaseCfg(v.Id)
		itemQuality := itemCfg.Quality
		// 记录每个品质对应itemId
		if qualityMap[itemQuality] == 0 {
			zodiacMap[itemQuality] = v.Id
		}
		qualityMap[itemQuality] += 1
	}
	// 向下兼容 预防品质多个
	quaKeySlice := make([]int, 0)
	for qua := range qualityMap {
		quaKeySlice = append(quaKeySlice, qua)
		for qua1 := range qualityMap {
			if qua > qua1 {
				qualityMap[qua1] += qualityMap[qua]
			}
		}
	}
	sort.Ints(quaKeySlice)

	effectMap := make(map[int]int)
	// 从高到低
	var isSex, isTwelve bool
	for i := len(quaKeySlice) - 1; i >= 0; i-- {
		quality := quaKeySlice[i]
		zodiacCfg := gamedb.GetZodiacEquipZodiacEquipCfg(zodiacMap[quality])
		if zodiacCfg == nil {
			continue
		}
		if qualityMap[quality] >= constEquip.ZODIAC_SUIT_SIX && !isSex {
			effectMap[zodiacCfg.Effectid[0]] = 0
			isSex = true
		}
		if qualityMap[quality] >= constEquip.ZODIAC_SUIT_TWELVE && !isTwelve {
			effectMap[zodiacCfg.Effectid[1]] = 0
			isTwelve = true
		}
	}
	effectSlice := addEffect(effectMap, p)
	user.Heros[heroIndex].ZodiacSuit = effectSlice
	return p
}

func calcClearProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	addMap := make(map[int]int)
	for _, clearUnit := range user.Heros[heroIndex].EquipClear {
		for _, equipClearUnit := range clearUnit {
			addMap[equipClearUnit.PropId] += equipClearUnit.Value
		}
	}
	p.Add(addMap)
	return p
}

func calcDictateProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()

	dictateMap := make(map[int]map[int]int)
	dictateMap[constEquip.DICTATE_TYPE_LEFT] = make(map[int]int)
	dictateMap[constEquip.DICTATE_TYPE_RIGHT] = make(map[int]int)

	for body, lv := range user.Heros[heroIndex].Dictates {
		if lv <= 0 {
			continue
		}
		dictateCfg := gamedb.GetDictateByBodyAndGrade(body, lv)
		if dictateCfg != nil {
			p.Add(dictateCfg.Properties)
		}
		// 大于5右边 小于5左边
		for i := 1; i <= lv; i++ {
			if body > 5 {
				dictateMap[constEquip.DICTATE_TYPE_RIGHT][i]++
			} else {
				dictateMap[constEquip.DICTATE_TYPE_LEFT][i]++
			}
		}
	}

	effectMap := make(map[int]int)
	for lr, dictateLvNum := range dictateMap {
		if len(dictateLvNum) <= 0 {
			continue
		}
		var isTwo, isThree, isFive bool
		lvSlice := make([]int, 0)
		for lv := range dictateLvNum {
			lvSlice = append(lvSlice, lv)
		}
		sort.Ints(lvSlice)
		for l := len(lvSlice) - 1; l >= 0; l-- {
			lv := lvSlice[l]
			dictateSuitCfg := gamedb.GetDictateSuitDictateSuitCfg(lv)
			effectId := dictateSuitCfg.Effectid1
			if lr == constEquip.DICTATE_TYPE_RIGHT {
				effectId = dictateSuitCfg.Effectid2
			}
			for i := dictateLvNum[lv]; i > 0; i-- {
				effect, ok := effectId[i]
				if ok {
					if i == constEquip.DICTATE_SUIT_TWO && !isTwo {
						effectMap[effect] = 0
						isTwo = true
					} else if i == constEquip.DICTATE_SUIT_THREE && !isThree {
						effectMap[effect] = 0
						isThree = true
					} else if i == constEquip.DICTATE_SUIT_FIVE && !isFive {
						effectMap[effect] = 0
						isFive = true
					}
				}
			}
		}
	}
	effectSlice := addEffect(effectMap, p)
	user.Heros[heroIndex].DictateSuit = effectSlice
	return p
}

func calcJewelProps(user *objs.User, heroIndex int) *prop.Prop {
	hero := user.Heros[heroIndex]
	p := prop.NewProp()
	countLv := 0
	addMap := make(map[int]int)
	for _, jewel := range hero.Jewel {
		if jewel.One != 0 {
			oneCfg := gamedb.GetJewelJewelCfg(jewel.One)
			if oneCfg != nil {
				countLv += oneCfg.Level
				p.Add(oneCfg.Attribute)
				for pid, pVal := range oneCfg.Attribute {
					addMap[pid] += pVal
				}
			}
		}
		if jewel.Two != 0 {
			twoCfg := gamedb.GetJewelJewelCfg(jewel.Two)
			if twoCfg != nil {
				countLv += twoCfg.Level
				p.Add(twoCfg.Attribute)
				for pid, pVal := range twoCfg.Attribute {
					addMap[pid] += pVal
				}
			}
		}
		if jewel.Three != 0 {
			threeCfg := gamedb.GetJewelJewelCfg(jewel.Three)
			if threeCfg != nil {
				countLv += threeCfg.Level
				p.Add(threeCfg.Attribute)
				for pid, pVal := range threeCfg.Attribute {
					addMap[pid] += pVal
				}
			}
		}
	}
	talentEffect8 := hero.TalentGeneral[pb.PROPERTYMODULE_JEWEL][constConstant.TALENTEFFECT_8]
	if talentEffect8 != 0 {
		for k, v := range addMap {
			p.AddOne(k, calcTalentGeneral(v, talentEffect8))
		}
	}
	// 所有部位等级属性
	jewelSuitCfgs := gamedb.GetJewelSuitCfgs()
	maxSumId := 0
	for _, v := range jewelSuitCfgs {
		if countLv >= v.Sum && v.Sum > maxSumId {
			maxSumId = v.Sum
		}
	}

	user.JewelAllLv = countLv
	jewelSuitCfg := gamedb.GetJewelSuitJewelSuitCfg(maxSumId)
	if jewelSuitCfg != nil {
		p.Add(jewelSuitCfg.Attribute)
	}
	return p
}

func calcFashionProps(user *objs.User, heroIndex int, t int) *prop.Prop {
	p := prop.NewProp()
	fashions := user.Heros[heroIndex].Fashions
	for _, v := range fashions {
		fashionConf := gamedb.GetFashionConf(v.Id, v.Lv)
		if fashionConf != nil && fashionConf.FashionType == t {
			p.Add(fashionConf.Attribute)
			p.Add(fashionConf.AttributeS)
		}
	}
	return p
}

func calcFashionWeaponProps(user *objs.User, heroIndex int) *prop.Prop {

	return calcFashionProps(user, heroIndex, pb.FASHIONTYPE_WEAPON)
}

func calcFashionClothProps(user *objs.User, heroIndex int) *prop.Prop {

	return calcFashionProps(user, heroIndex, pb.FASHIONTYPE_CLOTH)
}

func calcInsideProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	heroInside := hero.Inside
	addMap := make(map[int]int)
	for pos, id := range heroInside.Acupoint {
		insideCfg := gamedb.GetInsideArtInsideArtCfg(id)
		if insideCfg != nil {
			switch pos {
			case pb.INSIDETYPE_ONE:
				p.Add(insideCfg.Attribute1)
				for pid, pVal := range insideCfg.Attribute1 {
					addMap[pid] += pVal
				}
			case pb.INSIDETYPE_TWO:
				p.Add(insideCfg.Attribute2)
				for pid, pVal := range insideCfg.Attribute2 {
					addMap[pid] += pVal
				}
			case pb.INSIDETYPE_THREE:
				p.Add(insideCfg.Attribute3)
				for pid, pVal := range insideCfg.Attribute3 {
					addMap[pid] += pVal
				}
			case pb.INSIDETYPE_FOUR:
				p.Add(insideCfg.Attribute4)
				for pid, pVal := range insideCfg.Attribute4 {
					addMap[pid] += pVal
				}
			}
		}
	}
	effectMap := make(map[int]int)
	insideUp := 0
	for id, skill := range heroInside.Skill {
		insideSkillCfg := gamedb.GetInsideSkillBySidAndLv(id, skill.Level)
		effectMap[insideSkillCfg.Effect] = 0
		insideUp += insideSkillCfg.InsideUp
	}
	if len(effectMap) > 0 {
		hero.InsideEffects = make([]int, 0)
		for effect := range effectMap {
			effectCfg := gamedb.GetEffectEffectCfg(effect)
			if effectCfg == nil {
				continue
			}
			p.Add(effectCfg.Attribute)
			for pid, pVal := range effectCfg.Attribute {
				addMap[pid] += pVal
			}
			hero.InsideEffects = append(hero.InsideEffects, effect)
		}
	}
	if insideUp != 0 && len(addMap) > 0 {
		for pid, pVal := range addMap {
			p.AddOne(pid, calcTalentGeneral(pVal, insideUp))
		}
	}
	return p
}

func calcRingProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	hero.RingPhantomEffects = make([]int, 0)
	for _, ringUnit := range hero.Ring {
		if ringUnit.Rid == 0 {
			continue
		}
		ringCfg := gamedb.GetRingRingCfg(ringUnit.Rid)
		if ringCfg == nil {
			continue
		}
		p.Add(ringCfg.Attr)
		ringStrengthenCfg := gamedb.GetRingStrengthenRingStrengthenCfg(ringUnit.Strengthen)
		if ringStrengthenCfg == nil {
			continue
		}
		p.Add(ringStrengthenCfg.Attribute)
		ringPhantomCfg := gamedb.GetRingPhantomRingPhantomCfg(ringUnit.Pid)
		if ringPhantomCfg == nil {
			continue
		}
		p.Add(ringPhantomCfg.Attribute)
		// 幻灵技能效果
		effectMap := make(map[int]int)
		for _, ringPhantom := range ringUnit.Phantom {
			phantomCfg := gamedb.GetPhantomPhantomCfg(ringPhantom.Phantom)
			if phantomCfg != nil {
				effectMap[phantomCfg.Effect] = 0
			}
			for skillId, lv := range ringPhantom.Skill {
				if skillId == pb.RINGSKILLID_ONE {
					phantomCfg := gamedb.GetPhantomSkill1(ringPhantom.Phantom, pb.RINGSKILLID_ONE, lv)
					if phantomCfg == nil {
						continue
					}
					if phantomCfg.Effect1 != 0 {
						effectMap[phantomCfg.Effect1] = 0
					}
				} else {
					phantomCfg := gamedb.GetPhantomSkill2(ringPhantom.Phantom, pb.RINGSKILLID_TWO, lv)
					if phantomCfg == nil {
						continue
					}
					if phantomCfg.Effect2 != 0 {
						effectMap[phantomCfg.Effect2] = 0
					}
				}
			}
		}
		if len(effectMap) > 0 {
			for effect := range effectMap {
				effectCfg := gamedb.GetEffectEffectCfg(effect)
				if effectCfg == nil {
					continue
				}
				p.Add(effectCfg.Attribute)
				hero.RingPhantomEffects = append(hero.RingPhantomEffects, effect)
			}
		}
	}
	return p
}

func calcGodEquipProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	minLv := math.MaxInt32
	for _, godEquip := range hero.GodEquips {
		godEquipConf := gamedb.GetGodEquipLevelConfCfg(gamedb.GetRealId(godEquip.Id, godEquip.Lv))
		if godEquipConf == nil {
			common.ThrowPanic(base.Conf.Sandbox, gamedb.ERRSETTINGNOTFOUND)
		}
		p.Add(godEquipConf.Attribute)
		if godEquip.Lv < minLv {
			minLv = godEquip.Lv
		}
	}
	//套装属性
	if len(hero.GodEquips) == constEquip.GOD_EQUIP_NUM {
		suitConf := gamedb.GetGodEquipSuitConfCfg(minLv)
		if suitConf != nil {
			p.Add(suitConf.Attribute)
		}
	}
	return p
}

func calcAreaProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	for areaT, lv := range hero.Area {
		lvCfg := gamedb.GetAreaLevelAreaLevelCfg(gamedb.GetRealId(areaT, lv))
		if lvCfg == nil {
			continue
		}
		p.Add(lvCfg.Attribute)
		p.Add(lvCfg.Attribute1)
	}
	return p
}

func addProperties(user *objs.User, heroIndex int, p *prop.Prop, proP, proM, proT gamedb.IntMap) {
	switch user.Heros[heroIndex].Job {
	case pb.JOB_ZHANSHI:
		p.Add(proP)
	case pb.JOB_FASHI:
		p.Add(proM)
	case pb.JOB_DAOSHI:
		p.Add(proT)
	}
}

func calcMagicCircleProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	for _, id := range hero.MagicCircle {
		circleLevelCfg := gamedb.GetMagicCircleLevelMagicCircleLevelCfg(id)
		if circleLevelCfg == nil {
			continue
		}
		p.Add(circleLevelCfg.Attribute)
	}
	return p
}

// 经验池等级 add prop
func calcExpPoolProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	lvCfg := gamedb.GetExpPoolCfg(heroIndex, hero.ExpLvl)
	if lvCfg == nil {
		logger.Error("GetExpPoolCfg err heroIndex:%v, hero.ExpLvl:%v", heroIndex, hero.ExpLvl)
		return p
	}
	propInfo := make(map[int]int)
	if hero.Job == pb.JOB_ZHANSHI {
		propInfo = lvCfg.Attribute
	} else if hero.Job == pb.JOB_FASHI {
		propInfo = lvCfg.Attribute2
	} else if hero.Job == pb.JOB_DAOSHI {
		propInfo = lvCfg.Attribute3
	}
	p.Add(propInfo)
	return p
}

func calcTalentProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	hero.TalentEffects = make([]int, 0)
	for _, talentUnit := range hero.Talent.TalentList {
		for id, lv := range talentUnit.Talents {
			levelCfg := gamedb.GetTalentLevelTalentLevelCfg(gamedb.GetRealId(id, lv))
			if levelCfg == nil {
				continue
			}
			effectCfg := gamedb.GetEffectEffectCfg(levelCfg.Effect)
			if effectCfg != nil {
				p.Add(effectCfg.Attribute)
				if levelCfg.Skill == 0 {
					hero.TalentEffects = append(hero.TalentEffects, effectCfg.Id)
				}
			}
		}
	}
	return p
}

func calcHolyBeastProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	user.Heros[heroIndex].HolyBeastEffects = make([]int, 0)
	for _, info := range hero.HolyBeastInfos {

		for i := 1; i <= info.Star; i++ {
			cfg := gamedb.GetHolyBeastByTypesAndStar(info.Types, i)
			if cfg != nil {
				p.Add(cfg.Properties)
			}
		}
		if len(info.ChooseProp) > 0 {
			for star, effectIndex := range info.ChooseProp {
				cfg := gamedb.GetHolyBeastByTypesAndStar(info.Types, star)
				if cfg == nil {
					continue
				}

				if cfg.SelectProperties != nil {
					if len(cfg.SelectProperties) > 0 {
						selectProp := cfg.SelectProperties[effectIndex]
						sProp := make(map[int]int)
						sProp[selectProp[0]] = selectProp[1]
						p.Add(sProp)
					}

				}

				if cfg.Effect != nil {
					if len(cfg.Effect) > 0 {
						if effectIndex > len(cfg.Effect)-1 {
							continue
						}
						effectCfg := gamedb.GetEffectEffectCfg(cfg.Effect[effectIndex])
						if effectCfg == nil {
							continue
						}
						if effectCfg.Attribute != nil {
							p.Add(effectCfg.Attribute)
						}
						user.Heros[heroIndex].HolyBeastEffects = append(user.Heros[heroIndex].HolyBeastEffects, cfg.Effect[effectIndex])
					}
				}
			}

		}
	}

	return p
}

func calcVipProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	if vipCfg := gamedb.GetVipLvlCfg(user.VipLevel); vipCfg != nil {
		if effectId := vipCfg.Privilege[pb.VIPPRIVILEGE_ATTR]; effectId != 0 {
			if effectCfg := gamedb.GetEffectEffectCfg(effectId); effectCfg != nil {
				p.Add(effectCfg.Attribute)
			}
		}
	}
	return p
}

func calcSkillProps(user *objs.User, heroIndex int) *prop.Prop {

	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	for _, v := range hero.Skills {
		skillLvConf := gamedb.GetSkillLevelSkillCfg(gamedb.GetSkillLvId(v.Id, v.Lv))
		if skillLvConf == nil {
			logger.Error("获取技能配置异常,技能Id:%v，等级：%v", v.Id, v.Lv)
			continue
		}
		if len(skillLvConf.Attribute) > 0 {
			p.Add(skillLvConf.Attribute)
		}
	}
	return p
}

func calcChuanShiEquipProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	heroChuanShi := hero.ChuanShi
	suitMaps := make(map[int]int)
	for pos, id := range heroChuanShi {
		if id < 1 {
			continue
		}
		chuanShiEquipCfg := gamedb.GetChuanShiEquipChuanShiEquipCfg(id)
		if chuanShiEquipCfg != nil {
			p.Add(chuanShiEquipCfg.Properties)
			suitMaps[pos] = chuanShiEquipCfg.Level
		}
	}
	effectMap := make(map[int]int)
	suitCfgs := gamedb.GetChuanShiSuitCfgs()
	for t, cfg := range suitCfgs {
		lvNumMap := make(map[int]int)
		maxLv := math.MinInt32
		for _, pos := range cfg.Type {
			posLv := suitMaps[pos]
			if posLv == 0 {
				continue
			}
			if maxLv < posLv {
				maxLv = posLv
			}
			lvNumMap[posLv] += 1
		}
		for lv := range lvNumMap {
			for lv1 := range lvNumMap {
				if lv < lv1 {
					lvNumMap[lv] += lvNumMap[maxLv]
				}
			}
		}
		lvSlice := make([]int, 0)
		for lv := range lvNumMap {
			lvSlice = append(lvSlice, lv)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(lvSlice)))
		isTwo, isFour := false, false
		for _, lv := range lvSlice {
			lvNum := lvNumMap[lv]
			chuanShiSuitTypeCfg := gamedb.GetChuanShiSuitByTypeAndLv(t, lv)
			if lvNum > 1 {
				if !isTwo {
					effectMap[chuanShiSuitTypeCfg.Attribute1] = 0
					isTwo = true
				}
				if lvNum >= 4 && !isFour {
					effectMap[chuanShiSuitTypeCfg.Attribute2] = 0
					isFour = true
				}
			}
		}
	}
	effectSlice := addEffect(effectMap, p)
	hero.ChuanShiEquipEffects = effectSlice
	return p
}

func calcChuanShiStrengthen(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	equips := hero.ChuanShi
	strengthens := hero.ChuanshiStrengthen
	for pos, id := range equips {
		if id > 0 {
			strengthenCfg := gamedb.GetChuanShiStrengthenByPosAndLv(pos, strengthens[pos])
			if strengthenCfg != nil {
				addProperties(user, heroIndex, p, strengthenCfg.AttributeP, strengthenCfg.AttributeM, strengthenCfg.AttributeT)
			}
		}
	}
	effectMap := make(map[int]int)
	t1PosMap := map[int]int{pb.CHUANSHIPOS_WU_QI: 0, pb.CHUANSHIPOS_YI_FU: 0}
	t1MinLv, t2MinLv := math.MaxInt32, math.MaxInt32
	for _, pos := range pb.CHUANSHIPOS_ARRAY {
		if _, ok := t1PosMap[pos]; ok {
			if strengthens[pos] < t1MinLv {
				t1MinLv = strengthens[pos]
			}
		} else {
			if strengthens[pos] < t2MinLv {
				t2MinLv = strengthens[pos]
			}
		}
	}
	t1Map, t2Map := make(map[int]int), make(map[int]int)
	t1Slice, t2Slice := make([]int, 0), make([]int, 0)
	linkCfgs := gamedb.GetChuanShiStrengthenLinkCfgs()
	for _, cfg := range linkCfgs {
		if cfg.Type == 1 {
			t1Map[cfg.Condition] = cfg.Id
			t1Slice = append(t1Slice, cfg.Condition)
		} else {
			t2Map[cfg.Condition] = cfg.Id
			t2Slice = append(t2Slice, cfg.Condition)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(t1Slice)))
	sort.Sort(sort.Reverse(sort.IntSlice(t2Slice)))
	t1Id, t2Id := 0, 0
	for _, lv := range t1Slice {
		if t1MinLv >= lv {
			t1Id = t1Map[lv]
			break
		}
	}
	for _, lv := range t2Slice {
		if t2MinLv >= lv {
			t2Id = t2Map[lv]
			break
		}
	}
	linkCfg1 := gamedb.GetChuanShiStrengthenLinkChuanShiStrengthenLinkCfg(t1Id)
	if linkCfg1 != nil {
		effectMap[linkCfg1.Effect] = 0
	}
	linkCfg2 := gamedb.GetChuanShiStrengthenLinkChuanShiStrengthenLinkCfg(t2Id)
	if linkCfg2 != nil {
		effectMap[linkCfg2.Effect] = 0
	}
	effectSlice := addEffect(effectMap, p)
	hero.ChuanShiStrengthenEffects = effectSlice
	return p
}

func calcAncientSkillProps(user *objs.User, heroIndex int) *prop.Prop {
	hero := user.Heros[heroIndex]
	p := prop.NewProp()
	effectMap := make(map[int]int)
	heroAncientSkill := hero.AncientSkill
	lvCfg := gamedb.GetAncientSkillLevelAncientSkillLevelCfg(constConstant.COMPUTE_TEN_THOUSAND + heroAncientSkill.Level)
	if lvCfg != nil {
		effect := 0
		switch hero.Job {
		case pb.JOB_ZHANSHI:
			effect = lvCfg.Effect1
		case pb.JOB_FASHI:
			effect = lvCfg.Effect2
		case pb.JOB_DAOSHI:
			effect = lvCfg.Effect3
		}
		effectMap[effect] = 0
	}
	gradeCfg := gamedb.GetAncientSkillGradeAncientSkillGradeCfg(constConstant.COMPUTE_TEN_THOUSAND + heroAncientSkill.Grade)
	if gradeCfg != nil {
		effect := 0
		switch hero.Job {
		case pb.JOB_ZHANSHI:
			effect = gradeCfg.ZhanEffect
		case pb.JOB_FASHI:
			effect = gradeCfg.FaEffect
		case pb.JOB_DAOSHI:
			effect = gradeCfg.DaoEffect
		}
		effectMap[effect] = 0
	}
	effectSlice := addEffect(effectMap, p)
	hero.AncientSkillEffects = effectSlice
	return p
}

func calcTitleProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	userTitle := user.Title
	timeNow := int(time.Now().Unix())
	for id, info := range userTitle {
		titleCfg := gamedb.GetTitleTitleCfg(id)
		if titleCfg == nil || (info.EndTime != -1 && info.EndTime < timeNow) {
			continue
		}
		p.Add(titleCfg.Attribute)
	}
	return p
}

func calcMiJiProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	effectSlice := make([]int, 0)
	for id, info := range user.MiJi {
		cfgId := gamedb.GetRealId(id, info.MiJiLv)
		cfg := gamedb.GetMijiLevelMijiLevelCfg(cfgId)
		if cfg == nil {
			continue
		}
		p.Add(cfg.Attribute)

		effectCfg := gamedb.GetEffectEffectCfg(cfg.SkillLevel)
		if effectCfg != nil {
			p.Add(effectCfg.Attribute)
			effectSlice = append(effectSlice, cfg.SkillLevel)
		}
	}
	user.Heros[heroIndex].MiJiEffects = effectSlice
	return p
}

func getAncientTreasureAddPercent1(datas gamedb.IntSlice2, treasureTypesZhuLin, treasureTypesStar, treasureIdsZhuLin, treasureIdsStar map[int]int) (map[int]int, map[int]int, map[int]int, map[int]int) {

	data := datas

	//1:宝物类型  生效类型
	if data[0][0] == 1 {
		//宝物type 类型
		if data[2][0] == 1 {
			//注灵属性 万分比
			treasureTypesZhuLin[data[1][0]] += data[3][0]
		}
		if data[2][0] == 2 {
			//升星属性 万分比
			treasureTypesStar[data[1][0]] += data[3][0]
		}
	}

	//2:宝物id  生效类型
	if data[0][0] == 2 {
		//宝物id 类型
		if data[2][0] == 1 {
			//注灵属性 万分比
			treasureIdsZhuLin[data[1][0]] += data[3][0]
		}
		if data[2][0] == 2 {
			//升星属性 万分比
			treasureIdsStar[data[1][0]] += data[3][0]
		}

	}
	return treasureTypesZhuLin, treasureTypesStar, treasureIdsZhuLin, treasureIdsStar
}

// 获取远古宝物属性加成百分比
func getAncientTreasureAddPercent(user *objs.User) (treasureTypesZhuLin, treasureTypesStar, treasureIdsZhuLin, treasureIdsStar map[int]int) {

	//宝物类型  注灵 加成万分比
	treasureTypesZhuLin = make(map[int]int)

	//宝物类型  升星 加成万分比
	treasureTypesStar = make(map[int]int)

	//宝物id  注灵 加成万分比
	treasureIdsZhuLin = make(map[int]int)

	//宝物id  升星 加成万分比
	treasureIdsStar = make(map[int]int)

	for treasureId, info := range user.AncientTreasure {
		//升星属性
		starCfg := gamedb.GetAncientTreasureStarById(treasureId, info.Star)
		if starCfg != nil {
			if starCfg.Type == 3 {
				if starCfg.Attribute != nil && len(starCfg.Attribute) >= 4 && len(starCfg.Attribute[0]) >= 1 && len(starCfg.Attribute[1]) >= 1 && len(starCfg.Attribute[2]) >= 1 && len(starCfg.Attribute[3]) >= 1 {
					treasureTypesZhuLin, treasureTypesStar, treasureIdsZhuLin, treasureIdsStar = getAncientTreasureAddPercent1(starCfg.Attribute, treasureTypesZhuLin, treasureTypesStar, treasureIdsZhuLin, treasureIdsStar)
				}
			}
		}

		if info.JueXinLv > 0 {
			//觉醒生效
			jueXinCfg := gamedb.GetAncientTreasureJueXinCfg(treasureId)
			if jueXinCfg != nil {
				if jueXinCfg.Type == 3 {
					if jueXinCfg.Attribute != nil && len(jueXinCfg.Attribute) >= 4 && len(jueXinCfg.Attribute[0]) >= 1 && len(jueXinCfg.Attribute[1]) >= 1 && len(jueXinCfg.Attribute[2]) >= 1 && len(jueXinCfg.Attribute[3]) >= 1 {
						treasureTypesZhuLin, treasureTypesStar, treasureIdsZhuLin, treasureIdsStar = getAncientTreasureAddPercent1(jueXinCfg.Attribute, treasureTypesZhuLin, treasureTypesStar, treasureIdsZhuLin, treasureIdsStar)
					}
				}
			}
		}

	}

	//套装加成百分比
	allActiveTaoz := getAncientTreasureTaoZ(user)
	for taoId, num := range allActiveTaoz {
		taoCfg := gamedb.GetTreasureSuitTreasureSuitCfg(taoId)
		if taoCfg == nil {
			continue
		}

		for i := 1; i <= num; i++ {
			if i == 1 {
				if taoCfg.Type1 == 3 {
					treasureTypesZhuLin, treasureTypesStar, treasureIdsZhuLin, treasureIdsStar = getAncientTreasureAddPercent1(taoCfg.Attribute1, treasureTypesZhuLin, treasureTypesStar, treasureIdsZhuLin, treasureIdsStar)
				}
			}
			if i == 2 {
				if taoCfg.Type2 == 3 {
					treasureTypesZhuLin, treasureTypesStar, treasureIdsZhuLin, treasureIdsStar = getAncientTreasureAddPercent1(taoCfg.Attribute2, treasureTypesZhuLin, treasureTypesStar, treasureIdsZhuLin, treasureIdsStar)
				}
			}
			if i == 3 {
				if taoCfg.Type3 == 3 {
					treasureTypesZhuLin, treasureTypesStar, treasureIdsZhuLin, treasureIdsStar = getAncientTreasureAddPercent1(taoCfg.Attribute3, treasureTypesZhuLin, treasureTypesStar, treasureIdsZhuLin, treasureIdsStar)
				}
			}
		}

	}

	return
}

func getAddAllProp(user *objs.User, heroIndex, treasureId, conditionType, types int, props map[int]int, datas gamedb.IntSlice2) map[int]int {

	if types == 1 {
		//类型1 -- 直接加属性
		for _, data := range datas {
			for i, j := 0, len(data); i < j; i += 2 {
				props[data[i]] += data[i+1]
			}
		}
	}

	if treasureId > 0 {
		if types == 2 {
			//类型2 -- 特殊条件，增加属性   condition类型,条件数量 |加成属性id,属性数量
			if len(datas) > 0 {
				if len(datas[0]) >= 2 {

					if user.Heros[heroIndex] != nil {

						if user.Heros[heroIndex].AncientTreasureConditionAttribute[conditionType] != nil {
							if user.Heros[heroIndex].AncientTreasureConditionAttribute[conditionType][treasureId] != nil {
								for k, v := range user.Heros[heroIndex].AncientTreasureConditionAttribute[conditionType][treasureId] {
									props[k] += v
								}
							}
						}
					}

				}
			}
		}

	}
	return props
}

func calcAncientTreasure(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	zhuLinProp := make(map[int]int)
	starProp := make(map[int]int)
	jueXinProp := make(map[int]int)
	taoZhuangProp := make(map[int]int)

	//宝物类型  注灵 加成万分比,  宝物类型  升星 加成万分比,  宝物id  注灵 加成万分比  ,宝物id  升星 加成万分比
	treasureTypesZhuLin, treasureTypesStar, treasureIdsZhuLin, treasureIdsStar := getAncientTreasureAddPercent(user)

	for treasureId, info := range user.AncientTreasure {
		//注灵属性
		baseAncientTreasure := gamedb.GetTreasureTreasureCfg(treasureId)
		if baseAncientTreasure == nil {
			continue
		}
		addPercent := 0
		addPercent += treasureTypesZhuLin[baseAncientTreasure.Type]
		addPercent += treasureIdsZhuLin[treasureId]
		zhuLinCfg := gamedb.GetAncientTreasureZhuLinLvById(treasureId, info.ZhuLinLv)
		if zhuLinCfg != nil {
			for k, v := range zhuLinCfg.Attribute {
				zhuLinProp[k] += int(math.Ceil(float64(v) * (1 + float64(addPercent))))
			}
		}

		//升星属性
		addPercent = 0
		addPercent += treasureTypesStar[baseAncientTreasure.Type]
		addPercent += treasureIdsStar[treasureId]
		starCfg := gamedb.GetAncientTreasureStarById(treasureId, info.Star)
		if starCfg != nil {
			starProp = getAddAllProp(user, heroIndex, treasureId, pb.CONDITIONATTRIBUTETYPE_ANCIENT_TREASURE_STAR, starCfg.Type, starProp, starCfg.Attribute)
			for k, v := range starProp {
				starProp[k] = int(math.Ceil(float64(v) * (1 + float64(addPercent)/constConstant.COMPUTE_TEN_THOUSAND)))
			}
		}

		//觉醒属性加成
		if info.JueXinLv > 0 {
			jueXinCfg := gamedb.GetAncientTreasureJueXinCfg(treasureId)
			if jueXinCfg != nil {
				jueXinProp = getAddAllProp(user, heroIndex, treasureId, pb.CONDITIONATTRIBUTETYPE_ANCIENT_TREASURE_JUE_XING, jueXinCfg.Type, jueXinProp, jueXinCfg.Attribute)
			}
		}
	}

	//套装属性加成
	allActiveTaoZhuang := getAncientTreasureTaoZ(user)
	for taoId, num := range allActiveTaoZhuang {
		tapCfg := gamedb.GetTreasureSuitTreasureSuitCfg(taoId)
		if tapCfg == nil {
			continue
		}
		for i := 1; i <= num; i++ {
			if i == 1 {
				taoZhuangProp = getAddAllProp(user, heroIndex, -1, -1, tapCfg.Type1, taoZhuangProp, tapCfg.Attribute1)
			}

			if i == 2 {
				taoZhuangProp = getAddAllProp(user, heroIndex, -1, -1, tapCfg.Type2, taoZhuangProp, tapCfg.Attribute2)
			}

			if i == 3 {
				taoZhuangProp = getAddAllProp(user, heroIndex, -1, -1, tapCfg.Type3, taoZhuangProp, tapCfg.Attribute3)
			}
		}
	}

	//套装 条件属性
	heroInfo := user.Heros[heroIndex]
	if heroInfo != nil {
		if heroInfo.ConditionAttribute[pb.CONDITIONATTRIBUTETYPE_ANCIENT_TREASURE_TAO_ZHUANG] != nil {
			if len(heroInfo.ConditionAttribute[pb.CONDITIONATTRIBUTETYPE_ANCIENT_TREASURE_TAO_ZHUANG]) > 0 {
				p.Add(heroInfo.ConditionAttribute[pb.CONDITIONATTRIBUTETYPE_ANCIENT_TREASURE_TAO_ZHUANG])
			}
		}
	}

	p.Add(zhuLinProp)
	p.Add(starProp)
	p.Add(jueXinProp)
	p.Add(taoZhuangProp)
	return p
}

// 获取远古宝物  达成的  K:套装id  v:达成的特殊加成
func getAncientTreasureTaoZ(user *objs.User) map[int]int {
	jueXinMap := make(map[int]int) //k:套装id v:可以加成的效果

	for id, data := range gamedb.GetAllAncientTreasureSuit() {
		allThreeStar := make([]int, 0)
		allJueXin := make([]int, 0)
		needContinue := false
		for _, tid := range data.TruesureId {
			if user.AncientTreasure[tid] == nil {
				needContinue = true
				break
			}
			if user.AncientTreasure[tid].Star >= 3 {
				allThreeStar = append(allThreeStar, tid)
			}

			if user.AncientTreasure[tid].JueXinLv >= 1 {
				allJueXin = append(allJueXin, tid)
			}

		}
		if needContinue {
			continue
		}
		//全都激活 加成
		jueXinMap[id] = 1

		//全部达成3星
		if len(allThreeStar) >= len(data.TruesureId) {
			jueXinMap[id] = 2
		}

		//全都激活觉醒效果
		if len(allJueXin) >= len(data.TruesureId) {
			jueXinMap[id] = 3
		}
	}
	return jueXinMap
}

func calcPetAppendage(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	for pid, lv := range user.PetAppendage {
		appendageCfg := gamedb.GetPetAppendageByPidAndLv(pid, lv)
		if appendageCfg != nil {
			p.Add(appendageCfg.AttributePets)
		}
	}
	if len(user.PetAppendageEffects) > 0 {
		effectSlice := addEffect(user.PetAppendageEffects, p)
		hero.PetAppendageEffects = effectSlice
	}
	return p
}

func calcDaBaoEquipProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	effectMap := make(map[int]int)
	hero := user.Heros[heroIndex]
	hero.MapEffects = make(map[int]map[int]int)
	for t, lv := range user.DaBaoEquip {
		if equipCfg := gamedb.GetDaBaoEquipByTypeAndLv(t, lv); equipCfg != nil {
			p.Add(equipCfg.Attribute)
		}
		for i := 1; i <= 6; i++ {
			if effectSlice := gamedb.GetDaBaoEquipAdditionByTypeAndLv(t, i, lv); len(effectSlice) > 0 {
				effectId := effectSlice[0]
				effectMap[effectId] = 0
				for k, v := range effectSlice {
					if k == 0 {
						continue
					}
					if hero.MapEffects[v] == nil {
						hero.MapEffects[v] = make(map[int]int)
					}
					hero.MapEffects[v][effectId] = 0
				}
			}
		}
	}
	if len(effectMap) > 0 {
		addEffect(effectMap, p)
	}
	return p
}

func calcLabelProps(user *objs.User, heroIndex int) *prop.Prop {
	userLabel := user.Label
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	hero.LabelEffects = make([]int, 0)
	labelCfg := gamedb.GetLabelLabelCfg(userLabel.Id)
	if labelCfg != nil {
		p.Add(labelCfg.NormalAttr)
		effectMap := make(map[int]int)
		var attr map[int]int
		switch userLabel.Job {
		case pb.LABELTYPE_CIVIL_SERVICE:
			attr = labelCfg.WenAttr
			effectMap[labelCfg.WenEffect] = 0
		case pb.LABELTYPE_MILITARY_GENERAL:
			attr = labelCfg.WuAttr
			effectMap[labelCfg.WuEffect] = 0
		}
		if attr != nil && len(attr) > 0 {
			p.Add(attr)
		}
		if len(effectMap) > 0 && heroIndex == constUser.USER_HERO_MAIN_INDEX {
			effectSlice := addEffect(effectMap, p)
			hero.LabelEffects = effectSlice
		}
	}
	return p
}

func calcGodEquipBloodProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	hero := user.Heros[heroIndex]
	for _, godEquip := range hero.GodEquips {
		for i := 1; i <= godEquip.Blood; i++ {
			if godEquipCfg := gamedb.GetGodBloodGodBloodCfg(gamedb.GetRealId(godEquip.Id, i)); godEquipCfg != nil {
				p.Add(godEquipCfg.Attribute)
			}
		}
	}
	return p
}

func calcPrivilegeProps(user *objs.User, heroIndex int) *prop.Prop {
	p := prop.NewProp()
	userPrivilegeId := user.Privilege
	privilegeCfgs := gamedb.GetPrivilegeCfgs()
	for id, cfg := range privilegeCfgs {
		if _, ok := userPrivilegeId[id]; !ok {
			continue
		}
		if effectId := cfg.Privilege[pb.VIPPRIVILEGE_ATTR]; effectId != 0 {
			if effectCfg := gamedb.GetEffectEffectCfg(effectId); effectCfg != nil {
				p.Add(effectCfg.Attribute)
			}
		}
	}
	return p
}
