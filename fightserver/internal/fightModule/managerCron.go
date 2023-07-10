package fightModule

import (
	"cqserver/gamelibs/rmodelCross"
	"runtime/debug"

	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/robfig/cron"
	"math/rand"
	"time"
)

type CronJob struct {
	spec   string
	worker func() error
	name   string
}

type CronManager struct {
	util.DefaultModule
}

func NewCronManager() *CronManager {
	return &CronManager{}
}

func (this *CronManager) Init() error {
	mainNodejobs := []CronJob{
		//	{"*/30 * * * * ?", this.syncStatus, "SyncStatus"}, // 每30秒向watcher服务监控中心，发送同步信息
		{"0 */1 * * * ?", this.OneMinute, "OneMinute"},
		//	{"0 0 1 * * ?", this.OneHour, "OneHour"}, //定时每天1点重围
		//
	}

	crontab := cron.New()
	for _, job := range mainNodejobs {
		err := crontab.AddFunc(job.spec, this.wrap(job.name, job.worker))
		if err != nil {
			return err
		}
	}

	crontab.Start()
	go util.SafeRun(this.FiveMinuteMore)
	return nil
}

func (this *CronManager) OneMinute() error {
	fightManager.InitGameServers()
	return nil
}

//随机5-6分钟更新一次公共配置
func (this *CronManager) FiveMinuteMore() {

	defer func() {
		if err := recover(); err != nil {
			log.Error("cronjob OneMinuteMore panic: %v, %s", err, debug.Stack())
		}
	}()
	rand.Seed(time.Now().UnixNano())
	randSecond := 300 + rand.Intn(60)
	timer := time.NewTicker(time.Duration(randSecond) * time.Second)
	for {
		select {
		case <-timer.C:
			rmodelCross.GetSystemSeting().CronUpdate()
		}
	}
}

func (this *CronManager) syncStatus() error {

	return nil
}

func (this *CronManager) wrap(name string, f func() error) func() {
	return func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("cronjob %s panic: %v, %s", name, err, debug.Stack())
			}
		}()
		err := f()
		if err != nil {
			logger.Error("cronjob %s error: %s", name, err.Error())
		}
	}
}
