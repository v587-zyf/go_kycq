package materialStage

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

func (this *MaterialStage) EnterMaterialStageFight(user *objs.User, stageId int) error {
	if stageId < 1 {
		return gamedb.ERRPARAM
	}
	_, _, _, err := this.CheckFight(user, stageId)
	if err != nil {
		return err
	}

	fightId, err := this.GetFight().CreateFight(stageId, nil)
	if err != nil {
		return err
	}
	err = this.GetFight().EnterFightByFightId(user, stageId, fightId)
	if err != nil {
		return err
	}
	return nil
}

func (this *MaterialStage) MaterialStageFightResultNtf(user *objs.User, isWin bool, stageId int) error {
	materialCfg := gamedb.GetMaterialByStageId(stageId)
	materialType := materialCfg.Type
	ntf := &pb.MaterialStageFightResultNtf{
		StageId:      int32(stageId),
		MaterialType: int32(materialType),
		Result:       pb.RESULTFLAG_FAIL,
	}

	if isWin {
		user.MaterialStage.MaterialStages[materialType].DareNum++
		layer := user.MaterialStage.MaterialStages[materialType].NowLayer
		if materialCfg.Level >= layer {
			maxLayer := gamedb.GetMaxValById(materialType, constMax.MAX_MATERIAL_LEVEL)
			forNum := 1
			if user.MaterialStage.MaterialStages[materialType].LastLayer == layer {
				forNum = 2
			}
			for i := 0; i < forNum; i++ {
				nowLayer := user.MaterialStage.MaterialStages[materialType].NowLayer
				if nowLayer < maxLayer {
					nextLvCfg := gamedb.GetMaterialByTypeAndLv(materialType, nowLayer+1)
					if this.GetCondition().CheckMulti(user, -1, nextLvCfg.Conditon) {
						user.MaterialStage.MaterialStages[materialType].NowLayer++
					}
				}
			}
			if user.MaterialStage.MaterialStages[materialType].LastLayer == layer {
				user.MaterialStage.MaterialStages[materialType].LastLayer++
			} else {
				user.MaterialStage.MaterialStages[materialType].LastLayer = layer
			}
		}
		ntf.Result = pb.RESULTFLAG_SUCCESS

		op := ophelper.NewOpBagHelperDefault(constBag.OpTypeMaterialStage)
		rItemId, rCount := materialCfg.Reward.ItemId, materialCfg.Reward.Count
		//if monthCardPrivilege := this.GetVipManager().GetPrivilege(user, pb.MONTHCARDPRIVILEGE_MATERIALSTAGE); monthCardPrivilege != 0 {
		//	rCount = common.CalcTenThousand(monthCardPrivilege, rCount)
		//}
		err := this.GetBag().Add(user, op, rItemId, rCount)
		if err != nil {
			bags := make([]*model.Item, 0)
			bags = append(bags, &model.Item{
				ItemId: rItemId,
				Count:  rCount,
			})
			this.GetMail().SendSystemMail(user.Id, constMail.MATERIAL_REWARD, []string{gamedb.GetMaterialHomeMaterialHomeCfg(materialType).Name}, bags, 0)
		}
		ntf.Goods = op.ToChangeItems()
		this.GetUserManager().SendItemChangeNtf(user, op)
	}
	this.ChangeOperation(user, materialType, isWin)

	ntf.DareNum = int32(user.MaterialStage.MaterialStages[materialType].DareNum)
	ntf.NowLayer = int32(user.MaterialStage.MaterialStages[materialType].NowLayer)
	ntf.LastLayer = int32(user.MaterialStage.MaterialStages[materialType].LastLayer)
	this.GetUserManager().SendMessage(user, ntf, true)
	return nil
}

/**
 *  @Description: 材料副本扫荡
 *  @param user
 *  @param stageId
 *  @return error
 */
