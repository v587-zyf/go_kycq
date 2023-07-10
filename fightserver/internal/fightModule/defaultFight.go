package fightModule

import (
	"cqserver/fightserver/internal/ai"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/golibs/common"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"errors"
	"fmt"
	"math/rand"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"

	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
)

const (
	MaxFightLifeTime   = 60 * 20
	FIGHT_EVENT_LENGTH = 50
)
const (
	fightEndReasonNomal   = iota
	fightEndReasonTimeout = iota
)

const (
	FIGHT_STATUS_READY   = 1 //准备
	FIGHT_STATUS_RUNNING = 2 //运行
	FIGHT_STATUS_CLOSING = 3 //关闭中
)

type DefaultFight struct {
	id            uint32                    //战斗id
	StageConf     *gamedb.StageStageCfg     //战斗基础配置u
	userActors    map[int]base.Actor        //所有玩家数据
	userPets      map[int]base.Actor        //所有战宠
	userFits      map[int]base.Actor        //所有合体
	userSummons   map[int]base.Actor        //所有道士宠物
	monsterActors map[int]base.Actor        //所有怪物数据
	playerActors  map[int]*base.PlayerActor //玩家数据（userId=>玩家武将）
	playerPickUp  map[int]map[int]int       //玩家拾取东西

	context        base.Fight        //实际战斗类（子类）
	Scene          *scene.Scene      //地图场景
	Messages       chan base.Message //战斗chan
	fightEvents    chan *base.FightEvent
	loopChan       chan struct{} //控制器
	once           sync.Once
	status         int
	createTime     int64 //战斗创建时间
	lifeTime       int64 //战斗生命
	NextStatusTime int64
}

func NewDefaultFight(stageId int, context base.Fight) (*DefaultFight, error) {

	stageConf := gamedb.GetStageStageCfg(stageId)
	if stageConf == nil {
		return nil, gamedb.ERRSETTINGNOTFOUND
	}

	sceneVar, err := scene.NewScene(stageId)
	if err != nil {
		return nil, err
	}

	fight := &DefaultFight{
		createTime:    time.Now().Unix(),
		lifeTime:      MaxFightLifeTime,
		StageConf:     stageConf,
		userActors:    make(map[int]base.Actor),
		userPets:      make(map[int]base.Actor),
		userFits:      make(map[int]base.Actor),
		userSummons:   make(map[int]base.Actor),
		monsterActors: make(map[int]base.Actor),
		playerActors:  make(map[int]*base.PlayerActor),
		Scene:         sceneVar,
		Messages:      make(chan base.Message, 100),
		fightEvents:   make(chan *base.FightEvent, FIGHT_EVENT_LENGTH),
		context:       context,
		status:        FIGHT_STATUS_READY,
		playerPickUp:  make(map[int]map[int]int),
	}

	if stageConf.LifeTime > 0 {
		fight.lifeTime = int64(stageConf.LifeTime + 1)
	}

	logger.Info("创建战斗数据,配置数据：%v-%v", stageId, *stageConf)
	return fight, nil
}

func (this *DefaultFight) GetFightExtMark() int {
	return 0
}

func (this *DefaultFight) GetContext() base.Fight {
	return this.context
}

func (this *DefaultFight) SetId(id uint32) {
	this.id = id
}

func (this *DefaultFight) GetId() uint32 {
	return this.id
}

func (this *DefaultFight) GetStageConf() *gamedb.StageStageCfg {
	return this.StageConf
}

func (this *DefaultFight) SetLifeTime(time int64) {
	this.lifeTime = time
}

/**
 *  @Description: 设置战斗状态，进入下一个状态时间
 *  @param status
 *  @param nextStatusTime
 */
func (this *DefaultFight) SetFightStatusAndNextStatusTime(status int, nextStatusTime int64) {

	if nextStatusTime == 0 {
		this.NextStatusTime = 0
	} else {
		this.NextStatusTime = time.Now().Unix() + nextStatusTime
	}
	this.status = status
	if status == FIGHT_STATUS_CLOSING {
		this.Range(func(actor base.Actor) bool {
			actor.GetFSM().Stop()
			return false
		})
	} else if status == FIGHT_STATUS_RUNNING {
		this.Range(func(actor base.Actor) bool {
			if actor.GetType() == pb.SCENEOBJTYPE_MONSTER || actor.GetType() == pb.SCENEOBJTYPE_USER {
				actor.GetFSM().Recover()
			}
			return false
		})
	}
}

func (this *DefaultFight) SendMessage(message base.Message) {
	this.Messages <- message
}

func (this *DefaultFight) GetMessageChan() chan base.Message {
	return this.Messages
}

//=======================================战斗角色数据相关====================================================

func (this *DefaultFight) InitCollection() error {

	if len(this.StageConf.Collection) <= 0 {
		return nil
	}

	for _, v := range this.StageConf.Collection {
		collectionConf := gamedb.GetCollectionCollectionCfg(v)
		for i := 0; i < collectionConf.Goods; i++ {
			birthArea := collectionConf.Collection[rand.Intn(len(collectionConf.Collection))]
			birthPoint, err := this.Scene.GetBirthPointByAreaIndex(birthArea)
			if err != nil {
				return err
			}
			sceneItem := scene.NewSceneCollection(collectionConf.Id)
			err1 := sceneItem.EnterScene(this.Scene, birthPoint)
			if err1 != nil {
				return err1
			}
		}
	}
	return nil
}

