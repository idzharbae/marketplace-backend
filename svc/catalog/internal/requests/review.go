package requests

type ListReview struct {
	ShopID    int64
	ProductID int64
	Pagination
}

type GetTotalAndAverageReview struct {
	ShopID    int64
	ProductID int64
}

type TotalAndAverageReview struct {
	Total   int32
	Average float32
}
