package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuilderUserJuexue(user *objs.User) []*pb.Juexue {
	juexues := make([]*pb.Juexue, 0)
	for _, godEquip := range user.Juexues {
		juexues = append(juexues, BuilderJuexue(godEquip))
	}
	return juexues
}

func BuilderJuexue(juexue *model.Juexue) *pb.Juexue {
	return &pb.Juexue{
		Id:    int32(juexue.Id),
		Level: int32(juexue.Lv),
	}
}
