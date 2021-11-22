package product

import (
	"context"
	"fmt"

	"github.com/jolinGalal/match/internal/models"
	products "github.com/jolinGalal/match/internal/product/gen/products"
	"github.com/jolinGalal/match/internal/product/types"
	"github.com/jolinGalal/match/pkg/utils"
)

// list products
func (s *productssrvc) List(ctx context.Context, p *products.ListPayload) (res *products.ListResult, err error) {
	res = &products.ListResult{}
	s.db.Rest()
	var productObj = &models.Product{}
	err = s.db.Model(productObj).
		Sort(fmt.Sprintf("%s %s", types.ProductSortFields[p.SortKey], p.SortDirection)).
		Paging(&p.PageNumber, &p.PageSize).
		Find(&res.Products)
	if err != nil {
		return nil, err
	}
	s.db.Rest()
	totalCounts, err := s.db.Model(productObj).Count()
	paging := utils.GetPagination(&p.PageNumber, &p.PageSize, totalCounts)
	res.Pagination = &products.Pagination{
		CurrentPage: paging.GetCurrentPage(),
		PageSize:    paging.GetPageSize(),
		TotalPages:  paging.GetTotalPages(),
		TotalCount:  paging.GetCount(),
	}
	return res, err
}
