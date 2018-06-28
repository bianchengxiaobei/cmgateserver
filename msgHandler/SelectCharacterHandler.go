package msgHandler

import (
	"github.com/bianchengxiaobei/cmgo/network"
	"cmgateserver/message"
	"github.com/bianchengxiaobei/cmgo/log4g"
)

type SelectCharacterHandler struct {
	GateServer IGateServer
}

func (handler *SelectCharacterHandler) Action(session network.SocketSessionInterface,msg interface{}) {
	if protoMsg,ok := msg.(*message.C2G_SelectCharacter);ok{
		userId,_ := session.GetAttribute(network.USERID).(int64)
		userName,_ := session.GetAttribute(network.USERNAME).(string)
		serverId,_ := session.GetAttribute(network.SERVERID).(int32)
		//发送给游戏逻辑服登录玩家角色
		msg := new(message.G2M_LoginToGameServer)
		log4g.Info("Login")
		msg.RoleId = protoMsg.RoleId
		msg.UserId = userId
		msg.UserName = userName
		msg.ServerId = serverId
		handler.GateServer.SendMsgToGameServer(msg.ServerId,10001,msg)
	}
}