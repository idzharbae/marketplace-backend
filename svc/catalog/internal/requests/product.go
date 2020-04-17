package requests

import "github.com/lib/pq"

type ListProduct struct {
	ShopIDs    pq.Int64Array
	Search     string
	OrderBy    string
	OrderType  string
	Category   string
	Pagination Pagination
}
