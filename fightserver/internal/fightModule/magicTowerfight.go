package fightModule

import (
	"cqserver/fightserver/conf"
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"fmt"
	"time"
)

type MagicTowerMonsterRefresh struct {
	stageId   int
	time      int64
	status    int
	monsterId int
	monster   base.Actor
}

type MagicTowerFight struct {
	*DefaultFight
	userIntoScoreArea     map[int]int64                     //玩家进入积分区域时间
	lasttime              int64                             //上次积分计算时间
	addScore              int                               //积分增加
	addScoreInterval      int                               //积分增加间隔
	m                     *MagicTower                       //总管理类
	areaRelive            map[int][]int                     //记录区域怪物 复活
	monsterInitTime       map[int]*MagicTowerMonsterRefresh //怪物刷新信息
	fightInfoLocker       *util.Locker
	fightInfo             *pb.MagicTowerFightNtf
	fightInfoNotifyStatus bool                            //战斗信息是否全局推送
	userGetAward          map[int]bool                    //玩家是否领取当前层奖励
	magicTowerConf        *gamedb.MagicTowerMagicTowerCfg //当前层配置
	userTimeAward         map[int]int64                   //玩家计时奖励
}

func NewMagicTowerFight(stageId int) (*MagicTowerFight, error) {
	var err error
	f := &MagicTowerFight{
		userIntoScoreArea: make(map[int]int64),
		fightInfoLocker:   util.NewLocker(0),
		userGetAward:      make(map[int]bool),
		userTimeAward:     make(map[int]int64),
	}
	f.DefaultFight, err = NewDefaultFight(stageId, f)
	if err != nil {
		return nil, err
	}
	for _, v := range gamedb.GetConf().MagicTowerStageId {
		if v[0] == stageId {
			f.addScoreInterval = v[1]
			f.addScore = v[2]
			break
		}
	}

	gamedb.RangMagicTowerMagicTowerCfgs(func(conf *gamedb.MagicTowerMagicTowerCfg) bool {
		if conf.StageId1 == stageId {
			f.magicTowerConf = conf
			return false
		}
		return true
	})

	logger.Info("创建九层魔塔战斗成功")
	return f, nil
}

func (this *MagicTowerFight) InitStart(m *MagicTower) {
	this.m = m
	this.InitMonster()
	this.InitCollection()
	this.areaReliveInit()
	this.SetLifeTime(int64(m.stopTime) - time.Now().Unix())
	this.Start()
}

func (this *MagicTowerFight) InitMonster() error {

	this.monsterInitTime = make(map[int]*MagicTowerMonsterRefresh)
	for k, v := range this.StageConf.Monster_group {

		if len(v) == 3 {

			monsterGroupConf := gamedb.GetMonstergroupMonstergroupCfg(v[0])
			if len(monsterGroupConf.Monsterid) == 1 {
				monsterConf := gamedb.GetMonsterMonsterCfg(monsterGroupConf.Monsterid[0][0])
				if monsterConf.Type == constFight.MONSTER_TYPE_BOSS{
					this.monsterInitTime[k] = &MagicTowerMonsterRefresh{this.StageConf.Id, time.Now().Unix() + int64(v[2]), 0, 0, nil}
					this.monsterInitTime[k].monsterId = monsterConf.Monsterid
				}else{
					this.DefaultFight.InitMonsterByWave(k)
				}
			}else{
				this.DefaultFight.InitMonsterByWave(k)
			}

		} else {
			logger.Error("九重魔塔怪物组配置错误：%v",this.GetStageConf().Id)
		}
	}
	this.m.recordMonsterInfo(this.StageConf.Id, this.monsterInitTime)
	return nil
}

func (this *MagicTowerFight) areaReliveInit() {
	if len(this.StageConf.Monster_num) > 0 {

		this.areaRelive = make(map[int][]int)
		for _, v := range this.monsterActors {

			monster := v.(*actorPkg.MonsterActor)
			areaIndex := monster.GetBirthAreaIndex()
			if _, ok := this.StageConf.Monster_num[areaIndex]; !ok {
				continue
			}
			if _, ok := this.areaRelive[areaIndex]; !ok {
				this.areaRelive[areaIndex] = make([]int, 0)
			}
			this.areaRelive[areaIndex] = append(this.areaRelive[areaIndex], v.GetObjId())
		}
	}
}

