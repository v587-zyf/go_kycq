package managers

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constChallenge"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"fmt"
	"github.com/robfig/cron"
	"runtime/debug"
	"time"
)

type CronJob struct {
	Spec   string
	Worker func() error
	Name   string
}

type CronManager struct {
	util.DefaultModule
}

func NewCronManager() *CronManager {
	return &CronManager{}
}

func (this *CronManager) Init() error {
	logger.Info("Cron init ........................")
	//cfg := gamedb.GetCrossArenaTimeCrossArenaTimeCfg(1)
	mainNodeJobs := []CronJob{
		{"0 0 0 * * ?", this.DayZeroCron, "DayZeroCron"}, //定时每天0点执行
		{"0 0 1 * * ?", this.DayOneCron, "DayOneCron"},   //定时清理表里过期数据
		{"0 0 5 * * ?", this.DailyReset, "DailyReset"},   //定时每天5点执行
		{"*/1 * * * * ?", this.Second1, "Second1"},       //每1秒
		{"*/5 * * * * ?", this.Second5, "Second5"},       //每5秒
		{"*/30 * * * * ?", this.Second30, "Second30"},    //每30秒
		{"0 */1 * * * ?", this.OneMinute, "OneMinute"},   //每分钟
	}

	//擂台赛
	challengeJobs := this.GenCronJobs()
	if len(challengeJobs) > 0 {
		mainNodeJobs = append(mainNodeJobs, challengeJobs...)
	}

	cronTab := cron.New()
	for _, job := range mainNodeJobs {
		err := cronTab.AddFunc(job.Spec, this.wrap(job.Name, job.Worker))
		if err != nil {
			return err
		}
	}

	cronTab.Start()
	logger.Info("Cron CreateConnection ........................")
	return nil
}

func (this *CronManager) wrap(name string, f func() error) func() {
	return func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("cronJob %s panic: %v, %s", name, err, debug.Stack())
			}
		}()
		err := f()
		if err != nil {
			logger.Error("cronJob %s error: %s", name, err.Error())
		}
	}
}

//0点任务
func (this *CronManager) DayZeroCron() error {
	logger.Info("day zero cron")
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Cron DayZeroCron Panic Error. %T", err)
		}
	}()
	m.ServerListManager.CrossMatch()
	return nil
}

//每日1点任务
func (this *CronManager) DayOneCron() error {
	logger.Info(">>>daily one o'clock task")
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Cron DayOneCron Panic Error. %T", err)
		}
	}()
	nowTs := time.Now().Unix()
	// clear  table
	err := modelCross.GetChallengeModel().DeleteExpiredItem(nowTs)
	if err != nil {
		logger.Error("GetChallengeModel DeleteExpiredItem error: %v", err)
	}

	err = modelCross.GetChallengeDataModel().DeleteExpiredItem(nowTs)
	if err != nil {
		logger.Error("GetChallengeDataModel DeleteExpiredItem error: %v", err)
	}
	return nil
}

//5点任务
func (this *CronManager) DailyReset() error {
	logger.Info("-----DailyReset--------")
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Cron DailyReset Panic Error. %T", err)
		}
	}()
	return nil
}

// 1秒-主要是同步数据
func (this *CronManager) Second1() error {
	//回溯错误
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Cron Second1 Panic Error. %T", err)
		}
	}()

	return nil
}

// 5秒-主要是同步数据
func (this *CronManager) Second5() error {
	//回溯错误
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Cron Second5 Panic Error. %T", err)
		}
	}()

	return nil
}

// 30秒-主要是打印日志
func (this *CronManager) Second30() error {
	//回溯错误
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Cron Second30 Panic Error. %T", err)
		}
	}()

	return nil
}

//每分钟执行一次
func (this *CronManager) OneMinute() error {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Cron OneMinute Panic Error. %T", err)
		}
	}()
	return nil
}

func (this *CronManager) GenCronJobs() []CronJob {
	jobs := make([]CronJob, 0)
	cfg := gamedb.GetCrossArenaTimeCrossArenaTimeCfg(1)
	cronChallenge := fmt.Sprintf("%d %d %d * * *", 6, 6, 6)
	logger.Info("擂台赛 创建定时任务 提前插入假人:%v", cronChallenge)
	jobs = append(jobs, CronJob{Spec: cronChallenge, Worker: this.genChallengeJob1(), Name: fmt.Sprintf("跨服擂台赛分组")})

	for num := constChallenge.Group; num <= constChallenge.EIGHTH_ROUND; num++ {
		if num == constChallenge.Group {
			cronChallenge = fmt.Sprintf("%d %d %d * * *", cfg.SignUpEndTime.Second, cfg.SignUpEndTime.Minute, cfg.SignUpEndTime.Hour)
			logger.Info("num:%v 擂台赛 创建定时任务 cronChallenge:%v", num, cronChallenge)
			jobs = append(jobs, CronJob{Spec: cronChallenge, Worker: this.genChallengeJob(num), Name: fmt.Sprintf("跨服擂台赛分组")})
		} else {
			cfg1 := gamedb.GetCrossArenaCrossArenaCfg(num)
			cronChallenge = fmt.Sprintf("%d %d %d * * *", cfg1.IntervalTime.Second, cfg1.IntervalTime.Minute, cfg1.IntervalTime.Hour)
			logger.Info("num:%v 擂台赛 创建定时任务 cronChallenge:%v", num, cronChallenge)
			jobs = append(jobs, CronJob{Spec: cronChallenge, Worker: this.genChallengeJob(num), Name: fmt.Sprintf("跨服擂台赛分组")})
		}
	}

	return jobs
}

func (this *CronManager) genChallengeJob(stage int) func() error {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("cronjob genChallengeJob panic: %v, %s", err, debug.Stack())
		}
	}()

	return func() error {
		if stage == constChallenge.Group {
			m.ChallengeCcs.BeginGroup(stage)
		} else {
			m.ChallengeCcs.InSeasonCompetition(stage)
		}
		return nil
	}
}

func (this *CronManager) genChallengeJob1() func() error {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("cronjob genChallengeJob panic: %v, %s", err, debug.Stack())
		}
	}()

	return func() error {
		m.GetCcsChallenge().BeginInsertRobotUser()
		return nil
	}
}
