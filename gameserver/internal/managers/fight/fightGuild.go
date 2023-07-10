package fight

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constGuild"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

const SHABAKENEW = 1

func (this *Fight) EnterGuildBonfire(user *objs.User, guildId int) error {

	fightId, err := this.getFightIdByStageId(constFight.FIGHT_TYPE_GUILD_BONFIRE_STAGE, guildId)
	if err != nil {
		return err
	}
	if fightId <= 0 {
		//创建战斗
		fightId, err = this.GetFight().CreateFight(constFight.FIGHT_TYPE_GUILD_BONFIRE_STAGE, common.IntToBytes(guildId))
		if err != nil {
			return err
		}
	}
	//玩家进入战斗
	err = this.GetFight().EnterFightByFightId(user, constFight.FIGHT_TYPE_GUILD_BONFIRE_STAGE, fightId)
	if err != nil {
		return err
	}
	return nil
}

/**
 *  @Description: 公会篝火增加玩家经验
 *  @param endMsg
 */
func (this *Fight) GuildBonfireUserAddExp(endMsg *pbserver.GuildbonfireExpAddNtf) {

	users := common.ConvertInt32SliceToIntSlice(endMsg.UserIds)
	this.GetGuildBonfire().AddUserExp(users)
}

/**
 *  @Description: 公会篝火结束
 *  @param endMsg
 */
func (this *Fight) guildBonfireEnd(endMsg *pbserver.FSFightEndNtf) {
	var userIds []int
	if endMsg.Winners != nil {
		userIds = common.ConvertInt32SliceToIntSlice(endMsg.Winners)
	}
	this.GetGuildBonfire().GuildBonfireFightResult(userIds)
}

/**
 *  @Description: 进入沙巴克战斗
 *  @param user
 *  @return error
 */
func (this *Fight) EnterShabakeFight(user *objs.User) error {

	if SHABAKENEW == 1 {
		return this.EnterShabakeFightNew(user)
	}

	//this.fightMu.Lock()
	//isExist := false
	//var err error
	//if this.shabakeFightId > 0 {
	//	//检查战斗是否存在
	//	isExist = this.checkFightIsExistByFightId(this.shabakeFightId, constFight.FIGHT_TYPE_SHABAKE_STAGE)
	//}
	//if !isExist {
	//	//创建战斗
	//	this.shabakeFightId, err = this.GetFight().CreateFight(constFight.FIGHT_TYPE_SHABAKE_STAGE, nil)
	//	if err != nil {
	//		return err
	//	}
	//}
	//this.fightMu.Unlock()
	////玩家进入战斗
	//err = this.GetFight().EnterFightByFightId(user, constFight.FIGHT_TYPE_SHABAKE_STAGE, this.shabakeFightId)
	//if err != nil {
	//	return err
	//}
	return nil
}

/**
 *  @Description: 进入沙巴克战斗
 *  @param user
 *  @return error
 */
func (this *Fight) EnterShabakeFightNew(user *objs.User) error {

	request := &pbserver.GSTOFSGetFightIdReq{
		StageId: int32(constFight.FIGHT_TYPE_SHABAKE_NEW_STAGE),
	}
	replay := &pbserver.FSTOGSGetFightIdAck{}
	err := this.FSRpcCall(0, constFight.FIGHT_TYPE_SHABAKE_NEW_STAGE, request, replay)
	if err != nil {
		logger.Error("沙巴克战斗获取Id异常,stageId：%v,err:%v", constFight.FIGHT_TYPE_SHABAKE_NEW_STAGE, err)
		return gamedb.ERRFIGHTID
	}
	//玩家进入战斗
	err = this.GetFight().EnterFightByFightId(user, constFight.FIGHT_TYPE_SHABAKE_NEW_STAGE, int(replay.FightId))
	if err != nil {
		return err
	}
	return nil
}

