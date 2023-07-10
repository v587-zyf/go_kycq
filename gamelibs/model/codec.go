package model

import (
	"bytes"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	//"log"
	"strconv"
	"strings"
)

//Jsonable
//可以被json化的模型
type Jsonable interface {
	MarshalToJSON() string
}

func marshalBool(flag bool) int {
	n := 0
	if flag {
		n = 1
	}
	return n
}
func unmarshalBool(n int) bool {
	f := false
	if n == 1 {
		f = true
	}
	return f
}

func (this AccountBan) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k, banInfo := range this {

		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`"%d":["%s","%s","%s","%s"]`, k, strconv.Itoa(banInfo.BanType), banInfo.StartTime, banInfo.EndTime, banInfo.Reason))
		} else {
			buf.WriteString(fmt.Sprintf(`,"%d":["%s","%s","%s","%s"]`, k, strconv.Itoa(banInfo.BanType), banInfo.StartTime, banInfo.EndTime, banInfo.Reason))
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (this *AccountBan) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*BanInfo, 0)
	if len(data) == 0 {
		return nil
	}
	datas := make(map[int][]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for k, v := range datas {

		banInfo := &BanInfo{}
		banInfo.BanType, _ = strconv.Atoi(v[0].(string))
		banInfo.StartTime, _ = v[1].(string)
		banInfo.EndTime, _ = v[2].(string)
		banInfo.Reason, _ = v[3].(string)
		(*this)[k] = banInfo
	}
	return nil
}

func (this *AccountBan) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this AccountBan) Value() (driver.Value, error) {
	return json.Marshal(this)
}

//func (this CrossHeros) MarshalJSON() ([]byte, error) {
//	var buf bytes.Buffer
//	buf.WriteByte('[')
//	for k, hero := range this {
//		h,_ := json.Marshal(hero)
//		buf.Write(h)
//		if k != len(this)-1{
//			buf.Write([]bytes(","))
//		}
//	}
//	buf.WriteByte(']')
//	return buf.Bytes(), nil
//}
//
//func (this *CrossHeros) UnmarshalJSON(data []byte) error {
//	*this = make(map[int]*BanInfo, 0)
//	if len(data) == 0 {
//		return nil
//	}
//	datas := make(map[int][]interface{}, 0)
//	err := json.Unmarshal(data, &datas)
//	if err != nil {
//		return err
//	}
//	for k, v := range datas {
//
//		banInfo := &BanInfo{}
//		banInfo.BanType, _ = strconv.Atoi(v[0].(string))
//		banInfo.StartTime, _ = v[1].(string)
//		banInfo.EndTime, _ = v[2].(string)
//		banInfo.Reason, _ = v[3].(string)
//		(*this)[k] = banInfo
//	}
//	return nil
//}

func (this *CrossHeros) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), this)
	//return this.UnmarshalJSON(value.([]byte))
}

func (this CrossHeros) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this EquipBag) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for _, equip := range this {

		randPropStr := "["
		if len(equip.RandProps) > 0 {
			for _, v := range equip.RandProps {
				randPropStr += fmt.Sprintf("[%d,%d,%d],", v.PropId, v.Color, v.Value)
			}
			randPropStr = randPropStr[:len(randPropStr)-1]
		}
		randPropStr += "]"

		isLock := 0
		if equip.IsLock {
			isLock = 1
		}

		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%d,%s,%d,%d]`, equip.Index, equip.ItemId, randPropStr, isLock, equip.Lucky))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%d,%s,%d,%d]`, equip.Index, equip.ItemId, randPropStr, isLock, equip.Lucky))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *EquipBag) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*Equip, 0)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, v := range datas {
		index := int(v[0].(float64))
		equip := &Equip{Index: index, ItemId: int(v[1].(float64)), IsLock: false, Lucky: int(v[4].(float64))}
		if int(v[3].(float64)) == 1 {
			equip.IsLock = true
		}
		randProps := v[2].([]interface{})
		equip.RandProps = make([]*EquipRandProp, len(randProps))
		for kk, vv := range randProps {
			randProp := vv.([]interface{})
			equip.RandProps[kk] = &EquipRandProp{
				PropId: int(randProp[0].(float64)),
				Color:  int(randProp[1].(float64)),
				Value:  int(randProp[2].(float64)),
			}
		}

		(*this)[index] = equip
	}
	return nil
}

func (this *EquipBag) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this EquipBag) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Equips) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k, equip := range this {

		randPropStr := "["
		if len(equip.RandProps) > 0 {
			for _, v := range equip.RandProps {
				randPropStr += fmt.Sprintf("[%d,%d,%d],", v.PropId, v.Color, v.Value)
			}
			randPropStr = randPropStr[:len(randPropStr)-1]
		}
		randPropStr += "]"

		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`"%d":[%d,%d,%s,%d]`, k, equip.Index, equip.ItemId, randPropStr, equip.Lucky))
		} else {
			buf.WriteString(fmt.Sprintf(`,"%d":[%d,%d,%s,%d]`, k, equip.Index, equip.ItemId, randPropStr, equip.Lucky))
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (this *Equips) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*Equip, 0)
	if len(data) == 0 {
		return nil
	}
	datas := make(map[int][]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for k, v := range datas {
		index := int(v[0].(float64))
		equip := &Equip{Index: index, ItemId: int(v[1].(float64)), Lucky: int(v[3].(float64))}
		randProps := v[2].([]interface{})
		equip.RandProps = make([]*EquipRandProp, len(randProps))
		for kk, vv := range randProps {
			randProp := vv.([]interface{})
			equip.RandProps[kk] = &EquipRandProp{
				PropId: int(randProp[0].(float64)),
				Color:  int(randProp[1].(float64)),
				Value:  int(randProp[2].(float64)),
			}
		}
		(*this)[k] = equip
	}
	return nil
}

func (this *Equips) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Equips) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Bag) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for _, item := range this {
		if item == nil {
			continue
		}
		if item.ItemId <= 0 {
			continue
		}
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%d]`, item.ItemId, item.Count, item.Position, item.EquipIndex))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%d,%d,%d]`, item.ItemId, item.Count, item.Position, item.EquipIndex))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *Bag) UnmarshalJSON(data []byte) error {
	*this = make([]*Item, 0)
	if len(data) == 0 {
		return nil
	}
	mpBag := make([][]interface{}, 0)
	err := json.Unmarshal(data, &mpBag)
	if err != nil {
		return err
	}
	for _, v := range mpBag {
		item := &Item{ItemId: int(v[0].(float64)), Count: int(v[1].(float64)), Position: int(v[2].(float64)), EquipIndex: int(v[3].(float64))}
		*this = append(*this, item)
	}
	return nil
}

func (this *Bag) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Bag) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Fabaos) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for _, item := range this {
		if item.Id <= 0 {
			continue
		}
		skill := "["
		if len(item.Skill) > 0 {
			for _, v := range item.Skill {
				skill += fmt.Sprintf("%d,", v)
			}
			skill = skill[:len(skill)-1]
		}
		skill += "]"
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%s]`, item.Id, item.Level, item.Exp, skill))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%d,%d,%s]`, item.Id, item.Level, item.Exp, skill))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *Fabaos) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*Fabao, 0)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, v := range datas {
		id := int(v[0].(float64))
		fabao := &Fabao{Id: id, Level: int(v[1].(float64)), Exp: int(v[2].(float64))}
		skills := v[3].([]interface{})
		for _, vv := range skills {
			fabao.Skill = append(fabao.Skill, int(vv.(float64)))
		}
		(*this)[id] = fabao
	}
	return nil
}

func (this *Fabaos) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Fabaos) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this GodEquips) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for _, item := range this {
		if item.Id <= 0 {
			continue
		}
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%d,%d]`, item.Id, item.Lv, item.Blood))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%d,%d]`, item.Id, item.Lv, item.Blood))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *GodEquips) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*GodEquip, 0)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, v := range datas {
		id := int(v[0].(float64))
		godEquip := &GodEquip{Id: id, Lv: int(v[1].(float64))}
		if len(v) > 2 {
			godEquip.Blood = int(v[2].(float64))
		}
		(*this)[id] = godEquip
	}
	return nil
}

func (this *GodEquips) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this GodEquips) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Juexues) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for _, item := range this {
		if item.Id <= 0 {
			continue
		}
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%d]`, item.Id, item.Lv))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%d]`, item.Id, item.Lv))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *Juexues) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*Juexue, 0)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, v := range datas {
		id := int(v[0].(float64))
		juexue := &Juexue{Id: id, Lv: int(v[1].(float64))}
		(*this)[id] = juexue
	}
	return nil
}

func (this *Juexues) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Juexues) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Holyarms) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for id, item := range this {
		skill := "["
		if len(item.Skill) > 0 {
			for hlv, lv := range item.Skill {
				skill += fmt.Sprintf("[%d,%d],", hlv, lv)
			}
			skill = skill[:len(skill)-1]
		}
		skill += "]"
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%s]`, id, item.Level, item.Exp, skill))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%d,%d,%s]`, id, item.Level, item.Exp, skill))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *Holyarms) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*Holyarm)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, v := range datas {
		id := int(v[0].(float64))
		info := &Holyarm{Level: int(v[1].(float64)), Exp: int(v[2].(float64)), Skill: make(IntKv)}
		skills := v[3].([]interface{})
		for _, skill := range skills {
			skillArr := skill.([]interface{})
			info.Skill[int(skillArr[0].(float64))] = int(skillArr[1].(float64))
		}
		(*this)[id] = info
	}
	return nil
}

func (this *Holyarms) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Holyarms) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Wings) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for _, item := range this {
		if item.Id <= 0 {
			continue
		}
		isWear := 0
		if item.IsWear {
			isWear = 1
		}
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%d,%d]`, item.Id, item.Exp, isWear))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%d,%d]`, item.Id, item.Exp, isWear))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *Wings) UnMarshalJSON(data []byte) error {
	*this = make(map[int]*Wing)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for k, v := range datas {
		id := int(v[0].(float64))
		wing := &Wing{
			Id:  id,
			Exp: int(v[1].(float64)),
		}
		if int(v[2].(float64)) == 1 {
			wing.IsWear = true
		}
		(*this)[k] = wing
	}
	return nil
}

func (this *Wings) Scan(value interface{}) error {
	return this.UnMarshalJSON(value.([]byte))
}

func (this Wings) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Rein) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	buf.WriteString(fmt.Sprintf(`%d,%d`, this.Id, this.Exp))
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *Rein) UnMarshalJSON(data []byte) error {
	*this = Rein{}
	if len(data) == 0 {
		return nil
	}
	var arr []int
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}
	(*this).Id = arr[0]
	(*this).Exp = arr[1]
	return nil
}

func (this *Rein) Scan(value interface{}) error {
	return this.UnMarshalJSON(value.([]byte))
}

func (this Rein) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this ReinCosts) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for _, item := range this {
		if item.Id <= 0 {
			continue
		}
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%d,%d]`, item.Id, item.Num, item.Date))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%d,%d]`, item.Id, item.Num, item.Date))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *ReinCosts) UnMarshalJSON(data []byte) error {
	*this = make(map[int]*ReinCost)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, v := range datas {
		id := int(v[0].(float64))
		reinCost := &ReinCost{
			Id:   id,
			Num:  int(v[1].(float64)),
			Date: int(v[2].(float64)),
		}
		(*this)[id] = reinCost
	}
	return nil
}

func (this *ReinCosts) Scan(value interface{}) error {
	return this.UnMarshalJSON(value.([]byte))
}

func (this ReinCosts) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this FieldBoss) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	cdStr := "["
	if len(this.CD) > 0 {
		for stageId, cdTime := range this.CD {
			cdStr += fmt.Sprintf(`[%d,%d],`, stageId, cdTime)
		}
		cdStr = cdStr[:len(cdStr)-1]
	}
	cdStr += "]"
	first := 0
	if this.FirstReceive {
		first = 1
	}
	buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%s,%d]`, this.DareNum, this.BuyNum, this.ResetTime, cdStr, first))
	return buf.Bytes(), nil
}

func (this *FieldBoss) UnMarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	(*this).DareNum = int(datas[0].(float64))
	(*this).BuyNum = int(datas[1].(float64))
	(*this).ResetTime = int(datas[2].(float64))
	cdInterface := datas[3].([]interface{})
	cdMap := make(IntKv)
	for _, cd := range cdInterface {
		cdArr := cd.([]interface{})
		cdMap[int(cdArr[0].(float64))] = int(cdArr[1].(float64))
	}
	(*this).CD = cdMap
	if len(datas) > 4 {
		first := false
		if int(datas[4].(float64)) == 1 {
			first = true
		}
		(*this).FirstReceive = first
	}
	return nil
}

func (this *FieldBoss) Scan(value interface{}) error {
	return this.UnMarshalJSON(value.([]byte))
}

func (this FieldBoss) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this ExpStage) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	expStagesStr := "["
	if len(this.ExpStages) > 0 {
		for stageId, exp := range this.ExpStages {
			expStagesStr += fmt.Sprintf(`[%d,%d],`, stageId, exp)
		}
		expStagesStr = expStagesStr[:len(expStagesStr)-1]
	}
	expStagesStr += "]"
	appraiseStr := "["
	if len(this.Appraise) > 0 {
		for stageId, appraise := range this.Appraise {
			appraiseStr += fmt.Sprintf(`[%d,%d],`, stageId, appraise)
		}
		appraiseStr = appraiseStr[:len(appraiseStr)-1]
	}
	appraiseStr += "]"
	buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%s,%s,%d]`, this.DareNum, this.BuyNum, this.ResetTime, expStagesStr, appraiseStr, this.Layer))
	return buf.Bytes(), nil
}

