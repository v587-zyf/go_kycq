package main

import (
	"bufio"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/robots/conf"
	"cqserver/robots/internal/manager"
	"flag"
	"fmt"
	"os"
	"runtime"
)

var m *manager.Manager
var declareFlag uint32
var (
	confFile = flag.String("conf", "./robots.conf", "specify config file")
	robotName = flag.String("robotName", "robot_1", "specify config file")
)

func main() {
	logger.Init()
	defer logger.Close()
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := conf.Conf.Load(*confFile)
	if err != nil {
		logger.Error("program start error when loading config: %s", err.Error())
		return
	}
	m := manager.Get()
	m.Init(*robotName)
	m.Start()

	util.WaitForTerminate()
	logger.Info("robots stopping......")
	m.Stop()
}

func ReadStdin() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Read Stdin Error: %v\n", err)
		return ""
	}
	return str
}
