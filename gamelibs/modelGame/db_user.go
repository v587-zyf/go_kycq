package modelGame

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelCross"
	"cqserver/golibs/logger"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"cqserver/golibs/dbmodel"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gorp.v1"
)

type User struct {
	dbmodel.DbTable
	Id               int            `db:"id" orm:"pk;auto"`
	OpenId           string         `db:"openId" orm:"size(60);comment(openId)"`
	ChannelId        int            `db:"channelId" orm:"comment(渠道ID)"`
	NickName         string         `db:"nickName" orm:"size(50);comment(昵称)"`
	Avatar           string         `db:"avatar"  orm:"comment(头像)"`
	Gold             int            `db:"gold" orm:"type(int64);comment(金币)"`
	Honour           int            `db:"honour" orm:"type(int64);comment(荣誉币)"`
	Ingot            int            `db:"ingot" orm:"comment(元宝)"`
	BindingIngot     int            `db:"bindingIngot" orm:"comment(绑定元宝)"`
	ChuanqiBi        int            `db:"chuanqiBi" orm:"comment(传奇币)"`
	AllIngot         int            `db:"allIngot" orm:"comment(玩家总元宝数)"`
	Exp              int            `db:"exp" orm:"type(int64);comment(经验)"`
	VipLevel         int            `db:"vip"  orm:"comment(vip等级)"`
	VipScore         int            `db:"vipScore" orm:"comment(vip积分)"`
	VipGift          model.IntKv    `db:"vipGift" orm:"comment(vip特权礼包)"`
	Combat           int            `db:"combat" orm:"type(int64);comment(总战力)"`
	ServerIndex      int            `db:"serverIndex" orm:"comment(服务器索引，对应客户端选择的服务器id)"`
	ServerId         int            `db:"serverId" orm:"comment(服务器id)"`
	LastUpdateTime   time.Time      `db:"lastUpdateTime" orm:"comment(更新时间)"`
	CreateTime       time.Time      `db:"createTime" orm:"comment(创角时间)"`
	OfflineTime      time.Time      `db:"offlineTime" orm:"comment(下线时间)"`
	Bag              model.Bag      `db:"bag" orm:"type(text);comment(道具背包)"`
	BagInfo          model.BagInfo  `db:"bagInfo" orm:"type(text);comment(背包信息)"`
	EquipBag         model.EquipBag `db:"equipBag" orm:"type(text);comment(记录背包装备信息)"`
	WarehouseBag     model.Bag      `db:"warehouseBag" orm:"type(text);comment(仓库背包)"`
	WarehouseBagInfo model.BagInfo  `db:"warehouseBagInfo" orm:"type(text);comment(仓库信息)"`

	StageId                 int                        `db:"stageId" orm:"comment(当前在哪个副本)"`
	StageId2                int                        `db:"stageId2" orm:"comment(打完了哪个首领)"`
	StageWave               int                        `db:"stageWave" orm:"comment(副本波数)"`
	MainLineTask            *model.MainLineTask        `db:"mainLineTask" orm:"size(300);comment(当前主线任务进度)"`
	Preference              model.IntStringKv          `db:"preference" orm:"size(500);comment(设置)"`
	Conditions              model.Condition            `db:"conditions" orm:"type(text);comment(所有的条件记录)"`
	Tower                   *model.Tower               `db:"tower" orm:"size(200);comment(爬塔记录)"`
	Rein                    *model.Rein                `db:"rein" orm:"size(100);comment(转生信息)"`
	ReinCosts               model.ReinCosts            `db:"reinCosts" orm:"type(text);comment(转生道具)"`
	Shops                   *model.Shop                `db:"shops" orm:"size(2000);comment(商城列表记录)"`
	Atlases                 model.IntKv                `db:"atlases" orm:"type(text);comment(图鉴信息)"`
	AtlasGathers            model.IntKv                `db:"atlasGathers" orm:"type(text);comment(图鉴集合信息)"`
	FieldBoss               *model.FieldBoss           `db:"fieldBoss" orm:"size(300);comment(野外首领)"`
	ExpStage                *model.ExpStage            `db:"expStage" orm:"size(500);comment(经验副本)"`
	FightModel              int                        `db:"fightModel" orm:"comment(战斗模式)"`
	OnlineAward             *model.OnlineAward         `db:"onlineAward" orm:"null;size(100);comment(在线奖励)"`
	DayStateRecord          *model.DayStateRecord      `db:"dayStateRecord" orm:"null;size(100);comment(每日状态记录)"`
	Panaceas                model.Panaceas             `db:"panaceas" orm:"size(500);comment(灵丹)"`
	Sign                    *model.Sign                `db:"sign" orm:"size(200);comment(签到)"`
	Official                int                        `db:"official" orm:"size(5);comment(官职)"`
	Holyarms                model.Holyarms             `db:"holyarms" orm:"size(500);comment(至尊法器)"`
	Mining                  *model.Mining              `db:"mine" orm:"size(150);comment(挖矿)"`
	Pet                     model.Pets                 `db:"pet" orm:"size(500);comment(战宠)"`
	Wear                    *model.UserWear            `db:"wear" orm:"null;size(150);comment(穿戴)"`
	Juexues                 model.Juexues              `db:"juexues" orm:"size(300);comment(绝学)"`
	SeasonTimes             int                        `db:"seasonTimes" orm:"size(150);comment(竞技场赛季场次)"`
	SeasonWinTimes          int                        `db:"seasonWinTimes" orm:"size(150);comment(竞技场赛季胜场)"`
	SeasonLoseContinueTimes int                        `db:"seasonLoseContinueTimes" orm:"size(150);comment(竞技场连续失败次数)"`
	DarkPalace              *model.DarkPalace          `db:"darkPalace" orm:"size(70);comment(暗殿boss)"`
	PersonBosses            *model.PersonBosses        `db:"personBosses" orm:"size(300);comment(个人boss)"`
	VipBosses               *model.VipBosses           `db:"vipBosses" orm:"size(200);comment(vipBoss)"`
	MaterialStage           *model.MaterialStage       `db:"materialStage" orm:"size(200);comment(材料副本)"`
	GuildData               *model.GuildData           `db:"guildData" orm:"size(500);comment(门派信息)"`
	PaoDian                 *model.PaoDian             `db:"paoDian" orm:"size(100);comment(泡点Pk)"`
	Friend                  model.Friend               `db:"friend" orm:"type(text);comment(好友)"`
	Fit                     *model.Fit                 `db:"fit" orm:"type(text);comment(合体)"`
	RechargeAll             int                        `db:"rechargeAll" orm:"comment(充值总额)"`
	AccumulativeId          model.IntSlice             `db:"accumulativeId" orm:"comment(累计充值领取的奖励id)"`
	DailyTask               *model.DailyTaskInfo       `db:"dailyTask" orm:"type(text);comment(每日任务的活动数据)"`
	MonthCard               *model.MonthCard           `db:"monthCard" orm:"size(300);comment(月卡)"`
	FirstRecharge           *model.FirstRecharge       `db:"firstRecharge" orm:"size(100);comment(首充)"`
	SpendRebates            *model.SpendRebates        `db:"spendRebates" orm:"size(350);comment(累计消费)"`
	Achievement             *model.Achievement         `db:"achievement" orm:"type(text);comment(成就系统)"`
	LimitedGift             *model.LimitGift           `db:"limitedGift" orm:"size(800);comment(限时礼包)"`
	DailyPack               model.DailyPack            `db:"dailyPack" orm:"size(200);comment(每日礼包)"`
	GrowFund                *model.GrowFund            `db:"growFund" orm:"size(200);comment(成长基金)"`
	ChallengeApplyTime      int64                      `db:"challengeApplyTime" orm:"size(200);comment(擂台赛上次报名时间)"`
	RedPacketItem           *model.RedPacketItem       `db:"redPacketItem" orm:"size(200);comment(红包道具)"`
	WarOrder                *model.WarOrder            `db:"warOrder" orm:"type(text);comment(战令)"`
	Elf                     *model.Elf                 `db:"elf" orm:"type(150);comment(精灵)"`
	CardInfo                *model.CardInfo            `db:"cardInfo" orm:"type(text);comment(寻宝 抽卡)"`
	CutTreasure             int                        `db:"curTreasure" orm:"size(5);comment(切割等级)"`
	TreasureInfo            *model.TreasureInfo        `db:"treasureInfo" orm:"type(text);comment(寻宝探宝)"`
	Recharge                model.IntKv                `db:"recharge" orm:"size(100);comment(充值档位)"`
	CompetitiveInfo         *model.CompetitiveInfo     `db:"competitiveInfo" orm:"type(text);comment(竞技场)"`
	FieldFight              *model.FieldFight          `db:"fieldFight" orm:"type(text);comment(野战)"`
	DailyRankInfo           model.DailyRankInfos       `db:"dailyRankInfo" orm:"type(text);comment(每日排行信息)"`
	FitHolyEquip            *model.FitHolyEquip        `db:"fitHolyEquip" orm:"type(text);comment(合体圣装)"`
	HookMapBag              model.Bag                  `db:"hookMapBag" orm:"type(text);comment(挂机奖励)"`
	HookMapTime             int                        `db:"hookMapTime" orm:"type(text);comment(挂机时间)"`
	PreviewFunction         model.IntSlice             `db:"previewFunction" orm:"size(200);comment(功能预览已购买礼包)"`
	PreviewFunctionPoint    model.IntSlice             `db:"previewFunctionPoint" orm:"size(200);comment(功能预览已点击过的)"`
	SevenInvestment         *model.SevenInvestmentInfo `db:"sevenInvestmentInfo"orm:"size(200);comment(七日投资)"`
	ContRecharge            *model.ContRecharge        `db:"contRecharge" orm:"size(350);comment(连冲豪礼)"`
	OpenGift                model.IntKv                `db:"openGift" orm:"size(100);comment(开服礼包)"`
	VipCustomer             int                        `db:"vipCustomer" orm:"comment(vip客服)"`
	TaskChallengeTimes      model.IntKv                `db:"taskChallengeTimes" orm:"type(text);comment(记任务挑战次数)"`
	HaveUseRecharge         int                        `db:"haveUseRecharge" orm:"comment(兑换金锭用掉的充值额度)"`
	GoldIngot               int                        `db:"goldIngot" orm:"comment(金锭)"`
	AncientBoss             *model.AncientBoss         `db:"ancientBoss" orm:"size(70);comment(远古首领)"`
	Title                   model.Title                `db:"title" orm:"type(text);comment(称号)"`
	MiJi                    model.MiJi                 `db:"miJi" orm:"type(text);comment(秘籍)"`
	KillMonster             *model.KillMonster         `db:"killMonster" orm:"type(text);comment(首领首杀)"`
	AncientTreasure         model.AncientTreasuresInfo `db:"ancientTreasure" orm:"type(text);comment(远古宝物)"`
	TreasureShop            *model.TreasureShop        `db:"treasureShop" orm:"type(text);comment(多宝阁)"`
	PetAppendage            model.IntKv                `db:"petAppendage" orm:"size(100);comment(战宠附体)"`
	HellBoss                *model.HellBoss            `db:"hellBoss" orm:"size(70);comment(炼狱首领)"`
	RedPacketNum            int                        `db:"redPacketNum" orm:"comment(使用红包增加的额度)"`
	RedPacketUseNum         int                        `db:"redPacketGetNum" orm:"comment(使用红包的金额)"`
	RedPacketUseNumReset    int                        `db:"redPacketUseNumReset" orm:"comment(使用红包的金额 重置mark)"`
	LotteryInfo             *model.LotteryInfo         `db:"lotteryInfo" orm:"comment(摇彩信息)"`
	TrialTaskInfo           model.TrialTaskInfos       `db:"trialTaskInfos" orm:"comment(试炼之路)"`
	TrialTaskInfoStage      model.IntKv                `db:"trialTaskInfoStage" orm:"comment(试炼之路阶段奖励领取记录)"`
	DaBaoEquip              model.IntKv                `db:"daBaoEquip" orm:"comment(打宝神器)"`
	DaBaoMystery            *model.DaBaoMystery        `db:"daBaoMystery" orm:"size(100);comment(打宝秘境)"`
	AppletsInfo             *model.Applets             `db:"appletsInfo" orm:"size(500);comment(小程序信息)"`
	Label                   *model.Label               `db:"label" omr:"size(100);comment(官职)"`
	ModuleUpMax             *model.ModuleUpMaxLv       `db:"moduleUpMax" orm:"type(text);comment(模块升级到的最大等级)"`
	FirstDropItemInfo       model.IntKv                `db:"firstDropItemInfo" orm:"type(text);comment(装备首暴物品掉落状态)"` //k:firstDrop表id
	FirstDropItemGet        model.IntKv                `db:"firstDropItemGet" orm:"type(text);comment(装备首暴领取状态)"`    //k:firstDrop表id
	Privilege               model.IntKv                `db:"privilege" orm:"size(100);comment(特权)"`
}

