package ptsdk

import (
	"bytes"
	"cqserver/gamelibs/beans"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw/httpclient"
	"cqserver/protobuf/pbserver"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

const XY_STATUS_SUCCESS = "0"

var XY_STATUS = map[string]string{
	"0":   "成功",
	"101": "参数错误",
	"102": "签名错误",
	"103": "token过期",
}

type SdkXYLoginResult struct {
	Errno  string `json:"errno"`
	Msg    string `json:"msg"`
	Result struct {
		Uid   string `json:"uid"`
		UName string `json:"uname"`
	} `json:"data"`
}

type XYHttpResult struct {
	Status  int    `json:"status"` //是	请求状态码，0：成功；其他失败；
	Message string `json:"msg"`    //是	请求说明
}

type SdkXY struct {
	*BaseSDK
}

func initXY(sdkConf *beans.Sdkconfig) *SdkXY {

	sdk := &SdkXY{}
	sdk.BaseSDK = &BaseSDK{Sdkconfig: sdkConf}
	return sdk
}

func (this *SdkXY) GetOpenId(r *http.Request) (string, error) {

	openId := r.FormValue("openId")
	token := r.FormValue("token")
	if len(token) == 0 {
		return "", gamedb.ERRPARAM
	}
	param := make(map[string]interface{})
	param["uid"] = openId
	param["gid"] = this.Gameid
	param["time"] = time.Now().Unix()
	param["token"] = token
	param["sign"] = this.GetXYSign(param, false)

	rb, err := httpclient.DoGet(this.Verifyurl, ToUrlValues(param))
	logger.Info("XY sdk 登录验证，获取玩家openId,请求数据：%v,返回数据：%v,err:%v", param, string(rb), err)
	if err != nil {
		return "", err
	}
	var result SdkXYLoginResult
	err = json.Unmarshal(rb, &result)
	if err != nil {
		return "", err
	}
	if result.Errno != XY_STATUS_SUCCESS {
		if err, ok := XY_STATUS[result.Msg]; ok {
			return "", errors.New(err)
		}
		return "", gamedb.ERRUNKNOW
	}
	return openId, nil
}

func (this *SdkXY) GetXYSign(param map[string]interface{}, isPay bool) string {

	keys := make([]string, len(param))
	index := 0
	for k, _ := range param {
		keys[index] = k
		index++
	}
	sort.Sort(sort.StringSlice(keys))

	var buf bytes.Buffer
	lastKey := len(keys) - 1
	for index, key := range keys {
		buf.WriteString(fmt.Sprintf("%s=%v", key, param[key]))
		if index < lastKey {
			buf.WriteString("&")
		}
	}
	if isPay {
		buf.WriteString("#")
		buf.WriteString(this.Paykey)
	} else {
		buf.WriteString(this.Gamekey)
	}
	h := md5.New()
	io.WriteString(h, buf.String())
	sign := strings.ToLower(hex.EncodeToString(h.Sum(nil)))
	return sign
}

/**
 *  @Description: 平台通知充值结果验证
 *  @param w
 *  @param r
 *  @return *pbserver.RechageCcsToGsReq //充值数据
 **/
func (this *SdkXY) NotifierRecharge(w http.ResponseWriter, r *http.Request) (int, *pbserver.RechageCcsToGsReq) {
	param := make(map[string]interface{})

	r.ParseForm()
	for k, v := range r.Form {
		if len(v) > 0 {
			param[k] = v[0]
		}
	}
	logger.Info("接收到平台发送来的充值数据：%v", r.Form)
	sign := ""
	if _, ok := param["sign"]; ok {
		sign = param["sign"].(string)
		delete(param, "sign")
	} else {
		msg := this.getXYResultMsg(-2, "参数错误")
		this.HttpWriteReturnMsg(w, msg)
		return 0, nil
	}
	sign1 := this.GetXYSign(param, true)
	if sign != sign1 {
		logger.Error("签名错误,平台签名：%v,游戏加密签名：%v", sign, sign1)
		msg := this.getXYResultMsg(-2, "签名错误")
		this.HttpWriteReturnMsg(w, msg)
		return 0, nil
	}

	if _, ok := param["order_id"].(string); !ok {
		logger.Error("参数错误,游戏订单错误")
		msg := this.getXYResultMsg(-2, "参数错误,游戏订单错误")
		this.HttpWriteReturnMsg(w, msg)
		return 0, nil
	}
	if _, ok := param["app_order_id"].(string); !ok {
		logger.Error("参数错误,平台订单错误")
		msg := this.getXYResultMsg(-2, "参数错误,平台订单错误")
		this.HttpWriteReturnMsg(w, msg)
		return 0, nil
	}
	money, err := strconv.Atoi(param["money"].(string))
	if err != nil {
		logger.Error("参数错误,充值金额错误")
		msg := this.getXYResultMsg(-2, "参数错误,充值金额错误")
		this.HttpWriteReturnMsg(w, msg)
		return 0, nil
	}
	coins, err1 := strconv.Atoi(param["coins"].(string))
	if err1 != nil {
		logger.Error("参数错误,游戏元宝金额错误")
		msg := this.getXYResultMsg(-2, "参数错误,游戏元宝金额错误")
		this.HttpWriteReturnMsg(w, msg)
		return 0, nil
	}

	order := &pbserver.RechageCcsToGsReq{
		Oid:       param["order_id"].(string),
		GameOrder: param["app_order_id"].(string),
		Money:     int32(money),
		Coin:      int32(coins),
	}

	sidStr := param["sid"].(string)
	sid, err := strconv.Atoi(sidStr)
	if err != nil {
		logger.Error("参数错误,缺少服务器Id")
		msg := this.getXYResultMsg(-2, "参数错误,缺少服务器Id")
		this.HttpWriteReturnMsg(w, msg)
		return 0, nil
	}

	return sid, order
}

func (this *SdkXY) getXYResultMsg(status int, msg string) string {
	result := &XYHttpResult{
		Status:  status,
		Message: msg,
	}
	str, _ := json.Marshal(result)
	return string(str)
}

func (this *SdkXY) GetRechargeData(serverId int, userName string, Lv int, order *modelGame.OrderDb, trialServer bool) string {
	paydata := struct {
		PayRmb         string `json:"payRmb"`         //支付金额(整数人民币 元 单位)
		OrderId        string `json:"orderId"`        //游戏订单id(游戏方自定义参数,唯一且不重复)
		ProductId      string `json:"productId"`      //用户游戏内购买道具id(游戏方自定义参数)
		ProductName    string `json:"productName"`    //用户游戏内购买道具名称(游戏方自定义参数)
		ServerId       string `json:"serverId"`       //游戏大区id(游戏方自定义参数，只有一个区默认传1，不能为0，必须整数)
		ServerName     string `json:"serverName"`     //游戏大区名称(游戏方自定义参数)
		OpenId         string `json:"sdkuid"`         //XY平台用户id，由XYSDK传给游戏的用户id
		RoleId         string `json:"roleId"`         //用户游戏角色id(游戏方自定义参数)
		RoleName       string `json:"roleName"`       //用户游戏角色名(游戏方自定义参数)
		Level          string `json:"level"`          //用户游戏角色等级(游戏方自定义参数)
		AppCallbackUrl string `json:"appCallbackUrl"` //备用支付回调地址(一般情况下传空值，发货地址统一由游戏方提供给xy服务端配置--游戏方自定义参数)
		Extra1         string `json:"extra1"`         //支付透传参数字段1 (具体规则请查看发货文档)
		Extra2         string `json:"extra2"`         //支付透传参数字段2 (具体规则请查看发货文档)
	}{
		PayRmb:         strconv.Itoa(order.PayMoney),
		OrderId:        order.OrderNo,
		ProductId:      strconv.Itoa(order.PayModule),
		ProductName:    strconv.Itoa(order.PayModuleId),
		ServerId:       strconv.Itoa(serverId),
		ServerName:     fmt.Sprintf("%d服", serverId),
		OpenId:         order.OpenId,
		RoleId:         strconv.Itoa(order.UserId),
		RoleName:       userName,
		Level:          strconv.Itoa(Lv),
		AppCallbackUrl: "",
		Extra1:         "",
		Extra2:         "",
	}
	paydataStr, _ := json.Marshal(paydata)
	return string(paydataStr)
}
