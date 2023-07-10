package actorPkg

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/errex"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"strconv"
	"time"
)

type FitActor struct {
	*DefaultActor
	leader         base.Actor //归属玩家
	birthAreaIndex int        //出生区域索引（对应stage表monster_group下角标）
	guildId        int
	guildName      string
	fit            *pbserver.ActorFit
	createTime     int64
	heroHpFix      map[int]float64 //武将血量占比
	leftTime       int
}

func NewFitActor(owner base.Actor, fit *pbserver.ActorFit, aiCreator base.AICreator) *FitActor {
	actor := &FitActor{
		leader:     owner,
		fit:        fit,
		guildId:    owner.(base.ActorUser).GuildId(),
		guildName:  owner.(base.ActorUser).GuildName(),
		createTime: time.Now().Unix(),
	}
	actor.DefaultActor = NewDefaultActor(pb.SCENEOBJTYPE_FIT, owner.NickName(), "", pb.JOB_ZHANSHI, &pbserver.ActorDisplayInfo{}, aiCreator(actor), actor)
	actor.teamIndex = owner.TeamIndex()
	actor.SetHostId(owner.HostId())
	actor.SetSessionId(owner.SessionId())
	actor.InitFit()
	return actor
}

func (this *FitActor) InitFit() {

	actors := this.leader.GetFight().GetUserByUserId(this.leader.GetUserId())
	//propFix := float64(gamedb.GetConf().FitPropFix) / base.BASE_RATE
	hpFix := float64(gamedb.GetConf().FitPropHpFix) / base.BASE_RATE
	//合体血量=三角色血量之和*系数1
	//合体其他属性=三角色之中最大值*系数2
	hpNow := 0
	heroHp := make(map[int]int)

	heroCombat := make(map[int]int)
	for _, v := range actors {
		heroIndex := v.(base.ActorUser).GetHeroIndex()
		heroCombat[heroIndex] = v.GetProp().Get(pb.PROPERTY_COMBAT)
	}

	heroRank := common.SortKvIntMapDes(heroCombat)
	heroRankStr := ""
	for k, v := range heroRank {
		heroRankStr += strconv.Itoa(v.K) + ","
		heroActor := actors[v.K]
		heroIndex := v.K
		for _, pid := range pb.PROPERTY_ARRAY {

			switch pid {

			case pb.PROPERTY_HP:
				heroHp[heroIndex] = heroActor.GetProp().HpNow()
				hpNow += int(float64(heroActor.GetProp().HpNow()) * hpFix)
				this.GetProp().AddOne(pid, int(float64(heroActor.GetProp().Get(pid))*hpFix))
			case pb.PROPERTY_PATT_MIN, pb.PROPERTY_PATT_MAX:
				if heroActor.Job() == pb.JOB_ZHANSHI {
					this.calFitProp(heroIndex, pid, heroActor.GetProp().Get(pid))
				}
			case pb.PROPERTY_MATT_MIN, pb.PROPERTY_MATT_MAX:
				if heroActor.Job() == pb.JOB_FASHI {
					this.calFitProp(heroIndex, pid, heroActor.GetProp().Get(pid))
				}
			case pb.PROPERTY_TATT_MIN, pb.PROPERTY_TATT_MAX:
				if heroActor.Job() == pb.JOB_DAOSHI {
					this.calFitProp(heroIndex, pid, heroActor.GetProp().Get(pid))
				}
			//case pb.PROPERTY_DEF_MIN, pb.PROPERTY_DEF_MAX, pb.PROPERTY_ADF_MIN, pb.PROPERTY_ADF_MAX, pb.PROPERTY_MISS, pb.PROPERTY_HIT:
			//	this.calFitProp(heroIndex, pid, heroActor.GetProp().Get(pid))
			default:
				if k == 0 {
					this.GetProp().AddOne(pid, heroActor.GetProp().Get(pid))
				}
			}
		}
	}
	this.heroHpFix = make(map[int]float64)
	for heroIndex, v := range heroHp {
		this.heroHpFix[heroIndex] = float64(v) / float64(hpNow)
	}

	logger.Info("合体，玩家：%v,战力排名：%v,武将血量比例：%v,当前血量：%v", this.leader.NickName(), heroRankStr, this.heroHpFix, hpNow)
	fitLvConf := gamedb.GetFitLevelFitLevelCfg(gamedb.GetRealId(int(this.fit.Id), int(this.fit.Lv)))
	//this.GetProp().Add(fitLvConf.Attribute)
	this.leftTime = fitLvConf.Duration

	//fashionLvConf := gamedb.GetFitFashionLevelFitFashionLevelCfg(gamedb.GetRealId(int(this.fit.FashionId), int(this.fit.FashionLv)))
	//if fashionLvConf != nil {
	//this.GetProp().Add(fashionLvConf.Attribute)
	//}

	skills := make([]*base.Skill, 0)

	for _, v := range this.fit.Skills {

		//lvConf := gamedb.GetFitSkillLevelFitSkillLevelCfg(gamedb.GetRealId(int(v.Id), int(v.Lv)))
		//this.GetProp().Add(lvConf.Attribute)

		starConf := gamedb.GetFitSkillStarFitSkillStarCfg(gamedb.GetRealId(int(v.Id), int(v.Star)))
		//this.GetProp().Add(starConf.Attribute)

		if starConf.Skill > 0 {

			skillId, skillLv := gamedb.GetSkillIdAndLv(starConf.Skill)
			newSkill, err := base.NewSkill(skillId, skillLv, 0)
			if err != nil {
				logger.Error("Skill Error %v", err.Error())
				continue
			}
			skills = append(skills, newSkill)

		}
	}
	skillId, skillLv := gamedb.GetSkillIdAndLv(gamedb.GetConf().FitGeneralSkill)
	normalSkill, err := base.NewSkill(skillId, skillLv, 0)
	if err == nil {
		skills = append(skills, normalSkill)
	}
	skills = append(skills)

	if len(this.fit.Effect) > 0 {
		for _, v := range this.fit.Effect {
			suitEffect := gamedb.GetEffectEffectCfg(int(v))
			if suitEffect != nil {
				if len(suitEffect.Buffid) > 0 {
					for _, v := range suitEffect.Buffid {
						this.AddBuff(v, this, true)
					}
				}
				if suitEffect.Skillevelid > 0 {
					skillId, skillLv := gamedb.GetSkillIdAndLv(suitEffect.Skillevelid)
					newSkill, err := base.NewSkill(skillId, skillLv, 0)
					if err != nil {
						logger.Error("Skill Error %v", err.Error())
						continue
					}
					skills = append(skills, newSkill)
				}
			}
		}
	}

	this.SetSkills(skills)

	//属性计算
	this.GetProp().Calc(pb.JOB_ZHANSHI)
	this.GetProp().SetHpNow(hpNow)
	this.GetProp().SetMpNow(this.GetProp().Get(pb.PROPERTY_MP))
	//技能设置
	this.SetSkills(skills)
}

