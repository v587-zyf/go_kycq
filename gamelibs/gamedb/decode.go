package gamedb

import (
	"cqserver/golibs/logger"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"cqserver/golibs/common"
	"cqserver/protobuf/pb"
)

const (
	SEMICOLON = ";"
	COMMA     = ","
	COLON     = ":"
	PIPE      = "|"
	SPACE     = " "
	HLINE     = "-"
)

type ItemInfo struct {
	ItemId int `client:"key"`
	Count  int `client:"value"`
}

type PropInfo struct {
	K int `client:"key"`
	V int `client:"value"`
}

type ProbValue struct {
	Value int `client:"value"` //值
	Prob  int `client:"prob"`  //几率
}

type ProbItem struct {
	Id    int `client:"id"`    // 物品ID
	Count int `client:"count"` // 数量
	Prob  int `client:"prob"`  // 几率
}

type RangeNum struct {
	Min int `client:"min"`
	Max int `client:"max"`
}

type SignItem struct {
	Id    int `client:"id"`    //物品ID
	Count int `client:"count"` //数量
}

type HmsTime struct {
	Hour   int `client:"hour"`
	Minute int `client:"minute"`
	Second int `client:"second"`
}

type Condition struct {
	K    int         `client:"key"`
	V    int         `client:"value"`
	Subs map[int]int `client:"subs"`
}

type DateTime struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Second int
}

type ItemWithWeight struct {
	Id     int
	Count  int
	Weight int
}

type Conditions []*Condition
type ItemInfos []*ItemInfo
type ItemInfosSlice []ItemInfos
type PropInfos []*PropInfo
type ProbValues []*ProbValue
type ProbItems []*ProbItem
type RangeNums []*RangeNum
type FloatSlice []float64
type IntSlice []int
type IntSlice2 [][]int
type IntSlice3 [][][]int
type StringSlice []string
type StringSlice2 [][]string

type HmsTimes []*HmsTime
type IntMap map[int]int
type IntMapSlice []IntMap

type ItemsWithWeight []*ItemWithWeight

/***************************************************************/
/************************数据类型解析*****************************/
/***************************************************************/
func (this *IntSlice) Decode(str string) error {
	ints, err := common.IntSliceFromString(str, PIPE)
	if err != nil {
		return err
	}
	*this = IntSlice(ints)
	return nil
}

func (this *FloatSlice) Decode(str string) error {
	ints, err := common.Float64SliceFromString(str, PIPE)
	if err != nil {
		return err
	}
	*this = FloatSlice(ints)
	return nil
}

func (this IntSlice) ToInt32Slice() []int32 {
	l := len(this)
	ret := make([]int32, l)
	if l == 0 {
		return ret
	}
	for i := 0; i < l; i++ {
		ret[i] = int32(this[i])
	}
	return ret
}

func (this *StringSlice) Decode(str string) error {
	if len(strings.TrimSpace(str)) == 0 {
		*this = make([]string, 0)
		return nil
	}
	*this = strings.Split(str, PIPE)
	return nil
}

func (this ItemInfos) ToPbItems() []*pb.Item {
	result := make([]*pb.Item, 0, len(this))
	for _, itemInfo := range this {
		result = append(result, &pb.Item{ItemId: int32(itemInfo.ItemId), Count: int64(itemInfo.Count)})
	}
	return result
}
func (this ItemInfos) ToInt32Map() map[int32]int32 {
	result := make(map[int32]int32, len(this))
	for _, itemInfo := range this {
		result[int32(itemInfo.ItemId)] += int32(itemInfo.Count)
	}
	return result
}

func (this ItemInfos) Get(itemId int) *ItemInfo {
	for _, itemInfo := range this {
		if itemInfo.ItemId == itemId {
			return itemInfo
		}
	}
	return nil
}

func (this *ItemInfos) Decode(str string) error {
	*this = make(ItemInfos, 0)
	if len(str) == 0 {
		return nil
	}
	infoList := strings.Split(strings.Trim(strings.TrimSpace(str), PIPE), PIPE)
	if len(infoList) == 0 {
		return nil
	}

	for _, v := range infoList {
		list := strings.Split(strings.TrimSpace(v), COMMA)
		if len(list) < 2 {
			return errors.New(v + "物品信息格式错误")
		}
		itemId, _ := strconv.Atoi(list[0])
		var itemInfo ItemInfo
		itemInfo.ItemId = itemId
		itemCount, err := strconv.Atoi(list[1])
		if err != nil {
			return err
		}
		itemInfo.Count = itemCount
		*this = append(*this, &itemInfo)
	}

	return nil
}

