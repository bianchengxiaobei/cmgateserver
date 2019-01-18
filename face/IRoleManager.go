package face

type IRoleManager interface {
	GetOnlineRole(roleId int64) IOnlineRole
	AddOnlineRole(role IOnlineRole)
	NewOnlineRoleFormDB(roleId int64) IOnlineRole
	NewOnlineRole(roleId int64) IOnlineRole
	RegisterRole(serverId int32,userId int64,roleId int64)
	QuitRoleNoClearCache(roleId int64,serverId int32)
	QuitRoleWithClearCache(roleId int64)
}