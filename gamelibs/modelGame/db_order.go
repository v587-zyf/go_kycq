package modelGame

import (
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constOrder"
	"cqserver/golibs/dbmodel"
	"cqserver/golibs/logger"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"

	"gopkg.in/gorp.v1"
)

type OrderDb struct {
	Id              int       `db:"id" orm:"pk;auto"`
	OpenId          string    `db:"openId"  orm:"comment(玩家账号)"`
	UserId          int       `db:"userId"  orm:"comment(玩家id)"`
	ServerIndex     int       `db:"serverIndex"  orm:"comment(服务器索引)"`
	ServerId        int       `db:"serverId"  orm:"comment(服务器Id)"`
	RechargeId      int       `db:"rechargeId" orm:"comment(服务器Id)"`
	Ingot           int       `db:"ingot" orm:"comment(元宝数)"`
	PayMoney        int       `db:"paymoney" orm:"comment(支付金额)"`
	PlatformOrderNo string    `db:"platformOrderNo" orm:"comment(平台订单号)"`
	OrderNo         string    `db:"orderNo" orm:"comment(游戏订单号)"`
	Synced          int       `db:"synced" orm:"comment(是否同步玩家)"`
	Ip              string    `db:"ip"`
	FinishTime      time.Time `db:"finishedTime" orm:"comment(订单完成时间，平台同步过来时间)"`
	CreateTime      time.Time `db:"createdTime" orm:"comment(订单创建时间)"`
	PayModule       int       `db:"payModule" orm:"comment(订单类型)"`
	PayModuleId     int       `db:"payModuleId" orm:"comment(订单类型Id)"` //如每日礼包,dailyPack表id)
	IsPayToken      int       `db:"isPayToken" orm:"comment(是否用代币抵扣)"`
	OrderDis        string    //订单名称（描述）
}

func (this *OrderDb) TableName() string {
	return "orders"
}

type OrderModel struct {
	dbmodel.CommonModel
}

var (
	orderModel  = &OrderModel{}
	orderFields = model.GetAllFieldsAsString(OrderDb{})
)

func init() {
	dbmodel.Register(model.DB_SERVER, orderModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(OrderDb{}, "orders").SetKeys(true, "Id")
		orm.RegisterModelForAlias(model.DB_SERVER, new(OrderDb))
	})
}

func GetOrderModel() *OrderModel {
	return orderModel
}

func (this *OrderModel) TableName() string {
	return "orders"
}

func (this *OrderModel) Create(order *OrderDb) error {
	err := this.DbMap().Insert(order)
	return err
}

func (this *OrderModel) GetOrderById(userId int) (*OrderDb, error) {
	var order OrderDb
	sqlStr := fmt.Sprintf("select %s from orders where userId = %d  and year(finishedTime)>2000 and synced=0 order by id desc limit 1", orderFields, userId)
	err := this.DbMap().SelectOne(&order, sqlStr)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (this *OrderModel) Update(order *OrderDb) error {
	_, err := this.DbMap().Update(order)
	return err
}

func (this *OrderModel) GetNoRechargeOrder(userId int) ([]*OrderDb, error) {

	var orders []*OrderDb
	sqlStr := fmt.Sprintf("select %s from orders where userId = %d  and year(finishedTime)>2000 and synced= %d", orderFields, userId, constOrder.ORDER_SYNCED_NO)
	_, err := this.DbMap().Select(&orders, sqlStr)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (this *OrderModel) GetOrderByOrderNo(orderNo string) (*OrderDb, error) {
	var order OrderDb
	err := this.DbMap().SelectOne(&order, fmt.Sprintf("select %s from orders where orderNo = ?", orderFields), orderNo)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (this *OrderModel) GetOrderByPlatformOrderNo(orderNo string) (*OrderDb, error) {
	var order OrderDb
	err := this.DbMap().SelectOne(&order, fmt.Sprintf("select %s from orders where platformOrderNo = ?", orderFields), orderNo)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &order, nil
}

func (this *OrderModel) GetUserLastRechargeTime(userId int) time.Time {
	var t struct {
		FinishTime time.Time `db:"finishedTime" orm:"comment(订单完成时间，平台同步过来时间)"`
	}
	err := this.DbMap().SelectOne(&t, "select finishedTime from orders where userId=? and year(finishedTime)>2000 order by id desc limit 1", userId)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("获取玩家最后订单时间错误：%v", err)
	}
	return t.FinishTime
}

func (this *OrderModel) GetUserRechargeTotal(userId int) (int, int) {
	var t []struct {
		IsPayToken int `db:"isPayToken" orm:"comment(是否用代币抵扣)"`
		Total      int `db:"total" orm:"comment(总数)"`
	}
	_, err := this.DbMap().Select(&t, " SELECT isPayToken,sum(paymoney) as total from orders WHERE userId = ? and year(finishedTime)>2000  GROUP BY isPayToken;", userId)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("获取玩家最后订单时间错误：%v", err)
	}
	if err == sql.ErrNoRows {
		return 0, 0
	}
	realTotal, tokenToal := 0, 0
	for _, v := range t {
		if v.IsPayToken == 0 {
			realTotal = v.Total
		} else {
			tokenToal = v.Total
		}
	}
	return realTotal, tokenToal
}
