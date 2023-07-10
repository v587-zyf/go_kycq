package managers

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/ptsdk"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/managers/PreviewFunction"
	"cqserver/gameserver/internal/managers/achievement"
	"cqserver/gameserver/internal/managers/ancientBoss"
	"cqserver/gameserver/internal/managers/ancientTreasures"
	"cqserver/gameserver/internal/managers/announcement"
	"cqserver/gameserver/internal/managers/applets"
	"cqserver/gameserver/internal/managers/area"
	"cqserver/gameserver/internal/managers/arena"
	"cqserver/gameserver/internal/managers/atlas"
	"cqserver/gameserver/internal/managers/auction"
	"cqserver/gameserver/internal/managers/bag"
	"cqserver/gameserver/internal/managers/cardActivity"
	"cqserver/gameserver/internal/managers/challenge"
	"cqserver/gameserver/internal/managers/chat"
	"cqserver/gameserver/internal/managers/chuanshi"
	"cqserver/gameserver/internal/managers/competitve"
	"cqserver/gameserver/internal/managers/compose"
	"cqserver/gameserver/internal/managers/condition"
	"cqserver/gameserver/internal/managers/dabao"
	"cqserver/gameserver/internal/managers/dailyActivity"
	"cqserver/gameserver/internal/managers/dailyPack"
	"cqserver/gameserver/internal/managers/dailyRank"
	"cqserver/gameserver/internal/managers/dailyTask"
	"cqserver/gameserver/internal/managers/darkPalace"
	"cqserver/gameserver/internal/managers/dictate"
	"cqserver/gameserver/internal/managers/dragonEquip"
	"cqserver/gameserver/internal/managers/elf"
	"cqserver/gameserver/internal/managers/equip"
	"cqserver/gameserver/internal/managers/expPool"
	"cqserver/gameserver/internal/managers/expStage"
	"cqserver/gameserver/internal/managers/fabao"
	"cqserver/gameserver/internal/managers/fashion"
	"cqserver/gameserver/internal/managers/fieldBoss"
	"cqserver/gameserver/internal/managers/fieldFight"
	"cqserver/gameserver/internal/managers/fight"
	"cqserver/gameserver/internal/managers/firstDrop"
	"cqserver/gameserver/internal/managers/firstRecharge"
	"cqserver/gameserver/internal/managers/fit"
	"cqserver/gameserver/internal/managers/friend"
	"cqserver/gameserver/internal/managers/gift"
	"cqserver/gameserver/internal/managers/gm"
	"cqserver/gameserver/internal/managers/godEquip"
	"cqserver/gameserver/internal/managers/growFund"
	"cqserver/gameserver/internal/managers/guardPillar"
	"cqserver/gameserver/internal/managers/guild"
	"cqserver/gameserver/internal/managers/guildBonfire"
	"cqserver/gameserver/internal/managers/hellBoss"
	"cqserver/gameserver/internal/managers/holyBeast"
	"cqserver/gameserver/internal/managers/holyarms"
	"cqserver/gameserver/internal/managers/idGenerator"
	"cqserver/gameserver/internal/managers/inside"
	"cqserver/gameserver/internal/managers/jewel"
	"cqserver/gameserver/internal/managers/juexue"
	"cqserver/gameserver/internal/managers/killMonster"
	"cqserver/gameserver/internal/managers/kingarms"
	"cqserver/gameserver/internal/managers/label"
	"cqserver/gameserver/internal/managers/lottery"
	"cqserver/gameserver/internal/managers/magicCircle"
	"cqserver/gameserver/internal/managers/magicTower"
	"cqserver/gameserver/internal/managers/mail"
	"cqserver/gameserver/internal/managers/materialStage"
	"cqserver/gameserver/internal/managers/miJi"
	"cqserver/gameserver/internal/managers/mining"
	"cqserver/gameserver/internal/managers/monthCard"
	"cqserver/gameserver/internal/managers/official"
	"cqserver/gameserver/internal/managers/offline"
	"cqserver/gameserver/internal/managers/online"
	"cqserver/gameserver/internal/managers/panacea"
	"cqserver/gameserver/internal/managers/paodian"
	"cqserver/gameserver/internal/managers/personBoss"
	"cqserver/gameserver/internal/managers/pet"
	"cqserver/gameserver/internal/managers/privilege"
	"cqserver/gameserver/internal/managers/rank"
	"cqserver/gameserver/internal/managers/recharge"
	"cqserver/gameserver/internal/managers/rein"
	"cqserver/gameserver/internal/managers/ring"
	"cqserver/gameserver/internal/managers/sevenInvestment"
	"cqserver/gameserver/internal/managers/shabake"
	"cqserver/gameserver/internal/managers/shabakeCross"
	"cqserver/gameserver/internal/managers/shop"
	"cqserver/gameserver/internal/managers/sign"
	"cqserver/gameserver/internal/managers/skill"
	"cqserver/gameserver/internal/managers/spendRebates"
	"cqserver/gameserver/internal/managers/stage"
	"cqserver/gameserver/internal/managers/system"
	"cqserver/gameserver/internal/managers/talent"
	"cqserver/gameserver/internal/managers/task"
	"cqserver/gameserver/internal/managers/title"
	"cqserver/gameserver/internal/managers/tower"
	"cqserver/gameserver/internal/managers/treasure"
	"cqserver/gameserver/internal/managers/treasureShop"
	"cqserver/gameserver/internal/managers/trialTask"
	"cqserver/gameserver/internal/managers/user"
	"cqserver/gameserver/internal/managers/vip"
	"cqserver/gameserver/internal/managers/vipBoss"
	"cqserver/gameserver/internal/managers/warOrder"
	"cqserver/gameserver/internal/managers/wing"
	"cqserver/gameserver/internal/managers/worldBoss"
	"cqserver/gameserver/internal/managers/worldLeader"
	"cqserver/gameserver/internal/managers/zodiac"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/tlog"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pbserver"
	"errors"
	"flag"
	"fmt"
	"runtime/debug"

	"cqserver/gamelibs/model"
	"cqserver/golibs/dbmodel"
	"cqserver/golibs/util"

	"cqserver/golibs/logger"
)

