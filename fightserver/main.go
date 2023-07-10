package main

import (
	"cqserver/fightserver/conf"
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"

	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/golibs/version"
)

const (
	APP_NAME    = "fightserver"
)

var (
	APP_VERSION = "1.0.21.12823"
)

var (
	confFile       = flag.String("conf", "./fightserver.conf", "specify config file")
	showVersion    = flag.Bool("version", false, "print version string")
	gameDbBasePath = flag.String("gamedb", "../../yulong/data/configs", "specify gamedb path")
	genConf        = flag.Bool("genconf", false, "generate an init config file")
	profilePort    = flag.Int("profileport", 6066, "specify profile port")
)

func main() {

	//sysType := runtime.GOOS
	//if(sysType == "windows"){
	//	os.Setenv("ZONEINFO", ".\\zoneinfo.zip")
	//}

	logger.Init()
	defer func() {
		defer logger.Close()
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic main:%v,%s", r, stackBytes)
		}
	}()

	logger.Info("**********************************************************")
	logger.Info("                    cqh5 server start")
	logger.Info("                    APP_NAME:%s", APP_NAME)
	logger.Info("                    APP_VERSION:%v", APP_VERSION)
	logger.Info("**********************************************************")
	logger.Info("conf:%v", *confFile)

	flag.Parse()
	if *showVersion {
		fmt.Println(version.VersionDetail(APP_NAME, APP_VERSION))
		return
	}
	if *genConf {
		fmt.Println(conf.InitConfStr)
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
	err = conf.Conf.Load(*confFile)
	if err != nil {
		logger.Error("program start error when loading config: %s", err.Error())
		return
	}

	err = m.Init()
	if err != nil {
		logger.Error("program start error: %s", err.Error())
		return
	}
	m.Run()
	logger.Info("fightServer started")
	util.WaitForTerminate()
	m.Stop()
	logger.Info("fightserver stopped")
}
