package handler

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"time"

	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"

	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdMailReadReqId, HandleMailReadReq)
	pb.Register(pb.CmdMailRedeemReqId, HandleMailRedeemReq)
	pb.Register(pb.CmdMailLoadReqId, HandleMailLoadReq)
	pb.Register(pb.CmdMailRedeemAllReqId, HandleMailRedeemAllReq)
	pb.Register(pb.CmdMailDeleteReqId, HandleMailDeleteReq)
	pb.Register(pb.CmdMailDeleteAllReqId, HandleMailDeleteAllReq)
}

func HandleMailReadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.MailReadReq)
	user := conn.GetSession().(*managers.ClientSession).User
	err := m.Mail.ReadMail(user, int(req.Id))
	return &pb.MailReadAck{Id: req.Id}, nil, err
}

func HandleMailRedeemReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.MailRedeemReq)
	user := conn.GetSession().((*managers.ClientSession)).User

	ack := &pb.MailRedeemAck{}
	pbHelp := ophelper.NewOpBagHelperDefault(constBag.OpTypeMail)
	err := m.Mail.MailRedeem(user, int(req.Id), pbHelp, ack)
	if err != nil {
		return nil, nil, err
	}
	return ack, pbHelp, nil
}

func HandleMailLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().((*managers.ClientSession)).User

	mails, err := m.Mail.LoadMail(user)
	if err != nil {
		return nil, nil, err
	}

	logger.Info("个人邮件数量 玩家：%v,mailCount=%v", user.IdName(), len(mails))
	return &pb.MailLoadAck{Mails: builder.BuildMailNtfs(mails)}, nil, nil
}

func HandleMailRedeemAllReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().((*managers.ClientSession)).User

	ack := &pb.MailRedeemAllAck{}
	pbHelp := ophelper.NewOpBagHelperDefault(constBag.OpTypeMailAll)
	err := m.Mail.RedeemAllMail(user, pbHelp, ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, pbHelp, nil
}

func HandleMailDeleteReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.MailDeleteReq)
	user := conn.GetSession().((*managers.ClientSession)).User
	id := int(req.Id)
	mailModel := modelGame.GetMailModel()
	mail, err := mailModel.GetMailById(id, user.Id)
	if err != nil {
		return nil, nil, err
	}
	if mail == nil {
		return nil, nil, gamedb.ERRPARAM
	}
	if len(mail.Items) > 0 && mail.RedeemedAt.IsZero() {
		return nil, nil, gamedb.ERRUNKNOW
	}
	if mail.DeletedAt.IsZero() {
		mail.DeletedAt = time.Now()
		err := mailModel.Update(mail)
		if err != nil {
			return nil, nil, err
		}
	}
	return &pb.MailDeleteAck{Mail: builder.BuildMailNtf(mail)}, nil, nil
}

func HandleMailDeleteAllReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().((*managers.ClientSession)).User
	mailModel := modelGame.GetMailModel()
	// 先删除
	err := mailModel.AutoDeleteByUser(user.Id)
	if err != nil {
		return nil, nil, err
	}
	// 拉取剩下的
	mails, err := modelGame.GetMailModel().GetMailList(user.Id)
	if err != nil {
		logger.Error("加载个人邮件错误 %v", err.Error())
		return nil, nil, err
	}

	//m.Tlog.MailSystem(user, tlog.IMAILACTIONTYPE_TYPE_3)
	return &pb.MailDeleteAllAck{Mails: builder.BuildMailNtfs(mails)}, nil, nil
}
