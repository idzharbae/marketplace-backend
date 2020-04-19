package entity

type Cart struct {
	ID       int64
	Product  Product
	UserID   int64
	AmountKG float64
}
