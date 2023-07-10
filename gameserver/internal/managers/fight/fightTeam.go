package fight

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/objs"
	"fmt"
)

func (this *Fight) GetTeamIndex(user *objs.User, stageId int) int {

	stageConf := gamedb.GetStageStageCfg(stageId)
	switch stageConf.Type {
	case constFight.FIGHT_TYPE_STAGE, //界面关卡战斗
		constFight.FIGHT_TYPE_PERSON_BOSS,        //个人boss
		constFight.FIGHT_TYPE_TOWERBOSS,          //爬塔
		constFight.FIGHT_TYPE_MATERIAL,           //材料副本
		constFight.FIGHT_TYPE_VIPBOSS,            //vipboss副本
		constFight.FIGHT_TYPE_MAIN_CITY,          //主城
		constFight.FIGHT_TYPE_EXPBOSS,            //经验副本
		constFight.FIGHT_TYPE_ARENA,              //竞技场
		constFight.FIGHT_TYPE_FIELD,              //野战
		constFight.FIGHT_TYPE_MINING,             //矿洞
		constFight.FIGHT_TYPE_STAGE_BOSS,         //挂机boss
		constFight.FIGHT_TYPE_GUILD_BONFIRE,      //公会篝火
		constFight.FIGHT_TYPE_CROSS_WORLD_LEADER, //跨服世界首领
		constFight.FIGHT_TYPE_GUARDPILLAR,        //守卫龙柱
		constFight.FIGHT_TYPE_DABAO:              //打宝秘境
		return constFight.FIGHT_TEAM_ONE
	case constFight.FIGHT_TYPE_FIELDBOSS, //野外boss
		constFight.FIGHT_TYPE_PUBLIC_DABAO,        //公共打宝地图
		constFight.FIGHT_TYPE_DARKPALACE,          //暗殿
		constFight.FIGHT_TYPE_DARKPALACE_BOSS,     //暗殿
		constFight.FIGHT_TYPE_ANCIENT_BOSS,        //远古首领
		constFight.FIGHT_TYPE_PAODIAN,             //泡点PK
		constFight.FIGHT_TYPE_HELL_BOSS,           //炼狱boss
		constFight.FIGHT_TYPE_HELL,                //炼狱地图
		constFight.FIGHT_TYPE_PUBLIC_DABAO_SINGLE: //单人打宝秘境
		return user.Id
	case constFight.FIGHT_TYPE_MAGIC_TOWER: //九层魔塔
		if user.GuildData.NowGuildId > 0 {
			return user.GuildData.NowGuildId
		} else {
			return user.Id
		}
	case constFight.FIGHT_TYPE_SHABAKE, constFight.FIGHT_TYPE_SHABAKE_NEW: //沙巴克
		return user.GuildData.NowGuildId
	case constFight.FIGHT_TYPE_CROSS_SHABAKE: //跨服沙巴克
		return base.Conf.ServerId

	default:

		if base.Conf.Sandbox {
			//目前未实现，弃用战斗类型
			// constFight.FIGHT_TYPE_STAGE              ,
			//constFight.FIGHT_TYPE_WORLDBOSS          ,
			//constFight.FIGHT_TYPE_WORLDBOSS_NEW      ,
			panic(fmt.Sprintf("获取战斗队伍索引错误：%v", stageConf.Type))
		}
		return constFight.FIGHT_TEAM_ONE
	}
}
