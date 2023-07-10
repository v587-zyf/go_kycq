package ptsdk

import (
	"cqserver/gamelibs/beans"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/modelGame"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw/httpclient"
	"cqserver/protobuf/pbserver"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

const (
	ERR_BASE_PARAM        = "数据解析错误"
	ERR_BASE_LOSE_PARAM   = "缺少参数"
	ERR_BASE_SIGN         = "密钥验证错误"
	ERR_BASE_OPERATE_TYPE = "不支持的方式"
	ERR_BASE_PARAM_ERR    = "参数异常"
)

type BaseSDK struct {
	*beans.Sdkconfig
	sandbox bool
}

func (this *BaseSDK) SetSandBox(sandbox bool) {
	this.sandbox = sandbox
}

func (this *BaseSDK) GetPlatform() int {
	return this.Sdkconfig.KyPlatformId
}
func (this *BaseSDK) GetOpenId(r *http.Request) (string, error) {
	return "", errors.New("interface not found")
}

func (this *BaseSDK) CheckSignForPlatform(sign string, arg ...interface{}) bool {
	return false
}

func (this *BaseSDK) CheckSignForKy(sign string, arg ...interface{}) bool {

	if !this.sandbox {
		str := ""
		for _, v := range arg {
			str += v.(string)
		}
		h := md5.New()
		io.WriteString(h, str+this.KyToken)
		sign1 := hex.EncodeToString(h.Sum(nil))
		if sign1 != sign {
			logger.Error("封禁加密错误,数据：%v,加密后为：%v,平台sign：%v", str+this.KyToken, sign1, sign)
			return false
		}
	}
	return true
}

func (this *BaseSDK) getSign(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	sign := hex.EncodeToString(h.Sum(nil))
	return sign
}
func (this *BaseSDK) GetRechargeData(serverId int, userName string, Lv int, order *modelGame.OrderDb, trialServer bool) string {
	return ""
}

/**
 *  @Description: 平台通知充值结果验证
 *  @param w
 *  @param r
 *  @return *pbserver.RechageCcsToGsReq //充值数据
 **/
func (this *BaseSDK) NotifierRecharge(w http.ResponseWriter, r *http.Request) (int, *pbserver.RechageCcsToGsReq) {
	return 0, nil
}

/**
*  @Description: 平台申请充值
*  @receiver this
*  @param w
*  @param r
*  @return int
*  @return *pbserver.RechargeApplyReq
**/
func (this *BaseSDK) ApplyPay(w http.ResponseWriter, r *http.Request) (int, *pbserver.RechargeApplyReq) {
	return 0, nil
}

/**
 *  @Description: 回写http消息
 *  @param w
 *  @param msg
 **/
func (this *BaseSDK) HttpWriteReturnInfo(w http.ResponseWriter, code int, msg string, data interface{}) {

	ret := this.GetSdkResultMsg(code, msg, data)

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(ret))
}

/**
 *  @Description: 回写http消息
 *  @param w
 *  @param msg
 **/
func (this *BaseSDK) HttpWriteReturnMsg(w http.ResponseWriter, msg string) {

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(msg))
}

/**
 *  @Description: 聊天上报
 *  @param serverId		服务器Id
 *  @param channelId		聊天频道
 *  @param chatId		聊天Id
 *  @param chatMsg		聊天内容
 *  @param sender		发送者
 *  @param to			接受者
 **/
