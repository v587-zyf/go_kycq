package warOrder

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
)

func NewWarOrderManager(module managersI.IModule) *WarOrderManager {
	return &WarOrderManager{IModule: module}
}

type WarOrderManager struct {
	util.DefaultModule
	managersI.IModule

	WriteMu sync.Mutex
}

func (this *WarOrderManager) Online(user *objs.User) {
	this.ResetWarOrder(user, false)
}

func (this *WarOrderManager) ResetWarOrder(user *objs.User, reset bool) bool {
	warOrderCycles := gamedb.GetWarOrderCycle()
	nowTime := time.Now().Unix()
	var date time.Time
	var addDay int
	var season int
	for _, cfg := range warOrderCycles {
		y, _ := strconv.Atoi(cfg.StartTime[0])
		m, d := cfg.StartTime[1], cfg.StartTime[2]
		dateStr := fmt.Sprintf("%d-%s-%s 00:00:00", y, m, d)
		t, _ := time.ParseInLocation("2006-01-02 15:04:05", dateStr, time.Local)
		if nowTime < t.Unix() || t.Unix() < date.Unix() {
			continue
		}
		date = t
		addDay = cfg.Duration
		season = cfg.Id
	}
	startTime := int(date.Unix())
	resetFlag := false
	if user.WarOrder.StartTime != startTime {
		user.WarOrder = &model.WarOrder{
			Lv:        1,
			Season:    season,
			StartTime: startTime,
			EndTIme:   int(date.AddDate(0, 0, addDay).Unix()),
			Exchange:  make(model.IntKv),
			Task:      make(map[int]*model.WarOrderTask),
			WeekTask:  make(map[int]map[int]*model.WarOrderTask),
			Reward:    make(map[int]*model.WarOrderReward),
		}
		for _, v := range pb.WARORDERCONDITION_ARRAY {
			rmodel.WarOrder.DelTask(v, user.Id)
		}
		if reset {
			this.GetUserManager().SendMessage(user, &pb.WarOrderResetNtf{WarOrder: builder.BuildWarOrder(user)}, true)
		}
		resetFlag = true
		this.WriteWarOrderTask(user, pb.WARORDERCONDITION_WEEK_LOGIN_DAY, []int{1})
	} else {
		//由于机制问题,上线先记录签到天数 [放到下面会导致不记录登陆天数!!!]
		this.WriteWarOrderTask(user, pb.WARORDERCONDITION_WEEK_LOGIN_DAY, []int{1})
		for _, v := range pb.WARORDERCONDITION_ARRAY {
			addData := []int{0, 0}
			n := rmodel.WarOrder.GetTask(v, user.Id)
			if n != 0 {
				addData = []int{n}
				rmodel.WarOrder.DelTask(v, user.Id)
			}
			switch v {
			case pb.WARORDERCONDITION_SHABAKE_NUM: fallthrough
			case pb.WARORDERCONDITION_PAODIAN_NUM: fallthrough
			case pb.WARORDERCONDITION_GUILDBONFIRE_NUM:
				if n != 0 {
					this.WriteWarOrderTask(user, v, addData)
				}
			default:
				this.WriteWarOrderTask(user, v, addData)
			}
		}
	}
	return resetFlag
}

func (this *WarOrderManager) GetNowWeek(userWarOrder *model.WarOrder) int {
	startTime := time.Unix(int64(userWarOrder.StartTime), 0)
	day := common.GetTheDaysReduceHour(startTime, 5)
	return int(math.Ceil(float64(day) / float64(7)))
}

/**
 *  @Description: 豪华战令，校验是否已开通，支付金额
 *  @param user
 *  @param payNum	支付金额
 *  @return error
 */
func (this *WarOrderManager) WarOrderCheckBuyLuxury(user *objs.User, payNum int) error {
	if user.WarOrder.IsLuxury {
		return gamedb.ERRREPEATBUY
	}
	payMoney := gamedb.GetConf().WarOrderLuxury
	if payMoney != payNum {
		return gamedb.ERRBUYNUM
	}
	return nil
}

/**
 *  @Description: 豪华战令购买后续操作
 *  @param user
 */
func (this *WarOrderManager) WarOrderBuyLuxuryOperation(user *objs.User) {
	if user.WarOrder.IsLuxury {
		return
	}
	user.WarOrder.IsLuxury = true
	user.Dirty = true

	this.GetUserManager().SendMessage(user, &pb.WarOrderBuyLuxuryAck{IsLuxury: true}, true)
	this.GetAnnouncement().SendSystemChat(user, pb.SCROLINGTYPE_GOU_MAI_ZHAN_LIN, -1, -1)
}

/**
 *  @Description: 战令经验，校验支付金额
 *  @param user
 *  @param payNum	支付金额
 *  @return error
 */
func (this *WarOrderManager) WarOrderCheckBuyExp(user *objs.User, payNum int) error {
	payConf := gamedb.GetConf().WarOrderExpBuy
	payMoney := payConf[0]
	if payMoney != payNum {
		return gamedb.ERRBUYNUM
	}
	return nil
}

/**
 *  @Description: 战令经验购买后续操作
 *  @param user
 */
func (this *WarOrderManager) WarOrderBuyExpOperation(user *objs.User) {
	payConf := gamedb.GetConf().WarOrderExpBuy
	userWarOrder := user.WarOrder
	userWarOrder.Exp += payConf[1]
	lv, exp := this.GetWarOrder().AutoUpLv(user, userWarOrder, true)
	this.GetUserManager().SendMessage(user, &pb.WarOrderBuyExpAck{Lv: int32(lv), Exp: int32(exp)}, true)
}

