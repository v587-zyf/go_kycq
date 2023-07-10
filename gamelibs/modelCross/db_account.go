package modelCross

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/logger"
	"fmt"
	"time"

	"gopkg.in/gorp.v1"

	"cqserver/golibs/dbmodel"
)

type Account struct {
	Id                int              `db:"id"`
	OpenId            string           `db:"openId"`
	ChannelId         int              `db:"channelId"`
	LoginCount        int              `db:"loginCount"`
	Status            int              `db:"status"`
	Freeze            string           `db:"freeze"`
	LastLoginServerId int              `db:"lastLoginServerId"`
	LastLoginTime     time.Time        `db:"lastLoginTime"`
	CreateIp          string           `db:"createIp"`
	CreateTime        time.Time        `db:"createTime"`
	BanData           model.AccountBan `db:"banInfo"`
}

const (
	BEND_TYPE_CHAT         = 1 //禁止普通消息
	BEND_TYPE_SERVICE_CHAT = 2 //禁止发客服消息
)

type AccountModel struct {
	dbmodel.CommonModel
}

var (
	accountModel  = &AccountModel{}
	accountFields = dbmodel.GetAllFieldsAsString(Account{})
)

func init() {
	dbmodel.Register(model.DB_ACCOUNT, accountModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(Account{}, "account").SetKeys(true, "id")
	})
}

func GetAccountModel() *AccountModel {
	return accountModel
}

func (this *Account) IsLocked() bool {
	return false
}

func (this *AccountModel) GetByOpenId(openId string) (*Account, error) {
	var account Account
	err := this.DbMap().SelectOne(&account, fmt.Sprintf("select %s from account where openId = ?", accountFields), openId)
	if err != nil {
		logger.Error("获取玩家账号数据错误,账号：%v,err：%v", openId, err)
		return nil, err
	}
	return &account, nil
}

func (this *AccountModel) GetByAccountId(accountId int) (*Account, error) {
	var account Account
	err := this.DbMap().SelectOne(&account, fmt.Sprintf("select %s from account where id = ?", accountFields), accountId)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (this *AccountModel) SetStatus(accountId int, status int) error {
	account, err := this.GetByAccountId(accountId)
	if err != nil {
		return err
	}
	account.Status = status
	this.DbMap().Update(account)
	return nil
}

func (this *AccountModel) LockAccount(accountId int, status int, freeze string) error {
	account, err := this.GetByAccountId(accountId)
	if err != nil {
		return err
	}
	account.Status = status
	account.Freeze = freeze
	this.DbMap().Update(account)
	return nil
}

func (this *AccountModel) Update(account *Account) error {
	_, err := this.DbMap().Update(account)
	return err
}

func (this *AccountModel) SetLastLoginServerId(openId string, lastLoginServerId int) error {
	_, err := this.DbMap().Exec("update account set lastLoginServerId=? where openId=?", lastLoginServerId, openId)
	return err
}

func (this *AccountModel) GetAllAccountByOpenId(openId string, limit int, nowTime int64) ([]Account, error) {
	var infos []Account
	_, err := this.DbMap().Select(&infos, fmt.Sprintf("select %s from account where openId = ? and UNIX_TIMESTAMP(createTime)<? limit ?", accountFields), openId, nowTime-1, limit)
	if err != nil {
		return nil, err
	}
	return infos, nil
}
