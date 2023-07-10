package manager

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/robots/conf"
	"fmt"
	"strconv"
	"strings"
)

func (this *Robot) initData() {
	this.initDataPb = append(this.initDataPb, this.addUserGoods())
	mainHeroMsg := this.createHeroInfo(constUser.USER_HERO_MAIN_INDEX)
	if len(mainHeroMsg) > 0 {
		this.initDataPb = append(this.initDataPb, mainHeroMsg...)
	}
	herosMsg := this.CreateHeros()
	if len(herosMsg) > 0 {
		this.initDataPb = append(this.initDataPb, herosMsg...)
	}
	this.initDataPb = append(this.initDataPb, this.addUserItems())
}

func (this *Robot) CreateHeros() []nw.ProtoMessage {
	msgs := make([]nw.ProtoMessage, 0)
	heroNum := common.Interface2Int(conf.Conf.Create["user"]["heroNum"])
	fmt.Println(this.user)
	if len(this.user.Heros) == heroNum {
		return msgs
	}
	for heroIndex := 2; heroIndex <= heroNum; heroIndex++ {
		heroCfg := conf.Conf.Create["hero"+strconv.Itoa(heroIndex)]
		logger.Info("heroCfg:%v", heroCfg)
		msgs = append(msgs, this.MakeMsg(pb.CmdCreateHeroReqId, fmt.Sprintf(`%d,%d`, common.Interface2Int32(heroCfg["sex"]), common.Interface2Int32(heroCfg["job"]))))
		createHeroMsg := this.createHeroInfo(heroIndex)
		if len(createHeroMsg) > 0 {
			msgs = append(msgs, createHeroMsg...)
		}
	}
	return msgs
}

func (this *Robot) createHeroInfo(heroIndex int) []nw.ProtoMessage {
	msgs := make([]nw.ProtoMessage, 0)
	heroCfg := conf.Conf.Create["hero"+strconv.Itoa(heroIndex)]

	//先加等级 (-12，第几个武将 | -12，等级)
	msgs = append(msgs, this.DebugAdd([]int32{-12, -12}, []int32{int32(heroIndex), common.Interface2Int32(heroCfg["lv"])}))

	//武器
	if common.Interface2Int32(heroCfg["arms"]) > 0 {
		msgs = append(msgs, this.DebugAdd([]int32{common.Interface2Int32(heroCfg["arms"])}, []int32{1}))
		msgs = append(msgs, this.MakeMsg(pb.CmdEquipChangeReqId, fmt.Sprintf(`%d,%d,%d`, heroIndex, -1, 0)))
	}

	//衣服
	if common.Interface2Int32(heroCfg["clothes"]) > 0 {
		msgs = append(msgs, this.DebugAdd([]int32{common.Interface2Int32(heroCfg["clothes"])}, []int32{1}))
		msgs = append(msgs, this.MakeMsg(pb.CmdEquipChangeReqId, fmt.Sprintf(`%d,%d,%d`, heroIndex, -1, 0)))
	}

	//时装武器
	if common.Interface2Int32(heroCfg["f_arms"]) > 0 {
		fid := common.Interface2Int32(heroCfg["f_arms"])
		msgs = append(msgs, this.DebugAdd([]int32{fid}, []int32{1}))
		msgs = append(msgs, this.MakeMsg(pb.CmdFashionUpLevelReqId, fmt.Sprintf(`%d,%d`, heroIndex, fid)))
		msgs = append(msgs, this.MakeMsg(pb.CmdFashionWearReqId, fmt.Sprintf(`%d,%d,%s`, heroIndex, fid, "true")))
	}

	//时装衣服
	if common.Interface2Int32(heroCfg["f_clothes"]) > 0 {
		fid := common.Interface2Int32(heroCfg["f_clothes"])
		msgs = append(msgs, this.DebugAdd([]int32{fid}, []int32{1}))
		msgs = append(msgs, this.MakeMsg(pb.CmdFashionUpLevelReqId, fmt.Sprintf(`%d,%d`, heroIndex, fid)))
		msgs = append(msgs, this.MakeMsg(pb.CmdFashionWearReqId, fmt.Sprintf(`%d,%d,%s`, heroIndex, fid, "true")))
	}

	//翅膀
	if common.Interface2Int32(heroCfg["wing"]) > 0 {
		wingId := common.Interface2Int(heroCfg["wing"])
		for i := 1; i <= wingId; i++ {
			cfg := gamedb.GetWingNewWingNewCfg(i)
			msgs = append(msgs, this.DebugAdd([]int32{int32(cfg.Consume.ItemId)}, []int32{int32(cfg.Consume.Count)}))
			msgs = append(msgs, this.MakeMsg(pb.CmdWingUseMaterialReqId, fmt.Sprintf(`%d`, heroIndex)))
		}
		msgs = append(msgs, this.MakeMsg(pb.CmdWingWearReqId, fmt.Sprintf(`%d,%d`, heroIndex, wingId)))
	}

	//法阵
	if common.Interface2Int32(heroCfg["magicCircle"]) > 0 {
		magicCircleId := common.Interface2Int(heroCfg["magicCircle"])
		magicCircleCfg := gamedb.GetMagicCircleLevelMagicCircleLevelCfg(magicCircleId)
		for i := 0; i < magicCircleCfg.Rank; i++ {
			lv := magicCircleCfg.Level
			switch {
			case i == 0:
				lv = 1
			case magicCircleCfg.Rank > 1:
				lv = 10
			}
			for j := 0; j < lv; j++ {
				lvCfg := gamedb.GetMagicCircleLvCfg(magicCircleCfg.Type, i, j)
				if lvCfg != nil && len(lvCfg.Item) > 0 {
					debugIdSlice, debugCountSlice := this.MakeAddItemsSlice(lvCfg.Item)
					msgs = append(msgs, this.DebugAdd(debugIdSlice, debugCountSlice))
					msgs = append(msgs, this.MakeMsg(pb.CmdMagicCircleUpLvReqId, fmt.Sprintf(`%d,%d`, heroIndex, lvCfg.Type)))
				}
			}
		}
		msgs = append(msgs, this.MakeMsg(pb.CmdMagicCircleChangeWearReqId, fmt.Sprintf(`%d,%d`, heroIndex, magicCircleId)))
	}

	//技能
	if len(common.Interface2Str(heroCfg["skill"])) > 0 {
		skillArr := strings.Split(common.Interface2Str(heroCfg["skill"]), ",")
		for _, skillId := range skillArr {
			lvConf := gamedb.GetSkillLvConf(common.Str2Int(skillId), 0)
			debugIdSlice, debugCountSlice := this.MakeAddItemsSlice(lvConf.Consume)
			msgs = append(msgs, this.DebugAdd(debugIdSlice, debugCountSlice))
			msgs = append(msgs, this.MakeMsg(pb.CmdSkillUpLvReqId, fmt.Sprintf(`%d,%s`, heroIndex, skillId)))
			msgs = append(msgs, this.MakeMsg(pb.CmdSkillChangeWearReqId, fmt.Sprintf(`%d,%s`, heroIndex, skillId)))
		}
	}
	return msgs
}

