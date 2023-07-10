package ring

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewRingManager(module managersI.IModule) *RingManager {
	return &RingManager{IModule: module}
}

type RingManager struct {
	util.DefaultModule
	managersI.IModule
}

/**
 *  @Description: 特戒穿戴
 *  @param user
 *  @param heroIndex
 *  @param ringPos	戒指位置
 *  @param bagPos	背包位置
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *RingManager) Wear(user *objs.User, heroIndex, ringPos, bagPos int, op *ophelper.OpBagHelperDefault, ack *pb.RingWearAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	if _, ok := pb.RINGPOS_MAP[ringPos]; !ok {
		return gamedb.ERRPARAM
	}
	itemId := this.GetBag().GetItemByPosition(user, bagPos).ItemId
	itemCfg := gamedb.GetItemBaseCfg(itemId)
	if itemCfg == nil || itemCfg.Type != pb.ITEMTYPE_SPECIAL_CAUTION {
		return gamedb.ERRITEMCANNOTUSE
	}

	ring := hero.Ring[ringPos]
	if ring.Rid != 0 {
		if err := this.GetBag().Add(user, op, ring.Rid, 1); err != nil {
			return err
		}
	}
	if err := this.GetBag().Remove(user, op, itemId, 1); err != nil {
		return err
	}
	ring.Phantom = this.UpdatePhantom(itemId, ring)
	ring.Rid = itemId

	ack.HeroIndex = int32(heroIndex)
	ack.RingPos = int32(ringPos)
	ack.Ring = builder.BuildRing(ring)
	this.GetUserManager().UpdateCombat(user, heroIndex)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_HERO_WEAR_RING, []int{})
	return nil
}

/**
 *  @Description: 特戒卸下
 *  @param user
 *  @param heroIndex
 *  @param ringPos	特戒位置
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *RingManager) Remove(user *objs.User, heroIndex, ringPos int, op *ophelper.OpBagHelperDefault, ack *pb.RingRemoveAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	if _, ok := pb.RINGPOS_MAP[ringPos]; !ok {
		return gamedb.ERRPARAM
	}
	ring := hero.Ring[ringPos]
	if ring.Rid == 0 {
		return gamedb.ERRNOTWEARRING
	}
	if err := this.GetBag().Add(user, op, ring.Rid, 1); err != nil {
		return err
	}

	ring.Phantom = this.UpdatePhantom(ring.Rid, ring)
	ring.Rid = 0

	ack.HeroIndex = int32(heroIndex)
	ack.RingPos = int32(ringPos)
	ack.Ring = builder.BuildRing(ring)
	this.GetUserManager().UpdateCombat(user, heroIndex)
	return nil
}

/**
 *  @Description: 特戒强化
 *  @param user
 *  @param heroIndex
 *  @param ringPos	特戒位置
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *RingManager) Strengthen(user *objs.User, heroIndex, ringPos int, op *ophelper.OpBagHelperDefault, ack *pb.RingStrengthenAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	if _, ok := pb.RINGPOS_MAP[ringPos]; !ok {
		return gamedb.ERRPARAM
	}
	ring := hero.Ring[ringPos]
	if ring.Rid == 0 {
		return gamedb.ERRNOTWEARRING
	}
	if gamedb.GetRingStrengthenRingStrengthenCfg(ring.Strengthen+1) == nil {
		return gamedb.ERRLVENOUGH
	}
	strengthenCfg := gamedb.GetRingStrengthenRingStrengthenCfg(ring.Strengthen)
	if check := this.GetCondition().CheckMulti(user, heroIndex, strengthenCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}
	if err := this.GetBag().Remove(user, op, strengthenCfg.Consume.ItemId, strengthenCfg.Consume.Count); err != nil {
		return err
	}
	ring.Strengthen++

	kyEvent.RingStrengthen(user, heroIndex, ringPos, ring.Strengthen)
	this.GetCondition().RecordCondition(user, pb.CONDITION_TE_JIE_QIANG_HUA_TIMES, []int{1})
	ack.HeroIndex = int32(heroIndex)
	ack.RingPos = int32(ringPos)
	ack.Ring = builder.BuildRing(ring)
	this.GetUserManager().UpdateCombat(user, heroIndex)
	return nil
}

/**
 *  @Description: 特戒强化二
 *  @param user
 *  @param heroIndex
 *  @param ringPos 特戒位置
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *RingManager) RingPhantom(user *objs.User, heroIndex, ringPos int, op *ophelper.OpBagHelperDefault, ack *pb.RingPhantomAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	if _, ok := pb.RINGPOS_MAP[ringPos]; !ok {
		return gamedb.ERRPARAM
	}
	ring := hero.Ring[ringPos]
	if ring.Rid == 0 {
		return gamedb.ERRNOTWEARRING
	}
	if gamedb.GetRingPhantomRingPhantomCfg(ring.Pid+1) == nil {
		return gamedb.ERRLVENOUGH
	}
	ringPhantomCfg := gamedb.GetRingPhantomRingPhantomCfg(ring.Pid)
	if err := this.GetBag().Remove(user, op, ringPhantomCfg.Consume.ItemId, ringPhantomCfg.Consume.Count); err != nil {
		return err
	}
	ring.Pid++
	oldTalent, newTalent := ring.Talent, gamedb.GetRingPhantomRingPhantomCfg(ring.Pid).Talent
	talentNum := newTalent - oldTalent
	for pos := range ring.Phantom {
		ring.Phantom[pos].Talent += talentNum
	}
	ring.Talent = newTalent
	kyEvent.RingPhantomStarUp(user, heroIndex, ringPos, ring.Pid)

	ack.HeroIndex = int32(heroIndex)
	ack.RingPos = int32(ringPos)
	ack.Ring = builder.BuildRing(ring)
	this.GetUserManager().UpdateCombat(user, heroIndex)
	return nil
}

/**
 *  @Description: 戒灵技能激活升级
 *  @param user
 *  @param heroIndex
 *  @param ringPos		特戒位置
 *  @param phantomPos	戒灵id
 *  @param skillId		技能id
 *  @param ack
 *  @return error
 */
