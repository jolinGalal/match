package utils

import "math"

// Pagination is
type Pagination struct {
	currentPage int
	pageSize    int
	count       int64
	totalPages  int
}

// GetPagination ...
func GetPagination(currentPage, pageSize *int, count *int64) *Pagination {
	if currentPage != nil && pageSize != nil && count != nil {
		totalPages := int(math.Ceil(float64(*count) / float64(*pageSize)))
		return &Pagination{
			count:       *count,
			pageSize:    *pageSize,
			totalPages:  totalPages,
			currentPage: *currentPage,
		}
	}
	return &Pagination{
		count:       0,
		pageSize:    0,
		totalPages:  0,
		currentPage: 0,
	}
}

// GetCurrentPage ...
func (p *Pagination) GetCurrentPage() *int {
	return &p.currentPage
}

// GetPageSize ...
func (p *Pagination) GetPageSize() *int {
	return &p.pageSize
}

// GetCount ...
func (p *Pagination) GetCount() *int64 {
	return &p.count
}

// GetTotalPages ...
func (p *Pagination) GetTotalPages() *int {
	return &p.totalPages
}
