package gateserver

import (
	"cmgateserver/msgHandler"
	"errors"
	"github.com/bianchengxiaobei/cmgo/log4g"
	"github.com/bianchengxiaobei/cmgo/network"
	"github.com/golang/protobuf/proto"
)

type ServerMessageHandler struct {
	server     network.ISocket
	gateServer *GateServer
	pool       *HandlerPool
}

func (handler ServerMessageHandler) Init() {
	handler.pool.Register(1000, &msgHandler.UserLoginHandler{GateServer: handler.gateServer})
	handler.pool.Register(1002, &msgHandler.SelectCharacterHandler{GateServer: handler.gateServer})
}

func (handler ServerMessageHandler) MessageReceived(session network.SocketSessionInterface, message interface{}) error {

	if writeMsg, ok := message.(network.WriteMessage); !ok {
		return errors.New("不是WriteMessage类型")
	} else {
		//log4g.Infof("收到消息%d",writeMsg.MsgId)
		msgHandler := handler.pool.GetHandler(int32(writeMsg.MsgId))
		if msgHandler == nil {
			//说明是直接发给游戏服务器的
			roleId := session.GetAttribute(network.ROLEID).(int64)
			protoMsg := writeMsg.MsgData.(proto.Message)
			handler.gateServer.SendMsgToGameServerByRoleId(roleId, writeMsg.MsgId, protoMsg)
		} else {
			msgHandler.Action(session, writeMsg.MsgData)
		}
	}
	return nil
}

func (handler ServerMessageHandler) MessageSent(session network.SocketSessionInterface, message interface{}) error {

	return nil
}

func (handler ServerMessageHandler) SessionOpened(session network.SocketSessionInterface) error {
	if server, ok := handler.server.(*network.TcpServer); ok {
		log4g.Infof("Session总数:%d", len(server.Sessions))
	}
	return nil
}

func (handler ServerMessageHandler) SessionClosed(session network.SocketSessionInterface) {
	//客户端断开连接
	//通知游戏服务器断开连接
	defer func() {
		if err := recover();err != nil{
			log4g.Info(err.(error).Error())
		}
	}()
	quit := false
	roleId := session.GetAttribute(network.ROLEID)
	userId := session.GetAttribute(network.USERID)
	if roleId != nil{
		id := roleId.(int64)
		role := handler.gateServer.RoleManager.GetOnlineRole(id)
		if role != nil{
			handler.gateServer.RoleManager.QuitRoleNoClearCache(id,-100)
			quit = true
			handler.gateServer.RemoveRoleSession(id)
		}
	}
	if userId != nil{
		id := userId.(int64)
		serverId := session.GetAttribute(network.SERVERID).(int32)
		handler.gateServer.RemoveUserSession(id)
		if quit == false{
			handler.gateServer.RoleManager.QuitRoleNoClearCache(-1000,serverId)
		}
	}
	if server, ok := handler.server.(*network.TcpServer); ok {
		delete(server.Sessions, session.Id())
		log4g.Infof("Session[%d]关闭!角色[%d]退出网关!", session.Id(),roleId.(int64))
	}
}

func (handler ServerMessageHandler) SessionPeriod(session network.SocketSessionInterface) {

}

func (handler ServerMessageHandler) ExceptionCaught(session network.SocketSessionInterface, err error) {
	//log4g.Info(err.Error())
}
