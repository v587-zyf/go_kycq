package actorPkg

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/errex"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type UserActor struct {
	*DefaultActor
	player       *base.PlayerActor
	heroIndex    int
	fightModel   int
	collectionId int //标记为采集，值为采集物品ObjId
	userType     int

	leader     base.Actor //归属玩家
	reliveSelf int64      //自动复活时间

	isRobot   bool
	vip       int
	level     int
	guildId   int
	guildName string
	elf       *pbserver.ElfInfo
	elfSkills []*base.Skill
	cutSkill  *base.Skill
}

func (this *UserActor) GetPlayer() *base.PlayerActor {
	return this.player
}

func (this *UserActor) ReliveSelf() int64 {
	return this.reliveSelf
}

func (this *UserActor) SetCollectionId(collectionId int) {
	this.collectionId = collectionId
}

func (this *UserActor) GuildName() string {
	return this.guildName
}

func (this *UserActor) GuildId() int {
	return this.guildId
}

func (this *UserActor) HeroIndex() int {
	return this.heroIndex
}

func NewUserActor(player *base.PlayerActor, act *pbserver.User, heroIndex int, leader base.Actor, aiCreator base.AICreator) *UserActor {
	logger.Info("创建战斗玩家数据start,userId:%v,userName:%v", act.UserInfo.UserId, act.UserInfo.NickName, act.LocatedServerId, act.SessionId)
	actor := &UserActor{
		player:     player,
		vip:        int(act.UserInfo.Vip),
		fightModel: int(act.FightModel),
		heroIndex:  int(heroIndex),
		guildId:    int(act.UserInfo.GuildId),
		guildName:  act.UserInfo.GuildName,
		//reliveTest: 1,
	}
	heroInfo := act.UserInfo.Heros[int32(heroIndex)]
	//创建默认的战斗单元角色
	actor.DefaultActor = NewDefaultActor(pb.SCENEOBJTYPE_USER, heroInfo.NickName, act.UserInfo.Avatar, heroInfo.Job, heroInfo.DisplayInfo, aiCreator(actor), actor)
	actor.sex = int(heroInfo.Sex)
	actor.level = int(heroInfo.Level)
	actor.userId = act.UserInfo.UserId
	actor.SetOfficial(int(act.UserInfo.Official))
	if heroIndex == constUser.USER_HERO_MAIN_INDEX {
		actor.SetHostId(int(act.LocatedServerId))
		actor.SetSessionId(act.SessionId)
		actor.leader = actor
	} else {
		actor.leader = leader
	}
	actor.BuildPropByActor(heroInfo, true)
	actor.InitSkill(heroInfo.Skills)
	actor.SetTeamIndex(int(act.TeamId))
	if len(heroInfo.Buffs) > 0 {
		for _, v := range heroInfo.Buffs {
			actor.AddBuff(int(v), actor, true)
		}
	}
	actor.userType = int(act.UserType)
	//初始化精灵技能
	actor.initElf(act.UserInfo.Elf)
	//初始切割打包技能
	if act.UserInfo.CutSkill != nil && act.UserInfo.CutSkill.Id > 0 {
		newSkill, err := base.NewSkill(int(act.UserInfo.CutSkill.Id), int(act.UserInfo.CutSkill.Level), 0)
		if err == nil {
			actor.cutSkill = newSkill
		}
	}
	logger.Info("创建战斗玩家数据stop,userId:%v,userName:%v", actor.userId, actor.nickName)
	return actor
}

func (this *UserActor) initElf(elfInfo *pbserver.ElfInfo) {
	this.elf = elfInfo
	this.elfSkills = make([]*base.Skill, 0)
	for _, v := range elfInfo.Skills {
		newSkill, err := base.NewSkill(int(v.Id), int(v.Level), 0)
		if err != nil {
			continue
		}
		this.elfSkills = append(this.elfSkills, newSkill)
	}
}

