package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_mining struct {
	*kyEventLog.KyEventPropBase
	Miner   int         `json:"num1"`    //矿工等级
	Reward  map[int]int `json:"info"`   //挖矿收益
	WorkNum int         `json:"num2"` //挖矿次数
}

//挖矿记录
func Mining(user *objs.User, miner, workNum int, reward map[int]int) {
	data := &KyEvent_mining{
		KyEventPropBase: getEventBaseProp(user),
		Miner:           miner,
		Reward:          reward,
		WorkNum:         workNum,
	}
	writeUserEvent(user, "mining", data)
}