var (
	gameDbBasePath = flag.String("gamedb", "../../yulong/data/configs", "specify gamedb path")
)

var (
	serverDone = make(chan int)
)

type ModuleManager struct {
	*util.DefaultModuleManager
	serverInfo *modelCross.ServerInfo

	ServerStarted bool
	GameVesion    string
	IsTest        bool

	//用户
	UserManager   managersI.IUserManager
	VipManager    managersI.IVipManager
	Task          managersI.ITaskManager
	Bag           managersI.IBagManager
	Juexue        managersI.IJuexueManager
	Rein          managersI.IReinManager
	Atlas         managersI.IAtlasManager
	Shop          managersI.IShop
	Compose       managersI.IComposeManager
	Online        managersI.IOnline
	Panacea       managersI.IPanaceaManager
	Official      managersI.IOfficial
	Sign          managersI.ISignManager
	Inside        managersI.IInsideManager
	Pet           managersI.IPetManager
	Friend        managersI.IFriendManager
	Fit           managersI.IFitManager
	FirstRecharge managersI.IFirstRecharge
	SpendRebates  managersI.ISpendRebates
	Elf           managersI.IElfManager
	Title         managersI.ITitleManager
	ExpPool       managersI.IExpPoolManager
	Guild         managersI.IGuildManager
	Label         managersI.ILabelManager

	//武将
	Equip           managersI.IEquipManager
	Fabao           managersI.IFabaoManager
	GodEquip        managersI.IGodEquipManager
	Wing            managersI.IWingManager
	Zodiac          managersI.IZodiacManager
	Kingarms        managersI.IKingarmsManager
	Skill           managersI.ISkillManager
	Dictate         managersI.IDictateManager
	Jewel           managersI.IJewelManager
	Fashion         managersI.IFashion
	Holyarms        managersI.IHolyarmsManager
	Ring            managersI.IRingManager
	Area            managersI.IAreaManager
	DragonEquip     managersI.IDragonEquipManager
	MagicCircle     managersI.IMagicCircleManager
	Talent          managersI.ITalentManager
	ChuanShi        managersI.IChuanShiManager
	Offline         managersI.IOfflineManager
	Achievement     managersI.IAchievementManager
	HolyBeast       managersI.IHolyBeastManager
	IdGenerator     managersI.IIdGeneratorManager
	Announcement    managersI.IAnnouncementManager
	MiJi            managersI.IMiJi
	AncientTreasure managersI.IAncientTreasureManager
	Privilege       managersI.IPrivilege

	//战斗
	StageManager  managersI.IStageManger
	DailyActivity managersI.IDailyActivity
	PaoDian       managersI.IPaoDian
	AncientBoss   managersI.IAncientBossManager
	GuardPillar   managersI.IGuardPillar
	KillMonster   managersI.IKillMonsterManager
	HellBoss      managersI.IHellBoss
	DaBao         managersI.IDaBaoManager
	Fight         managersI.IFight
	Tower         managersI.ITower
	FieldBoss     managersI.IFieldBoss
	WorldBoss     managersI.IWorldBoss
	PersonBoss    managersI.IPersonBoss
	MaterialStage managersI.IMaterialStage
	VipBoss       managersI.IVipBossManager
	ExpStage      managersI.IExpStageManager
	Arena         managersI.IArenaManager
	Competitve    managersI.ICompetitiveManager
	FieldFight    managersI.IFieldFightManager
	DarkPalace    managersI.IDarkPalaceManager
	GuildBonfire  managersI.IGuildBonfireManager
	Shabake       managersI.IShabakeManager
	Mining        managersI.IMiningManager
	Challenge     managersI.IChallengeManager
	WorldLeader   managersI.IWorldLeaderManager
	ShaBaKeCross  managersI.IShaBaKeCrossManager

	//活动、排行榜
	DailyPack    managersI.IDailyPackManager
	MonthCard    managersI.IMonthCard
	GrowFund     managersI.IGrowFundManager
	WarOrder     managersI.IWarOrderManager
	TreasureShop managersI.ITreasureShop
	Auction      managersI.IAuctionManager
	DailyTask    managersI.IDailyTaskManager
	DailyRank    managersI.IDailyRankManager
	Gift         managersI.IGiftManager
	CardActivity managersI.ICardActivityManager
	Treasure     managersI.ITreasureManager
	FirstDrop    managersI.IFirstDropManager

	//系统
	Recharge        managersI.IRecharge
	Rank            managersI.IRank
	Chat            managersI.IChatManager
	Mail            managersI.IMail
	Condition       managersI.IConditionManager
	System          managersI.ISystemManager
	Gm              managersI.IGmManager
	PreviewFunction managersI.IPreviewFunctionManager
	SevenInvestment managersI.ISevenInvestmentManager
	MagicTower      managersI.IMagicTowerManager
	Lottery         managersI.ILotteryManager
	TrialTask       managersI.ITrialTaskManager
	Applets         managersI.IAppletsManager
	Tlog            *tlog.TLog
	Event           *EventManager
	Cron            *CronManager

	ClientManager *ClientManager
	CsManager     *CsManager
	FSManager     *FSManager
}

