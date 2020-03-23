package converter

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
)

func ListShopProtoToReq(in *catalogproto.ListShopsReq) requests.ListShop {
	return requests.ListShop{Pagination: requests.Pagination{
		Page:  int(in.GetPagination().GetPage()),
		Limit: int(in.GetPagination().GetLimit()),
	}}
}
