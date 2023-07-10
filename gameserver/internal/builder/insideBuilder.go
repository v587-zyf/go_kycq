package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildInsideInfo(hero *objs.Hero) *pb.InsideInfo {
	inside := hero.Inside
	return &pb.InsideInfo{
		Acupoint:    BuildAcupoint(inside.Acupoint),
		InsideSkill: BuildInsideSkill(inside.Skill),
	}
}

func BuildAcupoint(acupoint map[int]int) map[int32]int32 {
	pbAcupoint := make(map[int32]int32)
	for id, lv := range acupoint {
		pbAcupoint[int32(id)] = int32(lv)
	}
	return pbAcupoint
}

func BuildInsideSkill(skill map[int]*model.InsideSkill) map[int32]*pb.InsideSkill {
	pbInsideSkill := make(map[int32]*pb.InsideSkill)
	for id, skill := range skill {
		pbInsideSkill[int32(id)] = &pb.InsideSkill{
			Level: int32(skill.Level),
			Exp:   int32(skill.Exp),
		}
	}
	return pbInsideSkill
}
