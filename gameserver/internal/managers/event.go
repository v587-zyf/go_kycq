package managers

import (
	"time"

	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

const (
	SHOP          = 1
	LOTTERY       = 2
	TREASURE_COPY = 3
)

type EventManager struct {
	util.DefaultModule
	eventFIFO          []*pb.EventNtf
	eventChan          chan *pb.EventNtf
	lastSendAt         time.Time
	ntfSecondsPerCount int
}

func (this *EventManager) Init() error {
	this.eventChan = make(chan *pb.EventNtf, 200)
	go this.eventService()
	return nil
}

func (this *EventManager) eventService() {
	//ticker := time.NewTicker(time.Second * 2)
	//for {
	//	select {
	//	case msg := <-this.eventChan:
	//		this.eventFIFO = append(this.eventFIFO, msg)
	//		this.lastSendAt = time.Now()
	//	case <-ticker.C:
	//		l := len(this.eventFIFO)
	//		if l != 0 {
	//			now := time.Now()
	//			if int(time.Since(this.lastSendAt).Seconds()) > gameDb().GetConf().TipsRandSendSeconds {
	//				randomMsg := this.eventFIFO[rand.Intn(l)]
	//				m.ClientManager.BroadcastAll(randomMsg)
	//				this.lastSendAt = now
	//			}
	//			timePassed := now.Unix() - int64(this.eventFIFO[0].Ts)
	//			this.ntfSecondsPerCount = int(timePassed) / l
	//		} else {
	//			this.ntfSecondsPerCount = common.MaxInt
	//		}
	//	}
	//}
}

func (this *EventManager) ShouldTrigger(oldValue, newValue int, steps gamedb.PropInfos) bool {
	for _, v := range steps {
		//logger.Debug("ShouldTrigger:old:%d,new:%d,k:%d,v:%d, ntfSecondsPerCount:%d", oldValue, newValue, v.K, v.V, this.ntfSecondsPerCount)
		if newValue < v.K {
			return false
		}
		if oldValue < v.K && newValue >= v.K && this.ntfSecondsPerCount >= v.V {
			return true
		}
	}
	return false
}

func (this *EventManager) ShouldTriggerExactlyAt(limitArg int, steps gamedb.IntMap) bool {
	if v, ok := steps[limitArg]; ok {
		return this.ntfSecondsPerCount >= v
	}
	return false
}

func (this *EventManager) BroadcastWithSourceId(eventId, sourceId int, args ...interface{}) {
	stringSlice := common.InterfaceSlice2StringSlice(args)
	ntf := builder.BuildEventNtfWithSourceId(eventId, stringSlice, sourceId)
	m.ClientManager.BroadcastAll(ntf)
	select {
	case this.eventChan <- ntf:
	default:
		logger.Warn("eventManager : eventchan is full, please check .")
	}
}

func (this *EventManager) BroadcastEvent(eventId int, args ...interface{}) {
	this.BroadcastWithSourceId(eventId, 0, args...)
}

func (this *EventManager) BroadcastCheckGapArgs(eventId, sourceId, oldLevel, newLevel int, args ...interface{}) {
	//conf := gameDb().GetEventItem(eventId)
	//if conf == nil {
	//	logger.Error("BroadcastCheckGapArgs:nil eventId:%d", eventId)
	//	return
	//}
	//if this.ShouldTrigger(oldLevel, newLevel, conf.Arg) {
	//	this.BroadcastWithSourceId(eventId, sourceId, args...)
	//}
}
func (this *EventManager) BroadcastCheckArgs(eventId, sourceId, limitArg int, args ...interface{}) {
	//conf := gameDb().GetEventItem(eventId)
	//if conf == nil {
	//	logger.Error("BroadcastCheckArgs:nil eventId:%d", eventId)
	//	return
	//}
	//if this.ShouldTriggerExactlyAt(limitArg, conf.ArgMap) {
	//	this.BroadcastWithSourceId(eventId, sourceId, args...)
	//}
}

func (this *EventManager) BroadcastGotNiceItemEvent(user *objs.User, eventId int, item *model.Item, evenType int) {
	//itemConf := gameDb().GetItem(item.ItemIndex)
	//if itemConf == nil {
	//	logger.Error("BroadcastGotNiceItemEvent:nil item:%d", item.ItemIndex)
	//}
	//if itemConf.Type != pb.ITEMTYPE_WEAPON && itemConf.Type != pb.ITEMTYPE_WEAPON_CARD &&
	//	itemConf.Type != pb.ITEMTYPE_HERO && itemConf.Type != pb.ITEMTYPE_HERO_CARD &&
	//	itemConf.Type != pb.ITEMTYPE_MAGIC_CARD && itemConf.Type != pb.ITEMTYPE_MAGIC {
	//	return
	//}
	//eventTypeStr := ""
	//switch evenType {
	//case SHOP:
	//	eventTypeStr += base.GetCodeTextById(base.SHANG_DIAN)
	//case LOTTERY:
	//	eventTypeStr += base.GetCodeTextById(base.LIAN_YIN)
	//case TREASURE_COPY:
	//	eventTypeStr += base.GetCodeTextById(base.BAO_ZANG_FU_BEN)
	//default:
	//	return
	//}
	//
	//conf := gameDb().GetEventItem(eventId)
	//if conf == nil {
	//	logger.Error("BroadcastGotNiceItemEvent:nil eventId:%d", eventId)
	//	return
	//}
}

func (this *EventManager) BroadcastUserPatrol(dailyEventId, patrolDailyNum, totalEventId, patrolTotalNum int, user *objs.User) {
	this.BroadcastUserPatrolEvent(dailyEventId, patrolDailyNum, user)
	this.BroadcastUserPatrolEvent(totalEventId, patrolTotalNum, user)
}

func (this *EventManager) BroadcastUserPatrolEvent(eventId, patrolNum int, user *objs.User) {
	this.BroadcastCheckArgs(eventId, user.Id, patrolNum, user.NickName, patrolNum)
}

func (this *EventManager) BroadcastUserMrankFirstEvent(eventId, point int, user *objs.User) {
	this.BroadcastWithSourceId(eventId, user.Id, user.NickName, point)
}


