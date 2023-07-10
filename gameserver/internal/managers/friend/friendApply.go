package friend

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"fmt"
	"strconv"
	"time"
)

func (this *FriendManager) updateService() {
	for {
		select {
		case msg := <-this.updateChan:
			rmodel.Friend.AddFriendApply(msg.FriendId, msg.UserId)
			if user := this.GetUserManager().GetUser(msg.FriendId); user != nil {
				this.GetUserManager().SendMessage(user, &pb.FriendApplyAddNtf{FriendId: int32(msg.FriendId)}, true)
			}
		}
	}
}

func (this *FriendManager) ApplyAppend(friendId, userId int) {
	select {
	case this.updateChan <- FriendChange{UserId: userId, FriendId: friendId}:
	default:
		logger.Warn("friendManager: updateChan is full, please check .")
	}
}

/**
 *  @Description: 申请列表
 *  @param user
 *  @return []*pb.FriendApplyInfo
 */
func (this *FriendManager) ApplyList(user *objs.User) []*pb.FriendApplyInfo {
	pbSlice := make([]*pb.FriendApplyInfo, 0)
	applyIds := rmodel.Friend.GetFriendApply(user.Id)
	for _, friendId := range applyIds {
		pbSlice = append(pbSlice, this.BuildFriendApplyInfo(friendId))
	}
	return pbSlice
}

func (this *FriendManager) BuildFriendApplyInfo(friendId int) *pb.FriendApplyInfo {
	userBasicInfo := this.GetUserManager().GetUserBasicInfo(friendId)
	mainHero := userBasicInfo.HeroDisplay[constUser.USER_HERO_MAIN_INDEX]
	return &pb.FriendApplyInfo{
		UserId:   int32(userBasicInfo.Id),
		NickName: userBasicInfo.NickName,
		Lv:       int32(mainHero.ExpLvl),
		Avatar:   userBasicInfo.Avatar,
		Job:      int32(mainHero.Job),
		Sex:      int32(mainHero.Sex),
	}
}

func (this *FriendManager) checkFriendApply(friendId, userId int) bool {
	return rmodel.Friend.CheckFriendApply(friendId, userId)
}

func (this *FriendManager) delFriendApply(friendId, userId int) {
	rmodel.Friend.DelFriendApply(friendId, userId)
}

/**
 *  @Description: 好友申请
 *  @param user
 *  @param uId  要添加的好友id
 *  @return error
 */
func (this *FriendManager) ApplyAdd(user *objs.User, uId int) error {
	if uId < 1 {
		return gamedb.ERRPARAM
	}
	userId := user.Id
	if uId == userId {
		return gamedb.ERRPARAM
	}
	addFriendMaxNum := gamedb.GetConf().FriendsMaxNum
	if this.GetFriendNum(user) >= addFriendMaxNum {
		return gamedb.ERRFRIENDADDENOUGH
	}
	if friend, ok := user.Friend[uId]; ok {
		if friend.BlockTime == 0 && friend.DeletedAt == 0 {
			return gamedb.ERRFRIENDADD
		}
	}
	if this.checkFriendApply(uId, userId) {
		return gamedb.ERRREPEATAPPLY
	}
	this.ApplyAppend(uId, userId)
	return nil
}

/**
 *  @Description: 同意好友申请
 *  @param user
 *  @param uId
 *  @return error
 */
func (this *FriendManager) ApplyAgree(user *objs.User, uId int) error {
	if uId < 1 {
		return gamedb.ERRPARAM
	}
	userId := user.Id
	if uId == userId {
		return gamedb.ERRPARAM
	}
	if !this.checkFriendApply(userId, uId) {
		return gamedb.ERRPARAM
	}
	ack := &pb.FriendAddAck{}
	if err := this.Add(user, strconv.Itoa(uId), ack); err != nil && err != gamedb.ERRFRIENDADD {
		return err
	} else {
		this.GetUserManager().SendMessage(user, ack, true)
		this.GetUserManager().SendMessage(user, &pb.FriendApplyAgreeNtf{FriendId: int32(uId)}, true)
	}
	this.friendAddUser(uId, userId)
	this.delFriendApply(userId, uId)
	return nil
}

/**
 *  @Description: 拒绝好友申请
 *  @param user
 *  @param uId
 *  @return error
 */
func (this *FriendManager) ApplyRefuse(user *objs.User, uId int) error {
	if uId < 1 {
		return gamedb.ERRPARAM
	}
	userId := user.Id
	if uId == userId {
		return gamedb.ERRPARAM
	}
	if !this.checkFriendApply(userId, uId) {
		return gamedb.ERRPARAM
	}
	this.delFriendApply(userId, uId)
	this.GetUserManager().SendMessage(user, &pb.FriendApplyRefuseNtf{FriendId: int32(uId)}, true)
	return nil
}

func (this *FriendManager) friendAddUser(fid, userId int) {
	if friend := this.GetUserManager().GetUser(fid); friend != nil {
		ack := &pb.FriendAddAck{}
		if err := this.Add(friend, strconv.Itoa(userId), ack); err == nil {
			this.GetUserManager().SendMessage(friend, ack, true)
		}
	} else {
		applyStr := fmt.Sprintf("%d", userId)
		apply := rmodel.Friend.GetFriendApplyAdd(fid)
		if len([]rune(apply)) > 0 {
			applyStr = fmt.Sprintf(FRIEND_APPLY_TOO, applyStr, apply)
		}
		rmodel.Friend.SetFriendApplyAdd(fid, applyStr)
	}
}

func (this *FriendManager) friendDelUser(fid, userId int) {
	if friendInfo := this.GetUserManager().GetUser(fid); friendInfo != nil {
		if f, ok := friendInfo.Friend[userId]; ok && f.DeletedAt == 0 {
			f.DeletedAt = int(time.Now().Unix())
			this.GetUserManager().SendMessage(friendInfo, &pb.FriendDelAck{UserId: int64(userId)}, true)
		}
	} else {
		applyStr := fmt.Sprintf("%d", userId)
		apply := rmodel.Friend.GetFriendApplyDel(fid)
		if len([]rune(apply)) > 0 {
			applyStr = fmt.Sprintf(FRIEND_APPLY_TOO, applyStr, apply)
		}
		rmodel.Friend.SetFriendApplyDel(fid, applyStr)
	}
}