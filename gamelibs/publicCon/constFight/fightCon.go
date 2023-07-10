package constFight

import "cqserver/protobuf/pb"

const (
	FIGHT_TYPE_STAGE               = 1  //界面关卡战斗
	FIGHT_TYPE_PERSON_BOSS         = 2  //个人boss
	FIGHT_TYPE_TOWERBOSS           = 3  //爬塔
	FIGHT_TYPE_FIELDBOSS           = 4  //野外boss
	FIGHT_TYPE_WORLDBOSS           = 5  //世界boss
	FIGHT_TYPE_MATERIAL            = 6  //材料副本
	FIGHT_TYPE_VIPBOSS             = 7  //vipboss副本
	FIGHT_TYPE_MAIN_CITY           = 8  //主城
	FIGHT_TYPE_EXPBOSS             = 9  //经验副本
	FIGHT_TYPE_ARENA               = 10 //竞技场
	FIGHT_TYPE_PUBLIC_DABAO        = 11 //公共打宝地图
	FIGHT_TYPE_DARKPALACE          = 12 //暗殿
	FIGHT_TYPE_WORLDBOSS_NEW       = 13 //新世界boss
	FIGHT_TYPE_FIELD               = 14 //野战
	FIGHT_TYPE_MINING              = 15 //矿洞
	FIGHT_TYPE_DARKPALACE_BOSS     = 16 //暗殿
	FIGHT_TYPE_STAGE_BOSS          = 17 //挂机boss
	FIGHT_TYPE_GUILD_BONFIRE       = 19 //公会篝火
	FIGHT_TYPE_PAODIAN             = 20 //泡点PK
	FIGHT_TYPE_SHABAKE             = 21 //沙巴克
	FIGHT_TYPE_CROSS_WORLD_LEADER  = 22 //跨服世界首领
	FIGHT_TYPE_CROSS_SHABAKE       = 23 //跨服沙巴克
	FIGHT_TYPE_ANCIENT_BOSS        = 25 //远古首领
	FIGHT_TYPE_GUARDPILLAR         = 26 //守卫龙柱
	FIGHT_TYPE_MAGIC_TOWER         = 27 //九层魔塔
	FIGHT_TYPE_HELL_BOSS           = 28 //炼狱首领
	FIGHT_TYPE_HELL                = 29 //炼狱首领大地图
	FIGHT_TYPE_DABAO               = 31 //打宝秘境
	FIGHT_TYPE_SHABAKE_NEW         = 32 //沙巴克
	FIGHT_TYPE_PUBLIC_DABAO_SINGLE = 33 //单人打宝秘境
)

const (
	FIGHT_SHABAKE_NPC_ID_ONE   = 81 //传送门
	FIGHT_SHABAKE_NPC_ID_TWO   = 80 //传送Npc
	FIGHT_SHABAKE_NPC_ID_THREE = 82 //治疗Npc
)

const SCENE_NOTIFIER_CHAN_TYPE = false

const (
	SCENE_SIZE      = 50  //地图格子大小
	MOVE_CHECK_FIX  = 0.7 //玩家移动检查间隔容错
	MOVE_FAST_LIMIT = 2   //移动过快检查次数
)

const (
	FIGHT_USER_TYPE_PLAYER      = 0 //真实玩家
	FIGHT_USER_TYPE_PLAYER_data = 1 //玩家数据玩家
	FIGHT_USER_TYPE_CONF        = 2 //配置假人
)

const (
	FIGHT_TEAM_ZERO = 0
	FIGHT_TEAM_ONE  = 1
)

const MONSTER_DEFAULT_DIR = pb.SCENEDIR_BOTTOM

const (
	FIGHT_BIRTH_TYPE_LINE     = 1 //直线
	FIGHT_BIRTH_TYPE_TRIANGLE = 2 //三角形
)
const (
	WOWLD_BOSS_STATUS_INIT  = 0
	WOWLD_BOSS_STATUS_READY = 1
	WOWLD_BOSS_STATUS_RUN   = 2
	WOWLD_BOSS_STATUS_STOP  = 3
)

const (
	MONSTER_TYPE_NORMAL = 1 //小怪
	MONSTER_TYPE_BOSS   = 2 //boss
)

const (
	RELIVE_ADDR_TYPE_SITU  = 1 //原地复活
	RELIVE_ADDR_TYPE_BIRTH = 2 //出生点复活
	RELIVE_ADDR_BACK_CITY  = 3 //回城复活
)

