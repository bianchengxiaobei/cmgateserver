package bean
//订单

type TranType int32

const (
	DiamTranType TranType = 0
	CardTranType TranType = 1
)

type Tran struct {
	TranType TranType
	TranId   string
	TranValue int32
}
