package guild

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func NewGuildManager(m managersI.IModule) *GuildManager {
	guild := &GuildManager{}
	guild.IModule = m
	guild.guildInfo = make(map[int]*modelGame.Guild)
	guild.guildActivity = make(map[int]map[int]int)
	return guild
}

type GuildManager struct {
	util.DefaultModule
	managersI.IModule
	sync.RWMutex
	guildInfo map[int]*modelGame.Guild

	guildActivity map[int]map[int]int //公会id,活动id,状态
}

func (this *GuildManager) LoadGuild(user *objs.User, ack *pb.GuildLoadInfoAck) error {

	if this.GetExpPool().GetHeroMaxLv(user) < gamedb.GetConf().GuildOpenLv {
		return gamedb.ERRGUILDISNOTOPEN
	}

	if user.GuildData.NowGuildId <= 0 {
		return nil
	}

	guildInfo := this.GetGuild().GetGuildInfo(user.GuildData.NowGuildId)
	if guildInfo == nil {
		logger.Error("GetGuildInfo nil userId:%v  guildId:%v", user.Id, user.GuildData.NowGuildId)
		return gamedb.ERRHAVENOGUILD
	}

	ack.GuildInfo = this.BuildGuildInfo(user, guildInfo)
	return nil
}

func (this *GuildManager) CreateGuild(user *objs.User, op *ophelper.OpBagHelperDefault, req *pb.CreateGuildReq, ack *pb.CreateGuildAck) error {
	if user.GuildData.NowGuildId > 0 && user.GuildData.MyCreateId > 0 {
		return gamedb.ERRCREATEGUILDERR
	}

	err := base.CheckName(req.GuildName)
	if err != nil {
		return err
	}

	data, err1 := modelGame.GetGuildModel().GetAllGuildInfoByGuildName(req.GuildName)

	if err1 != nil {
		logger.Error("CreateGuild err:%v", err1)
		return err1
	}
	if len(data) > 0 {
		return gamedb.ERRGUILDNAME
	}

	guildInfo := &modelGame.Guild{
		GuildName:   req.GuildName,
		SettingId:   4,
		ChairmanId:  user.Id,
		Notice:      req.Notice,
		CreatedAt:   time.Now(),
		IsDelete:    0,
		Creator:     user.Id,
		Positions:   make(model.IntSlice, 0),
		ApplyList:   make(model.IntKv),
		DonateUsers: make(model.IntSlice, 0),
		DonateTimes: make(model.IntKv),
		JoinCombat:  gamedb.GetConf().GuildCombat,
		ServerId:    user.ServerId,
	}
	guildInfo.Positions = append(guildInfo.Positions, user.Id, pb.GUILDPOSITION_HUIZHANG)
	cfg := gamedb.GetConf().CreateGuild
	if cfg.ItemId == 0 || cfg.Count == 0 {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("CreateGuild")
	}

	enough, _ := this.GetBag().HasEnough(user, cfg.ItemId, cfg.Count)
	if !enough {
		return gamedb.ERRNOTENOUGHGOODS
	}
	err = this.GetBag().Remove(user, op, cfg.ItemId, cfg.Count)
	if err != nil {
		return err
	}

	err = modelGame.GetGuildModel().Create(guildInfo)
	if err != nil {
		logger.Error("创建门派失败 userId:%v", user.Id)
		return err
	}
	this.UpdateUserGuildDataInfo(user, guildInfo.GuildId, pb.GUILDPOSITION_HUIZHANG, false, true, true, false)
	this.SetGuildInfo(guildInfo)
	kyEvent.GuildCreate(user, guildInfo.Id, guildInfo.GuildName)
	ack.Success = true
	ack.GuildInfo = this.BuildGuildInfo(user, guildInfo)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ADD_GUILD, []int{1})
	this.GetCondition().RecordCondition(user, pb.CONDITION_JOIN_GUILD, []int{1})
	return nil
}

//
//  SetJoinGuildCombatLimit
//  @Description: 申请加入门派的战力限制修改  and  是否自动同意修改
//  @param combat
//
func (this *GuildManager) SetJoinGuildCombatLimit(user *objs.User, ack *pb.JoinGuildCombatLimitAck, combat, isAgree int) error {

	if user.GuildData.NowGuildId <= 0 {
		ack.Success = false
		return nil
	}
	if user.GuildData.Position > pb.GUILDPOSITION_ZHANGLAO || user.GuildData.Position < pb.GUILDPOSITION_HUIZHANG {
		logger.Error("玩家权限不足 userId:%v", user.Id)
		return gamedb.ERRNOPOWER
	}

	guildInfo := this.GetGuild().GetGuildInfo(user.GuildData.NowGuildId)
	if guildInfo == nil {
		logger.Error("GetGuildInfo nil userId:%v  guildId:%v", user.Id, user.GuildData.NowGuildId)
		return gamedb.ERRHAVENOGUILD
	}
	guildInfo.AutoAgree = isAgree
	guildInfo.JoinCombat = combat
	err := modelGame.GetGuildModel().Update(guildInfo)
	if err != nil {
		logger.Error("GetGuildInfo  update  userId:%v  guildId:%v  err:%v", user.Id, user.GuildData.NowGuildId, err)
		return err
	}
	this.SetGuildInfo(guildInfo)
	ack.Success = true
	ack.LimitCombat = int64(combat)
	ack.IsAgree = int32(isAgree)
	return nil
}

