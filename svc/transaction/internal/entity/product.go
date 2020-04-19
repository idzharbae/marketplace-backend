package entity

type Product struct {
	ID         int64
	ShopID     int64
	Name       string
	AmountKG   float64
	PricePerKG int32
	TotalPrice int64
}
