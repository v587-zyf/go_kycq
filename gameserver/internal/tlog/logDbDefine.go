package tlog

// (必填)服务器状态流水，每5分钟一条日志
type LogGameSvrState struct {
	DtEventTime string `db:"dtEventTime"` // (必填) 格式 YYYY-MM-DD HH:MM:SS, type: datetime, size:
	VGameIp     string `db:"vGameIP"`     // (必填)服务器IP, type: string, size: 32
	IZoneAreaId int    `db:"iZoneAreaID"` // (必填)针对分区分服的游戏填写分区id，用来唯一标示一个区；非分区分服游戏请填写0, type: int, size:
}

type LogCommon struct {
	Id          int    `db:"id"`          // (必填)自增
	DtEventTime string `db:"dtEventTime"` // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS, type: datetime, size:
	UserId      int    `db:"UserId"`      // (必填)玩家唯一编号
	ServerId    int    `db:"serverId"`    // (必填)登录的游戏服务器编号, type: string, size: 25
	Openid      string `db:"openid"`      // (必填)用户OPENID号, type: string, size: 64
}

// (必填)玩家注册
type LogPlayerRegister struct {
	LogCommon
}

// (必填)玩家登陆
type LogPlayerLogin struct {
	LogCommon
}

// (必填)玩家登出
type LogPlayerLogout struct {
	LogCommon
}

// (必填)货币流水
type LogMoneyFlow struct {
	Id          int    `db:"id"`          // (必填)自增
	UserId      int    `db:"UserId"`      // (必填)玩家唯一编号
	GameSvrId   string `db:"GameSvrId"`   // (必填)登录的游戏服务器编号, type: string, size: 25
	DtEventTime string `db:"dtEventTime"` // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS, type: datetime, size:
	VGameAppid  string `db:"vGameAppid"`  // (必填)游戏APPID, type: string, size: 32
	PlatId      int    `db:"PlatID"`      // (必填)ios 0/android 1, type: int, size:
	IZoneAreaId int    `db:"iZoneAreaID"` // (必填)针对分区分服的游戏填写分区id，用来唯一标示一个区；非分区分服游戏请填写0, type: int, size:
	Vopenid     string `db:"vopenid"`     // (必填)用户OPENID号, type: string, size: 64
	Sequence    int    `db:"Sequence"`    // (可选)用于关联一次动作产生多条不同类型的货币流动日志, type: int, size:
	Level       int    `db:"Level"`       // (必填)玩家等级, type: int, size:
	AfterMoney  int    `db:"AfterMoney"`  // (可选)动作后的金钱数, type: int, size:
	IMoney      int    `db:"iMoney"`      // (必填)动作涉及的金钱数, type: int, size:
	Reason      int    `db:"Reason"`      // (必填)货币流动一级原因, type: int, size:
	SubReason   int    `db:"SubReason"`   // (可选)货币流动二级原因, type: int, size:
	AddOrReduce int    `db:"AddOrReduce"` // (必填)增加 0/减少 1, type: int, size:
	IMoneyType  int    `db:"iMoneyType"`  // (必填)钱的类型MONEYTYPE,其它货币类型参考FAQ文档, type: int, size:
}

// (必填)道具流水表
type LogItemFlow struct {
	LogCommon
	UserLv      int    `db:"UserLv"`      // 用户等级
	IGoodsId    int    `db:"iGoodsId"`    // (必填)道具ID, type: int, size:
	Count       int    `db:"Count"`       // (必填)数量, type: int, size:
	AfterCount  int    `db:"AfterCount"`  // (必填)动作后的物品存量, type: int, size:
	Reason      string `db:"Reason"`      // (必填)道具流动一级原因, type: string, size:
	Reason2     string `db:"Reason2"`     // (必填)道具流动二级原因, type: string, size:
	AddOrReduce int    `db:"AddOrReduce"` // (必填)增加 0/减少 1, type: int, size:
}

