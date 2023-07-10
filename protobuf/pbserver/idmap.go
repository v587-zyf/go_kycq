// This is generated by git@gitlab.hd.com:yulong/genproto.git
// Do not modify here.

package pbserver
import (
    "reflect"
)

var msgPrototypes = make(map[uint16]reflect.Type)
var msgNames = make(map[uint16]string)
var msgLogLv = make(map[int]int)

func init() {
    msgPrototypes[CmdErrorAckId] = reflect.TypeOf((*ErrorAck)(nil)).Elem()
    msgPrototypes[CmdHandShakeReqId] = reflect.TypeOf((*HandShakeReq)(nil)).Elem()
    msgPrototypes[CmdHandShakeAckId] = reflect.TypeOf((*HandShakeAck)(nil)).Elem()
    msgPrototypes[CmdHandCloseNtfId] = reflect.TypeOf((*HandCloseNtf)(nil)).Elem()
    msgPrototypes[CmdLoginKeyVerifyReqId] = reflect.TypeOf((*LoginKeyVerifyReq)(nil)).Elem()
    msgPrototypes[CmdLoginKeyVerifyAckId] = reflect.TypeOf((*LoginKeyVerifyAck)(nil)).Elem()
    msgPrototypes[CmdLoginKeyVerifyUpdateReqId] = reflect.TypeOf((*LoginKeyVerifyUpdateReq)(nil)).Elem()
    msgPrototypes[CmdGSMessageToFSId] = reflect.TypeOf((*GSMessageToFS)(nil)).Elem()
    msgPrototypes[CmdFSMessageToGSId] = reflect.TypeOf((*FSMessageToGS)(nil)).Elem()
    msgPrototypes[CmdFSCallMessageToGSId] = reflect.TypeOf((*FSCallMessageToGS)(nil)).Elem()
    msgPrototypes[CmdGSCallMessageToFSId] = reflect.TypeOf((*GSCallMessageToFS)(nil)).Elem()
    msgPrototypes[CmdGsRouteMessageToFightId] = reflect.TypeOf((*GsRouteMessageToFight)(nil)).Elem()
    msgPrototypes[CmdFSCreateFightReqId] = reflect.TypeOf((*FSCreateFightReq)(nil)).Elem()
    msgPrototypes[CmdFSCreateFightAckId] = reflect.TypeOf((*FSCreateFightAck)(nil)).Elem()
    msgPrototypes[CmdFSEnterFightReqId] = reflect.TypeOf((*FSEnterFightReq)(nil)).Elem()
    msgPrototypes[CmdFSEnterFightAckId] = reflect.TypeOf((*FSEnterFightAck)(nil)).Elem()
    msgPrototypes[CmdFSLeaveFightReqId] = reflect.TypeOf((*FSLeaveFightReq)(nil)).Elem()
    msgPrototypes[CmdFSLeaveFightAckId] = reflect.TypeOf((*FSLeaveFightAck)(nil)).Elem()
    msgPrototypes[CmdFSUpdateUserInfoNtfId] = reflect.TypeOf((*FSUpdateUserInfoNtf)(nil)).Elem()
    msgPrototypes[CmdGSToFsUpdateUserFightModelId] = reflect.TypeOf((*GSToFsUpdateUserFightModel)(nil)).Elem()
    msgPrototypes[CmdGsToFSUserReliveReqId] = reflect.TypeOf((*GsToFSUserReliveReq)(nil)).Elem()
    msgPrototypes[CmdFSToGsUserReliveAckId] = reflect.TypeOf((*FSToGsUserReliveAck)(nil)).Elem()
    msgPrototypes[CmdGsToFSCheckUserReliveReqId] = reflect.TypeOf((*GsToFSCheckUserReliveReq)(nil)).Elem()
    msgPrototypes[CmdFsToGSCheckUserReliveAckId] = reflect.TypeOf((*FsToGSCheckUserReliveAck)(nil)).Elem()
    msgPrototypes[CmdGSTOFSCheckFightReqId] = reflect.TypeOf((*GSTOFSCheckFightReq)(nil)).Elem()
    msgPrototypes[CmdGSTOFSGetFightIdReqId] = reflect.TypeOf((*GSTOFSGetFightIdReq)(nil)).Elem()
    msgPrototypes[CmdGsToFsUpdateUserElfReqId] = reflect.TypeOf((*GsToFsUpdateUserElfReq)(nil)).Elem()
    msgPrototypes[CmdGsToFsFightNpcEventReqId] = reflect.TypeOf((*GsToFsFightNpcEventReq)(nil)).Elem()
    msgPrototypes[CmdFSFightEndNtfId] = reflect.TypeOf((*FSFightEndNtf)(nil)).Elem()
    msgPrototypes[CmdFsResidentFightNtfId] = reflect.TypeOf((*FsResidentFightNtf)(nil)).Elem()
    msgPrototypes[CmdGsToFsResidentFightReqId] = reflect.TypeOf((*GsToFsResidentFightReq)(nil)).Elem()
    msgPrototypes[CmdGsTOFsPickUpReqId] = reflect.TypeOf((*GsTOFsPickUpReq)(nil)).Elem()
    msgPrototypes[CmdFsTOGsPickUpAckId] = reflect.TypeOf((*FsTOGsPickUpAck)(nil)).Elem()
    msgPrototypes[CmdFSAddItemReqId] = reflect.TypeOf((*FSAddItemReq)(nil)).Elem()
    msgPrototypes[CmdFSAddItemAckId] = reflect.TypeOf((*FSAddItemAck)(nil)).Elem()
    msgPrototypes[CmdGsToFsPickRedPacketInfoId] = reflect.TypeOf((*GsToFsPickRedPacketInfo)(nil)).Elem()
    msgPrototypes[CmdUserDropReqId] = reflect.TypeOf((*UserDropReq)(nil)).Elem()
    msgPrototypes[CmdUserDropAckId] = reflect.TypeOf((*UserDropAck)(nil)).Elem()
    msgPrototypes[CmdFsRandomDeliveryNtfId] = reflect.TypeOf((*FsRandomDeliveryNtf)(nil)).Elem()
    msgPrototypes[CmdGsToFsUseItemNtfId] = reflect.TypeOf((*GsToFsUseItemNtf)(nil)).Elem()
    msgPrototypes[CmdGsToFsGmReqId] = reflect.TypeOf((*GsToFsGmReq)(nil)).Elem()
    msgPrototypes[CmdFsSkillUseNtfId] = reflect.TypeOf((*FsSkillUseNtf)(nil)).Elem()
    msgPrototypes[CmdFsTOGsClearSkillCdNtfId] = reflect.TypeOf((*FsTOGsClearSkillCdNtf)(nil)).Elem()
    msgPrototypes[CmdGsToFsGetCheerNumReqId] = reflect.TypeOf((*GsToFsGetCheerNumReq)(nil)).Elem()
    msgPrototypes[CmdGsToFsCheerReqId] = reflect.TypeOf((*GsToFsCheerReq)(nil)).Elem()
    msgPrototypes[CmdGsToFsGetPotionCdReqId] = reflect.TypeOf((*GsToFsGetPotionCdReq)(nil)).Elem()
    msgPrototypes[CmdGsToFsUsePotionReqId] = reflect.TypeOf((*GsToFsUsePotionReq)(nil)).Elem()
    msgPrototypes[CmdGsToFsCollectionReqId] = reflect.TypeOf((*GsToFsCollectionReq)(nil)).Elem()
    msgPrototypes[CmdFsToGsCollectionNtfId] = reflect.TypeOf((*FsToGsCollectionNtf)(nil)).Elem()
    msgPrototypes[CmdGsToFsCollectionCancelReqId] = reflect.TypeOf((*GsToFsCollectionCancelReq)(nil)).Elem()
    msgPrototypes[CmdGsToFsUseFitReqId] = reflect.TypeOf((*GsToFsUseFitReq)(nil)).Elem()
    msgPrototypes[CmdGsToFsFitCacelReqId] = reflect.TypeOf((*GsToFsFitCacelReq)(nil)).Elem()
    msgPrototypes[CmdGsToFsUpdatePetReqId] = reflect.TypeOf((*GsToFsUpdatePetReq)(nil)).Elem()
    msgPrototypes[CmdGsToFsUseCutTreasureReqId] = reflect.TypeOf((*GsToFsUseCutTreasureReq)(nil)).Elem()
    msgPrototypes[CmdGsToFsCheckForHelpReqId] = reflect.TypeOf((*GsToFsCheckForHelpReq)(nil)).Elem()
    msgPrototypes[CmdGsToFsChangeToHelperReqId] = reflect.TypeOf((*GsToFsChangeToHelperReq)(nil)).Elem()
    msgPrototypes[CmdGsToFsFightNumChangeReqId] = reflect.TypeOf((*GsToFsFightNumChangeReq)(nil)).Elem()
    msgPrototypes[CmdGsToFsFightScoreLessReqId] = reflect.TypeOf((*GsToFsFightScoreLessReq)(nil)).Elem()
    msgPrototypes[CmdGsToFsGamedbReloadReqId] = reflect.TypeOf((*GsToFsGamedbReloadReq)(nil)).Elem()
    msgPrototypes[CmdFsToGsShabakeKillBossNtfId] = reflect.TypeOf((*FsToGsShabakeKillBossNtf)(nil)).Elem()
    msgPrototypes[CmdFsToGsActorKillNtfId] = reflect.TypeOf((*FsToGsActorKillNtf)(nil)).Elem()
    msgPrototypes[CmdWorldBossStatusNtfId] = reflect.TypeOf((*WorldBossStatusNtf)(nil)).Elem()
    msgPrototypes[CmdFSContinueFightReqId] = reflect.TypeOf((*FSContinueFightReq)(nil)).Elem()
    msgPrototypes[CmdFsFieldBossInfoNtfId] = reflect.TypeOf((*FsFieldBossInfoNtf)(nil)).Elem()
    msgPrototypes[CmdExpStageKillMonsterNtfId] = reflect.TypeOf((*ExpStageKillMonsterNtf)(nil)).Elem()
    msgPrototypes[CmdMiningNewFightInfoReqId] = reflect.TypeOf((*MiningNewFightInfoReq)(nil)).Elem()
    msgPrototypes[CmdHangUpKillWaveNtfId] = reflect.TypeOf((*HangUpKillWaveNtf)(nil)).Elem()
    msgPrototypes[CmdGuildbonfireExpAddNtfId] = reflect.TypeOf((*GuildbonfireExpAddNtf)(nil)).Elem()
    msgPrototypes[CmdPaodianGoodsAddNtfId] = reflect.TypeOf((*PaodianGoodsAddNtf)(nil)).Elem()
    msgPrototypes[CmdWorldLeaderFightEndNtfId] = reflect.TypeOf((*WorldLeaderFightEndNtf)(nil)).Elem()
    msgPrototypes[CmdWorldLeaderFightRankNtfId] = reflect.TypeOf((*WorldLeaderFightRankNtf)(nil)).Elem()
    msgPrototypes[CmdFsFieldBossDieUserInfoNtfId] = reflect.TypeOf((*FsFieldBossDieUserInfoNtf)(nil)).Elem()
    msgPrototypes[CmdMagicTowerGetUserInfoReqId] = reflect.TypeOf((*MagicTowerGetUserInfoReq)(nil)).Elem()
    msgPrototypes[CmdDaBaoKillMonsterNtfId] = reflect.TypeOf((*DaBaoKillMonsterNtf)(nil)).Elem()
    msgPrototypes[CmdDaBaoResumeEnergyReqId] = reflect.TypeOf((*DaBaoResumeEnergyReq)(nil)).Elem()
    msgPrototypes[CmdBossFamilyBossInfoReqId] = reflect.TypeOf((*BossFamilyBossInfoReq)(nil)).Elem()
    msgPrototypes[CmdSyncUserInfoNtfId] = reflect.TypeOf((*SyncUserInfoNtf)(nil)).Elem()
    msgPrototypes[CmdCCSToGsCrossFsIdChangeNtfId] = reflect.TypeOf((*CCSToGsCrossFsIdChangeNtf)(nil)).Elem()
    msgPrototypes[CmdChallengeSendLoseRewardNtfId] = reflect.TypeOf((*ChallengeSendLoseRewardNtf)(nil)).Elem()
    msgPrototypes[CmdChallengeAppuserUpNtfId] = reflect.TypeOf((*ChallengeAppuserUpNtf)(nil)).Elem()
    msgPrototypes[CmdGsToCcsBackGuildInfoNtfId] = reflect.TypeOf((*GsToCcsBackGuildInfoNtf)(nil)).Elem()
    msgPrototypes[CmdCcsToGsBroadShaBakeFirstGuildInfoId] = reflect.TypeOf((*CcsToGsBroadShaBakeFirstGuildInfo)(nil)).Elem()
    msgPrototypes[CmdChallengeAppuserUpToGsNtfId] = reflect.TypeOf((*ChallengeAppuserUpToGsNtf)(nil)).Elem()
    msgPrototypes[CmdRechageCcsToGsReqId] = reflect.TypeOf((*RechageCcsToGsReq)(nil)).Elem()
    msgPrototypes[CmdRechageGsToCcsAckId] = reflect.TypeOf((*RechageGsToCcsAck)(nil)).Elem()
    msgPrototypes[CmdRechargeApplyReqId] = reflect.TypeOf((*RechargeApplyReq)(nil)).Elem()
    msgPrototypes[CmdBanInfoCcsToGsReqId] = reflect.TypeOf((*BanInfoCcsToGsReq)(nil)).Elem()
    msgPrototypes[CmdMailSendCCsToGsReqId] = reflect.TypeOf((*MailSendCCsToGsReq)(nil)).Elem()
    msgPrototypes[CmdFuncStateUpdateReqId] = reflect.TypeOf((*FuncStateUpdateReq)(nil)).Elem()
    msgPrototypes[CmdUpAnnouncementNowReqId] = reflect.TypeOf((*UpAnnouncementNowReq)(nil)).Elem()
    msgPrototypes[CmdUpPaoMaDengNowReqId] = reflect.TypeOf((*UpPaoMaDengNowReq)(nil)).Elem()
    msgPrototypes[CmdSetDayRechargeNumNtfId] = reflect.TypeOf((*SetDayRechargeNumNtf)(nil)).Elem()
    
    msgNames[CmdErrorAckId] = "ErrorAck"
    msgNames[CmdHandShakeReqId] = "HandShakeReq"
    msgNames[CmdHandShakeAckId] = "HandShakeAck"
    msgNames[CmdHandCloseNtfId] = "HandCloseNtf"
    msgNames[CmdLoginKeyVerifyReqId] = "LoginKeyVerifyReq"
    msgNames[CmdLoginKeyVerifyAckId] = "LoginKeyVerifyAck"
    msgNames[CmdLoginKeyVerifyUpdateReqId] = "LoginKeyVerifyUpdateReq"
    msgNames[CmdGSMessageToFSId] = "GSMessageToFS"
    msgNames[CmdFSMessageToGSId] = "FSMessageToGS"
    msgNames[CmdFSCallMessageToGSId] = "FSCallMessageToGS"
    msgNames[CmdGSCallMessageToFSId] = "GSCallMessageToFS"
    msgNames[CmdGsRouteMessageToFightId] = "GsRouteMessageToFight"
    msgNames[CmdFSCreateFightReqId] = "FSCreateFightReq"
    msgNames[CmdFSCreateFightAckId] = "FSCreateFightAck"
    msgNames[CmdFSEnterFightReqId] = "FSEnterFightReq"
    msgNames[CmdFSEnterFightAckId] = "FSEnterFightAck"
    msgNames[CmdFSLeaveFightReqId] = "FSLeaveFightReq"
    msgNames[CmdFSLeaveFightAckId] = "FSLeaveFightAck"
    msgNames[CmdFSUpdateUserInfoNtfId] = "FSUpdateUserInfoNtf"
    msgNames[CmdGSToFsUpdateUserFightModelId] = "GSToFsUpdateUserFightModel"
    msgNames[CmdGsToFSUserReliveReqId] = "GsToFSUserReliveReq"
    msgNames[CmdFSToGsUserReliveAckId] = "FSToGsUserReliveAck"
    msgNames[CmdGsToFSCheckUserReliveReqId] = "GsToFSCheckUserReliveReq"
    msgNames[CmdFsToGSCheckUserReliveAckId] = "FsToGSCheckUserReliveAck"
    msgNames[CmdGSTOFSCheckFightReqId] = "GSTOFSCheckFightReq"
    msgNames[CmdGSTOFSGetFightIdReqId] = "GSTOFSGetFightIdReq"
    msgNames[CmdGsToFsUpdateUserElfReqId] = "GsToFsUpdateUserElfReq"
    msgNames[CmdGsToFsFightNpcEventReqId] = "GsToFsFightNpcEventReq"
    msgNames[CmdFSFightEndNtfId] = "FSFightEndNtf"
    msgNames[CmdFsResidentFightNtfId] = "FsResidentFightNtf"
    msgNames[CmdGsToFsResidentFightReqId] = "GsToFsResidentFightReq"
    msgNames[CmdGsTOFsPickUpReqId] = "GsTOFsPickUpReq"
    msgNames[CmdFsTOGsPickUpAckId] = "FsTOGsPickUpAck"
    msgNames[CmdFSAddItemReqId] = "FSAddItemReq"
    msgNames[CmdFSAddItemAckId] = "FSAddItemAck"
    msgNames[CmdGsToFsPickRedPacketInfoId] = "GsToFsPickRedPacketInfo"
    msgNames[CmdUserDropReqId] = "UserDropReq"
    msgNames[CmdUserDropAckId] = "UserDropAck"
    msgNames[CmdFsRandomDeliveryNtfId] = "FsRandomDeliveryNtf"
    msgNames[CmdGsToFsUseItemNtfId] = "GsToFsUseItemNtf"
    msgNames[CmdGsToFsGmReqId] = "GsToFsGmReq"
    msgNames[CmdFsSkillUseNtfId] = "FsSkillUseNtf"
    msgNames[CmdFsTOGsClearSkillCdNtfId] = "FsTOGsClearSkillCdNtf"
    msgNames[CmdGsToFsGetCheerNumReqId] = "GsToFsGetCheerNumReq"
    msgNames[CmdGsToFsCheerReqId] = "GsToFsCheerReq"
    msgNames[CmdGsToFsGetPotionCdReqId] = "GsToFsGetPotionCdReq"
    msgNames[CmdGsToFsUsePotionReqId] = "GsToFsUsePotionReq"
    msgNames[CmdGsToFsCollectionReqId] = "GsToFsCollectionReq"
    msgNames[CmdFsToGsCollectionNtfId] = "FsToGsCollectionNtf"
    msgNames[CmdGsToFsCollectionCancelReqId] = "GsToFsCollectionCancelReq"
    msgNames[CmdGsToFsUseFitReqId] = "GsToFsUseFitReq"
    msgNames[CmdGsToFsFitCacelReqId] = "GsToFsFitCacelReq"
    msgNames[CmdGsToFsUpdatePetReqId] = "GsToFsUpdatePetReq"
    msgNames[CmdGsToFsUseCutTreasureReqId] = "GsToFsUseCutTreasureReq"
    msgNames[CmdGsToFsCheckForHelpReqId] = "GsToFsCheckForHelpReq"
    msgNames[CmdGsToFsChangeToHelperReqId] = "GsToFsChangeToHelperReq"
    msgNames[CmdGsToFsFightNumChangeReqId] = "GsToFsFightNumChangeReq"
    msgNames[CmdGsToFsFightScoreLessReqId] = "GsToFsFightScoreLessReq"
    msgNames[CmdGsToFsGamedbReloadReqId] = "GsToFsGamedbReloadReq"
    msgNames[CmdFsToGsShabakeKillBossNtfId] = "FsToGsShabakeKillBossNtf"
    msgNames[CmdFsToGsActorKillNtfId] = "FsToGsActorKillNtf"
    msgNames[CmdWorldBossStatusNtfId] = "WorldBossStatusNtf"
    msgNames[CmdFSContinueFightReqId] = "FSContinueFightReq"
    msgNames[CmdFsFieldBossInfoNtfId] = "FsFieldBossInfoNtf"
    msgNames[CmdExpStageKillMonsterNtfId] = "ExpStageKillMonsterNtf"
    msgNames[CmdMiningNewFightInfoReqId] = "MiningNewFightInfoReq"
    msgNames[CmdHangUpKillWaveNtfId] = "HangUpKillWaveNtf"
    msgNames[CmdGuildbonfireExpAddNtfId] = "GuildbonfireExpAddNtf"
    msgNames[CmdPaodianGoodsAddNtfId] = "PaodianGoodsAddNtf"
    msgNames[CmdWorldLeaderFightEndNtfId] = "WorldLeaderFightEndNtf"
    msgNames[CmdWorldLeaderFightRankNtfId] = "WorldLeaderFightRankNtf"
    msgNames[CmdFsFieldBossDieUserInfoNtfId] = "FsFieldBossDieUserInfoNtf"
    msgNames[CmdMagicTowerGetUserInfoReqId] = "MagicTowerGetUserInfoReq"
    msgNames[CmdDaBaoKillMonsterNtfId] = "DaBaoKillMonsterNtf"
    msgNames[CmdDaBaoResumeEnergyReqId] = "DaBaoResumeEnergyReq"
    msgNames[CmdBossFamilyBossInfoReqId] = "BossFamilyBossInfoReq"
    msgNames[CmdSyncUserInfoNtfId] = "SyncUserInfoNtf"
    msgNames[CmdCCSToGsCrossFsIdChangeNtfId] = "CCSToGsCrossFsIdChangeNtf"
    msgNames[CmdChallengeSendLoseRewardNtfId] = "ChallengeSendLoseRewardNtf"
    msgNames[CmdChallengeAppuserUpNtfId] = "ChallengeAppuserUpNtf"
    msgNames[CmdGsToCcsBackGuildInfoNtfId] = "GsToCcsBackGuildInfoNtf"
    msgNames[CmdCcsToGsBroadShaBakeFirstGuildInfoId] = "CcsToGsBroadShaBakeFirstGuildInfo"
    msgNames[CmdChallengeAppuserUpToGsNtfId] = "ChallengeAppuserUpToGsNtf"
    msgNames[CmdRechageCcsToGsReqId] = "RechageCcsToGsReq"
    msgNames[CmdRechageGsToCcsAckId] = "RechageGsToCcsAck"
    msgNames[CmdRechargeApplyReqId] = "RechargeApplyReq"
    msgNames[CmdBanInfoCcsToGsReqId] = "BanInfoCcsToGsReq"
    msgNames[CmdMailSendCCsToGsReqId] = "MailSendCCsToGsReq"
    msgNames[CmdFuncStateUpdateReqId] = "FuncStateUpdateReq"
    msgNames[CmdUpAnnouncementNowReqId] = "UpAnnouncementNowReq"
    msgNames[CmdUpPaoMaDengNowReqId] = "UpPaoMaDengNowReq"
    msgNames[CmdSetDayRechargeNumNtfId] = "SetDayRechargeNumNtf"
	
    msgLogLv[CmdErrorAckId] = 1
    msgLogLv[CmdHandShakeReqId] = 1
    msgLogLv[CmdHandShakeAckId] = 0
    msgLogLv[CmdHandCloseNtfId] = 0
    msgLogLv[CmdLoginKeyVerifyReqId] = 1
    msgLogLv[CmdLoginKeyVerifyAckId] = 1
    msgLogLv[CmdLoginKeyVerifyUpdateReqId] = 1
    msgLogLv[CmdGSMessageToFSId] = 1
    msgLogLv[CmdFSMessageToGSId] = 1
    msgLogLv[CmdFSCallMessageToGSId] = 1
    msgLogLv[CmdGSCallMessageToFSId] = 1
    msgLogLv[CmdGsRouteMessageToFightId] = 1
    msgLogLv[CmdFSCreateFightReqId] = 1
    msgLogLv[CmdFSCreateFightAckId] = 0
    msgLogLv[CmdFSEnterFightReqId] = 1
    msgLogLv[CmdFSEnterFightAckId] = 0
    msgLogLv[CmdFSLeaveFightReqId] = 1
    msgLogLv[CmdFSLeaveFightAckId] = 0
    msgLogLv[CmdFSUpdateUserInfoNtfId] = 1
    msgLogLv[CmdGSToFsUpdateUserFightModelId] = 1
    msgLogLv[CmdGsToFSUserReliveReqId] = 1
    msgLogLv[CmdFSToGsUserReliveAckId] = 0
    msgLogLv[CmdGsToFSCheckUserReliveReqId] = 1
    msgLogLv[CmdFsToGSCheckUserReliveAckId] = 0
    msgLogLv[CmdGSTOFSCheckFightReqId] = 1
    msgLogLv[CmdGSTOFSGetFightIdReqId] = 1
    msgLogLv[CmdGsToFsUpdateUserElfReqId] = 1
    msgLogLv[CmdGsToFsFightNpcEventReqId] = 1
    msgLogLv[CmdGsTOFsPickUpReqId] = 1
    msgLogLv[CmdFsTOGsPickUpAckId] = 0
    msgLogLv[CmdFSAddItemReqId] = 1
    msgLogLv[CmdFSAddItemAckId] = 0
    msgLogLv[CmdGsToFsPickRedPacketInfoId] = 1
    msgLogLv[CmdFSFightEndNtfId] = 1
    msgLogLv[CmdUserDropReqId] = 1
    msgLogLv[CmdUserDropAckId] = 0
    msgLogLv[CmdFsRandomDeliveryNtfId] = 1
    msgLogLv[CmdGsToFsUseItemNtfId] = 1
    msgLogLv[CmdGsToFsGmReqId] = 1
    msgLogLv[CmdFsSkillUseNtfId] = 0
    msgLogLv[CmdFsTOGsClearSkillCdNtfId] = 1
    msgLogLv[CmdFsResidentFightNtfId] = 1
    msgLogLv[CmdGsToFsResidentFightReqId] = 0
    msgLogLv[CmdGsToFsGetCheerNumReqId] = 1
    msgLogLv[CmdGsToFsCheerReqId] = 1
    msgLogLv[CmdGsToFsGetPotionCdReqId] = 1
    msgLogLv[CmdGsToFsUsePotionReqId] = 1
    msgLogLv[CmdGsToFsCollectionReqId] = 1
    msgLogLv[CmdFsToGsCollectionNtfId] = 0
    msgLogLv[CmdGsToFsCollectionCancelReqId] = 1
    msgLogLv[CmdGsToFsUseFitReqId] = 1
    msgLogLv[CmdGsToFsFitCacelReqId] = 1
    msgLogLv[CmdGsToFsUpdatePetReqId] = 1
    msgLogLv[CmdGsToFsUseCutTreasureReqId] = 1
    msgLogLv[CmdGsToFsCheckForHelpReqId] = 1
    msgLogLv[CmdGsToFsChangeToHelperReqId] = 1
    msgLogLv[CmdGsToFsFightNumChangeReqId] = 1
    msgLogLv[CmdGsToFsFightScoreLessReqId] = 1
    msgLogLv[CmdGsToFsGamedbReloadReqId] = 1
    msgLogLv[CmdFsToGsShabakeKillBossNtfId] = 1
    msgLogLv[CmdFsToGsActorKillNtfId] = 1
    msgLogLv[CmdWorldBossStatusNtfId] = 1
    msgLogLv[CmdFSContinueFightReqId] = 1
    msgLogLv[CmdFsFieldBossInfoNtfId] = 1
    msgLogLv[CmdExpStageKillMonsterNtfId] = 1
    msgLogLv[CmdMiningNewFightInfoReqId] = 1
    msgLogLv[CmdHangUpKillWaveNtfId] = 1
    msgLogLv[CmdGuildbonfireExpAddNtfId] = 1
    msgLogLv[CmdPaodianGoodsAddNtfId] = 1
    msgLogLv[CmdWorldLeaderFightEndNtfId] = 1
    msgLogLv[CmdWorldLeaderFightRankNtfId] = 1
    msgLogLv[CmdFsFieldBossDieUserInfoNtfId] = 1
    msgLogLv[CmdMagicTowerGetUserInfoReqId] = 1
    msgLogLv[CmdDaBaoKillMonsterNtfId] = 1
    msgLogLv[CmdDaBaoResumeEnergyReqId] = 1
    msgLogLv[CmdBossFamilyBossInfoReqId] = 1
    msgLogLv[CmdSyncUserInfoNtfId] = 1
    msgLogLv[CmdCCSToGsCrossFsIdChangeNtfId] = 0
    msgLogLv[CmdChallengeSendLoseRewardNtfId] = 0
    msgLogLv[CmdChallengeAppuserUpNtfId] = 0
    msgLogLv[CmdGsToCcsBackGuildInfoNtfId] = 0
    msgLogLv[CmdCcsToGsBroadShaBakeFirstGuildInfoId] = 0
    msgLogLv[CmdChallengeAppuserUpToGsNtfId] = 0
    msgLogLv[CmdRechageCcsToGsReqId] = 1
    msgLogLv[CmdRechageGsToCcsAckId] = 0
    msgLogLv[CmdRechargeApplyReqId] = 1
    msgLogLv[CmdBanInfoCcsToGsReqId] = 1
    msgLogLv[CmdMailSendCCsToGsReqId] = 1
    msgLogLv[CmdFuncStateUpdateReqId] = 1
    msgLogLv[CmdUpAnnouncementNowReqId] = 1
    msgLogLv[CmdUpPaoMaDengNowReqId] = 1
    msgLogLv[CmdSetDayRechargeNumNtfId] = 1
}

