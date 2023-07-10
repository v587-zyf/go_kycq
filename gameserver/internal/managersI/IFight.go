package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type IFight interface {

	/**
	 *  @Description: 进入战斗前检查
	 *  @param user
	 *  @param stageId
	 *  @return bool
	 **/
	CheckInFightBefore(user *objs.User, stageId int) bool
	/**
	 *  @Description: 进入战斗
	 *  @param user
	 *  @param fightId
	 *  @return error
	 */
	EnterFightByFightId(user *objs.User, stageId int, fightId int) error
	EnterFightByFightIdForUserRobot(userId int, fightId int, stageId int, teamId int) error
	NewHeroIntoFight(user *objs.User, heroIndex int)
	/**
	 *  @Description:  离开战斗
	 *  @param user
	 *  @param reason
	 *  @return error
	 */
	LeaveFight(user *objs.User, reason int) error
	/**
	 *  @Description: 创建战斗
	 *  @param stageId 战斗配置
	 *  @param cpData 关卡专属数据
	 *  @return int
	 *  @return error
	 */
	CreateFight(stageId int, cpData []byte) (int, error)

	/**
	 *  @Description:	更新玩家数据
	 *  @param user
	 *  @param heroIndex
	 */
	UpdateUserInfoToFight(user *objs.User, heroIndex map[int]bool, updateNow bool)

	/**
	 *  @Description: 更新玩家精灵数据
	 *  @param user
	 */
	UpdateUserElf(user *objs.User)

	/**
	 *  @Description:	更新玩家战斗模式
	 *  @param user
	 *  @param heroIndex
	 */
	UpdateUserFightModel(user *objs.User)
	/**
	 *  @Description:		获取野外boss信息
	 *  @param stageId		关卡Id
	 *  @return *pbserver.FsFieldBossInfoNtf
	 */
	GetFieldBossInfos(stageId int) *pbserver.FsFieldBossInfoNtf

	/**
	 *  @Description:		获取暗殿boss信息
	 *  @param stageId		关卡Id
	 *  @return *pbserver.FsFieldBossInfoNtf
	 */
	GetDarkPalaceBossInfos(stageId int) *pb.DarkPalaceBossNtf

	/**
	 *  @Description: 		获取远古首领信息
	 *  @param stageId		关卡Id
	 *  @return *pbserver.FsFieldBossInfoNtf
	 */
	GetAncientBossInfos(stageId int) *pbserver.FsFieldBossInfoNtf
	//获取炼狱首领信息
	GetHellBossInfos(stageId int) *pbserver.FsFieldBossInfoNtf
	/**
	 *  @Description:	通过stageId 进入常驻战斗
	 *  @param user
	 *  @param stageId
	 *  @return error
	 */
	EnterResidentFightByStageId(user *objs.User, stageId int, helpUserId int) error

	/**
	 *  @Description: 进入跨服沙巴克
	 *  @param user
	 *  @return error
	 */
	EnterShabakeCrossFight(user *objs.User) error

	/**
	 *  @Description: 进入公共副本（含主城）
	 *  @param stageId
	 */
	ClientEnterPublicCopy(user *objs.User, stageId int, condition int) (nw.ProtoMessage, error)

	/**
	 *  @Description:	战斗道具拾取
	 *  @param user
	 *  @param objsIds
	 *  @return *pb.FightPickUpAck
	 *  @return *ophelper.OpBagHelperDefault
	 */
	GsToFsPickUp(user *objs.User, objsIds []int32) (*pb.FightPickUpAck, *ophelper.OpBagHelperDefault,error)

	/**
	 *  @Description:	战斗玩家申请复活
	 *  @param user
	 *  @return *pb.FightUserReliveAck
	 *  @return *ophelper.OpBagHelperDefault
	 */
	GsToFsRelive(user *objs.User, SafeRelive bool) (*pb.FightUserReliveAck, *ophelper.OpBagHelperDefault, error)

	/**
	 *  @Description: 战斗道具使用
	 *  @param user
	 *  @param itemId
	 *  @return error
	 */
	UseItem(user *objs.User, itemId int) error

	/**
	 *  @Description: 鼓舞
	 *  @param user
	 *  @param op
	 */
	CheerReq(user *objs.User, op *ophelper.OpBagHelperDefault) (int, int, error)

	/**
	 *  @Description: 获取鼓舞使用次数
	 *  @param user
	 *  @return int
	 *  @return error
	 */
	CheerGetUseNum(user *objs.User) (int, int, error)

	/**
	 *  @Description: 使用药水
	 *  @param user
	 *  @param op
	 *  @return int
	 *  @return error
	 */
	UsePotion(user *objs.User, op *ophelper.OpBagHelperDefault) (int, error)

	/**
	 *  @Description: 使用药水CD
	 *  @param user
	 *  @return int
	 *  @return error
	 */
	UsePotionCdReq(user *objs.User) (int, error)

	/**
	 *  @Description: 申请采集
	 *  @param user
	 *  @param ack
	 */
	CollectionReq(user *objs.User, objId int, ack *pb.FightCollectionAck) error

	/**
	 *  @Description: 申请取消采集
	 *  @param user
	 *  @param objId
	 *  @param ack
	 *  @return error
	 */
	CollectionCancelReq(user *objs.User, objId int, ack *pb.FightCollectionCancelAck) error

	/**
	 *  @Description: 使用合体
	 *  @param user
	 */
	UserFitReq(user *objs.User) error

	/**
	 *  @Description: 取消合体
	 *  @param user
	 *  @return error
	 */
	UserFitCacelReq(user *objs.User) error

	/**
	 *  @Description: 战宠更新
	 *  @param user
	 *  @return error
	 */
	UserUpdatePet(user *objs.User) error

	/**
	 *  @Description: 使用切割打包
	 *  @param user
	 *  @return error
	 */
	UseCutTreasure(user *objs.User) error

	/**
	 *  @Description: gm命令
	 *  @param user
	 *  @param codes
	 *  @return string
	 */
	Gm(user *objs.User, codes string) string

	/**
	 *  @Description: 请求寻求帮助
	 *  @param user
	 *  @param HelpUserId
	 *  @return error
	 */
	ApplyForHelp(user *objs.User, HelpUserId int, source int) error

	/**
	 *  @Description: 回应帮助
	 *  @param user
	 *  @return error
	 */
	AskForHelpResult(user *objs.User, isAgree bool, reqHelpUserId, helpStageId int) error

	/**
	 *  @Description: 地图npc事件
	 *  @param user
	 *  @param npcId	地图npc标识Id
	 **/
	FightNpcEventReq(user *objs.User, npcId int) error

	/**
	 *  @Description: 更新玩家战斗次数
	 *  @param user
	 **/
	UpdateUserfightNum(user *objs.User)

	/**
	 *  @Description: 采集结束消息
	 *  @param overMsg
	 */
	CollectionOverNtf(overMsg *pbserver.FsToGsCollectionNtf)

	/**
	 *  @Description: 同步常驻战斗Id( 跨服的常驻战斗 )
	 **/
	SyncResidentFightId()

	/***************************************************************************/
	/***********************接收战斗服发送来消息处理逻辑接口**************************/
	/***************************************************************************/

	/*记录常驻战斗*/
	RecordResidentFight(fightInfos *pbserver.FsResidentFightNtf, fromCross bool)

	/**
	 *  @Description:		野外boss信息
	 *  @param fieldBossInfo
	 */
	HandlerFieldBossInfoNtf(fieldBossInfo *pbserver.FsFieldBossInfoNtf, isFromCross bool)

	/**
	 *  @Description:		野外boss  玩家死亡时间
	 *  @param fieldBossDieUserInfo
	 */
	HandlerFieldBossDieUserInfoNtf(fieldBossDieUserInfoNtf *pbserver.FsFieldBossDieUserInfoNtf)

	/**
	 *  @Description:	战斗服发来的添加道具
	 *  @param msg
	 *  @return nw.ProtoMessage
	 *  @return error
	 */
	FsToGsAddItem(msg *pbserver.FSAddItemReq) (nw.ProtoMessage, error)


	/**
	 *  @Description:		玩家技能使用
	 *  @param msg
	 */
	HandlerUserSkillUse(msg *pbserver.FsSkillUseNtf)

	/**
    *  @Description: 击杀消息
    *  @param msg
    **/
	ActorKillNtf(msg *pbserver.FsToGsActorKillNtf)

	/**
	 *  @Description:
	 *  @param userId
	 *  @param heroIndex
	 */
	ClearSkillCD(userId int, heroIndex int)

	/**
	 *  @Description: 战斗服通知game击杀第一个怪物
	 *  @param userId
	 */

	HandlerExpStageKillMonsterNtf(userId int)
	/**
	 *  @Description: 战斗结束
	 *  @param endMsg
	 */
	FightEnd(endMsg *pbserver.FSFightEndNtf, isCross bool)

	/**
	 *  @Description: 挂机地图击杀一波怪物
	 *  @param msg
	 */
	HangUpUserKillWave(msg *pbserver.HangUpKillWaveNtf)

	/**
	 *  @Description: 世界首例排行榜推送
	 *  @param endMsg
	 */
	CrossWorldLeaderRankNtf(endMsg *pbserver.WorldLeaderFightRankNtf)

	/**
	 *  @Description: 打宝秘境击杀怪物推送
	 *  @param msg
	 */
	DaBaoKillMonster(msg *pbserver.DaBaoKillMonsterNtf)

	/************************************泡点Pk*********************************************/

	/**
	 *  @Description: 进入泡点Pk
	 *  @param user
	 *  @param stageId
	 *  @return error
	 */
	EnterPaodian(user *objs.User, stageId int) error

	/**
	 *  @Description: 泡点增加奖励
	 *  @param endMsg
	 */
	PaodianGoodsAdd(endMsg *pbserver.PaodianGoodsAddNtf)

	/************************************公会相关战斗***********************************************/
	/**
	 *  @Description: 进入公会篝火
	 *  @param user
	 *  @param guildId
	 *  @return error
	 */
	EnterGuildBonfire(user *objs.User, guildId int) error
	/**
	 *  @Description:公会篝火增加玩家经验
	 *  @param endMsg
	 */
	GuildBonfireUserAddExp(endMsg *pbserver.GuildbonfireExpAddNtf)

	/**
	 *  @Description: 进入沙巴克战斗
	 *  @param user
	 *  @return error
	 */
	EnterShabakeFight(user *objs.User) error

	/**
	 *  @Description: 进入龙柱守卫
	 *  @param user
	 *  @return error
	 **/
	EnterGuardPillarFight(user *objs.User) error

	/**
	 *  @Description: 获取龙柱守卫结束时间
	 *  @param guildId
	 *  @return int
	 **/
	GetGuardPillarFightEndTime(guildId int) int

	/************************************九层魔塔相关战斗***********************************************/
	/**
	 *  @Description:
	 *  @param enterType
	 *  @return int
	 *  @return error
	 **/
	EnterMagicTowerByLayer(user *objs.User, enterType int) (int, error)

	/**
	 *  @Description: 领取九层魔塔玩家信息（是否领奖 当前积分）
	 *  @param user
	 *  @param op
	 **/
	MagicTowerGetUserInfo(user *objs.User) (int, int, error)

	/**
	 *  @Description: 领取九层魔塔层奖励
	 *  @param user
	 *  @param op
	 **/
	MagicTowerlayerAward(user *objs.User, op *ophelper.OpBagHelperDefault) error

	GetFightUserInfo(user *objs.User, teamId int, intoHeroIndex int, createPet bool, isCrossFight bool) *pbserver.User

	/**
	 *  @Description: 野外首领玩家死亡时间
	 *  @param guildId
	 *  @return int
	 **/
	GetFieldBossUserDieInfos(stageId int) int


	/**********************************bossfamily******************************/
    /**
    *  @Description: 获取bossfamily信息
    *  @param t
    *  @return map[int32]int32
    *  @return error
    **/
	GetBossFamilyInfo(t int) (map[int32]int32, error)

	/**
    *  @Description: 进入bossfamily
    *  @param user
	*  @param stageId
    *  @return error
    **/
	EnterBossFamily(user *objs.User,stageId int) error
}