/**
 *  @Description: 领取战令等级奖励
 *  @param user
 *  @param lv	等级
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *WarOrderManager) LvReward(user *objs.User, lv int, op *ophelper.OpBagHelperDefault, ack *pb.WarOrderLvRewardAck) error {
	if lv < 1 {
		return gamedb.ERRPARAM
	}
	userWarOrder := user.WarOrder
	if userWarOrder.Lv < lv {
		return gamedb.ERRWARORDERLVNOTENOUGH
	}
	reward, ok := userWarOrder.Reward[lv]
	if !ok {
		userWarOrder.Reward[lv] = &model.WarOrderReward{}
		reward = userWarOrder.Reward[lv]
	}
	if reward.Elite && reward.Luxury {
		return gamedb.ERRREPEATRECEIVE
	}

	conf := gamedb.GetWarOrderLevelWarOrderLevelCfg(gamedb.GetRealId(userWarOrder.Season, lv))
	if conf == nil {
		return gamedb.ERRPARAM
	}
	if userWarOrder.Exp < conf.WarOrderExp.Count {
		return gamedb.ERRWARORDERLVNOTENOUGH
	}
	if !reward.Elite {
		this.GetBag().AddItems(user, conf.Reward1, op)
		reward.Elite = true
	}
	if userWarOrder.IsLuxury && !reward.Luxury {
		this.GetBag().AddItems(user, conf.Reward2, op)
		reward.Luxury = true
	}
	user.Dirty = true

	ack.Lv = int32(lv)
	ack.Goods = op.ToChangeItems()
	return nil
}

/**
 *  @Description: 战令兑换
 *  @param user
 *  @param exchangeId	兑换id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *WarOrderManager) Exchange(user *objs.User, exchangeId, num int, op *ophelper.OpBagHelperDefault, ack *pb.WarOrderExchangeAck) error {
	if exchangeId < 1 || num < 1 {
		return gamedb.ERRPARAM
	}
	userWarOrder := user.WarOrder
	maxLv := gamedb.GetMaxValById(userWarOrder.Season, constMax.MAX_WARORDER_LEVEL)
	if userWarOrder.Lv < maxLv {
		return gamedb.ERRWARORDERLVNOTENOUGH
	}

	exchangeCfg := gamedb.GetWarOrderExchangeWarOrderExchangeCfg(gamedb.GetRealId(userWarOrder.Season, exchangeId))
	if exchangeCfg == nil {
		return gamedb.ERRPARAM
	}
	if exchangeCfg.Number != 0 {
		buyNum, ok := userWarOrder.Exchange[exchangeId]
		if !ok {
			buyNum = num
		} else {
			buyNum += num
		}
		if buyNum > exchangeCfg.Number {
			return gamedb.ERREXCHANGEENOUGH
		}
	}
	hasExp := userWarOrder.Exp - gamedb.GetWarOrderLevelWarOrderLevelCfg(gamedb.GetRealId(userWarOrder.Season, maxLv)).WarOrderExp.Count
	needExp := exchangeCfg.WarOrderExp.Count * num
	if hasExp < needExp {
		return gamedb.ERRWARORDEREXPNOTENOUGH
	}
	this.GetBag().Add(user, op, exchangeCfg.Item.ItemId, exchangeCfg.Item.Count*num)

	userWarOrder.Exp -= needExp
	userWarOrder.Exchange[exchangeId] += num
	user.Dirty = true

	ack.ExchangeId = int32(exchangeId)
	ack.Exp = int32(userWarOrder.Exp)
	ack.Goods = op.ToChangeItems()
	ack.Num = int32(userWarOrder.Exchange[exchangeId])
	return nil
}

/**
 *  @Description: 添加战令经验
 *  @param user
 *  @param op
 *  @param exp	经验值
 *  @return error
 */
func (this *WarOrderManager) AddExp(user *objs.User, op *ophelper.OpBagHelperDefault, exp int) error {
	userWarOrder := user.WarOrder
	userWarOrder.Exp += exp
	this.AutoUpLv(user, userWarOrder, true)
	op.OnGoodsChange(builder.BuildTopDataChange(pb.ITEMID_WAR_ORDER_EXP, exp, user.WarOrder.Exp), exp)
	return nil
}

func (this *WarOrderManager) AutoUpLv(user *objs.User, userWarOrder *model.WarOrder, isSend bool) (int, int) {
	nowLv, nowExp, season := userWarOrder.Lv, userWarOrder.Exp, userWarOrder.Season
	for i := nowLv; i < constConstant.COMPUTE_TEN_THOUSAND; i++ {
		if gamedb.GetWarOrderLevelWarOrderLevelCfg(gamedb.GetRealId(season, i+1)) == nil {
			break
		}
		lvCfg := gamedb.GetWarOrderLevelWarOrderLevelCfg(gamedb.GetRealId(season, i))
		if nowExp < lvCfg.WarOrderExp.Count {
			break
		}
		//nowExp -= lvCfg.WarOrderExp.Count
		nowLv++
	}
	userWarOrder.Lv = nowLv
	userWarOrder.Exp = nowExp
	user.Dirty = true

	if isSend {
		this.GetUserManager().SendMessage(user, &pb.WarOrderLvNtf{
			Lv:  int32(nowLv),
			Exp: int32(nowExp),
		}, false)
	}
	return nowLv, nowExp
}
