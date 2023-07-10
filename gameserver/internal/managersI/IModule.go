package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/tlog"
	"cqserver/golibs/nw"
)

type IModule interface {
	Test() bool
	GetUserManager() IUserManager
	GetVipManager() IVipManager
	GetStageManager() IStageManger
	GetTask() ITaskManager
	GetBag() IBagManager
	GetEquip() IEquipManager
	GetFabao() IFabaoManager
	GetGodEquip() IGodEquipManager
	GetJuexue() IJuexueManager
	GetWing() IWingManager
	GetRein() IReinManager
	GetAtlas() IAtlasManager
	GetShop() IShop
	GetZodiac() IZodiacManager
	GetKingarms() IKingarmsManager
	GetSkill() ISkillManager
	GetPanacea() IPanaceaManager
	GetFashion() IFashion
	GetOfficial() IOfficial
	GetSign() ISignManager
	GetMining() IMiningManager
	GetTalent() ITalentManager
	GetDailyActivity() IDailyActivity
	GetPaoDian() IPaoDian
	GetFriend() IFriendManager
	GetMonthCard() IMonthCard
	GetFirstRecharge() IFirstRecharge
	GetSpendRebates() ISpendRebates
	GetDailyPack() IDailyPackManager
	GetGrowFund() IGrowFundManager
	GetWarOrder() IWarOrderManager
	GetElf() IElfManager
	GetOffline() IOfflineManager
	GetAncientBoss() IAncientBossManager
	GetGuardPillar() IGuardPillar
	GetTreasureShop() ITreasureShop
	GetHellBoss() IHellBoss
	GetDaBao() IDaBaoManager
	GetLabel() ILabelManager
	GetPrivilegeModule() IPrivilege

	GetFight() IFight
	GetTower() ITower
	GetFieldBoss() IFieldBoss
	GetWorldBoss() IWorldBoss
	GetPersonBoss() IPersonBoss
	GetMaterialStage() IMaterialStage
	GetVipBoss() IVipBossManager
	GetExpStage() IExpStageManager
	GetArena() IArenaManager
	GetOnline() IOnline
	GetDarkPalace() IDarkPalaceManager
	GetGuildBonfire() IGuildBonfireManager
	GetShabake() IShabakeManager
	GetHolyBeast() IHolyBeastManager
	GetIdGenerator() IIdGeneratorManager
	GetPet() IPetManager
	GetTitle() ITitleManager
	GetKillMonster() IKillMonsterManager

	GetRecharge() IRecharge
	GetRank() IRank
	GetChat() IChatManager
	GetMail() IMail
	GetCondition() IConditionManager
	GetSystem() ISystemManager
	GetGm() IGmManager
	GetCompetitve() ICompetitiveManager
	GetFieldFight() IFieldFightManager
	GetExpPool() IExpPoolManager
	GetGuild() IGuildManager
	GetAuction() IAuctionManager
	GetDailyTask() IDailyTaskManager
	GetDailyRank() IDailyRankManager
	GetAchievement() IAchievementManager
	GetGift() IGiftManager
	GetChallenge() IChallengeManager
	GetWorldLeader() IWorldLeaderManager
	GetShaBaKeCross() IShaBaKeCrossManager
	GetCardActivity() ICardActivityManager
	GetTreasure() ITreasureManager
	GetPreviewFunction() IPreviewFunctionManager
	GetSevenInvestment() ISevenInvestmentManager
	GetAnnouncement() IAnnouncementManager
	GetMiJi() IMiJi
	GetAncientTreasure() IAncientTreasureManager
	GetMagicTower() IMagicTowerManager
	GetLottery() ILotteryManager
	GetTrialTask() ITrialTaskManager
	GetApplets() IAppletsManager
	GetFirstDrop() IFirstDropManager
	GetTlog() *tlog.TLog

	//GetEvent() *EventManager
	//GetCron() *CronManager
	// 
	DispatchEvent(userId int, data interface{}, callback func(userId int, user *objs.User, data interface{}))
	BroadcastAll(msg nw.ProtoMessage)
	BroadcastData(sessionIds []uint32, data []byte)
	PutOutMessage(sessionId uint32, msg nw.ProtoMessage, sendNow bool) error
	SyncUser(user *objs.User, status int)
	SendMsgToCCS(transId uint32, msg nw.ProtoMessage) //gs -> ccs
	CCsRpcCall(req nw.ProtoMessage, resp nw.ProtoMessage)

	//战斗服 RpcCall
	FSRpcCall(fightId int, stageId int, req nw.ProtoMessage, resp nw.ProtoMessage) error
	RpcCallByFightServerId(crossFightServerId int, req nw.ProtoMessage, resp nw.ProtoMessage) error
	//发送消息
	FSSendMessage(fightId int, stageId int, msg nw.ProtoMessage) error
	GetCrossFightServerId(stageId int) int
	//GetCsManager() *CsManager
	//GetFSManager() IFsManager
	ReloadGameDb()
}
