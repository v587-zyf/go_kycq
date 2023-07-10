package modelCross

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/dbmodel"
	"cqserver/golibs/logger"
	"fmt"
	"gopkg.in/gorp.v1"
)

var log = logger.Get("default", true)

//跨服战fightServer信息
type CrossFightServerInfo struct {
	Id       int    `db:"id"`
	Host     string `db:"host"`
	GatePort int    `db:"gatePort"`
	GsPort   int    `db:"gsPort"`
}

type CrossFightServerInfoAndServer struct {
	Id       int    `db:"id"`
	Host     string `db:"host"`
	GatePort int    `db:"gatePort"`
	GsPort   int    `db:"gsPort"`
	ServerId int    `db:"serverId"`
}

//跨服战fight服务器列表
type CrossFightServerInfoModel struct {
	dbmodel.CommonModel
}

var (
	cfsInfoModel  = &CrossFightServerInfoModel{}
	cfsInfoFields = model.GetAllFieldsAsStringWithTableName(CrossFightServerInfo{}, "crossfight_server_info")
)

func init() {
	dbmodel.Register(model.DB_ACCOUNT, cfsInfoModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(CrossFightServerInfo{}, "crossfight_server_info").SetKeys(true, "id")
	})
}

func GetCrossFightServerInfoModel() *CrossFightServerInfoModel {
	return cfsInfoModel
}

func (this *CrossFightServerInfoModel) GetCrossFightServerInfo(serverId int) (*CrossFightServerInfo, error) {
	var info CrossFightServerInfo
	err := this.DbMap().SelectOne(&info, fmt.Sprintf("select %s from crossfight_server_info where id=?;", cfsInfoFields),serverId)
	if err != nil {
		//log.Error("GetStaticCrossFightServerInfo failed : %v", err)
		return nil, err
	}
	return &info, nil
}

// 获取所有的动态战斗跨服
func (this *CrossFightServerInfoModel) GetAllCrossFightServerList() ([]CrossFightServerInfo, error) {
	var servers []CrossFightServerInfo

	_, err := this.DbMap().Select(&servers, fmt.Sprintf("select %s from crossfight_server_info where 1", cfsInfoFields))
	if err != nil {
		return nil, err
	}
	return servers, nil
}

// 获取所有的动态战斗跨服
func (this *CrossFightServerInfoModel) GetAllCrossFightServerListAndServer() ([]CrossFightServerInfoAndServer, error) {
	var servers []CrossFightServerInfoAndServer

	_, err := this.DbMap().Select(&servers, "select f.*,s.serverId from crossfight_server_info as f LEFT JOIN server_info as s on  f.id = s.crossFsId where 1;")
	if err != nil {
		return nil, err
	}
	return servers, nil
}
