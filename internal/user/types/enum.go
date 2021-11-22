package types

import "github.com/jolinGalal/match/internal/models"

var user = &models.User{}

// UserSort ...
var UserSort = struct {
	UserID    string
	UserName  string
	CreatedAt string
}{
	UserID:    "UserID",
	UserName:  "UserName",
	CreatedAt: "CreatedAt",
}

// UserSortFields map sorting field from the api to the database
var UserSortFields = map[string]string{
	UserSort.UserID:    user.ID(),
	UserSort.UserName:  user.Name(),
	UserSort.CreatedAt: user.CreatedAt(),
}

// SortDirection ...
var SortDirection = struct {
	Asc  string
	Desc string
}{
	Asc:  "asc",
	Desc: "desc",
}
