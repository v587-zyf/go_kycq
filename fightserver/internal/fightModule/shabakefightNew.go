package fightModule

import (
	"cqserver/fightserver/conf"
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"fmt"
	"time"
)

const (
	SHABAKE_STATUS_INIT = 0 //初始状态
	SHABAKE_STATUS_DOOR = 1 //城门破
	SHABAKE_STATUS_CITY = 2 //内城boss被击杀，开启皇宫传送阵
	SHABAKE_STATUS_HG   = 3 //计算归属
)

type ShabakeFightNew struct {
	*DefaultFight
	*base.FightCheerByGuild
	*base.FightUsePotion
	userScore          *userScores   //玩家积分
	userScoreNow       map[int]int   //玩家当前积分
	userIntoScoreArea  map[int]int64 //玩家进入积分区域时间
	lasttime           int64
	occupiedGuildId    int
	occupiedGuildName  string
	occupiedTime       int64
	occupiedStatus     int //皇宫区域状态（0：初始，1 城门破 2：皇宫传送门开启，3：计算归属）
	huanggongAreaIndex int
	doorBoss           *actorPkg.MonsterActor //城门boss
	cityBoss           *actorPkg.MonsterActor //内城boss
	areaScoreConf      map[int][]int          //区域积分配置(区域索引->{0:间隔秒数，1：获得积分})
	monsterEffect      map[int][]int          //怪物效果(怪物Id->{0:buff，1：获得积分})
}

func NewShabakeFightNew(stageId int) (*ShabakeFightNew, error) {
	var err error
	f := &ShabakeFightNew{
		userScore:         NewUserScores(),
		userScoreNow:      make(map[int]int),
		userIntoScoreArea: make(map[int]int64),
		FightUsePotion:    base.NewFightUsePotion(),
		FightCheerByGuild: base.NewFightCheerByGuild(),
	}
	f.DefaultFight, err = NewDefaultFight(stageId, f)
	if err != nil {
		return nil, err
	}
	err = f.InitMonster()
	if err != nil {
		return nil, err
	}
	f.InitCollection()
	f.Start()
	f.initConf()
	stopTime := gamedb.GetConf().ShabakeTime3[1]
	lifeTime := stopTime.GetSecondsFromZero() - common.GetTimeSeconds(time.Now())
	f.SetLifeTime(int64(lifeTime))
	f.huanggongAreaIndex = gamedb.GetConf().ShabakePalace
	logger.Info("创建沙巴克战斗成功，剩余战斗时间：%v", lifeTime)
	return f, nil
}

func (this *ShabakeFightNew) initConf() {
	this.areaScoreConf = make(map[int][]int)
	for _, v := range gamedb.GetConf().ShabakeScore3 {
		this.areaScoreConf[v[0]] = make([]int, 2)
		this.areaScoreConf[v[0]][0] = v[1]
		this.areaScoreConf[v[0]][1] = v[2]
	}
	this.monsterEffect = make(map[int][]int)
	for _, v := range gamedb.GetConf().ShabakeMonsterBUff {
		this.monsterEffect[v[0]] = make([]int, 2)
		this.monsterEffect[v[0]][0] = v[1]
		this.monsterEffect[v[0]][1] = v[2]
	}
	//设置区域不可走
	this.GetScene().SetSceneBlockByRectIndex(gamedb.GetConf().ShabakeStop, true)

}

