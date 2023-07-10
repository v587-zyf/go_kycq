package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildSkillBag(hero *objs.Hero, skillT int) map[int32]int32 {
	var skillBag model.IntKv
	if skillT == pb.SKILLTYPE_UNIQUE {
		skillBag = hero.UniqueSkillBag
	} else {
		skillBag = hero.SkillBag
	}
	pbSkillBag := make(map[int32]int32)
	for pos, skillId := range skillBag {
		pbSkillBag[int32(pos)] = int32(skillId)
	}
	return pbSkillBag
}

func BuildSkills(hero *objs.Hero, skillT int) []*pb.SkillUnit {
	var skills map[int]*model.SkillUnit
	if skillT == pb.SKILLTYPE_UNIQUE {
		skills = hero.UniqueSkills
	} else {
		skills = hero.Skills
	}
	pbSkillUnits := make([]*pb.SkillUnit, 0)
	for _, unit := range skills {
		pbSkillUnits = append(pbSkillUnits, BuildSkillUnit(unit))
	}
	return pbSkillUnits
}

func BuildSkillUnit(skillUnit *model.SkillUnit) *pb.SkillUnit {
	return &pb.SkillUnit{
		SkillId:   int32(skillUnit.Id),
		Level:     int32(skillUnit.Lv),
		StartTime: skillUnit.StartTime,
		EndTime:   skillUnit.EndTime,
	}
}

func BuildAncientSkill(hero *objs.Hero) *pb.AncientSkill {
	heroAncientSkill := hero.AncientSkill
	return &pb.AncientSkill{
		SkillId: int32(heroAncientSkill.SkillId),
		Level:   int32(heroAncientSkill.Level),
		Grade:   int32(heroAncientSkill.Grade),
	}
}
