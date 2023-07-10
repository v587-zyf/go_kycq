package fightModule

import (
	"cqserver/fightserver/conf"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type userScores struct {
	userScoreMap     map[int]*pb.ShabakeUserScore
	userScoreAddTime map[int]int64
	users            []int
}

func NewUserScores() *userScores {
	return &userScores{
		userScoreMap:     make(map[int]*pb.ShabakeUserScore),
		userScoreAddTime: make(map[int]int64),
	}
}

func (this *userScores) Len() int {
	return len(this.userScoreMap)
}

func (this *userScores) Less(i, j int) bool {
	if this.userScoreMap[this.users[i]].Score > this.userScoreMap[this.users[j]].Score {
		return true
	} else if this.userScoreMap[this.users[i]].Score < this.userScoreMap[this.users[j]].Score {
		return false
	} else {
		return this.userScoreAddTime[this.users[i]] < this.userScoreAddTime[this.users[j]]
	}
}

func (this *userScores) Swap(i, j int) {
	this.users[i], this.users[j] = this.users[j], this.users[i]
}

func (this *userScores) Sort() []int {
	this.users = make([]int, this.Len())
	i := 0
	for key, _ := range this.userScoreMap {
		this.users[i] = key
		i++
	}
	sort.Sort(this)
	return this.users
}

func (this *userScores) AddUser(userId int, name string) {
	if _, ok := this.userScoreMap[userId]; !ok {
		this.userScoreMap[userId] = &pb.ShabakeUserScore{UserId: int32(userId), UserName: name, Score: 0}
	}
}

func (this *userScores) addScore(userId int, addScore int) {

	if _, ok := this.userScoreMap[userId]; !ok {
		this.userScoreMap[userId] = &pb.ShabakeUserScore{UserId: int32(userId), UserName: "", Score: 0}
	}

	this.userScoreMap[userId].Score += int32(addScore)
	this.userScoreAddTime[userId] = time.Now().UnixNano()

}

func (this *userScores) rank(withoutZeroScore bool) []*pb.ShabakeUserScore {
	rank := make([]*pb.ShabakeUserScore, 0)
	for k, v := range this.users {
		if !withoutZeroScore || this.userScoreMap[v].Score > 0 {
			this.userScoreMap[v].Rank = int32(k + 1)
			rank = append(rank, this.userScoreMap[v])
		}
	}
	return rank
}

type guildScores struct {
	guildScoreMap map[int]*pb.ShabakeGuildScore
	guilds        []int
}

func (this *guildScores) Len() int {
	return len(this.guildScoreMap)
}

func (this *guildScores) Less(i, j int) bool {
	return this.guildScoreMap[this.guilds[i]].Score > this.guildScoreMap[this.guilds[j]].Score
}

func (this *guildScores) Swap(i, j int) {
	this.guilds[i], this.guilds[j] = this.guilds[j], this.guilds[i]
}

func (this *guildScores) Sort() []int {
	this.guilds = make([]int, this.Len())
	i := 0
	for key, _ := range this.guildScoreMap {
		this.guilds[i] = key
		i++
	}
	sort.Sort(this)
	return this.guilds
}

func (this *guildScores) rank(withoutZeroScore bool) []*pb.ShabakeGuildScore {
	rank := make([]*pb.ShabakeGuildScore, 0)
	for _, v := range this.guilds {
		if !withoutZeroScore || this.guildScoreMap[v].Score > 0 {
			rank = append(rank, this.guildScoreMap[v])
		}
	}
	return rank
}

type ShabakeFight struct {
	*DefaultFight
	*base.FightCheerByGuild
	*base.FightUsePotion
	userScore         *userScores   //玩家积分
	guildScore        *guildScores  //门派积分
	userIntoScoreArea map[int]int64 //玩家进入积分区域时间
	lasttime          int64
	occupiedUserId    int
	occupiedTime      int64
}

func NewShabakeFight(stageId int) (*ShabakeFight, error) {
	var err error
	f := &ShabakeFight{
		userScore:         NewUserScores(),
		guildScore:        &guildScores{guildScoreMap: make(map[int]*pb.ShabakeGuildScore)},
		userIntoScoreArea: make(map[int]int64),
		FightUsePotion:    base.NewFightUsePotion(),
		FightCheerByGuild: base.NewFightCheerByGuild(),
	}
	f.DefaultFight, err = NewDefaultFight(stageId, f)
	if err != nil {
		return nil, err
	}
	f.InitMonster()
	f.InitCollection()
	f.Start()

	stopTime := gamedb.GetConf().ShabakeTime3[1]
	lifeTime := stopTime.GetSecondsFromZero() - common.GetTimeSeconds(time.Now())
	f.SetLifeTime(int64(lifeTime))
	logger.Info("创建沙巴克战斗成功，剩余战斗时间：%v", lifeTime)
	return f, nil
}

func (this *ShabakeFight) UpdateFrame() {

	if this.status != FIGHT_STATUS_RUNNING {
		return
	}
	now := time.Now().Unix()
	if now-this.lasttime < 1 {
		return
	}
	this.lasttime = now

	mapConf := gamedb.GetMapMapCfg(this.StageConf.Mapid)
	objsMap := make(map[int]bool)
	for _, v := range mapConf.Special {
		objs := this.Scene.GetSceneRectObjs(v)
		for _, v := range objs {
			obj := this.GetActorByObjId(v)
			if obj != nil && obj.GetUserId() > 0 && obj.GetProp().HpNow() > 0 {
				objsMap[obj.GetUserId()] = true
			}
		}
	}

	//记录区域内的玩家，并计算分数
	for k, _ := range objsMap {
		if t, ok := this.userIntoScoreArea[k]; ok {
			if now-t >= int64(gamedb.GetConf().ShabakeScore[0]) {
				this.addScore(k, gamedb.GetConf().ShabakeScore[1])
				this.userIntoScoreArea[k] = now
			}
		} else {
			this.userIntoScoreArea[k] = now
		}
	}

	//删除未在区域的玩家
	for k, _ := range this.userIntoScoreArea {
		if _, ok := objsMap[k]; !ok {
			delete(this.userIntoScoreArea, k)
		}
	}

	if this.occupiedUserId > 0 && now-this.occupiedTime >= int64(gamedb.GetConf().ShabakeTime4*60) {
		this.OnEnd()
	}
}

func (this *ShabakeFight) OnDie(actor, killer base.Actor) {

	if killer == nil {
		if conf.Conf.Sandbox {
			panic(fmt.Sprintf("击杀者为空：%v,%v,%v", actor.NickName(), actor.GetObjId(), actor.GetUserId()))
		}
		return
	}

	if actor.GetType() == pb.SCENEOBJTYPE_USER || actor.GetType() == pb.SCENEOBJTYPE_FIT {
		userId := killer.GetUserId()
		if userId > 0 {
			this.addScore(userId, gamedb.GetConf().ShabakeScore[2])
		}
		if actor.GetUserId() == this.occupiedUserId {
			allDie := this.CheckUserAllDie(actor)
			if allDie {
				this.occupiedUserId = 0
				this.occupiedTime = 0
				this.AddCollectionFlag()
			}
		}
	}
}

func (this *ShabakeFight) AddCollectionFlag() {

	ntf := &pb.ShabakeOccupiedNtf{
		IsOccupy: false,
	}
	this.Scene.NotifyAll(ntf)

	for _, v := range this.StageConf.Collection {
		collectionConf := gamedb.GetCollectionCollectionCfg(v)
		if collectionConf.Type == constFight.COLLECTION_TYPE_THREE {

			birthArea := collectionConf.Collection[rand.Intn(len(collectionConf.Collection))]
			birthPoint, err := this.Scene.GetBirthPointByAreaIndex(birthArea)
			if err != nil {
				logger.Error("获取沙巴克旗帜位置异常")
				return
			}
			sceneItem := scene.NewSceneCollection(collectionConf.Id)
			err1 := sceneItem.EnterScene(this.Scene, birthPoint)
			if err1 != nil {
				logger.Error("添加沙巴克旗帜到场景异常")
				return
			}
			break
		}
	}
}

func (this *ShabakeFight) OnActorEnter(actor base.Actor) {

	this.DefaultFight.OnActorEnter(actor)
	this.FightCheerByGuild.OnGuildActorEnter(actor)
}

func (this *ShabakeFight) OnLeaveUser(userId int) {

	if this.occupiedUserId == userId {
		this.occupiedUserId = 0
		this.occupiedTime = 0
		this.AddCollectionFlag()
	}
}

func (this *ShabakeFight) OnEnterUser(userId int) {

	this.addScore(userId, 0)
	this.FightCheerByGuild.FightCheerUserInto(userId)
	this.FightUsePotion.FightPotionUserInto(userId)
	ntf := this.OccupieNtf()
	mainActor := this.GetUserMainActor(userId)
	net.GetGateConn().SendMessage(uint32(mainActor.HostId()), mainActor.SessionId(), 0, ntf)

}

func (this *ShabakeFight) addScore(userId int, score int) {

	actor := this.GetUserMainActor(userId)
	if _, ok := this.userScore.userScoreMap[userId]; !ok {
		this.userScore.AddUser(userId, actor.NickName())
	}
	this.userScore.addScore(userId, score)
	if user, ok := actor.(base.ActorUser); ok {
		guildId := user.GuildId()
		if _, ok := this.guildScore.guildScoreMap[guildId]; !ok {
			this.guildScore.guildScoreMap[guildId] = &pb.ShabakeGuildScore{GuildId: int32(guildId), GuildName: user.GuildName(), Score: 0}
		}
		this.guildScore.guildScoreMap[guildId].Score += int32(score)
		logger.Debug("沙巴克积分增加，玩家：%v,门派：%v,积分：%v", userId, guildId, score)
	}
	//通知客户端积分变化,新排名
	ntf := this.ShabakeScoreRank(true)
	this.Scene.NotifyAll(ntf)
}

func (this *ShabakeFight) OnEnd() {

	msg := &pbserver.FSFightEndNtf{
		FightType: int32(this.StageConf.Type),
		StageId:   int32(this.StageConf.Id),
		UseTime:   int32(time.Now().Unix() - this.createTime),
	}

	endRank := this.ShabakeScoreRank(false)
	endMsg := &pbserver.ShabakeFightEndNtf{
		UserRank:  make([]*pbserver.ShabakeRankScore, len(endRank.UserScores)),
		GuildRank: make([]*pbserver.ShabakeRankScore, len(endRank.GuildScores)),
	}
	for k, v := range endRank.UserScores {
		endMsg.UserRank[k] = &pbserver.ShabakeRankScore{Id: v.UserId, Score: v.Score}
	}
	for k, v := range endRank.GuildScores {
		endMsg.GuildRank[k] = &pbserver.ShabakeRankScore{Id: v.GuildId, Score: v.Score}
	}
	rb, _ := endMsg.Marshal()
	msg.CpData = rb
	logger.Info("发送game沙巴克战斗结束,服务器：%v，结果：%v", *msg)
	net.GetGsConn().SendMessage(msg)
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 15)
}

