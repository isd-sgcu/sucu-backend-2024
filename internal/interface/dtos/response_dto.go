package dtos

type PaginationResponse struct {
	Data      interface{} `json:"data,omitempty"`
	Page      string      `json:"page"`  // current page
	Limit     string      `json:"limit"` // page size
	TotalPage string      `json:"total_pages"`
}
