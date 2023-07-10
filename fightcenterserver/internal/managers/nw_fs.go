package managers

import (
	"cqserver/gamelibs/modelCross"
	"errors"
	"fmt"
	"time"

	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/tcpclient"
	"cqserver/golibs/util"
	"cqserver/protobuf/pbgt"
)

type FSManager struct {
	util.DefaultModule

	gateSessions         map[int]*FSSession //连接战斗服的gate端口 key:CrossServerId
	gateDialerCloseChans map[int]chan struct{}

	gsSessions         map[int]*FSSession //连接战斗服的gs端口 key:CrossServerId
	gsDialerCloseChans map[int]chan struct{}
}

func NewFSManager() *FSManager {
	fs := &FSManager{
		gateSessions:         make(map[int]*FSSession),
		gateDialerCloseChans: make(map[int]chan struct{}),

		gsSessions:         make(map[int]*FSSession),
		gsDialerCloseChans: make(map[int]chan struct{}),
	}

	return fs
}

func GetFSSessionCreator(session *FSSession) func(conn nw.Conn) nw.Session {
	return func(conn nw.Conn) nw.Session {
		session.Conn = conn
		return session
	}
}

func (this *FSManager) Init() error {

	go util.SafeRun(this.reloadDynamicCrossFightServer)

	return nil
}

func (this *FSManager) setGateDialerCloseChans(FSNo int, dialerCloseChan chan struct{}) {
	this.gateDialerCloseChans[FSNo] = dialerCloseChan

	logger.Info("FSManager gateDialerCloseChans now len=%v", len(this.gateDialerCloseChans))
}

func (this *FSManager) setGSDialerCloseChans(FSNo int, dialerCloseChan chan struct{}) {
	this.gsDialerCloseChans[FSNo] = dialerCloseChan

	logger.Info("FSManager gsDialerCloseChans now len=%v", len(this.gsDialerCloseChans))
}

// 连接战斗服
func (this *FSManager) connectFs(session *FSSession) chan struct{} {
	context := &nw.Context{
		SessionCreator: GetFSSessionCreator(session),
		Splitter:       pbgt.Split,
		ReadBufferSize: 1024 * 1024 * 2,
		ChanSize:       1024 * 8,
	}
	addr := fmt.Sprintf("%s:%d", session.Host, session.Port)
	logger.Info("FSManager.DialEx, addr=%v, FSNo=%v", addr, session.FSNo)
	return tcpclient.DialEx(addr, context, 3*time.Second)
}

// 重载动态战斗跨服组 间隔5秒
func (this *FSManager) reloadDynamicCrossFightServer() {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	this.reloadCfs()
	for {
		select {
		case <-ticker.C:
			this.reloadCfs()
		}
	}
}

func (this *FSManager) reloadCfs() {

	cfsInfo, err := modelCross.GetCrossFightServerInfoModel().GetAllCrossFightServerListAndServer()
	if err != nil {
		logger.Error("reloadCfs db error: %v", err)
		return
	}

	linkServer := make(map[int]map[int]bool)
	for _,v := range cfsInfo{
		if _,ok := linkServer[v.Id];!ok{
			linkServer[v.Id] = make(map[int]bool)
		}
		if v.ServerId>0 {
			linkServer[v.Id][v.ServerId] = true
		}
	}

	logger.Info("reload动态战斗服链接服务器 ...... %v", linkServer)
	for _, cInfo := range cfsInfo {
		CFSNo := cInfo.Id
		gateSession := this.gateSessions[CFSNo]
		if gateSession != nil {
			// 若地址或端口有变化，则断开已有连接，并连接新地址
			if gateSession.Host != cInfo.Host || gateSession.Port != cInfo.GatePort {
				gateSession.Close()
				if this.gateDialerCloseChans[CFSNo] != nil {
					close(this.gateDialerCloseChans[CFSNo])
					this.gateDialerCloseChans[CFSNo] = nil
				}

				gateSession.Host = cInfo.Host
				gateSession.Port = cInfo.GatePort
				logger.Info("已有动态战斗服gate地址发生变化，重新连接 %v->%v", fmt.Sprintf("%s:%d", gateSession.Host, gateSession.Port), fmt.Sprintf("%s:%d", cInfo.Host, cInfo.GatePort))
				this.setGateDialerCloseChans(CFSNo, this.connectFs(gateSession))
			}
		} else { // 发起新连接
			gateSession = NewFSSession(CFSNo, cInfo.Host, cInfo.GatePort, false)
			this.gateSessions[CFSNo] = gateSession
			this.setGateDialerCloseChans(CFSNo, this.connectFs(gateSession))
		}

		gsSession := this.gsSessions[CFSNo]
		if gsSession != nil {
			// 若地址或端口有变化，则断开已有连接，并连接新地址
			if gsSession.Host != cInfo.Host || gsSession.Port != cInfo.GsPort {
				gsSession.Close()
				if this.gsDialerCloseChans[CFSNo] != nil {
					close(this.gsDialerCloseChans[CFSNo])
					this.gsDialerCloseChans[CFSNo] = nil
				}

				gsSession.Host = cInfo.Host
				gsSession.Port = cInfo.GatePort
				logger.Info("已有动态战斗服gs地址发生变化，重新连接 %v->%v", fmt.Sprintf("%s:%d", gsSession.Host, gsSession.Port), fmt.Sprintf("%s:%d", cInfo.Host, cInfo.GsPort))
				this.setGSDialerCloseChans(CFSNo, this.connectFs(gsSession))
			}
		} else { // 发起新连接
			gsSession = NewFSSession(CFSNo, cInfo.Host, cInfo.GsPort, true)
			this.gsSessions[CFSNo] = gsSession
			this.setGSDialerCloseChans(CFSNo, this.connectFs(gsSession))
		}
	}

	for k,v := range linkServer{
		this.gsSessions[k].SetLinkServers(v)
	}
}

func (this *FSManager) Stop() {
	for _, session := range this.gateSessions {
		conn := session.Conn
		if conn != nil {
			conn.Close()
			conn.Wait()
		}
	}
	for _, v := range this.gateDialerCloseChans {
		close(v)
	}

	for _, session := range this.gsSessions {
		conn := session.Conn
		if conn != nil {
			conn.Close()
			conn.Wait()
		}
	}
	for _, v := range this.gsDialerCloseChans {
		close(v)
	}
	logger.Info("nw_fs stop")
}

// 获取gs连接
func (this *FSManager) GetGSSession(fsNo int) *FSSession {
	return this.gsSessions[fsNo]
}

func (this *FSManager) RouteGateMessageToFs(crossServerId int, data []byte) error {

	fsSession := this.gateSessions[crossServerId]
	if fsSession == nil {
		return errors.New("fightserver connect not fount")
	}
	return fsSession.RoutoGateMsgTofs(data)
}

func (this *FSManager) RouteGsMessageToFs(crossServerId int, data []byte) error {

	fsSession := this.gsSessions[crossServerId]
	if fsSession == nil {
		return errors.New("fightserver connect not fount")
	}
	return fsSession.RoutoGsMsgTofs(data)
}

func (this *FSManager) RouteGsMessageCallToFs(crossServerId int, req nw.ProtoMessage, resp nw.ProtoMessage) error {

	fsSession := this.gsSessions[crossServerId]
	if fsSession == nil {
		return errors.New("fightserver connect not fount")
	}
	return fsSession.RouteGsMsgCallFs(req, resp)
}
