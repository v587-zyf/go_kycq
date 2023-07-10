package gm

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/prop"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type GmManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewGmManager(module managersI.IModule) *GmManager {
	return &GmManager{IModule: module}
}

func (this *GmManager) GmChatCode(user *objs.User, codes string) string {

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
			return ""
		}
	}

	logger.Debug("接收到客户端发送来的gm指令：%v", codes)
	codeSlices := strings.Split(codes, "#")
	if len(codeSlices) < 2 {
		return gamedb.ERRPARAM.Message
	}
	cmd := strings.TrimSpace(strings.ToLower(codeSlices[1]))
	switch cmd {
	case "property":
		if len(codeSlices) <= 2 {
			user.GmProperty = nil
			return gamedb.ERRPARAM.Message
		}
		property := make(gamedb.IntMap)
		property.Decode(codeSlices[2])
		props := prop.NewProp()
		props.Add(property)
		props.Calc(user.Heros[constUser.USER_HERO_MAIN_INDEX].Job)
		user.GmProperty = props
		this.GetFight().UpdateUserInfoToFight(user, map[int]bool{constUser.USER_HERO_MAIN_INDEX: true}, true)
	case "bagclear":
		this.GetBag().BagClear(user)
	case "recharge":
		rechargeId, _ := strconv.Atoi(codeSlices[2])
		this.GetRecharge().TestPay(user, rechargeId)
	case "addbuff", "showfightprop":
		msg := this.GetFight().Gm(user, codes)
		return msg
	case "showprop":
		heroIndex := 0
		if len(codeSlices) > 2 {
			heroIndex, _ = strconv.Atoi(codeSlices[2])
		}
		isAll := true
		for _, v := range constUser.USER_HERO_INDEX {
			if v == heroIndex {
				isAll = false
			}
		}
		propMsg := ""
		for _, v := range user.Heros {
			if isAll || v.Index == heroIndex {
				propStr := v.Prop.AnalyzePercent()
				propMsg += fmt.Sprintf("昵称:%v,玩家Id:%v，详情：%s \n", v.Name, propStr)
			}
		}
		return propMsg
	case "atlas":
		this.GetAtlas().AtlasGm(user, codeSlices[2])
	case "clear":
		mo := "all"
		if len(codeSlices) > 2 {
			mo = codeSlices[2]
		}

		if mo == "all" || mo == "cutskill" {
			user.CutTreasureUseEndCd = 0
			this.GetUserManager().SendMessage(user, &pb.CutTreasureUseAck{UseTime: int32(time.Now().Unix()), CdEndTime: int32(0)}, true)
		}
		if mo == "all" || mo == "fit" {
			user.Fit.CdStart = 0
			user.Fit.CdEnd = 0
			ack := &pb.FitEnterAck{}
			ack.CdStartTime = int32(user.Fit.CdStart)
			ack.CdEndTime = int32(user.Fit.CdEnd)
			this.GetUserManager().SendMessage(user, ack, true)
		}
	case "removeitem":
		op := ophelper.NewOpBagHelperDefault(constBag.OpTypeDebugAddGoods)
		itemId, _ := strconv.Atoi(codeSlices[2])
		this.GetBag().RemoveAllByItemId(user, op, itemId)
		this.GetUserManager().SendItemChangeNtf(user, op)
	case "changeguild":
		guildName := codeSlices[2]
		modelGame.GetGuildModel().GetAllGuildInfoByGuildName(guildName)
		data, _ := modelGame.GetGuildModel().GetAllGuildInfoByGuildName(guildName)
		guildId := 0
		if len(data) > 0 {
			guildId = data[0].GuildId
		}
		this.GetUserManager().SendMessage(user, &pb.ChatMessageNtf{
			Type: int32(pb.CHATTYPE_PRIVATE),
			ToId: int32(user.Id),
			Msg:  fmt.Sprintf(`%s#%d`, "changeguildgm", guildId),
		}, true)
	case "reload":
		this.ReloadGameDb()
	}
	return codes + " 操作成功"
}
