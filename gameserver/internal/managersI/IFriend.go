package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

type IFriendManager interface {
	Online(user *objs.User)
	/**
    *  @Description: 获取列表信息
    *  @param user
    *  @param block	是否是黑名单列表
    *  @return []*pb.FriendInfo
    */
	List(user *objs.User, block bool) []*pb.FriendInfo
	/**
    *  @Description: 添加好友
    *  @param user
    *  @param fId
    *  @return error
    */
	Add(user *objs.User, friendIdStr string, ack *pb.FriendAddAck) error
	/**
    *  @Description: 删除好友
    *  @param user
    *  @param fId
    *  @return error
    */
	Del(user *objs.User, fId int) error
	/**
    *  @Description: 添加黑名单
    *  @param user
    *  @param fId
    *  @return error
    */
	BlockAdd(user *objs.User, fId int) error
	/**
    *  @Description: 删除黑名单
    *  @param user
    *  @param fId
    *  @return error
    */
	BlockDel(user *objs.User, fId int) error
	/**
    *  @Description: 搜索用户
    *  @param user
    *  @param name	搜索名称（不填为推荐好友）
    *  @param ack
    *  @return error
    */
	Search(user *objs.User, name string, ack *pb.FriendSearchAck) error
	/**
    *  @Description: 私聊信息
    *  @param userId
    *  @param friendId
    *  @param msg
    */
	WriteMsgLog(user *objs.User, friendId int, msg string)
	ReadMsg(user *objs.User, friendId int) error
	/**
    *  @Description: 获取好友私聊信息
    *  @param user
    *  @param friendId	好友id
    *  @param ack
    *  @return error
    */
	FriendMsg(user *objs.User, friendId int, ack *pb.FriendMsgAck) error

	CheckFriendBlock(userId, friendId int) bool

	GetFriendNum(user *objs.User) int

	WriteFriendUserInfo(user *objs.User)
	GetFriendUserInfo(friendId int) *pb.FriendUserInfo
	/**
    *  @Description: 获取所有好友申请
    *  @param user
    *  @return []*pb.FriendApplyInfo
    */
	ApplyList(user *objs.User) []*pb.FriendApplyInfo
	ApplyAdd(user *objs.User, uId int) error
	ApplyAgree(user *objs.User, uId int) error
	ApplyRefuse(user *objs.User, uId int) error
}
