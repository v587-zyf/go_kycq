package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_guild_create struct {
	*kyEventLog.KyEventPropBase
	GuildId   int    `json:"guild_id"`
	GuildName string `json:"guild_name"`
}

//公会创建记录
func GuildCreate(user *objs.User, guildId int, guildName string) {
	data := &KyEvent_guild_create{
		KyEventPropBase: getEventBaseProp(user),
		GuildId:         guildId,
		GuildName:       guildName,
	}
	writeUserEvent(user, "guild_create", data)
}
