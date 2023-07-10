package modelCross

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/logger"
	"fmt"
	"cqserver/golibs/dbmodel"
	"gopkg.in/gorp.v1"
)

const (
	CLIENT_TO_LOGIN         = "clientToLogin_%d"
	GATE_TO_LOGIN           = "gateToLogin"
	GAME_TO_CROSSCENTER     = "gameToCrosscenter"
	GATE_TO_FIGHTCENTER     = "gateToFightcenter"
	GAME_TO_FIGHTCENTER     = "gameToFightcenter"
	CROSS_CENTER_HTTP       = "crossCenterHttpPort"
	TRIAL_CROSS_CENTER_HTTP = "trialCrossCenterHttp"
)

// 单个服务器信息
type ServerPortInfo struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Host string `db:"host"`
	Port int    `db:"port"`
}

type ServerPortInfoModel struct {
	dbmodel.CommonModel
}

var (
	serverPortInfoModel  = &ServerPortInfoModel{}
	serverInfoPortFields = model.GetAllFieldsAsStringWithTableName(ServerPortInfo{}, "server_port")
)

func init() {
	dbmodel.Register(model.DB_ACCOUNT, serverPortInfoModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(ServerPortInfo{}, "server_port").SetKeys(false, "id")
	})

}

func GetServerPortInfoModel() *ServerPortInfoModel {
	return serverPortInfoModel
}

func (this *ServerPortInfoModel) GetServerPortInfo(name string) (*ServerPortInfo, error) {
	var info *ServerPortInfo
	err := this.DbMap().SelectOne(&info, fmt.Sprintf("select %s from server_port where name = ?", serverInfoPortFields), name)
	if err != nil {
		return nil, err
	}
	logger.Info("获取服务器端口name(%v)=>port(%v)", name, info.Port)
	return info, nil
}

func (this *ServerPortInfoModel) GetLoginServerPortInfos(name string) ([]ServerPortInfo, error) {
	var info []ServerPortInfo

	_, err := this.DbMap().Select(&info, fmt.Sprintf(`select %s from server_port where name like "%s%%"`, serverInfoPortFields, name))
	if err != nil {
		return nil, err
	}
	return info, nil
}
