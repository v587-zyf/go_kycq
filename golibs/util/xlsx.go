package util

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

type Decoder interface {
	Decode(str string) error
}

// 从excel中读取数据，并生成为制定的对象数组
// startRow, startCol为开始读的起始行和列，从1开始编号
// groupFinder为寻找名字键值对，可以在excel中填入字符串，但解析到具体对象上自动转化为该字符串映射的整形数值
// 若配置的字段实现了Decoder接口，则自动调用Decode进行解析，这种情况下group参数忽略
// 配置示例:
// type Item struct {
//     Id        int    `col:"id"`
//     Name      string `col:"name"`
//     Desc      string `col:"desc"`
//     Type      int    `col:"type" group:"item_type"`
// }
// 其中col代表excel中的字段名，group代表是否需要根据字符串转化成整型
// 该示例返回的是[]*Item
func ReadXlsxSheet(sheet *xlsx.Sheet, obj interface{}, startRow int, startCol int, groupFinder func(groupName string, fieldName string) (int, error)) ([]interface{}, error) {
	objT := reflect.TypeOf(obj)
	if !(objT.Kind() == reflect.Ptr && objT.Elem().Kind() == reflect.Struct) {
		return nil, errors.New("readSheet must pass a struct type")
	}
	if len(sheet.Rows) <= startRow || len(sheet.Cols) <= startCol {
		return nil, errors.New("empty sheet " + sheet.Name)
	}
	type FieldInfo struct {
		Index   int
		Field   *reflect.StructField
		Group   string
		ColName string
	}

	objT = objT.Elem()
	var colMap = make(map[int]*FieldInfo)
	var maxColumnIndex = 0 //it's the column index of first invalid column
	for i, cell := range sheet.Rows[startRow-1].Cells {
		if i < startCol-1 {
			continue
		} else if cell == nil || len(strings.TrimSpace(cell.Value)) == 0 { // break when meet first empty column
			break
		}
		maxColumnIndex = i
		cellValue := strings.TrimSpace(cell.Value)
		for j := 0; j < objT.NumField(); j++ {
			field := objT.Field(j)
			if field.Tag.Get("col") == cellValue {
				colMap[i] = &FieldInfo{Index: j, Field: &field, Group: field.Tag.Get("group"), ColName: cellValue}
			}
		}
	}
	if len(colMap) == 0 {
		return nil, errors.New("no column found for sheet " + sheet.Name)
	}

	errFunc := func(elem reflect.Type, fieldIndex, i, j int, sheet *xlsx.Sheet, err error) error {
		return fmt.Errorf("field %s at %c%d error for sheet %s: %s", elem.Field(fieldIndex).Name, 'A'+j%26, i+1, sheet.Name, err.Error())
	}
	var result = make([]interface{}, 0)
	for i, row := range sheet.Rows {
		if i < startRow {
			continue
		} else if row == nil || len(row.Cells) == 0 {
			break
		}
		objInstance := reflect.New(objT).Interface()
		objV := reflect.ValueOf(objInstance).Elem()
		for j, cell := range row.Cells {
			if j < startCol-1 {
				continue
			}
			fieldInfo := colMap[j]
			if fieldInfo == nil {
				continue
			}
			cellString := strings.TrimSpace(cell.String())
			if j == startCol-1 && i >= startRow && (cell == nil || len(cellString) == 0) {
				goto exit //finish when meet first empty row (the first column of this row is empty)
			}
			if j > maxColumnIndex {
				break
			}
			fieldV := objV.Field(fieldInfo.Index)
			if !fieldV.CanSet() {
				return nil, fmt.Errorf("field %s can not be set for sheet %s", objT.Field(fieldInfo.Index).Name, sheet.Name)
			}
			if decoder, ok := fieldV.Addr().Interface().(Decoder); ok {
				err := decoder.Decode(cellString)
				if err != nil {
					return nil, errFunc(objT, fieldInfo.Index, i, j, sheet, err)
				}
				continue
			}
			if len(cellString) == 0 {
				continue
			}
			switch objT.Field(fieldInfo.Index).Type.Kind() {
			case reflect.Bool:
				if cellString == "1" {
					fieldV.SetBool(true)
				} else if cellString == "0" {
					fieldV.SetBool(false)
				} else {
					b, err := strconv.ParseBool(cellString)
					if err != nil {
						return nil, errFunc(objT, fieldInfo.Index, i, j, sheet, err)
					}
					fieldV.SetBool(b)
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if len(fieldInfo.Group) > 0 && groupFinder != nil {
					x, err := groupFinder(fieldInfo.Group, cellString)
					if err != nil {
						return nil, errFunc(objT, fieldInfo.Index, i, j, sheet, err)
					}
					fieldV.SetInt(int64(x))
				} else {
					x, err := strconv.ParseInt(strings.Split(cellString, ".")[0], 10, 64) //需防止自动计算字段为float类型
					if err != nil {
						return nil, errFunc(objT, fieldInfo.Index, i, j, sheet, err)
					}
					fieldV.SetInt(x)
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				x, err := strconv.ParseUint(cellString, 10, 64)
				if err != nil {
					return nil, errFunc(objT, fieldInfo.Index, i, j, sheet, err)
				}
				fieldV.SetUint(x)
			case reflect.Float32, reflect.Float64:
				x, err := strconv.ParseFloat(cellString, 64)
				if err != nil {
					return nil, errFunc(objT, fieldInfo.Index, i, j, sheet, err)
				}
				fieldV.SetFloat(x)
			case reflect.String:
				fieldV.SetString(cellString)
			}
		}
		result = append(result, objInstance)
	}
exit:
	return result, nil
}
