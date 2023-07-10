package modelGame

import (
	"cqserver/gamelibs/model"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"

	"gopkg.in/gorp.v1"

	"cqserver/golibs/dbmodel"
	"cqserver/protobuf/pb"
)

type Mail struct {
	Id         int               `db:"id" orm:"pk;auto"`
	UserId     int               `db:"userId" orm:"size(11);comment(羽翼信息)"`
	MailID     int               `db:"mailId" orm:"size(5);comment(邮件配置Id)"`
	Sender     string            `db:"sender" orm:"size(255);comment(发送者)"` //发送者
	Title      string            `db:"title" orm:"size(255);comment(标题)"`
	Content    string            `db:"content" orm:"size(500);comment(内容)"`
	Args       model.StringSlice `db:"args" orm:"size(255);comment(替换参数)"` //可替换的内容
	Items      model.Bag         `db:"items" orm:"size(512);comment(道具)"`  //包括的内容j
	Status     int               `db:"status" orm:"size(5);comment(状态)"`   //  状态，参见->enums.MailStatus
	CreatedAt  time.Time         `db:"createdAt" orm:"comment(创建时间)"`      //创建时间j
	ExpireAt   time.Time         `db:"expireAt" orm:"comment(过期时间)"`       //过期时间
	RedeemedAt time.Time         `db:"redeemedAt" orm:"comment(道具领取时间)"`   //兑换的时间点j
	DeletedAt  time.Time         `db:"deletedAt" orm:"comment(删除时间)"`      //删除时间点j
	GmMailId   string            `db:"gmMailId" orm:"size(50);comment(GM id)"`
	ItemSource int               `db:"itemSource" orm:"type(int64);comment(道具来源，用于背包满的情况下)"`
}

func (this *Mail) TableName() string {
	return "mail"
}

type MailModel struct {
	dbmodel.CommonModel
}

var (
	mailModel  = &MailModel{}
	mailFields = model.GetAllFieldsAsString(Mail{})
)

func init() {
	dbmodel.Register(model.DB_SERVER, mailModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(Mail{}, "mail").SetKeys(true, "Id")
		orm.RegisterModelForAlias(model.DB_SERVER, new(Mail))
	})
}

func GetMailModel() *MailModel {
	return mailModel
}

func convertMailSlice2InterfaceSlice(mails []*Mail) []interface{} {
	mailInterfaces := make([]interface{}, len(mails))
	for i, v := range mails {
		mailInterfaces[i] = v
	}
	return mailInterfaces
}

func (this *MailModel) Create(mails ...*Mail) error {
	return this.DbMap().Insert(convertMailSlice2InterfaceSlice(mails)...)
}

func (this *MailModel) Update(mails ...*Mail) error {
	_, err := this.DbMap().Update(convertMailSlice2InterfaceSlice(mails)...)
	return err
}

func (this *MailModel) GetMailList(userId int) ([]*Mail, error) {
	var mails []*Mail
	_, err := this.DbMap().Select(&mails, fmt.Sprintf("select %s from mail where userId = ? and deletedAt = 0 order by createdAt desc limit 200", mailFields), userId)
	return mails, err
}

func (this *MailModel) GetMailById(mailId, userId int) (*Mail, error) {
	var mail Mail
	err := this.DbMap().SelectOne(&mail, fmt.Sprintf("select %s from mail where userId=? and id = ? and deletedAt = 0", mailFields), userId, mailId)
	return &mail, err
}

//得到所有的可以领取的邮件。最后300封
func (this *MailModel) GetRedeemableOrUnreadMailList(userId int) ([]*Mail, error) {
	var mails []*Mail
	_, err := this.DbMap().Select(&mails, fmt.Sprintf("select %s from mail where userId = ? and deletedAt = 0 and ((redeemedAt = 0 and items != '{}') or status=%d) order by createdAt desc limit 200", mailFields, pb.MAILSTATUS_UNREAD), userId)
	return mails, err
}

func (this *MailModel) DeleteMail(mail *Mail) error {
	_, err := this.DbMap().Delete(mail)
	return err
}

func (this *MailModel) AutoDeleteByUser(userId int) error {
	_, err := this.DbMap().Exec("update mail set deletedAt=Now() where userId=? and deletedAt=0 and ((status=? and items=?) or redeemedAt>0)", userId, pb.MAILSTATUS_READ, "{}")
	return err
}
