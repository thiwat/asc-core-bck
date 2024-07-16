package types

type ListOutput[T any] struct {
	Rows      []T `json:"rows"`
	Page      int64  `json:"page"`
	PageSize  int64  `json:"page_size"`
	Total     int64  `json:"total"`
	TotalPage int64  `json:"total_page"`
}
