package modelCross

import (
	"cqserver/gamelibs/model"
	"fmt"
	"gopkg.in/gorp.v1"

	"cqserver/golibs/dbmodel"
)

type Challenge struct {
	Id            int    `db:"id" orm:"pk;auto"`
	Season        string `db:"season"` //第几期
	UserId        int    `db:"userId"`
	OpenId        string `db:"openId"`
	NickName      string `db:"nickName"`
	Avatar        string `db:"avatar"`
	CrossFsId     int    `db:"crossFsId"`
	ServerId      int    `db:"serverId"`
	Combat        int64  `db:"combat"`
	IsLose        int    `db:"isLose"`    //是否输了
	LoseRound     int    `db:"loseRound"` //第几轮输了
	WinUserId     int    `db:"winUserId"` //击败自己的玩家id
	FightUserInfo string `db:"fightUserInfo"`
	Round         int    `db:"round"`      //第几轮
	ExpireTime    int64  `db:"expireTime"` //数据删除时间，物品售出或者流拍后数据保留 5天 (43200秒)
	GuildName     string `db:"guildName"`
}

type ChallengeModel struct {
	dbmodel.CommonModel
}

var (
	challengeModel  = &ChallengeModel{}
	challengeFields = dbmodel.GetAllFieldsAsString(Challenge{})
)

func init() {
	dbmodel.Register(model.DB_ACCOUNT, challengeModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(Challenge{}, "challenge").SetKeys(true, "Id")
		//orm.RegisterModelForAlias(model.DB_ACCOUNT, new(Challenge))
	})

}

func GetChallengeModel() *ChallengeModel {
	return challengeModel
}

// 获取指定服务器的报名玩家
func (this *ChallengeModel) GetAllServerApplyUsersList(crossFsId int, season string) ([]*Challenge, error) {
	var challenge []*Challenge

	_, err := this.DbMap().Select(&challenge, fmt.Sprintf("select %s from challenge where  crossFsId = %v and season = '%v' order by combat desc limit 256", challengeFields, crossFsId, season))
	if err != nil {
		return nil, err
	}
	return challenge, nil
}

func (this *ChallengeModel) GetAllServerApplyUserInfo(crossFsId, userId int, season string) (*Challenge, error) {
	var challenge *Challenge

	err := this.DbMap().SelectOne(&challenge, fmt.Sprintf("select %s from challenge where crossFsId = %v and userId = %v and season = '%v'", challengeFields, crossFsId, userId, season))
	if err != nil {
		return nil, err
	}
	return challenge, nil
}

//获取指定服申请战力前2的玩家
func (this *ChallengeModel) GetServerApplyUsersListByCombat(limit, crossFsId int, season string) ([]*Challenge, error) {
	var challenge []*Challenge

	_, err := this.DbMap().Select(&challenge, fmt.Sprintf("select %s from challenge where  season = '%v' and crossFsId = %v  and isLose = 0 order by combat desc limit %v ", challengeFields, season, crossFsId, limit))
	if err != nil {
		return nil, err
	}
	return challenge, nil
}

func (this *ChallengeModel) DeleteExpiredItem(ts int64) error {
	sql := fmt.Sprintf("DELETE FROM challenge where expireTime > 0 AND expireTime <= %d", ts)
	_, err := this.DbMap().Exec(sql)
	return err
}

// 获取指定服务器的报名玩家
func (this *ChallengeModel) GetChallengeNowRound(season string, crossFsId int) (int, error) {

	sql := fmt.Sprintf("select MAX(round) from challenge where season = '%v' and crossFsId = %v ;", season, crossFsId)
	var maxId int
	err := this.DbMap().SelectOne(&maxId, sql)
	if err != nil {
		return -1, err
	}
	return maxId, nil

}

// 获取指定服务器的报名玩家包含机器人
func (this *ChallengeModel) GetAllServerApplyUsersListByRound(season string, crossFsId int) ([]*Challenge, error) {
	var challenge []*Challenge

	_, err := this.DbMap().Select(&challenge, fmt.Sprintf("select %s from challenge where  season = '%v' and crossFsId = %v  order by combat desc", challengeFields, season, crossFsId))
	if err != nil {
		return nil, err
	}
	return challenge, nil
}

