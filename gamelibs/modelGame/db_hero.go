package modelGame

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/golibs/dbmodel"
	"cqserver/protobuf/pb"
	"fmt"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gorp.v1"
	"time"
)

type Hero struct {
	dbmodel.DbTable
	Id                 int                  `db:"id" orm:"pk;auto"`
	UserId             int                  `db:"userId" orm:"comment(玩家角色Id)"`
	Index              int                  `db:"index" orm:"size(60);comment(玩家第几角色)"`
	Sex                int                  `db:"sex" orm:"comment(性别)"`
	Job                int                  `db:"job" orm:"comment(职业)"`
	Name               string               `db:"name" orm:"comment(昵称)"`
	Combat             int                  `db:"combat" orm:"type(int64);comment(总战力)"` //总战力
	ExpLvl             int                  `db:"ExpLvl" orm:"size(100);comment(经验池等级);default(1)"`
	Equips             model.Equips         `db:"equips" orm:"type(text);comment(装备信息)"`
	EquipsStrength     model.IntKv          `db:"equipsStrength" orm:"size(100);comment(装备强化1信息)"`
	Wings              model.Wings          `db:"wings" orm:"size(100);comment(羽翼信息)"`
	WingSpecial        model.IntKv          `db:"wingSpecial" orm:"size(50);comment(羽翼特殊属性)"`
	Zodiacs            model.Zodiacs        `db:"zodiac" orm:"size(1000);comment(生肖装备)"`
	Display            *model.Display       `db:"display" orm:"type(text);comment(形象显示信息)"`
	Kingarms           model.Kingarms       `db:"kingarms" orm:"size(1000);comment(帝器)"`
	Dictates           model.IntKv          `db:"dictate" orm:"size(500);comment(主宰装备)"`
	Jewel              model.Jewels         `db:"jewel" orm:"size(300);comment(宝石)"`
	Fashions           model.Fashions       `db:"fashion" orm:"null;type(text);comment(时装)"`
	Wear               *model.Wear          `db:"wear" orm:"null;type(text);comment(穿戴时装 装备 称号)"`
	Inside             *model.Inside        `db:"inside" orm:"size(150);comment(内功)"`
	Ring               model.Rings          `db:"ring" orm:"type(text);comment(特戒)"`
	GodEquips          model.GodEquips      `db:"godEquips" orm:"size(300);comment(神兵)"`
	Skills             model.Skills         `db:"skills" orm:"size(400);comment(技能)"`
	SkillBag           model.IntKv          `db:"skillBag" orm:"size(100);comment(技能背包)"`
	UniqueSkills       model.Skills         `db:"uniqueSkills" orm:"size(400);comment(合击)"`
	UniqueSkillBag     model.IntKv          `db:"uniqueSkillBag" orm:"size(100);comment(合击背包)"`
	Area               model.IntKv          `db:"area" orm:"size(70);comment(领域)"`
	EquipClear         model.EquipClears    `db:"equipClear" orm:"type(text);comment(洗练)"`
	DragonEquip        model.IntKv          `db:"dragonEquip" orm:"size(200);comment(龙器)"`
	MagicCircle        model.IntKv          `db:"magicCircle" orm:"size(100);comment(法阵)"`
	Talent             *model.Talent        `db:"talent" orm:"size(300);comment(天赋)"`
	HolyBeastInfos     model.HolyBeastInfos `db:"holyBeastInfos"orm:"type(text);comment(圣兽)"`
	HolyAllPoint       int                  `db:"holyAllPoint" orm:"size(70);comment(圣灵点)"`
	CreatTime          time.Time            `db:"creatTime" orm:"comment(创建时间)"`
	ChuanShi           model.IntKv          `db:"chuanshi" orm:"size(200);comment(传世装备)"`
	AncientSkill       *model.AncientSkill  `db:"ancientSkill" orm:"size(100);comment(远古神技)"`
	ChuanshiStrengthen model.IntKv          `db:"chuanshiStrengthen" orm:"size(100);comment(传世装备强化)"`
}

type HeroDisplay struct {
	UserId  int            `db:"userId" orm:"comment(玩家角色Id)"`
	Index   int            `db:"index" orm:"size(60);comment(玩家第几角色)"`
	Sex     int            `db:"sex" orm:"comment(性别)"`
	Job     int            `db:"job" orm:"comment(职业)"`
	ExpLvl  int            `db:"ExpLvl" orm:"size(4);comment(经验池等级);default(1)"`
	Name    string         `db:"name" orm:"comment(昵称)"`
	Combat  int            `db:"combat" orm:"type(int64);comment(总战力)"`
	Display *model.Display `db:"display" orm:"size(100);comment(形象显示信息)"`
}

func (this *Hero) TableName() string {
	return "hero"
}

type HeroModel struct {
	dbmodel.CommonModel
}

var (
	heroModel         = &HeroModel{}
	heroFields        = model.GetAllFieldsAsString(Hero{})
	heroDisplayFields = model.GetAllFieldsAsString(HeroDisplay{})
)

func init() {

	dbmodel.Register(model.DB_SERVER, heroModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(Hero{}, "hero").SetKeys(true, "id")
		orm.RegisterModelForAlias(model.DB_SERVER, new(Hero))
	})
}

