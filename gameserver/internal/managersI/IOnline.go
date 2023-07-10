package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type IOnline interface {
	Online(user *objs.User)
	ResetOnline(user *objs.User)
	//玩家下线，记录玩家今天在线时间
	OffLine(user *objs.User,isOffLine bool)
	//发放奖励
	GetOnlineAward(user *objs.User, awardId int, op *ophelper.OpBagHelperDefault) error

}
