package bean

import "time"

type Role struct {
	RoleId   int64
	UserId   int64
	NickName string
	ServerId int32
	Level	int32
	AvatarId	int32
	Gold	int32
	Diam    int32
	Exp     int32
	Sex 	int32
	Sign    string
	MaxBagNum  int32
	Items	[]Item
	WinLevel []int32
	DayGetTask	[]int32
	Achievement []int32
	LoginTime time.Time
	GetSign			bool
	TaskSeed  int32
	FreeSoldierData   [4]FreeSoldierData
}
type FreeSoldierData struct {
	PlayerType		int32
	CarrierType     int32
	TouKuiId		int32
	BodyId			int32
	WeapId			int32
}
