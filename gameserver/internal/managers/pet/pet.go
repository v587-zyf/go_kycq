package pet

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewPetManager(m managersI.IModule) *PetManager {
	return &PetManager{
		IModule: m,
	}
}

type PetManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *PetManager) Online(user *objs.User) {
	this.CalcPetCombat(user)
	this.calcAppendageSkillEffects(user)
}

/**
 *  @Description: 战宠激活
 *  @param user
 *  @param id 	战宠id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *PetManager) Active(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.PetActiveAck) error {
	_, ok := user.Pet[id]
	if ok {
		return gamedb.ERRREPEATACTIVE
	}
	petCfg := gamedb.GetPetsConfCfg(id)
	if petCfg == nil {
		return gamedb.ERRPARAM
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, petCfg.Consume); err != nil {
		return err
	}
	user.Pet[id] = &model.Pet{
		Lv: 1,
		//Grade: 1,
		//Break: 1,
	}

	ack.Id = int32(id)
	ack.PetInfo = builder.BuildPetUnit(user.Pet[id])

	this.changeOperation(user)
	this.GetTask().AddTaskProcess(user, pb.CONDITION_ACTIVATE_ZHAN_CHONG, -1)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ACTIVATE_ZHAN_CHONG, []int{})
	return nil
}

/**
 *  @Description:战宠升级
 *  @param user
 *  @param id		战宠id
 *  @param itemId 	使用升级道具
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *PetManager) UpLv(user *objs.User, id, itemId, itemNum int, op *ophelper.OpBagHelperDefault, ack *pb.PetUpLvAck) error {
	pet, ok := user.Pet[id]
	if !ok {
		return gamedb.ERRNOTACTIVE
	}
	flag := false
	for _, id := range gamedb.GetConf().PetUpLvConsume {
		if id == itemId {
			flag = true
			break
		}
	}
	if !flag {
		return gamedb.ERRITEMCANNOTUSE
	}
	if gamedb.GetPetsLevelConfCfg(gamedb.GetRealId(id, pet.Lv+1)) == nil {
		return gamedb.ERRLVENOUGH
	}
	if err := this.GetBag().Remove(user, op, itemId, itemNum); err != nil {
		return err
	}
	petLvCfg := gamedb.GetPetsLevelConfCfg(gamedb.GetRealId(id, pet.Lv))
	itemCfg := gamedb.GetItemBaseCfg(itemId)
	pet.Exp += itemCfg.EffectVal * itemNum
	for pet.Exp >= petLvCfg.Exp {
		pet.Exp -= petLvCfg.Exp
		pet.Lv++
		nextCfg := gamedb.GetPetsLevelConfCfg(gamedb.GetRealId(id, pet.Lv))
		if nextCfg == nil {
			break
		}
		petLvCfg = nextCfg
	}

	ack.Id = int32(id)
	ack.Lv = int32(pet.Lv)
	ack.Exp = int32(pet.Exp)

	kyEvent.PetLvUp(user, id, pet.Lv)

	this.changeOperation(user)
	this.GetTask().AddTaskProcess(user, pb.CONDITION_UPGRADE_ZHAN_CHONG, 1)
	if err := this.GetFight().UserUpdatePet(user); err != nil {
		logger.Error("更新战斗中战宠数据异常,玩家：%v，异常：%v", user.IdName(), err)
	}
	return nil
}

/**
 *  OneKeyUpLv
 *  @Description: 战宠一键升级，满足本次升级道具
 *  @receiver this
 *  @param user
 *  @param petId
 *  @param itemId
 *  @param op
 *  @param ack
 *  @return error
**/
func (this *PetManager) OneKeyUpLv(user *objs.User, petId, itemId int, op *ophelper.OpBagHelperDefault, ack *pb.PetUpLvAck) error {
	pet, ok := user.Pet[petId]
	if !ok {
		return gamedb.ERRNOTACTIVE
	}
	flag := false
	for _, id := range gamedb.GetConf().PetUpLvConsume {
		if id == itemId {
			flag = true
			break
		}
	}
	if !flag {
		return gamedb.ERRITEMCANNOTUSE
	}
	if gamedb.GetPetsLevelConfCfg(gamedb.GetRealId(petId, pet.Lv+1)) == nil {
		return gamedb.ERRLVENOUGH
	}

	petLvCfg := gamedb.GetPetsLevelConfCfg(gamedb.GetRealId(petId, pet.Lv))
	itemCfg := gamedb.GetItemBaseCfg(itemId)
	needExp := petLvCfg.Exp - pet.Exp
	needItemNum := common.CeilFloat64(float64(needExp) / float64(itemCfg.EffectVal))
	isUpdate := false
	if hasNum, _ := this.GetBag().GetItemNum(user, itemId); hasNum >= needItemNum {
		if err := this.GetBag().Remove(user, op, itemId, needItemNum); err != nil {
			return err
		}
		pet.Lv++
		pet.Exp = itemCfg.EffectVal*needItemNum + pet.Exp - petLvCfg.Exp
		nextCfg := gamedb.GetPetsLevelConfCfg(gamedb.GetRealId(petId, pet.Lv))
		for pet.Exp >= nextCfg.Exp {
			pet.Exp -= nextCfg.Exp
			pet.Lv++
			if nextCfg = gamedb.GetPetsLevelConfCfg(gamedb.GetRealId(petId, pet.Lv)); nextCfg == nil {
				break
			}
		}

		isUpdate = true
	} else {
		if err := this.GetBag().Remove(user, op, itemId, hasNum); err != nil {
			return err
		}
		pet.Exp += itemCfg.EffectVal * hasNum
	}

	ack.Id = int32(petId)
	ack.Lv = int32(pet.Lv)
	ack.Exp = int32(pet.Exp)

	kyEvent.PetLvUp(user, petId, pet.Lv)

	this.changeOperation(user)
	this.GetTask().AddTaskProcess(user, pb.CONDITION_UPGRADE_ZHAN_CHONG, 1)
	if isUpdate {
		if err := this.GetFight().UserUpdatePet(user); err != nil {
			logger.Error("更新战斗中战宠数据异常,玩家：%v，异常：%v", user.IdName(), err)
		}
	}
	return nil
}

