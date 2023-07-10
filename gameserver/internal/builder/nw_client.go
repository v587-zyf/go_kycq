package builder

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/prop"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/golibs/logger"
	"fmt"
	"time"

	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildEnterGameAck(user *objs.User, token string, openServerDay, openServerTime int, openDay, mergeOpenServerDay int, serverName string) *pb.EnterGameAck {
	return &pb.EnterGameAck{User: BuildUserLoginInfo(user, openDay),
		Ts:                 int32(time.Now().Unix()),
		Version:            "v0.1",
		OpenServerTime:     int32(openServerTime),
		OpenServerDay:      int32(openServerDay),
		MergeOpenServerDay: int32(mergeOpenServerDay),
		RealServerId:       int32(base.Conf.ServerId),
		RealServerName:     serverName,
	}
}

func BuildUserLoginInfo(user *objs.User, openDay int) *pb.UserLoginInfo {
	r := &pb.UserLoginInfo{}
	r.Userid = int32(user.Id)
	r.NickName = user.NickName
	r.Avatar = user.Avatar
	r.CreateTime = int32(user.CreateTime.Unix())
	r.VipLevel = int32(user.VipLevel)
	r.VipScore = int32(user.VipScore)
	r.Exp = int32(user.Exp)
	r.Gold = int64(user.Gold)
	r.Honour = int64(user.Honour)
	r.Ingot = int32(user.Ingot)
	r.ChuanqiBi = int32(user.ChuanqiBi)
	r.StageId = int32(user.StageId)
	r.StageWave = int32(user.StageWave)
	r.Combat = int64(user.Combat)
	r.Heros = BuildUserHerosInfo(user)
	r.Rein, r.ReinCost = BuildUserRein(user)
	r.Fabao = BuilderFabao(user)
	r.FieldBossInfo = BuilderFieldBoss(user)
	r.ArenaFightNum = int32(user.Arena.DareNum + user.Arena.BuyDareNums)
	r.FightModel = int32(user.FightModel)
	r.Task = &pb.TaskInfoNtf{
		TaskId:      int32(user.MainLineTask.TaskId),
		Process:     int32(user.MainLineTask.Process),
		MarkProcess: int32(user.MainLineTask.MarkProcess),
	}
	cfg := gamedb.GetTaskConditionCfg(user.MainLineTask.TaskId)
	if cfg != nil {
		if cfg.ConditionType == pb.CONDITION_ONE_HERO_LV && cfg.ConditionValue[0] == 1 {
			r.Task.Process = 1
		}
	}
	r.MaterialStage = BuildMaterialStage(user)
	r.Panaceas = BuilderPanacea(user)
	r.SignInfo = BuildSignInfo(user)
	r.DayStateInfo = BuildDayStateInfo(user)
	r.Official = int32(user.Official)
	r.Holy = BuilderHolyArms(user)
	r.Atlases = BuilderAtlas(user)
	r.AtlasGathers = BuilderAtlasGather(user)
	r.MiningWorkTime = int64(user.Mining.WorkTime)
	r.Miner = int32(user.Mining.Miner)
	r.ExpStage = BuilderExpStage(user)
	r.Pets = BuildPet(user)
	r.Juexues = BuilderUserJuexue(user)
	r.UserWear = BuildUserWear(user)
	r.DarkPalaceInfo = BuildDarkPalace(user)
	r.PersonBoss = BuildPersonBoss(user)
	r.VipBoss = BuildVipBoss(user)
	r.ShopInfo = BuildShop(user)
	r.VipGift = BuildVipGift(user)
	r.Fit = BuildFit(user)
	r.MonthCard = BuildMonthCard(user)
	r.FirstRecharge = BuildFirstRecharge(user)
	r.SpendRebates = BuildSpendRebates(user)
	r.DailyPack = BuildDailyPack(user)
	r.RedPacketGetNum = int32(user.RedPacketItem.PickNum)
	r.GrowFund = BuildGrowFund(user)
	r.WarOrder = BuildWarOrder(user)
	r.Elf = BuildElf(user)
	r.CutTreasureLv = int32(user.CutTreasure)
	r.FitHolyEquip = BuildFitHolyEquip(user)
	r.Recharge = BuildRecharge(user)
	r.IsFriendApply = BuildIsFriendApply(user)
	r.ContRecharge = BuildContRecharge(user)
	r.OpenGift = BuildOpenGift(user)
	r.VipCustomer = BuildVipCustomer(user)
	r.AncientBossInfo = BuildAncient(user)
	r.TitleList = BuildTitleList(user, false)
	r.PetAppendage = BuildPetAppendage(user)
	r.RechargeAll = int32(user.RechargeAll)
	r.RedPacketNum = int32(user.RedPacketNum)
	r.HellBossInfo = BuildHellBoss(user)
	r.DaBaoEquip = BuildDaBaoEquip(user)
	r.DaBaoMysteryEnergy = int32(user.DaBaoMystery.Energy)
	r.Label = BuildLabel(user)
	r.Privilege = BuildPrivilege(user)
	r.DailyRecharge = int32(user.DayStateRecord.DailyRecharge)

	allIds := make([]int32, 0)
	for _, ids := range user.AccumulativeId {
		allIds = append(allIds, int32(ids))
	}
	r.AccumulativeAllGetIds = allIds
	if openDay == 1 {
		r.IsHaveGetDailyCompetitveReward = 1
		user.CompetitiveInfo.BeforeDayRewardGetState = 1
	} else {
		r.IsHaveGetDailyCompetitveReward = int32(user.CompetitiveInfo.BeforeDayRewardGetState)
	}
	r.ShaBakeIsEnd = int32(rmodel.Shabake.GetShaBakeIsEnd(base.Conf.ServerId))
	r.CrossShabakeIsEnd = int32(rmodel.Shabake.GetCrossShaBakeIsEnd(base.Conf.ServerId))
	r.BindingIngot = int32(user.BindingIngot)
	r.HookupTime = int32(user.HookMapTime)
	r.HaveUseRecharge = int32(user.HaveUseRecharge)
	r.GoldIngot = int32(user.GoldIngot)
	hookupbag := make([]*pb.ItemUnit, len(user.HookMapBag))
	for k, v := range user.HookMapBag {
		hookupbag[k] = &pb.ItemUnit{ItemId: int32(v.ItemId), Count: int64(v.Count)}
	}
	r.HookupBag = hookupbag
	r.AppletsEnergy = int32(user.AppletsInfo.Energy)
	r.AppletsResumeTime = int32(user.AppletsInfo.ResumeTime)
	r.AppletsInfos = make(map[int32]*pb.AppletsInfo)
	for k, v := range user.AppletsInfo.List {
		r.AppletsInfos[int32(k)] = &pb.AppletsInfo{StageId: int32(v.Stage)}
	}
	r.UseRedPacketNum = int32(user.RedPacketUseNum)
	return r
}

