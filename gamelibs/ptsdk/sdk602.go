package ptsdk

import (
	"bytes"
	"cqserver/gamelibs/beans"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
)

var SDK_602_CODE = map[int]string{
	0:    " 验证成功",
	1001: "    参数未提供完整",
	1002: "    gameId参数错误",
	1003: "    token参数错误",
	1004: "    gameId参数错误[2]",
	1005: "    sign参数错误",
	1006: "    token无效",
	1007: "    token过期",
	1008: "    验证失败",
}

type Skd602LoginResult struct {
	Err  int    `json:"err"`
	Msg  string `json:"msg"`
	Data struct {
		Uid string `json:"uid"`
	} `json:"data"`
}

type Sdk602 struct {
	*BaseSDK
}

func initSdk602(sdkConf *beans.Sdkconfig) *Sdk602 {

	sdk := &Sdk602{}
	sdk.BaseSDK = &BaseSDK{Sdkconfig: sdkConf}
	return sdk
}

func (this *Sdk602) GetOpenId(r *http.Request) (string, error) {

	r.ParseForm()
	auth := r.FormValue("auth")
	verify := r.FormValue("verify")
	//authEncode := base64.StdEncoding.EncodeToString([]byte("account=111123121"))
	//auth = url.QueryEscape(authEncode)
	//logger.Info("-------------------------------1", auth)

	enurlAuth, err := url.QueryUnescape(auth)
	if err != nil {
		logger.Error("sdk602 url decode错误,数据：%v，异常：%v", auth, err)
		return "", nil
	}
	authDec, err := base64.StdEncoding.DecodeString(enurlAuth)
	if err != nil {
		logger.Error("sdk602 base64 解析错误：%v,异常：%v", auth, err)
		return "", nil
	}


	var buf bytes.Buffer
	buf.WriteString(auth)
	buf.WriteString(this.Gamekey)

	h := md5.New()
	io.WriteString(h, buf.String())
	sign := hex.EncodeToString(h.Sum(nil))
	//
	//verify = sign
	//logger.Info("---------------------------------2", sign)
	logger.Error("登录验证：auth:%v，verify:%v,游戏加密：%v", auth, verify, sign)
	if !this.sandbox && sign != verify {
		return "", errors.New("-4")
	}



	authStr := string(authDec)
	authSlice := common.NewStringSlice(authStr, "&")
	if len(authSlice) <= 0 {
		return "", errors.New("-3")
	}
	param := make(map[string]string)
	for _, v := range authSlice {
		vv := common.NewStringSlice(v, "=")
		if len(vv) == 2 {
			param[vv[0]] = vv[1]
		}
	}

	if _, ok := param["account"]; ok {
		return param["account"], nil
	}
	return "", errors.New("-3")

}

func (this *Sdk602) sign(param map[string]interface{}) string {

	var keys = make([]string, len(param))
	var index int
	for k := range param {
		keys[index] = k
		index++
	}
	sort.Sort(sort.StringSlice(keys))

	var buf bytes.Buffer
	for _, key := range keys {
		buf.WriteString(fmt.Sprintf("%s=%v", key, param[key]))
	}
	buf.WriteString(this.Gamekey)
	h := md5.New()
	io.WriteString(h, buf.String())
	sign := hex.EncodeToString(h.Sum(nil))
	return sign
}
