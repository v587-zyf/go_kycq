package managersI

import "cqserver/protobuf/pbserver"

type IShaBake interface {


	BackGuildInfoNtf(message *pbserver.GsToCcsBackGuildInfoNtf)

}
