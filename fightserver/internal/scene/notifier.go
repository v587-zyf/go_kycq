package scene

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/nw"
)

type Notifier interface {
	Add(obj ISceneObj)                                                                  // 增加一个对象，并通知相关消息
	Adds(objs []ISceneObj, points []*Point, enterType int)                              // 增加多个对象，并通知相关消息
	Remove(obj ISceneObj)                                                               // 移除一个对象，并通知相关消息
	Move(obj ISceneObj, oldPoint *Point, moveType int, moveForce bool, sendClient bool) // 移动一个对象，并通知相关消息
	Update(obj ISceneObj)                                                               // 更新一个对象
	Relive(obj ISceneObj, oldPoint *Point, reliveType int)                              // 复活一个对象
	NotifyNearby(obj ISceneObj, msg nw.ProtoMessage, excludeSession map[uint32]bool)
	NotifyAll(msg nw.ProtoMessage)
}

type NotifierBase struct {
	StageId int
}

func (this *NotifierBase) Add(obj ISceneObj)                                     {} // 增加一个对象，并通知相关消息
func (this *NotifierBase) Adds(objs []ISceneObj, points []*Point, enterType int) {} // 增加多个对象，并通知相关消息
func (this *NotifierBase) Remove(obj ISceneObj)                                  {} // 移除一个对象，并通知相关消息
func (this *NotifierBase) Move(obj ISceneObj, oldPoint *Point, moveType int, moveForce bool, sendClient bool) {
}                                                                                   // 移动一个对象，并通知相关消息
func (this *NotifierBase) Update(obj ISceneObj)                                  {} // 更新一个对象
func (this *NotifierBase) Relive(obj ISceneObj, oldPoint *Point, reliveType int) {} // 复活一个对象
func (this *NotifierBase) NotifyNearby(obj ISceneObj, msg nw.ProtoMessage, excludeSession map[uint32]bool) {
}
func (this *NotifierBase) NotifyAll(msg nw.ProtoMessage) {}

const (
	MaxNotifierNum = 10
)

func NewNotifier(stageId int, sceneT *gamedb.SceneConf) Notifier {
	stageConf := gamedb.GetStageStageCfg(stageId)
	mapTypeConf := gamedb.GetMaptypeGameCfg(stageConf.Type)
	if mapTypeConf.Tower == 0 {
		return NewSimpleNotifier(stageId)
	} else {
		return NewBlockNotifier(stageId, sceneT)
	}
}

type notifyItem struct {
	isNotifyAll    bool
	obj            ISceneObj
	msg            nw.ProtoMessage
	excludeSession map[uint32]bool
}

func newNotifyItem(isNotifyAll bool, obj ISceneObj, msg nw.ProtoMessage, excludeSelf map[uint32]bool) *notifyItem {
	return &notifyItem{
		isNotifyAll:    isNotifyAll,
		obj:            obj,
		msg:            msg,
		excludeSession: excludeSelf,
	}
}
