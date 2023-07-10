package httpManager

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/ptsdk"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pbserver"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

var funcMu sync.Mutex

type funcStateData struct {
	FunctionId int    `json:"functionId"`
	Name       string `json:"name"`
	State      int    `json:"state"`
	ServerIds  []int  `json:"serverIds"`
}

type funcStateRequestData struct {
	Uuid              int `json:"uuid"` //是	记录id，主要用于签名
	FunctionStateList []struct {
		FunctionId int `json:"functionId"` //是	功能开关ID
		State      int `json:"state"`      //是	功能开关状态 0 关 1 开
	} `json:"functionStateList"`
	ChannelIds []int `json:"channelIds"` //array	否	渠道id集合(中心服模式必填)，注意当前参数为历史版本保留字段，新游戏可直接用serverIds参数分发，如[1,2,3]
	ServerIds  []int `json:"serverIds"`  //
}

func httpFuncShow(w http.ResponseWriter, r *http.Request) {

	timeParam := r.Header.Get("time")
	signParam := r.Header.Get("sign")
	if !ptsdk.GetSdk().CheckSignForKy(signParam, timeParam) {
		ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "加密验证错误", nil)
		return
	}
	data := make([]*funcStateData, 0)
	all := modelCross.GetFuncCloseDbModel().Getall()
	if all != nil {
		for _, v := range all {
			funcConf := gamedb.GetFunctionFunctionCfg(v.FuncId)
			if funcConf != nil {
				data = append(data, &funcStateData{
					FunctionId: v.FuncId,
					Name:       funcConf.Desc,
					State:      0,
					ServerIds:  v.ServerIds,
				})
			}
		}
	}

	ptsdk.GetSdk().HttpWriteReturnInfo(w, 200, "success", data)
}

func httpfuncUpdate(w http.ResponseWriter, r *http.Request) {

	defer func() {
		funcMu.Unlock()
	}()
	funcMu.Lock()

	var data funcStateRequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		r.Body.Close()
		logger.Error("解析平台发送来封禁数据错误：%v", r.Body, err)
		ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "功能模块数据解析错误", nil)
		return
	}
	timeParam := r.Header.Get("time")
	signParam := r.Header.Get("sign")
	if !ptsdk.GetSdk().CheckSignForKy(signParam, strconv.Itoa(data.Uuid)+timeParam) {
		ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "加密验证错误", nil)
		return
	}

	all := modelCross.GetFuncCloseDbModel().Getall()
	if all != nil {

		for _, v := range data.FunctionStateList {
			hasUpdate := false
			if all != nil {
				for _, has := range all {
					if has.FuncId == v.FunctionId {
						hasUpdate = true
						if v.State == 1 {
							modelCross.GetFuncCloseDbModel().Del(has)
						} else {
							has.ServerIds = data.ServerIds
							modelCross.GetFuncCloseDbModel().Update(has)
						}
						break
					}
				}
			}
			if !hasUpdate {
				info := &modelCross.FuncCloseDb{FuncId: v.FunctionId, ServerIds: data.ServerIds}
				modelCross.GetFuncCloseDbModel().Create(info)
			}
		}
	}
	ntf := &pbserver.FuncStateUpdateReq{}
	if len(data.ServerIds) > 0 {
		for _, v := range data.ServerIds {
			m.GetGsServers().SendMessage(v, ntf)
		}
	} else {
		m.GetGsServers().SendAllServerMessage(ntf)
	}
	ptsdk.GetSdk().HttpWriteReturnInfo(w,200,"success",nil)
}
