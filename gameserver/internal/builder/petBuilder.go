package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildPet(user *objs.User) map[int32]*pb.PetInfo {
	pbPetInfo := make(map[int32]*pb.PetInfo)
	for id, pet := range user.Pet {
		pbPetInfo[int32(id)] = BuildPetUnit(pet)
	}
	return pbPetInfo
}

func BuildPetUnit(pet *model.Pet) *pb.PetInfo {
	return &pb.PetInfo{
		Lv:    int32(pet.Lv),
		Exp:   int32(pet.Exp),
		Grade: int32(pet.Grade),
		Break: int32(pet.Break),
		Skill: BuildPetSkill(pet.Skill),
	}
}

func BuildPetSkill(skill map[int]int) []int32 {
	pbSkill := make([]int32, 0)
	for id := range skill {
		pbSkill = append(pbSkill, int32(id))
	}
	return pbSkill
}

func BuildPetAppendage(user *objs.User) map[int32]int32 {
	appendage := user.PetAppendage
	pbMap := make(map[int32]int32)
	for pid, lv := range appendage {
		pbMap[int32(pid)] = int32(lv)
	}
	return pbMap
}
