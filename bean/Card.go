package bean

import "time"

type CardType int32

const(
	NoneCard  CardType = 0
	YueCard CardType = 1
	NianCard CardType  = 2
	JiCard CardType = 3
	ZhongShenCard CardType = 4
)

type Card struct {
	CardType CardType
	BuyTime time.Time
}
