package managersI

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/golibs/nw"
)

type IGsServers interface {
	GetServerInfo(serverId int) *modelCross.ServerInfo

	/**
	 *  @Description: 给指定服务器发消息
	 *  @param serverId
	 *  @param msg
	 *  @return error
	 */
	SendMessage(serverId int, msg nw.ProtoMessage) error

	/**
    *  @Description: 推送全服消息
    *  @param msg
    **/
	SendAllServerMessage(msg nw.ProtoMessage)

	/**
    *  @Description:	发送指定服务器消息
    *  @param serverId		服务器Id
    *  @param requestMsg	发送消息
    *  @param resultMsg		返回消息
    *  @return error
    **/
	CallMessage(serverId int,requestMsg nw.ProtoMessage,resultMsg nw.ProtoMessage)error
	CrossMatch()

	//
	//  GetAllCrossGroupServerInfo
	//  @Description: 获取跨服组下服务器信息
	//  @return map[int][]*modelCross.ServerInfo
	//
	GetAllCrossGroupServerInfo() map[int]map[int]*modelCross.ServerInfo

	GetCrossGroupServerInfoByCrossFsId(crossFsId int) map[int]*modelCross.ServerInfo

	UpServerListTicker()
}
