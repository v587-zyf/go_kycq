package scene

import (
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"errors"
)

type ISceneObj interface {
	GetContext() ISceneObj
	GetObjId() int
	GetType() int
	IsSceneObj() bool

	SetBlockIndex(index int)

	Point() *Point         //获取当前坐标
	SetPoint(point *Point) //设置坐标
	GetDir() int
	SetDir(dir int)
	GetVisible() bool
	SetVisible(visible bool)
	GetScene() *Scene
	SetScene(scene *Scene)

	HostId() int
	SetHostId(hostId int)
	SessionId() uint32
	SetSessionId(sessionId uint32)

	EnterScene(scene *Scene, point *Point) error
	LeaveScene()
	MoveTo(point *Point, moveType int, moveForce, sendClient bool) error

	BuildSceneObjMessage() nw.ProtoMessage
	BuildAppearMessage() nw.ProtoMessage
	BuildRelliveMessage() nw.ProtoMessage
	BuildMoveMessage(moveType int, moveForce bool) nw.ProtoMessage
	BuildDisappearMessage() nw.ProtoMessage
	GetSpecialAI() int32
}

type BatchBuilder interface {
	AddObj(obj ISceneObj)
	Build() nw.ProtoMessage
	Reset()
	Length() int
	SetEnterType(enterType int)
}

type DefaultSceneObj struct {
	context   ISceneObj
	objId     int
	typ       int
	point     *Point
	dir       int
	visible   bool
	scene     *Scene
	hostId    int
	sessionId uint32

	blockIndex int
}

func (this *DefaultSceneObj) SetBlockIndex(index int) {
	//if index > 0 && this.blockIndex > 0 {
	//	logger.Error("-----------%v---异常设置区域Id-------当前：%v, 新区域：%v--------", this.objId, this.blockIndex, index)
	//	logger.Error("-----------%v---异常设置区域Id-------当前：%v, 新区域：%v--------", this.objId, this.blockIndex, index)
	//	logger.Error("-----------%v---异常设置区域Id-------当前：%v, 新区域：%v--------", this.objId, this.blockIndex, index)
	//	logger.Error("-----------%v---异常设置区域Id-------当前：%v, 新区域：%v--------", this.objId, this.blockIndex, index)
	//	logger.Error("-----------%v---异常设置区域Id-------当前：%v, 新区域：%v--------", this.objId, this.blockIndex, index)
	//	logger.Error("-----------%v---异常设置区域Id-------当前：%v, 新区域：%v--------", this.objId, this.blockIndex, index)
	//}
	this.blockIndex = index
}

func (this *DefaultSceneObj) Point() *Point {
	return this.point
}

func (this *DefaultSceneObj) SetPoint(point *Point) {
	this.point = point
}

type BatchBuilderType int

const (
	BatchBuilderTypeEnter         BatchBuilderType = 1
	BatchBuilderTypeLeave                          = 2
	BatchBuilderTypeEnterForTower                  = 3
	BatchBuilderTypeLeaveForTower                  = 4
)

var sceneObjIdAllocator = util.NewUint32IdAllocator(1)

func NewDefaultSceneObj(context ISceneObj, typ int) *DefaultSceneObj {
	return &DefaultSceneObj{
		objId:   int(sceneObjIdAllocator.Get()),
		typ:     typ,
		visible: true,
		context: context,
	}
}

func (this *DefaultSceneObj) GetContext() ISceneObj {
	return this.context
}

func (this *DefaultSceneObj) CreateNewObjId() {
	this.objId = int(sceneObjIdAllocator.Get())
}

func (this *DefaultSceneObj) GetObjId() int {
	return this.objId
}
func (this *DefaultSceneObj) GetType() int {
	return this.typ
}

func (this *DefaultSceneObj) IsSceneObj() bool {
	if this.typ == pb.SCENEOBJTYPE_ITEM || this.typ == pb.SCENEOBJTYPE_COLLECTION || this.typ == pb.SCENEOBJTYPE_BUFF {
		return true
	}
	return false
}

func (this *DefaultSceneObj) GetDir() int {
	return this.dir
}

func (this *DefaultSceneObj) SetDir(dir int) {
	this.dir = dir
}

func (this *DefaultSceneObj) GetVisible() bool {
	return this.visible
}

func (this *DefaultSceneObj) SetVisible(visible bool) {
	// scene := this.scene
	// if scene != nil {
	// 	if visible {
	// 		msg := this.BuildAppearMessage()
	// 		scene.NotifyNearby(this, msg, false)
	// 	} else {
	// 		msg := this.BuildDisappearMessage()
	// 		scene.NotifyNearby(this, msg, false)
	// 	}
	// }
	this.visible = visible
}

func (this *DefaultSceneObj) HostId() int {
	return this.hostId
}

func (this *DefaultSceneObj) SetHostId(hostId int) {
	this.hostId = hostId
}

func (this *DefaultSceneObj) SessionId() uint32 {
	return this.sessionId
}

func (this *DefaultSceneObj) SetSessionId(sessionId uint32) {
	this.sessionId = sessionId
}

func (this *DefaultSceneObj) GetScene() *Scene {
	return this.scene
}

func (this *DefaultSceneObj) SetScene(scene *Scene) {
	this.scene = scene
}

func (this *DefaultSceneObj) EnterScene(scene *Scene, point *Point) error {
	if err := scene.AddSceneObj(this, point); err != nil {
		return err
	}
	this.scene = scene
	logger.Info("sceneObj:%d enter scene,%v,bornPoint:%v", this.objId, this, point.ToString())
	return nil
}

func (this *DefaultSceneObj) LeaveScene() {
	scene := this.scene
	if scene == nil {
		return
	}
	logger.Info("sceneObj:%d leave scene", this.objId)
	scene.RemoveSceneObj(this)
	this.scene = nil
}

func (this *DefaultSceneObj) MoveTo(point *Point, moveType int, moveForce, sendClient bool) error {
	scene := this.scene
	if scene == nil {
		return errors.New("not on a scene")
	}
	dir := GetFaceDirByPoint(this.point, point)
	if err := scene.MoveSceneObj(this, point, moveType, moveForce, sendClient); err != nil {
		return err
	}
	this.SetDir(dir)
	return nil
}

func (this *DefaultSceneObj) BuildSceneObjMessage() nw.ProtoMessage {
	return nil
}

func (this *DefaultSceneObj) BuildAppearMessage() nw.ProtoMessage {
	return nil
}

func (this *DefaultSceneObj) BuildRelliveMessage() nw.ProtoMessage {
	return nil
}

func (this *DefaultSceneObj) BuildMoveMessage(moveType int, moveForce bool) nw.ProtoMessage {
	ack := &pb.SceneMoveNtf{
		Point:    this.Point().ToPbPoint(),
		ObjId:    int32(this.GetObjId()),
		MoveType: int32(moveType),
		Force:    moveForce,
	}
	return ack
}

func (this *DefaultSceneObj) BuildDisappearMessage() nw.ProtoMessage {
	r := &pb.SceneLeaveNtf{}
	r.ObjIds = append(r.ObjIds, int32(this.GetObjId()))
	return r
}

func (this *DefaultSceneObj) GetSpecialAI() int32 {
	return 0
}
