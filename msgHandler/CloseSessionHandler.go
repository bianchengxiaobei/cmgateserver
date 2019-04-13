package msgHandler

import (
	"cmgateserver/face"
	"github.com/bianchengxiaobei/cmgo/network"
	"cmgateserver/message"
	"github.com/bianchengxiaobei/cmgo/log4g"
)

type CloseSessionHandler struct {
	GateServer face.IGateServer
}

func (handler *CloseSessionHandler) Action(session network.SocketSessionInterface, msg interface{}) {
	if innerMsg,ok := msg.(network.InnerWriteMessage);ok {
		if protoMsg, ok := innerMsg.MsgData.(*message.M2G_CloseSession); ok {
			session := handler.GateServer.GetRoleSession(protoMsg.RoleId)
			if session != nil {
				if session.IsClosed() == false{
					//log4g.Infof("玩家[%d]关闭Session[%d]", protoMsg.RoleId, handler.GateServer.GetId())
					//rMsg := &message.G2C_QuitGame{}
					//handler.GateServer.SendMsgToClient(session,1003,rMsg)
					session.Close(0)
				}else{
					session.Close(0)
					log4g.Info("OtherClose")
				}
			}
		} else {
			log4g.Error("不是M2G_CloseSession！")
		}
	}
}
