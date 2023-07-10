package main

import (
	"bytes"
	//"fmt"
	"path"
	"reflect"
	"strings"
	"text/template"
)

var exportTemplate = template.New("codegen")

type RangePointer struct {
	Len   int
	Value int
}

func (this *RangePointer) Next() string {
	this.Value++
	return ""
}

func (this *RangePointer) IsFirst() bool {
	return this.Value == 1
}

func (this *RangePointer) IsLast() bool {
	return this.Value == this.Len
}

func GetRangePointer(obj interface{}) *RangePointer {
	return &RangePointer{
		Len: reflect.ValueOf(obj).Len(),
	}
}

// 根据一个变量导出其需要出现在模板中的字段列表如[{{.Id}},{{.Name}},[{{range .Properties}}[{{.K}}, {{.V}}]{{end}}]]
func TemplateExportFields(field *CollectionField) string {
	st := field.ExportStruct
	if st == nil {
		return "[]"
	}
	return st.ExportFields()
}

func counter() func() int {
	i := -1
	return func() int {
		i++
		return i
	}
}

func ToUpper(v string) string {
	return strings.ToUpper(v)
}

func InitTemplate(rootPath string, gameDb interface{}) error {
	var t = exportTemplate
	t.Funcs(template.FuncMap{
		"exportFields": TemplateExportFields,
		"pointer":      GetRangePointer,
		"counter":      counter,
		"toupper":      ToUpper,
	})
	_, err := t.Delims("$$", "$$").ParseFiles(path.Join(rootPath, "templates/data.tmpl"))
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	for _, group := range exportGroups {
		buf.Reset()
		err = t.ExecuteTemplate(&buf, "data.tmpl", group)
		if err != nil {
			return err
		}
		//fmt.Println(buf.String());
		_, err = t.New(group.Name).Delims("{{", "}}").Parse(buf.String())
		if err != nil {
			return err
		}
	}
	_, err = t.Delims("{{", "}}").ParseFiles(path.Join(rootPath, "templates/decl.tmpl"))
	if err != nil {
		return err
	}
	_, err = t.Delims("{{", "}}").ParseFiles(path.Join(rootPath, "templates/def.tmpl"))
	return err
}