func (this *UserActor) GetElfAttack() float64 {

	elfConf := gamedb.GetElfGrowElfGrowCfg(int(this.elf.Lv))
	return float64(elfConf.Buff[pb.PROPERTY_ATT_PETS]) * (1 + float64(elfConf.Buff[pb.PROPERTY_ATT_PETS_RATE])/base.BASE_RATE)
}

func (this *UserActor) SetFightModel(fightModel int) {
	this.fightModel = fightModel
}

func (this *UserActor) OnDie() bool {

	//死亡buff
	isRelive := this.buffManager.ReliveBuffCheck()
	if !isRelive {
		this.SetIsDeath(true, constFight.DEATH_REASON_ATTACK)
	} else {
		this.SetRelive()
	}

	isAllDie := this.GetFight().CheckUserAllDie(this)
	logger.Debug("玩家%v死亡,是否触发复活：%v,是否全部死亡：%v",this.nickName,isRelive,isAllDie)
	if isAllDie {

		petActor := this.GetFight().GetPetActor(this.GetUserId())
		if petActor != nil {
			leaveNtf := &pb.SceneLeaveNtf{
				ObjIds:    make([]int32, 0),
				LeaveType: 1,
			}
			petActor.SetVisible(false)
			leaveNtf.ObjIds = append(leaveNtf.ObjIds, int32(petActor.GetObjId()))
			this.GetScene().NotifyNearby(this, leaveNtf, nil)
		}
	}

	if this.job == pb.JOB_DAOSHI {
		summonActors := this.player.SummonActors()
		if summonActors != nil && len(summonActors) > 0 {
			for _, v := range summonActors {
				this.GetFight().Leave(v)
			}
		}
	}

	return isRelive
}

func (this *UserActor) SetRelive() {
	this.reliveSelf = common.GetNowMillisecond() + constFight.RELIVE_DELAY
}

func (this *UserActor) LeaveScene() {
	this.DefaultActor.LeaveScene()
	this.ResetCollectionStatus()
	this.fight = nil
}

func (this *UserActor) JustLeaveScene() {
	this.DefaultActor.LeaveScene()
}

func (this *UserActor) ResetCollectionStatus() {
	if this.collectionId > 0 {
		this.GetFight().ResetCollection(this.collectionId)
		this.collectionId = 0
	}
}

func (this *UserActor) GetUserId() int {
	return int(this.userId)
}

func (this *UserActor) GetHeroIndex() int {
	return this.heroIndex
}

func (this *UserActor) GetLvl() int {
	return this.official
}

func (this *UserActor) UserType() int {
	return this.userType
}

func (this *UserActor) GetLeader() base.Actor {
	return this.leader
}

func (this *UserActor) InitSkill(skills []*pbserver.Skill) {
	var userSkills []*base.Skill

	for _, skill := range skills {
		newSkill, err := base.NewSkill(int(skill.Id), int(skill.Level), skill.CdEndTime)
		if err != nil {
			continue
		}
		if skill.TalentEffect != nil {
			newSkill.SetTalentEffect(common.ConvertInt32SliceToIntSlice(skill.TalentEffect))
		}
		userSkills = append(userSkills, newSkill)
	}
	this.SetSkills(userSkills)
}

