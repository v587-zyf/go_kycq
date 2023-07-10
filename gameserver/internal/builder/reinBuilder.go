package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/protobuf/pb"
)

func BuilderRein(rein *model.Rein) *pb.Rein {
	return &pb.Rein{
		Id:  int32(rein.Id),
		Exp: int64(rein.Exp),
	}
}

func BuilderReinCost(reinCost *model.ReinCost) *pb.ReinCost {
	return &pb.ReinCost{
		Id:  int32(reinCost.Id),
		Num: int32(reinCost.Num),
	}
}