const (
	RELIVE_TYPE_NOMAL = 0 //普通复活
	RELIVE_TYPE_BUFF  = 1 //buff复活
	RELIVE_TYPE_SKILL = 2 //技能复活
	RELIVE_TYPE_COST  = 3 //消耗元宝复活
)

const (
	RELIVE_DELAY = 500 //复活延迟时间,毫秒
)

const (
	SKILL_PASSIVE_CONDITION_NORMAL   = 0  //普通
	SKILL_PASSIVE_CONDITION_ATK      = 1  //主动攻击
	SKILL_PASSIVE_CONDITION_BE_ATK   = 2  //被动攻击
	SKILL_PASSIVE_CONDITION_FATAL    = 3  //致命一击
	SKILL_PASSIVE_CONDITION_RAGE     = 4  //狂暴攻击
	SKILL_PASSIVE_CONDITION_BE_FATAL = 5  //受到致命一击
	SKILL_PASSIVE_CONDITION_CRIT     = 6  //暴击
	SKILL_PASSIVE_KILL_TAREGET       = 7  //击杀目标
	SKILL_PASSIVE_RELIVE             = 8  //复活
	SKILL_PASSIVE_SAME_SECOND        = 10 //间隔秒数
	SKILL_PASSIVE_DAMAGE             = 11 //伤害累加
)

const (
	SKILL_PASSIVE_HP_CHECK_ENMY_LOW  = 1 //1 目标生命值低于X%
	SKILL_PASSIVE_HP_CHECK_ENMY_HIGH = 2 //目标生命值高于X%
	SKILL_PASSIVE_HP_CHECK_SELF_LOW  = 3 //自生生命值低于X%
	SKILL_PASSIVE_HP_CHECK_SELF_HIGH = 4 //自生生命自高于X%
	SKILL_PASSIVE_HP_CHECK_LESS      = 5 //生命降低X%
	SKILL_PASSIVE_COMBAT_HIGH        = 6 //高于自身的万分比
	SKILL_PASSIVE_COMBAT_LOW         = 7 //低于自身的万分比
)

const (
	BUFF_RULE_REPLACE        = 1 //同类直接顶替
	BUFF_RULE_NO_REPLACE     = 2 //同类存在则不添加
	BUFF_RULE_NO_AFFECT      = 3 //同类存在时互补影响，共存
	BUFF_RULE_REPLACE_LOW_LV = 4 //同类存在，高级替代低级
)

const (
	BUFF_TARGET_ENMY = 1
	BUFF_TARGET_SELF = 2
)

const (
	BUFF_BY_SKILL_PRE   = 1
	BUFF_BY_SKILL_AFTER = 2
)

const (
	BUFF_KEY_ZERO  = 0
	BUFF_KEY_ONE   = 1
	BUFF_KEY_TWO   = 2
	BUFF_KEY_THREE = 3
	BUFF_KEY_FOUR  = 4
)

const (
	BUFF_IS_DEBUFF_TRUE  = 1 //buff为debuff
	BUFF_IS_DEBUFF_FALSE = 2 //buff不是debuff
)

const (
	BUFF_CAN_REMOVE = 1
)

const (
	MOVE_TYPE_HOME  = 1 //回家移动
	MOVE_TYPE_CHASE = 2 //追击敌人
)

const (
	FIGHT_TYPE_SHABAKE_STAGE       = 221
	FIGHT_TYPE_SHABAKE_NEW_STAGE   = 222
	FIGHT_TYPE_SHABAKE_CROSS_STAGE = 226
	FIGHT_TYPE_GUILD_BONFIRE_STAGE = 301
	FIGHT_TYPE_GUARDPILLAR_STAGE   = 302
	FIGHT_TYPE_MAIN_CITY_STAGE     = 5000
	FIGHT_TYPE_ARENA_STAGE         = 6000
	FIGHT_TYPE_FIELD_STAGE         = 70001
	FIGHT_TYPE_MINING_STAGE        = 80001
	FIGHT_TYPE_FIELD_BOSS_STAGE    = 8000000
)

