package requests

type ListProduct struct {
	ShopID     int64
	Province   string
	Search     string
	OrderBy    string
	OrderType  string
	Category   string
	Pagination Pagination
}