func (this *FitActor) OnDie() bool {
	this.FitCancel()
	return false
}
func (this *FitActor) calFitProp(heroIndex int, propId int, value int) {
	propFix := float64(gamedb.GetConf().FitPropFix[heroIndex-1]) / base.BASE_RATE
	newProp := int(float64(value) * propFix)
	newPid := propId
	if n, ok := constFight.ATTACK_PROP[propId]; ok {
		newPid = n
	}
	this.GetProp().AddOne(newPid, newProp)
}

func (this *FitActor) BuildSceneObjMessage() nw.ProtoMessage {
	r := base.BuildDefaltSceneObjMessage(this)
	r.Fit = &pb.SceneFit{
		UserId:    int32(this.leader.GetUserId()),
		FitId:     this.fit.Id,
		FashionId: this.fit.FashionId,
		FashionLv: this.fit.FashionLv,
		Name:      this.leader.NickName(),
		FitLv:     this.fit.Lv,
		GuildId:   int32(this.guildId),
		GuildName: this.guildName,
		LeaderJob: int32(this.leader.Job()),
		LeaderSex: int32(this.leader.Sex()),
	}
	return r
}

func (this *FitActor) BuildAppearMessage() nw.ProtoMessage {
	appearNtf := &pb.SceneEnterNtf{}
	appearNtf.Objs = append(appearNtf.Objs, this.BuildSceneObjMessage().(*pb.SceneObj))
	return appearNtf
}

func (this *FitActor) GetBirthAreaIndex() int {
	return this.birthAreaIndex
}

func (this *FitActor) TeamIndex() int {
	return this.leader.TeamIndex()
}

