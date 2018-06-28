package msgHandler

import (
	"github.com/bianchengxiaobei/cmgo/network"
	"cmgateserver/db"
	"github.com/golang/protobuf/proto"
)

type IGateServer interface {
	RegisterInnerGameServer(serverId int32,session network.SocketSessionInterface)
	RemoveInnerGameServer(serverId int32,session network.SocketSessionInterface)
	GetId()  int
	GetDBManager() *db.MongoBDManager
	RegisterUserSession(serverId int32,userId int64,session network.SocketSessionInterface)
	RemoveUserSession(userId int64)
	GetUserSession(userId int64)(session network.SocketSessionInterface)
	RegisterRoleSession(roleId int64,session network.SocketSessionInterface)
	RemoveRoleSession(roleId int64)
	GetRoleSession(roleId int64) (session network.SocketSessionInterface)
	SendMsgToGameServer(serverId int32,msgId int,msg proto.Message) error
}
