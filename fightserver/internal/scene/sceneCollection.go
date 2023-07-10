package scene

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"time"
)

type SceneCollection struct {
	*DefaultSceneObj
	collectionId    int
	collectionObjId int
	collectionTime  int64
	reShowTime      int64
}

func NewSceneCollection(collectionId int) *SceneCollection {

	dropItem := &SceneCollection{
		collectionId: collectionId,
	}
	dropItem.DefaultSceneObj = NewDefaultSceneObj(dropItem, pb.SCENEOBJTYPE_COLLECTION)
	return dropItem
}

func (this *SceneCollection) CanCollection(pickuserId int) bool {

	now := int(time.Now().Unix())
	if this.collectionObjId > 0 {
		logger.Error("物品不能被采集，物品:%v,被采集中：%v,当前时间：%v", this.objId, this.collectionObjId, now)
		return false
	}

	return true
}

func (this *SceneCollection) CanCancelCollection(userObjId int) bool {

	now := int(time.Now().Unix())
	if this.collectionObjId != userObjId {
		logger.Error("物品不能被取消采集，物品:%v,被采集中：%v,当前时间：%v", this.objId, this.collectionObjId, userObjId, now)
		return false
	}

	if this.disappeared(now) {
		return false
	}

	return true
}

func (this *SceneCollection) Collection(actorObjId int) {
	this.collectionObjId = actorObjId
	this.collectionTime = time.Now().Unix()
}

func (this *SceneCollection) CollectionId() int {
	return this.collectionId
}

func (this *SceneCollection) disappeared(now int) bool {

	if this.collectionObjId <= 0 {
		return false
	}
	disappearTime := int(this.GetEndTime())
	if now > disappearTime {
		return true
	}
	return false
}

func (this *SceneCollection) ReShow() bool {
	if this.reShowTime > 0 && this.reShowTime <= time.Now().Unix() {
		this.reShowTime = 0
		return true
	}
	return false
}

func (this *SceneCollection) Reset(isCollection bool) {
	this.collectionObjId = 0
	this.collectionTime = 0
	conf := gamedb.GetCollectionCollectionCfg(this.collectionId)
	if isCollection && conf.Type == constFight.COLLECTION_TYPE_ONE {
		this.reShowTime = time.Now().Unix() + int64(conf.Time)
	}
}

func (this *SceneCollection) getCollectionId() int {
	return this.collectionId
}

func (this *SceneCollection) GetEndTime() int64 {
	if this.collectionTime > 0 {
		return this.collectionTime + int64(gamedb.GetCollectionCollectionCfg(this.collectionId).Success)
	}
	return 0
}

func (this *SceneCollection) BuildSceneObjMessage() nw.ProtoMessage {
	r := &pb.SceneObj{}
	r.ObjType = int32(this.GetType())
	r.Point = this.Point().ToPbPoint()
	r.ObjId = int32(this.GetObjId())
	r.Collection = &pb.SceneCollection{
		Id:              int32(this.collectionId),
		CollectionObjId: int32(this.collectionObjId),
		ServerTime:      int32(time.Now().Unix()),
		EndTime:         int32(this.GetEndTime()),
	}
	return r
}

func (this *SceneCollection) BuildAppearMessage() nw.ProtoMessage {
	appearNtf := &pb.SceneEnterNtf{}
	appearNtf.Objs = append(appearNtf.Objs, this.BuildSceneObjMessage().(*pb.SceneObj))
	return appearNtf
}
