package fightModule

import (
	"cqserver/fightserver/conf"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"database/sql"
	"runtime/debug"
	"sync"
	"time"
)

var ServerId int

type FightIdAllocator interface {
	GenerateFightId() uint32
	GenerateMainCityId(kingdom int, lineNo int) uint32
}

type FightManager struct {
	util.DefaultModule
	idAllocator         FightIdAllocator
	fights              map[uint32]base.Fight
	residentFight       map[int]uint32 //常驻战斗记录 key为战斗类型加副本条件拼接，value为战斗Id
	shabakeCrossFightId uint32
	shabakeFightNewId   uint32
	fightMu             sync.RWMutex
	servers             map[int]string
	serverMu            sync.RWMutex
	magicTower          *MagicTower
	bossFamilyBossInfo  map[int]int //bossfamily boss数量信息
}

var fightManager *FightManager

func GetFightMgr() *FightManager {
	return fightManager
}

/**
 *  @Description:
 *  @param gs
 *  @param gate
 *  @return *FightManager
 */
func NewFightManager() *FightManager {

	fightManager = &FightManager{
		idAllocator:        base.NewIdAllocator(1),
		fights:             make(map[uint32]base.Fight),
		residentFight:      make(map[int]uint32),
		magicTower:         NewMagicTower(),
		bossFamilyBossInfo: make(map[int]int),
	}
	return fightManager
}

func (this *FightManager) Init() error {

	//常驻战斗初始化
	allStageCfg := gamedb.GetAllStageCfg()
	if conf.Conf.ServerType == 1 {
		//初始化常驻战斗
		for _, v := range allStageCfg {
			//初始野外boss
			var fight base.Fight
			var err error
			if v.Type == constFight.FIGHT_TYPE_FIELDBOSS {

				fight, err = NewFieldBossFight(v.Id)

			} else if v.Type == constFight.FIGHT_TYPE_WORLDBOSS {

				fight, err = NewWorldBossFight(v.Id)

			} else if v.Type == constFight.FIGHT_TYPE_MAIN_CITY {

				fight, err = NewMainCity(v.Id)

			} else if v.Type == constFight.FIGHT_TYPE_PUBLIC_DABAO {

				fight, err = NewPublicDabaoFight(v.Id)

				if err == nil {
					bossFamilyConf := gamedb.GetBossFamilyBossFamilyCfg(v.Id)
					if bossFamilyConf != nil {
						this.bossFamilyBossInfo[v.Id] = fight.GetBossAliveNum()
					}
				}

			} else if v.Type == constFight.FIGHT_TYPE_DARKPALACE {

				fight, err = NewDarkPalace(v.Id)
			} else if v.Type == constFight.FIGHT_TYPE_DARKPALACE_BOSS {

				fight, err = NewDarkPalaceBossFight(v.Id)
			} else if v.Type == constFight.FIGHT_TYPE_ANCIENT_BOSS {

				fight, err = NewAncientBossFight(v.Id)
			} else if v.Type == constFight.FIGHT_TYPE_HELL {

				fight, err = NewHellFight(v.Id)
			} else if v.Type == constFight.FIGHT_TYPE_HELL_BOSS {

				fight, err = NewHellBossFight(v.Id)
			} else if v.Type == constFight.FIGHT_TYPE_CROSS_WORLD_LEADER {

				fight, err = NewWorldLeaderFight(v.Id)
			} else {
				continue
			}
			if err != nil {
				return err
			}
			fightId := this.idAllocator.GenerateFightId()
			fight.SetId(fightId)
			this.fights[fightId] = fight
			this.residentFight[v.Id] = fightId
		}
	} else if conf.Conf.ServerType == 2 {
		//初始化常驻战斗
		for _, v := range allStageCfg {
			var fight base.Fight
			var err error
			if v.Type == constFight.FIGHT_TYPE_CROSS_WORLD_LEADER {
				//初始世界首领boss
				fight, err = NewWorldLeaderFight(v.Id)

			} else if v.Type == constFight.FIGHT_TYPE_HELL {
				//初始化炼狱首领
				fight, err = NewHellFight(v.Id)
			} else if v.Type == constFight.FIGHT_TYPE_HELL_BOSS {
				//初始化炼狱首领
				fight, err = NewHellBossFight(v.Id)
			} else {
				continue
			}
			if err != nil {
				return err
			}
			fightId := this.idAllocator.GenerateFightId()
			fight.SetId(fightId)
			this.fights[fightId] = fight
			this.residentFight[v.Id] = fightId
		}
		err := this.InitGameServers()
		if err != nil {
			return err
		}
	}
	this.Updateframe()
	return nil
}