func (this *DefaultFight) InitMonster() error {

	if len(this.StageConf.Monster_group) <= 0 {
		return nil
	}

	var err error
	for k, _ := range this.StageConf.Monster_group {
		_, err = this.InitMonsterByWave(k)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *DefaultFight) InitMonsterByWave2(waveIndex int) ([]base.Actor, error) {

	if waveIndex >= len(this.StageConf.Monster_group) {
		errStr := fmt.Sprintf("wave index err,count wave:%v", len(this.StageConf.Monster_group))
		logger.Error(errStr)
		return nil, errors.New(errStr)
	}

	monsterGroup := this.StageConf.Monster_group[waveIndex]
	if len(monsterGroup) < 2 {
		return nil, gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg(monsterGroup)
	}
	allMonsters := make([]base.Actor, 0)
	interval := 2
	if this.StageConf.Type == constFight.FIGHT_TYPE_MAGIC_TOWER {
		interval = 3
	}
	for i, l := 0, len(monsterGroup); i < l; i += interval {
		monsterGroupId, monsterBirthId := monsterGroup[i], monsterGroup[i+1]
		monsterGroupT := gamedb.GetMonstergroupMonstergroupCfg(monsterGroupId)
		if monsterGroupT == nil {
			return nil, gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("monsterGroup", monsterGroupId)
		}

		randIndex := 0
		if len(monsterGroupT.Monsterid) > 1 {
			weightSlice := make([]int, len(monsterGroupT.Monsterid))
			for i, w := range monsterGroupT.Monsterid {
				weightSlice[i] = w[2]
			}
			randIndex = common.RandWeightByIntSlice(weightSlice)
		}
		monsterId, monsterNum := monsterGroupT.Monsterid[randIndex][0], monsterGroupT.Monsterid[randIndex][1]
		monsters := make([]base.Actor, monsterNum)
		monsterPoints := make([]*scene.Point, monsterNum)
		for count := 0; count < monsterNum; count++ {
			birthPoint, err := this.Scene.GetBirthPointByAreaIndex(monsterBirthId)
			if err != nil {
				return nil, gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("monster born", monsterBirthId)
			}
			monsterActor := actorPkg.NewMonsterActor(monsterId, ai.NewMonsterAI, monsterBirthId)
			monsters[count] = monsterActor
			monsterPoints[count] = birthPoint
			allMonsters = append(allMonsters, monsterActor)
			//this.Enter(monsterActor, birthPoint)
			//birthPoint, _ = this.Scene.GetBirthPointByAreaIndex(monsterBirthId)
		}
		err11 := this.EnterMuli(monsters, monsterPoints, 0)
		if err11 != nil {
			return nil, err11
		}
		logger.Info("初始化战斗怪物数据：%v,随机怪物为：%v 数量：%v,出生区域：%v", this.StageConf.Monster_group[waveIndex], monsterId, monsterNum, monsterBirthId)
	}

	return allMonsters, nil
}

func (this *DefaultFight) InitMonsterByWave(waveIndex int) (int, error) {

	monsters, err := this.InitMonsterByWave2(waveIndex)
	if err != nil {
		return 0, err
	}
	return len(monsters), nil
}

func (this *DefaultFight) Range(f func(actor base.Actor) bool) {

	for _, monster := range this.monsterActors {
		if f(monster) {
			return
		}
	}
	for _, user := range this.userActors {
		if f(user) {
			return
		}
	}
	for _, pet := range this.userPets {
		if f(pet) {
			return
		}
	}
	for _, fit := range this.userFits {
		if f(fit) {
			return
		}
	}

	for _, summon := range this.userSummons {
		if f(summon) {
			return
		}
	}
}

func (this *DefaultFight) GetOnlineActor(sessionId uint32) base.Actor {

	for _, v := range this.playerActors {
		if v.SessionId() == sessionId {
			return v.GetHeroActor(constUser.USER_HERO_MAIN_INDEX)
		}
	}
	return nil
}

func (this *DefaultFight) GetActorByObjId(objId int) base.Actor {
	if actor, ok := this.userActors[objId]; ok {
		return actor
	}
	if actor, ok := this.monsterActors[objId]; ok {
		return actor
	}
	if actor, ok := this.userPets[objId]; ok {
		return actor
	}
	if actor, ok := this.userFits[objId]; ok {
		return actor
	}
	if actor, ok := this.userSummons[objId]; ok {
		return actor
	}
	return nil
}

func (this *DefaultFight) GetUserActors() map[int]base.Actor {
	return this.userActors
}

func (this *DefaultFight) GetPlayerUserids() []int {
	userIds := make([]int, len(this.playerActors))
	i := 0
	for k, _ := range this.playerActors {
		userIds[i] = k
		i++
	}
	return userIds
}

func (this *DefaultFight) CheckUserDieByUserId(userId int) bool {

	player := this.playerActors[userId]
	if player != nil {
		return player.CheckUserDie()
	}
	return true
}

func (this *DefaultFight) CheckUserAllDieByHp(dieActorUser base.Actor) bool {

	if dieActorUser.GetType() == pb.SCENEOBJTYPE_FIT {
		return true
	}
	if dieActorUser.GetType() != pb.SCENEOBJTYPE_USER {
		return false
	}
	actorsHeros := this.playerActors[dieActorUser.GetUserId()].HeroActors()
	for _, v := range actorsHeros {
		if v.GetProp().HpNow() > 0 {
			return false
		}
	}
	return true
}

func (this *DefaultFight) CheckUserAllDie(actorUser base.Actor) bool {

	if actorUser.GetType() == pb.SCENEOBJTYPE_FIT {
		return true
	}
	if actorUser.GetType() != pb.SCENEOBJTYPE_USER {
		return false
	}
	player := this.playerActors[actorUser.GetUserId()]
	if player != nil {
		return player.CheckUserDie()
	}
	return true
}

func (this *DefaultFight) GetUserBirthPoint(index int) (*scene.Point, error) {
	return this.Scene.GetUserBirthPoint(index)
}

func (this *DefaultFight) GetPlayer(fightUser *pbserver.User, stageType int, enterType int) map[int]*actorPkg.UserActor {

	userId := int(fightUser.UserInfo.UserId)
	player := this.playerActors[userId]
	if player == nil {
		player = base.NewPlayer(fightUser.SessionId, fightUser.UserInfo.RedPacket, fightUser.UserInfo.PublicSkill)
		player.SetStageFightNum(int(fightUser.UserInfo.StageFightNum))
		this.playerActors[userId] = player
	}
	player = this.playerActors[userId]
	player.SetUserCombat(int(fightUser.UserInfo.UserCombat))
	player.SetFightNum(int(fightUser.UserInfo.DarkPalaceTimes))
	player.SetDaBaoEnergy(int(fightUser.UserInfo.DabaoEnergy))
	userActors := make(map[int]*actorPkg.UserActor)

	//协助者队伍处理
	toHelpUserId := int(fightUser.ToHelpUserId)
	if toHelpUserId > 0 {
		if _, ok := this.playerActors[int(fightUser.ToHelpUserId)]; ok {
			allDie := this.CheckUserDieByUserId(toHelpUserId)
			if !allDie {
				player.SetToHelpUserId(toHelpUserId)
				//fightUser.TeamId = int32(this.GetUserMainActor(toHelpUserId).TeamIndex())
			}
		}
	}

	heros := fightUser.UserInfo.Heros
	mainActor := this.GetUserMainActor(int(fightUser.UserInfo.UserId))

	for _, heroIndex := range constUser.USER_HERO_INDEX {

		var userActor *actorPkg.UserActor
		if v, ok := heros[int32(heroIndex)]; ok {

			if fightUser.UserType == constFight.FIGHT_USER_TYPE_PLAYER_data || fightUser.UserType == constFight.FIGHT_USER_TYPE_CONF {

				userActor = actorPkg.NewUserActor(player, fightUser, int(v.Index), mainActor, ai.NewUserAI)
			} else {
				userActor = actorPkg.NewUserActor(player, fightUser, int(v.Index), mainActor, ai.NewUserFsm)
			}
			userActors[int(v.Index)] = userActor
			if mainActor == nil && heroIndex == constUser.USER_HERO_MAIN_INDEX {
				mainActor = userActor
			}
		}
	}
	//计算武将总血量
	hpTotal := 0
	for _, v := range constUser.USER_HERO_INDEX {
		if u, ok := userActors[v]; ok {
			hpTotal += u.GetProp().Get(pb.PROPERTY_HP)
		} else if u := this.playerActors[userId].GetHeroActor(v); u != nil {
			hpTotal += u.GetProp().Get(pb.PROPERTY_HP)
		}
	}
	mainActor.(base.ActorPlayer).GetPlayer().SetUserHpTotal(hpTotal)
	return userActors
}

func (this *DefaultFight) UpdateUserReqPacketInfo(userId int, redPacketInfo *pbserver.ActorRedPacket) {
	if p, ok := this.playerActors[userId]; ok {
		p.SetRedPacketInfo(redPacketInfo)
	}
}

func (this *DefaultFight) getBirthType() int {
	stageType := this.GetStageConf().Type
	birthType := constFight.FIGHT_BIRTH_TYPE_TRIANGLE
	if stageType == constFight.FIGHT_TYPE_ARENA || stageType == constFight.FIGHT_TYPE_FIELD {
		birthType = constFight.FIGHT_BIRTH_TYPE_LINE
	}
	return birthType
}

func (this *DefaultFight) getDir(userId, teamIndex int) int {
	mainActor := this.GetUserMainActor(userId)
	if mainActor != nil {
		return mainActor.GetDir()
	}
	stageType := this.GetStageConf().Type
	if stageType == constFight.FIGHT_TYPE_ARENA || stageType == constFight.FIGHT_TYPE_FIELD {
		if teamIndex == constFight.FIGHT_TEAM_ONE {
			return scene.DIR_RIGHT_TOP
		} else {
			return scene.DIR_LEFT_BOTTOM
		}
	}
	return scene.DIR_TOP
}

/**
*  @Description: 获取主武将出生坐标
*  @receiver this
*  @param userId
*  @param birthArea
*  @param teamIndex
*  @return *scene.Point
**/
func (this *DefaultFight) getMainActorBirthPoint(userId int, birthArea int, teamIndex int) *scene.Point {
	//新武将进入
	mainUserActor := this.GetUserMainActor(userId)
	if mainUserActor != nil {

		return mainUserActor.Point()
	} else {
		var birthPoint *scene.Point
		var err error
		if birthArea > 0 {
			birthPoint, err = this.Scene.GetBirthPointByAreaIndex(birthArea)
			if err != nil {
				logger.Error("获取出生点异常，出生区域,指定：%v", birthArea)
			}
		} else {
			birthAreaIndex := -1
			if this.GetStageConf().Type == constFight.FIGHT_TYPE_ARENA || this.GetStageConf().Type == constFight.FIGHT_TYPE_FIELD {
				birthAreaIndex = teamIndex
			}
			birthPoint, err = this.GetUserBirthPoint(birthAreaIndex)
			if err != nil {
				logger.Error("获取出生点异常，出生区域：%v", birthAreaIndex)
			}
		}
		return birthPoint
	}
}

func (this *DefaultFight) UserEnter(userEnterMsg *pbserver.FSEnterFightReq) error {

	userId := int(userEnterMsg.FightUser.UserInfo.UserId)
	logger.Info("接收到gs服发来的，请求进入游戏,fightId:%v,玩家：%v,指定出生区域：%v，队伍：%v", this.id, userId, userEnterMsg.FightUser.BirthArea, userEnterMsg.FightUser.TeamId)
	userActors := this.GetPlayer(userEnterMsg.FightUser, this.GetStageConf().Type, int(userEnterMsg.EnterType))
	birthPoint := this.getMainActorBirthPoint(userId, int(userEnterMsg.FightUser.BirthArea), int(userEnterMsg.FightUser.TeamId))
	if birthPoint == nil {
		return gamedb.ERRSCENEPOINT
	}
	var err error
	dir := this.getDir(userId, int(userEnterMsg.FightUser.TeamId))
	birthType := this.getBirthType()

	//武将进入
	for k, _ := range userEnterMsg.FightUser.UserInfo.Heros {
		heroIndex := int(k)
		if userActor, ok := userActors[heroIndex]; ok {
			heroBirthPoint := birthPoint
			if heroIndex != constUser.USER_HERO_MAIN_INDEX {
				heroBirthPoint = this.Scene.GetHeroBirthPoint(userActor, birthPoint, dir, heroIndex, birthType)
			}
			userActor.SetDir(dir)
			err = this.Enter(userActor, heroBirthPoint)
			if err != nil {
				return err
			}
		}
	}
	if userEnterMsg.EnterType == constFight.ENTER_FIGHT_TYPE_NOMAL {
		//战宠 进入
		this.petEnter(userId, userEnterMsg.FightUser.UserInfo.Pet)
		this.context.OnEnterUser(int(userEnterMsg.FightUser.UserInfo.UserId))
	}
	return nil
}

/**
 *  @Description: 道士宠物进入
 *  @param leaderActor
 *  @param summonId
 */
func (this *DefaultFight) EnterSummon(leaderActor base.Actor, summonId int) {
	actorFitUnit := actorPkg.NewSummonActor(leaderActor, summonId, ai.NewUserAI)
	birthPoint := this.GetScene().GetPointByPointRange(actorFitUnit, leaderActor.Point().X(), leaderActor.Point().Y(), nil)
	this.Enter(actorFitUnit, birthPoint)
}

func (this *DefaultFight) EnterMuli(actor1 []base.Actor, point []*scene.Point, enterType int) error {

	objs := make([]scene.ISceneObj, len(actor1))
	for k, v := range actor1 {
		objs[k] = v.GetContext()
	}
	err := this.Scene.AddSceneObjs(objs, point, enterType)
	if err != nil {
		logger.Error("多角色进入场景异常：%v", len(objs), len(point), err)
		return err
	}
	//角色进入，不需要再通知场景
	for k, v := range actor1 {
		this.enterNew(v, point[k], true)
	}
	return nil
}

func (this *DefaultFight) Enter(actor1 base.Actor, point *scene.Point) error {

	return this.enterNew(actor1, point, false)
}

func (this *DefaultFight) enterNew(actor1 base.Actor, point *scene.Point, justInFight bool) error {
	if actor1.GetType() == pb.SCENEOBJTYPE_USER {
		this.userActors[actor1.GetObjId()] = actor1
		this.playerActors[actor1.GetUserId()].AddHeroActors(actor1.(base.ActorUser).GetHeroIndex(), actor1)

	} else if actor1.GetType() == pb.SCENEOBJTYPE_MONSTER {
		this.monsterActors[actor1.GetObjId()] = actor1
	} else if actor1.GetType() == pb.SCENEOBJTYPE_PET {
		this.userPets[actor1.GetObjId()] = actor1
		this.playerActors[actor1.GetUserId()].SetPetActor(actor1)
	} else if actor1.GetType() == pb.SCENEOBJTYPE_FIT {
		this.userFits[actor1.GetObjId()] = actor1
		this.playerActors[actor1.GetUserId()].SetFitActor(actor1)
	} else if actor1.GetType() == pb.SCENEOBJTYPE_SUMMON {
		this.playerActors[actor1.GetUserId()].AddSummonActor(actor1.GetObjId(), actor1)
		this.userSummons[actor1.GetObjId()] = actor1
	}

	actor1.SetFight(this.context)
	if !justInFight {
		err := actor1.EnterScene(this.Scene, point)
		if err != nil {
			return err
		}
	}
	actor1.SetBirthPoint(point)
	this.GetContext().OnActorEnter(actor1)
	return nil
}

func (this *DefaultFight) Leave(actor base.Actor) {

	if actor.GetType() == pb.SCENEOBJTYPE_USER {
		delete(this.userActors, actor.GetObjId())
	} else if actor.GetType() == pb.SCENEOBJTYPE_MONSTER {
		delete(this.monsterActors, actor.GetObjId())
	} else if actor.GetType() == pb.SCENEOBJTYPE_PET {
		playerActor := this.playerActors[actor.GetUserId()]
		if playerActor != nil {
			playerActor.SetPetActor(nil)
		}
		delete(this.userPets, actor.GetObjId())
	} else if actor.GetType() == pb.SCENEOBJTYPE_FIT {
		playerActor := this.playerActors[actor.GetUserId()]
		if playerActor != nil {
			playerActor.SetFitActor(nil)
		}
		delete(this.userFits, actor.GetObjId())
	} else if actor.GetType() == pb.SCENEOBJTYPE_SUMMON {
		playerActor := this.playerActors[actor.GetUserId()]
		if playerActor != nil {
			playerActor.RemoveSummonActor(actor.GetObjId())
		}
		delete(this.userSummons, actor.GetObjId())
	}
	this.context.OnLeave(actor)
	actor.LeaveScene()
}

func (this *DefaultFight) LeaveUser(userId int) {

	if userActor, ok := this.playerActors[userId]; ok {
		heroActors := userActor.HeroActors()
		for _, actor := range heroActors {
			delete(this.userActors, actor.GetObjId())
			this.context.OnLeave(actor)
			actor.LeaveScene()
		}
		//战宠
		petActor := userActor.PetActor()
		if petActor != nil {
			this.context.OnLeave(petActor)
			petActor.LeaveScene()
			delete(this.userPets, petActor.GetObjId())
		}
		fitActor := userActor.FitActor()
		if fitActor != nil {
			this.context.OnLeave(fitActor)
			fitActor.LeaveScene()
			delete(this.userFits, fitActor.GetObjId())
		}
		summonActors := userActor.SummonActors()
		if summonActors != nil && len(summonActors) > 0 {
			for _, v := range summonActors {
				this.context.OnLeave(v)
				v.LeaveScene()
				delete(this.userSummons, v.GetObjId())
			}
		}

	}
	delete(this.playerActors, userId)

	//清理怪物归属
	for _, v := range this.monsterActors {
		if m, ok := v.(*actorPkg.MonsterActor); ok {
			if m.Owner() == userId {
				m.SetOwner(0)
			}
		}
	}

	//解除协助关系
	for _, v := range this.playerActors {
		if v.ToHelpUserId() == userId {
			v.SetToHelpUserId(0)
		}
	}

	this.context.OnLeaveUser(userId)
}

func (this *DefaultFight) GetAllUserIds() []int {

	userIds := make([]int, 0)
	for userId, _ := range this.playerActors {
		userIds = append(userIds, userId)
	}
	return userIds
}

func (this *DefaultFight) GetPlayerNum() int {
	return len(this.playerActors)
}

func (this *DefaultFight) GetUserByUserId(userId int) map[int]base.Actor {
	if _, ok := this.playerActors[userId]; ok {
		return this.playerActors[userId].HeroActors()
	}
	return nil
}

func (this *DefaultFight) GetUserFitActor(userId int) base.Actor {
	if _, ok := this.playerActors[userId]; !ok {
		return nil
	}
	return this.playerActors[userId].FitActor()
}

func (this *DefaultFight) GetUserMainActor(userId int) base.Actor {
	if _, ok := this.playerActors[userId]; ok {
		return this.playerActors[userId].GetHeroActor(constUser.USER_HERO_MAIN_INDEX)
	}
	return nil
}

/**
*  @Description: 获取玩家战宠数据
*  @receiver this
*  @param userId
*  @return base.Actor
**/
func (this *DefaultFight) GetPetActor(userId int) base.Actor {

	if _, ok := this.playerActors[userId]; ok {
		return this.playerActors[userId].PetActor()
	}
	return nil
}

func (this *DefaultFight) UpdateUserFigntInfo(userInfo *pbserver.Actor, heroIndex int) {

	userActors := this.GetUserByUserId(int(userInfo.UserId))
	if userActors != nil && len(userActors) > 0 {
		for _, userActor := range userActors {
			user := userActor.(*actorPkg.UserActor)
			if _, ok := userInfo.Heros[int32(user.HeroIndex())]; ok {
				user.UpdateUserInfo(userInfo, user.HeroIndex())
			}
		}
	}
}

/**
*  @Description: 变更玩家为协助者
*  @receiver this
*  @param userId
*  @param toHelpUserId
**/
func (this *DefaultFight) ChangeUserToHelper(userId, toHelpUserId int) {

	player := this.playerActors[userId]
	if player == nil {
		logger.Error("玩家申请变更为协助者，协助玩家数未找到：玩家：%v，被协助者：%v", userId, toHelpUserId)
		return
	}
	toHelpPlayer := this.playerActors[toHelpUserId]
	if toHelpPlayer == nil {
		logger.Info("玩家申请变更为协助者，被协助玩家数未找到：玩家：%v，被协助者：%v", userId, toHelpUserId)
		return
	}

	player.SetToHelpUserId(toHelpUserId)
	logger.Info("玩家申请变更为协助者，协助者：%v,被协助者：%v", userId, toHelpUserId)
	//teamIndex := toHelpPlayer.GetHeroActor(constUser.USER_HERO_MAIN_INDEX).TeamIndex()
	//player.ChangeTeam(teamIndex)
	//ntf := &pb.FightUserChangeToHelperNtf{
	//	UserId:       int32(userId),
	//	ToHelpUserId: int32(toHelpUserId),
	//}
	//this.Scene.NotifyAll(ntf)
}

func (this *DefaultFight) PlayerFightNumChange(req *pbserver.GsToFsFightNumChangeReq) {

	player := this.playerActors[int(req.UserId)]
	if player == nil {
		logger.Error("接收到更新玩家战斗次数，玩家数据未找到")
		return
	}
	player.SetFightNum(int(req.FightNumChange))
}

func (this *DefaultFight) FightScoreLess(userId int, lessNum int) (int, error) {
	return 0, gamedb.ERRUNKNOW
}

func (this *DefaultFight) RandomDelivery(userId int, rand bool) {

	for _, v := range this.userActors {
		if v.GetUserId() == userId {
			var point *scene.Point
			if !rand && this.StageConf.Type == constFight.FIGHT_TYPE_MAIN_CITY {
				point = v.BirthPoint()
			} else {
				point = this.Scene.RandomDelivery(v, false)
			}
			if point != nil {
				this.Scene.MoveSceneObj(v, point, pb.MOVETYPE_RANDOM_STONE, true, true)
			}
		}
	}
}

func (this *DefaultFight) KickActorByGate(serverId int) {
	this.Range(func(inactor base.Actor) bool {
		if inactor.GetType() == base.ActorTypeUser {
			if inactor.(*actorPkg.UserActor).HostId() == serverId {
				this.LeaveUser(inactor.GetUserId())
			}
		}
		return false
	})
}

/**
*  @Description: 获取或者boss的数量
*  @receiver this
*  @return int
**/
func (this *DefaultFight) GetBossAliveNum() int {

	num := 0
	for _, v := range this.monsterActors {
		if v.GetProp().HpNow() <= 0 {
			continue
		}
		if v.(*actorPkg.MonsterActor).MonsterT.Type == constFight.MONSTER_TYPE_BOSS {
			num += 1
		}
	}
	return num
}

func (this *DefaultFight) GetBossInfos() []*pb.FightBossInfoUnit {

	bossInfos := make([]*pb.FightBossInfoUnit, 0)
	for _, v := range this.monsterActors {
		monster := v.(*actorPkg.MonsterActor)
		if monster.MonsterT.Type == constFight.MONSTER_TYPE_BOSS {
			bossInfo := &pb.FightBossInfoUnit{
				ObjId:     int32(v.GetObjId()),
				MonsterId: int32(monster.MonsterT.Monsterid),
				Point:     v.BirthPoint().ToPbPoint(),
				Hp:        int64(v.GetProp().HpNow()),
				ReliveCD:  int32(v.ReliveTime() / 1000),
			}
			bossInfos = append(bossInfos, bossInfo)
		}
	}
	return bossInfos

}

func (this *DefaultFight) MonsterDrop(dropMonsterId int, dropX, dropY int, owner base.Actor, dropItems []*pbserver.ItemUnit) {

	//points := scene.RandomDropPoint(this.Scene, dropX, dropY)
	if this.playerPickUp[owner.GetUserId()] == nil {
		this.playerPickUp[owner.GetUserId()] = make(map[int]int)
	}
	intoBagItem := make(map[int32]int32)
	var points []*scene.Point
	for _, item := range dropItems {
		if constFight.DROP_ITEM_INTO_BAG[int(item.ItemId)] {
			this.playerPickUp[owner.GetUserId()][int(item.ItemId)] += int(item.ItemNum)
			intoBagItem[item.ItemId] = item.ItemNum
			continue
		}
		if len(points) == 0 {
			points = scene.RandomDropPoint(this.Scene, dropX, dropY)
		}
		if len(points) == 0 {
			logger.Error("物品掉落，随机坐标异常：归属：%v,战斗：%v,位置：%v-%v", this.StageConf.Id, dropX, dropY)
			break
		}

		point := points[0]
		points = points[1:]
		sceneItem := scene.NewSceneItem(owner.GetUserId(), owner.NickName(), dropMonsterId, int(item.ItemId), int(item.ItemNum), this.StageConf.DropItemDisappearTime)
		err := sceneItem.EnterScene(this.Scene, point)
		if err != nil {
			logger.Error("物品掉落添加到地图异常：归属：%v,战斗：%v,物品：%v-%v，异常：%v", owner, this.StageConf.Type, int(item.ItemId), int(item.ItemNum), err)
		}
	}
	if len(intoBagItem) > 0 {
		base.AddItems(owner, intoBagItem, int32(constBag.OpTypePickUp))
	}
}

func (this *DefaultFight) PickUp(userId int, objIds []int32, isPick bool) (map[int32]*pbserver.ItemUnitForPickUp, error) {

	actor := this.GetUserMainActor(userId)
	if actor == nil {
		return nil, gamedb.ERRUNFOUNDUSER
	}
	//point := actor.Point()
	//if point == nil {
	//	logger.Error("道具拾取获取玩家坐标失败：%v", actor.GetUserId())
	//	return nil, gamedb.ERRUSERPOINT
	//}
	if len(objIds) == 0 {
		objIds = common.ConvertIntSlice2Int32Slice(this.Scene.GetAllItemObjsByPlayer(actor.GetUserId()))
	}

	pickUpItems := make(map[int32]*pbserver.ItemUnitForPickUp)
	if len(objIds) == 0 {
		return pickUpItems, nil
	}

	for _, v := range objIds {

		sceneObj := this.Scene.GetSceneObj(int(v))
		if sceneObj == nil {
			logger.Debug("客户端申请拾取物品,玩家：%v,物品不存在：%v", userId, v)
			continue
		}
		dropItemObj := sceneObj.GetContext()
		if dropItemObj.GetType() != pb.SCENEOBJTYPE_ITEM {
			continue
		}
		dropItem := dropItemObj.(*scene.SceneItem)

		//if scene.DistanceByPoint(dropItem.Point(), point) > 2 {
		//	logger.Error("玩家距离物品过远，不能拾取:玩家位置:%v，物品位置：%v", point.ToString(), dropItem.Point().ToString())
		//	continue
		//}

		if !dropItem.CanPickUp(actor.GetUserId()) {
			continue
		}
		pickUpItems[v] = &pbserver.ItemUnitForPickUp{ItemId: 0, ItemNum: 0}
		pickUpItems[v].ItemId = int32(dropItem.ItemId())
		pickUpItems[v].ItemNum = int32(dropItem.Num())
		pickUpItems[v].Owner = dropItem.OwnerName()
		pickUpItems[v].DropDate = dropItem.DropTime()
		pickUpItems[v].MonsterId = int32(dropItem.DropMonsterId())

		if this.playerPickUp[userId] == nil {
			this.playerPickUp[userId] = make(map[int]int)
		}
		this.playerPickUp[userId][dropItem.ItemId()] += dropItem.Num()
		//真实拾取的时候才移除
		if isPick {
			this.Scene.RemoveSceneObj(dropItemObj)
		}
	}

	//只有拾取到东西，且地上没有东西了时候才会触发拾取完毕
	if len(pickUpItems) > 0 && this.Scene.GetItemObjsNum() == 0 {
		this.fightEvents <- &base.FightEvent{
			EventType: base.FIGHT_EVENT_TYPE_PICK_ALL,
			Data:      actor.GetObjId(),
		}
	}

	return pickUpItems, nil
}

func (this *DefaultFight) Collection(userId int, objId int) (int, error) {

	fitActor := this.GetUserFitActor(userId)
	if fitActor != nil {
		return 0, gamedb.ERRFITCANNOTCOLLECT
	}
	actor := this.GetUserMainActor(userId)
	if actor == nil {
		return 0, gamedb.ERRUNFOUNDUSER
	}
	if !actor.CanMove() {
		return 0, gamedb.ERRDISTANCE
	}
	point := actor.Point()
	if point == nil {
		logger.Error("道具拾取获取玩家坐标失败：%v", actor.GetUserId())
		return 0, gamedb.ERRUSERPOINT
	}

	obj := this.Scene.GetSceneObj(objId).GetContext()
	if obj.GetType() != pb.SCENEOBJTYPE_COLLECTION {
		return 0, gamedb.ERRSCENETYPE
	}
	collectionObj := obj.(*scene.SceneCollection)

	if scene.DistanceByPoint(collectionObj.Point(), point) > 2 {
		logger.Error("玩家距离物品过远，不能拾取:玩家位置:%v，物品位置：%v", point.ToString(), collectionObj.Point().ToString())
		return 0, gamedb.ERRDISTANCE
	}

	if !collectionObj.CanCollection(actor.GetUserId()) {
		return 0, gamedb.ERRCOLLECTION
	}

	collectionObj.Collection(actor.GetObjId())
	//标记玩家采集中
	if u, ok := actor.(base.ActorUser); ok {
		u.SetCollectionId(collectionObj.GetObjId())
	}
	ntf := &pb.CollectionStatusChangeNtf{
		ObjId:     int32(collectionObj.GetObjId()),
		UserObjId: int32(actor.GetObjId()),
		StartTime: time.Now().Unix(),
		EndTime:   collectionObj.GetEndTime(),
	}
	this.Scene.NotifyNearby(actor, ntf, nil)
	return int(collectionObj.GetEndTime()), nil
}

func (this *DefaultFight) ResetCollection(objId int) {
	obj := this.Scene.GetSceneObj(objId)
	if obj == nil {
		return
	}
	if obj.GetType() != pb.SCENEOBJTYPE_COLLECTION {
		return
	}
	collectionObj := obj.GetContext().(*scene.SceneCollection)
	collectionObj.Reset(false)
	ntf := &pb.CollectionStatusChangeNtf{
		ObjId: int32(collectionObj.GetObjId()),
	}
	this.Scene.NotifyNearby(obj, ntf, nil)
}

func (this *DefaultFight) CollectionCancel(userId int, objId int) error {
	obj := this.Scene.GetSceneObj(objId)
	if obj == nil || obj.GetType() != pb.SCENEOBJTYPE_COLLECTION {
		return gamedb.ERRPARAM
	}

	objCollection := this.Scene.GetSceneObj(objId).GetContext().(*scene.SceneCollection)
	mainActor := this.GetUserMainActor(userId)
	if mainActor == nil {
		return gamedb.ERRUNFOUNDUSER
	}
	if !objCollection.CanCancelCollection(mainActor.GetObjId()) {
		return gamedb.ERRCOLLECTION
	}
	mainActor.(base.ActorUser).ResetCollectionStatus()
	return nil
}

func (this *DefaultFight) UseFitReq(userId int, actorFit *pbserver.ActorFit) error {

	mainActor := this.GetUserMainActor(userId)
	if mainActor == nil {
		return gamedb.ERRUNFOUNDUSER
	}
	//刀刀切割状态不能合体
	if has, _ := mainActor.BuffHasType(pb.BUFFTYPE_CUT_SKILL, nil); has {
		return gamedb.ERRFITBYCUTTREASURE
	}

	//allDie := this.CheckUserAllDie(mainActor)
	//if allDie {
	//	return gamedb.ERRPLAYERDIE
	//}

	allActor := this.GetUserByUserId(userId)
	for _, v := range allActor {
		if v.GetProp().HpNow() <= 0 {
			return gamedb.ERRPLAYERDIE
		}
		if has, _ := v.BuffHasType(pb.BUFFTYPE_FIT_LIMIT, nil); has {
			return gamedb.ERRFITLIMIT
		}
	}
	for _, v := range allActor {
		v.ClearFitBuff()
		v.SetVisible(false)
		if v.Job() == pb.JOB_DAOSHI {
			summonActor := v.(base.ActorPlayer).GetPlayer().SummonActors()
			if len(summonActor) > 0 {
				for _, v := range summonActor {
					this.Leave(v)
				}
			}
		}
	}

	//合体进入
	actorFitUnit := actorPkg.NewFitActor(mainActor, actorFit, ai.NewUserFsm)
	actorFitUnit.SetDir(mainActor.GetDir())
	birthPoint := mainActor.Point()
	this.Enter(actorFitUnit, birthPoint)

	//玩家武将离开
	for _, v := range allActor {
		v.(*actorPkg.UserActor).JustLeaveScene()
	}
	return nil
}

func (this *DefaultFight) FitCacelReq(userId int) error {
	playerActor := this.playerActors[userId]
	if playerActor == nil || playerActor.FitActor() == nil {
		return gamedb.ERRUNFOUNDUSER
	}
	fitActor := playerActor.FitActor().(*actorPkg.FitActor)
	fitActor.FitCancel()
	return nil
}

func (this *DefaultFight) UpdatePet(userId int, pet *pbserver.ActorPet) error {

	playerActor := this.playerActors[userId]
	if playerActor == nil {
		return gamedb.ERRUNFOUNDUSER
	}
	if this.CheckUserDieByUserId(userId) {
		return gamedb.ERRPLAYERDIE
	}
	fightPet := playerActor.PetActor()
	if fightPet == nil {
		this.petEnter(userId, pet)
	} else {

		fightPetActor := fightPet.(*actorPkg.PetActor)
		if fightPetActor.PetId == int(pet.PetId) {
			//更新数据
			fightPetActor.UpdateInfo(pet)

		} else {
			this.Leave(fightPet)
			err := this.petEnter(userId, pet)
			return err
		}
	}
	return nil
}

func (this *DefaultFight) petEnter(userId int, pet *pbserver.ActorPet) error {
	if pet == nil || pet.PetId <= 0 {
		return nil
	}
	playerActor := this.playerActors[userId]
	mainActor := playerActor.GetHeroActor(constUser.USER_HERO_MAIN_INDEX)
	userPet := actorPkg.NewPetActor(mainActor, pet, ai.NewUserAI)
	userPet.SetTeamIndex(mainActor.TeamIndex())
	birthPoint := this.getPetEnterScenePoint(userId, userPet)
	err := this.Enter(userPet, birthPoint)
	if err != nil {
		logger.Error("战宠进入战斗错误，玩家：%v,异常：%v", userId, err)
	}
	return nil
}

func (this *DefaultFight) getPetEnterScenePoint(userId int, petActor base.Actor) *scene.Point {
	playerActor := this.playerActors[userId]
	for i := len(constUser.USER_HERO_INDEX) - 1; i >= 0; i-- {
		u := playerActor.GetHeroActor(constUser.USER_HERO_INDEX[i])
		if u != nil {
			birthPoint := this.GetScene().GetPointByPointRange(petActor, u.Point().X(), u.Point().Y(), nil)
			return birthPoint
		}
	}
	return nil
}

func (this *DefaultFight) UpdateElf(userId int, elfInfo *pbserver.ElfInfo) error {
	playerActor := this.playerActors[userId]
	if playerActor == nil {
		return gamedb.ERRUNFOUNDUSER
	}
	var obj base.Actor
	heroActors := playerActor.HeroActors()
	for _, v := range heroActors {
		if obj == nil && v.GetProp().HpNow() > 0 {
			obj = v
		}
		if u, ok := v.(base.ActorUser); ok {
			u.UpdateElf(elfInfo)
		}
	}
	if obj != nil {
		ntf := &pb.SceneUserElfUpdateNtf{
			UserId: int32(userId),
			ElfLv:  int32(elfInfo.Lv),
		}
		this.Scene.NotifyNearby(obj, ntf, nil)
	}
	return nil
}

func (this *DefaultFight) UserRelive(userId, reliveAddrType int) error {
	userActors := this.GetUserByUserId(userId)
	if userActors == nil {
		return gamedb.ERRUNFOUNDUSER
	}
	for _, v := range userActors {
		//设置坐标，更新到场景 通知客户端
		v.Relive(reliveAddrType, constFight.RELIVE_TYPE_COST)
	}
	petActor := this.GetPetActor(userId)
	if petActor != nil {
		point := this.getPetEnterScenePoint(userId, petActor)
		petActor.SetVisible(true)
		petActor.MoveTo(point, pb.MOVETYPE_WALK, true, false)
		appearNtf := petActor.BuildAppearMessage()
		this.GetScene().NotifyNearby(petActor, appearNtf, nil)
	}
	return nil
}

/**
*  @Description: 玩家强制移动
*  @receiver this
*  @param userId
*  @param moveToArea
*  @return error
**/
func (this *DefaultFight) UserMoveFoce(userId int, moveToArea int) error {

	point, err := this.Scene.GetBirthPointByAreaIndex(moveToArea)
	if err != nil {
		return err
	}

	fit := this.GetUserFitActor(userId)
	if fit != nil {
		return fit.MoveTo(point, pb.MOVETYPE_WALK, true, true)
	}

	userActors := this.GetUserByUserId(userId)
	if userActors == nil {
		return gamedb.ERRUNFOUNDUSER
	}
	mainActor := this.GetUserMainActor(userId)
	for _, v := range userActors {
		//设置坐标，更新到场景 通知客户端
		u := v.(base.ActorUser)
		birthPoint := point
		if u.GetHeroIndex() != constUser.USER_HERO_MAIN_INDEX {
			birthPoint = this.GetScene().GetHeroBirthPoint(v, birthPoint, mainActor.GetDir(), u.GetHeroIndex(), constFight.FIGHT_BIRTH_TYPE_TRIANGLE)
		}
		v.MoveTo(birthPoint, pb.MOVETYPE_WALK, true, true)
	}
	petActor := this.GetPetActor(userId)
	if petActor != nil {
		point := this.getPetEnterScenePoint(userId, petActor)
		petActor.MoveTo(point, pb.MOVETYPE_WALK, true, true)
	}

	summonActors := mainActor.(base.ActorPlayer).GetPlayer().SummonActors()
	if len(summonActors) > 0 {
		for _, v := range summonActors {
			this.Leave(v)
		}
	}

	return nil
}

/**
*  @Description: 玩家离开 死亡后协助逻辑
*  @receiver this
*  @param userId 死亡玩家userId
*  @return base.Actor
**/
func (this *DefaultFight) UserDieOrLeaveCheckHelp(userId int) base.Actor {

	//解除协助关系
	for _, v := range this.playerActors {
		if v.ToHelpUserId() == userId {
			v.SetToHelpUserId(0)
		}
	}

	////玩家自己是协助者
	//if p, ok := this.playerActors[userId]; ok {
	//	if p.ToHelpUserId() > 0 {
	//		return nil
	//	}
	//}
	////玩家本身不是协助者
	//ntf := &pb.FightTeamChangeNtf{
	//	UserTeamIndex: make(map[int32]int32),
	//}
	//firstHelper := 0
	//firstHerlpTime := 0
	//for keyUserId, p := range this.playerActors {
	//	if keyUserId == userId {
	//		continue
	//	}
	//	if p.ToHelpUserId() == userId {
	//		if p.FightNum() > 0 {
	//			if firstHelper == 0 || p.ToHelpTime() < int64(firstHerlpTime) {
	//				firstHelper = keyUserId
	//			}
	//		}
	//		ntf.UserTeamIndex[int32(keyUserId)] = int32(keyUserId)
	//		p.ChangeTeam(keyUserId)
	//		p.SetToHelpUserId(0)
	//
	//	}
	//}
	//if len(ntf.UserTeamIndex) > 0 {
	//	this.Scene.NotifyAll(ntf)
	//}
	//if firstHelper > 0 {
	//	return this.GetUserMainActor(firstHelper)
	//}
	return nil
}

//================================控制相关=========================================================

func (this *DefaultFight) RunAI() {
	if this.status != FIGHT_STATUS_CLOSING {
		//TODO 待优化 只有场景有玩家是才开启AI
		//if this.GetScene().GetPlayerObjsNum() > 0 {
		this.Range(func(actor base.Actor) bool { actor.RunAI(); return false })
		//}
	}
}

func (this *DefaultFight) CheckLife() bool {

	now := time.Now().Unix()
	lifeOver := false
	if this.lifeTime > -1 {
		if now-this.createTime >= this.lifeTime {
			//this.Stop()
			if this.status != FIGHT_STATUS_CLOSING {
				logger.Info("Fight %d Life Time Over, fight Type: %d lifeTime:%d", this.GetId(), this.StageConf.Id, this.lifeTime)
				this.GetContext().OnEnd()
				if this.status != FIGHT_STATUS_CLOSING {
					this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 10)
				}
			}
			lifeOver = true
		}
	}

	if this.NextStatusTime != 0 && now > this.NextStatusTime {
		if this.status == FIGHT_STATUS_CLOSING {
			this.Stop()
		} else if this.status == FIGHT_STATUS_READY {
			this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_RUNNING, 0)
		}
	}

	return lifeOver
}

