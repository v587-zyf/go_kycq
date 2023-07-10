package ptsdk

import (
	"cqserver/gamelibs/beans"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constPlatfrom"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"fmt"
	"net/http"
	"net/url"
)

const (
	STATUS_SUCCESS = 1
	STATUS_FAIL    = 0
)

//1 私聊；2 喇叭；3 邮件；4 世界；5 国家；6 工会/帮会；7 队伍；8 附近；9 其他;10 联盟
var KYCHAT_TYPE = map[int]int{
	pb.CHATTYPE_WORLD:   4,
	pb.CHATTYPE_TEAM:    7,
	pb.CHATTYPE_GUILD:   6,
	pb.CHATTYPE_PRIVATE: 1,
}

var sdkApi ISDK

type ISDK interface {
	SetSandBox(sandbox bool)
	GetPlatform() int
	GetOpenId(r *http.Request) (string, error)
	CheckSignForPlatform(sign string, arg ...interface{}) bool
	CheckSignForKy(sign string, arg ...interface{}) bool
	/**
	 *  @Description: 生成充值数据
	 *  @param serverId	服务器Id
	 *  @param doid		游戏订单号
	 *  @param money		充值金额（元）
	 *  @return string
	 **/
	GetRechargeData(serverId int, userName string, Lv int, order *modelGame.OrderDb, trialServer bool) string

	/**
	 *  @Description: 平台通知充值结果验证
	 *  @param w
	 *  @param r
	 *  @return *pbserver.RechageCcsToGsReq //充值数据
	 **/
	NotifierRecharge(w http.ResponseWriter, r *http.Request) (int, *pbserver.RechageCcsToGsReq)

	/**
	 *  @Description: 平台申请充值
	 *  @param w
	 *  @param r
	 *  @return int
	 *  @return *pbserver.RechargeApplyReq
	 **/
	ApplyPay(w http.ResponseWriter, r *http.Request) (int, *pbserver.RechargeApplyReq)
	/**
	 *  @Description: 聊天上报
	 *  @param serverId		服务器Id
	 *  @param channelId		聊天频道
	 *  @param chatId		聊天Id
	 *  @param chatMsg		聊天内容
	 *  @param sender		发送者
	 *  @param to			接受者
	 **/
	ChatReport(serverId int, channelId int, chatId int, chatMsg string, sender *modelGame.UserBasicInfo, to *modelGame.UserBasicInfo)

	/**
	 *  @Description:封禁
	 *  @param r
	 **/
	Ban(r *http.Request) (*KyBlock, string)

	/**
	 *  @Description:解封
	 *  @param r
	 **/
	BanRemove(r *http.Request) (*KyBlock, string)

	/**
	 *  @Description: 创建http返回消息
	 *  @param code
	 *  @param msg
	 *  @return string
	 **/
	GetSdkResultMsg(code int, msg string, data interface{}) string

	/**
	 *  @Description: 邮件发送
	 *  @param r
	 *  @param arg
	 *  @return *KyMail
	 *  @return string
	 **/
	MailSend(r *http.Request) (*KyMail, string)

	/**
	 *  @Description: 礼包兑换码
	 *  @param code
	 **/
	ExchangeCode(code string, userId int, userName string, serverId int, channelId int) (map[int]int, error)

	/**
	 *  @Description: 查询玩家数据
	 *  @param r
	 *  @return openId
	 *  @return userId
	 *  @return serverId
	 **/
	GetUserInfo(r *http.Request) (openId string, userId int, username string, serverId int, err string)

	/**
	 *  @Description: 设置白名单
	 *  @param w
	 *  @param r
	 *  @return whiteId
	 *  @return whiteType
	 *  @return whiteVal
	 *  @return err
	 **/
	SetWhiteBlock(w http.ResponseWriter, r *http.Request) (whiteId int, whiteType int, whiteVal string, err error)

	/**
	 *  @Description: 删除白名单
	 *  @param w
	 *  @param r
	 *  @return whiteId
	 *  @return whiteVal
	 *  @return err
	 **/
	DelWhiteBlock(w http.ResponseWriter, r *http.Request) (whiteId int, whiteVal string, err error)

	/**
	 *  @Description: 申请发布公告
	 *  @param w
	 *  @param r
	 *  @return *modelCross.Announcement
	 **/
	ApplyAnnouncement(w http.ResponseWriter, r *http.Request) (*modelCross.Announcement, *modelCross.PaoMaDeng)

	/**
	 *  @Description: 删除公告
	 *  @param w
	 *  @param r
	 *  @return int
	 **/
	DelAnnouncement(w http.ResponseWriter, r *http.Request) int

	/**
	 *  @Description: 回写http消息
	 *  @param w
	 *  @param msg
	 **/
	HttpWriteReturnInfo(w http.ResponseWriter, code int, msg string, data interface{})

	/**
	 *  @Description: 回写http消息
	 *  @param w
	 *  @param msg
	 **/
	HttpWriteReturnMsg(w http.ResponseWriter, msg string)

	/**
	 *  @Description: 订阅推送消息
	 *  @param openId
	 *  @param template
	 *  @param arg
	 **/
	Subscribe(openId string, template int, arg ...string)
}

