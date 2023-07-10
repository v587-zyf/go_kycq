package httpManager

import (
	"cqserver/crosscenterserver/internal/conf"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/ptsdk"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/protobuf/pb"
	"net/http"
	"strconv"
	"time"
)

type UserInfo struct {
	Account                  string    `json:"account"`                  //	是	账号
	AccountLastLoginAt       time.Time `json:"accountLastLoginAt"`       //datetime	是	账号最后登录时间
	AccountLastLogoutAt      time.Time `json:"accountLastLogoutAt"`      //	是	账号最后下线时间
	AccountLastDepositAt     time.Time `json:"accountLastDepositAt"`     //	账号最后充值时间
	AccountDepositTotal      int       `json:"accountDepositTotal"`      //	账号累计充值
	AccountDepositTokenTotal int       `json:"accountDepositTokenTotal"` //	账号累计充值
	AccountState             int       `json:"accountState"`             //	账号状态 0:正常，1：禁言，2：禁登
	AccountCreatedAt         time.Time `json:"accountCreatedAt"`         //	账号创建时间
	PlatformId               int       `json:"platformId"`               //	平台id

	RoleId                string    `json:"roleId"`                //	角色ID
	RoleName              string    `json:"roleName"`              //	角色名称
	ServerId              int       `json:"serverId"`              //	区服id
	ServerGroupId         int       `json:"serverGroupId"`         //	区服组id
	ChannelId             int       `json:"channelId"`             //	渠道id
	Exp                   int       `json:"exp"`                   //	经验
	IsFirstDeposit        bool      `json:"isFirstDeposit"`        //	是否首冲
	Vip                   int       `json:"vip"`                   //	VIP等级
	RoleLastLoginAt       time.Time `json:"roleLastLoginAt"`       //	角色最后登录时间
	RoleLastLogoutAt      time.Time `json:"roleLastLogoutAt"`      //	角色最后下线时间
	RoleLastDepositAt     time.Time `json:"roleLastDepositAt"`     //	角色最后充值时间
	RoleDepositTotal      int       `json:"roleDepositTotal"`      //	角色累计充值
	RoleDepositTokenTotal int       `json:"roleDepositTokenTotal"` //	角色累计充值
	RoleState             int       `json:"roleState"`             //	角色状态 0:正常，1：禁言，2：禁登
	RoleCreatedAt         time.Time `json:"roleCreatedAt"`         //	角色创建时间
	IsDeleted             bool      `json:"isDeleted"`             //	//否删除
	Currencies            []struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Count int    `json:"count"`
	} `json:"currencies"` //	货币详情，见：Currency

	Heros []UserHeroInfo `json:"heros"`
}

type UserHeroInfo struct {
	HeroIndex int `json:"hero_index"` //几号武将
	Level     int `json:"level"`      //	等级
	Career    int `json:"career"`     //	职业
	Sex       int `json:"sex"`        //	性别
}

/**
*  @Description: 获取玩家信息
*  @param w
*  @param r
**/
func httpUserInfo(w http.ResponseWriter, r *http.Request) {

	openId, userId, username, serverId, err := ptsdk.GetSdk().GetUserInfo(r)
	if len(err) > 0 {
		ptsdk.GetSdk().HttpWriteReturnMsg(w, err)
		return
	}
	if len(openId) <= 0 && userId <= 0 && len(username) <= 0 {
		ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "条件参数不能空", nil)
		return
	}
	var accout *modelCross.Account
	if len(openId) > 0 {
		accout, _ = modelCross.GetAccountModel().GetByOpenId(openId)
	}

	retInfos := make([]*UserInfo, 0)
	users := modelCross.GetUserCrossInfoModel().GetUserInfos(openId, userId, username, serverId)
	if users != nil {
		if accout == nil {

			var tempAccount *modelCross.Account
			for _, v := range users {

				if tempAccount != nil && v.OpenId == tempAccount.OpenId {
					userInfo := createUserInfoReturnInfo(tempAccount, &v)
					retInfos = append(retInfos, userInfo)

				} else {
					tempAccount, _ = modelCross.GetAccountModel().GetByOpenId(v.OpenId)
					if tempAccount != nil {
						userInfo := createUserInfoReturnInfo(tempAccount, &v)
						retInfos = append(retInfos, userInfo)
					}
				}
			}

			accountRechargeTotal := make(map[string]int)
			accountRechargeTokenTotal := make(map[string]int)
			for _, v := range retInfos {
				accountRechargeTotal[v.Account] += v.RoleDepositTotal
				accountRechargeTokenTotal[v.Account] += v.RoleDepositTokenTotal
			}
			for _, v := range retInfos {
				v.AccountDepositTotal = accountRechargeTotal[v.Account]
				v.AccountDepositTokenTotal = accountRechargeTokenTotal[v.Account]
			}

		} else {
			accountRechargeTotal := 0
			accountRechargeTokenTotal := 0
			for _, v := range users {
				userInfo := createUserInfoReturnInfo(accout, &v)
				retInfos = append(retInfos, userInfo)
				accountRechargeTotal += userInfo.RoleDepositTotal
				accountRechargeTokenTotal += userInfo.RoleDepositTokenTotal
			}
			for _, v := range retInfos {
				v.AccountDepositTotal = accountRechargeTotal
				v.AccountDepositTokenTotal = accountRechargeTokenTotal
			}
		}
	}

	ptsdk.GetSdk().HttpWriteReturnInfo(w, 200, "", retInfos)

}