func (this *FitActor) RunAI() {

	now := time.Now().Unix()
	if now-this.createTime > int64(this.leftTime) {
		this.FitCancel()
	} else {
		this.DefaultActor.RunAI()
	}
}

func (this *FitActor) FitCancel() {
	//角色进入
	this.userActorHeroInto()
	//合体离开
	this.GetFight().Leave(this)
}

func (this *FitActor) ResetHeroHpFix() {
	for k, _ := range this.heroHpFix {
		this.heroHpFix[k] = 0
	}
}

/**
 *  @Description:
 *  @param hpNow
 */
func (this *FitActor) userActorHeroInto() {

	//各角色血量=合体解除时血量*(合体前角色的血量/合体前3角色的总血量)
	hp := this.GetProp().HpNow()
	actors := this.GetFight().GetUserByUserId(this.GetUserId())
	enterActors := make([]base.Actor, 0)
	enterPoints := make([]*scene.Point, 0)
	for _, v := range actors {
		if v.GetProp().HpNow() > 0 {
			if u, ok := v.(*UserActor); ok {
				heroHpFix := this.heroHpFix[u.GetHeroIndex()]
				newHp := v.GetProp().Get(pb.PROPERTY_HP)
				if heroHpFix != 0 {
					newHp = int(float64(hp) * heroHpFix)
				}
				if newHp > v.GetProp().Get(pb.PROPERTY_HP) {
					newHp = v.GetProp().Get(pb.PROPERTY_HP)
				}
				v.GetProp().SetHpNow(newHp)
				if newHp <= 0 {
					v.(*UserActor).SetIsDeath(true, constFight.DEATH_REASON_FIT)
				}
				v.GetProp().SetMpNow(v.GetProp().Get(pb.PROPERTY_MP))
				v.SetVisible(true)
				v.SetDir(this.GetDir())

				birthPoint := this.Point()
				if u.GetHeroIndex() != constUser.USER_HERO_MAIN_INDEX {
					birthPoint = this.GetScene().GetHeroBirthPoint(v, this.Point(), this.GetDir(), u.GetHeroIndex(), constFight.FIGHT_BIRTH_TYPE_TRIANGLE)
				}
				//v.EnterScene(this.GetScene(), birthPoint)
				enterActors = append(enterActors, v)
				enterPoints = append(enterPoints, birthPoint)
			}
		} else {
			logger.Warn("武将死亡了，%v", v.NickName())
		}
	}
	//武将进入游戏
	this.GetFight().EnterMuli(enterActors, enterPoints, constFight.SCENE_ENTER_FIT)
}

func (this *FitActor) IsEnemy(target base.Actor) bool {
	return this.leader.IsEnemy(target)
}
func (this *FitActor) IsFriend(target base.Actor) bool {
	return this.leader.IsFriend(target)
}

func (this *FitActor) GetUserId() int {
	return this.leader.GetUserId()
}

func (this *FitActor) GetLeader() base.Actor {
	return this.leader
}

func (this *FitActor) GetPlayer() *base.PlayerActor {
	return this.leader.(base.ActorPlayer).GetPlayer()
}

func (this *FitActor) FitAttack(skillid int, point *pb.Point, dir int, targetIds []int) nw.ProtoMessage {

	//以下为原有的技能攻击逻辑增加进入技能碰撞模式检测
	skill := this.GetSkill(skillid)
	if skill == nil {
		logger.Error("UserActor:UserAttack:nil skill:%d", skillid)
		return base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), gamedb.ERRSKILLNOTFOUND.Code)
	}

	//普通技能
	if err, hasEffect := this.CanUseSkill(skill); err != nil {
		return base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), err.(*errex.ErrorItem).Code)
	} else if !hasEffect {
		//技能标记使用，记录cd
		this.UseSkill(skill, false)
		return base.SkillAttackEffect(this, skill, this.GetDir(), make([]*pb.HurtEffect, 0), 0)
	}

	if this.GetProp().HpNow() <= 0 {
		return base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), gamedb.ERRPLAYERDIE.Code)
	}

	attackEffectNtf, err := CastSkill(this, skill, dir, targetIds, false)
	if err != nil {
		attackEffectNtf = base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), err.(*errex.ErrorItem).Code)
	}
	return attackEffectNtf
}
