package shabake

import (
	"cqserver/crosscenterserver/internal/managersI"
	"cqserver/gamelibs/modelCross"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pbserver"
)

type ShaBake struct {
	util.DefaultModule
	managersI.IModule
}

func NewShaBakeManager(m managersI.IModule) *ShaBake {
	return &ShaBake{
		IModule: m,
	}
}

func (this *ShaBake) BackGuildInfoNtf(message *pbserver.GsToCcsBackGuildInfoNtf) {

	infos, err := modelCross.GetServerInfoModel().GetAllServerIdsByCrossFsIds(int(message.CrossFsId))
	logger.Info("沙巴克获胜门派展示信息存储 message.CrossFsId:%v, infos:%v  err:%v", message.CrossFsId, infos, err)
	if err == nil {
		for _, info := range infos {
			_ = this.GetGsServers().SendMessage(info.ServerId, &pbserver.CcsToGsBroadShaBakeFirstGuildInfo{FirstGuildInfo: message.FirstGuildInfo, BenFuShaBaKe: message.BenFuShaBaKe})
		}
	}

}
