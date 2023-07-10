package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildMaterialStage(user *objs.User) map[int32]*pb.MaterialStage {
	userMaterial := user.MaterialStage
	pbMaterial := make(map[int32]*pb.MaterialStage)
	for _, mateType := range pb.MATERIALSTAGETYPE_ARRAY {
		pbMaterial[int32(mateType)] = BuildMaterialStageUnit(userMaterial, mateType)
	}
	return pbMaterial
}

func BuildMaterialStageUnit(userMaterial *model.MaterialStage, mateType int) *pb.MaterialStage {
	mateInfo := userMaterial.MaterialStages[mateType]
	return &pb.MaterialStage{
		DareNum:   int32(mateInfo.DareNum),
		BuyNum:    int32(mateInfo.BuyNum),
		NowLayer:  int32(mateInfo.NowLayer),
		LastLayer: int32(mateInfo.LastLayer),
	}
}
