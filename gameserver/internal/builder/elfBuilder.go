package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildElf(user *objs.User) *pb.Elf {
	userElf := user.Elf
	return &pb.Elf{
		Lv:           int32(userElf.Lv),
		Exp:          int32(userElf.Exp),
		Skills:       BuildElfSkills(userElf.Skills),
		SkillBag:     BuildElfSkillBag(userElf.SkillBag),
		ReceiveLimit: BuildElfReceiveLimit(userElf.RecoverLimit),
	}
}

func BuildElfSkills(skills model.IntKv) map[int32]int32 {
	pbMap := make(map[int32]int32, 0)
	for id, lv := range skills {
		pbMap[int32(id)] = int32(lv)
	}
	return pbMap
}

func BuildElfSkillBag(skillBag model.IntKv) map[int32]int32 {
	pbMap := make(map[int32]int32)
	for pos, id := range skillBag {
		pbMap[int32(pos)] = int32(id)
	}
	return pbMap
}

func BuildElfReceiveLimit(items map[int]int) map[int32]int32 {
	pbMap := make(map[int32]int32)
	if items != nil {
		for itemId, count := range items {
			pbMap[int32(itemId)] = int32(count)
		}
	}
	return pbMap
}
