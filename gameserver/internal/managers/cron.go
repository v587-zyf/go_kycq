package managers

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gamelibs/rmodel"
	"cqserver/gamelibs/rmodelCross"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/protobuf/pb"
	"fmt"
	"runtime/debug"
	"time"

	"cqserver/golibs/logger"

	"github.com/robfig/cron"
	"math/rand"

	"cqserver/golibs/util"
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
	oaRandSecond := rand.Intn(90)
	oaRandSecond = oaRandSecond % 60
	randM := 15 + rand.Intn(5)

	dayRankSendRewardTime := gamedb.GetConf().DayRankSendRewardTime
	dailyPackResetTime := gamedb.GetConf().Dailypack
	spendRebateResetTime := gamedb.GetConf().SpendrebatesRefresh
	firstRechargeResetTime := gamedb.GetConf().FirstRechargeRefresh
	shabakeOpenTime := gamedb.GetConf().ShabakeTime3
	openAwardTime := gamedb.GetConf().RewardTime
	cfg := gamedb.GetCrossArenaTimeCrossArenaTimeCfg(1)

	mainNodejobs := []CronJob{
		{"0 0 0 * * ?", this.DailyMidnight, "DailyMidnight"}, //定时每天0点执行
		{"0 0 1 * * ?", this.DayOneCron, "DayOneCron"},       //定时清理表里过期数据
		{"0 0 5 * * ?", this.Hour5Reset, "Hour5Reset"},       //定时每天5点重围
		{"0 0 23 * * ?", this.Hour23Reset, "Hour23Reset"},    //定时每天23点
		{"0 0 */1 * * ?", this.OneHour, "OneHour"},
		{fmt.Sprintf("0 %d 0 * * ?", randM), this.ZeroMinutes, "15-20minutes"},
		{"0 */5 * * * ?", this.FiveMinutes, "FiveMinutes"},
		{"0 */1 * * * ?", this.OneMinute, "OneMinute"},
		{"*/5 * * * * ?", this.fiveSecond, "fiveSecond"},
		{"*/30 * * * * ?", this.Second30, "Second30"}, // 每30秒
		{"*/3 * * * * ?", this.Second3, "Second3"},    // 每3秒 拍卖行专用
		{fmt.Sprintf("%d %d %d * * ?", dayRankSendRewardTime.Second, dayRankSendRewardTime.Minute, dayRankSendRewardTime.Hour), m.DailyRank.SendEndMail, "dayRankSendRewardTime"}, //每日排行发奖
		{fmt.Sprintf("%d %d %d * * ?", dailyPackResetTime.Second, dailyPackResetTime.Minute, dailyPackResetTime.Hour), this.DailyPack, "dailyPackReset"},                          //每日礼包
		{fmt.Sprintf("%d %d %d * * ?", spendRebateResetTime.Second, spendRebateResetTime.Minute, spendRebateResetTime.Hour), this.SpendRebate, "spendRebateReset"},                //消费返利
		{fmt.Sprintf("%d %d %d * * ?", firstRechargeResetTime.Second, firstRechargeResetTime.Minute, firstRechargeResetTime.Hour), this.FirstRecharge, "firstRechargeReset"},      //首充
		{fmt.Sprintf("%d %d %d * * *", 0, cfg.SignUpEndTime.Minute+gamedb.GetConf().CrossArenaGroping/60, cfg.SignUpEndTime.Hour), this.Challenge, "SignUpEndTime"},
		{fmt.Sprintf("%d %d %d * * *", 1, cfg.SignUpBeginTime.Minute, cfg.SignUpBeginTime.Hour), this.NewChallengeOpen, "SignUpBeginTime"},
		{fmt.Sprintf("%d %d %d * * *", shabakeOpenTime[0].Second, shabakeOpenTime[0].Minute, shabakeOpenTime[0].Hour), this.NewShaBakeOpen, "shabakeOpen"},
		{fmt.Sprintf("%d %d %d * * *", openAwardTime[0].Second, openAwardTime[0].Minute, openAwardTime[0].Hour), this.LotteryOpenAward, "lotteryOpenAwardTime"},
	}

	crontab := cron.New()
	for _, job := range mainNodejobs {
		err := crontab.AddFunc(job.spec, this.wrap(job.name, job.worker))
		if err != nil {
			return err
		}
	}

	crontab.Start()

	go util.SafeRun(this.OneMinuteMore)
	go util.SafeRun(this.FiveMinuteMore)
	go util.SafeRun(this.checkIsSendHour0Mail)
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

// 每天0点执行
func (this *CronManager) DailyMidnight() error {
	logger.Info("执行0点任务开始")
	this.hour0SendMail()
	m.UserManager.TimingUpdate(0)

	m.UserManager.GetAllOnlineUserInfo()

	m.GetWorldLeader().Reset()

	//用来处理 展示排行期间模块战力提升 不加战力
	m.GetDailyRank().ResAddState()
	//远古首领归属者
	m.GetAncientBoss().ResetAncientBossOwner()

	m.GetGuild().DelRobotGuild()

	logger.Info("执行0点任务完成")
	return nil
}

// 每天23点执行
func (this *CronManager) Hour23Reset() error {
	logger.Info("每日23点任务开始")

	logger.Info("每日23点任务完成")
	return nil
}

func (this *CronManager) Hour5Reset() error {
	logger.Info("执行5点任务开始")

	m.UserManager.TimingUpdate(5)

	logger.Info("执行5点任务完成")
	return nil
}

