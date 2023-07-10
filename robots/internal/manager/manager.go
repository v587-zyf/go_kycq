package manager

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelCross"
	"cqserver/golibs/dbmodel"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw/httpclient"
	"cqserver/robots/conf"
	"encoding/json"
	"flag"
	"fmt"
)

var (
	gameDbBasePath = flag.String("gamedb", "../../yulong/data/configs", "specify gamedb path")
	platform       = flag.String("platform", "", "specify gamedb path")
)

type Manager struct {
	serverInfo *modelCross.ServerInfo
	robot      *Robot
	logginKey  string
}

var m = &Manager{}

func Get() *Manager {
	return m
}

type LoginRet struct {
	OpenId   string `json:"openId"`
	Cmd      string `json:"cmd"`
	LoginKey string `json:"loginkey"`
	Ip       string `json:"ip"`
}

func (this *Manager) Init(robot string) {

	err := gamedb.Load(*gameDbBasePath)
	if err != nil {
		logger.Error("加载gamedb错误")
		return
	}

	//先初始化accountdb
	if err := dbmodel.InitDbByKey(conf.Conf, model.DB_ACCOUNT, false, false, false); err != nil {
		panic("初始化数据库错误")
	}

	//获取服务器的db配置
	serverInfo, err := modelCross.GetServerInfoModel().GetServerInfoByServerId(conf.Conf.ServerId)
	if err != nil {
		panic("获取服务器异常")
	}
	this.serverInfo = serverInfo
	this.robot = NewRobot(robot, conf.Conf.ServerId)

	if len(*platform) > 0 && *platform == "602" {
		param := make(map[string]interface{})
		param["openId"] = robot
		param["debug"] = "1da5720b7a4e3e9f2de3c007804d4a02"
		rb, err := httpclient.DoPost("https://wmcqlogin.hhqaq.com/login", ToUrlValues(param))
		//rb, err := httpclient.DoPost("http://127.0.0.1:7100/login", ToUrlValues(param))
		if err != nil {
			logger.Error("登录异常：%v", err)
			panic(err)
		}
		logger.Info("--------", string(rb))
		var result LoginRet
		err = json.Unmarshal(rb, &result)
		if err != nil {
			panic(err)
		}
		this.logginKey = result.LoginKey
		logger.Info("----", result)
	}
}

func (this *Manager) GetAddr() string {

	if *platform == "602" {
		return "wss://wmcqgs.hhqaq.com/ws?host=10.1.84.13&port=7001"
	} else {
		return fmt.Sprintf("ws://%s/ws", this.serverInfo.Gates)
	}
}

func (this *Manager) Start() {
	go this.robot.Start()
}

func (this *Manager) Stop() {
	this.robot.Stop()
}
