package killMonster

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

/**
 *  @Description: 领取本服首杀全服奖励
 *  @param user
 *  @param stageId
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *KillMonster) DrawUni(user *objs.User, stageId int, op *ophelper.OpBagHelperDefault) error {
	this.Mu.Lock()
	defer this.Mu.Unlock()
	uniCfg := gamedb.GetKillMonsterUniByStageId(stageId)
	if uniCfg == nil {
		return gamedb.ERRPARAM
	}
	userKillMonster := user.KillMonster.Uni
	uni, ok := userKillMonster[stageId]
	if !ok {
		userKillMonster[stageId] = &model.KillMonsterUni{}
		uni = userKillMonster[stageId]
	}
	if this.KillMonsterData[stageId] == nil {
		return gamedb.ERRPARAM
	}
	if uni.Draw {
		return gamedb.ERRREPEATRECEIVE
	}

	this.GetBag().AddItems(user, uniCfg.RewardUni, op)
	uni.Draw = true
	user.Dirty = true
	return nil
}

/**
 *  @Description: 领取本服首杀归属奖励
 *  @param user
 *  @param stageId
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *KillMonster) DrawUniFirst(user *objs.User, stageId int, op *ophelper.OpBagHelperDefault) error {
	this.Mu.Lock()
	defer this.Mu.Unlock()
	uniCfg := gamedb.GetKillMonsterUniByStageId(stageId)
	if uniCfg == nil {
		return gamedb.ERRPARAM
	}
	userKillMonster := user.KillMonster.Uni
	uni, ok := userKillMonster[stageId]
	if !ok {
		userKillMonster[stageId] = &model.KillMonsterUni{}
		uni = userKillMonster[stageId]
	}
	db, ok1 := this.KillMonsterData[stageId]
	if !ok1 || db.FirstKillUserId != user.Id {
		return gamedb.ERRPARAM
	}
	if uni.FirstDraw {
		return gamedb.ERRREPEATRECEIVE
	}

	this.GetBag().AddItems(user, uniCfg.RewardFirst, op)
	uni.FirstDraw = true
	user.Dirty = true
	return nil
}

/**
 *  @Description: 领取个人首通
 *  @param user
 *  @param stageId
 *  @param op
 *  @return error
 */
func (this *KillMonster) DrawPer(user *objs.User, stageId int, op *ophelper.OpBagHelperDefault) error {
	perCfg := gamedb.GetKillMonsterPerByStageId(stageId)
	if perCfg == nil || user.StageId2 < stageId {
		return gamedb.ERRPARAM
	}
	if user.KillMonster.Per >= stageId {
		return gamedb.ERRPARAM
	}

	this.GetBag().AddItems(user, perCfg.RewardFirst, op)
	user.KillMonster.Per = stageId
	user.Dirty = true
	return nil
}

/**
 *  @Description: 领取里程碑
 *  @param user
 *  @param t
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *KillMonster) DrawMil(user *objs.User, t int, op *ophelper.OpBagHelperDefault, ack *pb.KillMonsterMilDrawAck) error {
	this.Mu.Lock()
	defer this.Mu.Unlock()
	userKillMonster := user.KillMonster.Mil
	mil, ok := userKillMonster[t]
	if !ok {
		userKillMonster[t] = &model.KillMonsterMil{Level: 1}
		mil = userKillMonster[t]
	}
	if mil.Draw {
		return gamedb.ERRREPEATRECEIVE
	}
	milCfg := gamedb.GetKillMonsterMilByTypeAndLv(t, mil.Level)
	if milCfg == nil {
		return gamedb.ERRPARAM
	}
	db, ok1 := this.KillMonsterData[milCfg.Stageid]
	if !ok1 {
		return gamedb.ERRPARAM
	}
	if db.KillNumAll < milCfg.Num {
		return gamedb.ERRPARAM
	}

	this.GetBag().AddItems(user, milCfg.Reward, op)
	if mil.Level >= gamedb.GetMaxValById(t, constMax.MAX_KILL_MONSTER_MIL_LEVEL) {
		mil.Draw = true
	}
	mil.Level++
	user.Dirty = true

	ack.Type = int32(t)
	ack.Level = int32(userKillMonster[t].Level)
	ack.Goods = op.ToChangeItems()
	return nil
}
