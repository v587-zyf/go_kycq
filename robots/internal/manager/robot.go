package manager

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw/wsclient"
	"cqserver/robots/conf"
	"fmt"
	"net/url"
	"strconv"
	"sync"
	"time"

	"cqserver/golibs/nw"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

var idAllocator = util.NewUint32IdAllocator(1)

type Robot struct {
	Id                uint32
	session           *ClientSession
	Messages          chan util.Message
	done              chan struct{}
	wg                sync.WaitGroup
	once              sync.Once
	lastConnectTime   int64
	status            int32
	inited            bool
	lastEnterGameTime int64

	sceneT             *gamedb.SceneConf
	inFightSceneStatus int //0:初始，1：进入中 2：战斗场景中

	openId      string
	serverId    int32
	user        *pb.UserLoginInfo
	initDataPb  []nw.ProtoMessage
	guildStatue int //0：初始，1：获取门派数据中，2：门派数据准备完成
	guildInfo   *pb.GuildInfo
}

func NewRobot(openId string, serverId int) *Robot {
	robot := &Robot{
		Id:       idAllocator.Get(),
		Messages: make(chan util.Message, 200),
		done:     make(chan struct{}),
		openId:   openId,
		serverId: int32(serverId),
	}
	return robot
}

func ToUrlValues(param map[string]interface{}) url.Values {
	values := make(url.Values)
	for k, v := range param {
		values.Add(k, fmt.Sprintf("%v", v))
	}
	return values
}

func (this *Robot) Start() {

	robot := this
	robot.wg.Add(1)
	ticker := time.NewTicker(time.Millisecond * 160)
	defer ticker.Stop()
	for {
		select {
		case msg := <-this.Messages:
			msg.Handle()
			msg.Done()
		case <-ticker.C:
			robot.RunAI()
		case <-robot.done:
			logger.Info("robot done come:%s", robot.openId)
			goto exit

		}
	}
exit:
	this.Stop()
}

func (this *Robot) Stop() {
	logger.Info("关闭机器人-------------------")
	this.once.Do(func() {
		close(this.done)
	})
	if this.session != nil {
		this.session.Conn.Close()
	}
}

func (this *Robot) RunAI() {
	if this.status == conf.STATUS_NONE {
		this.ConnectLs()
		return
	}
	if this.status == conf.STATUS_CONNECTED {
		this.SendMessage(0, this.MakeMsg(pb.CmdEnterGameReqId, ""))
		return
	}
	if this.status == conf.STATUS_ENTERGAME_ING { //10秒进不了，重进
		now := time.Now().Unix()
		if now > this.lastEnterGameTime+100 {
			this.status = conf.STATUS_NONE
			logger.Info("%d秒进不了，重进:%s", 10, this.openId)
			return
		}
	}
	if this.status == conf.STATUS_ENTERGAME_DONE {
		if this.user.NickName == "" {
			this.CreateUser()
		} else {
			this.status = conf.STATUS_IN_GAME
		}
	}
	//if len(this.initDataPb) > 0 {
	//	msg := this.initDataPb[0]
	//	this.initDataPb = this.initDataPb[1:len(this.initDataPb)]
	//	this.SendMessage(0, msg)
	//	logger.Info("创建角色中：%v,剩余角色数据：%v", msg, len(this.initDataPb))
	//	return
	//}

	if this.status == conf.STATUS_IN_GAME {
		//门派判断
		//if *fightType == constFight.FIGHT_TYPE_SHABAKE {
		//	if this.guildStatue == 0 {
		//		this.guildStatue = 1
		//		this.SendMessage(0, &pb.GuildLoadInfoReq{})
		//	}
		//} else {
		//	this.guildStatue = 2
		//}
		//this.lockTest()

		//if this.inFightSceneStatus == 0 {
		//	//发送进入场景
		//	this.inFightSceneTest()
		//} else if this.inFightSceneStatus == 2 {
		//	this.sceneAction()
		//}
		this.MakeMsg(0, "")
	}

}

//把cmdId传过来了----
func (this *Robot) SendMessage(transId uint32, msg nw.ProtoMessage) error {
	if this.session == nil {
		return nil
	}
	rb, err := pb.Marshal(pb.GetCmdIdFromType(msg), transId, msg)
	if err != nil {
		return err
	}
	_, err = this.session.Conn.Write(rb)
	return err
}

func (this *Robot) Write(data []byte) {
	if this.session != nil {
		this.session.Conn.Write(data)
	}
}

func (this *Robot) IsConnected() bool {
	return this.session != nil
}

func (this *Robot) ConnectLs() error {
	if this.status != conf.STATUS_NONE {
		return nil
	}
	now := time.Now().Unix()
	if now-this.lastConnectTime < 3 {
		return nil
	}
	this.lastConnectTime = now
	context := &nw.Context{
		SessionCreator: func(conn nw.Conn) nw.Session {
			sess := NewClientSession(conn)
			sess.Robot = this
			this.session = sess
			return sess
		},
		Splitter: pb.Split,
		ChanSize: 1024, //shs
	}
	var err error
	_, err = wsclient.Dial(m.GetAddr(), context)
	if err != nil {
		logger.Info("Connect fail:%v", err)
	} else {
		logger.Info("connect success")
	}
	return err
}

func (this *Robot) CreateUser() {
	heroCfg := conf.Conf.Create["hero"+strconv.Itoa(1)]
	this.SendMessage(pb.CmdCreateUserReqId,
		this.MakeMsg(pb.CmdCreateUserReqId, fmt.Sprintf(`%d,%d`, common.Interface2Int(heroCfg["sex"]), common.Interface2Int(heroCfg["job"]))))
	this.status = conf.STATUS_CREATE_ROLE_ING
	logger.Info("send CreateUser done:%s", this.openId)
}
