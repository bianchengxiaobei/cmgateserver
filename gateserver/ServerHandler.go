package gateserver

import (
	"github.com/bianchengxiaobei/cmgo/network"
	"fmt"
)

type ServerMessageHandler struct {
	server network.ISocket
	gateServer *GateServer
}

func (handler ServerMessageHandler) MessageReceived(session network.SocketSessionInterface, message interface{}) error {
	return nil
}

func (handler ServerMessageHandler) MessageSent(session network.SocketSessionInterface, message interface{}) error {
	return nil
}

func (handler ServerMessageHandler) SessionOpened(session network.SocketSessionInterface) error {
	if server,ok := handler.server.(*network.TcpServer);ok{
		fmt.Printf("Session总数:%d",len(server.Sessions))
	}
	return nil
}

func (handler ServerMessageHandler) SessionClosed(session network.SocketSessionInterface) {

}

func (handler ServerMessageHandler) SessionPeriod(session network.SocketSessionInterface) {

}

func (handler ServerMessageHandler) ExceptionCaught(session network.SocketSessionInterface, err error) {

}


