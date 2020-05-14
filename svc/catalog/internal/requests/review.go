package requests

type ListReview struct {
	ShopID    int64
	ProductID int64
	Pagination
}
