package util

import (
	"cqserver/golibs/logger"
	"runtime/debug"
	"strings"
	"time"
)

func SafeRun(fun func()) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("SafeRun panic:", time.Now(), err, string(debug.Stack()))
		}
	}()

	fun()
}

//将' " \ -- 替换为空, 用于插入数据库之前的字符过滤
var replacer *strings.Replacer

func init() {
	replacer = strings.NewReplacer("'", "", `"`, "", "--", "", "\\", "")
}

func ReplaceNull(str string) string {
	return replacer.Replace(str)
}
