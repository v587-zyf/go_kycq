package gamedb

import (
	"errors"
	"fmt"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

/*
	规则:
	gamedb的数据中
	1:如果是map,那么key必须是int类型或者string类型
	2:如果是array,那么索引就是表对应客户端的key值. 现有表:strengths,lvls,QqGameGrowUpGifts
*/

type FieldInfo struct {
	idx     int //字段索引
	isMap   bool
	isArray bool
}

type Demand struct {
	gamedb_rt reflect.Type
	gamedb_rv *reflect.Value
	fieldIdx  map[string]FieldInfo //在gamedb中对应的结构字段索引
}

var (
	demand *Demand
)

func InitDemandMap() {
	data := GetDb()
	rt := reflect.TypeOf(data)
	rv := reflect.ValueOf(data)

	rt2, rv2 := getElemTV(rt, &rv)

	demand = &Demand{
		gamedb_rt: rt2,
		gamedb_rv: rv2,
		fieldIdx:  make(map[string]FieldInfo),
	}

	num := rt2.NumField()
	for i := 0; i < num; i++ {
		f := rt2.Field(i)
		if f.Tag.Get("onDemandGroup") == "onDemandData" {
			tmp := f.Tag.Get("client")
			if tmp != "" {
				arr := strings.Split(tmp, ",")
				f_k := f.Type.Kind()
				demand.fieldIdx[strings.Trim(arr[0], " ")] = FieldInfo{
					idx:     i,
					isMap:   f_k == reflect.Map,
					isArray: f_k == reflect.Slice || f_k == reflect.Array,
				}
			}
		}
	}
}

//获取table所有值
func GetDemandDatas(table string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("GetDemandData panic:", time.Now(), err, string(debug.Stack()))
		}
	}()

	tableInfo, ok := demand.fieldIdx[table]
	if !ok {
		return "", errors.New("Not found " + table + " in DemandData")
	}
	result := ""

	fieldV := demand.gamedb_rv.Field(tableInfo.idx)

	if tableInfo.isMap {
		var strs []string
		keys := fieldV.MapKeys()
		for i := 0; i < len(keys); i++ {
			rv := keys[i]
			rt := rv.Type()
			keystr := demand.switchKind(rt, &rv)

			rv2 := fieldV.MapIndex(rv)
			rt2 := rv2.Type()
			valstr := demand.switchKind(rt2, &rv2)
			strs = append(strs, keystr+":"+valstr)
		}

		result = "[" + strings.Join(strs, ",") + "]"

	} else if tableInfo.isArray { //未实现

	} else {
		return "", errors.New("data is not map and array")
	}

	return result, nil
}

//获取table中key对应的json值
func GetDemandData(table string, key interface{}) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("GetDemandData panic:", time.Now(), err, string(debug.Stack()))
		}
	}()

	tableInfo, ok := demand.fieldIdx[table]
	if !ok {
		return "", errors.New("Not found " + table + " in DemandData")
	}
	result := ""

	fieldV := demand.gamedb_rv.Field(tableInfo.idx)

	key_rv := reflect.ValueOf(key)
	if tableInfo.isMap {
		result_rv := fieldV.MapIndex(key_rv)
		if !result_rv.IsValid() {
			return "", errors.New("key is not found")
		}
		result_rt := result_rv.Type()

		result = demand.switchKind(result_rt, &result_rv)

	} else if tableInfo.isArray {
		if isIntType(key_rv.Type()) {
			result_rv := fieldV.Index(int(key_rv.Int()))

			if !result_rv.IsValid() {
				return "", errors.New("key is not found")
			}

			result_rt := result_rv.Type()
			result = demand.switchKind(result_rt, &result_rv)
		} else {
			return "", errors.New("key is not integer")
		}
	} else {
		return "", errors.New("data is not map and array")
	}

	return result, nil
}

