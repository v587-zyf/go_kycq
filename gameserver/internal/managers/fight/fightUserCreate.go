package fight

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"fmt"
)

func (this *Fight) isCrossFight(stageId int) bool {
	crossServerId := this.GetCrossFightServerId(stageId)
	isCross := crossServerId != 0
	return isCross
}

func (this *Fight) GetUserHeroFightNickname(user *objs.User, heroIndex int, isCrossFight bool) string {
	prefix := this.GetSystem().GetPrefix()
	serverIndex := user.ServerIndex
	nickName := ""
	if isCrossFight {
		nickName = fmt.Sprintf("%s%d.%s", prefix, serverIndex, user.NickName)
	} else {
		nickName = fmt.Sprintf("%s", user.NickName)
	}
	if heroIndex != constUser.USER_HERO_MAIN_INDEX {
		nickName = nickName + fmt.Sprintf(".%s", user.Heros[heroIndex].Name)
	}
	return nickName
}

func (this *Fight) createFightUserInfo(user *objs.User, teamId int, intoHeroIndex int, createPet bool, stageId int, helpToUseId int) *pbserver.User {

	fightUser := &pbserver.User{}
	fightUser.SessionId = user.GateSessionId
	fightUser.LocatedServerId = uint32(base.Conf.ServerId)
	fightUser.TeamId = int32(teamId)
	fightUser.UserInfo = &pbserver.Actor{}
	fightUser.FightModel = int32(user.FightModel)
	fightUser.ToHelpUserId = int32(helpToUseId)

	//基础信息
	this.crateFightUserBaseInfo(user, fightUser.UserInfo, stageId)
	if createPet {
		petActor := this.createUserPet(user)
		fightUser.UserInfo.Pet = petActor
	}
	for heroIndex, heroInfo := range user.Heros {

		if intoHeroIndex != -1 && heroIndex != intoHeroIndex {
			continue
		}
		actorHero := &pbserver.ActorHero{}
		fightUser.UserInfo.Heros[int32(heroIndex)] = actorHero
		this.createFighUserHeroBaseInfo(user, heroInfo, actorHero, this.isCrossFight(stageId))
		//属性信息
		actorHero.Prop = this.createFightHeroProp(user, heroInfo)
		//主城不战斗不初始化技能 buff信息
		if stageId == constFight.FIGHT_TYPE_MAIN_CITY_STAGE {
			continue
		}
		//技能信息
		this.createFightUserSkillInfo(heroInfo, actorHero)
		//出生buff
		this.createFightUserInitBuff(heroInfo, actorHero)
		//地图额外buff
		this.createFightUserStageBuff(heroInfo, actorHero, stageId)
	}

	return fightUser
}

func (this *Fight) GetFightUserInfo(user *objs.User, teamId int, intoHeroIndex int, createPet bool, isCrossFight bool) *pbserver.User {

	fightUser := &pbserver.User{}
	fightUser.SessionId = user.GateSessionId
	fightUser.LocatedServerId = uint32(base.Conf.ServerId)
	fightUser.TeamId = int32(teamId)
	fightUser.UserInfo = &pbserver.Actor{}
	fightUser.FightModel = int32(user.FightModel)

	//基础信息
	this.crateFightUserBaseInfo(user, fightUser.UserInfo, 0)
	if createPet {
		petActor := this.createUserPet(user)
		fightUser.UserInfo.Pet = petActor
	}
	for heroIndex, heroInfo := range user.Heros {

		if intoHeroIndex != -1 && heroIndex != intoHeroIndex {
			continue
		}
		actorHero := &pbserver.ActorHero{}
		fightUser.UserInfo.Heros[int32(heroIndex)] = actorHero
		this.createFighUserHeroBaseInfo(user, heroInfo, actorHero, isCrossFight)
		//技能信息
		this.createFightUserSkillInfo(heroInfo, actorHero)
		//属性信息
		actorHero.Prop = this.createFightHeroProp(user, heroInfo)
		//出生buff
		this.createFightUserInitBuff(heroInfo, actorHero)
		//地图额外buff
		this.createFightUserStageBuff(heroInfo, actorHero, user.FightStageId)
	}

	return fightUser
}

