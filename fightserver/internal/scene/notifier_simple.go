package scene

import (
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

type SimpleNotifier struct {
	*NotifierBase
	objs map[int]ISceneObj
}

func NewSimpleNotifier(stageId int) *SimpleNotifier {
	notifier := &SimpleNotifier{
		NotifierBase: &NotifierBase{StageId: stageId},
		objs:         make(map[int]ISceneObj),
	}
	return notifier
}

func (this *SimpleNotifier) Add(obj ISceneObj) {
	obj = obj.GetContext()
	this.objs[obj.GetObjId()] = obj
	if obj.GetVisible() {
		appearMsg := obj.BuildAppearMessage()
		this.NotifyAll(appearMsg)
	}
	if obj.GetType() == pb.SCENEOBJTYPE_USER {
		this.sendOthersMsgTo(obj, BatchBuilderTypeEnter, nil)
		//推送玩家进入场景结束
		if obj.GetVisible() && obj.HostId() > 0 && obj.SessionId() > 0 {
			net.GetGateConn().SendMessage(uint32(obj.HostId()), obj.SessionId(), 0, &pb.SceneEnterOverNtf{})
		}
	}
}

// 增加多个对象，并通知相关消息{
func (this *SimpleNotifier) Adds(objs []ISceneObj, points []*Point, enterType int) {

	firstObj := objs[0].GetContext()
	var builder = GetBatchBuilder(this.StageId, BatchBuilderTypeEnter)
	builder.SetEnterType(enterType)
	withoutObjs := make(map[int]bool)
	for _, obj := range objs {
		obj = obj.GetContext()
		this.objs[obj.GetObjId()] = obj
		withoutObjs[obj.GetObjId()] = true
		if obj.GetVisible() {
			builder.AddObj(obj)
		}
	}

	appearMsg := builder.Build()
	this.NotifyAll(appearMsg)

	if enterType != constFight.SCENE_ENTER_FIT && firstObj.GetType() == pb.SCENEOBJTYPE_USER {
		this.sendOthersMsgTo(firstObj, BatchBuilderTypeEnter, withoutObjs)
		//推送玩家进入场景结束
		if firstObj.GetVisible() && firstObj.HostId() > 0 && firstObj.SessionId() > 0 {
			net.GetGateConn().SendMessage(uint32(firstObj.HostId()), firstObj.SessionId(), 0, &pb.SceneEnterOverNtf{})
		}
	}
}

func (this *SimpleNotifier) Update(obj ISceneObj) {
	obj = obj.GetContext()
	this.objs[obj.GetObjId()] = obj
	if obj.GetVisible() {
		appearMsg := obj.BuildAppearMessage()
		this.NotifyAll(appearMsg)
	}
}

func (this *SimpleNotifier) Relive(obj ISceneObj, oldPoint *Point, reliveType int) {
	obj = obj.GetContext()
	this.objs[obj.GetObjId()] = obj
	if obj.GetVisible() {
		appearMsg := obj.BuildRelliveMessage()
		this.NotifyAll(appearMsg)
	}
}

func (this *SimpleNotifier) Remove(obj ISceneObj) {
	obj = obj.GetContext()
	delete(this.objs, obj.GetObjId())
	//if obj.GetVisible() {
	disappearMsg := obj.BuildDisappearMessage()
	this.NotifyAll(disappearMsg)
	//}
}

func (this *SimpleNotifier) Move(obj ISceneObj, point *Point, moveType int, moveForce bool, sendClinet bool) {
	if !sendClinet {
		return
	}
	obj = obj.GetContext()
	if obj.GetVisible() {
		moveMsg := obj.BuildMoveMessage(moveType, moveForce)
		this.NotifyAll(moveMsg)
	}
}

func (this *SimpleNotifier) NotifyNearby(sceneObj ISceneObj, msg nw.ProtoMessage, excludeSession map[uint32]bool) {
	ids := make(map[int]map[int]int)
	for _, obj := range this.objs {
		if obj != nil && obj.SessionId() > 0 {
			if excludeSession != nil && excludeSession[obj.SessionId()] {
				continue
			}
			hostId := obj.HostId()
			if hostId == 0 {
				continue
			}
			if ids[hostId] == nil {
				ids[hostId] = make(map[int]int)
			}
			ids[hostId][int(obj.SessionId())] = hostId
		}
	}
	if len(ids) > 0 {
		net.GetGateConn().BroadcastToGate(ids, msg)
	}
}

func (this *SimpleNotifier) NotifyAll(msg nw.ProtoMessage) {
	ids := make(map[int]map[int]int)
	for _, obj := range this.objs {
		hostId := obj.HostId()
		if hostId == 0 {
			continue
		}
		if obj.SessionId() <= 0 {
			continue
		}
		if ids[hostId] == nil {
			ids[hostId] = make(map[int]int)
		}

		ids[hostId][int(obj.SessionId())] = hostId
	}
	if len(ids) > 0 {
		net.GetGateConn().BroadcastToGate(ids, msg)
	}
}

// sendOthersMsgTo 发送其他人的出现或消失消息
func (this *SimpleNotifier) sendOthersMsgTo(target ISceneObj, typ BatchBuilderType, withoutobjs map[int]bool) {
	hostId := target.HostId()
	if hostId <= 0 {
		return
	}
	logger.Info("推送角色：%v场景内角色,场景内玩家：%v,消息统计玩家：%v", target.GetObjId(), (this.objs))
	var builder = GetBatchBuilder(this.StageId, typ)
	var notifyCount = 0
	for _, obj := range this.objs {
		if withoutobjs != nil && withoutobjs[obj.GetObjId()] {
			continue
		}
		if obj != target && obj.GetVisible() {
			builder.AddObj(obj)
			notifyCount++
		}
		if notifyCount >= MaxNotifierNum {
			net.GetGateConn().SendMessage(uint32(hostId), target.SessionId(), 0, builder.Build())
			builder.Reset()
			notifyCount = 0
		}
	}
	if builder != nil && builder.Length() > 0 {
		net.GetGateConn().SendMessage(uint32(hostId), target.SessionId(), 0, builder.Build())
	}
}
