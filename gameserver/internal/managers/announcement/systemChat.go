package announcement

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constAnnouncement"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
	"fmt"
	"time"
)

func (this *AnnouncementManager) FightSendSystemChat(user *objs.User, items map[int]int, stageId, types int) {

	maxQuality := 0
	itemId := 0
	if types == pb.SCROLINGTYPE_KILL_MONSTER {
		types = pb.SCROLINGTYPE_ANCIENT_BOSS
	}
	cfg1 := gamedb.GetScrollingScrollingCfg(types)
	if cfg1 == nil {
		return
	}
	for item := range items {
		cfg := gamedb.GetItemBaseCfg(item)
		if cfg == nil {
			continue
		}
		if cfg.Quality >= cfg1.Condition && cfg.Quality > maxQuality {
			maxQuality = cfg.Quality
			itemId = item
		}
	}
	if types == pb.SCROLINGTYPE_ANCIENT_BOSS {
		types = pb.SCROLINGTYPE_KILL_MONSTER
	}
	if itemId > 0 {
		this.SendSystemChat(user, types, itemId, stageId)
	}
	return
}

func (this *AnnouncementManager) SendSystemChat(user *objs.User, types, itemId, stageId int) {

	content := this.BuildContent(user, types, itemId, stageId)
	if content == "" {
		return
	}
	ntf := &pb.PaoMaDengNtf{}
	ntf.PaoMaDengInfos = make([]*pb.PaoMaDengInfo, 0)
	ntf.PaoMaDengInfos = append(ntf.PaoMaDengInfos, &pb.PaoMaDengInfo{Type: constAnnouncement.ActivityInfo, Content: content})
	this.BroadcastAll(ntf)
	return
}

func (this *AnnouncementManager) BuildContent(user *objs.User, types, itemId, stageId int) string {

	content := ""
	switch types {

	case pb.SCROLINGTYPE_JI_HUO_BAI_YING_KA:
		content = this.getBaiYinKaContent(user.Id, user.NickName)
	case pb.SCROLINGTYPE_JI_HUO_BAI_HUANG_JIN_KA:
		content = this.getHuangJinKaContent(user.Id, user.NickName)
	case pb.SCROLINGTYPE_GOU_MAI_ZHAN_LIN:
		content = this.getZLinContent(user.Id, user.NickName)
	case pb.SCROLINGTYPE_GOU_MAI_QI_RI_TOU_ZHI:
		content = this.getQiRiTouZiContent(user.Id, user.NickName)
	case pb.SCROLINGTYPE_CHOUKA:
		content = this.getCardContent(user.Id, itemId, user.NickName)
	case pb.SCROLINGTYPE_KILL_GET:
		content = this.getKillDropContent(user.Id, itemId, stageId, user.NickName)
	case pb.SCROLINGTYPE_SHABAKE_BEGIN:
		content = this.getShaBakeOpenContent()
	case pb.SCROLINGTYPE_AN_DIAN_BOSS_DROP:
		content = this.getAnDianKillDropContent(user.Id, itemId, stageId, user.NickName)
	case pb.SCROLINGTYPE_WORLD_LEADER_OPEN:
		content = this.getWorldLeaderOpenContent(stageId)
	case pb.SCROLINGTYPE_KAI_FU_BUY:
		content = this.getKaiFuGiftContent(user.Id, user.NickName)
	case pb.SCROLINGTYPE_LIAN_CHONG_HAO_LI:
		content = this.getLianCongGiftContent(user.Id, user.NickName)
	case pb.SCROLINGTYPE_FIGHT_HELP:
		content = this.getFightHelpContent(user.FightStageId)
	case pb.SCROLINGTYPE_LOTTERY:
		content = this.getLotteryContent(user.Id, itemId, user.NickName)
	case pb.SCROLINGTYPE_ANCIENT_BOSS: //远古首领
		content = this.getAncientBossKillDropContent(user.Id, itemId, stageId, user.NickName)
	case pb.SCROLINGTYPE_HELL_BOSS: //炼狱首领
		content = this.getHellBossKillDropContent(user.Id, itemId, stageId, user.NickName)
	case pb.SCROLINGTYPE_DA_BAO:
		content = this.getDaBaoKillDropContent(user.Id, itemId, stageId, user.NickName)
	case pb.SCROLINGTYPE_PRIVILEGE_OPEN:
		content = this.getPrivilegeContent(user.Id, itemId, user.NickName)
	case pb.SCROLINGTYPE_KILL_MONSTER:
		content = this.getKillMonsterDropContent(user.Id, itemId, stageId, user.NickName)
	}

	return content
}

