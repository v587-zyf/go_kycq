package util

import (
	"cqserver/golibs/logger"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"sync"
	"syscall"
	"time"
)

var log = logger.Get("default", true)

type Module interface {
	Init() error
	Start() error
	Run()
	Stop()
	ModifyServerOpenTime() // 修改开服时间
	MuCheck() error
}

type DefaultModule struct {
}

func (this DefaultModule) Init() error {
	return nil
}

func (this DefaultModule) Start() error {
	return nil
}

func (this DefaultModule) Run() {

}

func (this DefaultModule) Stop() {

}

func (this DefaultModule) MuCheck() error {
	return nil
}

func (this DefaultModule) ModifyServerOpenTime() {

}

// DefaultModuleManager default module manager
type DefaultModuleManager struct {
	Module
	Modules []Module
}

func NewDefaultModuleManager() *DefaultModuleManager {
	return &DefaultModuleManager{
		Modules: make([]Module, 0, 5),
	}
}

func (this *DefaultModuleManager) Init() error {
	for i := 0; i < len(this.Modules); i++ {
		clsName := fmt.Sprintf("%T", this.Modules[i])
		dotIndex := strings.Index(clsName, ".") + 1
		log.Info(clsName[dotIndex:len(clsName)] + " Init")
		err := this.Modules[i].Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *DefaultModuleManager) Start() error {
	for i := 0; i < len(this.Modules); i++ {
		err := this.Modules[i].Start()
		if err != nil {
			return err
		}
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("mucheck panic:", time.Now(), err, string(debug.Stack()))
			}
		}()
		ticker := time.NewTicker(5 * time.Minute)
		for {
			select {
			case <-ticker.C:
				logger.Info("mucheck start")
				for i := 0; i < len(this.Modules); i++ {
					clsName := fmt.Sprintf("%T", this.Modules[i])
					dotIndex := strings.Index(clsName, ".") + 1
					log.Info(clsName[dotIndex:len(clsName)] + " mucheck start")
					err := this.Modules[i].MuCheck()
					if err != nil {
						logger.Error("mucheck error:%v", err)
						return
					}
					log.Info(clsName[dotIndex:len(clsName)] + " mucheck stop")
				}
				logger.Info("mucheck stop")
			}
		}
	}()

	return nil
}

func (this *DefaultModuleManager) Run() {
	for i := 0; i < len(this.Modules); i++ {
		this.Modules[i].Run()
	}
}

func (this *DefaultModuleManager) Stop() {
	var wg sync.WaitGroup
	for i := 0; i < len(this.Modules); i++ {
		wg.Add(1)
		go func(module Module) {
			module.Stop()
			wg.Done()
		}(this.Modules[i])
	}
	wg.Wait()
}

func (this DefaultModuleManager) MuCheck() error {
	return nil
}

func (this *DefaultModuleManager) ModifyServerOpenTime() {
	for i := 0; i < len(this.Modules); i++ {
		this.Modules[i].ModifyServerOpenTime()
	}
}

func (this *DefaultModuleManager) AppendModule(module Module) Module {
	this.Modules = append(this.Modules, module)
	return module
}

// WaitTerminateSignal wait signal to end the program
func WaitForTerminate() {
	exitChan := make(chan struct{})
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		close(exitChan)
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-exitChan
}