func BuildUserRein(user *objs.User) (*pb.Rein, []*pb.ReinCost) {
	rein := &pb.Rein{
		Id:  int32(user.Rein.Id),
		Exp: int64(user.Rein.Exp),
	}
	reinCost := make([]*pb.ReinCost, 0)
	for _, v := range user.ReinCosts {
		reinCost = append(reinCost, &pb.ReinCost{
			Id:  int32(v.Id),
			Num: int32(v.Num),
		})
	}
	return rein, reinCost
}

func BuildUserHerosInfo(user *objs.User) []*pb.HeroInfo {

	heros := make([]*pb.HeroInfo, 0)
	for _, v := range user.Heros {
		heros = append(heros, BuildHeroInfo(v))
	}
	return heros
}

func BuildHeroInfo(hero *objs.Hero) *pb.HeroInfo {
	heroInfo := &pb.HeroInfo{
		Index:              int32(hero.Index),
		Sex:                int32(hero.Sex),
		Job:                int32(hero.Job),
		ExpLvl:             int32(hero.ExpLvl),
		Equips:             BuildEquiqs(hero),
		EquipGrids:         BuilderEquipStrength(hero),
		HeroProp:           BuildHeroProp(hero),
		Wing:               BuildWings(hero),
		WingSpecial:        BuilderWingSpecials(hero),
		Zodiacs:            BuilderSpecialEquips(hero, pb.ITEMTYPE_ZODIAC),
		Kingarms:           BuilderSpecialEquips(hero, pb.ITEMTYPE_KINGARMS),
		Dictates:           BuilderDictates(hero),
		Jewels:             BuilderJewels(hero),
		Name:               hero.Name,
		InsideInfo:         BuildInsideInfo(hero),
		Fashions:           BuilderFashion(hero),
		Wears:              BuildWear(hero),
		Rings:              BuildRings(hero),
		Skills:             BuildSkills(hero, pb.SKILLTYPE_ORDINARY),
		SkillBag:           BuildSkillBag(hero, pb.SKILLTYPE_ORDINARY),
		UniqueSkills:       BuildSkills(hero, pb.SKILLTYPE_UNIQUE),
		UniqueSkillBag:     BuildSkillBag(hero, pb.SKILLTYPE_UNIQUE),
		GodEquips:          BuilderHeroGodEquip(hero),
		Area:               BuildArea(hero),
		EquipClears:        BuildEquipClears(hero),
		DragonEquip:        BuildDragonEquip(hero),
		MagicCircle:        BuildMagicCircle(hero),
		Talents:            BuildTalent(hero),
		ChuanShiEquip:      BuildChuanShiEquip(hero),
		AncientSkill:       BuildAncientSkill(hero),
		ChuanShiStrengthen: BuildChuanShiStrengthen(hero),
	}
	return heroInfo
}