func (this *UserActor) ResetSkill(skills []*pbserver.Skill) {

	tempSkills := this.skills
	tempPassiveSkills := this.passiveSkill
	this.skills = make([]*base.Skill, 0)
	this.passiveSkill = make([]*base.Skill, 0)

loop:
	for _, newSkillData := range skills {
		has := false
		for _, baseSkill := range tempSkills {
			if baseSkill.Skillid == int(newSkillData.Id) {
				has = true
				if baseSkill.LevelT.Level < int(newSkillData.Level) {
					baseSkill.ResetSkillLvel(int(newSkillData.Level))
				}
				if newSkillData.TalentEffect != nil {
					baseSkill.SetTalentEffect(common.ConvertInt32SliceToIntSlice(newSkillData.TalentEffect))
				}
				this.skills = append(this.skills, baseSkill)
				continue loop
			}
		}
		for _, baseSkill := range tempPassiveSkills {
			if baseSkill.Skillid == int(newSkillData.Id) {
				has = true
				if baseSkill.LevelT.Level < int(newSkillData.Level) {
					baseSkill.ResetSkillLvel(int(newSkillData.Level))
				}
				this.passiveSkill = append(this.passiveSkill, baseSkill)
				continue loop
			}
		}
		if !has {
			newSkill, err := base.NewSkill(int(newSkillData.Id), int(newSkillData.Level), newSkillData.CdEndTime)
			if err != nil {
				logger.Error("更新玩家技能数据异常，：%v", err)
				continue
			}
			if newSkillData.TalentEffect != nil {
				newSkill.SetTalentEffect(common.ConvertInt32SliceToIntSlice(newSkillData.TalentEffect))
			}
			logger.Debug("更新玩家数据，增加新的技能：%v", newSkill.Skillid, newSkill.LevelT.Level)
			if newSkill.Type == pb.SKILLTYPE_PASSIVE || newSkill.SkillSkillCfg.Type == pb.SKILLTYPE_PASSIVE2 {
				this.passiveSkill = append(this.passiveSkill, newSkill)
			} else {
				this.skills = append(this.skills, newSkill)
			}
		}
	}
}

func (this *UserActor) IsEnemy(target base.Actor) bool {

	//角色自己肯定不能是目标
	if target.GetObjId() == this.GetObjId() {
		return false
	}

	//玩家自己武将
	if this.GetUserId() == target.GetUserId() {
		return false
	}

	if target.TeamIndex() != this.TeamIndex() {

		stageType := this.GetFight().GetStageConf().Type
		switch stageType {
		case constFight.FIGHT_TYPE_DARKPALACE_BOSS, constFight.FIGHT_TYPE_HELL_BOSS:
			if this.player.FightNum() <= 0 && target.GetType() == pb.SCENEOBJTYPE_MONSTER {
				return false
			}
		}

		return true
	}

	return false

	//stageType := this.GetFight().GetStageConf().Type
	//switch stageType {
	//case constFight.FIGHT_TYPE_CROSS_WORLD_LEADER, //世界首领
	//	constFight.FIGHT_TYPE_GUILD_BONFIRE: //门派篝火
	//	if target.GetType() == base.ActorTypeMonster {
	//		return true
	//	} else {
	//		return false
	//	}
	//case constFight.FIGHT_TYPE_ARENA,
	//	constFight.FIGHT_TYPE_FIELD,
	//	constFight.FIGHT_TYPE_MINING,
	//	constFight.FIGHT_TYPE_SHABAKE,
	//	constFight.FIGHT_TYPE_CROSS_SHABAKE,
	//	constFight.FIGHT_TYPE_GUARDPILLAR:
	//	if target.TeamIndex() != this.TeamIndex() {
	//		return true
	//	}
	//	return false
	//default:
	//	if target.GetType() == base.ActorTypeMonster {
	//		return true
	//	}
	//}
	//
	//if target.GetType() == pb.SCENEOBJTYPE_FIT || target.GetType() == pb.SCENEOBJTYPE_SUMMON {
	//	target = target.(base.ActorLeader).GetLeader()
	//}
	//
	//if target.GetType() == pb.SCENEOBJTYPE_USER {
	//
	//	return true
	//	//if this.fightModel == pb.FIGHTMODEL_PEACE {
	//	//	//和平模式，攻击不了任何人
	//	//	return false
	//	//} else if this.fightModel == pb.FIGHTMODEL_WHOLE {
	//	//	//全体模式
	//	//	return true
	//	//}
	//}
	//return false
}
func (this *UserActor) IsFriend(target base.Actor) bool {
	//if target.GetType() == base.ActorTypeMonster {
	//	return false
	//}
	//if this.GetFight().GetStageConf().Type == constFight.FIGHT_TYPE_ARENA || this.GetFight().GetStageConf().Type == constFight.FIGHT_TYPE_FIELD {
	//	if target.TeamIndex() == this.TeamIndex() {
	//		return true
	//	}
	//}
	////目前只有自身是友方
	//if target.GetUserId() == this.GetUserId() {
	//	return true
	//}
	if target.TeamIndex() == this.TeamIndex() {
		return true
	}
	return false
}

