package fight

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

func (this *Fight) ApplyForHelp(user *objs.User, HelpUserId int, source int) error {

	if user.FightId <= 0 {
		return gamedb.ERRFIGHTID
	}

	return this.GetChat().ChatForFightHelp(user)
	//ntf := &pb.FightApplyForHelpNtf{
	//	ReqHelpUserId: int32(user.Id),
	//	ReqHelpName:   user.NickName,
	//	StageId:       int32(user.FightStageId),
	//	Source:        int32(source),
	//	ReqHelpUser:   this.GetUserManager().BuilderBrieUserInfo(user.Id),
	//}
	//
	//helpUser := this.GetUserManager().GetUser(HelpUserId)
	//if helpUser != nil {
	//	this.GetUserManager().SendMessage(helpUser, ntf, true)
	//} else {
	//	return gamedb.ERRUSEROFFLINE
	//}
	//return nil
}

func (this *Fight) AskForHelpResult(user *objs.User, isAgree bool, reqHelpUserId, helpStageId int) error {

	ntf := &pb.FightAskForHelpResultNtf{
		IsAgree:    isAgree,
		HelpUserId: int32(user.Id),
		Name:       user.NickName,
	}
	if !isAgree {
		this.GetUserManager().SendMessageByUserId(reqHelpUserId, ntf)
		return nil
	}

	//applyHelpUser := this.GetUserManager().GetUser(reqHelpUserId)
	//if applyHelpUser == nil {
	//	return gamedb.ERRPLAYERLEAVE
	//}
	//if applyHelpUser.FightStageId != helpStageId {
	//	return gamedb.ERRFIGHTEND
	//}

	if user.FightStageId != helpStageId {
		stageConf := gamedb.GetStageStageCfg(helpStageId)
		if stageConf.Type == constFight.FIGHT_TYPE_DARKPALACE_BOSS {
			//协助者进入战斗
			err := this.GetDarkPalace().EnterDarkPalaceFight(user, helpStageId, reqHelpUserId)
			return err
		} else if stageConf.Type == constFight.FIGHT_TYPE_HELL_BOSS {
			//协助者进入战斗
			err := this.GetHellBoss().EnterHellBossFight(user, helpStageId, reqHelpUserId)
			return err
		}
	} else {
		msg := &pbserver.GsToFsChangeToHelperReq{
			UserId:       int32(user.Id),
			ToHelpUserId: int32(reqHelpUserId),
		}
		err := this.FSSendMessage(user.FightId, user.FightStageId, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *Fight) UpdateUserfightNum(user *objs.User) {

	if user.FightId <= 0 {
		return
	}
	stageConf := gamedb.GetStageStageCfg(user.FightStageId)
	if stageConf == nil {
		return
	}
	if stageConf.Type != constFight.FIGHT_TYPE_DARKPALACE_BOSS && stageConf.Type != constFight.FIGHT_TYPE_HELL_BOSS {
		return
	}

	msg := &pbserver.GsToFsFightNumChangeReq{
		UserId:         int32(user.Id),
		FightNumChange: int32(this.GetDarkPalace().GetSurplusNum(user)),
	}
	if stageConf.Type == constFight.FIGHT_TYPE_HELL_BOSS {
		msg.FightNumChange = int32(this.GetHellBoss().GetSurplusNum(user))
	}
	this.FSSendMessage(user.FightId, user.FightStageId, msg)

}
