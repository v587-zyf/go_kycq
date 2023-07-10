package scene

import (
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

type SceneBuff struct {
	*DefaultSceneObj
	buffId      int
	ownerUserId int
}

func NewSceneBuff(buffId int, ownerUserId int) *SceneBuff {

	groundBuff := &SceneBuff{
		buffId:      buffId,
		ownerUserId: ownerUserId,
	}
	groundBuff.DefaultSceneObj = NewDefaultSceneObj(groundBuff, pb.SCENEOBJTYPE_BUFF)
	return groundBuff
}

func (this *SceneBuff) BuildSceneObjMessage() nw.ProtoMessage {
	r := &pb.SceneObj{}
	r.ObjType = int32(this.GetType())
	r.Point = this.Point().ToPbPoint()
	r.ObjId = int32(this.GetObjId())
	r.Buff = &pb.SceneBuff{
		BuffId: int32(this.buffId),
		UserId: int32(this.ownerUserId),
	}
	return r
}

func (this *SceneBuff) BuildAppearMessage() nw.ProtoMessage {
	appearNtf := &pb.SceneEnterNtf{}
	appearNtf.Objs = append(appearNtf.Objs, this.BuildSceneObjMessage().(*pb.SceneObj))
	return appearNtf
}
