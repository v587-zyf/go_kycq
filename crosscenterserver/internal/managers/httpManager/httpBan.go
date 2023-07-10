package httpManager

import (
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/ptsdk"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pbserver"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var blockMu sync.Mutex

func GetAccount(target int, targetValue string) (*modelCross.Account, *modelCross.UserCrossInfo) {

	openId := ""
	var user *modelCross.UserCrossInfo
	if target == 2 {
		var err error
		userId, _ := strconv.Atoi(targetValue)
		user, err = modelCross.GetUserCrossInfoModel().GetUserInfo(userId)
		if err != nil || user == nil {
			logger.Error("获取玩家数据错错误，%v,%v,err:%v,user:%v", target, targetValue, err, user)
			return nil, nil
		}
		openId = user.OpenId
	} else if target == 1 {
		openId = targetValue
	} else {
		logger.Error("暂不支持")
		return nil, nil
	}

	account, err1 := modelCross.GetAccountModel().GetByOpenId(openId)
	if err1 != nil || account == nil {
		logger.Error("玩家账号数据未找到，%v,%v,err:%v,account:%v", target, targetValue, err1, account)
		return nil, nil
	}
	return account, user
}

func httpBlock(w http.ResponseWriter, r *http.Request) {

	defer func() {
		blockMu.Unlock()
	}()
	blockMu.Lock()

	banData, err := ptsdk.GetSdk().Ban(r)
	if len(err) > 0 {
		ptsdk.GetSdk().HttpWriteReturnMsg(w, err)
		return
	}
	account, user := GetAccount(banData.Target, banData.TargetVal)
	if account == nil {
		ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "获取账号数据错误", nil)
		return
	}

	now := time.Now()
	if account.BanData != nil {

		for k, v := range account.BanData {
			endTime, err := common.GetTime(v.EndTime)
			if err != nil {
				continue
			}
			if endTime.Before(now) {
				delete(account.BanData, k)
			}
		}
	} else {
		account.BanData = make(model.AccountBan)
	}

	if banData.Target == 1 {

		if banData.BlockType == 1 {
			account.BanData = make(model.AccountBan)
		}
		endTime := time.Now()
		if banData.Duration > 0 {
			endTime = now.Add(time.Duration(banData.Duration) * time.Millisecond)
		} else {
			endTime = now.Add(time.Duration(86400000*365) * time.Millisecond)
		}

		account.BanData[constUser.ACCOUNT_BAN_KEY] = &model.BanInfo{
			StartTime: common.GetFormatTime2(now),
			EndTime:   common.GetFormatTime2(endTime),
			Reason:    banData.Reason,
			BanType:   banData.BlockType,
		}

	} else {
		endTime := time.Now()
		if banData.Duration > 0 {
			endTime = now.Add(time.Duration(banData.Duration) * time.Millisecond)
		} else {
			endTime = now.Add(time.Duration(86400000*365) * time.Millisecond)
		}

		account.BanData[user.UserId] = &model.BanInfo{
			StartTime: common.GetFormatTime2(now),
			EndTime:   common.GetFormatTime2(endTime),
			Reason:    banData.Reason,
			BanType:   banData.BlockType,
		}
	}

	err1 := modelCross.GetAccountModel().Update(account)
	if err1 != nil {
		logger.Error("更新账号数据异常：%v", err1)
	}
	ptsdk.GetSdk().HttpWriteReturnInfo(w, 200, "success", nil)

	//通知game
	banInfoSendGame(account, user, false)
}

func httpBlockRemove(w http.ResponseWriter, r *http.Request) {

	defer func() {
		blockMu.Unlock()
	}()
	blockMu.Lock()

	blockData, err := ptsdk.GetSdk().BanRemove(r)
	if len(err) > 0 {
		ptsdk.GetSdk().HttpWriteReturnMsg(w,err)
		return
	}

	account, user := GetAccount(blockData.Target, blockData.TargetVal)
	if account == nil {
		ptsdk.GetSdk().HttpWriteReturnInfo(w,400, "获取账号数据错误", nil)
		return
	}

	if account.BanData == nil {
		ptsdk.GetSdk().HttpWriteReturnInfo(w,200, "success", nil)
		return
	}

	if blockData.Target == 1 {
		account.BanData = make(model.AccountBan)
	} else {
		delete(account.BanData, user.UserId)
	}
	err1 := modelCross.GetAccountModel().Update(account)
	if err1 != nil {
		logger.Error("更新账号数据异常：%v", err1)
	}
	ptsdk.GetSdk().HttpWriteReturnInfo(w,200, "success", nil)
	//解封封禁登录的账号，不用通知玩家
	if blockData.BlockType == constUser.BAN_TYPE_LOGIIN {
		return
	}
	//通知game
	banInfoSendGame(account, user, true)
}

func banInfoSendGame(account *modelCross.Account, user *modelCross.UserCrossInfo, isRemove bool) {
	ntf := &pbserver.BanInfoCcsToGsReq{
		OpenId:   account.OpenId,
		IsRemove: isRemove,
	}
	if user != nil {
		ntf.UserId = int32(user.UserId)
		m.GetGsServers().SendMessage(user.ServerId, ntf)
		return
	}
	users, err := modelCross.GetUserCrossInfoModel().GetAllByOpenId(account.OpenId, 50)
	if err != nil {
		logger.Error("拉取玩家数据错误：%v,err:%v", account.OpenId, err)
		return
	}
	for _, v := range users {
		ntf.UserId = int32(v.UserId)
		m.GetGsServers().SendMessage(v.ServerId, ntf)
	}

}
