package mail

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"time"
)

func (this *MailManager) MailRedeem(user *objs.User, mailId int, op *ophelper.OpBagHelperDefault, ack *pb.MailRedeemAck) error {

	mailModel := modelGame.GetMailModel()
	mail, err := mailModel.GetMailById(mailId, user.Id)
	if err != nil {
		return err
	}
	if len(mail.Items) < 1 || !mail.RedeemedAt.IsZero() {
		return gamedb.ERRPARAM
	}

	itemInfos := make([]*gamedb.ItemInfo, len(mail.Items))
	for k, v := range mail.Items {
		itemInfos[k] = &gamedb.ItemInfo{
			ItemId: v.ItemId,
			Count:  v.Count,
		}
	}
	if !this.GetBag().CheckHasEnoughPos(user, itemInfos) {
		return gamedb.ERRBAGENOUGH
	}

	mail.RedeemedAt = time.Now()
	err = mailModel.Update(mail)
	if err != nil {
		return err
	}

	op.SetOpType(ophelper.GetMailOpType(mail.MailID))
	op.SetOpTypeSecond(mail.ItemSource)
	for _, item := range mail.Items {
		//err := user.addGoods(item.ItemIndex, item.Count, pbHelp)
		if item.Count == 0 {
			logger.Error("邮件物品奖励数量为空：%v，邮件：%v,道具：%v", user.Id, mail.Id, item.ItemId)
			continue
		}
		err := this.GetBag().Add(user, op, item.ItemId, item.Count)
		if err != nil {
			return err
		}

	}
	ack.Id = int32(mailId)
	ack.GoodsChanges = op.ToChangeItems()
	ack.Mail = builder.BuildMailNtf(mail)
	return nil
}

func (this *MailManager) RedeemAllMail(user *objs.User, pbHelp *ophelper.OpBagHelperDefault, ack *pb.MailRedeemAllAck) error {
	mails, err := modelGame.GetMailModel().GetRedeemableOrUnreadMailList(user.Id)

	if err != nil {
		logger.Error("load mail error: %v", err.Error())
		return err
	}

	if len(mails) < 1 {
		logger.Info("handler_mail:no mail to redeem")
		return nil
	}

	changedIds := make([]int32, 0)
	changeMails := make([]*pb.MailNtf, 0)
	for _, mail := range mails {
		if !mail.RedeemedAt.IsZero() {
			continue
		}

		if len(mail.Items) > 0 {
			mail.RedeemedAt = time.Now()
		} else {
			continue
		}

		itemInfos := make([]*gamedb.ItemInfo, len(mail.Items))
		for k, v := range mail.Items {
			itemInfos[k] = &gamedb.ItemInfo{
				ItemId: v.ItemId,
				Count:  v.Count,
			}
		}

		if !this.GetBag().CheckHasEnoughPos(user, itemInfos) {
			break
		}

		mail.Status = pb.MAILSTATUS_READ
		err = modelGame.GetMailModel().Update(mail)
		if err != nil {
			logger.Error("更新玩家邮件异常，user:%,mail:%v,err:%v", user.IdName(), mail.Id, err)
			return err
		}

		pbHelp.SetOpType(ophelper.GetMailOpType(mail.MailID))
		pbHelp.SetOpTypeSecond(mail.ItemSource)
		for _, item := range mail.Items {
			if item.Count == 0 {
				logger.Error("邮件物品奖励数量为空：%v，邮件：%v,道具：%v", user.Id, mail.Id, item.ItemId)
				continue
			}
			err := this.GetBag().Add(user, pbHelp, item.ItemId, item.Count)
			if err != nil {
				return err
			}
		}
		changedIds = append(changedIds, int32(mail.Id))
		changeMails = append(changeMails, builder.BuildMailNtf(mail))
	}
	ack.Ids = changedIds
	ack.Mail = changeMails
	ack.GoodsChanges = pbHelp.ToChangeItems()
	//m.Tlog.MailSystem(user, tlog.IMAILACTIONTYPE_TYPE_2)
	return nil
}