func (this *FightManager) InitGameServers() error {
	if conf.Conf.ServerType != 2 {
		return nil
	}
	this.serverMu.Lock()
	defer this.serverMu.Unlock()
	servers, err := modelCross.GetServerInfoModel().GetAllServerIdsByCrossFsIds(conf.Conf.ServerId)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	this.servers = make(map[int]string)
	for _, v := range servers {
		this.servers[v.ServerId] = v.Name
	}
	logger.Info("获取服务器名字信息：%v", this.servers)
	return nil
}

func (this *FightManager) GetServerName(serverId int) string {
	this.serverMu.Lock()
	defer this.serverMu.Unlock()
	return this.servers[serverId]
}

func (this *FightManager) GetFight(fightId uint32) base.Fight {
	this.fightMu.RLock()
	fight := this.fights[fightId]
	this.fightMu.RUnlock()
	return fight
}

func (this *FightManager) GetFightByStageId(stageId int,ext int) uint32 {

	stageConf := gamedb.GetStageStageCfg(stageId)
	if stageConf.Type == constFight.FIGHT_TYPE_CROSS_SHABAKE {
		return this.GetShabakeFight()
	} else if stageConf.Type == constFight.FIGHT_TYPE_MAGIC_TOWER {
		return this.magicTower.GetFightId(stageId)
	} else if stageConf.Type == constFight.FIGHT_TYPE_SHABAKE_NEW {
		return this.GetShabakeFightNew()
	}

	this.fightMu.RLock()
	defer this.fightMu.RUnlock()
	if fightId,ok := this.residentFight[stageId];ok{
		return fightId
	}
	for fightId,f := range this.fights{
		if f.GetStageConf().Id == stageId && f.GetFightExtMark() == ext{
			return fightId
		}
	}
	return 0
}

func (this *FightManager) RemoveFight(id uint32) {
	this.fightMu.Lock()
	if id == this.shabakeCrossFightId {
		this.shabakeCrossFightId = 0
	} else if id == this.shabakeFightNewId {
		this.shabakeFightNewId = 0
	}
	logger.Debug("删除战斗：%v", id)
	delete(this.fights, id)
	this.fightMu.Unlock()
}

/**
 *  @Description: 慎用方法  传入的方法中，不能有调用fightmanager的方法，会死锁
 *  @param f
 *  @return bool
 */
func (this *FightManager) Range(f func(actor base.Fight) bool) bool {
	this.fightMu.RLock()
	for _, fight := range this.fights {
		if f(fight) { // TODO: 检查回调中是否会调用RemoveFight，这样会造成死锁
			this.fightMu.RUnlock()
			return true
		}
	}
	this.fightMu.RUnlock()
	return false
}

func (this *FightManager) GsServerConnected() {

	if conf.Conf.ServerType == 2 {
		return
	}
	ntf := this.GetResidentFightInfo()
	net.GetGsConn().SendMessage(ntf)
}

