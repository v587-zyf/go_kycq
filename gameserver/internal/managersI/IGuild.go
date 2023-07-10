package managersI

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

type IGuildManager interface {
	LoadGuild(user *objs.User, ack *pb.GuildLoadInfoAck) error

	CreateGuild(user *objs.User, op *ophelper.OpBagHelperDefault, req *pb.CreateGuildReq, ack *pb.CreateGuildAck) error

	//设置加入公会的最低战力
	SetJoinGuildCombatLimit(user *objs.User, ack *pb.JoinGuildCombatLimitAck, combat, isAgree int) error

	//加入公会
	ApplyJoinGuild(user *objs.User, guildId int, ack *pb.ApplyJoinGuildAck) error

	//退出公会
	QuitGuild(user *objs.User, ack *pb.QuitGuildAck) error

	//任命
	GuildAssign(user *objs.User, ack *pb.GuildAssignAck, assUserId, appPosition int) error

	//处理申请列表是否同意玩家加入门派
	JoinGuildDispose(user *objs.User, ack *pb.JoinGuildDisposeAck, isAgree bool, applyUserId int) error

	//申请列表玩家
	GetAllApplyUserLists(user *objs.User, ack *pb.GetApplyUserListAck) error

	GetAllGuildInfo(user *objs.User, ack *pb.AllGuildInfosAck) error

	//踢人
	KickOut(user *objs.User, ack *pb.KickOutAck, kickUserId int) error

	//解散公会
	DissolveGuild(user *objs.User, ack *pb.DissolveGuildAck) error

	//修改公告
	ModifyBulletin(user *objs.User, ack *pb.ModifyBulletinAck, notice string) error

	//弹劾会长
	ImpeachPresident(user *objs.User, ack *pb.ImpeachPresidentAck) error

	//门派成员职位 key:职位 1:会长 2:副会长 3:长老 4:成员
	GetGuildMemberInfo(guildId int) (error, map[int][]int, int, []int)

	/**
	 *  @Description: 获取门派信息
	 *  @param guildId
	 *  @return *modelGame.Guild
	 *  @return error
	 */
	GetGuildInfo(guildId int) *modelGame.Guild

	GetGuildName(userId int) string

	BroadcastChatToGuildUsers(user *objs.User, protoMsg *pb.ChatMessageNtf) error

	//获取指定门派会长and副会长 userIds    [userId,position,userId1,position1]
	GetGuildHuiAndFuHuiUserIds(guildId int) []int

	//一键处理申请列表
	AllJoinGuildDispose(user *objs.User, ack *pb.AllJoinGuildDisposeAck, isAgree bool) error

	OperationGuildCheck1(user *objs.User) error

	SetGuildInfo(guildInfo *modelGame.Guild)

	ResetGuildBonfireDonateInfo()

	DelGuildInfo(guildInfo *modelGame.Guild)

	SetGuildActivityInfo(guildId, activityId, status int)
	GetGuildActivityInfo(guildId, activityId int) bool
	GuildActivityLoad(user *objs.User, activityId int, ack *pb.GuildActivityLoadAck) error
	SendMsgToAllUser(guildId int, msg nw.ProtoMessage, notSendUserIds []int)
	CheckActiveOpen(user *objs.User, guildActivityId int) (error, *gamedb.GuildActivityGuildActivityCfg)
	CheckActivityOpenPower(position, guildActivityId int) bool

	DelRobotGuild()
}
