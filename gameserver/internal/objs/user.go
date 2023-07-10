package objs

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/prop"
	"fmt"
	"sync"
	"time"

	"cqserver/golibs/common"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

type User struct {
	*modelGame.User
	Fabaos      model.Fabaos //弃用字段，因需配合客户端删除，暂保留在这里
	Arena       *model.Arena //弃用字段，因需配合客户端删除，暂保留在这里
	Heros       map[int]*Hero
	AccountInfo *modelCross.Account
	DeviceId    string
	Origin      string
	Ip          string

	OffLineWg     sync.WaitGroup
	conn          nw.Conn
	GateSessionId uint32

	OfflineAwardMark bool

	StageCumulationTime  int       //玩家关卡累计时间
	StageNormalStartTime time.Time //玩家关卡开始时间

	StageExpCumulationTime  int       //玩家关卡累计时间
	StageExpNormalStartTime time.Time //玩家关卡开始时间

	Dirty        bool
	OnlineTime   time.Time //登录时间
	LoginTime    time.Time
	RechargeTime time.Time

	SyncStatus common.Bitmask //向客户端同步的状态

	FightId                  int       //当前所在fightid
	FightStageId             int       //当前stageId
	FightStartTime           time.Time //战斗开始
	updateFightUserHeroState bool
	updateFightUserHeroIndex map[int]bool //更新武将索引

	FitHolyEquipEffects []int //合体圣装套装效果id
	CutTreasureUseEndCd int64 //切割技能Cd
	LastRechargeTime    int64

	QuickCd map[int]int //快捷使用cd

	AwakenRank int //神威榜排名
	JewelAllLv int //宝石等级
	WingCombat int //羽翼总战力
	SortTimes  model.IntKv

	FightLessTimes bool //标记是否已扣除次数

	PetCombat           int                     //战宠战力
	PetAppendageEffects map[int]int             //战宠附体技能效果id,类型
	PetAddSkills        []int32                 //战宠额外技能
	PetAddAttr          map[int]map[int32]int64 //战宠额外属性
	GmProperty          *prop.Prop

	CheckTitleExpire        bool  //是否检查称号到期
	ChallengeApplyTime      int64 //擂台赛上次报名时间
	CheckDaBaoMysteryEnergy bool  //是否检查打宝秘境体力
}

type UserDataChangeEvent struct {
	UserId   int
	Data     interface{}
	Callback func(int, *User, interface{})
}

func NewUser() *User {
	u := &User{
		Heros:                    make(map[int]*Hero),
		QuickCd:                  make(map[int]int),
		updateFightUserHeroIndex: make(map[int]bool),
		LastRechargeTime:         -1,
		Fabaos:                   make(model.Fabaos),
		Arena:                    &model.Arena{},
	}
	return u
}

func (user *User) MarkSyncStatus(status common.Bitmask) {
	user.SyncStatus.AddFlag(status)
}

func (user *User) getTopField(itemId int) *int {
	switch itemId {
	case pb.ITEMID_EXP:
		return &user.Exp
	case pb.ITEMID_LV:
		lv := user.GetMaxHeroLv()
		return &lv
	case pb.ITEMID_VIP_EXP:
		return &user.VipScore
	case pb.ITEMID_VIP_LV:
		return &user.VipLevel
	case pb.ITEMID_INGOT:
		return &user.Ingot
	case pb.ITEMID_CHUANQI_BI:
		return &user.ChuanqiBi
	case pb.ITEMID_GOLD:
		return &user.Gold
	case pb.ITEMID_HONOR:
		return &user.Honour
	case pb.ITEMID_WAR_ORDER_EXP:
		return &user.WarOrder.Exp
	case pb.ITEMID_BINDING_INGOT:
		return &user.BindingIngot
	case pb.ITEMID_GOLD_INGOT:
		return &user.GoldIngot
	default:
		return nil
	}
}

func (user *User) GetTopDataByItemId(itemId int) int {
	f := user.getTopField(itemId)
	return *f
}

func (user *User) AddTopDataByItemId(k, v int) (int, error) {
	f := user.getTopField(k)
	if f == nil {
		return 0, gamedb.ERRPARAM.CloneWithMsg("top data bad k: %d", k)
	}
	*f += v
	return *f, nil
}

func (user *User) SetTopDataByItemId(k, v int) (int, error) {
	f := user.getTopField(k)
	if f == nil {
		return 0, gamedb.ERRPARAM.CloneWithMsg("top data bad k: %d", k)
	}
	if v <= 0 {
		return 0, gamedb.ERRPARAM.CloneWithMsg("addProp k is %d, v:%d should great than 0", k, v)
	}
	*f = v
	return *f, nil
}

func (this *User) GetDisplayIdsArtifact() (int, int, int, int) {
	return 0, 0, 0, 0
}

func (this *User) SetConn(conn nw.Conn) {
	this.conn = conn
}

// 双登 0315
func (this *User) CloseConn(reason string) {
	if this.conn != nil {
		this.conn.SetCloseReason(reason)
		this.conn.Close()
		this.conn.Wait()
	}
}

func (user *User) String() string {
	return fmt.Sprintf("{nickname:%s,openId:%s,VipLevel:%d,serverIndex:%d}", user.NickName, user.OpenId, user.VipLevel, user.ServerIndex)
}

func (user *User) IdName() string {
	return fmt.Sprintf("【%d_%s_%s】", user.Id, user.OpenId, user.NickName)
}

func (user *User) UpdateFightUserHeroIndexFun(heroIndex int) {

	if !user.updateFightUserHeroState {
		user.updateFightUserHeroState = true
	}
	if heroIndex == -1 {
		for k, _ := range user.Heros {
			user.updateFightUserHeroIndex[k] = true
		}
	} else {
		if _, ok := user.Heros[heroIndex]; ok {
			user.updateFightUserHeroIndex[heroIndex] = true
		}
	}
}

func (user *User) GetUpdateFightUserHeroIndex() map[int]bool {
	return user.updateFightUserHeroIndex
}

func (user *User) GetUpdateFightUserHeroState() bool {
	return user.updateFightUserHeroState
}

func (user *User) ResetUpdateFightUserHeroInfo() {
	user.updateFightUserHeroState = false
	user.updateFightUserHeroIndex = make(map[int]bool)
}

func (user *User) GetMaxHeroLv() int {

	maxLv := 0
	for _, v := range user.Heros {
		if v.ExpLvl > maxLv {
			maxLv = v.ExpLvl
		}
	}
	return maxLv
}
