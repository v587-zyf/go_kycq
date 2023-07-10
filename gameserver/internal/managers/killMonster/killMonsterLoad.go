package killMonster

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

/**
 *  @Description: 加载本服首领首杀
 *  @param user
 *  @return error
 *  @return []*pb.KillMonsterUniInfo
 */
func (this *KillMonster) LoadUni(user *objs.User) (error, []*pb.KillMonsterUniInfo) {
	this.Mu.Lock()
	defer this.Mu.Unlock()
	if user.KillMonster.Uni == nil {
		user.KillMonster.Uni = make(map[int]*model.KillMonsterUni)
	}
	userKillMonster := user.KillMonster.Uni
	cfgs := gamedb.GetKillMonsterUniCfgs()
	pbSlice := make([]*pb.KillMonsterUniInfo, 0)
	for stageId := range cfgs {
		pbInfo := &pb.KillMonsterUniInfo{StageId: int32(stageId)}
		if db, ok := this.KillMonsterData[stageId]; ok {
			pbInfo.KlillUserId = int32(db.FirstKillUserId)
			pbInfo.KillUserName = this.GetUserManager().GetUserBasicInfo(db.FirstKillUserId).NickName
		}
		if uni, ok := userKillMonster[stageId]; ok {
			pbInfo.ServerFirstKill = uni.FirstDraw
			pbInfo.ServerKill = uni.Draw
		}
		pbSlice = append(pbSlice, pbInfo)
	}
	user.Dirty = true
	return nil, pbSlice
}

/**
 *  @Description: 加载个人首通
 *  @param user
 *  @return error
 *  @return map[int32]bool
 */
func (this *KillMonster) LoadPer(user *objs.User) (error, []*pb.KillMonsterPerInfo) {
	userKillMonster := user.StageId2
	userReceive := user.KillMonster.Per
	pbSlice := make([]*pb.KillMonsterPerInfo, 0)
	cfgs := gamedb.GetKillMonsterPerCfgs()
	for stageId := range cfgs {
		pbInfo := &pb.KillMonsterPerInfo{StageId: int32(stageId)}
		if userKillMonster >= stageId {
			pbInfo.Kill = true
		}
		if userReceive >= stageId {
			pbInfo.Receive = true
		}
		pbSlice = append(pbSlice, pbInfo)
	}
	user.Dirty = true
	return nil, pbSlice
}

/**
 *  @Description: 加载里程碑
 *  @param user
 *  @return error
 *  @return []*pb.KillMonsterMilInfo
 */
func (this *KillMonster) LoadMil(user *objs.User) (error, []*pb.KillMonsterMilInfo) {
	this.Mu.Lock()
	defer this.Mu.Unlock()
	if user.KillMonster.Mil == nil {
		user.KillMonster.Mil = make(map[int]*model.KillMonsterMil)
	}
	userKillMonster := user.KillMonster.Mil
	pbSlice := make([]*pb.KillMonsterMilInfo, 0)
	types := gamedb.GetKillMonsterType()
	for t := range types {
		pbInfo := &pb.KillMonsterMilInfo{Type: int32(t), Level: 1}
		if mil, ok := userKillMonster[t]; ok {
			pbInfo.Level = int32(mil.Level)
			pbInfo.Receive = mil.Draw
		}
		milCfg := gamedb.GetKillMonsterMilByTypeAndLv(t, 1)
		if db, ok := this.KillMonsterData[milCfg.Stageid]; ok {
			pbInfo.KillNum = int32(db.KillNumAll)
		}
		pbSlice = append(pbSlice, pbInfo)
	}
	return nil, pbSlice
}