func (this *UserActor) UpdateElf(elf *pbserver.ElfInfo) error {

	this.initElf(elf)
	return nil
}

func (this *UserActor) UseCutTreasure(cutTreasureLv int) error {

	cutTreasureConf := gamedb.GetCutTreasureByLv(cutTreasureLv)
	cutSkill := 0
	cutBuffId := 0
	if this.job == pb.JOB_ZHANSHI {
		cutSkill = cutTreasureConf.SkillZhan
		cutBuffId = cutTreasureConf.Buffzhan
	} else if this.job == pb.JOB_FASHI {
		cutSkill = cutTreasureConf.SkillFa
		cutBuffId = cutTreasureConf.Bufffa
	} else if this.job == pb.JOB_DAOSHI {
		cutSkill = cutTreasureConf.SkillDao
		cutBuffId = cutTreasureConf.Buffdao
	}
	skillId, skillLv := gamedb.GetSkillIdAndLv(cutSkill)
	if this.cutSkill == nil || this.cutSkill.Skillid != skillId || this.cutSkill.LevelT.Level != skillLv {
		newSkill, err := base.NewSkill(skillId, skillLv, 0)
		if err != nil {
			return err
		}
		this.cutSkill = newSkill
	}
	var err error
	if cutBuffId > 0 {
		_, err = this.AddBuff(cutBuffId, this.context, false)
	}
	logger.Debug("玩家使用切割打宝，玩家：%v,切割打包等级：%v,技能：%v,buff:%v", this.context.NickName(), cutTreasureLv, cutSkill, cutBuffId)
	return err
}

func (this *UserActor) GetSkill(skillId int) *base.Skill {

	skill := this.DefaultActor.GetSkill(skillId)
	if skill != nil {
		return skill
	}
	return this.player.GetSkill(skillId)
}

func (this *UserActor) TriggerPassiveSkillByHpChange(passiveType int, target base.Actor, oldHp, newHp int) {
	this.DefaultActor.TriggerPassiveSkillByHpChange(passiveType, target, oldHp, newHp)

	//玩家技能触发
	canTrigger := true
	for i := this.heroIndex - 1; i >= constUser.USER_HERO_MAIN_INDEX; i-- {
		if this.player.GetHeroActor(i).GetProp().HpNow() > 0 {
			canTrigger = false
			break
		}
	}
	if !canTrigger {
		return
	}
	playerPassiveSkills := this.player.PassiveSkills()
	TriggerPassiveSkillByHpChange(this, playerPassiveSkills, passiveType, target, oldHp, newHp)
}

func (this *UserActor) TriggerPassiveSkill(passiveType int, target base.Actor, skill *base.Skill) {
	this.DefaultActor.TriggerPassiveSkill(passiveType, target, skill)
	//玩家技能触发
	canTrigger := true
	for i := this.heroIndex - 1; i >= constUser.USER_HERO_MAIN_INDEX; i-- {
		if this.player.GetHeroActor(i).GetProp().HpNow() > 0 {
			canTrigger = false
			break
		}
	}
	if !canTrigger {
		return
	}
	playerPassiveSkills := this.player.PassiveSkills()
	TriggerPassiveSkill(this, playerPassiveSkills, this.passiveSkillConditionTimes, passiveType, target, skill)
}

