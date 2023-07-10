package builder

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuilderFashion( hero *objs.Hero ) map[int32]*pb.Fashion {
	fashionPb := make(map[int32]*pb.Fashion)
	for _,v := range hero.Fashions {
		fashionPb[int32(v.Id)] = &pb.Fashion{
			Id: int32(v.Id),
			Level: int32(v.Lv),
		}
	}
	return fashionPb
}