var m = &ModuleManager{
	DefaultModuleManager: util.NewDefaultModuleManager(),
}

func Get() *ModuleManager {
	return m
}

func (this *ModuleManager) Test() bool {
	return this.IsTest
}

func (this *ModuleManager) GetUserManager() managersI.IUserManager {
	return this.UserManager
}
func (this *ModuleManager) GetVipManager() managersI.IVipManager {
	return this.VipManager
}
func (this *ModuleManager) GetStageManager() managersI.IStageManger {
	return this.StageManager
}
func (this *ModuleManager) GetTask() managersI.ITaskManager {
	return this.Task
}
func (this *ModuleManager) GetBag() managersI.IBagManager {
	return this.Bag
}
func (this *ModuleManager) GetEquip() managersI.IEquipManager {
	return this.Equip
}
func (this *ModuleManager) GetFabao() managersI.IFabaoManager {
	return this.Fabao
}
func (this *ModuleManager) GetGodEquip() managersI.IGodEquipManager {
	return this.GodEquip
}
func (this *ModuleManager) GetJuexue() managersI.IJuexueManager {
	return this.Juexue
}
func (this *ModuleManager) GetWing() managersI.IWingManager {
	return this.Wing
}
func (this *ModuleManager) GetRein() managersI.IReinManager {
	return this.Rein
}
func (this *ModuleManager) GetAtlas() managersI.IAtlasManager {
	return this.Atlas
}
func (this *ModuleManager) GetShop() managersI.IShop {
	return this.Shop
}
func (this *ModuleManager) GetZodiac() managersI.IZodiacManager {
	return this.Zodiac
}
func (this *ModuleManager) GetKingarms() managersI.IKingarmsManager {
	return this.Kingarms
}
func (this *ModuleManager) GetSkill() managersI.ISkillManager {
	return this.Skill
}
func (this *ModuleManager) GetDragonEquip() managersI.IDragonEquipManager {
	return this.DragonEquip
}
func (this *ModuleManager) GetOnline() managersI.IOnline {
	return this.Online
}
func (this *ModuleManager) GetDictate() managersI.IDictateManager {
	return this.Dictate
}
func (this *ModuleManager) GetPanacea() managersI.IPanaceaManager {
	return this.Panacea
}
func (this *ModuleManager) GetJewel() managersI.IJewelManager {
	return this.Jewel
}
func (this *ModuleManager) GetFashion() managersI.IFashion {
	return this.Fashion
}
func (this *ModuleManager) GetOfficial() managersI.IOfficial {
	return this.Official
}
func (this *ModuleManager) GetSign() managersI.ISignManager {
	return this.Sign
}
func (this *ModuleManager) GetInside() managersI.IInsideManager {
	return this.Inside
}
func (this *ModuleManager) GetHolyarms() managersI.IHolyarmsManager {
	return this.Holyarms
}
func (this *ModuleManager) GetRing() managersI.IRingManager {
	return this.Ring
}
func (this *ModuleManager) GetMining() managersI.IMiningManager {
	return this.Mining
}
func (this *ModuleManager) GetPet() managersI.IPetManager {
	return this.Pet
}
func (this *ModuleManager) GetArea() managersI.IAreaManager {
	return this.Area
}
func (this *ModuleManager) GetMagicCircle() managersI.IMagicCircleManager {
	return this.MagicCircle
}
func (this *ModuleManager) GetTalent() managersI.ITalentManager {
	return this.Talent
}
func (this *ModuleManager) GetDailyActivity() managersI.IDailyActivity {
	return this.DailyActivity
}
func (this *ModuleManager) GetPaoDian() managersI.IPaoDian {
	return this.PaoDian
}
func (this *ModuleManager) GetFriend() managersI.IFriendManager {
	return this.Friend
}
func (this *ModuleManager) GetFit() managersI.IFitManager {
	return this.Fit
}
func (this *ModuleManager) GetFirstRecharge() managersI.IFirstRecharge {
	return this.FirstRecharge
}
func (this *ModuleManager) GetMonthCard() managersI.IMonthCard {
	return this.MonthCard
}
func (this *ModuleManager) GetSpendRebates() managersI.ISpendRebates {
	return this.SpendRebates
}
func (this *ModuleManager) GetDailyPack() managersI.IDailyPackManager {
	return this.DailyPack
}
func (this *ModuleManager) GetGrowFund() managersI.IGrowFundManager {
	return this.GrowFund
}
func (this *ModuleManager) GetWarOrder() managersI.IWarOrderManager {
	return this.WarOrder
}
func (this *ModuleManager) GetElf() managersI.IElfManager {
	return this.Elf
}
func (this *ModuleManager) GetChuanShi() managersI.IChuanShiManager {
	return this.ChuanShi
}
func (this *ModuleManager) GetAncientBoss() managersI.IAncientBossManager {
	return this.AncientBoss
}
func (this *ModuleManager) GetGuardPillar() managersI.IGuardPillar {
	return this.GuardPillar
}
func (this *ModuleManager) GetTitle() managersI.ITitleManager {
	return this.Title
}
func (this *ModuleManager) GetKillMonster() managersI.IKillMonsterManager {
	return this.KillMonster
}
func (this *ModuleManager) GetTreasureShop() managersI.ITreasureShop {
	return this.TreasureShop
}
func (this *ModuleManager) GetHellBoss() managersI.IHellBoss {
	return this.HellBoss
}
func (this *ModuleManager) GetDaBao() managersI.IDaBaoManager {
	return this.DaBao
}
func (this *ModuleManager) GetLabel() managersI.ILabelManager {
	return this.Label
}
func (this *ModuleManager) GetPrivilegeModule() managersI.IPrivilege {
	return this.Privilege
}

