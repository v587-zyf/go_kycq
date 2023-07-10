package managersI

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelGame"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type IMail interface {
	LoadMail(user *objs.User) ([]*modelGame.Mail, error)
	SendSystemMail(userId, mailId int, args []string, items model.Bag,itemSource int) error
	SendSystemMailWithItemInfosSignItemSource(userId, mailId int, args []string, itemMap gamedb.ItemInfos,itemSource int) error
	SendSystemMailWithItemInfos(userId, mailId int, args []string, itemMap gamedb.ItemInfos) error
	ReadMail(user *objs.User, mailId int) error
	MailRedeem(user *objs.User, mailId int, op *ophelper.OpBagHelperDefault, ack *pb.MailRedeemAck) error
	RedeemAllMail(user *objs.User, pbHelp *ophelper.OpBagHelperDefault, ack *pb.MailRedeemAllAck) error
	/**
    *  @Description: gm邮件
    *  @param req
    **/
	MailSendByGM(req *pbserver.MailSendCCsToGsReq)
}