// 白银卡 文案
func (this *AnnouncementManager) getBaiYinKaContent(userId int, nickName string) string {
	str := ""
	content := ""
	cfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_JI_HUO_BAI_YING_KA)
	if cfg == nil {
		return ""
	}
	str = cfg.Txt
	content = fmt.Sprintf(str, userId, nickName)
	return content
}

// 白银卡 文案
func (this *AnnouncementManager) getHuangJinKaContent(userId int, nickName string) string {
	str := ""
	content := ""
	cfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_JI_HUO_BAI_HUANG_JIN_KA)
	if cfg == nil {
		return ""
	}
	str = cfg.Txt
	content = fmt.Sprintf(str, userId, nickName)
	return content
}

// 战令 文案
func (this *AnnouncementManager) getZLinContent(userId int, nickName string) string {
	str := ""
	content := ""
	cfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_GOU_MAI_ZHAN_LIN)
	if cfg == nil {
		return ""
	}
	str = cfg.Txt
	content = fmt.Sprintf(str, userId, nickName)
	return content
}

// 七日投资 文案
func (this *AnnouncementManager) getQiRiTouZiContent(userId int, nickName string) string {
	str := ""
	content := ""
	cfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_GOU_MAI_QI_RI_TOU_ZHI)
	if cfg == nil {
		return ""
	}
	str = cfg.Txt
	content = fmt.Sprintf(str, userId, nickName)
	return content
}

// 抽卡  文案
func (this *AnnouncementManager) getCardContent(userId, itemId int, nickName string) string {
	str := ""
	content := ""
	cfg := gamedb.GetItemBaseCfg(itemId)
	if cfg == nil {
		return ""
	}

	scrollingCfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_CHOUKA)
	if scrollingCfg == nil {
		return ""
	}
	if cfg.Quality < scrollingCfg.Condition {
		return ""
	}
	str = scrollingCfg.Txt
	content = fmt.Sprintf(str, userId, nickName, itemId, cfg.Name)
	return content
}

// 击杀世界首领 掉落奖励  文案
func (this *AnnouncementManager) getKillDropContent(userId, itemId, stageId int, nickName string) string {
	str := ""
	content := ""
	cfg := gamedb.GetItemBaseCfg(itemId)
	if cfg == nil {
		return ""
	}

	scrollingCfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_KILL_GET)
	if scrollingCfg == nil {
		return ""
	}
	if !this.getSkillDropContent(itemId, cfg, scrollingCfg) {
		return ""
	}

	str = scrollingCfg.Txt

	monsterName := gamedb.GetItemSourceByStageId(stageId)

	content = fmt.Sprintf(str, userId, nickName, monsterName, itemId, cfg.Name)
	return content
}

func (this *AnnouncementManager) getShaBakeOpenContent() string {

	cfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_SHABAKE_BEGIN)
	if cfg == nil {
		return ""
	}

	limitedOpenDay := gamedb.GetConf().ShabakeTime1
	openDay := this.GetSystem().GetServerOpenDaysByServerId(base.Conf.ServerId)
	//开服天数限制
	if limitedOpenDay > openDay {
		return ""
	}

	limitedWeekDay := gamedb.GetConf().ShabakeTime2
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
			return ""
		}
	}
	return cfg.Txt

}