func (this *ModuleManager) GetHolyBeast() managersI.IHolyBeastManager {
	return this.HolyBeast
}

func (this *ModuleManager) GetFight() managersI.IFight {
	return this.Fight
}
func (this *ModuleManager) GetTower() managersI.ITower {
	return this.Tower
}
func (this *ModuleManager) GetFieldBoss() managersI.IFieldBoss {
	return this.FieldBoss
}
func (this *ModuleManager) GetWorldBoss() managersI.IWorldBoss {
	return this.WorldBoss
}
func (this *ModuleManager) GetPersonBoss() managersI.IPersonBoss {
	return this.PersonBoss
}
func (this *ModuleManager) GetMaterialStage() managersI.IMaterialStage {
	return this.MaterialStage
}
func (this *ModuleManager) GetVipBoss() managersI.IVipBossManager {
	return this.VipBoss
}
func (this *ModuleManager) GetExpStage() managersI.IExpStageManager {
	return this.ExpStage
}
func (this *ModuleManager) GetArena() managersI.IArenaManager {
	return this.Arena
}
func (this *ModuleManager) GetCompetitve() managersI.ICompetitiveManager {
	return this.Competitve
}
func (this *ModuleManager) GetFieldFight() managersI.IFieldFightManager {
	return this.FieldFight
}
func (this *ModuleManager) GetDarkPalace() managersI.IDarkPalaceManager {
	return this.DarkPalace
}

func (this *ModuleManager) GetExpPool() managersI.IExpPoolManager {
	return this.ExpPool
}
func (this *ModuleManager) GetGuild() managersI.IGuildManager {
	return this.Guild
}

func (this *ModuleManager) GetGuildBonfire() managersI.IGuildBonfireManager {
	return this.GuildBonfire
}

func (this *ModuleManager) GetShabake() managersI.IShabakeManager {
	return this.Shabake
}

func (this *ModuleManager) GetAuction() managersI.IAuctionManager {
	return this.Auction
}

func (this *ModuleManager) GetDailyTask() managersI.IDailyTaskManager {
	return this.DailyTask
}

func (this *ModuleManager) GetDailyRank() managersI.IDailyRankManager {
	return this.DailyRank
}

func (this *ModuleManager) GetGift() managersI.IGiftManager {
	return this.Gift
}

func (this *ModuleManager) GetOffline() managersI.IOfflineManager {
	return this.Offline
}

func (this *ModuleManager) GetAchievement() managersI.IAchievementManager {
	return this.Achievement
}

func (this *ModuleManager) GetRecharge() managersI.IRecharge {
	return this.Recharge
}

func (this *ModuleManager) GetChallenge() managersI.IChallengeManager {
	return this.Challenge
}

func (this *ModuleManager) GetWorldLeader() managersI.IWorldLeaderManager {
	return this.WorldLeader
}

func (this *ModuleManager) GetShaBaKeCross() managersI.IShaBaKeCrossManager {
	return this.ShaBaKeCross
}

func (this *ModuleManager) GetCardActivity() managersI.ICardActivityManager {
	return this.CardActivity
}

func (this *ModuleManager) GetTreasure() managersI.ITreasureManager {
	return this.Treasure
}

func (this *ModuleManager) GetFirstDrop() managersI.IFirstDropManager {
	return this.FirstDrop
}