func (this *DefaultFight) DoLoop() chan struct{} {

	done := make(chan struct{})
	go func() {
		rand.Seed(time.Now().UnixNano())
		var ticker = time.NewTicker(100 * time.Millisecond)
		var context = this.GetContext()
		var messages = context.GetMessageChan()
		defer func() {
			ticker.Stop()
			if r := recover(); r != nil {
				stackBytes := debug.Stack()
				logger.Error("panic when DoLoop:%v,%s,%d", r, stackBytes, this.GetId())
			}
		}()
		for {
			select {
			case <-ticker.C:
				if !this.CheckLife() {
					util.SafeRun(context.RunAI)
					util.SafeRun(context.UpdateFrame)
					util.SafeRun(this.privateUpdateFrame)
					util.SafeRun(this.updateCollection)
				}
			case fe := <-this.fightEvents:
				this.fightEventFunc(fe)

			case msg := <-messages:
				msg.Handle()
				if len(this.fightEvents) > 0 {
					this.fightEventFunc(<-this.fightEvents)
				}
			case <-done:
				return
			}
		}
	}()
	return done
}

func (this *DefaultFight) privateUpdateFrame() {
	hasItemDisappeared := this.Scene.UpdateFrame()
	if hasItemDisappeared {
		if this.Scene.GetItemObjsNum() <= 0 {
			this.context.OnPickAll(0)
		}
	}
}

