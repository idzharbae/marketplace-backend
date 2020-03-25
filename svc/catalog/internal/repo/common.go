package repo

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/repo/connection"
	"github.com/idzharbae/marketplace-backend/svc/catalog/internal/requests"
)

func applyPagination(pagination requests.Pagination, db connection.Gormw) connection.Gormw {
	if pagination.Limit > 0 {
		db = db.Limit(pagination.Limit)
	}
	if pagination.Page > 1 {
		db = db.Offset(pagination.OffsetFromPagination())
	}
	return db
}