// (可选)人物等级流水表
type LogPlayerLvlFlow struct {
	Id          int    `db:"id"`          // (必填)自增
	UserId      int    `db:"UserId"`      // (必填)玩家唯一编号
	GameSvrId   string `db:"GameSvrId"`   // (必填)登录的游戏服务器编号, type: string, size: 25
	DtEventTime string `db:"dtEventTime"` // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS, type: datetime, size:
	VGameAppid  string `db:"vGameAppid"`  // (必填)游戏APPID, type: string, size: 32
	PlatId      int    `db:"PlatID"`      // (必填)ios 0/android 1, type: int, size:
	IZoneAreaId int    `db:"iZoneAreaID"` // (必填)针对分区分服的游戏填写分区id，用来唯一标示一个区；非分区分服游戏请填写0, type: int, size:
	Vopenid     string `db:"vopenid"`     // (必填)玩家, type: string, size: 64
	ExpChange   int    `db:"ExpChange"`   // (必填)经验变化, type: int, size:
	BeforeLevel int    `db:"BeforeLevel"` // (可选)动作前等级, type: int, size:
	AfterLevel  int    `db:"AfterLevel"`  // (必填)动作后等级, type: int, size:
	Time        int    `db:"Time"`        // (必填)升级所用时间(秒), type: int, size:
	Reason      int    `db:"Reason"`      // (必填)经验流动一级原因, type: int, size:
	SubReason   int    `db:"SubReason"`   // (必填)经验流动二级原因, type: int, size:
}

// (可选)VIP等级流水表
type LogVipLevelFlow struct {
	Id              int    `db:"id"`              // (必填)自增
	UserId          int    `db:"UserId"`          // (必填)玩家唯一编号
	GameSvrId       string `db:"GameSvrId"`       // (必填)登录的游戏服务器编号, type: string, size: 25
	DtEventTime     string `db:"dtEventTime"`     // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS, type: datetime, size:
	VGameAppid      string `db:"vGameAppid"`      // (必填)游戏APPID, type: string, size: 32
	PlatId          int    `db:"PlatID"`          // (必填)ios 0/android 1, type: int, size:
	IZoneAreaId     int    `db:"iZoneAreaID"`     // (必填)针对分区分服的游戏填写分区id，用来唯一标示一个区；非分区分服游戏请填写0, type: int, size:
	Vopenid         string `db:"vopenid"`         // (必填)玩家, type: string, size: 64
	ILevel          int    `db:"iLevel"`          // (必填)玩家等级, type: int, size:
	IBeforeVipLevel int    `db:"iBeforeVipLevel"` // (必填)动作前等级, type: int, size:
	IAfterVipLevel  int    `db:"iAfterVipLevel"`  // (必填)动作后等级, type: int, size:
	VRoleId         string `db:"vRoleID"`         // (必填)玩家角色ID, type: string, size: 64
}

// 战斗力变化表
type LogCombatFlow struct {
	Id                int    `db:"id"`                // (必填)自增
	UserId            int    `db:"UserId"`            // (必填)玩家唯一编号
	GameSvrId         string `db:"GameSvrId"`         // (必填)登录的游戏服务器编号, type: string, size: 25
	VGameAppid        string `db:"vGameAppid"`        // (必填)游戏APPID, type: string, size: 64
	DtEventTime       string `db:"dtEventTime"`       // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS, type: datetime, size:
	VOpenId           string `db:"vOpenID"`           // (必填)操作玩家的openId, type: string, size: 64
	IPlatId           int    `db:"iPlatId"`           // (必填)平台：IOS（0），安卓（1）, type: int, size:
	IPartition        int    `db:"iPartition"`        // (必填)小区, type: int, size:
	IBeforeCombat     int    `db:"iBeforeCombat"`     // 变化前战斗力, type: int, size:
	IAfterCombat      int    `db:"iAfterCombat"`      // 变化后战斗力, type: int, size:
	ICombatType       int    `db:"iCombatType"`       // 战斗力变化类型
	IPartBeforeCombat int    `db:"iPartBeforeCombat"` // 部分变化前战斗力, type: int, size:
	IPartAfterCombat  int    `db:"iPartAfterCombat"`  // 部分变化后战斗力, type: int, size:
	ILevel            int    `db:"iLevel"`            // (必填)玩家等级, type: int, size:
}

