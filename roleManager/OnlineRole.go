package roleManager

import (
	"cmgateserver/bean"
)

type OnlineRole struct {
	Role     bean.Role
	GateId   int32
	UserName string
}
func (role *OnlineRole)GetRole() *bean.Role{
	return &role.Role
}
func (role *OnlineRole)GetRoleId() int64{
	return role.Role.RoleId
}
func (role *OnlineRole)GetServerId() int32{
	return role.Role.ServerId
}
func (role *OnlineRole)GetUserId() int64{
	return role.Role.UserId
}
func (role *OnlineRole)GetUseName() string{
	return role.UserName
}
func (role *OnlineRole)GetGateId() int32{
	return role.GateId
}
func (role *OnlineRole)SetRoleId(roleId int64) {
	role.Role.RoleId = roleId
}
func (role *OnlineRole)SetServerId(serverId int32) {
	role.Role.ServerId = serverId
}
func (role *OnlineRole)SetUserId(userId int64){
	role.Role.UserId = userId
}
func (role *OnlineRole)SetUseName(name string) {
	role.UserName = name
}
func (role *OnlineRole)SetGateId(gateId int32) {
	role.GateId = gateId
}
func (role *OnlineRole)GetAvatarId() int32{
	return role.Role.AvatarId
}
func (role *OnlineRole)SetAvatarId(avatarId int32){
	role.Role.AvatarId = avatarId
}
func (role *OnlineRole)GetNickName()string{
	return role.Role.NickName
}
func (role *OnlineRole)SetNickName(nickName string){
	role.Role.NickName = nickName
}
func (role *OnlineRole)GetGold()int32{
	return role.Role.Gold
}
func (role *OnlineRole)SetGold(gold int32){
	role.Role.Gold = gold
}
func (role *OnlineRole)GetExp()int32{
	return role.Role.Exp
}
func (role *OnlineRole)SetExp(exp int32){
	role.Role.Exp = exp
}
func (role *OnlineRole)GetLevel()int32{
	return role.Role.Level
}
func (role *OnlineRole)SetLevel(level int32){
	role.Role.Level = level
}
//func (role *OnlineRole)AddEmail(email bean.Email){
//	for k,v:= range role.Role.Emails{
//		if v == 0{
//			v.EmailId = email.EmailId
//			v.Get = email.Get
//			v.EmailTime = email.EmailTime
//			v.EmailIndex = int32(k)
//			role.Role.Emails[k] = v
//			break
//		}
//	}
//}