func (this *ModuleManager) Get() managersI.IWorldLeaderManager {
	return this.WorldLeader
}

func (this *ModuleManager) GetIdGenerator() managersI.IIdGeneratorManager {
	return this.IdGenerator
}

func (this *ModuleManager) GetRank() managersI.IRank {
	return this.Rank
}
func (this *ModuleManager) GetChat() managersI.IChatManager {
	return this.Chat
}
func (this *ModuleManager) GetMail() managersI.IMail {
	return this.Mail
}
func (this *ModuleManager) GetCondition() managersI.IConditionManager {
	return this.Condition
}
func (this *ModuleManager) GetSystem() managersI.ISystemManager {
	return this.System
}
func (this *ModuleManager) GetGm() managersI.IGmManager {
	return this.Gm
}
func (this *ModuleManager) GetTlog() *tlog.TLog {
	return this.Tlog
}

func (this *ModuleManager) GetPreviewFunction() managersI.IPreviewFunctionManager {
	return this.PreviewFunction
}

func (this *ModuleManager) GetAnnouncement() managersI.IAnnouncementManager {
	return this.Announcement
}

func (this *ModuleManager) GetMiJi() managersI.IMiJi {
	return this.MiJi
}

func (this *ModuleManager) GetAncientTreasure() managersI.IAncientTreasureManager {
	return this.AncientTreasure
}

func (this *ModuleManager) GetSevenInvestment() managersI.ISevenInvestmentManager {
	return this.SevenInvestment
}

func (this *ModuleManager) GetMagicTower() managersI.IMagicTowerManager {
	return this.MagicTower
}

func (this *ModuleManager) GetLottery() managersI.ILotteryManager {
	return this.Lottery
}

func (this *ModuleManager) GetTrialTask() managersI.ITrialTaskManager {
	return this.TrialTask
}

func (this *ModuleManager) GetApplets() managersI.IAppletsManager {
	return this.Applets
}

func (this *ModuleManager) DispatchEvent(userId int, data interface{}, callback func(userId int, user *objs.User, data interface{})) {
	this.ClientManager.DispatchEvent(userId, data, callback)
}

func (this *ModuleManager) BroadcastAll(msg nw.ProtoMessage) {
	this.ClientManager.BroadcastAll(msg)
}
func (this *ModuleManager) BroadcastData(sessionIds []uint32, data []byte) {
	this.ClientManager.BroadcastData(sessionIds, data)
}
func (this *ModuleManager) PutOutMessage(sessionId uint32, msg nw.ProtoMessage, sendNow bool) error {
	return this.ClientManager.PutOutMessage(sessionId, msg, sendNow)
}
func (this *ModuleManager) SyncUser(user *objs.User, status int) {
	lastRechargeTime := user.LastRechargeTime
	if user.LastRechargeTime <= 0 {
		t := modelGame.GetOrderModel().GetUserLastRechargeTime(user.Id)
		if t.IsZero() {
			lastRechargeTime = 0
		} else {
			lastRechargeTime = t.Unix()
		}
	}
	realRechargeTotal, tokenRechargeTotal := modelGame.GetOrderModel().GetUserRechargeTotal(user.Id)
	this.CsManager.SyncUser(user, status, lastRechargeTime, realRechargeTotal, tokenRechargeTotal)
}

func (this *ModuleManager) SendMsgToCCS(transId uint32, msg nw.ProtoMessage) {
	this.CsManager.SendMsgToCenterServer(transId, msg)
}

func (this *ModuleManager) CCsRpcCall(req nw.ProtoMessage, resp nw.ProtoMessage) {
	this.CsManager.RpcCall(req, resp)
}

// 战斗服 RpcCall
func (this *ModuleManager) FSRpcCall(fightId int, stageId int, req nw.ProtoMessage, resp nw.ProtoMessage) error {
	return this.FSManager.RpcCall(fightId, stageId, req, resp)
}

func (this *ModuleManager) RpcCallByFightServerId(crossFightServerId int, req nw.ProtoMessage, resp nw.ProtoMessage) error {
	return this.FSManager.RpcCallByFightServerId(crossFightServerId, req, resp)
}

func (this *ModuleManager) FSSendMessage(fightId, stageId int, msg nw.ProtoMessage) error {
	return this.FSManager.SendMessage(fightId, stageId, msg)
}

func (this *ModuleManager) GetCrossFightServerId(stageId int) int {
	return this.FSManager.GetCrossFightServerId(stageId)
}