func (this *ExpStage) UnMarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		logger.Error("ExpStage UnMarshalJSON error err is %v", err)
		return err
	}
	(*this).DareNum = int(datas[0].(float64))
	(*this).BuyNum = int(datas[1].(float64))
	(*this).ResetTime = int(datas[2].(float64))

	expStagesMap := make(IntKv)
	expStagesInterface := datas[3].([]interface{})
	for _, v := range expStagesInterface {
		expStageArr := v.([]interface{})
		expStagesMap[int(expStageArr[0].(float64))] = int(expStageArr[1].(float64))
	}
	(*this).ExpStages = expStagesMap
	appraiseMap := make(IntKv)
	if len(datas) > 4 {
		appraiseInterface := datas[4].([]interface{})
		for _, appraise := range appraiseInterface {
			appraiseArr := appraise.([]interface{})
			appraiseMap[int(appraiseArr[0].(float64))] = int(appraiseArr[1].(float64))
		}
	}
	(*this).Appraise = appraiseMap
	layer := 1
	if len(datas) > 5 {
		layer = int(datas[5].(float64))
	}
	(*this).Layer = layer
	return nil
}

func (this *ExpStage) Scan(value interface{}) error {
	return this.UnMarshalJSON(value.([]byte))
}

func (this ExpStage) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Arena) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%d]`, this.DareNum, this.DareDate, this.BuyDareNums, this.BuyDareNum))
	return buf.Bytes(), nil
}

func (this *Arena) UnMarshalJSON(data []byte) error {
	datas := make([]int, 0)
	if len(data) == 0 {
		return nil
	}
	err := json.Unmarshal(data, &datas)
	if err != nil {
		logger.Error("Arena UnMarshalJSON error err is %v", err)
		return err
	}
	(*this).DareNum = datas[0]
	(*this).DareDate = datas[1]
	(*this).BuyDareNums = datas[2]
	(*this).BuyDareNum = datas[3]
	return nil
}

func (this *Arena) Scan(value interface{}) error {
	return this.UnMarshalJSON(value.([]byte))
}

func (this Arena) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this BagInfo) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k, item := range this {
		bagInfoUnitStr, _ := json.Marshal(item)
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`"%d":%s`, k, bagInfoUnitStr))
		} else {
			buf.WriteString(fmt.Sprintf(`,"%d":%s`, k, bagInfoUnitStr))
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (this *BagInfo) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*BagInfoUnit, 0)
	if len(data) == 0 {
		return nil
	}
	mpBag := make(map[int]*BagInfoUnit, 0)
	err := json.Unmarshal(data, &mpBag)
	if err != nil {
		return err
	}
	for k, v := range mpBag {

		(*this)[k] = v
	}
	return nil
}

func (this *BagInfo) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this BagInfo) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this *MainLineTask) Scan(value interface{}) error {
	if len(value.([]byte)) == 0 {
		return nil
	}
	var mainTask []int
	err := json.Unmarshal(value.([]byte), &mainTask)
	if err != nil {
		return err
	}
	*this = MainLineTask{
		TaskId:  mainTask[0],
		Process: mainTask[1],
	}
	if len(mainTask) > 2 {
		*this = MainLineTask{
			TaskId:      mainTask[0],
			Process:     mainTask[1],
			MarkProcess: mainTask[2],
		}
	}

	return nil
}

func (this MainLineTask) Value() (driver.Value, error) {
	return fmt.Sprintf("[%d,%d,%d]", this.TaskId, this.Process, this.MarkProcess), nil
}

func (this *Counts) Scan(value interface{}) error {
	*this = make(map[string][2]int)
	if len(value.([]byte)) == 0 {
		return nil
	}
	return json.Unmarshal(value.([]byte), this)
}

func (this Counts) Value() (driver.Value, error) {
	return this.MarshalJSON()
}

func (this Counts) MarshalJSON() ([]byte, error) {
	if len(this) <= 0 {
		return []byte("{}"), nil
	}
	counts := make(map[string][2]int)
	for k, v := range this {
		counts[k] = v
	}
	return json.Marshal(counts)
}

func (this ExData) MarshalJSON() ([]byte, error) {
	if len(this) <= 0 {
		return []byte("{}"), nil
	}
	var exData = make(map[string]*json.RawMessage)
	for k, v := range this {
		ks := strconv.Itoa(k)
		exData[ks] = v
	}
	return json.Marshal(exData)
}

func (this *ExData) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*json.RawMessage, 0)
	if len(data) == 0 {
		return nil
	}
	var exData map[string]*json.RawMessage
	if err := json.Unmarshal(data, &exData); err != nil {
		return err
	}

	for k, v := range exData {
		if ki, err := strconv.Atoi(k); err != nil {
			return err
		} else {
			(*this)[ki] = v
		}
	}
	return nil
}

func (this *ExData) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this ExData) Value() (driver.Value, error) {
	return this.MarshalJSON()
}

func (this IntSlice) MarshalJSON() ([]byte, error) {
	var result = make([]string, len(this))
	for i, v := range this {
		result[i] = strconv.Itoa(v)
	}
	return []byte("[" + strings.Join(result, ",") + "]"), nil
}

func (this *IntSlice) UnmarshalJSON(data []byte) error {
	if len(data) <= 2 {
		*this = make([]int, 0)
		return nil
	}
	strs := strings.Split(string(data[1:len(data)-1]), ",")
	*this = make([]int, len(strs))
	for i, str := range strs {
		(*this)[i], _ = strconv.Atoi(strings.TrimSpace(str))
	}
	return nil
}

func (this *IntSlice) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this IntSlice) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this StringSlice) MarshalJSON() ([]byte, error) {
	return []byte("[" + strings.Join(this, ",") + "]"), nil
}

func (this *StringSlice) UnmarshalJSON(data []byte) error {
	if len(data) <= 2 {
		*this = make([]string, 0)
		return nil
	}
	*this = strings.Split(string(data[1:len(data)-1]), ",")
	return nil
}

func (this *StringSlice) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this StringSlice) Value() (driver.Value, error) {
	return this.MarshalJSON()
}

func (this IntKv) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k, v := range this {
		if buf.Len() > 1 {
			buf.WriteByte(',')
		}
		buf.WriteString(fmt.Sprintf(`"%d":%d`, k, v))
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (this *IntKv) UnmarshalJSON(data []byte) error {
	*this = make(map[int]int)
	if len(data) <= 2 {
		return nil
	}
	var mp map[string]int
	err := json.Unmarshal(data, &mp)
	if err != nil {
		return err
	}
	for k, v := range mp {
		key, _ := strconv.Atoi(k)
		(*this)[key] = v
	}
	return nil
}

func (this *IntKv) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this IntKv) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this IntKv) KeysInt32() []int32 {
	iret := make([]int32, len(this))
	i := 0
	for k, _ := range this {
		iret[i] = int32(k)
		i++
	}
	return iret
}

func (this Float64Slice) MarshalJSON() ([]byte, error) {
	var result = make([]string, len(this))
	for i, v := range this {
		result[i] = strconv.FormatFloat(v, 'E', -1, 64)
	}
	return []byte("[" + strings.Join(result, ",") + "]"), nil
}

func (this *Float64Slice) UnmarshalJSON(data []byte) error {
	if len(data) <= 2 {
		*this = make([]float64, 0)
		return nil
	}
	strs := strings.Split(string(data[1:len(data)-1]), ",")
	*this = make([]float64, len(strs))
	for i, str := range strs {
		(*this)[i], _ = strconv.ParseFloat(strings.TrimSpace(str), 64)
	}
	return nil
}

func (this *Float64Slice) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Float64Slice) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Int64Kv) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k, v := range this {
		if buf.Len() > 1 {
			buf.WriteByte(',')
		}
		buf.WriteString(fmt.Sprintf(`"%d":%d`, k, v))
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (this *Int64Kv) UnmarshalJSON(data []byte) error {
	*this = make(map[int]int64)
	if len(data) <= 2 {
		return nil
	}
	var mp map[string]int64
	err := json.Unmarshal(data, &mp)
	if err != nil {
		return err
	}
	for k, v := range mp {
		key, _ := strconv.Atoi(k)
		(*this)[key] = v
	}
	return nil
}

func (this *Int64Kv) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Int64Kv) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this *MapIntKv) MarshalJSON() ([]byte, error) {

	var buf bytes.Buffer
	buf.WriteByte('{')
	for k, v := range *this {
		if buf.Len() > 1 {
			buf.WriteByte(',')
		}
		inerIntKv, _ := json.Marshal(v)
		buf.WriteString(fmt.Sprintf(`"%d":%s`, k, string(inerIntKv)))
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (this *MapIntKv) UnmarshalJSON(data []byte) error {
	*this = make(map[int]IntKv)
	if len(data) <= 2 {
		return nil
	}
	var mp map[string]map[string]int
	err := json.Unmarshal(data, &mp)
	if err != nil {
		return err
	}
	for k, v := range mp {
		key, _ := strconv.Atoi(k)
		(*this)[key] = make(IntKv)

		for inerK, inerV := range v {
			inerkey, _ := strconv.Atoi(inerK)
			(*this)[key][inerkey] = inerV
		}
	}
	return nil
}
func (this *MapIntKv) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this MapIntKv) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func MarshalMapJsonable(m map[int]Jsonable) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k, item := range m {
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`"%d":%s`, k, item.MarshalToJSON()))
		} else {
			buf.WriteString(fmt.Sprintf(`,"%d":%s`, k, item.MarshalToJSON()))
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func UnmarshalJsonable(data []byte, fn func(intK int, v interface{}) Jsonable) (map[int]Jsonable, error) {
	if len(data) == 0 {
		return nil, nil
	}

	mpBag := make(map[string]interface{})
	err := json.Unmarshal(data, &mpBag)
	if err != nil {
		return nil, err
	}
	mapJsonable := make(map[int]Jsonable)
	for k, v := range mpBag {
		intK, _ := strconv.Atoi(k)
		mapJsonable[intK] = fn(intK, v)
	}
	return mapJsonable, nil
}

func (this *IntStringKv) Scan(value interface{}) error {

	*this = make(map[int]string)
	data := value.([]byte)
	if len(data) <= 2 {
		return nil
	}
	var mp map[string]string
	err := json.Unmarshal(data, &mp)
	if err != nil {
		return err
	}
	for k, v := range mp {
		key, _ := strconv.Atoi(k)
		(*this)[key] = v
	}
	return nil
}

func (this IntStringKv) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this StringIntKv) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k, v := range this {
		if buf.Len() > 1 {
			buf.WriteByte(',')
		}
		buf.WriteString(fmt.Sprintf(`"%s":%d`, k, v))
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (this *StringIntKv) UnmarshalJSON(data []byte) error {
	*this = make(map[string]int)
	if len(data) <= 2 {
		return nil
	}
	var mp map[string]string
	err := json.Unmarshal(data, &mp)
	if err != nil {
		return err
	}
	for k, v := range mp {
		val, _ := strconv.Atoi(v)
		(*this)[k] = val
	}
	return nil
}

func (this *StringIntKv) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this StringIntKv) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this *ItemSlice) Scan(value interface{}) error {
	*this = make(ItemSlice, 0)
	if value == nil {
		return nil
	}
	return json.Unmarshal(value.([]byte), this)
}

func (this ItemSlice) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this *ItemSlice2) Scan(value interface{}) error {
	*this = make(ItemSlice2, 0)
	if value == nil {
		return nil
	}
	return json.Unmarshal(value.([]byte), this)
}

func (this ItemSlice2) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Tower) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	get := "["
	if len(this.Lottery) > 0 {
		for _, v := range this.Lottery {
			get += fmt.Sprintf("%d,", v)
		}
		get = get[:len(get)-1]
	}
	get += "]"
	buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%d,%s]`, this.TowerLv, this.LotteryNum, this.DayAwardState, this.LotteryId, get))
	return buf.Bytes(), nil
}

func (this *Tower) UnmarshalJSON(data []byte) error {

	if len(data) == 0 {
		return nil
	}
	list := make([]interface{}, 0)
	err := json.Unmarshal(data, &list)
	if err != nil {
		return err
	}
	*this = Tower{
		TowerLv:       int(list[0].(float64)),
		LotteryNum:    int(list[1].(float64)),
		DayAwardState: int(list[2].(float64)),
		LotteryId:     int(list[3].(float64)),
	}
	lottery := list[4].([]interface{})
	for _, v := range lottery {
		(*this).Lottery = append((*this).Lottery, int(v.(float64)))
	}
	return nil
}