// 击杀暗殿首领 掉落奖励  文案
func (this *AnnouncementManager) getAnDianKillDropContent(userId, itemId, stageId int, nickName string) string {
	str := ""
	content := ""
	cfg := gamedb.GetItemBaseCfg(itemId)
	if cfg == nil {
		return ""
	}

	scrollingCfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_AN_DIAN_BOSS_DROP)
	if scrollingCfg == nil {
		return ""
	}
	if !this.getSkillDropContent(itemId, cfg, scrollingCfg) {
		return ""
	}

	str = scrollingCfg.Txt

	monsterName := gamedb.GetItemSourceByStageId(stageId)

	content = fmt.Sprintf(str, userId, nickName, monsterName, itemId, cfg.Name)
	return content
}

// 世界首领开启 广播 文案
func (this *AnnouncementManager) getWorldLeaderOpenContent(stageId int) string {

	cfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_WORLD_LEADER_OPEN)
	if cfg == nil {
		return ""
	}

	monsterName := gamedb.GetItemSourceByStageId(stageId)
	return fmt.Sprintf(cfg.Txt, monsterName)
}

// 购买开服礼包 文案
func (this *AnnouncementManager) getKaiFuGiftContent(userId int, nickName string) string {
	str := ""
	content := ""
	cfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_KAI_FU_BUY)
	if cfg == nil {
		return ""
	}
	str = cfg.Txt
	content = fmt.Sprintf(str, userId, nickName)
	return content
}

// 成功领取连充豪礼 文案
func (this *AnnouncementManager) getLianCongGiftContent(userId int, nickName string) string {
	str := ""
	content := ""
	cfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_LIAN_CHONG_HAO_LI)
	if cfg == nil {
		return ""
	}
	str = cfg.Txt
	content = fmt.Sprintf(str, userId, nickName)
	return content
}

/*战斗协助*/
func (this *AnnouncementManager) getFightHelpContent(stageId int) string {
	str := ""
	content := ""
	cfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_FIGHT_HELP)
	if cfg == nil {
		return ""
	}
	stageConf := gamedb.GetStageStageCfg(stageId)
	str = cfg.Txt
	content = fmt.Sprintf(str, stageConf.Name, stageId)
	return content
}

func (this *AnnouncementManager) getSkillDropContent(itemId int, cfg *gamedb.ItemBaseCfg, scrollingCfg *gamedb.ScrollingScrollingCfg) bool {

	if cfg.Type == pb.ITEMTYPE_EQUIP && gamedb.GetEquipEquipCfg(itemId) != nil {
		if !this.dropItemQualityAndLvJudge(cfg, gamedb.GetEquipEquipCfg(itemId)) {
			return false
		}
	} else {
		if cfg.Quality < scrollingCfg.Condition {
			return false
		}
	}
	return true
}

// 判断掉落物品的品质和等级
func (this *AnnouncementManager) dropItemQualityAndLvJudge(itemInfo *gamedb.ItemBaseCfg, equipCfg *gamedb.EquipEquipCfg) bool {

	limitCfg := gamedb.GetConf().ScrollingEquipt
	openDay := this.GetSystem().GetServerOpenDaysByServerId(base.Conf.ServerId)
	for index, data := range limitCfg {
		if data == nil || len(data) < 3 {
			continue
		}

		if index+1 == len(limitCfg) {
			if openDay >= data[0] {
				if itemInfo.Quality >= data[1] && equipCfg.Class >= data[2] {
					return true
				}
				return false
			}
		}

		if openDay >= data[0] && openDay < limitCfg[index+1][0] {
			if itemInfo.Quality >= data[1] && equipCfg.Class >= data[2] {
				return true
			}
			return false
		}

	}
	return false
}

