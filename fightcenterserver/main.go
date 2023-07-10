package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	//"os"

	"cqserver/fightcenterserver/internal/base"
	"cqserver/fightcenterserver/internal/managers"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/golibs/version"
)

const (
	APP_NAME    = "fightcenterserver"
)
var (
	APP_VERSION = "1.0.21.12823"
)

var (
	confFile    = flag.String("conf", "./fightcenterserver.conf", "specify config file")
	showVersion = flag.Bool("version", false, "print version string")
	showHelp    = flag.Bool("help", false, "show help")
	genConf     = flag.Bool("genconf", false, "generate an init config file")
	profilePort = flag.Int("profileport", 6166, "specify profile port")
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
	err = base.Conf.Load(*confFile)
	if err != nil {
		logger.Error("program start error when loading config: %s", err.Error())
		return
	}

	m := managers.Get()
	err = m.Init()
	if err != nil {
		logger.Error("program start error: %s", err.Error())
		return
	}
	logger.Info("fightcenterserver started")
	err = m.Start()
	if err != nil {
		logger.Error("fightcenterserver Start error at managers: %s", err.Error())
		return
	}

	m.Run()
	util.WaitForTerminate()
	m.Stop()
	logger.Info("fightcenterserver stopped")
}