func (this *Tower) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Tower) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Display) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%d,%d,%d]`, this.ClothItemId, this.ClothType, this.WeaponItemId, this.WeaponType, this.WingId, this.MagicCircleLvId))
	return buf.Bytes(), nil
}

func (this *Display) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	list := make([]interface{}, 0)
	err := json.Unmarshal(data, &list)
	if err != nil {
		return err
	}
	(*this).ClothItemId = int(list[0].(float64))
	(*this).ClothType = int(list[1].(float64))
	(*this).WeaponItemId = int(list[2].(float64))
	(*this).WeaponType = int(list[3].(float64))
	(*this).WingId = int(list[4].(float64))
	(*this).MagicCircleLvId = int(list[5].(float64))
	return nil
}

func (this *Display) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Display) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Shop) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	shopItemStr := "["
	if len(this.ShopItem) > 0 {
		for shopType, shopItem := range this.ShopItem {
			shopInfoStr := "["
			if len(shopItem) > 0 {
				for id, buyNum := range shopItem {
					shopInfoStr += fmt.Sprintf("[%d,%d],", id, buyNum)
				}
				shopInfoStr = shopInfoStr[:len(shopInfoStr)-1]
			}
			shopInfoStr += "]"
			shopItemStr += fmt.Sprintf("[%d,%s],", shopType, shopInfoStr)
		}
		shopItemStr = shopItemStr[:len(shopItemStr)-1]
	}
	shopItemStr += "]"
	buf.WriteString(fmt.Sprintf(`[%d,%s]`, this.ResetTime, shopItemStr))
	return buf.Bytes(), nil
}

func (this *Shop) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	shopItemInterface := datas[1].([]interface{})
	shopItemMap := make(MapIntKv)
	for _, shopItem := range shopItemInterface {
		shopItemArr := shopItem.([]interface{})
		shopInfoInterface := shopItemArr[1].([]interface{})
		shopInfoMap := make(IntKv)
		for _, shopInfo := range shopInfoInterface {
			shopInfoArr := shopInfo.([]interface{})
			shopInfoMap[int(shopInfoArr[0].(float64))] = int(shopInfoArr[1].(float64))
		}
		shopItemMap[int(shopItemArr[0].(float64))] = shopInfoMap
	}
	(*this).ResetTime = int(datas[0].(float64))
	(*this).ShopItem = shopItemMap
	return nil
}

func (this *Shop) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Shop) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Zodiacs) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for pos, equipUnit := range this {
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`"%d":[%d]`, pos, equipUnit.Id))
		} else {
			buf.WriteString(fmt.Sprintf(`,"%d":[%d]`, pos, equipUnit.Id))
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (this *Zodiacs) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*SpecialEquipUnit)
	if len(data) == 0 {
		return nil
	}
	list := make(map[int][]interface{}, 0)
	err := json.Unmarshal(data, &list)
	if err != nil {
		return err
	}
	for k, v := range list {
		id := int(v[0].(float64))
		zodiac := &SpecialEquipUnit{Id: id}
		(*this)[k] = zodiac
	}
	return nil
}

func (this *Zodiacs) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Zodiacs) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Kingarms) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for pos, equipUnit := range this {
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`"%d":[%d]`, pos, equipUnit.Id))
		} else {
			buf.WriteString(fmt.Sprintf(`,"%d":[%d]`, pos, equipUnit.Id))
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (this *Kingarms) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*SpecialEquipUnit)
	if len(data) == 0 {
		return nil
	}
	list := make(map[int][]interface{}, 0)
	err := json.Unmarshal(data, &list)
	if err != nil {
		return err
	}
	for k, v := range list {
		id := int(v[0].(float64))
		zodiac := &SpecialEquipUnit{Id: id}
		(*this)[k] = zodiac
	}
	return nil
}

func (this *Kingarms) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Kingarms) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Skills) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for _, unit := range this {
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%d]`, unit.Id, unit.Lv, unit.StartTime, unit.EndTime))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%d,%d,%d]`, unit.Id, unit.Lv, unit.StartTime, unit.EndTime))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *Skills) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*SkillUnit)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]int, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, v := range datas {
		skillId := v[0]
		(*this)[skillId] = &SkillUnit{
			Id:        skillId,
			Lv:        v[1],
			StartTime: int64(v[2]),
			EndTime:   int64(v[3]),
		}
	}
	return nil
}

func (this *Skills) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Skills) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this OnlineAward) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`[%d,%d,[%s]]`, this.Day, this.OnlineTime, common.JoinIntSlice(this.GetAwardIds, ",")))
	return buf.Bytes(), nil
}

func (this *OnlineAward) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	list := make([]interface{}, 0)
	err := json.Unmarshal(data, &list)
	if err != nil {
		return err
	}
	(*this).Day = int(list[0].(float64))
	(*this).OnlineTime = int(list[1].(float64))
	awardIds := list[2].([]interface{})
	for _, v := range awardIds {
		(*this).GetAwardIds = append((*this).GetAwardIds, int(v.(float64)))
	}
	return nil
}

func (this *OnlineAward) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this OnlineAward) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this DayStateRecord) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	monthCardStr := "["
	if len(this.MonthCardReceive) > 0 {
		for t := range this.MonthCardReceive {
			monthCardStr += fmt.Sprintf("%d,", t)
		}
		monthCardStr = monthCardStr[:len(monthCardStr)-1]
	}
	monthCardStr += "]"

	buf.WriteString(fmt.Sprintf(`[%d,%d,%s,%d]`, this.Day, this.RankWorship, monthCardStr, this.RechargeResetTime))
	return buf.Bytes(), nil
}

func (this *DayStateRecord) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	list := make([]interface{}, 0)
	err := json.Unmarshal(data, &list)
	if err != nil {
		return err
	}
	monthCardInterface := list[2].([]interface{})
	monthCardMap := make(IntKv)
	for _, monthCard := range monthCardInterface {
		monthCardMap[int(monthCard.(float64))] = 0
	}
	(*this).Day = int(list[0].(float64))
	(*this).RankWorship = int(list[1].(float64))
	(*this).MonthCardReceive = monthCardMap
	if len(list) > 3 {
		(*this).RechargeResetTime = int(list[3].(float64))
	}
	return nil
}

func (this *DayStateRecord) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this DayStateRecord) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Wear) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	atlasStr := "["
	if len(this.AtlasWear) > 0 {
		for id := range this.AtlasWear {
			atlasStr += fmt.Sprintf(`%d,`, id)
		}
		atlasStr = atlasStr[:len(atlasStr)-1]
	}
	atlasStr += "]"
	buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%s,%d,%d]`, this.FashionWeaponId, this.FashionClothId, this.WingId, atlasStr, this.MagicCircleLvId, this.TitleId))
	return buf.Bytes(), nil
}

func (this *Wear) UnmarshalJSON(data []byte) error {
	(*this).AtlasWear = make(IntKv)
	if len(data) == 0 {
		return nil
	}
	list := make([]interface{}, 0)
	err := json.Unmarshal(data, &list)
	if err != nil {
		return err
	}
	(*this).FashionWeaponId = int(list[0].(float64))
	(*this).FashionClothId = int(list[1].(float64))
	(*this).WingId = int(list[2].(float64))
	atlasMap := make(IntKv)
	atlasArr := list[3].([]interface{})
	for _, id := range atlasArr {
		atlasMap[int(id.(float64))] = int(id.(float64))
	}
	(*this).AtlasWear = atlasMap
	(*this).MagicCircleLvId = int(list[4].(float64))
	if len(list) > 5 {
		(*this).TitleId = int(list[5].(float64))
	}
	return nil
}

func (this *Wear) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Wear) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this UserWear) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`[%d,%d]`, this.PetId, this.FitFashionId))
	return buf.Bytes(), nil
}

func (this *UserWear) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	list := make([]interface{}, 0)
	err := json.Unmarshal(data, &list)
	if err != nil {
		return err
	}
	(*this).PetId = int(list[0].(float64))
	(*this).FitFashionId = int(list[1].(float64))
	return nil
}

func (this *UserWear) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this UserWear) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Panaceas) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for id, item := range this {
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%d,%d]`, id, item.Number, item.Numbers))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%d,%d]`, id, item.Number, item.Numbers))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *Panaceas) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*PanaceaUnit)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]int, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, item := range datas {
		(*this)[item[0]] = &PanaceaUnit{
			Number:  item[1],
			Numbers: item[2],
		}
	}

	return nil
}

func (this *Panaceas) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Panaceas) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Jewels) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for pos, jewel := range this {
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%d]`, pos, jewel.One, jewel.Two, jewel.Three))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%d,%d,%d]`, pos, jewel.One, jewel.Two, jewel.Three))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *Jewels) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*Jewel)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]int, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, val := range datas {
		(*this)[val[0]] = &Jewel{
			One:   val[1],
			Two:   val[2],
			Three: val[3],
		}
	}
	return nil
}

func (this *Jewels) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Jewels) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Fashions) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for _, v := range this {
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%d]`, v.Id, v.Lv))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%d]`, v.Id, v.Lv))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *Fashions) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*Fashion)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]int, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, val := range datas {
		(*this)[val[0]] = &Fashion{
			Id: val[0],
			Lv: val[1],
		}
	}
	return nil
}

func (this *Fashions) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Fashions) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Sign) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	signDayStr := "["
	if len(this.SignDay) > 0 {
		for day := range this.SignDay {
			signDayStr += fmt.Sprintf(`%d,`, day)
		}
		signDayStr = signDayStr[:len(signDayStr)-1]
	}
	signDayStr += "]"
	cumulativeStr := "["
	if len(this.Cumulative) > 0 {
		for day := range this.Cumulative {
			cumulativeStr += fmt.Sprintf(`%d,`, day)
		}
		cumulativeStr = cumulativeStr[:len(cumulativeStr)-1]
	}
	cumulativeStr += "]"
	buf.WriteString(fmt.Sprintf(`%d,%s,%d,%s,%d`, this.Count, signDayStr, this.ResetTime, cumulativeStr, this.ContinuitySign))
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *Sign) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}

	signDayMap := make(map[int]int)
	signDayInterface := datas[1].([]interface{})
	for _, signDay := range signDayInterface {
		signDayMap[int(signDay.(float64))] = int(signDay.(float64))
	}

	cumulativeMap := make(map[int]int)
	cumulativesInterface := datas[3].([]interface{})
	for _, cumulatives := range cumulativesInterface {
		cumulativeMap[int(cumulatives.(float64))] = int(cumulatives.(float64))
	}

	(*this).Count = int(datas[0].(float64))
	(*this).SignDay = signDayMap
	(*this).ResetTime = int(datas[2].(float64))
	(*this).Cumulative = cumulativeMap
	if len(datas) > 4 {
		(*this).ContinuitySign = int(datas[4].(float64))
	}
	return nil
}

func (this *Sign) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Sign) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Inside) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	acupointStr := "["
	if len(this.Acupoint) > 0 {
		for pos, id := range this.Acupoint {
			acupointStr += fmt.Sprintf(`[%d,%d],`, pos, id)
		}
		acupointStr = acupointStr[:len(acupointStr)-1]
	}
	acupointStr += "]"

	skillStr := "["
	if len(this.Skill) > 0 {
		for id, skill := range this.Skill {
			skillStr += fmt.Sprintf(`[%d,%d,%d],`, id, skill.Level, skill.Exp)
		}
		skillStr = skillStr[:len(skillStr)-1]
	}
	skillStr += "]"
	buf.WriteString(fmt.Sprintf(`[%s,%s]`, acupointStr, skillStr))
	return buf.Bytes(), nil
}

func (this *Inside) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([][][]int, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	acupointMap := make(map[int]int)
	for _, data := range datas[0] {
		acupointMap[data[0]] = data[1]
	}
	skillMap := make(map[int]*InsideSkill)
	for _, data := range datas[1] {
		skillMap[data[0]] = &InsideSkill{
			Level: data[1],
			Exp:   data[2],
		}
	}
	(*this).Acupoint = acupointMap
	(*this).Skill = skillMap
	return nil
}

func (this *Inside) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Inside) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Rings) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for pos, unit := range this {
		phantomStr := "["
		if len(unit.Phantom) > 0 {
			for pos, phantom := range unit.Phantom {
				skillStr := "["
				if len(phantom.Skill) > 0 {
					for id, lv := range phantom.Skill {
						skillStr += fmt.Sprintf(`[%d,%d],`, id, lv)
					}
					skillStr = skillStr[:len(skillStr)-1]
				}
				skillStr += "]"
				phantomStr += fmt.Sprintf(`[%d,%d,%d,%s],`, pos, phantom.Talent, phantom.Phantom, skillStr)
			}
			phantomStr = phantomStr[:len(phantomStr)-1]
		}
		phantomStr += "]"

		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%d,%d,%s]`, pos, unit.Rid, unit.Strengthen, unit.Pid, unit.Talent, phantomStr))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%d,%d,%d,%d,%s]`, pos, unit.Rid, unit.Strengthen, unit.Pid, unit.Talent, phantomStr))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *Rings) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*RingUnit)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]interface{}, 0)
	if err := json.Unmarshal(data, &datas); err != nil {
		return err
	}
	for _, data := range datas {
		pos := int(data[0].(float64))
		ringUnit := &RingUnit{
			Rid:        int(data[1].(float64)),
			Strengthen: int(data[2].(float64)),
			Pid:        int(data[3].(float64)),
			Talent:     int(data[4].(float64)),
			Phantom:    make(map[int]*RingPhantom),
		}
		phantomInterface := data[5].([]interface{})
		for _, phantom := range phantomInterface {
			phantomArr := phantom.([]interface{})
			skillInterface := phantomArr[3].([]interface{})
			skillMap := make(IntKv)
			for _, skill := range skillInterface {
				skillArr := skill.([]interface{})
				skillMap[int(skillArr[0].(float64))] = int(skillArr[1].(float64))
			}
			ringUnit.Phantom[int(phantomArr[0].(float64))] = &RingPhantom{
				Talent:  int(phantomArr[1].(float64)),
				Phantom: int(phantomArr[2].(float64)),
				Skill:   skillMap,
			}
		}
		(*this)[pos] = ringUnit
	}
	return nil
}