func (this *CronManager) hour0SendMail() {
	// 每日任务 发送邮件奖励
	m.GetDailyTask().SendReward()
	//竞技场发放赛季奖励
	m.GetCompetitve().SendSeasonEndReward()
	//野战记录 7天活跃 玩家当前战力
	m.GetFieldFight().FiveAmRecordUserCombat()
	//寻龙探宝阶段奖励发放
	m.GetTreasure().SendTreasureMail()
	//摇彩奖励发送
	m.GetLottery().SendReward()
	//试炼之路
	m.GetTrialTask().SendReward()
	_ = rmodel.User.SendHour0MailState(time.Now().Day())
	//试炼塔排行奖励
	m.Tower.SendRankReward()
}

// 保证5点出现意外时，玩家排名奖励可以正常领取
func (this *CronManager) checkIsSendHour0Mail() {
	if time.Now().Hour() >= 5 {
		day, _ := rmodel.User.GetHour0MailState()
		if int(day) != time.Now().Day() {
			this.hour0SendMail()
			_ = rmodel.User.SendHour0MailState(time.Now().Day())
		}
	}
}

func (this *CronManager) FiveMinutes() error {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("cronjob FiveMinutes panic: %v, %s", err, debug.Stack())
		}
	}()
	kyEvent.OnlineTotal(m.GetUserManager().GetOnlineTotal())
	return nil
}

func (this *CronManager) OneMinute() error {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("cronjob OneMinute panic: %v, %s", err, debug.Stack())
		}
	}()
	m.UserManager.CronSubscribe()
	return nil
}

func (this *CronManager) OneMinuteMore() {

	defer func() {
		if err := recover(); err != nil {
			logger.Error("cronjob OneMinuteMore panic: %v, %s", err, debug.Stack())
		}
	}()
}

// 随机5-6分钟更新一次公共配置
func (this *CronManager) FiveMinuteMore() {

	defer func() {
		if err := recover(); err != nil {
			logger.Error("cronjob OneMinuteMore panic: %v, %s", err, debug.Stack())
		}
	}()
	randSecond := 300 + rand.Intn(60)
	timer := time.NewTicker(time.Duration(randSecond) * time.Second)
	for {
		select {
		case <-timer.C:
			rmodelCross.GetSystemSeting().CronUpdate()
		}
	}
}

func (this *CronManager) fiveSecond() error {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("cronjob fiveSecond panic: %v, %s", err, debug.Stack())
		}
	}()

	return nil
}

func (this *CronManager) OneHour() error {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("cronjob OneHour panic: %v, %s", err, debug.Stack())
		}
	}()
	return nil
}

func (this *CronManager) ZeroMinutes() error {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("cronjob OneHour panic: %v, %s", err, debug.Stack())
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

// 3秒-主要拍卖行
func (this *CronManager) Second3() error {
	//回溯错误
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Cron Second3 Panic Error. %v, stack trace: %v", err, string(debug.Stack()))
		}
	}()
	m.Auction.CheckAuctionItemTask()
	return nil
}

// 每日1点任务
func (this *CronManager) DayOneCron() error {
	logger.Info(">>>daily one o'clock task")
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Cron DayOneCron Panic Error. %T", err)
		}
	}()
	nowTs := time.Now().Unix()
	// clear  table
	err := modelGame.GetAuctionItemModel().DeleteExpiredItem(nowTs)
	if err != nil {
		logger.Error("世界拍卖行 GetAuctionItemModel DeleteExpiredItem error: %v", err)
	}

	err = modelGame.GetGuildAuctionItemModel().DeleteExpiredItem(nowTs)
	if err != nil {
		logger.Error("门派拍卖行 GetGuildAuctionItemModel DeleteExpiredItem error: %v", err)
	}

	err = modelGame.GetCardModel().DeleteExpiredItem(nowTs)
	if err != nil {
		logger.Error("寻宝 GetCardModel DeleteExpiredItem error: %v", err)
	}

	err = modelGame.GetTreasureModel().DeleteExpiredItem(nowTs)
	if err != nil {
		logger.Error("寻龙探宝 GetTreasureModel DeleteExpiredItem error: %v", err)
	}

	return nil
}

func (this *CronManager) DailyPack() error {
	logger.Info("执行每日礼包重置任务开始")

	m.UserManager.TimingUpdate(constUser.RESET_DAILYPACK)

	logger.Info("执行每日礼包重置任务完成")
	return nil
}
func (this *CronManager) SpendRebate() error {
	logger.Info("执行消费返利重置任务开始")

	m.UserManager.TimingUpdate(constUser.RESET_SPENDREBATE)

	logger.Info("执行消费返利重置任务完成")
	return nil
}
func (this *CronManager) FirstRecharge() error {
	logger.Info("执行首充重置任务开始")

	m.UserManager.TimingUpdate(constUser.RESET_FIRSTRECHARGE)

	logger.Info("执行首充重置任务完成")
	return nil
}

func (this *CronManager) Challenge() error {
	logger.Info("擂台赛报名结束 广播")
	m.GetChallenge().Broadcast()
	return nil
}

func (this *CronManager) NewChallengeOpen() error {
	logger.Info("擂台赛新赛季 广播")
	m.GetChallenge().Broadcast()
	return nil
}

func (this *CronManager) NewShaBakeOpen() error {
	m.GetAnnouncement().SendSystemChat(nil, pb.SCROLINGTYPE_SHABAKE_BEGIN, -1, -1)
	return nil
}

// 摇彩开奖
func (this *CronManager) LotteryOpenAward() error {
	m.GetLottery().OpenAward()
	return nil
}