// 玩家基础信息
type UserBasicInfo struct {
	Id             int              `db:"id"`
	OpenId         string           `db:"openId"`
	NickName       string           `db:"nickName"`
	Vip            int              `db:"vip"`
	Combat         int              `db:"combat"` //总战力
	Avatar         string           `db:"avatar"` //头像
	ServerId       int              `db:"serverId"`
	LastUpdateTime time.Time        `db:"lastUpdateTime"`
	ChannelId      int              `db:"channelId" orm:"comment(渠道ID)"`
	GuildData      *model.GuildData `db:"guildData" orm:"size(500);comment(门派信息)"`
	HeroDisplay    map[int]*HeroDisplay
	Level          int
}

func (this *User) TableName() string {
	return "user"
}

type UserModel struct {
	dbmodel.CommonModel
}

var (
	userModel             = &UserModel{}
	userFields            = model.GetAllFieldsAsString(User{})
	localServerUserFields = model.GetAllFieldsAsString(UserBasicInfo{})
	idSeqUser             = &modelCross.IdSeq{Name: "user"}
	idSeqEquip            = &modelCross.IdSeq{Name: "equip"}
)

func init() {

	dbmodel.Register(model.DB_SERVER, userModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(User{}, "user").SetKeys(false, "id")
		orm.RegisterModelForAlias(model.DB_SERVER, new(User))
	})
}

