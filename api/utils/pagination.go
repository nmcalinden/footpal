package utils

import (
	"fmt"
	"github.com/nmcalinden/footpal/api/views"
)

func BuildPagination(total int, limit int, size int, initId, lastId int, query string) views.Pagination {

	return views.Pagination{
		Total:  total,
		Limit:  limit,
		Size:   size,
		LastId: lastId,
		Links: views.PaginationLinks{
			Previous: getPrevUrl(query, limit, initId),
			Next:     getNextUrl(query, size, limit, lastId),
		},
	}
}

func getPrevUrl(query string, limit int, initId int) string {
	return fmt.Sprintf(query, limit, initId)

}
func getNextUrl(query string, size int, limit int, lastId int) *string {
	if size == limit {
		next := fmt.Sprintf(query, limit, lastId)
		return &next
	}

	return nil
}
