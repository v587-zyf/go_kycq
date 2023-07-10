package httpManager

import (
	"cqserver/gamelibs/errex"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/ptsdk"
	"cqserver/gamelibs/publicCon/constPlatfrom"
	"cqserver/gamelibs/rmodelCross"
	"cqserver/golibs/common"
	"cqserver/golibs/nw/httpclient"
	"fmt"

	"cqserver/golibs/logger"
	"cqserver/protobuf/pbserver"
	"encoding/json"

	"net/http"
	"runtime/debug"
	"strconv"
)

func httpRechage(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic HandleLoginReq:%v,%s", r, stackBytes)
		}
	}()

	r.ParseForm()
	logger.Info("接收到http消息，充值请求：%v", r.Form)

	platfrom := ptsdk.GetSdk().GetPlatform()
	switch platfrom {
	case constPlatfrom.PLATFORM_602:
		rechage602(w, r)
	case constPlatfrom.PLATFORM_JUNSHANG:
		rechageJunshang(w, r)
	default:
		notifierRecharge(w, r)
	}
}

/**
*  @Description:
*  @param serverid
*  @return errMsg		返回错误消息，成功返回空字符串
**/
func callGameRechageResult(serverid int, orderData *pbserver.RechageCcsToGsReq) (int, string) {
	applyMsg := orderData
	replayMsg := &pbserver.RechageGsToCcsAck{}
	err := m.GetGsServers().CallMessage(serverid, applyMsg, replayMsg)
	if err != nil {
		logger.Error("发送game充值异常：%v", err)
		return gamedb.ERRUNKNOW.Code, gamedb.ERRUNKNOW.Message
	}
	logger.Info("收到game充值结果：%v", *replayMsg)
	if replayMsg.Result != ptsdk.STATUS_SUCCESS {
		code, _ := strconv.Atoi(replayMsg.Msg)
		return code, gamedb.GetGameTextErrorTextCfg(code).Chinese
	}
	return 0, ""
}

func rechage602(w http.ResponseWriter, r *http.Request) {

}

