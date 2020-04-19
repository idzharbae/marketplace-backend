package model

type Cart struct {
	ID        int64
	ProductID int64
	UserID    int64
	AmountKG  float64
}

func (c Cart) TableName() string {
	return "cart"
}
