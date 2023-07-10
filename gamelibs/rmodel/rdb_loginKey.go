package rmodel

import (
	"fmt"
	"strconv"
	"time"

	"cqserver/golibs/logger"
)

const (
	LOGINKEY       = "loginkey_%s"       //登录的token string,(openId)
	LOGINTICKETKEY = "loginticketkey_%s" //登录的token string,(loginticket)

	LOGINTICKET_NICKNAME  = "1"
	LOGINTICKET_CHANNELID = "2"
	LOGINTICKET_OPENID    = "3"
	LOGINTICKET_SEX       = "4"
	LOGINTICKET_AVATAR    = "5"
)

type LoginKeyModel struct {
}

var loginKey = &LoginKeyModel{}

func GetLoginKeyModel() *LoginKeyModel {
	return loginKey
}

var log = logger.Get("default", true)

func (this *LoginKeyModel) getLoginKey(openId string) string {
	return fmt.Sprintf(LOGINKEY, openId)
}

func (this *LoginKeyModel) CacheLoginKey(openId, loginKey string) error {

	value := fmt.Sprintf("%s", loginKey)
	key := this.getLoginKey(openId)
	RedisDb().SetWithExpire(key, value, 10*time.Minute)
	return nil
}

func (this *LoginKeyModel) ValidateLoginKey(openId, loginKey string) bool {
	key := this.getLoginKey(openId)
	loginKeyCache, err := RedisDb().Get(key).String()

	logger.Info("Verify key:%v loginKeyCache=%v,loginKey:%v", key, loginKeyCache, loginKey)

	if err != nil {
		return false
	}
	return loginKey == loginKeyCache
}

func (this *LoginKeyModel) ValidateLoginUpdate(openId string){
	key := this.getLoginKey(openId)
	logger.Info("更新玩家：%v,验证key过期时间",openId)
	RedisDb().Expire(key, 10*time.Minute)
}

func (this *LoginKeyModel) ValidateLogin(openId string) bool {
	key := this.getLoginKey(openId)
	loginKeyCache, err := RedisDb().Get(key).String()

	logger.Info("Verify key :%v,loginKeyCache=%v", key, loginKeyCache)

	if err != nil {
		return false
	}
	return len(loginKeyCache) > 0
}

func (this *LoginKeyModel) GetValidateLogin(openId string) string {
	key := this.getLoginKey(openId)
	loginKeyCache, err := RedisDb().Get(key).String()

	logger.Info("Verify key :%v,loginKeyCache=%v,err:%v", key, loginKeyCache, err)

	if err != nil {
		return ""
	}
	return loginKeyCache
}

func (this *LoginKeyModel) getLoginTicketKey(ticket string) string {
	return fmt.Sprintf(LOGINTICKETKEY, ticket)
}

func (this *LoginKeyModel) CacheLocalLoginTicket(ticket, nickName, openId, avatar string, channelId, sex int) error {
	key := this.getLoginTicketKey(ticket)
	RedisDb().Hmset(key, LOGINTICKET_NICKNAME, nickName, LOGINTICKET_CHANNELID, channelId, LOGINTICKET_OPENID, openId, LOGINTICKET_SEX, sex, LOGINTICKET_AVATAR, avatar)
	RedisDb().Expire(key, 10*time.Minute)
	return nil
}

func (this *LoginKeyModel) ValidateLoginTicket(ticket string) (string, string, string, int, int, bool) {
	key := this.getLoginTicketKey(ticket)
	strs, err := RedisDb().HmgetStr(key, LOGINTICKET_NICKNAME, LOGINTICKET_OPENID, LOGINTICKET_AVATAR, LOGINTICKET_CHANNELID, LOGINTICKET_SEX)

	if err != nil || len(strs) < 3 {
		return "", "", "", 0, 0, false
	}
	channelId, err := strconv.Atoi(strs[3])
	if err != nil {
		return "", "", "", 0, 0, false
	}
	sex, err := strconv.Atoi(strs[4])
	if err != nil {
		return "", "", "", 0, 0, false
	}
	return strs[0], strs[1], strs[2], channelId, sex, true
}
