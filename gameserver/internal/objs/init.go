package objs

import (
	"cqserver/golibs/common"
)

const (
	SyncStatusHeroChangeNtf common.Bitmask = 1 << iota
	SyncStatusDisplayidNtf
	SyncStatusTaskInfoNtf
	SyncStatusSceneUserNtf
	SyncStatusSceneHeroNtf
	SyncStatusComatNtf
)
