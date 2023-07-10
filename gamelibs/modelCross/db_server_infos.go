package modelCross

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/logger"
	"strconv"
	"time"

	"cqserver/golibs/dbmodel"
	"fmt"
	"gopkg.in/gorp.v1"
)

// 单个服务器信息
type ServerInfo struct {
	dbmodel.DbTable
	Id             int       `db:"id" `
	Name           string    `db:"name" orm:"size(50);comment(服务器名字)"`
	ServerIndex    int       `db:"serverIndex" orm:"comment(服务器索引)"`
	ServerId       int       `db:"serverId" orm:"comment(服务器Id)"`
	AppId          string    `db:"appId" orm:"size(50);comment(渠道标识)"`
	MergeServerId  int       `db:"mergeServerId" orm:"comment(合到服务器Id)"`
	MergeTime      time.Time `db:"mergeTime" orm:"comment(和服时间)"`
	CrossFsId      int       `db:"crossFsId" orm:"commen(跨服战斗服Id)"`
	CrossFirst     int       `db:"crossFirst" orm:"commen(第一次跨服)"`
	Gates          string    `db:"gates"`
	GsHostIn       string    `db:"gsHostIn"` // gameserver内网地址
	GsPort         int       `db:"gsPort"`
	GsfsPort       int       `db:"gsfsPort"`
	GatefsPort     int       `db:"gatefsPort"`
	GsHostWWW      string    `db:"gsHostWww"` // gameserver外网地址
	HttpPort       string    `db:"httpPort"`
	IsNew          int       `db:"isNew"`
	Status         int       `db:"status"` // 1:良好，2:正常，3:爆满
	OpenTime       time.Time `db:"openTime"`
	IsClose        int       `db:"isClose"`
	CloseExplain   string    `db:"closeExplain"`
	Prefix         string    `db:"prefix"`
	Version        string    `db:"version"`
	ClientVersion  string    `db:"clientVersion"`                        // 客户端版本号
	IpFilter       int       `db:"ipFilter"`                             //是否开启白名单登录0否，1是
	IsTrialVersion int       `db:"isTrialVersion" orm:"comment(是否是体验服)"` //是否是否是体验服0否，1是
	DbLink         string    `db:"dblink"`
	DbLinkLog      string    `db:"dblinkLog"`
	RedisAddr      string    `db:"redisAddr"`
}

func (this *ServerInfo) TableName() string {
	return "server_info"
}

type ServerInfoModel struct {
	dbmodel.CommonModel
}

var (
	serverInfoModel  = &ServerInfoModel{}
	serverInfoFields = model.GetAllFieldsAsStringWithTableName(ServerInfo{}, "server_info")
)

func init() {
	dbmodel.Register(model.DB_ACCOUNT, serverInfoModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(ServerInfo{}, "server_info").SetKeys(false, "id")
	})

}

func GetServerInfoModel() *ServerInfoModel {
	return serverInfoModel
}

// 获取所有的服务器列表
func (this *ServerInfoModel) GetServerListAll(serverIds []string) ([]ServerInfo, error) {

	ids := ""
	for _, v := range serverIds {
		ids += "'" + v + "',"
	}
	ids = ids[:len(ids)-1]
	var servers []ServerInfo
	_, err := this.DbMap().Select(&servers, fmt.Sprintf("select %s from server_info where serverId in (%s)", serverInfoFields, ids))
	if err != nil {
		return nil, err
	}
	return servers, nil
}

// 获取指定服务器信息
func (this *ServerInfoModel) GetServerInfoByServerId(serverId int) (*ServerInfo, error) {
	var server ServerInfo
	err := this.DbMap().SelectOne(&server, fmt.Sprintf("select %s from server_info where serverId=%d", serverInfoFields, serverId))
	if err != nil {
		return nil, err
	}
	return &server, nil
}

//根据和服服务器Id获取服务器信息
func (this *ServerInfoModel) GetServerInfoByMergerServerId(serverId int) ([]*ServerInfo, error) {
	var serverInfos []*ServerInfo
	sqlStr := fmt.Sprintf("select %s from server_info where mergeServerId = %d", serverInfoFields, serverId)
	_, err := this.DbMap().Select(&serverInfos, sqlStr)
	if err != nil {
		return nil, err
	}
	return serverInfos, nil
}

func (this *ServerInfoModel) GetCrossGroupByOpenDay(openDay int) ([]string, error) {
	var crossGroup []string
	_, err := this.DbMap().Select(&crossGroup, fmt.Sprintf("select crossGroup from `server_info` where DATEDIFF(NOW(),openTime) >= ?"), openDay)
	if err != nil {
		return nil, err
	}
	return crossGroup, nil
}

