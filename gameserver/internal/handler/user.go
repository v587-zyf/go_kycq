package handler

import (
	"cqserver/gamelibs/errex"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constAuction"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/gater"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"github.com/astaxie/beego/logs"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdPingReqId, HandlePingReq)
	pb.Register(pb.CmdEnterGameReqId, HandleEnterGameReq)
	pb.Register(pb.CmdCreateUserReqId, HandleCreateUserReq)
	pb.Register(pb.CmdCreateHeroReqId, HandleCreateHeroReq)
	pb.Register(pb.CmdDebugAddGoodsReqId, HandleDebugAddGoodsReq)
	pb.Register(pb.CmdRandNameReqId, HandlerRandNameReq)
	pb.Register(pb.CmdChangeFightModelReqId, HandlerChangeChangeFightModelReq)
	pb.Register(pb.CmdChangeHeroNameReqId, HandlerChangeHeroNameReq)
	pb.Register(pb.CmdVipCustomerReqId, HandlerVipCustomerReq)
	pb.Register(pb.CmdUserSubscribeReqId, HandlerUserSubscribeReq)
}

func HandlePingReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	return &pb.PingAck{
		Ts: common.GetNowMillisecond(),
	}, nil, nil
}

func HandleEnterGameReq(conn nw.Conn, p interface{}) (ack nw.ProtoMessage, op pb.OpGoodsHelper, enterErr error) {

	defer func() {
		if err := recover(); err != nil {
			logger.Error("SafeRun panic:", time.Now(), err, string(debug.Stack()))
			if base.Conf.Sandbox {
				enterErr = gamedb.ERRUNKNOW.CloneWithMsg(string(debug.Stack()))
			} else {
				enterErr = gamedb.ERRUNKNOW
			}
		}
	}()

	req := p.(*pb.EnterGameReq)
	logger.Info("user enterGame start openId:%v, origin:%d,serverId:%v,ip:%v", req.OpenId, req.Origin, req.ServerId, conn.RemoteAddr())
	if !m.System.ServerIdInLocalServer(int(req.ServerId)) {
		return nil, nil, gamedb.ERRPARAM
	}
	ipArr := strings.Split(conn.RemoteAddr().String(), ":")
	user, err := m.UserManager.LoadUser(req.OpenId, int(req.Channel), ipArr[0], int(req.ServerId), req.Origin, req.DeviceId)
	if err != nil {
		logger.Error("HandleEnterGameReq:err:%v", err)
		return nil, nil, err
	}

	session := conn.GetSession().((*managers.ClientSession))
	user.SetConn(conn)
	user.GateSessionId = session.Conn.(*gater.ClientConn).GetGateClientSessionId()
	session.User = user
	ackMsg := builder.BuildEnterGameAck(user, req.LoginKey,
		m.System.GetServerOpenDaysByServerId(user.ServerId),
		int(m.System.GetServerOpenTimeByServerId(user.ServerId).Unix()),
		m.System.GetServerOpenDaysByServerId(user.ServerId),
		m.System.GetMergerServerOpenDaysByServerId(user.ServerId),
		m.System.GetServerName(base.Conf.ServerId))
	closeFunc := m.System.GetFuncState()
	ackMsg.User.MiJiInfos = m.GetMiJi().GetMiJiInfos(user)
	ackMsg.CloseFuncIds = common.ConvertIntSlice2Int32Slice(closeFunc)
	ackMsg.User.WorldBossInfo = m.WorldBoss.GetWorldBossInfo()
	ackMsg.User.CrossChallengeIsApply = m.GetChallenge().IsApplyChallenge(user)
	ackMsg.User.AnnouncementInfos = m.GetAnnouncement().GetAnnouncement(user)
	ackMsg.User.AncientTreasureInfo = m.GetAncientTreasure().BuildAncientTreasureInfo(user)
	ackMsg.CrossBriefServerInfo = m.GetSystem().GetCrossServerBriefUserInfo()
	_, chatBanTime := m.UserManager.CheckBan(user, constUser.BAN_TYPE_CHAT)
	ackMsg.User.ChatBanTime = int32(chatBanTime)
	ackMsg.User.Subscribe = m.UserManager.GetSubscribe(user.Id)
	logger.Info("user enterGame ok openId:%v,userId:%v,nickName:%v,sessionId:%v,mem:%v", req.OpenId, user.Id, user.NickName, user.GateSessionId, unsafe.Sizeof(*user))

	m.UserManager.SendDisplay(user)
	m.Recharge.Online(user)
	ack = ackMsg
	return ack, op, enterErr
}

func HandleCreateUserReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.CreateUserReq)
	user := conn.GetSession().((*managers.ClientSession)).User

	if !pb.SEX_MAP[int(req.Sex)] {
		return nil, nil, gamedb.ERRPARAM
	}

	if !pb.JOB_MAP[int(req.Job)] {
		return nil, nil, gamedb.ERRPARAM
	}
	logger.Info("create role 玩家：%v-%v,当前武将数：%v,sex:%v,job:%v", user.Id, user.NickName, len(user.Heros), req.Sex, req.Job)
	err := m.UserManager.CreateRole(user, req.NickName, req.Avatar, int(req.Sex), int(req.Job))
	if err != nil {
		ei, ok := err.(*errex.ErrorItem)
		if !ok {
			ei = gamedb.ERRUNKNOW
		}
		return &pb.CreateUserAck{FailReason: ei.Message}, nil, nil
	}
	body := builder.BuildCreateUserAck(user, m.System.GetServerOpenDaysByServerId(user.ServerId))
	body.User.WorldBossInfo = m.WorldBoss.GetWorldBossInfo()
	m.UserManager.SendDisplay(user)
	return body, nil, nil
}

func HandleCreateHeroReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.CreateHeroReq)
	user := conn.GetSession().((*managers.ClientSession)).User
	logger.Info("create hero 玩家：%v-%v,当前武将数：%v,sex:%v,job:%v", user.Id, user.NickName, len(user.Heros), req.Sex, req.Job)
	if !pb.SEX_MAP[int(req.Sex)] {
		return nil, nil, gamedb.ERRPARAM
	}

	if !pb.JOB_MAP[int(req.Job)] {
		return nil, nil, gamedb.ERRPARAM
	}

	if len(user.NickName) == 0 || len(user.Heros) <= 0 {
		logger.Error("玩家角色角色信息还未创建", user.NickName, len(user.Heros))
		return nil, nil, gamedb.ERRUNLOGIN
	}

	index, err := m.UserManager.CreateHero(user, int(req.Sex), int(req.Job))
	if err != nil {
		return nil, nil, err
	}
	msg := &pb.CreateHeroAck{
		Hero: builder.BuildHeroInfo(user.Heros[index]),
	}
	//此线程只为控制消息先后顺序,保证客户端先收到创建英雄的消息，再收到武将进入战斗的消息
	go func() {
		time.Sleep(100 * time.Millisecond)
		m.ClientManager.DispatchEvent(user.Id, nil, func(userId int, user *objs.User, data interface{}) {
			if user != nil {
				if index != constUser.USER_HERO_MAIN_INDEX {
					m.GetFight().NewHeroIntoFight(user, index)
				}
			}
		})
	}()
	m.UserManager.SendDisplay(user)
	return msg, nil, nil

}

func HandleDebugAddGoodsReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.DebugAddGoodsReq)
	user := conn.GetSession().(*managers.ClientSession).User
	if !base.Conf.GmSwitch {
		all := modelCross.GetWhiteListDbModel().Getall()
		inWhite := false
		for _, v := range all {
			if v.Valtype == 2 && user.OpenId == v.Value {
				inWhite = true
				break
			}
		}
		if !inWhite {
			return nil, nil, nil
		}
	}

	opGoodsHelper := ophelper.NewOpBagHelperDefault(constBag.OpTypeDebugAddGoods)
	var err error
	logger.Info("接收到客户端命令：id:%v,count:%v,arg:%v", user.IdName(), user.Ip, req.Id, req.Count, req.Args)
	if len(req.Id) != len(req.Count) {
		return nil, nil, gamedb.ERRPARAM
	}
	for k, itemId := range req.Id {
		if itemId < 0 {
			switch itemId {
			case -1:
				//user.MainLineTask.TaskId = int(req.Count[0])
				//user.MainLineTask.Process = 0
				//m.Task.UpdateTaskProcess(user, true, true)
				m.UserManager.Subscribe(user, pb.SUBSCRIBE_HOOK)
			case -2:
				m.GetChat().ChatSendReq(user, &pb.ChatSendReq{
					Type: pb.CHATTYPE_WORLD,
					Msg:  "1111111111111111",
					ToId: 0,
				})
			case -3: //累计充值
				user.RechargeAll += int(req.Count[0])
				m.FirstRecharge.UpdateFirstRechargeStatus(user)
			case -4:
				//users, err := modelGame.GetUserModel().SearchActivityUsers(86400 * 5)
				user.DailyTask.DayExp += int(req.Count[0])
				user.DailyTask.WeekExp += int(req.Count[0])
				user.DailyTask.ResourcesBackExp += int(req.Count[0])
			case -5:
				op := ophelper.NewOpBagHelperDefault(constBag.OpTypeTask)
				m.Task.TaskDone(user, op)
			case -6:
				//err := m.Challenge.SetApplyUserInfo(user, &pb.ApplyChallengeAck{})
				//if err != nil {
				//	logger.Error("1111 err:%v", err)
				//}
				m.ReloadGameDb()

			case -7:
				user.MainLineTask.TaskId = int(req.Count[0])
				m.Task.UpdateTaskProcess(user, true, true)
			case -8:
				//邮件测试命令  count:邮件类型
				returnItem := gamedb.ItemInfos{&gamedb.ItemInfo{ItemId: pb.ITEMID_INGOT, Count: 100}}
				m.GetMail().SendSystemMailWithItemInfos(int(user.Id), int(req.Count[0]), []string{strconv.Itoa(1)}, returnItem)
			case -9:
				if user.Heros[int(req.Count[0])] != nil {
					user.Heros[int(req.Count[0])].HolyAllPoint += 100
				}
			case -10:
				if user.GuildData.NowGuildId > 0 {
					m.GetAuction().DropItemToGuildAuction(user.GuildData.NowGuildId, int(req.Count[0]), 5, constAuction.DropWorldLeader, []int{user.Id})
				}

			case -11:
				gamedb.GetDailyTaskDailyTaskCfgByTypeTest()
			case -12:
				//经验池 玩家等级 Gm -12,第几个hero|-12,指定等级
				if user.Heros[int(req.Count[0])] != nil {
					user.Heros[int(req.Count[0])].ExpLvl = int(req.Count[1])
				}
				ack := &pb.ExpPoolUpGradeAck{}
				ack.HeroIndex = int32(req.Count[0])
				ack.Lvl = int32(req.Count[1])
				m.GetUserManager().SendMessage(user, ack, true)
				user.Dirty = true
				return &pb.DebugAddGoodsAck{Result: 1}, opGoodsHelper, err
			case -13:
				//经验池 玩家等级 Gm
				for heroIndex, userInfo := range user.Heros {
					userInfo.ExpLvl = int(req.Count[0])
					ack := &pb.ExpPoolUpGradeAck{}
					ack.HeroIndex = int32(heroIndex)
					ack.Lvl = int32(req.Count[0])
					m.GetUserManager().SendMessage(user, ack, true)
				}
				user.Dirty = true

			case -14:
				//补充擂台赛假人
				//-14,机器人起始id|-14,serverId|-14,增加机器人的数量|-14,跨服组id
				//-14,1|-14,1|-14,20|-14,100000
				if len(req.Count) < 4 {
					return &pb.DebugAddGoodsAck{Result: 1}, opGoodsHelper, err
				}
				robotId := int(req.Count[0])
				serverId := int(req.Count[1])
				addNum := int(req.Count[2])
				crossFsId := int(req.Count[3])
				for i := robotId; i <= robotId+addNum; i++ {
					robotCfg := gamedb.GetCrossArenaRobotCrossArenaRobotCfg(i)
					if robotCfg == nil {
						logger.Error("GetCrossArenaRobotCrossArenaRobotCfg robotId:%v  nil 假人不足", robotId)
						break
					}
					info := &modelCross.Challenge{UserId: -robotCfg.Id, NickName: robotCfg.Name, Avatar: robotCfg.Icon, ServerId: serverId, Combat: int64(robotCfg.Combat), ExpireTime: time.Now().Unix() + 86400*14, Round: 0, CrossFsId: crossFsId, Season: m.GetChallenge().WeekByDate(time.Now().AddDate(0, 0, -1))}
					err := modelCross.GetChallengeModel().DbMap().Insert(info)
					if err != nil {
						logger.Error("insert challenge data: %v err: %v", *info, err)
					}
				}
				return &pb.DebugAddGoodsAck{Result: 1}, opGoodsHelper, err
			case -15:
				season := m.GetChallenge().WeekByDate(time.Now().AddDate(0, 0, -1))
				userInfo, err := modelCross.GetChallengeModel().GetAllServerApplyUserInfo(100000, -4, season)
				if err != nil {
					logger.Error("1111111  GetAllServerApplyUserInfo crossFsId:%v, userId:%v err:%v", 100000, -4, err)
				}
				logs.Info("%v", userInfo)
				maxRobotId, err := modelCross.GetChallengeModel().MaxRobotId(0, 10000, season)
				logs.Info("max:%v   :%v", maxRobotId, -maxRobotId)
			case -16:
				user.StageId = int(req.Count[0])
				user.StageWave = 0
				ntf := &pb.StageFightEndNtf{StageId: int32(user.StageId), Wave: int32(user.StageWave), OnlyUpdate: true, Result: pb.RESULTFLAG_SUCCESS}
				m.GetUserManager().SendMessage(user, ntf, true)
				user.Dirty = true
			case -18:
				m.GetLottery().OpenAward()
			}

			//特殊命令
		} else {
			if req.Count[k] > 0 {
				err = m.Bag.Add(user, opGoodsHelper, int(itemId), int(req.Count[k]))
				//source := gamedb.GetItemSourceByStageId(user.NickName, 401)
				//err = m.Bag.AddItem(user, opGoodsHelper, int(itemId), int(req.Count[k]), nil)
			} else {
				//err = m.Bag.Remove(user, opGoodsHelper, int(itemId), int(-req.Count[k]))
				m.Bag.ItemUse(user, -1, int(itemId), int(-req.Count[k]), opGoodsHelper)
			}
		}
	}

	return &pb.DebugAddGoodsAck{Result: 1}, opGoodsHelper, err
}

func HandlerRandNameReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.RandNameReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.RandNameAck{}
	err := m.UserManager.RandName(user, int(req.Sex), ack)
	if err != nil {
		return nil, nil, err
	}

	logger.Debug("HandlerRandNameReq ack is %v", ack)

	return ack, nil, nil
}

func HandlerChangeChangeFightModelReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ChangeFightModelReq)
	user := conn.GetSession().(*managers.ClientSession).User

	if user.FightModel != int(req.FightModel) {

		user.FightModel = int(req.FightModel)
		user.Dirty = true
		m.GetFight().UpdateUserFightModel(user)
	}

	return &pb.ChangeFightModelAck{FightModel: req.FightModel}, nil, nil
}

func HandlerChangeHeroNameReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ChangeHeroNameReq)
	user := conn.GetSession().(*managers.ClientSession).User

	err := m.UserManager.ChangeHeroName(user, int(req.HeroIndex), req.Name)
	if err != nil {
		return nil, nil, err
	}

	return &pb.ChangeHeroNameAck{HeroInfo: builder.BuildHeroInfo(user.Heros[int(req.HeroIndex)])}, nil, nil
}

func HandlerVipCustomerReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	user.VipCustomer = constUser.VIP_CUSTOMER_LOOK_YES
	user.Dirty = true
	return &pb.VipCustomerAck{Flag: true}, nil, nil
}

func HandlerUserSubscribeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.UserSubscribeReq)

	err := m.UserManager.Subscribe(user, int(req.SubscribeId))
	if err != nil {
		return nil, nil, err
	}
	return &pb.UserSubscribeAck{SubscribeId: req.SubscribeId}, nil, nil
}
