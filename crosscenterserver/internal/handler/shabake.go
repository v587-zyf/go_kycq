package handler

import (
	"cqserver/golibs/nw"
	"cqserver/protobuf/pbserver"
)

func init() {
	pbserver.Register(pbserver.CmdGsToCcsBackGuildInfoNtfId, HandleBackGuildInfoNtf)
}

func HandleBackGuildInfoNtf(conn nw.Conn, msgFrame *pbserver.MessageFrame) (nw.ProtoMessage, error) {
	req := msgFrame.Body.(*pbserver.GsToCcsBackGuildInfoNtf)
	m.GetShaBakeCcs().BackGuildInfoNtf(req)
	return nil, nil
}