func (this *ShabakeFight) ShabakeScoreRank(withoutZeroScore bool) *pb.ShabakeScoreRankNtf {

	this.userScore.Sort()
	this.guildScore.Sort()
	ntf := &pb.ShabakeScoreRankNtf{
		UserScores:  this.userScore.rank(withoutZeroScore),
		GuildScores: this.guildRank(withoutZeroScore),
	}
	return ntf
}

func (this *ShabakeFight) guildRank(withoutZeroScore bool) []*pb.ShabakeGuildScore {

	guildRank := this.guildScore.rank(withoutZeroScore)
	if this.occupiedUserId > 0 {
		mainActor := this.GetUserMainActor(this.occupiedUserId)
		guildId := mainActor.(base.ActorUser).GuildId()
		newGuildRank := make([]*pb.ShabakeGuildScore, len(guildRank))
		i := 1
		for _, v := range guildRank {
			if int(v.GuildId) == guildId {
				newGuildRank[0] = v
			} else {
				newGuildRank[i] = v
				i++
			}
		}
		return newGuildRank
	} else {
		return guildRank
	}
}

func (this *ShabakeFight) OnCheer(userId int) {

	userMainActor := this.GetUserMainActor(userId)
	if userMainActor == nil {
		logger.Error("玩家发送来鼓舞，鼓舞玩家信息未找到：%v", userId)
		return
	}
	guildId := userMainActor.(base.ActorUser).GuildId()
	buffId := gamedb.GetConf().ShabakeBuff[3]
	//记录门派鼓舞数据
	this.FightCheerByGuild.GuildCheer(this, guildId, buffId)
	actors := this.GetUserActors()
	for _, v := range actors {
		if u, ok := v.(base.ActorUser); ok {
			if u.GuildId() == guildId {
				v.AddBuff(gamedb.GetConf().ShabakeBuff[3], userMainActor, false)
			}
		}
	}
}

