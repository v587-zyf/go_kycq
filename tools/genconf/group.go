package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"time"
)

// 需要导出的gameDb中的一个array或map
type CollectionField struct {
	Type              reflect.Type // 原始的gameDb中的字段类型
	FieldT            reflect.StructField
	FieldV            reflect.Value
	FieldIndex        int
	FieldName         string // 原始的gameDb中的字段名
	ExportArrayName   string // 输出为array的名字
	ExportMapName     string // 输出为map的名字
	MapKeyName        string // 输出为map时，key使用的字段名
	MapKeyIndex       int    // 输出为map时，key使用的字段的序号
	MaxKeyName        string // 输出max...字段时的名字
	GroupName         string // 分组名称
	OnDemandGroupName string // 服务器动态分组名 存放客户端通过api动态获得的数据 该项设置时GroupName不生成具体内容，只有定义。
	IsOnDemandGroup   bool
	SkipOnDemandData  string //1:不导出onDemandData文件 0:导出
	IsExportArray     bool   // 是否输出为array
	IsExportMap       bool   // 是否输出为map
	IsPlainType       bool   // 原始的字段为map或array，且存储的值为string或int
	IsKeyString       bool   // 当IsPlainType为true时，指明map中的key是否是string
	IsValueString     bool   // 当IsPlainType为true时，指明map中的value是否是string
	ExportStruct      *ExportStruct
}

type ExportGroup struct {
	Name         string
	OnDemandName string
	Fields       []*CollectionField

	IsFirstGroup    bool
	Structs         []*ExportStruct // 本group需要导出的结构体
	IsReducedExport bool            // 是否缩减输出对象的层级，只用json.Marshal之后的string代替
}

var exportGroups = make(map[string]*ExportGroup)

const ArrMapFieldName = "__arrmap__"

func NewCollectionField() *CollectionField {
	return &CollectionField{}
}

func NewOrGetExportGroup(name string, isReducedExport bool) *ExportGroup {
	if len(name) == 0 {
		name = "data"
	}
	if group, ok := exportGroups[name]; ok {
		return group
	}
	group := &ExportGroup{}
	group.Name = name
	group.IsReducedExport = isReducedExport
	group.Fields = make([]*CollectionField, 0)
	group.Structs = make([]*ExportStruct, 0)
	exportGroups[name] = group
	return group
}

func CollectExports(gameDb interface{}, baseGroupName string, isReducedExport bool) error {

	now := time.Now()
	defer func() {
		fmt.Println("CollectExports use time", time.Since(now).Seconds())
	}()

	objT := getStructElem(reflect.TypeOf(gameDb))
	for i := 0; i < objT.NumField(); i++ {
		fieldT := objT.Field(i)
		group := NewOrGetExportGroup(fieldT.Tag.Get("group"), isReducedExport)
		group.AddExport(gameDb, i, false)
		/* 20190705 不输出文件
		if fieldT.Tag.Get("onDemandGroup") != "" {
			group := NewOrGetExportGroup(fieldT.Tag.Get("onDemandGroup"), isReducedExport)
			group.OnDemandName = group.Name
			group.AddExport(gameDb, i, true)
		}
		*/
	}
	var firstGroup *ExportGroup
	for _, group := range exportGroups {
		group.CollectStructs()
		group.IsFirstGroup = group.Name == baseGroupName
		if group.IsFirstGroup {
			firstGroup = group
		}
	}
	if firstGroup == nil { // 没有匹配到baseGroup则创建一个
		group := NewOrGetExportGroup(baseGroupName, isReducedExport)
		group.IsFirstGroup = true
		firstGroup = group
	}
	// 只有第一组导出class
	firstGroup.Structs = exportStructs.GetSortedStructs()
	return nil
}

func GetExportedStructs() []*ExportStruct {
	return exportStructs.GetSortedStructs()
}

func GetExportedFields() []*CollectionField {
	var fields = make([]*CollectionField, 0)
	for _, group := range exportGroups {
		fields = append(fields, group.Fields...)
	}
	return fields
}