func (this *MagicTowerFight) WriteFightInfo(info *pb.MagicTowerFightNtf) {
	this.fightInfoLocker.Lock()
	defer this.fightInfoLocker.Unlock()
	this.fightInfo = info
	this.fightInfoNotifyStatus = false
}

func (this *MagicTowerFight) UpdateFrame() {

	if this.status != FIGHT_STATUS_RUNNING {
		return
	}
	if this.fightInfo != nil && !this.fightInfoNotifyStatus {
		this.fightInfoLocker.Lock()
		this.Scene.NotifyAll(this.fightInfo)
		this.fightInfoNotifyStatus = true
		this.fightInfoLocker.Unlock()
	}

	now := time.Now().Unix()
	if now-this.lasttime < 1 {
		return
	}

	for k, v := range this.monsterInitTime {
		if v.status == 0 {
			if now > v.time {

				monster, err := this.DefaultFight.InitMonsterByWave2(k)
				if err != nil {
					logger.Error("初始化怪物数据异常：%v", err)
				}
				v.status = 1
				v.monster = monster[0]
				this.m.updateMonsterRefreshInfo(this.StageConf.Id, k, 1)
			}
		}
	}

	this.lasttime = now

	users := this.GetAllUserIds()
	userIdMap := make(map[int]bool)
	addScoreUser := make([]int, 0)
	//记录区域内的玩家，并计算分数
	for _, userId := range users {
		allDie := this.CheckUserDieByUserId(userId)
		if allDie {
			continue
		}
		if t, ok := this.userIntoScoreArea[userId]; ok {
			if now-t >= int64(this.addScoreInterval) {
				addScoreUser = append(addScoreUser, userId)
				this.userIntoScoreArea[userId] = now
			}
		} else {
			this.userIntoScoreArea[userId] = now
		}
		userIdMap[userId] = true
	}

	//删除未在区域的玩家
	for k, _ := range this.userIntoScoreArea {
		if !userIdMap[k] {
			delete(this.userIntoScoreArea, k)
		}
	}
	this.m.addUseScore(int(this.id), this.StageConf.Id, addScoreUser, this.addScore)
	//计时奖励
	this.timeAward()
}

func (this *MagicTowerFight) timeAward() {

	if this.magicTowerConf.RewardsSpecialtime <= 0 {
		return
	}
	intoBagItem := make(map[int32]int32)
	for _, v := range this.magicTowerConf.RewardsSpecial {
		intoBagItem[int32(v.ItemId)] = int32(v.Count)
	}
	now := time.Now().Unix()
	userIds := this.GetPlayerUserids()
	for _, v := range userIds {

		allDie := this.CheckUserDieByUserId(v)
		if allDie {
			continue
		}
		if t, ok := this.userTimeAward[v]; ok {
			if t+int64(this.magicTowerConf.RewardsSpecialtime) <= now {
				mainActor := this.GetUserMainActor(v)
				base.AddItems(mainActor, intoBagItem, int32(constBag.OpTypeMagicLayerTimeAward))
				this.userTimeAward[v] = now
			}
		} else {
			this.userTimeAward[v] = now
		}
	}
}

func (this *MagicTowerFight) OnDie(actor, killer base.Actor) {

	if killer == nil {
		if conf.Conf.Sandbox {
			panic(fmt.Sprintf("击杀者为空：%v,%v,%v", actor.NickName(), actor.GetObjId(), actor.GetUserId()))
		}
		return
	}

	userId := killer.GetUserId()
	if userId <= 0 {
		return
	}
	if actor.GetType() == pb.SCENEOBJTYPE_USER || actor.GetType() == pb.SCENEOBJTYPE_FIT {

		mainActor := this.GetUserMainActor(actor.GetUserId())
		allDie := this.CheckUserAllDie(mainActor)
		if allDie {
			this.m.addScore(this.StageConf.Id, userId, gamedb.GetConf().MagicTowerMark, true)
			this.BossOwnerChange(actor, killer)
		}
	} else if actor.GetType() == pb.SCENEOBJTYPE_MONSTER {

		monster := actor.(*actorPkg.MonsterActor)
		//积分添加
		for _, v := range this.monsterInitTime {
			if v.monster != nil && v.monster.GetObjId() == actor.GetObjId() {
				v.status = 2
			}
		}

		score := this.getMonsterScore(monster.GetMonsterT().Monsterid)
		this.m.addScore(this.StageConf.Id, userId, score, true)
		//怪物复活
		areaIndex := monster.GetBirthAreaIndex()
		if num, ok := this.StageConf.Monster_num[areaIndex]; ok {

			aliveNum := 0
			for _, v := range this.areaRelive[areaIndex] {
				monsterActor := this.GetActorByObjId(v)
				if monsterActor.GetProp().HpNow() > 0 {
					aliveNum += 1
				}
			}
			if aliveNum < num {
				for _, v := range this.areaRelive[areaIndex] {
					actorObj := this.GetActorByObjId(v)
					if actorObj.GetProp().HpNow() <= 0 {

						monsterActor := actorObj.(*actorPkg.MonsterActor)
						monsterActor.SetOwner(0)
						actorObj.Relive(monsterActor.MonsterT.ReliveAddrType, constFight.RELIVE_TYPE_NOMAL)
					}
				}
			}
		}
	}
}

