package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildWings(hero *objs.Hero) []*pb.Wing {
	wing := make([]*pb.Wing, 0)
	for _, v := range hero.Wings {
		if v.Id > 0 {
			wing = append(wing, BuildWing(v))
		}
	}
	return wing
}

func BuildWing(wing *model.Wing) *pb.Wing {
	return &pb.Wing{
		Id:  int32(wing.Id),
		Exp: int32(wing.Exp),
	}
}

func BuilderWingSpecials(hero *objs.Hero) []*pb.WingSpecialNtf {
	pbWingSpecials := make([]*pb.WingSpecialNtf, 0)
	for specialT, lv := range hero.WingSpecial {
		pbWingSpecials = append(pbWingSpecials, BuilderWingSpecial(specialT, lv))
	}
	return pbWingSpecials
}

func BuilderWingSpecial(specialT, level int) *pb.WingSpecialNtf {
	return &pb.WingSpecialNtf{
		SpecialType: int32(specialT),
		Level:       int32(level),
	}
}
