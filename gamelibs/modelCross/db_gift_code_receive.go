package modelCross

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/dbmodel"
	"fmt"
	"gopkg.in/gorp.v1"
	"time"
)

type GiftCodeReceiveDb struct {
	CodeId      int       `db:"codeId" orm:"comment(礼包码)"`
	UserId      int       `db:"userId" orm:"comment(用户id)"`
	BatchId     int       `db:"batchId" orm:"comment(批次)"`
	ReceiveTime time.Time `db:"receiveTime" orm:"comment(领取时间)"`
}

func (this *GiftCodeReceiveDb) TableName() string {
	return "gift_code_receive"
}

type GiftCodeReceiveModel struct {
	dbmodel.CommonModel
}

var (
	giftCodeReceiveDb     = &GiftCodeReceiveDb{}
	giftCodeReceiveModel  = &GiftCodeReceiveModel{}
	giftCodeReceiveFields = model.GetAllFieldsAsString(GiftCodeReceiveDb{})
)

func init() {
	dbmodel.Register(model.DB_ACCOUNT, giftCodeReceiveModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(GiftCodeReceiveDb{}, giftCodeReceiveDb.TableName()).SetKeys(false, "codeId")
	})
}

func GetGiftCodeReceiveDbModel() *GiftCodeReceiveModel {
	return giftCodeReceiveModel
}

func (this *GiftCodeReceiveModel) Create(data *GiftCodeReceiveDb) error {
	return this.DbMap().Insert(data)
}

func (this *GiftCodeReceiveModel) LoadGiftCodeReceiveAllData() ([]*GiftCodeReceiveDb, error) {
	sql := fmt.Sprintf("select %s from %s where 1", giftCodeReceiveFields, giftCodeReceiveDb.TableName())
	var data []*GiftCodeReceiveDb
	_, err := this.DbMap().Select(&data, sql)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *GiftCodeReceiveModel) GetGiftCodeReceiveAllDataByCode(codeId int) ([]*GiftCodeReceiveDb, error) {
	sql := fmt.Sprintf(`select %s from %s where codeId=%d`, giftCodeReceiveFields, giftCodeReceiveDb.TableName(), codeId)
	var data []*GiftCodeReceiveDb
	_, err := this.DbMap().Select(&data, sql)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *GiftCodeReceiveModel) GetGiftCodeReceiveAllDataByUserId(userId int) ([]*GiftCodeReceiveDb, error) {
	sql := fmt.Sprintf(`select %s from %s where userId=%d`, giftCodeReceiveFields, giftCodeReceiveDb.TableName(), userId)
	var data []*GiftCodeReceiveDb
	_, err := this.DbMap().Select(&data, sql)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *GiftCodeReceiveModel) GetGiftCodeReceiveAllDataByUserIdAndBatch(userId, batch int) ([]*GiftCodeReceiveDb, error) {
	sql := fmt.Sprintf(`select %s from %s where userId=%d and batchId=%d`, giftCodeReceiveFields, giftCodeReceiveDb.TableName(), userId, batch)
	var data []*GiftCodeReceiveDb
	_, err := this.DbMap().Select(&data, sql)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *GiftCodeReceiveModel) GetUserReceiveTimesByBatch(userId, batch int) (int, error) {
	sql := fmt.Sprintf(`select count(*) from %s where userId=%d and batchId=%d`, giftCodeReceiveDb.TableName(), userId, batch)
	result, err := this.DbMap().SelectInt(sql)
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

func (this *GiftCodeReceiveModel) GetGiftCodeReceiveNumByCodeId(codeId int) (int, error) {
	sql := fmt.Sprintf("select count(*) from %s where codeId=%d", giftCodeReceiveDb.TableName(), codeId)
	result, err := this.DbMap().SelectInt(sql)
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

func (this *GiftCodeReceiveModel) GetGiftCodeReceiveNumByCodeIdAndUserId(codeId, userId int) (int, error) {
	sql := fmt.Sprintf("select count(*) from %s where codeId=%d and userId=%d", giftCodeReceiveDb.TableName(), codeId, userId)
	result, err := this.DbMap().SelectInt(sql)
	if err != nil {
		return 0, err
	}
	return int(result), nil
}
