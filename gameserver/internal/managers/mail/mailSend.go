package mail

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gameserver/internal/builder"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"errors"
	"fmt"
	"time"
)

func (this *MailManager) SendSystemMail(userId, mailId int, args []string, items model.Bag, itemSource int) error {

	//玩家离线3天，不再发邮件
	userInfo := this.GetUserManager().GetUserBasicInfo(userId)
	if userInfo == nil {
		return gamedb.ERROPENIDGETERR
	}
	if time.Now().Sub(userInfo.LastUpdateTime) > 31*24*time.Hour {
		logger.Error("send systemMail user:%v,userLastupdatime:%v", userId, userInfo.LastUpdateTime)
		return errors.New(fmt.Sprintf("send systemMail user:%v,userLastupdatime:%v", userId, userInfo.LastUpdateTime))
	}

	mailConf := gamedb.GetMailMailCfg(mailId)
	if mailConf == nil {
		logger.Info("sendSystemMail:no mailconfig userId:%v mailId:%v", userId, mailId)
		return errors.New(fmt.Sprintf("sendSystemMail:no mailconfig for type:%d", mailId))
	}
	if args == nil { //?
		args = make([]string, 0)
	}
	this.SendMailByInfo(userId, mailId, "", "", "", args, items, "", itemSource)

	return nil
}

func (this *MailManager) SendMailByInfo(userId, mailId int, sender, title, content string, args []string, items model.Bag, gmMailId string, itemSource int) error {
	mailConf := gamedb.GetMailMailCfg(mailId)
	if len(title) <= 0 {
		title = mailConf.Title
	}
	if len(content) <= 0 {
		content = mailConf.Content
	}
	if len(sender) <= 0 {
		sender = mailConf.FromName
	}
	mail := &modelGame.Mail{
		UserId:     userId,
		MailID:     mailId,
		Sender:     sender,
		Title:      title,
		Content:    content,
		Status:     pb.MAILSTATUS_UNREAD,
		ExpireAt:   time.Now().AddDate(0, 0, mailConf.ExpireDays),
		CreatedAt:  time.Now(),
		Args:       args,
		ItemSource: itemSource,
	}
	if items == nil || len(items) <= 12 {
		if items != nil && len(items) > 0 {
			mail.Items = items
		}

		mailModel := modelGame.GetMailModel()
		err := mailModel.Create(mail)
		if err != nil {
			logger.Error("sendSystemMail:send mail err:", err)
			return err
		}
		this.GetUserManager().SendMessageByUserId(userId, builder.BuildMailNtf(mail))
		logger.Info("sendSystemMail: done:userId:%d,mailId:%d", userId, mailId)
	} else {

		mailIndex := 1
		mailTitle := mail.Title + "(%d)"
		for i := 0; i < len(items); i++ {

			mail.Items = append(mail.Items, items[i])
			if (i+1)%12 == 0 || i == len(items)-1 {

				mail.Title = fmt.Sprintf(mailTitle, mailIndex)
				mailModel := modelGame.GetMailModel()
				err := mailModel.Create(mail)
				if err != nil {
					logger.Error("sendSystemMail:send mail err:", err)
					return err
				}
				this.GetUserManager().SendMessageByUserId(userId, builder.BuildMailNtf(mail))
				logger.Info("sendSystemMail: done:userId:%d,mailId:%d", userId, mailId)
				mail.Items = make(model.Bag, 0)
				mailIndex += 1
			}
		}
	}
	return nil
}

func (this *MailManager) SendSystemMailWithItemInfosSignItemSource(userId, mailId int, args []string, itemMap gamedb.ItemInfos, itemSource int) error {
	l := len(itemMap)
	tempItemMap := make(map[int]*gamedb.ItemInfo)
	if l > 0 {
		for _, v := range itemMap {
			if tempItemMap[v.ItemId] != nil {
				tempItemMap[v.ItemId].Count += v.Count
			} else {
				tempItemMap[v.ItemId] = &gamedb.ItemInfo{ItemId: v.ItemId, Count: v.Count}
			}
		}
	}
	items := make(model.Bag, 0)
	if len(tempItemMap) > 0 {
		for _, v := range tempItemMap {
			items = append(items, &model.Item{ItemId: v.ItemId, Count: v.Count})
		}
	}
	return this.SendSystemMail(userId, mailId, args, items, itemSource)
}

func (this *MailManager) SendSystemMailWithItemInfos(userId, mailId int, args []string, itemMap gamedb.ItemInfos) error {
	return this.SendSystemMailWithItemInfosSignItemSource(userId, mailId, args, itemMap, 0)
}

/**
*  @Description: 人工发放邮件
*  @receiver this
*  @param req
**/
func (this *MailManager) MailSendByGM(req *pbserver.MailSendCCsToGsReq) {

	var items model.Bag
	if len(req.Items) > 0 {
		items = make(model.Bag, len(req.Items))
		for k, v := range req.Items {
			items[k] = &model.Item{
				ItemId: int(v.ItemId),
				Count:  int(v.ItemNum),
			}
		}
	}
	if req.UserIds != nil && len(req.UserIds) > 0 {

		for _, v := range req.UserIds {
			this.SendMailByInfo(int(v), constMail.GM_MAIL, "", req.Title, req.Content, []string{}, items, req.MailId, 0)
		}
	} else {
		//获取指定条件玩家数据
		userIds := modelGame.GetUserModel().GetUserIdsForGmMail(int(req.HighVip), int(req.LowVip), int(req.HighLevel), int(req.LowLevel), int(req.HighRecharge), int(req.LowRecharge))
		if len(userIds) > 0 {
			for _, v := range userIds {
				if req.IsOnline {
					if this.GetUserManager().GetUser(v) == nil {
						continue
					}
				}
				this.SendMailByInfo(v, constMail.GM_MAIL, "", req.Title, req.Content, []string{}, items, req.MailId, 0)
			}
		}
	}
}