func (this *ShabakeFightNew) InitMonster() error {

	monsters, err := this.DefaultFight.InitMonsterByWave2(0)
	if err != nil {
		return err
	}
	this.doorBoss = monsters[0].(*actorPkg.MonsterActor)
	monsters, err = this.DefaultFight.InitMonsterByWave2(1)
	if err != nil {
		return err
	}
	this.cityBoss = monsters[0].(*actorPkg.MonsterActor)
	for i := 2; i < len(this.StageConf.Monster_group); i++ {
		_, err = this.DefaultFight.InitMonsterByWave(i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *ShabakeFightNew) UpdateFrame() {

	if this.status != FIGHT_STATUS_RUNNING {
		return
	}
	now := time.Now()
	if now.Unix()-this.lasttime < 1 {
		return
	}
	this.lasttime = now.Unix()
	//区域积分计算
	this.areaScoreCalc(now)
	//皇宫占领计算
	this.occupiedCalc(now)

}

/**
*  @Description: 区域积分计算
*  @receiver this
*  @param nowTime
**/
func (this *ShabakeFightNew) areaScoreCalc(nowTime time.Time) {
	now := nowTime.Unix()
	objsMap := make(map[int]int) //玩家Id->区域
	for k, _ := range this.areaScoreConf {
		objs := this.Scene.GetSceneRectObjs(k)
		for _, v := range objs {
			obj := this.GetActorByObjId(v)
			if obj != nil && obj.GetUserId() > 0 && obj.GetProp().HpNow() > 0 {
				objsMap[obj.GetUserId()] = k
			}
		}
	}
	hasUser := false
	//记录区域内的玩家，并计算分数
	for userId, aIndex := range objsMap {
		if t, ok := this.userIntoScoreArea[userId]; ok {
			if now-t >= int64(this.areaScoreConf[aIndex][0]) {
				this.addScore(userId, this.areaScoreConf[aIndex][1], false)
				this.userIntoScoreArea[userId] = now
				hasUser = true
			}
		} else {
			this.userIntoScoreArea[userId] = now
		}
	}

	if hasUser {
		this.addScore(0, 0, true)
	}

	//删除未在区域的玩家
	for k, _ := range this.userIntoScoreArea {
		if _, ok := objsMap[k]; !ok {
			delete(this.userIntoScoreArea, k)
		}
	}
}

func (this *ShabakeFightNew) OnDie(actor, killer base.Actor) {

	if killer == nil {
		if conf.Conf.Sandbox {
			panic(fmt.Sprintf("击杀者为空：%v,%v,%v", actor.NickName(), actor.GetObjId(), actor.GetUserId()))
		}
		return
	}

	if actor.GetType() == pb.SCENEOBJTYPE_USER || actor.GetType() == pb.SCENEOBJTYPE_FIT {
		userId := killer.GetUserId()
		if userId > 0 {
			this.addScore(userId, gamedb.GetConf().ShabakeScore2, true)
		}
		mainActor := this.GetUserMainActor(actor.GetUserId())
		if mainActor != nil && mainActor.(base.ActorUser).GuildId() == this.occupiedGuildId {
			allDie := this.CheckUserAllDie(actor)
			if allDie {
				this.occupiedCalcByKill(mainActor, killer)
			}
		}

	} else if actor.GetType() == pb.SCENEOBJTYPE_MONSTER {
		statusChange := false
		if actor.GetObjId() == this.doorBoss.GetObjId() {
			//打破城门 通行
			this.GetScene().SetSceneBlockByRectIndex(gamedb.GetConf().ShabakeStop, false)
			this.occupiedStatus = SHABAKE_STATUS_DOOR
			statusChange = true
			this.broadcastSkillBoss(pb.SCROLINGTYPE_SHABAKE_GATE, killer.GetUserId(), killer.NickName())
		} else if actor.GetObjId() == this.cityBoss.GetObjId() {
			//首次打死内城boss，开启传送阵
			if this.occupiedStatus == SHABAKE_STATUS_DOOR {
				this.occupiedStatus = SHABAKE_STATUS_CITY
				this.occupiedTime = time.Now().Unix()
				statusChange = true
				this.broadcastSkillBoss(pb.SCROLINGTYPE_HUANGCHENG, killer.GetUserId(), killer.NickName())
			}
		}
		//怪物积分
		monsterActor := actor.(*actorPkg.MonsterActor)
		if effect, ok := this.monsterEffect[monsterActor.MonsterT.Monsterid]; ok {
			this.addScore(killer.GetUserId(), effect[1], true)
			//buff 添加
			if effect[0] > 0 {
				this.addGuildBuff(killer.GetUserId(), effect[0], false)
			}
		}
		if statusChange {
			this.Scene.NotifyAll(this.OccupieNtf())
		}
	}
}

func (this *ShabakeFightNew) occupiedCalc(now time.Time) {

	if this.occupiedStatus < SHABAKE_STATUS_CITY {
		return
	} else if this.occupiedStatus == SHABAKE_STATUS_CITY {
		if this.occupiedTime+int64(gamedb.GetConf().ShabakeTime7*60) > now.Unix() {
			return
		} else {
			this.occupiedTime = now.Unix()
			this.occupiedStatus = SHABAKE_STATUS_HG
			this.Scene.NotifyAll(this.OccupieNtf())
			return
		}

	} else if this.occupiedStatus == SHABAKE_STATUS_HG {

		if this.occupiedGuildId > 0 {
			if now.Unix()-this.occupiedTime >= int64(gamedb.GetConf().ShabakeTime6*60) {
				this.OnEnd()
				return
			}
		}
		objs := this.Scene.GetSceneRectObjs(this.huanggongAreaIndex)
		occupiedGuildId := 0
		occupiedGuildName := ""
		for _, v := range objs {
			obj := this.GetActorByObjId(v)
			if obj != nil && obj.GetUserId() > 0 && obj.GetProp().HpNow() > 0 {

				mainActor := this.GetUserMainActor(obj.GetUserId())
				if mainActor != nil {
					if occupiedGuildId == 0 {
						occupiedGuildId = mainActor.(base.ActorUser).GuildId()
						occupiedGuildName = mainActor.(base.ActorUser).GuildName()
					} else {
						if occupiedGuildId != mainActor.(base.ActorUser).GuildId() {
							occupiedGuildId = 0
							break
						}
					}
				}
			}
		}
		if occupiedGuildId != 0 && this.occupiedGuildId != occupiedGuildId {
			logger.Info("新的门派占领皇宫%v--%v", occupiedGuildId, this.occupiedGuildId)
			this.occupiedGuildId = occupiedGuildId
			this.occupiedGuildName = occupiedGuildName
			this.occupiedTime = now.Unix()
			this.Scene.NotifyAll(this.OccupieNtf())
		}
	}
}

func (this *ShabakeFightNew) occupiedCalcByKill(die base.Actor, killer base.Actor) {

	objs := this.Scene.GetSceneRectObjs(this.huanggongAreaIndex)
	//优先判断死亡者是否在皇宫内，不在则不用处理皇宫归属
	dieInHuangGong := false
	for _, v := range objs {
		if v == die.GetObjId() {
			dieInHuangGong = true
			break
		}
	}
	if !dieInHuangGong {
		return
	}
	//判断皇宫内是否还有占领门派活着的成员 有则不用处理皇宫归属
	hasOtherGuildPlayer := false
	for _, v := range objs {
		obj := this.GetActorByObjId(v)
		if obj != nil && obj.GetUserId() > 0 && obj.GetProp().HpNow() > 0 {

			mainActor := this.GetUserMainActor(obj.GetUserId())
			if mainActor != nil && mainActor.(base.ActorUser).GuildId() == this.occupiedGuildId {
				hasOtherGuildPlayer = true
				break
			}
		}
	}
	if hasOtherGuildPlayer {
		return
	}
	//判断击杀者是否再皇宫内（玩家可能因为buff死亡，击杀者已经不在皇宫了），在则成为新的归属
	for _, v := range objs {
		obj := this.GetActorByObjId(v)
		if obj != nil && obj.GetUserId() > 0 && obj.GetProp().HpNow() > 0 {
			if v == killer.GetObjId() {
				mainActor := this.GetUserMainActor(obj.GetUserId())
				this.occupiedGuildId = mainActor.(base.ActorUser).GuildId()
				this.occupiedGuildName = mainActor.(base.ActorUser).GuildName()
				this.occupiedTime = time.Now().Unix()
				this.Scene.NotifyAll(this.OccupieNtf())
				break
			}
		}
	}
}

func (this *ShabakeFightNew) OnActorEnter(actor base.Actor) {

	this.DefaultFight.OnActorEnter(actor)
	this.FightCheerByGuild.OnGuildActorEnter(actor)
}

func (this *ShabakeFightNew) OnEnterUser(userId int) {

	this.addScore(userId, 0, false)
	this.FightCheerByGuild.FightCheerUserInto(userId)
	this.FightUsePotion.FightPotionUserInto(userId)
	ntf := this.OccupieNtf()
	mainActor := this.GetUserMainActor(userId)
	net.GetGateConn().SendMessage(uint32(mainActor.HostId()), mainActor.SessionId(), 0, ntf)
	rankNtf := this.ShabakeScoreRank(true)
	net.GetGateConn().SendMessage(uint32(mainActor.HostId()), mainActor.SessionId(), 0, rankNtf)

}

func (this *ShabakeFightNew) addScore(userId int, score int, notifyClient bool) {

	if userId > 0 {
		actor := this.GetUserMainActor(userId)
		if _, ok := this.userScore.userScoreMap[userId]; !ok {
			this.userScore.AddUser(userId, actor.NickName())
		}
		this.userScore.addScore(userId, score)
		this.userScoreNow[userId] += score
		this.sendUserScoreChange(userId, score)
	}
	//通知客户端积分变化,新排名
	if notifyClient {
		ntf := this.ShabakeScoreRank(true)
		this.Scene.NotifyAll(ntf)
	}
}

/**
*  @Description: 推送玩家积分变化
*  @receiver this
*  @param userId
*  @param changeNum
**/
func (this *ShabakeFightNew) sendUserScoreChange(userId int, changeNum int) {

	mainActor := this.GetUserMainActor(userId)
	if mainActor == nil {
		return
	}
	ntf := &pb.FightUserScoreNtf{
		Score:       int32(this.userScoreNow[userId]),
		ChangeScore: int32(changeNum),
		RankScore:   int32(this.userScore.userScoreMap[userId].Score),
	}
	net.GetGateConn().SendMessage(uint32(mainActor.HostId()), mainActor.SessionId(), 0, ntf)

}

func (this *ShabakeFightNew) OnEnd() {

	msg := &pbserver.FSFightEndNtf{
		FightType: int32(this.StageConf.Type),
		StageId:   int32(this.StageConf.Id),
		UseTime:   int32(time.Now().Unix() - this.createTime),
	}

	endRank := this.ShabakeScoreRank(false)
	endMsg := &pbserver.ShabakeFightEndNtf{
		UserRank:  make([]*pbserver.ShabakeRankScore, len(endRank.UserScores)),
		GuildRank: make([]*pbserver.ShabakeRankScore, 0),
	}
	for k, v := range endRank.UserScores {
		endMsg.UserRank[k] = &pbserver.ShabakeRankScore{Id: v.UserId, Score: v.Score}
	}
	if this.occupiedGuildId > 0 {
		endMsg.GuildRank = append(endMsg.GuildRank, &pbserver.ShabakeRankScore{int32(this.occupiedGuildId), 1})
	}

	rb, _ := endMsg.Marshal()
	msg.CpData = rb
	logger.Info("发送game沙巴克战斗结束,服务器：%v，结果：%v", *msg)
	net.GetGsConn().SendMessage(msg)
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 15)
}

func (this *ShabakeFightNew) ShabakeScoreRank(withoutZeroScore bool) *pb.ShabakeScoreRankNtf {

	this.userScore.Sort()
	ntf := &pb.ShabakeScoreRankNtf{
		UserScores: this.userScore.rank(withoutZeroScore),
	}
	return ntf
}

func (this *ShabakeFightNew) OnCheer(userId int) {

	this.addGuildBuff(userId, gamedb.GetConf().ShabakeBuff[3], true)
}

/**
*  @Description: 增加门派buff
*  @receiver this
*  @param userId
*  @param buffId
*  @param fromCheer
**/
func (this *ShabakeFightNew) addGuildBuff(userId int, buffId int, fromCheer bool) {

	userMainActor := this.GetUserMainActor(userId)
	if userMainActor == nil {
		logger.Error("玩家发送来鼓舞，鼓舞玩家信息未找到：%v", userId)
		return
	}
	guildId := userMainActor.(base.ActorUser).GuildId()
	//记录门派鼓舞数据
	if fromCheer {
		this.FightCheerByGuild.GuildCheer(this, guildId, buffId)
	} else {
		this.FightCheerByGuild.GuildCheerBuff(guildId, buffId)
	}

	actors := this.GetUserActors()
	for _, v := range actors {
		if u, ok := v.(base.ActorUser); ok {
			if u.GuildId() == guildId {
				v.AddBuff(buffId, userMainActor, false)
			}
		}
	}
}

func (this *ShabakeFightNew) OnUsePotion(userId int) {

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

/**
*  @Description: 采集回调
*  @receiver this
*  @param colllection
**/
func (this *ShabakeFightNew) OnCollection(colllection map[int]int) {

	for k, v := range colllection {

		conf := gamedb.GetCollectionCollectionCfg(v)
		if len(conf.Effect) > 1 && conf.Effect[0] == constFight.COLLECTION_EFFECT_SCORE {
			actor := this.GetActorByObjId(k)
			if actor.GetUserId() > 0 {
				this.addScore(actor.GetUserId(), conf.Effect[1], true)
			}
		}
	}
}

func (this *ShabakeFightNew) OccupieNtf() *pb.ShabakeOccupiedNtf {
	ntf := &pb.ShabakeOccupiedNtf{}
	ntf.FightStatus = int32(this.occupiedStatus)
	if this.occupiedGuildId > 0 {
		ntf.IsOccupy = true
		ntf.GuildId = int32(this.occupiedGuildId)
		ntf.GuildName = this.occupiedGuildName
		ntf.StartTime = int32(this.occupiedTime)
		ntf.EndTime = int32(this.occupiedTime + int64(gamedb.GetConf().ShabakeTime6*60))

	} else {
		ntf.IsOccupy = false
	}
	return ntf
}

func (this *ShabakeFightNew) NpcEventReq(userId, npcId int) error {

	err := this.DefaultFight.NpcEventReq(userId, npcId)
	if err != nil {
		return err
	}

	if npcId == constFight.FIGHT_SHABAKE_NPC_ID_ONE {
		if this.occupiedStatus < SHABAKE_STATUS_CITY {
			return gamedb.ERRPARAM
		}
		//扣积分传送
		return this.npcByScore(userId, gamedb.GetConf().ShabakeNPC1)
	} else if npcId == constFight.FIGHT_SHABAKE_NPC_ID_TWO {
		if this.occupiedStatus < SHABAKE_STATUS_DOOR {
			return gamedb.ERRPARAM
		}
		//扣积分传送
		return this.npcByScore(userId, gamedb.GetConf().ShabakeNPC3)
	} else if npcId == constFight.FIGHT_SHABAKE_NPC_ID_THREE {
		//治疗
		if this.userScoreNow[userId] >= gamedb.GetConf().ShabakeNPC2[1] {
			this.userScoreNow[userId] -= gamedb.GetConf().ShabakeNPC2[1]
			this.sendUserScoreChange(userId, -gamedb.GetConf().ShabakeNPC2[1])
			heros := this.GetUserByUserId(userId)
			for _, v := range heros {
				if v.GetProp().HpNow() <= 0 {
					continue
				}
				temChangeHp, _ := v.ChangeHp(int(float64(v.GetProp().Get(pb.PROPERTY_HP)) * float64(gamedb.GetConf().ShabakeNPC2[2]) / base.BASE_RATE))
				HPChangeNtf := &pb.SceneObjHpNtf{
					ObjId:    int32(v.GetObjId()),
					Hp:       int64(v.GetProp().HpNow()),
					ChangeHp: int64(temChangeHp),
					TotalHp:  int64(v.GetProp().Get(pb.PROPERTY_HP)),
				}
				v.NotifyNearby(v, HPChangeNtf, nil)
			}
		} else {
			return gamedb.MARKNOTENOUGH
		}
	}
	return nil
}

func (this *ShabakeFightNew) npcByScore(userId int, npcSetting []int) error {
	//扣积分传送
	if this.userScoreNow[userId] >= npcSetting[1] {
		if npcSetting[1] > 0 {
			this.userScoreNow[userId] -= npcSetting[1]
			this.sendUserScoreChange(userId, -npcSetting[1])
		}
		return this.UserMoveFoce(userId, npcSetting[2])
	} else {
		return gamedb.MARKNOTENOUGH
	}
}

func (this *ShabakeFightNew) broadcastSkillBoss(sendType, userId int, nickName string) {

	msg := &pbserver.FsToGsShabakeKillBossNtf{}
	cfg := gamedb.GetScrollingScrollingCfg(sendType)
	if cfg == nil {
		return
	}
	monsterName := gamedb.GetMonsterMonsterCfg(cfg.Condition).Name
	content := fmt.Sprintf(cfg.Txt, userId, nickName, monsterName)
	msg.Infos = content
	net.GetGsConn().SendMessage(msg)
}