func (this *ModuleManager) init() error {

	var err error
	err = gamedb.Load(*gameDbBasePath)
	if err != nil {
		return err
	}

	//先初始化accountdb
	if err = dbmodel.InitDbByKey(base.Conf, model.DB_ACCOUNT, false, false, false); err != nil {
		return err
	}
	//获取服务器的db配置
	serverInfo, err := modelCross.GetServerInfoModel().GetServerInfoByServerId(base.Conf.ServerId)
	if err != nil {
		return err
	}
	this.serverInfo = serverInfo
	dbLink, dbLogLink, err := base.GetBeanDb(serverInfo.DbLink, serverInfo.DbLinkLog)
	if err != nil {
		return err
	}
	base.Conf.DbConfigs[model.DB_SERVER] = dbLink
	base.Conf.DbConfigs[model.DB_LOG] = dbLogLink
	// 日志初数据库始化
	tlog.Init()
	//serverDb log初始化
	if err = dbmodel.InitDb(base.Conf, []string{model.DB_SERVER}, model.DB_SERVER, []string{model.DB_SERVER}); err != nil {
		return err
	}

	//本地redis优化
	rc, err := base.GetRedisConf(serverInfo.RedisAddr)
	if err != nil {
		return err
	}
	if err = rmodel.Init(rc, base.Conf.ServerId); err != nil {
		return err
	}

	ptsdk.InitSdk(base.Conf.Sdkconfig, base.Conf.Sandbox)
	return nil
}

