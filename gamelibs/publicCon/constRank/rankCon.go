package constRank

import "cqserver/protobuf/pb"

const (
	MAX       = 100
	MAX_ARENA = 100
	RANK_SORT_TIME_FIX = 10000000000
)

var (
	SORT_TIME_RANK = map[int]bool{
		pb.RANKTYPE_TOWER: true,
	}
)
