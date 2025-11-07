package dto

type PaginationRes struct {
	Rows        any  `json:"rows"`
	TotalRows   int  `json:"total_rows"`
	RowPerPage  int  `json:"row_per_page"`
	TotalPage   int  `json:"total_page"`
	HasNextPage bool `json:"has_next_page"`
}
