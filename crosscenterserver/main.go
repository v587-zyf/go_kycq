package main

import (
	"cqserver/crosscenterserver/internal/managers"
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	//"os"

	"cqserver/crosscenterserver/internal/conf"
	_ "cqserver/crosscenterserver/internal/handler"
	_ "cqserver/gamelibs/modelCross"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/golibs/version"
)

const (
	APP_NAME    = "CrossCenterServer"
)
var (
	APP_VERSION = "1.0.21.12823"
)

var (
	confFile    = flag.String("conf", "./crosscenterserver.conf", "specify config file")
	showVersion = flag.Bool("version", false, "print version string")
	showHelp    = flag.Bool("help", false, "show help")
	genConf     = flag.Bool("genconf", false, "generate an init config file")
	profilePort = flag.Int("profileport", 6067, "specify profile port")
)

func main() {
	logger.Init()
	defer logger.Close()

	logger.Info("**********************************************************")
	logger.Info("                    cqh5 server start")
	logger.Info("                    APP_NAME:%s", APP_NAME)
	logger.Info("					APP_VERSION:%v", APP_VERSION)
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
	go func() {
		if *profilePort > 0 {
			logger.Info("CreateConnection profile: %v", *profilePort)
			err := http.ListenAndServe(fmt.Sprintf("localhost:%d", *profilePort), nil)
			if err != nil {
				logger.Error("profiler start error: %v", err)
				return
			}
			logger.Info("CreateConnection profile ok: %v", *profilePort)
		}
	}()

	var err error
	runtime.GOMAXPROCS(runtime.NumCPU())
	err = conf.Conf.Load(*confFile)
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

	//err = m.Start()
	//if err != nil {
	//	logger.Error("CrossCenterServer CreateConnection error at managers: %s", err.Error())
	//	return
	//}

	m.Run()
	logger.Info("CrossCenterServer start success!")
	util.WaitForTerminate()
	m.Stop()
	logger.Info("CrossCenterServer stopped")
}
