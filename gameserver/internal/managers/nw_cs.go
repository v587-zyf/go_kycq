package managers

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/objs"
	"errors"
	"fmt"
	"time"

	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/tcpclient"
	"cqserver/golibs/util"
	"cqserver/protobuf/pbserver"
)

type CsManager struct {
	util.DefaultModule
	csSession       *CsSession
	dialerCloseChan chan struct{}
}

func NewCsManager() *CsManager {
	csManager := &CsManager{
		csSession: NewCsSession(),
	}
	return csManager
}

func (this *CsManager) Init() error {
	serverPortInfo, err := modelCross.GetServerPortInfoModel().GetServerPortInfo(modelCross.GAME_TO_CROSSCENTER)
	if err != nil || serverPortInfo == nil {
		return errors.New(fmt.Sprintf("获取服务器启动端口号错误：%v", modelCross.GAME_TO_CROSSCENTER))
	}
	context := &nw.Context{
		SessionCreator: func(conn nw.Conn) nw.Session {
			this.csSession.Conn = conn
			return this.csSession
		},
		Splitter:       pbserver.Split,
		ReadBufferSize: 1024 * 64,
	}

	addr := fmt.Sprintf("%s:%d", serverPortInfo.Host, serverPortInfo.Port)
	this.dialerCloseChan = tcpclient.DialEx(addr, context, 3*time.Second)
	logger.Info("game crosscenter 开始建立链接：%v", addr)
	return nil
}

func (this *CsManager) Stop() {
	close(this.dialerCloseChan)
	this.csSession.Close()
	logger.Info("nw_cs stop")
}

func (this *CsManager) SendMsgToCenterServer(transId uint32, msg nw.ProtoMessage) {
	this.csSession.SendMessage(transId, msg)
}

func (this *CsManager) RpcCall(req nw.ProtoMessage, resp nw.ProtoMessage) error {
	return this.csSession.Call(req, resp)
}

func (this *CsManager) SyncUser(user *objs.User, status int, lastRechargeTime int64, realRechargeTotal, tokenRechargeTotal int) {

	syncMsg := builder.BuildSyncUserInfoNtf(user, status, lastRechargeTime)
	syncMsg.Recharge = int32(realRechargeTotal)
	syncMsg.TokenRecharge = int32(tokenRechargeTotal)
	this.SendMsgToCenterServer(0, syncMsg)
}