func (this *Demand) switchKind(rt reflect.Type, rv *reflect.Value) string {

	rt, rv = getElemTV(rt, rv)

	kind := rt.Kind()

	switch kind {
	case reflect.Struct:
		return this.structHandle(rt, rv)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return this.intHandle(rt, rv)
	case reflect.Float32, reflect.Float64:
		return this.floatHandle(rt, rv)
	case reflect.String:
		return this.stringHandle(rt, rv)
	case reflect.Slice:
		return this.sliceHandle(rt, rv)
	case reflect.Array:
		return this.arrayHandle(rt, rv)
	case reflect.Map:
		return this.mapHandle(rt, rv)
	}

	return ""
}

func isIntType(rt reflect.Type) bool {
	rt = getElemT(rt)

	switch rt.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

func isNumberType(rt reflect.Type) bool {
	rt = getElemT(rt)

	switch rt.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	case reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

func isStringType(rt reflect.Type) bool {
	rt = getElemT(rt)
	switch rt.Kind() {

	case reflect.String:
		return true
	}
	return false
}

func (this *Demand) intHandle(rt reflect.Type, rv *reflect.Value) string {
	return strconv.FormatInt(rv.Int(), 10)
}

func (this *Demand) floatHandle(rt reflect.Type, rv *reflect.Value) string {
	return strconv.FormatFloat(rv.Float(), 'f', 5, 64)
}

func (this *Demand) stringHandle(rt reflect.Type, rv *reflect.Value) string {
	return `"` + strings.Replace(rv.String(), `"`, `\"`, -1) + `"`
}

func (this *Demand) mapHandle(rt reflect.Type, rv *reflect.Value) string {
	var strs []string

	//num := rv.Len()
	keys := rv.MapKeys()
	for i := 0; i < len(keys); i++ {
		key_rv := keys[i]
		key_rt := key_rv.Type()

		key := this.switchKind(key_rt, &key_rv)
		if !isStringType(key_rt) {
			key = `"` + key + `"`
		}

		val_rv := rv.MapIndex(key_rv)
		val_rt := val_rv.Type()

		val := this.switchKind(val_rt, &val_rv)

		strs = append(strs, key+":"+val)

	}

	return "{" + strings.Join(strs, ",") + "}"
}

func (this *Demand) arrayHandle(rt reflect.Type, rv *reflect.Value) string {
	var strs []string

	num := rv.Len()
	for i := 0; i < num; i++ {

		e := rv.Index(i)

		field_rt := e.Type()
		field_rv := e

		//field_kind := field_rt.Kind()

		strs = append(strs, this.switchKind(field_rt, &field_rv))

	}

	return "[" + strings.Join(strs, ",") + "]"
}

func (this *Demand) sliceHandle(rt reflect.Type, rv *reflect.Value) string {
	var strs []string

	num := rv.Len()
	for i := 0; i < num; i++ {

		e := rv.Index(i)

		field_rt := e.Type()
		field_rv := e

		//field_kind := field_rt.Kind()

		strs = append(strs, this.switchKind(field_rt, &field_rv))

	}

	return "[" + strings.Join(strs, ",") + "]"
}

func (this *Demand) structHandle(rt reflect.Type, rv *reflect.Value) string {
	num := rt.NumField()
	var strs []string
	for i := 0; i < num; i++ {
		rt_Field := rt.Field(i)
		isclient := rt_Field.Tag.Get("client") != ""
		if isclient {
			field_rt := rt_Field.Type
			field_rv := rv.Field(i)

			//field_kind := field_rt.Kind()

			strs = append(strs, this.switchKind(field_rt, &field_rv))
		}
	}

	return "[" + strings.Join(strs, ",") + "]"

}

func getElemT(rt reflect.Type) reflect.Type {
	if rt.Kind() == reflect.Ptr {
		return rt.Elem()
	}
	return rt
}

func getElemTV(rt reflect.Type, rv *reflect.Value) (reflect.Type, *reflect.Value) {
	if rt.Kind() == reflect.Ptr {
		e := rv.Elem()
		return rt.Elem(), &e
	}
	return rt, rv
}
