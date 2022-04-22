package views

type Pagination struct {
	Total  int             `json:"total"`
	Limit  int             `json:"limit"`
	Size   int             `json:"size"`
	LastId int             `json:"lastId"`
	Links  PaginationLinks `json:"_links"`
}

type PaginationLinks struct {
	Previous string  `json:"previous"`
	Next     *string `json:"next,omitempty"`
}
