package types

import "github.com/jolinGalal/match/internal/models"

var product = &models.Product{}

// ProductSort ...
var ProductSort = struct {
	ProductID   string
	ProductName string
	CreatedAt   string
}{
	ProductID:   "ProductID",
	ProductName: "ProductName",
	CreatedAt:   "CreatedAt",
}

// ProductSortFields map sorting field from the api to the database
var ProductSortFields = map[string]string{
	ProductSort.ProductID:   product.ID(),
	ProductSort.ProductName: product.Name(),
	ProductSort.CreatedAt:   product.CreatedAt(),
}

// SortDirection ...
var SortDirection = struct {
	Asc  string
	Desc string
}{
	Asc:  "asc",
	Desc: "desc",
}