func (this *Rings) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Rings) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Mining) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%d,%d,%d,%d]`, this.WorkTime, this.WorkNum, this.RobNum, this.BuyNum, this.Miner, this.Luck, this.ResetTime))
	return buf.Bytes(), nil
}

func (this *Mining) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	if err := json.Unmarshal(data, &datas); err != nil {
		return err
	}
	(*this).WorkTime = int(datas[0].(float64))
	(*this).WorkNum = int(datas[1].(float64))
	(*this).RobNum = int(datas[2].(float64))
	(*this).BuyNum = int(datas[3].(float64))
	(*this).Miner = int(datas[4].(float64))
	(*this).Luck = int(datas[5].(float64))
	(*this).ResetTime = int(datas[6].(float64))
	return nil
}

func (this *Mining) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Mining) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Pets) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for id, item := range this {
		skillStr := "["
		if len(item.Skill) > 0 {
			for skillId := range item.Skill {
				skillStr += strconv.Itoa(skillId) + ","
			}
			skillStr = skillStr[:len(skillStr)-1]
		}
		skillStr += "]"
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%d,%d,%s]`, id, item.Lv, item.Exp, item.Grade, item.Break, skillStr))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%d,%d,%d,%d,%s]`, id, item.Lv, item.Exp, item.Grade, item.Break, skillStr))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *Pets) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*Pet)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, item := range datas {
		id := int(item[0].(float64))
		pet := &Pet{
			Lv:    int(item[1].(float64)),
			Exp:   int(item[2].(float64)),
			Grade: int(item[3].(float64)),
			Break: int(item[4].(float64)),
			Skill: make(IntKv),
		}
		skillInterface := item[5].([]interface{})
		for _, skillId := range skillInterface {
			pet.Skill[int(skillId.(float64))] = 0
		}
		(*this)[id] = pet
	}

	return nil
}

func (this *Pets) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Pets) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this EquipClears) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for pos, item := range this {
		propStr := "["
		if len(item) > 0 {
			for _, equipClearUnit := range item {
				propStr += fmt.Sprintf(`[%d,%d,%d,%d],`, equipClearUnit.Grade, equipClearUnit.Color, equipClearUnit.PropId, equipClearUnit.Value)
			}
			propStr = propStr[:len(propStr)-1]
		}
		propStr += "]"
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%s]`, pos, propStr))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%s]`, pos, propStr))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *EquipClears) UnmarshalJSON(data []byte) error {
	*this = make(map[int][]*EquipClearUnit)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, v := range datas {
		pos := int(v[0].(float64))
		propInterface := v[1].([]interface{})
		propSlice := make([]*EquipClearUnit, 0)
		for _, prop := range propInterface {
			propArr := prop.([]interface{})
			propSlice = append(propSlice, &EquipClearUnit{
				Grade:  int(propArr[0].(float64)),
				Color:  int(propArr[1].(float64)),
				PropId: int(propArr[2].(float64)),
				Value:  int(propArr[3].(float64)),
			})
		}
		(*this)[pos] = propSlice
	}
	return nil
}

func (this *EquipClears) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this EquipClears) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this DarkPalace) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%d]`, this.DareNum, this.BuyNum, this.ResetTime, this.HelpNum))
	return buf.Bytes(), nil
}

func (this *DarkPalace) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]int, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	(*this).DareNum = datas[0]
	(*this).BuyNum = datas[1]
	(*this).ResetTime = datas[2]
	(*this).HelpNum = datas[3]
	return nil
}

func (this *DarkPalace) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this DarkPalace) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this PersonBosses) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	dareStr := "["
	if len(this.DareNum) > 0 {
		for stageId, num := range this.DareNum {
			dareStr += fmt.Sprintf(`[%d,%d],`, stageId, num)
		}
		dareStr = dareStr[:len(dareStr)-1]
	}
	dareStr += "]"
	buf.WriteString(fmt.Sprintf(`[%s,%d]`, dareStr, this.ResetTime))
	return buf.Bytes(), nil
}

func (this *PersonBosses) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	dareInterface := datas[0].([]interface{})
	dareMap := make(map[int]int)
	for _, dare := range dareInterface {
		dareArr := dare.([]interface{})
		dareMap[int(dareArr[0].(float64))] = int(dareArr[1].(float64))
	}
	(*this).DareNum = dareMap
	(*this).ResetTime = int(datas[1].(float64))
	return nil
}

func (this *PersonBosses) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this PersonBosses) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this VipBosses) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	dareStr := "["
	if len(this.DareNum) > 0 {
		for stageId, num := range this.DareNum {
			dareStr += fmt.Sprintf(`[%d,%d],`, stageId, num)
		}
		dareStr = dareStr[:len(dareStr)-1]
	}
	dareStr += "]"
	buf.WriteString(fmt.Sprintf(`[%s,%d]`, dareStr, this.ResetTime))
	return buf.Bytes(), nil
}

func (this *VipBosses) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	dareInterface := datas[0].([]interface{})
	dareMap := make(map[int]int)
	for _, dare := range dareInterface {
		dareArr := dare.([]interface{})
		dareMap[int(dareArr[0].(float64))] = int(dareArr[1].(float64))
	}
	(*this).DareNum = dareMap
	(*this).ResetTime = int(datas[1].(float64))
	return nil
}

func (this *VipBosses) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this VipBosses) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this MaterialStage) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	materialStr := "["
	if len(this.MaterialStages) > 0 {
		for mateType, unit := range this.MaterialStages {
			materialStr += fmt.Sprintf(`[%d,%d,%d,%d,%d],`, mateType, unit.DareNum, unit.BuyNum, unit.NowLayer, unit.LastLayer)
		}
		materialStr = materialStr[:len(materialStr)-1]
	}
	materialStr += "]"
	buf.WriteString(fmt.Sprintf(`[%d,%s]`, this.ResetTime, materialStr))
	return buf.Bytes(), nil
}

func (this *MaterialStage) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	mateInterface := datas[1].([]interface{})
	mateMap := make(map[int]*MaterialStageUnit)
	for _, mate := range mateInterface {
		mateArr := mate.([]interface{})
		mateMap[int(mateArr[0].(float64))] = &MaterialStageUnit{
			DareNum:   int(mateArr[1].(float64)),
			BuyNum:    int(mateArr[2].(float64)),
			NowLayer:  int(mateArr[3].(float64)),
			LastLayer: int(mateArr[4].(float64)),
		}
	}
	(*this).ResetTime = int(datas[0].(float64))
	(*this).MaterialStages = mateMap
	return nil
}

func (this *MaterialStage) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this MaterialStage) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Talent) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	talentListStr := "["
	if len(this.TalentList) > 0 {
		for talentWayId, talentUnit := range this.TalentList {
			talentsStr := "["
			if len(talentUnit.Talents) > 0 {
				for id, lv := range talentUnit.Talents {
					talentsStr += fmt.Sprintf(`[%d,%d],`, id, lv)
				}
				talentsStr = talentsStr[:len(talentsStr)-1]
			}
			talentsStr += "]"
			talentListStr += fmt.Sprintf(`[%d,%d,%s],`, talentWayId, talentUnit.UsePoints, talentsStr)
		}
		talentListStr = talentListStr[:len(talentListStr)-1]
	}
	talentListStr += "]"
	buf.WriteString(fmt.Sprintf(`[%d,%d,%s]`, this.GetPoints, this.SurplusPoints, talentListStr))
	return buf.Bytes(), nil
}

func (this *Talent) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	talentListInterface := datas[2].([]interface{})
	talentListMap := make(map[int]*TalentUnit)
	for _, talentList := range talentListInterface {
		talentListArr := talentList.([]interface{})
		id := int(talentListArr[0].(float64))
		talentsInterface := talentListArr[2].([]interface{})
		talentsMap := make(map[int]int)
		for _, talents := range talentsInterface {
			talentsArr := talents.([]interface{})
			talentsMap[int(talentsArr[0].(float64))] = int(talentsArr[1].(float64))
		}
		talentListMap[id] = &TalentUnit{
			UsePoints: int(talentListArr[1].(float64)),
			Talents:   talentsMap,
		}
	}
	(*this).GetPoints = int(datas[0].(float64))
	(*this).SurplusPoints = int(datas[1].(float64))
	(*this).TalentList = talentListMap
	return nil
}

func (this *Talent) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Talent) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this *GuildData) Scan(value interface{}) error {
	data := value.([]byte)
	if len(data) > 0 {
		return json.Unmarshal(data, this)
	}
	return nil
}

func (this GuildData) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this PaoDian) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`[%d]`, this.EndTime))
	return buf.Bytes(), nil
}

func (this *PaoDian) UnMarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	(*this).EndTime = int(datas[0].(float64))
	return nil
}

func (this *PaoDian) Scan(value interface{}) error {
	return this.UnMarshalJSON(value.([]byte))
}

func (this PaoDian) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Friend) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	friendStr := "["
	if len(this) > 0 {
		for friendId, friendUnit := range this {
			msgStr := "["
			if len(friendUnit.MsgLog) > 0 {
				for _, log := range friendUnit.MsgLog {
					isMy := 0
					if log.IsMy {
						isMy = 1
					}
					msgStr += fmt.Sprintf(`["%s",%d,%d],`, log.Msg, log.Time, isMy)
				}
				msgStr = msgStr[:len(msgStr)-1]
			}
			msgStr += "]"
			isRead := 0
			if friendUnit.IsRead {
				isRead = 1
			}
			friendStr += fmt.Sprintf(`[%d,%s,%d,%d,%d,%d],`, friendId, msgStr, friendUnit.BlockTime, friendUnit.CreatedAt, friendUnit.DeletedAt, isRead)
		}
		friendStr = friendStr[:len(friendStr)-1]
	}
	friendStr += "]"
	buf.WriteString(fmt.Sprintf(`%s`, friendStr))
	return buf.Bytes(), nil
}

func (this *Friend) UnmarshalJSON(data []byte) error {
	*this = make(map[int]*FriendUnit)
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, val := range datas {
		valArr := val.([]interface{})
		msgLogInterface := valArr[1].([]interface{})
		msgMap := make(MsgLogs)
		for _, msgLog := range msgLogInterface {
			msgLogArr := msgLog.([]interface{})
			time := int(msgLogArr[1].(float64))
			isMy := false
			if int(msgLogArr[2].(float64)) == 1 {
				isMy = true
			}
			msgMap[time] = &MsgLog{
				Msg:  msgLogArr[0].(string),
				Time: time,
				IsMy: isMy,
			}
		}
		isRead := false
		if len(valArr) > 5 && int(valArr[5].(float64)) == 1 {
			isRead = true
		}
		(*this)[int(valArr[0].(float64))] = &FriendUnit{
			MsgLog:    msgMap,
			BlockTime: int(valArr[2].(float64)),
			CreatedAt: int(valArr[3].(float64)),
			DeletedAt: int(valArr[4].(float64)),
			IsRead:    isRead,
		}
	}
	return nil
}

func (this *Friend) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Friend) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Fit) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	fashionStr := "["
	if len(this.Fashion) > 0 {
		for id, lv := range this.Fashion {
			fashionStr += fmt.Sprintf(`[%d,%d],`, id, lv)
		}
		fashionStr = fashionStr[:len(fashionStr)-1]
	}
	fashionStr += "]"
	skillBagStr := "["
	if len(this.SkillBag) > 0 {
		for pos, id := range this.SkillBag {
			skillBagStr += fmt.Sprintf(`[%d,%d],`, pos, id)
		}
		skillBagStr = skillBagStr[:len(skillBagStr)-1]
	}
	skillBagStr += "]"
	lvStr := "["
	if len(this.Lv) > 0 {
		for id, lv := range this.Lv {
			lvStr += fmt.Sprintf(`[%d,%d],`, id, lv)
		}
		lvStr = lvStr[:len(lvStr)-1]
	}
	lvStr += "]"
	skillsStr := "["
	if len(this.Skills) > 0 {
		for id, skill := range this.Skills {
			skillsStr += fmt.Sprintf(`[%d,%d,%d],`, id, skill.Lv, skill.Star)
		}
		skillsStr = skillsStr[:len(skillsStr)-1]
	}
	skillsStr += "]"
	buf.WriteString(fmt.Sprintf(`[%d,%d,%s,%s,%s,%s]`, this.CdStart, this.CdEnd, fashionStr, skillBagStr, lvStr, skillsStr))
	return buf.Bytes(), nil
}

func (this *Fit) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	fashionInterface := datas[2].([]interface{})
	fashionMap := make(IntKv)
	for _, fashion := range fashionInterface {
		fashionArr := fashion.([]interface{})
		fashionMap[int(fashionArr[0].(float64))] = int(fashionArr[1].(float64))
	}
	skillBagInterface := datas[3].([]interface{})
	skillBagMap := make(IntKv)
	for _, skillBag := range skillBagInterface {
		skillBagArr := skillBag.([]interface{})
		skillBagMap[int(skillBagArr[0].(float64))] = int(skillBagArr[1].(float64))
	}
	lvInterface := datas[4].([]interface{})
	lvMap := make(IntKv)
	for _, lv := range lvInterface {
		lvArr := lv.([]interface{})
		lvMap[int(lvArr[0].(float64))] = int(lvArr[1].(float64))
	}
	skillsInterface := datas[5].([]interface{})
	skillsMap := make(map[int]*FitSkill)
	for _, skills := range skillsInterface {
		skillsArr := skills.([]interface{})
		skillsMap[int(skillsArr[0].(float64))] = &FitSkill{
			Lv:   int(skillsArr[1].(float64)),
			Star: int(skillsArr[2].(float64)),
		}
	}
	(*this).CdStart = int(datas[0].(float64))
	(*this).CdEnd = int(datas[1].(float64))
	(*this).Fashion = fashionMap
	(*this).SkillBag = skillBagMap
	(*this).Lv = lvMap
	(*this).Skills = skillsMap
	return nil
}