func (this *UserActor) UserAttack(skillid int, point *pb.Point, dir int, targetIds []int, iself bool) nw.ProtoMessage {

	//以下为原有的技能攻击逻辑增加进入技能碰撞模式检测
	var skill *base.Skill
	if iself {
		for _, v := range this.elfSkills {
			if v.Skillid == skillid {
				skill = v
				break
			}
		}
	} else {
		//判断是否切割打包技能
		if this.cutSkill != nil && skillid == this.cutSkill.Skillid {
			//检查是否有切割状态 优化，删除判断
			//if has, _ := this.buffManager.BuffHasType(pb.BUFFTYPE_CUT_SKILL, nil); !has {
			//	return base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), gamedb.ERRCUTTREASUREBUFF.Code)
			//}

			skill = this.cutSkill
		} else {
			skill = this.GetSkill(skillid)
		}
	}
	if skill == nil {
		logger.Error("UserActor:UserAttack:nil skill:%d", skillid)
		ntf := base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), gamedb.ERRSKILLNOTFOUND.Code)
		ntf.IsElf = iself
		return ntf
	}

	if this.GetProp().HpNow() <= 0 {
		ntf := base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), gamedb.ERRPLAYERDIE.Code)
		ntf.IsElf = iself
		return ntf
	}

	//普通技能
	if iself {
		if err := skill.CanUse(0); err != nil {
			ntf := base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), err.(*errex.ErrorItem).Code)
			ntf.IsElf = iself
			return ntf
		}
	} else {
		if err, hasEffect := this.CanUseSkill(skill); err != nil {
			return base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), err.(*errex.ErrorItem).Code)
		} else if !hasEffect {
			//技能标记使用，记录cd
			this.UseSkill(skill, false)
			return base.SkillAttackEffect(this, skill, this.GetDir(), make([]*pb.HurtEffect, 0), 0)
		}
	}

	//fight := this.GetFight()
	attackEffectNtf, err := CastSkill(this, skill, dir, targetIds, iself)
	if err == nil {
		//if !iself {
		//	killMonsterIds := make([]int32, 0)
		//	killUserNum := make(map[int]bool)
		//	for _, v := range attackEffectNtf.Hurts {
		//		if !v.IsDeath {
		//			continue
		//		}
		//		objs := fight.GetActorByObjId(int(v.ObjId))
		//		if objs == nil {
		//			continue
		//		}
		//		if objs.GetType() == pb.SCENEOBJTYPE_MONSTER {
		//			if m, ok := objs.(*MonsterActor); ok {
		//				killMonsterIds = append(killMonsterIds, int32(m.MonsterT.Monsterid))
		//			}
		//		} else if objs.GetType() == pb.SCENEOBJTYPE_USER || objs.GetType() == pb.SCENEOBJTYPE_FIT {
		//			isAllDie := this.GetFight().CheckUserAllDieByHp(objs)
		//			if isAllDie {
		//				if _, ok := killUserNum[objs.GetUserId()]; !ok {
		//					killUserNum[objs.GetUserId()] = true
		//				}
		//			}
		//		}
		//	}
		//	//说明技能正确使用了
		//	net.GetGsConn().SendMessageToGs(uint32(this.HostId()),
		//		&pbserver.FsSkillUseNtf{
		//			UseId:          int32(this.GetUserId()),
		//			HeroIndex:      int32(this.heroIndex),
		//			SkillId:        int32(skillid),
		//			CdStartTime:    common.GetNowMillisecond(),
		//			CdStopTime:     skill.GetNextAttackTime(),
		//			KillMonsterIds: killMonsterIds,
		//			KillUserNum:    int32(len(killUserNum)),
		//		})
		//}
	} else {
		attackEffectNtf = base.SkillAttackEffect(this, skill, 0, make([]*pb.HurtEffect, 0), err.(*errex.ErrorItem).Code)
		attackEffectNtf.IsElf = iself
	}
	return attackEffectNtf
}

func (this *UserActor) BuildSceneObjMessage() nw.ProtoMessage {
	r := base.BuildDefaltSceneObjMessage(this)
	r.User = this.BuildSceneUser()
	return r
}

