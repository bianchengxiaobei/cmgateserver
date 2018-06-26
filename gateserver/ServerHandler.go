package gateserver

import (
	"github.com/bianchengxiaobei/cmgo/network"
	"github.com/bianchengxiaobei/cmgo/log4g"
	"errors"
)

type ServerMessageHandler struct {
	server network.ISocket
	gateServer *GateServer
	pool	*HandlerPool
}

func (handler ServerMessageHandler)Init() {

}

func (handler ServerMessageHandler) MessageReceived(session network.SocketSessionInterface, message interface{}) error {
	if writeMsg,ok := message.(network.WriteMessage);!ok{
		return errors.New("不是WriteMessage类型")
	}else{
		if handler.pool.GetHandler(int32(writeMsg.MsgId)) == nil{

		}
	}
	return nil
}

func (handler ServerMessageHandler) MessageSent(session network.SocketSessionInterface, message interface{}) error {
	return nil
}

func (handler ServerMessageHandler) SessionOpened(session network.SocketSessionInterface) error {
	if server,ok := handler.server.(*network.TcpServer);ok{
		log4g.Infof("Session总数:%d",len(server.Sessions))
	}
	return nil
}

func (handler ServerMessageHandler) SessionClosed(session network.SocketSessionInterface) {

}

func (handler ServerMessageHandler) SessionPeriod(session network.SocketSessionInterface) {

}

func (handler ServerMessageHandler) ExceptionCaught(session network.SocketSessionInterface, err error) {

}


