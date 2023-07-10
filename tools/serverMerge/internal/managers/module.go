package managers

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/golibs/dbmodel"
	"cqserver/tools/serverMerge/internal/base"
	"flag"

	"cqserver/gamelibs/beans"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/tools/serverMerge/internal/model"
	"cqserver/tools/serverMerge/internal/rmodel"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ModuleManager struct {
	*util.DefaultModuleManager

	mergeSrverInfos []modelCross.ServerInfo

	DbDataMerge       *DbDataMerge
	RdbDataMerge      *RdbDataMerge
	BaseFunctionMerge *BaseFunctionMerge
}

var m = &ModuleManager{
	DefaultModuleManager: util.NewDefaultModuleManager(),
}
var (
	gameDbBasePath = flag.String("gamedb", "../gameserver/yulong/data/configs", "specify gamedb path")
)

func Get() *ModuleManager {
	return m
}

func (this *ModuleManager) init() error {

	var err error
	err = gamedb.Load(*gameDbBasePath)
	if err != nil {
		logger.Debug("err:%v  *gameDbBasePath:%v", err, *gameDbBasePath)
		return err
	}

	//检查要合服的服务器
	mergeServerIds := strings.Split(base.Conf.MergeServerIds, ",")
	if len(mergeServerIds) < 2 {
		return errors.New("一个服你合个锤子啊~~")
	}
	//初始化accountdb
	if err = dbmodel.InitDb(base.Conf, nil, "", nil); err != nil {
		return err
	}

	err = rmodel.Init(base.Conf.Redis, -1)
	if err != nil {
		return err
	}

	//获取要合并服务器的数据库配置
	mergeServers, err := modelCross.GetServerInfoModel().GetServerListAll(mergeServerIds)
	if err != nil {
		return err
	}
	if len(mergeServers) != len(mergeServerIds) {
		return errors.New("合并的服务器与数据库拉取到的服务器信息不匹配")
	}
	this.mergeSrverInfos = mergeServers
	for _, v := range mergeServers {
		dbConfig := this.GetBeanDb(v.DbLink)
		if dbConfig == nil {
			return errors.New(fmt.Sprintf("合并服务器：%v,获取服务器数据库配置错误", v.ServerId))
		}
		base.Conf.DbConfigs[model.DB_SERVER+"_"+strconv.Itoa(v.ServerId)] = dbConfig
	}
	//注册合并服务器连接
	model.Init()
	//db重新初始化一次（会导致account库连接重新初始化）
	if err = dbmodel.InitDb(base.Conf, []string{model.NEW_SERVER}, model.NEW_SERVER, []string{model.NEW_SERVER}); err != nil {
		return err
	}
	return nil
}

func (this *ModuleManager) Init() error {
	logger.Info("base init")
	err := this.init()
	if err != nil {
		return err
	}
	logger.Info("module init")
	this.DbDataMerge = this.AppendModule(NewDbDataMerge()).(*DbDataMerge)
	this.RdbDataMerge = this.AppendModule(NewRdbDataMerge()).(*RdbDataMerge)
	this.BaseFunctionMerge = this.AppendModule(NewBaseFunctionMerge()).(*BaseFunctionMerge)

	logger.Info("DefaultModuleManager init")
	err = this.DefaultModuleManager.Init()

	return err
}

func (this *ModuleManager) StartMerge() {

	logger.Info("置空新表数据")
	model.Clean()

	//数据合并
	logger.Info("-------------------开始合服-------------------------------")
	if ok := this.beforeCheck(); !ok {
		panic("合服前检查错误，程序退出")
		return
	}
	logger.Info("---------------db数据处理合并开始-------------------------")
	if ok := this.DbDataMerge.Merge(); !ok {
		panic("合服数据库数据处理错误，程序退出")
		return
	}
	logger.Info("---------------db数据处理合并结束-------------------------")
	//redis数据合并
	logger.Info("---------------redis数据合并开始-------------------------")
	if ok := this.RdbDataMerge.Merge(this.mergeSrverInfos); !ok {
		panic("合服redis数据处理错误，程序退出")
		return
	}
	logger.Info("---------------redis数据合并结束-------------------------")
	//更新serverInfo配置
	if ok := this.updateMergeServerInfo(); !ok {
		panic("合服更新主服serverInfo信息错误，程序退出")
		return
	}
	//删除account库玩家数据
	this.delAccountUser()
	logger.Info("-------------------合服结束------------------------------")
}