//
//  ApplyJoinGuild
//  @Description:申请加入门派
//
func (this *GuildManager) ApplyJoinGuild(user *objs.User, appGuildId int, ack *pb.ApplyJoinGuildAck) error {

	if user.GuildData.NowGuildId > 0 {
		return gamedb.ERRHAVEOWNGUILD
	}
	err := this.OperationGuildCheck1(user)
	if err != nil {
		return err
	}

	guildInfo := this.GetGuild().GetGuildInfo(appGuildId)
	if guildInfo == nil || guildInfo.IsDelete > 0 {
		logger.Error("GetGuildInfo nil userId:%v  guildId:%v", user.Id, appGuildId)
		return gamedb.ERRNOGUILD
	}
	isInGuild := this.CheckUserIsInGuild(user.Id, guildInfo.Positions)
	if isInGuild {
		return gamedb.ERRHAVEINGUILD
	}

	_, limitNum := gamedb.GetGuildMemberLimit(guildInfo.GuildContributionValue)
	memNum := len(guildInfo.Positions) / 2
	if memNum+1 > limitNum {
		return gamedb.ERRGUILDNUMBEROVER
	}

	if guildInfo.AutoAgree == 1 {
		if user.Combat < guildInfo.JoinCombat {
			logger.Error("玩家战力:%v 未达到加入门派的战力:%v限制", user.Combat, guildInfo.JoinCombat)
			return gamedb.ERRNOTENOUGHCOMBAT
		}
	}

	if user.GuildData.JoinCD > int(time.Now().Unix()) {
		return gamedb.ERRJOINGUILDCD
	}

	if rmodel.Guild.IsExistUserId(appGuildId, user.Id) {
		return gamedb.ERRHAVEAPPLY
	}

	//管理员设置了自动同意
	guildMemberInfos := make([]*pb.GuildMenberInfo, 0)
	infos := make([]int32, 0)
	if guildInfo.AutoAgree == 1 {
		guildInfo.Positions = append(guildInfo.Positions, user.Id, pb.GUILDPOSITION_CHENGYUAN)
		err := modelGame.GetGuildModel().Update(guildInfo)
		if err != nil {
			logger.Error("GetGuildInfo  update  userId:%v  guildId:%v  err:%v", user.Id, user.GuildData.NowGuildId, err)
			return err
		}
		this.SetGuildInfo(guildInfo)
		this.UpdateUserGuildDataInfo(user, appGuildId, pb.GUILDPOSITION_CHENGYUAN, true, true, false, false)
		this.GetCondition().RecordCondition(user, pb.CONDITION_JOIN_GUILD, []int{1})
		ack.Success = true
		_ = this.GetUserManager().SendMessage(user, &pb.JoinGuildSuccessNtf{UserId: int32(user.Id), GuildId: int32(appGuildId), Success: true}, true)
		infos = append(infos, int32(user.Id), pb.GUILDPOSITION_CHENGYUAN)
		guildMemberInfos = append(guildMemberInfos, this.GetMemberInfo(int32(user.Id), int32(pb.GUILDPOSITION_CHENGYUAN), int32(0), int32(user.GuildData.GuildCapital), int32(user.GuildData.ContributionValue),
			int32(this.GetExpPool().GetHeroMaxLv(user)), int32(user.Combat), user.NickName, user.Avatar, int32(user.Heros[constUser.USER_HERO_MAIN_INDEX].Job), int32(user.Heros[constUser.USER_HERO_MAIN_INDEX].Sex)))

		this.BroadcastGuildInfoNtf(user.GuildData.NowGuildId, pb.GUILDBROCASTTYPE_JOIN_GUILD, infos, guildMemberInfos)
		return nil
	}
	rmodel.Guild.SetGuildApplyUserId(appGuildId, user.Id)

	this.RedBotDispose(appGuildId)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ADD_GUILD, []int{1})
	return nil
}

func (this *GuildManager) QuitGuild(user *objs.User, ack *pb.QuitGuildAck) error {
	if user.GuildData.NowGuildId <= 0 {
		return gamedb.ERRHAVENOGUILD
	}
	err := this.OperationGuildCheck1(user)
	if err != nil {
		return err
	}

	err = this.UpdateGuildPosition(user.GuildData.NowGuildId, user.Id)
	if err != nil {
		return err
	}
	infos := make([]int32, 0)
	infos = append(infos, int32(user.Id), 0)
	this.BroadcastGuildInfoNtf(user.GuildData.NowGuildId, pb.GUILDBROCASTTYPE_APPLY_QUIT, infos, nil)
	this.UpdateUserGuildDataInfo(user, 0, 0, true, true, false, false)
	ack.Success = true
	return nil
}

//任命
func (this *GuildManager) GuildAssign(user *objs.User, ack *pb.GuildAssignAck, assUserId, applyPosition int) error {

	guildInfo := this.GetGuild().GetGuildInfo(user.GuildData.NowGuildId)
	if guildInfo == nil {
		logger.Error("GetGuildInfo nil userId:%v  guildId:%v", user.Id, user.GuildData.NowGuildId)
		return gamedb.ERRNOGUILD
	}

	data := gamedb.GetGuildGuildCfg(user.GuildData.Position)
	if data != nil {
		if data.ChangePosition != 1 {
			return gamedb.ERRNOPOWER
		}
	}
	isNeedUpdate := false
	assignUser := this.GetUserManager().GetUser(assUserId)
	if assignUser == nil {
		assignUser = this.GetUserManager().GetOfflineUserInfo(assUserId)
		isNeedUpdate = true
	}
	if assignUser == nil {
		return gamedb.ERRGETAPPLYUSERINFOERR
	}

	if assignUser.GuildData.Position == user.GuildData.Position {
		return gamedb.ERRNOPOWER
	}

	if applyPosition == pb.GUILDPOSITION_HUIZHANG {
		//转让会长
		if user.GuildData.Position != pb.GUILDPOSITION_HUIZHANG {
			return gamedb.ERRNOPOWER
		}
		err := this.makeOverHuiZ(user, assignUser, guildInfo, applyPosition, isNeedUpdate)
		if err != nil {
			return err
		}
	} else {

		if applyPosition <= user.GuildData.Position {
			return gamedb.ERRNOPOWER
		}

		count := gamedb.GetGuildGuildCfg(applyPosition).Count

		userPosition, positionCount, err := this.GetGuildPositionInfo(user.GuildData.NowGuildId)
		if err != nil {
			return err
		}
		haveNumber := int(positionCount[int32(applyPosition)])
		if count > 0 {
			if haveNumber >= count {
				return gamedb.ERRPOSITIONHAVEFULL
			}
		}

		assignUser.GuildData.Position = applyPosition
		userPosition[int32(assUserId)] = int32(applyPosition)
		positionSlice := make([]int, 0)
		for k, v := range userPosition {
			positionSlice = append(positionSlice, int(k), int(v))
		}

		guildInfo.Positions = positionSlice

		err = modelGame.GetGuildModel().Update(guildInfo)
		if err != nil {
			logger.Error("GetGuildModel().Update err:%v  guildId:%v", err, user.GuildData.NowGuildId)
			return gamedb.ERRGUILDASSIGN
		}
		this.SetGuildInfo(guildInfo)
		if isNeedUpdate {
			_, err := modelGame.GetUserModel().DbMap().Update(assignUser.User)
			if err != nil {
				logger.Error("存储离线玩家数据失败  userId:%v err:%v", assignUser.Id, err)
				return gamedb.ERRGUILDASSIGN
			}
		}
	}
	ack.Success = true
	ack.AssignUserId = int32(assUserId)
	ack.NowPosition = int32(applyPosition)
	ack.PositionCount = this.BuildGuildInfo(user, guildInfo).PositionCount
	infos := make([]int32, 0)
	infos = append(infos, int32(assUserId), ack.NowPosition)
	this.BroadcastGuildInfoNtf(user.GuildData.NowGuildId, pb.GUILDBROCASTTYPE_APPLY_USER, infos, nil)
	this.GetCondition().RecordCondition(user, pb.CONDITION_BECOME_GUILD_ELDERS, []int{0})
	return nil
}