func (this *Fight) createFightHeroProp(user *objs.User, hero *objs.Hero) *pbserver.ActorProp {
	//测试环境 自主设置的属性
	var gmProp *pbserver.ActorProp
	if base.Conf.Sandbox && hero.Index == constUser.USER_HERO_MAIN_INDEX {
		if user.GmProperty != nil {
			gmProp = user.GmProperty.ToFightActorProp()
		}
	}
	heroProp := hero.Prop.ToFightActorProp()
	if gmProp != nil {
		for k, v := range gmProp.Props {
			if v > 0 {
				heroProp.Props[k] = v
			}
		}
	}
	return heroProp
}

/**
 *  @Description: 玩家战斗基础信息
 *  @param user
 *  @param actor
 */
func (this *Fight) crateFightUserBaseInfo(user *objs.User, actor *pbserver.Actor, stageId int) {

	actor.UserId = uint32(user.Id)
	actor.Avatar = user.Avatar
	actor.Vip = int32(user.VipLevel)
	actor.Official = int32(user.Official)
	actor.NickName = user.NickName
	actor.UserCombat = int64(user.Combat)
	actor.DarkPalaceTimes = int32(this.GetDarkPalace().GetSurplusNum(user))
	actor.DabaoEnergy = int32(user.DaBaoMystery.Energy)
	stageConf := gamedb.GetStageStageCfg(stageId)
	if stageConf != nil && stageConf.Type == constFight.FIGHT_TYPE_HELL_BOSS {
		actor.DarkPalaceTimes = int32(this.GetHellBoss().GetSurplusNum(user))
	}
	actor.Heros = make(map[int32]*pbserver.ActorHero)
	actor.StageFightNum = int32(this.GetCondition().GetConditionData(user, pb.CONDITION_ALL_KILL_STAGE, stageId))
	guildInfo := this.GetGuild().GetGuildInfo(user.GuildData.NowGuildId)
	if guildInfo != nil {
		actor.GuildId = int32(guildInfo.GuildId)
		actor.GuildName = guildInfo.GuildName
	}
	actor.RedPacket = &pbserver.ActorRedPacket{
		PickNum:   int32(user.RedPacketItem.PickNum),
		PickMax:   int32(gamedb.GetRedPacketDropMax(this.GetSystem().GetServerOpenDaysByServerId(user.ServerId))),
		PickInfos: common.ConvertMapIntToInt32(user.RedPacketItem.PickInfo),
	}
	//精灵信息
	actor.Elf = this.createUserElfInfo(user)
	if user.Elf.Lv > 0 {
		for _, v := range user.Elf.SkillBag {
			skillLv := user.Elf.Skills[v]
			actor.Elf.Skills = append(actor.Elf.Skills, &pbserver.Skill{
				Id:    int32(v),
				Level: int32(skillLv),
			})
		}
	}

	//切割打包技能
	if user.CutTreasure > 0 {
		cutTreasureConf := gamedb.GetCutTreasureByLv(user.CutTreasure)
		cutSkillId := 0
		if user.Heros[constUser.USER_HERO_MAIN_INDEX].Job == pb.JOB_ZHANSHI {
			cutSkillId = cutTreasureConf.SkillZhan
		} else if user.Heros[constUser.USER_HERO_MAIN_INDEX].Job == pb.JOB_FASHI {
			cutSkillId = cutTreasureConf.SkillFa
		} else if user.Heros[constUser.USER_HERO_MAIN_INDEX].Job == pb.JOB_DAOSHI {
			cutSkillId = cutTreasureConf.SkillDao
		}
		if cutSkillId > 0 {
			skillId, SkillLv := gamedb.GetSkillIdAndLv(cutSkillId)
			actor.CutSkill = &pbserver.Skill{Id: int32(skillId), Level: int32(SkillLv)}
		}
	}

	//主城不初始化技能
	if stageConf == nil || stageConf.Type == constFight.FIGHT_TYPE_MAIN_CITY {
		return
	}
	//绝学
	if len(user.Heros[constUser.USER_HERO_MAIN_INDEX].JuexueEffects) > 0 {
		effects := user.Heros[constUser.USER_HERO_MAIN_INDEX].JuexueEffects
		for _, v := range effects {
			if skill := this.initEffectSkill(v); skill != nil {
				actor.PublicSkill = append(actor.PublicSkill, skill)
			}
		}
	}
	if len(user.Heros[constUser.USER_HERO_MAIN_INDEX].MiJiEffects) > 0 {
		effects := user.Heros[constUser.USER_HERO_MAIN_INDEX].MiJiEffects
		for _, v := range effects {
			if skill := this.initEffectSkill(v); skill != nil {
				actor.PublicSkill = append(actor.PublicSkill, skill)
			}
		}
	}
	if len(user.Heros[constUser.USER_HERO_MAIN_INDEX].PetAppendageEffects) > 0 {
		effects := user.Heros[constUser.USER_HERO_MAIN_INDEX].PetAppendageEffects
		for _, v := range effects {
			if skill := this.initEffectSkill(v); skill != nil {
				actor.PublicSkill = append(actor.PublicSkill, skill)
			}
		}
	}
}

