package mail

import (
	"cqserver/gamelibs/modelGame"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"time"

	"cqserver/golibs/util"
)

type MailManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewMailManager(module managersI.IModule) *MailManager {
	return &MailManager{IModule: module}
}

func (this *MailManager) LoadMail(user *objs.User) ([]*modelGame.Mail, error) {
	mailModel := modelGame.GetMailModel()
	mails, err := mailModel.GetMailList(user.Id)
	if err != nil {
		logger.Error("加载个人邮件错误 %v", err.Error())
		return nil, err
	}
	// 自动删除过期邮件
	now := time.Now()
	updateMails := make([]*modelGame.Mail, 0)
	sendMails := make([]*modelGame.Mail, 0)
	for _, mail := range mails {
		if mail.ExpireAt.Unix() < now.Unix() {
			if mail.DeletedAt.IsZero() {
				mail.DeletedAt = now
				updateMails = append(updateMails, mail)
			}
			continue
		}
		sendMails = append(sendMails, mail)
	}
	err = mailModel.Update(updateMails...)
	if err != nil {
		return nil, err
	}

	return sendMails, nil
}

func (this *MailManager) ReadMail(user *objs.User, mailId int) error {
	mailModel := modelGame.GetMailModel()
	mail, err := mailModel.GetMailById(mailId, user.Id)
	if err != nil {
		return err
	}
	mail.Status = pb.MAILSTATUS_READ
	// 如果没奖励 更新领取
	if len(mail.Items) <= 0 {
		mail.RedeemedAt = time.Now()
	}
	err = mailModel.Update(mail)
	if err != nil {
		return err
	}
	return err
}
