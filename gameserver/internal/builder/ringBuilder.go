package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildRings(hero *objs.Hero) map[int32]*pb.Ring {
	pbRing := make(map[int32]*pb.Ring)
	for pos, unit := range hero.Ring {
		pbRing[int32(pos)] = BuildRing(unit)
	}
	return pbRing
}

func BuildRing(unit *model.RingUnit) *pb.Ring {
	return &pb.Ring{
		Rid:        int32(unit.Rid),
		Strengthen: int32(unit.Strengthen),
		Pid:        int32(unit.Pid),
		Talent:     int32(unit.Talent),
		Phantom:    BuildRingPhantom(unit.Phantom),
	}
}

func BuildRingPhantom(ringPhantom map[int]*model.RingPhantom) map[int32]*pb.RingPhantom {
	pbRingPhantom := make(map[int32]*pb.RingPhantom)
	for pos, phantom := range ringPhantom {
		pbRingPhantom[int32(pos)] = &pb.RingPhantom{
			Talent:  int32(phantom.Talent),
			Phantom: int32(phantom.Phantom),
			Skill:   BuildRingSkill(phantom.Skill),
		}
	}
	return pbRingPhantom
}

func BuildRingSkill(skill map[int]int) map[int32]int32 {
	pbSkill := make(map[int32]int32)
	for id, lv := range skill {
		pbSkill[int32(id)] = int32(lv)
	}
	return pbSkill
}
