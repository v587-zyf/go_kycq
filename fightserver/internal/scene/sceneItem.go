package scene

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"time"
)

type SceneItem struct {
	*DefaultSceneObj
	itemId             int
	num                int
	dropTime           time.Time //掉落时间
	disappearTime      int       //消失时间
	owner              int       //玩家的userId
	ownerName          string    //归属玩家名字
	dropMonsterId      int       //掉落怪物Id
	ownerProtectedTime int       //归属者保护时间
}

func NewSceneItem(ownerUserId int, ownerName string, dropMonsterId, itemId, num int, disappearTime int) *SceneItem {

	if disappearTime <= 0 {
		disappearTime = 300
	}
	dropItem := &SceneItem{
		owner:              ownerUserId,
		ownerName:          ownerName,
		dropMonsterId:      dropMonsterId,
		itemId:             itemId,
		num:                num,
		dropTime:           time.Now(),
		disappearTime:      int(time.Now().Unix()) + disappearTime,
		ownerProtectedTime: int(time.Now().Unix()) + gamedb.GetConf().DropItemOwnerProtectTime,
	}
	dropItem.DefaultSceneObj = NewDefaultSceneObj(dropItem, pb.SCENEOBJTYPE_ITEM)
	return dropItem
}

func (this *SceneItem) Num() int {
	return this.num
}

func (this *SceneItem) ItemId() int {
	return this.itemId
}

func (this *SceneItem) DropTime() string {
	return common.GetFormatTime2(this.dropTime)
}

func (this *SceneItem) OwnerName() string {
	return this.ownerName
}

func (this *SceneItem) DropMonsterId() int {
	return this.dropMonsterId
}

func (this *SceneItem) CanPickUp(pickuserId int) bool {

	now := int(time.Now().Unix())
	if now > this.disappearTime {
		logger.Error("物品不能被拾取，物品已消失：%v,当前时间：%v,物品消失时间：%v", this.objId, now, this.disappearTime)
		return false
	}

	if now < this.ownerProtectedTime && pickuserId != this.owner {
		logger.Error("物品不能被拾取，物品不归属玩家：%v,拾取者：%v,归属者：%v,当前时间：%v,物品保护时间：%v", this.objId, pickuserId, this.owner, now, this.ownerProtectedTime)
		return false
	}

	return true
}

func (this *SceneItem) disappeared(now int) bool {

	if now > this.disappearTime {
		return true
	}
	return false
}

func (this *SceneItem) BuildSceneObjMessage() nw.ProtoMessage {
	r := &pb.SceneObj{}
	r.ObjType = int32(this.GetType())
	r.Point = this.Point().ToPbPoint()
	r.ObjId = int32(this.GetObjId())
	r.Item = &pb.SceneItem{
		ItemId:             int32(this.itemId),
		ItemNum:            int32(this.num),
		Owner:              int32(this.owner),
		OwnerProtectedTime: int32(this.ownerProtectedTime),
		DisappearTime:      int32(this.disappearTime),
	}
	return r
}

func (this *SceneItem) BuildAppearMessage() nw.ProtoMessage {
	appearNtf := &pb.SceneEnterNtf{}
	appearNtf.Objs = append(appearNtf.Objs, this.BuildSceneObjMessage().(*pb.SceneObj))
	return appearNtf
}
