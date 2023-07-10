package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type Kyevent_fit_lv_up struct {
	*kyEventLog.KyEventPropBase
	BeforeLv int `json:"before_lv"` //升级前等级
	Lv       int `json:"lv"`        //升级后等级
}

type KyEvent_fit_fashion_up struct {
	*kyEventLog.KyEventPropBase
	Id       int `json:"config_id1"` //时装id
	BeforeLv int `json:"before_lv"`  //升级前等级
	Lv       int `json:"lv"`         //升级后等级
}

type KyEvent_fit_skill_star_up struct {
	*kyEventLog.KyEventPropBase
	SkillId  int `json:"config_id1"` //技能id
	BeforeLv int `json:"before_lv"`  //升级前等级
	Lv       int `json:"lv"`         //升级后等级
}

type KyEvent_fit_skill_change struct {
	*kyEventLog.KyEventPropBase
	SkillId     int `json:"config_id1"` //合体技能id
	DownSkillId int `json:"config_id2"` //合体技能id
}

type KyEvent_fit_skill_lv_up struct {
	*kyEventLog.KyEventPropBase
	SkillId  int `json:"skill_id"`  //技能id
	BeforeLv int `json:"before_lv"` //升级前等级
	Lv       int `json:"lv"`        //升级后等级
}

type KyEvent_fit_holy_equip_lv_up struct {
	*kyEventLog.KyEventPropBase
	Type     int `json:"type"`       //圣装类型
	Id       int `json:"config_id1"` //圣装id
	BeforeLv int `json:"before_lv"`  //升级前等级
	Lv       int `json:"lv"`         //升级后等级
}

type KyEvent_fit_holy_equip_suit_skill_change struct {
	*kyEventLog.KyEventPropBase
	SuitType int `json:"config_id1"` //圣装套装类型
	Grade    int `json:"num1"`       //阶数
}

//合体升级
func FitLvUp(user *objs.User, lv int) {
	data := &Kyevent_fit_lv_up{
		KyEventPropBase: getEventBaseProp(user),
		BeforeLv:        lv - 1,
		Lv:              lv,
	}
	writeUserEvent(user, "fit_lv_up", data)
}

//合体时装升级
func FitFashionLvUp(user *objs.User, fashionId, lv int) {
	data := &KyEvent_fit_fashion_up{
		KyEventPropBase: getEventBaseProp(user),
		Id:              fashionId,
		BeforeLv:        lv - 1,
		Lv:              lv,
	}
	writeUserEvent(user, "fit_fashion_lv_up", data)
}

//合体技能升星
func FitSkillStarUp(user *objs.User, skillId, lv int) {
	data := &KyEvent_fit_skill_star_up{
		KyEventPropBase: getEventBaseProp(user),
		SkillId:         skillId,
		BeforeLv:        lv - 1,
		Lv:              lv,
	}
	writeUserEvent(user, "fit_skill_star_up", data)
}

//合体技能镶嵌
func FitSkillChange(user *objs.User, remSkillId, skillId int) {
	data := &KyEvent_fit_skill_change{
		KyEventPropBase: getEventBaseProp(user),
		DownSkillId:     remSkillId,
		SkillId:         skillId,
	}
	writeUserEvent(user, "fit_skill_change", data)
}

//合体技能升级
func FitSkillLvUp(user *objs.User, skillId, lv int) {
	data := &KyEvent_fit_skill_lv_up{
		KyEventPropBase: getEventBaseProp(user),
		SkillId:         skillId,
		BeforeLv:        lv - 1,
		Lv:              lv,
	}
	writeUserEvent(user, "fit_skill_lv_up", data)
}

//合体圣装升级
func FitHolyEquipLvUp(user *objs.User, suitT, id, lv int) {
	data := &KyEvent_fit_holy_equip_lv_up{
		KyEventPropBase: getEventBaseProp(user),
		Type:            suitT,
		Id:              id,
		BeforeLv:        lv - 1,
		Lv:              lv,
	}
	writeUserEvent(user, "fit_holy_equip_lv_up", data)
}

//合体圣装替换技能
func FitHolyEquipSuitSkillChange(user *objs.User, suitT, grade int) {
	data := &KyEvent_fit_holy_equip_suit_skill_change{
		KyEventPropBase: getEventBaseProp(user),
		SuitType:        suitT,
		Grade:           grade,
	}
	writeUserEvent(user, "fit_holy_equip_suit_skill_change", data)
}