func (this *Fit) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Fit) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this DailyTaskInfo) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	activityStr := "["
	if len(this.DailyTask) > 0 {
		for id, activity := range this.DailyTask {
			activityStr += fmt.Sprintf(`[%d,%d,%d,%d,%d],`, id, activity.ActivityId, activity.IsCanGetExp, activity.HaveChallengeTimes, activity.BuyChallengeTimes)
		}
		activityStr = activityStr[:len(activityStr)-1]
	}
	activityStr += "]"

	idsStr := "{}"
	if len(this.ResourcesHaveBackTimes) > 0 {
		infoBytes, _ := this.ResourcesHaveBackTimes.MarshalJSON()
		idsStr = string(infoBytes)
	}

	resourceCanBackTimes := "{}"
	if len(this.ResourceCanBackTimes) > 0 {
		infoBytes, _ := this.ResourceCanBackTimes.MarshalJSON()
		resourceCanBackTimes = string(infoBytes)
	}

	buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%s,[%s],[%s],%d,%s,%s]`, this.DayExp, this.WeekExp, this.ResourcesBackExp, activityStr, common.JoinIntSlice(this.GetDayRewardIds, ","), common.JoinIntSlice(this.GetWeekRewardIds, ","), this.ResetTime, idsStr, resourceCanBackTimes))
	return buf.Bytes(), nil
}

func (this *DailyTaskInfo) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}

	skillsInterface := datas[3].([]interface{})
	skillsMap := make(map[int]*DailyTaskActivityInfo)
	for _, skills := range skillsInterface {
		skillsArr := skills.([]interface{})
		skillsMap[int(skillsArr[0].(float64))] = &DailyTaskActivityInfo{
			ActivityId:         int(skillsArr[1].(float64)),
			IsCanGetExp:        int(skillsArr[2].(float64)),
			HaveChallengeTimes: int(skillsArr[3].(float64)),
			BuyChallengeTimes:  int(skillsArr[4].(float64)),
		}
	}
	(*this).DayExp = int(datas[0].(float64))
	(*this).WeekExp = int(datas[1].(float64))
	(*this).ResourcesBackExp = int(datas[2].(float64))
	(*this).DailyTask = skillsMap
	awardIds := datas[4].([]interface{})
	for _, v := range awardIds {
		(*this).GetDayRewardIds = append((*this).GetDayRewardIds, int(v.(float64)))
	}
	awardIds1 := datas[5].([]interface{})
	for _, v := range awardIds1 {
		(*this).GetWeekRewardIds = append((*this).GetWeekRewardIds, int(v.(float64)))
	}
	(*this).ResetTime = int(datas[6].(float64))
	if len(datas) > 7 {
		(*this).ResourcesHaveBackTimes = make(IntKv)
		infoMap := datas[7].(map[string]interface{})
		if len(infoMap) > 0 {
			for k, v := range infoMap {
				kk, _ := strconv.Atoi(k)
				vv := int(v.(float64))
				(*this).ResourcesHaveBackTimes[kk] = vv
			}
		}
	}
	if len(datas) > 8 {
		(*this).ResourceCanBackTimes = make(StringIntKv)
		infoMap := datas[8].(map[string]interface{})
		if len(infoMap) > 0 {
			for k, v := range infoMap {
				vv := int(v.(float64))
				(*this).ResourceCanBackTimes[k] = vv
			}
		}
	}
	return nil
}

func (this *DailyTaskInfo) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this DailyTaskInfo) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this MonthCard) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	monthCardStr := "["
	if len(this.MonthCards) > 0 {
		for id, unit := range this.MonthCards {
			monthCardStr += fmt.Sprintf(`[%d,%d,%d],`, id, unit.StartTime, unit.EndTime)
		}
		monthCardStr = monthCardStr[:len(monthCardStr)-1]
	}
	monthCardStr += "]"
	buf.WriteString(fmt.Sprintf(`[%d,%s]`, this.ResetTime, monthCardStr))
	return buf.Bytes(), nil
}

func (this *MonthCard) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	monthCardInterface := datas[1].([]interface{})
	monthCardMap := make(map[int]*MonthCardUnit)
	for _, monthCard := range monthCardInterface {
		monthCardArr := monthCard.([]interface{})
		monthCardMap[int(monthCardArr[0].(float64))] = &MonthCardUnit{
			StartTime: int(monthCardArr[1].(float64)),
			EndTime:   int(monthCardArr[2].(float64)),
		}
	}
	(*this).ResetTime = int(datas[0].(float64))
	(*this).MonthCards = monthCardMap
	return nil
}

func (this *MonthCard) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this MonthCard) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this FirstRecharge) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	dayStr := "["
	if len(this.Days) > 0 {
		for day := range this.Days {
			dayStr += fmt.Sprintf(`%d,`, day)
		}
		dayStr = dayStr[:len(dayStr)-1]
	}
	dayStr += "]"
	isRecharge := 0
	if this.IsRecharge {
		isRecharge = 1
	}
	buf.WriteString(fmt.Sprintf(`[%d,%s,%d,%d]`, isRecharge, dayStr, this.OpenDay, this.Discount))
	return buf.Bytes(), nil
}

func (this *FirstRecharge) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	dayInterface := datas[1].([]interface{})
	dayMap := make(IntKv)
	for _, day := range dayInterface {
		dayMap[int(day.(float64))] = 0
	}
	isRecharge := false
	if int(datas[0].(float64)) == 1 {
		isRecharge = true
	}
	(*this).IsRecharge = isRecharge
	(*this).Days = dayMap
	if len(datas) > 2 {
		(*this).OpenDay = int(datas[2].(float64))
	}
	if len(datas) > 3 {
		(*this).Discount = int(datas[3].(float64))
	}
	return nil
}

func (this *FirstRecharge) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this FirstRecharge) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this SpendRebates) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	rewardsStr := "["
	if len(this.Reward) > 0 {
		for id := range this.Reward {
			rewardsStr += fmt.Sprintf(`%d,`, id)
		}
		rewardsStr = rewardsStr[:len(rewardsStr)-1]
	}
	rewardsStr += "]"
	buf.WriteString(fmt.Sprintf(`[%d,%d,%s,%d]`, this.CountIngot, this.Ingot, rewardsStr, this.Cycle))
	return buf.Bytes(), nil
}

func (this *SpendRebates) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	rewardInterface := datas[2].([]interface{})
	rewardMap := make(IntKv)
	for _, reward := range rewardInterface {
		rewardMap[int(reward.(float64))] = 0
	}
	(*this).CountIngot = int(datas[0].(float64))
	(*this).Ingot = int(datas[1].(float64))
	(*this).Reward = rewardMap
	if len(datas) > 3 {
		(*this).Cycle = int(datas[3].(float64))
	}
	return nil
}

func (this *SpendRebates) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this SpendRebates) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Achievement) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	monthCardStr := "["
	if len(this.Task) > 0 {
		for id, unit := range this.Task {
			monthCardStr += fmt.Sprintf(`[%d,%d,%d,%d,%d],`, id, unit.NowTaskId, unit.NextTaskId, unit.Process, unit.IsGetAll)
		}
		monthCardStr = monthCardStr[:len(monthCardStr)-1]
	}
	monthCardStr += "]"

	buf.WriteString(fmt.Sprintf(`[%d,%s,[%s]]`, this.Point, monthCardStr, common.JoinIntSlice(this.Medal, ",")))
	return buf.Bytes(), nil
}

func (this *Achievement) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}

	skillsInterface := datas[1].([]interface{})
	skillsMap := make(map[int]*AchievementInfo)
	for _, skills := range skillsInterface {
		skillsArr := skills.([]interface{})
		skillsMap[int(skillsArr[0].(float64))] = &AchievementInfo{
			NowTaskId:  int(skillsArr[1].(float64)),
			NextTaskId: int(skillsArr[2].(float64)),
			Process:    int(skillsArr[3].(float64)),
			IsGetAll:   int(skillsArr[4].(float64)),
		}
	}

	(*this).Point = int(datas[0].(float64))
	(*this).Task = skillsMap
	awardIds := datas[2].([]interface{})
	for _, v := range awardIds {
		(*this).Medal = append((*this).Medal, int(v.(float64)))
	}

	return nil
}

func (this *Achievement) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Achievement) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this LimitGift) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	gradeStatusStr := "["
	if len(this.GradeStatus) > 0 {
		for t, g := range this.GradeStatus {
			gradeStatusStr += fmt.Sprintf(`[%d,%d],`, t, g)
		}
		gradeStatusStr = gradeStatusStr[:len(gradeStatusStr)-1]
	}
	gradeStatusStr += "]"
	isBuyStr := "["
	if len(this.IsBuy) > 0 {
		for t, f := range this.IsBuy {
			flag := 0
			if f {
				flag = 1
			}
			isBuyStr += fmt.Sprintf(`[%d,%d],`, t, flag)
		}
		isBuyStr = isBuyStr[:len(isBuyStr)-1]
	}
	isBuyStr += "]"
	tlvStr := "["
	if len(this.TLv) > 0 {
		for t, lv := range this.TLv {
			tlvStr += fmt.Sprintf(`[%d,%d],`, t, lv)
		}
		tlvStr = tlvStr[:len(tlvStr)-1]
	}
	tlvStr += "]"
	listStr := "["
	if len(this.List) > 0 {
		for t, lvUnits := range this.List {
			for lv, unit := range lvUnits {
				isBuy := 0
				if unit.IsBuy {
					isBuy = 1
				}
				listStr += fmt.Sprintf(`[%d,%d,%d,%d,%d,%d],`, t, lv, unit.Grade, unit.StartTime, unit.EndTime, isBuy)
			}
		}
		if len(listStr) != 1 {
			listStr = listStr[:len(listStr)-1]
		}
	}
	listStr += "]"
	mergeStr := "["
	if len(this.MergeData) > 0 {
		for t, grade := range this.MergeData {
			mergeStr += fmt.Sprintf("[%d,%d],", t, grade)
		}
		mergeStr = mergeStr[:len(mergeStr)-1]
	}
	mergeStr += "]"
	buf.WriteString(fmt.Sprintf(`[%s,%s,%s,%s,%s]`, gradeStatusStr, isBuyStr, tlvStr, listStr, mergeStr))
	return buf.Bytes(), nil
}

func (this *LimitGift) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}

	gradeStatusInterface := datas[0].([]interface{})
	gradeStatusMap := make(IntKv)
	for _, gradeStatus := range gradeStatusInterface {
		gradeStatusArr := gradeStatus.([]interface{})
		gradeStatusMap[int(gradeStatusArr[0].(float64))] = int(gradeStatusArr[1].(float64))
	}
	isBuyInterface := datas[1].([]interface{})
	isBuyMap := make(map[int]bool)
	for _, isBuy := range isBuyInterface {
		isBuyArr := isBuy.([]interface{})
		f := false
		if int(isBuyArr[1].(float64)) == 1 {
			f = true
		}
		isBuyMap[int(isBuyArr[0].(float64))] = f
	}
	tlvInterface := datas[2].([]interface{})
	tlvMap := make(IntKv)
	for _, tlv := range tlvInterface {
		tlvArr := tlv.([]interface{})
		tlvMap[int(tlvArr[0].(float64))] = int(tlvArr[1].(float64))
	}
	listInterface := datas[3].([]interface{})
	listMap := make(map[int]map[int]*LimitGiftUnit)
	for _, list := range listInterface {
		listArr := list.([]interface{})
		t := int(listArr[0].(float64))
		if listMap[t] == nil {
			listMap[t] = make(map[int]*LimitGiftUnit)
		}
		lv := int(listArr[1].(float64))
		isBuy := false
		if int(listArr[5].(float64)) == 1 {
			isBuy = true
		}
		listMap[t][lv] = &LimitGiftUnit{
			Lv:        lv,
			Grade:     int(listArr[2].(float64)),
			StartTime: int(listArr[3].(float64)),
			EndTime:   int(listArr[4].(float64)),
			IsBuy:     isBuy,
		}
	}
	mergeMap := make(IntKv)
	if len(datas) > 4 {
		mergeInterface := datas[4].([]interface{})
		for _, merge := range mergeInterface {
			mergeArr := merge.([]interface{})
			mergeMap[int(mergeArr[0].(float64))] = int(mergeArr[1].(float64))
		}
	}
	(*this).GradeStatus = gradeStatusMap
	(*this).IsBuy = isBuyMap
	(*this).TLv = tlvMap
	(*this).List = listMap
	(*this).MergeData = mergeMap
	return nil
}

func (this *LimitGift) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this LimitGift) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this DailyPack) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for t, dailyPack := range this {
		buyIdsStr := "["
		if len(dailyPack.BuyIds) > 0 {
			for id, num := range dailyPack.BuyIds {
				buyIdsStr += fmt.Sprintf("[%d,%d],", id, num)
			}
			buyIdsStr = buyIdsStr[:len(buyIdsStr)-1]
		}
		buyIdsStr += "]"
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%s,%d,%d]`, t, buyIdsStr, dailyPack.ResetTime, dailyPack.ResetWeek))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%s,%d,%d]`, t, buyIdsStr, dailyPack.ResetTime, dailyPack.ResetWeek))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *DailyPack) UnmarshalJSON(data []byte) error {
	(*this) = make(map[int]*DailyPackUnit)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, val := range datas {
		buyIdInterface := val[1].([]interface{})
		buyIdMap := make(IntKv)
		for _, buyId := range buyIdInterface {
			buyIdArr := buyId.([]interface{})
			buyIdMap[int(buyIdArr[0].(float64))] = int(buyIdArr[1].(float64))
		}
		(*this)[int(val[0].(float64))] = &DailyPackUnit{
			BuyIds:    buyIdMap,
			ResetTime: int(val[2].(float64)),
		}
		if len(val) > 3 {
			(*this)[int(val[0].(float64))].ResetWeek = int(val[3].(float64))
		}
	}
	return nil
}

func (this *DailyPack) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this DailyPack) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this GrowFund) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	idsStr := "["
	if len(this.Ids) > 0 {
		for id := range this.Ids {
			idsStr += fmt.Sprintf(`%d,`, id)
		}
		idsStr = idsStr[:len(idsStr)-1]
	}
	idsStr += "]"
	isBuy := 0
	if this.IsBuy {
		isBuy = 1
	}
	buf.WriteString(fmt.Sprintf(`[%d,%s]`, isBuy, idsStr))
	return buf.Bytes(), nil
}

func (this *GrowFund) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	idsInterface := datas[1].([]interface{})
	idsMap := make(IntKv)
	for _, id := range idsInterface {
		idsMap[int(id.(float64))] = 0
	}
	isBuy := false
	if int(datas[0].(float64)) == 1 {
		isBuy = true
	}
	(*this).IsBuy = isBuy
	(*this).Ids = idsMap
	return nil
}

func (this *GrowFund) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this GrowFund) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this RedPacketItem) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	info := "{}"
	if len(this.PickInfo) > 0 {
		infoBytes, _ := this.PickInfo.MarshalJSON()
		info = string(infoBytes)
	}

	buf.WriteString(fmt.Sprintf(`[%d,%d,%s]`, this.Day, this.PickNum, info))
	return buf.Bytes(), nil
}

func (this *RedPacketItem) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}

	(*this).Day = int(datas[0].(float64))
	(*this).PickNum = int(datas[1].(float64))
	(*this).PickInfo = make(IntKv)
	infoMap := datas[2].(map[string]interface{})
	if len(infoMap) > 0 {
		for k, v := range infoMap {
			kk, _ := strconv.Atoi(k)
			vv := int(v.(float64))
			(*this).PickInfo[kk] = vv
		}
	}
	return nil
}

func (this *RedPacketItem) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this RedPacketItem) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this WarOrder) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	exchangeStr := "["
	if len(this.Exchange) > 0 {
		for id, num := range this.Exchange {
			exchangeStr += fmt.Sprintf(`[%d,%d],`, id, num)
		}
		exchangeStr = exchangeStr[:len(exchangeStr)-1]
	}
	exchangeStr += "]"
	taskStr := "["
	if len(this.Task) > 0 {
		for id, task := range this.Task {
			valTwoStr := "["
			if len(task.Val.Two) > 0 {
				for k, v := range task.Val.Two {
					valTwoStr += fmt.Sprintf(`[%d,%d],`, k, v)
				}
				valTwoStr = valTwoStr[:len(valTwoStr)-1]
			}
			valTwoStr += "]"
			valThreeStr := "["
			if len(task.Val.Three) > 0 {
				for k, v := range task.Val.Three {
					valThreeStr += fmt.Sprintf(`[%d,%d],`, k, v)
				}
				valThreeStr = valThreeStr[:len(valThreeStr)-1]
			}
			valThreeStr += "]"
			valStr := fmt.Sprintf(`[%d,%s,%s]`, task.Val.One, valTwoStr, valThreeStr)
			finish := 0
			if task.Finish {
				finish = 1
			}
			reward := 0
			if task.Reward {
				reward = 1
			}
			dateStr := "["
			if len(task.Date) > 0 {
				for k, v := range task.Date {
					dateStr += fmt.Sprintf(`[%d,%d],`, k, v)
				}
				dateStr = dateStr[:len(dateStr)-1]
			}
			dateStr += "]"
			taskStr += fmt.Sprintf(`[%d,%s,%d,%d,%s],`, id, valStr, finish, reward, dateStr)
		}
		taskStr = taskStr[:len(taskStr)-1]
	}
	taskStr += "]"
	weekTaskStr := "["
	if len(this.WeekTask) > 0 {
		for week, weekTask := range this.WeekTask {
			for id, task := range weekTask {
				valTwoStr := "["
				if len(task.Val.Two) > 0 {
					for k, v := range task.Val.Two {
						valTwoStr += fmt.Sprintf(`[%d,%d],`, k, v)
					}
					valTwoStr = valTwoStr[:len(valTwoStr)-1]
				}
				valTwoStr += "]"
				valThreeStr := "["
				if len(task.Val.Three) > 0 {
					for k, v := range task.Val.Three {
						valThreeStr += fmt.Sprintf(`[%d,%d],`, k, v)
					}
					valThreeStr = valThreeStr[:len(valThreeStr)-1]
				}
				valThreeStr += "]"
				valStr := fmt.Sprintf(`[%d,%s,%s]`, task.Val.One, valTwoStr, valThreeStr)
				finish := 0
				if task.Finish {
					finish = 1
				}
				reward := 0
				if task.Reward {
					reward = 1
				}
				dateStr := "["
				if len(task.Date) > 0 {
					for k, v := range task.Date {
						dateStr += fmt.Sprintf(`[%d,%d],`, k, v)
					}
					dateStr = dateStr[:len(dateStr)-1]
				}
				dateStr += "]"
				weekTaskStr += fmt.Sprintf(`[%d,%d,%s,%d,%d,%s],`, week, id, valStr, finish, reward, dateStr)
			}
		}
		weekTaskStr = weekTaskStr[:len(weekTaskStr)-1]
	}
	weekTaskStr += "]"
	rewardStr := "["
	if len(this.Reward) > 0 {
		for lv, reward := range this.Reward {
			elite := 0
			if reward.Elite {
				elite = 1
			}
			luxury := 0
			if reward.Luxury {
				luxury = 1
			}
			rewardStr += fmt.Sprintf(`[%d,%d,%d],`, lv, elite, luxury)
		}
		rewardStr = rewardStr[:len(rewardStr)-1]
	}
	rewardStr += "]"
	isLuxury := 0
	if this.IsLuxury {
		isLuxury = 1
	}
	buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%d,%d,%d,%s,%s,%s,%s]`,
		this.Lv, this.Exp, this.Season, this.StartTime, this.EndTIme, isLuxury, exchangeStr, taskStr, weekTaskStr, rewardStr))
	return buf.Bytes(), nil
}

