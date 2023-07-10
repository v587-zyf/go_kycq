package shabakeCross

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"encoding/json"
	"time"
)

func (this *CrossShaBaKe) GetShaBakeOpenDay(serverId int) int {
	return this.GetSystem().GetMergerServerOpenDaysByServerId(serverId)
}

func (this *CrossShaBaKe) LoadCross(user *objs.User, ack *pb.ShaBaKeInfoCrossAck) error {
	data := rmodel.Shabake.GetCrossShabakeFirstGuildInfo()
	if data == "" {
		return nil
	}
	infos := &pb.ShaBaKeInfoCrossAck{}
	err := json.Unmarshal([]byte(data), infos)
	if err != nil {
		logger.Error("跨服沙巴克 load unMarshal err:%v", err)
		return nil
	}
	ack.IsEnd = int32(rmodel.Shabake.GetCrossShaBakeIsEnd(base.Conf.ServerId))
	ack.WinGuildUserInfo = infos.WinGuildUserInfo
	ack.FirstGuildName = infos.FirstGuildName
	return nil
}

func (this *CrossShaBaKe) CrossShaBakeOpenOrCloseNtf(isOpen bool) error {
	this.BroadcastAll(&pb.CrossShabakeOpenNtf{IsOpen: isOpen})
	return nil
}

func (this *CrossShaBaKe) JudgeCrossIsOpen(user *objs.User) error {

	state := rmodel.Shabake.GetCrossShaBakeIsEnd(base.Conf.ServerId)
	if state == 1 {
		return gamedb.ERRACTIVITYCLOSE
	}
	if this.GetSystem().GetCrossFightServerId() <= 0 {
		return gamedb.ERRACTIVITYNOTOPEN
	}
	limitedOpenDay := gamedb.GetConf().KuafushabakeTime1
	openDay := this.GetShaBakeOpenDay(user.ServerId)
	//开服天数限制
	if limitedOpenDay > openDay {
		return gamedb.ERRACTIVITYNOTOPEN
	}

	limitedWeekDay := gamedb.GetConf().KuafushabakeTime2
	//周几限制
	if limitedWeekDay != nil && len(limitedWeekDay) > 0 {
		t := int(time.Now().Weekday())
		if t == 0 {
			t = 7
		}
		//周几限制
		canIn := false
		for _, v := range limitedWeekDay {
			if v == t {
				canIn = true
				break
			}
		}
		if !canIn {
			return gamedb.ERRACTIVITYNOTOPEN
		}
	}

	limitedTimes := gamedb.GetConf().KuafushabakeTime3
	if limitedTimes == nil {
		return gamedb.ERRACTIVITYNOTOPEN
	}
	if len(limitedTimes) < 2 {
		return gamedb.ERRACTIVITYNOTOPEN
	}

	nowTime := time.Now()
	openTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), limitedTimes[0].Hour, limitedTimes[0].Minute, limitedTimes[0].Second, 0, time.Local)
	endTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), limitedTimes[1].Hour, limitedTimes[1].Minute, limitedTimes[1].Second, 0, time.Local)

	//每天时间段限制
	if nowTime.Unix() < openTime.Unix() || nowTime.Unix() > endTime.Unix() {
		logger.Error("跨服沙巴克不开启时间段内")
		return gamedb.ERRACTIVITYNOTOPEN
	}
	return nil
}

func (this *CrossShaBaKe) SendFirstGuildInfoToCcs(guildId, benFuShaBake int) {
	logger.Info("沙巴克获胜门派展示信息存储 guildId:%v, benFuShaBake:%v", guildId, benFuShaBake)
	guildInfo := this.GetGuild().GetGuildInfo(guildId)
	if guildInfo == nil {
		logger.Error("GetGuildInfo nil guildId:%v", guildId)
		return
	}
	crossFsId := this.GetSystem().GetServerIndexCrossFsId(base.Conf.ServerId)
	users := this.GetGuild().GetGuildHuiAndFuHuiUserIds(guildId)
	ntf := &pbserver.GsToCcsBackGuildInfoNtf{}
	ntf.CrossFsId = int32(crossFsId)
	ntf.BenFuShaBaKe = int32(benFuShaBake)
	winGuildUserInfo := this.BuildBackGuildInfo(users)

	kyInfo := make([]int, 0)
	kyInfo = append(kyInfo, guildInfo.ChairmanId, pb.GUILDPOSITION_HUIZHANG)
	infos := make([]*pb.Info, 0)
	data := &pb.ShaBaKeInfoCrossAck{}
	data.FirstGuildName = guildInfo.GuildName
	for _, info := range winGuildUserInfo {
		infos = append(infos, &pb.Info{NickName: info.NickName, Sex: info.Sex, Job: info.Job, Position: info.Position,
			Display: &pb.Display{ClothItemId: info.Display.ClothItemId, ClothType: info.Display.ClothType, WeaponItemId: info.Display.WeaponItemId,
				WeaponType: info.Display.WeaponType, WingId: info.Display.WingId, MagicCircleLvId: info.Display.MagicCircleLvId,
			}})
	}
	data.WinGuildUserInfo = infos
	data.IsEnd = 1
	data.ServerId = int32(base.Conf.ServerId)
	datas, err := json.Marshal(data)
	if err != nil {
		logger.Error("SetFirstGuildInfo err:%v", err)
		return
	}
	userInfo := this.GetUserManager().GetAllUserInfoIncludeOfflineUser(guildInfo.ChairmanId)
	if userInfo != nil {
		kyEvent.ShaBake(userInfo, guildId, kyInfo)
	}
	if crossFsId <= 0 {
		rmodel.Shabake.SetShabakeFirstGuildInfo(string(datas))
		return
	}
	ntf.FirstGuildInfo = string(datas)
	this.SendMsgToCCS(0, ntf)
}

func (this *CrossShaBaKe) BuildBackGuildInfo(users []int) []*pbserver.Info {

	infos := make([]*pbserver.Info, 0)
	for i, j := 0, len(users); i < j; i += 2 {
		userId := users[i]
		position := users[i+1]
		userInfo := this.GetUserManager().BuilderBrieUserInfo(userId)
		if userInfo != nil {
			infos = append(infos, &pbserver.Info{NickName: userInfo.Name, Sex: userInfo.Sex, Job: userInfo.Job, Position: int32(position),
				Display: &pbserver.Display{
					ClothItemId:     userInfo.Display[constUser.USER_HERO_MAIN_INDEX].ClothItemId,
					ClothType:       userInfo.Display[constUser.USER_HERO_MAIN_INDEX].ClothType,
					WeaponItemId:    userInfo.Display[constUser.USER_HERO_MAIN_INDEX].WeaponItemId,
					WeaponType:      userInfo.Display[constUser.USER_HERO_MAIN_INDEX].WeaponType,
					WingId:          userInfo.Display[constUser.USER_HERO_MAIN_INDEX].WingId,
					MagicCircleLvId: userInfo.Display[constUser.USER_HERO_MAIN_INDEX].MagicCircleLvId,
				}})
		}
	}
	return infos
}

func (this *CrossShaBaKe) SetFirstGuildInfo(msg *pbserver.CcsToGsBroadShaBakeFirstGuildInfo) {
	logger.Info("SetFirstGuildInfo msg:%v", msg)
	if msg.BenFuShaBaKe == 1 {
		rmodel.Shabake.SetShabakeFirstGuildInfo(msg.FirstGuildInfo)
		return
	}
	rmodel.Shabake.SetCrossShabakeFirstGuildInfo(msg.FirstGuildInfo)
}