//转让会长
func (this *GuildManager) makeOverHuiZ(user *objs.User, assignUser *objs.User, guildInfo *modelGame.Guild, applyPosition int, isNeedUpdate bool) error {

	positionSlice := make([]int, 0)
	for i, j := 0, len(guildInfo.Positions); i < j; i += 2 {
		userId := guildInfo.Positions[i]
		pos := guildInfo.Positions[i+1]
		//logger.Debug("makeOverHuiZ userId:%v pos:%v", userId, pos)
		if userId == guildInfo.ChairmanId {
			pos = pb.GUILDPOSITION_CHENGYUAN
		}
		if userId == assignUser.Id {
			pos = pb.GUILDPOSITION_HUIZHANG
		}
		positionSlice = append(positionSlice, userId, pos)
	}

	guildInfo.Positions = positionSlice
	guildInfo.ChairmanId = assignUser.Id
	this.GetGuild().SetGuildInfo(guildInfo)
	user.GuildData.Position = pb.GUILDPOSITION_CHENGYUAN
	assignUser.GuildData.Position = applyPosition
	if isNeedUpdate {
		_, err := modelGame.GetUserModel().DbMap().Update(assignUser.User)
		if err != nil {
			logger.Error("存储离线玩家数据失败  userId:%v err:%v", assignUser.Id, err)
			return gamedb.ERRGUILDASSIGN
		}
	}
	user.Dirty = true
	assignUser.Dirty = true

	infos := make([]int32, 0)
	infos = append(infos, int32(user.Id), pb.GUILDPOSITION_CHENGYUAN, int32(assignUser.Id), pb.GUILDPOSITION_HUIZHANG)
	this.BroadcastGuildInfoNtf(user.GuildData.NowGuildId, pb.GUILDBROCASTTYPE_ZHUAN_RANG_HUI_ZHANG, infos, nil)
	return nil
}

