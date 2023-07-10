package user

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

func (this *UserManager) CheckBan(user *objs.User, banType int) (bool, int) {
	if user.AccountInfo == nil || user.AccountInfo.BanData == nil {
		return false, 0
	}
	if ban, ok := user.AccountInfo.BanData[constUser.ACCOUNT_BAN_KEY]; ok {
		if ban.BanType == banType {
			endTime, err := common.GetTime(ban.EndTime)
			if err == nil {
				if time.Now().Before(endTime) {
					logger.Info("玩家当前被封禁：%v", &ban)
					return true, int(endTime.Unix())
				}
			}
		}
	}
	if ban, ok := user.AccountInfo.BanData[user.Id]; ok {
		if ban.BanType == banType {
			endTime, err := common.GetTime(ban.EndTime)
			if err == nil {
				if time.Now().Before(endTime) {
					logger.Info("玩家当前被封禁：%v", &ban)
					return true, int(endTime.Unix())
				}
			}
		}
	}
	return false, 0
}

/**
 *  @Description: 更新玩家封禁信息
 *  @param openId
 *  @param userId
 **/
func (this *UserManager) UserBanUpdate(req *pbserver.BanInfoCcsToGsReq) {
	logger.Info("接收跨服中心发送来的更新玩家封禁信息：%v", *req)
	if req.UserId >= 0 {
		this.DispatchEvent(int(req.UserId), req, this.banUser)
	} else {
		logger.Error("接收跨服中心发送来 玩家Id异常：%v", *req)
	}
}

func (this *UserManager) banUser(userId int, user *objs.User, data interface{}) {
	if user == nil {
		return
	}

	req := data.(*pbserver.BanInfoCcsToGsReq)
	accountdb, err := modelCross.GetAccountModel().GetByOpenId(req.OpenId)
	if err != nil || accountdb == nil {
		logger.Error("接收到跨服中心发送来更新封禁玩家信息，玩家账号数据获取异常：%v", err, accountdb)
		return
	}

	user.AccountInfo = accountdb

	//解除禁言
	if req.IsRemove {

		ntf := &pb.ChatBanRemoveNtf{}
		this.GetUserManager().SendMessage(user, ntf, true)
		return
	}

	banType := -1
	banEndTime := 0
	reason := ""
	now := time.Now()
	if user.AccountInfo.BanData[constUser.ACCOUNT_BAN_KEY] != nil {
		endTime, _ := common.GetTime(user.AccountInfo.BanData[constUser.ACCOUNT_BAN_KEY].EndTime)
		if endTime.After(now) {
			banType = user.AccountInfo.BanData[constUser.ACCOUNT_BAN_KEY].BanType
			banEndTime = int(endTime.Unix())
			reason = user.AccountInfo.BanData[constUser.ACCOUNT_BAN_KEY].Reason
			if banType == constUser.BAN_TYPE_LOGIIN {
				//踢玩家下线
				go this.GetUserManager().KickUserWithMsg(user, reason)
				return
			}
		}
	}

	if user.AccountInfo.BanData[user.Id] != nil {
		banType = user.AccountInfo.BanData[user.Id].BanType
		endTime, _ := common.GetTime(user.AccountInfo.BanData[user.Id].EndTime)
		if endTime.After(now) {
			banType = user.AccountInfo.BanData[user.Id].BanType
			banEndTime = int(endTime.Unix())
			reason = user.AccountInfo.BanData[user.Id].Reason
		}
	}
	if banType == constUser.BAN_TYPE_LOGIIN {
		go this.KickUserWithMsg(user, reason)
		//踢玩家下线
	} else {
		//禁言
		ntf := &pb.ChatBanNtf{
			EndTime: int32(banEndTime),
			Reason:  reason,
		}
		this.SendMessage(user, ntf, false)
	}
}
