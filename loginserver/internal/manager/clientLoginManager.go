package manager

import (
	"cqserver/gamelibs/errex"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/ptsdk"
	"cqserver/gamelibs/publicCon/constPlatfrom"
	"cqserver/gamelibs/rmodel"
	"cqserver/golibs/common"
	"cqserver/golibs/nw/httpserver"
	"cqserver/loginserver/conf"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"cqserver/golibs/logger"
	"cqserver/golibs/util"
)

func init() {
	// cmd register begin
	httpserver.Router().Handle("/login602", http.HandlerFunc(Handle602LoginReq))
	httpserver.Router().Handle("/login", http.HandlerFunc(HandleLoginReq))
	httpserver.Router().Handle("/serverInfos", http.HandlerFunc(HandleGetServerListReq))
	httpserver.Router().Handle("/upAnnouncement", http.HandlerFunc(HandleUpAnnouncementReq))
	httpserver.Router().Handle("/checkVersion", http.HandlerFunc(HandleCheckVersionReq))
	// cmd register end
}

func HttpServerInit() error {
	loginPortKey := fmt.Sprintf(modelCross.CLIENT_TO_LOGIN, conf.Conf.LoginServer)
	portInfo, err := modelCross.GetServerPortInfoModel().GetServerPortInfo(loginPortKey)
	if err != nil || portInfo == nil {
		return errors.New(fmt.Sprintf("获取服务启动端口错误:%s", loginPortKey))
	}
	err = httpserver.Start(fmt.Sprintf(":%d", portInfo.Port))
	if err != nil {
		logger.Error("http启动错误：%v", err)
		return err
	}
	logger.Info("登录服端口号初始化完成：%v", portInfo.Port)
	return nil
}

func originSet(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else {
		w.Header().Add("Access-Control-Allow-Origin", "*")
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods", "*")
}

type LoginRet struct {
	OpenId       string              `json:"openId"`
	Cmd          string              `json:"cmd"`
	Servers      []*serverInfoClient `json:"servers"`
	Announcement []*AnnouncementInfo `json:"announcement"`
	LoginKey     string              `json:"loginkey"`
	Ip           string              `json:"ip"`
	Code         int                 `json:"code"`
	Message      string              `json:"message"`
	NewPlayer    bool                `json:"newplayer"`
}
type ServerListRet struct {
	Cmd     string              `json:"cmd"`
	Servers []*serverInfoClient `json:"servers"`
}
type serverInfoClient struct {
	Name           string    `json:"name"`           //服务名字
	ServerId       int       `json:"serverId"`       //服务器Id
	ServerIndex    int       `json:"serverIndex"`    //服务器索引（区服）
	GateHost       string    `json:"gateHost"`       //服务器Ip
	GatePort       string    `json:"gatePort"`       //服务器port
	IsNew          int       `json:"isNew"`          //是否新服
	Status         int       `json:"status"`         //状态:, 1:良好，2:正常，3:爆满',
	IsClose        int       `json:"isClose"`        //是否维护
	OpenTime       time.Time `json:"openTime"`       //开服时间
	CloseExplain   string    `json:"closeExplain"`   //维护说明
	Prefix         string    `json:"prefix"`         //玩家名字前缀
	Version        string    `json:"version"`        //版本号
	ClientVersion  string    `json:"clientVersion"`  //客户端资源版本号
	IsTrialVersion int       `json:"isTrialVersion"` //是否是体验服
}

type AnnouncementInfo struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Announcement string `json:"announcement"`
}

