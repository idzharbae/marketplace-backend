package requests

import "errors"

type Pagination struct {
	Page  int
	Limit int
}

func (p Pagination) OffsetFromPagination() int {
	return (p.Page - 1) * p.Limit
}

func (p Pagination) Validate() error {
	if p.Limit < 0 {
		return errors.New("limit should not be negative")
	}
	if p.Page < 1 {
		return errors.New("page should be > 1")
	}
	return nil
}
