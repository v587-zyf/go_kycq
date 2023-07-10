package mining

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constMining"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"database/sql"
	"time"
)

const (
	MINMING_RSTATUS_OUT = 0
	MINMING_RSTATUS_IN  = 1
)

/**
 *  @Description: 进入矿洞
 */
func (this *MiningManager) In(user *objs.User) error {

	if !this.GetFight().CheckInFightBefore(user, constFight.FIGHT_TYPE_MINING_STAGE) {
		return gamedb.ERRUSERINFIGHT
	}

	fightId, err := this.GetFight().CreateFight(constFight.FIGHT_TYPE_MINING_STAGE, nil)
	if err != nil {
		return err
	}
	err = this.GetFight().EnterFightByFightId(user, constFight.FIGHT_TYPE_MINING_STAGE, fightId)
	if err != nil {
		return err
	}
	return nil
}

/**
 *  @Description: 掠夺进入战斗
 *  @param user
 *  @param robUid	被掠夺玩家id
 *  @return error
 */
func (this *MiningManager) Rob(user *objs.User, id int) error {
	miningInfo := this.GetMiningInfoById(id)
	if miningInfo == nil {
		return gamedb.ERRPARAM
	}
	if !miningInfo.Rtime.IsZero() {
		return gamedb.ERRMININGROB
	}
	fightUserId, _ := rmodel.Mining.GetMiningFight(id)
	if fightUserId > 0 {
		return gamedb.ERRUSERINFIGHT
	}
	if !miningInfo.ReceiveTime.IsZero() {
		return gamedb.ERRMININGOK
	}

	userMining := user.Mining
	maxRobNum := gamedb.GetConf().MiningRobMaxNum
	if userMining.RobNum >= maxRobNum {
		return gamedb.ERRNOTENOUGHTIMES
	}

	if user.FightId <= 0 {
		return gamedb.ERRFIGHTID
	}
	if user.FightStageId != constFight.FIGHT_TYPE_MINING_STAGE {
		return gamedb.ERRFIGHTTYPE
	}
	reqMsg := &pbserver.MiningNewFightInfoReq{
		MiningId:     int32(miningInfo.Id),
		MiningUserId: int32(miningInfo.UserId),
		IsRetake:     false,
	}
	replyMsg := &pbserver.MiningNewFightInfoAck{}
	err := this.FSRpcCall(user.FightId, user.FightStageId, reqMsg, replyMsg)
	if err != nil {
		return err
	}
	if !replyMsg.ReadyOk {
		return gamedb.ERRUSERINFIGHT
	}

	err = this.GetFight().EnterFightByFightIdForUserRobot(miningInfo.UserId, user.FightId, user.FightStageId, constFight.FIGHT_TEAM_ZERO)
	if err != nil {
		return err
	}

	userMining.RobNum++
	user.Dirty = true

	miningFightUserInfo := this.GetUserManager().GetUserBasicInfo(miningInfo.UserId)
	if miningFightUserInfo != nil {
		kyEvent.MiningFightStart(user, true, miningInfo.UserId, miningFightUserInfo.NickName, miningFightUserInfo.Combat)
	}
	rmodel.Rank.SetArenaRankFight(id, user.Id)

	this.MiningMu.Lock()
	defer this.MiningMu.Unlock()
	this.UpdateMiningDate(miningInfo, miningInfo.IsRobot == constMining.MINING_DATA_ROBOT_YES)
	return nil
}

/**
 *  @Description: 掠夺战斗回调
 *  @param user
 *  @param id 	  挖矿记录id(表id)
 *  @param isWin  战斗结果
 *  @param ack
 */
func (this *MiningManager) RobFightBack(user *objs.User, id int, isWin bool) {
	ack := &pb.MiningRobAck{
		Result: pb.RESULTFLAG_FAIL,
	}
	this.MiningMu.Lock()
	defer this.MiningMu.Unlock()
	miningInfo := this.GetMiningInfoById(id)
	rmodel.Mining.DelMiningFight(id)
	if user == nil {
		return
	}
	if isWin {
		miningInfo.Ruid = user.Id
		miningInfo.Rtime = time.Now()
		op := ophelper.NewOpBagHelperDefault(constBag.OpTypeMiningRobFight)
		lvCfg := gamedb.GetMiningLvCfg(miningInfo.Miner)
		items := make(gamedb.ItemInfos, 0)
		for itemId, count := range lvCfg.Lose {
			items = append(items, &gamedb.ItemInfo{
				ItemId: itemId,
				Count:  count,
			})
		}
		this.GetBag().AddItems(user, items, op)
		ack.Result = pb.RESULTFLAG_SUCCESS
		ack.Goods = op.ToChangeItems()
		this.GetUserManager().SendItemChangeNtf(user, op)
	}
	miningFightUserInfo := this.GetUserManager().GetUserBasicInfo(miningInfo.UserId)
	if miningFightUserInfo != nil {
		kyEvent.MiningFightEnd(user, true, int(ack.Result), miningInfo.UserId, miningFightUserInfo.NickName, miningFightUserInfo.Combat, miningInfo.Miner, user.Mining.RobNum)
	}
	ack.RobNum = int32(user.Mining.RobNum)
	if info, ok := this.MiningMap[miningInfo.UserId]; ok && info.Id == id {
		this.UpdateMiningDate(miningInfo, false)
	} else {
		if isWin && miningInfo.IsRobot == constMining.MINING_DATA_ROBOT_YES {
			timeNow := time.Now()
			miningInfo.ExpireTime = timeNow
			miningInfo.ReceiveTime = timeNow
			miningInfo.FindTime = timeNow
			miningInfo.DeletedAt = timeNow
			delete(this.RobotMap, miningInfo.UserId)
		}
		this.UpdateMiningDate(miningInfo, true)
	}
	this.GetUserManager().SendMessage(user, ack, true)
}

