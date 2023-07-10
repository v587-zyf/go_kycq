package model

import (
	"cqserver/gamelibs/modelGame"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"

	"cqserver/golibs/dbmodel"
	"cqserver/golibs/logger"
	"cqserver/tools/serverMerge/internal/base"
	"gopkg.in/gorp.v1"
	"strconv"
)

type UserModel struct {
	dbmodel.CommonModel
}

type ModuleState struct {
	ZmClick      int
	TaskAutoPick int
}

type UserModels struct {
	userModels   map[string]*UserModel
	newUserModel *UserModel
}

var (
	userModels = &UserModels{
		userModels:   make(map[string]*UserModel),
		newUserModel: &UserModel{},
	}
	userFields = GetAllFieldsAsString(modelGame.User{})
	userTableName = (&modelGame.User{}).TableName()
)

func dbUserInit() {
	logger.Info("userModels 初始化")
	for k, _ := range base.Conf.DbConfigs {
		k1 := strings.Split(k, "_")
		if len(k1) == 2 && k1[0] == DB_SERVER {
			if userModels.userModels[k] == nil {
				userModels.userModels[k] = &UserModel{}
			}
			dbmodel.Register(k, userModels.userModels[k], func(dbMap *gorp.DbMap) {
				dbMap.AddTableWithName(modelGame.User{}, userTableName).SetKeys(false, "id")
			})
		}
	}
	dbmodel.Register(NEW_SERVER, userModels.newUserModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(modelGame.User{}, userTableName).SetKeys(false, "id")
		orm.RegisterModelForAlias(NEW_SERVER, new(modelGame.User))
	})
}

func GetUserModel() *UserModels {
	return userModels
}

//获取有效玩家数据
func (this *UserModels) LoadAllUsers(dbKey string, activeDay int, rechargeMin int, levelMin int, combatMin int) ([]modelGame.User, error) {
	var users []modelGame.User
	var sqlStr = fmt.Sprintf("select %s from user where NickName != '' and ( rechargeAll > %d ", userFields, rechargeMin)
	if activeDay > 0 {
		t := time.Now().Add(time.Duration(-1*24*activeDay) * time.Hour).Format("2006-01-02 15:04:05")
		sqlStr += " or offlineTime >= '" + t + "' or lastUpdateTime >= '" + t + "'"
	}

	if combatMin > 0 {
		sqlStr += " or combat>" + strconv.Itoa(combatMin)
	}
	sqlStr += " )"

	//fmt.Println(sqlStr)
	_, err := this.userModels[dbKey].DbMap().Select(&users, sqlStr)
	if err != nil {
		return nil, err
	}
	return users, nil
}

//插入玩家数据
func (this *UserModels) InsertNewData(user *modelGame.User) error {
	return this.newUserModel.DbMap().Insert(user)
}

func (this *UserModels) Clean() {
	this.newUserModel.DbMap().Exec("delete from user where 1")
	this.newUserModel.DbMap().Exec("delete from hero where 1")
}