func (this *ServerInfoModel) GetServerInfos() ([]*ServerInfo, error) {
	var infos []*ServerInfo
	_, err := this.DbMap().Select(&infos, fmt.Sprintf("select %s from server_info where status != 4 order by id", serverInfoFields))
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func (this *ServerInfoModel) UpdateServerOpenTime(serverId int, open time.Time) error {
	sqlStr := fmt.Sprintf("update server_info set openTime = ? where serverId = ?")
	_, err := this.DbMap().Exec(sqlStr, open, serverId)
	return err
}

func (this *ServerInfoModel) UpdateServerInfosFromGM(param map[string]interface{}) error {

	sql := "Update server_info set"
	serverId := int(param["serverId"].(float64))
	hasUpdate := false
	if serverId <= 0 {

		//更新全服服务器，只更新白名单 维护状态
		if isClose, ok := param["isMaintain"]; ok {
			sql = fmt.Sprintf(sql+" isClose=%v,", int(isClose.(float64)))
			hasUpdate = true
		}
		if isOpenWhitelist, ok := param["isOpenWhitelist"]; ok {
			sql = fmt.Sprintf(sql+" ipFilter=%v,", int(isOpenWhitelist.(float64)))
			hasUpdate = true
		}
		if hasUpdate {
			sql = sql[0 : len(sql)-1]
		}
		sql += " where status<=3;"
	} else {

		if firstOpenTime, ok := param["firstOpenTime"]; ok {
			//openTime, _ := common.GetTime(firstOpenTime.(string))
			sql = fmt.Sprintf(sql+" openTime=\"%v\",", firstOpenTime)
			hasUpdate = true
		}

		if isClose, ok := param["isMaintain"]; ok {
			sql = fmt.Sprintf(sql+" isClose=%v,", int(isClose.(float64)))
			hasUpdate = true
		}
		if isOpenWhitelist, ok := param["isOpenWhitelist"]; ok {
			sql = fmt.Sprintf(sql+" ipFilter=%v,", int(isOpenWhitelist.(float64)))
			hasUpdate = true
		}

		if isNew, ok := param["isNew"]; ok {
			sql = fmt.Sprintf(sql+" isNew=%v,", int(isNew.(float64)))
			hasUpdate = true
		}

		if artificialLoad, ok := param["artificialLoad"]; ok {
			sql = fmt.Sprintf(sql+" status=%v,", int(artificialLoad.(float64)))
			hasUpdate = true
		}
		if hasUpdate {
			sql = sql[0 : len(sql)-1]
		}
		sql += " where serverId=" + strconv.Itoa(serverId) + ";"
	}
	var err error
	if hasUpdate {
		logger.Info("gm 更新服务器：%v", sql)
		_, err = this.DbMap().Exec(sql)
	}
	return err
}

// 获取所有的服务器列表
func (this *ServerInfoModel) UpdateServerMergeId(serverIds []string, gates string, dbLink string, redsAddr string, dbLog string,
	gsHostIn string, gsHostWWW string, gsPort string, httpPort string) ([]ServerInfo, error) {

	ids := ""
	for _, v := range serverIds {
		ids += "'" + v + "',"
	}
	ids = ids[:len(ids)-1]
	_, err := this.DbMap().Exec(fmt.Sprintf("update server_info set "+
		"mergeServerId = %s,"+"gates='%s',"+"dbLink = '%s',"+"redisAddr='%s',"+"dblinkLog='%s',"+
		"gsHostIn='%s',"+"gsHostWww='%s',"+"gsPort='%s',httpPort='%s',mergeTime=NOW() where mergeServerId in (%s)",
		serverIds[0], gates, dbLink, redsAddr, dbLog, gsHostIn, gsHostWWW, gsPort, httpPort, ids))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// 获取所有的服务器列表
func (this *ServerInfoModel) UpdateServerDbLinkAndRedis(serverIds []string, gates string, dbLink string, redisAddr string, dbLog string) ([]ServerInfo, error) {

	ids := ""
	for _, v := range serverIds {
		ids += "'" + v + "',"
	}
	ids = ids[:len(ids)-1]
	_, err := this.DbMap().Exec(fmt.Sprintf("update server_info set mergeServerId = %s,gates='%s',mergeTime=NOW() where mergeServerId in (%s)", serverIds[0], gates, ids))
	if err != nil {
		return nil, err
	}

	_, err = this.DbMap().Exec(fmt.Sprintf("update server_info set dbLink = '%s',redisAddr='%s',dblinkLog='%s' where serverId = %s", dbLink, redisAddr, dbLog, serverIds[0]))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//删除无效玩家
func (this *ServerInfoModel) DelInvalidUser(serverId string, userIds string) {

	this.DbMap().Exec(fmt.Sprintf("delete from user where serverId in (%s) and userId not in (%s)", serverId, userIds))
}

//批量更新跨服组信息
//serverIds: 逗号分隔的字符串
func (this *ServerInfoModel) UpdateCrossGroups(serverIds string, crossFsId int) error {
	var sql string
	sql = "Update server_info SET crossFsId=%d,crossFirst=1 WHERE serverId IN(%s)"
	sql = fmt.Sprintf(sql, crossFsId, serverIds)
	logger.Info("UpdateCrossGroups sql: %s", sql)
	_, err := this.DbMap().Exec(sql)
	if err != nil {
		logger.Error("ServerInfoModel UpdateCrossGroups Error:%v", err)
	}
	return err
}

//获取所有跨服组
func (this *ServerInfoModel) GetAllCrossFsIds() ([]ServerInfo, error) {
	var infos []ServerInfo
	_, err := this.DbMap().Select(&infos, fmt.Sprintf("select %s from server_info where crossFsId > 0", serverInfoFields))
	if err != nil {
		return nil, err
	}
	return infos, nil
}

//获取所有跨服组
func (this *ServerInfoModel) GetAllServerIdsByCrossFsIds(crossFsId int) ([]ServerInfo, error) {
	var infos []ServerInfo
	_, err := this.DbMap().Select(&infos, fmt.Sprintf("select %s from server_info where crossFsId = %v", serverInfoFields, crossFsId))
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func (this *ServerInfoModel) GetAllMergerServerIds(mergeServerId int) ([]ServerInfo, error) {
	var infos []ServerInfo
	_, err := this.DbMap().Select(&infos, fmt.Sprintf("select %s from server_info where mergeServerId = %v", serverInfoFields, mergeServerId))
	if err != nil {
		return nil, err
	}
	return infos, nil
}
