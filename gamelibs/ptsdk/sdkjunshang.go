package ptsdk

import (
	"cqserver/gamelibs/beans"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw/httpclient"
	"cqserver/protobuf/pb"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var JUNSHANG_CHANNEL = map[int]int{
	pb.CHATTYPE_WORLD:   1,
	pb.CHATTYPE_TEAM:    3,
	pb.CHATTYPE_GUILD:   4,
	pb.CHATTYPE_PRIVATE: 8,
}

const (
	PARAM_ERR = "缺少参数"
	SIGN_ERR  = "签名错误"
)

type SdkJunshangLoginResult struct {
	State int    `json:"state"`
	Msg   string `json:"msg"`
	Data  struct {
		Uid   string `json:"uid"`
		UName string `json:"uname"`
	} `json:"data"`
}

type sdkJunshangChatReprot struct {
	GameId int                 `json:"game_id"`
	Data   []*junshangChatData `json:"data"` // 二维数组，主体内容 否
	Ts     int                 `json:"ts"`   // ts Integer 时间戳
}

type junshangChatData struct {
	ChatId   int                  `json:"chatId"`   //Integer 聊天ID 否
	ServerId int                  `json:"serverId"` //Integer 游戏服Id 否
	Channel  int                  `json:"channel"`  //Integer 频道id 查看下面文档 否
	Chat     string               `json:"chat"`     //聊天内容 否
	Time     string               `json:"time"`     //timestamp 聊天日期时间 否
	From     *junshangChatUseInfo `json:"from"`     //from数据 发送者基本数据
	To       *junshangChatUseInfo `json:"to"`       //数据 接收者数据 （有数据则为私聊） 是
}

type junshangChatUseInfo struct {
	NickName string `json:"nickName"` //接收者角色名 是
	Account  string `json:"account"`  //接收者账号 是
	Role     string `json:"role"`     //接收者角色ID 是
	Vip      int    `json:"vip"`      //发送者vip等级 否
	Level    int    `json:"level"`    //Integer 发送者等级 否
}

type SdkJunshang struct {
	*BaseSDK
}

func initSdkJunshang(sdkConf *beans.Sdkconfig) *SdkJunshang {

	sdk := &SdkJunshang{}
	sdk.BaseSDK = &BaseSDK{Sdkconfig: sdkConf}
	return sdk
}

func (this *SdkJunshang) GetOpenId(r *http.Request) (string, error) {

	token := r.FormValue("openId")
	if len(token) == 0 {
		return "", gamedb.ERRPARAM
	}
	param := make(map[string]interface{})
	param["pid"] = this.PlatformId
	param["gid"] = this.Gameid
	param["time"] = time.Now().Unix()
	param["token"] = token
	singStr := fmt.Sprintf("%d%d%d%s", param["pid"], param["gid"], param["time"], this.Gamekey)
	h := md5.New()
	io.WriteString(h, singStr)
	sign := strings.ToLower(hex.EncodeToString(h.Sum(nil)))
	param["sign"] = sign

	rb, err := httpclient.DoGet(this.Verifyurl, ToUrlValues(param))
	logger.Info("君尚sdk 登录验证，获取玩家openId,请求数据：%v,返回数据：%v,err:%v", param, string(rb), err)
	if err != nil {
		return "", err
	}
	var result SdkJunshangLoginResult
	err = json.Unmarshal(rb, &result)
	if err != nil {
		return "", err
	}
	if result.State != STATUS_SUCCESS {
		return "", errors.New(result.Msg)
	}
	if len(result.Data.Uid) > 0 {
		return result.Data.Uid, nil
	}
	return "", gamedb.ERRUNKNOW

}

func (this *SdkJunshang) CheckSignForPlatform(checkSign string, arg ...interface{}) bool {

	var param map[string]string
	if len(arg) > 0 {
		if p, ok := arg[0].(map[string]string); ok {
			param = p
		}
	}
	if this.sandbox {
		return true
	}
	paramStr := fmt.Sprintf("%s%s%s%s%s%s%s%s", param["time"], this.Paykey, param["oid"], param["doid"], param["dsid"], param["uid"], param["money"], param["coin"])
	h := md5.New()
	io.WriteString(h, paramStr)
	sign := strings.ToLower(hex.EncodeToString(h.Sum(nil)))
	logger.Info("君尚sdk 加密验证，参数：%v,加密：%v,验证：%v", paramStr, sign, checkSign)
	return checkSign == sign
}

func (this *SdkJunshang) GetRechargeData(serverId int, userName string, Lv int, order *modelGame.OrderDb, trialServer bool) string {
	paydata := struct {
		Doid   string `json:"doid"`
		Dsid   int    `json:"dsid"`
		Dext   string `json:"dext"`
		Dmoney int    `json:"dmoney"`
		Money  int    `json:"money"`
	}{order.OrderNo, serverId, "dext", order.PayMoney, order.PayMoney}
	paydataStr, _ := json.Marshal(paydata)
	return string(paydataStr)
}

///**
// *  @Description: 聊天上报
// *  @param serverId		服务器Id
// *  @param channelId		聊天频道
// *  @param chatId		聊天Id
// *  @param chatMsg		聊天内容
// *  @param sender		发送者
// *  @param to			接受者
// **/
//func (this *SdkJunshang) ChatReport(serverId int, channelId int, chatId int, chatMsg string, sender *modelGame.UserBasicInfo, to *modelGame.UserBasicInfo) {
//	now := int(time.Now().Unix())
//	chatData := &sdkJunshangChatReprot{
//		GameId: this.Gameid,
//		Ts:     now,
//		Data:   make([]*junshangChatData, 1),
//	}
//	chatData.Data[0] = &junshangChatData{
//		Channel:  JUNSHANG_CHANNEL[channelId],
//		ServerId: serverId,
//		ChatId:   chatId,
//		Chat:     chatMsg,
//		From: &junshangChatUseInfo{
//			Account:  sender.OpenId,
//			NickName: sender.NickName,
//			Role:     strconv.Itoa(sender.Id),
//			Vip:      sender.Vip,
//			Level:    sender.Level,
//		},
//		To: &junshangChatUseInfo{},
//	}
//
//	if to != nil {
//		chatData.Data[0].To.Account = to.OpenId
//		chatData.Data[0].To.NickName = to.NickName
//		chatData.Data[0].To.Role = strconv.Itoa(to.Id)
//		chatData.Data[0].To.Vip = to.Vip
//		chatData.Data[0].To.Level = to.Level
//	}
//	chatData.Data = append(chatData.Data)
//
//	chatJson, err := json.Marshal(chatData)
//	if err != nil {
//		logger.Error("聊天数据json错误：%v,", err)
//		return
//	}
//	h := md5.New()
//	io.WriteString(h, string(chatJson))
//	postDataMd5 := strings.ToLower(hex.EncodeToString(h.Sum(nil)))
//	io.WriteString(h, postDataMd5+this.Gamekey)
//	sign := strings.ToLower(hex.EncodeToString(h.Sum(nil)))
//
//	param := make(map[string]interface{})
//	param["data"] = string(chatJson)
//	param["sign"] = sign
//
//	rb, err := httpclient.DoPost(this.ChatUrl, ToUrlValues(param))
//	if err != nil {
//		logger.Error("聊天上报异常：%v", err)
//		return
//	}
//	logger.Info("聊天上报结果：%v", string(rb))
//
//}