func GetUserModel() *UserModel {
	return userModel
}
func NewUser(openId string, channelId, serverId, serverIndex int) *User {
	nowTime := int(time.Now().Unix())
	user := &User{
		OpenId:               openId,
		ChannelId:            channelId,
		ServerId:             serverId,
		ServerIndex:          serverIndex,
		CreateTime:           time.Now(),
		LastUpdateTime:       time.Now(),
		Bag:                  make(model.Bag, 0),
		EquipBag:             make(model.EquipBag),
		Conditions:           make(model.Condition),
		Preference:           make(model.IntStringKv),
		Tower:                &model.Tower{TowerLv: 1, LotteryId: 1, Lottery: make([]int, 0)},
		Rein:                 &model.Rein{},
		ReinCosts:            make(model.ReinCosts),
		Atlases:              make(model.IntKv),
		AtlasGathers:         make(model.IntKv),
		FieldBoss:            &model.FieldBoss{CD: make(model.IntKv)},
		ExpStage:             &model.ExpStage{ExpStages: make(model.IntKv), Appraise: make(model.IntKv)},
		Panaceas:             make(model.Panaceas),
		Sign:                 &model.Sign{SignDay: make(model.IntKv), Cumulative: make(model.IntKv)},
		DayStateRecord:       &model.DayStateRecord{MonthCardReceive: make(model.IntKv)},
		Holyarms:             make(model.Holyarms),
		Mining:               &model.Mining{},
		Pet:                  make(model.Pets),
		Wear:                 &model.UserWear{},
		DarkPalace:           &model.DarkPalace{},
		PersonBosses:         &model.PersonBosses{DareNum: make(model.IntKv)},
		VipBosses:            &model.VipBosses{DareNum: make(model.IntKv)},
		Shops:                &model.Shop{ShopItem: make(model.MapIntKv)},
		GuildData:            &model.GuildData{},
		PaoDian:              &model.PaoDian{},
		VipGift:              make(model.IntKv),
		Friend:               make(model.Friend),
		Fit:                  &model.Fit{Fashion: make(model.IntKv), SkillBag: make(model.IntKv), Lv: make(model.IntKv), Skills: make(map[int]*model.FitSkill)},
		DailyTask:            &model.DailyTaskInfo{GetDayRewardIds: make([]int, 0), GetWeekRewardIds: make([]int, 0), ResourcesHaveBackTimes: make(map[int]int), DailyTask: make(map[int]*model.DailyTaskActivityInfo), ResourceCanBackTimes: make(map[string]int)},
		MonthCard:            &model.MonthCard{MonthCards: make(map[int]*model.MonthCardUnit)},
		FirstRecharge:        &model.FirstRecharge{Days: make(model.IntKv)},
		Juexues:              make(model.Juexues),
		SpendRebates:         &model.SpendRebates{Reward: make(model.IntKv)},
		Achievement:          &model.Achievement{},
		LimitedGift:          &model.LimitGift{GradeStatus: make(model.IntKv), IsBuy: make(map[int]bool), TLv: make(model.IntKv), List: make(map[int]map[int]*model.LimitGiftUnit)},
		DailyPack:            make(model.DailyPack),
		GrowFund:             &model.GrowFund{Ids: make(model.IntKv)},
		RedPacketItem:        &model.RedPacketItem{PickInfo: make(model.IntKv)},
		WarOrder:             &model.WarOrder{Lv: 1, Exchange: make(model.IntKv), Task: make(map[int]*model.WarOrderTask), WeekTask: make(map[int]map[int]*model.WarOrderTask), Reward: make(map[int]*model.WarOrderReward)},
		Elf:                  &model.Elf{Skills: make(model.IntKv), SkillBag: make(model.IntKv)},
		CardInfo:             &model.CardInfo{},
		TreasureInfo:         &model.TreasureInfo{AllGetRound: make(model.IntSlice, 0), BuyTimes: make(model.IntKv), ChooseItems: make(map[int]model.IntSlice), HaveRandomItems: make(map[int]model.IntSlice)},
		Recharge:             make(model.IntKv),
		CompetitiveInfo:      &model.CompetitiveInfo{},
		FieldFight:           &model.FieldFight{},
		DailyRankInfo:        make(model.DailyRankInfos),
		FitHolyEquip:         &model.FitHolyEquip{Equips: make(model.MapIntKv)},
		PreviewFunction:      make(model.IntSlice, 0),
		PreviewFunctionPoint: make(model.IntSlice, 0),
		SevenInvestment:      &model.SevenInvestmentInfo{GetAwardIds: make([]int, 0)},
		ContRecharge:         &model.ContRecharge{Receive: make(model.IntKv), Day: make(model.IntKv)},
		OpenGift:             make(model.IntKv),
		TaskChallengeTimes:   make(model.IntKv),
		MaterialStage:        &model.MaterialStage{MaterialStages: make(map[int]*model.MaterialStageUnit)},
		AncientBoss:          &model.AncientBoss{},
		Title:                make(model.Title),
		MiJi:                 make(model.MiJi),
		KillMonster:          &model.KillMonster{Uni: make(map[int]*model.KillMonsterUni), Mil: make(map[int]*model.KillMonsterMil)},
		AncientTreasure:      make(model.AncientTreasuresInfo),
		TreasureShop:         &model.TreasureShop{Shop: make(model.IntKv), Car: make(model.IntKv)},
		PetAppendage:         make(model.IntKv),
		HellBoss:             &model.HellBoss{},
		LotteryInfo:          &model.LotteryInfo{},
		TrialTaskInfo:        make(model.TrialTaskInfos),
		TrialTaskInfoStage:   make(model.IntKv),
		DaBaoEquip:           make(model.IntKv),
		DaBaoMystery:         &model.DaBaoMystery{Energy: gamedb.GetConf().DaBaoMysteryEnergy, ResumeTime: nowTime},
		AppletsInfo:          &model.Applets{Energy: gamedb.GetConf().XiaoYouXiEnergy, ResumeTime: time.Now().Unix(), List: make(map[int]*model.AppletsUnit)},
		Label:                &model.Label{Id: 1, TaskOver: make(model.IntKv)},
		ModuleUpMax:          &model.ModuleUpMaxLv{},
		FirstDropItemInfo:    make(model.IntKv),
		FirstDropItemGet:     make(model.IntKv),
		Privilege:            make(model.IntKv),
	}
	return user
}