func (this *BaseSDK) ChatReport(serverId int, channelId int, chatId int, chatMsg string, sender *modelGame.UserBasicInfo, to *modelGame.UserBasicInfo) {
	now := common.GetNowMillisecond()
	param := make(map[string]string)
	param["project"] = this.KyprojectName                //	是	string	项目ID，接入时和数据中心这边确认 单平台单项目:  gandan 单平台多项目:  plat_gameid   xyyou_193、xyplat_234
	param["event"] = "chat"                              //	是	string	聊天日志事件 chat
	param["plat"] = strconv.Itoa(this.KyPlatformId)      //是	string	平台id， 没有需要可以填充-1
	param["sid"] = strconv.Itoa(serverId)                // 是	string	区服id， 没有需要可以填充-1
	param["channel"] = strconv.Itoa(sender.ChannelId)    //是	string	渠道id， 没有需要可以填充-1
	param["ouid"] = sender.OpenId                        //是	string	用户ID
	param["roleid"] = strconv.Itoa(sender.Id)            // 是	string	角色id
	param["rolename"] = sender.NickName                  // 是	string	角色名
	param["type"] = strconv.Itoa(KYCHAT_TYPE[channelId]) //是	int	消息类型等 1 私聊；2 喇叭；3 邮件；4 世界；5 国家；6 工会/帮会；7 队伍；8 附近；9 其他;10 联盟
	param["info"] = chatMsg                              //是	string	聊天内容
	param["timestamp"] = strconv.Itoa(int(now))          //是	bigint	聊天的时间戳 — 13位
	param["vip"] = strconv.Itoa(sender.Vip)              //否	int	vip等级
	param["level"] = strconv.Itoa(sender.Level)          //否	int	角色等级
	if to != nil {
		param["toroleid"] = to.OpenId           //否	string	私聊时，发送给谁的角色id
		param["touserid"] = strconv.Itoa(to.Id) //否	string	私聊时，发送给谁的账号id
		param["torolename"] = to.NickName       //否	string	私聊时，发送给谁的角色名称
	}

	size := len(param)
	keys := make([]string, size+1)
	index := 0
	for k := range param {
		keys[index] = k
		index++
	}
	seconds := int(time.Now().Unix())
	keys[index] = "time"
	sort.Sort(sort.StringSlice(keys))
	paramvalues := ""
	for k, v := range keys {
		if v != "time" {
			paramvalues += fmt.Sprintf("%s=%s", v, url.QueryEscape(param[v]))
		} else {
			paramvalues += fmt.Sprintf("%s=%s", v, strconv.Itoa(seconds))
		}
		if k != size {
			paramvalues += "&"
		}
	}

	sign := this.getSign(this.KySecretKey + paramvalues)
	url := fmt.Sprintf(this.KyChatReport+"/ingestionservice/process?time=%d&sign=%s", time.Now().Unix(), sign)

	rb, err := httpclient.DoJsonPost(url, param)
	if err != nil {
		logger.Error("聊天上报异常：%v", err)
		return
	}
	logger.Info("聊天上报结果,地址：%v,加密数据：%v,结果：%v", url, this.KySecretKey+paramvalues, string(rb))
}

func (this *BaseSDK) checkAndGetBlockParam(r *http.Request) (*KyBlock, string) {

	var blockData KyBlock
	err := json.NewDecoder(r.Body).Decode(&blockData)
	if err != nil {
		r.Body.Close()
		logger.Error("解析平台发送来封禁数据错误：%v", r.Body, err)
		retJson := this.GetSdkResultMsg(400, "封禁数据解析错误", nil)
		return nil, retJson
	}

	timeParam := r.Header.Get("time")
	signParm := r.Header.Get("sign")
	if len(timeParam) <= 0 || len(signParm) <= 0 {
		retJson := this.GetSdkResultMsg(400, "缺少时间或加密参数", nil)
		return nil, retJson
	}

	logger.Info("接收到平台发送来的封禁数据：%v", blockData, timeParam, signParm)

	if blockData.Target == 2 && len(blockData.TargetVal) <= 0 {
		ret := this.GetSdkResultMsg(400, "参数错误", nil)
		return nil, ret
	}

	if !this.CheckSignForKy(signParm, blockData.TargetVal, timeParam) {
		retJson := this.GetSdkResultMsg(400, "密钥验证错误", nil)
		return nil, retJson
	}
	return &blockData, ""
}

/**
 *  @Description:
 *  @param r
 **/
func (this *BaseSDK) Ban(r *http.Request) (*KyBlock, string) {

	blockData, err := this.checkAndGetBlockParam(r)
	if len(err) > 0 {
		return nil, err
	}

	if blockData.Duration <= 0 {
		blockData.Duration = -1
	}

	return blockData, ""

}

/**
 *  @Description:
 *  @param r
 **/
func (this *BaseSDK) BanRemove(r *http.Request) (*KyBlock, string) {

	blockData, err := this.checkAndGetBlockParam(r)
	if len(err) > 0 {
		return nil, err
	}
	return blockData, ""
}

func (this *BaseSDK) GetSdkResultMsg(code int, msg string, data interface{}) string {
	ret := &KyHttpResult{Status: code, Data: "", Message: msg}
	if data != nil {
		ret.Data = data
	}
	retJson, err := json.Marshal(ret)
	if err != nil {
		logger.Error("GetSdkResultMsg json marshal 异常：%v", err)
	}
	return string(retJson)
}

