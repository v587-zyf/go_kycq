package common

import (
	"cqserver/golibs/logger"
	"fmt"
	"time"
)

const (
	DAILY_RESET_HOUR     = 5
	DAILY_RESET_HOUR_NEW = 24
)

func SameDay(a, b time.Time) bool {
	return a.Year() == b.Year() && a.YearDay() == b.YearDay()
}

func IsToday(ts int64) bool {
	return SameDay(time.Unix(ts, 0), time.Now())
}

func CheckTimeFormat(src, layout string) bool {
	_, err := time.Parse(layout, src)
	return err == nil
}

//获取时间字符串获取1970时间戳(timestr格式 2006-01-02 15:04:05)
func GetTimeUnix(timeStr string) (int, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return 0, err
	}
	time, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err != nil {
		return 0, err
	}
	return int(time.Unix()), nil
}

func GetTime(timeStr string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	time, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	return time, nil
}

//根据传入时间，计算当前第几天
func GetTheDays(startTime time.Time) int {
	serverOpenTime := startTime
	serverOpenZeroTime := ZeroTimeOfDay(serverOpenTime).Unix()
	nowTime := time.Now()
	nowZeroTime := ZeroTimeOfDay(nowTime).Unix()
	openDays := (nowZeroTime-serverOpenZeroTime)/(24*60*60) + 1
	return int(openDays)
}

//根据传入时间，计算当前第几天 (偏移几小时)
func GetTheDaysReduceHour(startTime time.Time, hour time.Duration) int {
	serverOpenTime := startTime
	serverOpenZeroTime := ZeroTimeOfDay(serverOpenTime).Unix()
	h, _ := time.ParseDuration("-1h")
	nowTime := time.Now().Add(h * hour)
	nowZeroTime := ZeroTimeOfDay(nowTime).Unix()
	openDays := (nowZeroTime-serverOpenZeroTime)/(24*60*60) + 1
	return int(openDays)
}

// 获取传入时间的零点时刻的时间戳
func GetZeroClockTimestamp(t time.Time) int64 {
	ts, _ := time.Parse("2006-01-02", t.Format("2006-01-02"))
	return ts.Unix()
}

// 获取传入时间的日期字符串（不包含时分秒）
func GetDateString(t time.Time) string {
	tStr := t.Format("2006-01-02")
	return tStr
}

func GetFormatTime2(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GetFormatTime() string {
	currentTime := time.Now().Local()
	return currentTime.Format("2006-01-02 15:04:05")
}

func GetDateStringWithoutSpace(t time.Time) string {
	tStr := t.Format("060102")
	return tStr
}

/**
 *  @Description: 以凌晨5点作为日期分割
 *  @param t
 *  @return int
 */
func GetDateIntByOffset5(t time.Time) int {
	y, m, d := t.Add(-time.Hour * DAILY_RESET_HOUR).Date()
	date := y*10000 + int(m)*100 + d
	return date
}

func GetDateInt(t time.Time) int {
	y, m, d := t.Date()
	date := y*10000 + int(m)*100 + d
	return date
}

// 获取传入时间时分秒换算成的秒数
func GetTimeSeconds(t time.Time) int {
	return t.Hour()*60*60 + t.Minute()*60 + t.Second()
}

// 获取传入时间的日期（不包含时分秒）
func GetDate(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Local")
	ts, _ := time.ParseInLocation("2006-01-02", t.Format("2006-01-02"), loc)
	return ts
}

// 获取传入时间的年份和星期组合(202150) 以凌晨5点作为日期分割
func GetYearWeekByOffset5(t time.Time) int {
	year, week := t.Add(-time.Hour * DAILY_RESET_HOUR).ISOWeek()
	date := year*100 + week
	return date
}

// 获取传入时间的年份和星期组合(202150)
func GetYearWeek(t time.Time) int {
	year, week := t.ISOWeek()
	date := year*100 + week
	return date
}

// 传入日期字符串，获取对应的日期时间
func GetTimeWitDateStr(dateStr string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	return time.ParseInLocation("2006-01-02", dateStr, loc)
}

//传入时间戳,获取对应的时间类型
func MakeTimeSe(se int64) time.Time {
	times := time.Unix(se, 0)
	return time.Date(times.Year(), times.Month(), times.Day(), 0, 0, 0, 0, time.Local)
}

//获取今日零点的时间戳
func GetZeroTimeUnixFrom1970() int {
	now := time.Now()
	return int(now.Unix()) - now.Hour()*3600 - now.Minute()*60 - now.Second()
}

func GetNowMillisecond() int64 {
	return time.Now().UnixNano() / 1000000
}

func GetMilliSecondsByString(strs []string) []int {

	millTimes := make([]int, 0)
	for _, v := range strs {

		t, err := IntSliceFromString(v, ":")
		if err != nil || len(t) <= 0 {
			logger.Error("时间字符串异常：%v", strs)
			continue
		}
		h := t[0]
		m := 0
		s := 0
		if len(t) > 1 {
			m = t[1]
		}
		if len(t) > 2 {
			m = t[2]
		}
		millTimes = append(millTimes, h*60*60*1000+m*60*1000+s*1000)

	}
	return millTimes
}

func GetResetTime(t time.Time) int {
	return GetDateInt(t)
}

//获取对应日零点的时间戳
func GetZeroTimeUnix(num int) int {
	now := time.Now().AddDate(0, 0, num)
	return int(now.Unix()) - now.Hour()*3600 - now.Minute()*60 - now.Second()
}

//获取两个时间 相差的天数
func GetNumberOfDaysDifference(beforeTime, afterTime time.Time) int {

	beforeTimeString := beforeTime.Format("2006-01-02")
	afterTimeString := afterTime.Format("2006-01-02")

	a, _ := time.Parse("2006-01-02", beforeTimeString)
	b, _ := time.Parse("2006-01-02", afterTimeString)
	d := b.Sub(a)

	c := d.Hours() / 24
	diff := int(c)
	if diff < 0 {
		diff = -diff
	}
	return diff + 1
}

//
//  CalcEndTime
//  @Description:
//  @param serverOpenTime 服务器开服时间
//  @param activityTimes 活动持续时间
//  @return time.Time
//
func CalcEndTime(serverOpenTime time.Time, activityTimes []int, hour int) int64 {
	openTime := activityTimes
	if openTime == nil || len(activityTimes) < 3 {
		return 0
	}

	addTime := time.Date(serverOpenTime.Year(), serverOpenTime.Month(), serverOpenTime.Day(), hour, 0, 0, 0, serverOpenTime.Location())
	addTime.Format("2006-01-02")
	addHour, _ := time.ParseDuration(fmt.Sprintf(`%dh`, openTime[1]))
	addMinute, _ := time.ParseDuration(fmt.Sprintf(`%dm`, openTime[2]))
	endTime := addTime.AddDate(0, 0, openTime[0]).Add(addHour).Add(addMinute)
	return endTime.Unix()
}