func (this *DefaultFight) fightEventFunc(fe *base.FightEvent) {
	if fe.EventType == base.FIGHT_EVENT_TYPE_PICK_ALL {
		lastPickObjId := fe.Data.(int)
		this.context.OnPickAll(lastPickObjId)
	}
}

func (this *DefaultFight) updateCollection() {
	collections := this.Scene.UpdateFrameCollection()
	for objId, collectionId := range collections {

		actor := this.GetActorByObjId(objId)
		if actor == nil {
			logger.Error("物品采集结束，未找到采集的玩家,玩家：%v,采集物品：%v", objId, collectionId)
			continue
		}
		if actotUser, ok := actor.(base.ActorUser); ok {
			actotUser.SetCollectionId(0)
		}
		collectionConf := gamedb.GetCollectionCollectionCfg(collectionId)
		collectionMsg := &pbserver.FsToGsCollectionNtf{
			UserId:    int32(actor.GetUserId()),
			StageType: int32(this.StageConf.Type),
		}
		if collectionConf.Effect[0] == constFight.COLLECTION_EFFECT_BUFF {
			for i, l := 1, len(collectionConf.Effect); i < l; i++ {
				actor.AddBuff(collectionConf.Effect[i], actor, false)
			}

		} else if collectionConf.Effect[0] == constFight.COLLECTION_EFFECT_ITEM {
			item := map[int32]int32{int32(collectionConf.Effect[1]): int32(collectionConf.Effect[2])}
			collectionMsg.Items = item
		} else if collectionConf.Effect[0] == constFight.COLLECTION_EFFECT_OTHER {
			actors := this.GetUserByUserId(actor.GetUserId())
			for _, v := range actors {
				if v.GetProp().HpNow() <= 0 {
					continue
				}
				for i, l := 1, len(collectionConf.Effect); i < l; i++ {
					v.AddBuff(collectionConf.Effect[i], actor, false)
				}
			}
		}
		net.GetGsConn().SendMessageToGs(uint32(actor.HostId()), collectionMsg)
	}

	this.GetContext().OnCollection(collections)
}

