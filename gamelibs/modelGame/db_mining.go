package modelGame

import (
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constMining"
	"cqserver/golibs/dbmodel"
	"fmt"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gorp.v1"
	"time"
)

type MiningDb struct {
	dbmodel.DbTable
	Id          int       `db:"id" orm:"pk;auto"`
	UserId      int       `db:"userId" orm:"comment(玩家id)"`
	Miner       int       `db:"miner" orm:"comment(矿工等级)"`
	Ruid        int       `db:"ruid" orm:"comment(抢夺玩家id)"`
	Rtime       time.Time `db:"rtime" orm:"comment(抢夺时间)"`
	Rstatus     int       `db:"rstatus" orm:"comment(抢夺状态)"`
	WorkTime    time.Time `db:"workTime" orm:"comment(挖矿时间)"`
	ExpireTime  time.Time `db:"expireTime" orm:"comment(挖矿结束时间)"`
	ReceiveTime time.Time `db:"receiveTime" orm:"comment(领取奖励时间)"`
	FindTime    time.Time `db:"ftime" orm:"comment(找回奖励时间)"`
	CreatedAt   time.Time `db:"createdAt" orm:"comment(创建时间)"`
	DeletedAt   time.Time `db:"deletedAt" orm:"comment(删除时间)"`
	IsRobot     int       `db:"isRobot" orm:"comment(是否机器人)"`
}

func (this *MiningDb) TableName() string {
	return "mining"
}

type MiningModel struct {
	dbmodel.CommonModel
}

var (
	miningModel  = &MiningModel{}
	miningFields = model.GetAllFieldsAsString(MiningDb{})
	idSeqMining  = &modelCross.IdSeq{Name: "mining"}
	notRobotSql = fmt.Sprintf(" and isRobot=%d ", constMining.MINING_DATA_ROBOT_NO)
)

func init() {
	dbmodel.Register(model.DB_SERVER, miningModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(MiningDb{}, "mining").SetKeys(true, "Id")
		orm.RegisterModelForAlias(model.DB_SERVER, new(MiningDb))
	})
}

func GetMiningModel() *MiningModel {
	return miningModel
}

func (this *MiningModel) Create(mining *MiningDb) error {
	return this.DbMap().Insert(mining)
}

func (this *MiningModel) Update(mining *MiningDb) error {
	_, err := this.DbMap().Update(mining)
	return err
}

func (this *MiningModel) GetMiningById(id int) (*MiningDb, error) {
	var mining MiningDb
	err := this.DbMap().SelectOne(&mining, fmt.Sprintf("select %s from mining where id = ?", miningFields), id)
	if err != nil {
		return nil, err
	}
	return &mining, nil
}

func (this *MiningModel) GetMiningListByUserId(userId int) ([]*MiningDb, error) {
	var mining []*MiningDb
	_, err := this.DbMap().Select(&mining, fmt.Sprintf("select %s from mining where userId=? and deletedAt=0 %s order by createdAt desc",
		miningFields, notRobotSql), userId)
	return mining, err
}

func (this *MiningModel) GetMiningByUserId(userId int) ([]*MiningDb, error) {
	var minings []*MiningDb
	_, err := this.DbMap().Select(&minings, fmt.Sprintf("select %s from mining where userId=%d %s", miningFields, userId, notRobotSql))
	if err != nil {
		return nil, err
	}
	return minings, nil
}

func (this *MiningModel) GetLastWorkTime(userId int) (*MiningDb, error) {
	var mining *MiningDb
	_, err := this.DbMap().Select(&mining, fmt.Sprintf("select %s from mining where userId=? and deletedAt=0 %s order by workTime desc limit 1",
		miningFields, notRobotSql), userId)
	return mining, err
}

func (this *MiningModel) GetMiningAll() ([]*MiningDb, error) {
	var mining []*MiningDb
	_, err := this.DbMap().Select(&mining, fmt.Sprintf("select %s from mining where receiveTime=0 and deletedAt=0", miningFields))
	return mining, err
}

func (this *MiningModel) GetRobListByUserId(userId int) ([]*MiningDb, error) {
	var minings []*MiningDb
	_, err := this.DbMap().Select(&minings, fmt.Sprintf("select %s from mining where userId=%d and rtime>=date(now()) and rtime<DATE_ADD(date(now()),INTERVAL 1 DAY) %s",
		miningFields, userId, notRobotSql))
	return minings, err
}

func (this *MiningModel) GetMiningId() (int, error) {
	return idSeqMining.Next()
}