func (this *UserActor) BuildSceneUser() *pb.SceneUser {
	return &pb.SceneUser{
		UserId: int32(this.userId),
		Name:   this.nickName,
		Vip:    int32(this.vip),
		Lvl:    int32(this.level),
		Display: &pb.Display{
			ClothItemId: this.displayInfo.ClothItemId, ClothType: this.displayInfo.ClothType,
			WeaponItemId: this.displayInfo.WeaponItemId, WeaponType: this.displayInfo.WeaponType,
			WingId: this.displayInfo.WingId, MagicCircleLvId: this.displayInfo.MagicCircleLvId,
			TitleId: this.displayInfo.TitleId, LabelId: this.displayInfo.LabelId,
			LabelJob: this.displayInfo.LabelJob,
		},
		Job:          int32(this.job),
		Sex:          int32(this.sex),
		Combat:       int64(this.GetProp().Combat()),
		Avatar:       this.avatar,
		HeroIndex:    int32(this.heroIndex),
		GuildId:      int32(this.guildId),
		GuildName:    this.guildName,
		ElfLv:        this.elf.Lv,
		Username:     this.leader.NickName(),
		Userjob:      int32(this.leader.Job()),
		Usersex:      int32(this.leader.Sex()),
		UserHpTotal:  int64(this.player.UserHpTotal()),
		ToHelpUserId: int32(this.player.ToHelpUserId()),
	}
}

func (this *UserActor) BuildAppearMessage() nw.ProtoMessage {
	appearNtf := &pb.SceneEnterNtf{
		StageId: int32(this.GetFight().GetStageConf().Id),
	}
	appearNtf.Objs = append(appearNtf.Objs, this.BuildSceneObjMessage().(*pb.SceneObj))
	return appearNtf
}

func (this *UserActor) BuildRelliveMessage() nw.ProtoMessage {

	reliveNtf := &pb.SceneUserReliveNtf{
		Obj: this.BuildSceneObjMessage().(*pb.SceneObj),
	}
	return reliveNtf
}

func (this *UserActor) UpdateUserInfo(newUser *pbserver.Actor, heroIndex int) *UserActor {

	heroInfo := newUser.Heros[int32(heroIndex)]
	logger.Debug("接收到game消息，更新玩家数据,玩家：%v,武将：%v", newUser.NickName, heroIndex)
	this.UpdateUserDisplayInfo(newUser, heroIndex)
	//更新玩家属性
	this.BuildPropByActor(heroInfo, false)
	//重新初始化玩家技能
	this.ResetSkill(heroInfo.Skills)
	//buff
	for _, v := range heroInfo.Buffs {
		this.AddBuff(int(v), this, false)
	}
	return this
}

func (this *UserActor) BuildPropByActor(userHeroInfo *pbserver.ActorHero, isNewUser bool) {
	actorProp := this.GetProp()
	nowTotolHp := actorProp.Get(pb.PROPERTY_HP)
	actorProp.ByFightActorProp(userHeroInfo.Prop)
	if isNewUser {
		actorProp.SetHpNow(actorProp.Get(pb.PROPERTY_HP))
		actorProp.SetMpNow(actorProp.Get(pb.PROPERTY_MP))
	} else {
		ratio := float64(actorProp.HpNow()) / float64(nowTotolHp)
		newTotalHp := actorProp.Get(pb.PROPERTY_HP)
		actorProp.SetHpNow(int(float64(newTotalHp) * ratio))
		if actorProp.HpNow() > newTotalHp {
			actorProp.SetHpNow(newTotalHp)
		}
		if nowTotolHp != newTotalHp {
			this.GetPlayer().CalcUserHpTotal()
			//推送血量变化
			HPChangeNtf := &pb.SceneObjHpNtf{
				ObjId:       int32(this.GetObjId()),
				Hp:          int64(this.GetProp().HpNow()),
				ChangeHp:    int64(0),
				TotalHp:     int64(newTotalHp),
				UserId:      int32(this.GetUserId()),
				UserHpTotal: int64(this.GetPlayer().UserHpTotal()),
			}
			this.NotifyNearby(this, HPChangeNtf, nil)
		}
	}
	//actorProp = this.GetBaseProp()
	//actorProp.ByFightActorProp(userHeroInfo.Prop)
}

