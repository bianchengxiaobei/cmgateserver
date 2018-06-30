package gateserver

import (
	"cmgateserver/msgHandler"
	"errors"
	"github.com/bianchengxiaobei/cmgo/network"
	"fmt"
)

type InnnerServerMessageHandler struct {
	server     network.ISocket
	gateServer *GateServer
	pool       *HandlerPool
}

func (handler InnnerServerMessageHandler) Init() {
	handler.pool.Register(10000, &msgHandler.RegisterGateHandler{GateServer: handler.gateServer})
	handler.pool.Register(10002,&msgHandler.LoginSuccessHandler{GateServer:handler.gateServer})
}

func (handler InnnerServerMessageHandler) MessageReceived(session network.SocketSessionInterface, message interface{}) error {
	if writeMsg, ok := message.(network.WriteMessage); !ok {
		return errors.New("不是WriteMessage类型")
	} else {
		action := handler.pool.GetHandler(int32(writeMsg.MsgId))
		if action == nil {
			//log4g.Errorf("找不到该Handler[%d]", writeMsg.MsgId)
			//return errors.New("找不到该Handler")
			//如果找不到handler说明是直接发给客户端的
			fmt.Println("jhjjffd")
			if innerMsg, ok := writeMsg.MsgData.(network.InnerWriteMessage); ok {
				if innerMsg.RoleId > 0 {
					//网关找到玩家session直接转发
					fmt.Println("jhjjffd")
					handler.gateServer.SendMsgToClientByRoleId(innerMsg.RoleId, writeMsg.MsgId, innerMsg.MsgData)
				}
			}
		} else {
			action.Action(session, writeMsg.MsgData)
		}
	}
	return nil
}

func (handler InnnerServerMessageHandler) MessageSent(session network.SocketSessionInterface, message interface{}) error {
	return nil
}

func (handler InnnerServerMessageHandler) SessionOpened(session network.SocketSessionInterface) error {
	return nil
}

func (handler InnnerServerMessageHandler) SessionClosed(session network.SocketSessionInterface) {

}

func (handler InnnerServerMessageHandler) SessionPeriod(session network.SocketSessionInterface) {

}

func (handler InnnerServerMessageHandler) ExceptionCaught(session network.SocketSessionInterface, err error) {

}
