package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuilderHolyArms(user *objs.User) []*pb.Holy {
	pbHoly := make([]*pb.Holy, 0)
	for id, holyarm := range user.Holyarms {
		pbHoly = append(pbHoly, BuilderHoly(id, holyarm))
	}
	return pbHoly
}

func BuilderHoly(id int, holyarm *model.Holyarm) *pb.Holy {
	return &pb.Holy{
		Id:     int32(id),
		Level:  int32(holyarm.Level),
		Exp:    int32(holyarm.Exp),
		Skills: BuilderHolySkill(holyarm.Skill),
	}
}

func BuilderHolySkill(skill map[int]int) map[int32]int32 {
	pbSkill := make(map[int32]int32)
	for id, lv := range skill {
		pbSkill[int32(id)] = int32(lv)
	}
	return pbSkill
}
