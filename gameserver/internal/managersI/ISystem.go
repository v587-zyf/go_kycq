package managersI

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
	"time"
)

type ISystemManager interface {
	ServerIdInLocalServer(serverId int) bool
	GetServerIndex(serverId int) int
	GetServerOpenTimeByServerId(serverId int) time.Time
	GetServerOpenDaysByServerId(serverId int) int
	GetServerMergeDayByServerId(serverId int) int
	GetServerName(serverId int) string
	GetPrefix() string
	GetCrossFightServerId() int
	IsCross() bool
	IsMerge() bool
	PreferenceSet(user *objs.User, preference []*pb.Preference, ack *pb.PreferenceSetAck) error
	PreferenceLoad(user *objs.User, ack *pb.PreferenceLoadAck) error
	GetServerOpenDaysByServerIdByExcursionTime(serverId int, reduceHour time.Duration) int //当前时间减多少小时 取开服天数
	GetServerIndexCrossFsId(serverId int) int
	UpdateFuncState(sendClient bool)
	GetFuncState() []int
	GetServerMergerIdAndMergerTime(serverId int) (int, time.Time)
	GetMergerServerOpenDaysByServerId(serverId int) int
	GetServerInfoByServerId(serverId int) *modelCross.ServerInfo
	GetCrossServerBriefUserInfo() map[int32]*pb.BriefServerInfo
}