func (this *FightManager) GetResidentFightInfo() *pbserver.FsResidentFightNtf {
	this.fightMu.Lock()
	defer this.fightMu.Unlock()
	ntf := &pbserver.FsResidentFightNtf{
		ResidentFights:     make(map[int32]uint32),
		FieldBossFightInfo: make(map[int32]*pbserver.FsFieldBossInfoNtf),
	}
	//将常驻战斗Id发送给game
	for k, v := range this.residentFight {
		stageId := int32(k)
		ntf.ResidentFights[stageId] = v
		f := this.fights[v]
		switch f.GetStageConf().Type {
		case constFight.FIGHT_TYPE_FIELDBOSS:
			fight := f.(*FieldBossFight)
			fightInfo := &pbserver.FsFieldBossInfoNtf{
				StageId:    stageId,
				Hp:         float32(fight.GetBossHpPoint()),
				ReliveTime: fight.GetReliveTime(),
			}
			ntf.FieldBossFightInfo[stageId] = fightInfo
		case constFight.FIGHT_TYPE_DARKPALACE_BOSS:
			fight := f.(*DarkPalaceBossFight)
			fightInfo := &pbserver.FsFieldBossInfoNtf{
				StageId:    int32(k),
				Hp:         float32(fight.GetBossHpPoint()),
				ReliveTime: fight.GetReliveTime(),
			}
			ntf.FieldBossFightInfo[int32(k)] = fightInfo
		case constFight.FIGHT_TYPE_ANCIENT_BOSS:
			fight := f.(*AncientBossFight)
			fightInfo := &pbserver.FsFieldBossInfoNtf{
				StageId:    int32(k),
				Hp:         float32(fight.GetBossHpPoint()),
				ReliveTime: fight.GetReliveTime(),
				UserCount:  int32(len(fight.GetPlayerUserids())),
			}
			ntf.FieldBossFightInfo[int32(k)] = fightInfo
		case constFight.FIGHT_TYPE_HELL_BOSS:
			fight := f.(*HellBossFight)
			fightInfo := &pbserver.FsFieldBossInfoNtf{
				StageId:    int32(k),
				Hp:         float32(fight.GetBossHpPoint()),
				ReliveTime: fight.GetReliveTime(),
				UserCount:  int32(len(fight.GetPlayerUserids())),
			}
			ntf.FieldBossFightInfo[int32(k)] = fightInfo
		}
	}
	return ntf
}

/**
 *  @Description:创建战斗
 *  @param fighType 战斗类型
 *  @param mapId 地图Id
 *  @param cpData 战斗副本专属数据
 *  @return int
 *  @return error
 */
func (this *FightManager) CreateFight(stageId int, cpData []byte) (uint32, error) {

	stageConf := gamedb.GetStageStageCfg(stageId)
	if stageConf == nil {
		logger.Error("创建stage:%v的战斗，战斗不存在", stageId)
		return 0, gamedb.ERRSETTINGNOTFOUND
	}

	var fight base.Fight
	var err error
	switch stageConf.Type {
	case constFight.FIGHT_TYPE_STAGE:
		fight, err = NewHangUpFight(stageId)
	case constFight.FIGHT_TYPE_STAGE_BOSS:
		fight, err = NewHangUpBossFight(stageId)
	case constFight.FIGHT_TYPE_PERSON_BOSS:
		fight, err = NewPersonBossFight(stageId)
	case constFight.FIGHT_TYPE_TOWERBOSS:
		fight, err = NewTowerFight(stageId)
	case constFight.FIGHT_TYPE_MATERIAL:
		fight, err = NewMaterialFight(stageId)
	case constFight.FIGHT_TYPE_VIPBOSS:
		fight, err = NewVipBossFight(stageId)
	case constFight.FIGHT_TYPE_EXPBOSS:
		fight, err = NewExpFight(stageId)
	case constFight.FIGHT_TYPE_ARENA:
		fight, err = NewArenaFight(stageId, cpData)
	case constFight.FIGHT_TYPE_FIELD:
		fight, err = NewFieldFight(stageId, cpData)
	case constFight.FIGHT_TYPE_MINING:
		fight, err = NewMiningFight(stageId, cpData)
	case constFight.FIGHT_TYPE_GUILD_BONFIRE:
		fight, err = NewGuildBonfireFight(stageId,common.BytesToInt(cpData))
	case constFight.FIGHT_TYPE_PAODIAN:
		fight, err = NewPaodianFight(stageId)
	case constFight.FIGHT_TYPE_SHABAKE:
		fight, err = NewShabakeFight(stageId)
	case constFight.FIGHT_TYPE_GUARDPILLAR:
		fight, err = NewGuardPillarFight(stageId,common.BytesToInt(cpData))
	case constFight.FIGHT_TYPE_MAGIC_TOWER:
		fight, err = NewMagicTowerFight(stageId)
	case constFight.FIGHT_TYPE_DABAO:
		fight, err = NewDaBaoFight(stageId)
	case constFight.FIGHT_TYPE_PUBLIC_DABAO_SINGLE:
		fight, err = NewPublicDabaoFight(stageId)
	default:
		err = gamedb.ERRFIGHTTYPE

	}
	if err != nil {
		return 0, err
	}
	fightId := this.idAllocator.GenerateFightId()
	fight.SetId(fightId)
	this.fightMu.Lock()
	this.fights[fightId] = fight
	this.fightMu.Unlock()
	return fightId, nil
}

