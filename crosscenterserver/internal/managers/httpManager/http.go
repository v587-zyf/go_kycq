package httpManager

import (
	"cqserver/crosscenterserver/internal/managersI"
	"cqserver/gamelibs/modelCross"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw/httpserver"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var m managersI.IModule

func init() {
	// cmd register begin
	httpserver.Router().Handle("/api/rechage", http.HandlerFunc(httpRechage))
	httpserver.Router().Handle("/api/applyPay", http.HandlerFunc(httpApplyPay))
	httpserver.Router().Handle("/api/block", http.HandlerFunc(httpBlock))
	httpserver.Router().Handle("/api/removeBlock", http.HandlerFunc(httpBlockRemove))
	httpserver.Router().Handle("/api/sendmail", http.HandlerFunc(httpMailSend))
	httpserver.Router().Handle("/api/rollbackMail", http.HandlerFunc(httpMailRollback))
	httpserver.Router().Handle("/api/setFunctionState", http.HandlerFunc(httpfuncUpdate))
	httpserver.Router().Handle("/api/functionSwitchList", http.HandlerFunc(httpFuncShow))
	httpserver.Router().Handle("/api/roles", http.HandlerFunc(httpUserInfo))
	httpserver.Router().Handle("/api/deleteBulletin", http.HandlerFunc(httpDelAnnouncement))
	httpserver.Router().Handle("/api/saveOrUpdateBulletin", http.HandlerFunc(httpApplyAnnouncement))
	httpserver.Router().Handle("/api/whiteBlock", http.HandlerFunc(httpWhiteBlock))
	httpserver.Router().Handle("/api/removeWhiteBlock", http.HandlerFunc(httpWhiteBlockRemove))
	httpserver.Router().Handle("/api/servers", http.HandlerFunc(httpServers))
	// cmd register end
}

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func HttpServerInit(module managersI.IModule) error {
	m = module
	portInfo, err := modelCross.GetServerPortInfoModel().GetServerPortInfo(modelCross.CROSS_CENTER_HTTP)
	if err != nil || portInfo == nil {
		return errors.New(fmt.Sprintf("获取服务启动端口错误:%s", modelCross.CROSS_CENTER_HTTP))
	}
	err = httpserver.Start(fmt.Sprintf(":%d", portInfo.Port))
	if err != nil {
		logger.Error("http启动错误：%v", err)
		return err
	}
	return nil
}

func writeHttpMsg(w http.ResponseWriter, code int, msg string) {
	reMsg := Message{
		Code:    code,
		Message: msg,
	}
	reData, err := json.Marshal(reMsg)
	if err != nil {
		logger.Error("回写http请求，数据编译异常：code:%v,msg:%v，err:%v", code, msg, err)
		w.Write([]byte(`{"code":-1,"msg":"unknow error"}`))
		return
	}
	logger.Info("回写http请求数据，%v", reMsg, string(reData))
	_, err1 := w.Write(reData)
	if err1 != nil {
		logger.Error("回写http请求,回写数据异常，code:%v，message:%v，err:%v", code, msg, err)
	}
}