//
//  JoinGuildDispose
//  @Description:处理玩家申请列表
//
func (this *GuildManager) JoinGuildDispose(user *objs.User, ack *pb.JoinGuildDisposeAck, isAgree bool, applyUserId int) error {

	if user.GuildData.NowGuildId <= 0 {
		return gamedb.ERRHAVENOGUILD
	}

	if isAgree {
		err := this.OperationGuildCheck1(user)
		if err != nil {
			return err
		}
	}

	data := gamedb.GetGuildGuildCfg(user.GuildData.Position)
	if data != nil {
		if data.ApplyMassage != 1 {
			return gamedb.ERRNOPOWER
		}
	}
	infos := make([]int32, 0)
	guildInfo := this.GetGuild().GetGuildInfo(user.GuildData.NowGuildId)

	if guildInfo == nil {
		logger.Error("GetGuildInfo  nil  userId:%v  appJoinGuildId:%v ", user.Id, user.GuildData.NowGuildId)
		return gamedb.ERRNOGUILD
	}
	isInGuild := this.CheckUserIsInGuild(applyUserId, guildInfo.Positions)
	if isInGuild {
		return gamedb.ERRHAVEINGUILD
	}
	logger.Debug(" guildInfo.Positions:%v", guildInfo.Positions)
	success := false
	guildMemberInfos := make([]*pb.GuildMenberInfo, 0)
	if isAgree {
		appUserInfo := this.GetUserManager().GetUser(applyUserId)
		if appUserInfo != nil {

			if appUserInfo.GuildData.NowGuildId > 0 {
				//清空该玩家申请信息
				_ = rmodel.Guild.DelApplyUserId(user.GuildData.NowGuildId, applyUserId)
			} else {
				if appUserInfo.GuildData.JoinCD > int(time.Now().Unix()) {
					return gamedb.ERRJOINGUILDCD
				}
			}

			if appUserInfo.GuildData.NowGuildId <= 0 {
				guildInfo.Positions = append(guildInfo.Positions, applyUserId, pb.GUILDPOSITION_CHENGYUAN)
				logger.Debug(" guildInfo.Positions:%v", guildInfo.Positions)
				success = true
				this.UpdateUserGuildDataInfo(appUserInfo, user.GuildData.NowGuildId, pb.GUILDPOSITION_CHENGYUAN, true, true, false, false)
				infos = append(infos, int32(applyUserId), pb.GUILDPOSITION_CHENGYUAN)
				guildMemberInfos = append(guildMemberInfos, this.GetMemberInfo(int32(applyUserId), int32(pb.GUILDPOSITION_CHENGYUAN), int32(0), int32(appUserInfo.GuildData.GuildCapital), int32(appUserInfo.GuildData.ContributionValue),
					int32(this.GetExpPool().GetHeroMaxLv(appUserInfo)), int32(appUserInfo.Combat), appUserInfo.NickName, appUserInfo.Avatar, int32(appUserInfo.Heros[constUser.USER_HERO_MAIN_INDEX].Job), int32(appUserInfo.Heros[constUser.USER_HERO_MAIN_INDEX].Sex)))
			} else {
				ack.IsHaveJoinGuild = true
			}
		} else {
			appUserInfo = this.GetUserManager().GetOfflineUserInfo(applyUserId)
			if appUserInfo == nil {
				logger.Error("玩家不存在 applyUserId:%v", applyUserId)
				return gamedb.ERRPARAM
			}

			if appUserInfo.GuildData.NowGuildId > 0 {
				//清空该玩家申请信息
				_ = rmodel.Guild.DelApplyUserId(user.GuildData.NowGuildId, applyUserId)
			} else {
				if appUserInfo.GuildData.JoinCD > int(time.Now().Unix()) {
					return gamedb.ERRJOINGUILDCD
				}
			}
			if appUserInfo.GuildData.NowGuildId <= 0 {
				guildInfo.Positions = append(guildInfo.Positions, applyUserId, pb.GUILDPOSITION_CHENGYUAN)
				logger.Debug(" guildInfo.Positions:%v", guildInfo.Positions)
				success = true
				this.UpdateUserGuildDataInfo(appUserInfo, user.GuildData.NowGuildId, pb.GUILDPOSITION_CHENGYUAN, true, true, false, true)
				infos = append(infos, int32(applyUserId), pb.GUILDPOSITION_CHENGYUAN)
				guildMemberInfos = append(guildMemberInfos, this.GetMemberInfo(int32(applyUserId), int32(pb.GUILDPOSITION_CHENGYUAN), int32(0), int32(appUserInfo.GuildData.GuildCapital), int32(appUserInfo.GuildData.ContributionValue),
					int32(this.GetExpPool().GetHeroMaxLv(appUserInfo)), int32(appUserInfo.Combat), appUserInfo.NickName, appUserInfo.Avatar, int32(appUserInfo.Heros[constUser.USER_HERO_MAIN_INDEX].Job), int32(appUserInfo.Heros[constUser.USER_HERO_MAIN_INDEX].Sex)))
			} else {
				ack.IsHaveJoinGuild = true
			}
		}
		if !ack.IsHaveJoinGuild {
			err := modelGame.GetGuildModel().Update(guildInfo)
			if err != nil {
				logger.Error("GetGuildInfo  update  userId:%v  guildId:%v  err:%v", user.Id, user.GuildData.NowGuildId, err)
				return err
			}
			this.SetGuildInfo(guildInfo)
		}

		this.GetCondition().RecordCondition(appUserInfo, pb.CONDITION_JOIN_GUILD, []int{1})
	}

	//清空该玩家申请信息
	_ = rmodel.Guild.DelApplyUserId(user.GuildData.NowGuildId, applyUserId)
	ack.ApplyUserInfo = this.GetApplyUserLists(user.GuildData.NowGuildId)
	ack.Success = success
	if success {
		this.BroadcastGuildInfoNtf(user.GuildData.NowGuildId, pb.GUILDBROCASTTYPE_JOIN_GUILD, infos, guildMemberInfos)
	}
	appUser := this.GetUserManager().GetUser(applyUserId)
	if appUser != nil {
		_ = this.GetUserManager().SendMessage(appUser, &pb.JoinGuildSuccessNtf{UserId: int32(applyUserId), GuildId: int32(appUser.GuildData.NowGuildId), Success: success}, true)
	}

	return nil
}

