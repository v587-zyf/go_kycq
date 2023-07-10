package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuilderPanacea(user *objs.User) map[int32]*pb.PanaceaInfo {
	pbPanaceas := make(map[int32]*pb.PanaceaInfo)
	for id, panaceaUnit := range user.Panaceas {
		pbPanaceas[int32(id)] = BuilderPanaceaInfo(panaceaUnit)
	}
	return pbPanaceas
}

func BuilderPanaceaInfo(panaceaUnit *model.PanaceaUnit) *pb.PanaceaInfo {
	return &pb.PanaceaInfo{
		Numbers: int32(panaceaUnit.Numbers),
		Number:  int32(panaceaUnit.Number),
	}
}