func GetHeroModel() *HeroModel {
	return heroModel
}
func NewHero(userId, sex, job int) *Hero {
	hero := &Hero{
		UserId:         userId,
		Sex:            sex,
		Job:            job,
		Equips:         make(model.Equips),
		EquipsStrength: make(model.IntKv),
		Wings: model.Wings{
			0: &model.Wing{
				Id:     1,
				IsWear: true,
			},
		},
		WingSpecial:        make(model.IntKv),
		Zodiacs:            make(model.Zodiacs),
		Display:            &model.Display{},
		Kingarms:           make(model.Kingarms),
		Dictates:           make(model.IntKv),
		Fashions:           make(model.Fashions),
		Wear:               &model.Wear{AtlasWear: make(model.IntKv)},
		Inside:             &model.Inside{Acupoint: make(model.IntKv), Skill: make(map[int]*model.InsideSkill)},
		Ring:               make(model.Rings),
		GodEquips:          make(model.GodEquips),
		Jewel:              make(model.Jewels),
		Skills:             make(model.Skills),
		SkillBag:           make(model.IntKv),
		UniqueSkills:       make(model.Skills),
		UniqueSkillBag:     make(model.IntKv),
		Area:               make(model.IntKv),
		EquipClear:         make(model.EquipClears),
		DragonEquip:        make(model.IntKv),
		ExpLvl:             1,
		MagicCircle:        make(model.IntKv),
		Talent:             &model.Talent{TalentList: make(map[int]*model.TalentUnit)},
		HolyBeastInfos:     make(model.HolyBeastInfos),
		CreatTime:          time.Now(),
		ChuanShi:           make(model.IntKv),
		AncientSkill:       &model.AncientSkill{},
		ChuanshiStrengthen: make(model.IntKv),
	}

	hero.Name = GetHeroDefName(sex, job)
	defSkillId := constFight.JOB_SKILL_MAP[job]
	hero.Skills[defSkillId] = &model.SkillUnit{
		Id: defSkillId,
		Lv: 1,
	}
	for _, v := range pb.SKILLPOS_ARRAY {
		hero.SkillBag[v] = 0
		if v == pb.SKILLPOS_SIX {
			hero.SkillBag[v] = defSkillId
		}
	}
	for _, v := range pb.UNIQUESKILLPOS_ARRAY {
		hero.UniqueSkillBag[v] = 0
	}
	for _, v := range pb.EQUIPPOS_ARRAY {
		hero.Equips[v] = &model.Equip{}
		hero.EquipsStrength[v] = 0
		hero.EquipClear[v] = make([]*model.EquipClearUnit, 0)
	}
	for _, v := range pb.ZODIACTYPE_ARRAY {
		hero.Zodiacs[v] = &model.SpecialEquipUnit{}
	}
	for _, v := range pb.KINGARMSTYPE_ARRAY {
		hero.Kingarms[v] = &model.SpecialEquipUnit{}
	}
	for _, v := range pb.DICTATETYPE_ARRAY {
		hero.Dictates[v] = 0
	}
	for _, v := range pb.INSIDETYPE_ARRAY {
		hero.Inside.Acupoint[v] = 1
	}
	for _, v := range pb.RINGPOS_ARRAY {
		phantom := make(map[int]*model.RingPhantom)
		for _, pos := range pb.RINGPHANTOMPOS_ARRAY {
			phantom[pos] = &model.RingPhantom{Skill: make(model.IntKv)}
		}
		hero.Ring[v] = &model.RingUnit{Strengthen: 1, Pid: 1, Phantom: phantom}
	}
	for _, v := range pb.AREATYPE_ARRAY {
		hero.Area[v] = 0
	}

	for _, v := range pb.HOLYBEASTTYPE_ARRAY {
		hero.HolyBeastInfos[v] = &model.HolyBeastInfo{Types: v, Star: 0, ChooseProp: make(map[int]int)}
	}
	return hero
}

func GetHeroDefName(sex int, job int) string {
	nameStr := "男"
	if sex == pb.SEX_FEMALE {
		nameStr = "女"
	}
	switch job {
	case pb.JOB_ZHANSHI:
		nameStr += gamedb.JOBZHANSHINAME
	case pb.JOB_FASHI:
		nameStr += gamedb.JOBFSHINAME
	case pb.JOB_DAOSHI:
		nameStr += gamedb.JOBDAOSHINAME
	}
	return nameStr
}

func (this *HeroModel) Create(hero *Hero) error {
	return this.DbMap().Insert(hero)
}

func (this *HeroModel) GetHerosByUserId(userId int) ([]*Hero, error) {
	var heros []*Hero
	_, err := this.DbMap().Select(&heros, fmt.Sprintf("select %s from hero where userId = ? ", heroFields), userId)
	if err != nil {
		return nil, err
	}
	return heros, nil
}

func (this *HeroModel) GetHerosDisplayByUserId(userIds []int) ([]*HeroDisplay, error) {
	userIdsStr := common.JoinIntSlice(userIds, ",")
	var heros []*HeroDisplay
	_, err := this.DbMap().Select(&heros, fmt.Sprintf("select %s from hero where userId in ( %s ) ", heroDisplayFields, userIdsStr))
	if err != nil {
		return nil, err
	}
	return heros, nil
}
