package msgHandler

import (
	"github.com/bianchengxiaobei/cmgo/network"
	"cmgateserver/message"
	"github.com/bianchengxiaobei/cmgo/log4g"
)

type RegisterGateHandler struct {
	GateServer IGateServer
}

func (handler *RegisterGateHandler) Action(session network.SocketSessionInterface,msg interface{}) {
	if protoMsg,ok := msg.(*message.M2G_RegisterGate);ok{
		handler.GateServer.RegisterInnerGameServer(protoMsg.Id,session)
		log4g.Infof("游戏服务器[%d]成功注册到网关服务器[%d]",protoMsg.Id,handler.GateServer.GetId())
	}
}

