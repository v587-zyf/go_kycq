package modelCross

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/common"
	"cqserver/golibs/dbmodel"
	"fmt"
	"gopkg.in/gorp.v1"
	"time"
)

type GiftCodeDb struct {
	dbmodel.DbTable
	Id           int         `db:"id" orm:"pk;auto"`
	Code         string      `db:"code" orm:"size(50);unique;comment(礼包码)"`
	BatchId      int         `db:"batchId" orm:"default:0";orm:"comment(批次)"`
	BatchName    string      `db:"batchName" orm:"null;size(100);comment(批次名)"`
	BatchNum     int         `db:"batchNum" orm:"default:1;comment(本批次可使用个数)"`
	Reward       model.IntKv `db:"reward" orm:"type(text);comment(奖励信息)"` //itemId,count
	ServerId     int         `db:"serverId" orm:"default:0;comment(可使用服务器id)"`
	Channel      int         `db:"channel" orm:"default:0;comment(可领取渠道)"`
	StartTime    time.Time   `db:"startTime" orm:"comment(开始时间)"`
	EndTime      time.Time   `db:"endTime" orm:"comment(结束时间)"`
	GiftCodeType int         `db:"codeType" orm:"default:1;comment(礼包码类型)"`
}

func (this *GiftCodeDb) TableName() string {
	return "gift_code"
}

type GiftCodeModel struct {
	dbmodel.CommonModel
}

var (
	giftCodeDb     = &GiftCodeDb{}
	giftCodeModel  = &GiftCodeModel{}
	giftCodeFields = model.GetAllFieldsAsString(GiftCodeDb{})
)

func init() {
	dbmodel.Register(model.DB_ACCOUNT, giftCodeModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(GiftCodeDb{}, giftCodeDb.TableName()).SetKeys(true, "Id")
	})
}

func GetGiftCodeDbModel() *GiftCodeModel {
	return giftCodeModel
}

func (this *GiftCodeModel) Create(data *GiftCodeDb) error {
	return this.DbMap().Insert(data)
}

func (this *GiftCodeModel) Update(data *GiftCodeDb) error {
	_, err := this.DbMap().Update(data)
	return err
}

//加载所有礼包码数据
func (this *GiftCodeModel) LoadGiftCodeAllData() ([]*GiftCodeDb, error) {
	sql := fmt.Sprintf("select %s from %s where 1", giftCodeFields, giftCodeDb.TableName())
	var data []*GiftCodeDb
	_, err := this.DbMap().Select(&data, sql)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//加载所有 code
func (this *GiftCodeModel) LoadGiftCodeAllCode() ([]*GiftCodeDb, error) {
	sql := fmt.Sprintf("select code from %s where 1", giftCodeDb.TableName())
	var data []*GiftCodeDb
	_, err := this.DbMap().Select(&data, sql)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//根据 code 查询所有礼包信息（模糊查询）
func (this *GiftCodeModel) GetGiftCodeByCode(code string) (*GiftCodeDb, error) {
	sql := fmt.Sprintf(`select %s from %s where code = "%s"`, giftCodeFields, giftCodeDb.TableName(), code)
	//sql := fmt.Sprintf(`select %s from %s where code = binary "%s"`, giftCodeFields, giftCodeDb.TableName(), code)	//严格模式，区分大小写
	var data GiftCodeDb
	err := this.DbMap().SelectOne(&data, sql)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

//查询有效期内的礼包码
func (this *GiftCodeModel) GetGiftCodeInTime(startT, endT string) ([]*GiftCodeDb, error) {
	var data []*GiftCodeDb
	startTime, _ := common.GetTime(startT)
	endTime, _ := common.GetTime(endT)
	err := this.DbMap().SelectOne(&data, fmt.Sprintf(`select %s from %s where startTime >= %s and endTime <= %s`, giftCodeFields, giftCodeDb.TableName(),
		startTime, endTime))
	if err != nil {
		return nil, err
	}
	return data, nil
}
