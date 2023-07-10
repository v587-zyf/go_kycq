package ptsdk

import (
	"bytes"
	"cqserver/gamelibs/beans"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw/httpclient"
	"cqserver/protobuf/pb"
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

var SUBSCRIBE_TEMPLATE_ID_MAP = map[int]string{pb.SUBSCRIBE_HOOK: "eqDnVuif6ucKTN6SBLkYOcfm-IBKn3ad2EfXnDML_Vc"}

type hookData struct {
	Thing1 map[string]string `json:"thing1"`
}

type SdkXYWXLoginResult struct {
	Errno  int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		OpenId   string `json:"openid"`
		Nickname string `json:"nick"`
		Avatar   string `json:"avatar"`
		Gender   string `json:"gender"`
		Age      string `json:"age"`
	} `json:"data"`
}

type SdkXYWX struct {
	*SdkXY
}

func initXYWX(sdkConf *beans.Sdkconfig) *SdkXYWX {

	sdk := &SdkXYWX{}
	sdk.SdkXY = initXY(sdkConf)
	return sdk
}

func (this *SdkXYWX) GetOpenId(r *http.Request) (string, error) {

	openId := r.FormValue("openId")
	param := make(map[string]interface{})
	param["info"] = openId
	//param["gid"] = this.Gameid
	param["sign"] = this.GetXYSign(param, false)

	rb, err := httpclient.DoPost(this.Verifyurl, ToUrlValues(param))
	logger.Info("XY sdk 登录验证，获取玩家openId,请求数据：%v,返回数据：%v,err:%v", param, string(rb), err)
	if err != nil {
		return "", err
	}
	var result SdkXYWXLoginResult
	err = json.Unmarshal(rb, &result)
	if err != nil {
		return "", err
	}
	if result.Errno != 0 {
		if err, ok := XY_STATUS[result.Msg]; ok {
			return "", errors.New(err)
		}
		return "", gamedb.ERRUNKNOW
	}
	return result.Result.OpenId, nil
}

