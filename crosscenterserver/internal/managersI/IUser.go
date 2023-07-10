package managersI

import "cqserver/protobuf/pbserver"

type IUser interface {
	UserInfoSync(message *pbserver.SyncUserInfoNtf)
}