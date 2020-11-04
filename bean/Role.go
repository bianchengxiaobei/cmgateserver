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
	MailAddress string
	QQ  int32
	WeiXin string
	Phone int32
	GuideFinished  bool //是否新手关卡完成了
	RankScore int32
	HeroCount int32
	MaxBagNum  int32
	MaxEmailNum int
	Card   []Card
	Trans  []Tran
	Items	[]Item
	Emails  []Email
	AchieveRecord Achievement
	WinLevel []int32
	DayGetTask	[]int32
	Achievement []int32
	LoginTime time.Time
	GetSign			bool
	TaskSeed  int32
	FreeSoldierData   [4]FreeSoldierData
	IsGM		bool //是否是游戏管理员
}
type FreeSoldierData struct {
	PlayerType		int32
	CarrierType     int32
	TouKuiId		int32
	BodyId			int32
	WeapId			int32
}
