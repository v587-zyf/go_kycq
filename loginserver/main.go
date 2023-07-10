package main

import (
	"cqserver/loginserver/conf"
	"cqserver/loginserver/internal/manager"
	_ "cqserver/loginserver/internal/handler"
	"flag"
	"fmt"

	//"net/http"
	_ "net/http/pprof"
	//"os"
	"runtime"

	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/golibs/version"
)

const (
	APP_NAME = "loginserver"
)

var (
	APP_VERSION = "1.0.21.12823"
)

var (
	confFile    = flag.String("conf", "./loginserver.conf", "specify config file")
	showVersion = flag.Bool("version", false, "print version string")
	showHelp    = flag.Bool("help", false, "show help")
	profilePort = flag.Int("profileport", 6061, "specify profile port")
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

	var err error
	runtime.GOMAXPROCS(runtime.NumCPU())
	err = conf.Conf.Load(*confFile)
	if err != nil {
		logger.Error("program start error when loading config: %s", err.Error())
		return
	}

	err = manager.Get().Init()
	if err != nil {
		logger.Error("program start error: %s", err.Error())
		return
	}
	logger.Info("LoginServer started")
	manager.Get().Run()
	util.WaitForTerminate()
	manager.Get().Stop()
	logger.Info("LoginServer stopped")
}
