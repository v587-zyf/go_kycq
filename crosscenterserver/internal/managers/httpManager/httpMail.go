package httpManager

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/ptsdk"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pbserver"
	"net/http"
	"strconv"
	"sync"
)

const (
	MAIL_SEND_TYPE_SERVERS = 3
	MAIL_SEND_TYPE_SERVER  = 4
	MAIL_SEND_TYPE_ROLE    = 5
)

var mailMu sync.Mutex
/**
*  @Description: 发邮件
*  @param w
*  @param r
**/
func httpMailSend(w http.ResponseWriter, r *http.Request) {

	defer func() {
		mailMu.Unlock()
	}()
	mailMu.Lock()

	kyMail, err := ptsdk.GetSdk().MailSend(r)
	if len(err) > 0 {
		ptsdk.GetSdk().HttpWriteReturnMsg(w,err)
		return
	}

	request := &pbserver.MailSendCCsToGsReq{
		MailId:       kyMail.Id,
		Title:        kyMail.Title,
		Content:      kyMail.Content,
		ValidityDay:  int32(kyMail.ValidityDay),
		HighVip:      int32(kyMail.Filter.HighVip),
		LowVip:       int32(kyMail.Filter.LowVip),
		HighLevel:    int32(kyMail.Filter.HighDegree),
		LowLevel:     int32(kyMail.Filter.LowDegree),
		IsOnline:     kyMail.Filter.IsOnline,
		HighRecharge: int32(kyMail.Filter.HighConsumption),
		LowRecharge:  int32(kyMail.Filter.LowConsumption),
	}

	if len(kyMail.ItemList) > 0 {
		for _, v := range kyMail.ItemList {
			request.Items = append(request.Items, &pbserver.ItemUnit{
				ItemId:  int32(v.Id),
				ItemNum: int32(v.Count),
			})
		}
	}

	if kyMail.Maitype == MAIL_SEND_TYPE_ROLE {

		if len(kyMail.Target.RoleIds) <= 0 {

			ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "发送个人邮件，个人Id数据为空", nil)
			return
		}

		userIds := make([]int, 0)
		if len(kyMail.Target.RoleIds) > 0 {

			for _, v := range kyMail.Target.RoleIds {
				userId, err := strconv.Atoi(v)
				if err != nil {
					logger.Error("转化userId异常：%v,err:%v", v, err)
					continue
				}
				userIds = append(userIds, userId)
				request.UserIds = append(request.UserIds, int32(userId))
			}
		}

		serverIds := modelCross.GetUserCrossInfoModel().GetServerIds(userIds)
		if len(serverIds) > 0 {
			for _, v := range serverIds {
				err := m.GetGsServers().SendMessage(v, request)
				logger.Info("推送game服 发送邮件：%v,err:%v", kyMail.Id, kyMail.Title, request.UserIds, err)
			}
		}

	} else if kyMail.Maitype == MAIL_SEND_TYPE_SERVER || kyMail.Maitype == MAIL_SEND_TYPE_SERVERS {

		if len(kyMail.Target.ServerIds) <= 0 {
			ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "发送区服（组）邮件，区服（组）数据为空", nil)
			return
		}
		for _, v := range kyMail.Target.ServerIds {
			err := m.GetGsServers().SendMessage(v, request)
			logger.Info("推送game服 发送邮件：%v,err:%v", kyMail.Id, kyMail.Title, request.UserIds, err)
		}
	} else {
		ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "暂不支持的类型", nil)
		return
	}
	ptsdk.GetSdk().HttpWriteReturnInfo(w, 200, "success", nil)
}

/**
*  @Description: 测回邮件
*  @param w
*  @param r
**/
func httpMailRollback(w http.ResponseWriter, r *http.Request) {

}
