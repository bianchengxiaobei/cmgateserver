package bean

type Email struct{
	//EmailId int32
	EmailTime int64
	EmailIndex int32
	Title string
	Content string
	AwardList []int32
	Get   bool
	Valid bool
}