func (this *FightManager) SendMessageToGs() {

}

func (this *FightManager) PaodianUserEnter() {

	var f base.Fight
	fightUser := make(map[int32]int32)
	this.fightMu.Lock()
	for _, v := range this.fights {
		if v.GetStageConf().Type == constFight.FIGHT_TYPE_PAODIAN {
			fightUser[int32(v.GetStageConf().Id)] = int32(v.GetPlayerNum())
			paodianReardConf := gamedb.GetPaodianConfByStageId(v.GetStageConf().Id)
			if paodianReardConf.Times == 1 {
				f = v
			}
		}
	}
	this.fightMu.Unlock()
	if f != nil {
		ntf := &pb.PaoDianUserNumNtf{
			UserNums: fightUser,
		}
		f.GetScene().NotifyAll(ntf)
	}
}

func (this *FightManager) GetShabakeFight() uint32 {

	this.fightMu.Lock()
	defer func() {
		this.fightMu.Unlock()
	}()

	if this.shabakeCrossFightId <= 0 {

		fight, err := NewShabakeCrossFight(constFight.FIGHT_TYPE_SHABAKE_CROSS_STAGE)
		if err != nil {
			logger.Error("创建跨服沙巴克异常：%v", err)
			return 0
		}
		fightId := this.idAllocator.GenerateFightId()
		fight.SetId(fightId)
		this.fights[fightId] = fight
		this.shabakeCrossFightId = fightId
	}
	return this.shabakeCrossFightId
}
func (this *FightManager) GetShabakeFightNew() uint32 {

	this.fightMu.Lock()
	defer func() {
		this.fightMu.Unlock()
	}()

	if this.shabakeFightNewId <= 0 {

		fight, err := NewShabakeFightNew(constFight.FIGHT_TYPE_SHABAKE_NEW_STAGE)
		if err != nil {
			logger.Error("创建新沙巴克异常：%v", err)
			return 0
		}
		fightId := this.idAllocator.GenerateFightId()
		fight.SetId(fightId)
		this.fights[fightId] = fight
		this.shabakeFightNewId = fightId
	}
	return this.shabakeFightNewId
}

func (this *FightManager) Updateframe() {
	go func() {
		var ticker = time.NewTicker(100 * time.Millisecond)
		defer func() {
			ticker.Stop()
			if r := recover(); r != nil {
				stackBytes := debug.Stack()
				logger.Error("panic when DoLoop:%v,%s", r, stackBytes)
			}
		}()
		for {
			select {
			case <-ticker.C:
				this.magicTower.UpdateFrame()
			}
		}
	}()
}

func (this *FightManager) GetBossFamilyInfo(bossFamilyType int) map[int]int {

	this.fightMu.Lock()
	defer func() {
		this.fightMu.Unlock()
	}()

	info := make(map[int]int)
	for k, v := range this.bossFamilyBossInfo {
		conf := gamedb.GetBossFamilyBossFamilyCfg(k)
		if conf.Function == bossFamilyType {
			info[k] = v
		}
	}
	return info
}

func (this *FightManager) UpdateBossFamilInfo(fight base.Fight) {

	this.fightMu.Lock()
	defer func() {
		this.fightMu.Unlock()
	}()
	stageid := fight.GetStageConf().Id
	bossFamilyConf := gamedb.GetBossFamilyBossFamilyCfg(stageid)
	if bossFamilyConf != nil {
		this.bossFamilyBossInfo[stageid] = fight.GetBossAliveNum()
	}
}
