package applets

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

func NewAppletsManager(m managersI.IModule) *Applets {
	return &Applets{
		IModule: m,
	}
}

type Applets struct {
	util.DefaultModule
	managersI.IModule
}

//
//  EnterAppletsReq
//  @Description: 进入小程序
//  @receiver this
//  @param user
//  @param appletsType 小程序类型
//  @param ntf
//  @return error
//
func (this *Applets) EnterAppletsReq(user *objs.User, appletsType int, ntf *pb.AppletsEnergyNtf) error {

	if !pb.APPLETSTYPE_MAP[appletsType] {
		return gamedb.ERRPARAM
	}

	if user.AppletsInfo.List[appletsType] == nil {
		user.AppletsInfo.List[appletsType] = &model.AppletsUnit{}
	}

	if user.AppletsInfo.Energy <= 0 {
		return gamedb.ERRENERGY
	}

	if user.AppletsInfo.Energy < gamedb.GetConf().XiaoYouXiCostEnergy {
		return gamedb.ERRENERGY
	}

	if user.AppletsInfo.Energy >= gamedb.GetConf().XiaoYouXiEnergy {
		user.AppletsInfo.ResumeTime = time.Now().Unix() + int64(gamedb.GetConf().XiaoYouXiEnergyResume[0])
	}

	for _, data := range user.AppletsInfo.List {
		if data.Stage > 0 {
			user.AppletsInfo.Energy -= gamedb.GetConf().XiaoYouXiCostEnergy
			if user.AppletsInfo.Energy < 0 {
				user.AppletsInfo.Energy = 0
			}
			break
		}
	}
	user.AppletsInfo.List[appletsType].IsInGame = 1
	user.Dirty = true
	ntf.Energy = int32(user.AppletsInfo.Energy)
	ntf.ResumeTime = user.AppletsInfo.ResumeTime
	return nil
}

//  AppletsReceiveReq
//  @Description: 领取魔法射击  杀怪奖励
//  @receiver this
//  @param user
//  @param receiveId archerElement 配置表 id
func (this *Applets) AppletsReceiveReq(user *objs.User, receiveId int, ack *pb.AppletsReceiveAck, op *ophelper.OpBagHelperDefault) error {

	if user.AppletsInfo.List[pb.APPLETSTYPE_BOW] == nil {
		user.AppletsInfo.List[pb.APPLETSTYPE_BOW] = &model.AppletsUnit{}
	}

	if user.AppletsInfo.List[pb.APPLETSTYPE_BOW].IsInGame <= 0 {
		return gamedb.ERRPARAM
	}

	cfg := gamedb.GetArcherElementMagicElementCfg(receiveId)
	if cfg == nil || cfg.Reward == nil || len(cfg.Reward) <= 0 {
		return gamedb.ERRPARAM
	}

	this.GetBag().AddItems(user, cfg.Reward, op)
	ack.ReceiveId = int32(receiveId)
	ack.Goods = op.ToChangeItems()
	return nil
}

//
//  CronGetAwardReq
//  @Description: 魔法射击定时奖励获取
//  @receiver this
//  @param user
//  @param id archer-magic 配置表id
//  @param index archer-magic   InsistReward 字段索引
//  @param ack
//  @param op
//  @return error
//
func (this *Applets) CronGetAwardReq(user *objs.User, id, index int, ack *pb.CronGetAwardAck, op *ophelper.OpBagHelperDefault) error {

	cfg := gamedb.GetArcherMagicCfg(id)
	if cfg == nil || cfg.InsistReward == nil || len(cfg.InsistReward) < index+1 {
		return gamedb.ERRPARAM
	}
	if user.AppletsInfo.List[pb.APPLETSTYPE_BOW] == nil {
		user.AppletsInfo.List[pb.APPLETSTYPE_BOW] = &model.AppletsUnit{}
	}

	if user.AppletsInfo.List[pb.APPLETSTYPE_BOW].IsInGame <= 0 {
		return gamedb.ERRPARAM
	}

	if int(time.Now().Unix()-user.AppletsInfo.List[pb.APPLETSTYPE_BOW].LastGetAwardTime) < cfg.InsistReward[index][0] {
		return gamedb.ERRREWARDTIME
	}

	items := cfg.InsistReward[index]
	if len(items) < 3 {
		return gamedb.ERRPARAM
	}
	this.GetBag().Add(user, op, items[1], items[2])
	user.AppletsInfo.List[pb.APPLETSTYPE_BOW].LastGetAwardTime = time.Now().Unix()
	user.Dirty = true
	ack.Goods = op.ToChangeItems()
	return nil
}