//
//  JoinGuildDispose
//  @Description:一键处理玩家申请列表
//
func (this *GuildManager) AllJoinGuildDispose(user *objs.User, ack *pb.AllJoinGuildDisposeAck, isAgree bool) error {
	if user.GuildData.NowGuildId <= 0 {
		return gamedb.ERRHAVENOGUILD
	}

	if isAgree {
		err := this.OperationGuildCheck1(user)
		if err != nil {
			return err
		}
	}

	data := gamedb.GetGuildGuildCfg(user.GuildData.Position)
	if data != nil {
		if data.ApplyMassage != 1 {
			return gamedb.ERRNOPOWER
		}
	}
	guildInfo := this.GetGuild().GetGuildInfo(user.GuildData.NowGuildId)
	if guildInfo == nil {
		logger.Error("GetGuildInfo  nil  userId:%v  appJoinGuildId:%v ", user.Id, user.GuildData.NowGuildId)
		return gamedb.ERRNOGUILD
	}
	_, limit := gamedb.GetGuildMemberLimit(guildInfo.GuildContributionValue)
	allCount := len(guildInfo.Positions) / 2
	infos := make([]int32, 0)
	applyUserList := this.GetApplyUserLists(user.GuildData.NowGuildId)
	isFullState := false
	guildMemberInfos := make([]*pb.GuildMenberInfo, 0)
	for _, appInfo := range applyUserList {
		applyUserId := int(appInfo.Id)
		isInGuild := this.CheckUserIsInGuild(applyUserId, guildInfo.Positions)
		if isInGuild {
			//清空该玩家申请信息
			_ = rmodel.Guild.DelApplyUserId(user.GuildData.NowGuildId, applyUserId)
			continue
		}
		logger.Debug(" guildInfo.Positions:%v", guildInfo.Positions)
		haveJoinGuild := false
		success := false
		if isAgree {
			if allCount >= limit {
				isFullState = true
			} else {
				appUserInfo := this.GetUserManager().GetUser(applyUserId)
				if appUserInfo != nil {
					if appUserInfo.GuildData.NowGuildId <= 0 {
						this.GetCondition().RecordCondition(appUserInfo, pb.CONDITION_JOIN_GUILD, []int{1})
						guildInfo.Positions = append(guildInfo.Positions, applyUserId, pb.GUILDPOSITION_CHENGYUAN)
						logger.Debug(" guildInfo.Positions:%v", guildInfo.Positions)
						this.UpdateUserGuildDataInfo(appUserInfo, user.GuildData.NowGuildId, pb.GUILDPOSITION_CHENGYUAN, true, true, false, false)
						allCount += 1
						success = true
						infos = append(infos, int32(applyUserId), pb.GUILDPOSITION_CHENGYUAN)
						guildMemberInfos = append(guildMemberInfos, this.GetMemberInfo(int32(applyUserId), int32(pb.GUILDPOSITION_CHENGYUAN), int32(0), int32(appUserInfo.GuildData.GuildCapital), int32(appUserInfo.GuildData.ContributionValue),
							int32(this.GetExpPool().GetHeroMaxLv(appUserInfo)), int32(appUserInfo.Combat), appUserInfo.NickName, appUserInfo.Avatar, int32(appUserInfo.Heros[constUser.USER_HERO_MAIN_INDEX].Job), int32(appUserInfo.Heros[constUser.USER_HERO_MAIN_INDEX].Sex)))
					} else {
						haveJoinGuild = true
					}
				} else {
					appUserInfo = this.GetUserManager().GetOfflineUserInfo(applyUserId)
					if appUserInfo.GuildData.NowGuildId <= 0 {
						this.GetCondition().RecordCondition(appUserInfo, pb.CONDITION_JOIN_GUILD, []int{1})
						guildInfo.Positions = append(guildInfo.Positions, applyUserId, pb.GUILDPOSITION_CHENGYUAN)
						logger.Debug(" guildInfo.Positions:%v", guildInfo.Positions)
						this.UpdateUserGuildDataInfo(appUserInfo, user.GuildData.NowGuildId, pb.GUILDPOSITION_CHENGYUAN, true, true, false, true)
						allCount += 1
						success = true
						infos = append(infos, int32(applyUserId), pb.GUILDPOSITION_CHENGYUAN)
						guildMemberInfos = append(guildMemberInfos, this.GetMemberInfo(int32(applyUserId), int32(pb.GUILDPOSITION_CHENGYUAN), int32(0), int32(appUserInfo.GuildData.GuildCapital), int32(appUserInfo.GuildData.ContributionValue),
							int32(this.GetExpPool().GetHeroMaxLv(appUserInfo)), int32(appUserInfo.Combat), appUserInfo.NickName, appUserInfo.Avatar, int32(appUserInfo.Heros[constUser.USER_HERO_MAIN_INDEX].Job), int32(appUserInfo.Heros[constUser.USER_HERO_MAIN_INDEX].Sex)))
					} else {
						haveJoinGuild = true
					}
				}
				if !haveJoinGuild {
					err := modelGame.GetGuildModel().Update(guildInfo)
					if err != nil {
						logger.Error("GetGuildInfo  update  userId:%v  guildId:%v  err:%v", user.Id, user.GuildData.NowGuildId, err)
						continue
					}
					this.SetGuildInfo(guildInfo)
				}
			}

		}
		//清空该玩家申请信息
		_ = rmodel.Guild.DelApplyUserId(user.GuildData.NowGuildId, applyUserId)
		appUser := this.GetUserManager().GetUser(applyUserId)
		if appUser != nil {
			_ = this.GetUserManager().SendMessage(appUser, &pb.JoinGuildSuccessNtf{UserId: int32(applyUserId), GuildId: int32(appUser.GuildData.NowGuildId), Success: success}, true)
		}
	}
	this.BroadcastGuildInfoNtf(user.GuildData.NowGuildId, pb.GUILDBROCASTTYPE_JOIN_GUILD, infos, guildMemberInfos)

	ack.IsFullState = isFullState
	ack.ApplyUserInfo = this.GetApplyUserLists(user.GuildData.NowGuildId)
	return nil
}

//申请列表玩家信息
func (this *GuildManager) GetAllApplyUserLists(user *objs.User, ack *pb.GetApplyUserListAck) error {

	//data := gamedb.GetGuildGuildCfg(user.GuildData.Position)
	//if data != nil {
	//	if data.ApplyMassage != 1 {
	//		return gamedb.ERRNOPOWER
	//	}
	//}

	ack.ApplyUserInfo = this.GetApplyUserLists(user.GuildData.NowGuildId)
	return nil
}

func (this *GuildManager) KickOut(user *objs.User, ack *pb.KickOutAck, kickUserId int) error {
	if user.GuildData.NowGuildId <= 0 {
		return gamedb.ERRHAVENOGUILD
	}
	err := this.OperationGuildCheck1(user)
	if err != nil {
		return err
	}

	data := gamedb.GetGuildGuildCfg(user.GuildData.Position)
	if data != nil {
		if data.OustGuild != 1 {
			logger.Debug("user.GuildData.Position:%v", user.GuildData.Position)
			return gamedb.ERRNOPOWER
		}
	}
	isOff := false
	appUserInfo := this.GetUserManager().GetUser(kickUserId)
	if appUserInfo == nil {
		appUserInfo = this.GetUserManager().GetOfflineUserInfo(kickUserId)
		isOff = true
	}

	if appUserInfo == nil {
		return gamedb.ERRPARAM
	}

	if appUserInfo.GuildData.Position <= user.GuildData.Position {
		logger.Debug("appUserInfo.GuildData.Position:%v <= user.GuildData.Position:%v", appUserInfo.GuildData.Position, user.GuildData.Position)
		return gamedb.ERRNOPOWER
	}

	err = this.UpdateGuildPosition(user.GuildData.NowGuildId, kickUserId)
	if err != nil {
		return err
	}

	this.UpdateUserGuildDataInfo(appUserInfo, 0, 0, true, true, false, isOff)
	kickUser := this.GetUserManager().GetUser(kickUserId)
	if kickUser != nil {
		this.GetUserManager().SendMessage(kickUser, &pb.JoinGuildSuccessNtf{UserId: int32(kickUserId), GuildId: 0, Success: false}, true)
	}
	ack.KickUserId = int32(kickUserId)
	infos := make([]int32, 0)
	infos = append(infos, int32(kickUserId), 0)
	this.BroadcastGuildInfoNtf(user.GuildData.NowGuildId, pb.GUILDBROCASTTYPE_KICK_USER, infos, nil)

	return nil

}