// (可选)任务流水
type LogTaskFlow struct {
	Id          int    `db:"id"`          // (必填)自增
	UserId      int    `db:"UserId"`      // (必填)玩家唯一编号
	GameSvrId   string `db:"GameSvrId"`   // (必填)登录的游戏服务器编号, type: string, size: 25
	DtEventTime string `db:"dtEventTime"` // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS, type: datetime, size:
	VGameAppid  string `db:"vGameAppid"`  // (必填)游戏APPID, type: string, size: 32
	PlatId      int    `db:"PlatID"`      // (必填)ios 0/android 1, type: int, size:
	IZoneAreaId int    `db:"iZoneAreaID"` // (必填)针对分区分服的游戏填写分区id，用来唯一标示一个区；非分区分服游戏请填写0, type: int, size:
	Vopenid     string `db:"vopenid"`     // (必填)玩家, type: string, size: 64
	ILevel      int    `db:"iLevel"`      // (必填)VIP等级, type: int, size:
	IVipLevel   int    `db:"iVipLevel"`   // (必填)VIP等级, type: int, size:
	ITaskType   int    `db:"iTaskType"`   // (必填)任务类型, type: int, size:
	ITaskId     int    `db:"iTaskID"`     // (必填)任务ID, type: int, size:
	IState      int    `db:"iState"`      // (必填)任务操作状态, type: int, size:
	VRoleId     string `db:"vRoleID"`     // (必填)玩家角色ID, type: string, size: 64
}

//装备强化表
type LogEquipStrengthen struct {
	Id           int    `db:"id"`           // (必填)自增
	UserId       int    `db:"UserId"`       // (必填)玩家唯一编号
	GameSvrId    string `db:"GameSvrId"`    // (必填)登录的游戏服务器编号, type: string, size: 25
	DtEventTime  string `db:"dtEventTime"`  // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS, type: datetime, size:
	VGameAppid   string `db:"vGameAppid"`   // (必填)游戏APPID, type: string, size: 64
	PlatId       int    `db:"PlatID"`       // (必填)0是ios，1是安卓, type: int, size:
	IZoneAreaId  int    `db:"iZoneAreaID"`  // (必填)针对分区分服的游戏填写分区id，用来唯一标示一个区；非分区分服游戏请填写0, type: int, size:
	Vopenid      string `db:"vopenid"`      // (必填)玩家, type: string, size: 64
	ILevel       int    `db:"iLevel"`       // (必填)玩家等级, type: int, size:
	IVipLevel    int    `db:"iVipLevel"`    // (必填)VIP等级, type: int, size:
	IEquipPos    int    `db:"iEquipPos"`    // (必填)装备部位 参考 iEQUIPPOSITION, type: int, size:
	IBeforeLevel int    `db:"iBeforeLevel"` // (必填)变化前等级, type: int, size:
	IAfterLevel  int    `db:"iAfterLevel"`  // (必填)变化后等级, type: int, size:
}

// 法宝激活
type LogFabaoActive struct {
}

// 法宝升级
type LogFabaoUpLevel struct {
}

// 法宝技能激活
type LogFabaoSkillActive struct {
}

// 转生激活
type LogReinActive struct {
}

// 转生
type LogReincarnation struct {
}

// 购买修为丹
type LogReinCostBuy struct {
}

// 使用修为丹
type LogReinCostUse struct {
}

// 神兵激活
type LogArtifactActive struct {
}

// 神兵穿戴
type LogArtifactWear struct {
}

// 神兵升阶
type LogArtifactUpAdvance struct {
}

// 神兵升级
type LogArtifactUpLevel struct {
}

// 羽翼升级|升阶
type LogWingUpLevel struct {
}

// 图鉴激活
type LogAtlasActive struct {
}

// 图鉴升星
type LogAtlasUpStar struct {
}

// 图鉴集合激活
type LogAtlasGatherActive struct {
}

// 图鉴集合升星
type LogAtlasGatherUpLevel struct {
}
