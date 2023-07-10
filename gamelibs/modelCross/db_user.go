package modelCross

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/logger"
	"fmt"
	"strconv"
	"time"

	"cqserver/golibs/dbmodel"
	"gopkg.in/gorp.v1"
)

type UserCrossInfo struct {
	UserId           int              `db:"userId"`
	OpenId           string           `db:"openId"`
	ChannelId        int              `db:"channelId"`
	ServerId         int              `db:"serverId"`
	ServerIndex      int              `db:"serverIndex"`
	NickName         string           `db:"nickname"`
	Avatar           string           `db:"avatar"`
	Gold             int              `db:"gold"`
	Ingot            int              `db:"ingot"`
	Vip              int              `db:"vip"`
	TaskId           int              `db:"taskId"`
	Combat           int              `db:"combat"`
	Recharge         int              `db:"recharge"`
	RechargeToken    int              `db:"tokenRecharge"`
	Exp              int              `db:"exp"`
	LoginTime        time.Time        `db:"loginTime"`
	CreateTime       time.Time        `db:"createTime"`
	UpdateTime       time.Time        `db:"updateTime"`
	OffLineTime      time.Time        `db:"offLineTime"`
	LastRechargeTime time.Time        `db:"lastRechargeTime"`
	Heros            model.CrossHeros `db:"heros"`
}

type UserCrossInfoModel struct {
	dbmodel.CommonModel
}

var (
	userServerInfoModel  = &UserCrossInfoModel{}
	userServerInfoFields = dbmodel.GetAllFieldsAsString(UserCrossInfo{})
)

func init() {
	dbmodel.Register(model.DB_ACCOUNT, userServerInfoModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(UserCrossInfo{}, "user").SetKeys(false, "userId")
	})
}

func GetUserCrossInfoModel() *UserCrossInfoModel {
	return userServerInfoModel
}

func (this *UserCrossInfoModel) Create(userInfo *UserCrossInfo) error {
	return this.DbMap().Insert(userInfo)
}

func (this *UserCrossInfoModel) Update(userInfo *UserCrossInfo) (int, error) {
	count, err := this.DbMap().Update(userInfo)
	if err != nil {
		fmt.Println("err", err)
	}
	return int(count), err
}

func (this *UserCrossInfoModel) GetUserInfo(userId int) (*UserCrossInfo, error) {
	var info UserCrossInfo
	err := this.DbMap().SelectOne(&info, fmt.Sprintf("select %s from user where UserId = ?", userServerInfoFields), userId)
	if err != nil {
		log.Error("GetUserInfo failed, UserId=%v, err=%v", userId, err)
		return nil, err
	}
	return &info, nil
}

func (this *UserCrossInfoModel) GetAllByOpenId(openId string, limit int) ([]UserCrossInfo, error) {
	var infos []UserCrossInfo
	sql := fmt.Sprintf("select %s from user where openId = ? order by updateTime  desc limit ?", userServerInfoFields)
	//fmt.Println("sql:", sql)
	_, err := this.DbMap().Select(&infos, sql, openId, limit)
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func (this *UserCrossInfoModel) GetAllUsers() ([]*UserCrossInfo, error) {
	var infos []*UserCrossInfo
	sql := fmt.Sprintf("select %s from user where 1 = 1 ", userServerInfoFields)
	_, err := this.DbMap().Select(&infos, sql)
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func (this *UserCrossInfoModel) GetOneByOpenId(openId string, serverIdx int) (*UserCrossInfo, error) {
	var info UserCrossInfo
	err := this.DbMap().SelectOne(&info, fmt.Sprintf("select %s from user where openId = ? and serverId = ?", userServerInfoFields), openId, serverIdx)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (this *UserCrossInfoModel) GetServerIds(userIds []int) []int {

	ids := ""
	for _, v := range userIds {
		ids += strconv.Itoa(v) + ","
	}
	ids = ids[:len(ids)-1]

	var serverIds []int
	_, err := this.DbMap().Select(&serverIds, "select serverId from user where userId in (?)", ids)
	if err != nil {
		logger.Error("获取玩家服务器Id数据错误：%v,err:%v", ids, err)
	}
	return serverIds
}

func (this *UserCrossInfoModel) GetUserInfos(openId string, userId int, userName string, serverid int) []UserCrossInfo {

	var users []UserCrossInfo

	sql := "select * from user where 1"
	if userId > 0 {
		sql += fmt.Sprintf(" and userId = %d", userId)
	}
	if serverid > 0 {
		sql += fmt.Sprintf(" and serverId=%d", serverid)
	}

	if len(openId) > 0 {
		sql += fmt.Sprintf(" and openId=\"%s\"", openId)
	}

	if len(userName) > 0 {
		sql += fmt.Sprintf(" and nickname=\"%s\"", userName)
	}

	logger.Info("查询玩家数据sql:%v", sql)
	_, err := this.DbMap().Select(&users, sql)
	if err != nil {
		logger.Error("查询玩家数据异常：%v", err)
		return nil
	}
	return users

}
