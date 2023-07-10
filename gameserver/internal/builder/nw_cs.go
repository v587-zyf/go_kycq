package builder

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pbserver"
)

func BuildSyncUserInfoNtf(user *objs.User, status int, lastRechargeTime int64) *pbserver.SyncUserInfoNtf {

	msg := &pbserver.SyncUserInfoNtf{
		UserId:           int32(user.Id),
		ServerId:         int32(user.ServerId),
		Nickname:         user.NickName,
		Vip:              int32(user.VipLevel),
		Combat:           int64(user.Combat),
		OpenId:           user.OpenId,
		ServerIndex:      int32(user.ServerIndex),
		CreateTime:       int32(user.CreateTime.Unix()),
		OfflineTime:      int32(user.OfflineTime.Unix()),
		Avatar:           user.Avatar,
		ChannelId:        int32(user.ChannelId),
		Recharge:         int32(user.RechargeAll),
		Gold:             int64(user.Gold),
		Ingot:            int32(user.Ingot),
		TaskId:           int32(user.MainLineTask.TaskId),
		LastRechargeTime: int32(lastRechargeTime),
		SyscStatus:       int32(status),
		Heros:            make([]*pbserver.SyscHeroInfo, len(user.Heros)),
		Exp:              int64(user.Exp),
		LoginTime:        int32(user.LoginTime.Unix()),
	}
	index := 0
	for _, v := range user.Heros {
		hero := &pbserver.SyscHeroInfo{
			HeroIndex: int32(v.Index),
			Sex:       int32(v.Sex),
			Level:     int32(v.ExpLvl),
			Job:       int32(v.Job),
		}
		msg.Heros[index] = hero
		index++
	}
	return msg
}