func (this *RingManager) PhantomSkill(user *objs.User, heroIndex, ringPos, phantomPos, skillId int, ack *pb.RingSkillUpAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	if _, ok := pb.RINGPOS_MAP[ringPos]; !ok {
		return gamedb.ERRPARAM
	}
	if _, ok := pb.RINGPHANTOMPOS_MAP[phantomPos]; !ok {
		return gamedb.ERRPARAM
	}
	ring := hero.Ring[ringPos]
	if ring.Rid == 0 || ring == nil {
		return gamedb.ERRNOTWEARRING
	}
	if _, ok := pb.RINGSKILLID_MAP[skillId]; !ok {
		return gamedb.ERRPARAM
	}
	if !pb.RINGPHANTOMPOS_MAP[phantomPos] {
		return gamedb.ERRPARAM
	}
	heroPhantom := ring.Phantom[phantomPos]
	var phantomCfg *gamedb.PhantomLevelPhantomLevelCfg
	var needTalent int
	if skillId == pb.RINGSKILLID_ONE {
		phantomCfg = gamedb.GetPhantomSkill1(heroPhantom.Phantom, skillId, heroPhantom.Skill[skillId])
	} else {
		phantomCfg = gamedb.GetPhantomSkill2(heroPhantom.Phantom, skillId, heroPhantom.Skill[skillId])
	}
	if phantomCfg == nil {
		return gamedb.ERRPARAM
	}

	if skillId == pb.RINGSKILLID_ONE {
		needTalent = phantomCfg.Talent1
	} else {
		needTalent = phantomCfg.Talent2
	}
	if needTalent > heroPhantom.Talent {
		return gamedb.ERRTALENTNOTENOUGH
	}
	heroPhantom.Skill[skillId]++
	heroPhantom.Talent -= needTalent

	ack.HeroIndex = int32(heroIndex)
	ack.RingPos = int32(ringPos)
	ack.Ring = builder.BuildRing(ring)
	this.GetUserManager().UpdateCombat(user, heroIndex)
	return nil
}

