package user

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
)

func (this *UserManager) SendMessage(user *objs.User, msg nw.ProtoMessage, sendNow bool) error {
	err := this.PutOutMessage(user.GateSessionId, msg, sendNow)
	if err != nil {
		logger.Warn("给玩家发送消息，session为空,玩家：%v,msg:%v", user.Id, msg)
		return err
	}
	return nil
}

func (this *UserManager) SendItemChangeNtf(user *objs.User, op *ophelper.OpBagHelperDefault) error {
	itemChangeNtf := op.ToGoodsChangeMessages()
	for _, v := range itemChangeNtf {
		this.SendMessage(user, v, false)
	}
	return nil
}

func (this *UserManager) SendMessageByUserId(userId int, msg nw.ProtoMessage) error {
	user := this.GetUser(userId)
	if user != nil {
		this.SendMessage(user, msg, true)
	}
	return nil
}

func (this *UserManager) SendMsgToIds(ids []int32, data []byte) {

	sessionids := make([]uint32, 0)
	for _, id := range ids {
		user := this.GetUserManager().GetUser(int(id))
		if user != nil {
			sessionids = append(sessionids, user.GateSessionId)
		}
	}
	if len(sessionids) > 0 {
		this.BroadcastData(sessionids, data)
	}
}
