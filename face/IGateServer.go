package face

import (
	"github.com/bianchengxiaobei/cmgo/db"
	"github.com/bianchengxiaobei/cmgo/network"
	"github.com/golang/protobuf/proto"
)

type IGateServer interface {
	RegisterInnerGameServer(serverId int32, session network.SocketSessionInterface)
	RemoveInnerGameServer(serverId int32, session network.SocketSessionInterface)
	GetId() int
	GetDBManager() *db.MongoBDManager
	GetRoleManager() IRoleManager
	RegisterUserSession(serverId int32, userId int64, session network.SocketSessionInterface)
	RemoveUserSession(userId int64)
	GetUserSession(userId int64) (session network.SocketSessionInterface)
	RegisterRoleSession(roleId int64, session network.SocketSessionInterface)
	RemoveRoleSession(roleId int64)
	GetRoleSession(roleId int64) (session network.SocketSessionInterface)
	SendMsgToGameServer(serverId int32, msgId int, msg proto.Message) error
	SendMsgToGameServerByRoleId(roleId int64, msgId int, msg proto.Message) error
	SendMsgToClientByRoleId(roleId int64, msgId int, msg interface{}) error
	SendMsgToClient(session network.SocketSessionInterface, msgId int, msg interface{}) error
}