/**
 *  @Description: 创建精灵信息
 *  @param user
 *  @return *pbserver.ElfInfo
 */
func (this *Fight) createUserElfInfo(user *objs.User) *pbserver.ElfInfo {
	//精灵信息
	elf := &pbserver.ElfInfo{
		Lv:     int32(user.Elf.Lv),
		Skills: make([]*pbserver.Skill, 0),
	}
	if user.Elf.Lv > 0 {
		for _, v := range user.Elf.SkillBag {
			skillLv := user.Elf.Skills[v]
			elf.Skills = append(elf.Skills, &pbserver.Skill{
				Id:    int32(v),
				Level: int32(skillLv),
			})
		}
	}
	return elf
}

/**
 *  @Description: 初始化战宠
 *  @param user
 *  @param actor
 */
func (this *Fight) createUserPet(user *objs.User) *pbserver.ActorPet {

	petActor := &pbserver.ActorPet{}
	wearPetId := user.Wear.PetId
	if wearPetId > 0 {
		pet := user.Pet[wearPetId]

		petActor.PetId = int32(wearPetId)
		petActor.Lv = int32(pet.Lv)
		petActor.Break = int32(pet.Break)
		petActor.Grade = int32(pet.Grade)
		petActor.AddSkill = user.PetAddSkills

		if petAddAttr, ok := user.PetAddAttr[wearPetId]; ok {
			petActor.AddAttr = petAddAttr
		}
	}
	return petActor
}

/**
 *  @Description: 玩家武将战斗基础信息
 *  @param user
 *  @param actor
 */
func (this *Fight) createFighUserHeroBaseInfo(user *objs.User, hero *objs.Hero, actor *pbserver.ActorHero, isCrossFight bool) {

	actor.Index = int32(hero.Index)
	actor.Job = int32(hero.Job)
	actor.Sex = int32(hero.Sex)
	actor.Level = int32(hero.ExpLvl)
	actor.NickName = this.GetUserHeroFightNickname(user, hero.Index, isCrossFight)
	heroDisplay := this.GetUserManager().GetHeroDisplay(hero)
	actor.DisplayInfo = &pbserver.ActorDisplayInfo{
		heroDisplay.ClothItemId, heroDisplay.ClothType,
		heroDisplay.WeaponItemId, heroDisplay.WeaponType,
		heroDisplay.WingId, heroDisplay.MagicCircleLvId,
		heroDisplay.TitleId, heroDisplay.LabelId, heroDisplay.LabelJob,
	}
}