func getServerInfoClientByServerInfo(serverInfo *modelCross.ServerInfo, inWhiteList bool) *serverInfoClient {
	strSlice := common.NewStringSlice(serverInfo.Gates, ":")
	u := &serverInfoClient{
		Name:           serverInfo.Name,           //服务名字
		ServerId:       serverInfo.ServerId,       //服务器Id
		ServerIndex:    serverInfo.ServerIndex,    //服务器索引（区服）
		GateHost:       strSlice[0],               //服务器Ip
		GatePort:       strSlice[1],               //服务器port
		IsNew:          serverInfo.IsNew,          //是否新服
		Status:         serverInfo.Status,         //状态:, 1:良好，2:正常，3:爆满',
		IsClose:        serverInfo.IsClose,        //是否维护
		OpenTime:       serverInfo.OpenTime,       //开服时间
		CloseExplain:   serverInfo.CloseExplain,   //维护说明
		Prefix:         serverInfo.Prefix,         //玩家名字前缀
		Version:        serverInfo.Version,        //版本号
		ClientVersion:  serverInfo.ClientVersion,  //客户端资源版本号
		IsTrialVersion: serverInfo.IsTrialVersion, //是否是体验服

	}
	if serverInfo.IpFilter == 1 && !inWhiteList {
		u.IsClose = 1
	}
	return u
}

//判断ip是否在白名单内,ip在白名单内,则返回true,不在返回false
func JudgeIpFilter(ip string) bool {
	ip = strings.Split(ip, ":")[0]
	var ipFilters int
	err := modelCross.GetServerInfoModel().DbMap().SelectOne(&ipFilters, "select id from ipFilter where ip = ?", ip)
	if err != nil {
		return false
	}
	return true
}

func HandleLoginReq(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic HandleLoginReq:%v,%s", r, stackBytes)
		}
	}()

	originSet(w, r)

	ret := &LoginRet{
		Cmd:          "login",
		Servers:      make([]*serverInfoClient, 0),
		Announcement: make([]*AnnouncementInfo, 0),
	}

	r.ParseForm()
	logger.Info("收到客户端登录请求：%v", r.Form)
	ip := common.GetIpAddress(r)
	startAt := time.Now()
	openId, err := getOpenId(r)
	if err != nil || len(openId) == 0 {
		writeClientErrMsg(w, ret, errex.ErrOpenIdEmpty.Code, errex.ErrOpenIdEmpty.Message)
		return
	}
	inWhiteList := m.ServerList.InWhiteList(ip, openId)
	if r.FormValue("debug") == SUPER_LOGIN_KEY && !inWhiteList {
		writeClientErrMsg(w, ret, errex.ErrOpenIdEmpty.Code, errex.ErrOpenIdEmpty.Message)
		return
	}
	//获取玩家登录标识
	loginKey, err1 := getLoginKey(openId)
	if len(err1) > 0 {
		err := errex.BuildClientErrorAck(err)
		writeClientErrMsg(w, ret, int(err.Code), err.Message)
		return
	}

	err2 := getAccount(openId, ip)
	if err2 != nil {
		writeClientErrMsg(w, ret, int(errex.ErrOpenIdGetErr.Code), errex.ErrOpenIdGetErr.Message)
		return
	}

	rmodel.GetLoginKeyModel().CacheLoginKey(openId, loginKey)

	ret.OpenId = openId
	ret.LoginKey = loginKey
	ret.Ip = ip

	//取游戏角色信息2
	dbInfos, err := modelCross.GetUserCrossInfoModel().GetAllByOpenId(openId, 1)
	if err != nil {
		logger.Error("sql GetAllByOpenId error:%v", err)
	}
	logger.Info("role count=%v openIdd=%v,ip:%v", len(dbInfos), openId, ip)
	//公告
	ret.Announcement = m.Announcement.GetAnnouncementInfos()
	isClose := false
	if len(dbInfos) > 0 {
		for _, itr := range dbInfos {
			serverInfo := m.ServerList.GetServerInfoById(itr.ServerId, "")
			if serverInfo == nil {
				if itr.ServerId == m.ServerList.GetTrialServer().ServerId {
					serverInfo = m.ServerList.GetTrialServer()
				} else {
					continue
				}
			}
			if serverInfo.IsClose == 1 {
				isClose = true
			}
			serverMsg := getServerInfoClientByServerInfo(serverInfo, inWhiteList)
			ret.Servers = append(ret.Servers, serverMsg)
		}
	} else {
		newestServer := m.ServerList.GetNewestServer()
		if newestServer.IsClose == 1 {
			isClose = true
		}
		if newestServer != nil {
			ret.Servers = append(ret.Servers, getServerInfoClientByServerInfo(newestServer, inWhiteList))
			ret.NewPlayer = true
		}
	}

	//维护中，又不在白名单中，不给登录验证key，防止维护中外挂 直接socket直接进入
	if isClose && !inWhiteList {
		ret.LoginKey = ""
	}

	logger.Info("HandleLoginReq costTime:%vms", time.Now().Sub(startAt).Nanoseconds()/1e6)
	reMsg, err := json.Marshal(ret)
	if err != nil {
		logger.Error("返回信息json异常：%v", err)
		writeClientErrMsg(w, ret, int(errex.ErrUnknow.Code), errex.ErrUnknow.Message)
		return
	}

	writeUserEvent(openId, r.FormValue("deviceId"), r.FormValue("channel"), "enter_sid_list")
	w.Write(reMsg)
}