func BuildCreateUserAck(user *objs.User, openDay int) *pb.CreateUserAck {
	return &pb.CreateUserAck{User: BuildUserLoginInfo(user, openDay)}
}

func BuildEquiqs(hero *objs.Hero) map[int32]*pb.EquipUnit {

	equips := make(map[int32]*pb.EquipUnit)
	for k, v := range hero.Equips {
		if v.ItemId > 0 {
			equips[int32(k)] = BuildPbEquipUnit(v)
		}
	}
	return equips
}

func BuildKickUserNtf(reason string) *pb.KickUserNtf {
	return &pb.KickUserNtf{
		Reason: reason,
	}
}

func BuildProperMsg(user *objs.User, heroIndex int) *pb.UserPropertyNtf {

	msg := &pb.UserPropertyNtf{
		HeroProps:  make(map[int32]*pb.HeroProp),
		UserCombat: int64(user.Combat),
	}
	for k, v := range user.Heros {
		if heroIndex > 0 && heroIndex != k {
			continue
		}
		msg.HeroProps[int32(k)] = BuildHeroProp(v)
	}
	return msg
}

func BuildHeroProp(hero *objs.Hero) *pb.HeroProp {

	heroProp := &pb.HeroProp{
		Props:         hero.Prop.BuildClient(),
		ModulesCombat: BuildModulesCombat(hero),
	}
	return heroProp
}

func BuildModulesCombat(hero *objs.Hero) map[int32]int64 {
	r := make(map[int32]int64, len(hero.ModuleCombat))
	for k, v := range hero.ModuleCombat {
		r[int32(k)] = int64(v)
	}
	return r
}

func BuildRobotUserInfo(robotId, serverId int) (*pb.BriefUserInfo, int32, []*pb.HeroInfo) {
	robotInfo := gamedb.GetRobotRobotCfg(robotId)
	if robotInfo == nil {
		logger.Error("获取假人信息失败  robot表 id:%v", robotId)
		return nil, 0, nil
	}
	userInfo := &pb.BriefUserInfo{}
	userInfo.Combat = GetCompetitiveRobotCombat(robotInfo.Id)
	userInfo.Id = int32(-robotInfo.Id)
	userInfo.Job = int32(robotInfo.Job[0])
	userInfo.Name = robotInfo.Name
	userInfo.Sex = int32(robotInfo.Gender[0])
	userInfo.Lvl = int32(robotInfo.Level)
	userInfo.Avatar = fmt.Sprintf("%v", robotInfo.Icon[0])
	userInfo.Vip = 0
	return userInfo, int32(robotInfo.LevelCompetive), BuildRobotUserHerosInfo(robotInfo.Id)
}

func GetCompetitiveRobotCombat(robotId int) int64 {

	allCombat := 0
	cfg := gamedb.GetRobotRobotCfg(robotId)
	if cfg == nil {
		return int64(allCombat)
	}
	prop := make(map[int]int, 0)
	for _, job := range cfg.Job {
		if job == pb.JOB_ZHANSHI {
			prop = cfg.Property1
		} else if job == pb.JOB_FASHI {
			prop = cfg.Property2
		} else if job == pb.JOB_DAOSHI {
			prop = cfg.Property3
		}
		allCombat += CalcCombat(job, -1, -1, prop, nil)
	}
	return int64(allCombat)
}