func DoExport(rootPath string, gameDb interface{}, baseGroupName string, isReducedExport bool) error {
	CollectExports(gameDb, baseGroupName, isReducedExport)
	err := InitTemplate(rootPath, gameDb)
	if err != nil {
		return err
	}
	for _, group := range exportGroups {
		err = group.Export(gameDb)
		if err != nil {
			return err
		}
	}
	err = ExportDecl(gameDb)
	if err != nil {
		return err
	}
	err = ExportDef(gameDb)
	return err
}

func ExportDecl(gameDb interface{}) error {
	var buf bytes.Buffer
	var data = struct {
		GameDb  interface{}
		Groups  map[string]*ExportGroup
		Structs []*ExportStruct
	}{
		GameDb:  gameDb,
		Groups:  exportGroups,
		Structs: exportStructs.GetSortedStructs(),
	}
	err := exportTemplate.ExecuteTemplate(&buf, "decl.tmpl", data)
	if err != nil {
		return err
	}
	var toFileName = "data.d.ts"
	return ioutil.WriteFile(toFileName, buf.Bytes(), os.ModePerm)
}

func ExportDef(gameDb interface{}) error {
	var buf bytes.Buffer
	var data = struct {
		GameDb  interface{}
		Structs []*ExportStruct
		Fields  []*CollectionField
	}{
		GameDb:  gameDb,
		Structs: GetExportedStructs(),
		Fields:  GetExportedFields(),
	}
	err := exportTemplate.ExecuteTemplate(&buf, "def.tmpl", data)
	if err != nil {
		return err
	}
	var toFileName = "data.js"
	return ioutil.WriteFile(toFileName, buf.Bytes(), os.ModePerm)
}

func (this *CollectionField) CollectStruct() {
	var fieldT = this.FieldT
	var elemT reflect.Type
	switch fieldT.Type.Kind() {
	case reflect.Map:
		if this.IsExportArray {
			fmt.Printf("can not export map as an array: gameDb.%s\n", fieldT.Name)
			return
		}
		elemT = getElem(fieldT.Type.Elem())
	case reflect.Slice, reflect.Array:
		elemT = getElem(fieldT.Type.Elem())
	default:
		fmt.Printf("can not auto export: gameDb.%s\n", fieldT.Name)
		return
	}
	if elemT.Kind() == reflect.Struct {
		st, err := exportStructs.CollectStruct(elemT)
		if err != nil {
			fmt.Printf("collect struct error for field gameDb.%s: %s\n", fieldT.Name, err)
		}
		this.ExportStruct = st
	} else {
		this.IsPlainType = true // 是string或int类型
		switch fieldT.Type.Kind() {
		case reflect.Map:
			this.IsKeyString = fieldT.Type.Key().Kind() == reflect.String
			this.IsValueString = elemT.Kind() == reflect.String
		case reflect.Slice, reflect.Array:
			this.IsValueString = elemT.Kind() == reflect.String
		}
	}
}

