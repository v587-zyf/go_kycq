package base

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/protobuf/pb"
)

type Buff interface {
	GetBuffT() *gamedb.BuffBuffCfg
	GetType() int
	GetSource() Actor
	GetOwenr() Actor
	GetStartTime() int64
	GetEndTime() int64
	OnAdd(buffHpChangeInfos *[]*pb.BuffHpChangeInfo, delBuffInfos *[]*pb.DelBuffInfo)
	OnRemove()
	IsExpire(nowTime int64) bool
	Run(buffHpChangeInfos *[]*pb.BuffHpChangeInfo,arg ...interface{})
	GetBuffIdx() int32
}
