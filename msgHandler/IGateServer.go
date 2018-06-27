package msgHandler

import (
	"github.com/bianchengxiaobei/cmgo/network"
	"cmgateserver/db"
)

type IGateServer interface {
	RegisterInnerGameServer(serverId int32,session network.SocketSessionInterface)
	RemoveInnerGameServer(serverId int32,session network.SocketSessionInterface)
	GetId()  int
	GetDBManager() *db.MongoBDManager
	RegisterUser(serverId int32,userId int64,session network.SocketSessionInterface)
	RemoveUserSession(userId int64)
}