func GetMsgPrototype(key uint16) reflect.Type {
    return msgPrototypes[key]
}

func GetMsgName(key uint16) string {
	return msgNames[key]
}

func GetMsgLogLv(key int) int {
	return msgLogLv[key]
}

const (
    CmdUnknownId uint16 = 0
    CmdErrorAckId = 1
    CmdHandShakeReqId = 2
    CmdHandShakeAckId = 3
    CmdHandCloseNtfId = 5
    CmdLoginKeyVerifyReqId = 1000
    CmdLoginKeyVerifyAckId = 1001
    CmdLoginKeyVerifyUpdateReqId = 1002
    CmdGSMessageToFSId = 2001
    CmdFSMessageToGSId = 2002
    CmdFSCallMessageToGSId = 2003
    CmdGSCallMessageToFSId = 2004
    CmdGsRouteMessageToFightId = 2005
    CmdFSCreateFightReqId = 2011
    CmdFSCreateFightAckId = 2012
    CmdFSEnterFightReqId = 2013
    CmdFSEnterFightAckId = 2014
    CmdFSLeaveFightReqId = 2015
    CmdFSLeaveFightAckId = 2016
    CmdFSUpdateUserInfoNtfId = 2017
    CmdGSToFsUpdateUserFightModelId = 2018
    CmdGsToFSUserReliveReqId = 2019
    CmdFSToGsUserReliveAckId = 2020
    CmdGsToFSCheckUserReliveReqId = 2021
    CmdFsToGSCheckUserReliveAckId = 2022
    CmdGSTOFSCheckFightReqId = 2023
    CmdGSTOFSGetFightIdReqId = 2024
    CmdGsToFsUpdateUserElfReqId = 2025
    CmdGsToFsFightNpcEventReqId = 2026
    CmdFSFightEndNtfId = 2031
    CmdFsResidentFightNtfId = 2042
    CmdGsToFsResidentFightReqId = 2043
    CmdGsTOFsPickUpReqId = 2126
    CmdFsTOGsPickUpAckId = 2127
    CmdFSAddItemReqId = 2128
    CmdFSAddItemAckId = 2129
    CmdGsToFsPickRedPacketInfoId = 2130
    CmdUserDropReqId = 2132
    CmdUserDropAckId = 2133
    CmdFsRandomDeliveryNtfId = 2134
    CmdGsToFsUseItemNtfId = 2135
    CmdGsToFsGmReqId = 2136
    CmdFsSkillUseNtfId = 2140
    CmdFsTOGsClearSkillCdNtfId = 2141
    CmdGsToFsGetCheerNumReqId = 2150
    CmdGsToFsCheerReqId = 2152
    CmdGsToFsGetPotionCdReqId = 2154
    CmdGsToFsUsePotionReqId = 2156
    CmdGsToFsCollectionReqId = 2158
    CmdFsToGsCollectionNtfId = 2160
    CmdGsToFsCollectionCancelReqId = 2161
    CmdGsToFsUseFitReqId = 2162
    CmdGsToFsFitCacelReqId = 2164
    CmdGsToFsUpdatePetReqId = 2166
    CmdGsToFsUseCutTreasureReqId = 2170
    CmdGsToFsCheckForHelpReqId = 2171
    CmdGsToFsChangeToHelperReqId = 2172
    CmdGsToFsFightNumChangeReqId = 2173
    CmdGsToFsFightScoreLessReqId = 2174
    CmdGsToFsGamedbReloadReqId = 2175
    CmdFsToGsShabakeKillBossNtfId = 2176
    CmdFsToGsActorKillNtfId = 2177
    CmdWorldBossStatusNtfId = 2201
    CmdFSContinueFightReqId = 2202
    CmdFsFieldBossInfoNtfId = 2203
    CmdExpStageKillMonsterNtfId = 2210
    CmdMiningNewFightInfoReqId = 2211
    CmdHangUpKillWaveNtfId = 2212
    CmdGuildbonfireExpAddNtfId = 2213
    CmdPaodianGoodsAddNtfId = 2214
    CmdWorldLeaderFightEndNtfId = 2220
    CmdWorldLeaderFightRankNtfId = 2221
    CmdFsFieldBossDieUserInfoNtfId = 2222
    CmdMagicTowerGetUserInfoReqId = 2223
    CmdDaBaoKillMonsterNtfId = 2224
    CmdDaBaoResumeEnergyReqId = 2225
    CmdBossFamilyBossInfoReqId = 2230
    CmdSyncUserInfoNtfId = 3001
    CmdCCSToGsCrossFsIdChangeNtfId = 3100
    CmdChallengeSendLoseRewardNtfId = 3101
    CmdChallengeAppuserUpNtfId = 3102
    CmdGsToCcsBackGuildInfoNtfId = 3103
    CmdCcsToGsBroadShaBakeFirstGuildInfoId = 3104
    CmdChallengeAppuserUpToGsNtfId = 3105
    CmdRechageCcsToGsReqId = 3200
    CmdRechageGsToCcsAckId = 3201
    CmdRechargeApplyReqId = 3202
    CmdBanInfoCcsToGsReqId = 3205
    CmdMailSendCCsToGsReqId = 3210
    CmdFuncStateUpdateReqId = 3215
    CmdUpAnnouncementNowReqId = 3216
    CmdUpPaoMaDengNowReqId = 3217
    CmdSetDayRechargeNumNtfId = 3218
)