func (this *WarOrder) UnmarshalJSON(data []byte) error {
	(*this).Exchange = make(IntKv)
	(*this).Task = make(map[int]*WarOrderTask)
	(*this).WeekTask = make(map[int]map[int]*WarOrderTask)
	(*this).Reward = make(map[int]*WarOrderReward)
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	exchangeInterface := datas[6].([]interface{})
	exchangeMap := make(IntKv)
	for _, exchange := range exchangeInterface {
		exchangeArr := exchange.([]interface{})
		exchangeMap[int(exchangeArr[0].(float64))] = int(exchangeArr[1].(float64))
	}
	taskInterface := datas[7].([]interface{})
	taskMap := make(map[int]*WarOrderTask)
	for _, task := range taskInterface {
		taskArr := task.([]interface{})
		taskValInterface := taskArr[1].([]interface{})
		taskValTwoInterface := taskValInterface[1].([]interface{})
		taskValTwoMap := make(IntKv)
		for _, taskValTwo := range taskValTwoInterface {
			taskValTwoArr := taskValTwo.([]interface{})
			taskValTwoMap[int(taskValTwoArr[0].(float64))] = int(taskValTwoArr[1].(float64))
		}
		taskValThreeInterface := taskValInterface[2].([]interface{})
		taskValThreeMap := make(IntKv)
		for _, taskValThree := range taskValThreeInterface {
			taskValThreeArr := taskValThree.([]interface{})
			taskValThreeMap[int(taskValThreeArr[0].(float64))] = int(taskValThreeArr[1].(float64))
		}
		finish := false
		if int(taskArr[2].(float64)) == 1 {
			finish = true
		}
		reward := false
		if int(taskArr[3].(float64)) == 1 {
			reward = true
		}
		taskDateInterface := taskArr[4].([]interface{})
		taskDateMap := make(IntKv)
		for _, taskDate := range taskDateInterface {
			taskDateArr := taskDate.([]interface{})
			taskDateMap[int(taskDateArr[0].(float64))] = int(taskDateArr[1].(float64))
		}
		taskMap[int(taskArr[0].(float64))] = &WarOrderTask{
			Val:    WarOrderTaskUnit{One: int(taskValInterface[0].(float64)), Two: taskValTwoMap, Three: taskValThreeMap},
			Finish: finish,
			Reward: reward,
			Date:   taskDateMap,
		}
	}
	weekTaskInterface := datas[8].([]interface{})
	weekTaskMap := make(map[int]map[int]*WarOrderTask)
	for _, weekTask := range weekTaskInterface {
		weekTaskArr := weekTask.([]interface{})
		week := int(weekTaskArr[0].(float64))
		id := int(weekTaskArr[1].(float64))
		if weekTaskMap[week] == nil {
			weekTaskMap[week] = make(map[int]*WarOrderTask)
		}
		taskValInterface := weekTaskArr[2].([]interface{})
		taskValTwoInterface := taskValInterface[1].([]interface{})
		taskValTwoMap := make(IntKv)
		for _, taskValTwo := range taskValTwoInterface {
			taskValTwoArr := taskValTwo.([]interface{})
			taskValTwoMap[int(taskValTwoArr[0].(float64))] = int(taskValTwoArr[1].(float64))
		}
		taskValThreeInterface := taskValInterface[2].([]interface{})
		taskValThreeMap := make(IntKv)
		for _, taskValThree := range taskValThreeInterface {
			taskValThreeArr := taskValThree.([]interface{})
			taskValThreeMap[int(taskValThreeArr[0].(float64))] = int(taskValThreeArr[1].(float64))
		}
		finish := false
		if int(weekTaskArr[3].(float64)) == 1 {
			finish = true
		}
		reward := false
		if int(weekTaskArr[4].(float64)) == 1 {
			reward = true
		}
		taskDateInterface := weekTaskArr[5].([]interface{})
		taskDateMap := make(IntKv)
		for _, taskDate := range taskDateInterface {
			taskDateArr := taskDate.([]interface{})
			taskDateMap[int(taskDateArr[0].(float64))] = int(taskDateArr[1].(float64))
		}
		weekTaskMap[week][id] = &WarOrderTask{
			Val:    WarOrderTaskUnit{One: int(taskValInterface[0].(float64)), Two: taskValTwoMap, Three: taskValThreeMap},
			Finish: finish,
			Reward: reward,
			Date:   taskDateMap,
		}
	}
	rewardInterface := datas[9].([]interface{})
	rewardMap := make(map[int]*WarOrderReward)
	for _, reward := range rewardInterface {
		rewardArr := reward.([]interface{})
		elite := false
		if int(rewardArr[1].(float64)) == 1 {
			elite = true
		}
		luxury := false
		if int(rewardArr[2].(float64)) == 1 {
			luxury = true
		}
		rewardMap[int(rewardArr[0].(float64))] = &WarOrderReward{
			Elite:  elite,
			Luxury: luxury,
		}
	}
	isLuxury := false
	if int(datas[5].(float64)) == 1 {
		isLuxury = true
	}
	(*this).Lv = int(datas[0].(float64))
	(*this).Exp = int(datas[1].(float64))
	(*this).Season = int(datas[2].(float64))
	(*this).StartTime = int(datas[3].(float64))
	(*this).EndTIme = int(datas[4].(float64))
	(*this).IsLuxury = isLuxury
	(*this).Exchange = exchangeMap
	(*this).Task = taskMap
	(*this).WeekTask = weekTaskMap
	(*this).Reward = rewardMap
	return nil
}

func (this *WarOrder) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this WarOrder) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Elf) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	skillStr := "["
	if len(this.Skills) > 0 {
		for id, lv := range this.Skills {
			skillStr += fmt.Sprintf(`[%d,%d],`, id, lv)
		}
		skillStr = skillStr[:len(skillStr)-1]
	}
	skillStr += "]"
	skillBagStr := "["
	if len(this.SkillBag) > 0 {
		for pos, id := range this.SkillBag {
			skillBagStr += fmt.Sprintf(`[%d,%d],`, pos, id)
		}
		skillBagStr = skillBagStr[:len(skillBagStr)-1]
	}
	skillBagStr += "]"
	receiveStr := "["
	if len(this.RecoverLimit) > 0 {
		for itemId, num := range this.RecoverLimit {
			receiveStr += fmt.Sprintf(`[%d,%d],`, itemId, num)
		}
		receiveStr = receiveStr[:len(receiveStr)-1]
	}
	receiveStr += "]"
	buf.WriteString(fmt.Sprintf(`[%d,%d,%s,%s,%s]`, this.Lv, this.Exp, skillStr, skillBagStr, receiveStr))
	return buf.Bytes(), nil
}

func (this *Elf) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	skillInterface := datas[2].([]interface{})
	skillMap := make(IntKv)
	for _, skill := range skillInterface {
		skillArr := skill.([]interface{})
		skillMap[int(skillArr[0].(float64))] = int(skillArr[1].(float64))
	}
	skillBagInterface := datas[3].([]interface{})
	skillBagMap := make(IntKv)
	for _, skillBag := range skillBagInterface {
		skillBagArr := skillBag.([]interface{})
		skillBagMap[int(skillBagArr[0].(float64))] = int(skillBagArr[1].(float64))
	}
	receiveMap := make(IntKv)
	if len(datas) > 4 {
		receiveInterface := datas[4].([]interface{})
		for _, receive := range receiveInterface {
			receiveArr := receive.([]interface{})
			receiveMap[int(receiveArr[0].(float64))] = int(receiveArr[1].(float64))
		}
	}
	(*this).Lv = int(datas[0].(float64))
	(*this).Exp = int(datas[1].(float64))
	(*this).Skills = skillMap
	(*this).SkillBag = skillBagMap
	(*this).RecoverLimit = receiveMap
	return nil
}

func (this *Elf) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Elf) Value() (driver.Value, error) {
	return json.Marshal(this)
}

//抽卡玩家数据
func (this CardInfo) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	getAwardIds := ""
	if len(this.GetAwardIds) > 0 {
		getAwardIds = common.JoinIntSlice(this.GetAwardIds, ";")
	}
	buf.WriteString(fmt.Sprintf(`["%v","%v","%v","%v","%v","%v","%v"]`, this.AddWeight, this.DrawTimes, this.Season, this.Integral, getAwardIds, this.DayResDay, this.MergeMark))
	return buf.Bytes(), nil
}

func (this *CardInfo) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		var wo []string
		err := json.Unmarshal(data, &wo)
		if err != nil {
			fmt.Println("err:", err)
		}
		this.AddWeight, _ = strconv.Atoi(wo[0])
		this.DrawTimes, _ = strconv.Atoi(wo[1])

		this.Season, _ = strconv.Atoi(wo[2])
		this.Integral, _ = strconv.Atoi(wo[3])
		var getAwardIds []int
		getAwardIds, _ = common.IntSliceFromString(wo[4], ";")
		this.GetAwardIds = getAwardIds
		if len(wo) > 5 {
			this.DayResDay, _ = strconv.Atoi(wo[5])
		}
		if len(wo) > 6 {
			this.MergeMark, _ = strconv.Atoi(wo[6])
		}
	}
	return nil
}

