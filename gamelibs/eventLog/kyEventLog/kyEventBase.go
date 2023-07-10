package kyEventLog

import (
	"cqserver/golibs/logger"
	"encoding/json"
)

var eventLog = logger.Get("event", true)

func init() {
	eventLog.EnableFuncCallDepth(false)
}

type KyEventBase struct {
	Ouid       string      `json:"ouid"`       //账号
	Timestamp  int         `json:"timestamp"`  //时间（毫秒）
	Did        string      `json:"did"`        //设备号
	Event      string      `json:"event"`      //事件名称
	Project    string      `json:"project"`    //项目
	Properties interface{} `json:"properties"` //事件数据
}
type KyEventPropBase struct {
	Terminal   string `json:"terminal"`
	Plat       int    `json:"plat"`       // 平台    数值    是
	Channel    string `json:"channel"`    // 渠道    字符    是
	Serverid   int    `json:"_serverid"`  // 服务器Id
	Source_sid int    `json:"source_sid"` //合服前的区服id    数值    是
	Openid     string `json:"openid"`     //平台账号id    字符    是
	Roleid     string `json:"roleid"`     //角色id    字符    是
	Rolename   string `json:"rolename"`   //角色名称    字符    是
	Vip        int    `json:"vip"`        //vip等级    数值    是
	Level      int    `json:"level"`      //等级    数值    是
	Combat     int    `json:"combat"`     //战斗力
	GuildId    int    `json:"guild_id"`
}

func WriteEvent(msg string) {
	eventLog.Write([]byte(msg ))
}

func WriteEventByData(data interface{}) {

	dataJson, err := json.Marshal(data)
	if err != nil {
		logger.Error("记录事件数据json marshal 错误：%v", err)
		return
	}
	eventLog.Write(dataJson)
}