func (this *Fight) createFightUserSkillInfo(hero *objs.Hero, actor *pbserver.ActorHero) {

	//普通技能
	actor.Skills = make([]*pbserver.Skill, 0)

	skillTalent := make(map[int][]int)
	for _, v := range hero.Talent.TalentList {
		for id, lv := range v.Talents {
			talentLvConf := gamedb.GetTalentLevelTalentLevelCfg(gamedb.GetRealId(id, lv))
			if talentLvConf != nil {
				if talentLvConf.Skill > 0 {
					if _, ok := skillTalent[talentLvConf.Skill]; !ok {
						skillTalent[talentLvConf.Skill] = make([]int, 0)
					}
					effectConf := gamedb.GetEffectEffectCfg(talentLvConf.Effect)
					if effectConf != nil {
						if effectConf.Skillevelid > 0 {
							skillId, skillLv := gamedb.GetSkillIdAndLv(effectConf.Skillevelid)
							actor.Skills = append(actor.Skills, &pbserver.Skill{
								Id:        int32(skillId),
								Level:     int32(skillLv),
								CdEndTime: 0,
							})
						}

						if len(effectConf.Attribute) > 0 || len(effectConf.Buffid) > 0 {

							skillTalent[talentLvConf.Skill] = append(skillTalent[talentLvConf.Skill], talentLvConf.Effect)
						}
					}
				}
			}
		}
	}

	for _, v := range hero.SkillBag {
		if skill, ok := hero.Skills[v]; ok {
			skillData := this.buildSkill(skill, skillTalent)
			actor.Skills = append(actor.Skills, skillData)
		}
	}
	//被动技能
	for _, v := range hero.Skills {
		skillConf := gamedb.GetSkillSkillCfg(v.Id)
		if skillConf != nil {
			if skillConf.Type == pb.SKILLTYPE_PASSIVE || skillConf.Type == pb.SKILLTYPE_PASSIVE2 {
				actor.Skills = append(actor.Skills, this.buildSkill(v, skillTalent))
			}
		}
	}

	//合击技能
	actor.Uniqueskills = make([]*pbserver.Skill, 0)
	for _, v := range hero.UniqueSkillBag {
		if skill, ok := hero.UniqueSkills[v]; ok {
			actor.Uniqueskills = append(actor.Uniqueskills, this.buildSkill(skill, skillTalent))
		}
	}
}

func (this *Fight) buildSkill(skill *model.SkillUnit, skillTalent map[int][]int) *pbserver.Skill {
	skillData := &pbserver.Skill{
		Id:        int32(skill.Id),
		Level:     int32(skill.Lv),
		CdEndTime: skill.EndTime,
	}
	if skillTalent != nil && skillTalent[skill.Id] != nil {
		skillData.TalentEffect = common.ConvertIntSlice2Int32Slice(skillTalent[skill.Id])
	}
	return skillData
}

func (this *Fight) initEffect(actor *pbserver.ActorHero, effectId []int, initEffectSkill bool) {
	for _, v := range effectId {
		suitEffect := gamedb.GetEffectEffectCfg(v)
		if suitEffect != nil {
			if len(suitEffect.Buffid) > 0 {
				for _, v := range suitEffect.Buffid {
					actor.Buffs = append(actor.Buffs, int32(v))
				}
			}
			if initEffectSkill {
				if skill := this.initEffectSkill(v); skill != nil {
					actor.Skills = append(actor.Skills, skill)
				}
			}
		}
	}
}

func (this *Fight) initEffectSkill(effectId int) *pbserver.Skill {
	suitEffect := gamedb.GetEffectEffectCfg(effectId)
	if suitEffect != nil {
		if suitEffect.Skillevelid > 0 {
			skillLvConf := gamedb.GetSkillLevelSkillCfg(suitEffect.Skillevelid)
			if skillLvConf != nil {
				skillId := skillLvConf.Skillid / 100
				return &pbserver.Skill{
					Id:        int32(skillId),
					Level:     int32(skillLvConf.Level),
					CdEndTime: 0,
				}
			} else {
				logger.Error("技能配置异常：%v", suitEffect.Skillevelid)
			}
		}
	}
	return nil
}