func createUserInfoReturnInfo(account *modelCross.Account, user *modelCross.UserCrossInfo) *UserInfo {

	crossGroup, accountStateNum, userStateNum := getBaseInfo(account, user)
	userInfo := &UserInfo{
		Account:               account.OpenId,
		AccountLastLoginAt:    account.LastLoginTime,
		AccountLastLogoutAt:   user.OffLineTime,
		AccountLastDepositAt:  user.LastRechargeTime,
		AccountDepositTotal:   user.Recharge,
		AccountState:          accountStateNum, //账号状态 0:正常，1：禁言，2：禁登
		AccountCreatedAt:      account.CreateTime,
		PlatformId:            conf.Conf.Sdkconfig.KyPlatformId,
		RoleId:                strconv.Itoa(user.UserId),        //	角色ID
		RoleName:              user.NickName,                    //	角色名称
		ServerId:              user.ServerId,                    //	区服id
		ServerGroupId:         crossGroup,                       //	区服组id
		ChannelId:             user.ChannelId,                   //	渠道id
		Exp:                   user.Exp,                         //	经验
		IsFirstDeposit:        user.LastRechargeTime.Unix() > 0, //	是否首冲
		Vip:                   user.Vip,                         //	VIP等级
		RoleLastLoginAt:       user.LoginTime,                   //	角色最后登录时间
		RoleLastLogoutAt:      user.OffLineTime,                 //	角色最后下线时间
		RoleLastDepositAt:     user.LastRechargeTime,            //	角色最后充值时间
		RoleDepositTotal:      user.Recharge,                    //	角色累计充值
		RoleDepositTokenTotal: user.RechargeToken,               //	角色累计充值 代币
		RoleState:             userStateNum,                     //	角色状态 0:正常，1：禁言，2：禁登
		RoleCreatedAt:         user.CreateTime,                  //	角色创建时间
		IsDeleted:             false,                            //	//否删除
		Currencies: make([]struct {
			Id    int    `json:"id"`
			Name  string `json:"name"`
			Count int    `json:"count"`
		}, 2), //	货币详情，见：Currency

		Heros: make([]UserHeroInfo, len(user.Heros)),
	}

	userInfo.Currencies[0] = struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Count int    `json:"count"`
	}{Id: pb.ITEMID_GOLD, Name: "金币", Count: user.Gold}
	userInfo.Currencies[1] = struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Count int    `json:"count"`
	}{Id: pb.ITEMID_INGOT, Name: "元宝", Count: user.Ingot}

	for k, v := range user.Heros {
		userInfo.Heros[k] = UserHeroInfo{
			HeroIndex: v.HeroIndex,
			Level:     v.Level,
			Career:    v.Job,
			Sex:       v.Sex,
		}
	}
	return userInfo
}

func getBaseInfo(account *modelCross.Account, user *modelCross.UserCrossInfo) (int, int, int) {

	crossGroup := 0
	serverInfo, _ := modelCross.GetServerInfoModel().GetServerInfoByServerId(user.ServerId)
	if serverInfo != nil {
		crossGroup = serverInfo.CrossFsId
	}
	accountStateNum := 0
	userStateNum := 0
	accountState := account.BanData[constUser.ACCOUNT_BAN_KEY]
	if accountState != nil {
		accountStateNum = accountState.BanType
	}
	userState := account.BanData[user.UserId]
	if userState != nil {
		userStateNum = userState.BanType
	}
	return crossGroup, accountStateNum, userStateNum
}
