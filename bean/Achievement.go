package bean

type TaskConditionType int32

const (
	KillFarmer TaskConditionType = 0
	KillSoldier TaskConditionType = iota
	KillHero TaskConditionType = iota
	KillBuild TaskConditionType = iota
	KillStoreCarNum TaskConditionType  = iota
	KillLadderNum TaskConditionType = iota
	KillMangoneNum TaskConditionType  = iota
	KillDockNum TaskConditionType  = iota
	WinNum TaskConditionType  = iota
	FailedNum TaskConditionType  = iota
	DemageNum TaskConditionType  = iota
	BeDemageNum TaskConditionType  = iota
	BreadNum TaskConditionType  = iota
	CerealNum TaskConditionType  = iota
	TreeNum TaskConditionType  = iota
	MineNum TaskConditionType  = iota
	MeatNum TaskConditionType  = iota
	WineNum TaskConditionType = iota
	RCaocao TaskConditionType = iota
	RZhugeliang TaskConditionType = iota
	RLiubei TaskConditionType = iota
	RXuRong TaskConditionType = iota
	RHuangFuSong TaskConditionType = iota
	RZhangJiao TaskConditionType = iota
	HeroKillSoldier TaskConditionType  = iota
	PaiWeiWinNum TaskConditionType = iota
	RoomWinNum TaskConditionType = iota
	AllGameNum TaskConditionType  = iota
	HighestRankLevel TaskConditionType  = iota
	ConditionTypeEnd TaskConditionType  = iota
)




type Achievement struct{
	//KillSoldierNum int32
	//KillFarmerNum int32
	//KillBuildNum int32
	//KillHeroNum int32
	//WinNum int32
	//PaiWeiWinNum int32
	//PaiWeiFailedNum int32
	//RoomWinNum int32
	//RoomFailedNum int32
	//SimulateWinNum int32
	//SimulateFailedNum int32
	//BattleGameWinNum int32
	//BattleGameFailedNum int32
	//FailedNum int32
	//AllGameNum int32
	//DemageNum int32
	//BeDemageNum int32
	//BreadNum int32
	//CerealNum int32
	//TreeNum int32
	//MineNum int32
	//MeatNum int32
	//WineNum int32
	//KillStoreCarNum int32
	//KillLadderNum int32
	//KillMangoneNum int32
	//KillDockNum int32
	//HighestRankLevel int32
	AllGameTime int32
	 ConditionType []int32
	ConditionValue []int32
}
