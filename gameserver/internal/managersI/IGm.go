package managersI

import "cqserver/gameserver/internal/objs"

type IGmManager interface {

	GmChatCode(user *objs.User,codes string)string
}
