package builder

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuilderAtlas(user *objs.User) []*pb.Atlas {
	pbAtlas := make([]*pb.Atlas, 0)
	for id, star := range user.Atlases {
		pbAtlas = append(pbAtlas, BuilderAtlasUnit(id, star))
	}
	return pbAtlas
}

func BuilderAtlasUnit(id int, star int) *pb.Atlas {
	return &pb.Atlas{
		Id:       int32(id),
		Star:     int32(star),
		IsActive: true,
	}
}

func BuilderAtlasGather(user *objs.User) []*pb.AtlasGather {
	atlasGathers := make([]*pb.AtlasGather, 0)
	for id, star := range user.AtlasGathers {
		atlasGathers = append(atlasGathers, BuilderAtlasGatherUnit(id, star))
	}
	return atlasGathers
}

func BuilderAtlasGatherUnit(id int, star int) *pb.AtlasGather {
	return &pb.AtlasGather{
		Id:       int32(id),
		Star:     int32(star),
		IsActive: true,
	}
}