// 摇彩  文案
func (this *AnnouncementManager) getLotteryContent(userId, itemId int, nickName string) string {
	str := ""
	content := ""
	cfg := gamedb.GetItemBaseCfg(itemId)
	if cfg == nil {
		return ""
	}

	scrollingCfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_LOTTERY)
	if scrollingCfg == nil {
		return ""
	}

	str = scrollingCfg.Txt
	content = fmt.Sprintf(str, userId, nickName, itemId, cfg.Name)
	return content
}

// 击杀远古首领 掉落奖励  文案
func (this *AnnouncementManager) getAncientBossKillDropContent(userId, itemId, stageId int, nickName string) string {
	str := ""
	content := ""
	cfg := gamedb.GetItemBaseCfg(itemId)
	if cfg == nil {
		return ""
	}

	scrollingCfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_ANCIENT_BOSS)
	if scrollingCfg == nil {
		return ""
	}
	if !this.getSkillDropContent(itemId, cfg, scrollingCfg) {
		return ""
	}

	str = scrollingCfg.Txt

	monsterName := gamedb.GetItemSourceByStageId(stageId)

	content = fmt.Sprintf(str, userId, nickName, monsterName, itemId, cfg.Name)
	return content
}

// 击杀炼狱首领 掉落奖励  文案
func (this *AnnouncementManager) getHellBossKillDropContent(userId, itemId, stageId int, nickName string) string {
	str := ""
	content := ""
	cfg := gamedb.GetItemBaseCfg(itemId)
	if cfg == nil {
		return ""
	}

	scrollingCfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_HELL_BOSS)
	if scrollingCfg == nil {
		return ""
	}
	if !this.getSkillDropContent(itemId, cfg, scrollingCfg) {
		return ""
	}

	str = scrollingCfg.Txt

	monsterName := gamedb.GetItemSourceByStageId(stageId)

	content = fmt.Sprintf(str, userId, nickName, monsterName, itemId, cfg.Name)
	return content
}

// 击杀打宝 掉落奖励  文案
func (this *AnnouncementManager) getDaBaoKillDropContent(userId, itemId, stageId int, nickName string) string {
	str := ""
	content := ""
	cfg := gamedb.GetItemBaseCfg(itemId)
	if cfg == nil {
		return ""
	}

	scrollingCfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_DA_BAO)
	if scrollingCfg == nil {
		return ""
	}
	if !this.getSkillDropContent(itemId, cfg, scrollingCfg) {
		return ""
	}

	str = scrollingCfg.Txt

	monsterName := gamedb.GetItemSourceByStageId(stageId)

	content = fmt.Sprintf(str, userId, nickName, monsterName, itemId, cfg.Name)
	return content
}

// 特权激活 文案
func (this *AnnouncementManager) getPrivilegeContent(userId, privilegeId int, nickName string) string {
	privilegeCfg := gamedb.GetPrivilegePrivilegeCfg(privilegeId)
	if privilegeCfg == nil {
		return ""
	}

	scrollingCfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_PRIVILEGE_OPEN)
	if scrollingCfg == nil {
		return ""
	}
	return fmt.Sprintf(scrollingCfg.Txt, userId, nickName, privilegeCfg.Name)
}

// 击杀怪物 掉落奖励  文案
func (this *AnnouncementManager) getKillMonsterDropContent(userId, itemId, monsterId int, nickName string) string {
	str := ""
	content := ""
	cfg := gamedb.GetItemBaseCfg(itemId)
	if cfg == nil {
		return ""
	}

	scrollingCfg := gamedb.GetScrollingScrollingCfg(pb.SCROLINGTYPE_ANCIENT_BOSS)
	if scrollingCfg == nil {
		return ""
	}
	if !this.getSkillDropContent(itemId, cfg, scrollingCfg) {
		return ""
	}
	monsterConf := gamedb.GetMonsterMonsterCfg(monsterId)
	if monsterConf == nil {
		return ""
	}
	str = scrollingCfg.Txt
	content = fmt.Sprintf(str, userId, nickName, monsterConf.Name, itemId, cfg.Name)
	return content
}
