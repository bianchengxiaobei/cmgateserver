package roleManager

import (
	"cmgateserver/bean"
	"gopkg.in/mgo.v2/bson"
	"sync"
	"github.com/bianchengxiaobei/cmgo/network"
	"cmgateserver/face"
	"cmgateserver/message"
)

type RoleManager struct {
	lock        sync.RWMutex
	onlineRoles map[int64]face.IOnlineRole
	GateServer   face.IGateServer
}
func NewRoleManager(server face.IGateServer) *RoleManager {
	return &RoleManager{
		onlineRoles: make(map[int64]face.IOnlineRole),
		GateServer:   server,
	}
}
func (manager *RoleManager) GetOnlineRole(roleId int64) face.IOnlineRole {
	manager.lock.RLock()
	defer manager.lock.RUnlock()
	return manager.onlineRoles[roleId]
}
func (manager *RoleManager) AddOnlineRole(role face.IOnlineRole) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.onlineRoles[role.GetRoleId()] = role
}
func (manager *RoleManager)RemoveOnlineRole(roleId int64)face.IOnlineRole{
	manager.lock.Lock()
	defer manager.lock.Unlock()
	role := manager.onlineRoles[roleId]
	if role != nil{
		delete(manager.onlineRoles,roleId)
	}
	return role
}
func (manager *RoleManager) NewOnlineRoleFormDB(roleId int64) face.IOnlineRole {
	var err error
	if manager.GateServer.GetDBManager() != nil {
		dbSession := manager.GateServer.GetDBManager().Get()
		if dbSession != nil {
			role := bean.Role{}
			c := dbSession.DB("sanguozhizhan").C("Role")
			err = c.Find(bson.M{"roleid": roleId}).One(&role)
			if err != nil {
				return nil
			}
			onlineRole := OnlineRole{
				Role: role,
			}
			manager.onlineRoles[roleId] = &onlineRole
			return &onlineRole
		}
	}
	return nil
}
func (manager *RoleManager)NewOnlineRole(roleId int64) face.IOnlineRole{
	onlineRole := &OnlineRole{
		Role:bean.Role{},
	}
	onlineRole.Role.RoleId = roleId
	manager.onlineRoles[roleId] = onlineRole
	return onlineRole
}
func (manager *RoleManager)RegisterRole(serverId int32,userId int64,roleId int64){
	role := manager.GetOnlineRole(roleId)
	if role == nil{
		role = manager.NewOnlineRoleFormDB(roleId)
		role.SetServerId(serverId)
		role.SetUserId(userId)
	}
	//如果玩家已经退出游戏了，就发送给游戏服务器断开连接
	session := manager.GateServer.GetUserSession(userId)
	if session == nil || session.IsClosed(){
		//发送给游戏逻辑服退出玩家
		return
	}
	aRoleId := session.GetAttribute(network.ROLEID)
	if aRoleId == nil{
		manager.GateServer.RegisterRoleSession(roleId,session)
	}
}
func (manager *RoleManager)QuitRoleNoClearCache(roleId int64,serverId int32){
	role := manager.GetOnlineRole(roleId)
	if role != nil{
		serverId := role.GetServerId()
		msg := new(message.G2M_RoleQuitGameServer)
		msg.RoleId = roleId
		manager.GateServer.SendMsgToGameServer(serverId, 10004, msg)
	}else{
		msg := new(message.G2M_RoleQuitGameServer)
		msg.RoleId = roleId
		manager.GateServer.SendMsgToGameServer(serverId, 10004, msg)
	}
}
func (manager *RoleManager)QuitRoleWithClearCache(roleId int64){
	role := manager.RemoveOnlineRole(roleId)
	if role != nil{
		session := manager.GateServer.GetUserSession(role.GetUserId())
		if session != nil{
			session.RemoveAttribute(network.ROLEID)
			session.RemoveAttribute(network.USERID)
			manager.GateServer.RemoveRoleSession(roleId)
			manager.GateServer.RemoveUserSession(role.GetUserId())
		}
	}
}

