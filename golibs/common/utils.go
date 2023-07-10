package common

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//TenThousand 1万
const TenThousand = 10000

const MaxInt = int(^uint(0) >> 1)

func RandNum(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func Sample(arr []int) (error, int) {
	l := len(arr)
	if l == 0 {
		return errors.New("common/util.go Sample can't take empty array"), 0
	}
	return nil, arr[rand.Intn(l)]
}

func ZeroTimeOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func NormalizeTimeOfDay(t time.Time, startHour int) time.Time {
	if t.Hour() < startHour {
		year, month, day := t.AddDate(0, 0, -1).Date()
		return time.Date(year, month, day, startHour, 0, 0, 0, time.Local)
	} else {
		year, month, day := t.Date()
		return time.Date(year, month, day, startHour, 0, 0, 0, time.Local)
	}
}

func DiffDays(endTime time.Time, startTime time.Time) int {
	year2, month2, day2 := endTime.Date()
	year1, month1, day1 := startTime.Date()
	d2 := time.Date(
		year2, month2, day2,
		0, 0, 0, 0, time.Local,
	)
	d1 := time.Date(
		year1, month1, day1,
		0, 0, 0, 0, time.Local,
	)
	return int(d2.Sub(d1) / (24 * time.Hour))
}

func GetIpAddress(r *http.Request) string {
	forwardedFor := r.Header.Get("X-Forwarded-For")
	if forwardedFor != "" {
		// X-Forwarded-For is potentially a list of addresses separated with ","
		parts := strings.Split(forwardedFor, ",")
		for _, part := range parts {
			ip := strings.TrimSpace(part)
			if ip != "" {
				return ip
			}
		}
	}
	ip := r.Header.Get("X-Real-Ip")
	if ip != "" {
		return ip
	}
	index := strings.LastIndex(r.RemoteAddr, ":")
	if index < 0 {
		return r.RemoteAddr
	}
	return r.RemoteAddr[:index]
}

//HitRateTenThousand
//是否命中万分比
func HitRateTenThousand(rate int) bool {
	return rand.Intn(TenThousand) < rate
}

func GetOutboundIp() string {
	netAddr, err := net.ResolveTCPAddr("tcp", "www.baidu.com:80")
	if err != nil {
		panic("can not get outbound ip " + err.Error())
	}
	conn, err := net.Dial("udp", netAddr.IP.String()+":80")
	if err != nil {
		panic("can not get outbound ip " + err.Error())
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")

	return localAddr[0:idx]

}

func GetTomorrowStamp() time.Time {
	tomorrow := time.Now().Add(24 * time.Hour)
	year, month, day := tomorrow.Date()
	//tomorrow_str := fmt.Sprintf("%d-%d-%d 00:00:00", year, month, day)
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func IntMinusEncode(num int) int {
	newNum := num<<16 | rand.Intn(9999)
	return -newNum
}

func IntMinusDecode(num int) int {
	newNum := -int(num)
	return newNum >> 16
}

//float32 四舍五入
func RoundFloat32(x float32) int {
	return RoundFloat64(float64(x))
}

//float64 四舍五入
func RoundFloat64(x float64) int {
	return int(math.Floor(x + 0.5))
}

//往上取整
func CeilFloat32(x float32) int {
	return CeilFloat64(float64(x))
}

//往上取整(保留两位小数)
func CeilFloat64(x float64) int {
	return int(math.Ceil(RoundFloat(x, 2)))
}

//往下取整
func FloorFloat32(x float32) int {
	return FloorFloat64(float64(x))
}

//往下取整
func FloorFloat64(x float64) int {
	return int(math.Floor(x))
}

func MinIntGet(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxIntGet(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func IntAbs(a int) int {
	return int(math.Abs(float64(a)))
}

func DecToMin(src *int, decValue int, min int) int {
	var diff = *src - min
	*src -= decValue
	if *src < min {
		*src = min
		return diff
	}
	return decValue
}

// f:需要处理的浮点数，n：要保留小数的位数
// Pow10（）返回10的n次方，最后一位四舍五入，对ｎ＋１位加０．５后四舍五入
func RoundFloat(f float64, n int) float64 {
	n10 := math.Pow10(n)
	if f < 0 {
		return math.Trunc((math.Abs(f)+0.5/n10)*n10*-1) / n10
	} else {
		return math.Trunc((math.Abs(f)+0.5/n10)*n10) / n10
	}
}

/**
 * 通过string获得一个[]string
 * @param str  10,1000,1000,1000
 * @param sep1 分隔符 ","
 */
func NewStringSlice(str string, sep string) []string {
	intSliceList := make([]string, 0)
	list := strings.Split(str, sep)
	for _, v := range list {
		intSliceList = append(intSliceList, v)
	}
	return intSliceList
}

//去除string中的非法字符
func StringValid(str string) string {
	name := ""
	for _, i := range str {
		if i > 31 && i != 92 {
			name += string(i)
		}
	}
	return name
}

func ThrowPanic(sandbox bool, err error) {
	if sandbox {
		panic(err)
	}
}

/**
 *  @Description: 计算万分比数值
 *  @param scale	万分比例
 *  @param num		计算的值
 *  @return int
 */
func CalcTenThousand(scale int, num int) int {
	res := float32(num) * (float32(scale) / float32(TenThousand))
	resultInt := CeilFloat32(float32(num) + res)
	return resultInt
}


func ToUrlValues(param map[string]interface{}) url.Values {
	values := make(url.Values)
	for k, v := range param {
		values.Add(k, fmt.Sprintf("%v", v))
	}
	return values
}


func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}


func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}