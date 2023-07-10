package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_pet_lv_up struct {
	*kyEventLog.KyEventPropBase
	PetId    int `json:"config_id1"` //战宠id
	BeforeLv int `json:"before_lv"`  //升级前等级
	Lv       int `json:"lv"`         //升级后等级
}

type KyEvent_pet_grade_up struct {
	*kyEventLog.KyEventPropBase
	PetId       int `json:"config_id1"` //战宠id
	BeforeGrade int `json:"before_lv"`  //升级前阶数
	Grade       int `json:"lv"`         //升级后阶数
}

//战宠升级记录
func PetLvUp(user *objs.User, petId, lv int) {
	data := &KyEvent_pet_lv_up{
		KyEventPropBase: getEventBaseProp(user),
		PetId:           petId,
		BeforeLv:        lv - 1,
		Lv:              lv,
	}
	writeUserEvent(user, "pet_lv_up", data)
}

//战宠升阶记录
func PetGradeUp(user *objs.User, petId, grade int) {
	data := &KyEvent_pet_grade_up{
		KyEventPropBase: getEventBaseProp(user),
		PetId:           petId,
		BeforeGrade:     grade - 1,
		Grade:           grade,
	}
	writeUserEvent(user, "pet_grade_up", data)
}
