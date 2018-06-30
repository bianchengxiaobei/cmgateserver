package face

type IOnlineRole interface {
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
}