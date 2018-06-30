package msgHandler

import (
	"cmgateserver/message"
	"github.com/bianchengxiaobei/cmgo/log4g"
	"github.com/bianchengxiaobei/cmgo/network"
	"cmgateserver/face"
)

type RegisterGateHandler struct {
	GateServer face.IGateServer
}

func (handler *RegisterGateHandler) Action(session network.SocketSessionInterface, msg interface{}) {
	if innerMsg,ok := msg.(network.InnerWriteMessage);ok {
		if protoMsg, ok := innerMsg.MsgData.(*message.M2G_RegisterGate); ok {
			handler.GateServer.RegisterInnerGameServer(protoMsg.Id, session)
			log4g.Infof("游戏服务器[%d]成功注册到网关服务器[%d]", protoMsg.Id, handler.GateServer.GetId())
		} else {
			log4g.Error("不是M2G_RegisterGate！")
		}
	}
}
