package model

import (
	"encoding/json"
)

type CrossHeros []CrossHero
type CrossHero struct {
	HeroIndex int `json:"heroIndex"`
	Sex       int `json:"sex"`
	Job       int `json:"job"`
	Level     int `json:"level"`
	Combat    int `json:"combat"`
}

type AccountBan map[int]*BanInfo

type BanInfo struct {
	StartTime string
	EndTime   string
	Reason    string
	BanType   int
}

type Property struct {
	Id    int
	Count int
}

type Equip struct {
	Index     int
	ItemId    int
	RandProps []*EquipRandProp
	IsLock    bool
	Lucky     int //武器幸运值
}
type EquipRandProp struct {
	PropId int
	Color  int
	Value  int
}

type EquipClearUnit struct {
	Grade  int
	Color  int
	PropId int
	Value  int
}
type EquipClears map[int][]*EquipClearUnit

type Item struct {
	ItemId     int
	Count      int
	Position   int
	EquipIndex int
}
type ItemMark struct {
	ItemId      int
	Count       int
	Index       int
	EquipIndex  int
	Class       int //阶数
	Quality     int //品质
	Star        int //星数
	ItemValue   int //item表value
	ItemQuality int //item表品质
	ItemCfgId   int //item表id
	Combat      int //装备战力
	EquipType   int //哪个部位的装备
	CountLimit  int
	ItemType    int
}

type ItemMarkSlice []*ItemMark

func (this ItemMarkSlice) Len() int {
	return len(this)
}

func (this ItemMarkSlice) Less(i, j int) bool {
	if this[i].Class > this[j].Class {
		return true
	} else if this[i].Class == this[j].Class {
		if this[i].Quality > this[j].Quality {
			return true
		} else if this[i].Quality == this[j].Quality {
			if this[i].Star > this[j].Star {
				return true
			} else if this[i].Star == this[j].Star {
				if this[i].Count >= this[j].Count {
					return true
				}
			}
		}
	}
	return false
}

