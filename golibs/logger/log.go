package logger

import (
	"github.com/astaxie/beego/logs"
)

type logger struct {
	*logs.BeeLogger
}

var defaultLog *logger
var defaultLogLevel = logs.LevelInformational

func DeafultLoggerInit() {
	defaultLog = &logger{
		BeeLogger: Get("default", true),
	}
	defaultLog.SetLogFuncCallDepth(3)
	defaultLogLevel = GetLogLevel("default")
}

//定时检查设置log等级
func SetDebugSwitch(openDebug string) {

	if defaultLogLevel == logs.LevelDebug {
		//关闭debug
		if openDebug == "0" {
			Info("关闭log debug")
			defaultLogLevel = logs.LevelInformational
			defaultLog.SetLevel(defaultLogLevel)
		}
	} else {
		//开启debug
		if openDebug == "1" {
			Info("开启log debug")
			defaultLogLevel = logs.LevelDebug
			defaultLog.SetLevel(defaultLogLevel)
		}
	}
}

func Debug(format string, v ...interface{}) {
	defaultLog.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	defaultLog.Info(format, v...)
}

func Warn(format string, v ...interface{}) {

	defaultLog.Warn(format, v...)
}

func Error(format string, v ...interface{}) {

	defaultLog.Error(format, v...)
}