//
//  DissolveGuild
//  @Description:解散公会
//
func (this *GuildManager) DissolveGuild(user *objs.User, ack *pb.DissolveGuildAck) error {
	if user.GuildData.NowGuildId <= 0 {
		return gamedb.ERRHAVENOGUILD
	}

	err := this.OperationGuildCheck1(user)
	if err != nil {
		return err
	}
	data := gamedb.GetGuildGuildCfg(user.GuildData.Position)
	if data != nil {
		if data.OustGuild != 1 {
			return gamedb.ERRNOPOWER
		}
	}

	guildInfo := this.GetGuild().GetGuildInfo(user.GuildData.NowGuildId)
	if guildInfo == nil {
		logger.Error("GetGuildInfo guildId:%v ", user.GuildData.NowGuildId)
		return gamedb.ERRNOGUILD
	}

	guildInfo.IsDelete = 1
	infos := make([]int32, 0)

	for i, j := 0, len(guildInfo.Positions); i < j; i += 2 {
		userId := guildInfo.Positions[i]
		infos = append(infos, int32(userId), 0)
		info := this.GetUserManager().GetAllUserInfoIncludeOfflineUser(userId)
		this.UpdateUserGuildDataInfo(info, 0, 0, true, true, false, true)
		kickUser := this.GetUserManager().GetUser(userId)
		if kickUser != nil {
			this.GetUserManager().SendMessage(kickUser, &pb.JoinGuildSuccessNtf{UserId: int32(userId), GuildId: 0, Success: false}, false)
		}
	}

	guildInfo.Positions = make(model.IntSlice, 0)
	err = modelGame.GetGuildModel().Update(guildInfo)
	if err != nil {
		logger.Error("GetGuildInfo  update  userId:%v  guildId:%v  err:%v", user.Id, user.GuildData.NowGuildId, err)
		return err
	}
	this.SetGuildInfo(guildInfo)
	this.DelGuildInfo(guildInfo)
	for i, j := 0, len(infos); i < j; i += 2 {
		userId := int(infos[i])
		kickUser := this.GetUserManager().GetUser(userId)
		if kickUser != nil {
			this.GetUserManager().SendMessage(kickUser, &pb.BroadcastGuildChangeNtf{Types: pb.GUILDBROCASTTYPE_DEL_GUILD, UserInfos: infos}, true)
		}
	}
	return nil
}

//弹劾会长
func (this *GuildManager) ImpeachPresident(user *objs.User, ack *pb.ImpeachPresidentAck) error {
	//data := gamedb.GetGuildGuildCfg(user.GuildData.Position)
	//if data != nil {
	//	if data.OustGuild != 1 {
	//		return gamedb.ERRNOPOWER
	//	}
	//}

	if user.GuildData.NowGuildId <= 0 {
		return gamedb.ERRNOGUILD
	}
	guildInfo := this.GetGuild().GetGuildInfo(user.GuildData.NowGuildId)
	if guildInfo == nil {
		logger.Error("GetGuildInfo guildId:%v ", user.GuildData.NowGuildId)
		return gamedb.ERRNOGUILD
	}
	if guildInfo.ChairmanId < 0 {
		return gamedb.ERRGUILDCHAIRMANISONLINE
	}
	userInfo := this.GetUserManager().GetUser(guildInfo.ChairmanId)
	if userInfo != nil {
		logger.Error("guildInfo.ChairmanId:%v", guildInfo.ChairmanId)
		return gamedb.ERRGUILDCHAIRMANISONLINE
	}
	userInfo = this.GetUserManager().GetOfflineUserInfo(guildInfo.ChairmanId)
	if userInfo == nil {
		logger.Error("guildId:%v guildName:%v chairmanId:%v", guildInfo.Id, guildInfo.GuildName, guildInfo.ChairmanId)
		return gamedb.ERRGUILDCHAIRMANISONLINE
	}
	if time.Now().Unix()-userInfo.OfflineTime.Unix() < int64(gamedb.GetConf().Impeach*60) {
		return gamedb.ERRGUILDCHAIRMANISONLINE
	}

	if rmodel.Guild.IsExistGuildImpeachPresident(guildInfo.ChairmanId) {
		ack.Success = true
		return nil
	}
	rmodel.Guild.SetGuildImpeachPresident(guildInfo.ChairmanId, int(time.Now().Unix()))
	go this.CheckChairmanIsLoggedIn(guildInfo.ChairmanId, guildInfo.GuildId)
	ack.Success = true
	return nil
}

func (this *GuildManager) ModifyBulletin(user *objs.User, ack *pb.ModifyBulletinAck, notice string) error {

	data := gamedb.GetGuildGuildCfg(user.GuildData.Position)
	if data != nil {
		if data.OustGuild != 1 {
			return gamedb.ERRNOPOWER
		}
	}

	guildInfo := this.GetGuild().GetGuildInfo(user.GuildData.NowGuildId)
	if guildInfo == nil {
		logger.Error("GetGuildInfo guildId:%v ", user.GuildData.NowGuildId)
		return gamedb.ERRNOGUILD
	}

	_, notice = base.CensorAndReplace(notice)
	guildInfo.Notice = notice
	ack.Success = true
	_ = modelGame.GetGuildModel().Update(guildInfo)
	return nil
}

