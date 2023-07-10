package scene

import (
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func GetBatchBuilder(stageId int, typ BatchBuilderType) BatchBuilder {
	switch typ {
	case BatchBuilderTypeEnter, BatchBuilderTypeEnterForTower:
		return NewEnterBuilder(stageId, typ)
	case BatchBuilderTypeLeave, BatchBuilderTypeLeaveForTower:
		return NewLeaveBuilder(typ)
	}
	return nil
}

type EnterBuilder struct {
	msg *pb.SceneEnterNtf
}

func NewEnterBuilder(stageId int, typ BatchBuilderType) BatchBuilder {
	msg := &EnterBuilder{
		msg: &pb.SceneEnterNtf{
			StageId: int32(stageId),
		},
	}
	if typ == BatchBuilderTypeEnterForTower {
		msg.msg.IsTower = true
	}
	return msg
}

func (this *EnterBuilder) AddObj(obj ISceneObj) {
	sceneObj := obj.BuildSceneObjMessage().(*pb.SceneObj)
	//合体进入，不论死活
	if this.msg.EnterType == constFight.SCENE_ENTER_FIT || sceneObj.Hp > 0 || obj.IsSceneObj() {
		this.msg.Objs = append(this.msg.Objs, sceneObj)
	}
}

func (this *EnterBuilder) Build() nw.ProtoMessage {
	return this.msg
}

func (this *EnterBuilder) Reset() {
	this.msg.Objs = append(this.msg.Objs[:0], this.msg.Objs[len(this.msg.Objs):]...)
}

func (this *EnterBuilder) Length() int {
	return len(this.msg.Objs)
}

func (this *EnterBuilder) SetEnterType(enterType int) {
	this.msg.EnterType = int32(enterType)
}

type LeaveBuilder struct {
	msg *pb.SceneLeaveNtf
}

func NewLeaveBuilder(typ BatchBuilderType) BatchBuilder {
	leaveMsg := &LeaveBuilder{
		msg: &pb.SceneLeaveNtf{},
	}
	if typ == BatchBuilderTypeLeaveForTower {
		leaveMsg.msg.IsTower = true
	}
	return leaveMsg
}

func (this *LeaveBuilder) AddObj(obj ISceneObj) {
	this.msg.ObjIds = append(this.msg.ObjIds, int32(obj.GetObjId()))
}

func (this *LeaveBuilder) Build() nw.ProtoMessage {
	return this.msg
}

func (this *LeaveBuilder) Reset() {
	this.msg.ObjIds = append(this.msg.ObjIds[:0], this.msg.ObjIds[len(this.msg.ObjIds):]...)
}

func (this *LeaveBuilder) Length() int {
	return len(this.msg.ObjIds)
}
func (this *LeaveBuilder) SetEnterType(enterType int) {
}