func (this *Robot) addUserGoods() nw.ProtoMessage {

	userCfg := conf.Conf.Create["user"]
	debugIdSlice := make([]int32, 0)
	debugCountSlice := make([]int32, 0)
	for key, val := range userCfg {
		switch key {
		case "gold":
			debugIdSlice = append(debugIdSlice, pb.ITEMID_GOLD)
			debugCountSlice = append(debugCountSlice, common.Interface2Int32(val))
		case "honour":
			debugIdSlice = append(debugIdSlice, pb.ITEMID_HONOR)
			debugCountSlice = append(debugCountSlice, common.Interface2Int32(val))
		case "ingot":
			debugIdSlice = append(debugIdSlice, pb.ITEMID_INGOT)
			debugCountSlice = append(debugCountSlice, common.Interface2Int32(val))
		case "vip":
			debugIdSlice = append(debugIdSlice, pb.ITEMID_VIP_LV)
			debugCountSlice = append(debugCountSlice, common.Interface2Int32(val))
		case "mainTask":
			debugIdSlice = append(debugIdSlice, -7)
			debugCountSlice = append(debugCountSlice, common.Interface2Int32(val))
		}
	}
	return this.DebugAdd(debugIdSlice, debugCountSlice)
}

func (this *Robot) addUserItems() nw.ProtoMessage {
	userCfg := conf.Conf.Create["user"]
	debugIdSlice := make([]int32, 0)
	debugCountSlice := make([]int32, 0)
	for key, val := range userCfg {
		switch key {
		case "addItem":
			if len(common.Interface2Str(val)) > 0 {
				items := strings.Split(common.Interface2Str(val), "|")
				for _, item := range items {
					itemArr := strings.Split(item, ",")
					debugIdSlice = append(debugIdSlice, common.Str2Int32(itemArr[0]))
					debugCountSlice = append(debugCountSlice, common.Str2Int32(itemArr[1]))
				}
			}
		}
	}
	return this.DebugAdd(debugIdSlice, debugCountSlice)
}

func (this *Robot) changeGuild() {
	userCfg := conf.Conf.Create["user"]
	if userCfg["guildName"] == nil {
		return
	}
	guildName := common.Interface2Str(userCfg["guildName"])
	if len(guildName) <= 0 {
		return
	}
	this.SendMessage(0, this.MakeMsg(pb.CmdChatSendReqId, fmt.Sprintf(`%d,%s,%d`, pb.CHATTYPE_WORLD, "GM#changeguild#"+guildName, 0)))
}

func (this *Robot) MakeAddItemsSlice(item gamedb.ItemInfos) ([]int32, []int32) {
	debugIdSlice := make([]int32, 0)
	debugCountSlice := make([]int32, 0)
	for _, info := range item {
		debugIdSlice = append(debugIdSlice, int32(info.ItemId))
		debugCountSlice = append(debugCountSlice, int32(info.Count))
	}
	return debugIdSlice, debugCountSlice
}
