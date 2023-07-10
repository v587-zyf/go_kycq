package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	//"os"
	"runtime"
	"time"

	_ "cqserver/gamelibs/retry"
	"cqserver/gameserver/internal/base"
	_ "cqserver/gameserver/internal/handler"
	//"cqserver/golibs/logger"
	"cqserver/gameserver/internal/managers"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/golibs/version"
)

const (
	APP_NAME = "gameserver"
)

var (
	APP_VERSION = "1.0.21.12823"
)

var (
	confFile    = flag.String("conf", "./gameserver.conf", "specify config file")
	showVersion = flag.Bool("version", false, "print version string")
	showHelp    = flag.Bool("help", false, "show help")
	genConf     = flag.Bool("genconf", false, "generate an init config file")
	profilePort = flag.Int("profileport", 6065, "specify profile port")
	isTest      = flag.Bool("isTest", false, "validate gamedb")
)

func main() {

	//sysType := runtime.GOOS
	//if(sysType == "windows"){
	//	os.Setenv("ZONEINFO", ".\\zoneinfo.zip")
	//}

	logger.Init()
	defer logger.Close()

	logger.Info("**********************************************************")
	logger.Info("                    cqh5 server start")
	logger.Info("                    APP_NAME:%s", APP_NAME)
	logger.Info("                    APP_VERSION:%s", APP_VERSION)
	logger.Info("**********************************************************")
	logger.Info("conf:%v", *confFile)

	//解析定义的flag参数
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
		fmt.Println(base.InitConfStr)
		return
	}
	go func() {
		if *profilePort > 0 {
			logger.Info("Start profile: %v", *profilePort)
			err := http.ListenAndServe(fmt.Sprintf("localhost:%d", *profilePort), nil)
			if err != nil {
				logger.Error("profiler start error: %v", err)
				return
			}
			logger.Info("Start profile ok: %v", *profilePort)
		}
	}()

	var err error
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())
	err = base.Conf.Load(*confFile)
	if err != nil {
		logger.Error("program start error when loading config: %s", err.Error())
		return
	}

	m := managers.Get()
	m.GameVesion = APP_VERSION
	m.IsTest = *isTest
	err = m.Init()
	if err != nil {
		logger.Error("GameServer Init error at managers: %s", err.Error())
		return
	}
	err = m.Start()
	if err != nil {
		logger.Error("GameServer Start error at managers: %s", err.Error())
		return
	}
	m.Run()
	logger.Info("GameServer started")
	util.WaitForTerminate()
	m.Stop()
	logger.Info("GameServer stopped")
}