func writeClientErrMsg(w http.ResponseWriter, ret *LoginRet, code int, message string) {
	ret.Code = code
	ret.Message = message
	rb, _ := json.Marshal(ret)
	w.Write(rb)
}

func HandleGetServerListReq(w http.ResponseWriter, r *http.Request) {

	originSet(w, r)

	startAt := time.Now()
	r.ParseForm()

	openId := strings.TrimSpace(r.FormValue("openId"))
	success := rmodel.GetLoginKeyModel().ValidateLogin(openId)
	if !success {
		w.Write([]byte(errex.ErrOpenIdGetErr.Message))
		return
	}

	serverInfos := m.ServerList.GetServerList()
	if serverInfos == nil || len(serverInfos) == 0 {
		w.Write([]byte(errex.ErrServerListEmpty.Message))
		logger.Error("服务器信息列表为空")
		return
	}
	//isFilter := JudgeIpFilter(conn.RemoteAddr().String())
	ip := common.GetIpAddress(r)
	inWhiteList := m.ServerList.InWhiteList(ip, openId)
	ret := ServerListRet{
		Cmd:     "serverInfos",
		Servers: make([]*serverInfoClient, 0),
	}
	for _, v := range serverInfos {
		ret.Servers = append(ret.Servers, getServerInfoClientByServerInfo(v, inWhiteList))
	}

	if inWhiteList {
		trialServer := m.ServerList.GetTrialServer()
		if trialServer != nil {
			ret.Servers = append(ret.Servers, getServerInfoClientByServerInfo(trialServer, inWhiteList))
		}
	}

	logger.Info("HandleGetServerListReq costTime=%vms appId=%v version=%v serversCount=%v serverIndex=%v", time.Now().Sub(startAt).Nanoseconds()/1e6)
	reMsg, err := json.Marshal(ret)
	if err != nil {
		logger.Error("获取服务器列表信息json异常：%v", err)
		w.Write([]byte(errex.ErrUnknow.Message))
	}
	w.Write(reMsg)
}

func getOpenId(r *http.Request) (string, error) {

	if conf.Conf.Sandbox || r.FormValue("debug") == SUPER_LOGIN_KEY {
		openId := r.FormValue("openId")
		if len(openId) <= 0 {
			logger.Error("获取登录参数openId错误")
			return "", errex.ErrOpenIdGetErr
		}
		return openId, nil
	} else {
		openId, err := ptsdk.GetSdk().GetOpenId(r)
		if err != nil {
			return "", err
		}
		return openId, nil
	}
}

