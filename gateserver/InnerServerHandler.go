package gateserver

import (
	"github.com/bianchengxiaobei/cmgo/network"
	"errors"
	"cmgateserver/msgHandler"
)
type InnnerServerMessageHandler struct {
	server network.ISocket
	gateServer *GateServer
	pool	*HandlerPool
}

func (handler InnnerServerMessageHandler)Init() {
	handler.pool.Register(10000,&msgHandler.RegisterGateHandler{GateServer:handler.gateServer,})
}


func (handler InnnerServerMessageHandler) MessageReceived(session network.SocketSessionInterface, message interface{}) error {
	if writeMsg,ok := message.(network.WriteMessage);!ok{
		return errors.New("不是WriteMessage类型")
	}else{
		action := handler.pool.GetHandler(int32(writeMsg.MsgId))
		if action == nil{
			return errors.New("找不到该Handler")
		}else{
			action.Action(session,writeMsg.MsgData)
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