func (this ItemMarkSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type ItemMarkSlice2 []*ItemMark

func (this ItemMarkSlice2) Len() int {
	return len(this)
}

func (this ItemMarkSlice2) Less(i, j int) bool {
	if this[i].ItemValue < this[j].ItemValue {
		return true
	} else if this[i].ItemValue == this[j].ItemValue {
		if this[i].ItemQuality > this[j].ItemQuality {
			return true
		} else if this[i].ItemQuality == this[j].ItemQuality {
			if this[i].ItemCfgId < this[j].ItemCfgId {
				return true
			}
		}
	}
	return false
}

func (this ItemMarkSlice2) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type ItemMarkSlice3 []*ItemMark

func (this ItemMarkSlice3) Len() int {
	return len(this)
}

func (this ItemMarkSlice3) Less(i, j int) bool {
	if this[i].Combat > this[j].Combat {
		return true
	}
	return false
}

func (this ItemMarkSlice3) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type Fabao struct {
	Id    int
	Level int
	Exp   int
	Skill []int
}
type Fabaos map[int]*Fabao

type GodEquip struct {
	Id    int
	Lv    int
	Blood int
}
type GodEquips map[int]*GodEquip

type Juexue struct {
	Id int
	Lv int
}
type Juexues map[int]*Juexue

type DarkPalace struct {
	DareNum   int
	BuyNum    int
	ResetTime int
	HelpNum   int //协助次数
}

type Holyarm struct {
	Level int
	Exp   int
	Skill IntKv
}
type Holyarms map[int]*Holyarm

type Mining struct {
	WorkTime  int //挖矿开始时间
	WorkNum   int //挖矿次数(已有)
	RobNum    int //掠夺次数(已有)
	BuyNum    int //购买次数(已有)
	Miner     int //矿工等级
	Luck      int //幸运值
	ResetTime int //每日重置
}

type Wing struct {
	Id     int
	Exp    int
	IsWear bool
}
type Wings map[int]*Wing

type Rein struct {
	Id  int
	Exp int
}
type ReinCost struct {
	Id   int
	Num  int
	Date int
}
type ReinCosts map[int]*ReinCost

type FieldBoss struct {
	DareNum      int
	BuyNum       int
	ResetTime    int
	CD           IntKv //stageId,cd时间
	FirstReceive bool
}

type AncientBoss struct {
	DareNum   int
	BuyNum    int
	ResetTime int
}

type AncientSkill struct {
	SkillId int
	Level   int
	Grade   int
}

type MaterialStageUnit struct {
	DareNum   int
	BuyNum    int
	NowLayer  int //当前到达层
	LastLayer int //通过层
}
type MaterialStage struct {
	ResetTime      int
	MaterialStages map[int]*MaterialStageUnit
}

type VipBosses struct {
	DareNum   IntKv
	ResetTime int
}

type ExpStage struct {
	DareNum   int
	BuyNum    int
	ResetTime int
	ExpStages IntKv
	Appraise  IntKv
	Layer     int
}

type PaoDian struct {
	EndTime int //最后进入时间
}

type MsgLog struct {
	Msg  string
	Time int
	IsMy bool //是否本人发送
}
type MsgLogs map[int]*MsgLog
type FriendUnit struct {
	MsgLog    MsgLogs
	BlockTime int
	CreatedAt int
	DeletedAt int
	IsRead    bool //是否有未阅读消息
}
type Friend map[int]*FriendUnit

type FitSkill struct {
	Lv   int
	Star int
}
type Fit struct {
	CdStart  int
	CdEnd    int
	Fashion  IntKv //时装id，等级
	SkillBag IntKv //位置，技能id
	Lv       IntKv //合体id,等级
	Skills   map[int]*FitSkill
}

type Arena struct {
	DareNum     int //剩余挑战次数
	DareDate    int //挑战时间
	BuyDareNums int //购买挑战次数
	BuyDareNum  int //今日购买次数
}

type OnlineAward struct {
	Day         int   //日期
	OnlineTime  int   //今日在线时间（秒）
	GetAwardIds []int //已领取奖励Id
}

type DayStateRecord struct {
	Day               int   //日期
	RankWorship       int   //排行榜膜拜
	MonthCardReceive  IntKv //月卡每日礼包(monthCardType=>0)
	RechargeResetTime int   //充值重置日期
	DailyRecharge     int   //每日充值额度
}

type PanaceaUnit struct {
	Number  int //已使用次数
	Numbers int //总次数
}
type Panaceas map[int]*PanaceaUnit

type Pet struct {
	Lv    int
	Exp   int //经验
	Grade int //阶级
	Break int //突破
	Skill IntKv
}
type Pets map[int]*Pet

type UserWear struct {
	PetId        int //出战战宠id
	FitFashionId int //合体时装
}

type TalentUnit struct {
	UsePoints int   //使用天赋点
	Talents   IntKv //天赋,等级
}
type Talent struct {
	GetPoints     int                 //获得天赋点
	SurplusPoints int                 //剩余天赋点
	TalentList    map[int]*TalentUnit //类型，天赋详情
}

type Bag []*Item

type EquipBag map[int]*Equip

type Equips map[int]*Equip

type ItemSlice []*Item
type ItemSlice2 [][]*Item

type MainLineTask struct {
	TaskId      int
	Process     int
	MarkProcess int
}

type Jewel struct {
	One   int
	Two   int
	Three int
}
type Jewels map[int]*Jewel

type Fashion struct {
	Id int
	Lv int
}
type Fashions map[int]*Fashion

type Wear struct {
	FashionWeaponId int   //穿戴时装武器Id
	FashionClothId  int   //穿戴时装衣服Id
	WingId          int   //穿戴神翼
	AtlasWear       IntKv //图鉴穿戴id
	MagicCircleLvId int   //法阵等级表id
	TitleId         int   //称号ID
}

type Inside struct {
	Acupoint IntKv                //位置,内功Id
	Skill    map[int]*InsideSkill //技能
}
type InsideSkill struct {
	Level int
	Exp   int
}

type RingPhantom struct {
	Talent  int
	Phantom int   //类型
	Skill   IntKv //技能id,等级
}
type RingUnit struct {
	Rid        int                  //戒指id
	Strengthen int                  //强化等级
	Pid        int                  //强化Id(ringPhantom表)
	Talent     int                  //技能点
	Phantom    map[int]*RingPhantom //类型,幻灵
}
type Rings map[int]*RingUnit

type BagInfo map[int]*BagInfoUnit

type BagInfoUnit struct {
	MaxNum   int   `json:"m"`
	SpaceAdd IntKv `json:"s"`
	BuyNum   int   `json:"buyNum"`
}

type Sign struct {
	Count          int   //签到总数
	SignDay        IntKv //签到具体哪天
	ResetTime      int   //重置时间
	Cumulative     IntKv //累计签到奖励
	ContinuitySign int   //连续签到
}

type PersonBosses struct {
	DareNum   IntKv //StageId,次数
	ResetTime int
}

type Tower struct {
	TowerLv       int   //通关层数
	LotteryNum    int   //剩余抽奖次数
	DayAwardState int   //每日奖励是否领取，记录的当天日期
	LotteryId     int   //当前转盘奖励Id
	Lottery       []int //已经抽的奖
	RankAwardTime int   //排行榜奖励领取，记录周
}

type Shop struct {
	ResetTime int      //重置时间
	ShopItem  MapIntKv //类型,id,购买次数
}

type Display struct {
	ClothItemId     int
	ClothType       int
	WeaponItemId    int
	WeaponType      int
	WingId          int
	MagicCircleLvId int
	TitleId         int
	LabelId         int
	LabelJob        int
}

type MonthCardUnit struct {
	StartTime int
	EndTime   int
}
type MonthCard struct {
	ResetTime  int
	MonthCards map[int]*MonthCardUnit //品质，信息
}

type FirstRecharge struct {
	IsRecharge bool
	Days       IntKv
	OpenDay    int
	Discount   int //优惠itemId
}

type SpendRebates struct {
	CountIngot int   //总消耗元宝
	Ingot      int   //消费元宝(除了拍卖行相关)
	Reward     IntKv //已领取id
	Cycle      int   //周期
}

type SpecialEquipUnit struct {
	Id int
}
type Zodiacs map[int]*SpecialEquipUnit
type Kingarms map[int]*SpecialEquipUnit

type SkillUnit struct {
	Id        int
	Lv        int
	StartTime int64 //cd开始时间
	EndTime   int64 //cd结束时间
}
type Skills map[int]*SkillUnit

type Counts map[string][2]int //{"2":[3,14333232323],...} 次数
type ExData map[int]*json.RawMessage
type IntSlice []int
type IntSlice2 [][]int
type StringSlice []string
type IntKv map[int]int
type Int64Kv map[int]int64
type MapIntKv map[int]IntKv
type IntStringKv map[int]string
type Float64Slice []float64
type StringIntKv map[string]int

func (this IntSlice) Index(element int) int {
	for i, v := range this {
		if v == element {
			return i
		}
	}
	return -1
}

func (intMap IntKv) ToInt32() map[int32]int32 {
	r := make(map[int32]int32)
	for k, v := range intMap {
		//if v > 0 {
		r[int32(k)] = int32(v)
		//}
	}
	return r
}

func (intMap IntKv) ToInt64() map[int32]int64 {
	r := make(map[int32]int64)
	for k, v := range intMap {
		r[int32(k)] = int64(v)
	}
	return r
}

func BuildIntMap(intMap IntKv) map[int32]int32 {
	return intMap.ToInt32()
}

func BuildIntMap64(intMap IntKv) map[int32]int64 {
	return intMap.ToInt64()
}

func BuildIntMapIncludeZero(intMap IntKv) map[int32]int32 {
	r := make(map[int32]int32)
	for k, v := range intMap {
		r[int32(k)] = int32(v)
	}
	return r
}

// 玩家的门派数据
type GuildData struct {
	NowGuildId        int    `json:"NowGuildId"`        //当前加入的自建门派id
	MyCreateId        int    `json:"MyCreateId"`        //我自己创建的自建门派id
	Position          int    `json:"Position"`          //帮派位置
	ContributionValue int    `json:"ContributionValue"` //贡献值
	BeforeName        string `json:"BeforeName"`
	JoinCD            int    `json:"JoinCD"`        //下次可以加入门派的时间戳
	BeforeGuildId     int    `json:"BeforeGuildId"` // 上个门派id
	GuildCapital      int    `json:"guildCapital"`  //门派资金用于门派商店
}

type DailyTaskInfo struct {
	DayExp                 int                            //日活跃度
	WeekExp                int                            //周活跃度
	ResourcesBackExp       int                            //资源回收日活跃度
	DailyTask              map[int]*DailyTaskActivityInfo //key:activityTypeId
	GetDayRewardIds        []int                          //日活跃已度领取的奖励
	GetWeekRewardIds       []int                          //周活跃已度领取的奖励
	ResetTime              int                            //重置时间
	ResourcesHaveBackTimes IntKv                          //资源回收每日已找回次数
	ResourceCanBackTimes   StringIntKv                    //记录每日找回 前三天每天可找回的次数
}

type DailyTaskActivityInfo struct {
	ActivityId         int `json:"activityId"`
	IsCanGetExp        int `json:"isCanGetExp"`        //是否完成了挑战 可以领取活跃度
	HaveChallengeTimes int `json:"haveChallengeTimes"` //已经挑战了几次
	BuyChallengeTimes  int `json:"buyChallengeTimes"`  //购买了几次挑战机会
}

type Achievement struct {
	Point int                      //积分
	Task  map[int]*AchievementInfo //key:condition
	Medal []int                    //激活的勋章
}

type AchievementInfo struct {
	NowTaskId  int //value:当前可领取的id  -1的话代表全都领取完了
	NextTaskId int
	Process    int
	IsGetAll   int //这个类型的奖励都领取完了
}

type LimitGiftUnit struct {
	Lv        int //等级
	Grade     int //档次
	StartTime int
	EndTime   int
	IsBuy     bool
}
type LimitGift struct {
	GradeStatus IntKv                          //模块id,上次档次
	IsBuy       map[int]bool                   //模块id,上次是否购买
	TLv         IntKv                          //模块id,等级
	List        map[int]map[int]*LimitGiftUnit //模块id，模块等级，信息
	MergeData   IntKv                          //合服未买礼包,只使用key => 模块id
}

type DailyPackUnit struct {
	BuyIds    IntKv //已购买id,次数
	ResetTime int   //重置时间
	ResetWeek int
}
type DailyPack map[int]*DailyPackUnit

type GrowFund struct {
	IsBuy bool
	Ids   IntKv
}

type RedPacketItem struct {
	Day      int
	PickNum  int
	PickInfo IntKv //红包类型->次数
}

type WarOrderTask struct {
	Val    WarOrderTaskUnit //任务完成数
	Finish bool             //是否完成
	Reward bool             //是否已领取奖励
	Date   IntKv            //记录任务时间，比如泡点一天中进多少次都算1
}
type WarOrderReward struct {
	Elite  bool //精英
	Luxury bool //豪华
}
type WarOrderTaskUnit struct {
	One   int   //只有1个值
	Two   IntKv //2个值(warOrderCondition表,key为类型,val为数量)
	Three IntKv //2个值(warOrderCondition表,key为道具id,val为数量)
}
type WarOrder struct {
	Lv        int                           //等级
	Exp       int                           //经验
	Season    int                           //赛季
	StartTime int                           //开始时间
	EndTIme   int                           //结束时间
	IsLuxury  bool                          //是否豪华
	Exchange  IntKv                         //id,数量
	Task      map[int]*WarOrderTask         //任务ID,信息
	WeekTask  map[int]map[int]*WarOrderTask //星期,任务ID,信息
	Reward    map[int]*WarOrderReward       //等级,奖励
}

type Elf struct {
	Lv           int
	Exp          int
	Skills       IntKv //k是技能id，v是技能等级
	SkillBag     IntKv //位置，技能id
	RecoverLimit IntKv //每日回收上限
}

type CardInfo struct {
	AddWeight   int      //高级奖池增加权重 随到一次高级重置为0
	DrawTimes   int      //总共抽了几次
	Season      int      //赛季
	Integral    int      //积分
	GetAwardIds IntSlice //已领取的积分奖励
	DayResDay   int
	MergeMark   int //是否需要使用创建服务的开服天数 1:需要  2:赛季重置后 取主服务器openDay
}

type TreasureInfo struct {
	Season          int
	PopUpState      int
	PopUpResOpenDay int
	AllUseTimes     int //一共抽了多少次
	AllGetRound     IntSlice
	BuyTimes        IntKv // key:5元宝  >5对应的充值金额
	ChooseItems     map[int]IntSlice
	HaveRandomItems map[int]IntSlice
	MergeMark       int //是否需要使用创建服务的开服天数 1:需要  2:赛季重置后 取主服务器openDay
}

type HolyBeastInfos map[int]*HolyBeastInfo //key:圣兽类型

type HolyBeastInfo struct {
	Types      int   //圣兽类型
	Star       int   //星数
	ChooseProp IntKv //自己选择的技能 k:星数 v:选择对应技能的下标
}

type CompetitiveInfo struct {
	HaveChallengeTimes      int //今天已经挑战的次数
	BuyTimes                int //今天已购买挑战次数
	DayResDay               int
	BeforeDayRewardGetState int //每日奖励领取状态
	NowSeason               int //当前赛季
	ContinuityWin           int //连续赢场数
}
type MoveInfo struct {
	ItemId      int
	Count       int
	EquipIndex  int
	BeforeCount int
	AllCount    int
	CountLimit  int
	IsNewPos    bool
}

type FieldFight struct {
	HaveChallengeTimes int //今天已经挑战的次数
	HaveBuyTimes       int //今天已购买挑战次数
	DayResDay          int
}

type FieldUserCombatSlice []*FieldUserCombat

func (this FieldUserCombatSlice) Len() int {
	return len(this)
}

func (this FieldUserCombatSlice) Less(i, j int) bool {
	if this[i].Combat > this[j].Combat {
		return true
	}
	return false
}

func (this FieldUserCombatSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type FieldUserCombat struct {
	UserId   int
	Combat   int
	Avatar   string
	NickName string
	UserLv   int32
	Job      int32
	Sex      int32
}

type DailyRankInfos map[int]*DailyRankInfo //每日排行领奖信息

type DailyRankInfo struct {
	GetDayRewardIds IntKv //已领取每日积分奖励
	BuyRewardInfo   IntKv //已购买的充值奖励
}

type FitHolyEquip struct {
	SuitId int      //套装效果id
	Equips MapIntKv //类型,位置,id
}

type SevenInvestmentInfo struct {
	BuyOpenDay  int      //哪天购买的
	GetAwardIds IntSlice //已领取的积分奖励
}

type ContRecharge struct {
	Cycle   int   //周期
	Day     IntKv //日期,充值金额
	Receive IntKv //已领取id,0
}

type TitleUnit struct {
	StartTime int
	EndTime   int
	IsLook    bool
	IsExpire  bool
}
type Title map[int]*TitleUnit

type Condition map[int][]int

type MiJiUnit struct {
	MiJiLv   int //秘籍等级
	MiJiType int //秘籍类型
}

// 秘籍Type
type MiJi map[int]*MiJiUnit

type KillMonsterUni struct {
	Draw      bool //本服奖励领取
	FirstDraw bool //本服首杀奖励领取
}
type KillMonsterMil struct {
	Level int
	Draw  bool
}
type KillMonster struct {
	Uni map[int]*KillMonsterUni //本服首杀
	Per int                     //个人首通(领到了哪一个)
	Mil map[int]*KillMonsterMil //里程碑
}

// 远古宝物  key:远古宝物id
type AncientTreasuresInfo map[int]*AncientTreasures

type AncientTreasures struct {
	ZhuLinLv int //注灵
	Star     int //升星
	JueXinLv int //是否激活觉醒
	Types    int //类型
}

type TreasureShop struct {
	RefreshFree bool  //刷新是否免费
	RefreshTime int   //下次刷新时间
	BuyNum      int   //已购买次数
	Shop        IntKv //商品
	Car         IntKv //购物车
}

type HellBoss struct {
	DareNum   int
	BuyNum    int
	ResetTime int
	HelpNum   int
}

type LotteryInfo struct {
	ResetDay      int //重置时间
	PopUpState    int //结算弹窗状态
	GoodLuckState int //接好运奖励状态
	IsGetAward    int //主动领取奖励状态
}

// 远古宝物  key:远古宝物id
type TrialTaskInfos map[int]*TrialTaskInfo

type TrialTaskInfo struct {
	MarkNum    int //活动结束是完成的数量
	IsGetAward int
}

type DaBaoMystery struct {
	Energy     int //体力
	ResumeTime int
}

type Applets struct {
	Energy     int   //体力
	ResumeTime int64 //下次恢复时间
	List       map[int]*AppletsUnit
}

type AppletsUnit struct {
	Stage            int   //通过关卡
	LastGetAwardTime int64 //上一次领奖时间
	IsInGame         int   //是否在游戏内
}

type GuildRobotUserInfoSlice []*GuildRobotUserInfo

func (this GuildRobotUserInfoSlice) Len() int {
	return len(this)
}

func (this GuildRobotUserInfoSlice) Less(i, j int) bool {
	if this[i].Lv > this[j].Lv {
		return true
	}
	return false
}

func (this GuildRobotUserInfoSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type GuildRobotUserInfo struct {
	UserId int
	Lv     int
}

type Label struct {
	Id            int
	Job           int
	Transfer      int   //转职次数
	RefTime       int   //刷新时间
	DayReward     bool  //每日奖励是否领取
	FirstTransfer bool  //首次转职
	TaskOver      IntKv //已完成任务id
}

type ModuleUpMaxLv struct {
	BaoSiLv int //宝石镶嵌的最大等级
}
