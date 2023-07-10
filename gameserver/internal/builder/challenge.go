package builder

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gameserver/internal/objs"
)

func BuildApplyChallengeUserInfoAck(user *objs.User, crossFsId int, serverName, guildName string) *modelCross.Challenge {

	ack := &modelCross.Challenge{}
	ack.UserId = user.Id
	ack.ServerId = user.ServerId
	ack.OpenId = user.OpenId
	ack.Combat = int64(user.Combat)
	ack.OpenId = user.OpenId
	ack.NickName = user.NickName
	ack.Avatar = user.Avatar
	ack.CrossFsId = crossFsId
	ack.GuildName = guildName
	return ack
}