/**
 *  @Description: 战宠升阶
 *  @param user
 *  @param id	战宠id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *PetManager) UpGrade(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.PetUpGradeAck) error {
	pet, ok := user.Pet[id]
	if !ok {
		return gamedb.ERRNOTACTIVE
	}

	if gamedb.GetPetsGradeConfCfg(gamedb.GetRealId(id, pet.Grade+1)) == nil {
		return gamedb.ERRLVENOUGH
	}
	petGradeCfg := gamedb.GetPetsGradeConfCfg(gamedb.GetRealId(id, pet.Grade))
	if pet.Lv < petGradeCfg.NeedLv {
		return gamedb.ERRCONDITION
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, petGradeCfg.Consume); err != nil {
		return err
	}
	pet.Grade++

	ack.Id = int32(id)
	ack.Grade = int32(pet.Grade)
	this.changeOperation(user)
	kyEvent.PetGradeUp(user, id, pet.Grade)

	this.GetTask().AddTaskProcess(user, pb.CONDITION_UPGRADE_ZHAN_CHONG_1, -1)
	err := this.GetFight().UserUpdatePet(user)
	if err != nil {
		logger.Error("更新战斗中战宠数据异常,玩家：%v，异常：%v", user.IdName(), err)
	}
	return nil
}

/**
 *  @Description: 战宠突破
 *  @param user
 *  @param id	战宠id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *PetManager) Break(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.PetBreakAck) error {
	pet, ok := user.Pet[id]
	if !ok {
		return gamedb.ERRNOTACTIVE
	}

	if gamedb.GetPetsBreakConfCfg(gamedb.GetRealId(id, pet.Break+1)) == nil {
		return gamedb.ERRLVENOUGH
	}
	petBreakCfg := gamedb.GetPetsBreakConfCfg(gamedb.GetRealId(id, pet.Break))
	if err := this.GetBag().RemoveItemsInfos(user, op, petBreakCfg.Consume); err != nil {
		return err
	}
	pet.Break++

	ack.Id = int32(id)
	ack.Break = int32(pet.Break)

	this.changeOperation(user)
	err := this.GetFight().UserUpdatePet(user)
	if err != nil {
		logger.Error("更新战斗中战宠数据异常,玩家：%v，异常：%v", user.IdName(), err)
	}
	return nil
}

/**
*  @Description: 战宠出战、休息
*  @param user
*  @param id	战宠id
*  @return error
 */
func (this *PetManager) ChangeWear(user *objs.User, id int, ack *pb.PetChangeWearAck) error {
	_, ok := user.Pet[id]
	if !ok {
		return gamedb.ERRNOTACTIVE
	}

	userWear := user.Wear
	if userWear.PetId != id {
		userWear.PetId = id
	} else {
		userWear.PetId = 0
	}
	user.Dirty = true

	ack.PetId = int32(userWear.PetId)
	err := this.GetFight().UserUpdatePet(user)
	return err
}

func (this *PetManager) CalcPetCombat(user *objs.User) {
	// 单个战宠攻击力：(等级攻击+进阶攻击)*(1+突破百分比)
	// 战斗力: 攻击力 * 战宠攻击力系数(属性表)
	combat := 0.0
	for id, pet := range user.Pet {
		lvAtt := 0.0
		if n, ok := gamedb.GetPetsLevelConfCfg(gamedb.GetRealId(id, pet.Lv)).AttributePets[pb.PROPERTY_ATT_PETS]; ok {
			lvAtt = float64(n)
		}
		gradeAtt := 0.0
		if n, ok := gamedb.GetPetsGradeConfCfg(gamedb.GetRealId(id, pet.Grade)).AttributePets[pb.PROPERTY_ATT_PETS]; ok {
			gradeAtt = float64(n)
		}
		breakAtt := 0
		if n, ok := gamedb.GetPetsBreakConfCfg(gamedb.GetRealId(id, pet.Break)).AttributePets[pb.PROPERTY_ATT_PETS_RATE]; ok {
			breakAtt = n
		}
		att, attRate := 0.0, 0
		if petAddAttr, ok := user.PetAddAttr[id]; ok {
			if n, ok := petAddAttr[int32(pb.PROPERTY_ATT_PETS)]; ok {
				att = float64(n)
			}
			if n, ok := petAddAttr[int32(pb.PROPERTY_ATT_PETS_RATE)]; ok {
				attRate = int(n)
			}
		}
		attack := (lvAtt + gradeAtt + att) * (1.0 + float64(breakAtt+attRate)/10000)
		combat += attack * gamedb.GetPropertyPropertyCfg(pb.PROPERTY_ATT_PETS).Combat[0]
	}
	for _, hero := range user.Heros {
		hero.ModuleCombat[pb.PROPERTYMODULE_PET] = int(combat)
	}
	user.PetCombat = int(combat)
	logger.Debug("--------------------pet combat:%v", combat)
}

func (this *PetManager) changeOperation(user *objs.User) {
	user.Dirty = true
	this.CalcPetCombat(user)
	this.GetUserManager().UpdateCombat(user, -1)
}