func (this *CardInfo) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this CardInfo) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this TreasureInfo) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	buyStr := "["
	if len(this.BuyTimes) > 0 {
		for pos, id := range this.BuyTimes {
			buyStr += fmt.Sprintf(`[%d,%d],`, pos, id)
		}
		buyStr = buyStr[:len(buyStr)-1]
	}
	buyStr += "]"

	chooseItems := "["
	if len(this.ChooseItems) > 0 {
		if len(this.ChooseItems) == 0 {
		}
		for k, v := range this.ChooseItems {
			if len(v) > 0 {
				chooseItems += fmt.Sprintf(`[%v,%v],`, k, common.JoinIntSlice(v, ","))
			} else {
				chooseItems += fmt.Sprintf(`[%v],`, k)
			}
		}
		chooseItems = chooseItems[:len(chooseItems)-1]
	}
	chooseItems += "]"

	haveRandomItems := "["
	if len(this.HaveRandomItems) > 0 {
		for k, v := range this.HaveRandomItems {
			if len(v) > 0 {
				haveRandomItems += fmt.Sprintf(`[%v,%v],`, k, common.JoinIntSlice(v, ","))
			} else {
				haveRandomItems += fmt.Sprintf(`[%v],`, k)
			}
		}
		haveRandomItems = haveRandomItems[:len(haveRandomItems)-1]
	}
	haveRandomItems += "]"
	buf.WriteString(fmt.Sprintf(`[%v,%v,%v,%v,[%v],%v,%v,%v,%v]`, this.Season, this.PopUpState, this.PopUpResOpenDay, this.AllUseTimes, common.JoinIntSlice(this.AllGetRound, ","), buyStr, chooseItems, haveRandomItems, this.MergeMark))
	return buf.Bytes(), nil
}

func (this *TreasureInfo) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	wo := make([]interface{}, 0)
	err := json.Unmarshal(data, &wo)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	(*this).Season = int(wo[0].(float64))
	(*this).PopUpState = int(wo[1].(float64))

	(*this).PopUpResOpenDay = int(wo[2].(float64))
	(*this).AllUseTimes = int(wo[3].(float64))
	awardIds := wo[4].([]interface{})
	for _, v := range awardIds {
		(*this).AllGetRound = append((*this).AllGetRound, int(v.(float64)))
	}

	if len(wo) >= 6 {
		buyInterface := wo[5].([]interface{})
		buyMap := make(IntKv)
		for _, data := range buyInterface {
			data1 := data.([]interface{})
			buyMap[int(data1[0].(float64))] = int(data1[1].(float64))
		}
		(*this).BuyTimes = buyMap
	}

	if len(wo) >= 7 {
		chooseItems := wo[6].([]interface{})
		chooseMap := make(map[int]IntSlice)
		for _, data := range chooseItems {
			data1 := data.([]interface{})
			for index, v := range data1 {
				if index == 0 {
					if chooseMap[int(data1[0].(float64))] == nil {
						chooseMap[int(data1[0].(float64))] = make(IntSlice, 0)
					}
				} else {
					chooseMap[int(data1[0].(float64))] = append(chooseMap[int(data1[0].(float64))], int(v.(float64)))
				}
			}
		}
		(*this).ChooseItems = chooseMap
	}

	if len(wo) >= 8 {
		haveGetItems := wo[7].([]interface{})
		haveGetMap := make(map[int]IntSlice)
		for _, data := range haveGetItems {
			data1 := data.([]interface{})
			for index, v := range data1 {
				if index == 0 {
					if haveGetMap[int(data1[0].(float64))] == nil {
						haveGetMap[int(data1[0].(float64))] = make(IntSlice, 0)
					}
				} else {
					haveGetMap[int(data1[0].(float64))] = append(haveGetMap[int(data1[0].(float64))], int(v.(float64)))
				}
			}
		}
		(*this).HaveRandomItems = haveGetMap
	}
	if len(wo) >= 9 {
		(*this).MergeMark = int(wo[8].(float64))
	}

	return nil
}

func (this *TreasureInfo) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this TreasureInfo) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this HolyBeastInfos) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for t, HolyBeastInfo := range this {
		buyIdsStr := "["
		if len(HolyBeastInfo.ChooseProp) > 0 {
			for star, index := range HolyBeastInfo.ChooseProp {
				buyIdsStr += fmt.Sprintf("[%d,%d],", star, index)
			}
			buyIdsStr = buyIdsStr[:len(buyIdsStr)-1]
		}
		buyIdsStr += "]"
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%s,%v,%v]`, t, buyIdsStr, HolyBeastInfo.Types, HolyBeastInfo.Star))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%s,%v,%v]`, t, buyIdsStr, HolyBeastInfo.Types, HolyBeastInfo.Star))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *HolyBeastInfos) UnmarshalJSON(data []byte) error {
	(*this) = make(map[int]*HolyBeastInfo)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, val := range datas {
		buyIdInterface := val[1].([]interface{})
		buyIdMap := make(IntKv)
		for _, buyId := range buyIdInterface {
			buyIdArr := buyId.([]interface{})
			buyIdMap[int(buyIdArr[0].(float64))] = int(buyIdArr[1].(float64))
		}
		(*this)[int(val[0].(float64))] = &HolyBeastInfo{
			ChooseProp: buyIdMap,
			Types:      int(val[2].(float64)),
			Star:       int(val[3].(float64)),
		}
	}
	return nil
}

func (this *HolyBeastInfos) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this HolyBeastInfos) Value() (driver.Value, error) {
	return json.Marshal(this)
}

//竞技场玩家数据
func (this *CompetitiveInfo) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this CompetitiveInfo) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this CompetitiveInfo) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(`["%v","%v","%v","%v","%v","%v"]`,
		this.HaveChallengeTimes, this.BuyTimes, this.DayResDay, this.BeforeDayRewardGetState, this.NowSeason, this.ContinuityWin))
	return buf.Bytes(), nil
}

func (this *CompetitiveInfo) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		var wo []string
		err := json.Unmarshal(data, &wo)
		if err != nil {
			fmt.Println("err:", err)
		}
		this.HaveChallengeTimes, _ = strconv.Atoi(wo[0])
		this.BuyTimes, _ = strconv.Atoi(wo[1])
		this.DayResDay, _ = strconv.Atoi(wo[2])
		this.BeforeDayRewardGetState, _ = strconv.Atoi(wo[3])
		this.NowSeason, _ = strconv.Atoi(wo[4])
		if len(wo) > 5 {
			this.ContinuityWin, _ = strconv.Atoi(wo[5])
		}
	}
	return nil
}

//竞技场玩家数据
func (this *FieldFight) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this FieldFight) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this FieldFight) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(`["%v","%v","%v"]`, this.HaveChallengeTimes, this.HaveBuyTimes, this.DayResDay))
	return buf.Bytes(), nil
}

func (this *FieldFight) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		var wo []string
		err := json.Unmarshal(data, &wo)
		if err != nil {
			fmt.Println("err:", err)
		}
		this.HaveChallengeTimes, _ = strconv.Atoi(wo[0])
		this.HaveBuyTimes, _ = strconv.Atoi(wo[1])
		this.DayResDay, _ = strconv.Atoi(wo[2])
	}
	return nil
}

func (this DailyRankInfos) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for t, dailyRankInfo := range this {
		buyIdsStr := "["
		if len(dailyRankInfo.GetDayRewardIds) > 0 {
			for star, index := range dailyRankInfo.GetDayRewardIds {
				buyIdsStr += fmt.Sprintf("[%d,%d],", star, index)
			}
			buyIdsStr = buyIdsStr[:len(buyIdsStr)-1]
		}
		buyIdsStr += "]"

		buyIdsStr1 := "["
		if len(dailyRankInfo.BuyRewardInfo) > 0 {
			for star, index := range dailyRankInfo.BuyRewardInfo {
				buyIdsStr1 += fmt.Sprintf("[%d,%d],", star, index)
			}
			buyIdsStr1 = buyIdsStr1[:len(buyIdsStr1)-1]
		}
		buyIdsStr1 += "]"

		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`[%d,%v,%v]`, t, buyIdsStr, buyIdsStr1))
		} else {
			buf.WriteString(fmt.Sprintf(`,[%d,%v,%v]`, t, buyIdsStr, buyIdsStr1))
		}
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func (this *DailyRankInfos) UnmarshalJSON(data []byte) error {
	(*this) = make(map[int]*DailyRankInfo)
	if len(data) == 0 {
		return nil
	}
	datas := make([][]interface{}, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for _, val := range datas {
		buyIdInterface := val[1].([]interface{})
		buyIdMap := make(IntKv)
		for _, buyId := range buyIdInterface {
			buyIdArr := buyId.([]interface{})
			buyIdMap[int(buyIdArr[0].(float64))] = int(buyIdArr[1].(float64))
		}

		buyIdInterface1 := val[2].([]interface{})
		buyIdMap1 := make(IntKv)
		for _, buyId := range buyIdInterface1 {
			buyIdArr := buyId.([]interface{})
			buyIdMap1[int(buyIdArr[0].(float64))] = int(buyIdArr[1].(float64))
		}

		(*this)[int(val[0].(float64))] = &DailyRankInfo{GetDayRewardIds: buyIdMap, BuyRewardInfo: buyIdMap1}
	}
	return nil
}

func (this *DailyRankInfos) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this DailyRankInfos) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this FitHolyEquip) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	equipStr := "["
	if len(this.Equips) > 0 {
		for t, kv := range this.Equips {
			for pos, id := range kv {
				equipStr += fmt.Sprintf(`[%d,%d,%d],`, t, pos, id)
			}
		}
		equipStr = equipStr[:len(equipStr)-1]
	}
	equipStr += "]"
	buf.WriteString(fmt.Sprintf(`[%d,%s]`, this.SuitId, equipStr))
	return buf.Bytes(), nil
}

func (this *FitHolyEquip) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	datas := make([]interface{}, 0)
	if err := json.Unmarshal(data, &datas); err != nil {
		return err
	}
	equipInterface := datas[1].([]interface{})
	equipMap := make(MapIntKv)
	for _, equip := range equipInterface {
		equipArr := equip.([]interface{})
		t := int(equipArr[0].(float64))
		if equipMap[t] == nil {
			equipMap[t] = make(IntKv)
		}
		equipMap[t][int(equipArr[1].(float64))] = int(equipArr[2].(float64))
	}
	(*this).SuitId = int(datas[0].(float64))
	(*this).Equips = equipMap
	return nil
}

func (this *FitHolyEquip) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this FitHolyEquip) Value() (driver.Value, error) {
	return json.Marshal(this)
}

//七日投资玩家数据
func (this SevenInvestmentInfo) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	getAwardIds := ""
	if len(this.GetAwardIds) > 0 {
		getAwardIds = common.JoinIntSlice(this.GetAwardIds, ";")
	}
	buf.WriteString(fmt.Sprintf(`["%v","%v"]`, this.BuyOpenDay, getAwardIds))
	return buf.Bytes(), nil
}

func (this *SevenInvestmentInfo) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		var wo []string
		err := json.Unmarshal(data, &wo)
		if err != nil {
			fmt.Println("err:", err)
		}
		this.BuyOpenDay, _ = strconv.Atoi(wo[0])
		var getAwardIds []int
		getAwardIds, _ = common.IntSliceFromString(wo[1], ";")
		this.GetAwardIds = getAwardIds

	}
	return nil
}

func (this *SevenInvestmentInfo) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this SevenInvestmentInfo) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this ContRecharge) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	dayStr := "["
	if len(this.Day) > 0 {
		for day, money := range this.Day {
			dayStr += fmt.Sprintf(`[%d,%d],`, day, money)
		}
		dayStr = dayStr[:len(dayStr)-1]
	}
	dayStr += "]"
	receiveStr := "["
	if len(this.Receive) > 0 {
		for id := range this.Receive {
			receiveStr += fmt.Sprintf(`%d,`, id)
		}
		receiveStr = receiveStr[:len(receiveStr)-1]
	}
	receiveStr += "]"
	buf.WriteString(fmt.Sprintf(`[%d,%s,%s]`, this.Cycle, dayStr, receiveStr))
	return buf.Bytes(), nil
}

func (this *ContRecharge) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		datas := make([]interface{}, 0)
		if err := json.Unmarshal(data, &datas); err != nil {
			return err
		}
		dayInterface := datas[1].([]interface{})
		dayMap := make(IntKv)
		for _, day := range dayInterface {
			dayArr := day.([]interface{})
			dayMap[int(dayArr[0].(float64))] = int(dayArr[1].(float64))
		}
		receiveInterface := datas[2].([]interface{})
		receiveMap := make(IntKv)
		for _, receive := range receiveInterface {
			receiveMap[int(receive.(float64))] = 0
		}
		(*this).Cycle = int(datas[0].(float64))
		(*this).Day = dayMap
		(*this).Receive = receiveMap
	}
	return nil
}

func (this *ContRecharge) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this ContRecharge) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this AncientBoss) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`[%d,%d,%d]`, this.DareNum, this.BuyNum, this.ResetTime))
	return buf.Bytes(), nil
}

func (this *AncientBoss) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		datas := make([]int, 0)
		if err := json.Unmarshal(data, &datas); err != nil {
			return err
		}
		(*this).DareNum = datas[0]
		(*this).BuyNum = datas[1]
		(*this).ResetTime = datas[2]
	}
	return nil
}

func (this *AncientBoss) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this AncientBoss) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this AncientSkill) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`[%d,%d,%d]`, this.SkillId, this.Level, this.Grade))
	return buf.Bytes(), nil
}

