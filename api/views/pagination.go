package views

type Pagination struct {
	Total  int              `json:"total"`
	Limit  int              `json:"limit"`
	Size   int              `json:"size"`
	Offset int              `json:"offset"`
	Links  *PaginationLinks `json:"_links,omitempty"`
}

type PaginationLinks struct {
	Previous string  `json:"previous"`
	Next     *string `json:"next,omitempty"`
}
