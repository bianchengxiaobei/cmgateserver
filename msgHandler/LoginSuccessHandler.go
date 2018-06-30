package msgHandler

import (
	"cmgateserver/message"
	"github.com/bianchengxiaobei/cmgo/log4g"
	"github.com/bianchengxiaobei/cmgo/network"
	"cmgateserver/face"
)

type LoginSuccessHandler struct {
	GateServer face.IGateServer
}

func (handler *LoginSuccessHandler) Action(session network.SocketSessionInterface, msg interface{}) {
	if innerMsg,ok := msg.(network.InnerWriteMessage);ok {
		if protoMsg, ok := innerMsg.MsgData.(*message.M2G_LoginSuccessNotifyGate); ok {
			handler.GateServer.GetRoleManager().RegisterRole(protoMsg.ServerId,protoMsg.UserId,protoMsg.RoleId)
			//发送给游戏逻辑服注册成功
			rMsg := &message.G2M_RoleRegisterToGateSuccess{}
			rMsg.RoleId = protoMsg.RoleId
			handler.GateServer.SendMsgToGameServerByRoleId(protoMsg.RoleId,10003,rMsg)
		} else {
			log4g.Error("不是M2G_LoginSuccessNotifyGate！")
		}
	}
}
