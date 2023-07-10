package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuilderFabao(user *objs.User) []*pb.Fabao {
	fabaos := make([]*pb.Fabao, 0)
	for _, v := range user.Fabaos {
		fabaos = append(fabaos, BuilderFabaoUnit(v))
	}
	return fabaos
}

func BuilderFabaoUnit(fabao *model.Fabao) *pb.Fabao {
	var skills []int32
	if len(fabao.Skill) > 0 {
		for _, v := range fabao.Skill {
			skills = append(skills, int32(v))
		}
	}
	return &pb.Fabao{
		Id:     int32(fabao.Id),
		Level:  int32(fabao.Level),
		Exp:    int32(fabao.Exp),
		Skills: skills,
	}
}
