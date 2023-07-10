package main

import (
	"cqserver/gateserver/conf"
	"cqserver/gateserver/internal/manager"
	"flag"
	"fmt"

	_ "cqserver/gateserver/internal/handler"
	_ "net/http/pprof"
	"os"
	"runtime"

	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/golibs/version"
)

const (
	APP_NAME    = "yulong_gateserver"
)
var (
	APP_VERSION = "1.0.21.12823"
)

var (
	confFile    = flag.String("conf", "./robots.conf", "specify config file")
	showVersion = flag.Bool("version", false, "print version string")
	showHelp    = flag.Bool("help", false, "show help")
	genConf     = flag.Bool("genconf", false, "generate an init config file")
	profilePort = flag.Int("profileport", 6062, "specify profile port")

	ENV_TEST       = os.Getenv("YULONG_TEST")
	GATE_TEST      = len(ENV_TEST) > 0 // 是否启用压力测试
	ENABLE_FS_TEST = len(ENV_TEST) > 1 // 如果是两位数，则代表启用fightserver, 否则只启用gameserver
)

func main() {
	logger.Init()
	defer logger.Close()

	logger.Info("**********************************************************")
	logger.Info("                    cqh5 server start")
	logger.Info("                    APP_NAME:%s", APP_NAME)
	logger.Info("                    APP_VERSION:%v", APP_VERSION)
	logger.Info("**********************************************************")
	logger.Info("conf:%v", *confFile)
	flag.Parse()
	if *showHelp {
		flag.Usage()
		return
	}
	if *showVersion {
		fmt.Println(version.VersionDetail(APP_NAME, APP_VERSION))
		return
	}
	if *genConf {
		fmt.Println(conf.InitConfStr)
		return
	}

	sysType := runtime.GOOS
	if sysType == "windows" {
		os.Setenv("ZONEINFO", ".\\zoneinfo.zip")
	}

	var err error
	runtime.GOMAXPROCS(runtime.NumCPU())
	err = conf.Conf.Load(*confFile)
	if err != nil {
		logger.Error("program start error when loading config: %s", err.Error())
		return
	}
	if !conf.Conf.Sandbox {
		GATE_TEST = false
	}

	m := manager.Get()
	err = m.Init()
	if err != nil {
		logger.Error("program start error when init: %v", err.Error())
		return
	}
	logger.Info("GateServer started")
	m.Run()
	util.WaitForTerminate()
	m.Stop()
	logger.Info("GateServer stopped")
}
