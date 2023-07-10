package main

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type ExportField struct {
	Name       string
	ExportName string
	Type       reflect.Type
	// ExportType    string
	DeepStruct *ExportStruct
	IsDeep     bool // 是否需要按struct导出此字段
}

type ExportStruct struct {
	Name       string
	Fields     []*ExportField
	GetterFrom string
}

type ExportStructs struct {
	structs map[string]*ExportStruct
	getters []*ExportStruct
}

var exportStructs = &ExportStructs{
	structs: make(map[string]*ExportStruct),
	getters: make([]*ExportStruct, 0),
}

func NewExportStruct(name string) *ExportStruct {
	return &ExportStruct{
		Name:   name,
		Fields: make([]*ExportField, 0),
	}
}

func NewExportField(name string) *ExportField {
	return &ExportField{
		Name: name,
	}
}

func isStructType(objT reflect.Type) bool {
	return objT.Kind() == reflect.Struct || (objT.Kind() == reflect.Ptr && objT.Elem().Kind() == reflect.Struct)
}

func isNumberType(objT reflect.Type) bool {
	switch objT.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func isBaseType(objT reflect.Type) bool {
	switch objT.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Bool, reflect.String:
		return true
	default:
		return false
	}
}

func getBaseClientType(objT reflect.Type) string {
	switch objT.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return "number"
	case reflect.Bool:
		return "boolean"
	case reflect.String:
		return "string"
	}
	return ""
}

func getElem(objT reflect.Type) reflect.Type {
	if objT.Kind() == reflect.Ptr {
		return objT.Elem()
	} else {
		return objT
	}
}

func getStructElem(objT reflect.Type) reflect.Type {
	if !isStructType(objT) {
		return nil
	}
	return getElem(objT)
}

func getErrorInfo(info string, objT reflect.Type) error {
	return fmt.Errorf("%s, for %s", info, objT)
}

func (this *ExportField) ClientExportType() string {
	if isBaseType(this.Type) {
		return getBaseClientType(this.Type)
	}
	switch this.Type.Kind() {
	case reflect.Slice, reflect.Array:
		elemType := getElem(this.Type.Elem())
		if isBaseType(elemType) {
			return getBaseClientType(elemType) + "[]"
		} else if this.DeepStruct != nil {
			return this.DeepStruct.Name + "[]"
		} else if elemType.Kind() == reflect.Slice && isNumberType(elemType.Elem()) { //支持[][]int
			return "number[][]"
		}
	case reflect.Map:
		return "{[index: string]: number}"
	case reflect.Struct:
		return this.DeepStruct.Name
	}
	return ""
}

func (this *ExportField) IsArrayStruct() bool {
	if this.DeepStruct == nil {
		return false
	}
	return this.Type.Kind() == reflect.Slice || this.Type.Kind() == reflect.Array
}

func (this *ExportField) IsStruct() bool {
	if this.DeepStruct == nil {
		return false
	}
	return this.Type.Kind() == reflect.Struct
}

func (this *ExportStruct) ExportFields() string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	var lastIndex = len(this.Fields) - 1
	var t string
	for i, field := range this.Fields {
		var fieldType = getElem(field.Type)
		switch fieldType.Kind() {
		case reflect.Map:

			//fmt.Printf("handle reflect.Map,%v,\n", field)
			//因为template 不支持转义，所以只能这样打印{}
			buf.WriteString("{{`{`}}")
			// counter 是为了消除第一个","  是一个标记func,标记是否执行过了一次，
			t = `{{$c := counter}}{{range $k,$v := .%s}}{{if call $c}},{{end}}"{{$k}}":{{$v}}{{end}}`
			buf.WriteString(fmt.Sprintf(t, field.Name))
			buf.WriteString("{{`}`}}")
			//fmt.Println(buf)

		case reflect.Slice, reflect.Array:
			elemType := getElem(fieldType.Elem())
			if isNumberType(elemType) || elemType.Kind() == reflect.Bool {
				t = `[{{range $i,$v := .%s}}{{if $i}},{{end}}{{$v}}{{end}}]`
				buf.WriteString(fmt.Sprintf(t, field.Name))
			} else if elemType.Kind() == reflect.String {
				t = `[{{range $i,$v := .%s}}{{if $i}},{{end}}"{{$v}}"{{end}}]`
				buf.WriteString(fmt.Sprintf(t, field.Name))
			} else if elemType.Kind() == reflect.Slice { //支持[][]int
				t = `[{{range $i,$v := .%s}}{{if $i}},{{end}}[{{range $j,$k := $v}}{{if $j}},{{end}}{{$k}}{{end}}]{{end}}]`
				buf.WriteString(fmt.Sprintf(t, field.Name))
			}
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Bool:
			buf.WriteString(fmt.Sprintf(`{{.%s}}`, field.Name))
		case reflect.String:
			buf.WriteString(fmt.Sprintf("\"{{.%s}}\"", field.Name))
		}
		if field.DeepStruct != nil {
			if field.Type.Kind() == reflect.Struct {
				buf.WriteString(fmt.Sprintf(`{{with .%s}}%s{{end}}`, field.Name, field.DeepStruct.ExportFields()))
			} else {
				t = `[{{$p2 := pointer .%s}}{{range .%s}}{{$p2.Next}}%s{{if not $p2.IsLast}},{{end}}{{end}}]`
				buf.WriteString(fmt.Sprintf(t, field.Name, field.Name, field.DeepStruct.ExportFields()))
			}
		}
		if i != lastIndex {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']')
	return buf.String()
}

func (this *ExportStructs) CollectStruct(objT reflect.Type) (*ExportStruct, error) {
	objT = getStructElem(objT)
	if objT == nil {
		return nil, getErrorInfo("must pass a struct type", objT)
	}
	if st, ok := this.structs[objT.Name()]; ok {
		return st, nil
	}
	var exportStruct = NewExportStruct(objT.Name())
	for i := 0; i < objT.NumField(); i++ {
		fieldT := objT.Field(i)
		tag := fieldT.Tag.Get("client")
		if len(tag) == 0 {
			continue
		}
		strs := strings.Split(tag, ",")
		exportField := NewExportField(fieldT.Name)
		exportField.ExportName = strs[0]
		exportField.Type = fieldT.Type
		for i := 1; i < len(strs); i++ {
			if strs[i] == "deep" {
				exportField.IsDeep = true
			}
		}
		exportStruct.Fields = append(exportStruct.Fields, exportField)

		var deepType reflect.Type
		switch fieldT.Type.Kind() {
		case reflect.Map, reflect.Slice, reflect.Array, reflect.Ptr:
			if isStructType(fieldT.Type.Elem()) {
				deepType = fieldT.Type.Elem()
			}
		case reflect.Struct:
			deepType = fieldT.Type
		}
		if deepType != nil {
			st, err := this.CollectStruct(deepType)
			if err != nil {
				return nil, err
			}
			exportField.DeepStruct = st
		}
	}
	this.structs[objT.Name()] = exportStruct
	return exportStruct, nil
}

func (this *ExportStructs) GetSortedStructNames() []string {
	var names = make([]string, 0)
	for name := range this.structs {
		names = append(names, name)
	}
	sort.Sort(sort.StringSlice(names))
	return names
}

func (this *ExportStructs) GetSortedStructs() []*ExportStruct {
	names := this.GetSortedStructNames()
	structs := make([]*ExportStruct, 0)
	for _, name := range names {
		structs = append(structs, this.GetStruct(name))
	}
	return structs
}

func (this *ExportStructs) GetStruct(name string) *ExportStruct {
	return this.structs[name]
}