func (this *ModuleManager) Init() (err error) {

	defer func() {
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			err = errors.New(fmt.Sprintf("%s", stackBytes))
		}
	}()

	err = this.init()

	if err != nil {
		return err
	}
	logger.Info("module init")
	this.UserManager = this.AppendModule(user.NewUserManager(this)).(managersI.IUserManager)
	this.VipManager = this.AppendModule(vip.NewVipManager(this)).(managersI.IVipManager)
	this.StageManager = this.AppendModule(stage.NewStageManager(this)).(managersI.IStageManger)
	this.Task = this.AppendModule(task.NewTaskManager(this)).(managersI.ITaskManager)
	this.Bag = this.AppendModule(bag.NewBagManager(this)).(managersI.IBagManager)
	this.Equip = this.AppendModule(equip.NewEquipManager(this)).(managersI.IEquipManager)
	this.PersonBoss = this.AppendModule(personBoss.NewPersonBoss(this)).(managersI.IPersonBoss)
	this.Fabao = this.AppendModule(fabao.NewFabaoManager(this)).(managersI.IFabaoManager)
	this.GodEquip = this.AppendModule(godEquip.NewGodEquipManager(this)).(managersI.IGodEquipManager)
	this.Juexue = this.AppendModule(juexue.NewJuexueManager(this)).(managersI.IJuexueManager)
	this.Wing = this.AppendModule(wing.NewWingManager(this)).(managersI.IWingManager)
	this.Rein = this.AppendModule(rein.NewReinManager(this)).(managersI.IReinManager)
	this.Atlas = this.AppendModule(atlas.NewAtlasManager(this)).(managersI.IAtlasManager)
	this.Shop = this.AppendModule(shop.NewShop(this)).(managersI.IShop)
	this.Arena = this.AppendModule(arena.NewArenaManager(this)).(managersI.IArenaManager)
	this.Competitve = this.AppendModule(competitve.NewCompetitveManager(this)).(managersI.ICompetitiveManager)
	this.Zodiac = this.AppendModule(zodiac.NewZodiacManager(this)).(managersI.IZodiacManager)
	this.Kingarms = this.AppendModule(kingarms.NewKingarmsManager(this)).(managersI.IKingarmsManager)
	this.Skill = this.AppendModule(skill.NewSkillManager(this)).(managersI.ISkillManager)
	this.Compose = this.AppendModule(compose.NewComposeManager(this)).(managersI.IComposeManager)
	this.Online = this.AppendModule(online.NewOnline(this)).(managersI.IOnline)
	this.Dictate = this.AppendModule(dictate.NewDictateManager(this)).(managersI.IDictateManager)
	this.Panacea = this.AppendModule(panacea.NewPanaceaManager(this)).(managersI.IPanaceaManager)
	this.Jewel = this.AppendModule(jewel.NewJewelManager(this)).(managersI.IJewelManager)
	this.Fashion = this.AppendModule(fashion.NewFashionManager(this)).(managersI.IFashion)
	this.Official = this.AppendModule(official.NewOfficialManager(this)).(managersI.IOfficial)
	this.Sign = this.AppendModule(sign.NewSignManager(this)).(managersI.ISignManager)
	this.Inside = this.AppendModule(inside.NewInsideManager(this)).(managersI.IInsideManager)
	this.Holyarms = this.AppendModule(holyarms.NewHolyarmsManager(this)).(managersI.IHolyarmsManager)
	this.Ring = this.AppendModule(ring.NewRingManager(this)).(managersI.IRingManager)
	this.Mining = this.AppendModule(mining.NewMiningManager(this)).(managersI.IMiningManager)
	this.Pet = this.AppendModule(pet.NewPetManager(this)).(managersI.IPetManager)
	this.Area = this.AppendModule(area.NewAreaManager(this)).(managersI.IAreaManager)
	this.FieldFight = this.AppendModule(fieldFight.NewFieldFightManager(this)).(managersI.IFieldFightManager)
	this.DragonEquip = this.AppendModule(dragonEquip.NewDragonEquipManager(this)).(managersI.IDragonEquipManager)
	this.ExpPool = this.AppendModule(expPool.NewExperiencePoolManager(this)).(managersI.IExpPoolManager)
	this.MagicCircle = this.AppendModule(magicCircle.NewMagicCircleManagerManager(this)).(managersI.IMagicCircleManager)
	this.Talent = this.AppendModule(talent.NewTalentManager(this)).(managersI.ITalentManager)
	this.Guild = this.AppendModule(guild.NewGuildManager(this)).(managersI.IGuildManager)
	this.GuildBonfire = this.AppendModule(guildBonfire.NewGuildBonfireManager(this)).(managersI.IGuildBonfireManager)
	this.DailyActivity = this.AppendModule(dailyActivity.NewDailyActivityManager(this)).(managersI.IDailyActivity)
	this.PaoDian = this.AppendModule(paodian.NewPaoDianManager(this)).(managersI.IPaoDian)
	this.Friend = this.AppendModule(friend.NewFriendManager(this)).(managersI.IFriendManager)
	this.Fit = this.AppendModule(fit.NewFitManager(this)).(managersI.IFitManager)
	this.FirstRecharge = this.AppendModule(firstRecharge.NewFirstRechargeManager(this)).(managersI.IFirstRecharge)
	this.MonthCard = this.AppendModule(monthCard.NewMonthCardManager(this)).(managersI.IMonthCard)
	this.SpendRebates = this.AppendModule(spendRebates.NewSpendRebatesManager(this)).(managersI.ISpendRebates)
	this.DailyPack = this.AppendModule(dailyPack.NewDailyPackManager(this)).(managersI.IDailyPackManager)
	this.GrowFund = this.AppendModule(growFund.NewGrowFundManager(this)).(managersI.IGrowFundManager)
	this.WarOrder = this.AppendModule(warOrder.NewWarOrderManager(this)).(managersI.IWarOrderManager)
	this.Elf = this.AppendModule(elf.NewElf(this)).(managersI.IElfManager)
	this.HolyBeast = this.AppendModule(holyBeast.NewHolyBeastManager(this)).(managersI.IHolyBeastManager)
	this.IdGenerator = this.AppendModule(idGenerator.NewIdGeneratorManager(this)).(managersI.IIdGeneratorManager)
	this.ChuanShi = this.AppendModule(chuanshi.NewChuanShiManager(this)).(managersI.IChuanShiManager)
	this.AncientBoss = this.AppendModule(ancientBoss.NewAncientBoss(this)).(managersI.IAncientBossManager)
	this.GuardPillar = this.AppendModule(guardPillar.NewGuardPillarManager(this)).(managersI.IGuardPillar)
	this.Title = this.AppendModule(title.NewTitleManager(this)).(managersI.ITitleManager)
	this.KillMonster = this.AppendModule(killMonster.NewKillMonsterManager(this)).(managersI.IKillMonsterManager)
	this.TreasureShop = this.AppendModule(treasureShop.NewTreasureShopManager(this)).(managersI.ITreasureShop)
	this.HellBoss = this.AppendModule(hellBoss.NewHellManager(this)).(managersI.IHellBoss)
	this.DaBao = this.AppendModule(dabao.NewDaBaoManager(this)).(managersI.IDaBaoManager)
	this.Label = this.AppendModule(label.NewLabelManager(this)).(managersI.ILabelManager)
	this.Privilege = this.AppendModule(privilege.NewPrivilegeManager(this)).(managersI.IPrivilege)

	this.Tower = this.AppendModule(tower.NewTower(this)).(managersI.ITower)
	this.Fight = this.AppendModule(fight.NewFight(this)).(managersI.IFight)
	this.FieldBoss = this.AppendModule(fieldBoss.NewFieldBoss(this)).(managersI.IFieldBoss)
	this.WorldBoss = this.AppendModule(worldBoss.NewWorldBoss(this)).(managersI.IWorldBoss)
	this.VipBoss = this.AppendModule(vipBoss.NewVipBossManager(this)).(managersI.IVipBossManager)
	this.MaterialStage = this.AppendModule(materialStage.NewMaterialStage(this)).(managersI.IMaterialStage)
	this.ExpStage = this.AppendModule(expStage.NewExpStageManager(this)).(managersI.IExpStageManager)
	this.DarkPalace = this.AppendModule(darkPalace.NewAreaManager(this)).(managersI.IDarkPalaceManager)
	this.Shabake = this.AppendModule(shabake.NewShabakeManager(this)).(managersI.IShabakeManager)
	this.Auction = this.AppendModule(auction.NewAuctionManager(this)).(managersI.IAuctionManager)
	this.DailyRank = this.AppendModule(dailyRank.NewDailyRankManager(this)).(managersI.IDailyRankManager)
	this.DailyTask = this.AppendModule(dailyTask.NewDailyTaskManager(this)).(managersI.IDailyTaskManager)
	this.Gift = this.AppendModule(gift.NewGiftManager(this)).(managersI.IGiftManager)
	this.Offline = this.AppendModule(offline.NewOfflineManager(this)).(managersI.IOfflineManager)
	this.Achievement = this.AppendModule(achievement.NewAchievementManager(this)).(managersI.IAchievementManager)
	this.Challenge = this.AppendModule(challenge.NewChallengeManager(this)).(managersI.IChallengeManager)
	this.WorldLeader = this.AppendModule(worldLeader.NewWorldLeaderBoss(this)).(managersI.IWorldLeaderManager)
	this.ShaBaKeCross = this.AppendModule(shabakeCross.NewCrossShaBaKeManager(this)).(managersI.IShaBaKeCrossManager)
	this.CardActivity = this.AppendModule(cardActivity.NewCardActivityManager(this)).(managersI.ICardActivityManager)
	this.Treasure = this.AppendModule(treasure.NewTreasure(this)).(managersI.ITreasureManager)
	this.Recharge = this.AppendModule(recharge.NewRechargeManager(this)).(managersI.IRecharge)
	this.Rank = this.AppendModule(rank.NewRankManager(this)).(managersI.IRank)
	this.Chat = this.AppendModule(chat.NewChatManager(this)).(managersI.IChatManager)
	this.Mail = this.AppendModule(mail.NewMailManager(this)).(managersI.IMail)
	this.Condition = this.AppendModule(condition.NewConditionManager(this)).(managersI.IConditionManager)
	this.System = this.AppendModule(system.NewSystemManager(this)).(managersI.ISystemManager)
	this.Gm = this.AppendModule(gm.NewGmManager(this)).(managersI.IGmManager)
	this.PreviewFunction = this.AppendModule(PreviewFunction.NewPreviewFunctionManager(this)).(managersI.IPreviewFunctionManager)
	this.SevenInvestment = this.AppendModule(sevenInvestment.NewSevenInvestmentManager(this)).(managersI.ISevenInvestmentManager)
	this.Announcement = this.AppendModule(announcement.NewAnnouncementManager(this)).(managersI.IAnnouncementManager)
	this.MiJi = this.AppendModule(miJi.NewMiJi(this)).(managersI.IMiJi)
	this.AncientTreasure = this.AppendModule(ancientTreasures.NewAncientTreasure(this)).(managersI.IAncientTreasureManager)
	this.MagicTower = this.AppendModule(magicTower.NewMagicTowerManager(this)).(managersI.IMagicTowerManager)
	this.Lottery = this.AppendModule(lottery.NewLotteryManager(this)).(managersI.ILotteryManager)
	this.TrialTask = this.AppendModule(trialTask.NewTrialTaskManager(this)).(managersI.ITrialTaskManager)
	this.Applets = this.AppendModule(applets.NewAppletsManager(this)).(managersI.IAppletsManager)
	this.FirstDrop = this.AppendModule(firstDrop.NewFirstDropManager(this)).(managersI.IFirstDropManager)
	this.Tlog = this.AppendModule(tlog.NewTLogManager()).(*tlog.TLog)

	this.Event = this.AppendModule(&EventManager{}).(*EventManager)
	this.Cron = this.AppendModule(NewCronManager()).(*CronManager)
	/********************************************************/
	/********************************************************/
	/*************以下代码不要上提，放最后初始化***************/
	/********************************************************/
	/********************************************************/
	this.ClientManager = this.AppendModule(NewClientManager(this.serverInfo.GsPort)).(*ClientManager)
	this.CsManager = this.AppendModule(NewCsManager()).(*CsManager)
	this.FSManager = this.AppendModule(NewFSManager(this.serverInfo.GsfsPort)).(*FSManager)
	logger.Info("DefaultModuleManager init")
	err = this.DefaultModuleManager.Init()
	if err != nil {
		return err
	}

	m.Fight.SyncResidentFightId()
	//试炼塔排行奖励
	m.Tower.SendRankReward()
	return err
}

