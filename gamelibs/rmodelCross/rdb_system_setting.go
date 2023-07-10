package rmodelCross

import (
	"cqserver/gamelibs/modelCross"
	"strconv"
	"sync"

	"cqserver/golibs/logger"
	"errors"
)

var settingCheckSlice []string = []string{}

func settingParamcheck(key string) string {
	settingCheckSlice = append(settingCheckSlice, key)
	return key
}

var (
	SYSTEM_SETTING_AREA_NAME                      = settingParamcheck("area_name")
	SYSTEM_SETTING_CROSS_OPEN_DAY                 = settingParamcheck("cross_open_day")
	SYSTEM_SETTING_CROSS_OPEN_ACTIVE_PLAYER       = settingParamcheck("cross_open_active_player")
	SYSTEM_SETTING_CROSS_OPEN_SERVER_DAY_RECHARGE = settingParamcheck("cross_open_server_day_recharge")
	SYSTEM_SETTING_CROSS_MAX_SERVER               = settingParamcheck("cross_max_server")
	SYSTEM_SETTING_CROSS_MIN_SERVER               = settingParamcheck("cross_min_server")
	SYSTEM_SETTING_MAX_ACTIVITY_DAY               = settingParamcheck("max_activity_day")
	SYSTEM_SETTING_CROSS_ACTIVITY_USER_DAY        = settingParamcheck("cross_activity_user_day") //跨服  取几天内活跃的玩家
	SYSTEM_SETTING_TRIAL_SERVER_VERSION           = settingParamcheck("trial_server_version")    //提审服版本
	SYSTEM_SETTING_TRIAL_SERVER                   = settingParamcheck("trial_server")            //标识是否提审服

)

type SystemSetting struct {
	settingData sync.Map
}

var systemSetting = &SystemSetting{}

func GetSystemSeting() *SystemSetting {
	return systemSetting
}

func (this *SystemSetting) GetSystemSetting(setting string) string {

	value, ok := this.settingData.Load(setting)
	if !ok || value == nil {
		logger.Debug("setting :%v,不存在，数据库拉取", setting)
		this.CronUpdate()
		value, ok = this.settingData.Load(setting)
		if !ok {
			return ""
		}
		return value.(string)
	}
	return value.(string)
}

func (this *SystemSetting) GetSystemSettingConverInt(setting string) int {

	value, ok := this.settingData.Load(setting)
	if !ok || value == nil {
		logger.Debug("setting :%v,不存在，数据库拉取", setting)
		this.CronUpdate()
		value, ok = this.settingData.Load(setting)
		if !ok {
			return -1
		}
	}
	valueStr := value.(string)
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		logger.Error("获取setting值转换int异常,setting:%v，值：%v", setting, value)
	}
	return valueInt
}

func (this *SystemSetting) InitCheck() error {
	this.CronUpdate()
	settingMiss := make([]string, 0)
	for _, v := range settingCheckSlice {
		if value, ok := this.settingData.Load(v); !ok || value == nil {
			settingMiss = append(settingMiss, v)
		}
	}
	if len(settingMiss) > 0 {
		logger.Error("system_setting配置缺失：%v", settingMiss)
		return errors.New("system_param或system_setting配置缺失")
	}
	return nil
}

//定时更新，crosscenter定时更新，拉取所有system_param system_setting数据到redis
func (this *SystemSetting) CronUpdate() {

	allSetting, err1 := modelCross.GetSystemSettingModel().GetAllSetting()
	if err1 != nil {
		logger.Error("获取所有system_setting错误:%v", err1)
	}
	for _, v := range allSetting {
		this.settingData.Store(v.Setting, v.Value)
	}
	logger.Debug("获取所有setting:%v", allSetting)
}