func (this *DefaultFight) Begin() {

}

func (this *DefaultFight) Start() {
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_RUNNING, 0)
	this.loopChan = this.GetContext().DoLoop()
}

func (this *DefaultFight) Stop() {
	this.once.Do(func() {
		if this.loopChan != nil {
			close(this.loopChan)
		}
		if this.Scene != nil {
			this.Scene.Destroy()
		}
	})
	//内存移除战斗
	fightManager.RemoveFight(this.GetId())
}

func (this *DefaultFight) GmReq(req *pbserver.GsToFsGmReq) string {
	mainActor := this.GetUserMainActor(int(req.UserId))
	if mainActor == nil {
		return gamedb.ERRUNFOUNDUSER.Message
	}

	cmdSlice := common.NewStringSlice(req.Cmd, "#")
	cmd := strings.TrimSpace(strings.ToLower(cmdSlice[1]))
	switch cmd {
	case "addbuff":
		buffIds, err := common.IntSliceFromString(cmdSlice[2], "|")
		if err != nil || len(buffIds) <= 0 {
			return gamedb.ERRPARAM.Message
		}
		for _, v := range buffIds {
			_, err := mainActor.AddBuff(v, mainActor, false)
			if err != nil {
				return err.Error()
			}
		}
		return fmt.Sprintf("添加 buff %s 成功", cmdSlice[2])
	case "showfightprop":
		heroIndex := 0
		propIds := make([]int, 0)
		if len(cmdSlice) > 2 {
			str := common.NewStringSlice(cmdSlice[2], "|")
			heroIndex, _ = strconv.Atoi(str[0])
			if len(str) > 1 {
				propIds, _ = common.IntSliceFromString(str[1], ",")
			}
		}
		propMsg := ""
		fit := this.GetUserFitActor(int(req.UserId))
		if fit != nil {
			propStr := fit.GetProp().ToString(propIds)
			propMsg = fmt.Sprintf("昵称:%v,玩家Id:%v，场景Id:%v,详情：%s", fit.NickName(), fit.GetUserId(), fit.GetObjId(), propStr)
		} else {
			isAll := true
			for _, v := range constUser.USER_HERO_INDEX {
				if v == heroIndex {
					isAll = false
				}
			}
			actors := this.GetUserByUserId(int(req.UserId))
			for _, v := range actors {
				if isAll || v.(base.ActorUser).GetHeroIndex() == heroIndex {
					propStr := v.GetProp().ToString(propIds)
					propMsg += fmt.Sprintf("昵称:%v,玩家Id:%v，场景Id:%v,详情：%s \n", v.NickName(), v.GetUserId(), v.GetObjId(), propStr)
				}
			}
		}
		logger.Debug(propMsg)
		return propMsg

	}
	return "未实现GM"
}