func rechageJunshang(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]string)
	param["pid"] = r.FormValue("pid")         //君尚平台联运商ID，详细见运营参数表格
	param["gid"] = r.FormValue("gid")         //君尚平台游戏ID，详细见运营参数表格
	param["time"] = r.FormValue("time")       //	Int Unix时间戳
	param["sign"] = r.FormValue("sign")       //	String
	param["oid"] = r.FormValue("oid")         //联运平台订单ID	String
	param["doid"] = r.FormValue("doid")       //	CP订单ID String
	param["dsid"] = r.FormValue("dsid")       //	CP游戏服ID String
	param["drid"] = r.FormValue("drid")       //CP角色ID
	param["drname"] = r.FormValue("drname")   //CP角色名
	param["drlevel"] = r.FormValue("drlevel") //CP角色等级
	param["uid"] = r.FormValue("uid")         //君尚平台用户UID
	param["money"] = r.FormValue("money")     //金额,例如:1.00，单位元
	param["coin"] = r.FormValue("coin")       //游戏,例如:10
	param["remark"] = r.FormValue("remark")   //简单的备注
	param["paid"] = r.FormValue("paid")       //君尚平台应用标识

	ret := &ptsdk.SdkJunshangLoginResult{
		State: ptsdk.STATUS_SUCCESS,
		Msg:   "",
	}
	if len(param["time"]) <= 0 || len(param["oid"]) <= 0 || len(param["doid"]) <= 0 || len(param["dsid"]) <= 0 || len(param["uid"]) <= 0 || len(param["money"]) <= 0 || len(param["coin"]) <= 0 {
		ret.State = ptsdk.STATUS_FAIL
		ret.Msg = ptsdk.PARAM_ERR
		msg, _ := json.Marshal(ret)
		w.Write([]byte(msg))
		logger.Error("充值异常，缺少参数：%v", param)
		return
	}
	if !ptsdk.GetSdk().CheckSignForPlatform(param["sign"], param) {
		ret.State = ptsdk.STATUS_FAIL
		ret.Msg = ptsdk.SIGN_ERR
		msg, _ := json.Marshal(ret)
		w.Write(msg)
		return
	}
	serverId, err := strconv.Atoi(param["dsid"])
	if err != nil {
		ret.State = ptsdk.STATUS_FAIL
		ret.Msg = ptsdk.PARAM_ERR
		msg, _ := json.Marshal(ret)
		w.Write([]byte(msg))
		logger.Error("充值异常，参数服务器Id错误：%v", param["dsid"])
		return
	}

	serverInfo := m.GetGsServers().GetServerInfo(serverId)
	if serverInfo == nil {
		ret.State = ptsdk.STATUS_FAIL
		ret.Msg = "server not found"
		msg, _ := json.Marshal(ret)
		w.Write([]byte(msg))
		logger.Error("充值异常，参数服务器未找到：%v", serverId)
		return
	}

	money, err := strconv.ParseFloat(param["money"], 64)
	if err != nil {
		if err != nil {
			ret.State = ptsdk.STATUS_FAIL
			ret.Msg = ptsdk.PARAM_ERR
			msg, _ := json.Marshal(ret)
			w.Write([]byte(msg))
			logger.Error("充值异常，参数充值金额错误：%v", param["money"])
			return
		}
	}
	moneyInt := int(money)

	coin, err := strconv.Atoi(param["coin"])
	if err != nil {
		if err != nil {
			ret.State = ptsdk.STATUS_FAIL
			ret.Msg = ptsdk.PARAM_ERR
			msg, _ := json.Marshal(ret)
			w.Write([]byte(msg))
			logger.Error("充值异常，参数获得错误：%v", param["coin"])
			return
		}
	}

	applyMsg := &pbserver.RechageCcsToGsReq{
		Oid:       param["oid"],
		GameOrder: param["doid"],
		Money:     int32(moneyInt),
		Coin:      int32(coin),
	}

	_, errmsg := callGameRechageResult(serverId, applyMsg)
	if len(errmsg) > 0 {
		ret.State = ptsdk.STATUS_FAIL
		ret.Msg = errmsg
		msg, _ := json.Marshal(ret)
		w.Write([]byte(msg))
		return
	}
	msg, _ := json.Marshal(ret)
	w.Write([]byte(msg))
}
func notifierRecharge(w http.ResponseWriter, r *http.Request) {

	sdk := ptsdk.GetSdk()
	serverId, orderdata := sdk.NotifierRecharge(w, r)

	if serverId == 0 || orderdata == nil {
		return
	}
	ret := &ptsdk.XYHttpResult{Status: 0, Message: ""}
	trialServer := rmodelCross.GetSystemSeting().GetSystemSettingConverInt(rmodelCross.SYSTEM_SETTING_TRIAL_SERVER)
	if trialServer == 0 {
		param := make(map[string]interface{})
		for k, v := range r.Form {
			if len(v) > 0 {
				param[k] = v[0]
			}
		}
		//判断是否提审服订单
		isTrialServer, err := strconv.ParseBool(param["extra1"].(string))
		if err == nil && isTrialServer {
			portInfo, err := modelCross.GetServerPortInfoModel().GetServerPortInfo(modelCross.TRIAL_CROSS_CENTER_HTTP)
			if err != nil {
				ret.Status = -1
				ret.Message = "trail server port error"
				msg, _ := json.Marshal(ret)
				w.Write(msg)
				logger.Error("获取提审服端口数据异常：%v", err)
				return
			}

			rb, err := httpclient.DoGet(fmt.Sprintf("http://%v:%v/api/rechage", portInfo.Host, portInfo.Port), common.ToUrlValues(param))

			if err != nil {
				logger.Error("推送提审服充值数据异常：%v", err)
				ret.Status = -1
				ret.Message = "trail server error"
				msg, _ := json.Marshal(ret)
				w.Write(msg)
				return
			}
			logger.Info("推送提审服结果数据：%v", string(rb))
			w.Write(rb)

			return
		}
	}

	errCode, errmsg := callGameRechageResult(serverId, orderdata)
	if sdk.GetPlatform() == constPlatfrom.PLATFORM_XY || sdk.GetPlatform() == constPlatfrom.PLATFORM_XY_WX {
		if len(errmsg) > 0 {
			ret.Status = -2
			ret.Message = errmsg
			if errCode == gamedb.ERRRECHARGEHASFINISH.Code {
				ret.Status = -8
			}
		}
	}

	msg, err := json.Marshal(ret)
	if err != nil {
		logger.Error("json marshal 异常：%v, err:%v", ret, err)
		return
	}

	w.Write([]byte(msg))
}

type XyWxRet struct {
	Status  int         `json:"status"` //是	请求状态码，0：成功；其他失败；
	Message string      `json:"msg"`    //是	请求说明
	Data    XyWxRetData `json:"data"`
}

type XyWxRetData struct {
	App_order_id string `json:"app_order_id"`
	Callback_url string `json:"callback_url"`
	Extra1       string `json:"extra1"`
}

func httpApplyPay(w http.ResponseWriter, r *http.Request) {

	sdk := ptsdk.GetSdk()
	serverId, applydata := sdk.ApplyPay(w, r)

	if serverId == 0 || applydata == nil {
		return
	}

	ret := XyWxRet{Status: 0, Message: "", Data: XyWxRetData{"", "", ""}}

	replayMsg := &pbserver.RechargeApplyAck{}
	err := m.GetGsServers().CallMessage(serverId, applydata, replayMsg)
	if err != nil {
		logger.Error("发送game申请充值异常：%v", err)
		errMsg := errex.GetErrorMessage(err)
		ret.Status = 1
		ret.Message = errMsg
	} else if len(replayMsg.OrderId) <= 0 {
		logger.Error("发送game申请充值，订单生成异常，订单为空")
		ret.Status = 1
		ret.Message = gamedb.ERRUNKNOW.Message
	} else {
		ret.Data.App_order_id = replayMsg.OrderId
	}

	msg, err := json.Marshal(ret)
	if err != nil {
		logger.Error("json marshal 异常：%v, err:%v", ret, err)
		return
	}
	w.Write(msg)
}