func getAccount(openId string, ip string) error {
	account, err := modelCross.GetAccountModel().GetByOpenId(openId)
	if err != nil && err != sql.ErrNoRows {

		logger.Error("获取玩家[%v]账号信息错误：%v", openId, err)
		return errex.ErrOpenIdGetErr
	}

	curTime := time.Now()
	if err == sql.ErrNoRows {
		account = &modelCross.Account{}
		account.OpenId = openId
		account.CreateTime = curTime
		account.LastLoginServerId = -1
		account.CreateIp = ip
	} else {
		if account.IsLocked() {
			logger.Error("账号：%v,锁定中", openId)
			return errex.ErrOpenIdGetErr
		}
	}

	account.LoginCount += 1
	account.LastLoginTime = time.Now()

	if err == sql.ErrNoRows {
		err = modelCross.GetAccountModel().DbMap().Insert(account)
	} else {
		_, err = modelCross.GetAccountModel().DbMap().Update(account)
	}
	if err != nil {
		logger.Error("获取玩家[%v]账号信息错误：%v", openId, err)
		return errex.ErrOpenIdGetErr
	}
	return nil
}

func getLoginKey(openId string) (string, string) {

	if ptsdk.GetSdk().GetPlatform() == constPlatfrom.PLATFORM_602 {
		loginKey := rmodel.GetLoginKeyModel().GetValidateLogin(openId)
		if len(loginKey) <= 0 {
			logger.Error("602平台验证玩家登录错误：%v", openId)
			return "", errex.ErrOpenIdGetErr.Message
		}
		return loginKey, ""

	} else {
		var err error
		loginKey, err := util.GenerateSessionId()
		if err != nil {
			logger.Error("生成loginkey openID:%v,err:%v", openId, err)
			return "", errex.ErrOpenIdGetErr.Message
		}
		return loginKey, ""
	}
}

func Handle602LoginReq(w http.ResponseWriter, r *http.Request) {

	if constPlatfrom.PLATFORM_602 != ptsdk.GetSdk().GetPlatform() {
		logger.Error("当前平台不是602平台")
		return
	}
	err := r.ParseForm()
	if err != nil {
		logger.Error("平台参数解析异常")
	}
	openId, err := ptsdk.GetSdk().GetOpenId(r)
	if err != nil {

		w.Write([]byte(fmt.Sprintf("%s", err.Error())))
		return
	}
	loginKey, err := util.GenerateSessionId()
	if err != nil {
		logger.Error("生成登录标识错误：%v", err)
		w.Write([]byte("-5"))
		return
	}
	rmodel.GetLoginKeyModel().CacheLoginKey(openId, loginKey)

	w.Header().Set("Cache-Control", "must-revalidate, no-store")
	w.Header().Set("Content-Type", " text/html;charset=UTF-8")
	w.Header().Set("Location", fmt.Sprintf(conf.Conf.ClientAddr+"?openId=%s", openId)) //跳转地址设置
	w.WriteHeader(http.StatusTemporaryRedirect)                                        //关键在这里！
}

func HandleUpAnnouncementReq(w http.ResponseWriter, r *http.Request) {

	originSet(w, r)
	m.Announcement.UpAnnouncementInfos()
	w.Write([]byte("SUCCESS"))
}

func HandleCheckVersionReq(w http.ResponseWriter, r *http.Request) {

	originSet(w, r)
	r.ParseForm()
	version := r.FormValue("version")

	ret := struct {
		Cmd        string `json:"cmd"`
		Addr       string `json:"addr"`
		ClientAddr string `json:"client_addr"`
		Version    string `json:"version"`
	}{
		Cmd:        "checkVersion",
		Addr:       conf.Conf.Wssaddr,
		ClientAddr: conf.Conf.ClientAddr,
		Version:    m.ServerList.TrailVersion(),
	}
	if m.ServerList.isTrailVersion(version) {
		ret.Addr = conf.Conf.WssaddrTrail
		ret.ClientAddr = conf.Conf.ClientAddrTrail
	}
	retData, err := json.Marshal(ret)
	if err != nil {
		logger.Error("json 异常：%v", err)
	}
	w.Write(retData)
}
