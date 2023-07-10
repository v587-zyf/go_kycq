package common

import (
	"strconv"
	"strings"
)

/*===========================================================================*/
/*=================================字符串转换=================================*/
/*===========================================================================*/
func Str2Int(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
func Str2Int32(str string) int32 {
	i, _ := strconv.Atoi(str)
	return int32(i)
}
func Str2Int64(str string) int64 {
	i, _ := strconv.Atoi(str)
	return int64(i)
}
func Str2Bool(str string) bool {
	b, _ := strconv.ParseBool(str)
	return b
}
func Str2Float64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 32)
	return f
}
func Str2Int32Arr(str string) []int32 {
	strArr := strings.Split(str, ",")
	int32Slice := make([]int32, len(strArr))
	for _, s := range strArr {
		int32Slice = append(int32Slice, Str2Int32(s))
	}
	return int32Slice
}

/*===========================================================================*/
/*================================interface转换==============================*/
/*===========================================================================*/
func Interface2Int(i interface{}) int {
	return int(i.(int64))
}
func Interface2Str(i interface{}) string {
	return i.(string)
}
func Interface2Int32(i interface{}) int32 {
	switch i.(type) {
	case int:
		return int32(i.(int))
	case int64:
		return int32(i.(int64))
	case float64:
		return int32(i.(float64))
	}
	return 0
}
func Interface2Bool(i interface{}) bool {
	return i.(bool)
}
