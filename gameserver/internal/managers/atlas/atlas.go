package atlas

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewAtlasManager(module managersI.IModule) *AtlasManager {
	atlas := &AtlasManager{}
	atlas.IModule = module
	return atlas
}

type AtlasManager struct {
	util.DefaultModule
	managersI.IModule
}

/**
 *  @Description: 图鉴激活
 *  @param user
 *  @param id	图鉴id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *AtlasManager) AtlasActive(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.AtlasActiveAck) error {
	if id <= 0 {
		return gamedb.ERRPARAM
	}
	if user.Atlases[id] != 0 {
		return gamedb.ERRREPEATACTIVE
	}

	conf := gamedb.GetAtlasAtlasCfg(id)
	if conf == nil {
		return gamedb.ERRPARAM
	}
	if err := this.GetBag().Remove(user, op, conf.Consume.ItemId, conf.Consume.Count); err != nil {
		return err
	}
	user.Atlases[id] = 1

	ack.Atlas = builder.BuilderAtlasUnit(id, user.Atlases[id])
	this.changeOperation(user)
	this.GetTask().AddTaskProcess(user, pb.CONDITION_ACTIVATE_TU_JIAN, -1)
	return nil
}

/**
 *  @Description: 图鉴升星
 *  @param user
 *  @param id	图鉴id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *AtlasManager) AtlasUpStar(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.AtlasUpStarAck) error {
	if _, ok := user.Atlases[id]; !ok {
		return gamedb.ERRATLASNOTACTIVE
	}

	maxStar := gamedb.GetMaxValById(id, constMax.MAX_ATLAS_STAR)
	if user.Atlases[id] >= maxStar {
		return gamedb.ERRATLASSTARENOUGH
	}

	conf := gamedb.GetAtlasStar(id, user.Atlases[id])
	if conf == nil {
		return gamedb.ERRPARAM
	}
	itemId, itemCount := conf.Consume.ItemId, conf.Consume.Count
	if err := this.GetBag().Remove(user, op, itemId, itemCount); err != nil {
		return gamedb.ERRNOTENOUGHGOODS
	}
	user.Atlases[id]++

	kyEvent.AtlasStarUp(user, conf.Type, id, user.Atlases[id])

	ack.Atlas = builder.BuilderAtlasUnit(id, user.Atlases[id])
	this.changeOperation(user)
	return nil
}

/**
 *  @Description: 图鉴连携激活
 *  @param user
 *  @param id
 *  @param ack
 *  @return error
 */
func (this *AtlasManager) AtlasGatherActive(user *objs.User, id int, ack *pb.AtlasGatherActiveAck) error {
	if user.AtlasGathers[id] != 0 {
		return gamedb.ERRATLASGATHERREPEATACTIVE
	}
	conf := gamedb.GetAtlasGather(id)
	if conf == nil {
		return gamedb.ERRPARAM
	}
	atlas := user.Atlases
	for _, v := range conf.Atlas_id {
		if _, ok := atlas[v]; !ok {
			return gamedb.ERRATLASNOTACTIVE
		}
	}
	user.AtlasGathers[id] = 1

	ack.AtlasGather = builder.BuilderAtlasGatherUnit(id, user.AtlasGathers[id])
	this.changeOperation(user)
	return nil
}

/**
 *  @Description: 图鉴连携升星
 *  @param user
 *  @param id
 *  @param ack
 *  @return error
 */
func (this *AtlasManager) AtlasGatherUpStar(user *objs.User, id int, ack *pb.AtlasGatherUpStarAck) error {
	maxStar := gamedb.GetMaxValById(id, constMax.MAX_ATLAS_UPGRADE)
	if user.AtlasGathers[id] >= maxStar {
		return gamedb.ERRATLASGATHERSTARENOUGH
	}
	atlasGatherConf := gamedb.GetAtlasGather(id)
	if atlasGatherConf == nil {
		return gamedb.ERRPARAM
	}
	needStar := gamedb.GetAtlasUpgrade(id, user.AtlasGathers[id]+1).Star
	atlas := user.Atlases
	for _, v := range atlasGatherConf.Atlas_id {
		if star, ok := atlas[v]; !ok || star < needStar {
			return gamedb.ERRATLASNOTACTIVE
		}
	}
	user.AtlasGathers[id]++

	ack.AtlasGather = builder.BuilderAtlasGatherUnit(id, user.AtlasGathers[id])
	this.changeOperation(user)
	return nil
}

