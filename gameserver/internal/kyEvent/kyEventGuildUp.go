package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_guild_up struct {
	*kyEventLog.KyEventPropBase
	GuildId     int    `json:"guild_id"`
	GuildName   string `json:"guild_name"`
	GuildNumber int    `json:"num1"`
	GuildLv     int    `json:"num2"` //公会贡献值
}

//公会升级记录
func GuildUp(user *objs.User, guildName string, guildId, guildNumber, guildLv int) {
	data := &KyEvent_guild_up{
		KyEventPropBase: getEventBaseProp(user),
		GuildId:         guildId,
		GuildName:       guildName,
		GuildNumber:     guildNumber,
		GuildLv:         guildLv,
	}
	writeUserEvent(user, "guild_up", data)
}
