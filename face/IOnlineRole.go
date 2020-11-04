package face

import "cmgateserver/bean"

type IOnlineRole interface {
	GetRole() *bean.Role
	GetRoleId() int64
	GetServerId() int32
	GetUserId() int64
	GetUseName() string
	GetGateId() int32
	SetRoleId(int64)
	SetServerId(int32)
	SetUserId(int64)
	SetUseName(string)
	SetGateId(int32)
	GetAvatarId() int32
	SetAvatarId(avatarId int32)
	GetNickName()string
	SetNickName(nickName string)
	GetGold() int32
	SetGold(gold int32)
	GetExp() int32
	SetExp(exp int32)
	GetLevel() int32
	SetLevel(level int32)
	//AddEmail(email bean.Email)
}