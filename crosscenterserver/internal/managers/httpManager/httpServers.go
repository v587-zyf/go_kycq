package httpManager

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/ptsdk"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func httpServers(w http.ResponseWriter, r *http.Request) {

	timeParam := r.Header.Get("time")
	signParam := r.Header.Get("sign")

	requestData := make(map[string]interface{})
	//ServerId        int       `json:"serverId"`        //是	区服id
	//State           int       `json:"state"`           //否	0="待配置" 1="已配置" 2="待清档" 3="已清档" 4="已开服" 5="已停服"
	//FirstOpenTime   time.Time `json:"firstOpenTime"`   //否	开服时间，示例：'2022-04-18 142708'
	//IsNew           int       `json:"isNew"`           //否	是否新服（0: 否 1: 是 ）
	//IsMaintain      int       `json:"isMaintain"`      //否	维护状态（0: 正常 1: 维护中 ）
	//ArtificialLoad  int       `json:"artificialLoad"`  //否	人工负载（0: 通畅 1: 拥挤 2: 爆满 ）, 负载状态可根据不同游戏定制
	//IsOpenWhitelist int       `json:"isOpenWhitelist"` //否	是否开启白名单（0: 否 1: 是 )

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		r.Body.Close()
		logger.Error("解析平台发送来删除公告数据错误：%v", r.Body, err)
		ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "数据解析异常", nil)
		return
	}
	logger.Info("接收到平台发送来的更新服务器信息：%v", requestData)
	serverId := ""
	if s, ok := requestData["serverId"]; ok {
		serverId = strconv.Itoa(int(s.(float64)))
	}

	if !ptsdk.GetSdk().CheckSignForKy(signParam, serverId+timeParam) {
		ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "加密验证错误", nil)
		return
	}

	if firstOpenTime, ok := requestData["firstOpenTime"]; ok {
		openTime, err := common.GetTime(firstOpenTime.(string))
		if err != nil {
			ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "开服时间参数错误", nil)
			return
		}
		if openTime.Before(time.Now()) {
			ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "开服时间不能设置比当前时间早", nil)
			return
		}
	}
	err = modelCross.GetServerInfoModel().UpdateServerInfosFromGM(requestData)
	if err != nil {
		ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "更新服务器异常，联系研发", nil)
		return
	}
	ptsdk.GetSdk().HttpWriteReturnInfo(w, 200, "success", nil)
}
