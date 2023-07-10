package modelCross

import (
	"cqserver/gamelibs/model"
	"fmt"
	"gopkg.in/gorp.v1"

	"cqserver/golibs/dbmodel"
)

type ChallengeData struct {
	Id         int            `db:"id" orm:"pk;auto"`
	Season     string         `db:"season"`
	CrossFsId  int            `db:"crossFsId"`
	Round      int            `db:"round"`
	UserIds    model.IntSlice `db:"userIds"  orm:"type(text)"`
	ExpireTime int64          `db:"expireTime"` //数据删除时间,数据保留5天 (43200秒)
}

type ChallengeDataModel struct {
	dbmodel.CommonModel
}

var (
	challengeDataModel  = &ChallengeDataModel{}
	challengeDataFields = dbmodel.GetAllFieldsAsString(ChallengeData{})
)

func init() {
	dbmodel.Register(model.DB_ACCOUNT, challengeDataModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(ChallengeData{}, "challenge_data").SetKeys(true, "Id")
		//orm.RegisterModelForAlias(model.DB_ACCOUNT, new(ChallengeData))
	})

}

func GetChallengeDataModel() *ChallengeDataModel {
	return challengeDataModel
}

// 获取指定服务器的报名玩家
func (this *ChallengeDataModel) GetRoundUserIdsByCrossFsId(crossFsId, round int, season string) (*ChallengeData, error) {
	var challenge *ChallengeData

	err := this.DbMap().SelectOne(&challenge, fmt.Sprintf("select %s from challenge_data where crossFsId = %v and round = %v and season = '%v'", challengeDataFields, crossFsId, round, season))
	if err != nil {
		return nil, err
	}
	return challenge, nil
}

func (this *ChallengeDataModel) DeleteExpiredItem(ts int64) error {
	sql := fmt.Sprintf("DELETE FROM challenge_data where expireTime > 0 AND expireTime <= %d", ts)
	_, err := this.DbMap().Exec(sql)
	return err
}