func (this *Fight) createFightUserInitBuff(hero *objs.Hero, actor *pbserver.ActorHero) {

	//出生buff
	if len(hero.ZodiacSuit) > 0 {
		this.initEffect(actor, hero.ZodiacSuit, true)
	}
	//主宰装备套装效果id
	if len(hero.DictateSuit) > 0 {
		this.initEffect(actor, hero.DictateSuit, true)
	}
	//内功技能效果id
	if len(hero.InsideEffects) > 0 {
		this.initEffect(actor, hero.InsideEffects, true)
	}
	//神兵技能效果id
	if len(hero.HolyEffects) > 0 {
		this.initEffect(actor, hero.HolyEffects, true)
	}
	//特戒幻灵效果id
	if len(hero.RingPhantomEffects) > 0 {
		this.initEffect(actor, hero.RingPhantomEffects, true)
	}
	//神翼
	if len(hero.WingEffects) > 0 {
		this.initEffect(actor, hero.WingEffects, true)
	}
	//vip
	if len(hero.VipEffects) > 0 {
		this.initEffect(actor, hero.VipEffects, true)
	}
	//圣兽
	if len(hero.HolyBeastEffects) > 0 {
		this.initEffect(actor, hero.HolyBeastEffects, true)
	}
	//天赋
	if len(hero.TalentEffects) > 0 {
		this.initEffect(actor, hero.TalentEffects, true)
	}
	////合体圣装
	//if len(hero.FitHolyEquipEffects) > 0 {
	//	this.initEffect(actor, hero.FitHolyEquipEffects)
	//}
	//绝学
	if len(hero.JuexueEffects) > 0 {
		this.initEffect(actor, hero.JuexueEffects, false)
	}
	//传世套装属性
	if len(hero.ChuanShiEquipEffects) > 0 {
		this.initEffect(actor, hero.ChuanShiEquipEffects, true)
	}
	//远古神技
	if len(hero.AncientSkillEffects) > 0 {
		this.initEffect(actor, hero.AncientSkillEffects, true)
	}
	//秘籍
	if len(hero.MiJiEffects) > 0 {
		this.initEffect(actor, hero.MiJiEffects, false)
	}
	//传世强化套装属性
	if len(hero.ChuanShiStrengthenEffects) > 0 {
		this.initEffect(actor, hero.ChuanShiStrengthenEffects, true)
	}
	//战宠附体
	if len(hero.PetAppendageEffects) > 0 {
		this.initEffect(actor, hero.PetAppendageEffects, false)
	}
	//头衔
	if len(hero.LabelEffects) > 0 {
		this.initEffect(actor, hero.LabelEffects, true)
	}
	//月卡
	if len(hero.MonthCardEffects) > 0 {
		this.initEffect(actor, hero.MonthCardEffects, false)
	}
	//特权
	if len(hero.PrivilegeEffects) > 0 {
		this.initEffect(actor, hero.PrivilegeEffects, false)
	}
}

func (this *Fight) createFightUserStageBuff(hero *objs.Hero, actor *pbserver.ActorHero, stageId int) {
	mapEffectSlice := []int{stageId, -1}
	for _, k := range mapEffectSlice {
		if mapEffectData, ok := hero.MapEffects[k]; ok && len(mapEffectData) > 0 {
			effectSlice := make([]int, 0)
			for effectId := range mapEffectData {
				effectSlice = append(effectSlice, effectId)
			}
			this.initEffect(actor, effectSlice, true)
		}
	}
}

func (this *Fight) UpdateUserElf(user *objs.User) {

	if user.FightId <= 0 {
		return
	}

	request := &pbserver.GsToFsUpdateUserElfReq{
		UserId: int32(user.Id),
		Elf:    this.createUserElfInfo(user),
	}
	replay := &pbserver.FsToGsUpdateUserElfAck{}
	err := this.FSRpcCall(user.FightId, user.FightStageId, request, replay)
	if err != nil {
		logger.Error("实时更新玩家精灵数据异常：%v", err)
	}

}

func (this *Fight) UpdateUserInfoToFight(user *objs.User, heroIndexs map[int]bool, updateNow bool) {

	if user.FightId <= 0 {
		return
	}
	//主城 只更新主武将
	if user.FightStageId == constFight.FIGHT_TYPE_MAIN_CITY_STAGE {
		for k, _ := range heroIndexs {
			if k != constUser.USER_HERO_MAIN_INDEX {
				delete(heroIndexs, k)
			}
		}
	}

	fightUser := &pbserver.Actor{}

	//基础信息
	this.crateFightUserBaseInfo(user, fightUser, user.FightStageId)

	for heroIndex, _ := range heroIndexs {
		actorHero := &pbserver.ActorHero{}
		fightUser.Heros[int32(heroIndex)] = actorHero
		heroInfo := user.Heros[heroIndex]
		this.createFighUserHeroBaseInfo(user, heroInfo, actorHero, this.isCrossFight(user.FightStageId))
		//技能信息
		this.createFightUserSkillInfo(heroInfo, actorHero)
		//属性信息
		actorHero.Prop = this.createFightHeroProp(user, heroInfo)
		//出生buff
		this.createFightUserInitBuff(heroInfo, actorHero)
		//地图额外buff
		this.createFightUserStageBuff(heroInfo, actorHero, user.FightStageId)
	}

	if updateNow {

		replaymsg := &pbserver.FSUpdateUserInfoAck{}
		err := this.FSRpcCall(user.FightId, user.FightStageId, &pbserver.FSUpdateUserInfoNtf{
			UserInfo: fightUser,
		}, replaymsg)
		if err != nil {
			logger.Error("实时更新玩家数据异常：%v", err)
		}
	} else {
		this.FSSendMessage(user.FightId, user.FightStageId, &pbserver.FSUpdateUserInfoNtf{
			UserInfo: fightUser,
		})
	}
}

