package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

type IWorldBoss interface {
	// 挑战世界Boss
	EnterWorldBossFight(user *objs.User, stageId int) error
	// 战斗之后，回调结果
	WorldBossFightEndAck(rank []int, lucker, stageId int) error
	// 推送世界Boss状态
	WorldBossInfoNtf() error

	GetWorldBossInfo() *pb.WorldBossInfoNtf
}
