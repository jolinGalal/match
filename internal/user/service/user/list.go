package user

import (
	"context"
	"fmt"

	"github.com/jolinGalal/match/internal/models"
	users "github.com/jolinGalal/match/internal/user/gen/users"
	"github.com/jolinGalal/match/internal/user/types"
	"github.com/jolinGalal/match/pkg/utils"
)

// list users
func (s *userssrvc) List(ctx context.Context, p *users.ListPayload) (res *users.ListResult, err error) {
	res = &users.ListResult{}
	s.db.Rest()
	var userObj = &models.User{}
	err = s.db.Model(userObj).
		Sort(fmt.Sprintf("%s %s", types.UserSortFields[p.SortKey], p.SortDirection)).
		Paging(&p.PageNumber, &p.PageSize).
		Find(&res.Users)
	if err != nil {
		return nil, err
	}
	s.db.Rest()
	totalCounts, err := s.db.Model(userObj).Count()
	paging := utils.GetPagination(&p.PageNumber, &p.PageSize, totalCounts)
	res.Pagination = &users.Pagination{
		CurrentPage: paging.GetCurrentPage(),
		PageSize:    paging.GetPageSize(),
		TotalPages:  paging.GetTotalPages(),
		TotalCount:  paging.GetCount(),
	}
	return res, err
}
