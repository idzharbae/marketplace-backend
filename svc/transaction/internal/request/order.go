package request

type ListOrderReq struct {
	UserID int64
}

type CheckoutReq struct {
	CartIDs       []int64
	PaymentAmount int64
}
