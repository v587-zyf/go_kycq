package builder

import (
	"cqserver/protobuf/pbserver"
)

func BuildLoginVerifyReq(openId string, loginKey string) *pbserver.LoginKeyVerifyReq {
	return &pbserver.LoginKeyVerifyReq{
		OpenId:   openId,
		LoginKey: loginKey,
	}
}