/**
 *  @Description: 特戒融合
 *  @param user
 *  @param id
 *  @param bagPos1	背包位置1
 *  @param bagPos2	背包位置2
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *RingManager) Fuse(user *objs.User, id, bagPos1, bagPos2 int, op *ophelper.OpBagHelperDefault, ack *pb.RingFuseAck) error {
	ringCfg := gamedb.GetRingRingCfg(id)
	if ringCfg == nil {
		return gamedb.ERRPARAM
	}
	item1 := this.GetBag().GetItemByPosition(user, bagPos1)
	item2 := this.GetBag().GetItemByPosition(user, bagPos2)
	item1Cfg := gamedb.GetItemBaseCfg(item1.ItemId)
	item2Cfg := gamedb.GetItemBaseCfg(item2.ItemId)
	if item1Cfg.Type != pb.ITEMTYPE_SPECIAL_CAUTION || item2Cfg.Type != pb.ITEMTYPE_SPECIAL_CAUTION {
		return gamedb.ERRITEMCANNOTUSE
	}

	consumeMap := make(map[int]int)
	for _, itemInfo := range ringCfg.Consume {
		if item1Cfg.Id != itemInfo.ItemId && item2Cfg.Id != itemInfo.ItemId {
			return gamedb.ERRPARAM
		}
		itemNum, _ := this.GetBag().GetItemNum(user, itemInfo.ItemId)
		if itemNum < itemInfo.Count {
			return gamedb.ERRNOTENOUGHGOODS
		}
		consumeMap[itemInfo.ItemId] += itemInfo.Count
	}
	itemNum, _ := this.GetBag().GetItemNum(user, ringCfg.Consume1.ItemId)
	if itemNum < ringCfg.Consume1.Count {
		return gamedb.ERRNOTENOUGHGOODS
	}
	consumeMap[ringCfg.Consume1.ItemId] += ringCfg.Consume1.Count

	for itemId, count := range consumeMap {
		this.GetBag().Remove(user, op, itemId, count)
	}
	if err := this.GetBag().Add(user, op, id, 1); err != nil {
		return err
	}
	ack.Id = int32(id)
	kyEvent.RingFuse(user, id, consumeMap)
	return nil
}

/**
 *  @Description: 特戒重置技能
 *  @param user
 *  @param heroIndex
 *  @param ringPos		特戒位置
 *  @param phantomPos	戒灵id
 *  @param ack
 *  @return error
 */
func (this *RingManager) ResetSkill(user *objs.User, heroIndex, ringPos, phantomPos int, ack *pb.RingSkillResetAck) error {
	if _, ok := pb.RINGPOS_MAP[ringPos]; !ok {
		return gamedb.ERRPARAM
	}
	if _, ok := pb.RINGPHANTOMPOS_MAP[phantomPos]; !ok {
		return gamedb.ERRPARAM
	}
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	ring := hero.Ring[ringPos]
	if ring.Rid == 0 || ring == nil {
		return gamedb.ERRNOTWEARRING
	}

	if !pb.RINGPHANTOMPOS_MAP[phantomPos] {
		return gamedb.ERRPARAM
	}
	skill := make(model.IntKv)
	for _, skillId := range pb.RINGSKILLID_ARRAY {
		skill[skillId] = 0
	}
	ring.Phantom[phantomPos] = &model.RingPhantom{Skill: skill, Talent: ring.Talent, Phantom: ring.Phantom[phantomPos].Phantom}

	ack.HeroIndex = int32(heroIndex)
	ack.RingPos = int32(ringPos)
	ack.Ring = builder.BuildRing(ring)
	this.GetUserManager().UpdateCombat(user, heroIndex)
	return nil
}

func (this *RingManager) UpdatePhantom(itemId int, ring *model.RingUnit) map[int]*model.RingPhantom {
	phantom := make(map[int]*model.RingPhantom)
	ringCfg := gamedb.GetRingRingCfg(itemId)
	phantomCfgLen := len(ringCfg.Phantom) - 1
	for i, id := range pb.RINGPHANTOMPOS_ARRAY {
		phantomNum := 0
		if phantomCfgLen >= i {
			phantomNum = ringCfg.Phantom[i]
		}
		skill := make(model.IntKv)
		for _, skillId := range pb.RINGSKILLID_ARRAY {
			skill[skillId] = 0
		}
		phantom[id] = &model.RingPhantom{Skill: skill, Talent: ring.Talent, Phantom: phantomNum}
	}
	return phantom
}