func (this *ShabakeFight) OnUsePotion(userId int) {

	userActors := this.GetUserByUserId(userId)
	if userActors == nil {
		logger.Error("玩家发送来药水使用，玩家信息未找到：%v", userId)
		return
	}

	for _, v := range userActors {
		//推送血量变化
		if v.GetProp().HpNow() <= 0 {
			continue
		}
		changeHp, _ := v.ChangeHp(int(float64(gamedb.GetConf().ShabakePotion[2]) / 100 * float64(v.GetProp().Get(pb.PROPERTY_HP))))
		HPChangeNtf := &pb.SceneObjHpNtf{
			ObjId:    int32(v.GetObjId()),
			Hp:       int64(v.GetProp().HpNow()),
			ChangeHp: int64(changeHp),
			TotalHp:  int64(v.GetProp().Get(pb.PROPERTY_HP)),
		}
		v.NotifyNearby(v, HPChangeNtf, nil)
	}
}

func (this *ShabakeFight) OnCollection(colllection map[int]int) {

	for k, v := range colllection {
		conf := gamedb.GetCollectionCollectionCfg(v)
		if conf.Type == constFight.COLLECTION_TYPE_THREE {

			actor := this.GetActorByObjId(k)
			this.occupiedUserId = actor.GetUserId()
			this.occupiedTime = time.Now().Unix()
			ntf := this.OccupieNtf()
			this.Scene.NotifyAll(ntf)
		}
	}
}

func (this *ShabakeFight) OccupieNtf() *pb.ShabakeOccupiedNtf {
	ntf := &pb.ShabakeOccupiedNtf{}
	if this.occupiedUserId > 0 {
		mainActor := this.GetUserMainActor(this.occupiedUserId)
		if mainActor != nil {
			ntf.IsOccupy = true
			ntf.UserId = int32(this.occupiedUserId)
			ntf.UserName = mainActor.NickName()
			ntf.GuildId = int32(mainActor.(base.ActorUser).GuildId())
			ntf.GuildName = mainActor.(base.ActorUser).GuildName()
			ntf.StartTime = int32(this.occupiedTime)
			ntf.EndTime = int32(this.occupiedTime + int64(gamedb.GetConf().ShabakeTime4*60))
		}
	} else {
		ntf.IsOccupy = false
	}
	return ntf
}