func GetCmdIdFromType(i interface{}) uint16 {
	switch i.(type) {
	case *ErrorAck:
	     return 1
	case *HandShakeReq:
	     return 2
	case *HandShakeAck:
	     return 3
	case *HandCloseNtf:
	     return 5
	case *LoginKeyVerifyReq:
	     return 1000
	case *LoginKeyVerifyAck:
	     return 1001
	case *LoginKeyVerifyUpdateReq:
	     return 1002
	case *GSMessageToFS:
	     return 2001
	case *FSMessageToGS:
	     return 2002
	case *FSCallMessageToGS:
	     return 2003
	case *GSCallMessageToFS:
	     return 2004
	case *GsRouteMessageToFight:
	     return 2005
	case *FSCreateFightReq:
	     return 2011
	case *FSCreateFightAck:
	     return 2012
	case *FSEnterFightReq:
	     return 2013
	case *FSEnterFightAck:
	     return 2014
	case *FSLeaveFightReq:
	     return 2015
	case *FSLeaveFightAck:
	     return 2016
	case *FSUpdateUserInfoNtf:
	     return 2017
	case *GSToFsUpdateUserFightModel:
	     return 2018
	case *GsToFSUserReliveReq:
	     return 2019
	case *FSToGsUserReliveAck:
	     return 2020
	case *GsToFSCheckUserReliveReq:
	     return 2021
	case *FsToGSCheckUserReliveAck:
	     return 2022
	case *GSTOFSCheckFightReq:
	     return 2023
	case *GSTOFSGetFightIdReq:
	     return 2024
	case *GsToFsUpdateUserElfReq:
	     return 2025
	case *GsToFsFightNpcEventReq:
	     return 2026
	case *FSFightEndNtf:
	     return 2031
	case *FsResidentFightNtf:
	     return 2042
	case *GsToFsResidentFightReq:
	     return 2043
	case *GsTOFsPickUpReq:
	     return 2126
	case *FsTOGsPickUpAck:
	     return 2127
	case *FSAddItemReq:
	     return 2128
	case *FSAddItemAck:
	     return 2129
	case *GsToFsPickRedPacketInfo:
	     return 2130
	case *UserDropReq:
	     return 2132
	case *UserDropAck:
	     return 2133
	case *FsRandomDeliveryNtf:
	     return 2134
	case *GsToFsUseItemNtf:
	     return 2135
	case *GsToFsGmReq:
	     return 2136
	case *FsSkillUseNtf:
	     return 2140
	case *FsTOGsClearSkillCdNtf:
	     return 2141
	case *GsToFsGetCheerNumReq:
	     return 2150
	case *GsToFsCheerReq:
	     return 2152
	case *GsToFsGetPotionCdReq:
	     return 2154
	case *GsToFsUsePotionReq:
	     return 2156
	case *GsToFsCollectionReq:
	     return 2158
	case *FsToGsCollectionNtf:
	     return 2160
	case *GsToFsCollectionCancelReq:
	     return 2161
	case *GsToFsUseFitReq:
	     return 2162
	case *GsToFsFitCacelReq:
	     return 2164
	case *GsToFsUpdatePetReq:
	     return 2166
	case *GsToFsUseCutTreasureReq:
	     return 2170
	case *GsToFsCheckForHelpReq:
	     return 2171
	case *GsToFsChangeToHelperReq:
	     return 2172
	case *GsToFsFightNumChangeReq:
	     return 2173
	case *GsToFsFightScoreLessReq:
	     return 2174
	case *GsToFsGamedbReloadReq:
	     return 2175
	case *FsToGsShabakeKillBossNtf:
	     return 2176
	case *FsToGsActorKillNtf:
	     return 2177
	case *WorldBossStatusNtf:
	     return 2201
	case *FSContinueFightReq:
	     return 2202
	case *FsFieldBossInfoNtf:
	     return 2203
	case *ExpStageKillMonsterNtf:
	     return 2210
	case *MiningNewFightInfoReq:
	     return 2211
	case *HangUpKillWaveNtf:
	     return 2212
	case *GuildbonfireExpAddNtf:
	     return 2213
	case *PaodianGoodsAddNtf:
	     return 2214
	case *WorldLeaderFightEndNtf:
	     return 2220
	case *WorldLeaderFightRankNtf:
	     return 2221
	case *FsFieldBossDieUserInfoNtf:
	     return 2222
	case *MagicTowerGetUserInfoReq:
	     return 2223
	case *DaBaoKillMonsterNtf:
	     return 2224
	case *DaBaoResumeEnergyReq:
	     return 2225
	case *BossFamilyBossInfoReq:
	     return 2230
	case *SyncUserInfoNtf:
	     return 3001
	case *CCSToGsCrossFsIdChangeNtf:
	     return 3100
	case *ChallengeSendLoseRewardNtf:
	     return 3101
	case *ChallengeAppuserUpNtf:
	     return 3102
	case *GsToCcsBackGuildInfoNtf:
	     return 3103
	case *CcsToGsBroadShaBakeFirstGuildInfo:
	     return 3104
	case *ChallengeAppuserUpToGsNtf:
	     return 3105
	case *RechageCcsToGsReq:
	     return 3200
	case *RechageGsToCcsAck:
	     return 3201
	case *RechargeApplyReq:
	     return 3202
	case *BanInfoCcsToGsReq:
	     return 3205
	case *MailSendCCsToGsReq:
	     return 3210
	case *FuncStateUpdateReq:
	     return 3215
	case *UpAnnouncementNowReq:
	     return 3216
	case *UpPaoMaDengNowReq:
	     return 3217
	case *SetDayRechargeNumNtf:
	     return 3218
	default:
		return 0
	}
}
