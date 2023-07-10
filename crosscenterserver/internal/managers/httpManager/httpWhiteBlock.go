package httpManager

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/ptsdk"
	"cqserver/golibs/logger"
	"net/http"
	"sync"
)

var whiteBlockMu sync.Mutex
func httpWhiteBlock(w http.ResponseWriter, r *http.Request) {

	defer func() {
		whiteBlockMu.Unlock()
	}()
	whiteBlockMu.Lock()

	whiteId, whiteType, whiteVal, err := ptsdk.GetSdk().SetWhiteBlock(w, r)
	if err != nil {
		return
	}

	whiteModelDb := &modelCross.WhiteListDb{
		GMId:    whiteId,
		Valtype: whiteType,
		Value:   whiteVal,
	}
	err1 := modelCross.GetWhiteListDbModel().Create(whiteModelDb)
	if err1 != nil {
		logger.Error("插入白名单数据异常：%", err1)
		ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "写入数据库异常", nil)
	} else {
		ptsdk.GetSdk().HttpWriteReturnInfo(w, 200, "success", nil)
	}
}

func httpWhiteBlockRemove(w http.ResponseWriter, r *http.Request) {

	defer func() {
		whiteBlockMu.Unlock()
	}()
	whiteBlockMu.Lock()

	whiteId, whiteVal, err := ptsdk.GetSdk().DelWhiteBlock(w, r)
	if err != nil {
		return
	}

	err1 := modelCross.GetWhiteListDbModel().Del(whiteId, whiteVal)
	if err1 != nil {
		logger.Error("删除白名单数据异常：%", err1)
		ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "数据库操作异常", nil)
	} else {
		ptsdk.GetSdk().HttpWriteReturnInfo(w, 200, "success", nil)
	}
}
