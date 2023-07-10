package fightModule

import (
	"cqserver/fightserver/conf"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"fmt"
	"time"
)

/**
*  @Description: 九层魔塔数据记录，不做任何逻辑
**/
type MagicTower struct {
	userScore        *userScores              //玩家累计积分
	userScoreMap     map[int]int              //玩家当前积分
	magicTowerfights map[int]*MagicTowerFight //stageId=>fight 每层战斗
	//locker            sync.Mutex
	locker            *util.Locker
	status            int                                       //开启状态
	startTime         int                                       //开始时间
	stopTime          int                                       //结束时间
	updateScoreByTime map[int]int                               //时间刷新玩家积分（战斗Id->是否更新） 所有战斗都更新后 推送所有战斗积分信息
	monsterInitTime   map[int]map[int]*MagicTowerMonsterRefresh //怪物刷新
}

func NewMagicTower() *MagicTower {

	f := &MagicTower{
		userScore:         NewUserScores(),
		userScoreMap:      make(map[int]int),
		updateScoreByTime: make(map[int]int),
		locker:            util.NewLocker(0),
		status:            0,
	}

	if conf.Conf.Sandbox {
		f.locker.SetDebug(true)
	}
	f.startTime, f.stopTime = f.GetNextActiveTime()

	return f
}

func (this *MagicTower) addRank(userId int, userName string) {

	this.locker.Lock()
	defer this.locker.Unlock()
	this.userScore.AddUser(userId, userName)
}

/**
*  @Description: 时间添加玩家积分
*  @receiver this
*  @param fightId
*  @param users
*  @param score
**/
func (this *MagicTower) addUseScore(fightId int, stageId int, users []int, score int) {

	hasUseAdd := false
	if len(users) > 0 && score > 0 {
		hasUseAdd = true
		for _, v := range users {
			this.addScore(stageId, v, score, false)
		}
	}

	this.locker.Lock()
	this.updateScoreByTime[fightId] = 0
	if hasUseAdd {
		this.updateScoreByTime[fightId] = len(users)
	}

	allUpdate := true
	userScoreSort := false
	for _, v := range this.magicTowerfights {
		if userScoreSortNum, ok := this.updateScoreByTime[int(v.id)]; !ok {
			allUpdate = false
			break
		} else {
			if userScoreSortNum > 0 {
				userScoreSort = true
			}
		}
	}
	if allUpdate {
		if userScoreSort {
			this.userScoreSort()
		}
		this.updateScoreByTime = make(map[int]int)
	}
	this.locker.Unlock()
}

/**
*  @Description: 增加单个玩家积分
*  @receiver this
*  @param userId
*  @param score
*  @param sortNow
**/
func (this *MagicTower) addScore(stageId int, userId int, score int, sortNow bool) {

	if this.status == 0 {
		return
	}
	this.locker.Lock()

	if score > 0 {
		this.userScore.addScore(userId, score)
		this.userScoreMap[userId] += score
		if sortNow {
			this.userScoreSort()
		}
		this.userScoreChangeNtf(stageId, userId, score)
	}
	this.locker.Unlock()
}

/**
*  @Description: 扣除玩家积分（进入下一层）
*  @receiver this
*  @param userId
*  @param lessNum
*  @return int
*  @return error
**/
func (this *MagicTower) LessUserScore(stageId int, userId, lessNum int) (int, error) {
	this.locker.Lock()
	defer this.locker.Unlock()
	if _, ok := this.userScoreMap[userId]; !ok {
		return 0, gamedb.ERRUNFOUNDUSER
	}
	score := this.userScoreMap[userId]
	if score < lessNum {
		return 0, gamedb.MARKNOTENOUGH
	}
	this.userScoreMap[userId] -= lessNum
	this.userScoreChangeNtf(stageId, userId, -lessNum)
	return this.userScoreMap[userId], nil
}

func (this *MagicTower) GetUserScore(userId int) int {
	this.locker.Lock()
	defer this.locker.Unlock()
	if score, ok := this.userScoreMap[userId]; ok {
		return score
	}
	return 0
}

/**
*  @Description: 积分排序
*  @receiver this
**/
func (this *MagicTower) userScoreSort() {

	this.userScore.Sort()
	//广播
	this.broadFightInfo()
}

func (this *MagicTower) broadFightInfo() {
	ntf := &pb.MagicTowerFightNtf{}
	ntf.UserScores = this.userScore.rank(true)
	for k, v := range this.monsterInitTime {
		fight := this.magicTowerfights[k]
		if fight == nil {
			continue
		}
		var magicTowerConf *gamedb.MagicTowerMagicTowerCfg
		gamedb.RangMagicTowerMagicTowerCfgs(func(conf *gamedb.MagicTowerMagicTowerCfg) bool {
			if conf.StageId1 == fight.GetStageConf().Id {
				magicTowerConf = conf
				return false
			}
			return true
		})
		if magicTowerConf == nil {
			logger.Error("获取配置Id异常：%v", fight.GetStageConf().Id)
		}
		for _, vv := range v {
			monsterConf := gamedb.GetMonsterMonsterCfg(vv.monsterId)
			bossInfo := &pb.MagicTowerBossInfo{
				BossName:  monsterConf.Name,
				Status:    int32(vv.status),
				Layer:     int32(magicTowerConf.Id),
				MonsterId: int32(monsterConf.Monsterid),
			}
			if vv.monster != nil {
				bossInfo.MonsterObjId = int32(vv.monster.GetObjId())
				if m, ok := vv.monster.(base.ActorMonster); ok {
					owner := fight.GetUserMainActor(m.Owner())
					if owner != nil {
						bossInfo.OwnerUseId = int32(m.Owner())
						bossInfo.OwnerName = owner.NickName()
					}
				}
			}
			if vv.status == 0 {
				bossInfo.RefreshTime = int32(vv.time)
			}
			ntf.BossInfos = append(ntf.BossInfos, bossInfo)
		}

	}
	for _, vv := range this.magicTowerfights {
		vv.WriteFightInfo(ntf)
	}
}