func (this *BaseSDK) MailSend(r *http.Request) (*KyMail, string) {

	r.ParseForm()

	var kyMail KyMail
	err := json.NewDecoder(r.Body).Decode(&kyMail)
	if err != nil {
		r.Body.Close()
		logger.Error("解析平台发送来发送邮件数据错误：%v", r.Body, err)
		retJson := this.GetSdkResultMsg(400, "封禁数据解析错误", nil)
		return nil, retJson
	}

	timeParam := r.Header.Get("time")
	signParm := r.Header.Get("sign")
	if len(timeParam) <= 0 || len(signParm) <= 0 {
		retJson := this.GetSdkResultMsg(400, "缺少时间或加密参数", nil)
		return nil, retJson
	}

	if !this.CheckSignForKy(signParm, kyMail.Id, timeParam) {
		retJson := this.GetSdkResultMsg(400, "密钥验证错误", nil)
		return nil, retJson
	}

	return &kyMail, ""

}

/**
 *  @Description: 礼包兑换码
 *  @param code
 **/
func (this *BaseSDK) ExchangeCode(code string, userId int, userName string, serverId int, channelId int) (map[int]int, error) {

	seconds := int(time.Now().Unix())
	param := make(map[string]interface{})
	param["code"] = code                    //	string	是	礼包兑换码
	param["roleId"] = strconv.Itoa(userId)  //	string	是	角色id
	param["serverId"] = serverId            //	int	否	区服id (区服游戏必填)
	param["serverGroupId"] = 1              //	int	否	区服组id (大区游戏必填 )
	param["appId"] = this.KyAppId           //	int	是	游戏项目id
	param["roleName"] = userName            //	string	是	角色名称
	param["channelId"] = channelId          //	Int	是	渠道id
	param["timeStamp"] = seconds            //	long	是	当前10位时间戳
	param["platformId"] = this.KyPlatformId //	int	是	平台id

	size := len(param)
	keys := make([]string, size)
	index := 0
	for k := range param {
		keys[index] = k
		index++
	}

	sort.Sort(sort.StringSlice(keys))
	paramvalues := ""
	for k, v := range keys {
		val, _ := json.Marshal(param[v])
		paramvalues += fmt.Sprintf("%s=%v", v, string(val))
		if k != size-1 {
			paramvalues += "&"
		}
	}

	sign := this.getSign(paramvalues + this.KyToken)
	param["sign"] = sign

	url := this.KyGmUrl + "/gift/openApi/exchange"
	rb, err := httpclient.DoJsonPost(url, param)
	if err != nil {
		logger.Error("激活码兑换异常：%v", err)
		return nil, gamedb.ERRUNKNOW
	}
	var result ExchangeCodeResultData
	err = json.Unmarshal(rb, &result)
	logger.Info("激活码兑换结果,地址：%v,加密数据：%v,结果：%v", url, paramvalues+this.KyToken, string(rb))

	if err != nil {
		return nil, err
	}

	if result.Code != 0 {
		return nil, gamedb.ERRUNKNOW.CloneWithMsg(result.Message)
	}

	rewards := make(map[int]int)
	for _, v := range result.Data.GiftItemList {
		rewards[v.ItemId] = v.ItemCount
	}
	return rewards, nil
}

/**
 *  @Description: 查询玩家数据
 *  @param r
 *  @return openId
 *  @return userId
 *  @return serverId
 **/
func (this *BaseSDK) GetUserInfo(r *http.Request) (openId string, userId int, username string, serverId int, err string) {

	r.ParseForm()

	openId = r.FormValue("loginName")
	userId, _ = strconv.Atoi(r.FormValue("roleId"))
	username = r.FormValue("roleName")
	serverId, _ = strconv.Atoi(r.FormValue("serverId"))

	timeParam := r.Header.Get("time")
	signParam := r.Header.Get("sign")
	if !this.CheckSignForKy(signParam, strconv.Itoa(serverId)+timeParam) {
		err = this.GetSdkResultMsg(400, "加密验证错误", nil)
		return
	}

	//var param struct {
	//	RoleId   string `json:"roleId"`
	//	RoleName string `json:"roleName"`
	//	OpenId   string `json:"account"`
	//	ServerId int    `json:"serverId"`
	//}
	//err1 := json.NewDecoder(r.Body).Decode(&param)
	//if err1 != nil {
	//	r.Body.Close()
	//	logger.Error("解析平台发送查询玩家数据错误：%v", r.Body, err)
	//	err = this.GetSdkResultMsg(400, "数据解析错误", nil)
	//	return
	//}
	//openId = param.OpenId
	//userId, _ = strconv.Atoi(param.RoleId)
	//username = param.RoleName
	//serverId = param.ServerId
	return
}

