package msglog

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbgt"
	"cqserver/protobuf/pbserver"
	"github.com/astaxie/beego/logs"
)

const (
	MSG_PB        = "pb"
	MSG_PBGT      = "pbgt"
	MSG_PB_SERVER = "pbserver"
)

type uniqueMsgInfo struct {
	Unique     string //统计用唯一标示
	FistTime   int64  //开始记录时间（第一条消息）
	UpdateTime int64  //最近一条消息时间
	MsgCount   int    //总消息数
	LastCount  int    //上次统计消息数
}

type msgRecord struct {
	MsgId int //消息ID
	Num   int //消息数
}

var msgLog *logs.BeeLogger = logger.Get("default", true)
var logLv int = 10
var msgStatisticsMap map[string]*uniqueMsgInfo = make(map[string]*uniqueMsgInfo)
var mu sync.Mutex
var msgRecordNum map[int]int = make(map[int]int)
var msgRecordNumChan chan int = make(chan int, 10)

func init() {
	go util.SafeRun(msgStatistics)
	go util.SafeRun(func() {

		timer := time.NewTicker(1 * time.Hour)
		for {
			select {
			case msgId := <-msgRecordNumChan:
				msgRecordNum[msgId] ++
			case <-timer.C:
				if len(msgRecordNum) > 0 {
					logger.Info("消息数量统计：%v", msgRecordNum)
				}
			}
		}
	})
}

func SetPrintMsgLv(lv int) {
	logLv = lv
}

/**
*打印消息
*msgType 消息类型（MSG_PB:pb MSG_PBGT:pbgt MSG_PB_SERVER:pbserver）
*msgId	消息Id
*msg	消息体内容
*unique 统计唯一标示（例：game玩家Id,gate玩家sessionid）
*otherInfo 其他备注消息
 */
func PrintMsg(unique string, msgType string, msgId int, msg nw.ProtoMessage, otherInfo string) {

	defer func() {
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			msgLog.Error("PrintMsg %s", stackBytes)
		}
	}()
	mu.Lock()
	if unique != "" && unique != "0" {
		if msgStatisticsMap[unique] == nil {
			msgStatisticsMap[unique] = &uniqueMsgInfo{
				Unique:   unique,
				FistTime: time.Now().Unix(),
			}
		}
		msgStatisticsMap[unique].UpdateTime = time.Now().Unix()
		msgStatisticsMap[unique].MsgCount += 1
	}
	mu.Unlock()
	msglogLv := -1
	switch msgType {
	case MSG_PB:
		msglogLv = pb.GetMsgLogLv(msgId)
	case MSG_PBGT:
		msglogLv = pbgt.GetMsgLogLv(msgId)
	case MSG_PB_SERVER:
		msglogLv = pbserver.GetMsgLogLv(msgId)
	}
	//判断消息log等级和打印设置等级
	if msglogLv >= logLv {
		msgLog.Info("unique[%v],msgId[%v],msg[%v],otherInfo[%v],", unique, msgId, msg.String(), otherInfo)
	}
}

//每分钟消息数
func msgStatistics() {
	timer := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-timer.C:
			now := time.Now().Unix()
			mu.Lock()
			for u, v := range msgStatisticsMap {
				av, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(v.MsgCount)/(float64(now-v.FistTime)/60)), 64)
				msgLog.Info("unique[%v],总消息数[%v],上次统计到本次消息数[%v],每分钟消息数[%v]", u, v.MsgCount, (v.MsgCount - v.LastCount), av)
				v.LastCount = v.MsgCount
				if now-v.UpdateTime > 300 {
					delete(msgStatisticsMap, u)
				}
			}
			mu.Unlock()
		}
	}
}

func MsgCostTime(startTime time.Time, msgType string, msgId uint16, arg ...interface{}) {

	usedTime := time.Now().Sub(startTime)
	if usedTime > 50*time.Millisecond {
		msgName := ""
		switch msgType {
		case MSG_PB:
			msgName = pb.GetMsgName(msgId)
		case MSG_PBGT:
			msgName = pbgt.GetMsgName(msgId)
		case MSG_PB_SERVER:
			msgName = pbserver.GetMsgName(msgId)
		}
		msgLog.Warn("msgCostTime cssesion.go Handle cost too much time.msgId:%v,msgName:%v,other:%v,costTime:%v", msgId, msgName, arg, usedTime)
	}
}

func MsgNumRecord(msgId int) {
	msgRecordNumChan <- msgId
}
