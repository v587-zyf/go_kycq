package pet

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

/**
 *  @Description: 战宠附体强化
 *  @param user
 *  @param pid
 *  @param op
 *  @return error
 */
func (this *PetManager) AppendageStrengthen(user *objs.User, pid int, op *ophelper.OpBagHelperDefault) error {
	petAppendage := user.PetAppendage
	lv := petAppendage[pid]
	if gamedb.GetPetAppendageByPidAndLv(pid, lv+1) == nil {
		return gamedb.ERRLVENOUGH
	}
	appendageCfg := gamedb.GetPetAppendageByPidAndLv(pid, lv)
	if appendageCfg == nil {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("pid:%v lv:%v", pid, lv)
	}
	for _, condition := range appendageCfg.Condition {
		if _, check := this.GetCondition().CheckBySlice(user, -1, condition); !check {
			return gamedb.ERRCONDITION
		}
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, appendageCfg.Consume); err != nil {
		return err
	}
	petAppendage[pid]++
	user.Dirty = true
	this.calcAppendageSkillEffects(user)
	this.CalcPetCombat(user)
	this.GetUserManager().UpdateCombat(user, -1)
	if err := this.GetFight().UserUpdatePet(user); err != nil {
		return err
	}
	return nil
}

func (this *PetManager) calcAppendageSkillEffects(user *objs.User) {
	effectMap := make(map[int]int)
	cfgs := gamedb.GetPetAppendageSkillCfgs()
	petAddSkills := make([]int32, 0)
	for _, cfg := range cfgs {
		check := true
		for _, condition := range cfg.Condition {
			if _, flag := this.GetCondition().CheckBySlice(user, -1, condition); !flag {
				check = false
				break
			}
		}
		if check {
			//战宠添加技能
			if cfg.Type == constUser.PET_APPENDAGE_EFFECT_PET {
				if effectCfg := gamedb.GetEffectEffectCfg(cfg.Effect); effectCfg != nil {
					petAddSkills = append(petAddSkills, int32(effectCfg.Skillevelid))
				}
			} else {
				effectMap[cfg.Effect] = 0
			}
		}
	}
	user.PetAddSkills = petAddSkills
	if len(effectMap) > 0 {
		user.PetAppendageEffects = effectMap
	}

	petAddAttr := make(map[int]map[int32]int64)
	wearPetId := user.Wear.PetId
	if wearPetId > 0 {
		if petAddAttr[wearPetId] == nil {
			petAddAttr[wearPetId] = make(map[int32]int64)
		}
		petsAddCfg := gamedb.GetPetAppendageByPidAndLv(wearPetId, user.PetAppendage[wearPetId])
		if petsAddCfg != nil {
			for pid, pVal := range petsAddCfg.AttributePets {
				if pid == pb.PROPERTY_ATT_PETS || pid == pb.PROPERTY_ATT_PETS_RATE {
					petAddAttr[wearPetId][int32(pid)] += int64(pVal)
				}
			}
		}
	}
	user.PetAddAttr = petAddAttr
}