func (this *MagicTower) recordMonsterInfo(stageId int, monsterRefreshInfo map[int]*MagicTowerMonsterRefresh) {
	fmt.Println(this.locker)
	this.locker.Lock()
	defer this.locker.Unlock()

	if this.monsterInitTime == nil {
		this.monsterInitTime = make(map[int]map[int]*MagicTowerMonsterRefresh)
	}
	if _, ok := this.monsterInitTime[stageId]; !ok {
		this.monsterInitTime[stageId] = make(map[int]*MagicTowerMonsterRefresh)
	}
	for k, v := range monsterRefreshInfo {
		this.monsterInitTime[stageId][k] = v
	}
}

func (this *MagicTower) updateMonsterRefreshInfo(stageId, monsterWave int, status int) {

	this.locker.Lock()
	defer this.locker.Unlock()
	if _, ok := this.monsterInitTime[stageId]; !ok {
		logger.Error("更新怪物刷新状态异常,stageId未找到：%v", stageId)
		return
	}
	if _, ok := this.monsterInitTime[stageId][monsterWave]; !ok {
		logger.Error("更新怪物刷新状态异常，波数异常,stageId:%v,波数：%v", stageId, monsterWave)
		return
	}
	this.monsterInitTime[stageId][monsterWave].status = status
	//广播
	this.broadFightInfo()
}

func (this *MagicTower) updateMonsterRefreshInfoByOwner(monster base.Actor) {

	//广播
	this.broadFightInfo()
}

func (this *MagicTower) GetFightId(stageId int) uint32 {
	this.locker.Lock()
	defer this.locker.Unlock()
	if this.status == 1 {
		return this.magicTowerfights[stageId].id
	}
	return 0
}

func (this *MagicTower) UpdateFrame() {

	now := time.Now()
	if this.status == 0 {
		if int(now.Unix()) >= this.startTime {

			this.magicTowerfights = make(map[int]*MagicTowerFight)
			this.userScoreMap = make(map[int]int) //玩家当前积分
			this.userScore = NewUserScores()
			gamedb.RangMagicTowerMagicTowerCfgs(func(conf *gamedb.MagicTowerMagicTowerCfg) bool {

				fightId, err := fightManager.CreateFight(conf.StageId1, nil)
				if err != nil {
					logger.Error("创建九层魔塔异常,stageId:%v，异常%v", conf.StageId1, err)
				}
				fight := GetFightMgr().GetFight(fightId).(*MagicTowerFight)
				fight.InitStart(this)
				this.locker.Lock()
				this.magicTowerfights[conf.StageId1] = fight
				this.locker.Unlock()
				return true
			})
			this.broadFightInfo()
			this.status = 1
			logger.Info("创建九层魔塔结束")

		}
	} else if this.status == 1 {
		if int(now.Unix()) > this.stopTime+1 {
			this.status = 0
			this.OnEnd()
			this.startTime, this.stopTime = this.GetNextActiveTime()
		}
	}
}

func (this *MagicTower) OnEnd() {

	this.locker.Lock()
	defer this.locker.Unlock()

	msg := &pbserver.FSFightEndNtf{
		FightType: int32(constFight.FIGHT_TYPE_MAGIC_TOWER),
		StageId:   int32(0),
		UseTime:   int32(0),
	}
	this.userScore.Sort()
	rank := this.userScore.rank(false)
	endMsg := &pbserver.MagicTowerFightEnd{
		UserRank: make([]*pbserver.ShabakeRankScore, len(rank)),
	}
	for k, v := range rank {
		endMsg.UserRank[k] = &pbserver.ShabakeRankScore{Id: v.UserId, Score: v.Score}
	}
	rb, _ := endMsg.Marshal()
	msg.CpData = rb

	for k, _ := range this.magicTowerfights {
		msg.StageId = int32(k)
		delete(this.magicTowerfights, k)
		logger.Info("九层魔塔战斗结束,stageid：%v", k)
	}
	net.GetGsConn().SendMessage(msg)

}

func (this *MagicTower) GetNextActiveTime() (int, int) {

	activiteConf := gamedb.GetDailyActivityDailyActivityCfg(pb.DAILYACTIVITYTYPE_MAGIC_TOWER)
	return gamedb.GetActiveTime(activiteConf.OpenTime, activiteConf.CloseTime, activiteConf.Week)
}

func (this *MagicTower) userScoreChangeNtf(stageId int, userId int, changeScore int) {

	fight := this.magicTowerfights[stageId]
	mainActor := fight.GetUserMainActor(userId)
	if mainActor == nil {
		return
	}
	ntf := &pb.FightUserScoreNtf{
		Score:       int32(this.userScoreMap[userId]),
		ChangeScore: int32(changeScore),
		RankScore:   this.userScore.userScoreMap[userId].Score,
	}
	net.GetGateConn().SendMessage(uint32(mainActor.HostId()), mainActor.SessionId(), 0, ntf)
}
