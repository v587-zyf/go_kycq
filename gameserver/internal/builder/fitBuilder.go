package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildFit(user *objs.User) *pb.Fit {
	userFit := user.Fit
	return &pb.Fit{
		CdStart:  int64(userFit.CdStart),
		CdEnd:    int64(userFit.CdEnd),
		Fashion:  BuildMapInt32(userFit.Fashion),
		SkillBag: BuildMapInt32(userFit.SkillBag),
		Lv:       BuildMapInt32(userFit.Lv),
		Skills:   BuildFitSkill(userFit.Skills),
	}
}

func BuildMapInt32(maps model.IntKv) map[int32]int32 {
	pbMap := make(map[int32]int32)
	for k, v := range maps {
		pbMap[int32(k)] = int32(v)
	}
	return pbMap
}

func BuildFitSkill(skill map[int]*model.FitSkill) map[int32]*pb.FitSkill {
	pbMap := make(map[int32]*pb.FitSkill)
	for skillId, fitSkill := range skill {
		pbMap[int32(skillId)] = &pb.FitSkill{
			Lv:   int32(fitSkill.Lv),
			Star: int32(fitSkill.Star),
		}
	}
	return pbMap
}

func BuildFitHolyEquip(user *objs.User) *pb.FitHolyEquip {
	userFitHolyEquip := user.FitHolyEquip
	return &pb.FitHolyEquip{
		SuitId: int32(userFitHolyEquip.SuitId),
		Equips: BuildFitHolyEquips(userFitHolyEquip.Equips),
	}
}
func BuildFitHolyEquips(equips model.MapIntKv) map[int32]*pb.FitHolyEquipUnit {
	pbMap := make(map[int32]*pb.FitHolyEquipUnit)
	for t, intKv := range equips {
		pbMap[int32(t)] = BuildFitHolyEquipUnit(intKv)
	}
	return pbMap
}
func BuildFitHolyEquipUnit(equips model.IntKv) *pb.FitHolyEquipUnit {
	pbMap := make(map[int32]int32)
	userFitHolyEquip := equips
	for pos, id := range userFitHolyEquip {
		pbMap[int32(pos)] = int32(id)
	}
	return &pb.FitHolyEquipUnit{Equip: pbMap}
}