const (
	SKILL_PUTONG_ID        = 100   //普通攻击（战士）
	SKILL_PUTONG_HUOQIU_ID = 110   //火球术（法师）
	SKILL_PUTONG_HUOFU_ID  = 120   //灵魂火符（道士）
	SKILL_JICHU_ID         = 200   //基础剑术
	SKILL_CISHA_ID         = 201   //刺杀剑术
	SKILL_YEMAN_ID         = 231   //野蛮冲撞
	SKILL_SHIHOU_ID        = 251   //狮子吼
	SKILL_HUOQIU_ID        = 301   //火球术
	SKILL_KANGJU_ID        = 311   //抗拒火环
	SKILL_MIETIANHUO_ID    = 351   //灭天火
	SKILL_LIUXINGHUOYU_ID  = 371   //流星火雨
	SKILL_CHONGFENGZHAN_ID = 11300 //冲锋斩
	SKILL_MIE_TIAN_HUO     = 20100
	SKILL_FENG_CHI         = 20200
	SKILL_CUT_ZHAN         = 120000 //切割技能 战士
	SKILL_CUT_FA           = 120100 //切割技能 法：
	SKILL_CUT_DAO          = 120200 //切割技能 道：
)

const (
	FIT_SKILL_ZHUDONG_ID  = 1 //合体主动技能
	FIT_SKILL_ZHUDONG_POS = 1 //合体主动技能部位
)

var (
	JOB_SKILL_MAP = map[int]int{
		pb.JOB_ZHANSHI: SKILL_PUTONG_ID,
		pb.JOB_FASHI:   SKILL_PUTONG_HUOQIU_ID,
		pb.JOB_DAOSHI:  SKILL_PUTONG_HUOFU_ID,
	}
)

const (
	SKILL_ATTACK_EFFECT_TYPE_1 = 1 //1:对某个目标提升效果
	SKILL_ATTACK_EFFECT_TYPE_2 = 2 //2:移除增益效果
	SKILL_ATTACK_EFFECT_TYPE_3 = 3 //3:攻击伤害x%回血量
	SKILL_ATTACK_EFFECT_TYPE_4 = 4 //4:治疗x%最为伤害施加给敌人
	SKILL_ATTACK_EFFECT_TYPE_5 = 5 //4:造成多倍伤害
)

const SKILL_KANGJU_MOVE_DIS = 2

const (
	ENTER_FIGHT_TYPE_NOMAL          = 0
	ENTER_FIGHT_TYPE_RELIVE_TO_CITY = 1
	ENTER_FIGHT_TYPE_NEW_HERO       = 2
)

const (
	SCENE_ENTER_FIT = 1
)

const (
	DEATH_REASON_ATTACK = 0
	DEATH_REASON_BUFF   = 1
	DEATH_REASON_FIT    = 2
)

const (
	LEAVE_FIGHT_TYPE_NOMAL   = 0
	LEAVE_FIGHT_TYPE_OFFLINE = 1
)

const (
	MONSTER_DROP_FOR_KILLER = 1 //怪物掉落归击杀者
	MONSTER_DROP_FOR_OWNER  = 2 //怪物掉落归归属者
)

var (
	DROP_ITEM_INTO_BAG = map[int]bool{pb.ITEMID_EXP: true, pb.ITEMID_GOLD: true}
)

var (
	HERO_BIRTH = map[int]map[string]int{2: {"left": 4, "back": 1}, 3: {"left": 1, "back": 4}}
)

const (
	COLLECTION_EFFECT_BUFF  = 1 //获得buff
	COLLECTION_EFFECT_ITEM  = 2 //获得物品
	COLLECTION_EFFECT_OTHER = 3 //特殊
	COLLECTION_EFFECT_SCORE = 4 //增加积分
)

const (
	COLLECTION_TYPE_ONE   = 1 //采集后刷新
	COLLECTION_TYPE_TWO   = 2 //采集后不刷新
	COLLECTION_TYPE_THREE = 3 //特殊刷新
)

const (
	FIT_ID = 1
)

var (
	ATTACK_PROP = map[int]int{
		pb.PROPERTY_PATT_MIN: pb.PROPERTY_PATT_MIN, pb.PROPERTY_PATT_MAX: pb.PROPERTY_PATT_MAX,
		pb.PROPERTY_MATT_MIN: pb.PROPERTY_PATT_MIN, pb.PROPERTY_MATT_MAX: pb.PROPERTY_PATT_MAX,
		pb.PROPERTY_TATT_MIN: pb.PROPERTY_PATT_MIN, pb.PROPERTY_TATT_MAX: pb.PROPERTY_PATT_MAX,
	}
)

const (
	MAGIC_TOWER_ENTER_TYPE_FIRST  = 0
	MAGIC_TOWER_ENTER_TYPE_NEXT   = 1
	MAGIC_TOWER_ENTER_TYPE_RELIVE = 2
)
