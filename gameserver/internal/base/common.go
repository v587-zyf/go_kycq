package base

import (
	textcensor "github.com/kai1987/go-text-censor"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"cqserver/gamelibs/gamedb"
)

var Time1970 = time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local)
var nameReg = regexp.MustCompile("^[\u4e00-\u9fa5_a-zA-Z0-9]+$")
var Time2028 = time.Date(2028, 1, 1, 0, 0, 0, 0, time.Local)

const (
	DAILY_RESET_HOUR    = 5
	DAILY_RESET_SECONDS = DAILY_RESET_HOUR * 3600
	TwoWeeksSeconds     = 86400 * 14
	OneDaySeconds       = 86400
)

type AsyncFunc func() error

func RandNum(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func BaseMulti(base int, ratio float64) int {
	return int(math.Floor(float64(base) * ratio))
}

func BaseDiv(base int, ratio float64) int {
	return int(math.Floor(float64(base) / ratio))
}

func AddToMax(src *int, addValue int, max int) {
	*src += addValue
	if *src > max {
		*src = max
	}
}

func DecToMin(src *int, decValue int, min int) {
	*src -= decValue
	if *src < min {
		*src = min
	}
}

//IsAfterResetPoint
//判断一个时间戳（秒）是否在重知点以后
func IsAfterResetPoint(timeStampSeconds int64) bool {
	return IsTmSameDay(time.Now().Unix(), timeStampSeconds, DAILY_RESET_HOUR)
}

//得到一个时间点,从每天5点开始,过去了几天.
func GetPassedDaysAtResetPoint(t time.Time) int {
	return GetPassedDaysAtResetPointViaTs(t.Unix())
}

//得到一个时间点,从每天5点开始,过去了几天.
func GetPassedDaysAtResetPointViaTs(ts int64) int {
	return int(GetPassedDaysVisTs(time.Now().Unix(), ts, DAILY_RESET_HOUR))
}

//GetNowDateWithOffset
//得到带重置时间的当天的YYYYMMHH
func GetNowDateWithOffset() string {
	return GetDateWithOffset(time.Now())
}
func GetDateWithOffset(t time.Time) string {
	return t.Add(-time.Hour * DAILY_RESET_HOUR).Format("060102")
}

func GetNowDateWithOffsetByUnix() int64 {
	return time.Now().Add(-time.Hour * DAILY_RESET_HOUR).Unix()
}

func GetDateWithOffsetByUnix(t time.Time) int64 {
	return t.Add(-time.Hour * DAILY_RESET_HOUR).Unix()
}

func GetDayBySecond(day int) int {
	return day * 60 * 60 * 24
}

func GetDifferDay(timeUnix, timeUnix1 int64) int {
	return int((timeUnix1 - timeUnix) / int64(GetDayBySecond(1)))
}

func IsSameDay(time1 time.Time, time2 time.Time) bool {
	return IsTmSameDay(time1.Unix(), time2.Unix(), DAILY_RESET_HOUR)
}

func IsSameDayWithStartHour(time1 time.Time, time2 time.Time, startHour int) bool {
	return IsTmSameDay(time1.Unix(), time2.Unix(), startHour)
}

func IsTmSameDay(tm1 int64, tm2 int64, startHour int) bool {
	return GetPassedDaysVisTs(tm1, tm2, startHour) == 0
	//return (tm1-Time1970.Unix()-int64(startHour)*3600)/86400 == (tm2-Time1970.Unix()-int64(startHour)*3600)/86400
}

func GetPassedDaysVisTs(tm1 int64, tm2 int64, startHour int) int64 {
	return (tm1-Time1970.Unix()-int64(startHour)*3600)/86400 - (tm2-Time1970.Unix()-int64(startHour)*3600)/86400
}

func IsSameWeek(time1 time.Time, time2 time.Time, startHour int) bool {
	_, week1 := time1.ISOWeek()
	_, week2 := time2.ISOWeek()
	return week1 == week2
}

func IsSameYear(time1 time.Time, time2 time.Time) bool {
	return time1.Year() == time2.Year()
}

func CheckNameByLike(name string) error {

	if !nameReg.MatchString(name) {
		return gamedb.ERRNICKNAMEINVALID
	}
	if pass := gamedb.CensorIsPass(name); !pass {
		return gamedb.ERRHASSENSITIVECHARACTER
	}
	return nil
}

func CheckName(name string) error {

	err := CheckNameSimple(name)
	if err != nil {
		return err
	}
	if !nameReg.MatchString(name) {
		return gamedb.ERRNICKNAMEINVALID
	}
	if pass := gamedb.CensorIsPass(name); !pass {
		return gamedb.ERRHASSENSITIVECHARACTER
	}
	return nil
}

func CensorAndReplace(text string) (bool, string) {
	return textcensor.CheckAndReplace(text, true, '*')
}

func CheckNameSimple(name string) error {

	nickNameRune := []rune(name)

	l := 0
	for _, v := range nickNameRune {
		if v < 128 {
			l++
		} else {
			l += 2
		}
	}

	if l < 3 || l > 14 {
		return gamedb.ERRNICKNAMELENGTHINVALID
	}

	return nil
}

func CalcSkillLvl(weaponGrade int) int {
	return weaponGrade + 1
}

func GetSecondsFromZeroTime(hour, min int) int {
	return hour*3600 + min*60
}

func GetFullSecondsFromZeroTime(hour, min, second int) int {
	return hour*3600 + min*60 + second
}

func GetSecondsFromZeroTimeNow() int {
	now := time.Now()
	hour := now.Hour()
	min := now.Minute()
	return hour*3600 + min*60 + now.Second()
}

//获取今日零点的时间戳
func GetZeroTimeUnixFrom1970() int {
	now := time.Now()
	return int(now.Unix()) - now.Hour()*3600 - now.Minute()*60 - now.Second()
}

func StringTimeToTimeInt(timeStr string) (int, bool) {
	timeList := strings.Split(timeStr, ":")
	if len(timeList) != 2 {
		return 0, false
	}
	hour, _ := strconv.Atoi(timeList[0])
	minutes, _ := strconv.Atoi(timeList[1])
	return hour*3600 + minutes*60, true
}

func ZeroTimeOfResetDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, DAILY_RESET_HOUR, 0, 0, 0, time.Local)
}

//得到当前时间到2028年的时间
//一般用在排行榜上,比如同等级,判断时间先后
func Time2TenYears() time.Duration {
	return Time2028.Sub(time.Now())
}