func (this *ItemInfosSlice) Decode(str string) error {
	*this = make(ItemInfosSlice, 0)
	if len(str) == 0 {
		return nil
	}
	infoList := strings.Split(strings.TrimSpace(str), PIPE)
	if len(infoList) == 0 {
		return nil
	}

	for _, v := range infoList {
		itemInfos := &ItemInfos{}
		err := itemInfos.Decode(v)
		if err != nil {
			return err
		}
		*this = append(*this, *itemInfos)
	}

	return nil
}

func (this *PropInfos) Decode(str string) error {
	*this = make(PropInfos, 0)
	if len(str) == 0 {
		return nil
	}
	infoList := strings.Split(strings.Trim(strings.TrimSpace(str), PIPE), PIPE)
	if len(infoList) == 0 {
		return nil
	}

	for _, v := range infoList {
		list := strings.Split(strings.TrimSpace(v), COMMA)
		if len(list) < 2 {
			return errors.New(v + "属性信息格式错误")
		}
		k, err := strconv.Atoi(list[0])
		if err != nil {
			return err
		}
		var propInfo PropInfo
		propInfo.K = k
		propInfo.V, err = strconv.Atoi(list[1])
		if err != nil {
			return err
		}
		*this = append(*this, &propInfo)
	}
	return nil
}

func (this *IntMap) Decode(str string) error {
	if len(strings.TrimSpace(str)) == 0 {
		return nil
	}
	//fmt.Printf("str = %+v\n", str)
	infoList := strings.Split(strings.Trim(strings.TrimSpace(str), PIPE), PIPE)
	if len(infoList) == 0 {
		return nil
	}

	*this = make(IntMap)
	for _, v := range infoList {
		list := strings.Split(strings.TrimSpace(v), COMMA)
		if len(list) != 2 {
			return errors.New(v + "IntMap 属性信息格式错误")
		}

		k, err := strconv.Atoi(list[0])
		if err != nil {
			return err
		}
		if _, ok := (*this)[k]; ok {
			return errors.New(v + "IntMap 属性重复")
		}
		v, err := strconv.Atoi(list[1])
		if err != nil {
			return err
		}
		(*this)[k] = v
	}
	//fmt.Printf("decode string is %s intmap:%v", str, *this)
	return nil

}

func (this *PropInfo) Decode(str string) error {
	if len(strings.TrimSpace(str)) == 0 {
		return nil
	}
	infoList := strings.Split(str, PIPE)
	if len(infoList) < 2 {
		return errors.New(str + "属性信息格式错误")
	}
	var propInfo PropInfo
	propInfo.K, _ = strconv.Atoi(infoList[0])
	propInfo.V, _ = strconv.Atoi(infoList[1])
	*this = propInfo
	return nil
}
func (this *ItemInfo) Decode(str string) error {
	if len(strings.TrimSpace(str)) == 0 {
		return nil
	}
	infoList := strings.Split(str, PIPE)
	if len(infoList) < 2 {
		return errors.New(str + "属性信息格式错误")
	}
	var itemInfo ItemInfo
	var err error
	itemInfo.ItemId, err = strconv.Atoi(infoList[0])
	if err != nil {
		return err
	}
	itemInfo.Count, err = strconv.Atoi(infoList[1])
	if err != nil {
		return err
	}
	*this = itemInfo
	return nil
}

func (this *IntMapSlice) Decode(str string) error {
	*this = make(IntMapSlice, 0)
	if len(str) == 0 {
		return nil
	}
	infoList := strings.Split(strings.TrimSpace(str), PIPE)
	if len(infoList) == 0 {
		return nil
	}

	for _, v := range infoList {
		intMap := &IntMap{}
		err := intMap.Decode(v)
		if err != nil {
			return err
		}
		*this = append(*this, *intMap)
	}
	return nil
}
func (this *ProbValue) Decode(str string) error {
	if len(strings.TrimSpace(str)) == 0 {
		return nil
	}
	infoList := strings.Split(strings.TrimSpace(str), COMMA)
	if len(infoList) < 2 {
		return errors.New(str + "属性信息格式错误")
	}
	var probInfo ProbValue
	probInfo.Value, _ = strconv.Atoi(infoList[0])
	probInfo.Prob, _ = strconv.Atoi(infoList[1])
	*this = probInfo
	return nil
}

