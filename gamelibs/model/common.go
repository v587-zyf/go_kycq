package model

import (
	"fmt"
	"reflect"
	"strings"
)

const DB_SERVER = "serverdb"
const DB_ACCOUNT = "accountdb"
const DB_LOG = "logdb"

// 获取model中的所有字段，防止select * 返回model中未定义的字段
func GetAllFieldsAsString(obj interface{}) string {
	return GetAllFieldsAsStringWithTableName(obj, "")
}
func GetAllFieldsAsStringWithTableName(obj interface{}, tableName string) string {
	objT := reflect.TypeOf(obj)
	var fields []string
	for i := 0; i < objT.NumField(); i++ {
		fieldT := objT.Field(i)
		tag := fieldT.Tag.Get("db")
		if tag == "" {
			continue
		}
		oneFileName := fmt.Sprintf("`%s`", tag)
		if tableName != "" {
			oneFileName = fmt.Sprintf("%s.`%s`", tableName, tag)
		}
		fields = append(fields, oneFileName)
	}
	return strings.Join(fields, ",")
}