func (this *DefaultFight) CanAttack() bool {
	return this.status == FIGHT_STATUS_RUNNING
}

//================================地图相关======================================
func (this *DefaultFight) GetSceneObj(objId int) scene.ISceneObj {
	return this.Scene.GetSceneObj(objId)
}

func (this *DefaultFight) GetSceneAllObj() map[int]scene.ISceneObj {
	return this.Scene.GetSceneAllObj()
}

func (this *DefaultFight) GetScene() *scene.Scene {
	return this.Scene
}

func (this *DefaultFight) CheckInNpcRange(actor base.Actor, npcId int) (bool, error) {

	for _, v := range this.StageConf.Door {
		if v[1] == npcId {
			return this.Scene.CheckInNpcRange(v[0], actor.Point())
		}
	}

	for _, v := range this.StageConf.Npc {
		if v[1] == npcId {
			return this.Scene.CheckInNpcRange(v[0], actor.Point())
		}
	}
	return false, gamedb.ERRPARAM

}

//=================================子类重写========================================

func (this *DefaultFight) OnActorEnter(obj base.Actor) {

	if obj.GetType() == base.ActorTypeMonster || obj.GetType() == pb.SCENEOBJTYPE_USER || obj.GetType() == pb.SCENEOBJTYPE_FIT {
		obj.GetFSM().Recover()
	}
}