func (this *HmsTime) String() string {
	return fmt.Sprintf("%d:%d:%d", this.Hour, this.Minute, this.Second)
}

func (this *HmsTimes) Decode(str string) error {
	*this = make(HmsTimes, 0)
	if len(strings.TrimSpace(str)) == 0 {
		return nil
	}
	infoList := strings.Split(strings.TrimSpace(str), PIPE)
	if len(infoList) < 1 {
		return errors.New(str + "属性信息格式错误")
	}
	for _, v := range infoList {
		one := &HmsTime{}
		err := one.Decode(v)
		if err != nil {
			return err
		}
		*this = append(*this, one)
	}
	return nil
}

func (this *HmsTime) Decode(str string) error {
	if len(strings.TrimSpace(str)) == 0 {
		return nil
	}
	infoList := strings.Split(strings.TrimSpace(str), COLON)
	if len(infoList) < 1 {
		return errors.New(str + "属性信息格式错误")
	}
	var hms HmsTime
	hms.Hour, _ = strconv.Atoi(infoList[0])
	if len(infoList) > 1 {
		hms.Minute, _ = strconv.Atoi(infoList[1])
	}
	if len(infoList) > 2 {
		hms.Second, _ = strconv.Atoi(infoList[2])
	}
	if hms.Hour < 0 || hms.Hour > 23 || hms.Minute < 0 || hms.Minute > 59 || hms.Second < 0 || hms.Second > 59 {
		return errors.New(str + "时分秒不对")
	}
	*this = hms
	return nil
}

func (this *HmsTime) GetSecondsFromZero() int { //从0点到该时刻的秒数
	return this.Hour*60*60 + this.Minute*60 + this.Second
}

func (this *DateTime) String() string {
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", this.Year, this.Month, this.Day, this.Hour, this.Minute, this.Second)
}

func (this *DateTime) Unix() int64 {
	if this.Month == 0 {
		return 0
	}
	seconds := time.Date(this.Year, time.Month(this.Month), this.Day, this.Hour, this.Minute, this.Second, 0, time.Local)
	return seconds.Unix()
}

func (this *DateTime) Unix32() int32 {
	return int32(this.Unix())
}

func (this *DateTime) Decode(str string) error {
	if len(strings.TrimSpace(str)) == 0 {
		return nil
	}
	infoList := strings.Split(strings.TrimSpace(str), SPACE)
	if len(infoList) < 1 {
		return errors.New(str + " DateTime 属性信息格式错误")
	}

	infoListLeft := strings.Split(strings.TrimSpace(infoList[0]), HLINE)
	if len(infoListLeft) < 1 {
		return errors.New(str + " DateTime 属性格式错误")
	}
	this.Year, _ = strconv.Atoi(infoListLeft[0])
	this.Month, _ = strconv.Atoi(infoListLeft[1])
	this.Day, _ = strconv.Atoi(infoListLeft[2])

	infoListRight := strings.Split(strings.TrimSpace(infoList[1]), COLON)
	if len(infoListRight) < 1 {
		return errors.New(str + " DateTime 属性信息格式错误")
	}
	this.Hour, _ = strconv.Atoi(infoListRight[0])
	this.Minute, _ = strconv.Atoi(infoListRight[1])
	this.Second, _ = strconv.Atoi(infoListRight[2])

	return nil
}

func (this *ProbValues) Decode(str string) error {
	*this = make(ProbValues, 0)
	if len(str) == 0 {
		return nil
	}
	infoList := strings.Split(strings.Trim(strings.TrimSpace(str), SEMICOLON), SEMICOLON)
	if len(infoList) == 0 {
		return nil
	}

	for _, v := range infoList {
		list := strings.Split(strings.TrimSpace(v), COMMA)
		if len(list) < 2 {
			continue
		}
		k, err := strconv.Atoi(list[0])
		if err != nil {
			return err
		}
		var probValue ProbValue
		probValue.Value = k
		probValue.Prob, err = strconv.Atoi(list[1])
		if err != nil {
			return err
		}
		*this = append(*this, &probValue)
	}
	return nil
}