func (this *GuildManager) GetAllGuildInfo(user *objs.User, ack *pb.AllGuildInfosAck) error {
	this.RLock()
	defer this.RUnlock()
	guildInfos := this.guildInfo
	if guildInfos == nil {
		logger.Error("GetAllGuildInfo guildInfos:%v", guildInfos)
		return nil
	}
	for _, guildInfo := range guildInfos {
		if guildInfo.IsDelete > 0 {
			continue
		}
		huiZLv, huiZName := this.getHuiZBaseUserInfo(guildInfo.ChairmanId)
		guildLv, _ := gamedb.GetGuildMemberLimit(guildInfo.GuildContributionValue)
		ack.GuildInfo = append(ack.GuildInfo, &pb.GuildInfo{
			GuildId:        int32(guildInfo.GuildId),
			GuildName:      guildInfo.GuildName,
			GuildLv:        int32(guildLv),
			Combat:         int64(guildInfo.JoinCombat),
			IsAutoAgree:    int32(guildInfo.AutoAgree),
			GuildPeopleNum: int32(len(guildInfo.Positions) / 2),
			HuiZhangLv:     int32(huiZLv),
			HuiZhangName:   huiZName,
		})
	}
	return nil
}

func (this *GuildManager) getHuiZBaseUserInfo(userId int) (int, string) {
	huiZLv := 0
	huiZName := ""

	if userId < 0 {
		cfg := gamedb.GetGuildRobotGuildRobotCfg(-userId)
		if cfg != nil {
			huiZLv = cfg.Level
			huiZName = cfg.Name
		}
		return huiZLv, huiZName
	}

	huiZUserInfo := this.GetUserManager().GetUserBasicInfo(userId)
	if huiZUserInfo != nil {
		huiZLv = huiZUserInfo.Level
		huiZName = huiZUserInfo.NickName
	}
	return huiZLv, huiZName
}

func (this *GuildManager) GetGuildMemberInfo(guildId int) (error, map[int][]int, int, []int) {
	memberInfo := make(map[int][]int) //key:职位
	guildUserIds := make([]int, 0)
	allNumber := 0

	guildInfo := this.GetGuild().GetGuildInfo(guildId)
	if guildInfo == nil {
		logger.Error("GetGuildInfo guildId:%v ", guildId)
		return gamedb.ERRNOGUILD, nil, 0, guildUserIds
	}

	for i, j := 0, len(guildInfo.Positions); i < j; i += 2 {
		userId := guildInfo.Positions[i]
		position := guildInfo.Positions[i+1]
		guildUserIds = append(guildUserIds, userId)
		if memberInfo[position] == nil {
			memberInfo[position] = make([]int, 0)
		}
		allNumber++
		memberInfo[position] = append(memberInfo[position], userId)
	}
	return nil, memberInfo, allNumber, guildUserIds
}

func (this *GuildManager) CheckUserIsInGuild(applyUserId int, position []int) bool {

	for i, j := 0, len(position); i < j; i += 2 {
		userId := position[i]
		if userId == applyUserId {
			return true
		}
	}
	return false

}

func (this *GuildManager) UpdateGuildPosition(guildId, userId int) error {

	guildInfo := this.GetGuild().GetGuildInfo(guildId)
	if guildInfo == nil {
		logger.Error("GetGuildInfo guildId:%v ", guildId)
		return gamedb.ERRNOGUILD
	}
	newData := make(model.IntSlice, 0)
	for i, j := 0, len(guildInfo.Positions); i < j; i += 2 {
		id := guildInfo.Positions[i]
		position := guildInfo.Positions[i+1]
		if userId == id {
			continue
		}
		newData = append(newData, id, position)
	}
	guildInfo.Positions = newData
	this.GetGuild().SetGuildInfo(guildInfo)
	return nil
}

//红点处理
func (this *GuildManager) RedBotDispose(guildId int) {
	ntf := &pb.ApplyJoinGuildReDotNtf{}
	_, guildUsers, _, _ := this.GetGuildMemberInfo(guildId)
	for pos, data := range guildUsers {
		if pos < pb.GUILDPOSITION_CHENGYUAN && pos > 0 {
			for _, userId := range data {
				userInfo := this.GetUserManager().GetUser(userId)
				if userInfo != nil {
					this.GetUserManager().SendMessage(userInfo, ntf, false)
				}
			}

		}
	}
}

//判断沙巴克 跨服沙巴克 期间操作公会权限
func (this *GuildManager) OperationGuildCheck1(user *objs.User) error {
	err := this.GetShabake().JudgeIsOpen(user)
	if err == nil {
		if rmodel.Shabake.GetShaBakeIsEnd(base.Conf.ServerId) == 1 {
			return nil
		}
		return gamedb.ERRGUILD1
	}

	err = this.GetShaBaKeCross().JudgeCrossIsOpen(user)
	if err == nil {
		if rmodel.Shabake.GetCrossShaBakeIsEnd(base.Conf.ServerId) == 1 {
			return nil
		}
		return gamedb.ERRGUILD2
	}

	return nil
}