/**
 *  @Description: 夺回奖励
 *  @param user
 *  @param id	挖矿记录id(表id)
 *  @return error
 */
func (this *MiningManager) RobBack(user *objs.User, id int) error {
	miningInfo := this.GetMiningInfoById(id)
	if miningInfo == nil {
		return gamedb.ERRPARAM
	}
	if !miningInfo.FindTime.IsZero() {
		return gamedb.ERRMININGFIND
	}

	if user.FightId <= 0 {
		return gamedb.ERRFIGHTID
	}
	if user.FightStageId != constFight.FIGHT_TYPE_MINING_STAGE {
		return gamedb.ERRFIGHTTYPE
	}
	reqMsg := &pbserver.MiningNewFightInfoReq{
		MiningId:     int32(miningInfo.Id),
		MiningUserId: int32(miningInfo.Ruid),
		IsRetake:     true,
	}
	replyMsg := &pbserver.MiningNewFightInfoAck{}
	err := this.FSRpcCall(user.FightId, user.FightStageId, reqMsg, replyMsg)
	if err != nil {
		return err
	}
	if !replyMsg.ReadyOk {
		return gamedb.ERRUSERINFIGHT
	}

	err = this.GetFight().EnterFightByFightIdForUserRobot(miningInfo.Ruid, user.FightId, user.FightStageId, constFight.FIGHT_TEAM_ZERO)
	if err != nil {
		return err
	}
	miningFightUserInfo := this.GetUserManager().GetUserBasicInfo(miningInfo.Ruid)
	if miningFightUserInfo != nil {
		kyEvent.MiningFightStart(user, false, miningInfo.UserId, miningFightUserInfo.NickName, miningFightUserInfo.Combat)
	}
	return nil
}

/**
 *  @Description: 掠夺奖励战斗回调
 *  @param user
 *  @param id	  挖矿记录id(表id)
 *  @param isWin  战斗结果
 *  @param ack
 */
func (this *MiningManager) RobBackFightBack(user *objs.User, id int, isWin bool) {
	ack := &pb.MiningRobBackAck{}
	miningInfo := this.GetMiningInfoById(id)
	if user == nil {
		return
	}
	ack.Result = pb.RESULTFLAG_FAIL
	if isWin {
		ack.Result = pb.RESULTFLAG_SUCCESS
		op := ophelper.NewOpBagHelperDefault(constBag.OpTypeMiningRobBackFight)
		lvCfg := gamedb.GetMiningLvCfg(miningInfo.Miner)
		items := make(gamedb.ItemInfos, 0)
		for itemId, count := range lvCfg.Lose {
			items = append(items, &gamedb.ItemInfo{
				ItemId: itemId,
				Count:  count,
			})
		}
		this.GetBag().AddItems(user, items, op)
		miningInfo.FindTime = time.Now()
		this.GetUserManager().SendItemChangeNtf(user, op)
		ack.Goods = op.ToChangeItems()
	}
	miningFightUserInfo := this.GetUserManager().GetUserBasicInfo(miningInfo.Ruid)
	if miningFightUserInfo != nil {
		kyEvent.MiningFightEnd(user, false, int(ack.Result), miningInfo.UserId, miningFightUserInfo.NickName, miningFightUserInfo.Combat, miningInfo.Miner, user.Mining.RobNum)
	}
	this.MiningMu.Lock()
	if info, ok := this.MiningMap[miningInfo.UserId]; ok && info.Id == id {
		this.UpdateMiningDate(miningInfo, false)
	} else {
		this.UpdateMiningDate(miningInfo, true)
	}
	this.MiningMu.Unlock()
	this.GetUserManager().SendMessage(user, ack, true)
}

func (this *MiningManager) GetMiningInfoById(id int) *modelGame.MiningDb {
	miningInfo, err := modelGame.GetMiningModel().GetMiningById(id)
	if miningInfo == nil || err != nil && err != sql.ErrNoRows {
		logger.Error("GetMining error id:%v err:%v", id, err)
		return nil
	}
	return miningInfo
}
