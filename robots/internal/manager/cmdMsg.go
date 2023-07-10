package manager

import (
	"bufio"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/robots/conf"
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

/**
 *  @Description: 控制台发送协议	协议id 参数...
 *  @param:		dataArr [0] => 协议id ...参数  (为空提取控制台输入)
 */
func (this *Robot) MakeMsg(cmdId int32, data string) nw.ProtoMessage {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("robot makeMsg panic:%v, %s", err, debug.Stack())
		}
	}()

	var msg nw.ProtoMessage
	var dataArr []string
	var send bool
	if cmdId < 1 {
		fmt.Println("please write cmd and data")
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			fmt.Println("scanner error")
		}
		data = strings.ReplaceAll(scanner.Text(), "，", ",")
		dataArr = strings.Split(data, " ")
		cmdId = common.Str2Int32(dataArr[0])
		send = true
	} else {
		data = strings.ReplaceAll(data, "，", ",")
		dataArr = make([]string, 1)
		dataArr[0] = ""
		dataArr = append(dataArr, strings.Split(data, ",")...)
	}
	if cmdId < 1 {
		fmt.Println("cmdId nothing")
		return msg
	}
	switch cmdId {
	case pb.CmdEnterGameReqId:
		msg = &pb.EnterGameReq{OpenId: this.openId, ServerId: int32(conf.Conf.ServerId), LoginKey: m.logginKey}
		this.status = conf.STATUS_ENTERGAME_ING
		this.lastEnterGameTime = time.Now().Unix()
	case pb.CmdCreateUserReqId:
		msg = &pb.CreateUserReq{NickName: this.openId, Sex: common.Str2Int32(dataArr[1]), Job: common.Str2Int32(dataArr[2])}
	case pb.CmdCreateHeroReqId:
		msg = &pb.CreateHeroReq{Sex: common.Str2Int32(dataArr[1]), Job: common.Str2Int32(dataArr[2])}
	case pb.CmdChatSendReqId:
		msg = &pb.ChatSendReq{Type: common.Str2Int32(dataArr[1]), Msg: dataArr[2], ToId: common.Str2Int32(dataArr[3])}
	case pb.CmdMoneyPayReqId:
		msg = &pb.MoneyPayReq{PayType: common.Str2Int32(dataArr[1]), PayNum: common.Str2Int32(dataArr[2]), TypeId: common.Str2Int32(dataArr[3])}
	case pb.CmdBagSortReqId:
		msg = &pb.BagSortReq{}
	case pb.CmdEquipRecoverReqId:
		msg = &pb.EquipRecoverReq{Positions: common.Str2Int32Arr(dataArr[1])}
	case pb.CmdDebugAddGoodsReqId:
		msg = &pb.DebugAddGoodsReq{Id: common.Str2Int32Arr(dataArr[1]), Count: common.Str2Int32Arr(dataArr[2])}
	case pb.CmdItemUseReqId:
		msg = &pb.ItemUseReq{ItemId: common.Str2Int32(dataArr[1]), ItemNum: common.Str2Int32(dataArr[2]), HeroIndex: common.Str2Int32(dataArr[3])}
	case pb.CmdRechargeApplyPayReqId:
		msg = &pb.RechargeApplyPayReq{RechargeId: common.Str2Int32(dataArr[1]), PayNum: common.Str2Int32(dataArr[2])}
	//普通装备
	case pb.CmdEquipChangeReqId:
		msg = &pb.EquipChangeReq{HeroIndex: common.Str2Int32(dataArr[1]), Pos: common.Str2Int32(dataArr[2]), BagPos: common.Str2Int32(dataArr[3])}
	case pb.CmdEquipRemoveReqId:
		msg = &pb.EquipRemoveReq{HeroIndex: common.Str2Int32(dataArr[1]), Pos: common.Str2Int32(dataArr[2])}
	case pb.CmdEquipStrengthenReqId:
		msg = &pb.EquipStrengthenReq{HeroIndex: common.Str2Int32(dataArr[1]), Pos: common.Str2Int32(dataArr[2])}
	case pb.CmdEquipStrengthenAutoReqId:
		msg = &pb.EquipStrengthenAutoReq{HeroIndex: common.Str2Int32(dataArr[1]), IsBreak: common.Str2Bool(dataArr[2])}
	//宝石
	case pb.CmdJewelMakeAllReqId:
		msg = &pb.JewelMakeAllReq{HeroIndex: common.Str2Int32(dataArr[1])}
	//合体圣装
	case pb.CmdFitHolyEquipComposeReqId:
		msg = &pb.FitHolyEquipComposeReq{EquipId: common.Str2Int32(dataArr[1])}
	case pb.CmdFitHolyEquipDeComposeReqId:
		msg = &pb.FitHolyEquipDeComposeReq{BagPos: common.Str2Int32(dataArr[1])}
	case pb.CmdFitHolyEquipWearReqId:
		msg = &pb.FitHolyEquipWearReq{BagPos: common.Str2Int32(dataArr[1])}
	case pb.CmdFitHolyEquipRemoveReqId:
		msg = &pb.FitHolyEquipRemoveReq{Pos: common.Str2Int32(dataArr[1]), SuitType: common.Str2Int32(dataArr[2])}
	case pb.CmdFitHolyEquipSuitSkillChangeReqId:
		msg = &pb.FitHolyEquipSuitSkillChangeReq{SuitId: common.Str2Int32(dataArr[1])}
	//传世装备
	case pb.CmdChuanShiWearReqId:
		msg = &pb.ChuanShiWearReq{HeroIndex: common.Str2Int32(dataArr[1]), BagPos: common.Str2Int32(dataArr[2])}
	case pb.CmdChuanShiRemoveReqId:
		msg = &pb.ChuanShiRemoveReq{HeroIndex: common.Str2Int32(dataArr[1]), EquipPos: common.Str2Int32(dataArr[2])}
	case pb.CmdChuanShiDeComposeReqId:
		msg = &pb.ChuanShiDeComposeReq{BagPos: common.Str2Int32(dataArr[1])}
	case pb.CmdComposeChuanShiEquipReqId:
		msg = &pb.ComposeChuanShiEquipReq{ComposeSubId: common.Str2Int32(dataArr[1])}
	//内功
	case pb.CmdInsideAutoUpReqId:
		msg = &pb.InsideAutoUpReq{HeroIndex: common.Str2Int32(dataArr[1])}
	//连续充值
	case pb.CmdContRechargeReceiveReqId:
		msg = &pb.ContRechargeReceiveReq{ContRechargeId: common.Str2Int32(dataArr[1])}
	//月卡
	case pb.CmdMonthCardDailyRewardReqId:
		msg = &pb.MonthCardDailyRewardReq{MonthCardType: common.Str2Int32(dataArr[1])}
	//时装
	case pb.CmdFashionUpLevelReqId:
		msg = &pb.FashionUpLevelReq{HeroIndex: common.Str2Int32(dataArr[1]), FashionId: common.Str2Int32(dataArr[2])}
	case pb.CmdFashionWearReqId:
		msg = &pb.FashionWearReq{HeroIndex: common.Str2Int32(dataArr[1]), FashionId: common.Str2Int32(dataArr[2]), IsWear: common.Str2Bool(dataArr[3])}
	//羽翼
	case pb.CmdWingUseMaterialReqId:
		msg = &pb.WingUseMaterialReq{HeroIndex: common.Str2Int32(dataArr[1])}
	case pb.CmdWingWearReqId:
		msg = &pb.WingWearReq{HeroIndex: common.Str2Int32(dataArr[1]), WingId: common.Str2Int32(dataArr[2])}
	//法阵
	case pb.CmdMagicCircleUpLvReqId:
		msg = &pb.MagicCircleUpLvReq{HeroIndex: common.Str2Int32(dataArr[1]), MagicCircleType: common.Str2Int32(dataArr[2])}
	case pb.CmdMagicCircleChangeWearReqId:
		msg = &pb.MagicCircleChangeWearReq{HeroIndex: common.Str2Int32(dataArr[1]), MagicCircleLvId: common.Str2Int32(dataArr[2])}
	//公会
	case pb.CmdGuildLoadInfoReqId:
		msg = &pb.GuildLoadInfoReq{}
	case pb.CmdAllGuildInfosReqId:
		msg = &pb.AllGuildInfosReq{}
	case pb.CmdCreateGuildReqId:
		msg = &pb.CreateGuildReq{GuildName: dataArr[1], GuildIcon: dataArr[2], Notice: dataArr[3]}
	case pb.CmdApplyJoinGuildReqId:
		msg = &pb.ApplyJoinGuildReq{GuildId: common.Str2Int32(dataArr[1])}
	case pb.CmdJoinGuildCombatLimitReqId:
		msg = &pb.JoinGuildCombatLimitReq{Combat: common.Str2Int64(dataArr[1]), IsAgree: common.Str2Int32(dataArr[2])}
	//技能
	case pb.CmdSkillUpLvReqId:
		msg = &pb.SkillUpLvReq{HeroIndex: common.Str2Int32(dataArr[1]), SkillId: common.Str2Int32(dataArr[2])}
	case pb.CmdSkillChangeWearReqId:
		msg = &pb.SkillChangeWearReq{HeroIndex: common.Str2Int32(dataArr[1]), SkillId: common.Str2Int32(dataArr[2])}
	//挖矿
	case pb.CmdMiningListReqId:
		msg = &pb.MiningListReq{}
	case pb.CmdMiningStartReqId:
		msg = &pb.MiningStartReq{}
	//世界首领
	case pb.CmdLoadWorldLeaderReqId:
		msg = &pb.LoadWorldLeaderReq{}
	case pb.CmdWorldLeaderEnterReqId:
		msg = &pb.WorldLeaderEnterReq{StageId: common.Str2Int32(dataArr[1])}
	//拍卖行
	case pb.CmdAuctionInfoReqId:
		msg = &pb.AuctionInfoReq{AuctionType: common.Str2Int32(dataArr[1])}
	case pb.CmdBidReqId:
		msg = &pb.BidReq{AuctionType: common.Str2Int32(dataArr[1]), AuctionId: common.Str2Int64(dataArr[2]), IsBuyNow: common.Str2Int32(dataArr[3])}
	//神机宝库
	case pb.CmdCardActivityApplyGetReqId:
		msg = &pb.CardActivityApplyGetReq{Times: common.Str2Int32(dataArr[1])}
	//远古首领
	case pb.CmdAncientBossLoadReqId:
		msg = &pb.AncientBossLoadReq{Area: common.Str2Int32(dataArr[1])}
	case pb.CmdAncientBossBuyNumReqId:
		msg = &pb.AncientBossBuyNumReq{Use: common.Str2Bool(dataArr[1]), BuyNum: common.Str2Int32(dataArr[2])}
	case pb.CmdEnterAncientBossFightReqId:
		msg = &pb.EnterAncientBossFightReq{StageId: common.Str2Int32(dataArr[1])}
	//远古神技
	case pb.CmdAncientSkillActiveReqId:
		msg = &pb.AncientSkillActiveReq{HeroIndex: common.Str2Int32(dataArr[1])}
	case pb.CmdAncientSkillUpLvReqId:
		msg = &pb.AncientSkillUpLvReq{HeroIndex: common.Str2Int32(dataArr[1])}
	case pb.CmdAncientSkillUpGradeReqId:
		msg = &pb.AncientSkillUpGradeReq{HeroIndex: common.Str2Int32(dataArr[1])}
	//称号
	case pb.CmdTitleActiveReqId:
		msg = &pb.TitleActiveReq{TitleId: common.Str2Int32(dataArr[1])}
	//多宝阁
	case pb.CmdTreasureShopLoadReqId:
		msg = &pb.TreasureShopLoadReq{}
	case pb.CmdTreasureShopRefreshReqId:
		msg = &pb.TreasureShopRefreshReq{}
	case pb.CmdTreasureShopBuyReqId:
		msg = &pb.TreasureShopBuyReq{}
	//传世强化
	case pb.CmdChuanshiStrengthenReqId:
		msg = &pb.ChuanshiStrengthenReq{HeroIndex: common.Str2Int32(dataArr[1]), EquipPos: common.Str2Int32(dataArr[2]), Stone: common.Str2Int32(dataArr[3])}
	//战宠附体
	case pb.CmdPetAppendageReqId:
		msg = &pb.PetAppendageReq{PetId: common.Str2Int32(dataArr[1])}
	//炼狱首领
	case pb.CmdHellBossLoadReqId:
		msg = &pb.HellBossLoadReq{Floor: common.Str2Int32(dataArr[1])}
	//排行榜
	case pb.CmdRankLoadReqId:
		msg = &pb.RankLoadReq{Type: common.Str2Int32(dataArr[1])}
	//特权
	case pb.CmdPrivilegeBuyReqId:
		msg = &pb.PrivilegeBuyReq{PrivilegeId: common.Str2Int32(dataArr[1])}
	default:
		fmt.Println("未找到协议")
	}
	//if msg != nil {
	//fmt.Println("msgId:", common.Str2Int32(dataArr[0]), "---- data:", dataArr[1:])
	//if err := this.SendMessage(0, msg); err != nil {
	//	fmt.Println("err:", err)
	//}
	//} else {
	//	fmt.Println("nothing send")
	//}
	logger.Info("cmd:%v data:%v", cmdId, dataArr)
	if msg != nil && send {
		this.SendMessage(0, msg)
		time.Sleep(time.Millisecond * 100)
	}

	return msg
}

func (this *Robot) DebugAdd(ids []int32, count []int32) nw.ProtoMessage {
	return &pb.DebugAddGoodsReq{Id: ids, Count: count}
}