// 获取指定服务器的报名玩家不包含机器人
func (this *ChallengeModel) GetAllServerApplyUsersListByRound1(season string, crossFsId int) ([]*Challenge, error) {
	var challenge []*Challenge

	_, err := this.DbMap().Select(&challenge, fmt.Sprintf("select %s from challenge where  season = '%v' and crossFsId = %v and userId > 0  order by combat desc", challengeFields, season, crossFsId))
	if err != nil {
		return nil, err
	}
	return challenge, nil
}

//// 获取指定服务器的报名玩家
//func (this *ChallengeModel) GetAllServerApplyUsersListByRound(serverId, round int) ([]*Challenge, error) {
//	var challenge []*Challenge
//
//	_, err := this.DbMap().Select(&challenge, fmt.Sprintf("select %s from challenge where serverId = %v and userId > 0", challengeFields, serverId))
//	if err != nil {
//		return nil, err
//	}
//	return challenge, nil
//}

// 获取指定跨服组的报名玩家
func (this *ChallengeModel) GetAllCrossFsIdApplyUsersListByRound(crossFsId, round int, season string) ([]*Challenge, error) {
	var challenge []*Challenge

	_, err := this.DbMap().Select(&challenge, fmt.Sprintf("select %s from challenge where crossFsId = %v and round = %v and season = '%v'", challengeFields, crossFsId, round, season))
	if err != nil {
		return nil, err
	}
	return challenge, nil
}

// 获取指定服务器的报名玩家
func (this *ChallengeModel) GetAllServerApplyUsersListWinUsersByRound(crossFsId, round int, season string) ([]*Challenge, error) {
	var challenge []*Challenge

	_, err := this.DbMap().Select(&challenge, fmt.Sprintf("select %s from challenge where crossFsId = %v and round = %v and isLose = 0 and season = '%v'", challengeFields, crossFsId, round, season))
	if err != nil {
		return nil, err
	}
	return challenge, nil
}

// 获取指定服务器的报名玩家
func (this *ChallengeModel) CheckUserIsInRound(serverId, round, userId, crossFsId int, season string) (*Challenge, error) {
	var challenge *Challenge

	err := this.DbMap().SelectOne(&challenge, fmt.Sprintf("select %s from challenge where serverId = %v and round = %v and userId = %v and crossFsId = %v and season = '%v'", challengeFields, serverId, round, userId, crossFsId, season))
	if err != nil {
		return nil, err
	}
	return challenge, nil
}

// 获取指定服务器的报名玩家
func (this *ChallengeModel) CheckUserIsInRound1(round, userId, crossFsId int, season string) (*Challenge, error) {
	var challenge *Challenge
	err := this.DbMap().SelectOne(&challenge, fmt.Sprintf("select %s from challenge where  round = %v and userId = %v and crossFsId = %v and season = '%v'", challengeFields, round, userId, crossFsId, season))
	if err != nil {
		return nil, err
	}
	return challenge, nil
}

func (this *ChallengeModel) UpdateRound(round, crossFsId int, season string) error {
	sql := fmt.Sprintf("update challenge SET round = %v WHERE round = 0  and crossFsId = %v and season = '%v';", round, crossFsId, season)
	_, err := this.DbMap().Exec(sql)
	return err
}

func (this *ChallengeModel) UpdateExpTime(expireTime int64) error {
	sql := fmt.Sprintf("update challenge SET expireTime = %v WHERE expireTime = 0;", expireTime)
	_, err := this.DbMap().Exec(sql)
	return err
}

func (this *ChallengeModel) MaxRobotId(round, crossFsId int, season string) (int, error) {
	sql := fmt.Sprintf("select min(userId) from challenge where season = '%v' and crossFsId = %v ;", season, crossFsId)
	var maxId int
	err := this.DbMap().SelectOne(&maxId, sql)
	if err != nil {
		return -1, err
	}
	return maxId, nil

}