func (this *MagicTowerFight) BossOwnerChange(dieActor, killer base.Actor) {

	for _, v := range this.monsterActors {
		if m, ok := v.(base.ActorMonster); ok {
			if m.Owner() == dieActor.GetUserId() {
				v.AddOwner(killer, true)
			}
		}
	}
}

func (this *MagicTowerFight) OnEnterUser(userId int) {

	actor := this.GetUserMainActor(userId)
	this.m.addRank(userId, actor.NickName())
	if this.fightInfo != nil {
		net.GetGateConn().SendMessage(uint32(actor.HostId()), actor.SessionId(), 0, this.fightInfo)
	}
}

func (this *MagicTowerFight) OnLeaveUser(userId int) {
	delete(this.userTimeAward, userId)
}

func (this *MagicTowerFight) OnEnd() {

	logger.Info("九层魔塔战斗结束,stageid：%v", this.StageConf.Id)
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 15)
}

func (this *MagicTowerFight) OnCollection(colllection map[int]int) {

	for k, v := range colllection {

		conf := gamedb.GetCollectionCollectionCfg(v)
		if len(conf.Effect) > 1 && conf.Effect[0] == constFight.COLLECTION_EFFECT_SCORE {
			actor := this.GetActorByObjId(k)
			if actor.GetUserId() > 0 {
				this.m.addScore(this.StageConf.Id, actor.GetUserId(), conf.Effect[1], true)
			}
		}
	}
}

func (this *MagicTowerFight) GetFightInfos() *pb.MagicTowerFightNtf {
	return this.fightInfo
}

func (this *MagicTowerFight) MonsterDrop(dropMonsterId int, dropX, dropY int, owner base.Actor, dropItems []*pbserver.ItemUnit) {

	intoBagItem := make(map[int32]int32)

	for _, v := range dropItems {
		intoBagItem[v.ItemId] += v.ItemNum
	}

	if len(intoBagItem) > 0 {
		base.AddItems(owner, intoBagItem, int32(constBag.OpTypePickUp))
	}

}

func (this *MagicTowerFight) getMonsterScore(monsterId int) int {

	for _, v := range gamedb.GetConf().MagicTowerMonster {
		if v[0] == monsterId {
			return v[1]
		}
	}
	logger.Warn("怪物未配置奖励积分：%v", monsterId)
	return 0
}

func (this *MagicTowerFight) FightScoreLess(userId int, lessNum int) (int, error) {

	allDie := this.CheckUserDieByUserId(userId)
	if allDie {
		return 0, gamedb.ERRPLAYERDIE
	}

	return this.m.LessUserScore(this.StageConf.Id, userId, lessNum)
}

func (this *MagicTowerFight) OnBossOwnerChange(monster base.Actor) {
	this.m.updateMonsterRefreshInfoByOwner(monster)
}

/**
*  @Description: 获取玩家积分 是否领取奖励
*  @receiver this
*  @param userId
*  @param nowIsGetAward 当前请求是否是领取奖励请求
*  @return int		积分
*  @return bool		是否已领去奖励（请求领奖时并可以领奖时，也返回已领取）
*  @return bool		标记玩家是否可以领奖
**/
func (this *MagicTowerFight) GetMagicUserInfo(userId int, nowIsGetAward bool) (int, bool, bool) {

	userScore := this.m.GetUserScore(userId)
	isGetAward := this.userGetAward[userId]
	if nowIsGetAward {
		if isGetAward {
			return userScore, isGetAward, false
		} else {
			this.userGetAward[userId] = true
			logger.Info("标记玩家领取当前层奖励", this.GetStageConf().Id)
			return userScore, true, true
		}
	} else {
		return userScore, isGetAward, false
	}
}
