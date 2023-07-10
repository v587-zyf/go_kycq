package modelCross

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/logger"

	//	"database/sql"
	"fmt"

	//	"strings"

	"gopkg.in/gorp.v1"

	"cqserver/golibs/dbmodel"
)

type SystemSetting struct {
	Id      int    `db:"id"`
	Setting string `db:"setting"`
	Value   string `db:"value"`
}

type SystemSettingModel struct {
	dbmodel.CommonModel
}

var (
	systemSettingModel  = &SystemSettingModel{}
	systemSettingFileds = model.GetAllFieldsAsString(SystemSetting{})
)

func init() {

	dbmodel.Register(model.DB_ACCOUNT, systemSettingModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(SystemSetting{}, "system_setting").SetKeys(true, "id")
	})
}

func GetSystemSettingModel() *SystemSettingModel {
	return systemSettingModel
}

func (this *SystemSettingModel) GetAllSetting() ([]SystemSetting, error) {
	var systemSetting []SystemSetting
	_, err := this.DbMap().Select(&systemSetting, fmt.Sprintf("select %s from system_setting  where 1", systemSettingFileds))
	if err != nil {
		return nil, err
	}
	return systemSetting, nil
}

func (this *SystemSettingModel) GetSetting(settingName string) (string, error) {
	var systemSetting SystemSetting
	err := this.DbMap().SelectOne(&systemSetting, fmt.Sprintf("select %s from system_setting  where setting=?", systemSettingFileds), settingName)
	if err != nil {
		logger.Error("获取setting配置异常：%v,err:%v", settingName, err)
		return "", err
	}
	return systemSetting.Value, nil
}