// 增加gameDb中的输出字段
func (this *ExportGroup) AddExport(gameDb interface{}, fieldIndex int, isOnDemandGroup bool) error {
	objT := getElem(reflect.TypeOf(gameDb))
	fieldT := objT.Field(fieldIndex)
	tag := fieldT.Tag.Get("client")
	if len(tag) == 0 {
		return nil
	}
	mapKey := fieldT.Tag.Get("mapKey")
	if len(mapKey) == 0 {
		mapKey = "Id"
	}
	maxKeyName := fieldT.Tag.Get("maxKeyName")
	if len(maxKeyName) == 0 {
		maxKeyName = "max" + fieldT.Name
	}
	fields := strings.Split(tag, ",")
	if len(fields) < 2 {
		return nil
	}
	var field = NewCollectionField()
	field.FieldName = fieldT.Name
	field.FieldIndex = fieldIndex
	field.FieldT = fieldT
	field.IsExportArray = fields[1] == "array" || fields[1] == "arrmap"
	field.IsExportMap = fields[1] == "map" || fields[1] == "arrmap"
	field.MapKeyName = mapKey
	field.MaxKeyName = maxKeyName
	var names = strings.Split(fields[0], ":")
	if len(names) > 1 {
		field.ExportArrayName = names[0]
		field.ExportMapName = names[1]
	} else {
		if field.IsExportArray && field.IsExportMap {
			return fmt.Errorf("collect struct error for field %s: %s", fieldT.Name, "missing export name")
		}
		if field.IsExportArray {
			field.ExportArrayName = names[0]
		} else if field.IsExportMap {
			field.ExportMapName = names[0]
		}
	}
	onDemandGroup := fieldT.Tag.Get("onDemandGroup")
	field.OnDemandGroupName = onDemandGroup
	field.IsOnDemandGroup = isOnDemandGroup
	field.SkipOnDemandData = fieldT.Tag.Get("skipOnDemandData")
	this.Fields = append(this.Fields, field)
	return nil
}

func (this *ExportGroup) CollectStructs() error {
	for _, field := range this.Fields {
		field.CollectStruct()
		// 填充MapKeyIndex
		if field.ExportStruct != nil && field.MapKeyName != "" {
			for i, f := range field.ExportStruct.Fields {
				if f.Name == field.MapKeyName {
					field.MapKeyIndex = i
				}
			}
		}
	}
	return nil
}

func (this *ExportGroup) Export(gameDb interface{}) error {

	now := time.Now()
	defer func() {
		fmt.Println(this.Name, "Export use time", time.Since(now).Seconds())
	}()

	var buf bytes.Buffer
	err := exportTemplate.ExecuteTemplate(&buf, this.Name, gameDb)
	if err != nil {
		return err
	}

	var data map[string]interface{}
	var rb = buf.Bytes()
	err = json.Unmarshal(rb, &data)
	if err != nil {
		err = toDetailedJsonError(err, rb)
	} else if this.IsReducedExport {
		arrmap, ok := data[ArrMapFieldName]
		delete(data, ArrMapFieldName)
		data = convertJsonField(data, 2).(map[string]interface{})
		if ok {
			data[ArrMapFieldName] = arrmap
		}
		rb, err = json.Marshal(data)
	}
	var toFileName = this.Name + ".json"
	ioutil.WriteFile(toFileName, rb, os.ModePerm)
	return err
}

// 递归遍历data到level层级，如果level等于0且字段值不是基础类型，则该字段转换为json.Marshal后的string
func convertJsonField(data interface{}, level int) interface{} {
	if level < 0 {
		return data
	}
	rv := reflect.ValueOf(data) // reflect value
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	switch rv.Type().Kind() {
	case reflect.Map:
		keys := rv.MapKeys()
		for _, mk := range keys {
			mv := rv.MapIndex(mk)
			if mv.Type().Kind() == reflect.Ptr {
				mv = mv.Elem()
			}
			rv.SetMapIndex(mk, reflect.ValueOf(convertJsonField(mv.Interface(), level-1)))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			rv.Index(i).Set(reflect.ValueOf(convertJsonField(rv.Index(i).Interface(), level-1)))
		}
	default:
		return data
	}
	if level == 0 {
		rb, _ := json.Marshal(data)
		return string(rb)
	}
	return data
}

func toDetailedJsonError(err error, dataSource []byte) error {
	var offset int64
	if jsonError, ok := err.(*json.UnmarshalTypeError); ok {
		offset = jsonError.Offset
	} else if jsonError, ok := err.(*json.SyntaxError); ok {
		offset = jsonError.Offset
	}
	if offset > 0 {
		start := offset - 200
		if start < 0 {
			start = 0
		}
		return errors.New(fmt.Sprintf("生成json时json格式不对，%v,\n可能出错的地方,注意最后一个字符:%s", err, dataSource[start:offset]))
	}
	return err
}