func (this *AncientSkill) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		datas := make([]int, 0)
		if err := json.Unmarshal(data, &datas); err != nil {
			return err
		}
		(*this).SkillId = datas[0]
		(*this).Level = datas[1]
		(*this).Grade = datas[2]
	}
	return nil
}

func (this *AncientSkill) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this AncientSkill) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Title) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	titleStr := ""
	if len(this) > 0 {
		for id, info := range this {
			isLook := 0
			if info.IsLook {
				isLook = 1
			}
			isExpire := 0
			if info.IsExpire {
				isExpire = 1
			}
			titleStr += fmt.Sprintf("[%d,%d,%d,%d,%d],", id, info.StartTime, info.EndTime, isLook, isExpire)
		}
		titleStr = titleStr[:len(titleStr)-1]
	}
	buf.WriteString(fmt.Sprintf(`[%s]`, titleStr))
	return buf.Bytes(), nil
}

func (this *Title) UnmarshalJSON(data []byte) error {
	(*this) = make(map[int]*TitleUnit)
	if len(data) > 0 {
		datas := make([][]int, 0)
		if err := json.Unmarshal(data, &datas); err != nil {
			return err
		}
		for _, info := range datas {
			isLook := false
			if info[3] == 1 {
				isLook = true
			}
			isExpire := false
			if len(info) > 4 && info[4] == 1 {
				isExpire = true
			}
			(*this)[info[0]] = &TitleUnit{
				StartTime: info[1],
				EndTime:   info[2],
				IsLook:    isLook,
				IsExpire:  isExpire,
			}
		}
	}
	return nil
}

func (this *Title) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Title) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this Condition) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k, val := range this {
		bagInfoUnitStr, _ := json.Marshal(val)
		if buf.Len() == 1 {
			buf.WriteString(fmt.Sprintf(`"%d":%s`, k, bagInfoUnitStr))
		} else {
			buf.WriteString(fmt.Sprintf(`,"%d":%s`, k, bagInfoUnitStr))
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (this *Condition) UnmarshalJSON(data []byte) error {
	*this = make(map[int][]int, 0)
	if len(data) == 0 {
		return nil
	}
	datas := make(map[int][]int, 0)
	err := json.Unmarshal(data, &datas)
	if err != nil {
		return err
	}
	for k, v := range datas {
		(*this)[k] = v
	}
	return nil
}

func (this *Condition) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this Condition) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this MiJi) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	titleStr := ""
	if len(this) > 0 {
		for id, info := range this {
			titleStr += fmt.Sprintf("[%d,%d,%d],", id, info.MiJiType, info.MiJiLv)
		}
		titleStr = titleStr[:len(titleStr)-1]
	}
	buf.WriteString(fmt.Sprintf(`[%s]`, titleStr))
	return buf.Bytes(), nil
}

func (this *MiJi) UnmarshalJSON(data []byte) error {
	(*this) = make(map[int]*MiJiUnit)
	if len(data) > 0 {
		datas := make([][]int, 0)
		if err := json.Unmarshal(data, &datas); err != nil {
			return err
		}
		for _, info := range datas {
			(*this)[info[0]] = &MiJiUnit{
				MiJiType: info[1],
				MiJiLv:   info[2],
			}
		}
	}
	return nil
}

func (this *MiJi) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this MiJi) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this KillMonster) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	uniStr := "["
	if len(this.Uni) > 0 {
		for stageId, uni := range this.Uni {
			uniStr += fmt.Sprintf(`[%d,%d,%d],`, stageId, marshalBool(uni.Draw), marshalBool(uni.FirstDraw))
		}
		uniStr = uniStr[:len(uniStr)-1]
	}
	uniStr += "]"
	milStr := "["
	if len(this.Mil) > 0 {
		for t, mil := range this.Mil {
			milStr += fmt.Sprintf(`[%d,%d,%d],`, t, mil.Level, marshalBool(mil.Draw))
		}
		milStr = milStr[:len(milStr)-1]
	}
	milStr += "]"
	buf.WriteString(fmt.Sprintf(`[%s,[[%d]],%s]`, uniStr, this.Per, milStr))
	return buf.Bytes(), nil
}
func (this *KillMonster) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		datas := make([][][]int, 0)
		if err := json.Unmarshal(data, &datas); err != nil {
			return err
		}
		uniMap := make(map[int]*KillMonsterUni)
		for _, info := range datas[0] {
			uniMap[info[0]] = &KillMonsterUni{
				Draw:      unmarshalBool(info[1]),
				FirstDraw: unmarshalBool(info[2]),
			}
		}
		milMap := make(map[int]*KillMonsterMil)
		for _, info := range datas[2] {
			milMap[info[0]] = &KillMonsterMil{
				Level: info[1],
				Draw:  unmarshalBool(info[2]),
			}
		}
		(*this).Uni = uniMap
		(*this).Per = datas[1][0][0]
		(*this).Mil = milMap
	}
	return nil
}
func (this *KillMonster) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}
func (this KillMonster) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this TreasureShop) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	shopStr := "["
	if len(this.Shop) > 0 {
		for shopId, addType := range this.Shop {
			shopStr += fmt.Sprintf(`[%d,%d],`, shopId, addType)
		}
		shopStr = shopStr[:len(shopStr)-1]
	}
	shopStr += "]"
	carStr := "["
	if len(this.Car) > 0 {
		for shopId, num := range this.Car {
			carStr += fmt.Sprintf(`[%d,%d],`, shopId, num)
		}
		carStr = carStr[:len(carStr)-1]
	}
	carStr += "]"
	refreshFree := 0
	if this.RefreshFree {
		refreshFree = 1
	}
	buf.WriteString(fmt.Sprintf(`[[[%d]],[[%d]],[[%d]],%s,%s]`, refreshFree, this.RefreshTime, this.BuyNum, shopStr, carStr))
	return buf.Bytes(), nil
}
func (this *TreasureShop) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		datas := make([][][]int, 0)
		if err := json.Unmarshal(data, &datas); err != nil {
			return err
		}
		refreshFree := false
		if datas[0][0][0] == 1 {
			refreshFree = true
		}
		shopMap := make(IntKv)
		for _, ints := range datas[3] {
			val := 0
			if len(ints) > 1 {
				val = ints[1]
			}
			shopMap[ints[0]] = val
		}
		carMap := make(IntKv)
		for _, ints := range datas[4] {
			val := 0
			if len(ints) > 1 {
				val = ints[1]
			}
			carMap[ints[0]] = val
		}
		(*this).RefreshFree = refreshFree
		(*this).RefreshTime = datas[1][0][0]
		(*this).BuyNum = datas[2][0][0]
		(*this).Shop = shopMap
		(*this).Car = carMap
	}
	return nil
}
func (this *TreasureShop) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}
func (this TreasureShop) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this AncientTreasuresInfo) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	titleStr := ""
	if len(this) > 0 {
		for id, info := range this {
			titleStr += fmt.Sprintf("[%d,%d,%d,%d,%d],", id, info.ZhuLinLv, info.Star, info.JueXinLv, info.Types)
		}
		titleStr = titleStr[:len(titleStr)-1]
	}
	buf.WriteString(fmt.Sprintf(`[%s]`, titleStr))
	return buf.Bytes(), nil
}

func (this *AncientTreasuresInfo) UnmarshalJSON(data []byte) error {
	(*this) = make(map[int]*AncientTreasures)
	if len(data) > 0 {
		datas := make([][]int, 0)
		if err := json.Unmarshal(data, &datas); err != nil {
			return err
		}
		for _, info := range datas {
			(*this)[info[0]] = &AncientTreasures{
				ZhuLinLv: info[1],
				Star:     info[2],
				JueXinLv: info[3],
				Types:    info[4],
			}
		}
	}
	return nil
}

func (this *AncientTreasuresInfo) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this AncientTreasuresInfo) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this HellBoss) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%d]`, this.DareNum, this.BuyNum, this.ResetTime, this.HelpNum))
	return buf.Bytes(), nil
}
func (this *HellBoss) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		datas := make([]int, 0)
		if err := json.Unmarshal(data, &datas); err != nil {
			return err
		}
		(*this).DareNum = datas[0]
		(*this).BuyNum = datas[1]
		(*this).ResetTime = datas[2]
		(*this).HelpNum = datas[3]
	}
	return nil
}

func (this *HellBoss) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this HellBoss) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this LotteryInfo) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%d]`, this.ResetDay, this.PopUpState, this.GoodLuckState, this.IsGetAward))
	return buf.Bytes(), nil
}
func (this *LotteryInfo) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		datas := make([]int, 0)
		if err := json.Unmarshal(data, &datas); err != nil {
			return err
		}
		(*this).ResetDay = datas[0]
		(*this).PopUpState = datas[1]
		(*this).GoodLuckState = datas[2]
		if len(datas) > 3 {
			(*this).IsGetAward = datas[3]
		}
	}
	return nil
}

func (this *LotteryInfo) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this LotteryInfo) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this TrialTaskInfos) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	titleStr := ""
	if len(this) > 0 {
		for id, info := range this {
			titleStr += fmt.Sprintf("[%d,%d,%d],", id, info.MarkNum, info.IsGetAward)
		}
		titleStr = titleStr[:len(titleStr)-1]
	}
	buf.WriteString(fmt.Sprintf(`[%s]`, titleStr))
	return buf.Bytes(), nil
}

func (this *TrialTaskInfos) UnmarshalJSON(data []byte) error {
	(*this) = make(map[int]*TrialTaskInfo)
	if len(data) > 0 {
		datas := make([][]int, 0)
		if err := json.Unmarshal(data, &datas); err != nil {
			return err
		}
		for _, info := range datas {
			(*this)[info[0]] = &TrialTaskInfo{
				MarkNum:    info[1],
				IsGetAward: info[2],
			}
		}
	}
	return nil
}

func (this *TrialTaskInfos) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this TrialTaskInfos) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this *DaBaoMystery) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`[%d,%d]`, this.Energy, this.ResumeTime))
	return buf.Bytes(), nil
}
func (this *DaBaoMystery) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		datas := make([]int, 0)
		if err := json.Unmarshal(data, &datas); err != nil {
			return err
		}
		(*this).Energy = datas[0]
		(*this).ResumeTime = datas[1]
	}
	return nil
}
func (this *DaBaoMystery) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}
func (this DaBaoMystery) Value() (driver.Value, error) {
	return this.MarshalJSON()
}

func (this *Applets) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	listStr := "["
	if len(this.List) > 0 {
		for t, info := range this.List {
			listStr += fmt.Sprintf(`[%d,%d,%d,%d],`, t, info.Stage, info.LastGetAwardTime, info.IsInGame)
		}
		listStr = listStr[:len(listStr)-1]
	}
	listStr += "]"
	buf.WriteString(fmt.Sprintf(`[%d,%d,%s]`, this.Energy, this.ResumeTime, listStr))
	return buf.Bytes(), nil
}
func (this *Applets) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		datas := make([]interface{}, 0)
		if err := json.Unmarshal(data, &datas); err != nil {
			return err
		}
		listInterface := datas[2].([]interface{})
		listMap := make(map[int]*AppletsUnit)
		for _, list := range listInterface {
			listArr := list.([]interface{})
			listMap[int(listArr[0].(float64))] = &AppletsUnit{
				Stage:            int(listArr[1].(float64)),
				LastGetAwardTime: int64(listArr[2].(float64)),
			}
			if len(listArr) > 3 {
				listMap[int(listArr[0].(float64))].IsInGame = int(listArr[3].(float64))
			}
		}
		(*this).Energy = int(datas[0].(float64))
		(*this).ResumeTime = int64(datas[1].(float64))
		(*this).List = listMap
	}
	return nil
}
func (this *Applets) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}
func (this Applets) Value() (driver.Value, error) {
	return this.MarshalJSON()
}

func (this *Label) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	taskStr := "["
	if len(this.TaskOver) > 0 {
		for taskId := range this.TaskOver {
			taskStr += fmt.Sprintf(`%d,`, taskId)
		}
		taskStr = taskStr[:len(taskStr)-1]
	}
	taskStr += "]"
	buf.WriteString(fmt.Sprintf(`[%d,%d,%d,%d,%d,%d,%s]`,
		this.Id, this.Job, this.Transfer, this.RefTime, marshalBool(this.DayReward), marshalBool(this.FirstTransfer), taskStr))
	return buf.Bytes(), nil
}
func (this *Label) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		datas := make([]interface{}, 0)
		if err := json.Unmarshal(data, &datas); err != nil {
			return err
		}
		(*this).Id = int(datas[0].(float64))
		(*this).Job = int(datas[1].(float64))
		(*this).Transfer = int(datas[2].(float64))
		(*this).RefTime = int(datas[3].(float64))
		(*this).DayReward = unmarshalBool(int(datas[4].(float64)))
		if len(datas) > 5 {
			(*this).FirstTransfer = unmarshalBool(int(datas[5].(float64)))
		}
		taskOverMap := make(IntKv)
		if len(datas) > 6 {
			taskIdInterface := datas[6].([]interface{})
			for _, taskId := range taskIdInterface {
				tId := int(taskId.(float64))
				taskOverMap[tId] = 0
			}
		}
		(*this).TaskOver = taskOverMap
	}
	return nil
}
func (this *Label) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}
func (this Label) Value() (driver.Value, error) {
	return this.MarshalJSON()
}

func (this *ModuleUpMaxLv) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`[%d]`, this.BaoSiLv))
	return buf.Bytes(), nil
}

func (this *ModuleUpMaxLv) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		datas := make([]int, 0)
		if err := json.Unmarshal(data, &datas); err != nil {
			return err
		}
		(*this).BaoSiLv = datas[0]
	}
	return nil
}

func (this *ModuleUpMaxLv) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this ModuleUpMaxLv) Value() (driver.Value, error) {
	return this.MarshalJSON()
}