func (this *UserActor) UpdateUserDisplayInfo(newUser *pbserver.Actor, heroIndex int) {

	hasChange := false
	if newUser.NickName != this.nickName {
		this.nickName = newUser.Heros[int32(heroIndex)].NickName
		hasChange = true
	}
	if int(newUser.Heros[int32(heroIndex)].Level) != this.level {
		this.level = int(newUser.Heros[int32(heroIndex)].Level)
		hasChange = true
	}
	if this.displayInfo.ClothItemId != newUser.Heros[int32(heroIndex)].DisplayInfo.ClothItemId ||
		this.displayInfo.WeaponItemId != newUser.Heros[int32(heroIndex)].DisplayInfo.WeaponItemId ||
		this.displayInfo.WingId != newUser.Heros[int32(heroIndex)].DisplayInfo.WingId ||
		this.displayInfo.MagicCircleLvId != newUser.Heros[int32(heroIndex)].DisplayInfo.MagicCircleLvId ||
		this.displayInfo.TitleId != newUser.Heros[int32(heroIndex)].DisplayInfo.TitleId ||
		this.displayInfo.LabelId != newUser.Heros[int32(heroIndex)].DisplayInfo.LabelId ||
		this.displayInfo.LabelJob != newUser.Heros[int32(heroIndex)].DisplayInfo.LabelJob {
		this.displayInfo = newUser.Heros[int32(heroIndex)].DisplayInfo
		hasChange = true
	}
	if hasChange {
		this.NotifyNearby(this, &pb.SceneUserUpdateNtf{
			ObjId:   int32(this.GetObjId()),
			ObjUser: this.BuildSceneUser(),
		}, nil)
	}
}

// 复制增益buff
func (this *UserActor) CloneGainBuff(source base.Actor) {

	this.buffManager.CloneGainBuff(source)
}

func (this *UserActor) UserItem(itemId int) {

	itemConf := gamedb.GetItemBaseCfg(itemId)
	if itemConf == nil {
		return
	}
	if itemConf.Type == pb.ITEMTYPE_POTION || itemConf.Type == pb.ITEMTYPE_HP_RECOVER || itemConf.Type == pb.ITEMTYPE_MP_RECOVER {
		_, err := this.AddBuff(itemConf.EffectVal, this, false)
		if err != nil {
			logger.Error("使用血瓶添加buff异常：玩家：%v-%v，错误：%v", this.userId, this.nickName, err)
		}
	}
}

func (this *UserActor) ClearSkillCD() {
	this.DefaultActor.ClearSkillCD()
	mainActor := this.GetFight().GetUserMainActor(this.GetUserId())
	if mainActor != nil {

		err := net.GetGsConn().SendMessageToGs(uint32(mainActor.HostId()), &pbserver.FsTOGsClearSkillCdNtf{UserId: int32(this.userId), HeroIndex: int32(this.heroIndex)})
		if err != nil {
			logger.Error("发送game清理技能CD异常")
		}
	}
	logger.Debug("清理技能Cd:%v", this.NickName())
}

func (this *UserActor) SendMessage(msg nw.ProtoMessage) {

	mainActor := this.GetFight().GetUserMainActor(this.GetUserId())
	if mainActor != nil {
		net.GetGateConn().SendMessage(uint32(mainActor.HostId()), mainActor.SessionId(), 0, msg)
	}
}

func (this *UserActor) Relive(reliveAddr int, reliveType int) {
	this.DefaultActor.Relive(reliveAddr, reliveType)
	this.reliveSelf = 0
}

func (this *UserActor) ChangeHp(changeHp int) (realChange int, isDeath bool) {
	this.ResetCollectionStatus()
	return this.DefaultActor.ChangeHp(changeHp)
}