/**
 *  @Description: 图鉴穿戴
 *  @param user
 *  @param heroIndex
 *  @param id
 *  @param ack
 *  @return error
 */
func (this *AtlasManager) Change(user *objs.User, heroIndex, id int, ack *pb.AtlasWearChangeAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	if _, ok := user.Atlases[id]; !ok {
		return gamedb.ERRNOTACTIVE
	}
	atlasWears := hero.Wear.AtlasWear
	atlasWearLen := len(atlasWears)
	if atlasWearLen > 0 {
		if atlasWears[id] != 0 {
			delete(atlasWears, id)
			ack.RemoveId = int32(id)
		} else {
			if check := this.GetCondition().CheckMulti(user, heroIndex, gamedb.GetAtlasPosAtlasPosCfg(atlasWearLen+1).Condition); !check {
				return gamedb.ERRWEARENOUGH
			}
			if atlasWearLen >= gamedb.GetMaxValById(0, constMax.MAX_ATLAS_WEAR) {
				return gamedb.ERRWEARENOUGH
			}
			//已经穿戴过，校验类型是否一致
			for hasId := range atlasWears {
				hasCfgT := gamedb.GetAtlasAtlasCfg(hasId).Type
				atlasCfgT := gamedb.GetAtlasAtlasCfg(id).Type
				if hasCfgT != atlasCfgT {
					return gamedb.ERRATLASTYPE
				}
			}
			//其他英雄已穿戴过不能再穿
			for index, hero := range user.Heros {
				if heroIndex == index {
					continue
				}
				if _, ok := hero.Wear.AtlasWear[id]; ok {
					return gamedb.ERRREPEATWEAR
				}
			}
			atlasWears[id] = id
		}
	} else {
		atlasWears[id] = id
	}
	kyEvent.AtlasChange(user, heroIndex, id)

	ack.HeroIndex = int32(heroIndex)
	ack.AtlasWear = builder.BuildAtlasWear(atlasWears)
	this.GetUserManager().UpdateCombat(user, heroIndex)
	this.GetTask().AddTaskProcess(user, pb.CONDITION_WEAR_TU_JIAN, 1)
	return nil
}

func (this *AtlasManager) AtlasGm(user *objs.User, gmT string) {
	userAtlas := user.Atlases
	userAtlasGathers := user.AtlasGathers
	atlasCfgs := gamedb.GetAtlasCfgs()
	atlasGathersCfgs := gamedb.GetAtlasGatherCfgs()
	for id := range atlasCfgs {
		star := 0
		switch gmT {
		case "active":
		case "up":
			star = gamedb.GetMaxValById(id, constMax.MAX_ATLAS_STAR)
		}
		userAtlas[id] = star
		if star == 0 {
			this.GetUserManager().SendMessage(user, &pb.AtlasActiveAck{Atlas: builder.BuilderAtlasUnit(id, star)}, true)
		} else {
			this.GetUserManager().SendMessage(user, &pb.AtlasUpStarAck{Atlas: builder.BuilderAtlasUnit(id, star)}, true)
		}
	}
	for id := range atlasGathersCfgs {
		star := 0
		switch gmT {
		case "active":
		case "up":
			star = gamedb.GetMaxValById(id, constMax.MAX_ATLAS_UPGRADE)
		}
		userAtlasGathers[id] = star
		if star == 0 {
			this.GetUserManager().SendMessage(user, &pb.AtlasGatherActiveAck{AtlasGather: builder.BuilderAtlasGatherUnit(id, star)}, true)
		} else {
			this.GetUserManager().SendMessage(user, &pb.AtlasGatherUpStarAck{AtlasGather: builder.BuilderAtlasGatherUnit(id, star)}, true)
		}
	}
	this.changeOperation(user)
}

func (this *AtlasManager) changeOperation(user *objs.User) {
	this.GetUserManager().UpdateCombat(user, -1)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ATLAS_STAR, []int{})
}