func InitSdk(sdkconfig *beans.Sdkconfig, sandBox bool) {

	switch sdkconfig.KyPlatformId {
	case constPlatfrom.PLATFORM_602:
		sdkApi = initSdk602(sdkconfig)
	case constPlatfrom.PLATFORM_JUNSHANG:
		sdkApi = initSdkJunshang(sdkconfig)
	case constPlatfrom.PLATFORM_XY:
		sdkApi = initXY(sdkconfig)
	case constPlatfrom.PLATFORM_XY_WX:
		sdkApi = initXYWX(sdkconfig)
	default:
		sdkApi = &BaseSDK{Sdkconfig: sdkconfig, sandbox: sandBox}
		logger.Error("平台sdk未实现")
		return
	}
	sdkApi.SetSandBox(sandBox)
	logger.Info("平台sdk初始化：%v", sdkconfig.KyPlatformId, sdkconfig.KyAppId)
}

func ToUrlValues(param map[string]interface{}) url.Values {
	values := make(url.Values)
	for k, v := range param {
		values.Add(k, fmt.Sprintf("%v", v))
	}
	return values
}

func GetSdk() ISDK {
	return sdkApi
}

type KyHttpResult struct {
	Status  int         `json:"status"`  //是	请求状态码，200：成功；其他失败；
	Data    interface{} `json:"data"`    //否
	Message string      `json:"message"` //是	请求说明
}

type KyBlock struct {
	Target     int    `json:"target"`     //是	封禁目标 1:帐号, 2:角色， 3：ip
	TargetVal  string `json:"targetVal"`  //是	封禁对象值（如帐号，角色Id,ip值)
	Duration   int    `json:"duration"`   //是	封禁时长（单位毫秒 -1:永久封禁)
	EndTime    int    `json:"endTime"`    //否	封禁截止时间（单位毫秒 -1:永久封禁)
	BlockType  int    `json:"type"`       //是	封禁操作 1 禁登 2 禁言
	Reason     string `json:"reason"`     //是	封禁原因
	ServerId   int    `json:"serverId"`   //否	区服id， target=2时必传
	ChannelIds []int  `json:"channelIds"` //否	渠道id集合(中心服模式必填)，注意当前参数为历史版本保留字段，新游戏可直接用serverIds参数分发，如[1,2,3]
	ServerIds  []int  `json:"serverIds"`  //否	区服id集合(中心服模式必填)，如[1,2,3]
}

type KyMail struct {
	Id          string `json:"id"`          //6201331312232", #邮件唯一id
	Title       string `json:"title"`       //": "主题", #邮件主题
	Content     string `json:"content"`     // "内容", #邮件内容
	Maitype     int    `json:"type"`        //": 1,  #1为平台 2为渠道 3为区服组 4区服 5为个人
	ValidityDay int    `json:"validityDay"` //": 1, #有效期 天数
	Target      struct { //发送目标
		PlatformIds []int    `json:"target"`     //":[1],  #平台id
		RoleIds     []string `json:"roleIds"`    //": ["100001"], #角色id
		ChannelIds  []int    `json:"channelIds"` //":[1], #渠道id
		ServerIds   []int    `json:"serverIds"`  //":[1], #区服id
	} `json:"target"`
	Filter struct {
		HighVip         int  `json:"highVip"`         //": 20,  #vip区间 高Vip
		LowVip          int  `json:"lowVip"`          //1, #低vip
		HighDegree      int  `json:"highDegree"`      //100, #等级区间 高等级
		LowDegree       int  `json:"lowDegree"`       //1, #低等级
		IsOnline        bool `json:"isOnline"`        //true, #是否在线玩家 false 否 true 是
		HighConsumption int  `json:"highConsumption"` //100, #充值区间 高
		LowConsumption  int  `json:"lowConsumption"`  //0, #充值区间 低
		SenderId        int  `json:"senderId"`        //0 #发件人id
	} `json:"filter"`
	ItemList []struct {
		Id    int `json:"id"`    //物品id
		Count int `json:"count"` //数量
		Bind  int `json:"bind"`  //是否绑定（1 是 2 否）
	} `json:"itemList"`
}

type ExchangeCodeResultData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RoleId       string `json:"roleId"`
		GiftName     string `json:"giftName"`
		GiftItemList []struct {
			ItemId    int `json:"id"`
			ItemCount int `json:"count"`
		} `json:"giftItemList"`
	} `json:"data"`
}