func (this *ModuleManager) Start() error {
	//return this.DefaultModuleManager.Start()
	return nil
}

func (this *ModuleManager) Run() {
	this.DefaultModuleManager.Run()
	this.ServerStarted = true
}

func (this *ModuleManager) Stop() {
	close(serverDone)
	this.DefaultModuleManager.Stop()
}

// 修改了开服时间
func (this *ModuleManager) ModifyServerOpenTime() {
	this.DefaultModuleManager.ModifyServerOpenTime()
}

func GetM() *ModuleManager {
	return m
}

func (this *ModuleManager) ReloadGameDb() {
	//game重新加载配置
	gamedb.Reload()
	//游戏服重启加载配置
	request := &pbserver.GsToFsGamedbReloadReq{}
	replay := &pbserver.GsToFsGamedbReloadAck{}
	err := this.FSRpcCall(0, 0, request, replay)
	if err != nil {
		logger.Error("重加载表配置 战斗服异常,err:%v", err)
		return
	}
	logger.Error("重加载表配置完成")
}

// 初始化跨服redis连接
func (this *ModuleManager) InitCrossRedis() error {

	//confs, err := model.GetServerInfoModel().GetCrossRedisConfs()
	//if err != nil {
	//	return err
	//}
	//for _, conf := range confs {
	//	err = rmodel.InitCrossMap(conf.Id, conf.Network, conf.Address, conf.Password, conf.Db)
	//	if err != nil {
	//		logger.Info("ModuleManager.InitCrossRedis err=%v", err)
	//		return err
	//	}
	//}
	return nil
}
