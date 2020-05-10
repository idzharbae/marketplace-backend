package request

type Transfer struct {
	SenderID       int64
	ReceiverID     int64
	TransferAmount int64
}

type TopUp struct {
	UserID   int64
	UserType int32
	Amount   int64
}
