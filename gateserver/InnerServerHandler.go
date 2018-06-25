package gateserver

import (
	"github.com/bianchengxiaobei/cmgo/network"
)
type InnnerServerMessageHandler struct {
	server network.ISocket
	gateServer *GateServer
}

func (handler InnnerServerMessageHandler) MessageReceived(session network.SocketSessionInterface, message interface{}) error {
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