func (this *GuildManager) CreateRobotGuild() {
	openDay := this.GetSystem().GetServerOpenDaysByServerId(base.Conf.ServerId)
	data, err := modelGame.GetGuildModel().GetAllGuildInfoBySystem()
	if err != nil {
		logger.Error("GetAllGuildInfo  err:%v", err)
		return
	}
	if len(data) > 0 || openDay > 1 {
		return
	}
	allId := 0
	allAutoCreate := gamedb.GetAllGuildAutoCreateGuildAutoCreateCfg()
	allGuildName := gamedb.GetAllGuildNameGuildNameCfgs()
	robotGuildName := make(map[int]bool)
	guildName := ""
	for i := 1; i <= len(allAutoCreate); i++ {
		robotId := 0
		autoCreateCfg := gamedb.GetGuildAutoCreateGuildAutoCreateCfg(i)
		if autoCreateCfg == nil {
			continue
		}
		rand.Seed(time.Now().Unix() + int64(i))
		num := rand.Intn(len(allGuildName)) + 1
		if robotGuildName[num] {
			for i1 := 1; i1 <= len(allGuildName); i1++ {
				if !robotGuildName[i1] {
					num = i1
					break
				}
			}
		}

		cfg := gamedb.GetGuildNameGuildNameCfg(num)
		if cfg != nil {
			if robotGuildName[num] {
				guildName = fmt.Sprintf("%v%v", cfg.Name, i)
			} else {
				guildName = cfg.Name
			}
		}
		if guildName == "" {
			continue
		}
		guildInfo := &modelGame.Guild{GuildName: guildName, SettingId: 4, ChairmanId: -1, Notice: gamedb.GetConf().GuildNotice, CreatedAt: time.Now(), IsDelete: 0, Creator: -1, Positions: make(model.IntSlice, 0), ApplyList: make(model.IntKv), DonateUsers: make(model.IntSlice, 0),
			DonateTimes: make(model.IntKv), JoinCombat: gamedb.GetConf().GuildCombat, ServerId: base.Conf.ServerId, IsSystem: 1, AutoAgreeJoin: 1, AutoAgree: 1}
		userPosInfoByLv := make(model.GuildRobotUserInfoSlice, 0)

		for i1 := robotId + 1; i1 <= autoCreateCfg.RobotNum; i1++ {
			cfgId := i1 + allId
			cfg := gamedb.GetGuildRobotGuildRobotCfg(cfgId)
			if cfg == nil {
				continue
			}
			userPosInfoByLv = append(userPosInfoByLv, &model.GuildRobotUserInfo{UserId: -(cfgId), Lv: cfg.Level})
		}
		sort.Sort(userPosInfoByLv)

		huiz := -1
		for i3 := pb.GUILDPOSITION_HUIZHANG; i3 <= pb.GUILDPOSITION_CHENGYUAN; i3++ {
			cfg := gamedb.GetGuildGuildCfg(i3)
			if cfg == nil {
				continue
			}
			addNum := 0
			for i2 := robotId; i2 < len(userPosInfoByLv); i2++ {
				if cfg.Count > 0 {
					if addNum >= cfg.Count {
						continue
					}
				}
				if i3 == pb.GUILDPOSITION_HUIZHANG {
					huiz = userPosInfoByLv[i2].UserId
				}
				guildInfo.Positions = append(guildInfo.Positions, userPosInfoByLv[i2].UserId, i3)
				robotId++
				addNum++
				allId++
			}
		}
		guildInfo.ChairmanId = huiz
		err = modelGame.GetGuildModel().Create(guildInfo)
		if err != nil {
			logger.Error("创建门派失败")
			continue
		}
		this.SetGuildInfo(guildInfo)
	}
	return
}

func (this *GuildManager) DelRobotGuild() {

	openDay := this.GetSystem().GetServerOpenDaysByServerId(base.Conf.ServerId)
	if openDay != 2 {
		return
	}

	datas, err := modelGame.GetGuildModel().GetAllGuildInfoBySystem()
	if err != nil || len(datas) <= 0 {
		return
	}

	for _, data := range datas {

		userLvInfo := make(model.GuildRobotUserInfoSlice, 0)
		delMark := true
		for i, j := 0, len(data.Positions); i < j; i += 2 {
			userId := data.Positions[i]
			if userId > 0 {
				userInfo := this.GetUserManager().GetUserBasicInfo(userId)
				if userInfo != nil {
					delMark = false
					userLvInfo = append(userLvInfo, &model.GuildRobotUserInfo{UserId: userId, Lv: userInfo.Level})
				}
				continue
			}
		}

		if delMark {
			modelGame.GetGuildModel().DeleteGuild(data)
			this.DelGuildInfo(data)
			continue
		}

		sort.Sort(userLvInfo)

		data.Positions = make(model.IntSlice, 0)
		positionMark := make(map[int]int) //userId:position
		allNum := 0
		huiZ := -1
		for i3 := pb.GUILDPOSITION_HUIZHANG; i3 <= pb.GUILDPOSITION_CHENGYUAN; i3++ {
			cfg := gamedb.GetGuildGuildCfg(i3)
			if cfg == nil {
				continue
			}
			addNum := 0
			for i2 := allNum; i2 < len(userLvInfo); i2++ {
				if addNum >= cfg.Count {
					continue
				}
				if i3 == pb.GUILDPOSITION_HUIZHANG {
					huiZ = userLvInfo[i2].UserId
				}
				positionMark[userLvInfo[i2].UserId] = i3
				data.Positions = append(data.Positions, userLvInfo[i2].UserId, i3)
				addNum++
				allNum++
			}
		}

		data.ChairmanId = huiZ
		this.GetGuild().SetGuildInfo(data)

		for _, info := range userLvInfo {
			this.changePositionInfo(info.UserId, positionMark)
			uInfo := this.GetUserManager().GetUser(info.UserId)
			if uInfo != nil {
				ack := &pb.GuildLoadInfoAck{}
				this.LoadGuild(uInfo, ack)
				this.GetUserManager().SendMessage(uInfo, ack, true)
			}
		}
	}
	ack := &pb.AllGuildInfosAck{}
	this.GetAllGuildInfo(nil, ack)
	this.BroadcastAll(ack)
}

func (this *GuildManager) changePositionInfo(userId int, positionMark map[int]int) {

	uInfo := this.GetUserManager().GetUser(userId)
	afterPos := positionMark[userId]
	if uInfo != nil {
		uInfo.GuildData.Position = afterPos
		if afterPos <= 0 {
			uInfo.GuildData.NowGuildId = 0
			uInfo.GuildData.ContributionValue = 0
		}
		uInfo.Dirty = true
		infos := make([]int32, 0)
		infos = append(infos, int32(userId), int32(afterPos))
		this.BroadcastGuildInfoNtf(uInfo.GuildData.NowGuildId, pb.GUILDBROCASTTYPE_APPLY_USER, infos, nil)
	} else {
		offUserInfo := this.GetUserManager().GetOfflineUserInfo(userId)
		if offUserInfo != nil {
			offUserInfo.GuildData.Position = afterPos
			if afterPos <= 0 {
				offUserInfo.GuildData.NowGuildId = 0
				offUserInfo.GuildData.ContributionValue = 0
			}
			err := this.GetUserManager().Save(offUserInfo, true)
			if err != nil {
				logger.Error("系统门派 离线变更职位信息err:%v", err)
			}
		}
	}
	return
}
