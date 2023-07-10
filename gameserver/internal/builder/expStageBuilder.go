package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuilderExpStage(user *objs.User) *pb.ExpStage {
	expStage := user.ExpStage
	return &pb.ExpStage{
		DareNum:   int32(expStage.DareNum),
		BuyNum:    int32(expStage.BuyNum),
		ExpStages: BuildExpStages(expStage.ExpStages),
		Appraise:  BuildAppraise(expStage.Appraise),
		Layer:     int32(expStage.Layer),
	}
}

func BuildExpStages(expStages model.IntKv) map[int32]int64 {
	pbExpStages := make(map[int32]int64)
	for stageId, exp := range expStages {
		pbExpStages[int32(stageId)] = int64(exp)
	}
	return pbExpStages
}

func BuildAppraise(appraise model.IntKv) map[int32]int32 {
	pbMap := make(map[int32]int32)
	for stageId, appr := range appraise {
		pbMap[int32(stageId)] = int32(appr)
	}
	return pbMap
}
