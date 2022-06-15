package utils

import (
	"fmt"

	"github.com/nmcalinden/footpal/views"
)

// BuildPagination
// total = Total amount of records
// limit = Limit of records for result
// Size = Total amount of records retrieved
// Init Offset = Requests initial search - Offset or ID
// Next Offset = Next Offset for next search - Offset or ID
// Query = Query path for Next and Prev Results
func BuildPagination(total, limit, size, initOffset, nextOffset int, query string) views.Pagination {
	pg := views.Pagination{
		Total:  total,
		Limit:  limit,
		Size:   size,
		Offset: initOffset,
		Links: &views.PaginationLinks{
			Previous: getPrevUrl(query, limit, initOffset),
		},
	}

	if size == limit && size+initOffset < total {
		pg.Links.Next = getNextUrl(query, limit, nextOffset)
	}
	return pg
}

func getPrevUrl(query string, limit int, initId int) string {
	return fmt.Sprintf(query, limit, initId)

}
func getNextUrl(query string, limit int, lastId int) *string {
	next := fmt.Sprintf(query, limit, lastId)
	return &next
}