//
//  EndResultReq
//  @Description: 通关结算奖励
//  @receiver this
//  @param user
//  @param appletsType 小程序类型,枚举AppletsType
//  @param id 通关id  配置表id
//  @param ack
//  @param op
//  @return error
//
func (this *Applets) EndResultReq(user *objs.User, appletsType, id int, ack *pb.EndResultAck, op *ophelper.OpBagHelperDefault) error {

	if !pb.APPLETSTYPE_MAP[appletsType] {
		return gamedb.ERRPARAM
	}

	if user.AppletsInfo.List[appletsType] == nil {
		user.AppletsInfo.List[appletsType] = &model.AppletsUnit{}
	}
	if user.AppletsInfo.List[appletsType].IsInGame <= 0 {
		return gamedb.ERRPARAM
	}

	err, items := this.buildGetAppletsEndResultReward(user, appletsType, id)
	if err != nil {
		return err
	}

	this.GetBag().AddItems(user, items, op)
	if id > user.AppletsInfo.List[appletsType].Stage {
		user.AppletsInfo.List[appletsType].Stage = id
	}

	user.AppletsInfo.List[appletsType].IsInGame = 0
	user.Dirty = true
	ack.Id = int32(id)
	ack.Energy = int32(user.AppletsInfo.Energy)
	ack.AppletsType = int32(appletsType)
	ack.Goods = op.ToChangeItems()
	return nil
}

//
//  CronAddPhysicalPower
//  @Description: 定时增加体力
//  @receiver this
//  @param user
//
func (this *Applets) CronAddPhysicalPower(user *objs.User) {
	if user == nil {
		return
	}
	if user.AppletsInfo == nil {
		return
	}
	if user.AppletsInfo.Energy >= gamedb.GetConf().XiaoYouXiEnergy {
		return
	}

	if user.AppletsInfo.ResumeTime > time.Now().Unix() {
		return
	}
	if user.AppletsInfo.Energy < gamedb.GetConf().XiaoYouXiEnergy {
		user.AppletsInfo.Energy += gamedb.GetConf().XiaoYouXiEnergyResume[1]
		if user.AppletsInfo.Energy > gamedb.GetConf().XiaoYouXiEnergy {
			user.AppletsInfo.Energy = gamedb.GetConf().XiaoYouXiEnergy
		}
		user.AppletsInfo.ResumeTime = time.Now().Unix() + int64(gamedb.GetConf().XiaoYouXiEnergyResume[0])
		user.Dirty = true
		ntf := &pb.AppletsEnergyNtf{
			Energy:     int32(user.AppletsInfo.Energy),
			ResumeTime: user.AppletsInfo.ResumeTime,
		}
		_ = this.GetUserManager().SendMessage(user, ntf, true)
	}

	return
}

//
//  OnlineAddPhysicalPower
//  @Description: 上线检查体力恢复
//  @receiver this
//  @param user
//
func (this *Applets) OnlineAddPhysicalPower(user *objs.User) {

	if user.AppletsInfo.List != nil && len(user.AppletsInfo.List) > 0 {
		for _, types := range pb.APPLETSTYPE_ARRAY {
			if user.AppletsInfo.List[types] != nil {
				user.AppletsInfo.List[types].IsInGame = 0
			}
		}
		user.Dirty = true
	}

	if user.AppletsInfo.Energy >= gamedb.GetConf().XiaoYouXiEnergy {
		return
	}

	if user.AppletsInfo.ResumeTime > time.Now().Unix() {
		return
	}

	nums := int(time.Now().Unix()-(user.AppletsInfo.ResumeTime-int64(gamedb.GetConf().XiaoYouXiEnergyResume[0]))) / gamedb.GetConf().XiaoYouXiEnergyResume[0]
	addNum := nums * gamedb.GetConf().XiaoYouXiEnergyResume[1]
	logger.Debug("now:%v  ResumeTime:%v  Energy:%v   addNum:%v", time.Now().Unix(), user.AppletsInfo.ResumeTime, user.AppletsInfo.Energy, addNum)
	if addNum < 0 {
		addNum = 0
	}
	user.AppletsInfo.Energy += addNum
	if user.AppletsInfo.Energy > gamedb.GetConf().XiaoYouXiEnergy {
		user.AppletsInfo.Energy = gamedb.GetConf().XiaoYouXiEnergy
	}

	user.AppletsInfo.ResumeTime = time.Now().Unix() + int64(gamedb.GetConf().XiaoYouXiEnergyResume[0])
	user.Dirty = true
	return
}

func (this *Applets) buildGetAppletsEndResultReward(user *objs.User, appletsType, id int) (error, gamedb.ItemInfos) {
	items := gamedb.ItemInfos{}
	switch appletsType {
	case pb.APPLETSTYPE_PURSUIT:
		cfg := gamedb.GetAttackEnemyAttackEnemyCfg(id)
		if cfg == nil || cfg.Reward == nil || len(cfg.Reward) <= 0 {
			return gamedb.ERRPARAM, nil
		}
		items = cfg.Reward
	case pb.APPLETSTYPE_BOW:
		cfg := gamedb.GetArcherMagicCfg(id)
		if cfg == nil || cfg.PassReward == nil || len(cfg.PassReward) <= 0 {
			return gamedb.ERRPARAM, nil
		}
		items = cfg.PassReward

	case pb.APPLETSTYPE_NUMBER:
		cfg := gamedb.GetXiaoyouxiTowerXiaoyouxiTowerCfg(id)
		if cfg == nil || cfg.Reward == nil || len(cfg.Reward) <= 0 {
			return gamedb.ERRPARAM, nil
		}
		state := this.GetCondition().CheckMulti(user, -1, cfg.Condition)
		if !state {
			return gamedb.ERRCONDITION, nil
		}
		items = cfg.Reward

	default:
		return gamedb.ERRPARAM, nil
	}
	return nil, items
}
