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
	"time"
)

type ShabakeCrossFight struct {
	*DefaultFight
	*base.FightCheerByGuild
	*base.FightUsePotion
	guildScore        *guildScores  //门派积分
	userIntoScoreArea map[int]int64 //玩家进入积分区域时间
	lasttime          int64
	occupiedUserId    int
	occupiedTime      int64
}

func NewShabakeCrossFight(stageId int) (*ShabakeCrossFight, error) {
	var err error
	f := &ShabakeCrossFight{
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

	stopTime := gamedb.GetConf().KuafushabakeTime3[1]
	lifeTime := stopTime.GetSecondsFromZero() - common.GetTimeSeconds(time.Now())
	if lifeTime < 10 {
		return nil, gamedb.ERRFIGHTID
	}
	//lifeTime := 60 * 60
	f.SetLifeTime(int64(lifeTime))
	logger.Info("创建跨服沙巴克战斗成功，剩余战斗时间：%v", lifeTime)
	return f, nil
}

func (this *ShabakeCrossFight) UpdateFrame() {

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
			if now-t >= int64(gamedb.GetConf().KuafushabakeScore[0]) {
				this.addScore(k, gamedb.GetConf().KuafushabakeScore[1])
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

	if this.occupiedUserId > 0 && now-this.occupiedTime >= int64(gamedb.GetConf().KuafushabakeTime4*60) {
		this.OnEnd()
	}
}

func (this *ShabakeCrossFight) OnDie(actor, killer base.Actor) {

	if killer == nil {
		if conf.Conf.Sandbox {
			panic(fmt.Sprintf("击杀者为空：%v,%v,%v", actor.NickName(), actor.GetObjId(), actor.GetUserId()))
		}
		return
	}

	if actor.GetType() == pb.SCENEOBJTYPE_USER || actor.GetType() == pb.SCENEOBJTYPE_FIT {
		userId := killer.GetUserId()
		if userId > 0 {
			this.addScore(userId, gamedb.GetConf().KuafushabakeScore[2])
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

func (this *ShabakeCrossFight) AddCollectionFlag() {

	ntf := &pb.ShabakeCrossOccupiedNtf{
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

func (this *ShabakeCrossFight) OnActorEnter(actor base.Actor) {

	this.DefaultFight.OnActorEnter(actor)
	this.FightCheerByGuild.OnGuildActorEnter(actor)
}

func (this *ShabakeCrossFight) OnLeaveUser(userId int) {

	if this.occupiedUserId == userId {
		this.occupiedUserId = 0
		this.occupiedTime = 0
		this.AddCollectionFlag()
	}
}

func (this *ShabakeCrossFight) OnEnterUser(userId int) {

	this.addScore(userId, 0)
	this.FightCheerByGuild.FightCheerUserInto(userId)
	this.FightUsePotion.FightPotionUserInto(userId)

	ntf := this.OccupieNtf()
	mainActor := this.GetUserMainActor(userId)
	net.GetGateConn().SendMessage(uint32(mainActor.HostId()), mainActor.SessionId(), 0, ntf)

}

func (this *ShabakeCrossFight) addScore(userId int, score int) {

	actor := this.GetUserMainActor(userId)
	if user, ok := actor.(base.ActorUser); ok {
		guildId := user.GuildId()
		if _, ok := this.guildScore.guildScoreMap[guildId]; !ok {
			this.guildScore.guildScoreMap[guildId] = &pb.ShabakeGuildScore{GuildId: int32(guildId), GuildName: fmt.Sprintf("s%v.%v", actor.HostId(), user.GuildName()), Score: 0, ServerId: int32(actor.HostId())}
		}
		this.guildScore.guildScoreMap[guildId].Score += int32(score)
		logger.Debug("沙巴克跨服积分增加，玩家：%v,门派：%v,积分：%v", userId, guildId, score)
	}
	//通知客户端积分变化,新排名
	ntf := this.ShabakeScoreRank(true)
	this.Scene.NotifyAll(ntf)
}

func (this *ShabakeCrossFight) OnEnd() {

	guildUser := make(map[int][]int32)
	allUser := this.GetPlayerUserids()
	for _, v := range allUser {
		mainActor := this.GetUserMainActor(v)
		if mainActor != nil {
			if u, ok := mainActor.(base.ActorUser); ok {
				guildId := u.GuildId()
				if guildUser[guildId] == nil {
					guildUser[guildId] = make([]int32, 0)
				}
				guildUser[guildId] = append(guildUser[guildId], int32(v))
			}
		}
	}
	msg := &pbserver.FSFightEndNtf{
		FightType: int32(this.StageConf.Type),
		StageId:   int32(this.StageConf.Id),
		UseTime:   int32(time.Now().Unix() - this.createTime),
	}

	endRank := this.ShabakeScoreRank(false)
	endMsg := &pbserver.ShabakeCrossFightEndNtf{
		ServerRank: make([]*pbserver.ShabakeRankScore, len(endRank.ServerScores)),
		GuildRank:  make([]*pbserver.ShabakeCrossRankScore, len(endRank.GuildScores)),
	}
	for k, v := range endRank.ServerScores {
		endMsg.ServerRank[k] = &pbserver.ShabakeRankScore{Id: v.ServerId, Score: v.Score}
	}
	for k, v := range endRank.GuildScores {
		endMsg.GuildRank[k] = &pbserver.ShabakeCrossRankScore{GuildId: v.GuildId, ServerId: v.ServerId, Score: v.Score}
		if guildUser[int(v.GuildId)] != nil {
			endMsg.GuildRank[k].Users = guildUser[int(v.GuildId)]
		}
	}
	rb, _ := endMsg.Marshal()
	msg.CpData = rb
	logger.Info("发送game跨服沙巴克战斗结束,服务器：%v，结果：%v", *msg)
	net.GetGsConn().SendMessage(msg)
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 15)
}

func (this *ShabakeCrossFight) ShabakeScoreRank(withoutZeroScore bool) *pb.ShabakeCrossScoreRankNtf {

	this.guildScore.Sort()
	ntf := &pb.ShabakeCrossScoreRankNtf{
		ServerScores: this.getServerRank(),
		GuildScores:  this.guildScore.rank(withoutZeroScore),
	}
	return ntf
}

func (this *ShabakeCrossFight) getServerRank() []*pb.ShabakeCrossServerScore {

	serverScore := make(map[int]int)
	for _, v := range this.guildScore.guildScoreMap {
		serverScore[int(v.ServerId)] += int(v.Score)
	}
	sortData := common.SortKvIntMapDes(serverScore)
	serverRank := make([]*pb.ShabakeCrossServerScore, len(sortData))
	for k, v := range sortData {
		serverRank[k] = &pb.ShabakeCrossServerScore{
			ServerId:   int32(v.K),
			ServerName: fightManager.GetServerName(v.K),
			Score:      int32(v.V),
		}
	}
	if this.occupiedUserId > 0 {
		mainActor := this.GetUserMainActor(this.occupiedUserId)
		newServerRank := make([]*pb.ShabakeCrossServerScore, len(serverRank))
		i := 1
		for _, v := range serverRank {
			if int(v.ServerId) == mainActor.HostId() {
				newServerRank[0] = v
			} else {
				newServerRank[i] = v
				i++
			}
		}
		return newServerRank
	}
	return serverRank
}

func (this *ShabakeCrossFight) OnCheer(userId int) {

	userMainActor := this.GetUserMainActor(userId)
	if userMainActor == nil {
		logger.Error("玩家发送来鼓舞，鼓舞玩家信息未找到：%v", userId)
		return
	}
	guildId := userMainActor.(base.ActorUser).GuildId()
	buffId := gamedb.GetConf().KuafushabakeBuff[3]
	//记录门派鼓舞数据
	this.FightCheerByGuild.GuildCheer(this, guildId, buffId)
	actors := this.GetUserActors()
	for _, v := range actors {
		if u, ok := v.(base.ActorUser); ok {
			if u.GuildId() == guildId {
				v.AddBuff(gamedb.GetConf().KuafushabakeBuff[3], userMainActor, false)
			}
		}
	}
}

func (this *ShabakeCrossFight) OnUsePotion(userId int) {

	userActors := this.GetUserByUserId(userId)
	if userActors == nil {
		logger.Error("玩家发送来药水使用，玩家信息未找到：%v", userId)
		return
	}

	for _, v := range userActors {

		if v.GetProp().HpNow() <= 0 {
			continue
		}
		//推送血量变化
		changeHp, _ := v.ChangeHp(int(float64(gamedb.GetConf().KuafushabakePotion[2]) / 100 * float64(v.GetProp().Get(pb.PROPERTY_HP))))
		HPChangeNtf := &pb.SceneObjHpNtf{
			ObjId:    int32(v.GetObjId()),
			Hp:       int64(v.GetProp().HpNow()),
			ChangeHp: int64(changeHp),
			TotalHp:  int64(v.GetProp().Get(pb.PROPERTY_HP)),
		}
		v.NotifyNearby(v, HPChangeNtf, nil)
	}
}

func (this *ShabakeCrossFight) OnCollection(colllection map[int]int) {

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

func (this *ShabakeCrossFight) OccupieNtf() *pb.ShabakeCrossOccupiedNtf {
	ntf := &pb.ShabakeCrossOccupiedNtf{}
	if this.occupiedUserId > 0 {
		mainActor := this.GetUserMainActor(this.occupiedUserId)
		if mainActor != nil {
			ntf.IsOccupy = true
			ntf.UserId = int32(this.occupiedUserId)
			ntf.UserName = mainActor.NickName()
			ntf.GuildId = int32(mainActor.(base.ActorUser).GuildId())
			ntf.GuildName = mainActor.(base.ActorUser).GuildName()
			ntf.StartTime = int32(this.occupiedTime)
			ntf.EndTime = int32(this.occupiedTime + int64(gamedb.GetConf().KuafushabakeTime4*60))
			ntf.ServerId = int32(mainActor.HostId())
			ntf.ServerName = fightManager.GetServerName(mainActor.HostId())
		}
	} else {
		ntf.IsOccupy = false
	}
	return ntf
}