func (this *ProbItems) Decode(str string) error {
	*this = make(ProbItems, 0)
	if len(str) == 0 {
		return nil
	}
	infoList := strings.Split(strings.Trim(strings.TrimSpace(str), SEMICOLON), SEMICOLON)
	if len(infoList) == 0 {
		return nil
	}

	for _, v := range infoList {
		list := strings.Split(strings.TrimSpace(v), COMMA)
		if len(list) != 3 {
			return errors.New(v + "属性信息格式错误")
		}
		id, err := strconv.Atoi(list[0])
		if err != nil {
			return err
		}
		var probItem ProbItem
		probItem.Id = id
		probItem.Count, err = strconv.Atoi(list[1])
		probItem.Prob, err = strconv.Atoi(list[2])
		if err != nil {
			return err
		}
		*this = append(*this, &probItem)
	}
	return nil
}

func (this *RangeNums) Decode(str string) error {
	*this = make(RangeNums, 0)
	if len(str) == 0 {
		return nil
	}
	infoList := strings.Split(str, SEMICOLON)
	for _, v := range infoList {
		ranges := strings.Split(v, COLON)
		if len(ranges) < 2 {
			return errors.New(str + "属性信息格式错误")
		}
		var rangeNum RangeNum
		rangeNum.Min, _ = strconv.Atoi(ranges[0])
		rangeNum.Max, _ = strconv.Atoi(ranges[1])
		*this = append(*this, &rangeNum)
	}
	return nil
}

func (this *SignItem) Decode(str string) error {
	infoList := strings.Split(str, COMMA)
	if len(infoList) < 2 {
		return errors.New(str + "属性信息格式错误")
	}
	var signItem SignItem
	signItem.Id, _ = strconv.Atoi(infoList[0])
	signItem.Count, _ = strconv.Atoi(infoList[1])
	*this = signItem
	return nil
}

func (this *RangeNum) Decode(str string) error {
	infoList := strings.Split(str, COLON)
	if len(infoList) < 2 {
		return errors.New(str + "属性信息格式错误")
	}
	var rangeNum RangeNum
	rangeNum.Min, _ = strconv.Atoi(infoList[0])
	rangeNum.Max, _ = strconv.Atoi(infoList[1])
	*this = rangeNum
	return nil
}

func (this *IntSlice2) Decode(str string) error {
	*this = make(IntSlice2, 0)
	if len(str) == 0 {
		return nil
	}
	infoList := strings.Split(strings.Trim(strings.TrimSpace(str), PIPE), PIPE)
	if len(infoList) == 0 {
		return nil
	}

	for _, v := range infoList {
		ints, err := common.IntSliceFromString(v, COMMA)
		if err != nil {
			return err
		}
		*this = append(*this, ints)
	}
	return nil
}

func (this *IntMap) Add(delta IntMap) {
	for k, v := range delta {
		(*this)[k] += v
	}
}
func (this IntMap) Clone() IntMap {
	ret := make(IntMap, len(this))
	for k, v := range this {
		ret[k] = v
	}
	return ret
}

func (this IntSlice) String(sep string) string {
	var arrStr = make([]string, len(this))
	for i, v := range this {
		arrStr[i] = strconv.Itoa(v)
	}
	return strings.Join(arrStr, sep)
}

func (this *Condition) Decode(str string) error {
	if len(strings.TrimSpace(str)) == 0 {
		return nil
	}
	infoList, err := common.IntSliceFromString(str, COMMA)
	if err != nil {
		return err
	}
	c, err := NewCondition(infoList)
	if err != nil {
		return err
	}

	*this = c
	return nil
}

func NewCondition(infoList []int) (Condition, error) {
	l := len(infoList)
	var c Condition
	if l < 2 {
		return c, errors.New("condition 属性信息格式错误")
	}
	c.K, c.V = infoList[0], infoList[1]
	if l > 2 {
		if l%2 != 0 {
			return c, errors.New("condition 长度必须是偶数")
		}
		subs := make(map[int]int, l/2)
		for i := 2; i < l; i++ {
			subKey, subValue := infoList[i], infoList[i+1]
			subs[subKey] = subValue
			i++
		}
		c.Subs = subs
	}
	return c, nil
}