func (this *MaterialStage) MaterialStageSweep(user *objs.User, stageId int, op *ophelper.OpBagHelperDefault, ack *pb.MaterialStageSweepAck) error {
	if stageId < 1 {
		return gamedb.ERRPARAM
	}
	userMaterial, _, stageCfg, err := this.CheckFight(user, stageId)
	if err != nil {
		return err
	}
	materialType := stageCfg.Type
	if userMaterial.MaterialStages[materialType].LastLayer < stageCfg.Level {
		return gamedb.ERRNOTOPEN
	}

	if privilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COPY_SWEEP); privilege < 1 {
		return gamedb.ERRVIPLVNOTENOUGH
	}
	this.GetBag().Add(user, op, stageCfg.Reward.ItemId, stageCfg.Reward.Count)

	userMaterial.MaterialStages[materialType].DareNum++
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_KILL_STAGE, []int{stageId, 1})
	this.ChangeOperation(user, materialType, true)

	ack.MaterialType = int32(materialType)
	ack.SweepNum = int32(userMaterial.MaterialStages[materialType].DareNum)
	ack.Goods = op.ToChangeItems()
	return nil
}

func (this *MaterialStage) CheckFight(user *objs.User, stageId int) (*model.MaterialStage, *gamedb.MaterialHomeMaterialHomeCfg, *gamedb.MaterialStageMaterialStageCfg, error) {
	userMaterial := user.MaterialStage

	stageCfg := gamedb.GetMaterialByStageId(stageId)
	if stageCfg == nil {
		return userMaterial, nil, nil, gamedb.ERRPARAM
	}
	materialType := stageCfg.Type
	homeCfg := gamedb.GetMaterialHomeMaterialHomeCfg(materialType)
	if homeCfg == nil {
		return userMaterial, nil, nil, gamedb.ERRPARAM
	}

	fightNum := userMaterial.MaterialStages[materialType].DareNum
	buyNum := userMaterial.MaterialStages[materialType].BuyNum
	if fightNum >= homeCfg.Challenge+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_MATERIALSTAGE_FIGHTNUM)+buyNum {
		return userMaterial, nil, nil, gamedb.ERRFIGHTNUMNOTENOUGH
	}
	if check := this.GetCondition().CheckMulti(user, -1, stageCfg.Conditon); !check {
		return userMaterial, nil, nil, gamedb.ERRCONDITION
	}

	if !this.GetFight().CheckInFightBefore(user, stageId) {
		return userMaterial, nil, nil, gamedb.ERRUSERINFIGHT
	}

	return userMaterial, homeCfg, stageCfg, nil
}

func (this *MaterialStage) ChangeOperation(user *objs.User, materialType int, isWin bool) {
	user.Dirty = true
	var dailyTaskActivityType int
	//fightNum := user.MaterialStage.MaterialStages[materialType].DareNum
	switch materialType {
	case pb.MATERIALSTAGETYPE_WING:
		dailyTaskActivityType = pb.DAILYTASKACTIVITYTYPE_YU_YI_FU_BEN
		//通知任务系统，挑战了一次神翼副本(材料副本)
		this.GetCondition().RecordCondition(user, pb.CONDITION_TIAO_ZHAN_SHEN_YI_FU_BEN, []int{1})
		this.GetTask().AddTaskProcess(user, pb.CONDITION_TIAO_ZHAN_SHEN_YI_FU_BEN, -1)
	case pb.MATERIALSTAGETYPE_GOLD:
		dailyTaskActivityType = pb.DAILYTASKACTIVITYTYPE_JIN_BI_FU_BEN
		//通知任务系统，挑战了一次金币副本(材料副本)
		//this.GetTask().AddTaskProcess(user, pb.CONDITION_CHALLENGE_QIANG_HUA_FU_BEN, 1)
	case pb.MATERIALSTAGETYPE_STRENGTH:
		//通知任务系统，挑战了一次强化副本(材料副本)、
		//dailyTaskActivityType = pb.DAILYTASKACTIVITYTYPE_JIN_BI_FU_BEN
		this.GetCondition().RecordCondition(user, pb.CONDITION_CHALLENGE_QIANG_HUA_FU_BEN, []int{1})
		this.GetTask().AddTaskProcess(user, pb.CONDITION_CHALLENGE_QIANG_HUA_FU_BEN, -1)
		this.GetDailyTask().CompletionOfTask(user, pb.DAILYTASKACTIVITYTYPE_QIANG_HUA_FU_BEN, 1)
	case pb.MATERIALSTAGETYPE_FASHION:
		// 时装副本
	}
	if isWin {
		this.GetDailyTask().CompletionOfTask(user, dailyTaskActivityType, 1)
		this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_COMPLETE_COPY, []int{1, constFight.FIGHT_TYPE_MATERIAL})
		this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_FINISH_COPY, []int{1})
	}
}