func (this *ModuleManager) beforeCheck() bool {

	return true
}

//更新合服serverInfo信息
func (this *ModuleManager) updateMergeServerInfo() bool {
	logger.Info("更新合并服务器server_info信息")

	mergeServerIds := strings.Split(base.Conf.MergeServerIds, ",")
	mainServerId, _ := strconv.Atoi(mergeServerIds[0])
	var mainServerInfo *modelCross.ServerInfo
	for _, v := range this.mergeSrverInfos {
		if v.ServerId == mainServerId {
			//gates = v.Gates
			mainServerInfo = &v
			break
		}
	}
	if mainServerInfo == nil {
		logger.Error("更新合并服务器信息，新gates获取失败")
		return false
	}
	dbLink := "server=" + base.Conf.DbConfigs[model.NEW_SERVER].Host +
		";port=" + strconv.Itoa(base.Conf.DbConfigs[model.NEW_SERVER].Port) +
		";database=" + base.Conf.DbConfigs[model.NEW_SERVER].DbName +
		";uid=" + base.Conf.DbConfigs[model.NEW_SERVER].Uid +
		";pwd=" + base.Conf.DbConfigs[model.NEW_SERVER].Pwd + ";charset=utf8"
	redsAddr := base.Conf.Redis.Address + ":" + base.Conf.Redis.Password + ":" + strconv.Itoa(base.Conf.Redis.DB)
	dbLog := base.Conf.DbConfigs[model.NEW_SERVER].Host + ":" + strconv.Itoa(base.Conf.DbConfigs[model.NEW_SERVER].Port) + "/" + base.Conf.DbLogName

	var err error
	if base.Conf.SandBox {
		_, err = modelCross.GetServerInfoModel().UpdateServerDbLinkAndRedis(mergeServerIds, mainServerInfo.Gates, dbLink, redsAddr, dbLog)
	} else {
		_, err = modelCross.GetServerInfoModel().UpdateServerMergeId(mergeServerIds,
			mainServerInfo.Gates,
			dbLink,
			redsAddr,
			dbLog,
			mainServerInfo.GsHostIn,
			mainServerInfo.GsHostWWW,
			strconv.Itoa(mainServerInfo.GsPort),
			mainServerInfo.HttpPort)
	}
	if err != nil {
		logger.Error("更新合并服务器的信息错误：%v", err)
		return false
	}
	logger.Info("更新合并服务器server_info信息完成")
	return true
}

func (this *ModuleManager) delAccountUser() {

	for sid, users := range this.DbDataMerge.mergeServerUsers {
		userIds := ""
		//玩家所在原始服务器
		usidMap := make(map[int]bool, 0)
		for uid, usid := range users {
			userIds += "'" + strconv.Itoa(uid) + "',"
			usidMap[usid] = true
		}
		usidStr := ""
		for usid, _ := range usidMap {
			usidStr += "'" + strconv.Itoa(usid) + "',"
		}
		//所有有效玩家
		userIds = userIds[:len(userIds)-1]
		//有效玩家所在原始服务器
		usidStr = usidStr[:len(usidStr)-1]
		modelCross.GetServerInfoModel().DelInvalidUser(usidStr, userIds)
		logger.Info("合并服务器删除服务器[%v]无效账号，有效账号[%v]", sid, userIds)
	}
}

func (this *ModuleManager) GetBeanDb(dblink string) *beans.DbConfig {

	dbLinkMap := make(map[string]string)
	dbSlice := strings.Split(dblink, ";")
	for _, v := range dbSlice {
		vv := strings.Split(v, "=")
		dbLinkMap[vv[0]] = vv[1]
	}
	dbConfig := &beans.DbConfig{
		Host:   dbLinkMap["server"],
		DbName: dbLinkMap["database"],
		Uid:    dbLinkMap["uid"],
		Pwd:    dbLinkMap["pwd"],
	}
	port, err := strconv.Atoi(dbLinkMap["port"])
	if err != nil {
		logger.Error("合并服务器：%v,端口转换错误：%v", dblink, err)
		return nil
	}
	dbConfig.Port = port
	return dbConfig
}