func (this *UserModel) Create(user *User) error {
	var err error
	user.Id, err = idSeqUser.Next()
	if err != nil {
		return err
	}
	return this.DbMap().Insert(user)
}

func (this *UserModel) GetByOpenId(openId string, serverId int) (*User, error) {
	var user User
	err := this.DbMap().SelectOne(&user, fmt.Sprintf("select %s from user where openId = ? and serverId= ?", userFields), openId, serverId)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (this *UserModel) GetByUserId(userId int) (*User, error) {
	var user User
	err := this.DbMap().SelectOne(&user, fmt.Sprintf("select %s from user where id = ?", userFields), userId)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (this *UserModel) GetUserIdsByOpenId(openId string) []int {
	var userIds []int
	_, err := this.DbMap().Select(&userIds, "select id from user where openId = ?", openId)
	if err != nil {
		logger.Error("获取玩家Id数据异常,openId:%v,err：%v", openId, err)
		return userIds
	}
	return userIds
}

// 加载本服所有玩家的UserId
func (this *UserModel) LoadAllUsers() ([]*UserBasicInfo, error) {
	sql := fmt.Sprintf("select %s from user where 1", localServerUserFields)
	var users []*UserBasicInfo
	_, err := this.DbMap().Select(&users, sql)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (this *UserModel) LoadAllUsersCombat(activityTime int) ([]*UserBasicInfo, error) {
	sql := fmt.Sprintf("select id,combat from user where UNIX_TIMESTAMP(now()) - UNIX_TIMESTAMP(offlineTime) <= %v", activityTime)
	var users []*UserBasicInfo
	_, err := this.DbMap().Select(&users, sql)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (this *UserModel) GetEquipId() (int, error) {
	return idSeqEquip.Next()
}

func (this *UserModel) SearchName(name string) ([]*UserBasicInfo, error) {
	name = strings.ReplaceAll(name, `"`, "")
	sql := fmt.Sprintf(`select id from user where nickName like "%%%s%%"`, name)
	var users []*UserBasicInfo
	_, err := this.DbMap().Select(&users, sql)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (this *UserModel) SearchActivityUsers(times int) ([]*User, error) {
	sql := fmt.Sprintf(`SELECT %v FROM user WHERE  UNIX_TIMESTAMP(NOW())  - UNIX_TIMESTAMP(lastUpdateTime)   < %v  `, userFields, times)
	var users []*User
	_, err := this.DbMap().Select(&users, sql)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (this *UserModel) GetFriendInfoByUserId(userId int) (model.Friend, error) {
	var friend model.Friend
	dataStr, err := this.DbMap().SelectStr(fmt.Sprintf("select %s from user where id = ?", "friend"), userId)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(dataStr), &friend)
	if err != nil {
		return nil, err
	}
	return friend, nil
}

func (this *UserModel) GetUserIdsForGmMail(highVip, lowVip int, highLevel, lowLevel int, highRecharge, lowRecharge int) []int {

	var userIds []int

	t := time.Now().Add(time.Duration(-1*24*10) * time.Hour).Format("2006-01-02 15:04:05")
	sql := "SELECT u.id from `user` as u LEFT JOIN hero as h on u.id = h.userId WHERE u.lastUpdateTime>\"" + t + "\""

	if highVip >= lowVip {
		if highVip > 0 {
			sql += fmt.Sprintf(" and u.vip <= %d", highVip)
		}
		if lowVip > 0 {
			sql += fmt.Sprintf(" and u.vip >= %d", lowVip)
		}
	}

	if highRecharge >= lowRecharge {
		if highRecharge > 0 {
			sql += fmt.Sprintf(" and u.rechargeAll<=%d", highRecharge*100)
		}
		if lowRecharge > 0 {
			sql += fmt.Sprintf(" and u.rechargeAll >= %d", lowRecharge*100)
		}
	}

	if highLevel >= lowLevel {
		if highLevel > 0 {
			sql += fmt.Sprintf(" and h.ExpLvl<=%d", highLevel)
		}
		if lowLevel > 0 {
			sql += fmt.Sprintf(" and h.ExpLvl >= %d", lowLevel)
		}
	}

	sql += " group by u.id;"

	logger.Info("sql:%v", sql)
	_, err := this.DbMap().Select(&userIds, sql)
	if err != nil {
		logger.Error("获取GM邮件指定条件玩家Id数据异常,openId:%v,err：%v", highVip, lowVip, highLevel, lowLevel, highRecharge, lowRecharge, err)
	}
	return userIds

}
