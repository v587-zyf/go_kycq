package ancientBoss

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
	"fmt"
	"strconv"
	"strings"
	"time"
)

/**
 *  @Description: 远古首领进入战斗
 *  @param user
 *  @param stageId
 *  @return error
 */
func (this *AncientBoss) EnterAncientBossFight(user *objs.User, stageId int) error {
	if stageId < 1 {
		return gamedb.ERRPARAM
	}
	if this.getSurplusNum(user) <= 0 {
		return gamedb.ERRNOTENOUGHTIMES
	}

	bossInfo := gamedb.GetAncientBossAncientBossCfg(stageId)
	if bossInfo == nil {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("ancientBoss stageId:%v", stageId)
	}
	if check := this.GetCondition().CheckMulti(user, -1, bossInfo.Condition); !check {
		return gamedb.ERRCONDITION
	}

	err := this.GetFight().EnterResidentFightByStageId(user, stageId, 0)
	if err != nil {
		return err
	}

	user.AncientBoss.DareNum++
	this.GetCondition().RecordCondition(user, pb.CONDITION_TIAO_ZHAN_YUAN_GU_BOSS, []int{1})
	user.Dirty = true
	return nil
}

/**
 *  @Description: 远古首领战斗结算
 *  @param user
 *  @param winUserId	归属者
 *  @param stageId
 *  @param items		拾取奖励
 *  @return error
 */
func (this *AncientBoss) AncientBossFightResult(user *objs.User, winUserId, stageId int, items map[int]int) {
	ntf := &pb.AncientBossFightResultNtf{
		StageId: int32(stageId),
		Result:  pb.RESULTFLAG_FAIL,
	}
	if winUserId != user.Id {
		op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAncientFight)
		stageCfg := gamedb.GetAncientBossAncientBossCfg(stageId)
		this.GetBag().AddItems(user, stageCfg.JoinDrop, op)
		ntf.Goods = op.ToChangeItems()
		this.GetUserManager().SendItemChangeNtf(user, op)
	} else {
		this.WriteOwner(winUserId, stageId)

		ntf.Result = pb.RESULTFLAG_SUCCESS
		ntf.Goods = ophelper.CreateGoodsChangeNtf(items)
		this.GetFirstDrop().CheckIsFirstDrop(user, items)
		//this.GetAnnouncement().FightSendSystemChat(user, items, stageId, pb.SCROLINGTYPE_ANCIENT_BOSS)
	}
	ntf.Winner = this.GetUserManager().BuilderBrieUserInfo(winUserId)
	ntf.DareNum = int32(user.AncientBoss.DareNum)
	this.GetUserManager().SendMessage(user, ntf, true)
}

/**
 *  @Description: 获取最近的归属者列表
 *  @param user
 *  @param stageId
 *  @param ack
 */
func (this *AncientBoss) GetOwnerList(user *objs.User, stageId int) []*pb.AncientBossOwnerInfo {
	rmodelBoss := rmodel.Boss
	ownerData := rmodelBoss.GetBossOwner(stageId)
	pbSlice := make([]*pb.AncientBossOwnerInfo, 0)
	if len([]rune(ownerData)) > 0 {
		ownerDataArr := strings.Split(ownerData, "{-}")
		for _, ownerDataUnit := range ownerDataArr {
			info := strings.Split(ownerDataUnit, "[-]")
			userId, _ := strconv.Atoi(info[0])
			times, _ := strconv.Atoi(info[1])

			pbSlice = append(pbSlice, &pb.AncientBossOwnerInfo{
				Name: this.GetUserManager().GetUserBasicInfo(userId).NickName,
				Time: int64(times),
			})
		}
	}
	return pbSlice
}

func (this *AncientBoss) WriteOwner(winUserId, stageId int) {
	rmodelBoss := rmodel.Boss
	ownerStr := fmt.Sprintf(OWNER_STR, winUserId, int(time.Now().Unix()))
	ownerData := rmodelBoss.GetBossOwner(stageId)
	if len([]rune(ownerData)) > 0 {
		if ownerArr := strings.Split(ownerData, "{-}"); len(ownerArr) >= 5 {
			ownerArr = ownerArr[1:]
			ownerData = ""
			for _, str := range ownerArr {
				ownerData += str + "{-}"
			}
			ownerData = ownerData[:len(ownerData)-3]
		}
		ownerStr = fmt.Sprintf(OWNER_STR_TOO, ownerData, ownerStr)
	}
	rmodelBoss.SetBossOwner(stageId, ownerStr)
}