/**
 *  @Description: 设置白名单
 *  @param w
 *  @param r
 *  @return whiteId
 *  @return whiteType
 *  @return whiteVal
 *  @return err
 **/
func (this *BaseSDK) SetWhiteBlock(w http.ResponseWriter, r *http.Request) (whiteId int, whiteType int, whiteVal string, err error) {

	var whiteBlockData struct {
		Id        int    `json:"id"`        //是	主键ID(主要用于更新和删除)
		Target    int    `json:"target"`    //是	白名单标签 1:帐号, 2:角色， 3:ip
		TargetVal string `json:"targetVal"` //是
	}
	err = json.NewDecoder(r.Body).Decode(&whiteBlockData)
	if err != nil {
		r.Body.Close()
		logger.Error("解析平台发送来发送邮件数据错误：%v", r.Body, err)
		this.HttpWriteReturnInfo(w, 400, ERR_BASE_PARAM, nil)
		return
	}

	timeParam := r.Header.Get("time")
	signParam := r.Header.Get("sign")
	if !this.CheckSignForKy(signParam, strconv.Itoa(whiteBlockData.Id)+timeParam) {
		err = errors.New(ERR_BASE_SIGN)
		this.HttpWriteReturnInfo(w, 400, ERR_BASE_SIGN, nil)
		return
	}
	whiteId, whiteType, whiteVal, err = whiteBlockData.Id, whiteBlockData.Target, whiteBlockData.TargetVal, nil
	if whiteType == 1 {
		whiteType = 2
	} else if whiteType == 3 {
		whiteType = 1
	} else {
		this.HttpWriteReturnInfo(w, 400, ERR_BASE_OPERATE_TYPE, nil)
		err = errors.New(ERR_BASE_OPERATE_TYPE)
	}
	return
}

/**
 *  @Description: 删除白名单
 *  @param w
 *  @param r
 *  @return whiteId
 *  @return whiteVal
 *  @return err
 **/
func (this *BaseSDK) DelWhiteBlock(w http.ResponseWriter, r *http.Request) (whiteId int, whiteVal string, err error) {
	var whiteBlockData struct {
		Id int `json:"id"` //是	主键ID(主要用于更新和删除)
	}
	err = json.NewDecoder(r.Body).Decode(&whiteBlockData)
	if err != nil {
		r.Body.Close()
		logger.Error("解析平台发送来发送邮件数据错误：%v", r.Body, err)
		this.HttpWriteReturnInfo(w, 400, ERR_BASE_PARAM, nil)
		return
	}

	timeParam := r.Header.Get("time")
	signParam := r.Header.Get("sign")
	if !this.CheckSignForKy(signParam, strconv.Itoa(whiteBlockData.Id)+timeParam) {
		err = errors.New(ERR_BASE_SIGN)
		this.HttpWriteReturnInfo(w, 400, ERR_BASE_SIGN, nil)
		return
	}
	whiteId, whiteVal, err = whiteBlockData.Id, "", nil
	return
}