func BuildRobotUserHerosInfo(robotId int) []*pb.HeroInfo {

	heros := make([]*pb.HeroInfo, 0)
	cfg := gamedb.GetRobotRobotCfg(robotId)
	if cfg == nil {
		return heros
	}

	gender := 1
	for index, job := range cfg.Job {
		if len(cfg.Gender) >= index+1 {
			gender = cfg.Gender[index]
		}
		heros = append(heros, BuildRobotHeroInfo(index+1, job, gender, cfg))
	}
	return heros
}

func BuildRobotHeroInfo(index, job, sex int, cfg *gamedb.RobotRobotCfg) *pb.HeroInfo {

	skillIds := make([]int, 0)
	if cfg.Skills[index-1] != nil {
		skillIds = cfg.Skills[index-1]
	}
	heroInfo := &pb.HeroInfo{
		Index:    int32(index),
		Sex:      int32(sex),
		Job:      int32(job),
		ExpLvl:   int32(cfg.Level),
		HeroProp: BuildRobotHeroProp(job, cfg),
		Name:     GetHeroDefName(sex, job),
		Wears:    &pb.Wears{FashionWeaponId: int32(cfg.Model1[index-1]), FashionClothId: int32(cfg.Model2[index-1])},
		Skills:   BuildRobotSkills(skillIds),
	}
	return heroInfo
}

func BuildRobotHeroProp(job int, cfg *gamedb.RobotRobotCfg) *pb.HeroProp {

	m := make(map[int32]int64, len(pb.PROPERTY_ARRAY))
	prop := make(map[int]int, 0)
	if job == pb.JOB_ZHANSHI {
		prop = cfg.Property1
	} else if job == pb.JOB_FASHI {
		prop = cfg.Property2
	} else if job == pb.JOB_DAOSHI {
		prop = cfg.Property3
	}
	combat := CalcCombat(job, -1, -1, prop, nil)

	m[int32(pb.PROPERTY_COMBAT)] = int64(combat)
	heroProp := &pb.HeroProp{
		Props:         m,
		ModulesCombat: make(map[int32]int64),
	}
	return heroProp
}

func BuildRobotSkills(skills []int) []*pb.SkillUnit {

	pbSkillUnits := make([]*pb.SkillUnit, 0)
	for _, id := range skills {
		skillId, skillLv := gamedb.GetSkillIdAndLv(id)
		cfg := gamedb.GetSkillLvConf(skillId, skillLv)
		if cfg == nil {
			continue
		}
		pbSkillUnits = append(pbSkillUnits, &pb.SkillUnit{SkillId: int32(skillId), Level: int32(skillLv), StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + int64(cfg.CD)})
	}
	return pbSkillUnits
}

func GetRobotName(robotName string, serverId int) string {

	serverIndex := base.Conf.ServerId
	return fmt.Sprintf("%d.%s", serverIndex, robotName)
}

func CalcCombat(heroIndex, itemId, bagIndex int, fixProp map[int]int, randomProp []*model.EquipRandProp) int {

	prop1 := prop.NewProp()
	allProp := make(map[int]int)
	if fixProp != nil {
		for k, v := range fixProp {
			allProp[k] += v
		}
	}
	randProp := make(map[int]int)
	logger.Debug("固定属性 fixProp:%v", fixProp)

	if randomProp != nil {
		for _, data := range randomProp {
			allProp[data.PropId] += data.Value
			randProp[data.PropId] += data.Value
		}
	}
	logger.Debug("装备战力计算    heroIndex:%v bagIndex:%v   itemId:%v  allProp:%v fixProp:%v  randProp:%v", heroIndex, bagIndex, itemId, allProp, fixProp, randProp)
	prop1.Add(allProp)
	//p.Calc(constUser.USER_HERO_MAIN_INDEX)
	prop1.Calc(heroIndex)
	logger.Debug("calcCombat combat:%v", prop1.Combat)
	return prop1.Combat
}

func GetHeroDefName(sex int, job int) string {
	nameStr := "男"
	if sex == pb.SEX_FEMALE {
		nameStr = "女"
	}
	switch job {
	case pb.JOB_ZHANSHI:
		nameStr += gamedb.JOBZHANSHINAME
	case pb.JOB_FASHI:
		nameStr += gamedb.JOBFSHINAME
	case pb.JOB_DAOSHI:
		nameStr += gamedb.JOBDAOSHINAME
	}
	return nameStr
}