func (this *SdkXYWX) GetXYSign(param map[string]interface{}, isPay bool) string {

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
	buf.WriteString("#")
	if isPay {
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
func (this *SdkXYWX) NotifierRecharge(w http.ResponseWriter, r *http.Request) (int, *pbserver.RechageCcsToGsReq) {
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

/**
 *  @Description: 平台申请充值验证
 *  @param w
 *  @param r
 *  @return *pbserver.RechageCcsToGsReq //充值数据
 **/
func (this *SdkXYWX) ApplyPay(w http.ResponseWriter, r *http.Request) (int, *pbserver.RechargeApplyReq) {
	param := make(map[string]interface{})

	r.ParseForm()
	for k, v := range r.Form {
		if len(v) > 0 {
			param[k] = v[0]
		}
	}
	logger.Info("收到平台发送的充值申请：%v", param)
	sign := ""
	if _, ok := param["sign"]; ok {
		sign = param["sign"].(string)
		delete(param, "sign")
	} else {
		msg := this.getXYResultMsg(-2, "参数错误")
		this.HttpWriteReturnMsg(w, msg)
		return 0, nil
	}

	if _, ok := param["role_name"]; ok {
		delete(param, "role_name")
	}

	sign1 := this.GetXYSign(param, true)
	if sign != sign1 {
		logger.Error("签名错误,平台签名：%v,游戏加密签名：%v", sign, sign1)
		msg := this.getXYResultMsg(-2, "签名错误")
		this.HttpWriteReturnMsg(w, msg)
		return 0, nil
	}
	var err error
	product_type := 0
	product_id := 0
	if p, ok := param["product_id"].(string); !ok {
		logger.Error("参数错误,充值类型异常 :%v", param["product_id"])
		msg := this.getXYResultMsg(-2, "参数错误,平台充值类型错误")
		this.HttpWriteReturnMsg(w, msg)
		return 0, nil
	} else {
		payModuleData, err := strconv.Atoi(p)
		if err != nil {
			logger.Error("参数错误,充值类型异常 :%v", param["product_id"])
			msg := this.getXYResultMsg(-2, "参数错误,平台充值类型错误")
			this.HttpWriteReturnMsg(w, msg)
			return 0, nil
		}
		product_type = payModuleData / 1000000
		product_id = payModuleData % 1000000
	}

	//if p, ok := param["product_name"].(string); !ok {
	//	logger.Error("参数错误,充值类型ID异常:%v", param["product_name"])
	//	msg := this.getXYResultMsg(-2, "参数错误,平台充值类型Id错误")
	//	this.HttpWriteReturnMsg(w, msg)
	//	return 0, nil
	//} else {
	//	product_id, err = strconv.Atoi(p)
	//	if err != nil {
	//		logger.Error("参数错误,充值类型Id异常:%v", param["product_name"])
	//		msg := this.getXYResultMsg(-2, "参数错误,平台充值类型Id错误")
	//		this.HttpWriteReturnMsg(w, msg)
	//		return 0, nil
	//	}
	//}

	userId := 0
	if p, ok := param["role_id"].(string); !ok {
		logger.Error("参数错误,充值玩家ID异常:%v", param["role_id"])
		msg := this.getXYResultMsg(-2, "参数错误,平台充值角色Id错误")
		this.HttpWriteReturnMsg(w, msg)
		return 0, nil
	} else {
		userId, err = strconv.Atoi(p)
		if err != nil {
			logger.Error("参数错误,充值角色Id异常:%v", param["role_id"])
			msg := this.getXYResultMsg(-2, "参数错误,平台充值角色Id错误")
			this.HttpWriteReturnMsg(w, msg)
			return 0, nil
		}
	}

	money, err := strconv.Atoi(param["money"].(string))
	if err != nil {
		logger.Error("参数错误,充值金额错误")
		msg := this.getXYResultMsg(-2, "参数错误,充值金额错误")
		this.HttpWriteReturnMsg(w, msg)
		return 0, nil
	}

	sidStr := param["sid"].(string)
	sid, err := strconv.Atoi(sidStr)
	if err != nil {
		logger.Error("参数错误,缺少服务器Id")
		msg := this.getXYResultMsg(-2, "参数错误,缺少服务器Id")
		this.HttpWriteReturnMsg(w, msg)
		return 0, nil
	}

	order := &pbserver.RechargeApplyReq{
		UserId:    int32(userId),
		PayType:   int32(product_type),
		PayTypeId: int32(product_id),
		PayNum:    int32(money / 100),
	}
	return sid, order
}

func (this *SdkXYWX) GetRechargeData(serverId int, userName string, Lv int, order *modelGame.OrderDb, trialServer bool) string {
	//paydata := struct {
	//	Openid       string `json:"openid"`       //是 平台的用户ID，登录校验时得到的
	//	Money        string `json:"money"`        //是 商品总金额，单位： 分
	//	Product_name string `json:"product_name"` //是 商品名称
	//	Product_id   string `json:"product_id"`   //是 游戏方定义的商品ID
	//	Sid          string `json:"sid"`          //是 服务器id
	//	Sname        string `json:"sname"`        //是 服务器名称
	//	Role_id      string `json:"role_id"`      //是 角色id
	//	Role_name    string `json:"role_name"`    //是 角色名称，不参与签名计算
	//	Role_level   string `json:"role_level"`   //是 角色等级
	//	Extra1       string `json:"extra1"`       //否 自定义透传参数，回调发货时传回发货回调接口，长度200
	//	App_order_id string `json:"app_order_id"` //否 游戏订单号。不传表示CP未创建订单，回调发货时以回调信息创建订单。
	//	Time         string `json:"time"`         //否 时间戳，10位数字
	//	Sign         string `json:"sign"`         //是 本签名使用 paykey 生成
	//	 }
	paydata := make(map[string]interface{})
	paydata["openid"] = order.OpenId
	paydata["money"] = strconv.Itoa(order.PayMoney * 100)
	paydata["product_name"] = order.OrderDis
	paydata["product_id"] = strconv.Itoa(order.PayModule*1000000 + order.PayModuleId)
	paydata["sid"] = strconv.Itoa(serverId)
	paydata["sname"] = fmt.Sprintf("%d服", serverId)
	paydata["role_id"] = strconv.Itoa(order.UserId)
	paydata["role_level"] = strconv.Itoa(Lv)
	paydata["extra1"] = strconv.FormatBool(trialServer)
	paydata["app_order_id"] = order.OrderNo
	paydata["time"] = strconv.Itoa(int(time.Now().Unix()))

	sign := this.GetXYSign(paydata, true)

	paydata["role_name"] = userName
	paydata["sign"] = sign

	paydataStr, _ := json.Marshal(paydata)
	return string(paydataStr)
}

func (this *SdkXYWX) Subscribe(openId string, template int, arg ...string) {

	if _, ok := SUBSCRIBE_TEMPLATE_ID_MAP[template]; !ok {
		logger.Error("订阅未实现：%v", template)
		return
	}

	now := time.Now().Unix()
	data := this.getSubscribeData(template, arg...)
	if len(data) == 0 {
		return
	}
	params := make(map[string]interface{})
	params["appid"] = strconv.Itoa(this.Gameid)
	params["plat"] = "weixinx"
	params["template_id"] = SUBSCRIBE_TEMPLATE_ID_MAP[template]
	params["touser"] = openId
	params["data"] = data
	params["miniprogram_state"] = "formal"
	params["lang"] = "zh_CN"
	params["time"] = strconv.Itoa(int(now))

	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(int(now)))
	buf.WriteString("#")
	buf.WriteString(this.Gamekey)

	h := md5.New()
	io.WriteString(h, buf.String())
	sign := strings.ToLower(hex.EncodeToString(h.Sum(nil)))
	params["sign"] = sign
	rb, err := httpclient.DoPost("https://fa.xy.com/opensdk/subscribe", ToUrlValues(params))
	logger.Info("XY sdk 订阅通知，玩家openId,请求数据：%v,返回数据：%v,err:%v", params, string(rb), err)
	if err != nil {
		return
	}
}

func (this *SdkXYWX) getSubscribeData(template int, arg ...string) string {

	data := make(map[string]map[string]string)

	if template == pb.SUBSCRIBE_HOOK {
		data["thing2"] = make(map[string]string)
		data["thing2"]["value"] = arg[0]
		data["thing5"] = make(map[string]string)
		data["thing5"]["value"] = arg[1]
	}

	dataJ, err := json.Marshal(data)
	if err != nil {
		logger.Error("json marshal 数据错误：%v", err)
		return ""
	}
	return string(dataJ)
}
