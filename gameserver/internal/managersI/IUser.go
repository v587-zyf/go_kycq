package managersI

import (
	"cqserver/gamelibs/modelGame"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type IUserManager interface {
	GetUser(userId int) *objs.User
	/**
	 *  @Description: 踢玩家下线，玩家自己线程不能踢自己下线
	 *  @param user
	 *  @param reason
	 *  @return error
	 **/
	KickUserWithMsg(user *objs.User, reason string) error
	Save(user *objs.User, force bool) error
	UserDisconnect(user *objs.User)
	GetUserBasicInfo(userId int) *modelGame.UserBasicInfo
	// 随机名称
	RandName(user *objs.User, sex int, ack *pb.RandNameAck) error
	//加载玩家数据
	LoadUser(openId string, channelId int, clientIp string, serverId int, origin string, deviceId string) (*objs.User, error)
	//创建玩家角色
	CreateRole(user *objs.User, nickName string, avatar string, sex, job int) error
	//创建武将
	CreateHero(user *objs.User, sex, job int) (int, error)
	//推送玩家消息
	SendMessage(user *objs.User, msg nw.ProtoMessage, sendNow bool) error
	//推送玩家消息
	SendItemChangeNtf(user *objs.User, op *ophelper.OpBagHelperDefault) error
	//推送玩家消息
	SendMessageByUserId(userId int, msg nw.ProtoMessage) error

	/**
	 *  @Description: 		更新玩家武将战斗力
	 *  @param user
	 *  @param heroIndex    -1更新所有武将 用于更新共享模块
	 */
	UpdateCombat(user *objs.User, heroIndex int)
	UpdateCombatRobot(user *objs.User, heroIndex int)
	/**
	 *  @Description: 在线玩家定点任务
	 */
	TimingUpdate(t int)
	/**
	*	@Description:	角色外显
	 */
	SendDisplay(user *objs.User)
	/**
	 *  @Description: 获取武将外显
	 *  @param hero
	 *  @return *pb.Display
	 */
	GetHeroDisplay(hero *objs.Hero) *pb.Display

	/**
	 *  @Description: 获取玩家排行榜显示数据
	 *  @param userId
	 *  @param heroIndex
	 *  @param rank
	 *  @param score
	 *  @return *pb.RankInfo
	 */
	BuildUserRankInfo(userId int, heroIndex int, rank, score int) *pb.RankInfo

	/**
	 *  @Description: 获取玩家显示信息
	 *  @param userId
	 *  @return *pb.BriefUserInfo
	 */
	BuilderBrieUserInfo(userId int) *pb.BriefUserInfo

	BuilderAllUserInfoAndOffline(user *objs.User, rivalUserId int) *pb.BriefUserInfo
	//修改角色昵称
	ChangeHeroName(user *objs.User, heroIndex int, name string) error

	GetAllOnlineUserInfo() map[int]*objs.User

	GetOfflineUserInfo(userId int) *objs.User

	GetAllUserInfoIncludeOfflineUser(userId int) *objs.User

	GetUserOnlineStatus(userId int) (online bool, lastUpdateTime int64)
	RandGetUserId(getNum int, hasMap map[int]int) map[int]int

	GetAllUsersBasicInfo() map[int]*modelGame.UserBasicInfo

	CheckBan(user *objs.User, banType int) (bool, int)

	/**
	 *  @Description: 更新玩家封禁信息
	 *  @param openId
	 *  @param userId
	 **/
	UserBanUpdate(req *pbserver.BanInfoCcsToGsReq)

	GetOnlineTotal() int

	CrossFightOpen()

	/**
	 *  @Description: 订阅
	 *  @param user
	 *  @param subscribeId
	 **/
	Subscribe(user *objs.User, subscribeId int) error
    /**
    *  @Description: 获取玩家订阅记录
    *  @param userId
    *  @return []int32
    **/
	GetSubscribe(userId int) []int32

	/**
    *  @Description: 订阅定时任务
    **/
	CronSubscribe()
}