func (this *Fight) guildShabakeEnd(endMsg *pbserver.FSFightEndNtf) {

	this.shabakeFightId = 0
	result := &pbserver.ShabakeFightEndNtf{}
	err := result.Unmarshal(endMsg.CpData)
	if err != nil {
		logger.Error("解析沙巴克结果异常：%v", err)
		return
	}
	userRank := make([]int, 0)
	otherUser := make([]int, 0)
	guidRank := make([]int, 0)
	otherGuild := make([]int, 0)
	for _, v := range result.UserRank {
		if v.Score > 0 {
			userRank = append(userRank, int(v.Id))
		} else {
			otherUser = append(otherUser, int(v.Id))
		}

	}
	for _, v := range result.GuildRank {
		if v.Score > 0 {
			guidRank = append(guidRank, int(v.Id))
		} else {
			otherGuild = append(otherGuild, int(v.Id))
		}
	}
	logger.Info("沙巴克结果,玩家排行：%v，门派排行：%v,无积分玩家：%v,无积分门派：%v", userRank, guidRank, otherUser, otherGuild)
	this.GetShabake().ShabakeFightEndAck(userRank, guidRank, otherUser)
}

/**
 *  @Description: 进入龙柱守护战斗
 *  @param user
 *  @return error
 */
func (this *Fight) EnterGuardPillarFight(user *objs.User) error {

	guildId := user.GuildData.NowGuildId
	fightId, err := this.getFightIdByStageId(constFight.FIGHT_TYPE_GUARDPILLAR_STAGE, guildId)
	if err != nil {
		return nil
	}
	if fightId <= 0 {
		//创建战斗
		fightId, err = this.GetFight().CreateFight(constFight.FIGHT_TYPE_GUARDPILLAR_STAGE, common.IntToBytes(guildId))
		if err != nil {
			return err
		}
		stageConf := gamedb.GetStageStageCfg(constFight.FIGHT_TYPE_GUARDPILLAR_STAGE)
		this.fightMu.Lock()
		this.guiardPillarFightEndTime[guildId] = int(time.Now().Unix()) + stageConf.LifeTime
		this.fightMu.Unlock()
		this.GetGuild().SendMsgToAllUser(guildId, &pb.GuildActivityOpenNtf{
			GuildActivityId: constGuild.GUILD_ACTIVITY_GUARD_PILLAR,
			EndTime:         int64(this.guiardPillarFightEndTime[guildId]),
		}, []int{user.Id})
	}
	//玩家进入战斗
	err = this.GetFight().EnterFightByFightId(user, constFight.FIGHT_TYPE_GUARDPILLAR_STAGE, fightId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Fight) GetGuardPillarFightEndTime(guildId int) int {

	//fightId := this.guiardPillarFight[guildId]
	//if fightId <= 0 {
	//	return 0
	//} else {
	//	isExist := this.checkFightIsExistByFightId(fightId, constFight.FIGHT_TYPE_GUARDPILLAR_STAGE)
	//	if !isExist {
	//		return 0
	//	}
	//}
	return this.guiardPillarFightEndTime[guildId]
}

/**
 *  @Description: 龙柱守护战斗结束
 *  @param endMsg
 */
func (this *Fight) guardPillarFightEnd(endMsg *pbserver.FSFightEndNtf) {

	result := &pbserver.GuardPillarFightEnd{}
	err := result.Unmarshal(endMsg.CpData)
	if err != nil {
		logger.Error("解析龙柱守卫结果异常：%v", err)
		return
	}
	rounds := int(result.Wave)
	for rank, v := range result.Users {
		userId := int(v)
		this.DispatchEvent(userId, rank, func(userId int, user *objs.User, data interface{}) {
			//if user == nil {
			//	logger.Warn("接收到战斗服发来的战斗结束，未找到相应玩家:%v", userId)
			//	return
			//}
			rankIndex := data.(int)
			this.GetGuardPillar().GuardPillarResult(userId, constFight.FIGHT_TYPE_GUARDPILLAR_STAGE, rounds, rankIndex+1)
			if user != nil {
				kyEvent.StageEnd(user, int(endMsg.StageId), pb.RESULTFLAG_SUCCESS, user.FightStartTime, nil)
			}
		})
	}
	logger.Info("接收到龙柱守卫的战斗结果,波数：%v，排名：%v", result.Wave, result.Users)
}