func (this *Fight) UpdateUserFightModel(user *objs.User) {

	if user.FightId <= 0 {
		return
	}

	msg := &pbserver.GSToFsUpdateUserFightModel{
		UserId:     int32(user.Id),
		FigthModel: int32(user.FightModel),
	}
	this.FSSendMessage(user.FightId, user.FightStageId, msg)

}

func (this *Fight) UserFitReq(user *objs.User) error {

	stageConf := gamedb.GetStageStageCfg(user.FightStageId)
	if stageConf == nil {
		return nil
	}
	mapTypeConf := gamedb.GetMaptypeGameCfg(stageConf.Type)
	if mapTypeConf.CanUseHt == 0 {
		return gamedb.ERRSKILLCANNOTUSE
	}

	request := &pbserver.GsToFsUseFitReq{
		UserId: int32(user.Id),
		Fit: &pbserver.ActorFit{
			Id:        int32(constFight.FIT_ID),
			Lv:        int32(user.Fit.Lv[constFight.FIT_ID]),
			FashionId: int32(user.Wear.FitFashionId),
			FashionLv: int32(user.Fit.Fashion[user.Wear.FitFashionId]),
			Skills:    make([]*pbserver.FitSkill, 0),
			Effect:    common.ConvertIntSlice2Int32Slice(user.FitHolyEquipEffects),
		},
	}
	for _, v := range user.Fit.SkillBag {
		skill := &pbserver.FitSkill{
			Id:   int32(v),
			Lv:   int32(user.Fit.Skills[v].Lv),
			Star: int32(user.Fit.Skills[v].Star),
		}
		request.Fit.Skills = append(request.Fit.Skills, skill)
	}

	for k, v := range user.Fit.Skills {
		skillConf := gamedb.GetFitSkillFitSkillCfg(k)
		if skillConf.Type == pb.FITSKILLTYPE_ZHUDONG {
			skill := &pbserver.FitSkill{
				Id:   int32(k),
				Lv:   int32(v.Lv),
				Star: int32(v.Star),
			}
			request.Fit.Skills = append(request.Fit.Skills, skill)
		}
	}

	replay := &pbserver.FsToGsUseFitAck{}
	err := this.FSRpcCall(user.FightId, user.FightStageId, request, replay)
	if err != nil {
		logger.Error("玩家：%v 使用合体异常，err:%v", user.IdName(), err)
		return err
	}
	return nil
}

func (this *Fight) UserFitCacelReq(user *objs.User) error {

	request := &pbserver.GsToFsFitCacelReq{
		UserId: int32(user.Id),
	}
	replay := &pbserver.FsToGsFitCacelAck{}
	err := this.FSRpcCall(user.FightId, user.FightStageId, request, replay)
	if err != nil {
		logger.Error("玩家：%v 取消合体异常，err:%v", user.IdName(), err)
		return err
	}
	return nil
}

/**
 *  @Description: 更新战宠
 *  @param user
 *  @return error
 */
func (this *Fight) UserUpdatePet(user *objs.User) error {

	if user.FightStageId == constFight.FIGHT_TYPE_MAIN_CITY_STAGE {
		return nil
	}
	request := &pbserver.GsToFsUpdatePetReq{
		UserId: int32(user.Id),
		Pet:    this.createUserPet(user),
	}
	replay := &pbserver.FsToGsUpdatePetAck{}
	err := this.FSRpcCall(user.FightId, user.FightStageId, request, replay)
	if err != nil {
		return err
	}
	return nil
}
