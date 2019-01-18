package msgHandler

import (
	"cmgateserver/face"
	"github.com/bianchengxiaobei/cmgo/network"
	"cmgateserver/message"
	"github.com/bianchengxiaobei/cmgo/log4g"
)

type RoleQuitHandler struct {
	GateServer face.IGateServer
}

func (handler *RoleQuitHandler) Action(session network.SocketSessionInterface, msg interface{}) {
	if innerMsg,ok := msg.(network.InnerWriteMessage);ok {
		if protoMsg, ok := innerMsg.MsgData.(*message.M2G_RoleQuitGate); ok {
			handler.GateServer.GetRoleManager().QuitRoleWithClearCache(protoMsg.RoleId)
			log4g.Infof("玩家[%d]退出网关服务器[%d]", protoMsg.RoleId, handler.GateServer.GetId())
		} else {
			log4g.Error("不是M2G_RoleQuitGate！")
		}
	}
}