func (this *DefaultFight) OnLeave(actor base.Actor) {

}

func (this *DefaultFight) OnEnterUser(userId int) {

}

func (this *DefaultFight) OnLeaveUser(userId int) {

}

func (this *DefaultFight) PostDamage(attacker, defender base.Actor, damage int) {

}

func (this *DefaultFight) OnDie(actor base.Actor, killer base.Actor) {

}

func (this *DefaultFight) CheckEnd() bool {
	return false
}

func (this *DefaultFight) OnEnd() {

}

func (this *DefaultFight) OnRelive(actor base.Actor, reliveType int) {

}

func (this *DefaultFight) UpdateFrame() {

}

func (this *DefaultFight) OnPickAll(lastPickObjId int) {

}

func (this *DefaultFight) OnCollection(collections map[int]int) {

}

func (this *DefaultFight) OnCheer(userId int) {

}

func (this *DefaultFight) OnUsePotion(userId int) {

}

func (this *DefaultFight) GetPowerRoll() string {
	return "0"
}

func (this *DefaultFight) OnBossOwnerChange(monster base.Actor) {

}

func (this *DefaultFight) NpcEventReq(userId, npcId int) error {
	mainActor := this.GetUserMainActor(userId)
	if mainActor == nil {
		return gamedb.ERRUNFOUNDUSER
	}

	firstReliveActor := mainActor.(base.ActorPlayer).GetPlayer().GetFirstReliveHero()
	if firstReliveActor == nil {
		return gamedb.ERRPLAYERDIE
	}

	inRange, err := this.CheckInNpcRange(firstReliveActor, npcId)
	if err != nil {
		return err
	}
	if !inRange {
		return gamedb.ERRDISTANCE
	}
	return nil
}