func (this *Conditions) Decode(str string) error {
	*this = make(Conditions, 0)
	if len(str) == 0 {
		return nil
	}
	infoList := strings.Split(strings.Trim(strings.TrimSpace(str), SEMICOLON), SEMICOLON)
	if len(infoList) == 0 {
		return nil
	}
	for _, one := range infoList {
		var c Condition
		err := c.Decode(one)
		if err != nil {
			return err
		}
		*this = append(*this, &c)
	}
	return nil
}
func (this *IntSlice3) Decode(str string) error {
	*this = make(IntSlice3, 0)
	if len(str) == 0 {
		return nil
	}
	infoList := strings.Split(strings.Trim(strings.TrimSpace(str), PIPE), PIPE)
	if len(infoList) == 0 {
		return nil
	}

	for _, val := range infoList {
		intSlice2 := make(IntSlice2, 0)
		intStr := strings.Split(strings.Trim(strings.TrimSpace(val), SEMICOLON), SEMICOLON)
		for _, v := range intStr {
			int3, err := common.IntSliceFromString(v, COMMA)
			if err != nil {
				return err
			}
			intSlice2 = append(intSlice2, int3)
		}
		*this = append(*this, intSlice2)
	}
	return nil
}

func (this *StringSlice2) Decode(str string) error {

	*this = make(StringSlice2, 0)
	if len(str) == 0 {
		return nil
	}
	stringSlice := strings.Split(strings.TrimSpace(str), SEMICOLON)
	for _, v := range stringSlice {
		allData := make(StringSlice, 0)
		data := strings.Split(strings.TrimSpace(v), COMMA)
		for _, v1 := range data {
			allData = append(allData, v1)
		}
		*this = append(*this, allData)
	}
	return nil
}

func (this *ItemsWithWeight) Decode(str string) error {
	*this = make(ItemsWithWeight, 0)
	if len(str) == 0 {
		return nil
	}
	infoList := strings.Split(str, SEMICOLON)
	for _, v := range infoList {
		ranges := strings.Split(v, COMMA)
		if len(ranges) < 3 {
			return errors.New(str + "属性信息格式错误")
		}
		var itemWithWeight ItemWithWeight
		itemWithWeight.Id, _ = strconv.Atoi(ranges[0])
		itemWithWeight.Count, _ = strconv.Atoi(ranges[1])
		itemWithWeight.Weight, _ = strconv.Atoi(ranges[2])
		*this = append(*this, &itemWithWeight)
	}
	return nil
}

func DecodeConfValues(obj interface{}, gameConfigs map[string]*GameBaseCfg) error {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	if !(objT.Kind() == reflect.Ptr && objT.Elem().Kind() == reflect.Struct) {
		return fmt.Errorf("%v must be a struct pointer", obj)
	}
	var values = make(map[string]string, 0)
	for _, v := range gameConfigs {
		values[v.Name] = v.Value
	}

	objT = objT.Elem()
	objV = objV.Elem()
	for i := 0; i < objT.NumField(); i++ {
		fieldV := objV.Field(i)
		if !fieldV.CanSet() {
			continue
		}
		fieldT := objT.Field(i)

		if fieldT.Type.Kind() == reflect.Ptr {
			//param := reflect.New(fieldT.Type.Elem())
			//groupConfigName := strings.TrimSpace(fieldT.Tag.Get("confgroup"))
			//if len(groupConfigName) > 0 {
			//	err := DecodeConfValues(groupConfigName, param.Interface(), gameConfigs)
			//	if err != nil {
			//		return err
			//	}
			//	fieldV.Set(param)
			//}
			logger.Warn("未知的指针类型name:%v,type:%v", fieldT.Name, fieldT.Type.Kind())
			continue
		}

		configName := strings.TrimSpace(fieldT.Tag.Get("conf"))
		if len(configName) == 0 {
			continue
		}
		defaultDefine := strings.TrimSpace(fieldT.Tag.Get("default"))
		value := values[configName]
		if len(value) == 0 {
			value = defaultDefine
		}
		cellString := strings.TrimSpace(value)
		if decoder, ok := fieldV.Addr().Interface().(Decoder); ok {
			err := decoder.Decode(cellString)
			if err != nil {
				return err
			}
			continue
		}
		switch fieldT.Type.Kind() {
		case reflect.Bool:
			b, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			fieldV.SetBool(b)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				fmt.Println("-------------------", configName)
				return err
			}
			fieldV.SetInt(x)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			x, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				fmt.Println("-------------------", configName)
				return err
			}
			fieldV.SetUint(x)
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			fieldV.SetFloat(x)
		case reflect.Interface:
			fieldV.Set(reflect.ValueOf(value))
		case reflect.String:
			fieldV.SetString(value)
		}
	}
	return nil
}
