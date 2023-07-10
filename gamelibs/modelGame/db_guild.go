package modelGame

import (
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelCross"
	"fmt"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gorp.v1"
	"time"

	"cqserver/golibs/dbmodel"
)

type Guild struct {
	Id                     int            `db:"id"`
	GuildId                int            ` db:"guildId" orm:"comment(门派id)"`
	GuildName              string         `db:"guildName" orm:"comment(门派名字)"`
	SettingId              int            `db:"settingId" orm:"comment(0-3 默认门派, 4或其它自建门派)"` //0-3 默认门派, 4或其它自建门派
	CrossGroup             string         `db:"crossGroup" `                                 //所属跨服组
	WeChat                 string         `db:"weChat" orm:"comment(wx)"`
	LogoId                 int            `db:"logoId"   orm:"comment(公会旗帜icon)" `
	AutoAgreeJoin          int            `db:"autoAgreeJoin"`
	Positions              model.IntSlice `db:"positions"   orm:"type(text);comment(玩家职位)" ` //玩家id,职位
	ChairmanId             int            `db:"chairmanId" orm:"comment(掌门id)" `
	CanChallenge           int            `db:"canChallenge"  orm:"comment(是否可挑战掌门)"`
	GuildContributionValue int            `db:"guildContributionValue"orm:"comment(门派贡献值)"`
	Notice                 string         `db:"notice" orm:"size(300);comment(门派公告)"`
	CreatedAt              time.Time      `db:"createdAt" orm:"comment(门派创建时间)"`
	IsDelete               int            `db:"isDelete"   orm:"comment(是否解散 0:正常 1:势力长解散 2:自动解散 3:被合并 4:被吞并)"`
	DeletedAt              time.Time      `db:"deletedAt" orm:"comment(门派解散时间)"`
	Creator                int            `db:"creator"  orm:"comment(门派创建人)"`
	JoinCombat             int            `db:"joinCombat"  orm:"comment(加入门派需要的战力)"` //战斗力
	ApplyList              model.IntKv    `db:"applyList" orm:"comment(门派申请加入列表)"`
	AutoAgree              int            `db:"autoAgree" orm:"comment(门派自动同意加入)"`
	DonateUsers            model.IntSlice `db:"donateUsers" orm:"size(500);comment(捐献玩家信息 userId,捐献类型1,userId1,捐献类型2)" `
	DonateTimes            model.IntKv    `db:"donateTimes" orm:"size(100);comment(门派一共捐献次数 k:type v:times)" `
	ServerId               int            `db:"serverId"`
	IsSystem               int            `db:"isSystem"`
}

func (this *Guild) TableName() string {
	return "guild"
}

type GuildModel struct {
	dbmodel.CommonModel
}

var (
	guildModel  = &GuildModel{}
	guildFields = model.GetAllFieldsAsString(Guild{})
	idSeqGuild  = &modelCross.IdSeq{Name: "guild"}
)

func init() {
	dbmodel.Register(model.DB_SERVER, guildModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(Guild{}, "guild").SetKeys(true, "Id")
		orm.RegisterModelForAlias(model.DB_SERVER, new(Guild))
	})
}

func GetGuildModel() *GuildModel {
	return guildModel
}

func (this *GuildModel) Create(guild *Guild) error {
	var err error
	guild.GuildId, err = idSeqGuild.Next()
	if err != nil {
		return err
	}
	return this.DbMap().Insert(guild)
}

func (this *GuildModel) GetGuildInfo(guildId int) (*Guild, error) {

	var guild Guild
	err := this.DbMap().SelectOne(&guild, fmt.Sprintf("select %s from guild where guildId = ? and isDelete = 0", guildFields), guildId)
	if err != nil {
		return nil, err
	}
	return &guild, nil
}

func (this *GuildModel) GetAllGuildInfoIncludeDel(guildId int) (*Guild, error) {

	var guild Guild
	err := this.DbMap().SelectOne(&guild, fmt.Sprintf("select %s from guild where guildId = ? ", guildFields), guildId)
	if err != nil {
		return nil, err
	}
	return &guild, nil
}

func (this *GuildModel) GetAllGuildInfos() ([]*Guild, error) {

	var guilds []*Guild
	_, err := this.DbMap().Select(&guilds, fmt.Sprintf("select %s from guild where isDelete = 0", guildFields))
	if err != nil {
		return nil, err
	}
	return guilds, nil
}

func (this *GuildModel) GetAllGuildInfo() ([]*Guild, error) {

	var guild []*Guild
	_, err := this.DbMap().Select(&guild, fmt.Sprintf("select %s from guild where isDelete = 0", guildFields))
	if err != nil {
		return nil, err
	}
	return guild, nil
}

func (this *GuildModel) GetAllGuildInfoByGuildName(guildName string) ([]*Guild, error) {

	var guild []*Guild
	_, err := this.DbMap().Select(&guild, fmt.Sprintf("select %s from guild where isDelete = 0 and guildName = ?", guildFields), guildName)
	if err != nil {
		return nil, err
	}
	return guild, nil
}

func (this *GuildModel) Update(guild *Guild) error {
	_, err := this.DbMap().Update(guild)
	return err
}

func (this *GuildModel) GetAllGuildInfoBySystem() ([]*Guild, error) {

	var guild []*Guild
	_, err := this.DbMap().Select(&guild, fmt.Sprintf("select %s from guild where isDelete = 0 and isSystem = 1", guildFields))
	if err != nil {
		return nil, err
	}
	return guild, nil
}

func (this *GuildModel) DeleteGuild(guild *Guild) error {
	_, err := this.DbMap().Delete(guild)
	return err
}