func (this *BaseSDK) ApplyAnnouncement(w http.ResponseWriter, r *http.Request) (*modelCross.Announcement, *modelCross.PaoMaDeng) {
	var KyAnnouncementstruct struct {
		Id         int    `json:"id"`         //是	公告ID(主要用于更新和删除)
		AnmType    int    `json:"type"`       //是	公告类型: 1 登录公告 2 滚屏公告 3 h5公告
		Weight     int    `json:"weight"`     //是	公告显示顺序：值越大越靠前(登录公告必传)
		Title      string `json:"title"`      //是	标题
		Content    string `json:"content"`    //是	内容
		Begintime  int    `json:"begintime"`  //是	公告上线时间戳,单位为秒
		Endtime    int    `json:"endtime"`    //公告下线时间戳，单位为秒
		Duration   int    `json:"duration"`   //间隔时间（单位分钟,滚屏公告必传）
		Position   []int  `json:"position"`   //否	显示区域(1:跑马灯, 2:聊天框 滚屏公告必传), 如[1,2]
		ChannelIds []int  `json:"channelIds"` //否	渠道id集合(中心服必填)，如[1,2,3]
		ServerIds  []int  `json:"serverIds"`  //否	区服id集合(中心必填)，如[1,2,3]
	}

	err := json.NewDecoder(r.Body).Decode(&KyAnnouncementstruct)
	if err != nil {
		r.Body.Close()
		logger.Error("解析平台发送来发布公告数据错误：%v", r.Body, err)
		this.HttpWriteReturnInfo(w, 400, ERR_BASE_PARAM, nil)
		return nil, nil
	}
	logger.Info("接收到平台发送来申请发布公告：%v", KyAnnouncementstruct)
	if KyAnnouncementstruct.AnmType <= 0 || len(KyAnnouncementstruct.Content) <= 0 || len(KyAnnouncementstruct.Title) <= 0 || KyAnnouncementstruct.Begintime <= 0 || KyAnnouncementstruct.Endtime <= 0 {
		this.HttpWriteReturnInfo(w, 400, ERR_BASE_LOSE_PARAM, nil)
		return nil, nil
	}

	timeParam := r.Header.Get("time")
	signParam := r.Header.Get("sign")
	if !this.CheckSignForKy(signParam, strconv.Itoa(KyAnnouncementstruct.Id)+timeParam) {
		err = errors.New(ERR_BASE_SIGN)
		this.HttpWriteReturnInfo(w, 400, ERR_BASE_SIGN, nil)
		return nil, nil
	}

	if time.Now().Unix() >= int64(KyAnnouncementstruct.Endtime) {
		logger.Error("公告结束时间异常")
		this.HttpWriteReturnInfo(w, 400, ERR_BASE_PARAM_ERR, nil)
		return nil, nil
	}

	if KyAnnouncementstruct.AnmType == 1 || KyAnnouncementstruct.AnmType == 3 {
		announcement := &modelCross.Announcement{
			GmId:         KyAnnouncementstruct.Id,
			Type:         KyAnnouncementstruct.AnmType,
			Title:        KyAnnouncementstruct.Title,
			StartTime:    time.Unix(int64(KyAnnouncementstruct.Begintime), 0),
			EndTime:      time.Unix(int64(KyAnnouncementstruct.Endtime), 0),
			ChannelIds:   KyAnnouncementstruct.ChannelIds,
			ServerIds:    KyAnnouncementstruct.ServerIds,
			Announcement: KyAnnouncementstruct.Content,
		}
		if announcement.Type == 3 {
			announcement.Type = 2
		}
		return announcement, nil
	} else {

		paoMaDeng := &modelCross.PaoMaDeng{
			GmId:          KyAnnouncementstruct.Id,
			Types:         -1,
			CycleTimes:    -1,
			IntervalTimes: KyAnnouncementstruct.Duration,
			Content:       KyAnnouncementstruct.Content,
			StartTime:     time.Unix(int64(KyAnnouncementstruct.Begintime), 0),
			EndTime:       time.Unix(int64(KyAnnouncementstruct.Endtime), 0),
			ChannelIds:    KyAnnouncementstruct.ChannelIds,
			ServerIds:     KyAnnouncementstruct.ServerIds,
		}
		return nil, paoMaDeng
	}

}

/**
*  @Description: 删除指定公告
*  @receiver this
*  @param w
*  @param r
*  @return int
**/
func (this *BaseSDK) DelAnnouncement(w http.ResponseWriter, r *http.Request) int {

	var KyAnnouncementstruct struct {
		Id         int   `json:"id"`          //是	公告ID(主要用于更新和删除)
		ChannelIds []int `json:"channel_ids"` //否	渠道id集合(中心服必填)，如[1,2,3]
	}

	err := json.NewDecoder(r.Body).Decode(&KyAnnouncementstruct)
	if err != nil {
		r.Body.Close()
		logger.Error("解析平台发送来删除公告数据错误：%v", r.Body, err)
		this.HttpWriteReturnInfo(w, 400, ERR_BASE_PARAM, nil)
		return -1
	}
	logger.Info("接收到平台发送来删除发布公告：%v", KyAnnouncementstruct)

	timeParam := r.Header.Get("time")
	signParam := r.Header.Get("sign")
	if !this.CheckSignForKy(signParam, strconv.Itoa(KyAnnouncementstruct.Id)+timeParam) {
		err = errors.New(ERR_BASE_SIGN)
		this.HttpWriteReturnInfo(w, 400, ERR_BASE_SIGN, nil)
		return -1
	}
	return KyAnnouncementstruct.Id
}

/**
 *  @Description: 订阅推送消息
 *  @param openId
 *  @param template
 *  @param arg
 **/
func (this *BaseSDK) Subscribe(openId string, template int, arg ...string){

}
