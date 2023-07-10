package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildTalent(hero *objs.Hero) *pb.TalentInfo {
	heroTalent := hero.Talent
	return &pb.TalentInfo{
		GetPoints:     int32(heroTalent.GetPoints),
		SurplusPoints: int32(heroTalent.SurplusPoints),
		Talents:       BuildTalentUnit(heroTalent.TalentList),
	}
}

func BuildTalentUnit(talents map[int]*model.TalentUnit) map[int32]*pb.TalentUnit {
	pbTalents := make(map[int32]*pb.TalentUnit)
	for id, talentUnit := range talents {
		pbTalents[int32(id)] = &pb.TalentUnit{
			UsePoints: int32(talentUnit.UsePoints),
			Talents:   BuildTalentUnitMap(talentUnit.Talents),
		}
	}
	return pbTalents
}

func BuildTalentUnitMap(talents map[int]int) map[int32]int32 {
	pbTalents := make(map[int32]int32)
	for id, lv := range talents {
		pbTalents[int32(id)] = int32(lv)
	}
	return pbTalents
}
