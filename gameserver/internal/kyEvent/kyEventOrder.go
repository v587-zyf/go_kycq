package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constOrder"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_startOrder struct {
	*kyEventLog.KyEventPropBase
	Order_id      string `json:"order_id"`      //订单id	字符	是
	Pay_amount    int    `json:"pay_amount"`    //充值金额	数值	是	分
	Order_channel int    `json:"order_channel"` //充值渠道	数值	是
	Diamond       int    `json:"diamond"`       //剩余付费货币数量	数值	是
	OrderType     int    `json:"type"`          //订单类型(1:正式 2:福利充值(测试))	数值	是
	Client_ip     string `json:"client_ip"`     //用户的客户端ip	字符	是
	PayModule     int    `json:"main_type"`     //充值模块
	PayModuleId   int    `json:"config_id1"`    //充值模块Id
}

type KyEvent_order struct {
	*kyEventLog.KyEventPropBase
	Order_id      string `json:"order_id"`      //订单id	字符	是
	Id            string `json:"id"`            //平台订单id	字符	是
	Pay_amount    int    `json:"pay_amount"`    //充值金额	数值	是	分
	Order_channel int    `json:"order_channel"` //充值渠道	数值	是
	Diamond       int    `json:"diamond"`       //剩余付费货币数量	数值	是
	OrderType     int    `json:"type"`          //订单类型(1:正式 2:福利充值(测试))	数值	是
	Client_ip     string `json:"client_ip"`     //用户的客户端ip	字符	是
	PayModule     int    `json:"main_type"`     //充值模块
	PayModuleId   int    `json:"config_id1"`    //充值模块Id
}

func UserStartOrder(user *objs.User, order *modelGame.OrderDb) {
	data := &KyEvent_startOrder{
		KyEventPropBase: getEventBaseProp(user),
		Order_id:        order.OrderNo,
		Pay_amount:      order.PayMoney * 100,
		Order_channel:   0,
		Diamond:         user.Ingot,
		Client_ip:       user.Ip,
		PayModule:       order.PayModule,
		PayModuleId:     order.PayModuleId,
	}
	if order.IsPayToken == constOrder.PAY_TOKEN_NO {
		data.OrderType = 1
	} else {
		data.OrderType = 2
	}
	writeUserEvent(user, "start_order", data)
}

func UserOrder(user *objs.User, order *modelGame.OrderDb) {
	data := &KyEvent_order{
		KyEventPropBase: getEventBaseProp(user),
		Order_id:        order.OrderNo,
		Id:              order.PlatformOrderNo,
		Pay_amount:      order.PayMoney * 100,
		Order_channel:   0,
		Diamond:         user.Ingot,
		Client_ip:       user.Ip,
		PayModule:       order.PayModule,
		PayModuleId:     order.PayModuleId,
	}
	if order.IsPayToken == constOrder.PAY_TOKEN_NO {
		data.OrderType = 1
	} else {
		data.OrderType = 2
	}
	writeUserEvent(user, "order", data)
}
