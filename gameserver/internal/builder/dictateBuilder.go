package builder

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuilderDictates(hero *objs.Hero) []*pb.DictateInfo {
	for _, v := range pb.DICTATETYPE_ARRAY {
		if _, ok := hero.Dictates[v]; !ok {
			hero.Dictates[v] = 0
		}
	}

	pbDictateInfos := make([]*pb.DictateInfo, 0)
	for body, lv := range hero.Dictates {
		pbDictateInfos = append(pbDictateInfos, BuildDictateInfo(body, lv))
	}
	return pbDictateInfos
}

func BuildDictateInfo(body, lv int) *pb.DictateInfo {
	return &pb.DictateInfo{
		Type:  int32(body),
		Level: int32(lv),
	}
}
